# ProviderContract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the Provider for this contract. | 
**Name** | **string** | The Provider&#39;s Name as it appears in Trinsic&#39;s Dashboard and Widget | 
**Subtext** | **string** | The Provider&#39;s subtext recommended to be shown next to the name.              This is flavor text, not a full, human-readable description of the provider. | 
**Description** | **string** | The Provider&#39;s description as it appears in Trinsic&#39;s Widget.              This is flavor text, not a full, human-readable description of the provider. | 
**LogoUrl** | **string** | A URL pointing to the Provider&#39;s logo on Trinsic&#39;s CDN.              May be a PNG, JPG, or SVG image. | 
**Available** | **bool** | Whether the Provider is available for use in your App.              If &#x60;false&#x60;, you will need to contact Trinsic to enable this Provider for your App. | 
**Geography** | **[]string** | The geographies within the Regions the Provider is available. | 
**Regions** | **[]string** | The regions within which the Provider is available. | 
**LaunchMethod** | [**IntegrationLaunchMethod**](IntegrationLaunchMethod.md) | Relevant only to Advanced Provider Sessions.              The &#x60;LaunchMethod&#x60; which must be supported to launch the Provider Session in Advanced Provider Sessions. | 
**CollectionMethod** | [**ResultCollectionMethod**](ResultCollectionMethod.md) | Relevant only to Advanced Provider Sessions.              The &#x60;CollectionMethod&#x60; which must be supported to launch the Provider Session in Advanced Provider Sessions. | 
**ResultsMayBeDelayedAfterRedirect** | **bool** | If &#x60;true&#x60;, then the results for this Provider may not be available immediately after the user is redirected back to your application. In this case, the &#x60;GetSessionResults&#x60; API must be called until results are available.              This is an uncommon scenario, and only applies to Providers which cannot guarantee the availability of results immediately after the user is redirected back to your application. | 
**HasRefreshableContent** | **bool** | Relevant only to Advanced Provider Sessions.              Whether the Provider requires the &#x60;RefreshStepContent&#x60; capability.              For example, Samsung Wallet&#39;s deep links expire every 30 seconds, and must be refreshed periodically for a resilient user flow. | 
**RequiresInput** | **bool** | Relevant to Hosted Provider Sessions and Advanced Provider Sessions.              If &#x60;true&#x60;, this Provider requires provider-specific input on Session creation. If this input is not provided, Trinsic&#39;s Hosted UI will be invoked to collect the input from the user. | 
**HasTrinsicInterface** | **bool** | Whether there exists a Trinsic-hosted UI for this Provider.              This is &#x60;true&#x60; for any Provider which is not a simple, OIDC-like redirect flow. | 
**SupportsAdvancedProviderSessions** | **bool** | Whether this Provider can be fully whitelabeled/OEMed through the Advanced Provider Sessions API.              If &#x60;false&#x60;, the Provider may still be launched through Advanced Provider Sessions; however, it will necessarily require a Trinsic-hosted UI to function. | 
**AvailableFields** | Pointer to [**[]ContractField**](ContractField.md) | Information about the fields that this Provider will return in verification results. | [optional] 
**SubProviders** | Pointer to [**[]SubProviderMetadata**](SubProviderMetadata.md) | Metadata about the sub-providers which are available for this Provider.              For example, Italy&#39;s SPID is a Provider which aggregates access to multiple sub-providers. | [optional] 
**Health** | [**ProviderHealth**](ProviderHealth.md) | The health for an integration to be able to successfully perform a verification session. | 

## Methods

### NewProviderContract

`func NewProviderContract(id string, name string, subtext string, description string, logoUrl string, available bool, geography []string, regions []string, launchMethod IntegrationLaunchMethod, collectionMethod ResultCollectionMethod, resultsMayBeDelayedAfterRedirect bool, hasRefreshableContent bool, requiresInput bool, hasTrinsicInterface bool, supportsAdvancedProviderSessions bool, health ProviderHealth, ) *ProviderContract`

NewProviderContract instantiates a new ProviderContract object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderContractWithDefaults

`func NewProviderContractWithDefaults() *ProviderContract`

