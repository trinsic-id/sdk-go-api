# PhilippinesDigitalNidProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PhilsysCardNumber** | **string** | The PhilSys Card Number (PCN). Every citizen or resident alien registered in PhilSys has a PhilSys Number (PSN). This number is tokenized into a card number to protect the PSN. The PhilSys Card Number is 12 characters long, and often is written in octets with dashes in between. | 
**GivenName** | Pointer to **NullableString** | The given (first) name of the individual. | [optional] 
**MiddleName** | Pointer to **NullableString** | The middle name of the individual. | [optional] 
**FamilyName** | Pointer to **NullableString** | The family (last) name of the individual. | [optional] 
**DateOfBirth** | Pointer to **NullableString** | The date of birth of the individual. | [optional] 
**Suffix** | Pointer to **NullableString** | The name suffix of the individual (e.g. Jr., III). | [optional] 
**Sex** | Pointer to **NullableString** | The sex of the individual.              Possible values: - Male - Female | [optional] 
**PlaceOfBirth** | Pointer to **NullableString** | The place of birth of the individual as recorded on the PhilSys credential. | [optional] 
**DocumentIssueDate** | Pointer to **NullableString** | The date the document was issued. | [optional] 
**DocumentNumber** | Pointer to **NullableString** | The document number of the digital National ID. | [optional] 
**IssuingAuthority** | Pointer to **NullableString** | The issuing authority of the document. | [optional] 

## Methods

### NewPhilippinesDigitalNidProviderOutput

`func NewPhilippinesDigitalNidProviderOutput(philsysCardNumber string, ) *PhilippinesDigitalNidProviderOutput`

NewPhilippinesDigitalNidProviderOutput instantiates a new PhilippinesDigitalNidProviderOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhilippinesDigitalNidProviderOutputWithDefaults

`func NewPhilippinesDigitalNidProviderOutputWithDefaults() *PhilippinesDigitalNidProviderOutput`

NewPhilippinesDigitalNidProviderOutputWithDefaults instantiates a new PhilippinesDigitalNidProviderOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhilsysCardNumber

`func (o *PhilippinesDigitalNidProviderOutput) GetPhilsysCardNumber() string`

GetPhilsysCardNumber returns the PhilsysCardNumber field if non-nil, zero value otherwise.

### GetPhilsysCardNumberOk

`func (o *PhilippinesDigitalNidProviderOutput) GetPhilsysCardNumberOk() (*string, bool)`

GetPhilsysCardNumberOk returns a tuple with the PhilsysCardNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhilsysCardNumber

`func (o *PhilippinesDigitalNidProviderOutput) SetPhilsysCardNumber(v string)`

SetPhilsysCardNumber sets PhilsysCardNumber field to given value.


### GetGivenName

`func (o *PhilippinesDigitalNidProviderOutput) GetGivenName() string`

GetGivenName returns the GivenName field if non-nil, zero value otherwise.

### GetGivenNameOk

`func (o *PhilippinesDigitalNidProviderOutput) GetGivenNameOk() (*string, bool)`

GetGivenNameOk returns a tuple with the GivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenName

`func (o *PhilippinesDigitalNidProviderOutput) SetGivenName(v string)`

SetGivenName sets GivenName field to given value.

### HasGivenName

`func (o *PhilippinesDigitalNidProviderOutput) HasGivenName() bool`

HasGivenName returns a boolean if a field has been set.

### SetGivenNameNil

`func (o *PhilippinesDigitalNidProviderOutput) SetGivenNameNil(b bool)`

 SetGivenNameNil sets the value for GivenName to be an explicit nil

### UnsetGivenName
`func (o *PhilippinesDigitalNidProviderOutput) UnsetGivenName()`

UnsetGivenName ensures that no value is present for GivenName, not even an explicit nil
### GetMiddleName

