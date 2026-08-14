# EvrotrustProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Names** | Pointer to **NullableString** | The full name of the individual. | [optional] 
**LatinNames** | Pointer to **NullableString** | The full Latin-script name of the individual. | [optional] 
**FirstName** | Pointer to **NullableString** | The given name of the individual. | [optional] 
**MiddleName** | Pointer to **NullableString** | The middle name of the individual. | [optional] 
**LastName** | Pointer to **NullableString** | The family name of the individual. | [optional] 
**FirstNameLatin** | Pointer to **NullableString** | The Latin-script given name of the individual. | [optional] 
**MiddleNameLatin** | Pointer to **NullableString** | The Latin-script middle name of the individual. | [optional] 
**LastNameLatin** | Pointer to **NullableString** | The Latin-script family name of the individual. | [optional] 
**Nationality** | Pointer to **NullableString** | The nationality label of the individual. | [optional] 
**NationalityCountry** | Pointer to **NullableString** | The nationality country of the individual. | [optional] 
**NationalityCountryCode** | Pointer to **NullableString** | The nationality country code of the individual.              Evrotrust uses ISO 3166-1 alpha-3 country codes. | [optional] 
**Gender** | Pointer to **NullableString** | The gender label of the individual.              Known values in Evrotrust API docs: - Male: the individual is male. - Female: the individual is female. | [optional] 
**GenderLatin** | Pointer to **NullableString** | The Latin-script gender label of the individual.              Known values in Evrotrust API docs: - Male: the individual is male. - Female: the individual is female. | [optional] 
**UserGender** | Pointer to **NullableString** | The gender value from Evrotrust&#39;s user data object.              Evrotrust documents this separately from the identification metadata gender labels. Known values: - Male: the individual is male. - Female: the individual is female. | [optional] 
**DateOfBirth** | Pointer to **NullableString** | The date of birth of the individual in YYYY-MM-DD format. | [optional] 
**PlaceOfBirth** | Pointer to **NullableString** | The place of birth of the individual, if available. | [optional] 
**IdentificationNumber** | Pointer to **NullableString** | The national identification number of the individual.              The format depends on the verified document and issuing country. | [optional] 
**DocumentType** | Pointer to **NullableString** | The localized document type label of the identity document.              For Evrotrust&#39;s canonical document type code, see &#x60;docType&#x60;. | [optional] 
**DocType** | Pointer to **NullableString** | The canonical Evrotrust document type code.              Known values: - IDcard: an identity card. - Passport: a passport. | [optional] 
**DocumentNumber** | Pointer to **NullableString** | The identity document number. | [optional] 
**DocumentIssuerName** | Pointer to **NullableString** | The name of the authority that issued the identity document. | [optional] 
**DocumentIssueDate** | Pointer to **NullableString** | The identity document issue date in YYYY-MM-DD format.              Evrotrust documents timestamp &#x60;0&#x60; for this field as no value; it is omitted from provider-specific output. | [optional] 
**DocumentValidDate** | Pointer to **NullableString** | The identity document expiration date in YYYY-MM-DD format.              Evrotrust documents timestamp &#x60;0&#x60; for this field as no value; it is omitted from provider-specific output. | [optional] 
**DocumentCountry** | Pointer to **NullableString** | The identity document issuing country. | [optional] 
**DocumentCountryCode** | Pointer to **NullableString** | The identity document issuing country code.              Evrotrust uses ISO 3166-1 alpha-3 country codes. | [optional] 
**Address** | Pointer to **NullableString** | The full address of the individual. | [optional] 
**AddressComponents** | Pointer to [**NullableEvrotrustAddressComponents**](EvrotrustAddressComponents.md) | The structured address fragments of the individual, if available. | [optional] 
**EmailAddresses** | Pointer to **[]string** | The email addresses of the individual, if available. | [optional] 
**PhoneNumbers** | Pointer to **[]string** | The phone numbers of the individual in E.164 format, if available. | [optional] 
**IdentificationReason** | Pointer to **NullableString** | The relying party-provided reason for requesting identification. | [optional] 
**IdentificationBefore** | Pointer to **NullableString** | The deadline or contextual value supplied for the identification request. | [optional] 

## Methods

### NewEvrotrustProviderOutput

`func NewEvrotrustProviderOutput() *EvrotrustProviderOutput`

NewEvrotrustProviderOutput instantiates a new EvrotrustProviderOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvrotrustProviderOutputWithDefaults

