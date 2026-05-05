# ProviderSupportedLanguage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | BCP 47 language code. | 
**Label** | Pointer to **NullableString** | Human-readable display label for the language. | [optional] 

## Methods

### NewProviderSupportedLanguage

`func NewProviderSupportedLanguage(code string, ) *ProviderSupportedLanguage`

NewProviderSupportedLanguage instantiates a new ProviderSupportedLanguage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderSupportedLanguageWithDefaults

`func NewProviderSupportedLanguageWithDefaults() *ProviderSupportedLanguage`

NewProviderSupportedLanguageWithDefaults instantiates a new ProviderSupportedLanguage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ProviderSupportedLanguage) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ProviderSupportedLanguage) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ProviderSupportedLanguage) SetCode(v string)`

SetCode sets Code field to given value.


### GetLabel

`func (o *ProviderSupportedLanguage) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ProviderSupportedLanguage) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ProviderSupportedLanguage) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ProviderSupportedLanguage) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### SetLabelNil

`func (o *ProviderSupportedLanguage) SetLabelNil(b bool)`

 SetLabelNil sets the value for Label to be an explicit nil

### UnsetLabel
`func (o *ProviderSupportedLanguage) UnsetLabel()`

UnsetLabel ensures that no value is present for Label, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


