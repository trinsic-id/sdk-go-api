# EudiPidCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GivenName** | Pointer to **NullableString** | Current first name(s), including middle name(s) where applicable, of the individual to whom the PID relates. | [optional] 
**GivenNameBirth** | Pointer to **NullableString** | First name(s), including middle name(s), of the individual at the time of birth. | [optional] 
**FamilyName** | Pointer to **NullableString** | Current last name(s) or surname(s) of the individuals. | [optional] 
**FamilyNameBirth** | Pointer to **NullableString** | Last name(s) or surname(s) of the individual at the time of birth. | [optional] 
**DateOfBirth** | Pointer to **NullableString** | The date of birth of the individual. | [optional] 
**BirthPlace** | Pointer to **NullableString** | Country (ISO 3166-1 alpha-2), state, province, district, local area, municipality, city, town, or village where the individual was born. | [optional] 
**Nationality** | Pointer to **[]string** | One or more ISO 3166-1 alpha-2 country codes representing the nationality of the individual. | [optional] 
**ExpiryDate** | Pointer to **NullableString** | Date when the PID expires.              Because PIDs are not necessarily backed by a physical document, this does not correspond to the expiration date of such. | [optional] 
**IssuingAuthority** | Pointer to **NullableString** | Name of the administrative authority that issued the PID, or the ISO 3166 alpha-2 country code of the respective Member State if there is no separate authority entitled to issue PIDs. | [optional] 
**IssuingCountry** | Pointer to **NullableString** | ISO 3166-1 alpha-2 country code of the country or territory of the issuer of the PID. | [optional] 
**IssuanceDate** | Pointer to **NullableString** | Date when the PID was issued and/or the administrative validity period began. | [optional] 
**TrustAnchor** | Pointer to **NullableString** | URL where a machine-readable trust anchor for verifying the PID can be found. | [optional] 
**AttestationLegalCategory** | Pointer to **NullableString** | Indicator that this credential has been issued specifically as PID.              When present, typically equal to \&quot;PID\&quot;. | [optional] 
**DocumentNumber** | Pointer to **NullableString** | A number assigned to the PID by the issuer.              This does not correspond to the document number of a physical document. | [optional] 
**IssuingJurisdiction** | Pointer to **NullableString** | ISO 3166-2 country-subdivision code for the jurisdiction that issued the PID. | [optional] 
**LocationStatus** | Pointer to **NullableString** | Location of validity status information for the PID, where revocation information is published. | [optional] 
**Sex** | Pointer to **NullableInt32** | Sex of the individual.              This is a superset of the ISO/IEC 5218 sex code standard, defining additional \&quot;other\&quot;, \&quot;inter\&quot;, \&quot;diverse\&quot;, and \&quot;open\&quot; values.              Possible values: - 0: Unknown - 1: Male - 2: Female - 3: Other - 4: Inter - 5: Diverse - 6: Open - 9: Not Applicable              For values 0, 1, 2, and 9, ISO/IEC 5218 applies. | [optional] 
**EmailAddress** | Pointer to **NullableString** | The individual&#39;s email address. | [optional] 
**MobilePhoneNumber** | Pointer to **NullableString** | The mobile phone number of the individual, in international format. | [optional] 
**PersonalAdministrativeNumber** | Pointer to **NullableString** | A value assigned to the individual that is unique among all personal administrative numbers issued by the provider. | [optional] 
**ResidentAddress** | Pointer to **NullableString** | The full address of the individual&#39;s permanent residence. | [optional] 
**ResidentStreet** | Pointer to **NullableString** | The street name of the individual&#39;s permanent residence. | [optional] 
**ResidentHouseNumber** | Pointer to **NullableString** | The house number of the individual&#39;s permanent residence. | [optional] 
**ResidentCity** | Pointer to **NullableString** | The municipality, city, town, or village of the individual&#39;s permanent residence. | [optional] 
**ResidentState** | Pointer to **NullableString** | The state, province, district, or local area of the individual&#39;s permanent residence. | [optional] 
**ResidentPostalCode** | Pointer to **NullableString** | The postal code of the individual&#39;s permanent residence. | [optional] 
**ResidentCountry** | Pointer to **NullableString** | The ISO 3166-1 alpha-2 country code of the individual&#39;s permanent residence. | [optional] 

