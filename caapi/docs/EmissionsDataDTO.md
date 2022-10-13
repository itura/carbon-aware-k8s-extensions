# EmissionsDataDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Location** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Duration** | Pointer to **int32** |  | [optional] 
**Value** | Pointer to **float64** |  | [optional] 

## Methods

### NewEmissionsDataDTO

`func NewEmissionsDataDTO() *EmissionsDataDTO`

NewEmissionsDataDTO instantiates a new EmissionsDataDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmissionsDataDTOWithDefaults

`func NewEmissionsDataDTOWithDefaults() *EmissionsDataDTO`

NewEmissionsDataDTOWithDefaults instantiates a new EmissionsDataDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocation

`func (o *EmissionsDataDTO) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *EmissionsDataDTO) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *EmissionsDataDTO) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *EmissionsDataDTO) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *EmissionsDataDTO) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *EmissionsDataDTO) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetTimestamp

`func (o *EmissionsDataDTO) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *EmissionsDataDTO) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *EmissionsDataDTO) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *EmissionsDataDTO) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetDuration

`func (o *EmissionsDataDTO) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *EmissionsDataDTO) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *EmissionsDataDTO) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *EmissionsDataDTO) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetValue

`func (o *EmissionsDataDTO) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *EmissionsDataDTO) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *EmissionsDataDTO) SetValue(v float64)`

SetValue sets Value field to given value.

### HasValue

`func (o *EmissionsDataDTO) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


