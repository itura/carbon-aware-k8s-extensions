package main

import "k8s.io/api/core/v1"

const (
	labelK8sRegion  = "topology.kubernetes.io/region"
	labelIntensity  = "greensoftware.foundation/carbon-intensity"
	optHigh         = "high"
	optMedium       = "medium"
	optLow          = "low"
	optAcceptable   = "acceptable"
	optUnAcceptable = "unacceptable"
	optCAAPI        = "CarbonAwareAPI"
	optCCF          = "CloudCarbonFootprint"
	optIntensity    = "intensity"
	optRating       = "rating"
	optWorst        = "worst"
	optTest         = "test"
	optBinary       = "binary"
	optScale        = "scale"
	optLessThan     = "lessThan"
	optNone         = "none"
	optStub         = "stub"
	optGcp          = "gcp"
)

func taintHighIntensity(effect v1.TaintEffect) v1.Taint {
	return v1.Taint{
		Key:    labelIntensity,
		Value:  optHigh,
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

func Id[T any](x T) T {
	return x
}
