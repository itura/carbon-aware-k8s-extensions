/*
 * CarbonAware.WebApi, Version=1.0.0.0, Culture=neutral, PublicKeyToken=null
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package caapi

import (
	"encoding/json"
	"time"
)

// EmissionsDataDTO struct for EmissionsDataDTO
type EmissionsDataDTO struct {
	Location  NullableString `json:"location,omitempty"`
	Timestamp *time.Time     `json:"timestamp,omitempty"`
	Duration  *int32         `json:"duration,omitempty"`
	Value     *float64       `json:"value,omitempty"`
}

// NewEmissionsDataDTO instantiates a new EmissionsDataDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmissionsDataDTO() *EmissionsDataDTO {
	this := EmissionsDataDTO{}
	return &this
}

// NewEmissionsDataDTOWithDefaults instantiates a new EmissionsDataDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmissionsDataDTOWithDefaults() *EmissionsDataDTO {
	this := EmissionsDataDTO{}
	return &this
}

// GetLocation returns the Location field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmissionsDataDTO) GetLocation() string {
	if o == nil || o.Location.Get() == nil {
		var ret string
		return ret
	}
	return *o.Location.Get()
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmissionsDataDTO) GetLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Location.Get(), o.Location.IsSet()
}

// HasLocation returns a boolean if a field has been set.
func (o *EmissionsDataDTO) HasLocation() bool {
	if o != nil && o.Location.IsSet() {
		return true
	}

	return false
}

// SetLocation gets a reference to the given NullableString and assigns it to the Location field.
func (o *EmissionsDataDTO) SetLocation(v string) {
	o.Location.Set(&v)
}

// SetLocationNil sets the value for Location to be an explicit nil
func (o *EmissionsDataDTO) SetLocationNil() {
	o.Location.Set(nil)
}

// UnsetLocation ensures that no value is present for Location, not even an explicit nil
func (o *EmissionsDataDTO) UnsetLocation() {
	o.Location.Unset()
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *EmissionsDataDTO) GetTimestamp() time.Time {
	if o == nil || o.Timestamp == nil {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmissionsDataDTO) GetTimestampOk() (*time.Time, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *EmissionsDataDTO) HasTimestamp() bool {
	if o != nil && o.Timestamp != nil {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *EmissionsDataDTO) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *EmissionsDataDTO) GetDuration() int32 {
	if o == nil || o.Duration == nil {
		var ret int32
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmissionsDataDTO) GetDurationOk() (*int32, bool) {
	if o == nil || o.Duration == nil {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *EmissionsDataDTO) HasDuration() bool {
	if o != nil && o.Duration != nil {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int32 and assigns it to the Duration field.
func (o *EmissionsDataDTO) SetDuration(v int32) {
	o.Duration = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *EmissionsDataDTO) GetValue() float64 {
	if o == nil || o.Value == nil {
		var ret float64
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmissionsDataDTO) GetValueOk() (*float64, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *EmissionsDataDTO) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given float64 and assigns it to the Value field.
func (o *EmissionsDataDTO) SetValue(v float64) {
	o.Value = &v
}

func (o EmissionsDataDTO) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Location.IsSet() {
		toSerialize["location"] = o.Location.Get()
	}
	if o.Timestamp != nil {
		toSerialize["timestamp"] = o.Timestamp
	}
	if o.Duration != nil {
		toSerialize["duration"] = o.Duration
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableEmissionsDataDTO struct {
	value *EmissionsDataDTO
	isSet bool
}

func (v NullableEmissionsDataDTO) Get() *EmissionsDataDTO {
	return v.value
}

func (v *NullableEmissionsDataDTO) Set(val *EmissionsDataDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableEmissionsDataDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableEmissionsDataDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmissionsDataDTO(val *EmissionsDataDTO) *NullableEmissionsDataDTO {
	return &NullableEmissionsDataDTO{value: val, isSet: true}
}

func (v NullableEmissionsDataDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmissionsDataDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
