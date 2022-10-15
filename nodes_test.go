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

	s.Equal([]v1.Node{n1}, nodes.ForRegion("us-east1"))
	s.Equal([]v1.Node{n2, n3}, nodes.ForRegion("us-central1"))
}

func (s *NodesSuite) TestGetRegions() {
	n1 := NewNodeBuilder("n1").SetRegion("us-east1").Build()
	n2 := NewNodeBuilder("n2").SetRegion("us-central1").Build()
	n3 := NewNodeBuilder("n3").SetRegion("us-central1").Build()

	nodes := NewNodes([]v1.Node{n1, n2, n3})

	s.Equal([]string{"us-east1", "us-central1"}, nodes.GetRegions())
}

func (s *NodesSuite) TestUpdateMetadata() {
	n1 := NewNodeBuilder("n1").SetRegion("us-east1").Build()
	policy := NewPolicy()

	nodes := NewNodes([]v1.Node{n1})
	nodes.UpdateMetadata(policy)
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
