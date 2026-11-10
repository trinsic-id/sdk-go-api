# SaudiArabiaNafathAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuildingNumber** | Pointer to **NullableString** | Building number. | [optional] 
**Street** | Pointer to **NullableString** | Street name in the session locale.              Arabic when the session locale is ar, and English when it is en. streetSecondary is the other language. | [optional] 
**District** | Pointer to **NullableString** | District or neighborhood in the session locale.              Arabic when the session locale is ar, and English when it is en. districtSecondary is the other language. | [optional] 
**City** | Pointer to **NullableString** | City in the session locale.              Arabic when the session locale is ar, and English when it is en. citySecondary is the other language. | [optional] 
**PostalCode** | Pointer to **NullableString** | Postal code. | [optional] 
**AdditionalNumber** | Pointer to **NullableString** | Additional number, used alongside the district to locate buildings within the Saudi addressing system. | [optional] 
**UnitNumber** | Pointer to **NullableString** | Apartment or unit number. | [optional] 
**LocationCoordinates** | Pointer to **NullableString** | Geographic coordinates of the registered address.              Format: Two space-separated decimal numbers, longitude then latitude. | [optional] 
**IsPrimaryAddress** | Pointer to **NullableBool** | Whether this address is the individual&#39;s primary registered address. | [optional] 
**RegionName** | Pointer to **NullableString** | Region name in the session locale.              Arabic when the session locale is ar, and English when it is en. regionNameSecondary is the other language. | [optional] 
**RegionNameSecondary** | Pointer to **NullableString** | Region name in the other language from regionName.              English when the session locale is ar, and Arabic when it is en. | [optional] 
**CitySecondary** | Pointer to **NullableString** | City name in the other language from city.              English when the session locale is ar, and Arabic when it is en. | [optional] 
**StreetSecondary** | Pointer to **NullableString** | Street name in the other language from street.              English when the session locale is ar, and Arabic when it is en. | [optional] 
**DistrictSecondary** | Pointer to **NullableString** | District name in the other language from district.              English when the session locale is ar, and Arabic when it is en. | [optional] 
**CityId** | Pointer to **NullableString** | City identifier. | [optional] 
**RegionId** | Pointer to **NullableString** | Region identifier. | [optional] 
**DistrictId** | Pointer to **NullableString** | District identifier. | [optional] 

## Methods

### Named construction

Use a keyed struct literal (`SaudiArabiaNafathAddress{FieldName: value}`), or call
`NewSaudiArabiaNafathAddressWithDefaults()` and set fields by name. Positional field
factories and unkeyed struct literals are not supported.

### NewSaudiArabiaNafathAddressWithDefaults

`func NewSaudiArabiaNafathAddressWithDefaults() *SaudiArabiaNafathAddress`

NewSaudiArabiaNafathAddressWithDefaults instantiates a new SaudiArabiaNafathAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuildingNumber

`func (o *SaudiArabiaNafathAddress) GetBuildingNumber() string`

GetBuildingNumber returns the BuildingNumber field if non-nil, zero value otherwise.

### GetBuildingNumberOk

`func (o *SaudiArabiaNafathAddress) GetBuildingNumberOk() (*string, bool)`

GetBuildingNumberOk returns a tuple with the BuildingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildingNumber

`func (o *SaudiArabiaNafathAddress) SetBuildingNumber(v string)`

SetBuildingNumber sets BuildingNumber field to given value.

### HasBuildingNumber

`func (o *SaudiArabiaNafathAddress) HasBuildingNumber() bool`

HasBuildingNumber returns a boolean if a field has been set.

### SetBuildingNumberNil

`func (o *SaudiArabiaNafathAddress) SetBuildingNumberNil(b bool)`

 SetBuildingNumberNil sets the value for BuildingNumber to be an explicit nil

### UnsetBuildingNumber
`func (o *SaudiArabiaNafathAddress) UnsetBuildingNumber()`

UnsetBuildingNumber ensures that no value is present for BuildingNumber, not even an explicit nil
### GetStreet

