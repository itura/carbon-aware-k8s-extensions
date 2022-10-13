# CarbonIntensityDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Location** | **string** | the location name where workflow is run | 
**StartTime** | **time.Time** | the time at which the workflow we are measuring carbon intensity for started | 
**EndTime** | **time.Time** | the time at which the workflow we are measuring carbon intensity for ended | 
**CarbonIntensity** | Pointer to **float64** | Value of the marginal carbon intensity in grams per kilowatt-hour. | [optional] 

## Methods

### NewCarbonIntensityDTO

`func NewCarbonIntensityDTO(location string, startTime time.Time, endTime time.Time, ) *CarbonIntensityDTO`

NewCarbonIntensityDTO instantiates a new CarbonIntensityDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCarbonIntensityDTOWithDefaults

`func NewCarbonIntensityDTOWithDefaults() *CarbonIntensityDTO`

NewCarbonIntensityDTOWithDefaults instantiates a new CarbonIntensityDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocation

`func (o *CarbonIntensityDTO) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *CarbonIntensityDTO) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *CarbonIntensityDTO) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetStartTime

`func (o *CarbonIntensityDTO) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *CarbonIntensityDTO) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *CarbonIntensityDTO) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.


### GetEndTime

`func (o *CarbonIntensityDTO) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *CarbonIntensityDTO) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *CarbonIntensityDTO) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.


### GetCarbonIntensity

`func (o *CarbonIntensityDTO) GetCarbonIntensity() float64`

GetCarbonIntensity returns the CarbonIntensity field if non-nil, zero value otherwise.

### GetCarbonIntensityOk

`func (o *CarbonIntensityDTO) GetCarbonIntensityOk() (*float64, bool)`

GetCarbonIntensityOk returns a tuple with the CarbonIntensity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarbonIntensity

`func (o *CarbonIntensityDTO) SetCarbonIntensity(v float64)`

SetCarbonIntensity sets CarbonIntensity field to given value.

### HasCarbonIntensity

`func (o *CarbonIntensityDTO) HasCarbonIntensity() bool`

HasCarbonIntensity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


