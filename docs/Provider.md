# Provider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier of the provider | 
**Name** | **string** | The display name of the provider | 
**LogoUrl** | **string** | The URL of the provider&#39;s logo | 
**Subtext** | **string** | The Provider&#39;s subtext recommended to be shown next to the name.              This is flavor text, not a full, human-readable description of the provider. | 
**Description** | **string** | A description of the provider&#39;s capabilities | [readonly] 
**Geography** | **[]string** | Geographic regions where this provider operates | 
**Regions** | **[]string** | Specific regions supported by this provider | 
**Countries** | **[]string** | The countries where this Provider is available (as alpha-2 ISO codes). | 
**Subdivisions** | **[]string** | The subdivisions (states, provinces, etc.) where this Provider is available (as alpha-2 ISO codes). | 
**Licensed** | **bool** | Whether this provider is licensed for use by your organization | 
**LaunchMethod** | [**IntegrationLaunchMethod**](IntegrationLaunchMethod.md) | Relevant only to Direct Provider Sessions.              The &#x60;LaunchMethod&#x60; which must be supported to launch the Provider Session in Direct Provider Sessions. | 
**CollectionMethod** | [**ResultCollectionMethod**](ResultCollectionMethod.md) | Relevant only to Direct Provider Sessions.              The &#x60;CollectionMethod&#x60; which must be supported to launch the Provider Session in Direct Provider Sessions. | 
**ResultsMayBeDelayedAfterRedirect** | **bool** | If &#x60;true&#x60;, then the results for this Provider may not be available immediately after the user is redirected back to your application. In this case, the &#x60;GetSessionResults&#x60; API must be called until results are available.              This is an uncommon scenario, and only applies to Providers which cannot guarantee the availability of results immediately after the user is redirected back to your application. | 
**HasRefreshableContent** | **bool** | Relevant only to Direct Provider Sessions.              Whether the Provider requires the &#x60;RefreshStepContent&#x60; capability.              For example, Samsung Wallet&#39;s deep links expire every 30 seconds, and must be refreshed periodically for a resilient user flow. | 
**RequiresInput** | **bool** | Relevant to Hosted Provider Sessions and Direct Provider Sessions.              If &#x60;true&#x60;, this Provider requires provider-specific input on Session creation. If this input is not provided, Trinsic&#39;s Hosted UI will be invoked to collect the input from the user. | 
**HasTrinsicInterface** | **bool** | Whether there exists a Trinsic-hosted UI for this Provider.              This is &#x60;true&#x60; for any Provider which is not a simple, OIDC-like redirect flow. | 
**SupportsDirectProviderSessions** | **bool** | Whether this Provider can be fully whitelabeled/OEMed through the Direct Provider Sessions API.              If &#x60;false&#x60;, the Provider may still be launched through Direct Provider Sessions; however, it will necessarily require a Trinsic-hosted UI to function. | 
**AvailableAttributes** | Pointer to [**[]ContractAttribute**](ContractAttribute.md) | Information about the user attributes that this Provider will return in verification results. | [optional] 
**AvailableAttachments** | Pointer to [**[]ContractAttachment**](ContractAttachment.md) | Information about the attachments that this Provider will return in verification results. | [optional] 
**SubProviders** | Pointer to [**[]SubProviderMetadata**](SubProviderMetadata.md) | Metadata about the sub-providers which are available for this Provider in the current Environment.              For example, Italy&#39;s SPID is a Provider which aggregates access to multiple sub-providers. | [optional] 
**LiveHealth** | [**ProviderHealth**](ProviderHealth.md) | The health for a provider in the live environment | 
**TestHealth** | [**ProviderHealth**](ProviderHealth.md) | The health for a provider in the test environment | 

## Methods

### NewProvider

`func NewProvider(id string, name string, logoUrl string, subtext string, description string, geography []string, regions []string, countries []string, subdivisions []string, licensed bool, launchMethod IntegrationLaunchMethod, collectionMethod ResultCollectionMethod, resultsMayBeDelayedAfterRedirect bool, hasRefreshableContent bool, requiresInput bool, hasTrinsicInterface bool, supportsDirectProviderSessions bool, liveHealth ProviderHealth, testHealth ProviderHealth, ) *Provider`

NewProvider instantiates a new Provider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderWithDefaults

`func NewProviderWithDefaults() *Provider`

NewProviderWithDefaults instantiates a new Provider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Provider) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Provider) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Provider) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Provider) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Provider) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Provider) SetName(v string)`

SetName sets Name field to given value.


### GetLogoUrl

`func (o *Provider) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *Provider) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *Provider) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.


### GetSubtext

`func (o *Provider) GetSubtext() string`

GetSubtext returns the Subtext field if non-nil, zero value otherwise.

### GetSubtextOk

`func (o *Provider) GetSubtextOk() (*string, bool)`

GetSubtextOk returns a tuple with the Subtext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtext

`func (o *Provider) SetSubtext(v string)`

SetSubtext sets Subtext field to given value.


### GetDescription

