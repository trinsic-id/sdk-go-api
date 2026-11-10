# BrazilCpfCheckInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpfNumber** | Pointer to **NullableString** | The user&#39;s 11-digit, numeric CPF Number | [optional] 
**DateOfBirth** | Pointer to **NullableString** | The user&#39;s date of birth, in &#x60;YYYY-MM-DD&#x60; format. Optional in this request; required to complete CPF lookups. | [optional] 
**SelfieImage** | Pointer to **NullableString** | The selfie image collected from the user. When using an SDK, this can be a raw byte array. In raw JSON form this is Base64. | [optional] 
**SelfieImageContentType** | Pointer to **NullableString** | The MIME Type of the file contained in &#x60;SelfieImage&#x60;.              Must be one of &#x60;image/jpeg&#x60; or &#x60;image/png&#x60;. | [optional] 

## Methods

### Named construction

Use a keyed struct literal (`BrazilCpfCheckInput{FieldName: value}`), or call
`NewBrazilCpfCheckInputWithDefaults()` and set fields by name. Positional field
factories and unkeyed struct literals are not supported.

### NewBrazilCpfCheckInputWithDefaults

`func NewBrazilCpfCheckInputWithDefaults() *BrazilCpfCheckInput`

NewBrazilCpfCheckInputWithDefaults instantiates a new BrazilCpfCheckInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpfNumber

`func (o *BrazilCpfCheckInput) GetCpfNumber() string`

GetCpfNumber returns the CpfNumber field if non-nil, zero value otherwise.

### GetCpfNumberOk

`func (o *BrazilCpfCheckInput) GetCpfNumberOk() (*string, bool)`

GetCpfNumberOk returns a tuple with the CpfNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpfNumber

`func (o *BrazilCpfCheckInput) SetCpfNumber(v string)`

SetCpfNumber sets CpfNumber field to given value.

### HasCpfNumber

`func (o *BrazilCpfCheckInput) HasCpfNumber() bool`

HasCpfNumber returns a boolean if a field has been set.

### SetCpfNumberNil

`func (o *BrazilCpfCheckInput) SetCpfNumberNil(b bool)`

 SetCpfNumberNil sets the value for CpfNumber to be an explicit nil

### UnsetCpfNumber
`func (o *BrazilCpfCheckInput) UnsetCpfNumber()`

UnsetCpfNumber ensures that no value is present for CpfNumber, not even an explicit nil
### GetDateOfBirth

`func (o *BrazilCpfCheckInput) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *BrazilCpfCheckInput) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *BrazilCpfCheckInput) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *BrazilCpfCheckInput) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### SetDateOfBirthNil

`func (o *BrazilCpfCheckInput) SetDateOfBirthNil(b bool)`

 SetDateOfBirthNil sets the value for DateOfBirth to be an explicit nil

### UnsetDateOfBirth
`func (o *BrazilCpfCheckInput) UnsetDateOfBirth()`

UnsetDateOfBirth ensures that no value is present for DateOfBirth, not even an explicit nil
### GetSelfieImage

`func (o *BrazilCpfCheckInput) GetSelfieImage() string`

GetSelfieImage returns the SelfieImage field if non-nil, zero value otherwise.

### GetSelfieImageOk

`func (o *BrazilCpfCheckInput) GetSelfieImageOk() (*string, bool)`

GetSelfieImageOk returns a tuple with the SelfieImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfieImage

`func (o *BrazilCpfCheckInput) SetSelfieImage(v string)`

SetSelfieImage sets SelfieImage field to given value.

### HasSelfieImage

`func (o *BrazilCpfCheckInput) HasSelfieImage() bool`

HasSelfieImage returns a boolean if a field has been set.

### SetSelfieImageNil

`func (o *BrazilCpfCheckInput) SetSelfieImageNil(b bool)`

 SetSelfieImageNil sets the value for SelfieImage to be an explicit nil

### UnsetSelfieImage
`func (o *BrazilCpfCheckInput) UnsetSelfieImage()`

UnsetSelfieImage ensures that no value is present for SelfieImage, not even an explicit nil
### GetSelfieImageContentType

`func (o *BrazilCpfCheckInput) GetSelfieImageContentType() string`

GetSelfieImageContentType returns the SelfieImageContentType field if non-nil, zero value otherwise.

### GetSelfieImageContentTypeOk

`func (o *BrazilCpfCheckInput) GetSelfieImageContentTypeOk() (*string, bool)`

GetSelfieImageContentTypeOk returns a tuple with the SelfieImageContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfieImageContentType

`func (o *BrazilCpfCheckInput) SetSelfieImageContentType(v string)`

SetSelfieImageContentType sets SelfieImageContentType field to given value.

### HasSelfieImageContentType

`func (o *BrazilCpfCheckInput) HasSelfieImageContentType() bool`

HasSelfieImageContentType returns a boolean if a field has been set.

### SetSelfieImageContentTypeNil

`func (o *BrazilCpfCheckInput) SetSelfieImageContentTypeNil(b bool)`

 SetSelfieImageContentTypeNil sets the value for SelfieImageContentType to be an explicit nil

### UnsetSelfieImageContentType
`func (o *BrazilCpfCheckInput) UnsetSelfieImageContentType()`

UnsetSelfieImageContentType ensures that no value is present for SelfieImageContentType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


