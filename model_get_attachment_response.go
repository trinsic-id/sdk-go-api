/*
Trinsic API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package trinsic_api

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the GetAttachmentResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAttachmentResponse{}

// GetAttachmentResponse struct for GetAttachmentResponse
type GetAttachmentResponse struct {
	// The raw file contents of the Attachment
	Content string `json:"content"`
	// The MIME type of the Attachment data
	ContentType string `json:"contentType"`
}

type _GetAttachmentResponse GetAttachmentResponse

// NewGetAttachmentResponse instantiates a new GetAttachmentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAttachmentResponse(content string, contentType string) *GetAttachmentResponse {
	this := GetAttachmentResponse{}
	this.Content = content
	this.ContentType = contentType
	return &this
}

// NewGetAttachmentResponseWithDefaults instantiates a new GetAttachmentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAttachmentResponseWithDefaults() *GetAttachmentResponse {
	this := GetAttachmentResponse{}
	return &this
}

// GetContent returns the Content field value
func (o *GetAttachmentResponse) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *GetAttachmentResponse) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *GetAttachmentResponse) SetContent(v string) {
	o.Content = v
}

// GetContentType returns the ContentType field value
func (o *GetAttachmentResponse) GetContentType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value
// and a boolean to check if the value has been set.
func (o *GetAttachmentResponse) GetContentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContentType, true
}

// SetContentType sets field value
func (o *GetAttachmentResponse) SetContentType(v string) {
	o.ContentType = v
}

func (o GetAttachmentResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAttachmentResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["content"] = o.Content
	toSerialize["contentType"] = o.ContentType
	return toSerialize, nil
}

func (o *GetAttachmentResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"content",
		"contentType",
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

	varGetAttachmentResponse := _GetAttachmentResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetAttachmentResponse)

	if err != nil {
		return err
	}

	*o = GetAttachmentResponse(varGetAttachmentResponse)

	return err
}

type NullableGetAttachmentResponse struct {
	value *GetAttachmentResponse
	isSet bool
}

func (v NullableGetAttachmentResponse) Get() *GetAttachmentResponse {
	return v.value
}

func (v *NullableGetAttachmentResponse) Set(val *GetAttachmentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAttachmentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAttachmentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAttachmentResponse(val *GetAttachmentResponse) *NullableGetAttachmentResponse {
	return &NullableGetAttachmentResponse{value: val, isSet: true}
}

func (v NullableGetAttachmentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAttachmentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


