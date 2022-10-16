package main

import (
	"github.com/stretchr/testify/suite"
	v1 "k8s.io/api/core/v1"
	"testing"
)

func TestNodes(t *testing.T) {
	suite.Run(t, new(NodesSuite))
}

type NodesSuite struct {
	suite.Suite
}

func (s *NodesSuite) TestNodesAreGroupedBasedOnLocationLabels() {
	n1 := NewNodeBuilder("n1").SetRegion("us-east1").Build()
	n2 := NewNodeBuilder("n2").SetRegion("us-central1").Build()
	n3 := NewNodeBuilder("n3").SetRegion("us-central1").Build()

	nodes := NewNodes([]v1.Node{n1, n2, n3})

	s.ElementsMatch([]v1.Node{n1}, nodes.ForLocation("us-east1"))
	s.ElementsMatch([]v1.Node{n2, n3}, nodes.ForLocation("us-central1"))
}

func (s *NodesSuite) TestGetRegions() {
	n1 := NewNodeBuilder("n1").SetRegion("us-east1").Build()
	n2 := NewNodeBuilder("n2").SetRegion("us-central1").Build()
	n3 := NewNodeBuilder("n3").SetRegion("us-central1").Build()

	nodes := NewNodes([]v1.Node{n1, n2, n3})

	s.ElementsMatch([]string{"us-east1", "us-central1"}, nodes.GetRegions())
}

func (s *NodesSuite) TestGetAll() {
	n1 := NewNodeBuilder("n1").SetRegion("us-east1").Build()
	n2 := NewNodeBuilder("n2").SetRegion("us-central1").Build()
	n3 := NewNodeBuilder("n3").SetRegion("us-central1").Build()

	nodes := NewNodes([]v1.Node{n1, n2, n3})

	var results []v1.Node
	for node := range nodes.Iterator() {
		results = append(results, node)
	}
	s.ElementsMatch([]v1.Node{n1, n2, n3}, results)
}

func (s *NodesSuite) TestUpdateMetadataAppliesATaintToNodesInTheLeastGreenLocation() {
	n1 := NewNodeBuilder("n1").SetRegion("us-east1").Build()
	n2 := NewNodeBuilder("n2").SetRegion("us-east1").Build()
	n3 := NewNodeBuilder("n3").SetRegion("us-central1").Build()
	policy := NewCarbonPolicy(CarbonPolicySpec{
		SortBy:     policySortByCurrentIntensity,
		DataSource: dataSourceCAAPI,
	}).
		SetLocations(NewLocations([]Location{{
			Name:      "us-east1",
			Intensity: 2.0,
		}, {
			Name:      "us-central1",
			Intensity: 1.0,
		}}))

	nodes := NewNodes([]v1.Node{n1, n2, n3})
	nodes.UpdateMetadata(policy)

	taints := []v1.Taint{taintHighIntensity(v1.TaintEffectNoSchedule)}
	emptyTaints := []v1.Taint{}
	s.Equal(taints, nodes.Get(0).Spec.Taints)
	s.Equal(taints, nodes.Get(1).Spec.Taints)
	s.Equal(emptyTaints, nodes.Get(2).Spec.Taints)
}

// changing this behavior would entail deep copying reference fields
func (s *NodesSuite) TestUpdateNodeMutatesTheReferencesOfOriginalNode() {
	n1 := NewNodeBuilder("n1").
		SetRegion("us-east1").
		Build()

	s.Equal("us-east1", n1.Labels[labelK8sRegion])
	_, exists := n1.Annotations["blah"]
	s.Equal(false, exists)
	s.Equal("n1", n1.Name)

	n2 := UpdateNode(n1).
		SetRegion("us-central1").
		SetAnnotation("blah", "blah").
		SetName("n2").
		Build()

	s.Equal("us-central1", n1.Labels[labelK8sRegion]) //reference updated
	_, exists = n1.Annotations["blah"]
	s.Equal(true, exists)  //reference updated
	s.Equal("n1", n1.Name) // value preserved

	s.Equal("us-central1", n2.Labels[labelK8sRegion])
	_, exists = n2.Annotations["blah"]
	s.Equal(true, exists)
	s.Equal("n2", n2.Name)
}
