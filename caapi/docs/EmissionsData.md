# EmissionsData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Location** | Pointer to **NullableString** |  | [optional] 
**Time** | Pointer to **time.Time** |  | [optional] 
**Rating** | Pointer to **float64** |  | [optional] 
**Duration** | Pointer to **string** |  | [optional] 

## Methods

### NewEmissionsData

`func NewEmissionsData() *EmissionsData`

NewEmissionsData instantiates a new EmissionsData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmissionsDataWithDefaults

`func NewEmissionsDataWithDefaults() *EmissionsData`

NewEmissionsDataWithDefaults instantiates a new EmissionsData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocation

`func (o *EmissionsData) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *EmissionsData) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *EmissionsData) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *EmissionsData) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *EmissionsData) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *EmissionsData) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetTime

`func (o *EmissionsData) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *EmissionsData) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *EmissionsData) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *EmissionsData) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetRating

`func (o *EmissionsData) GetRating() float64`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *EmissionsData) GetRatingOk() (*float64, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *EmissionsData) SetRating(v float64)`

SetRating sets Rating field to given value.

### HasRating

`func (o *EmissionsData) HasRating() bool`

HasRating returns a boolean if a field has been set.

### GetDuration

`func (o *EmissionsData) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *EmissionsData) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *EmissionsData) SetDuration(v string)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *EmissionsData) HasDuration() bool`

HasDuration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