`func (o *SaudiArabiaNafathAddress) GetStreet() string`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *SaudiArabiaNafathAddress) GetStreetOk() (*string, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet

`func (o *SaudiArabiaNafathAddress) SetStreet(v string)`

SetStreet sets Street field to given value.

### HasStreet

`func (o *SaudiArabiaNafathAddress) HasStreet() bool`

HasStreet returns a boolean if a field has been set.

### SetStreetNil

`func (o *SaudiArabiaNafathAddress) SetStreetNil(b bool)`

 SetStreetNil sets the value for Street to be an explicit nil

### UnsetStreet
`func (o *SaudiArabiaNafathAddress) UnsetStreet()`

UnsetStreet ensures that no value is present for Street, not even an explicit nil
### GetDistrict

`func (o *SaudiArabiaNafathAddress) GetDistrict() string`

GetDistrict returns the District field if non-nil, zero value otherwise.

### GetDistrictOk

`func (o *SaudiArabiaNafathAddress) GetDistrictOk() (*string, bool)`

GetDistrictOk returns a tuple with the District field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistrict

`func (o *SaudiArabiaNafathAddress) SetDistrict(v string)`

SetDistrict sets District field to given value.

### HasDistrict

`func (o *SaudiArabiaNafathAddress) HasDistrict() bool`

HasDistrict returns a boolean if a field has been set.

### SetDistrictNil

`func (o *SaudiArabiaNafathAddress) SetDistrictNil(b bool)`

 SetDistrictNil sets the value for District to be an explicit nil

### UnsetDistrict
`func (o *SaudiArabiaNafathAddress) UnsetDistrict()`

UnsetDistrict ensures that no value is present for District, not even an explicit nil
### GetCity

`func (o *SaudiArabiaNafathAddress) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *SaudiArabiaNafathAddress) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *SaudiArabiaNafathAddress) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *SaudiArabiaNafathAddress) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *SaudiArabiaNafathAddress) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *SaudiArabiaNafathAddress) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetPostalCode

`func (o *SaudiArabiaNafathAddress) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *SaudiArabiaNafathAddress) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *SaudiArabiaNafathAddress) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *SaudiArabiaNafathAddress) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### SetPostalCodeNil

`func (o *SaudiArabiaNafathAddress) SetPostalCodeNil(b bool)`

 SetPostalCodeNil sets the value for PostalCode to be an explicit nil

### UnsetPostalCode
`func (o *SaudiArabiaNafathAddress) UnsetPostalCode()`

UnsetPostalCode ensures that no value is present for PostalCode, not even an explicit nil
### GetAdditionalNumber

`func (o *SaudiArabiaNafathAddress) GetAdditionalNumber() string`

GetAdditionalNumber returns the AdditionalNumber field if non-nil, zero value otherwise.

### GetAdditionalNumberOk

`func (o *SaudiArabiaNafathAddress) GetAdditionalNumberOk() (*string, bool)`

GetAdditionalNumberOk returns a tuple with the AdditionalNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalNumber

`func (o *SaudiArabiaNafathAddress) SetAdditionalNumber(v string)`

SetAdditionalNumber sets AdditionalNumber field to given value.

### HasAdditionalNumber

`func (o *SaudiArabiaNafathAddress) HasAdditionalNumber() bool`

HasAdditionalNumber returns a boolean if a field has been set.

### SetAdditionalNumberNil

`func (o *SaudiArabiaNafathAddress) SetAdditionalNumberNil(b bool)`

 SetAdditionalNumberNil sets the value for AdditionalNumber to be an explicit nil

### UnsetAdditionalNumber
`func (o *SaudiArabiaNafathAddress) UnsetAdditionalNumber()`

UnsetAdditionalNumber ensures that no value is present for AdditionalNumber, not even an explicit nil
### GetUnitNumber

`func (o *SaudiArabiaNafathAddress) GetUnitNumber() string`

