# BrazilCnhMatchFieldDateOnly

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | **string** | Decoded CNH QR Code value used for this match field. | 
**MatchValue** | Pointer to **NullableBool** | Whether the decoded CNH QR Code value matched the government database. | [optional] 

## Methods

### Named construction

Use a keyed struct literal (`BrazilCnhMatchFieldDateOnly{FieldName: value}`), or call
`NewBrazilCnhMatchFieldDateOnlyWithDefaults()` and set fields by name. Positional field
factories and unkeyed struct literals are not supported.

### NewBrazilCnhMatchFieldDateOnlyWithDefaults

`func NewBrazilCnhMatchFieldDateOnlyWithDefaults() *BrazilCnhMatchFieldDateOnly`

NewBrazilCnhMatchFieldDateOnlyWithDefaults instantiates a new BrazilCnhMatchFieldDateOnly object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *BrazilCnhMatchFieldDateOnly) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *BrazilCnhMatchFieldDateOnly) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *BrazilCnhMatchFieldDateOnly) SetContent(v string)`

SetContent sets Content field to given value.


### GetMatchValue

`func (o *BrazilCnhMatchFieldDateOnly) GetMatchValue() bool`

GetMatchValue returns the MatchValue field if non-nil, zero value otherwise.

### GetMatchValueOk

`func (o *BrazilCnhMatchFieldDateOnly) GetMatchValueOk() (*bool, bool)`

GetMatchValueOk returns a tuple with the MatchValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchValue

`func (o *BrazilCnhMatchFieldDateOnly) SetMatchValue(v bool)`

SetMatchValue sets MatchValue field to given value.

### HasMatchValue

`func (o *BrazilCnhMatchFieldDateOnly) HasMatchValue() bool`

HasMatchValue returns a boolean if a field has been set.

### SetMatchValueNil

`func (o *BrazilCnhMatchFieldDateOnly) SetMatchValueNil(b bool)`

 SetMatchValueNil sets the value for MatchValue to be an explicit nil

### UnsetMatchValue
`func (o *BrazilCnhMatchFieldDateOnly) UnsetMatchValue()`

UnsetMatchValue ensures that no value is present for MatchValue, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


