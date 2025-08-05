# IndonesiaDukcapilBiometricMatchInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FullName** | Pointer to **NullableString** | The user&#39;s full name | [optional] 
**DateOfBirth** | Pointer to **NullableString** | The user&#39;s date of birth, in &#x60;YYYY-MM-DD&#x60; format | [optional] 
**NikIdNumber** | Pointer to **NullableString** | The user&#39;s Indonesia NIK ID number | [optional] 
**Email** | Pointer to **NullableString** | The email address of the individual.              Either email or phone number must be provided. | [optional] 
**PhoneNumber** | Pointer to **NullableString** | The phone number of the individual.              Either email or phone number must be provided. | [optional] 
**SelfieImage** | Pointer to **NullableString** | The raw bytes of the selfie image of the individual.              Must be JPEG format; 1MB maximum. | [optional] 
**DocumentImage** | Pointer to **NullableString** | The raw bytes of the image of the individual&#39;s KTP government ID.              Must be JPEG format; 1MB maximum.              Optional. | [optional] 
**ConsentGivenAt** | Pointer to **NullableTime** | The timestamp when consent was given by the user for the verification. | [optional] 

## Methods

### NewIndonesiaDukcapilBiometricMatchInput

`func NewIndonesiaDukcapilBiometricMatchInput() *IndonesiaDukcapilBiometricMatchInput`

NewIndonesiaDukcapilBiometricMatchInput instantiates a new IndonesiaDukcapilBiometricMatchInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndonesiaDukcapilBiometricMatchInputWithDefaults

`func NewIndonesiaDukcapilBiometricMatchInputWithDefaults() *IndonesiaDukcapilBiometricMatchInput`

NewIndonesiaDukcapilBiometricMatchInputWithDefaults instantiates a new IndonesiaDukcapilBiometricMatchInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFullName

`func (o *IndonesiaDukcapilBiometricMatchInput) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *IndonesiaDukcapilBiometricMatchInput) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *IndonesiaDukcapilBiometricMatchInput) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *IndonesiaDukcapilBiometricMatchInput) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### SetFullNameNil

`func (o *IndonesiaDukcapilBiometricMatchInput) SetFullNameNil(b bool)`

 SetFullNameNil sets the value for FullName to be an explicit nil

### UnsetFullName
`func (o *IndonesiaDukcapilBiometricMatchInput) UnsetFullName()`

UnsetFullName ensures that no value is present for FullName, not even an explicit nil
### GetDateOfBirth

`func (o *IndonesiaDukcapilBiometricMatchInput) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *IndonesiaDukcapilBiometricMatchInput) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *IndonesiaDukcapilBiometricMatchInput) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *IndonesiaDukcapilBiometricMatchInput) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### SetDateOfBirthNil

`func (o *IndonesiaDukcapilBiometricMatchInput) SetDateOfBirthNil(b bool)`

 SetDateOfBirthNil sets the value for DateOfBirth to be an explicit nil

### UnsetDateOfBirth
`func (o *IndonesiaDukcapilBiometricMatchInput) UnsetDateOfBirth()`

UnsetDateOfBirth ensures that no value is present for DateOfBirth, not even an explicit nil
### GetNikIdNumber

`func (o *IndonesiaDukcapilBiometricMatchInput) GetNikIdNumber() string`

GetNikIdNumber returns the NikIdNumber field if non-nil, zero value otherwise.

### GetNikIdNumberOk

`func (o *IndonesiaDukcapilBiometricMatchInput) GetNikIdNumberOk() (*string, bool)`

GetNikIdNumberOk returns a tuple with the NikIdNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNikIdNumber

`func (o *IndonesiaDukcapilBiometricMatchInput) SetNikIdNumber(v string)`

SetNikIdNumber sets NikIdNumber field to given value.

### HasNikIdNumber

`func (o *IndonesiaDukcapilBiometricMatchInput) HasNikIdNumber() bool`

HasNikIdNumber returns a boolean if a field has been set.

### SetNikIdNumberNil

`func (o *IndonesiaDukcapilBiometricMatchInput) SetNikIdNumberNil(b bool)`

 SetNikIdNumberNil sets the value for NikIdNumber to be an explicit nil

### UnsetNikIdNumber
`func (o *IndonesiaDukcapilBiometricMatchInput) UnsetNikIdNumber()`

UnsetNikIdNumber ensures that no value is present for NikIdNumber, not even an explicit nil
### GetEmail