`func NewEvrotrustProviderOutputWithDefaults() *EvrotrustProviderOutput`

NewEvrotrustProviderOutputWithDefaults instantiates a new EvrotrustProviderOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNames

`func (o *EvrotrustProviderOutput) GetNames() string`

GetNames returns the Names field if non-nil, zero value otherwise.

### GetNamesOk

`func (o *EvrotrustProviderOutput) GetNamesOk() (*string, bool)`

GetNamesOk returns a tuple with the Names field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNames

`func (o *EvrotrustProviderOutput) SetNames(v string)`

SetNames sets Names field to given value.

### HasNames

`func (o *EvrotrustProviderOutput) HasNames() bool`

HasNames returns a boolean if a field has been set.

### SetNamesNil

`func (o *EvrotrustProviderOutput) SetNamesNil(b bool)`

 SetNamesNil sets the value for Names to be an explicit nil

### UnsetNames
`func (o *EvrotrustProviderOutput) UnsetNames()`

UnsetNames ensures that no value is present for Names, not even an explicit nil
### GetLatinNames

`func (o *EvrotrustProviderOutput) GetLatinNames() string`

GetLatinNames returns the LatinNames field if non-nil, zero value otherwise.

### GetLatinNamesOk

`func (o *EvrotrustProviderOutput) GetLatinNamesOk() (*string, bool)`

GetLatinNamesOk returns a tuple with the LatinNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatinNames

`func (o *EvrotrustProviderOutput) SetLatinNames(v string)`

SetLatinNames sets LatinNames field to given value.

### HasLatinNames

`func (o *EvrotrustProviderOutput) HasLatinNames() bool`

HasLatinNames returns a boolean if a field has been set.

### SetLatinNamesNil

`func (o *EvrotrustProviderOutput) SetLatinNamesNil(b bool)`

 SetLatinNamesNil sets the value for LatinNames to be an explicit nil

### UnsetLatinNames
`func (o *EvrotrustProviderOutput) UnsetLatinNames()`

UnsetLatinNames ensures that no value is present for LatinNames, not even an explicit nil
### GetFirstName

`func (o *EvrotrustProviderOutput) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *EvrotrustProviderOutput) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *EvrotrustProviderOutput) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *EvrotrustProviderOutput) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *EvrotrustProviderOutput) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *EvrotrustProviderOutput) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetMiddleName

`func (o *EvrotrustProviderOutput) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *EvrotrustProviderOutput) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *EvrotrustProviderOutput) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *EvrotrustProviderOutput) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### SetMiddleNameNil

`func (o *EvrotrustProviderOutput) SetMiddleNameNil(b bool)`

 SetMiddleNameNil sets the value for MiddleName to be an explicit nil

### UnsetMiddleName
`func (o *EvrotrustProviderOutput) UnsetMiddleName()`

UnsetMiddleName ensures that no value is present for MiddleName, not even an explicit nil
### GetLastName

