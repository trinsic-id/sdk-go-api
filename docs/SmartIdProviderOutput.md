# SmartIdProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GivenName** | Pointer to **NullableString** | The given name (first name) of the individual, extracted from the Smart ID authentication certificate&#39;s Subject Distinguished Name \&quot;G\&quot; (givenName) attribute. | [optional] 
**FamilyName** | Pointer to **NullableString** | The family name (surname) of the individual, extracted from the Smart ID authentication certificate&#39;s Subject Distinguished Name \&quot;SN\&quot; (surname) attribute. | [optional] 
**DateOfBirth** | Pointer to **NullableString** | The individual&#39;s date of birth, derived from the personal code.              NOTE: Newer Latvian personal codes starting with \&quot;32\&quot; do not contain date of birth information. | [optional] 
**Sex** | Pointer to **NullableString** | The individual&#39;s sex, derived from the first digit of the personal code.              NOTE: Not available for Latvian personal codes.              Possible values: - Male - Female | [optional] 
**Country** | Pointer to **NullableString** | The ISO 3166-1 alpha-2 country code extracted from the Smart ID authentication certificate&#39;s Subject Distinguished Name \&quot;C\&quot; (country) attribute.              This represents the country of the certificate issuer, not the person&#39;s nationality or citizenship. For Smart ID, this will be \&quot;EE\&quot; (Estonia), \&quot;LT\&quot; (Lithuania), or \&quot;LV\&quot; (Latvia). | [optional] 
**IdentityType** | Pointer to **NullableString** | The identity document type, extracted from the first 3 characters of the SERIALNUMBER field (before the country code).              Possible values: - PNO: Personal Number (national civic registration number) - PAS: Passport - IDC: National identity card | [optional] 
**PersonalCode** | Pointer to **NullableString** | The personal code (Estonian: isikukood, Lithuanian: asmens kodas, Latvian: personas kods) of the individual, extracted from the SERIALNUMBER field of the Smart ID authentication certificate. This is the portion after the identity type and country prefix (e.g., \&quot;48501010014\&quot; from \&quot;PNOEE-48501010014\&quot;).              Estonian and Lithuanian personal codes are 11 digits in the format GYYMMDDSSSC where: - G &#x3D; century/gender (3-4 &#x3D; 1900s, 5-6 &#x3D; 2000s; odd &#x3D; male, even &#x3D; female) - YYMMDD &#x3D; date of birth - SSS &#x3D; sequence number - C &#x3D; checksum digit              Latvian personal codes use the format DDMMYY-NNNNN (6 digits, a dash, and 5 digits).              See: - https://en.wikipedia.org/wiki/National_identification_number#Estonia - https://en.wikipedia.org/wiki/National_identification_number#Lithuania - https://en.wikipedia.org/wiki/National_identification_number#Latvia | [optional] 
**SerialNumber** | Pointer to **NullableString** | The SERIALNUMBER attribute from the Smart ID authentication certificate&#39;s Subject Distinguished Name. Format: \&quot;{identity-type}{country-code}-{identifier}\&quot;              Components: - identity-type (3 chars): PNO (Personal Number), PAS (Passport), IDC (ID Card) - country-code (2 chars): ISO 3166-1 alpha-2 (EE, LT, LV) - identifier: The personal code | [optional] 
**CertificateSubject** | Pointer to **NullableString** | The full Subject Distinguished Name (Subject DN) from the Smart ID authentication certificate. Contains comma-separated RDN (Relative Distinguished Name) components including C (Country), CN (Common Name), SN (Surname), G (Given Name), and SERIALNUMBER (Personal identifier). | [optional] 
**CertificateLevel** | Pointer to **NullableString** | The certificate level indicating the authentication assurance level.              Possible values: - QUALIFIED: Highest assurance level (eIDAS QES - Qualified Electronic Signature),   legally equivalent to a handwritten signature. This is the standard production level. - ADVANCED: Lower assurance level (eIDAS AdES), also called \&quot;Smart-ID Basic\&quot;.   Only available in test environments for test accounts. | [optional] 
**DocumentNumber** | Pointer to **NullableString** | A unique identifier for the Smart ID account/document used for this authentication. Format: PNO{CC}-{personal-code}-{device-id}-{qualification}              Components: - PNO &#x3D; Personal Number (identity type) - CC &#x3D; Country code (EE, LT, LV) - personal-code &#x3D; The individual&#39;s personal code - device-id &#x3D; Random 4-character device identifier - qualification &#x3D; NQ (non-qualified) or Q (qualified certificate)              This is a stable identifier that can be used to recognize returning users across sessions. | [optional] 
**InteractionFlowUsed** | Pointer to **NullableString** | Indicates which interaction flow was used during the Smart ID authentication process. This reflects which of the allowedInteractionsOrder options was actually used.              Possible values: - displayTextAndPIN: User saw text and entered PIN - verificationCodeChoice: User selected from multiple verification codes - confirmationMessage: User confirmed a message - confirmationMessageAndVerificationCodeChoice: Combination of both | [optional] 
**DeviceIpAddress** | Pointer to **NullableString** | The IP address of the device where the Smart ID app was used for authentication. Can be IPv4 or IPv6 format. | [optional] 

