package main

import (
	"fmt"
	v1 "k8s.io/api/core/v1"
)

type CarbonPolicySpec struct {
	DataSource string      `yaml:"dataSource"`
	SortBy     string      `yaml:"sortBy"`
	Taints     TaintPolicy `yaml:"taints"`
}

type TaintPolicy struct {
	Type   string         `yaml:"type"`
	Effect v1.TaintEffect `yaml:"type"`
}

type CarbonPolicy struct {
	locations *Locations
	nodes     *Nodes
	Spec      CarbonPolicySpec
}

func NewCarbonPolicy(spec CarbonPolicySpec) *CarbonPolicy {
	return &CarbonPolicy{
		Spec: setDefaults(spec),
	}
}

func setDefaults(spec CarbonPolicySpec) CarbonPolicySpec {
	if spec.DataSource == "" {
		spec.DataSource = dataSourceCAAPI
	}
	if spec.SortBy == "" {
		spec.SortBy = policySortByCurrentIntensity
	}
	if spec.Taints.Type == "" {
		spec.Taints.Type = policyTaintTypeWorst
	}
	if spec.Taints.Effect == "" {
		spec.Taints.Effect = v1.TaintEffectNoSchedule
	}
	return spec
}

func (p *CarbonPolicy) SetLocations(locations *Locations) *CarbonPolicy {
	p.locations = locations
	return p
}

func (p *CarbonPolicy) SetNodes(nodes *Nodes) *CarbonPolicy {
	p.nodes = nodes
	return p
}

func (p *CarbonPolicy) UpdateNodes() (*Nodes, error) {
	if p.locations == nil {
		return nil, fmt.Errorf("location data missing from policy")
	}
	if p.nodes == nil {
		return nil, fmt.Errorf("node data missing from policy")
	}

	locations := p.nodes.GetRegions()
	if len(locations) < 2 && p.Spec.Taints.Type != policyTaintTypeTest {
		return p.nodes, nil
	}

	if p.Spec.SortBy == policySortByCurrentIntensity {
		p.locations.SortByIntensity()
	} else if p.Spec.SortBy == policySortByRating {
		p.locations.SortByRating()
	} else {
		return nil, fmt.Errorf("invalid value for .SortBy")
	}

	var updates Mapping[v1.Node]
	if p.Spec.Taints.Type == policyTaintTypeWorst {
		updates = p.taintWorst(locations)
	} else if p.Spec.Taints.Type == policyTaintTypeTest {
		updates = p.taintWorst(locations)
	} else {
		return nil, fmt.Errorf("invalid value for .Taints.Type")
	}

	p.nodes.Update(updates)

	return p.nodes, nil
}

func (p *CarbonPolicy) taintWorst(locationNames []string) Mapping[v1.Node] {
	dirtiestLocation := p.locations.GetLast().Name
	result := Mapping[v1.Node]{}
	for _, location := range locationNames {
		if location == dirtiestLocation {
			update := NewNodeBuilder("").
				AddTaint(taintHighIntensity(p.Spec.Taints.Effect)).
				Build()
			result[location] = update
		} else {
			result[location] = NewNodeBuilder("").Build()
		}
	}
	return result
}
