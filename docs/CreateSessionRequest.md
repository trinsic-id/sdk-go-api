# CreateSessionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LaunchProviderDirectly** | Pointer to **bool** | Whether to immediately launch the identity provider, without invoking the Trinsic Widget UI.                Users will not be shown the Widget; therefore, reuse of credentials, selection of an identity provider, and saving a verification for future reuse  are not available to the end user in this mode.                Sessions created with this option enabled must be created with a &#x60;RedirectUrl&#x60; specified, and cannot be invoked using the frontend SDK at this time. | [optional] 
**EnableRememberMe** | Pointer to **bool** | Whether to enable Trinsic&#39;s \&quot;Remember Me\&quot; feature, which allows users to save their credentials for future use.                This option is only relevant when &#x60;LaunchProviderDirectly&#x60; is unspecified or set to &#x60;false&#x60;.  If &#x60;LaunchProviderDirectly&#x60; is &#x60;true&#x60;, this field must be unspecified or set to &#x60;false&#x60;.                If this field is set to &#x60;true&#x60;, then:    - The user will be prompted to authenticate with their phone number at the start of the flow    - If the user has previously saved a verification for reuse with Trinsic, they will be offered the ability to reuse it    - After the user has verified their identity (and if the identity provider in question supports it), they will be prompted to save their credentials for future use                If this field is set to &#x60;false&#x60;, then:    - The user will not be prompted to authenticate with their phone number at the start of the flow.      - Instead, the user will be immediately shown the list of available providers    - The user will not be offered the ability to reuse a previously-saved Trinsic credential    - After the user has verified their identity, they will not be prompted to save their credentials for future use      - Instead, they will immediately return to your product | [optional] 
**Providers** | Pointer to **[]string** | The list of allowed identity providers. If not specified, all available providers will be allowed.                If &#x60;LaunchMethodDirectly&#x60; is &#x60;true&#x60;, this field must be set, and must have only a single entry.  If &#x60;LaunchMethodDirectly&#x60; is not specified or is &#x60;false&#x60;, this field may have any number of entries. | [optional] 
**KnownIdentityData** | Pointer to [**KnownIdentityData**](KnownIdentityData.md) | Known identity data of an individual being verified.                Provide this to Trinsic during Session creation to enable improved identity provider selection recommendations. | [optional] 
**DisclosedFields** | Pointer to [**DisclosedFieldsRequest**](DisclosedFieldsRequest.md) | Specific identity attributes to request. If not provided, all available attributes will be requested. | [optional] 

## Methods

### NewCreateSessionRequest

`func NewCreateSessionRequest() *CreateSessionRequest`

NewCreateSessionRequest instantiates a new CreateSessionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSessionRequestWithDefaults

`func NewCreateSessionRequestWithDefaults() *CreateSessionRequest`

NewCreateSessionRequestWithDefaults instantiates a new CreateSessionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLaunchProviderDirectly

`func (o *CreateSessionRequest) GetLaunchProviderDirectly() bool`

GetLaunchProviderDirectly returns the LaunchProviderDirectly field if non-nil, zero value otherwise.

### GetLaunchProviderDirectlyOk

`func (o *CreateSessionRequest) GetLaunchProviderDirectlyOk() (*bool, bool)`

GetLaunchProviderDirectlyOk returns a tuple with the LaunchProviderDirectly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchProviderDirectly

`func (o *CreateSessionRequest) SetLaunchProviderDirectly(v bool)`

SetLaunchProviderDirectly sets LaunchProviderDirectly field to given value.

### HasLaunchProviderDirectly

`func (o *CreateSessionRequest) HasLaunchProviderDirectly() bool`

HasLaunchProviderDirectly returns a boolean if a field has been set.

### GetEnableRememberMe

`func (o *CreateSessionRequest) GetEnableRememberMe() bool`

GetEnableRememberMe returns the EnableRememberMe field if non-nil, zero value otherwise.

### GetEnableRememberMeOk

`func (o *CreateSessionRequest) GetEnableRememberMeOk() (*bool, bool)`

GetEnableRememberMeOk returns a tuple with the EnableRememberMe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRememberMe

`func (o *CreateSessionRequest) SetEnableRememberMe(v bool)`

SetEnableRememberMe sets EnableRememberMe field to given value.

### HasEnableRememberMe

`func (o *CreateSessionRequest) HasEnableRememberMe() bool`

HasEnableRememberMe returns a boolean if a field has been set.

### GetProviders

`func (o *CreateSessionRequest) GetProviders() []string`

GetProviders returns the Providers field if non-nil, zero value otherwise.

### GetProvidersOk

`func (o *CreateSessionRequest) GetProvidersOk() (*[]string, bool)`

GetProvidersOk returns a tuple with the Providers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviders

`func (o *CreateSessionRequest) SetProviders(v []string)`

SetProviders sets Providers field to given value.

### HasProviders

`func (o *CreateSessionRequest) HasProviders() bool`

HasProviders returns a boolean if a field has been set.

### GetKnownIdentityData

`func (o *CreateSessionRequest) GetKnownIdentityData() KnownIdentityData`

GetKnownIdentityData returns the KnownIdentityData field if non-nil, zero value otherwise.

### GetKnownIdentityDataOk

`func (o *CreateSessionRequest) GetKnownIdentityDataOk() (*KnownIdentityData, bool)`

GetKnownIdentityDataOk returns a tuple with the KnownIdentityData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnownIdentityData

`func (o *CreateSessionRequest) SetKnownIdentityData(v KnownIdentityData)`

SetKnownIdentityData sets KnownIdentityData field to given value.

### HasKnownIdentityData

`func (o *CreateSessionRequest) HasKnownIdentityData() bool`

HasKnownIdentityData returns a boolean if a field has been set.

### GetDisclosedFields

`func (o *CreateSessionRequest) GetDisclosedFields() DisclosedFieldsRequest`

GetDisclosedFields returns the DisclosedFields field if non-nil, zero value otherwise.

### GetDisclosedFieldsOk

`func (o *CreateSessionRequest) GetDisclosedFieldsOk() (*DisclosedFieldsRequest, bool)`

GetDisclosedFieldsOk returns a tuple with the DisclosedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisclosedFields

`func (o *CreateSessionRequest) SetDisclosedFields(v DisclosedFieldsRequest)`

SetDisclosedFields sets DisclosedFields field to given value.

### HasDisclosedFields

`func (o *CreateSessionRequest) HasDisclosedFields() bool`

HasDisclosedFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


