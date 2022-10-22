package main

type EmissionsDataSource interface {
	GetLocationData(locationNames []string) (*Locations, error)
}

func NewStubEDS() EmissionsDataSource {
	return &StubDataSource{}
}

type StubDataSource struct{}

func (s *StubDataSource) GetLocationData(locationNames []string) (*Locations, error) {
	return NewLocations([]Location{{
		Name:      "eastus",
		Rating:    50.0,
		Intensity: 34.3,
	}, {
		Name:      "centus",
		Rating:    77.7,
		Intensity: 68.2,
	}}), nil
}
