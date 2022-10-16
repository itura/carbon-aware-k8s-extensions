package main

import (
	v1 "k8s.io/api/core/v1"
)

type CarbonPolicySpec struct {
	DataSource string
	SortBy     string
	Taints     TaintPolicy
}

type TaintPolicy struct {
	Type   string
	Effect v1.TaintEffect
}

type CarbonPolicy struct {
	locations *Locations
	Spec      CarbonPolicySpec
}

func NewCarbonPolicy(spec CarbonPolicySpec) *CarbonPolicy {
	if spec.DataSource == "" {
		spec.DataSource = dataSourceCAAPI
	}
	if spec.SortBy == "" {
		spec.SortBy = policySortByCurrentIntensity
	}
	if spec.Taints.Type == "" {
		spec.Taints.Type = policyTaintType

	}
	if spec.Taints.Effect == "" {
		spec.Taints.Effect = v1.TaintEffectNoSchedule
	}
	return &CarbonPolicy{
		Spec: spec,
	}
}

func (p *CarbonPolicy) SetLocations(locations *Locations) *CarbonPolicy {
	p.locations = locations
	return p
}

func (p *CarbonPolicy) NodeUpdatesByLocation(locationNames []string) Mapping[v1.Node] {
	if p.Spec.SortBy == policySortByCurrentIntensity {
		p.locations.SortByIntensity()
	} else if p.Spec.SortBy == policySortByRating {
		p.locations.SortByRating()
	} else {
		panic("invalid value for .SortBy")
	}

	var results Mapping[v1.Node]
	if p.Spec.Taints.Type == policyTaintType {
		results = p.taintWorst(locationNames)
	} else {
		panic("invalid value for .Taints.Type")
	}

	return results
}

func (p *CarbonPolicy) taintWorst(locationNames []string) Mapping[v1.Node] {
	dirtiestLocation := p.locations.GetLast().Name
	result := Mapping[v1.Node]{}
	for _, location := range locationNames {
		if location == dirtiestLocation {
			meta := NewNodeBuilder("").
				AddTaint(taintHighIntensity(p.Spec.Taints.Effect)).
				Build()
			result[location] = meta
		} else {
			result[location] = NewNodeBuilder("").Build()
		}
	}
	return result
}
