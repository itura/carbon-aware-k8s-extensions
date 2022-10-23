package main

import (
	"fmt"
	v1 "k8s.io/api/core/v1"
)

type CarbonPolicy struct {
	locations *Locations
	nodes     *Nodes
	Spec      CarbonPolicySpec
}

func NewCarbonPolicy(spec CarbonPolicySpec) *CarbonPolicy {
	return &CarbonPolicy{
		Spec: spec.SetDefaults(),
	}
}

func (p *CarbonPolicy) SetLocations(locations *Locations) *CarbonPolicy {
	var fn LocationMapper
	if p.Spec.DataSource.Locations.Preset == optGcp {
		fn = gcpRegions
	} else {
		fn = Id[Location]
	}
	p.locations = locations.Map(fn)
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

	if p.Spec.DataSource.SortBy == optIntensity {
		p.locations.SortByIntensity()
	} else if p.Spec.DataSource.SortBy == optRating {
		p.locations.SortByRating()
	} else {
		return nil, fmt.Errorf("invalid value for .SortBy")
	}

	p.nodes.Unset(labelIntensity)

	err := p.applyTaints()
	if err != nil {
		return nil, err
	}

	err = p.applyLabels()
	if err != nil {
		return nil, err
	}

	return p.nodes, nil
}

func (p *CarbonPolicy) applyTaints() error {
	locationNames := p.nodes.GetRegions()
	var taintUpdates Mapping[v1.Node]

	switch p.Spec.Taints.Type {
	case optNone:
		return nil
	case optWorst:
		if len(locationNames) < 2 && !p.Spec.Taints.ShouldTaintOnlyLocation {
			return nil
		}
		taintUpdates = p.applyTaintToWorstLocation(locationNames)
	default:
		return fmt.Errorf("invalid value for .Taints.Type")
	}

	p.nodes.Update(taintUpdates)
	return nil
}

func (p *CarbonPolicy) applyLabels() error {
	var labelUpdates Mapping[v1.Node]

	switch p.Spec.Labels.Type {
	case optNone:
		return nil
	case optBinary:
		labelUpdates = p.applyBinaryLabels(p.Spec.Labels.Thresholds[optAcceptable].Value)
	default:
		return fmt.Errorf("invalid value for .Labels.Type")
	}

	p.nodes.Update(labelUpdates)
	return nil
}

func (p *CarbonPolicy) applyTaintToWorstLocation(locationNames []string) Mapping[v1.Node] {
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

func (p *CarbonPolicy) applyBinaryLabels(threshold float64) Mapping[v1.Node] {
	result := Mapping[v1.Node]{}
	addAcceptableLabel := NewNodeBuilder("").
		SetLabel(labelIntensity, optAcceptable).
		Build()
	addUnacceptableLabel := NewNodeBuilder("").
		SetLabel(labelIntensity, optUnAcceptable).
		Build()

	update := addAcceptableLabel
	for location := range p.locations.Iterator() {
		if location.Intensity > threshold {
			update = addUnacceptableLabel
		}
		result[location.Name] = update
	}
	return result
}
