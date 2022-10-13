# EmissionsForecastDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestedAt** | **time.Time** | For current requests, this value is the timestamp the request for forecast data was made.  For historical forecast requests, this value is the timestamp used to access the most   recently generated forecast as of that time. | 
**Location** | **string** | The location of the forecast | 
**DataStartAt** | Pointer to **time.Time** | Start time boundary of forecasted data points. Ignores forecast data points before this time.  Defaults to the earliest time in the forecast data. | [optional] 
**DataEndAt** | Pointer to **time.Time** | End time boundary of forecasted data points. Ignores forecast data points after this time.  Defaults to the latest time in the forecast data. | [optional] 
**WindowSize** | Pointer to **int32** | The estimated duration (in minutes) of the workload.  Defaults to the duration of a single forecast data point. | [optional] 
**GeneratedAt** | Pointer to **time.Time** | Timestamp when the forecast was generated. | [optional] 
**OptimalDataPoints** | Pointer to [**[]EmissionsDataDTO**](EmissionsDataDTO.md) | The optimal forecasted data points within the &#39;forecastData&#39; array.  Returns empty array if &#39;forecastData&#39; array is empty. | [optional] 
**ForecastData** | Pointer to [**[]EmissionsDataDTO**](EmissionsDataDTO.md) | The forecasted data points transformed and filtered to reflect the specified time and window parameters.  Points are ordered chronologically; Empty array if all data points were filtered out.  E.G. dataStartAt and dataEndAt times outside the forecast period; windowSize greater than total duration of forecast data; | [optional] 

## Methods

### NewEmissionsForecastDTO

`func NewEmissionsForecastDTO(requestedAt time.Time, location string, ) *EmissionsForecastDTO`

NewEmissionsForecastDTO instantiates a new EmissionsForecastDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmissionsForecastDTOWithDefaults

`func NewEmissionsForecastDTOWithDefaults() *EmissionsForecastDTO`

NewEmissionsForecastDTOWithDefaults instantiates a new EmissionsForecastDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestedAt

`func (o *EmissionsForecastDTO) GetRequestedAt() time.Time`

GetRequestedAt returns the RequestedAt field if non-nil, zero value otherwise.

### GetRequestedAtOk

`func (o *EmissionsForecastDTO) GetRequestedAtOk() (*time.Time, bool)`

GetRequestedAtOk returns a tuple with the RequestedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedAt

`func (o *EmissionsForecastDTO) SetRequestedAt(v time.Time)`

SetRequestedAt sets RequestedAt field to given value.


### GetLocation