`func (o *Provider) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Provider) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Provider) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetGeography

`func (o *Provider) GetGeography() []string`

GetGeography returns the Geography field if non-nil, zero value otherwise.

### GetGeographyOk

`func (o *Provider) GetGeographyOk() (*[]string, bool)`

GetGeographyOk returns a tuple with the Geography field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeography

`func (o *Provider) SetGeography(v []string)`

SetGeography sets Geography field to given value.


### GetRegions

`func (o *Provider) GetRegions() []string`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *Provider) GetRegionsOk() (*[]string, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *Provider) SetRegions(v []string)`

SetRegions sets Regions field to given value.


### GetCountries

`func (o *Provider) GetCountries() []string`

GetCountries returns the Countries field if non-nil, zero value otherwise.

### GetCountriesOk

`func (o *Provider) GetCountriesOk() (*[]string, bool)`

GetCountriesOk returns a tuple with the Countries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountries

`func (o *Provider) SetCountries(v []string)`

SetCountries sets Countries field to given value.


### GetSubdivisions

`func (o *Provider) GetSubdivisions() []string`

GetSubdivisions returns the Subdivisions field if non-nil, zero value otherwise.

### GetSubdivisionsOk

`func (o *Provider) GetSubdivisionsOk() (*[]string, bool)`

GetSubdivisionsOk returns a tuple with the Subdivisions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdivisions

`func (o *Provider) SetSubdivisions(v []string)`

SetSubdivisions sets Subdivisions field to given value.


### GetLicensed

`func (o *Provider) GetLicensed() bool`

GetLicensed returns the Licensed field if non-nil, zero value otherwise.

### GetLicensedOk

`func (o *Provider) GetLicensedOk() (*bool, bool)`

GetLicensedOk returns a tuple with the Licensed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensed

`func (o *Provider) SetLicensed(v bool)`

SetLicensed sets Licensed field to given value.


### GetLaunchMethod

`func (o *Provider) GetLaunchMethod() IntegrationLaunchMethod`

GetLaunchMethod returns the LaunchMethod field if non-nil, zero value otherwise.

### GetLaunchMethodOk

`func (o *Provider) GetLaunchMethodOk() (*IntegrationLaunchMethod, bool)`

GetLaunchMethodOk returns a tuple with the LaunchMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchMethod

`func (o *Provider) SetLaunchMethod(v IntegrationLaunchMethod)`

SetLaunchMethod sets LaunchMethod field to given value.


### GetCollectionMethod

`func (o *Provider) GetCollectionMethod() ResultCollectionMethod`

GetCollectionMethod returns the CollectionMethod field if non-nil, zero value otherwise.

### GetCollectionMethodOk

`func (o *Provider) GetCollectionMethodOk() (*ResultCollectionMethod, bool)`

GetCollectionMethodOk returns a tuple with the CollectionMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionMethod

`func (o *Provider) SetCollectionMethod(v ResultCollectionMethod)`

SetCollectionMethod sets CollectionMethod field to given value.


### GetResultsMayBeDelayedAfterRedirect

`func (o *Provider) GetResultsMayBeDelayedAfterRedirect() bool`

GetResultsMayBeDelayedAfterRedirect returns the ResultsMayBeDelayedAfterRedirect field if non-nil, zero value otherwise.

### GetResultsMayBeDelayedAfterRedirectOk

`func (o *Provider) GetResultsMayBeDelayedAfterRedirectOk() (*bool, bool)`

GetResultsMayBeDelayedAfterRedirectOk returns a tuple with the ResultsMayBeDelayedAfterRedirect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsMayBeDelayedAfterRedirect

`func (o *Provider) SetResultsMayBeDelayedAfterRedirect(v bool)`

SetResultsMayBeDelayedAfterRedirect sets ResultsMayBeDelayedAfterRedirect field to given value.


### GetHasRefreshableContent

`func (o *Provider) GetHasRefreshableContent() bool`

GetHasRefreshableContent returns the HasRefreshableContent field if non-nil, zero value otherwise.

### GetHasRefreshableContentOk

`func (o *Provider) GetHasRefreshableContentOk() (*bool, bool)`

GetHasRefreshableContentOk returns a tuple with the HasRefreshableContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasRefreshableContent

`func (o *Provider) SetHasRefreshableContent(v bool)`

SetHasRefreshableContent sets HasRefreshableContent field to given value.


### GetRequiresInput

`func (o *Provider) GetRequiresInput() bool`

GetRequiresInput returns the RequiresInput field if non-nil, zero value otherwise.

### GetRequiresInputOk

`func (o *Provider) GetRequiresInputOk() (*bool, bool)`

GetRequiresInputOk returns a tuple with the RequiresInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresInput

`func (o *Provider) SetRequiresInput(v bool)`

SetRequiresInput sets RequiresInput field to given value.


### GetHasTrinsicInterface

`func (o *Provider) GetHasTrinsicInterface() bool`

GetHasTrinsicInterface returns the HasTrinsicInterface field if non-nil, zero value otherwise.

### GetHasTrinsicInterfaceOk

`func (o *Provider) GetHasTrinsicInterfaceOk() (*bool, bool)`

