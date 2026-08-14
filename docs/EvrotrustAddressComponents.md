# EvrotrustAddressComponents

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**District** | Pointer to **NullableString** | The district fragment of the address. | [optional] 
**DistrictLocal** | Pointer to **NullableString** | The local-language district fragment of the address. | [optional] 
**Municipality** | Pointer to **NullableString** | The municipality fragment of the address. | [optional] 
**MunicipalityLocal** | Pointer to **NullableString** | The local-language municipality fragment of the address. | [optional] 
**Settlement** | Pointer to **NullableString** | The settlement fragment of the address. | [optional] 
**SettlementLocal** | Pointer to **NullableString** | The local-language settlement fragment of the address. | [optional] 
**Location** | Pointer to **NullableString** | The location or street fragment of the address. | [optional] 
**LocationLocal** | Pointer to **NullableString** | The local-language location or street fragment of the address. | [optional] 
**BuildingNumber** | Pointer to **NullableString** | The building number fragment of the address. | [optional] 
**BuildingNumberLocal** | Pointer to **NullableString** | The local-language building number fragment of the address. | [optional] 
**Entrance** | Pointer to **NullableString** | The entrance fragment of the address. | [optional] 
**EntranceLocal** | Pointer to **NullableString** | The local-language entrance fragment of the address. | [optional] 
**Floor** | Pointer to **NullableString** | The floor fragment of the address. | [optional] 
**FloorLocal** | Pointer to **NullableString** | The local-language floor fragment of the address. | [optional] 
**Apartment** | Pointer to **NullableString** | The apartment fragment of the address. | [optional] 
**ApartmentLocal** | Pointer to **NullableString** | The local-language apartment fragment of the address. | [optional] 

## Methods

### NewEvrotrustAddressComponents

`func NewEvrotrustAddressComponents() *EvrotrustAddressComponents`

NewEvrotrustAddressComponents instantiates a new EvrotrustAddressComponents object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvrotrustAddressComponentsWithDefaults

`func NewEvrotrustAddressComponentsWithDefaults() *EvrotrustAddressComponents`

NewEvrotrustAddressComponentsWithDefaults instantiates a new EvrotrustAddressComponents object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDistrict

`func (o *EvrotrustAddressComponents) GetDistrict() string`

GetDistrict returns the District field if non-nil, zero value otherwise.

### GetDistrictOk

`func (o *EvrotrustAddressComponents) GetDistrictOk() (*string, bool)`

GetDistrictOk returns a tuple with the District field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistrict

`func (o *EvrotrustAddressComponents) SetDistrict(v string)`

SetDistrict sets District field to given value.

### HasDistrict

`func (o *EvrotrustAddressComponents) HasDistrict() bool`

HasDistrict returns a boolean if a field has been set.

### SetDistrictNil

`func (o *EvrotrustAddressComponents) SetDistrictNil(b bool)`

 SetDistrictNil sets the value for District to be an explicit nil

### UnsetDistrict
`func (o *EvrotrustAddressComponents) UnsetDistrict()`

UnsetDistrict ensures that no value is present for District, not even an explicit nil
### GetDistrictLocal

`func (o *EvrotrustAddressComponents) GetDistrictLocal() string`

GetDistrictLocal returns the DistrictLocal field if non-nil, zero value otherwise.

### GetDistrictLocalOk

`func (o *EvrotrustAddressComponents) GetDistrictLocalOk() (*string, bool)`

GetDistrictLocalOk returns a tuple with the DistrictLocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistrictLocal

`func (o *EvrotrustAddressComponents) SetDistrictLocal(v string)`

SetDistrictLocal sets DistrictLocal field to given value.

### HasDistrictLocal

`func (o *EvrotrustAddressComponents) HasDistrictLocal() bool`

HasDistrictLocal returns a boolean if a field has been set.

### SetDistrictLocalNil

`func (o *EvrotrustAddressComponents) SetDistrictLocalNil(b bool)`

 SetDistrictLocalNil sets the value for DistrictLocal to be an explicit nil

### UnsetDistrictLocal
`func (o *EvrotrustAddressComponents) UnsetDistrictLocal()`

UnsetDistrictLocal ensures that no value is present for DistrictLocal, not even an explicit nil
### GetMunicipality

`func (o *EvrotrustAddressComponents) GetMunicipality() string`

GetMunicipality returns the Municipality field if non-nil, zero value otherwise.

### GetMunicipalityOk

`func (o *EvrotrustAddressComponents) GetMunicipalityOk() (*string, bool)`

GetMunicipalityOk returns a tuple with the Municipality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMunicipality

