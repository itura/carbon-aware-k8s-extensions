package main

import (
	"ch22/caapi"
	"context"
	"net/http"
	"time"
)

const dateLayout = "2006-01-02"

func dateToTime(date string) time.Time {
	r, err := time.Parse(dateLayout, date)
	if err != nil {
		panic(err.Error())
	}
	return r
}

type CAClient struct {
	underlying *caapi.APIClient
}

func NewCAClient(host string) *CAClient {
	config := caapi.NewConfiguration()
	config.Servers = caapi.ServerConfigurations{
		{
			URL: host,
			Variables: map[string]caapi.ServerVariable{
				"basePath": {
					DefaultValue: "emissions",
				},
			},
		},
	}
	return &CAClient{
		underlying: caapi.NewAPIClient(config),
	}
}

func (c *CAClient) GetAverageCarbonIntensity(location string, startDate string, endDate string) (float64, error) {
	ca := c.underlying.CarbonAwareApi
	data, response, err := ca.GetAverageCarbonIntensity(context.Background()).
		Location(location).
		StartTime(dateToTime(startDate)).
		EndTime(dateToTime(endDate)).
		Execute()

	if response.StatusCode != http.StatusOK {
		return 0, err
	}

	return data.GetCarbonIntensity(), nil
}