## Methods

### NewEudiPidCredential

`func NewEudiPidCredential() *EudiPidCredential`

NewEudiPidCredential instantiates a new EudiPidCredential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEudiPidCredentialWithDefaults

`func NewEudiPidCredentialWithDefaults() *EudiPidCredential`

NewEudiPidCredentialWithDefaults instantiates a new EudiPidCredential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGivenName

`func (o *EudiPidCredential) GetGivenName() string`

GetGivenName returns the GivenName field if non-nil, zero value otherwise.

### GetGivenNameOk

`func (o *EudiPidCredential) GetGivenNameOk() (*string, bool)`

GetGivenNameOk returns a tuple with the GivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenName

`func (o *EudiPidCredential) SetGivenName(v string)`

SetGivenName sets GivenName field to given value.

### HasGivenName

`func (o *EudiPidCredential) HasGivenName() bool`

HasGivenName returns a boolean if a field has been set.

### SetGivenNameNil

`func (o *EudiPidCredential) SetGivenNameNil(b bool)`

 SetGivenNameNil sets the value for GivenName to be an explicit nil

### UnsetGivenName
`func (o *EudiPidCredential) UnsetGivenName()`

UnsetGivenName ensures that no value is present for GivenName, not even an explicit nil
### GetGivenNameBirth

`func (o *EudiPidCredential) GetGivenNameBirth() string`

GetGivenNameBirth returns the GivenNameBirth field if non-nil, zero value otherwise.

### GetGivenNameBirthOk

`func (o *EudiPidCredential) GetGivenNameBirthOk() (*string, bool)`

GetGivenNameBirthOk returns a tuple with the GivenNameBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenNameBirth

`func (o *EudiPidCredential) SetGivenNameBirth(v string)`

SetGivenNameBirth sets GivenNameBirth field to given value.

### HasGivenNameBirth

`func (o *EudiPidCredential) HasGivenNameBirth() bool`

HasGivenNameBirth returns a boolean if a field has been set.

### SetGivenNameBirthNil

`func (o *EudiPidCredential) SetGivenNameBirthNil(b bool)`

 SetGivenNameBirthNil sets the value for GivenNameBirth to be an explicit nil

### UnsetGivenNameBirth
`func (o *EudiPidCredential) UnsetGivenNameBirth()`

UnsetGivenNameBirth ensures that no value is present for GivenNameBirth, not even an explicit nil
### GetFamilyName

`func (o *EudiPidCredential) GetFamilyName() string`

GetFamilyName returns the FamilyName field if non-nil, zero value otherwise.

### GetFamilyNameOk

`func (o *EudiPidCredential) GetFamilyNameOk() (*string, bool)`

GetFamilyNameOk returns a tuple with the FamilyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyName

`func (o *EudiPidCredential) SetFamilyName(v string)`

SetFamilyName sets FamilyName field to given value.

### HasFamilyName

`func (o *EudiPidCredential) HasFamilyName() bool`

HasFamilyName returns a boolean if a field has been set.

### SetFamilyNameNil

`func (o *EudiPidCredential) SetFamilyNameNil(b bool)`

 SetFamilyNameNil sets the value for FamilyName to be an explicit nil

### UnsetFamilyName
`func (o *EudiPidCredential) UnsetFamilyName()`

UnsetFamilyName ensures that no value is present for FamilyName, not even an explicit nil
### GetFamilyNameBirth

`func (o *EudiPidCredential) GetFamilyNameBirth() string`

GetFamilyNameBirth returns the FamilyNameBirth field if non-nil, zero value otherwise.

### GetFamilyNameBirthOk

`func (o *EudiPidCredential) GetFamilyNameBirthOk() (*string, bool)`

