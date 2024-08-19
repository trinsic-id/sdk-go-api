/*
Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package connect

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ExchangeAttachmentAccessKeyRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExchangeAttachmentAccessKeyRequest{}

// ExchangeAttachmentAccessKeyRequest struct for ExchangeAttachmentAccessKeyRequest
type ExchangeAttachmentAccessKeyRequest struct {
	// The Attachment Access Key to exchange for the raw file contents of the related Attachment
	AttachmentAccessKey string `json:"attachmentAccessKey"`
}

type _ExchangeAttachmentAccessKeyRequest ExchangeAttachmentAccessKeyRequest

// NewExchangeAttachmentAccessKeyRequest instantiates a new ExchangeAttachmentAccessKeyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExchangeAttachmentAccessKeyRequest(attachmentAccessKey string) *ExchangeAttachmentAccessKeyRequest {
	this := ExchangeAttachmentAccessKeyRequest{}
	this.AttachmentAccessKey = attachmentAccessKey
	return &this
}

// NewExchangeAttachmentAccessKeyRequestWithDefaults instantiates a new ExchangeAttachmentAccessKeyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExchangeAttachmentAccessKeyRequestWithDefaults() *ExchangeAttachmentAccessKeyRequest {
	this := ExchangeAttachmentAccessKeyRequest{}
	return &this
}

// GetAttachmentAccessKey returns the AttachmentAccessKey field value
func (o *ExchangeAttachmentAccessKeyRequest) GetAttachmentAccessKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AttachmentAccessKey
}

// GetAttachmentAccessKeyOk returns a tuple with the AttachmentAccessKey field value
// and a boolean to check if the value has been set.
func (o *ExchangeAttachmentAccessKeyRequest) GetAttachmentAccessKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AttachmentAccessKey, true
}

// SetAttachmentAccessKey sets field value
func (o *ExchangeAttachmentAccessKeyRequest) SetAttachmentAccessKey(v string) {
	o.AttachmentAccessKey = v
}

func (o ExchangeAttachmentAccessKeyRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExchangeAttachmentAccessKeyRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["attachmentAccessKey"] = o.AttachmentAccessKey
	return toSerialize, nil
}

func (o *ExchangeAttachmentAccessKeyRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"attachmentAccessKey",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varExchangeAttachmentAccessKeyRequest := _ExchangeAttachmentAccessKeyRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varExchangeAttachmentAccessKeyRequest)

	if err != nil {
		return err
	}

	*o = ExchangeAttachmentAccessKeyRequest(varExchangeAttachmentAccessKeyRequest)

	return err
}

type NullableExchangeAttachmentAccessKeyRequest struct {
	value *ExchangeAttachmentAccessKeyRequest
	isSet bool
}

func (v NullableExchangeAttachmentAccessKeyRequest) Get() *ExchangeAttachmentAccessKeyRequest {
	return v.value
}

func (v *NullableExchangeAttachmentAccessKeyRequest) Set(val *ExchangeAttachmentAccessKeyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableExchangeAttachmentAccessKeyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableExchangeAttachmentAccessKeyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExchangeAttachmentAccessKeyRequest(val *ExchangeAttachmentAccessKeyRequest) *NullableExchangeAttachmentAccessKeyRequest {
	return &NullableExchangeAttachmentAccessKeyRequest{value: val, isSet: true}
}

func (v NullableExchangeAttachmentAccessKeyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExchangeAttachmentAccessKeyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


