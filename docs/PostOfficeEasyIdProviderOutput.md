# PostOfficeEasyIdProviderOutput

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

### NewPostOfficeEasyIdProviderOutput

`func NewPostOfficeEasyIdProviderOutput() *PostOfficeEasyIdProviderOutput`

NewPostOfficeEasyIdProviderOutput instantiates a new PostOfficeEasyIdProviderOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostOfficeEasyIdProviderOutputWithDefaults

`func NewPostOfficeEasyIdProviderOutputWithDefaults() *PostOfficeEasyIdProviderOutput`

NewPostOfficeEasyIdProviderOutputWithDefaults instantiates a new PostOfficeEasyIdProviderOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRememberMeId

`func (o *PostOfficeEasyIdProviderOutput) GetRememberMeId() string`

GetRememberMeId returns the RememberMeId field if non-nil, zero value otherwise.

### GetRememberMeIdOk

`func (o *PostOfficeEasyIdProviderOutput) GetRememberMeIdOk() (*string, bool)`

GetRememberMeIdOk returns a tuple with the RememberMeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRememberMeId

`func (o *PostOfficeEasyIdProviderOutput) SetRememberMeId(v string)`

SetRememberMeId sets RememberMeId field to given value.

### HasRememberMeId

`func (o *PostOfficeEasyIdProviderOutput) HasRememberMeId() bool`

HasRememberMeId returns a boolean if a field has been set.

### SetRememberMeIdNil

`func (o *PostOfficeEasyIdProviderOutput) SetRememberMeIdNil(b bool)`

 SetRememberMeIdNil sets the value for RememberMeId to be an explicit nil

### UnsetRememberMeId
`func (o *PostOfficeEasyIdProviderOutput) UnsetRememberMeId()`

UnsetRememberMeId ensures that no value is present for RememberMeId, not even an explicit nil
### GetEmail

`func (o *PostOfficeEasyIdProviderOutput) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PostOfficeEasyIdProviderOutput) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PostOfficeEasyIdProviderOutput) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *PostOfficeEasyIdProviderOutput) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *PostOfficeEasyIdProviderOutput) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *PostOfficeEasyIdProviderOutput) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetGivenName

`func (o *PostOfficeEasyIdProviderOutput) GetGivenName() string`

GetGivenName returns the GivenName field if non-nil, zero value otherwise.

### GetGivenNameOk

`func (o *PostOfficeEasyIdProviderOutput) GetGivenNameOk() (*string, bool)`

GetGivenNameOk returns a tuple with the GivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenName

`func (o *PostOfficeEasyIdProviderOutput) SetGivenName(v string)`

SetGivenName sets GivenName field to given value.

### HasGivenName

`func (o *PostOfficeEasyIdProviderOutput) HasGivenName() bool`

HasGivenName returns a boolean if a field has been set.

### SetGivenNameNil

`func (o *PostOfficeEasyIdProviderOutput) SetGivenNameNil(b bool)`

 SetGivenNameNil sets the value for GivenName to be an explicit nil

### UnsetGivenName
`func (o *PostOfficeEasyIdProviderOutput) UnsetGivenName()`

UnsetGivenName ensures that no value is present for GivenName, not even an explicit nil
### GetFamilyName

`func (o *PostOfficeEasyIdProviderOutput) GetFamilyName() string`

GetFamilyName returns the FamilyName field if non-nil, zero value otherwise.

### GetFamilyNameOk

`func (o *PostOfficeEasyIdProviderOutput) GetFamilyNameOk() (*string, bool)`

GetFamilyNameOk returns a tuple with the FamilyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyName

`func (o *PostOfficeEasyIdProviderOutput) SetFamilyName(v string)`

SetFamilyName sets FamilyName field to given value.

### HasFamilyName

`func (o *PostOfficeEasyIdProviderOutput) HasFamilyName() bool`

HasFamilyName returns a boolean if a field has been set.

### SetFamilyNameNil

`func (o *PostOfficeEasyIdProviderOutput) SetFamilyNameNil(b bool)`

 SetFamilyNameNil sets the value for FamilyName to be an explicit nil

### UnsetFamilyName
`func (o *PostOfficeEasyIdProviderOutput) UnsetFamilyName()`

UnsetFamilyName ensures that no value is present for FamilyName, not even an explicit nil
### GetFullName

`func (o *PostOfficeEasyIdProviderOutput) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *PostOfficeEasyIdProviderOutput) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *PostOfficeEasyIdProviderOutput) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *PostOfficeEasyIdProviderOutput) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### SetFullNameNil

