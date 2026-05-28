# Iso180135StandardNamespaceOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GivenName** | Pointer to **NullableString** | First name(s), other name(s), or secondary identifier of the individual. | [optional] 
**GivenNameNationalCharacter** | Pointer to **NullableString** | The given name of the individual using the full UTF-8 character set. | [optional] 
**FamilyName** | Pointer to **NullableString** | Last name, surname, or primary identifier of the individual. | [optional] 
**FamilyNameNationalCharacter** | Pointer to **NullableString** | The family name of the individual using the full UTF-8 character set. | [optional] 
**DateOfBirth** | Pointer to **NullableString** | The date of birth of the individual. | [optional] 
**IssueDate** | Pointer to **NullableString** | The date when the mDL was issued.              This date marks the beginning of the Administrative Validity Period of the mDL, which is usually (but not necessarily always) the same as the issue date of the underlying physical document used to create the mDL, if one exists. | [optional] 
**ExpiryDate** | Pointer to **NullableString** | The date when the mDL expires.              This date marks the end of the Administrative Validity Period of the mDL, which is usually (but not necessarily always) the same as the expiration date of the underlying physical document used to create the mDL, if one exists. | [optional] 
**IssuingCountry** | Pointer to **NullableString** | Alpha-2 country code of the issuing authority&#39;s country or territory. | [optional] 
**IssuingAuthority** | Pointer to **NullableString** | Name or identifier of the mDL issuing authority.              This field&#39;s contents are arbitrary; it has no guaranteed format. | [optional] 
**IssuingJurisdiction** | Pointer to **NullableString** | ISO 3166-2 country-subdivision code for the jurisdiction that issued the mDL. | [optional] 
**Nationality** | Pointer to **NullableString** | Nationality of the individual as an ISO 3166-1 alpha-2 country code. | [optional] 
**DocumentNumber** | Pointer to **NullableString** | The document number assigned to the mDL.              This is typically the same as the document number of the underlying physical document used to create the mDL, if one exists. | [optional] 
**DrivingPrivileges** | Pointer to [**[]Iso180135DrivingPrivilege**](Iso180135DrivingPrivilege.md) | Driving privileges of the individual, parsed per ISO 18013-5. | [optional] 
**UnDistinguishingSign** | Pointer to **NullableString** | The UN Distinguishing sign of the issuing country according to ISO/IEC 18013-1. | [optional] 
**AdministrativeNumber** | Pointer to **NullableString** | Audit control number assigned by the issuing authority.              The meaning and semantics of this field are defined by the issuing authority in question. | [optional] 
**Sex** | Pointer to **NullableInt32** | The individual&#39;s sex as an ISO/IEC 5218 code.              Possible values:              - 0: Unknown - 1: Male - 2: Female - 9: Not Applicable | [optional] 
**HeightCentimeters** | Pointer to **NullableInt32** | The individual&#39;s height in centimeters. | [optional] 
**WeightKilograms** | Pointer to **NullableInt32** | The individual&#39;s weight in kilograms. | [optional] 
**EyeColor** | Pointer to **NullableString** | The individual&#39;s eye color.              Possible values:              - \&quot;black\&quot; - \&quot;blue\&quot; - \&quot;brown\&quot; - \&quot;dichromatic\&quot; - \&quot;grey\&quot; - \&quot;green\&quot; - \&quot;hazel\&quot; - \&quot;maroon\&quot; - \&quot;pink\&quot; - \&quot;unknown\&quot; | [optional] 
**HairColor** | Pointer to **NullableString** | The individual&#39;s hair color.              Possible values: - \&quot;bald\&quot; - \&quot;black\&quot; - \&quot;blond\&quot; - \&quot;brown\&quot; - \&quot;grey\&quot; - \&quot;red\&quot; - \&quot;auburn\&quot; - \&quot;sandy\&quot; - \&quot;white\&quot; - \&quot;unknown\&quot; | [optional] 
**BirthPlace** | Pointer to **NullableString** | Country and municipality or state/province where the individual was born. | [optional] 
**ResidentAddress** | Pointer to **NullableString** | Address where the individual resides.              The exact format of this field is variable and depends on the issuer. It may map exactly to a \&quot;Line1\&quot;, or it may include city, state, and/or zip code as well. | [optional] 
**ResidentCity** | Pointer to **NullableString** | City where the individual resides. | [optional] 
**ResidentState** | Pointer to **NullableString** | State, province, or district where the individual resides. | [optional] 
**ResidentPostalCode** | Pointer to **NullableString** | Postal code where the individual resides. | [optional] 
**ResidentCountry** | Pointer to **NullableString** | Country where the individual resides as an ISO 3166-1 alpha-2 country code. | [optional] 
**PortraitCaptureDate** | Pointer to **NullableTime** | Date and time when the portrait image was captured. | [optional] 
**AgeInYears** | Pointer to **NullableInt32** | Age of the individual in years. | [optional] 
**AgeBirthYear** | Pointer to **NullableInt32** | Year of birth of the individual. | [optional] 
**AgeOver** | Pointer to [**[]AgeOverOutput**](AgeOverOutput.md) | Processed age-over claims returned by the mDL. | [optional] 
**BiometricTemplateFace** | Pointer to [**NullableIso180132BiometricGroupTemplate**](Iso180132BiometricGroupTemplate.md) | Facial biometric template group from &#x60;biometric_template_face&#x60;, parsed per ISO 18013-2 Annex C. | [optional] 
**BiometricTemplateVoice** | Pointer to [**NullableIso180132BiometricGroupTemplate**](Iso180132BiometricGroupTemplate.md) | Vocal biometric template group from &#x60;biometric_template_voice&#x60;, parsed per ISO 18013-2 Annex C. | [optional] 
**BiometricTemplateFinger** | Pointer to [**NullableIso180132BiometricGroupTemplate**](Iso180132BiometricGroupTemplate.md) | Fingerprint biometric template group from &#x60;biometric_template_finger&#x60;, parsed per ISO 18013-2 Annex C. | [optional] 
**BiometricTemplateIris** | Pointer to [**NullableIso180132BiometricGroupTemplate**](Iso180132BiometricGroupTemplate.md) | Iris biometric template group from &#x60;biometric_template_iris&#x60;, parsed per ISO 18013-2 Annex C. | [optional] 
**BiometricTemplateRetina** | Pointer to [**NullableIso180132BiometricGroupTemplate**](Iso180132BiometricGroupTemplate.md) | Retinal biometric template group from &#x60;biometric_template_retina&#x60;, parsed per ISO 18013-2 Annex C. | [optional] 