GetUnitNumber returns the UnitNumber field if non-nil, zero value otherwise.

### GetUnitNumberOk

`func (o *SaudiArabiaNafathAddress) GetUnitNumberOk() (*string, bool)`

GetUnitNumberOk returns a tuple with the UnitNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitNumber

`func (o *SaudiArabiaNafathAddress) SetUnitNumber(v string)`

SetUnitNumber sets UnitNumber field to given value.

### HasUnitNumber

`func (o *SaudiArabiaNafathAddress) HasUnitNumber() bool`

HasUnitNumber returns a boolean if a field has been set.

### SetUnitNumberNil

`func (o *SaudiArabiaNafathAddress) SetUnitNumberNil(b bool)`

 SetUnitNumberNil sets the value for UnitNumber to be an explicit nil

### UnsetUnitNumber
`func (o *SaudiArabiaNafathAddress) UnsetUnitNumber()`

UnsetUnitNumber ensures that no value is present for UnitNumber, not even an explicit nil
### GetLocationCoordinates

`func (o *SaudiArabiaNafathAddress) GetLocationCoordinates() string`

GetLocationCoordinates returns the LocationCoordinates field if non-nil, zero value otherwise.

### GetLocationCoordinatesOk

`func (o *SaudiArabiaNafathAddress) GetLocationCoordinatesOk() (*string, bool)`

GetLocationCoordinatesOk returns a tuple with the LocationCoordinates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationCoordinates

`func (o *SaudiArabiaNafathAddress) SetLocationCoordinates(v string)`

SetLocationCoordinates sets LocationCoordinates field to given value.

### HasLocationCoordinates

`func (o *SaudiArabiaNafathAddress) HasLocationCoordinates() bool`

HasLocationCoordinates returns a boolean if a field has been set.

### SetLocationCoordinatesNil

`func (o *SaudiArabiaNafathAddress) SetLocationCoordinatesNil(b bool)`

 SetLocationCoordinatesNil sets the value for LocationCoordinates to be an explicit nil

### UnsetLocationCoordinates
`func (o *SaudiArabiaNafathAddress) UnsetLocationCoordinates()`

UnsetLocationCoordinates ensures that no value is present for LocationCoordinates, not even an explicit nil
### GetIsPrimaryAddress

`func (o *SaudiArabiaNafathAddress) GetIsPrimaryAddress() bool`

GetIsPrimaryAddress returns the IsPrimaryAddress field if non-nil, zero value otherwise.

### GetIsPrimaryAddressOk

`func (o *SaudiArabiaNafathAddress) GetIsPrimaryAddressOk() (*bool, bool)`

GetIsPrimaryAddressOk returns a tuple with the IsPrimaryAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrimaryAddress

`func (o *SaudiArabiaNafathAddress) SetIsPrimaryAddress(v bool)`

SetIsPrimaryAddress sets IsPrimaryAddress field to given value.

### HasIsPrimaryAddress

`func (o *SaudiArabiaNafathAddress) HasIsPrimaryAddress() bool`

HasIsPrimaryAddress returns a boolean if a field has been set.

### SetIsPrimaryAddressNil

`func (o *SaudiArabiaNafathAddress) SetIsPrimaryAddressNil(b bool)`

 SetIsPrimaryAddressNil sets the value for IsPrimaryAddress to be an explicit nil

### UnsetIsPrimaryAddress
`func (o *SaudiArabiaNafathAddress) UnsetIsPrimaryAddress()`

UnsetIsPrimaryAddress ensures that no value is present for IsPrimaryAddress, not even an explicit nil
### GetRegionName

`func (o *SaudiArabiaNafathAddress) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### GetRegionNameOk

`func (o *SaudiArabiaNafathAddress) GetRegionNameOk() (*string, bool)`

GetRegionNameOk returns a tuple with the RegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionName

`func (o *SaudiArabiaNafathAddress) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.

### HasRegionName

`func (o *SaudiArabiaNafathAddress) HasRegionName() bool`

