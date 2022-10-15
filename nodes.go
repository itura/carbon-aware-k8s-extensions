package main

import (
	"k8s.io/api/core/v1"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Mapping[T any] map[string]T

func (m Mapping[T]) Merge(other Mapping[T]) Mapping[T] {
	result := Mapping[T]{}
	for k, v := range m {
		result[k] = v
	}
	for k, v := range other {
		result[k] = v
	}
	return result
}

const labelK8sRegion = "topology.kubernetes.io/region"

type Nodes struct {
	nodes []v1.Node
}

func NewNodes(nodes []v1.Node) *Nodes {
	return &Nodes{
		nodes: nodes,
	}
}

func (n *Nodes) ForLocation(location string) []v1.Node {
	var results []v1.Node
	for _, node := range n.nodes {
		if getLocation(node) == location {
			results = append(results, node)
		}
	}
	return results
}

func (n *Nodes) Get(i int) v1.Node {
	return n.nodes[i]
}

func (n *Nodes) GetAll() []v1.Node {
	return n.nodes
}

func (n *Nodes) Iterator() <-chan v1.Node {
	ch := make(chan v1.Node, len(n.nodes))
	go func() {
		for _, node := range n.nodes {
			ch <- node
		}
		close(ch)
	}()
	return ch
}

func (n *Nodes) Size() int {
	return len(n.nodes)
}

func (n *Nodes) GetRegions() []string {
	_result := Mapping[string]{}
	for _, node := range n.nodes {
		_result[getLocation(node)] = ""
	}
	var result []string
	for k, _ := range _result {
		result = append(result, k)
	}
	return result
}

type NodeMetadata struct {
	Labels      Mapping[string]
	Annotations Mapping[string]
	Taints      []v1.Taint
}

func (n *Nodes) UpdateMetadata(policy *CarbonPolicy) {
	locationMetadata := policy.ClassifyLocations(n.GetRegions())
	for i, node := range n.nodes {
		location := getLocation(node)
		metadata, exists := locationMetadata[location]
		if exists {
			updater := UpdateNode(node)
			if len(metadata.Spec.Taints) > 0 {
				updater.AddAllTaints(metadata.Spec.Taints)
			}
			if len(metadata.Labels) > 0 {
				updater.SetAllLabels(metadata.Labels)
			}
			n.nodes[i] = updater.Build()
		}
	}
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

func (b *NodeBuilder) SetAllLabels(labels Mapping[string]) *NodeBuilder {
	var m Mapping[string]
	m = b.node.Labels
	b.node.Labels = m.Merge(labels)
	return b
}

func (b *NodeBuilder) AddTaint(t v1.Taint) *NodeBuilder {
	b.node.Spec.Taints = append(b.node.Spec.Taints, t)
	return b
}

func (b *NodeBuilder) AddAllTaints(ts []v1.Taint) *NodeBuilder {
	b.node.Spec.Taints = append(b.node.Spec.Taints, ts...)
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

func getLocation(n v1.Node) string {
	result, exists := n.Labels[labelK8sRegion]
	if exists {
		return result
	} else {
		return ""
	}
}
