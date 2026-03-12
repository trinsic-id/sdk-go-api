# FrejaProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | The individual&#39;s full name. | [optional] 
**GivenName** | Pointer to **NullableString** | The individual&#39;s first name. | [optional] 
**FamilyName** | Pointer to **NullableString** | The individual&#39;s last name. | [optional] 
**DateOfBirth** | Pointer to **NullableString** | The date of birth of the individual.              Formatted as an ISO 8601 Date. | [optional] 
**Email** | Pointer to **NullableString** | The individual&#39;s primary email address within Freja. | [optional] 
**EmailVerified** | Pointer to **NullableBool** | A boolean which indicates whether the individual&#39;s primary email address has been verified by Freja. | [optional] 
**AllEmailAddresses** | Pointer to **[]string** | An array of all associated email addresses of the individual. | [optional] 
**PrimaryPhysicalAddress** | Pointer to [**NullableOutputFrejaAddress**](OutputFrejaAddress.md) | The individual&#39;s primary address on file. | [optional] 
**AllPhysicalAddresses** | Pointer to [**[]OutputFrejaAddress**](OutputFrejaAddress.md) | A list of all associated addresses of the individual. | [optional] 
**Age** | Pointer to **NullableInt32** | The individual&#39;s age in years. | [optional] 
**PhoneNumber** | Pointer to **NullableString** | The individual&#39;s phone number in the E.164 format. | [optional] 
**PhoneNumberVerified** | Pointer to **NullableBool** | Whether the individual&#39;s phone number has been verified by Freja. | [optional] 
**PersonalIdentityNumber** | Pointer to **NullableString** | The individual&#39;s personal identity number. The actual value of this field depends on the country of origin used to create the Freja credential. It is typically a Social Security Number, National Identification Number, or equivalent personal identifier. | [optional] 
**Country** | Pointer to **NullableString** | The ISO 3166-1 alpha-2 country code associated with the individual&#39;s country of origin. | [optional] 
**Document** | Pointer to [**NullableOutputFrejaDocument**](OutputFrejaDocument.md) | The underlying document, such as a passport, used to create the Freja credential. | [optional] 
**RegistrationLevel** | Pointer to **NullableString** | The Freja registration level associated with this individual. This can be BASIC, EXTENDED, or PLUS. * BASIC: Individual has a registered account with Freja. * EXTENDED: Individual has an official identity document verified by Freja. * PLUS: Individual has undergone in-person verification with Freja. | [optional] 
**RelyingPartyUserId** | Pointer to **NullableString** | The relying party user identifier for this individual. This is an identifier specific to the individual and the relying party (your service). | [optional] 
**TransactionReference** | Pointer to **NullableString** | The Freja transaction reference for this verification. This is an identifier specific to the verification transaction. | [optional] 

## Methods

### NewFrejaProviderOutput

`func NewFrejaProviderOutput() *FrejaProviderOutput`

NewFrejaProviderOutput instantiates a new FrejaProviderOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFrejaProviderOutputWithDefaults

`func NewFrejaProviderOutputWithDefaults() *FrejaProviderOutput`

NewFrejaProviderOutputWithDefaults instantiates a new FrejaProviderOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FrejaProviderOutput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FrejaProviderOutput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FrejaProviderOutput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FrejaProviderOutput) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *FrejaProviderOutput) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *FrejaProviderOutput) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetGivenName

`func (o *FrejaProviderOutput) GetGivenName() string`

GetGivenName returns the GivenName field if non-nil, zero value otherwise.

### GetGivenNameOk

`func (o *FrejaProviderOutput) GetGivenNameOk() (*string, bool)`

GetGivenNameOk returns a tuple with the GivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenName

`func (o *FrejaProviderOutput) SetGivenName(v string)`

SetGivenName sets GivenName field to given value.

### HasGivenName

`func (o *FrejaProviderOutput) HasGivenName() bool`

HasGivenName returns a boolean if a field has been set.

