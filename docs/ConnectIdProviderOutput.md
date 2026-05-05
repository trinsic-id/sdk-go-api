# ConnectIdProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Birthdate** | Pointer to **NullableString** | The date of birth of an individual in YYYY-MM-DD format. | [optional] 
**GivenName** | Pointer to **NullableString** | The given name of an individual. | [optional] 
**MiddleName** | Pointer to **NullableString** | The middle name of an individual. | [optional] 
**FamilyName** | Pointer to **NullableString** | The family/last name of an individual. | [optional] 
**PhoneNumber** | Pointer to **NullableString** | The phone number of an individual. | [optional] 
**Email** | Pointer to **NullableString** | The email of an individual. | [optional] 
**Sub** | Pointer to **NullableString** | The OpenID Connect (OIDC) subject identifier (sub). | [optional] 
**AgeVerification** | Pointer to [**NullableConnectIdAgeVerification**](ConnectIdAgeVerification.md) | The age verification claim and whether the individual meets the age requirement. | [optional] 
**Transaction** | Pointer to **NullableString** | The verification session&#39;s transaction number.              This is a unique identifier assigned to a single ConnectID transaction flow. It can be used for audit purposes or to flag fraudulent activity. | [optional] 
**Address** | Pointer to [**NullableConnectIdAddress**](ConnectIdAddress.md) | The address. | [optional] 

## Methods

### NewConnectIdProviderOutput

`func NewConnectIdProviderOutput() *ConnectIdProviderOutput`

NewConnectIdProviderOutput instantiates a new ConnectIdProviderOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectIdProviderOutputWithDefaults

`func NewConnectIdProviderOutputWithDefaults() *ConnectIdProviderOutput`

NewConnectIdProviderOutputWithDefaults instantiates a new ConnectIdProviderOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBirthdate

`func (o *ConnectIdProviderOutput) GetBirthdate() string`

GetBirthdate returns the Birthdate field if non-nil, zero value otherwise.

### GetBirthdateOk

`func (o *ConnectIdProviderOutput) GetBirthdateOk() (*string, bool)`

GetBirthdateOk returns a tuple with the Birthdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthdate

`func (o *ConnectIdProviderOutput) SetBirthdate(v string)`

SetBirthdate sets Birthdate field to given value.

### HasBirthdate

`func (o *ConnectIdProviderOutput) HasBirthdate() bool`

HasBirthdate returns a boolean if a field has been set.

### SetBirthdateNil

`func (o *ConnectIdProviderOutput) SetBirthdateNil(b bool)`

 SetBirthdateNil sets the value for Birthdate to be an explicit nil

### UnsetBirthdate
`func (o *ConnectIdProviderOutput) UnsetBirthdate()`

UnsetBirthdate ensures that no value is present for Birthdate, not even an explicit nil
### GetGivenName

`func (o *ConnectIdProviderOutput) GetGivenName() string`

GetGivenName returns the GivenName field if non-nil, zero value otherwise.

### GetGivenNameOk

`func (o *ConnectIdProviderOutput) GetGivenNameOk() (*string, bool)`

GetGivenNameOk returns a tuple with the GivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenName

`func (o *ConnectIdProviderOutput) SetGivenName(v string)`

SetGivenName sets GivenName field to given value.

### HasGivenName

`func (o *ConnectIdProviderOutput) HasGivenName() bool`

HasGivenName returns a boolean if a field has been set.

### SetGivenNameNil

`func (o *ConnectIdProviderOutput) SetGivenNameNil(b bool)`

 SetGivenNameNil sets the value for GivenName to be an explicit nil

### UnsetGivenName
`func (o *ConnectIdProviderOutput) UnsetGivenName()`

UnsetGivenName ensures that no value is present for GivenName, not even an explicit nil
### GetMiddleName

`func (o *ConnectIdProviderOutput) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *ConnectIdProviderOutput) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *ConnectIdProviderOutput) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *ConnectIdProviderOutput) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### SetMiddleNameNil

`func (o *ConnectIdProviderOutput) SetMiddleNameNil(b bool)`

 SetMiddleNameNil sets the value for MiddleName to be an explicit nil

### UnsetMiddleName
`func (o *ConnectIdProviderOutput) UnsetMiddleName()`

UnsetMiddleName ensures that no value is present for MiddleName, not even an explicit nil
### GetFamilyName

`func (o *ConnectIdProviderOutput) GetFamilyName() string`

GetFamilyName returns the FamilyName field if non-nil, zero value otherwise.

### GetFamilyNameOk

`func (o *ConnectIdProviderOutput) GetFamilyNameOk() (*string, bool)`

GetFamilyNameOk returns a tuple with the FamilyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyName

`func (o *ConnectIdProviderOutput) SetFamilyName(v string)`

SetFamilyName sets FamilyName field to given value.

### HasFamilyName

`func (o *ConnectIdProviderOutput) HasFamilyName() bool`

HasFamilyName returns a boolean if a field has been set.

### SetFamilyNameNil

`func (o *ConnectIdProviderOutput) SetFamilyNameNil(b bool)`

 SetFamilyNameNil sets the value for FamilyName to be an explicit nil

### UnsetFamilyName
`func (o *ConnectIdProviderOutput) UnsetFamilyName()`

UnsetFamilyName ensures that no value is present for FamilyName, not even an explicit nil
### GetPhoneNumber

`func (o *ConnectIdProviderOutput) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *ConnectIdProviderOutput) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *ConnectIdProviderOutput) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *ConnectIdProviderOutput) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### SetPhoneNumberNil

