# InteracProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sub** | Pointer to **NullableString** | The OpenID Connect subject identifier for the verified individual. | [optional] 
**GivenName** | Pointer to **NullableString** | The given name of the verified individual. | [optional] 
**MiddleName** | Pointer to **NullableString** | The middle name of the verified individual. | [optional] 
**FamilyName** | Pointer to **NullableString** | The family name of the verified individual. | [optional] 
**Title** | Pointer to **NullableString** | The title or honorific of the verified individual, when supplied by the source. | [optional] 
**DateOfBirth** | Pointer to **NullableString** | The date of birth of the verified individual. | [optional] 
**Email** | Pointer to **NullableString** | The email address of the verified individual. | [optional] 
**PhoneNumber** | Pointer to **NullableString** | The phone number of the verified individual in E.164 format. | [optional] 
**Address** | Pointer to [**NullableInteracAddress**](InteracAddress.md) | The verified address. | [optional] 

## Methods

### NewInteracProviderOutput

`func NewInteracProviderOutput() *InteracProviderOutput`

NewInteracProviderOutput instantiates a new InteracProviderOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInteracProviderOutputWithDefaults

`func NewInteracProviderOutputWithDefaults() *InteracProviderOutput`

NewInteracProviderOutputWithDefaults instantiates a new InteracProviderOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSub

`func (o *InteracProviderOutput) GetSub() string`

GetSub returns the Sub field if non-nil, zero value otherwise.

### GetSubOk

`func (o *InteracProviderOutput) GetSubOk() (*string, bool)`

GetSubOk returns a tuple with the Sub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSub

`func (o *InteracProviderOutput) SetSub(v string)`

SetSub sets Sub field to given value.

### HasSub

`func (o *InteracProviderOutput) HasSub() bool`

HasSub returns a boolean if a field has been set.

### SetSubNil

`func (o *InteracProviderOutput) SetSubNil(b bool)`

 SetSubNil sets the value for Sub to be an explicit nil

### UnsetSub
`func (o *InteracProviderOutput) UnsetSub()`

UnsetSub ensures that no value is present for Sub, not even an explicit nil
### GetGivenName

`func (o *InteracProviderOutput) GetGivenName() string`

GetGivenName returns the GivenName field if non-nil, zero value otherwise.

### GetGivenNameOk

`func (o *InteracProviderOutput) GetGivenNameOk() (*string, bool)`

GetGivenNameOk returns a tuple with the GivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenName

`func (o *InteracProviderOutput) SetGivenName(v string)`

SetGivenName sets GivenName field to given value.

### HasGivenName

`func (o *InteracProviderOutput) HasGivenName() bool`

HasGivenName returns a boolean if a field has been set.

### SetGivenNameNil

`func (o *InteracProviderOutput) SetGivenNameNil(b bool)`

 SetGivenNameNil sets the value for GivenName to be an explicit nil

### UnsetGivenName
`func (o *InteracProviderOutput) UnsetGivenName()`

UnsetGivenName ensures that no value is present for GivenName, not even an explicit nil
### GetMiddleName

`func (o *InteracProviderOutput) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *InteracProviderOutput) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *InteracProviderOutput) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *InteracProviderOutput) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### SetMiddleNameNil

`func (o *InteracProviderOutput) SetMiddleNameNil(b bool)`

 SetMiddleNameNil sets the value for MiddleName to be an explicit nil

### UnsetMiddleName
`func (o *InteracProviderOutput) UnsetMiddleName()`

UnsetMiddleName ensures that no value is present for MiddleName, not even an explicit nil
### GetFamilyName

`func (o *InteracProviderOutput) GetFamilyName() string`

GetFamilyName returns the FamilyName field if non-nil, zero value otherwise.

### GetFamilyNameOk

`func (o *InteracProviderOutput) GetFamilyNameOk() (*string, bool)`

GetFamilyNameOk returns a tuple with the FamilyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyName

`func (o *InteracProviderOutput) SetFamilyName(v string)`

SetFamilyName sets FamilyName field to given value.

### HasFamilyName

`func (o *InteracProviderOutput) HasFamilyName() bool`

HasFamilyName returns a boolean if a field has been set.

### SetFamilyNameNil

`func (o *InteracProviderOutput) SetFamilyNameNil(b bool)`

 SetFamilyNameNil sets the value for FamilyName to be an explicit nil

### UnsetFamilyName
`func (o *InteracProviderOutput) UnsetFamilyName()`

UnsetFamilyName ensures that no value is present for FamilyName, not even an explicit nil
### GetTitle

`func (o *InteracProviderOutput) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *InteracProviderOutput) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *InteracProviderOutput) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *InteracProviderOutput) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *InteracProviderOutput) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *InteracProviderOutput) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDateOfBirth

`func (o *InteracProviderOutput) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *InteracProviderOutput) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *InteracProviderOutput) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *InteracProviderOutput) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### SetDateOfBirthNil

`func (o *InteracProviderOutput) SetDateOfBirthNil(b bool)`

 SetDateOfBirthNil sets the value for DateOfBirth to be an explicit nil

### UnsetDateOfBirth
`func (o *InteracProviderOutput) UnsetDateOfBirth()`

UnsetDateOfBirth ensures that no value is present for DateOfBirth, not even an explicit nil
### GetEmail

`func (o *InteracProviderOutput) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *InteracProviderOutput) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *InteracProviderOutput) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *InteracProviderOutput) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *InteracProviderOutput) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *InteracProviderOutput) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetPhoneNumber

`func (o *InteracProviderOutput) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *InteracProviderOutput) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *InteracProviderOutput) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *InteracProviderOutput) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### SetPhoneNumberNil

`func (o *InteracProviderOutput) SetPhoneNumberNil(b bool)`

 SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil

### UnsetPhoneNumber
`func (o *InteracProviderOutput) UnsetPhoneNumber()`

UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
### GetAddress

`func (o *InteracProviderOutput) GetAddress() InteracAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *InteracProviderOutput) GetAddressOk() (*InteracAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *InteracProviderOutput) SetAddress(v InteracAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *InteracProviderOutput) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *InteracProviderOutput) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *InteracProviderOutput) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