### SetGivenNameNil

`func (o *FrejaProviderOutput) SetGivenNameNil(b bool)`

 SetGivenNameNil sets the value for GivenName to be an explicit nil

### UnsetGivenName
`func (o *FrejaProviderOutput) UnsetGivenName()`

UnsetGivenName ensures that no value is present for GivenName, not even an explicit nil
### GetFamilyName

`func (o *FrejaProviderOutput) GetFamilyName() string`

GetFamilyName returns the FamilyName field if non-nil, zero value otherwise.

### GetFamilyNameOk

`func (o *FrejaProviderOutput) GetFamilyNameOk() (*string, bool)`

GetFamilyNameOk returns a tuple with the FamilyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyName

`func (o *FrejaProviderOutput) SetFamilyName(v string)`

SetFamilyName sets FamilyName field to given value.

### HasFamilyName

`func (o *FrejaProviderOutput) HasFamilyName() bool`

HasFamilyName returns a boolean if a field has been set.

### SetFamilyNameNil

`func (o *FrejaProviderOutput) SetFamilyNameNil(b bool)`

 SetFamilyNameNil sets the value for FamilyName to be an explicit nil

### UnsetFamilyName
`func (o *FrejaProviderOutput) UnsetFamilyName()`

UnsetFamilyName ensures that no value is present for FamilyName, not even an explicit nil
### GetDateOfBirth

`func (o *FrejaProviderOutput) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *FrejaProviderOutput) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *FrejaProviderOutput) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *FrejaProviderOutput) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### SetDateOfBirthNil

`func (o *FrejaProviderOutput) SetDateOfBirthNil(b bool)`

 SetDateOfBirthNil sets the value for DateOfBirth to be an explicit nil

### UnsetDateOfBirth
`func (o *FrejaProviderOutput) UnsetDateOfBirth()`

UnsetDateOfBirth ensures that no value is present for DateOfBirth, not even an explicit nil
### GetEmail

`func (o *FrejaProviderOutput) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *FrejaProviderOutput) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *FrejaProviderOutput) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *FrejaProviderOutput) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *FrejaProviderOutput) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *FrejaProviderOutput) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetEmailVerified

`func (o *FrejaProviderOutput) GetEmailVerified() bool`

GetEmailVerified returns the EmailVerified field if non-nil, zero value otherwise.

### GetEmailVerifiedOk

`func (o *FrejaProviderOutput) GetEmailVerifiedOk() (*bool, bool)`

GetEmailVerifiedOk returns a tuple with the EmailVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailVerified

`func (o *FrejaProviderOutput) SetEmailVerified(v bool)`

SetEmailVerified sets EmailVerified field to given value.

### HasEmailVerified

`func (o *FrejaProviderOutput) HasEmailVerified() bool`

HasEmailVerified returns a boolean if a field has been set.

### SetEmailVerifiedNil

`func (o *FrejaProviderOutput) SetEmailVerifiedNil(b bool)`

 SetEmailVerifiedNil sets the value for EmailVerified to be an explicit nil

### UnsetEmailVerified
`func (o *FrejaProviderOutput) UnsetEmailVerified()`

UnsetEmailVerified ensures that no value is present for EmailVerified, not even an explicit nil
### GetAllEmailAddresses

`func (o *FrejaProviderOutput) GetAllEmailAddresses() []string`

GetAllEmailAddresses returns the AllEmailAddresses field if non-nil, zero value otherwise.

### GetAllEmailAddressesOk

`func (o *FrejaProviderOutput) GetAllEmailAddressesOk() (*[]string, bool)`

GetAllEmailAddressesOk returns a tuple with the AllEmailAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllEmailAddresses

`func (o *FrejaProviderOutput) SetAllEmailAddresses(v []string)`

SetAllEmailAddresses sets AllEmailAddresses field to given value.

### HasAllEmailAddresses

`func (o *FrejaProviderOutput) HasAllEmailAddresses() bool`