`func (o *ConnectIdProviderOutput) SetPhoneNumberNil(b bool)`

 SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil

### UnsetPhoneNumber
`func (o *ConnectIdProviderOutput) UnsetPhoneNumber()`

UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
### GetEmail

`func (o *ConnectIdProviderOutput) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ConnectIdProviderOutput) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ConnectIdProviderOutput) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ConnectIdProviderOutput) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *ConnectIdProviderOutput) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *ConnectIdProviderOutput) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetSub

`func (o *ConnectIdProviderOutput) GetSub() string`

GetSub returns the Sub field if non-nil, zero value otherwise.

### GetSubOk

`func (o *ConnectIdProviderOutput) GetSubOk() (*string, bool)`

GetSubOk returns a tuple with the Sub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSub

`func (o *ConnectIdProviderOutput) SetSub(v string)`

SetSub sets Sub field to given value.

### HasSub

`func (o *ConnectIdProviderOutput) HasSub() bool`

HasSub returns a boolean if a field has been set.

### SetSubNil

`func (o *ConnectIdProviderOutput) SetSubNil(b bool)`

 SetSubNil sets the value for Sub to be an explicit nil

### UnsetSub
`func (o *ConnectIdProviderOutput) UnsetSub()`

UnsetSub ensures that no value is present for Sub, not even an explicit nil
### GetAgeVerification

`func (o *ConnectIdProviderOutput) GetAgeVerification() ConnectIdAgeVerification`

GetAgeVerification returns the AgeVerification field if non-nil, zero value otherwise.

### GetAgeVerificationOk

`func (o *ConnectIdProviderOutput) GetAgeVerificationOk() (*ConnectIdAgeVerification, bool)`

GetAgeVerificationOk returns a tuple with the AgeVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgeVerification

`func (o *ConnectIdProviderOutput) SetAgeVerification(v ConnectIdAgeVerification)`

SetAgeVerification sets AgeVerification field to given value.

### HasAgeVerification

`func (o *ConnectIdProviderOutput) HasAgeVerification() bool`

HasAgeVerification returns a boolean if a field has been set.

### SetAgeVerificationNil

`func (o *ConnectIdProviderOutput) SetAgeVerificationNil(b bool)`

 SetAgeVerificationNil sets the value for AgeVerification to be an explicit nil

### UnsetAgeVerification
`func (o *ConnectIdProviderOutput) UnsetAgeVerification()`

UnsetAgeVerification ensures that no value is present for AgeVerification, not even an explicit nil
### GetTransaction

`func (o *ConnectIdProviderOutput) GetTransaction() string`

GetTransaction returns the Transaction field if non-nil, zero value otherwise.

### GetTransactionOk

`func (o *ConnectIdProviderOutput) GetTransactionOk() (*string, bool)`

GetTransactionOk returns a tuple with the Transaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransaction

`func (o *ConnectIdProviderOutput) SetTransaction(v string)`

SetTransaction sets Transaction field to given value.

### HasTransaction

`func (o *ConnectIdProviderOutput) HasTransaction() bool`

HasTransaction returns a boolean if a field has been set.

### SetTransactionNil

`func (o *ConnectIdProviderOutput) SetTransactionNil(b bool)`

 SetTransactionNil sets the value for Transaction to be an explicit nil

### UnsetTransaction
`func (o *ConnectIdProviderOutput) UnsetTransaction()`

UnsetTransaction ensures that no value is present for Transaction, not even an explicit nil
### GetAddress

`func (o *ConnectIdProviderOutput) GetAddress() ConnectIdAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ConnectIdProviderOutput) GetAddressOk() (*ConnectIdAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ConnectIdProviderOutput) SetAddress(v ConnectIdAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ConnectIdProviderOutput) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *ConnectIdProviderOutput) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *ConnectIdProviderOutput) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


