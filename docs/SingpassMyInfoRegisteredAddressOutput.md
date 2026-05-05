# SingpassMyInfoRegisteredAddressOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **NullableString** | The address type.              Possible values: - SG (Structured Address) - UNFORMATTED              Structured addresses most likely will be a Singapore address. Unformatted are typically a non-Singapore address, however it might still be an unformatted Singapore address. | [optional] 
**Block** | Pointer to **NullableString** | The block number.              This will only be present if the type is \&quot;SG\&quot; | [optional] 
**Building** | Pointer to **NullableString** | The building name.              This will only be present if the type is \&quot;SG\&quot; | [optional] 
**Floor** | Pointer to **NullableString** | The floor number.              This will only be present if the type is \&quot;SG\&quot; | [optional] 
**Unit** | Pointer to **NullableString** | The unit number.              This will only be present if the type is \&quot;SG\&quot; | [optional] 
**Street** | Pointer to **NullableString** | The street name.              This will only be present if the type is \&quot;SG\&quot; | [optional] 
**Postal** | Pointer to **NullableString** | The postal code.              This will only be present if the type is \&quot;SG\&quot; | [optional] 
**Country** | Pointer to **NullableString** | The ISO-2 country code and description for the address.              This will only be present if the type is \&quot;SG\&quot; | [optional] 
**Line1** | Pointer to **NullableString** | The first line of the unformatted address.              This will only be present if the type is \&quot;UNFORMATTED\&quot; | [optional] 
**Line2** | Pointer to **NullableString** | The second line of the unformatted address.              This will only be present if the type is \&quot;UNFORMATTED\&quot; | [optional] 

## Methods

### NewSingpassMyInfoRegisteredAddressOutput

`func NewSingpassMyInfoRegisteredAddressOutput() *SingpassMyInfoRegisteredAddressOutput`

NewSingpassMyInfoRegisteredAddressOutput instantiates a new SingpassMyInfoRegisteredAddressOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSingpassMyInfoRegisteredAddressOutputWithDefaults

`func NewSingpassMyInfoRegisteredAddressOutputWithDefaults() *SingpassMyInfoRegisteredAddressOutput`

NewSingpassMyInfoRegisteredAddressOutputWithDefaults instantiates a new SingpassMyInfoRegisteredAddressOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SingpassMyInfoRegisteredAddressOutput) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SingpassMyInfoRegisteredAddressOutput) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SingpassMyInfoRegisteredAddressOutput) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SingpassMyInfoRegisteredAddressOutput) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *SingpassMyInfoRegisteredAddressOutput) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *SingpassMyInfoRegisteredAddressOutput) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetBlock

`func (o *SingpassMyInfoRegisteredAddressOutput) GetBlock() string`

GetBlock returns the Block field if non-nil, zero value otherwise.

### GetBlockOk

`func (o *SingpassMyInfoRegisteredAddressOutput) GetBlockOk() (*string, bool)`

GetBlockOk returns a tuple with the Block field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlock

`func (o *SingpassMyInfoRegisteredAddressOutput) SetBlock(v string)`

SetBlock sets Block field to given value.

### HasBlock

`func (o *SingpassMyInfoRegisteredAddressOutput) HasBlock() bool`

HasBlock returns a boolean if a field has been set.

### SetBlockNil

`func (o *SingpassMyInfoRegisteredAddressOutput) SetBlockNil(b bool)`

 SetBlockNil sets the value for Block to be an explicit nil

### UnsetBlock
`func (o *SingpassMyInfoRegisteredAddressOutput) UnsetBlock()`

UnsetBlock ensures that no value is present for Block, not even an explicit nil
### GetBuilding

`func (o *SingpassMyInfoRegisteredAddressOutput) GetBuilding() string`

GetBuilding returns the Building field if non-nil, zero value otherwise.

### GetBuildingOk

`func (o *SingpassMyInfoRegisteredAddressOutput) GetBuildingOk() (*string, bool)`

GetBuildingOk returns a tuple with the Building field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuilding

`func (o *SingpassMyInfoRegisteredAddressOutput) SetBuilding(v string)`

SetBuilding sets Building field to given value.

### HasBuilding

`func (o *SingpassMyInfoRegisteredAddressOutput) HasBuilding() bool`

HasBuilding returns a boolean if a field has been set.

### SetBuildingNil

`func (o *SingpassMyInfoRegisteredAddressOutput) SetBuildingNil(b bool)`

 SetBuildingNil sets the value for Building to be an explicit nil

### UnsetBuilding
`func (o *SingpassMyInfoRegisteredAddressOutput) UnsetBuilding()`

UnsetBuilding ensures that no value is present for Building, not even an explicit nil
### GetFloor

`func (o *SingpassMyInfoRegisteredAddressOutput) GetFloor() string`

GetFloor returns the Floor field if non-nil, zero value otherwise.

### GetFloorOk

`func (o *SingpassMyInfoRegisteredAddressOutput) GetFloorOk() (*string, bool)`

GetFloorOk returns a tuple with the Floor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloor

`func (o *SingpassMyInfoRegisteredAddressOutput) SetFloor(v string)`

SetFloor sets Floor field to given value.

### HasFloor

`func (o *SingpassMyInfoRegisteredAddressOutput) HasFloor() bool`

HasFloor returns a boolean if a field has been set.