`func (o *EvrotrustProviderOutput) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *EvrotrustProviderOutput) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *EvrotrustProviderOutput) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *EvrotrustProviderOutput) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *EvrotrustProviderOutput) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *EvrotrustProviderOutput) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetFirstNameLatin

`func (o *EvrotrustProviderOutput) GetFirstNameLatin() string`

GetFirstNameLatin returns the FirstNameLatin field if non-nil, zero value otherwise.

### GetFirstNameLatinOk

`func (o *EvrotrustProviderOutput) GetFirstNameLatinOk() (*string, bool)`

GetFirstNameLatinOk returns a tuple with the FirstNameLatin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstNameLatin

`func (o *EvrotrustProviderOutput) SetFirstNameLatin(v string)`

SetFirstNameLatin sets FirstNameLatin field to given value.

### HasFirstNameLatin

`func (o *EvrotrustProviderOutput) HasFirstNameLatin() bool`

HasFirstNameLatin returns a boolean if a field has been set.

### SetFirstNameLatinNil

`func (o *EvrotrustProviderOutput) SetFirstNameLatinNil(b bool)`

 SetFirstNameLatinNil sets the value for FirstNameLatin to be an explicit nil

### UnsetFirstNameLatin
`func (o *EvrotrustProviderOutput) UnsetFirstNameLatin()`

UnsetFirstNameLatin ensures that no value is present for FirstNameLatin, not even an explicit nil
### GetMiddleNameLatin

`func (o *EvrotrustProviderOutput) GetMiddleNameLatin() string`

GetMiddleNameLatin returns the MiddleNameLatin field if non-nil, zero value otherwise.

### GetMiddleNameLatinOk

`func (o *EvrotrustProviderOutput) GetMiddleNameLatinOk() (*string, bool)`

GetMiddleNameLatinOk returns a tuple with the MiddleNameLatin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleNameLatin

`func (o *EvrotrustProviderOutput) SetMiddleNameLatin(v string)`

SetMiddleNameLatin sets MiddleNameLatin field to given value.

### HasMiddleNameLatin

`func (o *EvrotrustProviderOutput) HasMiddleNameLatin() bool`

HasMiddleNameLatin returns a boolean if a field has been set.

### SetMiddleNameLatinNil

`func (o *EvrotrustProviderOutput) SetMiddleNameLatinNil(b bool)`

 SetMiddleNameLatinNil sets the value for MiddleNameLatin to be an explicit nil

### UnsetMiddleNameLatin
`func (o *EvrotrustProviderOutput) UnsetMiddleNameLatin()`

UnsetMiddleNameLatin ensures that no value is present for MiddleNameLatin, not even an explicit nil
### GetLastNameLatin

`func (o *EvrotrustProviderOutput) GetLastNameLatin() string`

GetLastNameLatin returns the LastNameLatin field if non-nil, zero value otherwise.

### GetLastNameLatinOk

`func (o *EvrotrustProviderOutput) GetLastNameLatinOk() (*string, bool)`

GetLastNameLatinOk returns a tuple with the LastNameLatin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastNameLatin

`func (o *EvrotrustProviderOutput) SetLastNameLatin(v string)`

SetLastNameLatin sets LastNameLatin field to given value.

### HasLastNameLatin

`func (o *EvrotrustProviderOutput) HasLastNameLatin() bool`

HasLastNameLatin returns a boolean if a field has been set.

### SetLastNameLatinNil

`func (o *EvrotrustProviderOutput) SetLastNameLatinNil(b bool)`

 SetLastNameLatinNil sets the value for LastNameLatin to be an explicit nil

### UnsetLastNameLatin
`func (o *EvrotrustProviderOutput) UnsetLastNameLatin()`

UnsetLastNameLatin ensures that no value is present for LastNameLatin, not even an explicit nil
### GetNationality

`func (o *EvrotrustProviderOutput) GetNationality() string`

GetNationality returns the Nationality field if non-nil, zero value otherwise.

### GetNationalityOk

`func (o *EvrotrustProviderOutput) GetNationalityOk() (*string, bool)`

GetNationalityOk returns a tuple with the Nationality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationality

`func (o *EvrotrustProviderOutput) SetNationality(v string)`

SetNationality sets Nationality field to given value.

### HasNationality

`func (o *EvrotrustProviderOutput) HasNationality() bool`

HasNationality returns a boolean if a field has been set.

### SetNationalityNil

`func (o *EvrotrustProviderOutput) SetNationalityNil(b bool)`

 SetNationalityNil sets the value for Nationality to be an explicit nil

### UnsetNationality
`func (o *EvrotrustProviderOutput) UnsetNationality()`

UnsetNationality ensures that no value is present for Nationality, not even an explicit nil
### GetNationalityCountry

`func (o *EvrotrustProviderOutput) GetNationalityCountry() string`

GetNationalityCountry returns the NationalityCountry field if non-nil, zero value otherwise.

### GetNationalityCountryOk

`func (o *EvrotrustProviderOutput) GetNationalityCountryOk() (*string, bool)`

GetNationalityCountryOk returns a tuple with the NationalityCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationalityCountry

`func (o *EvrotrustProviderOutput) SetNationalityCountry(v string)`

SetNationalityCountry sets NationalityCountry field to given value.

### HasNationalityCountry

`func (o *EvrotrustProviderOutput) HasNationalityCountry() bool`

HasNationalityCountry returns a boolean if a field has been set.

### SetNationalityCountryNil

`func (o *EvrotrustProviderOutput) SetNationalityCountryNil(b bool)`

 SetNationalityCountryNil sets the value for NationalityCountry to be an explicit nil

### UnsetNationalityCountry
`func (o *EvrotrustProviderOutput) UnsetNationalityCountry()`

UnsetNationalityCountry ensures that no value is present for NationalityCountry, not even an explicit nil
### GetNationalityCountryCode

`func (o *EvrotrustProviderOutput) GetNationalityCountryCode() string`

GetNationalityCountryCode returns the NationalityCountryCode field if non-nil, zero value otherwise.

### GetNationalityCountryCodeOk

`func (o *EvrotrustProviderOutput) GetNationalityCountryCodeOk() (*string, bool)`

GetNationalityCountryCodeOk returns a tuple with the NationalityCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationalityCountryCode

`func (o *EvrotrustProviderOutput) SetNationalityCountryCode(v string)`

SetNationalityCountryCode sets NationalityCountryCode field to given value.

### HasNationalityCountryCode

`func (o *EvrotrustProviderOutput) HasNationalityCountryCode() bool`

HasNationalityCountryCode returns a boolean if a field has been set.

### SetNationalityCountryCodeNil

`func (o *EvrotrustProviderOutput) SetNationalityCountryCodeNil(b bool)`

 SetNationalityCountryCodeNil sets the value for NationalityCountryCode to be an explicit nil

### UnsetNationalityCountryCode
`func (o *EvrotrustProviderOutput) UnsetNationalityCountryCode()`

UnsetNationalityCountryCode ensures that no value is present for NationalityCountryCode, not even an explicit nil
### GetGender

`func (o *EvrotrustProviderOutput) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *EvrotrustProviderOutput) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *EvrotrustProviderOutput) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *EvrotrustProviderOutput) HasGender() bool`

HasGender returns a boolean if a field has been set.

### SetGenderNil

`func (o *EvrotrustProviderOutput) SetGenderNil(b bool)`

 SetGenderNil sets the value for Gender to be an explicit nil

### UnsetGender
`func (o *EvrotrustProviderOutput) UnsetGender()`

UnsetGender ensures that no value is present for Gender, not even an explicit nil
### GetGenderLatin

`func (o *EvrotrustProviderOutput) GetGenderLatin() string`

GetGenderLatin returns the GenderLatin field if non-nil, zero value otherwise.

### GetGenderLatinOk

`func (o *EvrotrustProviderOutput) GetGenderLatinOk() (*string, bool)`

GetGenderLatinOk returns a tuple with the GenderLatin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenderLatin

`func (o *EvrotrustProviderOutput) SetGenderLatin(v string)`

SetGenderLatin sets GenderLatin field to given value.

### HasGenderLatin

`func (o *EvrotrustProviderOutput) HasGenderLatin() bool`

HasGenderLatin returns a boolean if a field has been set.

### SetGenderLatinNil

`func (o *EvrotrustProviderOutput) SetGenderLatinNil(b bool)`

 SetGenderLatinNil sets the value for GenderLatin to be an explicit nil

### UnsetGenderLatin
`func (o *EvrotrustProviderOutput) UnsetGenderLatin()`

UnsetGenderLatin ensures that no value is present for GenderLatin, not even an explicit nil
### GetUserGender

`func (o *EvrotrustProviderOutput) GetUserGender() string`

GetUserGender returns the UserGender field if non-nil, zero value otherwise.

### GetUserGenderOk

`func (o *EvrotrustProviderOutput) GetUserGenderOk() (*string, bool)`

GetUserGenderOk returns a tuple with the UserGender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGender

`func (o *EvrotrustProviderOutput) SetUserGender(v string)`

SetUserGender sets UserGender field to given value.

### HasUserGender

`func (o *EvrotrustProviderOutput) HasUserGender() bool`

HasUserGender returns a boolean if a field has been set.

### SetUserGenderNil

`func (o *EvrotrustProviderOutput) SetUserGenderNil(b bool)`

 SetUserGenderNil sets the value for UserGender to be an explicit nil

### UnsetUserGender
`func (o *EvrotrustProviderOutput) UnsetUserGender()`

UnsetUserGender ensures that no value is present for UserGender, not even an explicit nil
### GetDateOfBirth

`func (o *EvrotrustProviderOutput) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *EvrotrustProviderOutput) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *EvrotrustProviderOutput) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *EvrotrustProviderOutput) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### SetDateOfBirthNil

