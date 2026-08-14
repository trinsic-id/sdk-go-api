# IdemiaMobileIdProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sub** | Pointer to **NullableString** | The OpenID Connect subject identifier for the verified individual. | [optional] 
**NameSuffix** | Pointer to **NullableString** | The individual&#39;s name suffix. | [optional] 
**FullName** | Pointer to **NullableString** | The individual&#39;s full name. | [optional] 
**GivenName** | Pointer to **NullableString** | The individual&#39;s given name. | [optional] 
**FamilyName** | Pointer to **NullableString** | The individual&#39;s family name. | [optional] 
**BirthDate** | Pointer to **NullableString** | The individual&#39;s birth date. | [optional] 
**AgeOver18** | Pointer to **NullableString** | Whether the individual is over 18 years old. | [optional] 
**AgeOver21** | Pointer to **NullableString** | Whether the individual is over 21 years old. | [optional] 
**AgeOver25** | Pointer to **NullableString** | Whether the individual is over 25 years old. | [optional] 
**ResidentAddress** | Pointer to **NullableString** | The individual&#39;s resident street address. | [optional] 
**ResidentState** | Pointer to **NullableString** | The individual&#39;s resident state. | [optional] 
**ResidentCountry** | Pointer to **NullableString** | The individual&#39;s resident country. | [optional] 
**ResidentPostalCode** | Pointer to **NullableString** | The individual&#39;s resident postal code. | [optional] 
**ResidentCity** | Pointer to **NullableString** | The individual&#39;s resident city. | [optional] 
**Email** | Pointer to **NullableString** | The individual&#39;s email address. | [optional] 
**EmailVerified** | Pointer to **NullableBool** | Whether the individual&#39;s email address has been verified. | [optional] 
**DocumentNumber** | Pointer to **NullableString** | The individual&#39;s document number. | [optional] 
**IssueDate** | Pointer to **NullableString** | The date the document was issued. | [optional] 
**IssuingCountry** | Pointer to **NullableString** | The country that issued the document. | [optional] 
**IssuingAuthority** | Pointer to **NullableString** | The authority that issued the document. | [optional] 
**ExpiryDate** | Pointer to **NullableString** | The date the document expires. | [optional] 
**IssuingJurisdiction** | Pointer to **NullableString** | The jurisdiction that issued the document. | [optional] 
**DocumentType** | Pointer to **NullableString** | The document type. | [optional] 
**Sex** | Pointer to **NullableString** | The individual&#39;s sex. | [optional] 

## Methods

### NewIdemiaMobileIdProviderOutput

`func NewIdemiaMobileIdProviderOutput() *IdemiaMobileIdProviderOutput`

NewIdemiaMobileIdProviderOutput instantiates a new IdemiaMobileIdProviderOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdemiaMobileIdProviderOutputWithDefaults

`func NewIdemiaMobileIdProviderOutputWithDefaults() *IdemiaMobileIdProviderOutput`

NewIdemiaMobileIdProviderOutputWithDefaults instantiates a new IdemiaMobileIdProviderOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSub

`func (o *IdemiaMobileIdProviderOutput) GetSub() string`

GetSub returns the Sub field if non-nil, zero value otherwise.

### GetSubOk

`func (o *IdemiaMobileIdProviderOutput) GetSubOk() (*string, bool)`

GetSubOk returns a tuple with the Sub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSub

`func (o *IdemiaMobileIdProviderOutput) SetSub(v string)`

SetSub sets Sub field to given value.

### HasSub

`func (o *IdemiaMobileIdProviderOutput) HasSub() bool`

HasSub returns a boolean if a field has been set.

### SetSubNil

`func (o *IdemiaMobileIdProviderOutput) SetSubNil(b bool)`

 SetSubNil sets the value for Sub to be an explicit nil

### UnsetSub
`func (o *IdemiaMobileIdProviderOutput) UnsetSub()`

UnsetSub ensures that no value is present for Sub, not even an explicit nil
### GetNameSuffix