GetFamilyNameBirthOk returns a tuple with the FamilyNameBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyNameBirth

`func (o *EudiPidCredential) SetFamilyNameBirth(v string)`

SetFamilyNameBirth sets FamilyNameBirth field to given value.

### HasFamilyNameBirth

`func (o *EudiPidCredential) HasFamilyNameBirth() bool`

HasFamilyNameBirth returns a boolean if a field has been set.

### SetFamilyNameBirthNil

`func (o *EudiPidCredential) SetFamilyNameBirthNil(b bool)`

 SetFamilyNameBirthNil sets the value for FamilyNameBirth to be an explicit nil

### UnsetFamilyNameBirth
`func (o *EudiPidCredential) UnsetFamilyNameBirth()`

UnsetFamilyNameBirth ensures that no value is present for FamilyNameBirth, not even an explicit nil
### GetDateOfBirth

`func (o *EudiPidCredential) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *EudiPidCredential) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *EudiPidCredential) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *EudiPidCredential) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### SetDateOfBirthNil

`func (o *EudiPidCredential) SetDateOfBirthNil(b bool)`

 SetDateOfBirthNil sets the value for DateOfBirth to be an explicit nil

### UnsetDateOfBirth
`func (o *EudiPidCredential) UnsetDateOfBirth()`

UnsetDateOfBirth ensures that no value is present for DateOfBirth, not even an explicit nil
### GetBirthPlace

`func (o *EudiPidCredential) GetBirthPlace() string`

GetBirthPlace returns the BirthPlace field if non-nil, zero value otherwise.

### GetBirthPlaceOk

`func (o *EudiPidCredential) GetBirthPlaceOk() (*string, bool)`

GetBirthPlaceOk returns a tuple with the BirthPlace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthPlace

`func (o *EudiPidCredential) SetBirthPlace(v string)`

SetBirthPlace sets BirthPlace field to given value.

### HasBirthPlace

`func (o *EudiPidCredential) HasBirthPlace() bool`

HasBirthPlace returns a boolean if a field has been set.

### SetBirthPlaceNil

`func (o *EudiPidCredential) SetBirthPlaceNil(b bool)`

 SetBirthPlaceNil sets the value for BirthPlace to be an explicit nil

### UnsetBirthPlace
`func (o *EudiPidCredential) UnsetBirthPlace()`

UnsetBirthPlace ensures that no value is present for BirthPlace, not even an explicit nil
### GetNationality

`func (o *EudiPidCredential) GetNationality() []string`

GetNationality returns the Nationality field if non-nil, zero value otherwise.

### GetNationalityOk

`func (o *EudiPidCredential) GetNationalityOk() (*[]string, bool)`

GetNationalityOk returns a tuple with the Nationality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationality

`func (o *EudiPidCredential) SetNationality(v []string)`

SetNationality sets Nationality field to given value.

### HasNationality

`func (o *EudiPidCredential) HasNationality() bool`

HasNationality returns a boolean if a field has been set.

### SetNationalityNil

`func (o *EudiPidCredential) SetNationalityNil(b bool)`

 SetNationalityNil sets the value for Nationality to be an explicit nil

### UnsetNationality
`func (o *EudiPidCredential) UnsetNationality()`

UnsetNationality ensures that no value is present for Nationality, not even an explicit nil
### GetExpiryDate

`func (o *EudiPidCredential) GetExpiryDate() string`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *EudiPidCredential) GetExpiryDateOk() (*string, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *EudiPidCredential) SetExpiryDate(v string)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *EudiPidCredential) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### SetExpiryDateNil

`func (o *EudiPidCredential) SetExpiryDateNil(b bool)`

 SetExpiryDateNil sets the value for ExpiryDate to be an explicit nil

### UnsetExpiryDate
`func (o *EudiPidCredential) UnsetExpiryDate()`

UnsetExpiryDate ensures that no value is present for ExpiryDate, not even an explicit nil
### GetIssuingAuthority

`func (o *EudiPidCredential) GetIssuingAuthority() string`

GetIssuingAuthority returns the IssuingAuthority field if non-nil, zero value otherwise.

### GetIssuingAuthorityOk

`func (o *EudiPidCredential) GetIssuingAuthorityOk() (*string, bool)`

GetIssuingAuthorityOk returns a tuple with the IssuingAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingAuthority

`func (o *EudiPidCredential) SetIssuingAuthority(v string)`

SetIssuingAuthority sets IssuingAuthority field to given value.

### HasIssuingAuthority

`func (o *EudiPidCredential) HasIssuingAuthority() bool`

HasIssuingAuthority returns a boolean if a field has been set.

### SetIssuingAuthorityNil

`func (o *EudiPidCredential) SetIssuingAuthorityNil(b bool)`

 SetIssuingAuthorityNil sets the value for IssuingAuthority to be an explicit nil

### UnsetIssuingAuthority
`func (o *EudiPidCredential) UnsetIssuingAuthority()`

UnsetIssuingAuthority ensures that no value is present for IssuingAuthority, not even an explicit nil
### GetIssuingCountry

`func (o *EudiPidCredential) GetIssuingCountry() string`

GetIssuingCountry returns the IssuingCountry field if non-nil, zero value otherwise.

### GetIssuingCountryOk

`func (o *EudiPidCredential) GetIssuingCountryOk() (*string, bool)`

GetIssuingCountryOk returns a tuple with the IssuingCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingCountry

`func (o *EudiPidCredential) SetIssuingCountry(v string)`

SetIssuingCountry sets IssuingCountry field to given value.

### HasIssuingCountry

`func (o *EudiPidCredential) HasIssuingCountry() bool`

HasIssuingCountry returns a boolean if a field has been set.

### SetIssuingCountryNil

`func (o *EudiPidCredential) SetIssuingCountryNil(b bool)`

 SetIssuingCountryNil sets the value for IssuingCountry to be an explicit nil

### UnsetIssuingCountry
`func (o *EudiPidCredential) UnsetIssuingCountry()`

UnsetIssuingCountry ensures that no value is present for IssuingCountry, not even an explicit nil
### GetIssuanceDate

`func (o *EudiPidCredential) GetIssuanceDate() string`

GetIssuanceDate returns the IssuanceDate field if non-nil, zero value otherwise.

### GetIssuanceDateOk

`func (o *EudiPidCredential) GetIssuanceDateOk() (*string, bool)`

GetIssuanceDateOk returns a tuple with the IssuanceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuanceDate

`func (o *EudiPidCredential) SetIssuanceDate(v string)`

SetIssuanceDate sets IssuanceDate field to given value.

### HasIssuanceDate

`func (o *EudiPidCredential) HasIssuanceDate() bool`

HasIssuanceDate returns a boolean if a field has been set.

### SetIssuanceDateNil

`func (o *EudiPidCredential) SetIssuanceDateNil(b bool)`

 SetIssuanceDateNil sets the value for IssuanceDate to be an explicit nil

### UnsetIssuanceDate
`func (o *EudiPidCredential) UnsetIssuanceDate()`

UnsetIssuanceDate ensures that no value is present for IssuanceDate, not even an explicit nil
### GetTrustAnchor

`func (o *EudiPidCredential) GetTrustAnchor() string`

GetTrustAnchor returns the TrustAnchor field if non-nil, zero value otherwise.

### GetTrustAnchorOk

`func (o *EudiPidCredential) GetTrustAnchorOk() (*string, bool)`

GetTrustAnchorOk returns a tuple with the TrustAnchor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustAnchor

`func (o *EudiPidCredential) SetTrustAnchor(v string)`

SetTrustAnchor sets TrustAnchor field to given value.

### HasTrustAnchor

`func (o *EudiPidCredential) HasTrustAnchor() bool`

HasTrustAnchor returns a boolean if a field has been set.

### SetTrustAnchorNil

`func (o *EudiPidCredential) SetTrustAnchorNil(b bool)`

 SetTrustAnchorNil sets the value for TrustAnchor to be an explicit nil

### UnsetTrustAnchor
`func (o *EudiPidCredential) UnsetTrustAnchor()`

UnsetTrustAnchor ensures that no value is present for TrustAnchor, not even an explicit nil
### GetAttestationLegalCategory

`func (o *EudiPidCredential) GetAttestationLegalCategory() string`

GetAttestationLegalCategory returns the AttestationLegalCategory field if non-nil, zero value otherwise.

### GetAttestationLegalCategoryOk

`func (o *EudiPidCredential) GetAttestationLegalCategoryOk() (*string, bool)`

GetAttestationLegalCategoryOk returns a tuple with the AttestationLegalCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttestationLegalCategory

`func (o *EudiPidCredential) SetAttestationLegalCategory(v string)`

SetAttestationLegalCategory sets AttestationLegalCategory field to given value.

### HasAttestationLegalCategory

`func (o *EudiPidCredential) HasAttestationLegalCategory() bool`

HasAttestationLegalCategory returns a boolean if a field has been set.

### SetAttestationLegalCategoryNil

`func (o *EudiPidCredential) SetAttestationLegalCategoryNil(b bool)`

 SetAttestationLegalCategoryNil sets the value for AttestationLegalCategory to be an explicit nil

### UnsetAttestationLegalCategory
`func (o *EudiPidCredential) UnsetAttestationLegalCategory()`

UnsetAttestationLegalCategory ensures that no value is present for AttestationLegalCategory, not even an explicit nil
### GetDocumentNumber

`func (o *EudiPidCredential) GetDocumentNumber() string`

GetDocumentNumber returns the DocumentNumber field if non-nil, zero value otherwise.

### GetDocumentNumberOk

`func (o *EudiPidCredential) GetDocumentNumberOk() (*string, bool)`

GetDocumentNumberOk returns a tuple with the DocumentNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentNumber

`func (o *EudiPidCredential) SetDocumentNumber(v string)`

SetDocumentNumber sets DocumentNumber field to given value.

### HasDocumentNumber

`func (o *EudiPidCredential) HasDocumentNumber() bool`

HasDocumentNumber returns a boolean if a field has been set.

### SetDocumentNumberNil

`func (o *EudiPidCredential) SetDocumentNumberNil(b bool)`

 SetDocumentNumberNil sets the value for DocumentNumber to be an explicit nil

### UnsetDocumentNumber
`func (o *EudiPidCredential) UnsetDocumentNumber()`

UnsetDocumentNumber ensures that no value is present for DocumentNumber, not even an explicit nil
### GetIssuingJurisdiction

`func (o *EudiPidCredential) GetIssuingJurisdiction() string`

GetIssuingJurisdiction returns the IssuingJurisdiction field if non-nil, zero value otherwise.

### GetIssuingJurisdictionOk

`func (o *EudiPidCredential) GetIssuingJurisdictionOk() (*string, bool)`

GetIssuingJurisdictionOk returns a tuple with the IssuingJurisdiction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingJurisdiction

`func (o *EudiPidCredential) SetIssuingJurisdiction(v string)`

SetIssuingJurisdiction sets IssuingJurisdiction field to given value.

### HasIssuingJurisdiction

`func (o *EudiPidCredential) HasIssuingJurisdiction() bool`

HasIssuingJurisdiction returns a boolean if a field has been set.

### SetIssuingJurisdictionNil

`func (o *EudiPidCredential) SetIssuingJurisdictionNil(b bool)`

 SetIssuingJurisdictionNil sets the value for IssuingJurisdiction to be an explicit nil

### UnsetIssuingJurisdiction
`func (o *EudiPidCredential) UnsetIssuingJurisdiction()`

UnsetIssuingJurisdiction ensures that no value is present for IssuingJurisdiction, not even an explicit nil
### GetLocationStatus

`func (o *EudiPidCredential) GetLocationStatus() string`

GetLocationStatus returns the LocationStatus field if non-nil, zero value otherwise.

### GetLocationStatusOk

`func (o *EudiPidCredential) GetLocationStatusOk() (*string, bool)`

GetLocationStatusOk returns a tuple with the LocationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationStatus

`func (o *EudiPidCredential) SetLocationStatus(v string)`

SetLocationStatus sets LocationStatus field to given value.

### HasLocationStatus

`func (o *EudiPidCredential) HasLocationStatus() bool`

HasLocationStatus returns a boolean if a field has been set.

### SetLocationStatusNil

`func (o *EudiPidCredential) SetLocationStatusNil(b bool)`

 SetLocationStatusNil sets the value for LocationStatus to be an explicit nil

### UnsetLocationStatus
`func (o *EudiPidCredential) UnsetLocationStatus()`

UnsetLocationStatus ensures that no value is present for LocationStatus, not even an explicit nil
### GetSex

`func (o *EudiPidCredential) GetSex() int32`

GetSex returns the Sex field if non-nil, zero value otherwise.

### GetSexOk

`func (o *EudiPidCredential) GetSexOk() (*int32, bool)`

GetSexOk returns a tuple with the Sex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSex

`func (o *EudiPidCredential) SetSex(v int32)`

SetSex sets Sex field to given value.

### HasSex

`func (o *EudiPidCredential) HasSex() bool`

HasSex returns a boolean if a field has been set.

### SetSexNil

`func (o *EudiPidCredential) SetSexNil(b bool)`

 SetSexNil sets the value for Sex to be an explicit nil

### UnsetSex
`func (o *EudiPidCredential) UnsetSex()`

UnsetSex ensures that no value is present for Sex, not even an explicit nil
### GetEmailAddress

`func (o *EudiPidCredential) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *EudiPidCredential) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *EudiPidCredential) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *EudiPidCredential) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### SetEmailAddressNil