`func (o *EvrotrustProviderOutput) SetDateOfBirthNil(b bool)`

 SetDateOfBirthNil sets the value for DateOfBirth to be an explicit nil

### UnsetDateOfBirth
`func (o *EvrotrustProviderOutput) UnsetDateOfBirth()`

UnsetDateOfBirth ensures that no value is present for DateOfBirth, not even an explicit nil
### GetPlaceOfBirth

`func (o *EvrotrustProviderOutput) GetPlaceOfBirth() string`

GetPlaceOfBirth returns the PlaceOfBirth field if non-nil, zero value otherwise.

### GetPlaceOfBirthOk

`func (o *EvrotrustProviderOutput) GetPlaceOfBirthOk() (*string, bool)`

GetPlaceOfBirthOk returns a tuple with the PlaceOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceOfBirth

`func (o *EvrotrustProviderOutput) SetPlaceOfBirth(v string)`

SetPlaceOfBirth sets PlaceOfBirth field to given value.

### HasPlaceOfBirth

`func (o *EvrotrustProviderOutput) HasPlaceOfBirth() bool`

HasPlaceOfBirth returns a boolean if a field has been set.

### SetPlaceOfBirthNil

`func (o *EvrotrustProviderOutput) SetPlaceOfBirthNil(b bool)`

 SetPlaceOfBirthNil sets the value for PlaceOfBirth to be an explicit nil

