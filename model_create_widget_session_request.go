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

// checks if the CreateWidgetSessionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateWidgetSessionRequest{}

// CreateWidgetSessionRequest struct for CreateWidgetSessionRequest
type CreateWidgetSessionRequest struct {
	// The URL to redirect the user to after the widget session is complete.                *Note*: this should NOT be set if you intend to use Trinsic's Web UI SDK to launch the Widget  as an embedded iFrame or popup; in that case, session resolution is handled by our SDK, not via redirect.
	RedirectUrl NullableString `json:"redirectUrl,omitempty"`
	// The list of allowed identity providers. If not specified, all available providers will be allowed.
	Providers []string `json:"providers,omitempty"`
	// Known identity data of an individual being verified.                Provide this to Trinsic during Session creation to enable improved identity provider selection recommendations.
	KnownIdentityData NullableKnownIdentityData `json:"knownIdentityData,omitempty"`
}

// NewCreateWidgetSessionRequest instantiates a new CreateWidgetSessionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateWidgetSessionRequest() *CreateWidgetSessionRequest {
	this := CreateWidgetSessionRequest{}
	return &this
}

// NewCreateWidgetSessionRequestWithDefaults instantiates a new CreateWidgetSessionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateWidgetSessionRequestWithDefaults() *CreateWidgetSessionRequest {
	this := CreateWidgetSessionRequest{}
	return &this
}

// GetRedirectUrl returns the RedirectUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateWidgetSessionRequest) GetRedirectUrl() string {
	if o == nil || IsNil(o.RedirectUrl.Get()) {
		var ret string
		return ret
	}
	return *o.RedirectUrl.Get()
}

// GetRedirectUrlOk returns a tuple with the RedirectUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateWidgetSessionRequest) GetRedirectUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RedirectUrl.Get(), o.RedirectUrl.IsSet()
}

// HasRedirectUrl returns a boolean if a field has been set.
func (o *CreateWidgetSessionRequest) HasRedirectUrl() bool {
	if o != nil && o.RedirectUrl.IsSet() {
		return true
	}

	return false
}

// SetRedirectUrl gets a reference to the given NullableString and assigns it to the RedirectUrl field.
func (o *CreateWidgetSessionRequest) SetRedirectUrl(v string) {
	o.RedirectUrl.Set(&v)
}
// SetRedirectUrlNil sets the value for RedirectUrl to be an explicit nil
func (o *CreateWidgetSessionRequest) SetRedirectUrlNil() {
	o.RedirectUrl.Set(nil)
}

// UnsetRedirectUrl ensures that no value is present for RedirectUrl, not even an explicit nil
func (o *CreateWidgetSessionRequest) UnsetRedirectUrl() {
	o.RedirectUrl.Unset()
}

// GetProviders returns the Providers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateWidgetSessionRequest) GetProviders() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Providers
}

// GetProvidersOk returns a tuple with the Providers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateWidgetSessionRequest) GetProvidersOk() ([]string, bool) {
	if o == nil || IsNil(o.Providers) {
		return nil, false
	}
	return o.Providers, true
}

// HasProviders returns a boolean if a field has been set.
func (o *CreateWidgetSessionRequest) HasProviders() bool {
	if o != nil && !IsNil(o.Providers) {
		return true
	}

	return false
}

// SetProviders gets a reference to the given []string and assigns it to the Providers field.
func (o *CreateWidgetSessionRequest) SetProviders(v []string) {
	o.Providers = v
}

// GetKnownIdentityData returns the KnownIdentityData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateWidgetSessionRequest) GetKnownIdentityData() KnownIdentityData {
	if o == nil || IsNil(o.KnownIdentityData.Get()) {
		var ret KnownIdentityData
		return ret
	}
	return *o.KnownIdentityData.Get()
}

// GetKnownIdentityDataOk returns a tuple with the KnownIdentityData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateWidgetSessionRequest) GetKnownIdentityDataOk() (*KnownIdentityData, bool) {
	if o == nil {
		return nil, false
	}
	return o.KnownIdentityData.Get(), o.KnownIdentityData.IsSet()
}

// HasKnownIdentityData returns a boolean if a field has been set.
func (o *CreateWidgetSessionRequest) HasKnownIdentityData() bool {
	if o != nil && o.KnownIdentityData.IsSet() {
		return true
	}

	return false
}

// SetKnownIdentityData gets a reference to the given NullableKnownIdentityData and assigns it to the KnownIdentityData field.
func (o *CreateWidgetSessionRequest) SetKnownIdentityData(v KnownIdentityData) {
	o.KnownIdentityData.Set(&v)
}
// SetKnownIdentityDataNil sets the value for KnownIdentityData to be an explicit nil
func (o *CreateWidgetSessionRequest) SetKnownIdentityDataNil() {
	o.KnownIdentityData.Set(nil)
}

// UnsetKnownIdentityData ensures that no value is present for KnownIdentityData, not even an explicit nil
func (o *CreateWidgetSessionRequest) UnsetKnownIdentityData() {
	o.KnownIdentityData.Unset()
}

func (o CreateWidgetSessionRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateWidgetSessionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.RedirectUrl.IsSet() {
		toSerialize["redirectUrl"] = o.RedirectUrl.Get()
	}
	if o.Providers != nil {
		toSerialize["providers"] = o.Providers
	}
	if o.KnownIdentityData.IsSet() {
		toSerialize["knownIdentityData"] = o.KnownIdentityData.Get()
	}
	return toSerialize, nil
}

type NullableCreateWidgetSessionRequest struct {
	value *CreateWidgetSessionRequest
	isSet bool
}

func (v NullableCreateWidgetSessionRequest) Get() *CreateWidgetSessionRequest {
	return v.value
}

func (v *NullableCreateWidgetSessionRequest) Set(val *CreateWidgetSessionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateWidgetSessionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateWidgetSessionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateWidgetSessionRequest(val *CreateWidgetSessionRequest) *NullableCreateWidgetSessionRequest {
	return &NullableCreateWidgetSessionRequest{value: val, isSet: true}
}

func (v NullableCreateWidgetSessionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateWidgetSessionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


