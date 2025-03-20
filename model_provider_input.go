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

// checks if the ProviderInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProviderInput{}

// ProviderInput struct for ProviderInput
type ProviderInput struct {
	// Input for the `f-indonesia-nik` integration
	IndonesiaNik NullableIndonesiaNikInput `json:"indonesiaNik,omitempty"`
	// Input for the `f-mexico-curp` integration
	MexicoCurp NullableMexicoCurpInput `json:"mexicoCurp,omitempty"`
	// Input for the `f-south-africa-nid` integration
	SouthAfricaNid NullableSouthAfricaNidInput `json:"southAfricaNid,omitempty"`
	// Input for the `f-kenya-nid` integration
	KenyaNid NullableKenyaNidInput `json:"kenyaNid,omitempty"`
	// Input for the `f-nigeria-nin` integration
	NigeriaNin NullableNigeriaNinInput `json:"nigeriaNin,omitempty"`
	// Input for the `f-india-aadhaar-match` integration
	Aadhaar NullableAadhaarInput `json:"aadhaar,omitempty"`
	// Input for the `bangladesh-nid` integration
	BangladeshNationalId NullableBangladeshNationalIdInput `json:"bangladeshNationalId,omitempty"`
	// Input for the `g-brazil-cpf` integration
	BrazilCpfCheck NullableBrazilCpfCheckInput `json:"brazilCpfCheck,omitempty"`
	// Input for the `g-brazil-digital-cnh` integration
	BrazilDigitalCnh NullableBrazilDigitalCnhInput `json:"brazilDigitalCnh,omitempty"`
	// Input for the `b-philsys-biometric` integration
	PhilippineMatch NullablePhilippineMatchInput `json:"philippineMatch,omitempty"`
	// Input for the `b-philippine-qr-digital-national-id` and `b-philippine-qr-ephill-id` integrations
	PhilippineQR NullablePhilippineQRInput `json:"philippineQR,omitempty"`
}

// NewProviderInput instantiates a new ProviderInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProviderInput() *ProviderInput {
	this := ProviderInput{}
	return &this
}

// NewProviderInputWithDefaults instantiates a new ProviderInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProviderInputWithDefaults() *ProviderInput {
	this := ProviderInput{}
	return &this
}

// GetIndonesiaNik returns the IndonesiaNik field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProviderInput) GetIndonesiaNik() IndonesiaNikInput {
	if o == nil || IsNil(o.IndonesiaNik.Get()) {
		var ret IndonesiaNikInput
		return ret
	}
	return *o.IndonesiaNik.Get()
}

// GetIndonesiaNikOk returns a tuple with the IndonesiaNik field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProviderInput) GetIndonesiaNikOk() (*IndonesiaNikInput, bool) {
	if o == nil {
		return nil, false
	}
	return o.IndonesiaNik.Get(), o.IndonesiaNik.IsSet()
}

// HasIndonesiaNik returns a boolean if a field has been set.
func (o *ProviderInput) HasIndonesiaNik() bool {
	if o != nil && o.IndonesiaNik.IsSet() {
		return true
	}

	return false
}

// SetIndonesiaNik gets a reference to the given NullableIndonesiaNikInput and assigns it to the IndonesiaNik field.
func (o *ProviderInput) SetIndonesiaNik(v IndonesiaNikInput) {
	o.IndonesiaNik.Set(&v)
}
// SetIndonesiaNikNil sets the value for IndonesiaNik to be an explicit nil
func (o *ProviderInput) SetIndonesiaNikNil() {
	o.IndonesiaNik.Set(nil)
}

// UnsetIndonesiaNik ensures that no value is present for IndonesiaNik, not even an explicit nil
func (o *ProviderInput) UnsetIndonesiaNik() {
	o.IndonesiaNik.Unset()
}

// GetMexicoCurp returns the MexicoCurp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProviderInput) GetMexicoCurp() MexicoCurpInput {
	if o == nil || IsNil(o.MexicoCurp.Get()) {
		var ret MexicoCurpInput
		return ret
	}
	return *o.MexicoCurp.Get()
}

// GetMexicoCurpOk returns a tuple with the MexicoCurp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProviderInput) GetMexicoCurpOk() (*MexicoCurpInput, bool) {
	if o == nil {
		return nil, false
	}
	return o.MexicoCurp.Get(), o.MexicoCurp.IsSet()
}

