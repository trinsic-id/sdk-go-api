# ProviderInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the provider | 
**Name** | **string** | The friendly, human-readable name of the provider | 
**LogoUrl** | **string** | A URL pointing to the provider&#39;s logo | 
**Subtext** | **string** | The Provider&#39;s subtext recommended to be shown next to the name.              This is flavor text, not a full, human-readable description of the provider. | 
**Description** | **string** | The Provider&#39;s description recommended to be shown next to the name.              This is flavor text, not a full, human-readable description of the provider. | 
**Health** | **string** | The current health status of the provider | 
**SubProviders** | Pointer to [**[]SubProviderMetadata**](SubProviderMetadata.md) | Metadata about the sub-providers which are available for this Provider.              For example, Italy&#39;s SPID is a Provider which aggregates access to multiple sub-providers. | [optional] 

## Methods

### NewProviderInfo

`func NewProviderInfo(id string, name string, logoUrl string, subtext string, description string, health string, ) *ProviderInfo`

NewProviderInfo instantiates a new ProviderInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderInfoWithDefaults

`func NewProviderInfoWithDefaults() *ProviderInfo`

NewProviderInfoWithDefaults instantiates a new ProviderInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProviderInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProviderInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProviderInfo) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ProviderInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProviderInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProviderInfo) SetName(v string)`

SetName sets Name field to given value.


### GetLogoUrl

`func (o *ProviderInfo) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *ProviderInfo) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *ProviderInfo) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.


### GetSubtext

`func (o *ProviderInfo) GetSubtext() string`

GetSubtext returns the Subtext field if non-nil, zero value otherwise.

### GetSubtextOk

`func (o *ProviderInfo) GetSubtextOk() (*string, bool)`

GetSubtextOk returns a tuple with the Subtext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtext

`func (o *ProviderInfo) SetSubtext(v string)`

SetSubtext sets Subtext field to given value.


### GetDescription

`func (o *ProviderInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProviderInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProviderInfo) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetHealth

`func (o *ProviderInfo) GetHealth() string`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *ProviderInfo) GetHealthOk() (*string, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *ProviderInfo) SetHealth(v string)`

SetHealth sets Health field to given value.


### GetSubProviders

`func (o *ProviderInfo) GetSubProviders() []SubProviderMetadata`

GetSubProviders returns the SubProviders field if non-nil, zero value otherwise.

### GetSubProvidersOk

`func (o *ProviderInfo) GetSubProvidersOk() (*[]SubProviderMetadata, bool)`

GetSubProvidersOk returns a tuple with the SubProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubProviders

`func (o *ProviderInfo) SetSubProviders(v []SubProviderMetadata)`

SetSubProviders sets SubProviders field to given value.

### HasSubProviders

`func (o *ProviderInfo) HasSubProviders() bool`

HasSubProviders returns a boolean if a field has been set.

### SetSubProvidersNil

`func (o *ProviderInfo) SetSubProvidersNil(b bool)`

 SetSubProvidersNil sets the value for SubProviders to be an explicit nil

### UnsetSubProviders
`func (o *ProviderInfo) UnsetSubProviders()`

UnsetSubProviders ensures that no value is present for SubProviders, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