## Methods

### NewIso180135StandardNamespaceOutput

`func NewIso180135StandardNamespaceOutput() *Iso180135StandardNamespaceOutput`

NewIso180135StandardNamespaceOutput instantiates a new Iso180135StandardNamespaceOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIso180135StandardNamespaceOutputWithDefaults

`func NewIso180135StandardNamespaceOutputWithDefaults() *Iso180135StandardNamespaceOutput`

NewIso180135StandardNamespaceOutputWithDefaults instantiates a new Iso180135StandardNamespaceOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGivenName

`func (o *Iso180135StandardNamespaceOutput) GetGivenName() string`

GetGivenName returns the GivenName field if non-nil, zero value otherwise.

### GetGivenNameOk

`func (o *Iso180135StandardNamespaceOutput) GetGivenNameOk() (*string, bool)`

GetGivenNameOk returns a tuple with the GivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenName

`func (o *Iso180135StandardNamespaceOutput) SetGivenName(v string)`

SetGivenName sets GivenName field to given value.

### HasGivenName

`func (o *Iso180135StandardNamespaceOutput) HasGivenName() bool`

HasGivenName returns a boolean if a field has been set.

### SetGivenNameNil

`func (o *Iso180135StandardNamespaceOutput) SetGivenNameNil(b bool)`

 SetGivenNameNil sets the value for GivenName to be an explicit nil

### UnsetGivenName
`func (o *Iso180135StandardNamespaceOutput) UnsetGivenName()`

UnsetGivenName ensures that no value is present for GivenName, not even an explicit nil
### GetGivenNameNationalCharacter

`func (o *Iso180135StandardNamespaceOutput) GetGivenNameNationalCharacter() string`

GetGivenNameNationalCharacter returns the GivenNameNationalCharacter field if non-nil, zero value otherwise.

### GetGivenNameNationalCharacterOk

`func (o *Iso180135StandardNamespaceOutput) GetGivenNameNationalCharacterOk() (*string, bool)`

GetGivenNameNationalCharacterOk returns a tuple with the GivenNameNationalCharacter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenNameNationalCharacter

`func (o *Iso180135StandardNamespaceOutput) SetGivenNameNationalCharacter(v string)`

SetGivenNameNationalCharacter sets GivenNameNationalCharacter field to given value.

### HasGivenNameNationalCharacter

`func (o *Iso180135StandardNamespaceOutput) HasGivenNameNationalCharacter() bool`

HasGivenNameNationalCharacter returns a boolean if a field has been set.

### SetGivenNameNationalCharacterNil

`func (o *Iso180135StandardNamespaceOutput) SetGivenNameNationalCharacterNil(b bool)`

 SetGivenNameNationalCharacterNil sets the value for GivenNameNationalCharacter to be an explicit nil

### UnsetGivenNameNationalCharacter
`func (o *Iso180135StandardNamespaceOutput) UnsetGivenNameNationalCharacter()`

UnsetGivenNameNationalCharacter ensures that no value is present for GivenNameNationalCharacter, not even an explicit nil
### GetFamilyName

`func (o *Iso180135StandardNamespaceOutput) GetFamilyName() string`

GetFamilyName returns the FamilyName field if non-nil, zero value otherwise.

### GetFamilyNameOk

`func (o *Iso180135StandardNamespaceOutput) GetFamilyNameOk() (*string, bool)`

GetFamilyNameOk returns a tuple with the FamilyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyName

`func (o *Iso180135StandardNamespaceOutput) SetFamilyName(v string)`

SetFamilyName sets FamilyName field to given value.

### HasFamilyName

`func (o *Iso180135StandardNamespaceOutput) HasFamilyName() bool`

HasFamilyName returns a boolean if a field has been set.

### SetFamilyNameNil

`func (o *Iso180135StandardNamespaceOutput) SetFamilyNameNil(b bool)`

 SetFamilyNameNil sets the value for FamilyName to be an explicit nil

### UnsetFamilyName
`func (o *Iso180135StandardNamespaceOutput) UnsetFamilyName()`

UnsetFamilyName ensures that no value is present for FamilyName, not even an explicit nil
### GetFamilyNameNationalCharacter

`func (o *Iso180135StandardNamespaceOutput) GetFamilyNameNationalCharacter() string`

GetFamilyNameNationalCharacter returns the FamilyNameNationalCharacter field if non-nil, zero value otherwise.

### GetFamilyNameNationalCharacterOk

`func (o *Iso180135StandardNamespaceOutput) GetFamilyNameNationalCharacterOk() (*string, bool)`

GetFamilyNameNationalCharacterOk returns a tuple with the FamilyNameNationalCharacter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyNameNationalCharacter

`func (o *Iso180135StandardNamespaceOutput) SetFamilyNameNationalCharacter(v string)`

SetFamilyNameNationalCharacter sets FamilyNameNationalCharacter field to given value.

### HasFamilyNameNationalCharacter

`func (o *Iso180135StandardNamespaceOutput) HasFamilyNameNationalCharacter() bool`

HasFamilyNameNationalCharacter returns a boolean if a field has been set.

### SetFamilyNameNationalCharacterNil

`func (o *Iso180135StandardNamespaceOutput) SetFamilyNameNationalCharacterNil(b bool)`

 SetFamilyNameNationalCharacterNil sets the value for FamilyNameNationalCharacter to be an explicit nil

### UnsetFamilyNameNationalCharacter
`func (o *Iso180135StandardNamespaceOutput) UnsetFamilyNameNationalCharacter()`

UnsetFamilyNameNationalCharacter ensures that no value is present for FamilyNameNationalCharacter, not even an explicit nil
### GetDateOfBirth

