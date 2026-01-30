# YotiProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RememberMeId** | Pointer to **NullableString** | The unique id for your service of the individual. This is not shared between different services so if the user logs into a different service, the user will have a new id for the other service. | [optional] 
**Email** | Pointer to **NullableString** | The email address of the individual | [optional] 
**GivenName** | Pointer to **NullableString** | The given name of the individual. This can correspond to first and middle names in English. | [optional] 
**FamilyName** | Pointer to **NullableString** | The family name of the individual | [optional] 
**FullName** | Pointer to **NullableString** | Full name of the individual which is composed of the given and family name. Unverified full name may be safe to use. For example, Aadhaar card provides full name, but due to certain limitations of verifying Aadhaar, Yoti is unable to verify this property so it is listed as \&quot;unverified\&quot;. | [optional] 
**DateOfBirth** | Pointer to **NullableString** | The date of birth of the individual. Unverified date of birth may be safe to use. For example, Aadhaar card provides date of birth, but due to certain limitations of verifying Aadhaar, Yoti is unable to verify this property so it is listed as \&quot;unverified\&quot;. | [optional] 
**Gender** | Pointer to **NullableString** | The gender of the individual. Depending on the source or country, gender may not be verified. This is a setting that can be configured while setting up scopes.              Unverified gender may be safe to use. For example, Aadhaar card provides gender, but due to certain limitations of verifying Aadhaar, Yoti is unable to verify this property so it is listed as \&quot;unverified\&quot;. In some countries Possible values: - \&quot;MALE\&quot; - \&quot;FEMALE\&quot; - \&quot;TRANSGENDER\&quot; - \&quot;OTHER\&quot; | [optional] 
**Nationality** | Pointer to **NullableString** | The nationality of the individual as a ISO alpha-3 code | [optional] 
**MobileNumber** | Pointer to **NullableString** | The mobile phone number of the individual. This number was verified with a one time password (OTP) during the user&#39;s registration with Yoti. In some cases, Yoti may do additional checks against its sources to confirm the user&#39;s identity. | [optional] 
**StructuredPostalAddress** | Pointer to [**NullableYotiStructuredPostalAddress**](YotiStructuredPostalAddress.md) | A structured postal address for the individual that comes from the underlying document or can be manually added. In the case of being manually added, the address is considered \&quot;unverified\&quot;. Unverified addresses may be safe to use. For example, Aadhaar card provides an address, but due to certain limitations of verifying Aadhaar, Yoti is unable to verify this property so it is listed as \&quot;unverified\&quot;. | [optional] 
**PostalAddress** | Pointer to **NullableString** | The whole address for the individual that comes from the underlying document or can be manually added. In the case of being manually added, the address is considered \&quot;unverified\&quot;. The format of this will be different for each country. Refer to &#x60;StructuredPostalAddress&#x60; for a detailed version of the address. | [optional] 
**DocumentDetails** | Pointer to [**NullableYotiDocumentDetails**](YotiDocumentDetails.md) | The details of the underlying document used to help create the Yoti credential | [optional] 

## Methods

### NewYotiProviderOutput

`func NewYotiProviderOutput() *YotiProviderOutput`

NewYotiProviderOutput instantiates a new YotiProviderOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewYotiProviderOutputWithDefaults

`func NewYotiProviderOutputWithDefaults() *YotiProviderOutput`

NewYotiProviderOutputWithDefaults instantiates a new YotiProviderOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRememberMeId

`func (o *YotiProviderOutput) GetRememberMeId() string`

GetRememberMeId returns the RememberMeId field if non-nil, zero value otherwise.

### GetRememberMeIdOk

`func (o *YotiProviderOutput) GetRememberMeIdOk() (*string, bool)`

GetRememberMeIdOk returns a tuple with the RememberMeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRememberMeId

`func (o *YotiProviderOutput) SetRememberMeId(v string)`

SetRememberMeId sets RememberMeId field to given value.

### HasRememberMeId

`func (o *YotiProviderOutput) HasRememberMeId() bool`

HasRememberMeId returns a boolean if a field has been set.

### SetRememberMeIdNil

`func (o *YotiProviderOutput) SetRememberMeIdNil(b bool)`

 SetRememberMeIdNil sets the value for RememberMeId to be an explicit nil

### UnsetRememberMeId
`func (o *YotiProviderOutput) UnsetRememberMeId()`

UnsetRememberMeId ensures that no value is present for RememberMeId, not even an explicit nil
### GetEmail