`func (o *EudiPidCredential) SetEmailAddressNil(b bool)`

 SetEmailAddressNil sets the value for EmailAddress to be an explicit nil

### UnsetEmailAddress
`func (o *EudiPidCredential) UnsetEmailAddress()`

UnsetEmailAddress ensures that no value is present for EmailAddress, not even an explicit nil
### GetMobilePhoneNumber

`func (o *EudiPidCredential) GetMobilePhoneNumber() string`

GetMobilePhoneNumber returns the MobilePhoneNumber field if non-nil, zero value otherwise.

### GetMobilePhoneNumberOk

`func (o *EudiPidCredential) GetMobilePhoneNumberOk() (*string, bool)`

GetMobilePhoneNumberOk returns a tuple with the MobilePhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobilePhoneNumber

`func (o *EudiPidCredential) SetMobilePhoneNumber(v string)`

SetMobilePhoneNumber sets MobilePhoneNumber field to given value.

### HasMobilePhoneNumber

`func (o *EudiPidCredential) HasMobilePhoneNumber() bool`

HasMobilePhoneNumber returns a boolean if a field has been set.

### SetMobilePhoneNumberNil

`func (o *EudiPidCredential) SetMobilePhoneNumberNil(b bool)`

 SetMobilePhoneNumberNil sets the value for MobilePhoneNumber to be an explicit nil