`func (o *Iso180135StandardNamespaceOutput) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *Iso180135StandardNamespaceOutput) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *Iso180135StandardNamespaceOutput) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *Iso180135StandardNamespaceOutput) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### SetDateOfBirthNil

`func (o *Iso180135StandardNamespaceOutput) SetDateOfBirthNil(b bool)`

 SetDateOfBirthNil sets the value for DateOfBirth to be an explicit nil

### UnsetDateOfBirth
`func (o *Iso180135StandardNamespaceOutput) UnsetDateOfBirth()`

UnsetDateOfBirth ensures that no value is present for DateOfBirth, not even an explicit nil
### GetIssueDate

`func (o *Iso180135StandardNamespaceOutput) GetIssueDate() string`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *Iso180135StandardNamespaceOutput) GetIssueDateOk() (*string, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *Iso180135StandardNamespaceOutput) SetIssueDate(v string)`

SetIssueDate sets IssueDate field to given value.

### HasIssueDate

`func (o *Iso180135StandardNamespaceOutput) HasIssueDate() bool`

HasIssueDate returns a boolean if a field has been set.

### SetIssueDateNil

`func (o *Iso180135StandardNamespaceOutput) SetIssueDateNil(b bool)`

 SetIssueDateNil sets the value for IssueDate to be an explicit nil

### UnsetIssueDate
`func (o *Iso180135StandardNamespaceOutput) UnsetIssueDate()`

UnsetIssueDate ensures that no value is present for IssueDate, not even an explicit nil
### GetExpiryDate

`func (o *Iso180135StandardNamespaceOutput) GetExpiryDate() string`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *Iso180135StandardNamespaceOutput) GetExpiryDateOk() (*string, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *Iso180135StandardNamespaceOutput) SetExpiryDate(v string)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *Iso180135StandardNamespaceOutput) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### SetExpiryDateNil

`func (o *Iso180135StandardNamespaceOutput) SetExpiryDateNil(b bool)`

 SetExpiryDateNil sets the value for ExpiryDate to be an explicit nil

### UnsetExpiryDate
`func (o *Iso180135StandardNamespaceOutput) UnsetExpiryDate()`

UnsetExpiryDate ensures that no value is present for ExpiryDate, not even an explicit nil
### GetIssuingCountry

`func (o *Iso180135StandardNamespaceOutput) GetIssuingCountry() string`

GetIssuingCountry returns the IssuingCountry field if non-nil, zero value otherwise.

### GetIssuingCountryOk

`func (o *Iso180135StandardNamespaceOutput) GetIssuingCountryOk() (*string, bool)`

GetIssuingCountryOk returns a tuple with the IssuingCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingCountry

`func (o *Iso180135StandardNamespaceOutput) SetIssuingCountry(v string)`

SetIssuingCountry sets IssuingCountry field to given value.

### HasIssuingCountry

`func (o *Iso180135StandardNamespaceOutput) HasIssuingCountry() bool`

HasIssuingCountry returns a boolean if a field has been set.

### SetIssuingCountryNil

`func (o *Iso180135StandardNamespaceOutput) SetIssuingCountryNil(b bool)`

 SetIssuingCountryNil sets the value for IssuingCountry to be an explicit nil

### UnsetIssuingCountry
`func (o *Iso180135StandardNamespaceOutput) UnsetIssuingCountry()`

UnsetIssuingCountry ensures that no value is present for IssuingCountry, not even an explicit nil
### GetIssuingAuthority

`func (o *Iso180135StandardNamespaceOutput) GetIssuingAuthority() string`

GetIssuingAuthority returns the IssuingAuthority field if non-nil, zero value otherwise.

### GetIssuingAuthorityOk

`func (o *Iso180135StandardNamespaceOutput) GetIssuingAuthorityOk() (*string, bool)`

GetIssuingAuthorityOk returns a tuple with the IssuingAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingAuthority

`func (o *Iso180135StandardNamespaceOutput) SetIssuingAuthority(v string)`

SetIssuingAuthority sets IssuingAuthority field to given value.

### HasIssuingAuthority

`func (o *Iso180135StandardNamespaceOutput) HasIssuingAuthority() bool`

HasIssuingAuthority returns a boolean if a field has been set.

### SetIssuingAuthorityNil

`func (o *Iso180135StandardNamespaceOutput) SetIssuingAuthorityNil(b bool)`

 SetIssuingAuthorityNil sets the value for IssuingAuthority to be an explicit nil

### UnsetIssuingAuthority
`func (o *Iso180135StandardNamespaceOutput) UnsetIssuingAuthority()`

UnsetIssuingAuthority ensures that no value is present for IssuingAuthority, not even an explicit nil
### GetIssuingJurisdiction

`func (o *Iso180135StandardNamespaceOutput) GetIssuingJurisdiction() string`

GetIssuingJurisdiction returns the IssuingJurisdiction field if non-nil, zero value otherwise.

### GetIssuingJurisdictionOk

`func (o *Iso180135StandardNamespaceOutput) GetIssuingJurisdictionOk() (*string, bool)`

GetIssuingJurisdictionOk returns a tuple with the IssuingJurisdiction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingJurisdiction

`func (o *Iso180135StandardNamespaceOutput) SetIssuingJurisdiction(v string)`

SetIssuingJurisdiction sets IssuingJurisdiction field to given value.

### HasIssuingJurisdiction

`func (o *Iso180135StandardNamespaceOutput) HasIssuingJurisdiction() bool`

HasIssuingJurisdiction returns a boolean if a field has been set.

### SetIssuingJurisdictionNil

`func (o *Iso180135StandardNamespaceOutput) SetIssuingJurisdictionNil(b bool)`

 SetIssuingJurisdictionNil sets the value for IssuingJurisdiction to be an explicit nil

### UnsetIssuingJurisdiction
`func (o *Iso180135StandardNamespaceOutput) UnsetIssuingJurisdiction()`

UnsetIssuingJurisdiction ensures that no value is present for IssuingJurisdiction, not even an explicit nil
### GetNationality

`func (o *Iso180135StandardNamespaceOutput) GetNationality() string`

GetNationality returns the Nationality field if non-nil, zero value otherwise.

### GetNationalityOk

