# IndonesiaDukcapilMatchInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FullName** | Pointer to **NullableString** | The user&#39;s full name | [optional] 
**DateOfBirth** | Pointer to **NullableString** | The user&#39;s date of birth, in &#x60;YYYY-MM-DD&#x60; format | [optional] 
**NikIdNumber** | Pointer to **NullableString** | The user&#39;s Indonesia NIK ID number | [optional] 
**Email** | Pointer to **NullableString** | The email address of the individual.              Either email or phone number must be provided. | [optional] 
**PhoneNumber** | Pointer to **NullableString** | The phone number of the individual.              Either email or phone number must be provided. | [optional] 
**SelfieImage** | Pointer to **NullableString** | The raw bytes of the selfie image of the individual.              Must be JPEG or PNG format; 10MB maximum. | [optional] 
**DocumentImage** | Pointer to **NullableString** | The raw bytes of the image of the individual&#39;s KTP government ID.              Must be JPEG format; 1MB maximum.              Optional. | [optional] 
**ConsentGivenAt** | Pointer to **NullableTime** | The timestamp when consent was given by the user for the verification. | [optional] 

## Methods

### NewIndonesiaDukcapilMatchInput

`func NewIndonesiaDukcapilMatchInput() *IndonesiaDukcapilMatchInput`

NewIndonesiaDukcapilMatchInput instantiates a new IndonesiaDukcapilMatchInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndonesiaDukcapilMatchInputWithDefaults

`func NewIndonesiaDukcapilMatchInputWithDefaults() *IndonesiaDukcapilMatchInput`

NewIndonesiaDukcapilMatchInputWithDefaults instantiates a new IndonesiaDukcapilMatchInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFullName

`func (o *IndonesiaDukcapilMatchInput) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *IndonesiaDukcapilMatchInput) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *IndonesiaDukcapilMatchInput) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *IndonesiaDukcapilMatchInput) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### SetFullNameNil

`func (o *IndonesiaDukcapilMatchInput) SetFullNameNil(b bool)`

 SetFullNameNil sets the value for FullName to be an explicit nil

### UnsetFullName
`func (o *IndonesiaDukcapilMatchInput) UnsetFullName()`

UnsetFullName ensures that no value is present for FullName, not even an explicit nil
### GetDateOfBirth

`func (o *IndonesiaDukcapilMatchInput) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *IndonesiaDukcapilMatchInput) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *IndonesiaDukcapilMatchInput) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *IndonesiaDukcapilMatchInput) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### SetDateOfBirthNil

`func (o *IndonesiaDukcapilMatchInput) SetDateOfBirthNil(b bool)`

 SetDateOfBirthNil sets the value for DateOfBirth to be an explicit nil

### UnsetDateOfBirth
`func (o *IndonesiaDukcapilMatchInput) UnsetDateOfBirth()`

UnsetDateOfBirth ensures that no value is present for DateOfBirth, not even an explicit nil
### GetNikIdNumber

`func (o *IndonesiaDukcapilMatchInput) GetNikIdNumber() string`

GetNikIdNumber returns the NikIdNumber field if non-nil, zero value otherwise.

### GetNikIdNumberOk

`func (o *IndonesiaDukcapilMatchInput) GetNikIdNumberOk() (*string, bool)`

GetNikIdNumberOk returns a tuple with the NikIdNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNikIdNumber

`func (o *IndonesiaDukcapilMatchInput) SetNikIdNumber(v string)`

SetNikIdNumber sets NikIdNumber field to given value.

### HasNikIdNumber

`func (o *IndonesiaDukcapilMatchInput) HasNikIdNumber() bool`

HasNikIdNumber returns a boolean if a field has been set.

### SetNikIdNumberNil

`func (o *IndonesiaDukcapilMatchInput) SetNikIdNumberNil(b bool)`

 SetNikIdNumberNil sets the value for NikIdNumber to be an explicit nil

### UnsetNikIdNumber
`func (o *IndonesiaDukcapilMatchInput) UnsetNikIdNumber()`

UnsetNikIdNumber ensures that no value is present for NikIdNumber, not even an explicit nil
### GetEmail

