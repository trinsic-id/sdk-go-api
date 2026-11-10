# JapanMyNumberProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | The individual&#39;s full name as recorded on the card, preserving its original script and spacing. | [optional] 
**Address** | Pointer to **NullableString** | The individual&#39;s full address as recorded on the card, preserving its original script and spacing. | [optional] 
**DateOfBirth** | Pointer to **NullableString** | The individual&#39;s Gregorian date of birth in YYYY-MM-DD format. | [optional] 
**Sex** | Pointer to **NullableString** | The original sex code returned for the verified individual.              Known values: - 1: Male. - 2: Female. - 9: Other; omitted from normalized sex because it has no equivalent value. An absent code means it was not supplied. | [optional] 

## Methods

### Named construction

Use a keyed struct literal (`JapanMyNumberProviderOutput{FieldName: value}`), or call
`NewJapanMyNumberProviderOutputWithDefaults()` and set fields by name. Positional field
factories and unkeyed struct literals are not supported.

### NewJapanMyNumberProviderOutputWithDefaults

`func NewJapanMyNumberProviderOutputWithDefaults() *JapanMyNumberProviderOutput`

NewJapanMyNumberProviderOutputWithDefaults instantiates a new JapanMyNumberProviderOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *JapanMyNumberProviderOutput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *JapanMyNumberProviderOutput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *JapanMyNumberProviderOutput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *JapanMyNumberProviderOutput) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *JapanMyNumberProviderOutput) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *JapanMyNumberProviderOutput) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetAddress

`func (o *JapanMyNumberProviderOutput) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *JapanMyNumberProviderOutput) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *JapanMyNumberProviderOutput) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *JapanMyNumberProviderOutput) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *JapanMyNumberProviderOutput) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *JapanMyNumberProviderOutput) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetDateOfBirth

`func (o *JapanMyNumberProviderOutput) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *JapanMyNumberProviderOutput) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *JapanMyNumberProviderOutput) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *JapanMyNumberProviderOutput) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### SetDateOfBirthNil

`func (o *JapanMyNumberProviderOutput) SetDateOfBirthNil(b bool)`

 SetDateOfBirthNil sets the value for DateOfBirth to be an explicit nil

### UnsetDateOfBirth
`func (o *JapanMyNumberProviderOutput) UnsetDateOfBirth()`

UnsetDateOfBirth ensures that no value is present for DateOfBirth, not even an explicit nil
### GetSex

`func (o *JapanMyNumberProviderOutput) GetSex() string`

GetSex returns the Sex field if non-nil, zero value otherwise.

### GetSexOk

`func (o *JapanMyNumberProviderOutput) GetSexOk() (*string, bool)`

GetSexOk returns a tuple with the Sex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSex

`func (o *JapanMyNumberProviderOutput) SetSex(v string)`

SetSex sets Sex field to given value.

### HasSex

`func (o *JapanMyNumberProviderOutput) HasSex() bool`

HasSex returns a boolean if a field has been set.

### SetSexNil

`func (o *JapanMyNumberProviderOutput) SetSexNil(b bool)`

 SetSexNil sets the value for Sex to be an explicit nil

### UnsetSex
`func (o *JapanMyNumberProviderOutput) UnsetSex()`

UnsetSex ensures that no value is present for Sex, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


