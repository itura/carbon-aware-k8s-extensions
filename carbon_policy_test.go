package main

import (
	"github.com/stretchr/testify/suite"
	v1 "k8s.io/api/core/v1"
	"testing"
)

func TestCarbonPolicy(t *testing.T) {
	suite.Run(t, new(CarbonPolicySuite))
}

type CarbonPolicySuite struct {
	suite.Suite
}

func (s *CarbonPolicySuite) TestNodeUpdatesForDefaultPolicy() {
	policy := NewCarbonPolicy(CarbonPolicySpec{}).
		SetLocations(NewLocations([]Location{{
			Name:      "us-east1",
			Intensity: 2.0,
		}, {
			Name:      "us-central1",
			Intensity: 1.0,
		}}))

	nodeUpdates := policy.NodeUpdatesByLocation([]string{"us-east1", "us-central1"})

	s.Equal(
		Mapping[v1.Node]{
			"us-east1": NewNodeBuilder("").
				AddTaint(taintHighIntensity(v1.TaintEffectNoSchedule)).
				Build(),
			"us-central1": NewNodeBuilder("").Build(),
		},
		nodeUpdates,
	)
}

// 0 rating == good, 100 rating == bad
func (s *CarbonPolicySuite) TestNodeUpdatesForSortByRating() {
	policy := NewCarbonPolicy(CarbonPolicySpec{
		SortBy: policySortByRating,
	}).
		SetLocations(NewLocations([]Location{{
			Name:   "us-east1",
			Rating: 2.0,
		}, {
			Name:   "us-central1",
			Rating: 1.0,
		}}))

	nodeUpdates := policy.NodeUpdatesByLocation([]string{"us-east1", "us-central1"})

	s.Equal(
		Mapping[v1.Node]{
			"us-east1": NewNodeBuilder("").
				AddTaint(taintHighIntensity(v1.TaintEffectNoSchedule)).
				Build(),
			"us-central1": NewNodeBuilder("").Build(),
		},
		nodeUpdates,
	)
}
