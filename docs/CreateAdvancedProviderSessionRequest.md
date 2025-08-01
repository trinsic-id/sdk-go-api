# CreateAdvancedProviderSessionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | **string** | The ID of the provider to launch | 
**RedirectUrl** | Pointer to **NullableString** | The Redirect URL to which the user should be sent after the session is complete.              This field is required for providers which employ a redirect-based flow. | [optional] 
**Capabilities** | [**[]IntegrationCapability**](IntegrationCapability.md) | The list of capabilities your integration supports. Capabilities are the core of Trinsic&#39;s whitelabel-with-optional-fallback offering.              Most capabilities align with either an &#x60;IntegrationLaunchMethod&#x60; or an &#x60;IntegrationCollectionMethod&#x60;. The exception being refresh content to support updating the content of the launch method.              For example, to support a basic redirect-based flow, you must include the &#x60;LaunchRedirect&#x60; and &#x60;CaptureRedirect&#x60; capabilities. To support a mobile deeplink / polling flow, you must include the &#x60;DeeplinkToMobile&#x60; and &#x60;PollForResults&#x60; capabilities.              If &#x60;FallbackToHostedUi&#x60; is &#x60;true&#x60;, Trinsic will automatically fall back to a Trinsic-hosted UI to cover any gaps in your integration&#39;s capabilities. If &#x60;FallbackToHostedUi&#x60; is &#x60;false&#x60;, gaps in your integration&#39;s capabilities will result in an error during Session creation.              Read more on how to integrate at &lt;a href&#x3D;\&quot;https://docs.trinsic.id/docs/advanced-provider-sessions\&quot;&gt;the guide on Advanced Provider Sessions&lt;/a&gt; | 
**FallbackToHostedUI** | Pointer to **NullableBool** | Whether the session should fall back to a Trinsic-hosted UI in certain instances.              Specifically, fallback will occur if any of the following are true: - You attempted to launch a provider which requires a capability you did not express support for     - In this case, Trinsic&#39;s hosted UI will perform the necessary capability - You attempted to launch a provider which requires input, and the input was either not provided or incomplete     - In this case, Trinsic&#39;s hosted UI will collect the necessary input from the user              If fallback occurs, the session&#39;s NextStep will always be LaunchBrowser, and the CollectionMethod will always be CaptureRedirect.              If this field is set to &#x60;true&#x60;, you must also: 1. Set the &#x60;RedirectUrl&#x60; field to a non-empty value 2. Include the &#x60;LaunchBrowser&#x60; and &#x60;CaptureRedirect&#x60; capabilities in the &#x60;Capabilities&#x60; field | [optional] 
**ProviderInput** | Pointer to [**NullableProviderInput**](ProviderInput.md) | Provider-specific input for those providers which require it. | [optional] 

## Methods

### NewCreateAdvancedProviderSessionRequest

`func NewCreateAdvancedProviderSessionRequest(provider string, capabilities []IntegrationCapability, ) *CreateAdvancedProviderSessionRequest`

NewCreateAdvancedProviderSessionRequest instantiates a new CreateAdvancedProviderSessionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAdvancedProviderSessionRequestWithDefaults

`func NewCreateAdvancedProviderSessionRequestWithDefaults() *CreateAdvancedProviderSessionRequest`

NewCreateAdvancedProviderSessionRequestWithDefaults instantiates a new CreateAdvancedProviderSessionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *CreateAdvancedProviderSessionRequest) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CreateAdvancedProviderSessionRequest) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CreateAdvancedProviderSessionRequest) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetRedirectUrl

`func (o *CreateAdvancedProviderSessionRequest) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *CreateAdvancedProviderSessionRequest) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *CreateAdvancedProviderSessionRequest) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *CreateAdvancedProviderSessionRequest) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### SetRedirectUrlNil

`func (o *CreateAdvancedProviderSessionRequest) SetRedirectUrlNil(b bool)`

 SetRedirectUrlNil sets the value for RedirectUrl to be an explicit nil

### UnsetRedirectUrl
`func (o *CreateAdvancedProviderSessionRequest) UnsetRedirectUrl()`

UnsetRedirectUrl ensures that no value is present for RedirectUrl, not even an explicit nil
### GetCapabilities

`func (o *CreateAdvancedProviderSessionRequest) GetCapabilities() []IntegrationCapability`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *CreateAdvancedProviderSessionRequest) GetCapabilitiesOk() (*[]IntegrationCapability, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *CreateAdvancedProviderSessionRequest) SetCapabilities(v []IntegrationCapability)`

SetCapabilities sets Capabilities field to given value.


### GetFallbackToHostedUI

`func (o *CreateAdvancedProviderSessionRequest) GetFallbackToHostedUI() bool`

GetFallbackToHostedUI returns the FallbackToHostedUI field if non-nil, zero value otherwise.

### GetFallbackToHostedUIOk

`func (o *CreateAdvancedProviderSessionRequest) GetFallbackToHostedUIOk() (*bool, bool)`

GetFallbackToHostedUIOk returns a tuple with the FallbackToHostedUI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackToHostedUI

`func (o *CreateAdvancedProviderSessionRequest) SetFallbackToHostedUI(v bool)`

SetFallbackToHostedUI sets FallbackToHostedUI field to given value.

### HasFallbackToHostedUI

`func (o *CreateAdvancedProviderSessionRequest) HasFallbackToHostedUI() bool`

HasFallbackToHostedUI returns a boolean if a field has been set.

### SetFallbackToHostedUINil

`func (o *CreateAdvancedProviderSessionRequest) SetFallbackToHostedUINil(b bool)`

 SetFallbackToHostedUINil sets the value for FallbackToHostedUI to be an explicit nil

### UnsetFallbackToHostedUI
`func (o *CreateAdvancedProviderSessionRequest) UnsetFallbackToHostedUI()`

UnsetFallbackToHostedUI ensures that no value is present for FallbackToHostedUI, not even an explicit nil
### GetProviderInput

`func (o *CreateAdvancedProviderSessionRequest) GetProviderInput() ProviderInput`

GetProviderInput returns the ProviderInput field if non-nil, zero value otherwise.

### GetProviderInputOk

`func (o *CreateAdvancedProviderSessionRequest) GetProviderInputOk() (*ProviderInput, bool)`

GetProviderInputOk returns a tuple with the ProviderInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderInput

`func (o *CreateAdvancedProviderSessionRequest) SetProviderInput(v ProviderInput)`

SetProviderInput sets ProviderInput field to given value.

### HasProviderInput

`func (o *CreateAdvancedProviderSessionRequest) HasProviderInput() bool`

HasProviderInput returns a boolean if a field has been set.

### SetProviderInputNil

`func (o *CreateAdvancedProviderSessionRequest) SetProviderInputNil(b bool)`

 SetProviderInputNil sets the value for ProviderInput to be an explicit nil

### UnsetProviderInput
`func (o *CreateAdvancedProviderSessionRequest) UnsetProviderInput()`

UnsetProviderInput ensures that no value is present for ProviderInput, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