`func (o *IndonesiaDukcapilMatchInput) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *IndonesiaDukcapilMatchInput) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *IndonesiaDukcapilMatchInput) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *IndonesiaDukcapilMatchInput) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *IndonesiaDukcapilMatchInput) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *IndonesiaDukcapilMatchInput) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetPhoneNumber

`func (o *IndonesiaDukcapilMatchInput) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *IndonesiaDukcapilMatchInput) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *IndonesiaDukcapilMatchInput) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *IndonesiaDukcapilMatchInput) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### SetPhoneNumberNil

`func (o *IndonesiaDukcapilMatchInput) SetPhoneNumberNil(b bool)`

 SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil

### UnsetPhoneNumber
`func (o *IndonesiaDukcapilMatchInput) UnsetPhoneNumber()`

UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
### GetSelfieImage

`func (o *IndonesiaDukcapilMatchInput) GetSelfieImage() string`

GetSelfieImage returns the SelfieImage field if non-nil, zero value otherwise.

### GetSelfieImageOk

`func (o *IndonesiaDukcapilMatchInput) GetSelfieImageOk() (*string, bool)`

GetSelfieImageOk returns a tuple with the SelfieImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfieImage

`func (o *IndonesiaDukcapilMatchInput) SetSelfieImage(v string)`

SetSelfieImage sets SelfieImage field to given value.

### HasSelfieImage

`func (o *IndonesiaDukcapilMatchInput) HasSelfieImage() bool`

HasSelfieImage returns a boolean if a field has been set.

### SetSelfieImageNil

`func (o *IndonesiaDukcapilMatchInput) SetSelfieImageNil(b bool)`

 SetSelfieImageNil sets the value for SelfieImage to be an explicit nil

### UnsetSelfieImage
`func (o *IndonesiaDukcapilMatchInput) UnsetSelfieImage()`

UnsetSelfieImage ensures that no value is present for SelfieImage, not even an explicit nil
### GetDocumentImage

`func (o *IndonesiaDukcapilMatchInput) GetDocumentImage() string`

GetDocumentImage returns the DocumentImage field if non-nil, zero value otherwise.

### GetDocumentImageOk

`func (o *IndonesiaDukcapilMatchInput) GetDocumentImageOk() (*string, bool)`

GetDocumentImageOk returns a tuple with the DocumentImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentImage

`func (o *IndonesiaDukcapilMatchInput) SetDocumentImage(v string)`

SetDocumentImage sets DocumentImage field to given value.

### HasDocumentImage

`func (o *IndonesiaDukcapilMatchInput) HasDocumentImage() bool`

HasDocumentImage returns a boolean if a field has been set.

### SetDocumentImageNil

`func (o *IndonesiaDukcapilMatchInput) SetDocumentImageNil(b bool)`

 SetDocumentImageNil sets the value for DocumentImage to be an explicit nil

### UnsetDocumentImage
`func (o *IndonesiaDukcapilMatchInput) UnsetDocumentImage()`

UnsetDocumentImage ensures that no value is present for DocumentImage, not even an explicit nil
### GetConsentGivenAt

`func (o *IndonesiaDukcapilMatchInput) GetConsentGivenAt() time.Time`

GetConsentGivenAt returns the ConsentGivenAt field if non-nil, zero value otherwise.

### GetConsentGivenAtOk

`func (o *IndonesiaDukcapilMatchInput) GetConsentGivenAtOk() (*time.Time, bool)`

GetConsentGivenAtOk returns a tuple with the ConsentGivenAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentGivenAt

`func (o *IndonesiaDukcapilMatchInput) SetConsentGivenAt(v time.Time)`

SetConsentGivenAt sets ConsentGivenAt field to given value.

### HasConsentGivenAt

`func (o *IndonesiaDukcapilMatchInput) HasConsentGivenAt() bool`

HasConsentGivenAt returns a boolean if a field has been set.

### SetConsentGivenAtNil

`func (o *IndonesiaDukcapilMatchInput) SetConsentGivenAtNil(b bool)`

 SetConsentGivenAtNil sets the value for ConsentGivenAt to be an explicit nil

### UnsetConsentGivenAt
`func (o *IndonesiaDukcapilMatchInput) UnsetConsentGivenAt()`

UnsetConsentGivenAt ensures that no value is present for ConsentGivenAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