## Methods

### NewSmartIdProviderOutput

`func NewSmartIdProviderOutput() *SmartIdProviderOutput`

NewSmartIdProviderOutput instantiates a new SmartIdProviderOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmartIdProviderOutputWithDefaults

`func NewSmartIdProviderOutputWithDefaults() *SmartIdProviderOutput`

NewSmartIdProviderOutputWithDefaults instantiates a new SmartIdProviderOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGivenName

`func (o *SmartIdProviderOutput) GetGivenName() string`

GetGivenName returns the GivenName field if non-nil, zero value otherwise.

### GetGivenNameOk

`func (o *SmartIdProviderOutput) GetGivenNameOk() (*string, bool)`

GetGivenNameOk returns a tuple with the GivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenName

`func (o *SmartIdProviderOutput) SetGivenName(v string)`

SetGivenName sets GivenName field to given value.

### HasGivenName

`func (o *SmartIdProviderOutput) HasGivenName() bool`

HasGivenName returns a boolean if a field has been set.

### SetGivenNameNil

`func (o *SmartIdProviderOutput) SetGivenNameNil(b bool)`

 SetGivenNameNil sets the value for GivenName to be an explicit nil

### UnsetGivenName
`func (o *SmartIdProviderOutput) UnsetGivenName()`

UnsetGivenName ensures that no value is present for GivenName, not even an explicit nil
### GetFamilyName

`func (o *SmartIdProviderOutput) GetFamilyName() string`

GetFamilyName returns the FamilyName field if non-nil, zero value otherwise.

### GetFamilyNameOk

`func (o *SmartIdProviderOutput) GetFamilyNameOk() (*string, bool)`

GetFamilyNameOk returns a tuple with the FamilyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyName

`func (o *SmartIdProviderOutput) SetFamilyName(v string)`

SetFamilyName sets FamilyName field to given value.

### HasFamilyName

`func (o *SmartIdProviderOutput) HasFamilyName() bool`

HasFamilyName returns a boolean if a field has been set.

### SetFamilyNameNil

`func (o *SmartIdProviderOutput) SetFamilyNameNil(b bool)`

 SetFamilyNameNil sets the value for FamilyName to be an explicit nil

### UnsetFamilyName
`func (o *SmartIdProviderOutput) UnsetFamilyName()`

UnsetFamilyName ensures that no value is present for FamilyName, not even an explicit nil
### GetDateOfBirth

`func (o *SmartIdProviderOutput) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *SmartIdProviderOutput) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *SmartIdProviderOutput) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *SmartIdProviderOutput) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### SetDateOfBirthNil

`func (o *SmartIdProviderOutput) SetDateOfBirthNil(b bool)`

 SetDateOfBirthNil sets the value for DateOfBirth to be an explicit nil

### UnsetDateOfBirth
`func (o *SmartIdProviderOutput) UnsetDateOfBirth()`

UnsetDateOfBirth ensures that no value is present for DateOfBirth, not even an explicit nil
### GetSex

`func (o *SmartIdProviderOutput) GetSex() string`

GetSex returns the Sex field if non-nil, zero value otherwise.

### GetSexOk

`func (o *SmartIdProviderOutput) GetSexOk() (*string, bool)`

GetSexOk returns a tuple with the Sex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSex

`func (o *SmartIdProviderOutput) SetSex(v string)`

SetSex sets Sex field to given value.

### HasSex

`func (o *SmartIdProviderOutput) HasSex() bool`

HasSex returns a boolean if a field has been set.

### SetSexNil

`func (o *SmartIdProviderOutput) SetSexNil(b bool)`

 SetSexNil sets the value for Sex to be an explicit nil

### UnsetSex
`func (o *SmartIdProviderOutput) UnsetSex()`

UnsetSex ensures that no value is present for Sex, not even an explicit nil
### GetCountry

`func (o *SmartIdProviderOutput) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *SmartIdProviderOutput) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *SmartIdProviderOutput) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *SmartIdProviderOutput) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *SmartIdProviderOutput) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *SmartIdProviderOutput) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetIdentityType

