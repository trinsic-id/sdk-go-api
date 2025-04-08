# BrazilDigitalCnhInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpfNumber** | **string** | The user&#39;s 11-digit, numeric CPF Number | 
**DigitalCnhFile** | Pointer to **NullableString** | The raw bytes of the digital CNH file collected from the user.              TODO: Lucas or JP help me describe these below vvvvvvvv This can be: - An image containing a physical or digital QR code - A PDF file exported from the CNH app - Some scary third thing? | [optional] 
**DigitalCnhFileContentType** | Pointer to **NullableString** | The MIME Type of the file contained in &#x60;DigitalCnhFile&#x60;.              Must be one of &#x60;application/pdf&#x60;, &#x60;image/jpeg&#x60;, or &#x60;image/png&#x60;. | [optional] 
**FacialBiometryPhoto** | Pointer to **NullableString** | The raw bytes of the image of the user&#39;s face, collected for biometric comparison. | [optional] 

## Methods

### NewBrazilDigitalCnhInput

`func NewBrazilDigitalCnhInput(cpfNumber string, ) *BrazilDigitalCnhInput`

NewBrazilDigitalCnhInput instantiates a new BrazilDigitalCnhInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrazilDigitalCnhInputWithDefaults

`func NewBrazilDigitalCnhInputWithDefaults() *BrazilDigitalCnhInput`

NewBrazilDigitalCnhInputWithDefaults instantiates a new BrazilDigitalCnhInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpfNumber

`func (o *BrazilDigitalCnhInput) GetCpfNumber() string`

GetCpfNumber returns the CpfNumber field if non-nil, zero value otherwise.

### GetCpfNumberOk

`func (o *BrazilDigitalCnhInput) GetCpfNumberOk() (*string, bool)`

GetCpfNumberOk returns a tuple with the CpfNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpfNumber

`func (o *BrazilDigitalCnhInput) SetCpfNumber(v string)`

SetCpfNumber sets CpfNumber field to given value.


### GetDigitalCnhFile

`func (o *BrazilDigitalCnhInput) GetDigitalCnhFile() string`

GetDigitalCnhFile returns the DigitalCnhFile field if non-nil, zero value otherwise.

### GetDigitalCnhFileOk

`func (o *BrazilDigitalCnhInput) GetDigitalCnhFileOk() (*string, bool)`

GetDigitalCnhFileOk returns a tuple with the DigitalCnhFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigitalCnhFile

`func (o *BrazilDigitalCnhInput) SetDigitalCnhFile(v string)`

SetDigitalCnhFile sets DigitalCnhFile field to given value.

### HasDigitalCnhFile

`func (o *BrazilDigitalCnhInput) HasDigitalCnhFile() bool`

HasDigitalCnhFile returns a boolean if a field has been set.

### SetDigitalCnhFileNil

`func (o *BrazilDigitalCnhInput) SetDigitalCnhFileNil(b bool)`

 SetDigitalCnhFileNil sets the value for DigitalCnhFile to be an explicit nil

### UnsetDigitalCnhFile
`func (o *BrazilDigitalCnhInput) UnsetDigitalCnhFile()`

UnsetDigitalCnhFile ensures that no value is present for DigitalCnhFile, not even an explicit nil
### GetDigitalCnhFileContentType

`func (o *BrazilDigitalCnhInput) GetDigitalCnhFileContentType() string`

GetDigitalCnhFileContentType returns the DigitalCnhFileContentType field if non-nil, zero value otherwise.

### GetDigitalCnhFileContentTypeOk

`func (o *BrazilDigitalCnhInput) GetDigitalCnhFileContentTypeOk() (*string, bool)`

GetDigitalCnhFileContentTypeOk returns a tuple with the DigitalCnhFileContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigitalCnhFileContentType

`func (o *BrazilDigitalCnhInput) SetDigitalCnhFileContentType(v string)`

SetDigitalCnhFileContentType sets DigitalCnhFileContentType field to given value.

### HasDigitalCnhFileContentType

`func (o *BrazilDigitalCnhInput) HasDigitalCnhFileContentType() bool`

HasDigitalCnhFileContentType returns a boolean if a field has been set.

### SetDigitalCnhFileContentTypeNil

`func (o *BrazilDigitalCnhInput) SetDigitalCnhFileContentTypeNil(b bool)`

 SetDigitalCnhFileContentTypeNil sets the value for DigitalCnhFileContentType to be an explicit nil

### UnsetDigitalCnhFileContentType
`func (o *BrazilDigitalCnhInput) UnsetDigitalCnhFileContentType()`

UnsetDigitalCnhFileContentType ensures that no value is present for DigitalCnhFileContentType, not even an explicit nil
### GetFacialBiometryPhoto

`func (o *BrazilDigitalCnhInput) GetFacialBiometryPhoto() string`

GetFacialBiometryPhoto returns the FacialBiometryPhoto field if non-nil, zero value otherwise.

### GetFacialBiometryPhotoOk

`func (o *BrazilDigitalCnhInput) GetFacialBiometryPhotoOk() (*string, bool)`

GetFacialBiometryPhotoOk returns a tuple with the FacialBiometryPhoto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacialBiometryPhoto

`func (o *BrazilDigitalCnhInput) SetFacialBiometryPhoto(v string)`

SetFacialBiometryPhoto sets FacialBiometryPhoto field to given value.

### HasFacialBiometryPhoto

`func (o *BrazilDigitalCnhInput) HasFacialBiometryPhoto() bool`

HasFacialBiometryPhoto returns a boolean if a field has been set.

### SetFacialBiometryPhotoNil

`func (o *BrazilDigitalCnhInput) SetFacialBiometryPhotoNil(b bool)`

 SetFacialBiometryPhotoNil sets the value for FacialBiometryPhoto to be an explicit nil

### UnsetFacialBiometryPhoto
`func (o *BrazilDigitalCnhInput) UnsetFacialBiometryPhoto()`

UnsetFacialBiometryPhoto ensures that no value is present for FacialBiometryPhoto, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