`func (o *Iso180135StandardNamespaceOutput) GetNationalityOk() (*string, bool)`

GetNationalityOk returns a tuple with the Nationality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationality

`func (o *Iso180135StandardNamespaceOutput) SetNationality(v string)`

SetNationality sets Nationality field to given value.

### HasNationality

`func (o *Iso180135StandardNamespaceOutput) HasNationality() bool`

HasNationality returns a boolean if a field has been set.

### SetNationalityNil

`func (o *Iso180135StandardNamespaceOutput) SetNationalityNil(b bool)`

 SetNationalityNil sets the value for Nationality to be an explicit nil

### UnsetNationality
`func (o *Iso180135StandardNamespaceOutput) UnsetNationality()`

UnsetNationality ensures that no value is present for Nationality, not even an explicit nil
### GetDocumentNumber

`func (o *Iso180135StandardNamespaceOutput) GetDocumentNumber() string`

GetDocumentNumber returns the DocumentNumber field if non-nil, zero value otherwise.

### GetDocumentNumberOk

`func (o *Iso180135StandardNamespaceOutput) GetDocumentNumberOk() (*string, bool)`

GetDocumentNumberOk returns a tuple with the DocumentNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentNumber

`func (o *Iso180135StandardNamespaceOutput) SetDocumentNumber(v string)`

SetDocumentNumber sets DocumentNumber field to given value.

### HasDocumentNumber

`func (o *Iso180135StandardNamespaceOutput) HasDocumentNumber() bool`

HasDocumentNumber returns a boolean if a field has been set.

### SetDocumentNumberNil

`func (o *Iso180135StandardNamespaceOutput) SetDocumentNumberNil(b bool)`

 SetDocumentNumberNil sets the value for DocumentNumber to be an explicit nil

### UnsetDocumentNumber
`func (o *Iso180135StandardNamespaceOutput) UnsetDocumentNumber()`

UnsetDocumentNumber ensures that no value is present for DocumentNumber, not even an explicit nil
### GetDrivingPrivileges

`func (o *Iso180135StandardNamespaceOutput) GetDrivingPrivileges() []Iso180135DrivingPrivilege`

GetDrivingPrivileges returns the DrivingPrivileges field if non-nil, zero value otherwise.

### GetDrivingPrivilegesOk

`func (o *Iso180135StandardNamespaceOutput) GetDrivingPrivilegesOk() (*[]Iso180135DrivingPrivilege, bool)`

GetDrivingPrivilegesOk returns a tuple with the DrivingPrivileges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrivingPrivileges

`func (o *Iso180135StandardNamespaceOutput) SetDrivingPrivileges(v []Iso180135DrivingPrivilege)`

SetDrivingPrivileges sets DrivingPrivileges field to given value.

### HasDrivingPrivileges

`func (o *Iso180135StandardNamespaceOutput) HasDrivingPrivileges() bool`

HasDrivingPrivileges returns a boolean if a field has been set.

### SetDrivingPrivilegesNil

`func (o *Iso180135StandardNamespaceOutput) SetDrivingPrivilegesNil(b bool)`

 SetDrivingPrivilegesNil sets the value for DrivingPrivileges to be an explicit nil

### UnsetDrivingPrivileges
`func (o *Iso180135StandardNamespaceOutput) UnsetDrivingPrivileges()`

UnsetDrivingPrivileges ensures that no value is present for DrivingPrivileges, not even an explicit nil
### GetUnDistinguishingSign

`func (o *Iso180135StandardNamespaceOutput) GetUnDistinguishingSign() string`

GetUnDistinguishingSign returns the UnDistinguishingSign field if non-nil, zero value otherwise.

### GetUnDistinguishingSignOk

`func (o *Iso180135StandardNamespaceOutput) GetUnDistinguishingSignOk() (*string, bool)`

GetUnDistinguishingSignOk returns a tuple with the UnDistinguishingSign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnDistinguishingSign

`func (o *Iso180135StandardNamespaceOutput) SetUnDistinguishingSign(v string)`

SetUnDistinguishingSign sets UnDistinguishingSign field to given value.

### HasUnDistinguishingSign

`func (o *Iso180135StandardNamespaceOutput) HasUnDistinguishingSign() bool`

HasUnDistinguishingSign returns a boolean if a field has been set.

### SetUnDistinguishingSignNil

`func (o *Iso180135StandardNamespaceOutput) SetUnDistinguishingSignNil(b bool)`

 SetUnDistinguishingSignNil sets the value for UnDistinguishingSign to be an explicit nil

### UnsetUnDistinguishingSign
`func (o *Iso180135StandardNamespaceOutput) UnsetUnDistinguishingSign()`

UnsetUnDistinguishingSign ensures that no value is present for UnDistinguishingSign, not even an explicit nil
### GetAdministrativeNumber

`func (o *Iso180135StandardNamespaceOutput) GetAdministrativeNumber() string`

GetAdministrativeNumber returns the AdministrativeNumber field if non-nil, zero value otherwise.

### GetAdministrativeNumberOk

`func (o *Iso180135StandardNamespaceOutput) GetAdministrativeNumberOk() (*string, bool)`

GetAdministrativeNumberOk returns a tuple with the AdministrativeNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministrativeNumber

`func (o *Iso180135StandardNamespaceOutput) SetAdministrativeNumber(v string)`

SetAdministrativeNumber sets AdministrativeNumber field to given value.

### HasAdministrativeNumber

`func (o *Iso180135StandardNamespaceOutput) HasAdministrativeNumber() bool`

HasAdministrativeNumber returns a boolean if a field has been set.

### SetAdministrativeNumberNil

`func (o *Iso180135StandardNamespaceOutput) SetAdministrativeNumberNil(b bool)`

 SetAdministrativeNumberNil sets the value for AdministrativeNumber to be an explicit nil

### UnsetAdministrativeNumber
`func (o *Iso180135StandardNamespaceOutput) UnsetAdministrativeNumber()`

UnsetAdministrativeNumber ensures that no value is present for AdministrativeNumber, not even an explicit nil
### GetSex

`func (o *Iso180135StandardNamespaceOutput) GetSex() int32`

