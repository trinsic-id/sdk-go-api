# PhilippineQRInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QrCodeText** | Pointer to **NullableString** | The raw text of the user&#39;s QR code after decoding it. | [optional] 
**QrCodeImage** | Pointer to **NullableString** | The raw bytes of the image containing the user&#39;s QR code. | [optional] 

## Methods

### Named construction

Use a keyed struct literal (`PhilippineQRInput{FieldName: value}`), or call
`NewPhilippineQRInputWithDefaults()` and set fields by name. Positional field
factories and unkeyed struct literals are not supported.

### NewPhilippineQRInputWithDefaults

`func NewPhilippineQRInputWithDefaults() *PhilippineQRInput`

NewPhilippineQRInputWithDefaults instantiates a new PhilippineQRInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQrCodeText

`func (o *PhilippineQRInput) GetQrCodeText() string`

GetQrCodeText returns the QrCodeText field if non-nil, zero value otherwise.

### GetQrCodeTextOk

`func (o *PhilippineQRInput) GetQrCodeTextOk() (*string, bool)`

GetQrCodeTextOk returns a tuple with the QrCodeText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQrCodeText

`func (o *PhilippineQRInput) SetQrCodeText(v string)`

SetQrCodeText sets QrCodeText field to given value.

### HasQrCodeText

`func (o *PhilippineQRInput) HasQrCodeText() bool`

HasQrCodeText returns a boolean if a field has been set.

### SetQrCodeTextNil

`func (o *PhilippineQRInput) SetQrCodeTextNil(b bool)`

 SetQrCodeTextNil sets the value for QrCodeText to be an explicit nil

### UnsetQrCodeText
`func (o *PhilippineQRInput) UnsetQrCodeText()`

UnsetQrCodeText ensures that no value is present for QrCodeText, not even an explicit nil
### GetQrCodeImage

`func (o *PhilippineQRInput) GetQrCodeImage() string`

GetQrCodeImage returns the QrCodeImage field if non-nil, zero value otherwise.

### GetQrCodeImageOk

`func (o *PhilippineQRInput) GetQrCodeImageOk() (*string, bool)`

GetQrCodeImageOk returns a tuple with the QrCodeImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQrCodeImage

`func (o *PhilippineQRInput) SetQrCodeImage(v string)`

SetQrCodeImage sets QrCodeImage field to given value.

### HasQrCodeImage

`func (o *PhilippineQRInput) HasQrCodeImage() bool`

HasQrCodeImage returns a boolean if a field has been set.

### SetQrCodeImageNil

`func (o *PhilippineQRInput) SetQrCodeImageNil(b bool)`

 SetQrCodeImageNil sets the value for QrCodeImage to be an explicit nil

### UnsetQrCodeImage
`func (o *PhilippineQRInput) UnsetQrCodeImage()`

UnsetQrCodeImage ensures that no value is present for QrCodeImage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