`func (o *PhilippinesDigitalNidProviderOutput) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *PhilippinesDigitalNidProviderOutput) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *PhilippinesDigitalNidProviderOutput) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *PhilippinesDigitalNidProviderOutput) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### SetMiddleNameNil

`func (o *PhilippinesDigitalNidProviderOutput) SetMiddleNameNil(b bool)`

 SetMiddleNameNil sets the value for MiddleName to be an explicit nil

### UnsetMiddleName
`func (o *PhilippinesDigitalNidProviderOutput) UnsetMiddleName()`

UnsetMiddleName ensures that no value is present for MiddleName, not even an explicit nil
### GetFamilyName

`func (o *PhilippinesDigitalNidProviderOutput) GetFamilyName() string`

GetFamilyName returns the FamilyName field if non-nil, zero value otherwise.

### GetFamilyNameOk

`func (o *PhilippinesDigitalNidProviderOutput) GetFamilyNameOk() (*string, bool)`

GetFamilyNameOk returns a tuple with the FamilyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyName

`func (o *PhilippinesDigitalNidProviderOutput) SetFamilyName(v string)`

SetFamilyName sets FamilyName field to given value.

### HasFamilyName

`func (o *PhilippinesDigitalNidProviderOutput) HasFamilyName() bool`

HasFamilyName returns a boolean if a field has been set.

### SetFamilyNameNil

`func (o *PhilippinesDigitalNidProviderOutput) SetFamilyNameNil(b bool)`

 SetFamilyNameNil sets the value for FamilyName to be an explicit nil

### UnsetFamilyName
`func (o *PhilippinesDigitalNidProviderOutput) UnsetFamilyName()`

UnsetFamilyName ensures that no value is present for FamilyName, not even an explicit nil
### GetDateOfBirth

`func (o *PhilippinesDigitalNidProviderOutput) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *PhilippinesDigitalNidProviderOutput) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *PhilippinesDigitalNidProviderOutput) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *PhilippinesDigitalNidProviderOutput) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### SetDateOfBirthNil

`func (o *PhilippinesDigitalNidProviderOutput) SetDateOfBirthNil(b bool)`

 SetDateOfBirthNil sets the value for DateOfBirth to be an explicit nil

### UnsetDateOfBirth
`func (o *PhilippinesDigitalNidProviderOutput) UnsetDateOfBirth()`

UnsetDateOfBirth ensures that no value is present for DateOfBirth, not even an explicit nil
### GetSuffix

`func (o *PhilippinesDigitalNidProviderOutput) GetSuffix() string`

GetSuffix returns the Suffix field if non-nil, zero value otherwise.

### GetSuffixOk

`func (o *PhilippinesDigitalNidProviderOutput) GetSuffixOk() (*string, bool)`

GetSuffixOk returns a tuple with the Suffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuffix

`func (o *PhilippinesDigitalNidProviderOutput) SetSuffix(v string)`

SetSuffix sets Suffix field to given value.

### HasSuffix

`func (o *PhilippinesDigitalNidProviderOutput) HasSuffix() bool`

HasSuffix returns a boolean if a field has been set.

### SetSuffixNil

`func (o *PhilippinesDigitalNidProviderOutput) SetSuffixNil(b bool)`

 SetSuffixNil sets the value for Suffix to be an explicit nil

### UnsetSuffix
`func (o *PhilippinesDigitalNidProviderOutput) UnsetSuffix()`

UnsetSuffix ensures that no value is present for Suffix, not even an explicit nil
### GetSex

`func (o *PhilippinesDigitalNidProviderOutput) GetSex() string`

GetSex returns the Sex field if non-nil, zero value otherwise.

### GetSexOk

`func (o *PhilippinesDigitalNidProviderOutput) GetSexOk() (*string, bool)`

GetSexOk returns a tuple with the Sex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSex

`func (o *PhilippinesDigitalNidProviderOutput) SetSex(v string)`

SetSex sets Sex field to given value.

### HasSex

`func (o *PhilippinesDigitalNidProviderOutput) HasSex() bool`

HasSex returns a boolean if a field has been set.

### SetSexNil

`func (o *PhilippinesDigitalNidProviderOutput) SetSexNil(b bool)`

 SetSexNil sets the value for Sex to be an explicit nil

### UnsetSex
`func (o *PhilippinesDigitalNidProviderOutput) UnsetSex()`

UnsetSex ensures that no value is present for Sex, not even an explicit nil
### GetPlaceOfBirth

`func (o *PhilippinesDigitalNidProviderOutput) GetPlaceOfBirth() string`

GetPlaceOfBirth returns the PlaceOfBirth field if non-nil, zero value otherwise.

### GetPlaceOfBirthOk

`func (o *PhilippinesDigitalNidProviderOutput) GetPlaceOfBirthOk() (*string, bool)`

GetPlaceOfBirthOk returns a tuple with the PlaceOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceOfBirth

`func (o *PhilippinesDigitalNidProviderOutput) SetPlaceOfBirth(v string)`

SetPlaceOfBirth sets PlaceOfBirth field to given value.

### HasPlaceOfBirth

