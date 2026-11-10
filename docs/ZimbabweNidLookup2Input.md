# ZimbabweNidLookup2Input

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdNumber** | Pointer to **NullableString** | Zimbabwe National ID number (NID) issued by the Zimbabwean government.              Format: - 8-9 digits followed by 1 letter and 2 digits - Regex: /^[0-9]{8,9}[A-Za-z]\\d{2}$/ - There is no publicly documented encoding scheme for encoding personal information in the NID - No check digit or algorithm has been publicly documented by the Zimbabwean government | [optional] 

## Methods

### Named construction

Use a keyed struct literal (`ZimbabweNidLookup2Input{FieldName: value}`), or call
`NewZimbabweNidLookup2InputWithDefaults()` and set fields by name. Positional field
factories and unkeyed struct literals are not supported.

### NewZimbabweNidLookup2InputWithDefaults

`func NewZimbabweNidLookup2InputWithDefaults() *ZimbabweNidLookup2Input`

NewZimbabweNidLookup2InputWithDefaults instantiates a new ZimbabweNidLookup2Input object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdNumber

`func (o *ZimbabweNidLookup2Input) GetIdNumber() string`

GetIdNumber returns the IdNumber field if non-nil, zero value otherwise.

### GetIdNumberOk

`func (o *ZimbabweNidLookup2Input) GetIdNumberOk() (*string, bool)`

GetIdNumberOk returns a tuple with the IdNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdNumber

`func (o *ZimbabweNidLookup2Input) SetIdNumber(v string)`

SetIdNumber sets IdNumber field to given value.

### HasIdNumber

`func (o *ZimbabweNidLookup2Input) HasIdNumber() bool`

HasIdNumber returns a boolean if a field has been set.

### SetIdNumberNil

`func (o *ZimbabweNidLookup2Input) SetIdNumberNil(b bool)`

 SetIdNumberNil sets the value for IdNumber to be an explicit nil

### UnsetIdNumber
`func (o *ZimbabweNidLookup2Input) UnsetIdNumber()`

UnsetIdNumber ensures that no value is present for IdNumber, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


