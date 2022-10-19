package main

import (
	"fmt"
	v1 "k8s.io/api/core/v1"
)

type CarbonPolicySpec struct {
	DataSource string    `yaml:"dataSource"`
	SortBy     string    `yaml:"sortBy"`
	Taints     TaintSpec `yaml:"taints"`
	Labels     LabelSpec `yaml:"labels"`
}

type TaintSpec struct {
	Type   string         `yaml:"type"`
	Effect v1.TaintEffect `yaml:"type"`
}

type LabelSpec struct {
	Type       string                   `yaml:"type"`
	Thresholds map[string]ThresholdSpec `yaml:"thresholds"`
}

func DefaultLabelSpec(spec LabelSpec) LabelSpec {
	if spec.Type == "" {
		spec.Type = policyLabelTypeBinary
	}
	if len(spec.Thresholds) == 0 {
		spec.Thresholds = map[string]ThresholdSpec{
			intensityAcceptable: DefaultThresholdSpec(ThresholdSpec{}),
		}
	}
	return spec
}

type ThresholdSpec struct {
	Type  string  `yaml:"type"`
	Value float64 `yaml:"value"`
}

func DefaultThresholdSpec(spec ThresholdSpec) ThresholdSpec {
	if spec.Type == "" {
		spec.Type = comparisonLessThan
	}
	if spec.Value == 0 {
		spec.Value = 10
	}

	return spec
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
	spec.Labels = DefaultLabelSpec(spec.Labels)
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

	locationNames := p.nodes.GetRegions()
	if len(locationNames) < 2 && p.Spec.Taints.Type != policyTaintTypeTest {
		return p.nodes, nil
	}

	if p.Spec.SortBy == policySortByCurrentIntensity {
		p.locations.SortByIntensity()
	} else if p.Spec.SortBy == policySortByRating {
		p.locations.SortByRating()
	} else {
		return nil, fmt.Errorf("invalid value for .SortBy")
	}

	var taintUpdates Mapping[v1.Node]
	if p.Spec.Taints.Type == policyTaintTypeWorst || p.Spec.Taints.Type == policyTaintTypeTest {
		taintUpdates = p.applyTaintToWorstLocation(locationNames)
	} else {
		return nil, fmt.Errorf("invalid value for .Taints.Type")
	}
	p.nodes.Update(taintUpdates)

	var labelUpdates Mapping[v1.Node]
	if p.Spec.Labels.Type == policyLabelTypeBinary {
		labelUpdates = p.applyBinaryLabels(p.Spec.Labels.Thresholds[intensityAcceptable].Value)
	}
	p.nodes.Update(labelUpdates)

	return p.nodes, nil
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
		SetLabel(labelIntensity, intensityAcceptable).
		Build()
	addUnacceptableLabel := NewNodeBuilder("").
		SetLabel(labelIntensity, intensityUnaccaptable).
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
