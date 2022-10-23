package main

import "k8s.io/api/core/v1"

type CarbonPolicySpec struct {
	DataSource DataSourceSpec `yaml:"dataSource"`
	Taints     TaintSpec      `yaml:"taints"`
	Labels     LabelSpec      `yaml:"labels"`
}

func (spec CarbonPolicySpec) SetDefaults() CarbonPolicySpec {
	spec.DataSource = DefaultDataSourceSpec(spec.DataSource)
	spec.Taints = DefaultTaintSpec(spec.Taints)
	spec.Labels = DefaultLabelSpec(spec.Labels)
	return spec
}

type DataSourceSpec struct {
	Type      string       `yaml:"type"`
	SortBy    string       `yaml:"sortBy"`
	Locations LocationSpec `yaml:"locations"`
}

func DefaultDataSourceSpec(spec DataSourceSpec) DataSourceSpec {
	if spec.Type == "" {
		spec.Type = optStub
	}
	if spec.SortBy == "" {
		spec.SortBy = optIntensity
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

type LocationSpec struct {
	Preset string `yaml:"preset"`
}
