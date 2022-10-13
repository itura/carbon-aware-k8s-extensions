# \CarbonAwareApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchForecastDataAsync**](CarbonAwareApi.md#BatchForecastDataAsync) | **Post** /emissions/forecasts/batch | Given an array of historical forecasts, retrieves the data that contains  forecasts metadata, the optimal forecast and a range of forecasts filtered by the attributes [start...end] if provided.
[**GetAverageCarbonIntensity**](CarbonAwareApi.md#GetAverageCarbonIntensity) | **Get** /emissions/average-carbon-intensity | Retrieves the measured carbon intensity data between the time boundaries and calculates the average carbon intensity during that period.
[**GetAverageCarbonIntensityBatch**](CarbonAwareApi.md#GetAverageCarbonIntensityBatch) | **Post** /emissions/average-carbon-intensity/batch | Given an array of request objects, each with their own location and time boundaries, calculate the average carbon intensity for that location and time period   and return an array of carbon intensity objects.
[**GetBestEmissionsDataForLocationsByTime**](CarbonAwareApi.md#GetBestEmissionsDataForLocationsByTime) | **Get** /emissions/bylocations/best | Calculate the best emission data by list of locations for a specified time period.
[**GetCurrentForecastData**](CarbonAwareApi.md#GetCurrentForecastData) | **Get** /emissions/forecasts/current | Retrieves the most recent forecasted data and calculates the optimal marginal carbon intensity window.
[**GetEmissionsDataForLocationByTime**](CarbonAwareApi.md#GetEmissionsDataForLocationByTime) | **Get** /emissions/bylocation | Calculate the best emission data by location for a specified time period.
[**GetEmissionsDataForLocationsByTime**](CarbonAwareApi.md#GetEmissionsDataForLocationsByTime) | **Get** /emissions/bylocations | Calculate the observed emission data by list of locations for a specified time period.



## BatchForecastDataAsync

> []EmissionsForecastDTO BatchForecastDataAsync(ctx).EmissionsForecastBatchDTO(emissionsForecastBatchDTO).Execute()

Given an array of historical forecasts, retrieves the data that contains  forecasts metadata, the optimal forecast and a range of forecasts filtered by the attributes [start...end] if provided.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    emissionsForecastBatchDTO := []openapiclient.EmissionsForecastBatchDTO{*openapiclient.NewEmissionsForecastBatchDTO(time.Now(), "eastus")} // []EmissionsForecastBatchDTO | Array of requested forecasts. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CarbonAwareApi.BatchForecastDataAsync(context.Background()).EmissionsForecastBatchDTO(emissionsForecastBatchDTO).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `CarbonAwareApi.BatchForecastDataAsync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchForecastDataAsync`: []EmissionsForecastDTO
    fmt.Fprintf(os.Stdout, "Response from `CarbonAwareApi.BatchForecastDataAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchForecastDataAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **emissionsForecastBatchDTO** | [**[]EmissionsForecastBatchDTO**](EmissionsForecastBatchDTO.md) | Array of requested forecasts. | 

### Return type

[**[]EmissionsForecastDTO**](EmissionsForecastDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/json, application/_*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAverageCarbonIntensity

> CarbonIntensityDTO GetAverageCarbonIntensity(ctx).Location(location).StartTime(startTime).EndTime(endTime).Execute()

Retrieves the measured carbon intensity data between the time boundaries and calculates the average carbon intensity during that period.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    location := "location_example" // string | The location name of the region that we are measuring carbon usage in.
    startTime := time.Now() // time.Time | The time at which the workload and corresponding carbon usage begins.
    endTime := time.Now() // time.Time | The time at which the workload and corresponding carbon usage ends.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CarbonAwareApi.GetAverageCarbonIntensity(context.Background()).Location(location).StartTime(startTime).EndTime(endTime).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `CarbonAwareApi.GetAverageCarbonIntensity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAverageCarbonIntensity`: CarbonIntensityDTO
    fmt.Fprintf(os.Stdout, "Response from `CarbonAwareApi.GetAverageCarbonIntensity`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAverageCarbonIntensityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **location** | **string** | The location name of the region that we are measuring carbon usage in. | 
 **startTime** | **time.Time** | The time at which the workload and corresponding carbon usage begins. | 
 **endTime** | **time.Time** | The time at which the workload and corresponding carbon usage ends. | 

### Return type

[**CarbonIntensityDTO**](CarbonIntensityDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAverageCarbonIntensityBatch

> []CarbonIntensityDTO GetAverageCarbonIntensityBatch(ctx).CarbonIntensityBatchDTO(carbonIntensityBatchDTO).Execute()

Given an array of request objects, each with their own location and time boundaries, calculate the average carbon intensity for that location and time period   and return an array of carbon intensity objects.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    carbonIntensityBatchDTO := []openapiclient.CarbonIntensityBatchDTO{*openapiclient.NewCarbonIntensityBatchDTO("eastus", time.Now(), time.Now())} // []CarbonIntensityBatchDTO | Array of inputs where each contains a \"location\", \"startDate\", and \"endDate\" for which to calculate average marginal carbon intensity. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CarbonAwareApi.GetAverageCarbonIntensityBatch(context.Background()).CarbonIntensityBatchDTO(carbonIntensityBatchDTO).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `CarbonAwareApi.GetAverageCarbonIntensityBatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAverageCarbonIntensityBatch`: []CarbonIntensityDTO
    fmt.Fprintf(os.Stdout, "Response from `CarbonAwareApi.GetAverageCarbonIntensityBatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAverageCarbonIntensityBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **carbonIntensityBatchDTO** | [**[]CarbonIntensityBatchDTO**](CarbonIntensityBatchDTO.md) | Array of inputs where each contains a \&quot;location\&quot;, \&quot;startDate\&quot;, and \&quot;endDate\&quot; for which to calculate average marginal carbon intensity. | 

### Return type

[**[]CarbonIntensityDTO**](CarbonIntensityDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/json, application/_*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBestEmissionsDataForLocationsByTime

> []EmissionsData GetBestEmissionsDataForLocationsByTime(ctx).Location(location).Time(time).ToTime(toTime).DurationMinutes(durationMinutes).Execute()

Calculate the best emission data by list of locations for a specified time period.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    location := []string{"Inner_example"} // []string | String array of named locations.
    time := time.Now() // time.Time | [Optional] Start time for the data query. (optional)
    toTime := time.Now() // time.Time | [Optional] End time for the data query. (optional)
    durationMinutes := int32(56) // int32 | [Optional] Duration for the data query. (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CarbonAwareApi.GetBestEmissionsDataForLocationsByTime(context.Background()).Location(location).Time(time).ToTime(toTime).DurationMinutes(durationMinutes).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `CarbonAwareApi.GetBestEmissionsDataForLocationsByTime``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBestEmissionsDataForLocationsByTime`: []EmissionsData
    fmt.Fprintf(os.Stdout, "Response from `CarbonAwareApi.GetBestEmissionsDataForLocationsByTime`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBestEmissionsDataForLocationsByTimeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **location** | **[]string** | String array of named locations. | 
 **time** | **time.Time** | [Optional] Start time for the data query. | 
 **toTime** | **time.Time** | [Optional] End time for the data query. | 
 **durationMinutes** | **int32** | [Optional] Duration for the data query. | [default to 0]

### Return type

[**[]EmissionsData**](EmissionsData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurrentForecastData

> []EmissionsForecastDTO GetCurrentForecastData(ctx).Location(location).DataStartAt(dataStartAt).DataEndAt(dataEndAt).WindowSize(windowSize).Execute()

Retrieves the most recent forecasted data and calculates the optimal marginal carbon intensity window.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    location := []string{"Inner_example"} // []string | String array of named locations.
    dataStartAt := time.Now() // time.Time | Start time boundary of forecasted data points. Ignores current forecast data points before this time.  Defaults to the earliest time in the forecast data. (optional)
    dataEndAt := time.Now() // time.Time | End time boundary of forecasted data points. Ignores current forecast data points after this time.  Defaults to the latest time in the forecast data. (optional)
    windowSize := int32(56) // int32 | The estimated duration (in minutes) of the workload.  Defaults to the duration of a single forecast data point. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CarbonAwareApi.GetCurrentForecastData(context.Background()).Location(location).DataStartAt(dataStartAt).DataEndAt(dataEndAt).WindowSize(windowSize).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `CarbonAwareApi.GetCurrentForecastData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCurrentForecastData`: []EmissionsForecastDTO
    fmt.Fprintf(os.Stdout, "Response from `CarbonAwareApi.GetCurrentForecastData`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentForecastDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **location** | **[]string** | String array of named locations. | 
 **dataStartAt** | **time.Time** | Start time boundary of forecasted data points. Ignores current forecast data points before this time.  Defaults to the earliest time in the forecast data. | 
 **dataEndAt** | **time.Time** | End time boundary of forecasted data points. Ignores current forecast data points after this time.  Defaults to the latest time in the forecast data. | 
 **windowSize** | **int32** | The estimated duration (in minutes) of the workload.  Defaults to the duration of a single forecast data point. | 

### Return type

[**[]EmissionsForecastDTO**](EmissionsForecastDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEmissionsDataForLocationByTime

> []EmissionsData GetEmissionsDataForLocationByTime(ctx).Location(location).Time(time).ToTime(toTime).DurationMinutes(durationMinutes).Execute()

Calculate the best emission data by location for a specified time period.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    location := "location_example" // string | String named location.
    time := time.Now() // time.Time | [Optional] Start time for the data query. (optional)
    toTime := time.Now() // time.Time | [Optional] End time for the data query. (optional)
    durationMinutes := int32(56) // int32 | [Optional] Duration for the data query. (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CarbonAwareApi.GetEmissionsDataForLocationByTime(context.Background()).Location(location).Time(time).ToTime(toTime).DurationMinutes(durationMinutes).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `CarbonAwareApi.GetEmissionsDataForLocationByTime``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEmissionsDataForLocationByTime`: []EmissionsData
    fmt.Fprintf(os.Stdout, "Response from `CarbonAwareApi.GetEmissionsDataForLocationByTime`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEmissionsDataForLocationByTimeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **location** | **string** | String named location. | 
 **time** | **time.Time** | [Optional] Start time for the data query. | 
 **toTime** | **time.Time** | [Optional] End time for the data query. | 
 **durationMinutes** | **int32** | [Optional] Duration for the data query. | [default to 0]

### Return type

[**[]EmissionsData**](EmissionsData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEmissionsDataForLocationsByTime

> []EmissionsData GetEmissionsDataForLocationsByTime(ctx).Location(location).Time(time).ToTime(toTime).DurationMinutes(durationMinutes).Execute()

Calculate the observed emission data by list of locations for a specified time period.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    location := []string{"Inner_example"} // []string | String array of named locations.
    time := time.Now() // time.Time | [Optional] Start time for the data query. (optional)
    toTime := time.Now() // time.Time | [Optional] End time for the data query. (optional)
    durationMinutes := int32(56) // int32 | [Optional] Duration for the data query. (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CarbonAwareApi.GetEmissionsDataForLocationsByTime(context.Background()).Location(location).Time(time).ToTime(toTime).DurationMinutes(durationMinutes).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `CarbonAwareApi.GetEmissionsDataForLocationsByTime``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEmissionsDataForLocationsByTime`: []EmissionsData
    fmt.Fprintf(os.Stdout, "Response from `CarbonAwareApi.GetEmissionsDataForLocationsByTime`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEmissionsDataForLocationsByTimeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **location** | **[]string** | String array of named locations. | 
 **time** | **time.Time** | [Optional] Start time for the data query. | 
 **toTime** | **time.Time** | [Optional] End time for the data query. | 
 **durationMinutes** | **int32** | [Optional] Duration for the data query. | [default to 0]

### Return type

[**[]EmissionsData**](EmissionsData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

