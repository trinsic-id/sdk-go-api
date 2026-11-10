# UkEvisaLookupInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShareCode** | **string** | The 9-character alphanumeric UK Visas and Immigration (UKVI) share code. Must start with W for right to work, R for right to rent, or S for immigration status, matching the configured lookup type. Whitespace is removed and lowercase ASCII letters are converted to uppercase. | 
**DateOfBirth** | **string** | The user&#39;s date of birth, in &#x60;YYYY-MM-DD&#x60; format. | 

## Methods

### Named construction

Use a keyed struct literal (`UkEvisaLookupInput{FieldName: value}`), or call
`NewUkEvisaLookupInputWithDefaults()` and set fields by name. Positional field
factories and unkeyed struct literals are not supported.

### NewUkEvisaLookupInputWithDefaults

`func NewUkEvisaLookupInputWithDefaults() *UkEvisaLookupInput`

NewUkEvisaLookupInputWithDefaults instantiates a new UkEvisaLookupInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShareCode

`func (o *UkEvisaLookupInput) GetShareCode() string`

GetShareCode returns the ShareCode field if non-nil, zero value otherwise.

### GetShareCodeOk

`func (o *UkEvisaLookupInput) GetShareCodeOk() (*string, bool)`

GetShareCodeOk returns a tuple with the ShareCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareCode

`func (o *UkEvisaLookupInput) SetShareCode(v string)`

SetShareCode sets ShareCode field to given value.


### GetDateOfBirth

`func (o *UkEvisaLookupInput) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *UkEvisaLookupInput) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *UkEvisaLookupInput) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


