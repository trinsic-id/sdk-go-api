# CreateHostedProviderSessionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | **string** | The ID of the provider to launch | 
**RedirectUrl** | **string** | The Redirect URL to which the user should be sent after the session is complete. | 
**ProviderInput** | Pointer to [**NullableProviderInput**](ProviderInput.md) | Provider-specific input for those providers which require it. | [optional] 

## Methods

### NewCreateHostedProviderSessionRequest

`func NewCreateHostedProviderSessionRequest(provider string, redirectUrl string, ) *CreateHostedProviderSessionRequest`

NewCreateHostedProviderSessionRequest instantiates a new CreateHostedProviderSessionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateHostedProviderSessionRequestWithDefaults

`func NewCreateHostedProviderSessionRequestWithDefaults() *CreateHostedProviderSessionRequest`

NewCreateHostedProviderSessionRequestWithDefaults instantiates a new CreateHostedProviderSessionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *CreateHostedProviderSessionRequest) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CreateHostedProviderSessionRequest) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CreateHostedProviderSessionRequest) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetRedirectUrl

`func (o *CreateHostedProviderSessionRequest) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *CreateHostedProviderSessionRequest) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *CreateHostedProviderSessionRequest) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.


### GetProviderInput

`func (o *CreateHostedProviderSessionRequest) GetProviderInput() ProviderInput`

GetProviderInput returns the ProviderInput field if non-nil, zero value otherwise.

### GetProviderInputOk

`func (o *CreateHostedProviderSessionRequest) GetProviderInputOk() (*ProviderInput, bool)`

GetProviderInputOk returns a tuple with the ProviderInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderInput

`func (o *CreateHostedProviderSessionRequest) SetProviderInput(v ProviderInput)`

SetProviderInput sets ProviderInput field to given value.

### HasProviderInput

`func (o *CreateHostedProviderSessionRequest) HasProviderInput() bool`

HasProviderInput returns a boolean if a field has been set.

### SetProviderInputNil

`func (o *CreateHostedProviderSessionRequest) SetProviderInputNil(b bool)`

 SetProviderInputNil sets the value for ProviderInput to be an explicit nil

### UnsetProviderInput
`func (o *CreateHostedProviderSessionRequest) UnsetProviderInput()`

UnsetProviderInput ensures that no value is present for ProviderInput, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