`func (o *PhilippinesDigitalNidProviderOutput) HasPlaceOfBirth() bool`

HasPlaceOfBirth returns a boolean if a field has been set.

### SetPlaceOfBirthNil

`func (o *PhilippinesDigitalNidProviderOutput) SetPlaceOfBirthNil(b bool)`

 SetPlaceOfBirthNil sets the value for PlaceOfBirth to be an explicit nil

### UnsetPlaceOfBirth
`func (o *PhilippinesDigitalNidProviderOutput) UnsetPlaceOfBirth()`

UnsetPlaceOfBirth ensures that no value is present for PlaceOfBirth, not even an explicit nil
### GetDocumentIssueDate

`func (o *PhilippinesDigitalNidProviderOutput) GetDocumentIssueDate() string`

GetDocumentIssueDate returns the DocumentIssueDate field if non-nil, zero value otherwise.

### GetDocumentIssueDateOk

`func (o *PhilippinesDigitalNidProviderOutput) GetDocumentIssueDateOk() (*string, bool)`

GetDocumentIssueDateOk returns a tuple with the DocumentIssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentIssueDate

`func (o *PhilippinesDigitalNidProviderOutput) SetDocumentIssueDate(v string)`

SetDocumentIssueDate sets DocumentIssueDate field to given value.

### HasDocumentIssueDate

`func (o *PhilippinesDigitalNidProviderOutput) HasDocumentIssueDate() bool`

HasDocumentIssueDate returns a boolean if a field has been set.

### SetDocumentIssueDateNil

`func (o *PhilippinesDigitalNidProviderOutput) SetDocumentIssueDateNil(b bool)`

 SetDocumentIssueDateNil sets the value for DocumentIssueDate to be an explicit nil

### UnsetDocumentIssueDate
`func (o *PhilippinesDigitalNidProviderOutput) UnsetDocumentIssueDate()`

UnsetDocumentIssueDate ensures that no value is present for DocumentIssueDate, not even an explicit nil
### GetDocumentNumber

`func (o *PhilippinesDigitalNidProviderOutput) GetDocumentNumber() string`

GetDocumentNumber returns the DocumentNumber field if non-nil, zero value otherwise.

### GetDocumentNumberOk

`func (o *PhilippinesDigitalNidProviderOutput) GetDocumentNumberOk() (*string, bool)`

GetDocumentNumberOk returns a tuple with the DocumentNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentNumber

`func (o *PhilippinesDigitalNidProviderOutput) SetDocumentNumber(v string)`

SetDocumentNumber sets DocumentNumber field to given value.

### HasDocumentNumber

`func (o *PhilippinesDigitalNidProviderOutput) HasDocumentNumber() bool`

HasDocumentNumber returns a boolean if a field has been set.

### SetDocumentNumberNil

`func (o *PhilippinesDigitalNidProviderOutput) SetDocumentNumberNil(b bool)`

 SetDocumentNumberNil sets the value for DocumentNumber to be an explicit nil

### UnsetDocumentNumber
`func (o *PhilippinesDigitalNidProviderOutput) UnsetDocumentNumber()`

UnsetDocumentNumber ensures that no value is present for DocumentNumber, not even an explicit nil
### GetIssuingAuthority

`func (o *PhilippinesDigitalNidProviderOutput) GetIssuingAuthority() string`

GetIssuingAuthority returns the IssuingAuthority field if non-nil, zero value otherwise.

### GetIssuingAuthorityOk

`func (o *PhilippinesDigitalNidProviderOutput) GetIssuingAuthorityOk() (*string, bool)`

GetIssuingAuthorityOk returns a tuple with the IssuingAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingAuthority

`func (o *PhilippinesDigitalNidProviderOutput) SetIssuingAuthority(v string)`

SetIssuingAuthority sets IssuingAuthority field to given value.

### HasIssuingAuthority

`func (o *PhilippinesDigitalNidProviderOutput) HasIssuingAuthority() bool`

HasIssuingAuthority returns a boolean if a field has been set.

### SetIssuingAuthorityNil

`func (o *PhilippinesDigitalNidProviderOutput) SetIssuingAuthorityNil(b bool)`

 SetIssuingAuthorityNil sets the value for IssuingAuthority to be an explicit nil

### UnsetIssuingAuthority
`func (o *PhilippinesDigitalNidProviderOutput) UnsetIssuingAuthority()`

UnsetIssuingAuthority ensures that no value is present for IssuingAuthority, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