HasRegionName returns a boolean if a field has been set.

### SetRegionNameNil

`func (o *SaudiArabiaNafathAddress) SetRegionNameNil(b bool)`

 SetRegionNameNil sets the value for RegionName to be an explicit nil

### UnsetRegionName
`func (o *SaudiArabiaNafathAddress) UnsetRegionName()`

UnsetRegionName ensures that no value is present for RegionName, not even an explicit nil
### GetRegionNameSecondary

`func (o *SaudiArabiaNafathAddress) GetRegionNameSecondary() string`

GetRegionNameSecondary returns the RegionNameSecondary field if non-nil, zero value otherwise.

### GetRegionNameSecondaryOk

`func (o *SaudiArabiaNafathAddress) GetRegionNameSecondaryOk() (*string, bool)`

GetRegionNameSecondaryOk returns a tuple with the RegionNameSecondary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionNameSecondary

`func (o *SaudiArabiaNafathAddress) SetRegionNameSecondary(v string)`

SetRegionNameSecondary sets RegionNameSecondary field to given value.

### HasRegionNameSecondary

`func (o *SaudiArabiaNafathAddress) HasRegionNameSecondary() bool`

HasRegionNameSecondary returns a boolean if a field has been set.

### SetRegionNameSecondaryNil

`func (o *SaudiArabiaNafathAddress) SetRegionNameSecondaryNil(b bool)`

 SetRegionNameSecondaryNil sets the value for RegionNameSecondary to be an explicit nil

### UnsetRegionNameSecondary
`func (o *SaudiArabiaNafathAddress) UnsetRegionNameSecondary()`

UnsetRegionNameSecondary ensures that no value is present for RegionNameSecondary, not even an explicit nil
### GetCitySecondary

`func (o *SaudiArabiaNafathAddress) GetCitySecondary() string`

GetCitySecondary returns the CitySecondary field if non-nil, zero value otherwise.

### GetCitySecondaryOk

`func (o *SaudiArabiaNafathAddress) GetCitySecondaryOk() (*string, bool)`

GetCitySecondaryOk returns a tuple with the CitySecondary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitySecondary

`func (o *SaudiArabiaNafathAddress) SetCitySecondary(v string)`

SetCitySecondary sets CitySecondary field to given value.

### HasCitySecondary

`func (o *SaudiArabiaNafathAddress) HasCitySecondary() bool`

HasCitySecondary returns a boolean if a field has been set.

### SetCitySecondaryNil

`func (o *SaudiArabiaNafathAddress) SetCitySecondaryNil(b bool)`

 SetCitySecondaryNil sets the value for CitySecondary to be an explicit nil

### UnsetCitySecondary
`func (o *SaudiArabiaNafathAddress) UnsetCitySecondary()`

UnsetCitySecondary ensures that no value is present for CitySecondary, not even an explicit nil
### GetStreetSecondary

`func (o *SaudiArabiaNafathAddress) GetStreetSecondary() string`

GetStreetSecondary returns the StreetSecondary field if non-nil, zero value otherwise.

### GetStreetSecondaryOk

`func (o *SaudiArabiaNafathAddress) GetStreetSecondaryOk() (*string, bool)`

GetStreetSecondaryOk returns a tuple with the StreetSecondary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetSecondary

`func (o *SaudiArabiaNafathAddress) SetStreetSecondary(v string)`

SetStreetSecondary sets StreetSecondary field to given value.

### HasStreetSecondary

`func (o *SaudiArabiaNafathAddress) HasStreetSecondary() bool`

HasStreetSecondary returns a boolean if a field has been set.

### SetStreetSecondaryNil

`func (o *SaudiArabiaNafathAddress) SetStreetSecondaryNil(b bool)`

 SetStreetSecondaryNil sets the value for StreetSecondary to be an explicit nil

### UnsetStreetSecondary
`func (o *SaudiArabiaNafathAddress) UnsetStreetSecondary()`