// HasMexicoCurp returns a boolean if a field has been set.
func (o *ProviderInput) HasMexicoCurp() bool {
	if o != nil && o.MexicoCurp.IsSet() {
		return true
	}

	return false
}

// SetMexicoCurp gets a reference to the given NullableMexicoCurpInput and assigns it to the MexicoCurp field.
func (o *ProviderInput) SetMexicoCurp(v MexicoCurpInput) {
	o.MexicoCurp.Set(&v)
}
// SetMexicoCurpNil sets the value for MexicoCurp to be an explicit nil
func (o *ProviderInput) SetMexicoCurpNil() {
	o.MexicoCurp.Set(nil)
}

// UnsetMexicoCurp ensures that no value is present for MexicoCurp, not even an explicit nil
func (o *ProviderInput) UnsetMexicoCurp() {
	o.MexicoCurp.Unset()
}

// GetSouthAfricaNid returns the SouthAfricaNid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProviderInput) GetSouthAfricaNid() SouthAfricaNidInput {
	if o == nil || IsNil(o.SouthAfricaNid.Get()) {
		var ret SouthAfricaNidInput
		return ret
	}
	return *o.SouthAfricaNid.Get()
}

// GetSouthAfricaNidOk returns a tuple with the SouthAfricaNid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProviderInput) GetSouthAfricaNidOk() (*SouthAfricaNidInput, bool) {
	if o == nil {
		return nil, false
	}
	return o.SouthAfricaNid.Get(), o.SouthAfricaNid.IsSet()
}

// HasSouthAfricaNid returns a boolean if a field has been set.
func (o *ProviderInput) HasSouthAfricaNid() bool {
	if o != nil && o.SouthAfricaNid.IsSet() {
		return true
	}

	return false
}

// SetSouthAfricaNid gets a reference to the given NullableSouthAfricaNidInput and assigns it to the SouthAfricaNid field.
func (o *ProviderInput) SetSouthAfricaNid(v SouthAfricaNidInput) {
	o.SouthAfricaNid.Set(&v)
}
// SetSouthAfricaNidNil sets the value for SouthAfricaNid to be an explicit nil
func (o *ProviderInput) SetSouthAfricaNidNil() {
	o.SouthAfricaNid.Set(nil)
}

// UnsetSouthAfricaNid ensures that no value is present for SouthAfricaNid, not even an explicit nil
func (o *ProviderInput) UnsetSouthAfricaNid() {
	o.SouthAfricaNid.Unset()
}

// GetKenyaNid returns the KenyaNid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProviderInput) GetKenyaNid() KenyaNidInput {
	if o == nil || IsNil(o.KenyaNid.Get()) {
		var ret KenyaNidInput
		return ret
	}
	return *o.KenyaNid.Get()
}

// GetKenyaNidOk returns a tuple with the KenyaNid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProviderInput) GetKenyaNidOk() (*KenyaNidInput, bool) {
	if o == nil {
		return nil, false
	}
	return o.KenyaNid.Get(), o.KenyaNid.IsSet()
}

// HasKenyaNid returns a boolean if a field has been set.
func (o *ProviderInput) HasKenyaNid() bool {
	if o != nil && o.KenyaNid.IsSet() {
		return true
	}

	return false
}

// SetKenyaNid gets a reference to the given NullableKenyaNidInput and assigns it to the KenyaNid field.
func (o *ProviderInput) SetKenyaNid(v KenyaNidInput) {
	o.KenyaNid.Set(&v)
}
// SetKenyaNidNil sets the value for KenyaNid to be an explicit nil
func (o *ProviderInput) SetKenyaNidNil() {
	o.KenyaNid.Set(nil)
}

// UnsetKenyaNid ensures that no value is present for KenyaNid, not even an explicit nil
func (o *ProviderInput) UnsetKenyaNid() {
	o.KenyaNid.Unset()
}

// GetNigeriaNin returns the NigeriaNin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProviderInput) GetNigeriaNin() NigeriaNinInput {
	if o == nil || IsNil(o.NigeriaNin.Get()) {
		var ret NigeriaNinInput
		return ret
	}
	return *o.NigeriaNin.Get()
}

