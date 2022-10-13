# EmissionsForecastBatchDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestedAt** | **time.Time** | For historical forecast requests, this value is the timestamp used to access the most  recently generated forecast as of that time. | 
**DataStartAt** | Pointer to **time.Time** | Start time boundary of forecasted data points. Ignores forecast data points before this time.  Defaults to the earliest time in the forecast data. | [optional] 
**DataEndAt** | Pointer to **time.Time** | End time boundary of forecasted data points. Ignores forecast data points after this time.  Defaults to the latest time in the forecast data. | [optional] 
**WindowSize** | Pointer to **int32** | The estimated duration (in minutes) of the workload.  Defaults to the duration of a single forecast data point. | [optional] 
**Location** | **string** | The location of the forecast | 

## Methods

### NewEmissionsForecastBatchDTO

`func NewEmissionsForecastBatchDTO(requestedAt time.Time, location string, ) *EmissionsForecastBatchDTO`

NewEmissionsForecastBatchDTO instantiates a new EmissionsForecastBatchDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmissionsForecastBatchDTOWithDefaults

`func NewEmissionsForecastBatchDTOWithDefaults() *EmissionsForecastBatchDTO`

NewEmissionsForecastBatchDTOWithDefaults instantiates a new EmissionsForecastBatchDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestedAt

`func (o *EmissionsForecastBatchDTO) GetRequestedAt() time.Time`

GetRequestedAt returns the RequestedAt field if non-nil, zero value otherwise.

### GetRequestedAtOk

`func (o *EmissionsForecastBatchDTO) GetRequestedAtOk() (*time.Time, bool)`

GetRequestedAtOk returns a tuple with the RequestedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedAt

`func (o *EmissionsForecastBatchDTO) SetRequestedAt(v time.Time)`

SetRequestedAt sets RequestedAt field to given value.


### GetDataStartAt

`func (o *EmissionsForecastBatchDTO) GetDataStartAt() time.Time`

GetDataStartAt returns the DataStartAt field if non-nil, zero value otherwise.

### GetDataStartAtOk

`func (o *EmissionsForecastBatchDTO) GetDataStartAtOk() (*time.Time, bool)`

GetDataStartAtOk returns a tuple with the DataStartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataStartAt

`func (o *EmissionsForecastBatchDTO) SetDataStartAt(v time.Time)`

SetDataStartAt sets DataStartAt field to given value.

### HasDataStartAt

`func (o *EmissionsForecastBatchDTO) HasDataStartAt() bool`

HasDataStartAt returns a boolean if a field has been set.

### GetDataEndAt

`func (o *EmissionsForecastBatchDTO) GetDataEndAt() time.Time`

GetDataEndAt returns the DataEndAt field if non-nil, zero value otherwise.

### GetDataEndAtOk

`func (o *EmissionsForecastBatchDTO) GetDataEndAtOk() (*time.Time, bool)`

GetDataEndAtOk returns a tuple with the DataEndAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataEndAt

`func (o *EmissionsForecastBatchDTO) SetDataEndAt(v time.Time)`

SetDataEndAt sets DataEndAt field to given value.

### HasDataEndAt

`func (o *EmissionsForecastBatchDTO) HasDataEndAt() bool`

HasDataEndAt returns a boolean if a field has been set.

### GetWindowSize

`func (o *EmissionsForecastBatchDTO) GetWindowSize() int32`

GetWindowSize returns the WindowSize field if non-nil, zero value otherwise.

### GetWindowSizeOk

`func (o *EmissionsForecastBatchDTO) GetWindowSizeOk() (*int32, bool)`

GetWindowSizeOk returns a tuple with the WindowSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowSize

`func (o *EmissionsForecastBatchDTO) SetWindowSize(v int32)`

SetWindowSize sets WindowSize field to given value.

### HasWindowSize

`func (o *EmissionsForecastBatchDTO) HasWindowSize() bool`

HasWindowSize returns a boolean if a field has been set.

### GetLocation

`func (o *EmissionsForecastBatchDTO) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *EmissionsForecastBatchDTO) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *EmissionsForecastBatchDTO) SetLocation(v string)`

SetLocation sets Location field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


