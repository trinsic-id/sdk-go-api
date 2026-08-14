# AadhaarLanguage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LanguageCode** | Pointer to **NullableString** | The language code for the localized claims. | [optional] 
**LanguageName** | Pointer to **NullableString** | The language name for the localized claims.              This is Trinsic mapping the language from the code. The mapping comes from the spec, however, the code may not always reflect with the actual language used in the localized claims. Use this as a reference. Possible values: - Assamese (01) - Bengali (02) - Gujarati (05) - Hindi (06) - Kannada (07) - Malayalam (11) - Manipuri (12) - Marathi (13) - Oriya (15) - Punjabi (16) - Tamil (20) - Telugu (21) - Urdu (22) | [optional] [readonly] 

## Methods

### NewAadhaarLanguage

`func NewAadhaarLanguage() *AadhaarLanguage`

NewAadhaarLanguage instantiates a new AadhaarLanguage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAadhaarLanguageWithDefaults

`func NewAadhaarLanguageWithDefaults() *AadhaarLanguage`

NewAadhaarLanguageWithDefaults instantiates a new AadhaarLanguage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLanguageCode

`func (o *AadhaarLanguage) GetLanguageCode() string`

GetLanguageCode returns the LanguageCode field if non-nil, zero value otherwise.

### GetLanguageCodeOk

`func (o *AadhaarLanguage) GetLanguageCodeOk() (*string, bool)`

GetLanguageCodeOk returns a tuple with the LanguageCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageCode

`func (o *AadhaarLanguage) SetLanguageCode(v string)`

SetLanguageCode sets LanguageCode field to given value.

### HasLanguageCode

`func (o *AadhaarLanguage) HasLanguageCode() bool`

HasLanguageCode returns a boolean if a field has been set.

### SetLanguageCodeNil

`func (o *AadhaarLanguage) SetLanguageCodeNil(b bool)`

 SetLanguageCodeNil sets the value for LanguageCode to be an explicit nil

### UnsetLanguageCode
`func (o *AadhaarLanguage) UnsetLanguageCode()`

UnsetLanguageCode ensures that no value is present for LanguageCode, not even an explicit nil
### GetLanguageName

`func (o *AadhaarLanguage) GetLanguageName() string`

GetLanguageName returns the LanguageName field if non-nil, zero value otherwise.

### GetLanguageNameOk

`func (o *AadhaarLanguage) GetLanguageNameOk() (*string, bool)`

GetLanguageNameOk returns a tuple with the LanguageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageName

`func (o *AadhaarLanguage) SetLanguageName(v string)`

SetLanguageName sets LanguageName field to given value.

### HasLanguageName

`func (o *AadhaarLanguage) HasLanguageName() bool`

HasLanguageName returns a boolean if a field has been set.

### SetLanguageNameNil

`func (o *AadhaarLanguage) SetLanguageNameNil(b bool)`

 SetLanguageNameNil sets the value for LanguageName to be an explicit nil

### UnsetLanguageName
`func (o *AadhaarLanguage) UnsetLanguageName()`

UnsetLanguageName ensures that no value is present for LanguageName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