// GetNigeriaNinOk returns a tuple with the NigeriaNin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProviderInput) GetNigeriaNinOk() (*NigeriaNinInput, bool) {
	if o == nil {
		return nil, false
	}
	return o.NigeriaNin.Get(), o.NigeriaNin.IsSet()
}

// HasNigeriaNin returns a boolean if a field has been set.
func (o *ProviderInput) HasNigeriaNin() bool {
	if o != nil && o.NigeriaNin.IsSet() {
		return true
	}

	return false
}

// SetNigeriaNin gets a reference to the given NullableNigeriaNinInput and assigns it to the NigeriaNin field.
func (o *ProviderInput) SetNigeriaNin(v NigeriaNinInput) {
	o.NigeriaNin.Set(&v)
}
// SetNigeriaNinNil sets the value for NigeriaNin to be an explicit nil
func (o *ProviderInput) SetNigeriaNinNil() {
	o.NigeriaNin.Set(nil)
}

// UnsetNigeriaNin ensures that no value is present for NigeriaNin, not even an explicit nil
func (o *ProviderInput) UnsetNigeriaNin() {
	o.NigeriaNin.Unset()
}

// GetAadhaar returns the Aadhaar field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProviderInput) GetAadhaar() AadhaarInput {
	if o == nil || IsNil(o.Aadhaar.Get()) {
		var ret AadhaarInput
		return ret
	}
	return *o.Aadhaar.Get()
}

// GetAadhaarOk returns a tuple with the Aadhaar field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProviderInput) GetAadhaarOk() (*AadhaarInput, bool) {
	if o == nil {
		return nil, false
	}
	return o.Aadhaar.Get(), o.Aadhaar.IsSet()
}

// HasAadhaar returns a boolean if a field has been set.
func (o *ProviderInput) HasAadhaar() bool {
	if o != nil && o.Aadhaar.IsSet() {
		return true
	}

	return false
}

// SetAadhaar gets a reference to the given NullableAadhaarInput and assigns it to the Aadhaar field.
func (o *ProviderInput) SetAadhaar(v AadhaarInput) {
	o.Aadhaar.Set(&v)
}
// SetAadhaarNil sets the value for Aadhaar to be an explicit nil
func (o *ProviderInput) SetAadhaarNil() {
	o.Aadhaar.Set(nil)
}

// UnsetAadhaar ensures that no value is present for Aadhaar, not even an explicit nil
func (o *ProviderInput) UnsetAadhaar() {
	o.Aadhaar.Unset()
}

// GetBangladeshNationalId returns the BangladeshNationalId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProviderInput) GetBangladeshNationalId() BangladeshNationalIdInput {
	if o == nil || IsNil(o.BangladeshNationalId.Get()) {
		var ret BangladeshNationalIdInput
		return ret
	}
	return *o.BangladeshNationalId.Get()
}

// GetBangladeshNationalIdOk returns a tuple with the BangladeshNationalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProviderInput) GetBangladeshNationalIdOk() (*BangladeshNationalIdInput, bool) {
	if o == nil {
		return nil, false
	}
	return o.BangladeshNationalId.Get(), o.BangladeshNationalId.IsSet()
}

// HasBangladeshNationalId returns a boolean if a field has been set.
func (o *ProviderInput) HasBangladeshNationalId() bool {
	if o != nil && o.BangladeshNationalId.IsSet() {
		return true
	}

	return false
}

// SetBangladeshNationalId gets a reference to the given NullableBangladeshNationalIdInput and assigns it to the BangladeshNationalId field.
func (o *ProviderInput) SetBangladeshNationalId(v BangladeshNationalIdInput) {
	o.BangladeshNationalId.Set(&v)
}
// SetBangladeshNationalIdNil sets the value for BangladeshNationalId to be an explicit nil
func (o *ProviderInput) SetBangladeshNationalIdNil() {
	o.BangladeshNationalId.Set(nil)
}

// UnsetBangladeshNationalId ensures that no value is present for BangladeshNationalId, not even an explicit nil
func (o *ProviderInput) UnsetBangladeshNationalId() {
	o.BangladeshNationalId.Unset()
}