### UnsetPlaceOfBirth
`func (o *EvrotrustProviderOutput) UnsetPlaceOfBirth()`

UnsetPlaceOfBirth ensures that no value is present for PlaceOfBirth, not even an explicit nil
### GetIdentificationNumber

`func (o *EvrotrustProviderOutput) GetIdentificationNumber() string`

GetIdentificationNumber returns the IdentificationNumber field if non-nil, zero value otherwise.

### GetIdentificationNumberOk

`func (o *EvrotrustProviderOutput) GetIdentificationNumberOk() (*string, bool)`

GetIdentificationNumberOk returns a tuple with the IdentificationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentificationNumber

`func (o *EvrotrustProviderOutput) SetIdentificationNumber(v string)`

SetIdentificationNumber sets IdentificationNumber field to given value.

### HasIdentificationNumber

`func (o *EvrotrustProviderOutput) HasIdentificationNumber() bool`

HasIdentificationNumber returns a boolean if a field has been set.

### SetIdentificationNumberNil

`func (o *EvrotrustProviderOutput) SetIdentificationNumberNil(b bool)`

 SetIdentificationNumberNil sets the value for IdentificationNumber to be an explicit nil

### UnsetIdentificationNumber
`func (o *EvrotrustProviderOutput) UnsetIdentificationNumber()`

UnsetIdentificationNumber ensures that no value is present for IdentificationNumber, not even an explicit nil
### GetDocumentType

`func (o *EvrotrustProviderOutput) GetDocumentType() string`

GetDocumentType returns the DocumentType field if non-nil, zero value otherwise.

### GetDocumentTypeOk

`func (o *EvrotrustProviderOutput) GetDocumentTypeOk() (*string, bool)`

GetDocumentTypeOk returns a tuple with the DocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentType

`func (o *EvrotrustProviderOutput) SetDocumentType(v string)`

SetDocumentType sets DocumentType field to given value.

### HasDocumentType

`func (o *EvrotrustProviderOutput) HasDocumentType() bool`

HasDocumentType returns a boolean if a field has been set.

### SetDocumentTypeNil

`func (o *EvrotrustProviderOutput) SetDocumentTypeNil(b bool)`

 SetDocumentTypeNil sets the value for DocumentType to be an explicit nil

### UnsetDocumentType
`func (o *EvrotrustProviderOutput) UnsetDocumentType()`

UnsetDocumentType ensures that no value is present for DocumentType, not even an explicit nil
### GetDocType

`func (o *EvrotrustProviderOutput) GetDocType() string`

GetDocType returns the DocType field if non-nil, zero value otherwise.

### GetDocTypeOk

`func (o *EvrotrustProviderOutput) GetDocTypeOk() (*string, bool)`

GetDocTypeOk returns a tuple with the DocType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocType

`func (o *EvrotrustProviderOutput) SetDocType(v string)`

SetDocType sets DocType field to given value.

### HasDocType

`func (o *EvrotrustProviderOutput) HasDocType() bool`

HasDocType returns a boolean if a field has been set.

### SetDocTypeNil

`func (o *EvrotrustProviderOutput) SetDocTypeNil(b bool)`

 SetDocTypeNil sets the value for DocType to be an explicit nil

### UnsetDocType
`func (o *EvrotrustProviderOutput) UnsetDocType()`

UnsetDocType ensures that no value is present for DocType, not even an explicit nil
### GetDocumentNumber

