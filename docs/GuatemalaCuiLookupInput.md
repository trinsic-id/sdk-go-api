# GuatemalaCuiLookupInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocumentNumber** | Pointer to **NullableString** | The Guatemalan Código Único de Identificación (CUI) number.              Assigned and maintained by RENAP (Registro Nacional de las Personas). Official format: exactly 13 numeric digits. 8 RENAP-assigned serial digits, 1 verifier digit (dígito verificador), and 4 geographic digits for department and municipality of birth. The CUI is printed on the Documento Personal de Identificación (DPI) in three groups (4–5–4) separated by spaces.              Trinsic normalizes to digits-only before lookup, automatically removing spaces, dots, hyphens, and other non-alphanumeric characters.              No verifier algorithm appears in publicly accessible RENAP resources. Community-maintained validators often use modulus-11 (non-official). | [optional] 

## Methods

### NewGuatemalaCuiLookupInput

`func NewGuatemalaCuiLookupInput() *GuatemalaCuiLookupInput`

NewGuatemalaCuiLookupInput instantiates a new GuatemalaCuiLookupInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuatemalaCuiLookupInputWithDefaults

`func NewGuatemalaCuiLookupInputWithDefaults() *GuatemalaCuiLookupInput`

NewGuatemalaCuiLookupInputWithDefaults instantiates a new GuatemalaCuiLookupInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocumentNumber

`func (o *GuatemalaCuiLookupInput) GetDocumentNumber() string`

GetDocumentNumber returns the DocumentNumber field if non-nil, zero value otherwise.

### GetDocumentNumberOk

`func (o *GuatemalaCuiLookupInput) GetDocumentNumberOk() (*string, bool)`

GetDocumentNumberOk returns a tuple with the DocumentNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentNumber

`func (o *GuatemalaCuiLookupInput) SetDocumentNumber(v string)`

SetDocumentNumber sets DocumentNumber field to given value.

### HasDocumentNumber

`func (o *GuatemalaCuiLookupInput) HasDocumentNumber() bool`

HasDocumentNumber returns a boolean if a field has been set.

### SetDocumentNumberNil

`func (o *GuatemalaCuiLookupInput) SetDocumentNumberNil(b bool)`

 SetDocumentNumberNil sets the value for DocumentNumber to be an explicit nil

### UnsetDocumentNumber
`func (o *GuatemalaCuiLookupInput) UnsetDocumentNumber()`

UnsetDocumentNumber ensures that no value is present for DocumentNumber, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


