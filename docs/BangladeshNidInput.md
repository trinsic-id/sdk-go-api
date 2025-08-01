# BangladeshNidInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NationalIdNumber** | **string** | The user&#39;s Bangladesh National ID number. | 
**DateOfBirth** | **string** | The user&#39;s date of birth, in &#x60;YYYY-MM-DD&#x60; format | 
**Name** | **string** | The user&#39;s full name | 
**PhotoByes** | Pointer to **NullableString** | The raw bytes of the photo file collected from the user. | [optional] 
**PhotoImageMimeType** | Pointer to **NullableString** | The MIME Type of the file contained in &#x60;PhotoByes&#x60;.              Must be one of &#x60;image/jpeg&#x60;, or &#x60;image/png&#x60;. | [optional] 

## Methods

### NewBangladeshNidInput

`func NewBangladeshNidInput(nationalIdNumber string, dateOfBirth string, name string, ) *BangladeshNidInput`

NewBangladeshNidInput instantiates a new BangladeshNidInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBangladeshNidInputWithDefaults

`func NewBangladeshNidInputWithDefaults() *BangladeshNidInput`

NewBangladeshNidInputWithDefaults instantiates a new BangladeshNidInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNationalIdNumber

`func (o *BangladeshNidInput) GetNationalIdNumber() string`

GetNationalIdNumber returns the NationalIdNumber field if non-nil, zero value otherwise.

### GetNationalIdNumberOk

`func (o *BangladeshNidInput) GetNationalIdNumberOk() (*string, bool)`

GetNationalIdNumberOk returns a tuple with the NationalIdNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationalIdNumber

`func (o *BangladeshNidInput) SetNationalIdNumber(v string)`

SetNationalIdNumber sets NationalIdNumber field to given value.


### GetDateOfBirth

`func (o *BangladeshNidInput) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *BangladeshNidInput) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *BangladeshNidInput) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.


### GetName

`func (o *BangladeshNidInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BangladeshNidInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BangladeshNidInput) SetName(v string)`

SetName sets Name field to given value.


### GetPhotoByes

`func (o *BangladeshNidInput) GetPhotoByes() string`

GetPhotoByes returns the PhotoByes field if non-nil, zero value otherwise.

### GetPhotoByesOk

`func (o *BangladeshNidInput) GetPhotoByesOk() (*string, bool)`

GetPhotoByesOk returns a tuple with the PhotoByes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhotoByes

`func (o *BangladeshNidInput) SetPhotoByes(v string)`

SetPhotoByes sets PhotoByes field to given value.

### HasPhotoByes

`func (o *BangladeshNidInput) HasPhotoByes() bool`

HasPhotoByes returns a boolean if a field has been set.

### SetPhotoByesNil

`func (o *BangladeshNidInput) SetPhotoByesNil(b bool)`

 SetPhotoByesNil sets the value for PhotoByes to be an explicit nil

### UnsetPhotoByes
`func (o *BangladeshNidInput) UnsetPhotoByes()`

UnsetPhotoByes ensures that no value is present for PhotoByes, not even an explicit nil
### GetPhotoImageMimeType

`func (o *BangladeshNidInput) GetPhotoImageMimeType() string`

GetPhotoImageMimeType returns the PhotoImageMimeType field if non-nil, zero value otherwise.

### GetPhotoImageMimeTypeOk

`func (o *BangladeshNidInput) GetPhotoImageMimeTypeOk() (*string, bool)`

GetPhotoImageMimeTypeOk returns a tuple with the PhotoImageMimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhotoImageMimeType

`func (o *BangladeshNidInput) SetPhotoImageMimeType(v string)`

SetPhotoImageMimeType sets PhotoImageMimeType field to given value.

### HasPhotoImageMimeType

`func (o *BangladeshNidInput) HasPhotoImageMimeType() bool`

HasPhotoImageMimeType returns a boolean if a field has been set.

### SetPhotoImageMimeTypeNil

`func (o *BangladeshNidInput) SetPhotoImageMimeTypeNil(b bool)`

 SetPhotoImageMimeTypeNil sets the value for PhotoImageMimeType to be an explicit nil

### UnsetPhotoImageMimeType
`func (o *BangladeshNidInput) UnsetPhotoImageMimeType()`

UnsetPhotoImageMimeType ensures that no value is present for PhotoImageMimeType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


