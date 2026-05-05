# PanamaCedulaLookupInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocumentNumber** | Pointer to **NullableString** | The number from the holder&#39;s cédula (Cédula de Identidad Personal).              Send the full number as {firstSegment}-{libro}-{tomo} (libro 1–4 digits, tomo 1–6), using only ASCII letters, digits, and hyphens. Hyphens are not inserted for you between segments. Trinsic uppercases letters, collapses repeated hyphens, trims leading and trailing hyphens, and merges a redundant hyphen between province and AV or PI (for example 10-AV-1234-12345 becomes 10AV-1234-12345).              Citizen category and format: - Born in Panama format: {province}-{libro}-{tomo} ({province} is official code 1 through 13). Examples   8-1234-12345, 4-56-789, 12-12-12345. - Panamanian born abroad format: PE-{libro}-{tomo}. Example PE-1234-12345. - Foreign national with cédula format: E-{libro}-{tomo}. Examples E-1234-12345, E-8-102017. - Naturalized citizen format: N-{libro}-{tomo}. Example N-1234-12345. - Pre-2006 civil registry (AV) format: {province}AV-{libro}-{tomo}. Example 10AV-1234-12345. - Indigenous (PI) format: {province}PI-{libro}-{tomo}. Example 1PI-1234-12345. | [optional] 
**DateOfBirth** | Pointer to **NullableString** | The user&#39;s date of birth, in &#x60;YYYY-MM-DD&#x60; format. | [optional] 

## Methods

### NewPanamaCedulaLookupInput

`func NewPanamaCedulaLookupInput() *PanamaCedulaLookupInput`

NewPanamaCedulaLookupInput instantiates a new PanamaCedulaLookupInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPanamaCedulaLookupInputWithDefaults

`func NewPanamaCedulaLookupInputWithDefaults() *PanamaCedulaLookupInput`

NewPanamaCedulaLookupInputWithDefaults instantiates a new PanamaCedulaLookupInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocumentNumber

`func (o *PanamaCedulaLookupInput) GetDocumentNumber() string`

GetDocumentNumber returns the DocumentNumber field if non-nil, zero value otherwise.

### GetDocumentNumberOk

`func (o *PanamaCedulaLookupInput) GetDocumentNumberOk() (*string, bool)`

GetDocumentNumberOk returns a tuple with the DocumentNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentNumber

`func (o *PanamaCedulaLookupInput) SetDocumentNumber(v string)`

SetDocumentNumber sets DocumentNumber field to given value.

### HasDocumentNumber

`func (o *PanamaCedulaLookupInput) HasDocumentNumber() bool`

HasDocumentNumber returns a boolean if a field has been set.

### SetDocumentNumberNil

`func (o *PanamaCedulaLookupInput) SetDocumentNumberNil(b bool)`

 SetDocumentNumberNil sets the value for DocumentNumber to be an explicit nil

### UnsetDocumentNumber
`func (o *PanamaCedulaLookupInput) UnsetDocumentNumber()`

UnsetDocumentNumber ensures that no value is present for DocumentNumber, not even an explicit nil
### GetDateOfBirth

`func (o *PanamaCedulaLookupInput) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *PanamaCedulaLookupInput) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *PanamaCedulaLookupInput) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *PanamaCedulaLookupInput) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### SetDateOfBirthNil

`func (o *PanamaCedulaLookupInput) SetDateOfBirthNil(b bool)`

 SetDateOfBirthNil sets the value for DateOfBirth to be an explicit nil

### UnsetDateOfBirth
`func (o *PanamaCedulaLookupInput) UnsetDateOfBirth()`

UnsetDateOfBirth ensures that no value is present for DateOfBirth, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