`func (o *EvrotrustProviderOutput) GetDocumentNumber() string`

GetDocumentNumber returns the DocumentNumber field if non-nil, zero value otherwise.

### GetDocumentNumberOk

`func (o *EvrotrustProviderOutput) GetDocumentNumberOk() (*string, bool)`

GetDocumentNumberOk returns a tuple with the DocumentNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentNumber

`func (o *EvrotrustProviderOutput) SetDocumentNumber(v string)`

SetDocumentNumber sets DocumentNumber field to given value.

### HasDocumentNumber

`func (o *EvrotrustProviderOutput) HasDocumentNumber() bool`

HasDocumentNumber returns a boolean if a field has been set.

### SetDocumentNumberNil

`func (o *EvrotrustProviderOutput) SetDocumentNumberNil(b bool)`

 SetDocumentNumberNil sets the value for DocumentNumber to be an explicit nil

### UnsetDocumentNumber
`func (o *EvrotrustProviderOutput) UnsetDocumentNumber()`

UnsetDocumentNumber ensures that no value is present for DocumentNumber, not even an explicit nil
### GetDocumentIssuerName

`func (o *EvrotrustProviderOutput) GetDocumentIssuerName() string`

GetDocumentIssuerName returns the DocumentIssuerName field if non-nil, zero value otherwise.

### GetDocumentIssuerNameOk

`func (o *EvrotrustProviderOutput) GetDocumentIssuerNameOk() (*string, bool)`

GetDocumentIssuerNameOk returns a tuple with the DocumentIssuerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentIssuerName

`func (o *EvrotrustProviderOutput) SetDocumentIssuerName(v string)`

SetDocumentIssuerName sets DocumentIssuerName field to given value.

### HasDocumentIssuerName

`func (o *EvrotrustProviderOutput) HasDocumentIssuerName() bool`

HasDocumentIssuerName returns a boolean if a field has been set.

### SetDocumentIssuerNameNil

`func (o *EvrotrustProviderOutput) SetDocumentIssuerNameNil(b bool)`

 SetDocumentIssuerNameNil sets the value for DocumentIssuerName to be an explicit nil

### UnsetDocumentIssuerName
`func (o *EvrotrustProviderOutput) UnsetDocumentIssuerName()`

UnsetDocumentIssuerName ensures that no value is present for DocumentIssuerName, not even an explicit nil
### GetDocumentIssueDate

`func (o *EvrotrustProviderOutput) GetDocumentIssueDate() string`

GetDocumentIssueDate returns the DocumentIssueDate field if non-nil, zero value otherwise.

### GetDocumentIssueDateOk

`func (o *EvrotrustProviderOutput) GetDocumentIssueDateOk() (*string, bool)`

GetDocumentIssueDateOk returns a tuple with the DocumentIssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentIssueDate

`func (o *EvrotrustProviderOutput) SetDocumentIssueDate(v string)`

SetDocumentIssueDate sets DocumentIssueDate field to given value.

### HasDocumentIssueDate

`func (o *EvrotrustProviderOutput) HasDocumentIssueDate() bool`

HasDocumentIssueDate returns a boolean if a field has been set.

### SetDocumentIssueDateNil

`func (o *EvrotrustProviderOutput) SetDocumentIssueDateNil(b bool)`

 SetDocumentIssueDateNil sets the value for DocumentIssueDate to be an explicit nil

### UnsetDocumentIssueDate
`func (o *EvrotrustProviderOutput) UnsetDocumentIssueDate()`

UnsetDocumentIssueDate ensures that no value is present for DocumentIssueDate, not even an explicit nil
### GetDocumentValidDate

`func (o *EvrotrustProviderOutput) GetDocumentValidDate() string`

GetDocumentValidDate returns the DocumentValidDate field if non-nil, zero value otherwise.

### GetDocumentValidDateOk

`func (o *EvrotrustProviderOutput) GetDocumentValidDateOk() (*string, bool)`

GetDocumentValidDateOk returns a tuple with the DocumentValidDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentValidDate

`func (o *EvrotrustProviderOutput) SetDocumentValidDate(v string)`

SetDocumentValidDate sets DocumentValidDate field to given value.

### HasDocumentValidDate

`func (o *EvrotrustProviderOutput) HasDocumentValidDate() bool`

HasDocumentValidDate returns a boolean if a field has been set.

### SetDocumentValidDateNil