`func (o *IndonesiaDukcapilBiometricMatchInput) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *IndonesiaDukcapilBiometricMatchInput) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *IndonesiaDukcapilBiometricMatchInput) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *IndonesiaDukcapilBiometricMatchInput) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *IndonesiaDukcapilBiometricMatchInput) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *IndonesiaDukcapilBiometricMatchInput) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetPhoneNumber

`func (o *IndonesiaDukcapilBiometricMatchInput) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *IndonesiaDukcapilBiometricMatchInput) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *IndonesiaDukcapilBiometricMatchInput) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *IndonesiaDukcapilBiometricMatchInput) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### SetPhoneNumberNil

`func (o *IndonesiaDukcapilBiometricMatchInput) SetPhoneNumberNil(b bool)`

 SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil

### UnsetPhoneNumber
`func (o *IndonesiaDukcapilBiometricMatchInput) UnsetPhoneNumber()`

UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
### GetSelfieImage

`func (o *IndonesiaDukcapilBiometricMatchInput) GetSelfieImage() string`

GetSelfieImage returns the SelfieImage field if non-nil, zero value otherwise.

### GetSelfieImageOk

`func (o *IndonesiaDukcapilBiometricMatchInput) GetSelfieImageOk() (*string, bool)`

GetSelfieImageOk returns a tuple with the SelfieImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfieImage

`func (o *IndonesiaDukcapilBiometricMatchInput) SetSelfieImage(v string)`

SetSelfieImage sets SelfieImage field to given value.

### HasSelfieImage

`func (o *IndonesiaDukcapilBiometricMatchInput) HasSelfieImage() bool`

HasSelfieImage returns a boolean if a field has been set.

### SetSelfieImageNil

`func (o *IndonesiaDukcapilBiometricMatchInput) SetSelfieImageNil(b bool)`

 SetSelfieImageNil sets the value for SelfieImage to be an explicit nil

### UnsetSelfieImage
`func (o *IndonesiaDukcapilBiometricMatchInput) UnsetSelfieImage()`

UnsetSelfieImage ensures that no value is present for SelfieImage, not even an explicit nil
### GetDocumentImage

`func (o *IndonesiaDukcapilBiometricMatchInput) GetDocumentImage() string`

GetDocumentImage returns the DocumentImage field if non-nil, zero value otherwise.

### GetDocumentImageOk

`func (o *IndonesiaDukcapilBiometricMatchInput) GetDocumentImageOk() (*string, bool)`

GetDocumentImageOk returns a tuple with the DocumentImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentImage

`func (o *IndonesiaDukcapilBiometricMatchInput) SetDocumentImage(v string)`

SetDocumentImage sets DocumentImage field to given value.

### HasDocumentImage

`func (o *IndonesiaDukcapilBiometricMatchInput) HasDocumentImage() bool`

HasDocumentImage returns a boolean if a field has been set.

### SetDocumentImageNil

`func (o *IndonesiaDukcapilBiometricMatchInput) SetDocumentImageNil(b bool)`

 SetDocumentImageNil sets the value for DocumentImage to be an explicit nil

### UnsetDocumentImage
`func (o *IndonesiaDukcapilBiometricMatchInput) UnsetDocumentImage()`

UnsetDocumentImage ensures that no value is present for DocumentImage, not even an explicit nil
### GetConsentGivenAt

`func (o *IndonesiaDukcapilBiometricMatchInput) GetConsentGivenAt() time.Time`

GetConsentGivenAt returns the ConsentGivenAt field if non-nil, zero value otherwise.

### GetConsentGivenAtOk

`func (o *IndonesiaDukcapilBiometricMatchInput) GetConsentGivenAtOk() (*time.Time, bool)`

GetConsentGivenAtOk returns a tuple with the ConsentGivenAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentGivenAt

`func (o *IndonesiaDukcapilBiometricMatchInput) SetConsentGivenAt(v time.Time)`

SetConsentGivenAt sets ConsentGivenAt field to given value.

### HasConsentGivenAt

`func (o *IndonesiaDukcapilBiometricMatchInput) HasConsentGivenAt() bool`

HasConsentGivenAt returns a boolean if a field has been set.

### SetConsentGivenAtNil

`func (o *IndonesiaDukcapilBiometricMatchInput) SetConsentGivenAtNil(b bool)`

 SetConsentGivenAtNil sets the value for ConsentGivenAt to be an explicit nil

### UnsetConsentGivenAt
`func (o *IndonesiaDukcapilBiometricMatchInput) UnsetConsentGivenAt()`

UnsetConsentGivenAt ensures that no value is present for ConsentGivenAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