`func (o *PostOfficeEasyIdProviderOutput) SetFullNameNil(b bool)`

 SetFullNameNil sets the value for FullName to be an explicit nil

### UnsetFullName
`func (o *PostOfficeEasyIdProviderOutput) UnsetFullName()`

UnsetFullName ensures that no value is present for FullName, not even an explicit nil
### GetDateOfBirth

`func (o *PostOfficeEasyIdProviderOutput) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *PostOfficeEasyIdProviderOutput) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *PostOfficeEasyIdProviderOutput) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *PostOfficeEasyIdProviderOutput) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### SetDateOfBirthNil

`func (o *PostOfficeEasyIdProviderOutput) SetDateOfBirthNil(b bool)`

 SetDateOfBirthNil sets the value for DateOfBirth to be an explicit nil

### UnsetDateOfBirth
`func (o *PostOfficeEasyIdProviderOutput) UnsetDateOfBirth()`

UnsetDateOfBirth ensures that no value is present for DateOfBirth, not even an explicit nil
### GetGender

`func (o *PostOfficeEasyIdProviderOutput) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *PostOfficeEasyIdProviderOutput) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *PostOfficeEasyIdProviderOutput) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *PostOfficeEasyIdProviderOutput) HasGender() bool`

HasGender returns a boolean if a field has been set.

### SetGenderNil

`func (o *PostOfficeEasyIdProviderOutput) SetGenderNil(b bool)`

 SetGenderNil sets the value for Gender to be an explicit nil

### UnsetGender
`func (o *PostOfficeEasyIdProviderOutput) UnsetGender()`

UnsetGender ensures that no value is present for Gender, not even an explicit nil
### GetNationality

`func (o *PostOfficeEasyIdProviderOutput) GetNationality() string`

GetNationality returns the Nationality field if non-nil, zero value otherwise.

### GetNationalityOk

`func (o *PostOfficeEasyIdProviderOutput) GetNationalityOk() (*string, bool)`

GetNationalityOk returns a tuple with the Nationality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationality

`func (o *PostOfficeEasyIdProviderOutput) SetNationality(v string)`

SetNationality sets Nationality field to given value.

### HasNationality

`func (o *PostOfficeEasyIdProviderOutput) HasNationality() bool`

HasNationality returns a boolean if a field has been set.

### SetNationalityNil

`func (o *PostOfficeEasyIdProviderOutput) SetNationalityNil(b bool)`

 SetNationalityNil sets the value for Nationality to be an explicit nil

### UnsetNationality
`func (o *PostOfficeEasyIdProviderOutput) UnsetNationality()`

UnsetNationality ensures that no value is present for Nationality, not even an explicit nil
### GetMobileNumber

`func (o *PostOfficeEasyIdProviderOutput) GetMobileNumber() string`

GetMobileNumber returns the MobileNumber field if non-nil, zero value otherwise.

### GetMobileNumberOk

`func (o *PostOfficeEasyIdProviderOutput) GetMobileNumberOk() (*string, bool)`

GetMobileNumberOk returns a tuple with the MobileNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileNumber

`func (o *PostOfficeEasyIdProviderOutput) SetMobileNumber(v string)`

SetMobileNumber sets MobileNumber field to given value.

### HasMobileNumber

`func (o *PostOfficeEasyIdProviderOutput) HasMobileNumber() bool`

HasMobileNumber returns a boolean if a field has been set.

### SetMobileNumberNil

`func (o *PostOfficeEasyIdProviderOutput) SetMobileNumberNil(b bool)`

 SetMobileNumberNil sets the value for MobileNumber to be an explicit nil

### UnsetMobileNumber
`func (o *PostOfficeEasyIdProviderOutput) UnsetMobileNumber()`

UnsetMobileNumber ensures that no value is present for MobileNumber, not even an explicit nil
### GetStructuredPostalAddress

`func (o *PostOfficeEasyIdProviderOutput) GetStructuredPostalAddress() YotiStructuredPostalAddress`

GetStructuredPostalAddress returns the StructuredPostalAddress field if non-nil, zero value otherwise.

### GetStructuredPostalAddressOk

`func (o *PostOfficeEasyIdProviderOutput) GetStructuredPostalAddressOk() (*YotiStructuredPostalAddress, bool)`

GetStructuredPostalAddressOk returns a tuple with the StructuredPostalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructuredPostalAddress

`func (o *PostOfficeEasyIdProviderOutput) SetStructuredPostalAddress(v YotiStructuredPostalAddress)`

SetStructuredPostalAddress sets StructuredPostalAddress field to given value.

### HasStructuredPostalAddress

`func (o *PostOfficeEasyIdProviderOutput) HasStructuredPostalAddress() bool`

HasStructuredPostalAddress returns a boolean if a field has been set.

### SetStructuredPostalAddressNil

`func (o *PostOfficeEasyIdProviderOutput) SetStructuredPostalAddressNil(b bool)`

 SetStructuredPostalAddressNil sets the value for StructuredPostalAddress to be an explicit nil

### UnsetStructuredPostalAddress
`func (o *PostOfficeEasyIdProviderOutput) UnsetStructuredPostalAddress()`

UnsetStructuredPostalAddress ensures that no value is present for StructuredPostalAddress, not even an explicit nil
### GetPostalAddress

`func (o *PostOfficeEasyIdProviderOutput) GetPostalAddress() string`

GetPostalAddress returns the PostalAddress field if non-nil, zero value otherwise.

### GetPostalAddressOk

`func (o *PostOfficeEasyIdProviderOutput) GetPostalAddressOk() (*string, bool)`

GetPostalAddressOk returns a tuple with the PostalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalAddress

`func (o *PostOfficeEasyIdProviderOutput) SetPostalAddress(v string)`

SetPostalAddress sets PostalAddress field to given value.

### HasPostalAddress

`func (o *PostOfficeEasyIdProviderOutput) HasPostalAddress() bool`

HasPostalAddress returns a boolean if a field has been set.

### SetPostalAddressNil

`func (o *PostOfficeEasyIdProviderOutput) SetPostalAddressNil(b bool)`

 SetPostalAddressNil sets the value for PostalAddress to be an explicit nil

### UnsetPostalAddress
`func (o *PostOfficeEasyIdProviderOutput) UnsetPostalAddress()`

UnsetPostalAddress ensures that no value is present for PostalAddress, not even an explicit nil
### GetDocumentDetails

`func (o *PostOfficeEasyIdProviderOutput) GetDocumentDetails() YotiDocumentDetails`

GetDocumentDetails returns the DocumentDetails field if non-nil, zero value otherwise.

### GetDocumentDetailsOk

`func (o *PostOfficeEasyIdProviderOutput) GetDocumentDetailsOk() (*YotiDocumentDetails, bool)`

GetDocumentDetailsOk returns a tuple with the DocumentDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentDetails

`func (o *PostOfficeEasyIdProviderOutput) SetDocumentDetails(v YotiDocumentDetails)`

SetDocumentDetails sets DocumentDetails field to given value.

### HasDocumentDetails

`func (o *PostOfficeEasyIdProviderOutput) HasDocumentDetails() bool`

HasDocumentDetails returns a boolean if a field has been set.

### SetDocumentDetailsNil

`func (o *PostOfficeEasyIdProviderOutput) SetDocumentDetailsNil(b bool)`

 SetDocumentDetailsNil sets the value for DocumentDetails to be an explicit nil

### UnsetDocumentDetails
`func (o *PostOfficeEasyIdProviderOutput) UnsetDocumentDetails()`

UnsetDocumentDetails ensures that no value is present for DocumentDetails, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


