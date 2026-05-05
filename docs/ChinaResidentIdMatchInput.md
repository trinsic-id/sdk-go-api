# ChinaResidentIdMatchInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FullName** | **string** | The person&#39;s full name in Chinese characters. | 
**DateOfBirth** | **string** | The person&#39;s date of birth, in &#x60;YYYY-MM-DD&#x60; format. | 
**NationalIdNumber** | **string** | The person&#39;s 18-character PRC resident identity card number. | 
**Consent** | **bool** | Consent from the end user for verification of their data. | 

## Methods

### NewChinaResidentIdMatchInput

`func NewChinaResidentIdMatchInput(fullName string, dateOfBirth string, nationalIdNumber string, consent bool, ) *ChinaResidentIdMatchInput`

NewChinaResidentIdMatchInput instantiates a new ChinaResidentIdMatchInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChinaResidentIdMatchInputWithDefaults

`func NewChinaResidentIdMatchInputWithDefaults() *ChinaResidentIdMatchInput`

NewChinaResidentIdMatchInputWithDefaults instantiates a new ChinaResidentIdMatchInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFullName

`func (o *ChinaResidentIdMatchInput) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *ChinaResidentIdMatchInput) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *ChinaResidentIdMatchInput) SetFullName(v string)`

SetFullName sets FullName field to given value.


### GetDateOfBirth

`func (o *ChinaResidentIdMatchInput) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *ChinaResidentIdMatchInput) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *ChinaResidentIdMatchInput) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.


### GetNationalIdNumber

`func (o *ChinaResidentIdMatchInput) GetNationalIdNumber() string`

GetNationalIdNumber returns the NationalIdNumber field if non-nil, zero value otherwise.

### GetNationalIdNumberOk

`func (o *ChinaResidentIdMatchInput) GetNationalIdNumberOk() (*string, bool)`

GetNationalIdNumberOk returns a tuple with the NationalIdNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationalIdNumber

`func (o *ChinaResidentIdMatchInput) SetNationalIdNumber(v string)`

SetNationalIdNumber sets NationalIdNumber field to given value.


### GetConsent

`func (o *ChinaResidentIdMatchInput) GetConsent() bool`

GetConsent returns the Consent field if non-nil, zero value otherwise.

### GetConsentOk

`func (o *ChinaResidentIdMatchInput) GetConsentOk() (*bool, bool)`

GetConsentOk returns a tuple with the Consent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsent

`func (o *ChinaResidentIdMatchInput) SetConsent(v bool)`

SetConsent sets Consent field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


