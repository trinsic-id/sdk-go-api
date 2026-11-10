# RevolutProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subject** | Pointer to **NullableString** | The Revolut subject identifier. | [optional] 
**Name** | Pointer to **NullableString** | The individual&#39;s full name. | [optional] 
**Email** | Pointer to **NullableString** | The email address returned by Revolut. | [optional] 
**EmailVerified** | Pointer to **NullableBool** | Whether Revolut verified the email address. | [optional] 
**Phone** | Pointer to **NullableString** | The phone number in E.164 format. | [optional] 
**PhoneVerified** | Pointer to **NullableBool** | Whether Revolut verified the phone number. | [optional] 

## Methods

### Named construction

Use a keyed struct literal (`RevolutProviderOutput{FieldName: value}`), or call
`NewRevolutProviderOutputWithDefaults()` and set fields by name. Positional field
factories and unkeyed struct literals are not supported.

### NewRevolutProviderOutputWithDefaults

`func NewRevolutProviderOutputWithDefaults() *RevolutProviderOutput`

NewRevolutProviderOutputWithDefaults instantiates a new RevolutProviderOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubject

`func (o *RevolutProviderOutput) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *RevolutProviderOutput) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *RevolutProviderOutput) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *RevolutProviderOutput) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### SetSubjectNil

`func (o *RevolutProviderOutput) SetSubjectNil(b bool)`

 SetSubjectNil sets the value for Subject to be an explicit nil

### UnsetSubject
`func (o *RevolutProviderOutput) UnsetSubject()`

UnsetSubject ensures that no value is present for Subject, not even an explicit nil
### GetName

`func (o *RevolutProviderOutput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RevolutProviderOutput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RevolutProviderOutput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RevolutProviderOutput) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *RevolutProviderOutput) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *RevolutProviderOutput) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetEmail

`func (o *RevolutProviderOutput) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *RevolutProviderOutput) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *RevolutProviderOutput) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *RevolutProviderOutput) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *RevolutProviderOutput) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *RevolutProviderOutput) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetEmailVerified

`func (o *RevolutProviderOutput) GetEmailVerified() bool`

GetEmailVerified returns the EmailVerified field if non-nil, zero value otherwise.

### GetEmailVerifiedOk

`func (o *RevolutProviderOutput) GetEmailVerifiedOk() (*bool, bool)`

GetEmailVerifiedOk returns a tuple with the EmailVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailVerified

`func (o *RevolutProviderOutput) SetEmailVerified(v bool)`

SetEmailVerified sets EmailVerified field to given value.

### HasEmailVerified

`func (o *RevolutProviderOutput) HasEmailVerified() bool`

HasEmailVerified returns a boolean if a field has been set.

### SetEmailVerifiedNil

`func (o *RevolutProviderOutput) SetEmailVerifiedNil(b bool)`

 SetEmailVerifiedNil sets the value for EmailVerified to be an explicit nil

### UnsetEmailVerified
`func (o *RevolutProviderOutput) UnsetEmailVerified()`

UnsetEmailVerified ensures that no value is present for EmailVerified, not even an explicit nil
### GetPhone

`func (o *RevolutProviderOutput) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *RevolutProviderOutput) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *RevolutProviderOutput) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *RevolutProviderOutput) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### SetPhoneNil

`func (o *RevolutProviderOutput) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *RevolutProviderOutput) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil
### GetPhoneVerified

`func (o *RevolutProviderOutput) GetPhoneVerified() bool`

GetPhoneVerified returns the PhoneVerified field if non-nil, zero value otherwise.

### GetPhoneVerifiedOk

`func (o *RevolutProviderOutput) GetPhoneVerifiedOk() (*bool, bool)`

GetPhoneVerifiedOk returns a tuple with the PhoneVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneVerified

`func (o *RevolutProviderOutput) SetPhoneVerified(v bool)`

SetPhoneVerified sets PhoneVerified field to given value.

### HasPhoneVerified

`func (o *RevolutProviderOutput) HasPhoneVerified() bool`

HasPhoneVerified returns a boolean if a field has been set.

### SetPhoneVerifiedNil

`func (o *RevolutProviderOutput) SetPhoneVerifiedNil(b bool)`

 SetPhoneVerifiedNil sets the value for PhoneVerified to be an explicit nil

### UnsetPhoneVerified
`func (o *RevolutProviderOutput) UnsetPhoneVerified()`

UnsetPhoneVerified ensures that no value is present for PhoneVerified, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


