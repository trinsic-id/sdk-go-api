# VerificationProfileResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the verification profile | 
**Alias** | **string** | An alias of the verification profile shown to developers and administrators. | 
**BrandName** | **string** | The brand name of the verification profile shown to end-users. | 
**LogoUrl** | **string** | The URL of the verification profile&#39;s logo. | 
**PrimaryColor** | **string** | The primary color of the verification profile. | 
**EnabledProviders** | **[]string** | The providers that are currently enabled for the verification profile. | 

## Methods

### NewVerificationProfileResponse

`func NewVerificationProfileResponse(id string, alias string, brandName string, logoUrl string, primaryColor string, enabledProviders []string, ) *VerificationProfileResponse`

NewVerificationProfileResponse instantiates a new VerificationProfileResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerificationProfileResponseWithDefaults

`func NewVerificationProfileResponseWithDefaults() *VerificationProfileResponse`

NewVerificationProfileResponseWithDefaults instantiates a new VerificationProfileResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VerificationProfileResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VerificationProfileResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VerificationProfileResponse) SetId(v string)`

SetId sets Id field to given value.


### GetAlias

`func (o *VerificationProfileResponse) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *VerificationProfileResponse) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *VerificationProfileResponse) SetAlias(v string)`

SetAlias sets Alias field to given value.


### GetBrandName

`func (o *VerificationProfileResponse) GetBrandName() string`

GetBrandName returns the BrandName field if non-nil, zero value otherwise.

### GetBrandNameOk

`func (o *VerificationProfileResponse) GetBrandNameOk() (*string, bool)`

GetBrandNameOk returns a tuple with the BrandName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandName

`func (o *VerificationProfileResponse) SetBrandName(v string)`

SetBrandName sets BrandName field to given value.


### GetLogoUrl

`func (o *VerificationProfileResponse) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *VerificationProfileResponse) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *VerificationProfileResponse) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.


### GetPrimaryColor

`func (o *VerificationProfileResponse) GetPrimaryColor() string`

GetPrimaryColor returns the PrimaryColor field if non-nil, zero value otherwise.

### GetPrimaryColorOk

`func (o *VerificationProfileResponse) GetPrimaryColorOk() (*string, bool)`

GetPrimaryColorOk returns a tuple with the PrimaryColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryColor

`func (o *VerificationProfileResponse) SetPrimaryColor(v string)`

SetPrimaryColor sets PrimaryColor field to given value.


### GetEnabledProviders

`func (o *VerificationProfileResponse) GetEnabledProviders() []string`

GetEnabledProviders returns the EnabledProviders field if non-nil, zero value otherwise.

### GetEnabledProvidersOk

`func (o *VerificationProfileResponse) GetEnabledProvidersOk() (*[]string, bool)`

GetEnabledProvidersOk returns a tuple with the EnabledProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledProviders

`func (o *VerificationProfileResponse) SetEnabledProviders(v []string)`

SetEnabledProviders sets EnabledProviders field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