`func (o *EvrotrustAddressComponents) SetMunicipality(v string)`

SetMunicipality sets Municipality field to given value.

### HasMunicipality

`func (o *EvrotrustAddressComponents) HasMunicipality() bool`

HasMunicipality returns a boolean if a field has been set.

### SetMunicipalityNil

`func (o *EvrotrustAddressComponents) SetMunicipalityNil(b bool)`

 SetMunicipalityNil sets the value for Municipality to be an explicit nil

### UnsetMunicipality
`func (o *EvrotrustAddressComponents) UnsetMunicipality()`

UnsetMunicipality ensures that no value is present for Municipality, not even an explicit nil
### GetMunicipalityLocal

`func (o *EvrotrustAddressComponents) GetMunicipalityLocal() string`

GetMunicipalityLocal returns the MunicipalityLocal field if non-nil, zero value otherwise.

### GetMunicipalityLocalOk

`func (o *EvrotrustAddressComponents) GetMunicipalityLocalOk() (*string, bool)`

GetMunicipalityLocalOk returns a tuple with the MunicipalityLocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMunicipalityLocal

`func (o *EvrotrustAddressComponents) SetMunicipalityLocal(v string)`

SetMunicipalityLocal sets MunicipalityLocal field to given value.

### HasMunicipalityLocal

`func (o *EvrotrustAddressComponents) HasMunicipalityLocal() bool`

HasMunicipalityLocal returns a boolean if a field has been set.

### SetMunicipalityLocalNil

`func (o *EvrotrustAddressComponents) SetMunicipalityLocalNil(b bool)`

 SetMunicipalityLocalNil sets the value for MunicipalityLocal to be an explicit nil

### UnsetMunicipalityLocal
`func (o *EvrotrustAddressComponents) UnsetMunicipalityLocal()`

UnsetMunicipalityLocal ensures that no value is present for MunicipalityLocal, not even an explicit nil
### GetSettlement

`func (o *EvrotrustAddressComponents) GetSettlement() string`

GetSettlement returns the Settlement field if non-nil, zero value otherwise.

### GetSettlementOk

`func (o *EvrotrustAddressComponents) GetSettlementOk() (*string, bool)`

GetSettlementOk returns a tuple with the Settlement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlement

`func (o *EvrotrustAddressComponents) SetSettlement(v string)`

SetSettlement sets Settlement field to given value.

### HasSettlement

`func (o *EvrotrustAddressComponents) HasSettlement() bool`

HasSettlement returns a boolean if a field has been set.

### SetSettlementNil

`func (o *EvrotrustAddressComponents) SetSettlementNil(b bool)`

 SetSettlementNil sets the value for Settlement to be an explicit nil

### UnsetSettlement
`func (o *EvrotrustAddressComponents) UnsetSettlement()`

UnsetSettlement ensures that no value is present for Settlement, not even an explicit nil
### GetSettlementLocal

`func (o *EvrotrustAddressComponents) GetSettlementLocal() string`

GetSettlementLocal returns the SettlementLocal field if non-nil, zero value otherwise.

### GetSettlementLocalOk

`func (o *EvrotrustAddressComponents) GetSettlementLocalOk() (*string, bool)`

GetSettlementLocalOk returns a tuple with the SettlementLocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementLocal

`func (o *EvrotrustAddressComponents) SetSettlementLocal(v string)`

SetSettlementLocal sets SettlementLocal field to given value.

### HasSettlementLocal

`func (o *EvrotrustAddressComponents) HasSettlementLocal() bool`

HasSettlementLocal returns a boolean if a field has been set.

### SetSettlementLocalNil

`func (o *EvrotrustAddressComponents) SetSettlementLocalNil(b bool)`

 SetSettlementLocalNil sets the value for SettlementLocal to be an explicit nil

### UnsetSettlementLocal
`func (o *EvrotrustAddressComponents) UnsetSettlementLocal()`

UnsetSettlementLocal ensures that no value is present for SettlementLocal, not even an explicit nil
### GetLocation