NewProviderContractWithDefaults instantiates a new ProviderContract object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProviderContract) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProviderContract) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProviderContract) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ProviderContract) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProviderContract) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProviderContract) SetName(v string)`

SetName sets Name field to given value.


### GetSubtext

`func (o *ProviderContract) GetSubtext() string`

GetSubtext returns the Subtext field if non-nil, zero value otherwise.

### GetSubtextOk

`func (o *ProviderContract) GetSubtextOk() (*string, bool)`

GetSubtextOk returns a tuple with the Subtext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtext

`func (o *ProviderContract) SetSubtext(v string)`

SetSubtext sets Subtext field to given value.


### GetDescription

`func (o *ProviderContract) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProviderContract) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProviderContract) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetLogoUrl

`func (o *ProviderContract) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *ProviderContract) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *ProviderContract) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.


### GetAvailable

`func (o *ProviderContract) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *ProviderContract) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *ProviderContract) SetAvailable(v bool)`

SetAvailable sets Available field to given value.


### GetGeography

`func (o *ProviderContract) GetGeography() []string`

GetGeography returns the Geography field if non-nil, zero value otherwise.

### GetGeographyOk

`func (o *ProviderContract) GetGeographyOk() (*[]string, bool)`

GetGeographyOk returns a tuple with the Geography field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeography

`func (o *ProviderContract) SetGeography(v []string)`

SetGeography sets Geography field to given value.


### GetRegions

`func (o *ProviderContract) GetRegions() []string`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *ProviderContract) GetRegionsOk() (*[]string, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *ProviderContract) SetRegions(v []string)`

SetRegions sets Regions field to given value.


### GetLaunchMethod

`func (o *ProviderContract) GetLaunchMethod() IntegrationLaunchMethod`

GetLaunchMethod returns the LaunchMethod field if non-nil, zero value otherwise.

### GetLaunchMethodOk

`func (o *ProviderContract) GetLaunchMethodOk() (*IntegrationLaunchMethod, bool)`

GetLaunchMethodOk returns a tuple with the LaunchMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchMethod

`func (o *ProviderContract) SetLaunchMethod(v IntegrationLaunchMethod)`

SetLaunchMethod sets LaunchMethod field to given value.


### GetCollectionMethod

`func (o *ProviderContract) GetCollectionMethod() ResultCollectionMethod`

GetCollectionMethod returns the CollectionMethod field if non-nil, zero value otherwise.

### GetCollectionMethodOk

`func (o *ProviderContract) GetCollectionMethodOk() (*ResultCollectionMethod, bool)`

GetCollectionMethodOk returns a tuple with the CollectionMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionMethod

`func (o *ProviderContract) SetCollectionMethod(v ResultCollectionMethod)`

SetCollectionMethod sets CollectionMethod field to given value.


### GetResultsMayBeDelayedAfterRedirect

`func (o *ProviderContract) GetResultsMayBeDelayedAfterRedirect() bool`

GetResultsMayBeDelayedAfterRedirect returns the ResultsMayBeDelayedAfterRedirect field if non-nil, zero value otherwise.

### GetResultsMayBeDelayedAfterRedirectOk

`func (o *ProviderContract) GetResultsMayBeDelayedAfterRedirectOk() (*bool, bool)`

GetResultsMayBeDelayedAfterRedirectOk returns a tuple with the ResultsMayBeDelayedAfterRedirect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsMayBeDelayedAfterRedirect

`func (o *ProviderContract) SetResultsMayBeDelayedAfterRedirect(v bool)`

SetResultsMayBeDelayedAfterRedirect sets ResultsMayBeDelayedAfterRedirect field to given value.


### GetHasRefreshableContent

`func (o *ProviderContract) GetHasRefreshableContent() bool`

GetHasRefreshableContent returns the HasRefreshableContent field if non-nil, zero value otherwise.

### GetHasRefreshableContentOk

`func (o *ProviderContract) GetHasRefreshableContentOk() (*bool, bool)`

GetHasRefreshableContentOk returns a tuple with the HasRefreshableContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasRefreshableContent

`func (o *ProviderContract) SetHasRefreshableContent(v bool)`

SetHasRefreshableContent sets HasRefreshableContent field to given value.


