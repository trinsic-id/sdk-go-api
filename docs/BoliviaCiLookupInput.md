# BoliviaCiLookupInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocumentNumber** | Pointer to **NullableString** | The holder&#39;s CI (\&quot;Cédula de Identidad\&quot;) number from the Bolivian identity card (\&quot;carnet de identidad\&quot;). This is the identifier assigned by the Servicio General de Identificación Personal (SEGIP) in the Registro Único de Identificación (RUI). Is entirely numeric values. Any non-alphanumeric characters (dots, hyphens, spaces, etc.) will be stripped before lookup.              Published regulations do not define a fixed length; digit count may vary. | [optional] 
**DateOfBirth** | Pointer to **NullableString** | The holder&#39;s date of birth. Must match the CI record. | [optional] 

## Methods

### NewBoliviaCiLookupInput

`func NewBoliviaCiLookupInput() *BoliviaCiLookupInput`

NewBoliviaCiLookupInput instantiates a new BoliviaCiLookupInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBoliviaCiLookupInputWithDefaults

`func NewBoliviaCiLookupInputWithDefaults() *BoliviaCiLookupInput`

NewBoliviaCiLookupInputWithDefaults instantiates a new BoliviaCiLookupInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocumentNumber

`func (o *BoliviaCiLookupInput) GetDocumentNumber() string`

GetDocumentNumber returns the DocumentNumber field if non-nil, zero value otherwise.

### GetDocumentNumberOk

`func (o *BoliviaCiLookupInput) GetDocumentNumberOk() (*string, bool)`

GetDocumentNumberOk returns a tuple with the DocumentNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentNumber

`func (o *BoliviaCiLookupInput) SetDocumentNumber(v string)`

SetDocumentNumber sets DocumentNumber field to given value.

### HasDocumentNumber

`func (o *BoliviaCiLookupInput) HasDocumentNumber() bool`

HasDocumentNumber returns a boolean if a field has been set.

### SetDocumentNumberNil

`func (o *BoliviaCiLookupInput) SetDocumentNumberNil(b bool)`

 SetDocumentNumberNil sets the value for DocumentNumber to be an explicit nil

### UnsetDocumentNumber
`func (o *BoliviaCiLookupInput) UnsetDocumentNumber()`

UnsetDocumentNumber ensures that no value is present for DocumentNumber, not even an explicit nil
### GetDateOfBirth

`func (o *BoliviaCiLookupInput) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *BoliviaCiLookupInput) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *BoliviaCiLookupInput) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *BoliviaCiLookupInput) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### SetDateOfBirthNil

`func (o *BoliviaCiLookupInput) SetDateOfBirthNil(b bool)`

 SetDateOfBirthNil sets the value for DateOfBirth to be an explicit nil

### UnsetDateOfBirth
`func (o *BoliviaCiLookupInput) UnsetDateOfBirth()`

UnsetDateOfBirth ensures that no value is present for DateOfBirth, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


