# PhilippinesPhysicalNidProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PhilsysCardNumber** | Pointer to **NullableString** | The PhilSys Card Number (PCN). Every citizen or resident alien registered in PhilSys has a PhilSys Number (PSN). This number is tokenized into a card number to protect the PSN. The PhilSys Card Number is 12 characters long, and often is written in octets with dashes in between. | [optional] 
**GivenName** | Pointer to **NullableString** | The given (first) name of the individual. | [optional] 
**MiddleName** | Pointer to **NullableString** | The middle name of the individual. | [optional] 
**FamilyName** | Pointer to **NullableString** | The family (last) name of the individual. | [optional] 
**DateOfBirth** | Pointer to **NullableString** | The date of birth of the individual. | [optional] 
**Suffix** | Pointer to **NullableString** | The name suffix of the individual (e.g. Jr., III). | [optional] 
**Sex** | Pointer to **NullableString** | The sex of the individual.              Possible values: - Male - Female | [optional] 
**PlaceOfBirth** | Pointer to **NullableString** | The place of birth of the individual as recorded on the PhilSys credential. | [optional] 
**DocumentIssueDate** | Pointer to **NullableString** | The date the document was issued. | [optional] 

## Methods

### NewPhilippinesPhysicalNidProviderOutput

`func NewPhilippinesPhysicalNidProviderOutput() *PhilippinesPhysicalNidProviderOutput`

NewPhilippinesPhysicalNidProviderOutput instantiates a new PhilippinesPhysicalNidProviderOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhilippinesPhysicalNidProviderOutputWithDefaults

`func NewPhilippinesPhysicalNidProviderOutputWithDefaults() *PhilippinesPhysicalNidProviderOutput`

NewPhilippinesPhysicalNidProviderOutputWithDefaults instantiates a new PhilippinesPhysicalNidProviderOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhilsysCardNumber

`func (o *PhilippinesPhysicalNidProviderOutput) GetPhilsysCardNumber() string`

GetPhilsysCardNumber returns the PhilsysCardNumber field if non-nil, zero value otherwise.

### GetPhilsysCardNumberOk

`func (o *PhilippinesPhysicalNidProviderOutput) GetPhilsysCardNumberOk() (*string, bool)`

GetPhilsysCardNumberOk returns a tuple with the PhilsysCardNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhilsysCardNumber

`func (o *PhilippinesPhysicalNidProviderOutput) SetPhilsysCardNumber(v string)`

SetPhilsysCardNumber sets PhilsysCardNumber field to given value.

### HasPhilsysCardNumber

`func (o *PhilippinesPhysicalNidProviderOutput) HasPhilsysCardNumber() bool`

HasPhilsysCardNumber returns a boolean if a field has been set.

### SetPhilsysCardNumberNil

`func (o *PhilippinesPhysicalNidProviderOutput) SetPhilsysCardNumberNil(b bool)`

 SetPhilsysCardNumberNil sets the value for PhilsysCardNumber to be an explicit nil

### UnsetPhilsysCardNumber
`func (o *PhilippinesPhysicalNidProviderOutput) UnsetPhilsysCardNumber()`

UnsetPhilsysCardNumber ensures that no value is present for PhilsysCardNumber, not even an explicit nil
### GetGivenName

`func (o *PhilippinesPhysicalNidProviderOutput) GetGivenName() string`

GetGivenName returns the GivenName field if non-nil, zero value otherwise.

### GetGivenNameOk

`func (o *PhilippinesPhysicalNidProviderOutput) GetGivenNameOk() (*string, bool)`

GetGivenNameOk returns a tuple with the GivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenName

`func (o *PhilippinesPhysicalNidProviderOutput) SetGivenName(v string)`

SetGivenName sets GivenName field to given value.

### HasGivenName

`func (o *PhilippinesPhysicalNidProviderOutput) HasGivenName() bool`

HasGivenName returns a boolean if a field has been set.

### SetGivenNameNil

`func (o *PhilippinesPhysicalNidProviderOutput) SetGivenNameNil(b bool)`

 SetGivenNameNil sets the value for GivenName to be an explicit nil

### UnsetGivenName
`func (o *PhilippinesPhysicalNidProviderOutput) UnsetGivenName()`

UnsetGivenName ensures that no value is present for GivenName, not even an explicit nil
### GetMiddleName

`func (o *PhilippinesPhysicalNidProviderOutput) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *PhilippinesPhysicalNidProviderOutput) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *PhilippinesPhysicalNidProviderOutput) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *PhilippinesPhysicalNidProviderOutput) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### SetMiddleNameNil

`func (o *PhilippinesPhysicalNidProviderOutput) SetMiddleNameNil(b bool)`

 SetMiddleNameNil sets the value for MiddleName to be an explicit nil

### UnsetMiddleName
`func (o *PhilippinesPhysicalNidProviderOutput) UnsetMiddleName()`

UnsetMiddleName ensures that no value is present for MiddleName, not even an explicit nil
### GetFamilyName

`func (o *PhilippinesPhysicalNidProviderOutput) GetFamilyName() string`

GetFamilyName returns the FamilyName field if non-nil, zero value otherwise.

### GetFamilyNameOk

`func (o *PhilippinesPhysicalNidProviderOutput) GetFamilyNameOk() (*string, bool)`

GetFamilyNameOk returns a tuple with the FamilyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyName

`func (o *PhilippinesPhysicalNidProviderOutput) SetFamilyName(v string)`

SetFamilyName sets FamilyName field to given value.

### HasFamilyName

`func (o *PhilippinesPhysicalNidProviderOutput) HasFamilyName() bool`

HasFamilyName returns a boolean if a field has been set.

### SetFamilyNameNil

`func (o *PhilippinesPhysicalNidProviderOutput) SetFamilyNameNil(b bool)`

 SetFamilyNameNil sets the value for FamilyName to be an explicit nil

### UnsetFamilyName
`func (o *PhilippinesPhysicalNidProviderOutput) UnsetFamilyName()`

UnsetFamilyName ensures that no value is present for FamilyName, not even an explicit nil
### GetDateOfBirth

`func (o *PhilippinesPhysicalNidProviderOutput) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *PhilippinesPhysicalNidProviderOutput) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *PhilippinesPhysicalNidProviderOutput) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *PhilippinesPhysicalNidProviderOutput) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### SetDateOfBirthNil

