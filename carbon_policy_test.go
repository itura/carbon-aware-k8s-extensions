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

func (s *CarbonPolicySuite) TestUpdateNodesWithDefaultPolicy() {
	policy := NewCarbonPolicy(CarbonPolicySpec{}).
		SetNodes(NewNodes([]v1.Node{
			NewNodeBuilder("n1").
				SetRegion("us-east1").
				Build(),
			NewNodeBuilder("n2").
				SetRegion("us-central1").
				Build(),
		})).
		SetLocations(NewLocations([]Location{{
			Name:      "us-east1",
			Intensity: 2.0,
		}, {
			Name:      "us-central1",
			Intensity: 1.0,
		}}))

	updatedNodes, err := policy.UpdateNodes()

	s.Nil(err)
	s.Equal(
		NewNodeBuilder("n1").
			SetRegion("us-east1").
			AddTaint(taintHighIntensity(v1.TaintEffectNoSchedule)).
			Build(),
		updatedNodes.Get(0),
	)
	s.Equal(
		NewNodeBuilder("n2").
			SetRegion("us-central1").
			Build(),
		updatedNodes.Get(1),
	)
}

// 0 rating == good, 100 rating == bad
func (s *CarbonPolicySuite) TestUpdateNodesBasedOnRating() {
	policy := NewCarbonPolicy(CarbonPolicySpec{
		SortBy: policySortByRating,
	}).
		SetNodes(NewNodes([]v1.Node{
			NewNodeBuilder("n1").
				SetRegion("us-east1").
				Build(),
			NewNodeBuilder("n2").
				SetRegion("us-central1").
				Build(),
		})).
		SetLocations(NewLocations([]Location{{
			Name:   "us-east1",
			Rating: 2.0,
		}, {
			Name:   "us-central1",
			Rating: 1.0,
		}}))

	updatedNodes, err := policy.UpdateNodes()

	s.Nil(err)
	s.Equal(
		NewNodeBuilder("n1").
			SetRegion("us-east1").
			AddTaint(taintHighIntensity(v1.TaintEffectNoSchedule)).
			Build(),
		updatedNodes.Get(0),
	)
	s.Equal(
		NewNodeBuilder("n2").
			SetRegion("us-central1").
			Build(),
		updatedNodes.Get(1),
	)
}

func (s *CarbonPolicySuite) TestUpdateNodesWithTaintEffect() {
	policy := NewCarbonPolicy(CarbonPolicySpec{
		Taints: TaintPolicy{
			Effect: v1.TaintEffectPreferNoSchedule,
		},
	}).
		SetNodes(NewNodes([]v1.Node{
			NewNodeBuilder("n1").
				SetRegion("us-east1").
				Build(),
			NewNodeBuilder("n2").
				SetRegion("us-central1").
				Build(),
		})).
		SetLocations(NewLocations([]Location{{
			Name:      "us-east1",
			Intensity: 2.0,
		}, {
			Name:      "us-central1",
			Intensity: 1.0,
		}}))

	updatedNodes, err := policy.UpdateNodes()

	s.Nil(err)
	s.Equal(
		NewNodeBuilder("n1").
			SetRegion("us-east1").
			AddTaint(taintHighIntensity(v1.TaintEffectPreferNoSchedule)).
			Build(),
		updatedNodes.Get(0),
	)
	s.Equal(
		NewNodeBuilder("n2").
			SetRegion("us-central1").
			Build(),
		updatedNodes.Get(1),
	)
}

func (s *CarbonPolicySuite) TestUpdateNodesWithOnlyOneLocation() {
	policy := NewCarbonPolicy(CarbonPolicySpec{}).
		SetNodes(NewNodes([]v1.Node{
			NewNodeBuilder("n1").
				SetRegion("us-east1").
				Build(),
		})).
		SetLocations(NewLocations([]Location{{
			Name:      "us-east1",
			Intensity: 2.0,
		}}))

	updatedNodes, err := policy.UpdateNodes()

	s.Nil(err)
	// taint is not applied
	s.Equal(
		NewNodeBuilder("n1").
			SetRegion("us-east1").
			Build(),
		updatedNodes.Get(0),
	)
}

func (s *CarbonPolicySuite) TestUpdateNodesWithoutLocations() {
	policy := NewCarbonPolicy(CarbonPolicySpec{}).
		SetNodes(NewNodes([]v1.Node{
			NewNodeBuilder("n1").
				SetRegion("us-east1").
				Build(),
		}))

	updatedNodes, err := policy.UpdateNodes()

	s.Error(err, "location data missing from policy")
	s.Nil(updatedNodes)
}

func (s *CarbonPolicySuite) TestUpdateNodesWithoutNodes() {
	policy := NewCarbonPolicy(CarbonPolicySpec{}).
		SetLocations(NewLocations([]Location{{
			Name:      "us-east1",
			Intensity: 2.0,
		}}))

	updatedNodes, err := policy.UpdateNodes()

	s.Error(err, "node data missing from policy")
	s.Nil(updatedNodes)
}