### UnsetMobilePhoneNumber
`func (o *EudiPidCredential) UnsetMobilePhoneNumber()`

UnsetMobilePhoneNumber ensures that no value is present for MobilePhoneNumber, not even an explicit nil
### GetPersonalAdministrativeNumber

`func (o *EudiPidCredential) GetPersonalAdministrativeNumber() string`

GetPersonalAdministrativeNumber returns the PersonalAdministrativeNumber field if non-nil, zero value otherwise.

### GetPersonalAdministrativeNumberOk

`func (o *EudiPidCredential) GetPersonalAdministrativeNumberOk() (*string, bool)`

GetPersonalAdministrativeNumberOk returns a tuple with the PersonalAdministrativeNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalAdministrativeNumber

`func (o *EudiPidCredential) SetPersonalAdministrativeNumber(v string)`

SetPersonalAdministrativeNumber sets PersonalAdministrativeNumber field to given value.

### HasPersonalAdministrativeNumber

`func (o *EudiPidCredential) HasPersonalAdministrativeNumber() bool`

HasPersonalAdministrativeNumber returns a boolean if a field has been set.

### SetPersonalAdministrativeNumberNil

`func (o *EudiPidCredential) SetPersonalAdministrativeNumberNil(b bool)`

 SetPersonalAdministrativeNumberNil sets the value for PersonalAdministrativeNumber to be an explicit nil