`func (o *YotiProviderOutput) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *YotiProviderOutput) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *YotiProviderOutput) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *YotiProviderOutput) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *YotiProviderOutput) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *YotiProviderOutput) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetGivenName

`func (o *YotiProviderOutput) GetGivenName() string`

GetGivenName returns the GivenName field if non-nil, zero value otherwise.

### GetGivenNameOk

`func (o *YotiProviderOutput) GetGivenNameOk() (*string, bool)`

GetGivenNameOk returns a tuple with the GivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenName

`func (o *YotiProviderOutput) SetGivenName(v string)`

SetGivenName sets GivenName field to given value.

### HasGivenName

`func (o *YotiProviderOutput) HasGivenName() bool`

HasGivenName returns a boolean if a field has been set.

### SetGivenNameNil

`func (o *YotiProviderOutput) SetGivenNameNil(b bool)`

 SetGivenNameNil sets the value for GivenName to be an explicit nil

### UnsetGivenName
`func (o *YotiProviderOutput) UnsetGivenName()`

UnsetGivenName ensures that no value is present for GivenName, not even an explicit nil
### GetFamilyName

`func (o *YotiProviderOutput) GetFamilyName() string`

GetFamilyName returns the FamilyName field if non-nil, zero value otherwise.

### GetFamilyNameOk

`func (o *YotiProviderOutput) GetFamilyNameOk() (*string, bool)`

GetFamilyNameOk returns a tuple with the FamilyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyName

`func (o *YotiProviderOutput) SetFamilyName(v string)`

SetFamilyName sets FamilyName field to given value.

### HasFamilyName

`func (o *YotiProviderOutput) HasFamilyName() bool`

HasFamilyName returns a boolean if a field has been set.

### SetFamilyNameNil

`func (o *YotiProviderOutput) SetFamilyNameNil(b bool)`

 SetFamilyNameNil sets the value for FamilyName to be an explicit nil

### UnsetFamilyName
`func (o *YotiProviderOutput) UnsetFamilyName()`

UnsetFamilyName ensures that no value is present for FamilyName, not even an explicit nil
### GetFullName

`func (o *YotiProviderOutput) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *YotiProviderOutput) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *YotiProviderOutput) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *YotiProviderOutput) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### SetFullNameNil

`func (o *YotiProviderOutput) SetFullNameNil(b bool)`

 SetFullNameNil sets the value for FullName to be an explicit nil

### UnsetFullName
`func (o *YotiProviderOutput) UnsetFullName()`

UnsetFullName ensures that no value is present for FullName, not even an explicit nil
### GetDateOfBirth