HasAllEmailAddresses returns a boolean if a field has been set.

### SetAllEmailAddressesNil

`func (o *FrejaProviderOutput) SetAllEmailAddressesNil(b bool)`

 SetAllEmailAddressesNil sets the value for AllEmailAddresses to be an explicit nil

### UnsetAllEmailAddresses
`func (o *FrejaProviderOutput) UnsetAllEmailAddresses()`

UnsetAllEmailAddresses ensures that no value is present for AllEmailAddresses, not even an explicit nil
### GetPrimaryPhysicalAddress

`func (o *FrejaProviderOutput) GetPrimaryPhysicalAddress() OutputFrejaAddress`

GetPrimaryPhysicalAddress returns the PrimaryPhysicalAddress field if non-nil, zero value otherwise.

### GetPrimaryPhysicalAddressOk

`func (o *FrejaProviderOutput) GetPrimaryPhysicalAddressOk() (*OutputFrejaAddress, bool)`

GetPrimaryPhysicalAddressOk returns a tuple with the PrimaryPhysicalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryPhysicalAddress

`func (o *FrejaProviderOutput) SetPrimaryPhysicalAddress(v OutputFrejaAddress)`

SetPrimaryPhysicalAddress sets PrimaryPhysicalAddress field to given value.

### HasPrimaryPhysicalAddress

`func (o *FrejaProviderOutput) HasPrimaryPhysicalAddress() bool`

HasPrimaryPhysicalAddress returns a boolean if a field has been set.

### SetPrimaryPhysicalAddressNil

`func (o *FrejaProviderOutput) SetPrimaryPhysicalAddressNil(b bool)`

 SetPrimaryPhysicalAddressNil sets the value for PrimaryPhysicalAddress to be an explicit nil

### UnsetPrimaryPhysicalAddress
`func (o *FrejaProviderOutput) UnsetPrimaryPhysicalAddress()`

UnsetPrimaryPhysicalAddress ensures that no value is present for PrimaryPhysicalAddress, not even an explicit nil
### GetAllPhysicalAddresses

`func (o *FrejaProviderOutput) GetAllPhysicalAddresses() []OutputFrejaAddress`

GetAllPhysicalAddresses returns the AllPhysicalAddresses field if non-nil, zero value otherwise.

### GetAllPhysicalAddressesOk

`func (o *FrejaProviderOutput) GetAllPhysicalAddressesOk() (*[]OutputFrejaAddress, bool)`

GetAllPhysicalAddressesOk returns a tuple with the AllPhysicalAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllPhysicalAddresses

`func (o *FrejaProviderOutput) SetAllPhysicalAddresses(v []OutputFrejaAddress)`

SetAllPhysicalAddresses sets AllPhysicalAddresses field to given value.

### HasAllPhysicalAddresses

`func (o *FrejaProviderOutput) HasAllPhysicalAddresses() bool`

HasAllPhysicalAddresses returns a boolean if a field has been set.

### SetAllPhysicalAddressesNil

`func (o *FrejaProviderOutput) SetAllPhysicalAddressesNil(b bool)`

 SetAllPhysicalAddressesNil sets the value for AllPhysicalAddresses to be an explicit nil

### UnsetAllPhysicalAddresses
`func (o *FrejaProviderOutput) UnsetAllPhysicalAddresses()`

UnsetAllPhysicalAddresses ensures that no value is present for AllPhysicalAddresses, not even an explicit nil
### GetAge

`func (o *FrejaProviderOutput) GetAge() int32`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *FrejaProviderOutput) GetAgeOk() (*int32, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *FrejaProviderOutput) SetAge(v int32)`

SetAge sets Age field to given value.

### HasAge

`func (o *FrejaProviderOutput) HasAge() bool`

HasAge returns a boolean if a field has been set.

### SetAgeNil

`func (o *FrejaProviderOutput) SetAgeNil(b bool)`

 SetAgeNil sets the value for Age to be an explicit nil

### UnsetAge
`func (o *FrejaProviderOutput) UnsetAge()`

UnsetAge ensures that no value is present for Age, not even an explicit nil
### GetPhoneNumber

`func (o *FrejaProviderOutput) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *FrejaProviderOutput) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *FrejaProviderOutput) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *FrejaProviderOutput) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### SetPhoneNumberNil

