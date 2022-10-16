package main

import "k8s.io/api/core/v1"

const (
	labelK8sRegion               = "topology.kubernetes.io/region"
	labelIntensity               = "greensoftware.foundation/carbon-intensity"
	intensityHigh                = "high"
	intensityLow                 = "low"
	dataSourceCAAPI              = "CarbonAwareAPI"
	dataSourceCCF                = "CloudCarbonFootprint"
	policySortByCurrentIntensity = "currentIntensity"
	policySortByRating           = "rating"
	policyTaintTypeWorst         = "worst"
	policyTaintTypeTest          = "test"
)

func taintHighIntensity(effect v1.TaintEffect) v1.Taint {
	return v1.Taint{
		Key:    labelIntensity,
		Value:  intensityHigh,
		Effect: effect,
	}
}

type Mapping[T any] map[string]T

func (m Mapping[T]) Merge(other Mapping[T]) Mapping[T] {
	result := Mapping[T]{}
	for k, v := range m {
		result[k] = v
	}
	for k, v := range other {
		result[k] = v
	}
	return result
}