`func (o *YotiProviderOutput) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *YotiProviderOutput) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *YotiProviderOutput) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *YotiProviderOutput) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### SetDateOfBirthNil

`func (o *YotiProviderOutput) SetDateOfBirthNil(b bool)`

 SetDateOfBirthNil sets the value for DateOfBirth to be an explicit nil

### UnsetDateOfBirth
`func (o *YotiProviderOutput) UnsetDateOfBirth()`

UnsetDateOfBirth ensures that no value is present for DateOfBirth, not even an explicit nil
### GetGender

`func (o *YotiProviderOutput) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *YotiProviderOutput) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *YotiProviderOutput) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *YotiProviderOutput) HasGender() bool`

HasGender returns a boolean if a field has been set.

### SetGenderNil

`func (o *YotiProviderOutput) SetGenderNil(b bool)`

 SetGenderNil sets the value for Gender to be an explicit nil

### UnsetGender
`func (o *YotiProviderOutput) UnsetGender()`

UnsetGender ensures that no value is present for Gender, not even an explicit nil
### GetNationality

`func (o *YotiProviderOutput) GetNationality() string`

GetNationality returns the Nationality field if non-nil, zero value otherwise.

### GetNationalityOk

`func (o *YotiProviderOutput) GetNationalityOk() (*string, bool)`

GetNationalityOk returns a tuple with the Nationality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationality

`func (o *YotiProviderOutput) SetNationality(v string)`

SetNationality sets Nationality field to given value.

### HasNationality

`func (o *YotiProviderOutput) HasNationality() bool`

HasNationality returns a boolean if a field has been set.

### SetNationalityNil

`func (o *YotiProviderOutput) SetNationalityNil(b bool)`

 SetNationalityNil sets the value for Nationality to be an explicit nil

### UnsetNationality
`func (o *YotiProviderOutput) UnsetNationality()`

UnsetNationality ensures that no value is present for Nationality, not even an explicit nil
### GetMobileNumber

`func (o *YotiProviderOutput) GetMobileNumber() string`

GetMobileNumber returns the MobileNumber field if non-nil, zero value otherwise.

### GetMobileNumberOk

`func (o *YotiProviderOutput) GetMobileNumberOk() (*string, bool)`

GetMobileNumberOk returns a tuple with the MobileNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileNumber

`func (o *YotiProviderOutput) SetMobileNumber(v string)`

SetMobileNumber sets MobileNumber field to given value.

### HasMobileNumber

`func (o *YotiProviderOutput) HasMobileNumber() bool`

HasMobileNumber returns a boolean if a field has been set.

### SetMobileNumberNil

`func (o *YotiProviderOutput) SetMobileNumberNil(b bool)`

 SetMobileNumberNil sets the value for MobileNumber to be an explicit nil

### UnsetMobileNumber
`func (o *YotiProviderOutput) UnsetMobileNumber()`

UnsetMobileNumber ensures that no value is present for MobileNumber, not even an explicit nil
### GetStructuredPostalAddress

`func (o *YotiProviderOutput) GetStructuredPostalAddress() YotiStructuredPostalAddress`

GetStructuredPostalAddress returns the StructuredPostalAddress field if non-nil, zero value otherwise.

### GetStructuredPostalAddressOk

`func (o *YotiProviderOutput) GetStructuredPostalAddressOk() (*YotiStructuredPostalAddress, bool)`

GetStructuredPostalAddressOk returns a tuple with the StructuredPostalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructuredPostalAddress

`func (o *YotiProviderOutput) SetStructuredPostalAddress(v YotiStructuredPostalAddress)`

SetStructuredPostalAddress sets StructuredPostalAddress field to given value.

### HasStructuredPostalAddress

`func (o *YotiProviderOutput) HasStructuredPostalAddress() bool`

HasStructuredPostalAddress returns a boolean if a field has been set.

### SetStructuredPostalAddressNil

`func (o *YotiProviderOutput) SetStructuredPostalAddressNil(b bool)`

 SetStructuredPostalAddressNil sets the value for StructuredPostalAddress to be an explicit nil

### UnsetStructuredPostalAddress
`func (o *YotiProviderOutput) UnsetStructuredPostalAddress()`

UnsetStructuredPostalAddress ensures that no value is present for StructuredPostalAddress, not even an explicit nil
### GetPostalAddress

`func (o *YotiProviderOutput) GetPostalAddress() string`

GetPostalAddress returns the PostalAddress field if non-nil, zero value otherwise.

### GetPostalAddressOk

`func (o *YotiProviderOutput) GetPostalAddressOk() (*string, bool)`

GetPostalAddressOk returns a tuple with the PostalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalAddress

`func (o *YotiProviderOutput) SetPostalAddress(v string)`

SetPostalAddress sets PostalAddress field to given value.

### HasPostalAddress

`func (o *YotiProviderOutput) HasPostalAddress() bool`

HasPostalAddress returns a boolean if a field has been set.

### SetPostalAddressNil

`func (o *YotiProviderOutput) SetPostalAddressNil(b bool)`

 SetPostalAddressNil sets the value for PostalAddress to be an explicit nil

### UnsetPostalAddress
`func (o *YotiProviderOutput) UnsetPostalAddress()`

UnsetPostalAddress ensures that no value is present for PostalAddress, not even an explicit nil
### GetDocumentDetails

`func (o *YotiProviderOutput) GetDocumentDetails() YotiDocumentDetails`

GetDocumentDetails returns the DocumentDetails field if non-nil, zero value otherwise.

### GetDocumentDetailsOk

`func (o *YotiProviderOutput) GetDocumentDetailsOk() (*YotiDocumentDetails, bool)`

GetDocumentDetailsOk returns a tuple with the DocumentDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentDetails

`func (o *YotiProviderOutput) SetDocumentDetails(v YotiDocumentDetails)`

SetDocumentDetails sets DocumentDetails field to given value.

### HasDocumentDetails

`func (o *YotiProviderOutput) HasDocumentDetails() bool`

HasDocumentDetails returns a boolean if a field has been set.

### SetDocumentDetailsNil

`func (o *YotiProviderOutput) SetDocumentDetailsNil(b bool)`

 SetDocumentDetailsNil sets the value for DocumentDetails to be an explicit nil

### UnsetDocumentDetails
`func (o *YotiProviderOutput) UnsetDocumentDetails()`

UnsetDocumentDetails ensures that no value is present for DocumentDetails, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


