# CoteDIvoireNidLookup2Input

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdNumber** | Pointer to **NullableString** | The user&#39;s National ID number (NNI) or document number | [optional] 

## Methods

### Named construction

Use a keyed struct literal (`CoteDIvoireNidLookup2Input{FieldName: value}`), or call
`NewCoteDIvoireNidLookup2InputWithDefaults()` and set fields by name. Positional field
factories and unkeyed struct literals are not supported.

### NewCoteDIvoireNidLookup2InputWithDefaults

`func NewCoteDIvoireNidLookup2InputWithDefaults() *CoteDIvoireNidLookup2Input`

NewCoteDIvoireNidLookup2InputWithDefaults instantiates a new CoteDIvoireNidLookup2Input object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdNumber

`func (o *CoteDIvoireNidLookup2Input) GetIdNumber() string`

GetIdNumber returns the IdNumber field if non-nil, zero value otherwise.

### GetIdNumberOk

`func (o *CoteDIvoireNidLookup2Input) GetIdNumberOk() (*string, bool)`

GetIdNumberOk returns a tuple with the IdNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdNumber

`func (o *CoteDIvoireNidLookup2Input) SetIdNumber(v string)`

SetIdNumber sets IdNumber field to given value.

### HasIdNumber

`func (o *CoteDIvoireNidLookup2Input) HasIdNumber() bool`

HasIdNumber returns a boolean if a field has been set.

### SetIdNumberNil

`func (o *CoteDIvoireNidLookup2Input) SetIdNumberNil(b bool)`

 SetIdNumberNil sets the value for IdNumber to be an explicit nil

### UnsetIdNumber
`func (o *CoteDIvoireNidLookup2Input) UnsetIdNumber()`

UnsetIdNumber ensures that no value is present for IdNumber, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