`func (o *FrejaProviderOutput) SetPhoneNumberNil(b bool)`

 SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil

### UnsetPhoneNumber
`func (o *FrejaProviderOutput) UnsetPhoneNumber()`

UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
### GetPhoneNumberVerified

`func (o *FrejaProviderOutput) GetPhoneNumberVerified() bool`

GetPhoneNumberVerified returns the PhoneNumberVerified field if non-nil, zero value otherwise.

### GetPhoneNumberVerifiedOk

`func (o *FrejaProviderOutput) GetPhoneNumberVerifiedOk() (*bool, bool)`

GetPhoneNumberVerifiedOk returns a tuple with the PhoneNumberVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberVerified

`func (o *FrejaProviderOutput) SetPhoneNumberVerified(v bool)`

SetPhoneNumberVerified sets PhoneNumberVerified field to given value.

### HasPhoneNumberVerified

`func (o *FrejaProviderOutput) HasPhoneNumberVerified() bool`

HasPhoneNumberVerified returns a boolean if a field has been set.

### SetPhoneNumberVerifiedNil

`func (o *FrejaProviderOutput) SetPhoneNumberVerifiedNil(b bool)`

 SetPhoneNumberVerifiedNil sets the value for PhoneNumberVerified to be an explicit nil

### UnsetPhoneNumberVerified
`func (o *FrejaProviderOutput) UnsetPhoneNumberVerified()`

UnsetPhoneNumberVerified ensures that no value is present for PhoneNumberVerified, not even an explicit nil
### GetPersonalIdentityNumber

`func (o *FrejaProviderOutput) GetPersonalIdentityNumber() string`

GetPersonalIdentityNumber returns the PersonalIdentityNumber field if non-nil, zero value otherwise.

### GetPersonalIdentityNumberOk

`func (o *FrejaProviderOutput) GetPersonalIdentityNumberOk() (*string, bool)`

GetPersonalIdentityNumberOk returns a tuple with the PersonalIdentityNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalIdentityNumber

`func (o *FrejaProviderOutput) SetPersonalIdentityNumber(v string)`

SetPersonalIdentityNumber sets PersonalIdentityNumber field to given value.

### HasPersonalIdentityNumber

`func (o *FrejaProviderOutput) HasPersonalIdentityNumber() bool`

HasPersonalIdentityNumber returns a boolean if a field has been set.

### SetPersonalIdentityNumberNil

`func (o *FrejaProviderOutput) SetPersonalIdentityNumberNil(b bool)`

 SetPersonalIdentityNumberNil sets the value for PersonalIdentityNumber to be an explicit nil

### UnsetPersonalIdentityNumber
`func (o *FrejaProviderOutput) UnsetPersonalIdentityNumber()`

UnsetPersonalIdentityNumber ensures that no value is present for PersonalIdentityNumber, not even an explicit nil
### GetCountry

