# IndiaPanLookupInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PermanentAccountNumber** | **string** | Permanent Account Number (PAN) to verify. PAN is a ten-character tax identifier issued by the Income Tax Department of India. | 

## Methods

### Named construction

Use a keyed struct literal (`IndiaPanLookupInput{FieldName: value}`), or call
`NewIndiaPanLookupInputWithDefaults()` and set fields by name. Positional field
factories and unkeyed struct literals are not supported.

### NewIndiaPanLookupInputWithDefaults

`func NewIndiaPanLookupInputWithDefaults() *IndiaPanLookupInput`

NewIndiaPanLookupInputWithDefaults instantiates a new IndiaPanLookupInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermanentAccountNumber

`func (o *IndiaPanLookupInput) GetPermanentAccountNumber() string`

GetPermanentAccountNumber returns the PermanentAccountNumber field if non-nil, zero value otherwise.

### GetPermanentAccountNumberOk

`func (o *IndiaPanLookupInput) GetPermanentAccountNumberOk() (*string, bool)`

GetPermanentAccountNumberOk returns a tuple with the PermanentAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermanentAccountNumber

`func (o *IndiaPanLookupInput) SetPermanentAccountNumber(v string)`

SetPermanentAccountNumber sets PermanentAccountNumber field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