`func (o *EmissionsForecastDTO) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *EmissionsForecastDTO) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *EmissionsForecastDTO) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetDataStartAt

`func (o *EmissionsForecastDTO) GetDataStartAt() time.Time`

GetDataStartAt returns the DataStartAt field if non-nil, zero value otherwise.

### GetDataStartAtOk

`func (o *EmissionsForecastDTO) GetDataStartAtOk() (*time.Time, bool)`

GetDataStartAtOk returns a tuple with the DataStartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataStartAt

`func (o *EmissionsForecastDTO) SetDataStartAt(v time.Time)`

SetDataStartAt sets DataStartAt field to given value.

### HasDataStartAt

`func (o *EmissionsForecastDTO) HasDataStartAt() bool`

HasDataStartAt returns a boolean if a field has been set.

### GetDataEndAt

`func (o *EmissionsForecastDTO) GetDataEndAt() time.Time`

GetDataEndAt returns the DataEndAt field if non-nil, zero value otherwise.

### GetDataEndAtOk

`func (o *EmissionsForecastDTO) GetDataEndAtOk() (*time.Time, bool)`

GetDataEndAtOk returns a tuple with the DataEndAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataEndAt

`func (o *EmissionsForecastDTO) SetDataEndAt(v time.Time)`

SetDataEndAt sets DataEndAt field to given value.

### HasDataEndAt

`func (o *EmissionsForecastDTO) HasDataEndAt() bool`

HasDataEndAt returns a boolean if a field has been set.

### GetWindowSize

`func (o *EmissionsForecastDTO) GetWindowSize() int32`

GetWindowSize returns the WindowSize field if non-nil, zero value otherwise.

### GetWindowSizeOk

`func (o *EmissionsForecastDTO) GetWindowSizeOk() (*int32, bool)`

GetWindowSizeOk returns a tuple with the WindowSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowSize

`func (o *EmissionsForecastDTO) SetWindowSize(v int32)`

SetWindowSize sets WindowSize field to given value.

### HasWindowSize

`func (o *EmissionsForecastDTO) HasWindowSize() bool`

HasWindowSize returns a boolean if a field has been set.

### GetGeneratedAt

`func (o *EmissionsForecastDTO) GetGeneratedAt() time.Time`

GetGeneratedAt returns the GeneratedAt field if non-nil, zero value otherwise.

### GetGeneratedAtOk

`func (o *EmissionsForecastDTO) GetGeneratedAtOk() (*time.Time, bool)`

GetGeneratedAtOk returns a tuple with the GeneratedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratedAt

`func (o *EmissionsForecastDTO) SetGeneratedAt(v time.Time)`

SetGeneratedAt sets GeneratedAt field to given value.

### HasGeneratedAt

`func (o *EmissionsForecastDTO) HasGeneratedAt() bool`

HasGeneratedAt returns a boolean if a field has been set.

### GetOptimalDataPoints

`func (o *EmissionsForecastDTO) GetOptimalDataPoints() []EmissionsDataDTO`

GetOptimalDataPoints returns the OptimalDataPoints field if non-nil, zero value otherwise.

### GetOptimalDataPointsOk

`func (o *EmissionsForecastDTO) GetOptimalDataPointsOk() (*[]EmissionsDataDTO, bool)`

GetOptimalDataPointsOk returns a tuple with the OptimalDataPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptimalDataPoints

`func (o *EmissionsForecastDTO) SetOptimalDataPoints(v []EmissionsDataDTO)`

SetOptimalDataPoints sets OptimalDataPoints field to given value.

### HasOptimalDataPoints

`func (o *EmissionsForecastDTO) HasOptimalDataPoints() bool`

HasOptimalDataPoints returns a boolean if a field has been set.

### SetOptimalDataPointsNil

`func (o *EmissionsForecastDTO) SetOptimalDataPointsNil(b bool)`

 SetOptimalDataPointsNil sets the value for OptimalDataPoints to be an explicit nil

### UnsetOptimalDataPoints
`func (o *EmissionsForecastDTO) UnsetOptimalDataPoints()`

UnsetOptimalDataPoints ensures that no value is present for OptimalDataPoints, not even an explicit nil
### GetForecastData

`func (o *EmissionsForecastDTO) GetForecastData() []EmissionsDataDTO`

GetForecastData returns the ForecastData field if non-nil, zero value otherwise.

### GetForecastDataOk

`func (o *EmissionsForecastDTO) GetForecastDataOk() (*[]EmissionsDataDTO, bool)`

GetForecastDataOk returns a tuple with the ForecastData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForecastData

`func (o *EmissionsForecastDTO) SetForecastData(v []EmissionsDataDTO)`

SetForecastData sets ForecastData field to given value.

### HasForecastData

`func (o *EmissionsForecastDTO) HasForecastData() bool`

HasForecastData returns a boolean if a field has been set.

### SetForecastDataNil

`func (o *EmissionsForecastDTO) SetForecastDataNil(b bool)`

 SetForecastDataNil sets the value for ForecastData to be an explicit nil

### UnsetForecastData
`func (o *EmissionsForecastDTO) UnsetForecastData()`

UnsetForecastData ensures that no value is present for ForecastData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