`func (o *PhilippinesPhysicalNidProviderOutput) SetDateOfBirthNil(b bool)`

 SetDateOfBirthNil sets the value for DateOfBirth to be an explicit nil

### UnsetDateOfBirth
`func (o *PhilippinesPhysicalNidProviderOutput) UnsetDateOfBirth()`

UnsetDateOfBirth ensures that no value is present for DateOfBirth, not even an explicit nil
### GetSuffix

`func (o *PhilippinesPhysicalNidProviderOutput) GetSuffix() string`

GetSuffix returns the Suffix field if non-nil, zero value otherwise.

### GetSuffixOk

`func (o *PhilippinesPhysicalNidProviderOutput) GetSuffixOk() (*string, bool)`

GetSuffixOk returns a tuple with the Suffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuffix

`func (o *PhilippinesPhysicalNidProviderOutput) SetSuffix(v string)`

SetSuffix sets Suffix field to given value.

### HasSuffix

`func (o *PhilippinesPhysicalNidProviderOutput) HasSuffix() bool`

HasSuffix returns a boolean if a field has been set.

### SetSuffixNil

`func (o *PhilippinesPhysicalNidProviderOutput) SetSuffixNil(b bool)`

 SetSuffixNil sets the value for Suffix to be an explicit nil

### UnsetSuffix
`func (o *PhilippinesPhysicalNidProviderOutput) UnsetSuffix()`

UnsetSuffix ensures that no value is present for Suffix, not even an explicit nil
### GetSex

`func (o *PhilippinesPhysicalNidProviderOutput) GetSex() string`

GetSex returns the Sex field if non-nil, zero value otherwise.

### GetSexOk

`func (o *PhilippinesPhysicalNidProviderOutput) GetSexOk() (*string, bool)`

GetSexOk returns a tuple with the Sex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSex

`func (o *PhilippinesPhysicalNidProviderOutput) SetSex(v string)`

SetSex sets Sex field to given value.

### HasSex

`func (o *PhilippinesPhysicalNidProviderOutput) HasSex() bool`

HasSex returns a boolean if a field has been set.

### SetSexNil

`func (o *PhilippinesPhysicalNidProviderOutput) SetSexNil(b bool)`

 SetSexNil sets the value for Sex to be an explicit nil

### UnsetSex
`func (o *PhilippinesPhysicalNidProviderOutput) UnsetSex()`

UnsetSex ensures that no value is present for Sex, not even an explicit nil
### GetPlaceOfBirth

`func (o *PhilippinesPhysicalNidProviderOutput) GetPlaceOfBirth() string`

GetPlaceOfBirth returns the PlaceOfBirth field if non-nil, zero value otherwise.

### GetPlaceOfBirthOk

`func (o *PhilippinesPhysicalNidProviderOutput) GetPlaceOfBirthOk() (*string, bool)`

GetPlaceOfBirthOk returns a tuple with the PlaceOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceOfBirth

`func (o *PhilippinesPhysicalNidProviderOutput) SetPlaceOfBirth(v string)`

SetPlaceOfBirth sets PlaceOfBirth field to given value.

### HasPlaceOfBirth

`func (o *PhilippinesPhysicalNidProviderOutput) HasPlaceOfBirth() bool`

HasPlaceOfBirth returns a boolean if a field has been set.

### SetPlaceOfBirthNil

`func (o *PhilippinesPhysicalNidProviderOutput) SetPlaceOfBirthNil(b bool)`

 SetPlaceOfBirthNil sets the value for PlaceOfBirth to be an explicit nil

### UnsetPlaceOfBirth
`func (o *PhilippinesPhysicalNidProviderOutput) UnsetPlaceOfBirth()`

UnsetPlaceOfBirth ensures that no value is present for PlaceOfBirth, not even an explicit nil
### GetDocumentIssueDate

`func (o *PhilippinesPhysicalNidProviderOutput) GetDocumentIssueDate() string`

GetDocumentIssueDate returns the DocumentIssueDate field if non-nil, zero value otherwise.

### GetDocumentIssueDateOk

`func (o *PhilippinesPhysicalNidProviderOutput) GetDocumentIssueDateOk() (*string, bool)`

GetDocumentIssueDateOk returns a tuple with the DocumentIssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentIssueDate

`func (o *PhilippinesPhysicalNidProviderOutput) SetDocumentIssueDate(v string)`

SetDocumentIssueDate sets DocumentIssueDate field to given value.

### HasDocumentIssueDate

`func (o *PhilippinesPhysicalNidProviderOutput) HasDocumentIssueDate() bool`

HasDocumentIssueDate returns a boolean if a field has been set.

### SetDocumentIssueDateNil

`func (o *PhilippinesPhysicalNidProviderOutput) SetDocumentIssueDateNil(b bool)`

 SetDocumentIssueDateNil sets the value for DocumentIssueDate to be an explicit nil

### UnsetDocumentIssueDate
`func (o *PhilippinesPhysicalNidProviderOutput) UnsetDocumentIssueDate()`

UnsetDocumentIssueDate ensures that no value is present for DocumentIssueDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


