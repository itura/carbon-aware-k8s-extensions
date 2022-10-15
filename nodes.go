package main

import (
	"k8s.io/api/core/v1"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Mapping[T any] map[string]T

const labelK8sRegion = "topology.kubernetes.io/region"

type Policy struct {
}

func NewPolicy() *Policy {
	return &Policy{}
}

type Nodes struct {
	nodes []v1.Node
	m     Mapping[[]v1.Node]
}

func NewNodes(nodes []v1.Node) *Nodes {
	m := Mapping[[]v1.Node]{}
	for _, node := range nodes {
		region := node.Labels[labelK8sRegion]
		_, exists := m[region]
		if !exists {
			m[region] = []v1.Node{node}
		} else {
			m[region] = append(m[region], node)
		}
	}

	return &Nodes{
		nodes: nodes,
		m:     m,
	}
}

func (n *Nodes) ForRegion(region string) []v1.Node {
	nodes, present := n.m[region]
	if !present {
		return []v1.Node{}
	}

	return nodes
}

func (n *Nodes) GetAll() []v1.Node {
	return n.nodes
}

func (n *Nodes) Size() int {
	return len(n.nodes)
}

func (n *Nodes) GetRegions() []string {
	var result []string
	for k, _ := range n.m {
		result = append(result, k)
	}
	return result
}

func (n *Nodes) UpdateMetadata(policy *Policy) {

}

type NodeBuilder struct {
	node v1.Node
}

func NewNodeBuilder(name string) *NodeBuilder {
	return &NodeBuilder{
		node: v1.Node{
			ObjectMeta: v12.ObjectMeta{
				Name:        name,
				Labels:      map[string]string{},
				Annotations: map[string]string{},
			},
			Spec: v1.NodeSpec{
				Taints: []v1.Taint{},
			},
		}}
}

func UpdateNode(node v1.Node) *NodeBuilder {
	return &NodeBuilder{node: node}
}

func (b *NodeBuilder) Build() v1.Node {
	return b.node
}

func (b *NodeBuilder) SetName(n string) *NodeBuilder {
	b.node.Name = n
	return b
}

func (b *NodeBuilder) SetLabel(k, v string) *NodeBuilder {
	b.node.Labels[k] = v
	return b
}

func (b *NodeBuilder) SetTaint(t v1.Taint) *NodeBuilder {
	b.node.Spec.Taints = append(b.node.Spec.Taints, t)
	return b
}

func (b *NodeBuilder) RemoveTaint(key string) *NodeBuilder {
	var updated []v1.Taint
	for _, taint := range b.node.Spec.Taints {
		if taint.Key != key {
			updated = append(updated, taint)
		}
	}
	b.node.Spec.Taints = updated
	return b
}

func (b *NodeBuilder) SetAnnotation(k, v string) *NodeBuilder {
	b.node.Annotations[k] = v
	return b
}

func (b *NodeBuilder) SetRegion(r string) *NodeBuilder {
	return b.SetLabel(labelK8sRegion, r)
}