### UnsetPersonalAdministrativeNumber
`func (o *EudiPidCredential) UnsetPersonalAdministrativeNumber()`

UnsetPersonalAdministrativeNumber ensures that no value is present for PersonalAdministrativeNumber, not even an explicit nil
### GetResidentAddress

`func (o *EudiPidCredential) GetResidentAddress() string`

GetResidentAddress returns the ResidentAddress field if non-nil, zero value otherwise.

### GetResidentAddressOk

`func (o *EudiPidCredential) GetResidentAddressOk() (*string, bool)`

GetResidentAddressOk returns a tuple with the ResidentAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidentAddress

`func (o *EudiPidCredential) SetResidentAddress(v string)`

SetResidentAddress sets ResidentAddress field to given value.

### HasResidentAddress

`func (o *EudiPidCredential) HasResidentAddress() bool`

HasResidentAddress returns a boolean if a field has been set.

### SetResidentAddressNil

`func (o *EudiPidCredential) SetResidentAddressNil(b bool)`

 SetResidentAddressNil sets the value for ResidentAddress to be an explicit nil

### UnsetResidentAddress
`func (o *EudiPidCredential) UnsetResidentAddress()`

UnsetResidentAddress ensures that no value is present for ResidentAddress, not even an explicit nil
### GetResidentStreet

