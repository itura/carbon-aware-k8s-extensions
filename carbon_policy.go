package main

import (
	v1 "k8s.io/api/core/v1"
)

type CarbonPolicySpec struct {
	DataSource string
	SortBy     string
}

type CarbonPolicy struct {
	locations *Locations
	Spec      CarbonPolicySpec
}

func NewPolicy(spec CarbonPolicySpec) *CarbonPolicy {
	return &CarbonPolicy{
		Spec: spec,
	}
}

func (p *CarbonPolicy) SetLocations(locations *Locations) *CarbonPolicy {
	p.locations = locations
	return p
}

func (p *CarbonPolicy) ClassifyLocations(locationNames []string) Mapping[v1.Node] {
	if p.Spec.SortBy == policySortByIntensity {
		p.locations.SortByIntensity()
	} else {
		p.locations.SortByRating()
	}
	dirtiestLocation := p.locations.GetLast().Name

	result := Mapping[v1.Node]{}
	for _, location := range locationNames {
		if location == dirtiestLocation {
			meta := NewNodeBuilder("").
				AddTaint(taintHighIntensity).
				Build()
			result[location] = meta
		} else {
			result[location] = NewNodeBuilder("").Build()
		}
	}

	return result
}