`func (o *EvrotrustAddressComponents) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *EvrotrustAddressComponents) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *EvrotrustAddressComponents) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *EvrotrustAddressComponents) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *EvrotrustAddressComponents) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *EvrotrustAddressComponents) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetLocationLocal

`func (o *EvrotrustAddressComponents) GetLocationLocal() string`

GetLocationLocal returns the LocationLocal field if non-nil, zero value otherwise.

### GetLocationLocalOk

`func (o *EvrotrustAddressComponents) GetLocationLocalOk() (*string, bool)`

GetLocationLocalOk returns a tuple with the LocationLocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationLocal

`func (o *EvrotrustAddressComponents) SetLocationLocal(v string)`

SetLocationLocal sets LocationLocal field to given value.

### HasLocationLocal

`func (o *EvrotrustAddressComponents) HasLocationLocal() bool`

HasLocationLocal returns a boolean if a field has been set.

### SetLocationLocalNil

`func (o *EvrotrustAddressComponents) SetLocationLocalNil(b bool)`

 SetLocationLocalNil sets the value for LocationLocal to be an explicit nil

### UnsetLocationLocal
`func (o *EvrotrustAddressComponents) UnsetLocationLocal()`

UnsetLocationLocal ensures that no value is present for LocationLocal, not even an explicit nil
### GetBuildingNumber

`func (o *EvrotrustAddressComponents) GetBuildingNumber() string`

GetBuildingNumber returns the BuildingNumber field if non-nil, zero value otherwise.

### GetBuildingNumberOk

`func (o *EvrotrustAddressComponents) GetBuildingNumberOk() (*string, bool)`

GetBuildingNumberOk returns a tuple with the BuildingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildingNumber

`func (o *EvrotrustAddressComponents) SetBuildingNumber(v string)`

SetBuildingNumber sets BuildingNumber field to given value.

### HasBuildingNumber

`func (o *EvrotrustAddressComponents) HasBuildingNumber() bool`

HasBuildingNumber returns a boolean if a field has been set.

### SetBuildingNumberNil

`func (o *EvrotrustAddressComponents) SetBuildingNumberNil(b bool)`

 SetBuildingNumberNil sets the value for BuildingNumber to be an explicit nil

### UnsetBuildingNumber
`func (o *EvrotrustAddressComponents) UnsetBuildingNumber()`

UnsetBuildingNumber ensures that no value is present for BuildingNumber, not even an explicit nil
### GetBuildingNumberLocal

`func (o *EvrotrustAddressComponents) GetBuildingNumberLocal() string`

GetBuildingNumberLocal returns the BuildingNumberLocal field if non-nil, zero value otherwise.

### GetBuildingNumberLocalOk

`func (o *EvrotrustAddressComponents) GetBuildingNumberLocalOk() (*string, bool)`

GetBuildingNumberLocalOk returns a tuple with the BuildingNumberLocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildingNumberLocal

`func (o *EvrotrustAddressComponents) SetBuildingNumberLocal(v string)`

SetBuildingNumberLocal sets BuildingNumberLocal field to given value.

### HasBuildingNumberLocal

`func (o *EvrotrustAddressComponents) HasBuildingNumberLocal() bool`

HasBuildingNumberLocal returns a boolean if a field has been set.

### SetBuildingNumberLocalNil

`func (o *EvrotrustAddressComponents) SetBuildingNumberLocalNil(b bool)`

 SetBuildingNumberLocalNil sets the value for BuildingNumberLocal to be an explicit nil

### UnsetBuildingNumberLocal
`func (o *EvrotrustAddressComponents) UnsetBuildingNumberLocal()`

UnsetBuildingNumberLocal ensures that no value is present for BuildingNumberLocal, not even an explicit nil
### GetEntrance

`func (o *EvrotrustAddressComponents) GetEntrance() string`

GetEntrance returns the Entrance field if non-nil, zero value otherwise.

### GetEntranceOk

`func (o *EvrotrustAddressComponents) GetEntranceOk() (*string, bool)`

GetEntranceOk returns a tuple with the Entrance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntrance

`func (o *EvrotrustAddressComponents) SetEntrance(v string)`

SetEntrance sets Entrance field to given value.

### HasEntrance

`func (o *EvrotrustAddressComponents) HasEntrance() bool`

HasEntrance returns a boolean if a field has been set.

### SetEntranceNil

`func (o *EvrotrustAddressComponents) SetEntranceNil(b bool)`

 SetEntranceNil sets the value for Entrance to be an explicit nil

### UnsetEntrance
`func (o *EvrotrustAddressComponents) UnsetEntrance()`

UnsetEntrance ensures that no value is present for Entrance, not even an explicit nil
### GetEntranceLocal

`func (o *EvrotrustAddressComponents) GetEntranceLocal() string`

GetEntranceLocal returns the EntranceLocal field if non-nil, zero value otherwise.

### GetEntranceLocalOk

`func (o *EvrotrustAddressComponents) GetEntranceLocalOk() (*string, bool)`

GetEntranceLocalOk returns a tuple with the EntranceLocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntranceLocal

`func (o *EvrotrustAddressComponents) SetEntranceLocal(v string)`

SetEntranceLocal sets EntranceLocal field to given value.

### HasEntranceLocal

`func (o *EvrotrustAddressComponents) HasEntranceLocal() bool`

HasEntranceLocal returns a boolean if a field has been set.

### SetEntranceLocalNil

`func (o *EvrotrustAddressComponents) SetEntranceLocalNil(b bool)`

 SetEntranceLocalNil sets the value for EntranceLocal to be an explicit nil

### UnsetEntranceLocal
`func (o *EvrotrustAddressComponents) UnsetEntranceLocal()`

UnsetEntranceLocal ensures that no value is present for EntranceLocal, not even an explicit nil
### GetFloor

`func (o *EvrotrustAddressComponents) GetFloor() string`

GetFloor returns the Floor field if non-nil, zero value otherwise.

### GetFloorOk

`func (o *EvrotrustAddressComponents) GetFloorOk() (*string, bool)`

GetFloorOk returns a tuple with the Floor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloor

`func (o *EvrotrustAddressComponents) SetFloor(v string)`

SetFloor sets Floor field to given value.

### HasFloor

`func (o *EvrotrustAddressComponents) HasFloor() bool`

HasFloor returns a boolean if a field has been set.

### SetFloorNil

`func (o *EvrotrustAddressComponents) SetFloorNil(b bool)`

 SetFloorNil sets the value for Floor to be an explicit nil

### UnsetFloor
`func (o *EvrotrustAddressComponents) UnsetFloor()`

UnsetFloor ensures that no value is present for Floor, not even an explicit nil
### GetFloorLocal

`func (o *EvrotrustAddressComponents) GetFloorLocal() string`

GetFloorLocal returns the FloorLocal field if non-nil, zero value otherwise.

### GetFloorLocalOk

`func (o *EvrotrustAddressComponents) GetFloorLocalOk() (*string, bool)`

GetFloorLocalOk returns a tuple with the FloorLocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloorLocal

`func (o *EvrotrustAddressComponents) SetFloorLocal(v string)`

SetFloorLocal sets FloorLocal field to given value.

### HasFloorLocal

`func (o *EvrotrustAddressComponents) HasFloorLocal() bool`

HasFloorLocal returns a boolean if a field has been set.

### SetFloorLocalNil

`func (o *EvrotrustAddressComponents) SetFloorLocalNil(b bool)`

 SetFloorLocalNil sets the value for FloorLocal to be an explicit nil

### UnsetFloorLocal
`func (o *EvrotrustAddressComponents) UnsetFloorLocal()`

UnsetFloorLocal ensures that no value is present for FloorLocal, not even an explicit nil
### GetApartment

`func (o *EvrotrustAddressComponents) GetApartment() string`

GetApartment returns the Apartment field if non-nil, zero value otherwise.

### GetApartmentOk

`func (o *EvrotrustAddressComponents) GetApartmentOk() (*string, bool)`

GetApartmentOk returns a tuple with the Apartment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApartment

`func (o *EvrotrustAddressComponents) SetApartment(v string)`

SetApartment sets Apartment field to given value.

### HasApartment

`func (o *EvrotrustAddressComponents) HasApartment() bool`

HasApartment returns a boolean if a field has been set.

### SetApartmentNil

`func (o *EvrotrustAddressComponents) SetApartmentNil(b bool)`

 SetApartmentNil sets the value for Apartment to be an explicit nil

### UnsetApartment
`func (o *EvrotrustAddressComponents) UnsetApartment()`

UnsetApartment ensures that no value is present for Apartment, not even an explicit nil
### GetApartmentLocal

`func (o *EvrotrustAddressComponents) GetApartmentLocal() string`

GetApartmentLocal returns the ApartmentLocal field if non-nil, zero value otherwise.

### GetApartmentLocalOk

`func (o *EvrotrustAddressComponents) GetApartmentLocalOk() (*string, bool)`

GetApartmentLocalOk returns a tuple with the ApartmentLocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApartmentLocal

`func (o *EvrotrustAddressComponents) SetApartmentLocal(v string)`

SetApartmentLocal sets ApartmentLocal field to given value.

### HasApartmentLocal

`func (o *EvrotrustAddressComponents) HasApartmentLocal() bool`

HasApartmentLocal returns a boolean if a field has been set.

### SetApartmentLocalNil

`func (o *EvrotrustAddressComponents) SetApartmentLocalNil(b bool)`

 SetApartmentLocalNil sets the value for ApartmentLocal to be an explicit nil

### UnsetApartmentLocal
`func (o *EvrotrustAddressComponents) UnsetApartmentLocal()`

UnsetApartmentLocal ensures that no value is present for ApartmentLocal, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