`func (o *SmartIdProviderOutput) GetIdentityType() string`

GetIdentityType returns the IdentityType field if non-nil, zero value otherwise.

### GetIdentityTypeOk

`func (o *SmartIdProviderOutput) GetIdentityTypeOk() (*string, bool)`

GetIdentityTypeOk returns a tuple with the IdentityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityType

`func (o *SmartIdProviderOutput) SetIdentityType(v string)`

SetIdentityType sets IdentityType field to given value.

### HasIdentityType

`func (o *SmartIdProviderOutput) HasIdentityType() bool`

HasIdentityType returns a boolean if a field has been set.

### SetIdentityTypeNil

`func (o *SmartIdProviderOutput) SetIdentityTypeNil(b bool)`

 SetIdentityTypeNil sets the value for IdentityType to be an explicit nil

### UnsetIdentityType
`func (o *SmartIdProviderOutput) UnsetIdentityType()`

UnsetIdentityType ensures that no value is present for IdentityType, not even an explicit nil
### GetPersonalCode

`func (o *SmartIdProviderOutput) GetPersonalCode() string`

GetPersonalCode returns the PersonalCode field if non-nil, zero value otherwise.

### GetPersonalCodeOk

`func (o *SmartIdProviderOutput) GetPersonalCodeOk() (*string, bool)`

GetPersonalCodeOk returns a tuple with the PersonalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalCode

`func (o *SmartIdProviderOutput) SetPersonalCode(v string)`

SetPersonalCode sets PersonalCode field to given value.

### HasPersonalCode

`func (o *SmartIdProviderOutput) HasPersonalCode() bool`

HasPersonalCode returns a boolean if a field has been set.

### SetPersonalCodeNil

`func (o *SmartIdProviderOutput) SetPersonalCodeNil(b bool)`

 SetPersonalCodeNil sets the value for PersonalCode to be an explicit nil

### UnsetPersonalCode
`func (o *SmartIdProviderOutput) UnsetPersonalCode()`

UnsetPersonalCode ensures that no value is present for PersonalCode, not even an explicit nil
### GetSerialNumber