// GetBrazilCpfCheck returns the BrazilCpfCheck field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProviderInput) GetBrazilCpfCheck() BrazilCpfCheckInput {
	if o == nil || IsNil(o.BrazilCpfCheck.Get()) {
		var ret BrazilCpfCheckInput
		return ret
	}
	return *o.BrazilCpfCheck.Get()
}

// GetBrazilCpfCheckOk returns a tuple with the BrazilCpfCheck field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProviderInput) GetBrazilCpfCheckOk() (*BrazilCpfCheckInput, bool) {
	if o == nil {
		return nil, false
	}
	return o.BrazilCpfCheck.Get(), o.BrazilCpfCheck.IsSet()
}

// HasBrazilCpfCheck returns a boolean if a field has been set.
func (o *ProviderInput) HasBrazilCpfCheck() bool {
	if o != nil && o.BrazilCpfCheck.IsSet() {
		return true
	}

	return false
}

// SetBrazilCpfCheck gets a reference to the given NullableBrazilCpfCheckInput and assigns it to the BrazilCpfCheck field.
func (o *ProviderInput) SetBrazilCpfCheck(v BrazilCpfCheckInput) {
	o.BrazilCpfCheck.Set(&v)
}
// SetBrazilCpfCheckNil sets the value for BrazilCpfCheck to be an explicit nil
func (o *ProviderInput) SetBrazilCpfCheckNil() {
	o.BrazilCpfCheck.Set(nil)
}

// UnsetBrazilCpfCheck ensures that no value is present for BrazilCpfCheck, not even an explicit nil
func (o *ProviderInput) UnsetBrazilCpfCheck() {
	o.BrazilCpfCheck.Unset()
}

// GetBrazilDigitalCnh returns the BrazilDigitalCnh field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProviderInput) GetBrazilDigitalCnh() BrazilDigitalCnhInput {
	if o == nil || IsNil(o.BrazilDigitalCnh.Get()) {
		var ret BrazilDigitalCnhInput
		return ret
	}
	return *o.BrazilDigitalCnh.Get()
}

// GetBrazilDigitalCnhOk returns a tuple with the BrazilDigitalCnh field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProviderInput) GetBrazilDigitalCnhOk() (*BrazilDigitalCnhInput, bool) {
	if o == nil {
		return nil, false
	}
	return o.BrazilDigitalCnh.Get(), o.BrazilDigitalCnh.IsSet()
}

// HasBrazilDigitalCnh returns a boolean if a field has been set.
func (o *ProviderInput) HasBrazilDigitalCnh() bool {
	if o != nil && o.BrazilDigitalCnh.IsSet() {
		return true
	}

	return false
}

// SetBrazilDigitalCnh gets a reference to the given NullableBrazilDigitalCnhInput and assigns it to the BrazilDigitalCnh field.
func (o *ProviderInput) SetBrazilDigitalCnh(v BrazilDigitalCnhInput) {
	o.BrazilDigitalCnh.Set(&v)
}
// SetBrazilDigitalCnhNil sets the value for BrazilDigitalCnh to be an explicit nil
func (o *ProviderInput) SetBrazilDigitalCnhNil() {
	o.BrazilDigitalCnh.Set(nil)
}

// UnsetBrazilDigitalCnh ensures that no value is present for BrazilDigitalCnh, not even an explicit nil
func (o *ProviderInput) UnsetBrazilDigitalCnh() {
	o.BrazilDigitalCnh.Unset()
}

// GetPhilippineMatch returns the PhilippineMatch field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProviderInput) GetPhilippineMatch() PhilippineMatchInput {
	if o == nil || IsNil(o.PhilippineMatch.Get()) {
		var ret PhilippineMatchInput
		return ret
	}
	return *o.PhilippineMatch.Get()
}

// GetPhilippineMatchOk returns a tuple with the PhilippineMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProviderInput) GetPhilippineMatchOk() (*PhilippineMatchInput, bool) {
	if o == nil {
		return nil, false
	}
	return o.PhilippineMatch.Get(), o.PhilippineMatch.IsSet()
}

// HasPhilippineMatch returns a boolean if a field has been set.
func (o *ProviderInput) HasPhilippineMatch() bool {
	if o != nil && o.PhilippineMatch.IsSet() {
		return true
	}

	return false
}