### GetRequiresInput

`func (o *ProviderContract) GetRequiresInput() bool`

GetRequiresInput returns the RequiresInput field if non-nil, zero value otherwise.

### GetRequiresInputOk

`func (o *ProviderContract) GetRequiresInputOk() (*bool, bool)`

GetRequiresInputOk returns a tuple with the RequiresInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresInput

`func (o *ProviderContract) SetRequiresInput(v bool)`

SetRequiresInput sets RequiresInput field to given value.


### GetHasTrinsicInterface

`func (o *ProviderContract) GetHasTrinsicInterface() bool`

GetHasTrinsicInterface returns the HasTrinsicInterface field if non-nil, zero value otherwise.

### GetHasTrinsicInterfaceOk

`func (o *ProviderContract) GetHasTrinsicInterfaceOk() (*bool, bool)`

GetHasTrinsicInterfaceOk returns a tuple with the HasTrinsicInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasTrinsicInterface

`func (o *ProviderContract) SetHasTrinsicInterface(v bool)`

SetHasTrinsicInterface sets HasTrinsicInterface field to given value.


### GetSupportsAdvancedProviderSessions

`func (o *ProviderContract) GetSupportsAdvancedProviderSessions() bool`

GetSupportsAdvancedProviderSessions returns the SupportsAdvancedProviderSessions field if non-nil, zero value otherwise.

### GetSupportsAdvancedProviderSessionsOk

`func (o *ProviderContract) GetSupportsAdvancedProviderSessionsOk() (*bool, bool)`

GetSupportsAdvancedProviderSessionsOk returns a tuple with the SupportsAdvancedProviderSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsAdvancedProviderSessions

`func (o *ProviderContract) SetSupportsAdvancedProviderSessions(v bool)`

SetSupportsAdvancedProviderSessions sets SupportsAdvancedProviderSessions field to given value.


### GetAvailableFields

`func (o *ProviderContract) GetAvailableFields() []ContractField`

GetAvailableFields returns the AvailableFields field if non-nil, zero value otherwise.

### GetAvailableFieldsOk

`func (o *ProviderContract) GetAvailableFieldsOk() (*[]ContractField, bool)`

GetAvailableFieldsOk returns a tuple with the AvailableFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableFields

`func (o *ProviderContract) SetAvailableFields(v []ContractField)`

SetAvailableFields sets AvailableFields field to given value.

### HasAvailableFields

`func (o *ProviderContract) HasAvailableFields() bool`

HasAvailableFields returns a boolean if a field has been set.

### SetAvailableFieldsNil

`func (o *ProviderContract) SetAvailableFieldsNil(b bool)`

 SetAvailableFieldsNil sets the value for AvailableFields to be an explicit nil

### UnsetAvailableFields
`func (o *ProviderContract) UnsetAvailableFields()`

UnsetAvailableFields ensures that no value is present for AvailableFields, not even an explicit nil
### GetSubProviders

`func (o *ProviderContract) GetSubProviders() []SubProviderMetadata`

GetSubProviders returns the SubProviders field if non-nil, zero value otherwise.

### GetSubProvidersOk

`func (o *ProviderContract) GetSubProvidersOk() (*[]SubProviderMetadata, bool)`

GetSubProvidersOk returns a tuple with the SubProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubProviders

`func (o *ProviderContract) SetSubProviders(v []SubProviderMetadata)`

SetSubProviders sets SubProviders field to given value.

### HasSubProviders

`func (o *ProviderContract) HasSubProviders() bool`

HasSubProviders returns a boolean if a field has been set.

### SetSubProvidersNil

`func (o *ProviderContract) SetSubProvidersNil(b bool)`

 SetSubProvidersNil sets the value for SubProviders to be an explicit nil

### UnsetSubProviders
`func (o *ProviderContract) UnsetSubProviders()`

UnsetSubProviders ensures that no value is present for SubProviders, not even an explicit nil
### GetHealth

`func (o *ProviderContract) GetHealth() ProviderHealth`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *ProviderContract) GetHealthOk() (*ProviderHealth, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *ProviderContract) SetHealth(v ProviderHealth)`

SetHealth sets Health field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