GetHasTrinsicInterfaceOk returns a tuple with the HasTrinsicInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasTrinsicInterface

`func (o *Provider) SetHasTrinsicInterface(v bool)`

SetHasTrinsicInterface sets HasTrinsicInterface field to given value.


### GetSupportsDirectProviderSessions

`func (o *Provider) GetSupportsDirectProviderSessions() bool`

GetSupportsDirectProviderSessions returns the SupportsDirectProviderSessions field if non-nil, zero value otherwise.

### GetSupportsDirectProviderSessionsOk

`func (o *Provider) GetSupportsDirectProviderSessionsOk() (*bool, bool)`

GetSupportsDirectProviderSessionsOk returns a tuple with the SupportsDirectProviderSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsDirectProviderSessions

`func (o *Provider) SetSupportsDirectProviderSessions(v bool)`

SetSupportsDirectProviderSessions sets SupportsDirectProviderSessions field to given value.


### GetAvailableAttributes

`func (o *Provider) GetAvailableAttributes() []ContractAttribute`

GetAvailableAttributes returns the AvailableAttributes field if non-nil, zero value otherwise.

### GetAvailableAttributesOk

`func (o *Provider) GetAvailableAttributesOk() (*[]ContractAttribute, bool)`

GetAvailableAttributesOk returns a tuple with the AvailableAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableAttributes

`func (o *Provider) SetAvailableAttributes(v []ContractAttribute)`

SetAvailableAttributes sets AvailableAttributes field to given value.

### HasAvailableAttributes

`func (o *Provider) HasAvailableAttributes() bool`

HasAvailableAttributes returns a boolean if a field has been set.

### SetAvailableAttributesNil

`func (o *Provider) SetAvailableAttributesNil(b bool)`

 SetAvailableAttributesNil sets the value for AvailableAttributes to be an explicit nil

### UnsetAvailableAttributes
`func (o *Provider) UnsetAvailableAttributes()`

UnsetAvailableAttributes ensures that no value is present for AvailableAttributes, not even an explicit nil
### GetAvailableAttachments

`func (o *Provider) GetAvailableAttachments() []ContractAttachment`

GetAvailableAttachments returns the AvailableAttachments field if non-nil, zero value otherwise.

### GetAvailableAttachmentsOk

`func (o *Provider) GetAvailableAttachmentsOk() (*[]ContractAttachment, bool)`

GetAvailableAttachmentsOk returns a tuple with the AvailableAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableAttachments

`func (o *Provider) SetAvailableAttachments(v []ContractAttachment)`

SetAvailableAttachments sets AvailableAttachments field to given value.

### HasAvailableAttachments

`func (o *Provider) HasAvailableAttachments() bool`

HasAvailableAttachments returns a boolean if a field has been set.

### SetAvailableAttachmentsNil

`func (o *Provider) SetAvailableAttachmentsNil(b bool)`

 SetAvailableAttachmentsNil sets the value for AvailableAttachments to be an explicit nil

### UnsetAvailableAttachments
`func (o *Provider) UnsetAvailableAttachments()`

UnsetAvailableAttachments ensures that no value is present for AvailableAttachments, not even an explicit nil
### GetSubProviders

`func (o *Provider) GetSubProviders() []SubProviderMetadata`

GetSubProviders returns the SubProviders field if non-nil, zero value otherwise.

### GetSubProvidersOk

`func (o *Provider) GetSubProvidersOk() (*[]SubProviderMetadata, bool)`

GetSubProvidersOk returns a tuple with the SubProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubProviders

`func (o *Provider) SetSubProviders(v []SubProviderMetadata)`

SetSubProviders sets SubProviders field to given value.

### HasSubProviders

`func (o *Provider) HasSubProviders() bool`

HasSubProviders returns a boolean if a field has been set.

### SetSubProvidersNil

`func (o *Provider) SetSubProvidersNil(b bool)`

 SetSubProvidersNil sets the value for SubProviders to be an explicit nil

### UnsetSubProviders
`func (o *Provider) UnsetSubProviders()`

UnsetSubProviders ensures that no value is present for SubProviders, not even an explicit nil
### GetLiveHealth

`func (o *Provider) GetLiveHealth() ProviderHealth`

GetLiveHealth returns the LiveHealth field if non-nil, zero value otherwise.

### GetLiveHealthOk

`func (o *Provider) GetLiveHealthOk() (*ProviderHealth, bool)`

GetLiveHealthOk returns a tuple with the LiveHealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveHealth

`func (o *Provider) SetLiveHealth(v ProviderHealth)`

SetLiveHealth sets LiveHealth field to given value.


### GetTestHealth

`func (o *Provider) GetTestHealth() ProviderHealth`

GetTestHealth returns the TestHealth field if non-nil, zero value otherwise.

### GetTestHealthOk

`func (o *Provider) GetTestHealthOk() (*ProviderHealth, bool)`

GetTestHealthOk returns a tuple with the TestHealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestHealth

`func (o *Provider) SetTestHealth(v ProviderHealth)`

SetTestHealth sets TestHealth field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


