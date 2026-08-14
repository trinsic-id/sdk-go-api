# ClearProviderOutputWatchlistHit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityType** | Pointer to **NullableString** | The watchlist entity type.              Known values: - person | [optional] 
**Details** | Pointer to [**NullableClearProviderOutputWatchlistHitDetails**](ClearProviderOutputWatchlistHitDetails.md) | More details about the watchlist hit entity. | [optional] 
**SourceLists** | Pointer to [**[]ClearProviderOutputWatchlistSource**](ClearProviderOutputWatchlistSource.md) | The watchlist source lists where CLEAR found the hit. | [optional] 
**HitTypes** | Pointer to **[]string** | The types of watchlist hits.              Known values: - sanction - pep - warning - adverse_media | [optional] 

## Methods

### NewClearProviderOutputWatchlistHit

`func NewClearProviderOutputWatchlistHit() *ClearProviderOutputWatchlistHit`

NewClearProviderOutputWatchlistHit instantiates a new ClearProviderOutputWatchlistHit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClearProviderOutputWatchlistHitWithDefaults

`func NewClearProviderOutputWatchlistHitWithDefaults() *ClearProviderOutputWatchlistHit`

NewClearProviderOutputWatchlistHitWithDefaults instantiates a new ClearProviderOutputWatchlistHit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityType

`func (o *ClearProviderOutputWatchlistHit) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *ClearProviderOutputWatchlistHit) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *ClearProviderOutputWatchlistHit) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *ClearProviderOutputWatchlistHit) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### SetEntityTypeNil

`func (o *ClearProviderOutputWatchlistHit) SetEntityTypeNil(b bool)`

 SetEntityTypeNil sets the value for EntityType to be an explicit nil

### UnsetEntityType
`func (o *ClearProviderOutputWatchlistHit) UnsetEntityType()`

UnsetEntityType ensures that no value is present for EntityType, not even an explicit nil
### GetDetails

`func (o *ClearProviderOutputWatchlistHit) GetDetails() ClearProviderOutputWatchlistHitDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ClearProviderOutputWatchlistHit) GetDetailsOk() (*ClearProviderOutputWatchlistHitDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ClearProviderOutputWatchlistHit) SetDetails(v ClearProviderOutputWatchlistHitDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *ClearProviderOutputWatchlistHit) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### SetDetailsNil

`func (o *ClearProviderOutputWatchlistHit) SetDetailsNil(b bool)`

 SetDetailsNil sets the value for Details to be an explicit nil

### UnsetDetails
`func (o *ClearProviderOutputWatchlistHit) UnsetDetails()`

UnsetDetails ensures that no value is present for Details, not even an explicit nil
### GetSourceLists

`func (o *ClearProviderOutputWatchlistHit) GetSourceLists() []ClearProviderOutputWatchlistSource`

GetSourceLists returns the SourceLists field if non-nil, zero value otherwise.

### GetSourceListsOk

`func (o *ClearProviderOutputWatchlistHit) GetSourceListsOk() (*[]ClearProviderOutputWatchlistSource, bool)`

GetSourceListsOk returns a tuple with the SourceLists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceLists

`func (o *ClearProviderOutputWatchlistHit) SetSourceLists(v []ClearProviderOutputWatchlistSource)`

SetSourceLists sets SourceLists field to given value.

### HasSourceLists

`func (o *ClearProviderOutputWatchlistHit) HasSourceLists() bool`

HasSourceLists returns a boolean if a field has been set.

### SetSourceListsNil

`func (o *ClearProviderOutputWatchlistHit) SetSourceListsNil(b bool)`

 SetSourceListsNil sets the value for SourceLists to be an explicit nil

### UnsetSourceLists
`func (o *ClearProviderOutputWatchlistHit) UnsetSourceLists()`

UnsetSourceLists ensures that no value is present for SourceLists, not even an explicit nil
### GetHitTypes

`func (o *ClearProviderOutputWatchlistHit) GetHitTypes() []string`

GetHitTypes returns the HitTypes field if non-nil, zero value otherwise.

### GetHitTypesOk

`func (o *ClearProviderOutputWatchlistHit) GetHitTypesOk() (*[]string, bool)`

GetHitTypesOk returns a tuple with the HitTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHitTypes

`func (o *ClearProviderOutputWatchlistHit) SetHitTypes(v []string)`

SetHitTypes sets HitTypes field to given value.

### HasHitTypes

`func (o *ClearProviderOutputWatchlistHit) HasHitTypes() bool`

HasHitTypes returns a boolean if a field has been set.

### SetHitTypesNil

`func (o *ClearProviderOutputWatchlistHit) SetHitTypesNil(b bool)`

 SetHitTypesNil sets the value for HitTypes to be an explicit nil

### UnsetHitTypes
`func (o *ClearProviderOutputWatchlistHit) UnsetHitTypes()`

UnsetHitTypes ensures that no value is present for HitTypes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