GetSex returns the Sex field if non-nil, zero value otherwise.

### GetSexOk

`func (o *Iso180135StandardNamespaceOutput) GetSexOk() (*int32, bool)`

GetSexOk returns a tuple with the Sex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSex

`func (o *Iso180135StandardNamespaceOutput) SetSex(v int32)`

SetSex sets Sex field to given value.

### HasSex

`func (o *Iso180135StandardNamespaceOutput) HasSex() bool`

HasSex returns a boolean if a field has been set.

### SetSexNil

`func (o *Iso180135StandardNamespaceOutput) SetSexNil(b bool)`

 SetSexNil sets the value for Sex to be an explicit nil

### UnsetSex
`func (o *Iso180135StandardNamespaceOutput) UnsetSex()`

UnsetSex ensures that no value is present for Sex, not even an explicit nil
### GetHeightCentimeters

`func (o *Iso180135StandardNamespaceOutput) GetHeightCentimeters() int32`

GetHeightCentimeters returns the HeightCentimeters field if non-nil, zero value otherwise.

### GetHeightCentimetersOk

`func (o *Iso180135StandardNamespaceOutput) GetHeightCentimetersOk() (*int32, bool)`

GetHeightCentimetersOk returns a tuple with the HeightCentimeters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeightCentimeters

`func (o *Iso180135StandardNamespaceOutput) SetHeightCentimeters(v int32)`

SetHeightCentimeters sets HeightCentimeters field to given value.

### HasHeightCentimeters

`func (o *Iso180135StandardNamespaceOutput) HasHeightCentimeters() bool`

HasHeightCentimeters returns a boolean if a field has been set.

### SetHeightCentimetersNil

`func (o *Iso180135StandardNamespaceOutput) SetHeightCentimetersNil(b bool)`

 SetHeightCentimetersNil sets the value for HeightCentimeters to be an explicit nil

### UnsetHeightCentimeters
`func (o *Iso180135StandardNamespaceOutput) UnsetHeightCentimeters()`

UnsetHeightCentimeters ensures that no value is present for HeightCentimeters, not even an explicit nil
### GetWeightKilograms

`func (o *Iso180135StandardNamespaceOutput) GetWeightKilograms() int32`

GetWeightKilograms returns the WeightKilograms field if non-nil, zero value otherwise.

### GetWeightKilogramsOk

`func (o *Iso180135StandardNamespaceOutput) GetWeightKilogramsOk() (*int32, bool)`

GetWeightKilogramsOk returns a tuple with the WeightKilograms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightKilograms

`func (o *Iso180135StandardNamespaceOutput) SetWeightKilograms(v int32)`

SetWeightKilograms sets WeightKilograms field to given value.

### HasWeightKilograms

`func (o *Iso180135StandardNamespaceOutput) HasWeightKilograms() bool`

HasWeightKilograms returns a boolean if a field has been set.

### SetWeightKilogramsNil

`func (o *Iso180135StandardNamespaceOutput) SetWeightKilogramsNil(b bool)`

 SetWeightKilogramsNil sets the value for WeightKilograms to be an explicit nil

### UnsetWeightKilograms
`func (o *Iso180135StandardNamespaceOutput) UnsetWeightKilograms()`

UnsetWeightKilograms ensures that no value is present for WeightKilograms, not even an explicit nil
### GetEyeColor

`func (o *Iso180135StandardNamespaceOutput) GetEyeColor() string`

GetEyeColor returns the EyeColor field if non-nil, zero value otherwise.

### GetEyeColorOk

`func (o *Iso180135StandardNamespaceOutput) GetEyeColorOk() (*string, bool)`

GetEyeColorOk returns a tuple with the EyeColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEyeColor

`func (o *Iso180135StandardNamespaceOutput) SetEyeColor(v string)`

SetEyeColor sets EyeColor field to given value.

### HasEyeColor

`func (o *Iso180135StandardNamespaceOutput) HasEyeColor() bool`

HasEyeColor returns a boolean if a field has been set.

### SetEyeColorNil

`func (o *Iso180135StandardNamespaceOutput) SetEyeColorNil(b bool)`

 SetEyeColorNil sets the value for EyeColor to be an explicit nil

### UnsetEyeColor
`func (o *Iso180135StandardNamespaceOutput) UnsetEyeColor()`

UnsetEyeColor ensures that no value is present for EyeColor, not even an explicit nil
### GetHairColor

`func (o *Iso180135StandardNamespaceOutput) GetHairColor() string`

GetHairColor returns the HairColor field if non-nil, zero value otherwise.

### GetHairColorOk

`func (o *Iso180135StandardNamespaceOutput) GetHairColorOk() (*string, bool)`

GetHairColorOk returns a tuple with the HairColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHairColor

`func (o *Iso180135StandardNamespaceOutput) SetHairColor(v string)`

SetHairColor sets HairColor field to given value.

### HasHairColor

`func (o *Iso180135StandardNamespaceOutput) HasHairColor() bool`

HasHairColor returns a boolean if a field has been set.

### SetHairColorNil

`func (o *Iso180135StandardNamespaceOutput) SetHairColorNil(b bool)`

 SetHairColorNil sets the value for HairColor to be an explicit nil

### UnsetHairColor
`func (o *Iso180135StandardNamespaceOutput) UnsetHairColor()`

UnsetHairColor ensures that no value is present for HairColor, not even an explicit nil
### GetBirthPlace

`func (o *Iso180135StandardNamespaceOutput) GetBirthPlace() string`

GetBirthPlace returns the BirthPlace field if non-nil, zero value otherwise.

### GetBirthPlaceOk

`func (o *Iso180135StandardNamespaceOutput) GetBirthPlaceOk() (*string, bool)`

GetBirthPlaceOk returns a tuple with the BirthPlace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthPlace

`func (o *Iso180135StandardNamespaceOutput) SetBirthPlace(v string)`

SetBirthPlace sets BirthPlace field to given value.

### HasBirthPlace

`func (o *Iso180135StandardNamespaceOutput) HasBirthPlace() bool`

HasBirthPlace returns a boolean if a field has been set.

### SetBirthPlaceNil

