package main

import (
	"fmt"
	"k8s.io/api/core/v1"
	"reflect"
)

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

func (s TaintSpec) Init() (TaintSpec, error) {
	errs := Validate(&s,
		NewOption("Type").
			SetValues(optNone, optWorst).
			SetDefault(optWorst),
		NewOption("Effect").
			SetValues(v1.TaintEffectPreferNoSchedule, v1.TaintEffectNoSchedule, v1.TaintEffectNoExecute).
			SetDefault(v1.TaintEffectPreferNoSchedule),
		NewOption("ShouldTaintOnlyLocation").
			SetDefault(false),
	)

	if len(errs) > 0 {
		return s, fmt.Errorf("ouch")
	}

	return s, nil
}

func Validate(spec interface{}, options ...*Option) Mapping[ValidationErrors] {
	val := reflect.ValueOf(spec).Elem()

	errors := Mapping[ValidationErrors]{}
	for _, option := range options {
		result := option.SetInitial(val).Build()
		if result != nil {
			errors[option.Key] = ValidationErrors{result}
		}
	}
	if len(errors) > 0 {
		return errors
	}
	return nil
}

type ValidationErrors []error

type Option struct {
	currentValue  reflect.Value
	defaultValue  reflect.Value
	allowedValues []reflect.Value
	Key           string
}

func NewOption(key string) *Option {
	return &Option{Key: key}
}

func (f *Option) SetInitial(v reflect.Value) *Option {
	f.currentValue = v.FieldByName(f.Key)
	return f
}

func (f *Option) SetValues(vs ...any) *Option {
	for _, v := range vs {
		value := reflect.ValueOf(v)
		f.allowedValues = append(f.allowedValues, value)
	}
	return f
}

func (f *Option) SetDefault(v any) *Option {
	value := reflect.ValueOf(v)
	f.defaultValue = value
	return f
}

func (f *Option) Build() error {
	var value any
	if f.currentValue.IsZero() && !f.defaultValue.IsZero() {
		value = f.defaultValue.Interface()
	} else {
		value = f.currentValue.Interface()
	}
	if f.allowedValues != nil {
		isAllowed := false
		for _, v := range f.allowedValues {
			if v.Interface() == value {
				isAllowed = true
				break
			}
		}
		if !isAllowed {
			return fmt.Errorf("invalid value (%v), expected one of %v", value, f.allowedValues)
		}
	}
	if f.currentValue.CanSet() {
		f.currentValue.Set(reflect.ValueOf(value))
	}
	return nil
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
