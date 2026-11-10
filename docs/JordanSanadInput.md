# JordanSanadInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QrCodeText** | Pointer to **NullableString** | The decoded contents of the user&#39;s SANAD QR code. | [optional] 

## Methods

### Named construction

Use a keyed struct literal (`JordanSanadInput{FieldName: value}`), or call
`NewJordanSanadInputWithDefaults()` and set fields by name. Positional field
factories and unkeyed struct literals are not supported.

### NewJordanSanadInputWithDefaults

`func NewJordanSanadInputWithDefaults() *JordanSanadInput`

NewJordanSanadInputWithDefaults instantiates a new JordanSanadInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQrCodeText

`func (o *JordanSanadInput) GetQrCodeText() string`

GetQrCodeText returns the QrCodeText field if non-nil, zero value otherwise.

### GetQrCodeTextOk

`func (o *JordanSanadInput) GetQrCodeTextOk() (*string, bool)`

GetQrCodeTextOk returns a tuple with the QrCodeText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQrCodeText

`func (o *JordanSanadInput) SetQrCodeText(v string)`

SetQrCodeText sets QrCodeText field to given value.

### HasQrCodeText

`func (o *JordanSanadInput) HasQrCodeText() bool`

HasQrCodeText returns a boolean if a field has been set.

### SetQrCodeTextNil

`func (o *JordanSanadInput) SetQrCodeTextNil(b bool)`

 SetQrCodeTextNil sets the value for QrCodeText to be an explicit nil

### UnsetQrCodeText
`func (o *JordanSanadInput) UnsetQrCodeText()`

UnsetQrCodeText ensures that no value is present for QrCodeText, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


