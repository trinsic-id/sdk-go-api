# PeruDniAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Region** | Pointer to **NullableString** | Region in Peru (also known as department). This is the first level subdivision in the country.              Format: - All uppercase. | [optional] 
**Province** | Pointer to **NullableString** | Province within the region. This is the second level subdivision in the country.              Format: - All uppercase. | [optional] 
**District** | Pointer to **NullableString** | District within the province. This is the third level subdivision in the country.              Format: - All uppercase. | [optional] 

## Methods

### NewPeruDniAddress

`func NewPeruDniAddress() *PeruDniAddress`

NewPeruDniAddress instantiates a new PeruDniAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPeruDniAddressWithDefaults

`func NewPeruDniAddressWithDefaults() *PeruDniAddress`

NewPeruDniAddressWithDefaults instantiates a new PeruDniAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegion

`func (o *PeruDniAddress) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *PeruDniAddress) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *PeruDniAddress) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *PeruDniAddress) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *PeruDniAddress) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *PeruDniAddress) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetProvince

`func (o *PeruDniAddress) GetProvince() string`

GetProvince returns the Province field if non-nil, zero value otherwise.

### GetProvinceOk

`func (o *PeruDniAddress) GetProvinceOk() (*string, bool)`

GetProvinceOk returns a tuple with the Province field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvince

`func (o *PeruDniAddress) SetProvince(v string)`

SetProvince sets Province field to given value.

### HasProvince

`func (o *PeruDniAddress) HasProvince() bool`

HasProvince returns a boolean if a field has been set.

### SetProvinceNil

`func (o *PeruDniAddress) SetProvinceNil(b bool)`

 SetProvinceNil sets the value for Province to be an explicit nil

### UnsetProvince
`func (o *PeruDniAddress) UnsetProvince()`

UnsetProvince ensures that no value is present for Province, not even an explicit nil
### GetDistrict

`func (o *PeruDniAddress) GetDistrict() string`

GetDistrict returns the District field if non-nil, zero value otherwise.

### GetDistrictOk

`func (o *PeruDniAddress) GetDistrictOk() (*string, bool)`

GetDistrictOk returns a tuple with the District field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistrict

`func (o *PeruDniAddress) SetDistrict(v string)`

SetDistrict sets District field to given value.

### HasDistrict

`func (o *PeruDniAddress) HasDistrict() bool`

HasDistrict returns a boolean if a field has been set.

### SetDistrictNil

`func (o *PeruDniAddress) SetDistrictNil(b bool)`

 SetDistrictNil sets the value for District to be an explicit nil

### UnsetDistrict
`func (o *PeruDniAddress) UnsetDistrict()`

UnsetDistrict ensures that no value is present for District, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


