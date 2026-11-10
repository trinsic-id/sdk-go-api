# RecommendProviderInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the provider | [readonly] 
**Name** | **string** | The friendly, human-readable name of the provider | [readonly] 
**LocalizedDisplayNames** | Pointer to **map[string]string** | Localized display names for the Provider, if any, keyed by lowercase BCP 47 language codes (eg \&quot;fr\&quot;).              If a particular language code is not present, fall back to &#x60;name&#x60;.              This should not be confused for the list of languages the Provider officially supports. | [optional] [readonly] 
**LogoUrl** | **string** | A URL pointing to the provider&#39;s logo | [readonly] 
**Subtext** | **string** | The Provider&#39;s subtext recommended to be shown next to the name.              This is flavor text, not a full, human-readable description of the provider. | [readonly] 
**Health** | **string** | The current health status of the provider | [readonly] 
**SubProviders** | Pointer to [**[]SubProviderMetadata**](SubProviderMetadata.md) | Metadata about the sub-providers which are available for this Provider.              For example, Italy&#39;s SPID is a Provider which aggregates access to multiple sub-providers. | [optional] 
**Regions** | **[]string** | The regions a provider is available in. | [readonly] 
**Countries** | **[]string** | A list of countries, in alpha-2 ISO 3166-1 format, that the provider is available in. | [readonly] 
**Subdivisions** | **[]string** | A list of subdivisions, in ISO 3166-2 format, that the provider is available in. | [readonly] 

## Methods

### Named construction

Use a keyed struct literal (`RecommendProviderInformation{FieldName: value}`), or call
`NewRecommendProviderInformationWithDefaults()` and set fields by name. Positional field
factories and unkeyed struct literals are not supported.

### NewRecommendProviderInformationWithDefaults

`func NewRecommendProviderInformationWithDefaults() *RecommendProviderInformation`

NewRecommendProviderInformationWithDefaults instantiates a new RecommendProviderInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RecommendProviderInformation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RecommendProviderInformation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RecommendProviderInformation) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *RecommendProviderInformation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RecommendProviderInformation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RecommendProviderInformation) SetName(v string)`

SetName sets Name field to given value.


### GetLocalizedDisplayNames

`func (o *RecommendProviderInformation) GetLocalizedDisplayNames() map[string]string`

GetLocalizedDisplayNames returns the LocalizedDisplayNames field if non-nil, zero value otherwise.

### GetLocalizedDisplayNamesOk

`func (o *RecommendProviderInformation) GetLocalizedDisplayNamesOk() (*map[string]string, bool)`

GetLocalizedDisplayNamesOk returns a tuple with the LocalizedDisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizedDisplayNames

`func (o *RecommendProviderInformation) SetLocalizedDisplayNames(v map[string]string)`

SetLocalizedDisplayNames sets LocalizedDisplayNames field to given value.

### HasLocalizedDisplayNames

`func (o *RecommendProviderInformation) HasLocalizedDisplayNames() bool`

HasLocalizedDisplayNames returns a boolean if a field has been set.

### SetLocalizedDisplayNamesNil

`func (o *RecommendProviderInformation) SetLocalizedDisplayNamesNil(b bool)`

 SetLocalizedDisplayNamesNil sets the value for LocalizedDisplayNames to be an explicit nil

### UnsetLocalizedDisplayNames
`func (o *RecommendProviderInformation) UnsetLocalizedDisplayNames()`

UnsetLocalizedDisplayNames ensures that no value is present for LocalizedDisplayNames, not even an explicit nil
### GetLogoUrl

`func (o *RecommendProviderInformation) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *RecommendProviderInformation) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *RecommendProviderInformation) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.


### GetSubtext

`func (o *RecommendProviderInformation) GetSubtext() string`

GetSubtext returns the Subtext field if non-nil, zero value otherwise.

### GetSubtextOk

`func (o *RecommendProviderInformation) GetSubtextOk() (*string, bool)`

GetSubtextOk returns a tuple with the Subtext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtext

`func (o *RecommendProviderInformation) SetSubtext(v string)`

SetSubtext sets Subtext field to given value.


### GetHealth

`func (o *RecommendProviderInformation) GetHealth() string`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *RecommendProviderInformation) GetHealthOk() (*string, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *RecommendProviderInformation) SetHealth(v string)`

SetHealth sets Health field to given value.


### GetSubProviders

`func (o *RecommendProviderInformation) GetSubProviders() []SubProviderMetadata`

GetSubProviders returns the SubProviders field if non-nil, zero value otherwise.

### GetSubProvidersOk

`func (o *RecommendProviderInformation) GetSubProvidersOk() (*[]SubProviderMetadata, bool)`

GetSubProvidersOk returns a tuple with the SubProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubProviders

`func (o *RecommendProviderInformation) SetSubProviders(v []SubProviderMetadata)`

SetSubProviders sets SubProviders field to given value.

### HasSubProviders

`func (o *RecommendProviderInformation) HasSubProviders() bool`

HasSubProviders returns a boolean if a field has been set.

### SetSubProvidersNil

`func (o *RecommendProviderInformation) SetSubProvidersNil(b bool)`

 SetSubProvidersNil sets the value for SubProviders to be an explicit nil

### UnsetSubProviders
`func (o *RecommendProviderInformation) UnsetSubProviders()`

UnsetSubProviders ensures that no value is present for SubProviders, not even an explicit nil
### GetRegions

`func (o *RecommendProviderInformation) GetRegions() []string`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *RecommendProviderInformation) GetRegionsOk() (*[]string, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *RecommendProviderInformation) SetRegions(v []string)`

SetRegions sets Regions field to given value.


### GetCountries

`func (o *RecommendProviderInformation) GetCountries() []string`

GetCountries returns the Countries field if non-nil, zero value otherwise.

### GetCountriesOk

`func (o *RecommendProviderInformation) GetCountriesOk() (*[]string, bool)`

GetCountriesOk returns a tuple with the Countries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountries

`func (o *RecommendProviderInformation) SetCountries(v []string)`

SetCountries sets Countries field to given value.


### GetSubdivisions

`func (o *RecommendProviderInformation) GetSubdivisions() []string`

GetSubdivisions returns the Subdivisions field if non-nil, zero value otherwise.

### GetSubdivisionsOk

`func (o *RecommendProviderInformation) GetSubdivisionsOk() (*[]string, bool)`

GetSubdivisionsOk returns a tuple with the Subdivisions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdivisions

`func (o *RecommendProviderInformation) SetSubdivisions(v []string)`

SetSubdivisions sets Subdivisions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