### SetFloorNil

`func (o *SingpassMyInfoRegisteredAddressOutput) SetFloorNil(b bool)`

 SetFloorNil sets the value for Floor to be an explicit nil

### UnsetFloor
`func (o *SingpassMyInfoRegisteredAddressOutput) UnsetFloor()`

UnsetFloor ensures that no value is present for Floor, not even an explicit nil
### GetUnit

`func (o *SingpassMyInfoRegisteredAddressOutput) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *SingpassMyInfoRegisteredAddressOutput) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *SingpassMyInfoRegisteredAddressOutput) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *SingpassMyInfoRegisteredAddressOutput) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### SetUnitNil

`func (o *SingpassMyInfoRegisteredAddressOutput) SetUnitNil(b bool)`

 SetUnitNil sets the value for Unit to be an explicit nil

### UnsetUnit
`func (o *SingpassMyInfoRegisteredAddressOutput) UnsetUnit()`

UnsetUnit ensures that no value is present for Unit, not even an explicit nil
### GetStreet

`func (o *SingpassMyInfoRegisteredAddressOutput) GetStreet() string`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *SingpassMyInfoRegisteredAddressOutput) GetStreetOk() (*string, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet

`func (o *SingpassMyInfoRegisteredAddressOutput) SetStreet(v string)`

SetStreet sets Street field to given value.

### HasStreet

`func (o *SingpassMyInfoRegisteredAddressOutput) HasStreet() bool`

HasStreet returns a boolean if a field has been set.

### SetStreetNil

`func (o *SingpassMyInfoRegisteredAddressOutput) SetStreetNil(b bool)`

 SetStreetNil sets the value for Street to be an explicit nil

### UnsetStreet
`func (o *SingpassMyInfoRegisteredAddressOutput) UnsetStreet()`

UnsetStreet ensures that no value is present for Street, not even an explicit nil
### GetPostal

`func (o *SingpassMyInfoRegisteredAddressOutput) GetPostal() string`

GetPostal returns the Postal field if non-nil, zero value otherwise.

### GetPostalOk

`func (o *SingpassMyInfoRegisteredAddressOutput) GetPostalOk() (*string, bool)`

GetPostalOk returns a tuple with the Postal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostal

`func (o *SingpassMyInfoRegisteredAddressOutput) SetPostal(v string)`

SetPostal sets Postal field to given value.

### HasPostal

`func (o *SingpassMyInfoRegisteredAddressOutput) HasPostal() bool`

HasPostal returns a boolean if a field has been set.

### SetPostalNil

`func (o *SingpassMyInfoRegisteredAddressOutput) SetPostalNil(b bool)`

 SetPostalNil sets the value for Postal to be an explicit nil

### UnsetPostal
`func (o *SingpassMyInfoRegisteredAddressOutput) UnsetPostal()`

UnsetPostal ensures that no value is present for Postal, not even an explicit nil
### GetCountry

`func (o *SingpassMyInfoRegisteredAddressOutput) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *SingpassMyInfoRegisteredAddressOutput) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *SingpassMyInfoRegisteredAddressOutput) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *SingpassMyInfoRegisteredAddressOutput) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *SingpassMyInfoRegisteredAddressOutput) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *SingpassMyInfoRegisteredAddressOutput) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetLine1

`func (o *SingpassMyInfoRegisteredAddressOutput) GetLine1() string`

GetLine1 returns the Line1 field if non-nil, zero value otherwise.

### GetLine1Ok

`func (o *SingpassMyInfoRegisteredAddressOutput) GetLine1Ok() (*string, bool)`

GetLine1Ok returns a tuple with the Line1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine1

`func (o *SingpassMyInfoRegisteredAddressOutput) SetLine1(v string)`

SetLine1 sets Line1 field to given value.

### HasLine1

`func (o *SingpassMyInfoRegisteredAddressOutput) HasLine1() bool`

HasLine1 returns a boolean if a field has been set.

### SetLine1Nil

`func (o *SingpassMyInfoRegisteredAddressOutput) SetLine1Nil(b bool)`

 SetLine1Nil sets the value for Line1 to be an explicit nil

### UnsetLine1
`func (o *SingpassMyInfoRegisteredAddressOutput) UnsetLine1()`

UnsetLine1 ensures that no value is present for Line1, not even an explicit nil
### GetLine2

`func (o *SingpassMyInfoRegisteredAddressOutput) GetLine2() string`

GetLine2 returns the Line2 field if non-nil, zero value otherwise.

### GetLine2Ok

`func (o *SingpassMyInfoRegisteredAddressOutput) GetLine2Ok() (*string, bool)`

GetLine2Ok returns a tuple with the Line2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine2

`func (o *SingpassMyInfoRegisteredAddressOutput) SetLine2(v string)`

SetLine2 sets Line2 field to given value.

### HasLine2

`func (o *SingpassMyInfoRegisteredAddressOutput) HasLine2() bool`

HasLine2 returns a boolean if a field has been set.

### SetLine2Nil

`func (o *SingpassMyInfoRegisteredAddressOutput) SetLine2Nil(b bool)`

 SetLine2Nil sets the value for Line2 to be an explicit nil

### UnsetLine2
`func (o *SingpassMyInfoRegisteredAddressOutput) UnsetLine2()`

UnsetLine2 ensures that no value is present for Line2, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


