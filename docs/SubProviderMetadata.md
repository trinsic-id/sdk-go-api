# SubProviderMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the sub-provider.              This cannot be used as a standalone Provider ID when creating a Session. It must be passed in via the Provider-specific input. | 
**Name** | **string** | The name of the sub-provider | 
**Subtext** | **string** | The Provider&#39;s subtext recommended to be shown next to the name.              This is flavor text, not a full, human-readable description of the provider. | 
**Description** | **string** | The Provider&#39;s subtext recommended to be shown next to the name.              This is flavor text, not a full, human-readable description of the provider. | 
**LogoUrl** | **string** | A URL pointing to the logo on Trinsic&#39;s CDN.              May be a PNG, JPG, or SVG image. | 

## Methods

### NewSubProviderMetadata

`func NewSubProviderMetadata(id string, name string, subtext string, description string, logoUrl string, ) *SubProviderMetadata`

NewSubProviderMetadata instantiates a new SubProviderMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubProviderMetadataWithDefaults

`func NewSubProviderMetadataWithDefaults() *SubProviderMetadata`

NewSubProviderMetadataWithDefaults instantiates a new SubProviderMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SubProviderMetadata) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubProviderMetadata) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubProviderMetadata) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *SubProviderMetadata) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SubProviderMetadata) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SubProviderMetadata) SetName(v string)`

SetName sets Name field to given value.


### GetSubtext

`func (o *SubProviderMetadata) GetSubtext() string`

GetSubtext returns the Subtext field if non-nil, zero value otherwise.

### GetSubtextOk

`func (o *SubProviderMetadata) GetSubtextOk() (*string, bool)`

GetSubtextOk returns a tuple with the Subtext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtext

`func (o *SubProviderMetadata) SetSubtext(v string)`

SetSubtext sets Subtext field to given value.


### GetDescription

`func (o *SubProviderMetadata) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SubProviderMetadata) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SubProviderMetadata) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetLogoUrl

`func (o *SubProviderMetadata) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *SubProviderMetadata) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *SubProviderMetadata) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