`func (o *Iso180135StandardNamespaceOutput) SetBirthPlaceNil(b bool)`

 SetBirthPlaceNil sets the value for BirthPlace to be an explicit nil

### UnsetBirthPlace
`func (o *Iso180135StandardNamespaceOutput) UnsetBirthPlace()`

UnsetBirthPlace ensures that no value is present for BirthPlace, not even an explicit nil
### GetResidentAddress

`func (o *Iso180135StandardNamespaceOutput) GetResidentAddress() string`

GetResidentAddress returns the ResidentAddress field if non-nil, zero value otherwise.

### GetResidentAddressOk

`func (o *Iso180135StandardNamespaceOutput) GetResidentAddressOk() (*string, bool)`

GetResidentAddressOk returns a tuple with the ResidentAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidentAddress

`func (o *Iso180135StandardNamespaceOutput) SetResidentAddress(v string)`

SetResidentAddress sets ResidentAddress field to given value.

### HasResidentAddress

`func (o *Iso180135StandardNamespaceOutput) HasResidentAddress() bool`

HasResidentAddress returns a boolean if a field has been set.

### SetResidentAddressNil

`func (o *Iso180135StandardNamespaceOutput) SetResidentAddressNil(b bool)`

 SetResidentAddressNil sets the value for ResidentAddress to be an explicit nil

### UnsetResidentAddress
`func (o *Iso180135StandardNamespaceOutput) UnsetResidentAddress()`

UnsetResidentAddress ensures that no value is present for ResidentAddress, not even an explicit nil
### GetResidentCity

`func (o *Iso180135StandardNamespaceOutput) GetResidentCity() string`

GetResidentCity returns the ResidentCity field if non-nil, zero value otherwise.

### GetResidentCityOk

`func (o *Iso180135StandardNamespaceOutput) GetResidentCityOk() (*string, bool)`

GetResidentCityOk returns a tuple with the ResidentCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidentCity

`func (o *Iso180135StandardNamespaceOutput) SetResidentCity(v string)`

SetResidentCity sets ResidentCity field to given value.

### HasResidentCity

`func (o *Iso180135StandardNamespaceOutput) HasResidentCity() bool`

HasResidentCity returns a boolean if a field has been set.

### SetResidentCityNil

`func (o *Iso180135StandardNamespaceOutput) SetResidentCityNil(b bool)`

 SetResidentCityNil sets the value for ResidentCity to be an explicit nil

### UnsetResidentCity
`func (o *Iso180135StandardNamespaceOutput) UnsetResidentCity()`

UnsetResidentCity ensures that no value is present for ResidentCity, not even an explicit nil
### GetResidentState

`func (o *Iso180135StandardNamespaceOutput) GetResidentState() string`

GetResidentState returns the ResidentState field if non-nil, zero value otherwise.

### GetResidentStateOk

`func (o *Iso180135StandardNamespaceOutput) GetResidentStateOk() (*string, bool)`

GetResidentStateOk returns a tuple with the ResidentState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidentState

`func (o *Iso180135StandardNamespaceOutput) SetResidentState(v string)`

SetResidentState sets ResidentState field to given value.

### HasResidentState

`func (o *Iso180135StandardNamespaceOutput) HasResidentState() bool`

HasResidentState returns a boolean if a field has been set.

### SetResidentStateNil

`func (o *Iso180135StandardNamespaceOutput) SetResidentStateNil(b bool)`

 SetResidentStateNil sets the value for ResidentState to be an explicit nil

### UnsetResidentState
`func (o *Iso180135StandardNamespaceOutput) UnsetResidentState()`

UnsetResidentState ensures that no value is present for ResidentState, not even an explicit nil
### GetResidentPostalCode

`func (o *Iso180135StandardNamespaceOutput) GetResidentPostalCode() string`

GetResidentPostalCode returns the ResidentPostalCode field if non-nil, zero value otherwise.

### GetResidentPostalCodeOk

`func (o *Iso180135StandardNamespaceOutput) GetResidentPostalCodeOk() (*string, bool)`

GetResidentPostalCodeOk returns a tuple with the ResidentPostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidentPostalCode

`func (o *Iso180135StandardNamespaceOutput) SetResidentPostalCode(v string)`

SetResidentPostalCode sets ResidentPostalCode field to given value.

### HasResidentPostalCode

`func (o *Iso180135StandardNamespaceOutput) HasResidentPostalCode() bool`

HasResidentPostalCode returns a boolean if a field has been set.

### SetResidentPostalCodeNil

`func (o *Iso180135StandardNamespaceOutput) SetResidentPostalCodeNil(b bool)`

 SetResidentPostalCodeNil sets the value for ResidentPostalCode to be an explicit nil

### UnsetResidentPostalCode
`func (o *Iso180135StandardNamespaceOutput) UnsetResidentPostalCode()`

UnsetResidentPostalCode ensures that no value is present for ResidentPostalCode, not even an explicit nil
### GetResidentCountry

`func (o *Iso180135StandardNamespaceOutput) GetResidentCountry() string`

GetResidentCountry returns the ResidentCountry field if non-nil, zero value otherwise.

### GetResidentCountryOk

`func (o *Iso180135StandardNamespaceOutput) GetResidentCountryOk() (*string, bool)`

GetResidentCountryOk returns a tuple with the ResidentCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidentCountry

`func (o *Iso180135StandardNamespaceOutput) SetResidentCountry(v string)`

SetResidentCountry sets ResidentCountry field to given value.

### HasResidentCountry

`func (o *Iso180135StandardNamespaceOutput) HasResidentCountry() bool`

HasResidentCountry returns a boolean if a field has been set.

### SetResidentCountryNil

`func (o *Iso180135StandardNamespaceOutput) SetResidentCountryNil(b bool)`

 SetResidentCountryNil sets the value for ResidentCountry to be an explicit nil

### UnsetResidentCountry
`func (o *Iso180135StandardNamespaceOutput) UnsetResidentCountry()`

UnsetResidentCountry ensures that no value is present for ResidentCountry, not even an explicit nil
### GetPortraitCaptureDate

`func (o *Iso180135StandardNamespaceOutput) GetPortraitCaptureDate() time.Time`

