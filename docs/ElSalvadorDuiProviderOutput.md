# ElSalvadorDuiProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FullName** | **string** | Full name as it appears on the DUI (Documento Único de Identidad), as returned by Verifik from official records administered by the Registro Nacional de las Personas Naturales (RNPN). | 
**DocumentNumber** | **string** | The DUI (Documento Único de Identidad) document number as returned by Verifik for the matched record.              Nine numeric digits after sanitization. Commonly printed as ########-# (hyphen before the final digit). The output will be stripped of the hyphen and will preserve leading zeros.              The ninth digit is a check digit. This is not publicly documented by the Salvadoran government, but the algorithm is available in the public domain for those who seek it. | 
**DateOfBirth** | **string** | The date of birth that was supplied for the lookup and confirmed as matching the DUI record. | 

## Methods

### NewElSalvadorDuiProviderOutput

`func NewElSalvadorDuiProviderOutput(fullName string, documentNumber string, dateOfBirth string, ) *ElSalvadorDuiProviderOutput`

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