`func (o *IdemiaMobileIdProviderOutput) GetNameSuffix() string`

GetNameSuffix returns the NameSuffix field if non-nil, zero value otherwise.

### GetNameSuffixOk

`func (o *IdemiaMobileIdProviderOutput) GetNameSuffixOk() (*string, bool)`

GetNameSuffixOk returns a tuple with the NameSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameSuffix

`func (o *IdemiaMobileIdProviderOutput) SetNameSuffix(v string)`

SetNameSuffix sets NameSuffix field to given value.

### HasNameSuffix

`func (o *IdemiaMobileIdProviderOutput) HasNameSuffix() bool`

HasNameSuffix returns a boolean if a field has been set.

### SetNameSuffixNil

`func (o *IdemiaMobileIdProviderOutput) SetNameSuffixNil(b bool)`

 SetNameSuffixNil sets the value for NameSuffix to be an explicit nil

### UnsetNameSuffix
`func (o *IdemiaMobileIdProviderOutput) UnsetNameSuffix()`

UnsetNameSuffix ensures that no value is present for NameSuffix, not even an explicit nil
### GetFullName

`func (o *IdemiaMobileIdProviderOutput) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *IdemiaMobileIdProviderOutput) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *IdemiaMobileIdProviderOutput) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *IdemiaMobileIdProviderOutput) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### SetFullNameNil

`func (o *IdemiaMobileIdProviderOutput) SetFullNameNil(b bool)`

 SetFullNameNil sets the value for FullName to be an explicit nil

### UnsetFullName
`func (o *IdemiaMobileIdProviderOutput) UnsetFullName()`

UnsetFullName ensures that no value is present for FullName, not even an explicit nil
### GetGivenName

`func (o *IdemiaMobileIdProviderOutput) GetGivenName() string`

GetGivenName returns the GivenName field if non-nil, zero value otherwise.

### GetGivenNameOk

`func (o *IdemiaMobileIdProviderOutput) GetGivenNameOk() (*string, bool)`

GetGivenNameOk returns a tuple with the GivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenName

`func (o *IdemiaMobileIdProviderOutput) SetGivenName(v string)`

SetGivenName sets GivenName field to given value.

### HasGivenName

`func (o *IdemiaMobileIdProviderOutput) HasGivenName() bool`

HasGivenName returns a boolean if a field has been set.

### SetGivenNameNil

`func (o *IdemiaMobileIdProviderOutput) SetGivenNameNil(b bool)`

 SetGivenNameNil sets the value for GivenName to be an explicit nil

### UnsetGivenName
`func (o *IdemiaMobileIdProviderOutput) UnsetGivenName()`

UnsetGivenName ensures that no value is present for GivenName, not even an explicit nil
### GetFamilyName

`func (o *IdemiaMobileIdProviderOutput) GetFamilyName() string`

GetFamilyName returns the FamilyName field if non-nil, zero value otherwise.

### GetFamilyNameOk

`func (o *IdemiaMobileIdProviderOutput) GetFamilyNameOk() (*string, bool)`

GetFamilyNameOk returns a tuple with the FamilyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyName

`func (o *IdemiaMobileIdProviderOutput) SetFamilyName(v string)`

SetFamilyName sets FamilyName field to given value.

### HasFamilyName

`func (o *IdemiaMobileIdProviderOutput) HasFamilyName() bool`

HasFamilyName returns a boolean if a field has been set.

### SetFamilyNameNil

`func (o *IdemiaMobileIdProviderOutput) SetFamilyNameNil(b bool)`

 SetFamilyNameNil sets the value for FamilyName to be an explicit nil

### UnsetFamilyName
`func (o *IdemiaMobileIdProviderOutput) UnsetFamilyName()`

UnsetFamilyName ensures that no value is present for FamilyName, not even an explicit nil
### GetBirthDate

`func (o *IdemiaMobileIdProviderOutput) GetBirthDate() string`

GetBirthDate returns the BirthDate field if non-nil, zero value otherwise.

### GetBirthDateOk

`func (o *IdemiaMobileIdProviderOutput) GetBirthDateOk() (*string, bool)`

GetBirthDateOk returns a tuple with the BirthDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthDate

`func (o *IdemiaMobileIdProviderOutput) SetBirthDate(v string)`

SetBirthDate sets BirthDate field to given value.

### HasBirthDate

`func (o *IdemiaMobileIdProviderOutput) HasBirthDate() bool`

HasBirthDate returns a boolean if a field has been set.

### SetBirthDateNil

`func (o *IdemiaMobileIdProviderOutput) SetBirthDateNil(b bool)`

 SetBirthDateNil sets the value for BirthDate to be an explicit nil

### UnsetBirthDate
`func (o *IdemiaMobileIdProviderOutput) UnsetBirthDate()`

UnsetBirthDate ensures that no value is present for BirthDate, not even an explicit nil
### GetAgeOver18

`func (o *IdemiaMobileIdProviderOutput) GetAgeOver18() string`

GetAgeOver18 returns the AgeOver18 field if non-nil, zero value otherwise.

### GetAgeOver18Ok

`func (o *IdemiaMobileIdProviderOutput) GetAgeOver18Ok() (*string, bool)`

GetAgeOver18Ok returns a tuple with the AgeOver18 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgeOver18

`func (o *IdemiaMobileIdProviderOutput) SetAgeOver18(v string)`

SetAgeOver18 sets AgeOver18 field to given value.

### HasAgeOver18

`func (o *IdemiaMobileIdProviderOutput) HasAgeOver18() bool`

HasAgeOver18 returns a boolean if a field has been set.

### SetAgeOver18Nil

`func (o *IdemiaMobileIdProviderOutput) SetAgeOver18Nil(b bool)`

 SetAgeOver18Nil sets the value for AgeOver18 to be an explicit nil

### UnsetAgeOver18
`func (o *IdemiaMobileIdProviderOutput) UnsetAgeOver18()`

UnsetAgeOver18 ensures that no value is present for AgeOver18, not even an explicit nil
### GetAgeOver21

`func (o *IdemiaMobileIdProviderOutput) GetAgeOver21() string`

GetAgeOver21 returns the AgeOver21 field if non-nil, zero value otherwise.

### GetAgeOver21Ok

`func (o *IdemiaMobileIdProviderOutput) GetAgeOver21Ok() (*string, bool)`

GetAgeOver21Ok returns a tuple with the AgeOver21 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgeOver21

`func (o *IdemiaMobileIdProviderOutput) SetAgeOver21(v string)`

SetAgeOver21 sets AgeOver21 field to given value.

### HasAgeOver21

`func (o *IdemiaMobileIdProviderOutput) HasAgeOver21() bool`

HasAgeOver21 returns a boolean if a field has been set.

### SetAgeOver21Nil

`func (o *IdemiaMobileIdProviderOutput) SetAgeOver21Nil(b bool)`

 SetAgeOver21Nil sets the value for AgeOver21 to be an explicit nil

### UnsetAgeOver21
`func (o *IdemiaMobileIdProviderOutput) UnsetAgeOver21()`

UnsetAgeOver21 ensures that no value is present for AgeOver21, not even an explicit nil
### GetAgeOver25

`func (o *IdemiaMobileIdProviderOutput) GetAgeOver25() string`

GetAgeOver25 returns the AgeOver25 field if non-nil, zero value otherwise.

### GetAgeOver25Ok

`func (o *IdemiaMobileIdProviderOutput) GetAgeOver25Ok() (*string, bool)`

GetAgeOver25Ok returns a tuple with the AgeOver25 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgeOver25

`func (o *IdemiaMobileIdProviderOutput) SetAgeOver25(v string)`

SetAgeOver25 sets AgeOver25 field to given value.

### HasAgeOver25

`func (o *IdemiaMobileIdProviderOutput) HasAgeOver25() bool`

HasAgeOver25 returns a boolean if a field has been set.

### SetAgeOver25Nil

`func (o *IdemiaMobileIdProviderOutput) SetAgeOver25Nil(b bool)`

 SetAgeOver25Nil sets the value for AgeOver25 to be an explicit nil

### UnsetAgeOver25
`func (o *IdemiaMobileIdProviderOutput) UnsetAgeOver25()`

UnsetAgeOver25 ensures that no value is present for AgeOver25, not even an explicit nil
### GetResidentAddress

`func (o *IdemiaMobileIdProviderOutput) GetResidentAddress() string`

GetResidentAddress returns the ResidentAddress field if non-nil, zero value otherwise.

### GetResidentAddressOk

`func (o *IdemiaMobileIdProviderOutput) GetResidentAddressOk() (*string, bool)`

GetResidentAddressOk returns a tuple with the ResidentAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidentAddress

`func (o *IdemiaMobileIdProviderOutput) SetResidentAddress(v string)`

SetResidentAddress sets ResidentAddress field to given value.

### HasResidentAddress

`func (o *IdemiaMobileIdProviderOutput) HasResidentAddress() bool`

HasResidentAddress returns a boolean if a field has been set.

### SetResidentAddressNil

`func (o *IdemiaMobileIdProviderOutput) SetResidentAddressNil(b bool)`

 SetResidentAddressNil sets the value for ResidentAddress to be an explicit nil

### UnsetResidentAddress
`func (o *IdemiaMobileIdProviderOutput) UnsetResidentAddress()`

UnsetResidentAddress ensures that no value is present for ResidentAddress, not even an explicit nil
### GetResidentState

`func (o *IdemiaMobileIdProviderOutput) GetResidentState() string`

GetResidentState returns the ResidentState field if non-nil, zero value otherwise.

### GetResidentStateOk

`func (o *IdemiaMobileIdProviderOutput) GetResidentStateOk() (*string, bool)`

GetResidentStateOk returns a tuple with the ResidentState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidentState

`func (o *IdemiaMobileIdProviderOutput) SetResidentState(v string)`

SetResidentState sets ResidentState field to given value.

### HasResidentState

`func (o *IdemiaMobileIdProviderOutput) HasResidentState() bool`

HasResidentState returns a boolean if a field has been set.

### SetResidentStateNil

`func (o *IdemiaMobileIdProviderOutput) SetResidentStateNil(b bool)`

 SetResidentStateNil sets the value for ResidentState to be an explicit nil

### UnsetResidentState
`func (o *IdemiaMobileIdProviderOutput) UnsetResidentState()`

UnsetResidentState ensures that no value is present for ResidentState, not even an explicit nil
### GetResidentCountry

`func (o *IdemiaMobileIdProviderOutput) GetResidentCountry() string`

GetResidentCountry returns the ResidentCountry field if non-nil, zero value otherwise.

### GetResidentCountryOk

`func (o *IdemiaMobileIdProviderOutput) GetResidentCountryOk() (*string, bool)`

GetResidentCountryOk returns a tuple with the ResidentCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidentCountry

`func (o *IdemiaMobileIdProviderOutput) SetResidentCountry(v string)`

SetResidentCountry sets ResidentCountry field to given value.

### HasResidentCountry

`func (o *IdemiaMobileIdProviderOutput) HasResidentCountry() bool`

HasResidentCountry returns a boolean if a field has been set.

### SetResidentCountryNil

`func (o *IdemiaMobileIdProviderOutput) SetResidentCountryNil(b bool)`

 SetResidentCountryNil sets the value for ResidentCountry to be an explicit nil

### UnsetResidentCountry
`func (o *IdemiaMobileIdProviderOutput) UnsetResidentCountry()`

UnsetResidentCountry ensures that no value is present for ResidentCountry, not even an explicit nil
### GetResidentPostalCode

`func (o *IdemiaMobileIdProviderOutput) GetResidentPostalCode() string`

GetResidentPostalCode returns the ResidentPostalCode field if non-nil, zero value otherwise.

### GetResidentPostalCodeOk

`func (o *IdemiaMobileIdProviderOutput) GetResidentPostalCodeOk() (*string, bool)`

GetResidentPostalCodeOk returns a tuple with the ResidentPostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidentPostalCode

`func (o *IdemiaMobileIdProviderOutput) SetResidentPostalCode(v string)`

SetResidentPostalCode sets ResidentPostalCode field to given value.

### HasResidentPostalCode

`func (o *IdemiaMobileIdProviderOutput) HasResidentPostalCode() bool`

HasResidentPostalCode returns a boolean if a field has been set.

### SetResidentPostalCodeNil

`func (o *IdemiaMobileIdProviderOutput) SetResidentPostalCodeNil(b bool)`

 SetResidentPostalCodeNil sets the value for ResidentPostalCode to be an explicit nil

### UnsetResidentPostalCode
`func (o *IdemiaMobileIdProviderOutput) UnsetResidentPostalCode()`

UnsetResidentPostalCode ensures that no value is present for ResidentPostalCode, not even an explicit nil
### GetResidentCity

`func (o *IdemiaMobileIdProviderOutput) GetResidentCity() string`

GetResidentCity returns the ResidentCity field if non-nil, zero value otherwise.

### GetResidentCityOk

`func (o *IdemiaMobileIdProviderOutput) GetResidentCityOk() (*string, bool)`

GetResidentCityOk returns a tuple with the ResidentCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidentCity

`func (o *IdemiaMobileIdProviderOutput) SetResidentCity(v string)`

SetResidentCity sets ResidentCity field to given value.

### HasResidentCity

`func (o *IdemiaMobileIdProviderOutput) HasResidentCity() bool`

HasResidentCity returns a boolean if a field has been set.

### SetResidentCityNil

`func (o *IdemiaMobileIdProviderOutput) SetResidentCityNil(b bool)`

 SetResidentCityNil sets the value for ResidentCity to be an explicit nil

### UnsetResidentCity
`func (o *IdemiaMobileIdProviderOutput) UnsetResidentCity()`

UnsetResidentCity ensures that no value is present for ResidentCity, not even an explicit nil
### GetEmail

`func (o *IdemiaMobileIdProviderOutput) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *IdemiaMobileIdProviderOutput) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *IdemiaMobileIdProviderOutput) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *IdemiaMobileIdProviderOutput) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *IdemiaMobileIdProviderOutput) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *IdemiaMobileIdProviderOutput) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetEmailVerified

