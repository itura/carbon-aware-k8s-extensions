package main

import (
	v1 "k8s.io/api/core/v1"
	"sort"
)

const (
	labelIntensity = "greensoftware.foundation/carbon-intesity"
	intensityHigh  = "high"
	intensityLow   = "low"
)

type CarbonPolicy struct {
	currentIntensity Mapping[float64]
}

func NewPolicy() *CarbonPolicy {
	return &CarbonPolicy{
		currentIntensity: Mapping[float64]{},
	}
}

func (p *CarbonPolicy) SetIntensity(location string, intensity float64) *CarbonPolicy {
	p.currentIntensity[location] = intensity
	return p
}

func (p *CarbonPolicy) IsGreener(l1 string, l2 string) bool {
	return p.currentIntensity[l1] < p.currentIntensity[l2]
}

func (p *CarbonPolicy) ClassifyLocations(locations []string) Mapping[v1.Node] {
	sortedLocations := NewCurrentIntensitySorter(p).Sort()
	dirtestLocation := sortedLocations[len(sortedLocations)-1]

	result := Mapping[v1.Node]{}
	for _, location := range locations {
		if location == dirtestLocation {
			meta := NewNodeBuilder("meta").
				AddTaint(v1.Taint{
					Key:   labelIntensity,
					Value: intensityHigh,
				}).
				Build()
			result[location] = meta
		} else {
			result[location] = NewNodeBuilder("meta").Build()
		}
	}

	return result
}

type CurrentIntensitySorter struct {
	locations []string
	policy    *CarbonPolicy
}

func NewCurrentIntensitySorter(policy *CarbonPolicy) *CurrentIntensitySorter {
	var locations []string
	for location, _ := range policy.currentIntensity {
		locations = append(locations, location)
	}
	return &CurrentIntensitySorter{
		locations: locations,
		policy:    policy,
	}
}

func (s *CurrentIntensitySorter) Len() int {
	return len(s.locations)
}

func (s *CurrentIntensitySorter) Less(i, j int) bool {
	return s.policy.currentIntensity[s.locations[i]] < s.policy.currentIntensity[s.locations[j]]
}

func (s *CurrentIntensitySorter) Swap(i, j int) {
	s.locations[i], s.locations[j] = s.locations[j], s.locations[i]
}

func (s *CurrentIntensitySorter) Sort() []string {
	sort.Sort(s)
	return s.locations
}

func (p *CarbonPolicy) NodeSorter(nodes *Nodes) sort.Interface {
	return NewNodeSorter(nodes, p)
}

type NodeSorter struct {
	nodes  []v1.Node
	policy *CarbonPolicy
}

func NewNodeSorter(nodes *Nodes, policy *CarbonPolicy) *NodeSorter {
	return &NodeSorter{nodes.nodes, policy}
}

func (n *NodeSorter) Len() int {
	return len(n.nodes)
}

func (n *NodeSorter) Swap(i, j int) {
	n.nodes[i], n.nodes[j] = n.nodes[j], n.nodes[i]
}

func (n *NodeSorter) Less(i, j int) bool {
	return n.policy.IsGreener(getLocation(n.nodes[i]), getLocation(n.nodes[j]))
}
