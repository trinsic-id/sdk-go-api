# ProviderInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the provider | 
**Name** | **string** | The friendly, human-readable name of the provider | 
**LogoUrl** | **string** | A URL pointing to the provider&#39;s logo | 

## Methods

### NewProviderInfo

`func NewProviderInfo(id string, name string, logoUrl string, ) *ProviderInfo`

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


