# ClearProviderOutputWatchlistHits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hits** | Pointer to [**[]ClearProviderOutputWatchlistHit**](ClearProviderOutputWatchlistHit.md) | The watchlist hits. | [optional] 
**UpdatedAt** | Pointer to **NullableTime** | The last time CLEAR fetched watchlist hit data from its data sources, as a UTC timestamp. | [optional] 

## Methods

### NewClearProviderOutputWatchlistHits

`func NewClearProviderOutputWatchlistHits() *ClearProviderOutputWatchlistHits`

NewClearProviderOutputWatchlistHits instantiates a new ClearProviderOutputWatchlistHits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClearProviderOutputWatchlistHitsWithDefaults

`func NewClearProviderOutputWatchlistHitsWithDefaults() *ClearProviderOutputWatchlistHits`

NewClearProviderOutputWatchlistHitsWithDefaults instantiates a new ClearProviderOutputWatchlistHits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHits

`func (o *ClearProviderOutputWatchlistHits) GetHits() []ClearProviderOutputWatchlistHit`

GetHits returns the Hits field if non-nil, zero value otherwise.

### GetHitsOk

`func (o *ClearProviderOutputWatchlistHits) GetHitsOk() (*[]ClearProviderOutputWatchlistHit, bool)`

GetHitsOk returns a tuple with the Hits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHits

`func (o *ClearProviderOutputWatchlistHits) SetHits(v []ClearProviderOutputWatchlistHit)`

SetHits sets Hits field to given value.

### HasHits

`func (o *ClearProviderOutputWatchlistHits) HasHits() bool`

HasHits returns a boolean if a field has been set.

### SetHitsNil

`func (o *ClearProviderOutputWatchlistHits) SetHitsNil(b bool)`

 SetHitsNil sets the value for Hits to be an explicit nil

### UnsetHits
`func (o *ClearProviderOutputWatchlistHits) UnsetHits()`

UnsetHits ensures that no value is present for Hits, not even an explicit nil
### GetUpdatedAt

`func (o *ClearProviderOutputWatchlistHits) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ClearProviderOutputWatchlistHits) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ClearProviderOutputWatchlistHits) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ClearProviderOutputWatchlistHits) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *ClearProviderOutputWatchlistHits) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *ClearProviderOutputWatchlistHits) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