`func (o *EudiPidCredential) GetResidentStreet() string`

GetResidentStreet returns the ResidentStreet field if non-nil, zero value otherwise.

### GetResidentStreetOk

`func (o *EudiPidCredential) GetResidentStreetOk() (*string, bool)`

GetResidentStreetOk returns a tuple with the ResidentStreet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidentStreet

`func (o *EudiPidCredential) SetResidentStreet(v string)`

SetResidentStreet sets ResidentStreet field to given value.

### HasResidentStreet

`func (o *EudiPidCredential) HasResidentStreet() bool`

HasResidentStreet returns a boolean if a field has been set.

### SetResidentStreetNil

`func (o *EudiPidCredential) SetResidentStreetNil(b bool)`

 SetResidentStreetNil sets the value for ResidentStreet to be an explicit nil

### UnsetResidentStreet
`func (o *EudiPidCredential) UnsetResidentStreet()`

UnsetResidentStreet ensures that no value is present for ResidentStreet, not even an explicit nil
### GetResidentHouseNumber

`func (o *EudiPidCredential) GetResidentHouseNumber() string`

GetResidentHouseNumber returns the ResidentHouseNumber field if non-nil, zero value otherwise.

### GetResidentHouseNumberOk

`func (o *EudiPidCredential) GetResidentHouseNumberOk() (*string, bool)`

GetResidentHouseNumberOk returns a tuple with the ResidentHouseNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidentHouseNumber

`func (o *EudiPidCredential) SetResidentHouseNumber(v string)`

SetResidentHouseNumber sets ResidentHouseNumber field to given value.

### HasResidentHouseNumber

`func (o *EudiPidCredential) HasResidentHouseNumber() bool`

HasResidentHouseNumber returns a boolean if a field has been set.

### SetResidentHouseNumberNil

`func (o *EudiPidCredential) SetResidentHouseNumberNil(b bool)`

 SetResidentHouseNumberNil sets the value for ResidentHouseNumber to be an explicit nil

### UnsetResidentHouseNumber
`func (o *EudiPidCredential) UnsetResidentHouseNumber()`

UnsetResidentHouseNumber ensures that no value is present for ResidentHouseNumber, not even an explicit nil
### GetResidentCity