`func (o *EvrotrustProviderOutput) SetDocumentValidDateNil(b bool)`

 SetDocumentValidDateNil sets the value for DocumentValidDate to be an explicit nil

### UnsetDocumentValidDate
`func (o *EvrotrustProviderOutput) UnsetDocumentValidDate()`

UnsetDocumentValidDate ensures that no value is present for DocumentValidDate, not even an explicit nil
### GetDocumentCountry

`func (o *EvrotrustProviderOutput) GetDocumentCountry() string`

GetDocumentCountry returns the DocumentCountry field if non-nil, zero value otherwise.

### GetDocumentCountryOk

`func (o *EvrotrustProviderOutput) GetDocumentCountryOk() (*string, bool)`

GetDocumentCountryOk returns a tuple with the DocumentCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentCountry

`func (o *EvrotrustProviderOutput) SetDocumentCountry(v string)`

SetDocumentCountry sets DocumentCountry field to given value.

### HasDocumentCountry

`func (o *EvrotrustProviderOutput) HasDocumentCountry() bool`

HasDocumentCountry returns a boolean if a field has been set.

### SetDocumentCountryNil

`func (o *EvrotrustProviderOutput) SetDocumentCountryNil(b bool)`

 SetDocumentCountryNil sets the value for DocumentCountry to be an explicit nil

### UnsetDocumentCountry
`func (o *EvrotrustProviderOutput) UnsetDocumentCountry()`

UnsetDocumentCountry ensures that no value is present for DocumentCountry, not even an explicit nil
### GetDocumentCountryCode

`func (o *EvrotrustProviderOutput) GetDocumentCountryCode() string`

GetDocumentCountryCode returns the DocumentCountryCode field if non-nil, zero value otherwise.

### GetDocumentCountryCodeOk

`func (o *EvrotrustProviderOutput) GetDocumentCountryCodeOk() (*string, bool)`

GetDocumentCountryCodeOk returns a tuple with the DocumentCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentCountryCode

`func (o *EvrotrustProviderOutput) SetDocumentCountryCode(v string)`

SetDocumentCountryCode sets DocumentCountryCode field to given value.

### HasDocumentCountryCode

`func (o *EvrotrustProviderOutput) HasDocumentCountryCode() bool`

HasDocumentCountryCode returns a boolean if a field has been set.

### SetDocumentCountryCodeNil

`func (o *EvrotrustProviderOutput) SetDocumentCountryCodeNil(b bool)`

 SetDocumentCountryCodeNil sets the value for DocumentCountryCode to be an explicit nil

### UnsetDocumentCountryCode
`func (o *EvrotrustProviderOutput) UnsetDocumentCountryCode()`

UnsetDocumentCountryCode ensures that no value is present for DocumentCountryCode, not even an explicit nil
### GetAddress