`func (o *SmartIdProviderOutput) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *SmartIdProviderOutput) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *SmartIdProviderOutput) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *SmartIdProviderOutput) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### SetSerialNumberNil

`func (o *SmartIdProviderOutput) SetSerialNumberNil(b bool)`

 SetSerialNumberNil sets the value for SerialNumber to be an explicit nil

### UnsetSerialNumber
`func (o *SmartIdProviderOutput) UnsetSerialNumber()`

UnsetSerialNumber ensures that no value is present for SerialNumber, not even an explicit nil
### GetCertificateSubject

`func (o *SmartIdProviderOutput) GetCertificateSubject() string`

GetCertificateSubject returns the CertificateSubject field if non-nil, zero value otherwise.

### GetCertificateSubjectOk

`func (o *SmartIdProviderOutput) GetCertificateSubjectOk() (*string, bool)`

GetCertificateSubjectOk returns a tuple with the CertificateSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateSubject

`func (o *SmartIdProviderOutput) SetCertificateSubject(v string)`

SetCertificateSubject sets CertificateSubject field to given value.

### HasCertificateSubject

`func (o *SmartIdProviderOutput) HasCertificateSubject() bool`

HasCertificateSubject returns a boolean if a field has been set.

### SetCertificateSubjectNil

`func (o *SmartIdProviderOutput) SetCertificateSubjectNil(b bool)`

 SetCertificateSubjectNil sets the value for CertificateSubject to be an explicit nil

### UnsetCertificateSubject
`func (o *SmartIdProviderOutput) UnsetCertificateSubject()`

UnsetCertificateSubject ensures that no value is present for CertificateSubject, not even an explicit nil
### GetCertificateLevel

`func (o *SmartIdProviderOutput) GetCertificateLevel() string`

GetCertificateLevel returns the CertificateLevel field if non-nil, zero value otherwise.

### GetCertificateLevelOk

`func (o *SmartIdProviderOutput) GetCertificateLevelOk() (*string, bool)`

GetCertificateLevelOk returns a tuple with the CertificateLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateLevel

`func (o *SmartIdProviderOutput) SetCertificateLevel(v string)`

SetCertificateLevel sets CertificateLevel field to given value.

### HasCertificateLevel

`func (o *SmartIdProviderOutput) HasCertificateLevel() bool`

HasCertificateLevel returns a boolean if a field has been set.

### SetCertificateLevelNil

`func (o *SmartIdProviderOutput) SetCertificateLevelNil(b bool)`

 SetCertificateLevelNil sets the value for CertificateLevel to be an explicit nil

### UnsetCertificateLevel
`func (o *SmartIdProviderOutput) UnsetCertificateLevel()`

UnsetCertificateLevel ensures that no value is present for CertificateLevel, not even an explicit nil
### GetDocumentNumber

`func (o *SmartIdProviderOutput) GetDocumentNumber() string`

GetDocumentNumber returns the DocumentNumber field if non-nil, zero value otherwise.

### GetDocumentNumberOk

`func (o *SmartIdProviderOutput) GetDocumentNumberOk() (*string, bool)`

GetDocumentNumberOk returns a tuple with the DocumentNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentNumber

`func (o *SmartIdProviderOutput) SetDocumentNumber(v string)`

SetDocumentNumber sets DocumentNumber field to given value.

### HasDocumentNumber

`func (o *SmartIdProviderOutput) HasDocumentNumber() bool`

HasDocumentNumber returns a boolean if a field has been set.

### SetDocumentNumberNil

`func (o *SmartIdProviderOutput) SetDocumentNumberNil(b bool)`

 SetDocumentNumberNil sets the value for DocumentNumber to be an explicit nil

### UnsetDocumentNumber
`func (o *SmartIdProviderOutput) UnsetDocumentNumber()`

UnsetDocumentNumber ensures that no value is present for DocumentNumber, not even an explicit nil
### GetInteractionFlowUsed

`func (o *SmartIdProviderOutput) GetInteractionFlowUsed() string`

GetInteractionFlowUsed returns the InteractionFlowUsed field if non-nil, zero value otherwise.

### GetInteractionFlowUsedOk

`func (o *SmartIdProviderOutput) GetInteractionFlowUsedOk() (*string, bool)`

GetInteractionFlowUsedOk returns a tuple with the InteractionFlowUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInteractionFlowUsed

`func (o *SmartIdProviderOutput) SetInteractionFlowUsed(v string)`

SetInteractionFlowUsed sets InteractionFlowUsed field to given value.

### HasInteractionFlowUsed

`func (o *SmartIdProviderOutput) HasInteractionFlowUsed() bool`

HasInteractionFlowUsed returns a boolean if a field has been set.

### SetInteractionFlowUsedNil

`func (o *SmartIdProviderOutput) SetInteractionFlowUsedNil(b bool)`

 SetInteractionFlowUsedNil sets the value for InteractionFlowUsed to be an explicit nil

### UnsetInteractionFlowUsed
`func (o *SmartIdProviderOutput) UnsetInteractionFlowUsed()`

UnsetInteractionFlowUsed ensures that no value is present for InteractionFlowUsed, not even an explicit nil
### GetDeviceIpAddress

`func (o *SmartIdProviderOutput) GetDeviceIpAddress() string`

GetDeviceIpAddress returns the DeviceIpAddress field if non-nil, zero value otherwise.

### GetDeviceIpAddressOk

`func (o *SmartIdProviderOutput) GetDeviceIpAddressOk() (*string, bool)`

GetDeviceIpAddressOk returns a tuple with the DeviceIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceIpAddress

`func (o *SmartIdProviderOutput) SetDeviceIpAddress(v string)`

SetDeviceIpAddress sets DeviceIpAddress field to given value.

### HasDeviceIpAddress

`func (o *SmartIdProviderOutput) HasDeviceIpAddress() bool`

HasDeviceIpAddress returns a boolean if a field has been set.

### SetDeviceIpAddressNil

`func (o *SmartIdProviderOutput) SetDeviceIpAddressNil(b bool)`

 SetDeviceIpAddressNil sets the value for DeviceIpAddress to be an explicit nil

### UnsetDeviceIpAddress
`func (o *SmartIdProviderOutput) UnsetDeviceIpAddress()`

UnsetDeviceIpAddress ensures that no value is present for DeviceIpAddress, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


