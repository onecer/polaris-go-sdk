/*
Polaris Server

一个支持多语言、多框架的云原生服务发现和治理中心  提供高性能SDK和无侵入Sidecar两种接入方式  

API version: v0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package polaris

import (
	"encoding/json"
)

// V1Amount struct for V1Amount
type V1Amount struct {
	MaxAmount *int32 `json:"maxAmount,omitempty"`
	MinAmount *int32 `json:"minAmount,omitempty"`
	Precision *int32 `json:"precision,omitempty"`
	StartAmount *int32 `json:"startAmount,omitempty"`
	ValidDuration *DurationpbDuration `json:"validDuration,omitempty"`
}

// NewV1Amount instantiates a new V1Amount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1Amount() *V1Amount {
	this := V1Amount{}
	return &this
}

// NewV1AmountWithDefaults instantiates a new V1Amount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1AmountWithDefaults() *V1Amount {
	this := V1Amount{}
	return &this
}

// GetMaxAmount returns the MaxAmount field value if set, zero value otherwise.
func (o *V1Amount) GetMaxAmount() int32 {
	if o == nil || o.MaxAmount == nil {
		var ret int32
		return ret
	}
	return *o.MaxAmount
}

// GetMaxAmountOk returns a tuple with the MaxAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Amount) GetMaxAmountOk() (*int32, bool) {
	if o == nil || o.MaxAmount == nil {
		return nil, false
	}
	return o.MaxAmount, true
}

// HasMaxAmount returns a boolean if a field has been set.
func (o *V1Amount) HasMaxAmount() bool {
	if o != nil && o.MaxAmount != nil {
		return true
	}

	return false
}

// SetMaxAmount gets a reference to the given int32 and assigns it to the MaxAmount field.
func (o *V1Amount) SetMaxAmount(v int32) {
	o.MaxAmount = &v
}

// GetMinAmount returns the MinAmount field value if set, zero value otherwise.
func (o *V1Amount) GetMinAmount() int32 {
	if o == nil || o.MinAmount == nil {
		var ret int32
		return ret
	}
	return *o.MinAmount
}

// GetMinAmountOk returns a tuple with the MinAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Amount) GetMinAmountOk() (*int32, bool) {
	if o == nil || o.MinAmount == nil {
		return nil, false
	}
	return o.MinAmount, true
}

// HasMinAmount returns a boolean if a field has been set.
func (o *V1Amount) HasMinAmount() bool {
	if o != nil && o.MinAmount != nil {
		return true
	}

	return false
}

// SetMinAmount gets a reference to the given int32 and assigns it to the MinAmount field.
func (o *V1Amount) SetMinAmount(v int32) {
	o.MinAmount = &v
}

// GetPrecision returns the Precision field value if set, zero value otherwise.
func (o *V1Amount) GetPrecision() int32 {
	if o == nil || o.Precision == nil {
		var ret int32
		return ret
	}
	return *o.Precision
}

// GetPrecisionOk returns a tuple with the Precision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Amount) GetPrecisionOk() (*int32, bool) {
	if o == nil || o.Precision == nil {
		return nil, false
	}
	return o.Precision, true
}

// HasPrecision returns a boolean if a field has been set.
func (o *V1Amount) HasPrecision() bool {
	if o != nil && o.Precision != nil {
		return true
	}

	return false
}

// SetPrecision gets a reference to the given int32 and assigns it to the Precision field.
func (o *V1Amount) SetPrecision(v int32) {
	o.Precision = &v
}

// GetStartAmount returns the StartAmount field value if set, zero value otherwise.
func (o *V1Amount) GetStartAmount() int32 {
	if o == nil || o.StartAmount == nil {
		var ret int32
		return ret
	}
	return *o.StartAmount
}

// GetStartAmountOk returns a tuple with the StartAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Amount) GetStartAmountOk() (*int32, bool) {
	if o == nil || o.StartAmount == nil {
		return nil, false
	}
	return o.StartAmount, true
}

// HasStartAmount returns a boolean if a field has been set.
func (o *V1Amount) HasStartAmount() bool {
	if o != nil && o.StartAmount != nil {
		return true
	}

	return false
}

// SetStartAmount gets a reference to the given int32 and assigns it to the StartAmount field.
func (o *V1Amount) SetStartAmount(v int32) {
	o.StartAmount = &v
}

// GetValidDuration returns the ValidDuration field value if set, zero value otherwise.
func (o *V1Amount) GetValidDuration() DurationpbDuration {
	if o == nil || o.ValidDuration == nil {
		var ret DurationpbDuration
		return ret
	}
	return *o.ValidDuration
}

// GetValidDurationOk returns a tuple with the ValidDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Amount) GetValidDurationOk() (*DurationpbDuration, bool) {
	if o == nil || o.ValidDuration == nil {
		return nil, false
	}
	return o.ValidDuration, true
}

// HasValidDuration returns a boolean if a field has been set.
func (o *V1Amount) HasValidDuration() bool {
	if o != nil && o.ValidDuration != nil {
		return true
	}

	return false
}

// SetValidDuration gets a reference to the given DurationpbDuration and assigns it to the ValidDuration field.
func (o *V1Amount) SetValidDuration(v DurationpbDuration) {
	o.ValidDuration = &v
}

func (o V1Amount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MaxAmount != nil {
		toSerialize["maxAmount"] = o.MaxAmount
	}
	if o.MinAmount != nil {
		toSerialize["minAmount"] = o.MinAmount
	}
	if o.Precision != nil {
		toSerialize["precision"] = o.Precision
	}
	if o.StartAmount != nil {
		toSerialize["startAmount"] = o.StartAmount
	}
	if o.ValidDuration != nil {
		toSerialize["validDuration"] = o.ValidDuration
	}
	return json.Marshal(toSerialize)
}

type NullableV1Amount struct {
	value *V1Amount
	isSet bool
}

func (v NullableV1Amount) Get() *V1Amount {
	return v.value
}

func (v *NullableV1Amount) Set(val *V1Amount) {
	v.value = val
	v.isSet = true
}

func (v NullableV1Amount) IsSet() bool {
	return v.isSet
}

func (v *NullableV1Amount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1Amount(val *V1Amount) *NullableV1Amount {
	return &NullableV1Amount{value: val, isSet: true}
}

func (v NullableV1Amount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1Amount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