// SetPhilippineMatch gets a reference to the given NullablePhilippineMatchInput and assigns it to the PhilippineMatch field.
func (o *ProviderInput) SetPhilippineMatch(v PhilippineMatchInput) {
	o.PhilippineMatch.Set(&v)
}
// SetPhilippineMatchNil sets the value for PhilippineMatch to be an explicit nil
func (o *ProviderInput) SetPhilippineMatchNil() {
	o.PhilippineMatch.Set(nil)
}

// UnsetPhilippineMatch ensures that no value is present for PhilippineMatch, not even an explicit nil
func (o *ProviderInput) UnsetPhilippineMatch() {
	o.PhilippineMatch.Unset()
}

// GetPhilippineQR returns the PhilippineQR field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProviderInput) GetPhilippineQR() PhilippineQRInput {
	if o == nil || IsNil(o.PhilippineQR.Get()) {
		var ret PhilippineQRInput
		return ret
	}
	return *o.PhilippineQR.Get()
}

// GetPhilippineQROk returns a tuple with the PhilippineQR field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProviderInput) GetPhilippineQROk() (*PhilippineQRInput, bool) {
	if o == nil {
		return nil, false
	}
	return o.PhilippineQR.Get(), o.PhilippineQR.IsSet()
}

// HasPhilippineQR returns a boolean if a field has been set.
func (o *ProviderInput) HasPhilippineQR() bool {
	if o != nil && o.PhilippineQR.IsSet() {
		return true
	}

	return false
}

// SetPhilippineQR gets a reference to the given NullablePhilippineQRInput and assigns it to the PhilippineQR field.
func (o *ProviderInput) SetPhilippineQR(v PhilippineQRInput) {
	o.PhilippineQR.Set(&v)
}
// SetPhilippineQRNil sets the value for PhilippineQR to be an explicit nil
func (o *ProviderInput) SetPhilippineQRNil() {
	o.PhilippineQR.Set(nil)
}

// UnsetPhilippineQR ensures that no value is present for PhilippineQR, not even an explicit nil
func (o *ProviderInput) UnsetPhilippineQR() {
	o.PhilippineQR.Unset()
}

func (o ProviderInput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProviderInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.IndonesiaNik.IsSet() {
		toSerialize["indonesiaNik"] = o.IndonesiaNik.Get()
	}
	if o.MexicoCurp.IsSet() {
		toSerialize["mexicoCurp"] = o.MexicoCurp.Get()
	}
	if o.SouthAfricaNid.IsSet() {
		toSerialize["southAfricaNid"] = o.SouthAfricaNid.Get()
	}
	if o.KenyaNid.IsSet() {
		toSerialize["kenyaNid"] = o.KenyaNid.Get()
	}
	if o.NigeriaNin.IsSet() {
		toSerialize["nigeriaNin"] = o.NigeriaNin.Get()
	}
	if o.Aadhaar.IsSet() {
		toSerialize["aadhaar"] = o.Aadhaar.Get()
	}
	if o.BangladeshNationalId.IsSet() {
		toSerialize["bangladeshNationalId"] = o.BangladeshNationalId.Get()
	}
	if o.BrazilCpfCheck.IsSet() {
		toSerialize["brazilCpfCheck"] = o.BrazilCpfCheck.Get()
	}
	if o.BrazilDigitalCnh.IsSet() {
		toSerialize["brazilDigitalCnh"] = o.BrazilDigitalCnh.Get()
	}
	if o.PhilippineMatch.IsSet() {
		toSerialize["philippineMatch"] = o.PhilippineMatch.Get()
	}
	if o.PhilippineQR.IsSet() {
		toSerialize["philippineQR"] = o.PhilippineQR.Get()
	}
	return toSerialize, nil
}

type NullableProviderInput struct {
	value *ProviderInput
	isSet bool
}

func (v NullableProviderInput) Get() *ProviderInput {
	return v.value
}

func (v *NullableProviderInput) Set(val *ProviderInput) {
	v.value = val
	v.isSet = true
}

func (v NullableProviderInput) IsSet() bool {
	return v.isSet
}

func (v *NullableProviderInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProviderInput(val *ProviderInput) *NullableProviderInput {
	return &NullableProviderInput{value: val, isSet: true}
}

func (v NullableProviderInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProviderInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