`func (o *IdemiaMobileIdProviderOutput) GetEmailVerified() bool`

GetEmailVerified returns the EmailVerified field if non-nil, zero value otherwise.

### GetEmailVerifiedOk

`func (o *IdemiaMobileIdProviderOutput) GetEmailVerifiedOk() (*bool, bool)`

GetEmailVerifiedOk returns a tuple with the EmailVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailVerified

`func (o *IdemiaMobileIdProviderOutput) SetEmailVerified(v bool)`

SetEmailVerified sets EmailVerified field to given value.

### HasEmailVerified

`func (o *IdemiaMobileIdProviderOutput) HasEmailVerified() bool`

HasEmailVerified returns a boolean if a field has been set.

### SetEmailVerifiedNil

`func (o *IdemiaMobileIdProviderOutput) SetEmailVerifiedNil(b bool)`

 SetEmailVerifiedNil sets the value for EmailVerified to be an explicit nil

### UnsetEmailVerified
`func (o *IdemiaMobileIdProviderOutput) UnsetEmailVerified()`

UnsetEmailVerified ensures that no value is present for EmailVerified, not even an explicit nil
### GetDocumentNumber

`func (o *IdemiaMobileIdProviderOutput) GetDocumentNumber() string`

GetDocumentNumber returns the DocumentNumber field if non-nil, zero value otherwise.

