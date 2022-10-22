package main

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestLocations(t *testing.T) {
	suite.Run(t, new(LocationsSuite))
}

type LocationsSuite struct {
	suite.Suite
}

func (s *LocationsSuite) TestNoTranslator() {
	_locations := []Location{{
		Name: "us-east1",
	}, {
		Name: "us-central1",
	}}
	locations := NewLocations(_locations)

	s.Equal("us-central1", locations.GetLast().Name)
	var results []string
	for l := range locations.Iterator() {
		results = append(results, l.Name)
	}
	s.Equal("us-east1", results[0])
	s.Equal("us-central1", results[1])
}

func (s *LocationsSuite) TestTranslator() {
	_locations := []Location{{
		Name: "eastus",
	}, {
		Name: "centus",
	}}
	locations := NewLocations(_locations).Map(gcpRegions)

	s.Equal("us-central1", locations.GetLast().Name)
	var results []string
	for l := range locations.Iterator() {
		results = append(results, l.Name)
	}
	s.Equal("us-east1", results[0])
	s.Equal("us-central1", results[1])
}
