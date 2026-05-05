# FaydaProviderAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Region** | Pointer to **NullableString** | The region, which is the primary administrative division of the address.              This attribute is only available if registered directly. | [optional] 
**Zone** | Pointer to **NullableString** | The zone, which is the administrative area within the region.              This attribute is only available if registered directly. | [optional] 
**Woreda** | Pointer to **NullableString** | The woreda, which is the local district within the zone.              This attribute is only available if registered directly. | [optional] 

## Methods

### NewFaydaProviderAddress

`func NewFaydaProviderAddress() *FaydaProviderAddress`

NewFaydaProviderAddress instantiates a new FaydaProviderAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFaydaProviderAddressWithDefaults

`func NewFaydaProviderAddressWithDefaults() *FaydaProviderAddress`

NewFaydaProviderAddressWithDefaults instantiates a new FaydaProviderAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegion

`func (o *FaydaProviderAddress) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *FaydaProviderAddress) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *FaydaProviderAddress) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *FaydaProviderAddress) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *FaydaProviderAddress) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *FaydaProviderAddress) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetZone

`func (o *FaydaProviderAddress) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *FaydaProviderAddress) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *FaydaProviderAddress) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *FaydaProviderAddress) HasZone() bool`

HasZone returns a boolean if a field has been set.

### SetZoneNil

`func (o *FaydaProviderAddress) SetZoneNil(b bool)`

 SetZoneNil sets the value for Zone to be an explicit nil

### UnsetZone
`func (o *FaydaProviderAddress) UnsetZone()`

UnsetZone ensures that no value is present for Zone, not even an explicit nil
### GetWoreda

`func (o *FaydaProviderAddress) GetWoreda() string`

GetWoreda returns the Woreda field if non-nil, zero value otherwise.

### GetWoredaOk

`func (o *FaydaProviderAddress) GetWoredaOk() (*string, bool)`

GetWoredaOk returns a tuple with the Woreda field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWoreda

`func (o *FaydaProviderAddress) SetWoreda(v string)`

SetWoreda sets Woreda field to given value.

### HasWoreda

`func (o *FaydaProviderAddress) HasWoreda() bool`

HasWoreda returns a boolean if a field has been set.

### SetWoredaNil

`func (o *FaydaProviderAddress) SetWoredaNil(b bool)`

 SetWoredaNil sets the value for Woreda to be an explicit nil

### UnsetWoreda
`func (o *FaydaProviderAddress) UnsetWoreda()`

UnsetWoreda ensures that no value is present for Woreda, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


