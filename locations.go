package main

import "sort"

type LocationMapper func(location Location) Location

var _gcpRegions = Mapping[string]{
	"eastus": "us-east1",
	"centus": "us-central1",
	"westus": "us-west1",
}

func gcpRegions(l Location) Location {
	l.Name = _gcpRegions[l.Name]
	return l
}

type Location struct {
	Name      string
	Rating    float64
	Intensity float64
}

type Locations struct {
	locations  []Location
	lessThanFn func(i, j int) bool
}

func NewLocations(locations []Location) *Locations {
	return &Locations{
		locations: locations,
		lessThanFn: func(i, j int) bool {
			return false
		},
	}
}

func (l *Locations) Map(fn LocationMapper) *Locations {
	for i, location := range l.locations {
		l.locations[i] = fn(location)
	}
	return l
}

func (l *Locations) Len() int {
	return len(l.locations)
}

func (l *Locations) Less(i, j int) bool {
	return l.lessThanFn(i, j)
}

func (l *Locations) Swap(i, j int) {
	l.locations[i], l.locations[j] = l.locations[j], l.locations[i]
}

func (l *Locations) SortByIntensity() *Locations {
	l.lessThanFn = func(i, j int) bool {
		return l.locations[i].Intensity < l.locations[j].Intensity
	}
	sort.Sort(l)
	return l
}

func (l *Locations) SortByRating() *Locations {
	l.lessThanFn = func(i, j int) bool {
		return l.locations[i].Rating < l.locations[j].Rating
	}
	sort.Sort(l)
	return l
}

func (l *Locations) GetLast() Location {
	return l.locations[len(l.locations)-1]
}

func (l *Locations) Iterator() <-chan Location {
	ch := make(chan Location, len(l.locations))
	go func() {
		for _, location := range l.locations {
			ch <- location
		}
		close(ch)
	}()
	return ch
}
