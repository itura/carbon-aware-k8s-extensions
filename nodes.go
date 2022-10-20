package main

import (
	"k8s.io/api/core/v1"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

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

func (n *Nodes) Len() int {
	return len(n.nodes)
}

func (n *Nodes) GetRegions() []string {
	_result := Mapping[string]{}
	for _, node := range n.nodes {
		_result[getLocation(node)] = ""
	}
	var result []string
	for k := range _result {
		result = append(result, k)
	}
	return result
}

func (n *Nodes) Update(updates Mapping[v1.Node]) {
	for i, node := range n.nodes {
		location := getLocation(node)
		update, exists := updates[location]
		if exists {
			updater := UpdateNode(node)
			if len(update.Spec.Taints) > 0 {
				updater.AddAllTaints(update.Spec.Taints)
			}
			if len(update.Labels) > 0 {
				updater.SetAllLabels(update.Labels)
			}
			if len(update.Annotations) > 0 {
				updater.SetAllAnnotations(update.Annotations)
			}
			n.nodes[i] = updater.Build()
		}
	}
}

func (n *Nodes) Unset(label string) {
	for i, node := range n.nodes {
		n.nodes[i] = UpdateNode(node).
			RemoveLabel(label).
			RemoveTaint(label).
			RemoveAnnotation(label).
			Build()
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

func (b *NodeBuilder) RemoveLabel(k string) *NodeBuilder {
	delete(b.node.Labels, k)
	return b
}

func (b *NodeBuilder) SetAllLabels(labels Mapping[string]) *NodeBuilder {
	var m Mapping[string]
	m = b.node.Labels
	b.node.Labels = m.Merge(labels)
	return b
}

func (b *NodeBuilder) AddTaint(t v1.Taint) *NodeBuilder {
	for _, taint := range b.node.Spec.Taints {
		if taint.Key == t.Key && taint.Effect == t.Effect {
			return b
		}
	}
	b.node.Spec.Taints = append(b.node.Spec.Taints, t)
	return b
}

func (b *NodeBuilder) AddAllTaints(ts []v1.Taint) *NodeBuilder {
	for _, taint := range ts {
		b.AddTaint(taint)
	}
	return b
}

func (b *NodeBuilder) RemoveTaint(k string) *NodeBuilder {
	updated := []v1.Taint{}
	for _, taint := range b.node.Spec.Taints {
		if taint.Key != k {
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

func (b *NodeBuilder) RemoveAnnotation(k string) *NodeBuilder {
	delete(b.node.Annotations, k)
	return b
}

func (b *NodeBuilder) SetAllAnnotations(annotations Mapping[string]) *NodeBuilder {
	var m Mapping[string]
	m = b.node.Annotations
	b.node.Annotations = m.Merge(annotations)
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
