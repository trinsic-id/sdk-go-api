# ProviderInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the provider | 
**Name** | **string** | The friendly, human-readable name of the provider | 
**LogoUrl** | **string** | A URL pointing to the provider&#39;s logo | 
**Subtext** | **string** | The Provider&#39;s subtext recommended to be shown next to the name.              This is flavor text, not a full, human-readable description of the provider. | 
**ProviderId** | **string** | The ID of the provider | 
**ProviderDisplayName** | **string** | The friendly, human-readable name of the provider | 
**ProviderLogo** | **string** | A URL pointing to the provider&#39;s logo | 
**Health** | **string** | The current health status of the provider | 
**SubProviders** | Pointer to [**[]SubProviderMetadata**](SubProviderMetadata.md) | Metadata about the sub-providers which are available for this Provider.              For example, Italy&#39;s SPID is a Provider which aggregates access to multiple sub-providers. | [optional] 

## Methods

### NewProviderInformation

`func NewProviderInformation(id string, name string, logoUrl string, subtext string, providerId string, providerDisplayName string, providerLogo string, health string, ) *ProviderInformation`

NewProviderInformation instantiates a new ProviderInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderInformationWithDefaults

`func NewProviderInformationWithDefaults() *ProviderInformation`

NewProviderInformationWithDefaults instantiates a new ProviderInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProviderInformation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProviderInformation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProviderInformation) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ProviderInformation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProviderInformation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProviderInformation) SetName(v string)`

SetName sets Name field to given value.


### GetLogoUrl

`func (o *ProviderInformation) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *ProviderInformation) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *ProviderInformation) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.


### GetSubtext

`func (o *ProviderInformation) GetSubtext() string`

GetSubtext returns the Subtext field if non-nil, zero value otherwise.

### GetSubtextOk

`func (o *ProviderInformation) GetSubtextOk() (*string, bool)`

GetSubtextOk returns a tuple with the Subtext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtext

`func (o *ProviderInformation) SetSubtext(v string)`

SetSubtext sets Subtext field to given value.


### GetProviderId

`func (o *ProviderInformation) GetProviderId() string`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *ProviderInformation) GetProviderIdOk() (*string, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *ProviderInformation) SetProviderId(v string)`

SetProviderId sets ProviderId field to given value.


### GetProviderDisplayName

`func (o *ProviderInformation) GetProviderDisplayName() string`

GetProviderDisplayName returns the ProviderDisplayName field if non-nil, zero value otherwise.

### GetProviderDisplayNameOk

`func (o *ProviderInformation) GetProviderDisplayNameOk() (*string, bool)`

GetProviderDisplayNameOk returns a tuple with the ProviderDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderDisplayName

`func (o *ProviderInformation) SetProviderDisplayName(v string)`

SetProviderDisplayName sets ProviderDisplayName field to given value.


### GetProviderLogo

`func (o *ProviderInformation) GetProviderLogo() string`

GetProviderLogo returns the ProviderLogo field if non-nil, zero value otherwise.

### GetProviderLogoOk

`func (o *ProviderInformation) GetProviderLogoOk() (*string, bool)`

GetProviderLogoOk returns a tuple with the ProviderLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderLogo

`func (o *ProviderInformation) SetProviderLogo(v string)`

SetProviderLogo sets ProviderLogo field to given value.


### GetHealth

`func (o *ProviderInformation) GetHealth() string`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *ProviderInformation) GetHealthOk() (*string, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *ProviderInformation) SetHealth(v string)`

SetHealth sets Health field to given value.


### GetSubProviders

`func (o *ProviderInformation) GetSubProviders() []SubProviderMetadata`

GetSubProviders returns the SubProviders field if non-nil, zero value otherwise.

### GetSubProvidersOk

`func (o *ProviderInformation) GetSubProvidersOk() (*[]SubProviderMetadata, bool)`

GetSubProvidersOk returns a tuple with the SubProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubProviders

`func (o *ProviderInformation) SetSubProviders(v []SubProviderMetadata)`

SetSubProviders sets SubProviders field to given value.

### HasSubProviders

`func (o *ProviderInformation) HasSubProviders() bool`

HasSubProviders returns a boolean if a field has been set.

### SetSubProvidersNil

`func (o *ProviderInformation) SetSubProvidersNil(b bool)`

 SetSubProvidersNil sets the value for SubProviders to be an explicit nil

### UnsetSubProviders
`func (o *ProviderInformation) UnsetSubProviders()`

UnsetSubProviders ensures that no value is present for SubProviders, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