`func (o *EvrotrustProviderOutput) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *EvrotrustProviderOutput) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *EvrotrustProviderOutput) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *EvrotrustProviderOutput) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *EvrotrustProviderOutput) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *EvrotrustProviderOutput) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetAddressComponents

`func (o *EvrotrustProviderOutput) GetAddressComponents() EvrotrustAddressComponents`

GetAddressComponents returns the AddressComponents field if non-nil, zero value otherwise.

### GetAddressComponentsOk

`func (o *EvrotrustProviderOutput) GetAddressComponentsOk() (*EvrotrustAddressComponents, bool)`

GetAddressComponentsOk returns a tuple with the AddressComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressComponents

`func (o *EvrotrustProviderOutput) SetAddressComponents(v EvrotrustAddressComponents)`

SetAddressComponents sets AddressComponents field to given value.

### HasAddressComponents

`func (o *EvrotrustProviderOutput) HasAddressComponents() bool`

HasAddressComponents returns a boolean if a field has been set.

### SetAddressComponentsNil

`func (o *EvrotrustProviderOutput) SetAddressComponentsNil(b bool)`

 SetAddressComponentsNil sets the value for AddressComponents to be an explicit nil

### UnsetAddressComponents
`func (o *EvrotrustProviderOutput) UnsetAddressComponents()`

UnsetAddressComponents ensures that no value is present for AddressComponents, not even an explicit nil
### GetEmailAddresses

`func (o *EvrotrustProviderOutput) GetEmailAddresses() []string`

GetEmailAddresses returns the EmailAddresses field if non-nil, zero value otherwise.

### GetEmailAddressesOk

`func (o *EvrotrustProviderOutput) GetEmailAddressesOk() (*[]string, bool)`

GetEmailAddressesOk returns a tuple with the EmailAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddresses

`func (o *EvrotrustProviderOutput) SetEmailAddresses(v []string)`

SetEmailAddresses sets EmailAddresses field to given value.

### HasEmailAddresses

`func (o *EvrotrustProviderOutput) HasEmailAddresses() bool`

HasEmailAddresses returns a boolean if a field has been set.

### SetEmailAddressesNil

`func (o *EvrotrustProviderOutput) SetEmailAddressesNil(b bool)`

 SetEmailAddressesNil sets the value for EmailAddresses to be an explicit nil

### UnsetEmailAddresses
`func (o *EvrotrustProviderOutput) UnsetEmailAddresses()`

UnsetEmailAddresses ensures that no value is present for EmailAddresses, not even an explicit nil
### GetPhoneNumbers

`func (o *EvrotrustProviderOutput) GetPhoneNumbers() []string`

GetPhoneNumbers returns the PhoneNumbers field if non-nil, zero value otherwise.

### GetPhoneNumbersOk

`func (o *EvrotrustProviderOutput) GetPhoneNumbersOk() (*[]string, bool)`

GetPhoneNumbersOk returns a tuple with the PhoneNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumbers

`func (o *EvrotrustProviderOutput) SetPhoneNumbers(v []string)`

SetPhoneNumbers sets PhoneNumbers field to given value.

### HasPhoneNumbers

`func (o *EvrotrustProviderOutput) HasPhoneNumbers() bool`

HasPhoneNumbers returns a boolean if a field has been set.

### SetPhoneNumbersNil

`func (o *EvrotrustProviderOutput) SetPhoneNumbersNil(b bool)`

 SetPhoneNumbersNil sets the value for PhoneNumbers to be an explicit nil

### UnsetPhoneNumbers
`func (o *EvrotrustProviderOutput) UnsetPhoneNumbers()`

UnsetPhoneNumbers ensures that no value is present for PhoneNumbers, not even an explicit nil
### GetIdentificationReason

`func (o *EvrotrustProviderOutput) GetIdentificationReason() string`

GetIdentificationReason returns the IdentificationReason field if non-nil, zero value otherwise.

### GetIdentificationReasonOk

`func (o *EvrotrustProviderOutput) GetIdentificationReasonOk() (*string, bool)`

GetIdentificationReasonOk returns a tuple with the IdentificationReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentificationReason

`func (o *EvrotrustProviderOutput) SetIdentificationReason(v string)`

SetIdentificationReason sets IdentificationReason field to given value.

### HasIdentificationReason

`func (o *EvrotrustProviderOutput) HasIdentificationReason() bool`

HasIdentificationReason returns a boolean if a field has been set.

### SetIdentificationReasonNil

`func (o *EvrotrustProviderOutput) SetIdentificationReasonNil(b bool)`

 SetIdentificationReasonNil sets the value for IdentificationReason to be an explicit nil

### UnsetIdentificationReason
`func (o *EvrotrustProviderOutput) UnsetIdentificationReason()`

UnsetIdentificationReason ensures that no value is present for IdentificationReason, not even an explicit nil
### GetIdentificationBefore

`func (o *EvrotrustProviderOutput) GetIdentificationBefore() string`

GetIdentificationBefore returns the IdentificationBefore field if non-nil, zero value otherwise.

### GetIdentificationBeforeOk

`func (o *EvrotrustProviderOutput) GetIdentificationBeforeOk() (*string, bool)`

GetIdentificationBeforeOk returns a tuple with the IdentificationBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentificationBefore

`func (o *EvrotrustProviderOutput) SetIdentificationBefore(v string)`

SetIdentificationBefore sets IdentificationBefore field to given value.

### HasIdentificationBefore

`func (o *EvrotrustProviderOutput) HasIdentificationBefore() bool`

HasIdentificationBefore returns a boolean if a field has been set.

### SetIdentificationBeforeNil

`func (o *EvrotrustProviderOutput) SetIdentificationBeforeNil(b bool)`

 SetIdentificationBeforeNil sets the value for IdentificationBefore to be an explicit nil

### UnsetIdentificationBefore
`func (o *EvrotrustProviderOutput) UnsetIdentificationBefore()`

UnsetIdentificationBefore ensures that no value is present for IdentificationBefore, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


