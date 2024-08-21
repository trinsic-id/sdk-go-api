# CreateSessionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LaunchMethodDirectly** | Pointer to **bool** | Whether to immediately launch the identity provider, without invoking the Trinsic Connect Widget UI.                Users will not be shown the Connect Widget; therefore, reuse of Connect credentials, selection of an identity provider, and saving a verification for future reuse  are not available to the end user in this mode.                Sessions created with this option enabled must be created with a &#x60;RedirectUrl&#x60; specified, and cannot be invoked using the frontend SDK at this time. | [optional] 
**Providers** | Pointer to **[]string** | The list of allowed identity providers. If not specified, all available providers will be allowed.                If &#x60;LaunchMethodDirectly&#x60; is &#x60;true&#x60;, this field must be set, and must have only a single entry.  If &#x60;LaunchMethodDirectly&#x60; is not specified or is &#x60;false&#x60;, this field may have any number of entries. | [optional] 
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

### GetLaunchMethodDirectly

`func (o *CreateSessionRequest) GetLaunchMethodDirectly() bool`

GetLaunchMethodDirectly returns the LaunchMethodDirectly field if non-nil, zero value otherwise.

### GetLaunchMethodDirectlyOk

`func (o *CreateSessionRequest) GetLaunchMethodDirectlyOk() (*bool, bool)`

GetLaunchMethodDirectlyOk returns a tuple with the LaunchMethodDirectly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchMethodDirectly

`func (o *CreateSessionRequest) SetLaunchMethodDirectly(v bool)`

SetLaunchMethodDirectly sets LaunchMethodDirectly field to given value.

### HasLaunchMethodDirectly

`func (o *CreateSessionRequest) HasLaunchMethodDirectly() bool`

HasLaunchMethodDirectly returns a boolean if a field has been set.

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


