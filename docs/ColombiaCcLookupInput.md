# ColombiaCcLookupInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocumentNumber** | Pointer to **NullableString** | The CC (Cédula de Ciudadanía, Citizenship Document) document number.              Format: - Cédulas after 2004 use the NUIP (Número Único de Identificación Personal), which is 10 digits.   Older documents may have fewer than 10 digits and are still valid. - In Colombia the number is often written with dots as thousands separators (e.g. 1.234.567.890). If dots,   hyphens, spaces, or other non-alphanumeric characters are included, they will be sanitized before lookup. | [optional] 
**IssueDate** | Pointer to **NullableString** | Document issue date, in &#x60;YYYY-MM-DD&#x60; format | [optional] 

## Methods

### NewColombiaCcLookupInput

`func NewColombiaCcLookupInput() *ColombiaCcLookupInput`

NewColombiaCcLookupInput instantiates a new ColombiaCcLookupInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewColombiaCcLookupInputWithDefaults

`func NewColombiaCcLookupInputWithDefaults() *ColombiaCcLookupInput`

NewColombiaCcLookupInputWithDefaults instantiates a new ColombiaCcLookupInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocumentNumber

`func (o *ColombiaCcLookupInput) GetDocumentNumber() string`

GetDocumentNumber returns the DocumentNumber field if non-nil, zero value otherwise.

### GetDocumentNumberOk

`func (o *ColombiaCcLookupInput) GetDocumentNumberOk() (*string, bool)`

GetDocumentNumberOk returns a tuple with the DocumentNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentNumber

`func (o *ColombiaCcLookupInput) SetDocumentNumber(v string)`

SetDocumentNumber sets DocumentNumber field to given value.

### HasDocumentNumber

`func (o *ColombiaCcLookupInput) HasDocumentNumber() bool`

HasDocumentNumber returns a boolean if a field has been set.

### SetDocumentNumberNil

`func (o *ColombiaCcLookupInput) SetDocumentNumberNil(b bool)`

 SetDocumentNumberNil sets the value for DocumentNumber to be an explicit nil

### UnsetDocumentNumber
`func (o *ColombiaCcLookupInput) UnsetDocumentNumber()`

UnsetDocumentNumber ensures that no value is present for DocumentNumber, not even an explicit nil
### GetIssueDate

`func (o *ColombiaCcLookupInput) GetIssueDate() string`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *ColombiaCcLookupInput) GetIssueDateOk() (*string, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *ColombiaCcLookupInput) SetIssueDate(v string)`

SetIssueDate sets IssueDate field to given value.

### HasIssueDate

`func (o *ColombiaCcLookupInput) HasIssueDate() bool`

HasIssueDate returns a boolean if a field has been set.

### SetIssueDateNil

`func (o *ColombiaCcLookupInput) SetIssueDateNil(b bool)`

 SetIssueDateNil sets the value for IssueDate to be an explicit nil

### UnsetIssueDate
`func (o *ColombiaCcLookupInput) UnsetIssueDate()`

UnsetIssueDate ensures that no value is present for IssueDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