GetPortraitCaptureDate returns the PortraitCaptureDate field if non-nil, zero value otherwise.

### GetPortraitCaptureDateOk

`func (o *Iso180135StandardNamespaceOutput) GetPortraitCaptureDateOk() (*time.Time, bool)`

GetPortraitCaptureDateOk returns a tuple with the PortraitCaptureDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortraitCaptureDate

`func (o *Iso180135StandardNamespaceOutput) SetPortraitCaptureDate(v time.Time)`

SetPortraitCaptureDate sets PortraitCaptureDate field to given value.

### HasPortraitCaptureDate

`func (o *Iso180135StandardNamespaceOutput) HasPortraitCaptureDate() bool`

HasPortraitCaptureDate returns a boolean if a field has been set.

### SetPortraitCaptureDateNil

`func (o *Iso180135StandardNamespaceOutput) SetPortraitCaptureDateNil(b bool)`

 SetPortraitCaptureDateNil sets the value for PortraitCaptureDate to be an explicit nil

### UnsetPortraitCaptureDate
`func (o *Iso180135StandardNamespaceOutput) UnsetPortraitCaptureDate()`

UnsetPortraitCaptureDate ensures that no value is present for PortraitCaptureDate, not even an explicit nil
### GetAgeInYears

`func (o *Iso180135StandardNamespaceOutput) GetAgeInYears() int32`

GetAgeInYears returns the AgeInYears field if non-nil, zero value otherwise.

### GetAgeInYearsOk

`func (o *Iso180135StandardNamespaceOutput) GetAgeInYearsOk() (*int32, bool)`

GetAgeInYearsOk returns a tuple with the AgeInYears field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgeInYears

`func (o *Iso180135StandardNamespaceOutput) SetAgeInYears(v int32)`

SetAgeInYears sets AgeInYears field to given value.

### HasAgeInYears

`func (o *Iso180135StandardNamespaceOutput) HasAgeInYears() bool`

HasAgeInYears returns a boolean if a field has been set.

### SetAgeInYearsNil

`func (o *Iso180135StandardNamespaceOutput) SetAgeInYearsNil(b bool)`

 SetAgeInYearsNil sets the value for AgeInYears to be an explicit nil

### UnsetAgeInYears
`func (o *Iso180135StandardNamespaceOutput) UnsetAgeInYears()`

UnsetAgeInYears ensures that no value is present for AgeInYears, not even an explicit nil
### GetAgeBirthYear

`func (o *Iso180135StandardNamespaceOutput) GetAgeBirthYear() int32`

GetAgeBirthYear returns the AgeBirthYear field if non-nil, zero value otherwise.

### GetAgeBirthYearOk

`func (o *Iso180135StandardNamespaceOutput) GetAgeBirthYearOk() (*int32, bool)`

GetAgeBirthYearOk returns a tuple with the AgeBirthYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgeBirthYear

`func (o *Iso180135StandardNamespaceOutput) SetAgeBirthYear(v int32)`

SetAgeBirthYear sets AgeBirthYear field to given value.

### HasAgeBirthYear

`func (o *Iso180135StandardNamespaceOutput) HasAgeBirthYear() bool`

HasAgeBirthYear returns a boolean if a field has been set.

### SetAgeBirthYearNil

`func (o *Iso180135StandardNamespaceOutput) SetAgeBirthYearNil(b bool)`

 SetAgeBirthYearNil sets the value for AgeBirthYear to be an explicit nil

### UnsetAgeBirthYear
`func (o *Iso180135StandardNamespaceOutput) UnsetAgeBirthYear()`

UnsetAgeBirthYear ensures that no value is present for AgeBirthYear, not even an explicit nil
### GetAgeOver

`func (o *Iso180135StandardNamespaceOutput) GetAgeOver() []AgeOverOutput`

GetAgeOver returns the AgeOver field if non-nil, zero value otherwise.

### GetAgeOverOk

`func (o *Iso180135StandardNamespaceOutput) GetAgeOverOk() (*[]AgeOverOutput, bool)`

GetAgeOverOk returns a tuple with the AgeOver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgeOver

`func (o *Iso180135StandardNamespaceOutput) SetAgeOver(v []AgeOverOutput)`

SetAgeOver sets AgeOver field to given value.

### HasAgeOver

`func (o *Iso180135StandardNamespaceOutput) HasAgeOver() bool`

HasAgeOver returns a boolean if a field has been set.

### SetAgeOverNil

`func (o *Iso180135StandardNamespaceOutput) SetAgeOverNil(b bool)`

 SetAgeOverNil sets the value for AgeOver to be an explicit nil

### UnsetAgeOver
`func (o *Iso180135StandardNamespaceOutput) UnsetAgeOver()`

UnsetAgeOver ensures that no value is present for AgeOver, not even an explicit nil
### GetBiometricTemplateFace

`func (o *Iso180135StandardNamespaceOutput) GetBiometricTemplateFace() Iso180132BiometricGroupTemplate`

GetBiometricTemplateFace returns the BiometricTemplateFace field if non-nil, zero value otherwise.

### GetBiometricTemplateFaceOk

`func (o *Iso180135StandardNamespaceOutput) GetBiometricTemplateFaceOk() (*Iso180132BiometricGroupTemplate, bool)`

GetBiometricTemplateFaceOk returns a tuple with the BiometricTemplateFace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiometricTemplateFace

`func (o *Iso180135StandardNamespaceOutput) SetBiometricTemplateFace(v Iso180132BiometricGroupTemplate)`

SetBiometricTemplateFace sets BiometricTemplateFace field to given value.

### HasBiometricTemplateFace

`func (o *Iso180135StandardNamespaceOutput) HasBiometricTemplateFace() bool`

HasBiometricTemplateFace returns a boolean if a field has been set.

### SetBiometricTemplateFaceNil

`func (o *Iso180135StandardNamespaceOutput) SetBiometricTemplateFaceNil(b bool)`

 SetBiometricTemplateFaceNil sets the value for BiometricTemplateFace to be an explicit nil

### UnsetBiometricTemplateFace
`func (o *Iso180135StandardNamespaceOutput) UnsetBiometricTemplateFace()`