`func (o *FrejaProviderOutput) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *FrejaProviderOutput) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *FrejaProviderOutput) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *FrejaProviderOutput) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *FrejaProviderOutput) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *FrejaProviderOutput) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetDocument

`func (o *FrejaProviderOutput) GetDocument() OutputFrejaDocument`

GetDocument returns the Document field if non-nil, zero value otherwise.

### GetDocumentOk

`func (o *FrejaProviderOutput) GetDocumentOk() (*OutputFrejaDocument, bool)`

GetDocumentOk returns a tuple with the Document field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocument

`func (o *FrejaProviderOutput) SetDocument(v OutputFrejaDocument)`

SetDocument sets Document field to given value.

### HasDocument

`func (o *FrejaProviderOutput) HasDocument() bool`

HasDocument returns a boolean if a field has been set.

### SetDocumentNil

`func (o *FrejaProviderOutput) SetDocumentNil(b bool)`

 SetDocumentNil sets the value for Document to be an explicit nil

### UnsetDocument
`func (o *FrejaProviderOutput) UnsetDocument()`

UnsetDocument ensures that no value is present for Document, not even an explicit nil
### GetRegistrationLevel

`func (o *FrejaProviderOutput) GetRegistrationLevel() string`

GetRegistrationLevel returns the RegistrationLevel field if non-nil, zero value otherwise.

### GetRegistrationLevelOk

`func (o *FrejaProviderOutput) GetRegistrationLevelOk() (*string, bool)`

GetRegistrationLevelOk returns a tuple with the RegistrationLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationLevel

`func (o *FrejaProviderOutput) SetRegistrationLevel(v string)`

SetRegistrationLevel sets RegistrationLevel field to given value.

### HasRegistrationLevel

`func (o *FrejaProviderOutput) HasRegistrationLevel() bool`

HasRegistrationLevel returns a boolean if a field has been set.

### SetRegistrationLevelNil

`func (o *FrejaProviderOutput) SetRegistrationLevelNil(b bool)`

 SetRegistrationLevelNil sets the value for RegistrationLevel to be an explicit nil

### UnsetRegistrationLevel
`func (o *FrejaProviderOutput) UnsetRegistrationLevel()`

UnsetRegistrationLevel ensures that no value is present for RegistrationLevel, not even an explicit nil
### GetRelyingPartyUserId

`func (o *FrejaProviderOutput) GetRelyingPartyUserId() string`

GetRelyingPartyUserId returns the RelyingPartyUserId field if non-nil, zero value otherwise.

### GetRelyingPartyUserIdOk

`func (o *FrejaProviderOutput) GetRelyingPartyUserIdOk() (*string, bool)`

GetRelyingPartyUserIdOk returns a tuple with the RelyingPartyUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelyingPartyUserId

`func (o *FrejaProviderOutput) SetRelyingPartyUserId(v string)`

SetRelyingPartyUserId sets RelyingPartyUserId field to given value.

### HasRelyingPartyUserId

`func (o *FrejaProviderOutput) HasRelyingPartyUserId() bool`

HasRelyingPartyUserId returns a boolean if a field has been set.

### SetRelyingPartyUserIdNil

`func (o *FrejaProviderOutput) SetRelyingPartyUserIdNil(b bool)`

 SetRelyingPartyUserIdNil sets the value for RelyingPartyUserId to be an explicit nil

### UnsetRelyingPartyUserId
`func (o *FrejaProviderOutput) UnsetRelyingPartyUserId()`

UnsetRelyingPartyUserId ensures that no value is present for RelyingPartyUserId, not even an explicit nil
### GetTransactionReference

`func (o *FrejaProviderOutput) GetTransactionReference() string`

GetTransactionReference returns the TransactionReference field if non-nil, zero value otherwise.

### GetTransactionReferenceOk

`func (o *FrejaProviderOutput) GetTransactionReferenceOk() (*string, bool)`

GetTransactionReferenceOk returns a tuple with the TransactionReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionReference

`func (o *FrejaProviderOutput) SetTransactionReference(v string)`

SetTransactionReference sets TransactionReference field to given value.

### HasTransactionReference

`func (o *FrejaProviderOutput) HasTransactionReference() bool`

HasTransactionReference returns a boolean if a field has been set.

### SetTransactionReferenceNil

`func (o *FrejaProviderOutput) SetTransactionReferenceNil(b bool)`

 SetTransactionReferenceNil sets the value for TransactionReference to be an explicit nil

### UnsetTransactionReference
`func (o *FrejaProviderOutput) UnsetTransactionReference()`

UnsetTransactionReference ensures that no value is present for TransactionReference, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


