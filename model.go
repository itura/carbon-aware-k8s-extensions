package main

import "k8s.io/api/core/v1"

type Mapping[T any] map[string]T

type Nodes struct {
	nodes []v1.Node
	m     Mapping[[]*v1.Node]
}

func NewNodes(nodes []v1.Node) *Nodes {
	m := Mapping[[]*v1.Node]{}
	for _, node := range nodes {
		region := node.Labels["topology.kubernetes.io/region"]
		_, exists := m[region]
		if !exists {
			m[region] = []*v1.Node{&node}
		} else {
			m[region] = append(m[region], &node)
		}
	}

	return &Nodes{
		nodes: nodes,
		m:     m,
	}
}

func (n *Nodes) ForRegion(region string) []*v1.Node {
	nodes, present := n.m[region]
	if !present {
		return []*v1.Node{}
	}

	return nodes
}

func (n *Nodes) GetAll() []v1.Node {
	return n.nodes
}

func (n *Nodes) Size() int {
	return len(n.nodes)
}
