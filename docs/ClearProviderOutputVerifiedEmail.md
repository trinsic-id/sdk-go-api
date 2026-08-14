# ClearProviderOutputVerifiedEmail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **NullableString** | The verified email address. | [optional] 
**VerifiedAt** | Pointer to **NullableTime** | The time the email was verified, as a UTC timestamp. | [optional] 
**VerificationMethod** | Pointer to **NullableString** | The method CLEAR used to verify the email. | [optional] 

## Methods

### NewClearProviderOutputVerifiedEmail

`func NewClearProviderOutputVerifiedEmail() *ClearProviderOutputVerifiedEmail`

NewClearProviderOutputVerifiedEmail instantiates a new ClearProviderOutputVerifiedEmail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClearProviderOutputVerifiedEmailWithDefaults

`func NewClearProviderOutputVerifiedEmailWithDefaults() *ClearProviderOutputVerifiedEmail`

NewClearProviderOutputVerifiedEmailWithDefaults instantiates a new ClearProviderOutputVerifiedEmail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *ClearProviderOutputVerifiedEmail) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ClearProviderOutputVerifiedEmail) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ClearProviderOutputVerifiedEmail) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ClearProviderOutputVerifiedEmail) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *ClearProviderOutputVerifiedEmail) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *ClearProviderOutputVerifiedEmail) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetVerifiedAt

`func (o *ClearProviderOutputVerifiedEmail) GetVerifiedAt() time.Time`

GetVerifiedAt returns the VerifiedAt field if non-nil, zero value otherwise.

### GetVerifiedAtOk

`func (o *ClearProviderOutputVerifiedEmail) GetVerifiedAtOk() (*time.Time, bool)`

GetVerifiedAtOk returns a tuple with the VerifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedAt

`func (o *ClearProviderOutputVerifiedEmail) SetVerifiedAt(v time.Time)`

SetVerifiedAt sets VerifiedAt field to given value.

### HasVerifiedAt

`func (o *ClearProviderOutputVerifiedEmail) HasVerifiedAt() bool`

HasVerifiedAt returns a boolean if a field has been set.

### SetVerifiedAtNil

`func (o *ClearProviderOutputVerifiedEmail) SetVerifiedAtNil(b bool)`

 SetVerifiedAtNil sets the value for VerifiedAt to be an explicit nil

### UnsetVerifiedAt
`func (o *ClearProviderOutputVerifiedEmail) UnsetVerifiedAt()`

UnsetVerifiedAt ensures that no value is present for VerifiedAt, not even an explicit nil
### GetVerificationMethod

`func (o *ClearProviderOutputVerifiedEmail) GetVerificationMethod() string`

GetVerificationMethod returns the VerificationMethod field if non-nil, zero value otherwise.

### GetVerificationMethodOk

`func (o *ClearProviderOutputVerifiedEmail) GetVerificationMethodOk() (*string, bool)`

GetVerificationMethodOk returns a tuple with the VerificationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationMethod

`func (o *ClearProviderOutputVerifiedEmail) SetVerificationMethod(v string)`

SetVerificationMethod sets VerificationMethod field to given value.

### HasVerificationMethod

`func (o *ClearProviderOutputVerifiedEmail) HasVerificationMethod() bool`

HasVerificationMethod returns a boolean if a field has been set.

### SetVerificationMethodNil

`func (o *ClearProviderOutputVerifiedEmail) SetVerificationMethodNil(b bool)`

 SetVerificationMethodNil sets the value for VerificationMethod to be an explicit nil

### UnsetVerificationMethod
`func (o *ClearProviderOutputVerifiedEmail) UnsetVerificationMethod()`

UnsetVerificationMethod ensures that no value is present for VerificationMethod, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