`func (o *EudiPidCredential) GetResidentCity() string`

GetResidentCity returns the ResidentCity field if non-nil, zero value otherwise.

### GetResidentCityOk

`func (o *EudiPidCredential) GetResidentCityOk() (*string, bool)`

GetResidentCityOk returns a tuple with the ResidentCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidentCity

`func (o *EudiPidCredential) SetResidentCity(v string)`

SetResidentCity sets ResidentCity field to given value.

### HasResidentCity

`func (o *EudiPidCredential) HasResidentCity() bool`

HasResidentCity returns a boolean if a field has been set.

### SetResidentCityNil

`func (o *EudiPidCredential) SetResidentCityNil(b bool)`

 SetResidentCityNil sets the value for ResidentCity to be an explicit nil

### UnsetResidentCity
`func (o *EudiPidCredential) UnsetResidentCity()`

UnsetResidentCity ensures that no value is present for ResidentCity, not even an explicit nil
### GetResidentState

`func (o *EudiPidCredential) GetResidentState() string`

GetResidentState returns the ResidentState field if non-nil, zero value otherwise.

### GetResidentStateOk

`func (o *EudiPidCredential) GetResidentStateOk() (*string, bool)`

GetResidentStateOk returns a tuple with the ResidentState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidentState

`func (o *EudiPidCredential) SetResidentState(v string)`

SetResidentState sets ResidentState field to given value.

### HasResidentState

`func (o *EudiPidCredential) HasResidentState() bool`

HasResidentState returns a boolean if a field has been set.

### SetResidentStateNil

`func (o *EudiPidCredential) SetResidentStateNil(b bool)`

 SetResidentStateNil sets the value for ResidentState to be an explicit nil

### UnsetResidentState
`func (o *EudiPidCredential) UnsetResidentState()`

UnsetResidentState ensures that no value is present for ResidentState, not even an explicit nil
### GetResidentPostalCode

`func (o *EudiPidCredential) GetResidentPostalCode() string`

GetResidentPostalCode returns the ResidentPostalCode field if non-nil, zero value otherwise.

### GetResidentPostalCodeOk

`func (o *EudiPidCredential) GetResidentPostalCodeOk() (*string, bool)`

GetResidentPostalCodeOk returns a tuple with the ResidentPostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidentPostalCode

`func (o *EudiPidCredential) SetResidentPostalCode(v string)`

SetResidentPostalCode sets ResidentPostalCode field to given value.

### HasResidentPostalCode

`func (o *EudiPidCredential) HasResidentPostalCode() bool`

HasResidentPostalCode returns a boolean if a field has been set.

### SetResidentPostalCodeNil

`func (o *EudiPidCredential) SetResidentPostalCodeNil(b bool)`

 SetResidentPostalCodeNil sets the value for ResidentPostalCode to be an explicit nil

### UnsetResidentPostalCode
`func (o *EudiPidCredential) UnsetResidentPostalCode()`

UnsetResidentPostalCode ensures that no value is present for ResidentPostalCode, not even an explicit nil
### GetResidentCountry

`func (o *EudiPidCredential) GetResidentCountry() string`

GetResidentCountry returns the ResidentCountry field if non-nil, zero value otherwise.

### GetResidentCountryOk

`func (o *EudiPidCredential) GetResidentCountryOk() (*string, bool)`

GetResidentCountryOk returns a tuple with the ResidentCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidentCountry

`func (o *EudiPidCredential) SetResidentCountry(v string)`

SetResidentCountry sets ResidentCountry field to given value.

### HasResidentCountry

`func (o *EudiPidCredential) HasResidentCountry() bool`

HasResidentCountry returns a boolean if a field has been set.

### SetResidentCountryNil

`func (o *EudiPidCredential) SetResidentCountryNil(b bool)`

 SetResidentCountryNil sets the value for ResidentCountry to be an explicit nil

### UnsetResidentCountry
`func (o *EudiPidCredential) UnsetResidentCountry()`

UnsetResidentCountry ensures that no value is present for ResidentCountry, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


