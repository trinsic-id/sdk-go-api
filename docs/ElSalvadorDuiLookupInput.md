# ElSalvadorDuiLookupInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocumentNumber** | Pointer to **NullableString** | The DUI (Documento Único de Identidad) number for the holder.              Nine numeric digits after sanitization. Commonly printed as ########-# (hyphen before the final digit). The input will automatically be sanitized of dots, hyphens, spaces, or other non-alphanumeric characters before lookup.              The ninth digit is a check digit. This is not publicly documented by the Salvadoran government, but the algorithm is available in the public domain for those who seek it. | [optional] 
**DateOfBirth** | Pointer to **NullableString** | The DUI holder&#39;s date of birth. Required to match the correct person in official records. | [optional] 

## Methods

### NewElSalvadorDuiLookupInput

`func NewElSalvadorDuiLookupInput() *ElSalvadorDuiLookupInput`

NewElSalvadorDuiLookupInput instantiates a new ElSalvadorDuiLookupInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewElSalvadorDuiLookupInputWithDefaults

`func NewElSalvadorDuiLookupInputWithDefaults() *ElSalvadorDuiLookupInput`

NewElSalvadorDuiLookupInputWithDefaults instantiates a new ElSalvadorDuiLookupInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocumentNumber

`func (o *ElSalvadorDuiLookupInput) GetDocumentNumber() string`

GetDocumentNumber returns the DocumentNumber field if non-nil, zero value otherwise.

### GetDocumentNumberOk

`func (o *ElSalvadorDuiLookupInput) GetDocumentNumberOk() (*string, bool)`

GetDocumentNumberOk returns a tuple with the DocumentNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentNumber

`func (o *ElSalvadorDuiLookupInput) SetDocumentNumber(v string)`

SetDocumentNumber sets DocumentNumber field to given value.

### HasDocumentNumber

`func (o *ElSalvadorDuiLookupInput) HasDocumentNumber() bool`

HasDocumentNumber returns a boolean if a field has been set.

### SetDocumentNumberNil

`func (o *ElSalvadorDuiLookupInput) SetDocumentNumberNil(b bool)`

 SetDocumentNumberNil sets the value for DocumentNumber to be an explicit nil

### UnsetDocumentNumber
`func (o *ElSalvadorDuiLookupInput) UnsetDocumentNumber()`

UnsetDocumentNumber ensures that no value is present for DocumentNumber, not even an explicit nil
### GetDateOfBirth

`func (o *ElSalvadorDuiLookupInput) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *ElSalvadorDuiLookupInput) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *ElSalvadorDuiLookupInput) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *ElSalvadorDuiLookupInput) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### SetDateOfBirthNil

`func (o *ElSalvadorDuiLookupInput) SetDateOfBirthNil(b bool)`

 SetDateOfBirthNil sets the value for DateOfBirth to be an explicit nil

### UnsetDateOfBirth
`func (o *ElSalvadorDuiLookupInput) UnsetDateOfBirth()`

UnsetDateOfBirth ensures that no value is present for DateOfBirth, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


