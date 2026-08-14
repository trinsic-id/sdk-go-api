# ElSalvadorDuiProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FullName** | Pointer to **NullableString** | Full name as it appears on the DUI (Documento Único de Identidad), as returned by Verifik from official records administered by the Registro Nacional de las Personas Naturales (RNPN). | [optional] 
**DocumentNumber** | Pointer to **NullableString** | The DUI (Documento Único de Identidad) document number as returned by Verifik for the matched record.              Nine numeric digits after sanitization. Commonly printed as ########-# (hyphen before the final digit). The output will be stripped of the hyphen and will preserve leading zeros.              The ninth digit is a check digit. This is not publicly documented by the Salvadoran government, but the algorithm is available in the public domain for those who seek it. | [optional] 
**DateOfBirth** | Pointer to **NullableString** | The date of birth that was supplied for the lookup and confirmed as matching the DUI record. | [optional] 

## Methods

### NewElSalvadorDuiProviderOutput

`func NewElSalvadorDuiProviderOutput() *ElSalvadorDuiProviderOutput`

NewElSalvadorDuiProviderOutput instantiates a new ElSalvadorDuiProviderOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewElSalvadorDuiProviderOutputWithDefaults

`func NewElSalvadorDuiProviderOutputWithDefaults() *ElSalvadorDuiProviderOutput`

NewElSalvadorDuiProviderOutputWithDefaults instantiates a new ElSalvadorDuiProviderOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFullName

`func (o *ElSalvadorDuiProviderOutput) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *ElSalvadorDuiProviderOutput) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *ElSalvadorDuiProviderOutput) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *ElSalvadorDuiProviderOutput) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### SetFullNameNil

`func (o *ElSalvadorDuiProviderOutput) SetFullNameNil(b bool)`

 SetFullNameNil sets the value for FullName to be an explicit nil

### UnsetFullName
`func (o *ElSalvadorDuiProviderOutput) UnsetFullName()`

UnsetFullName ensures that no value is present for FullName, not even an explicit nil
### GetDocumentNumber

`func (o *ElSalvadorDuiProviderOutput) GetDocumentNumber() string`

GetDocumentNumber returns the DocumentNumber field if non-nil, zero value otherwise.

### GetDocumentNumberOk

`func (o *ElSalvadorDuiProviderOutput) GetDocumentNumberOk() (*string, bool)`

GetDocumentNumberOk returns a tuple with the DocumentNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentNumber

`func (o *ElSalvadorDuiProviderOutput) SetDocumentNumber(v string)`

SetDocumentNumber sets DocumentNumber field to given value.

### HasDocumentNumber

`func (o *ElSalvadorDuiProviderOutput) HasDocumentNumber() bool`

HasDocumentNumber returns a boolean if a field has been set.

### SetDocumentNumberNil

`func (o *ElSalvadorDuiProviderOutput) SetDocumentNumberNil(b bool)`

 SetDocumentNumberNil sets the value for DocumentNumber to be an explicit nil

### UnsetDocumentNumber
`func (o *ElSalvadorDuiProviderOutput) UnsetDocumentNumber()`

UnsetDocumentNumber ensures that no value is present for DocumentNumber, not even an explicit nil
### GetDateOfBirth

`func (o *ElSalvadorDuiProviderOutput) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *ElSalvadorDuiProviderOutput) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *ElSalvadorDuiProviderOutput) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *ElSalvadorDuiProviderOutput) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### SetDateOfBirthNil

`func (o *ElSalvadorDuiProviderOutput) SetDateOfBirthNil(b bool)`

 SetDateOfBirthNil sets the value for DateOfBirth to be an explicit nil

### UnsetDateOfBirth
`func (o *ElSalvadorDuiProviderOutput) UnsetDateOfBirth()`

UnsetDateOfBirth ensures that no value is present for DateOfBirth, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


