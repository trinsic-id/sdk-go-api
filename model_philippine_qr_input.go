/*
Trinsic API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package trinsic_api

import (
	"encoding/json"
)

// checks if the PhilippineQRInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PhilippineQRInput{}

// PhilippineQRInput struct for PhilippineQRInput
type PhilippineQRInput struct {
	// The raw text of the user's QR code after decoding it.
	QrCodeText NullableString `json:"qrCodeText,omitempty"`
	// The raw bytes of the image containing the user's QR code.
	QrCodeImage NullableString `json:"qrCodeImage,omitempty"`
}

// NewPhilippineQRInput instantiates a new PhilippineQRInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPhilippineQRInput() *PhilippineQRInput {
	this := PhilippineQRInput{}
	return &this
}

// NewPhilippineQRInputWithDefaults instantiates a new PhilippineQRInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPhilippineQRInputWithDefaults() *PhilippineQRInput {
	this := PhilippineQRInput{}
	return &this
}

// GetQrCodeText returns the QrCodeText field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PhilippineQRInput) GetQrCodeText() string {
	if o == nil || IsNil(o.QrCodeText.Get()) {
		var ret string
		return ret
	}
	return *o.QrCodeText.Get()
}

// GetQrCodeTextOk returns a tuple with the QrCodeText field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PhilippineQRInput) GetQrCodeTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.QrCodeText.Get(), o.QrCodeText.IsSet()
}

// HasQrCodeText returns a boolean if a field has been set.
func (o *PhilippineQRInput) HasQrCodeText() bool {
	if o != nil && o.QrCodeText.IsSet() {
		return true
	}

	return false
}

// SetQrCodeText gets a reference to the given NullableString and assigns it to the QrCodeText field.
func (o *PhilippineQRInput) SetQrCodeText(v string) {
	o.QrCodeText.Set(&v)
}
// SetQrCodeTextNil sets the value for QrCodeText to be an explicit nil
func (o *PhilippineQRInput) SetQrCodeTextNil() {
	o.QrCodeText.Set(nil)
}

// UnsetQrCodeText ensures that no value is present for QrCodeText, not even an explicit nil
func (o *PhilippineQRInput) UnsetQrCodeText() {
	o.QrCodeText.Unset()
}

// GetQrCodeImage returns the QrCodeImage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PhilippineQRInput) GetQrCodeImage() string {
	if o == nil || IsNil(o.QrCodeImage.Get()) {
		var ret string
		return ret
	}
	return *o.QrCodeImage.Get()
}

// GetQrCodeImageOk returns a tuple with the QrCodeImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PhilippineQRInput) GetQrCodeImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.QrCodeImage.Get(), o.QrCodeImage.IsSet()
}

// HasQrCodeImage returns a boolean if a field has been set.
func (o *PhilippineQRInput) HasQrCodeImage() bool {
	if o != nil && o.QrCodeImage.IsSet() {
		return true
	}

	return false
}

// SetQrCodeImage gets a reference to the given NullableString and assigns it to the QrCodeImage field.
func (o *PhilippineQRInput) SetQrCodeImage(v string) {
	o.QrCodeImage.Set(&v)
}
// SetQrCodeImageNil sets the value for QrCodeImage to be an explicit nil
func (o *PhilippineQRInput) SetQrCodeImageNil() {
	o.QrCodeImage.Set(nil)
}

// UnsetQrCodeImage ensures that no value is present for QrCodeImage, not even an explicit nil
func (o *PhilippineQRInput) UnsetQrCodeImage() {
	o.QrCodeImage.Unset()
}

func (o PhilippineQRInput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PhilippineQRInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.QrCodeText.IsSet() {
		toSerialize["qrCodeText"] = o.QrCodeText.Get()
	}
	if o.QrCodeImage.IsSet() {
		toSerialize["qrCodeImage"] = o.QrCodeImage.Get()
	}
	return toSerialize, nil
}

type NullablePhilippineQRInput struct {
	value *PhilippineQRInput
	isSet bool
}

func (v NullablePhilippineQRInput) Get() *PhilippineQRInput {
	return v.value
}

func (v *NullablePhilippineQRInput) Set(val *PhilippineQRInput) {
	v.value = val
	v.isSet = true
}

func (v NullablePhilippineQRInput) IsSet() bool {
	return v.isSet
}

func (v *NullablePhilippineQRInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePhilippineQRInput(val *PhilippineQRInput) *NullablePhilippineQRInput {
	return &NullablePhilippineQRInput{value: val, isSet: true}
}

func (v NullablePhilippineQRInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePhilippineQRInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