UnsetStreetSecondary ensures that no value is present for StreetSecondary, not even an explicit nil
### GetDistrictSecondary

`func (o *SaudiArabiaNafathAddress) GetDistrictSecondary() string`

GetDistrictSecondary returns the DistrictSecondary field if non-nil, zero value otherwise.

### GetDistrictSecondaryOk

`func (o *SaudiArabiaNafathAddress) GetDistrictSecondaryOk() (*string, bool)`

GetDistrictSecondaryOk returns a tuple with the DistrictSecondary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistrictSecondary

`func (o *SaudiArabiaNafathAddress) SetDistrictSecondary(v string)`

SetDistrictSecondary sets DistrictSecondary field to given value.

### HasDistrictSecondary

`func (o *SaudiArabiaNafathAddress) HasDistrictSecondary() bool`

HasDistrictSecondary returns a boolean if a field has been set.

### SetDistrictSecondaryNil

`func (o *SaudiArabiaNafathAddress) SetDistrictSecondaryNil(b bool)`

 SetDistrictSecondaryNil sets the value for DistrictSecondary to be an explicit nil

### UnsetDistrictSecondary
`func (o *SaudiArabiaNafathAddress) UnsetDistrictSecondary()`

UnsetDistrictSecondary ensures that no value is present for DistrictSecondary, not even an explicit nil
### GetCityId

`func (o *SaudiArabiaNafathAddress) GetCityId() string`

GetCityId returns the CityId field if non-nil, zero value otherwise.

### GetCityIdOk

`func (o *SaudiArabiaNafathAddress) GetCityIdOk() (*string, bool)`

GetCityIdOk returns a tuple with the CityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityId

`func (o *SaudiArabiaNafathAddress) SetCityId(v string)`

SetCityId sets CityId field to given value.

### HasCityId

`func (o *SaudiArabiaNafathAddress) HasCityId() bool`

HasCityId returns a boolean if a field has been set.

### SetCityIdNil

`func (o *SaudiArabiaNafathAddress) SetCityIdNil(b bool)`

 SetCityIdNil sets the value for CityId to be an explicit nil

### UnsetCityId
`func (o *SaudiArabiaNafathAddress) UnsetCityId()`

UnsetCityId ensures that no value is present for CityId, not even an explicit nil
### GetRegionId

`func (o *SaudiArabiaNafathAddress) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *SaudiArabiaNafathAddress) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *SaudiArabiaNafathAddress) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.

### HasRegionId

`func (o *SaudiArabiaNafathAddress) HasRegionId() bool`

HasRegionId returns a boolean if a field has been set.

### SetRegionIdNil

`func (o *SaudiArabiaNafathAddress) SetRegionIdNil(b bool)`

 SetRegionIdNil sets the value for RegionId to be an explicit nil

### UnsetRegionId
`func (o *SaudiArabiaNafathAddress) UnsetRegionId()`

UnsetRegionId ensures that no value is present for RegionId, not even an explicit nil
### GetDistrictId

`func (o *SaudiArabiaNafathAddress) GetDistrictId() string`

GetDistrictId returns the DistrictId field if non-nil, zero value otherwise.

### GetDistrictIdOk

`func (o *SaudiArabiaNafathAddress) GetDistrictIdOk() (*string, bool)`

GetDistrictIdOk returns a tuple with the DistrictId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistrictId

`func (o *SaudiArabiaNafathAddress) SetDistrictId(v string)`

SetDistrictId sets DistrictId field to given value.

### HasDistrictId

`func (o *SaudiArabiaNafathAddress) HasDistrictId() bool`

HasDistrictId returns a boolean if a field has been set.

### SetDistrictIdNil

`func (o *SaudiArabiaNafathAddress) SetDistrictIdNil(b bool)`

 SetDistrictIdNil sets the value for DistrictId to be an explicit nil

### UnsetDistrictId
`func (o *SaudiArabiaNafathAddress) UnsetDistrictId()`

UnsetDistrictId ensures that no value is present for DistrictId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