UnsetBiometricTemplateFace ensures that no value is present for BiometricTemplateFace, not even an explicit nil
### GetBiometricTemplateVoice

`func (o *Iso180135StandardNamespaceOutput) GetBiometricTemplateVoice() Iso180132BiometricGroupTemplate`

GetBiometricTemplateVoice returns the BiometricTemplateVoice field if non-nil, zero value otherwise.

### GetBiometricTemplateVoiceOk

`func (o *Iso180135StandardNamespaceOutput) GetBiometricTemplateVoiceOk() (*Iso180132BiometricGroupTemplate, bool)`

GetBiometricTemplateVoiceOk returns a tuple with the BiometricTemplateVoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiometricTemplateVoice

`func (o *Iso180135StandardNamespaceOutput) SetBiometricTemplateVoice(v Iso180132BiometricGroupTemplate)`

SetBiometricTemplateVoice sets BiometricTemplateVoice field to given value.

### HasBiometricTemplateVoice

`func (o *Iso180135StandardNamespaceOutput) HasBiometricTemplateVoice() bool`

HasBiometricTemplateVoice returns a boolean if a field has been set.

### SetBiometricTemplateVoiceNil

`func (o *Iso180135StandardNamespaceOutput) SetBiometricTemplateVoiceNil(b bool)`

 SetBiometricTemplateVoiceNil sets the value for BiometricTemplateVoice to be an explicit nil

### UnsetBiometricTemplateVoice
`func (o *Iso180135StandardNamespaceOutput) UnsetBiometricTemplateVoice()`

UnsetBiometricTemplateVoice ensures that no value is present for BiometricTemplateVoice, not even an explicit nil
### GetBiometricTemplateFinger

`func (o *Iso180135StandardNamespaceOutput) GetBiometricTemplateFinger() Iso180132BiometricGroupTemplate`

GetBiometricTemplateFinger returns the BiometricTemplateFinger field if non-nil, zero value otherwise.

### GetBiometricTemplateFingerOk

`func (o *Iso180135StandardNamespaceOutput) GetBiometricTemplateFingerOk() (*Iso180132BiometricGroupTemplate, bool)`

GetBiometricTemplateFingerOk returns a tuple with the BiometricTemplateFinger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiometricTemplateFinger

`func (o *Iso180135StandardNamespaceOutput) SetBiometricTemplateFinger(v Iso180132BiometricGroupTemplate)`

SetBiometricTemplateFinger sets BiometricTemplateFinger field to given value.

### HasBiometricTemplateFinger

`func (o *Iso180135StandardNamespaceOutput) HasBiometricTemplateFinger() bool`

HasBiometricTemplateFinger returns a boolean if a field has been set.

### SetBiometricTemplateFingerNil

`func (o *Iso180135StandardNamespaceOutput) SetBiometricTemplateFingerNil(b bool)`

 SetBiometricTemplateFingerNil sets the value for BiometricTemplateFinger to be an explicit nil

### UnsetBiometricTemplateFinger
`func (o *Iso180135StandardNamespaceOutput) UnsetBiometricTemplateFinger()`

UnsetBiometricTemplateFinger ensures that no value is present for BiometricTemplateFinger, not even an explicit nil
### GetBiometricTemplateIris

`func (o *Iso180135StandardNamespaceOutput) GetBiometricTemplateIris() Iso180132BiometricGroupTemplate`

GetBiometricTemplateIris returns the BiometricTemplateIris field if non-nil, zero value otherwise.

### GetBiometricTemplateIrisOk

`func (o *Iso180135StandardNamespaceOutput) GetBiometricTemplateIrisOk() (*Iso180132BiometricGroupTemplate, bool)`

GetBiometricTemplateIrisOk returns a tuple with the BiometricTemplateIris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiometricTemplateIris

`func (o *Iso180135StandardNamespaceOutput) SetBiometricTemplateIris(v Iso180132BiometricGroupTemplate)`

SetBiometricTemplateIris sets BiometricTemplateIris field to given value.

### HasBiometricTemplateIris

`func (o *Iso180135StandardNamespaceOutput) HasBiometricTemplateIris() bool`

HasBiometricTemplateIris returns a boolean if a field has been set.

### SetBiometricTemplateIrisNil

`func (o *Iso180135StandardNamespaceOutput) SetBiometricTemplateIrisNil(b bool)`

 SetBiometricTemplateIrisNil sets the value for BiometricTemplateIris to be an explicit nil

### UnsetBiometricTemplateIris
`func (o *Iso180135StandardNamespaceOutput) UnsetBiometricTemplateIris()`

UnsetBiometricTemplateIris ensures that no value is present for BiometricTemplateIris, not even an explicit nil
### GetBiometricTemplateRetina

`func (o *Iso180135StandardNamespaceOutput) GetBiometricTemplateRetina() Iso180132BiometricGroupTemplate`

GetBiometricTemplateRetina returns the BiometricTemplateRetina field if non-nil, zero value otherwise.

### GetBiometricTemplateRetinaOk

`func (o *Iso180135StandardNamespaceOutput) GetBiometricTemplateRetinaOk() (*Iso180132BiometricGroupTemplate, bool)`

GetBiometricTemplateRetinaOk returns a tuple with the BiometricTemplateRetina field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiometricTemplateRetina

`func (o *Iso180135StandardNamespaceOutput) SetBiometricTemplateRetina(v Iso180132BiometricGroupTemplate)`

SetBiometricTemplateRetina sets BiometricTemplateRetina field to given value.

### HasBiometricTemplateRetina

`func (o *Iso180135StandardNamespaceOutput) HasBiometricTemplateRetina() bool`

HasBiometricTemplateRetina returns a boolean if a field has been set.

### SetBiometricTemplateRetinaNil

`func (o *Iso180135StandardNamespaceOutput) SetBiometricTemplateRetinaNil(b bool)`

 SetBiometricTemplateRetinaNil sets the value for BiometricTemplateRetina to be an explicit nil

### UnsetBiometricTemplateRetina
`func (o *Iso180135StandardNamespaceOutput) UnsetBiometricTemplateRetina()`

UnsetBiometricTemplateRetina ensures that no value is present for BiometricTemplateRetina, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


