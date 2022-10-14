package main

import (
	"github.com/stretchr/testify/suite"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"testing"
)

func TestNodes(t *testing.T) {
	suite.Run(t, new(NodesSuite))
}

type NodesSuite struct {
	suite.Suite
}

func (s *NodesSuite) TestNodesAreGroupedBasedOnLocationLabels() {
	n1 := stubNodeWithLabels("n1", Mapping[string]{
		labelK8sRegion: "us-east1",
	})
	n2 := stubNodeWithLabels("n2", Mapping[string]{
		labelK8sRegion: "us-central1",
	})
	n3 := stubNodeWithLabels("n3", Mapping[string]{
		labelK8sRegion: "us-central1",
	})

	nodes := NewNodes([]v1.Node{n1, n2, n3})

	s.Equal([]v1.Node{n1}, nodes.ForRegion("us-east1"))
	s.Equal([]v1.Node{n2, n3}, nodes.ForRegion("us-central1"))
}

func stubNode(name string) *v1.Node {
	return &v1.Node{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
	}
}

func stubNodeWithLabels(name string, labels Mapping[string]) v1.Node {
	return v1.Node{
		ObjectMeta: metav1.ObjectMeta{
			Name:   name,
			Labels: labels,
		},
	}
}
