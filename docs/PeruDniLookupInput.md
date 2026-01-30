# PeruDniLookupInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocumentNumber** | Pointer to **NullableString** | The user&#39;s DNI number (8 digits).              Format: - Must not include verification digit. On the DNI card, a ninth digit appears after the first eight, with value   0-9 or A-K. This is not included in the DNI number when verifying against Peru&#39;s database. - Peru DNI is sometimes represented with dots. Though uncommon, if dots are included, they will be sanitized. | [optional] 

## Methods

### NewPeruDniLookupInput

`func NewPeruDniLookupInput() *PeruDniLookupInput`

NewPeruDniLookupInput instantiates a new PeruDniLookupInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPeruDniLookupInputWithDefaults

`func NewPeruDniLookupInputWithDefaults() *PeruDniLookupInput`

NewPeruDniLookupInputWithDefaults instantiates a new PeruDniLookupInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocumentNumber

`func (o *PeruDniLookupInput) GetDocumentNumber() string`

GetDocumentNumber returns the DocumentNumber field if non-nil, zero value otherwise.

### GetDocumentNumberOk

`func (o *PeruDniLookupInput) GetDocumentNumberOk() (*string, bool)`

GetDocumentNumberOk returns a tuple with the DocumentNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentNumber

`func (o *PeruDniLookupInput) SetDocumentNumber(v string)`

SetDocumentNumber sets DocumentNumber field to given value.

### HasDocumentNumber

`func (o *PeruDniLookupInput) HasDocumentNumber() bool`

HasDocumentNumber returns a boolean if a field has been set.

### SetDocumentNumberNil

`func (o *PeruDniLookupInput) SetDocumentNumberNil(b bool)`

 SetDocumentNumberNil sets the value for DocumentNumber to be an explicit nil

### UnsetDocumentNumber
`func (o *PeruDniLookupInput) UnsetDocumentNumber()`

UnsetDocumentNumber ensures that no value is present for DocumentNumber, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


