package main

import (
	"fmt"
	v1 "k8s.io/api/core/v1"
)

type CarbonPolicySpec struct {
	DataSource DataSourceSpec `yaml:"dataSource"`
	Taints     TaintSpec      `yaml:"taints"`
	Labels     LabelSpec      `yaml:"labels"`
}

type DataSourceSpec struct {
	Type   string `yaml:"type"`
	SortBy string `yaml:"sortBy"`
}

func DefaultDataSourceSpec(spec DataSourceSpec) DataSourceSpec {
	if spec.Type == "" {
		spec.Type = optCAAPI
	}
	if spec.SortBy == "" {
		spec.SortBy = optCurrentIntensity
	}
	return spec
}

type TaintSpec struct {
	Type                    string         `yaml:"type"`
	Effect                  v1.TaintEffect `yaml:"effect"`
	ShouldTaintOnlyLocation bool           `yaml:"shouldTaintOnlyLocation"`
}

func DefaultTaintSpec(spec TaintSpec) TaintSpec {
	if spec.Type == "" {
		spec.Type = optWorst
	}
	if spec.Effect == "" {
		spec.Effect = v1.TaintEffectPreferNoSchedule
	}
	return spec
}

type LabelSpec struct {
	Type       string                   `yaml:"type"`
	Thresholds map[string]ThresholdSpec `yaml:"thresholds"`
}

func DefaultLabelSpec(spec LabelSpec) LabelSpec {
	if spec.Type == "" {
		spec.Type = optBinary
	}
	if len(spec.Thresholds) == 0 {
		spec.Thresholds = map[string]ThresholdSpec{
			optAcceptable: DefaultThresholdSpec(ThresholdSpec{}),
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
		spec.Type = optLessThan
	}
	if spec.Value == 0 {
		spec.Value = 50
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
	spec.DataSource = DefaultDataSourceSpec(spec.DataSource)
	spec.Taints = DefaultTaintSpec(spec.Taints)
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

	if p.Spec.DataSource.SortBy == optCurrentIntensity {
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
