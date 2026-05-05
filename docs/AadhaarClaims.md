# AadhaarClaims

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | The full name. | [optional] 
**DateOfBirth** | Pointer to **NullableString** | The date of birth.              The format is YYYY-MM-DD. | [optional] 
**Gender** | Pointer to **NullableString** | The gender of the individual.              Possible values: - M (Male) - F (Female) - T (Transgender) | [optional] 
**Address** | Pointer to [**NullableAadhaarAddress**](AadhaarAddress.md) | The structured Indian address. | [optional] 

## Methods

### NewAadhaarClaims

`func NewAadhaarClaims() *AadhaarClaims`

NewAadhaarClaims instantiates a new AadhaarClaims object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAadhaarClaimsWithDefaults

`func NewAadhaarClaimsWithDefaults() *AadhaarClaims`

NewAadhaarClaimsWithDefaults instantiates a new AadhaarClaims object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AadhaarClaims) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AadhaarClaims) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AadhaarClaims) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AadhaarClaims) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AadhaarClaims) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AadhaarClaims) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDateOfBirth

`func (o *AadhaarClaims) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *AadhaarClaims) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *AadhaarClaims) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *AadhaarClaims) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### SetDateOfBirthNil

`func (o *AadhaarClaims) SetDateOfBirthNil(b bool)`

 SetDateOfBirthNil sets the value for DateOfBirth to be an explicit nil

### UnsetDateOfBirth
`func (o *AadhaarClaims) UnsetDateOfBirth()`

UnsetDateOfBirth ensures that no value is present for DateOfBirth, not even an explicit nil
### GetGender

`func (o *AadhaarClaims) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *AadhaarClaims) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *AadhaarClaims) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *AadhaarClaims) HasGender() bool`

HasGender returns a boolean if a field has been set.

### SetGenderNil

`func (o *AadhaarClaims) SetGenderNil(b bool)`

 SetGenderNil sets the value for Gender to be an explicit nil

### UnsetGender
`func (o *AadhaarClaims) UnsetGender()`

UnsetGender ensures that no value is present for Gender, not even an explicit nil
### GetAddress

`func (o *AadhaarClaims) GetAddress() AadhaarAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *AadhaarClaims) GetAddressOk() (*AadhaarAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *AadhaarClaims) SetAddress(v AadhaarAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *AadhaarClaims) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *AadhaarClaims) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *AadhaarClaims) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


