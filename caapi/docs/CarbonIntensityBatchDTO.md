# CarbonIntensityBatchDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Location** | **string** | the location name where workflow is run | 
**StartTime** | **time.Time** | the time at which the workflow we are measuring carbon intensity for started | 
**EndTime** | **time.Time** | the time at which the workflow we are measuring carbon intensity for ended | 

## Methods

### NewCarbonIntensityBatchDTO

`func NewCarbonIntensityBatchDTO(location string, startTime time.Time, endTime time.Time, ) *CarbonIntensityBatchDTO`

NewCarbonIntensityBatchDTO instantiates a new CarbonIntensityBatchDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCarbonIntensityBatchDTOWithDefaults

`func NewCarbonIntensityBatchDTOWithDefaults() *CarbonIntensityBatchDTO`

NewCarbonIntensityBatchDTOWithDefaults instantiates a new CarbonIntensityBatchDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocation

`func (o *CarbonIntensityBatchDTO) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *CarbonIntensityBatchDTO) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *CarbonIntensityBatchDTO) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetStartTime

`func (o *CarbonIntensityBatchDTO) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *CarbonIntensityBatchDTO) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *CarbonIntensityBatchDTO) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.


### GetEndTime

`func (o *CarbonIntensityBatchDTO) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *CarbonIntensityBatchDTO) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *CarbonIntensityBatchDTO) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