### GetDocumentNumberOk

`func (o *IdemiaMobileIdProviderOutput) GetDocumentNumberOk() (*string, bool)`

GetDocumentNumberOk returns a tuple with the DocumentNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentNumber

`func (o *IdemiaMobileIdProviderOutput) SetDocumentNumber(v string)`

SetDocumentNumber sets DocumentNumber field to given value.

### HasDocumentNumber

`func (o *IdemiaMobileIdProviderOutput) HasDocumentNumber() bool`

HasDocumentNumber returns a boolean if a field has been set.

### SetDocumentNumberNil

`func (o *IdemiaMobileIdProviderOutput) SetDocumentNumberNil(b bool)`

 SetDocumentNumberNil sets the value for DocumentNumber to be an explicit nil

### UnsetDocumentNumber
`func (o *IdemiaMobileIdProviderOutput) UnsetDocumentNumber()`

UnsetDocumentNumber ensures that no value is present for DocumentNumber, not even an explicit nil
### GetIssueDate

`func (o *IdemiaMobileIdProviderOutput) GetIssueDate() string`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *IdemiaMobileIdProviderOutput) GetIssueDateOk() (*string, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *IdemiaMobileIdProviderOutput) SetIssueDate(v string)`

SetIssueDate sets IssueDate field to given value.

### HasIssueDate

`func (o *IdemiaMobileIdProviderOutput) HasIssueDate() bool`

HasIssueDate returns a boolean if a field has been set.

### SetIssueDateNil

`func (o *IdemiaMobileIdProviderOutput) SetIssueDateNil(b bool)`

 SetIssueDateNil sets the value for IssueDate to be an explicit nil

### UnsetIssueDate
`func (o *IdemiaMobileIdProviderOutput) UnsetIssueDate()`

UnsetIssueDate ensures that no value is present for IssueDate, not even an explicit nil
### GetIssuingCountry

`func (o *IdemiaMobileIdProviderOutput) GetIssuingCountry() string`

GetIssuingCountry returns the IssuingCountry field if non-nil, zero value otherwise.

### GetIssuingCountryOk

`func (o *IdemiaMobileIdProviderOutput) GetIssuingCountryOk() (*string, bool)`

GetIssuingCountryOk returns a tuple with the IssuingCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingCountry

`func (o *IdemiaMobileIdProviderOutput) SetIssuingCountry(v string)`

SetIssuingCountry sets IssuingCountry field to given value.

### HasIssuingCountry

`func (o *IdemiaMobileIdProviderOutput) HasIssuingCountry() bool`

HasIssuingCountry returns a boolean if a field has been set.

### SetIssuingCountryNil

`func (o *IdemiaMobileIdProviderOutput) SetIssuingCountryNil(b bool)`

 SetIssuingCountryNil sets the value for IssuingCountry to be an explicit nil

### UnsetIssuingCountry
`func (o *IdemiaMobileIdProviderOutput) UnsetIssuingCountry()`

UnsetIssuingCountry ensures that no value is present for IssuingCountry, not even an explicit nil
### GetIssuingAuthority

`func (o *IdemiaMobileIdProviderOutput) GetIssuingAuthority() string`

GetIssuingAuthority returns the IssuingAuthority field if non-nil, zero value otherwise.

### GetIssuingAuthorityOk

`func (o *IdemiaMobileIdProviderOutput) GetIssuingAuthorityOk() (*string, bool)`

GetIssuingAuthorityOk returns a tuple with the IssuingAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingAuthority

`func (o *IdemiaMobileIdProviderOutput) SetIssuingAuthority(v string)`

SetIssuingAuthority sets IssuingAuthority field to given value.

### HasIssuingAuthority

`func (o *IdemiaMobileIdProviderOutput) HasIssuingAuthority() bool`

HasIssuingAuthority returns a boolean if a field has been set.

### SetIssuingAuthorityNil

`func (o *IdemiaMobileIdProviderOutput) SetIssuingAuthorityNil(b bool)`

 SetIssuingAuthorityNil sets the value for IssuingAuthority to be an explicit nil

### UnsetIssuingAuthority
`func (o *IdemiaMobileIdProviderOutput) UnsetIssuingAuthority()`

UnsetIssuingAuthority ensures that no value is present for IssuingAuthority, not even an explicit nil
### GetExpiryDate

`func (o *IdemiaMobileIdProviderOutput) GetExpiryDate() string`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *IdemiaMobileIdProviderOutput) GetExpiryDateOk() (*string, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *IdemiaMobileIdProviderOutput) SetExpiryDate(v string)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *IdemiaMobileIdProviderOutput) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### SetExpiryDateNil

`func (o *IdemiaMobileIdProviderOutput) SetExpiryDateNil(b bool)`

 SetExpiryDateNil sets the value for ExpiryDate to be an explicit nil

### UnsetExpiryDate
`func (o *IdemiaMobileIdProviderOutput) UnsetExpiryDate()`

UnsetExpiryDate ensures that no value is present for ExpiryDate, not even an explicit nil
### GetIssuingJurisdiction

`func (o *IdemiaMobileIdProviderOutput) GetIssuingJurisdiction() string`

GetIssuingJurisdiction returns the IssuingJurisdiction field if non-nil, zero value otherwise.

### GetIssuingJurisdictionOk

`func (o *IdemiaMobileIdProviderOutput) GetIssuingJurisdictionOk() (*string, bool)`

GetIssuingJurisdictionOk returns a tuple with the IssuingJurisdiction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingJurisdiction

`func (o *IdemiaMobileIdProviderOutput) SetIssuingJurisdiction(v string)`

SetIssuingJurisdiction sets IssuingJurisdiction field to given value.

### HasIssuingJurisdiction

`func (o *IdemiaMobileIdProviderOutput) HasIssuingJurisdiction() bool`

HasIssuingJurisdiction returns a boolean if a field has been set.

### SetIssuingJurisdictionNil

`func (o *IdemiaMobileIdProviderOutput) SetIssuingJurisdictionNil(b bool)`

 SetIssuingJurisdictionNil sets the value for IssuingJurisdiction to be an explicit nil

### UnsetIssuingJurisdiction
`func (o *IdemiaMobileIdProviderOutput) UnsetIssuingJurisdiction()`

UnsetIssuingJurisdiction ensures that no value is present for IssuingJurisdiction, not even an explicit nil
### GetDocumentType

`func (o *IdemiaMobileIdProviderOutput) GetDocumentType() string`

GetDocumentType returns the DocumentType field if non-nil, zero value otherwise.

### GetDocumentTypeOk

`func (o *IdemiaMobileIdProviderOutput) GetDocumentTypeOk() (*string, bool)`

GetDocumentTypeOk returns a tuple with the DocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentType

`func (o *IdemiaMobileIdProviderOutput) SetDocumentType(v string)`

SetDocumentType sets DocumentType field to given value.

### HasDocumentType

`func (o *IdemiaMobileIdProviderOutput) HasDocumentType() bool`

HasDocumentType returns a boolean if a field has been set.

### SetDocumentTypeNil

`func (o *IdemiaMobileIdProviderOutput) SetDocumentTypeNil(b bool)`

 SetDocumentTypeNil sets the value for DocumentType to be an explicit nil

### UnsetDocumentType
`func (o *IdemiaMobileIdProviderOutput) UnsetDocumentType()`

UnsetDocumentType ensures that no value is present for DocumentType, not even an explicit nil
### GetSex

`func (o *IdemiaMobileIdProviderOutput) GetSex() string`

GetSex returns the Sex field if non-nil, zero value otherwise.

### GetSexOk

`func (o *IdemiaMobileIdProviderOutput) GetSexOk() (*string, bool)`

GetSexOk returns a tuple with the Sex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSex

`func (o *IdemiaMobileIdProviderOutput) SetSex(v string)`

SetSex sets Sex field to given value.

### HasSex

`func (o *IdemiaMobileIdProviderOutput) HasSex() bool`

HasSex returns a boolean if a field has been set.

### SetSexNil

`func (o *IdemiaMobileIdProviderOutput) SetSexNil(b bool)`

 SetSexNil sets the value for Sex to be an explicit nil

### UnsetSex
`func (o *IdemiaMobileIdProviderOutput) UnsetSex()`

UnsetSex ensures that no value is present for Sex, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


