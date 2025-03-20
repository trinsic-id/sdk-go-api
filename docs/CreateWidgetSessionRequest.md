# CreateWidgetSessionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RedirectUrl** | Pointer to **NullableString** | The URL to redirect the user to after the widget session is complete.                *Note*: this should NOT be set if you intend to use Trinsic&#39;s Web UI SDK to launch the Widget  as an embedded iFrame or popup; in that case, session resolution is handled by our SDK, not via redirect. | [optional] 
**Providers** | Pointer to **[]string** | The list of allowed identity providers. If not specified, all available providers will be allowed. | [optional] 
**KnownIdentityData** | Pointer to [**NullableKnownIdentityData**](KnownIdentityData.md) | Known identity data of an individual being verified.                Provide this to Trinsic during Session creation to enable improved identity provider selection recommendations. | [optional] 

## Methods

### NewCreateWidgetSessionRequest

`func NewCreateWidgetSessionRequest() *CreateWidgetSessionRequest`

NewCreateWidgetSessionRequest instantiates a new CreateWidgetSessionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateWidgetSessionRequestWithDefaults

`func NewCreateWidgetSessionRequestWithDefaults() *CreateWidgetSessionRequest`

NewCreateWidgetSessionRequestWithDefaults instantiates a new CreateWidgetSessionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRedirectUrl

`func (o *CreateWidgetSessionRequest) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *CreateWidgetSessionRequest) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *CreateWidgetSessionRequest) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *CreateWidgetSessionRequest) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### SetRedirectUrlNil

`func (o *CreateWidgetSessionRequest) SetRedirectUrlNil(b bool)`

 SetRedirectUrlNil sets the value for RedirectUrl to be an explicit nil

### UnsetRedirectUrl
`func (o *CreateWidgetSessionRequest) UnsetRedirectUrl()`

UnsetRedirectUrl ensures that no value is present for RedirectUrl, not even an explicit nil
### GetProviders

`func (o *CreateWidgetSessionRequest) GetProviders() []string`

GetProviders returns the Providers field if non-nil, zero value otherwise.

### GetProvidersOk

`func (o *CreateWidgetSessionRequest) GetProvidersOk() (*[]string, bool)`

GetProvidersOk returns a tuple with the Providers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviders

`func (o *CreateWidgetSessionRequest) SetProviders(v []string)`

SetProviders sets Providers field to given value.

### HasProviders

`func (o *CreateWidgetSessionRequest) HasProviders() bool`

HasProviders returns a boolean if a field has been set.

### SetProvidersNil

`func (o *CreateWidgetSessionRequest) SetProvidersNil(b bool)`

 SetProvidersNil sets the value for Providers to be an explicit nil

### UnsetProviders
`func (o *CreateWidgetSessionRequest) UnsetProviders()`

UnsetProviders ensures that no value is present for Providers, not even an explicit nil
### GetKnownIdentityData

`func (o *CreateWidgetSessionRequest) GetKnownIdentityData() KnownIdentityData`

GetKnownIdentityData returns the KnownIdentityData field if non-nil, zero value otherwise.

### GetKnownIdentityDataOk

`func (o *CreateWidgetSessionRequest) GetKnownIdentityDataOk() (*KnownIdentityData, bool)`

GetKnownIdentityDataOk returns a tuple with the KnownIdentityData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnownIdentityData

`func (o *CreateWidgetSessionRequest) SetKnownIdentityData(v KnownIdentityData)`

SetKnownIdentityData sets KnownIdentityData field to given value.

### HasKnownIdentityData

`func (o *CreateWidgetSessionRequest) HasKnownIdentityData() bool`

HasKnownIdentityData returns a boolean if a field has been set.

### SetKnownIdentityDataNil

`func (o *CreateWidgetSessionRequest) SetKnownIdentityDataNil(b bool)`

 SetKnownIdentityDataNil sets the value for KnownIdentityData to be an explicit nil

### UnsetKnownIdentityData
`func (o *CreateWidgetSessionRequest) UnsetKnownIdentityData()`

UnsetKnownIdentityData ensures that no value is present for KnownIdentityData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


