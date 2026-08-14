# CzechBankIdAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **NullableString** | The address type.              Possible values are: - PERMANENT_RESIDENCE - SECONDARY_RESIDENCE - UNKNOWN | [optional] 
**Street** | Pointer to **NullableString** | The street name.              This is usually present, but unused for small villages. | [optional] 
**BuildingApartment** | Pointer to **NullableString** | The address land registry number.              This is usually present, but unused for small houses. | [optional] 
**StreetNumber** | Pointer to **NullableString** | The additional address house number.              This is usually present, but unused for small villages. | [optional] 
**EvidenceNumber** | Pointer to **NullableString** | The house evidence number.              This is rarely used in suburbs. An evidence number is used for addresses that are residential, temporary or non-residential buildings, e.g. cabins, garages etc. | [optional] 
**City** | Pointer to **NullableString** | The city name. | [optional] 
**CityArea** | Pointer to **NullableString** | The city area name.              This is usually present. | [optional] 
**Zipcode** | Pointer to **NullableString** | The zip of the address. | [optional] 
**Country** | Pointer to **NullableString** | The country code in ISO 3166-1 alpha-2 format. | [optional] 
**RuianReference** | Pointer to **NullableString** | The address identifier in the Czech RUIAN address register.              RUIAN is the Czech Register of Territorial Identification, Addresses, and Real Estate. This value identifies the address record in that register and can be used to reconcile the address against Czech government address data. | [optional] 

## Methods

### NewCzechBankIdAddress

`func NewCzechBankIdAddress() *CzechBankIdAddress`

NewCzechBankIdAddress instantiates a new CzechBankIdAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCzechBankIdAddressWithDefaults

`func NewCzechBankIdAddressWithDefaults() *CzechBankIdAddress`

NewCzechBankIdAddressWithDefaults instantiates a new CzechBankIdAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CzechBankIdAddress) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CzechBankIdAddress) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CzechBankIdAddress) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CzechBankIdAddress) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *CzechBankIdAddress) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *CzechBankIdAddress) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetStreet

`func (o *CzechBankIdAddress) GetStreet() string`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *CzechBankIdAddress) GetStreetOk() (*string, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet

`func (o *CzechBankIdAddress) SetStreet(v string)`

SetStreet sets Street field to given value.

### HasStreet

`func (o *CzechBankIdAddress) HasStreet() bool`

HasStreet returns a boolean if a field has been set.

### SetStreetNil

`func (o *CzechBankIdAddress) SetStreetNil(b bool)`

 SetStreetNil sets the value for Street to be an explicit nil

### UnsetStreet
`func (o *CzechBankIdAddress) UnsetStreet()`

UnsetStreet ensures that no value is present for Street, not even an explicit nil
### GetBuildingApartment

`func (o *CzechBankIdAddress) GetBuildingApartment() string`

GetBuildingApartment returns the BuildingApartment field if non-nil, zero value otherwise.

### GetBuildingApartmentOk

`func (o *CzechBankIdAddress) GetBuildingApartmentOk() (*string, bool)`

GetBuildingApartmentOk returns a tuple with the BuildingApartment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildingApartment

`func (o *CzechBankIdAddress) SetBuildingApartment(v string)`

SetBuildingApartment sets BuildingApartment field to given value.

### HasBuildingApartment

`func (o *CzechBankIdAddress) HasBuildingApartment() bool`

HasBuildingApartment returns a boolean if a field has been set.

### SetBuildingApartmentNil

`func (o *CzechBankIdAddress) SetBuildingApartmentNil(b bool)`

 SetBuildingApartmentNil sets the value for BuildingApartment to be an explicit nil

### UnsetBuildingApartment
`func (o *CzechBankIdAddress) UnsetBuildingApartment()`

UnsetBuildingApartment ensures that no value is present for BuildingApartment, not even an explicit nil
### GetStreetNumber

`func (o *CzechBankIdAddress) GetStreetNumber() string`

GetStreetNumber returns the StreetNumber field if non-nil, zero value otherwise.

### GetStreetNumberOk

`func (o *CzechBankIdAddress) GetStreetNumberOk() (*string, bool)`

GetStreetNumberOk returns a tuple with the StreetNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetNumber

`func (o *CzechBankIdAddress) SetStreetNumber(v string)`

SetStreetNumber sets StreetNumber field to given value.

### HasStreetNumber

`func (o *CzechBankIdAddress) HasStreetNumber() bool`

HasStreetNumber returns a boolean if a field has been set.

### SetStreetNumberNil

`func (o *CzechBankIdAddress) SetStreetNumberNil(b bool)`

 SetStreetNumberNil sets the value for StreetNumber to be an explicit nil

### UnsetStreetNumber
`func (o *CzechBankIdAddress) UnsetStreetNumber()`

UnsetStreetNumber ensures that no value is present for StreetNumber, not even an explicit nil
### GetEvidenceNumber

`func (o *CzechBankIdAddress) GetEvidenceNumber() string`

GetEvidenceNumber returns the EvidenceNumber field if non-nil, zero value otherwise.

### GetEvidenceNumberOk

`func (o *CzechBankIdAddress) GetEvidenceNumberOk() (*string, bool)`

GetEvidenceNumberOk returns a tuple with the EvidenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceNumber

`func (o *CzechBankIdAddress) SetEvidenceNumber(v string)`

SetEvidenceNumber sets EvidenceNumber field to given value.

### HasEvidenceNumber

`func (o *CzechBankIdAddress) HasEvidenceNumber() bool`

HasEvidenceNumber returns a boolean if a field has been set.

### SetEvidenceNumberNil

`func (o *CzechBankIdAddress) SetEvidenceNumberNil(b bool)`

 SetEvidenceNumberNil sets the value for EvidenceNumber to be an explicit nil

### UnsetEvidenceNumber
`func (o *CzechBankIdAddress) UnsetEvidenceNumber()`

UnsetEvidenceNumber ensures that no value is present for EvidenceNumber, not even an explicit nil
### GetCity

`func (o *CzechBankIdAddress) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *CzechBankIdAddress) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *CzechBankIdAddress) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *CzechBankIdAddress) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *CzechBankIdAddress) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *CzechBankIdAddress) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetCityArea

`func (o *CzechBankIdAddress) GetCityArea() string`

GetCityArea returns the CityArea field if non-nil, zero value otherwise.

### GetCityAreaOk

`func (o *CzechBankIdAddress) GetCityAreaOk() (*string, bool)`

GetCityAreaOk returns a tuple with the CityArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityArea

`func (o *CzechBankIdAddress) SetCityArea(v string)`

SetCityArea sets CityArea field to given value.

### HasCityArea

`func (o *CzechBankIdAddress) HasCityArea() bool`

HasCityArea returns a boolean if a field has been set.

### SetCityAreaNil

`func (o *CzechBankIdAddress) SetCityAreaNil(b bool)`

 SetCityAreaNil sets the value for CityArea to be an explicit nil

### UnsetCityArea
`func (o *CzechBankIdAddress) UnsetCityArea()`

UnsetCityArea ensures that no value is present for CityArea, not even an explicit nil
### GetZipcode

`func (o *CzechBankIdAddress) GetZipcode() string`

GetZipcode returns the Zipcode field if non-nil, zero value otherwise.

### GetZipcodeOk

`func (o *CzechBankIdAddress) GetZipcodeOk() (*string, bool)`

GetZipcodeOk returns a tuple with the Zipcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipcode

`func (o *CzechBankIdAddress) SetZipcode(v string)`

SetZipcode sets Zipcode field to given value.

### HasZipcode

`func (o *CzechBankIdAddress) HasZipcode() bool`

HasZipcode returns a boolean if a field has been set.

### SetZipcodeNil

`func (o *CzechBankIdAddress) SetZipcodeNil(b bool)`

 SetZipcodeNil sets the value for Zipcode to be an explicit nil

### UnsetZipcode
`func (o *CzechBankIdAddress) UnsetZipcode()`

UnsetZipcode ensures that no value is present for Zipcode, not even an explicit nil
### GetCountry

`func (o *CzechBankIdAddress) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *CzechBankIdAddress) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *CzechBankIdAddress) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *CzechBankIdAddress) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *CzechBankIdAddress) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *CzechBankIdAddress) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetRuianReference

`func (o *CzechBankIdAddress) GetRuianReference() string`

GetRuianReference returns the RuianReference field if non-nil, zero value otherwise.

### GetRuianReferenceOk

`func (o *CzechBankIdAddress) GetRuianReferenceOk() (*string, bool)`

GetRuianReferenceOk returns a tuple with the RuianReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuianReference

`func (o *CzechBankIdAddress) SetRuianReference(v string)`

SetRuianReference sets RuianReference field to given value.

### HasRuianReference

`func (o *CzechBankIdAddress) HasRuianReference() bool`

HasRuianReference returns a boolean if a field has been set.

### SetRuianReferenceNil

`func (o *CzechBankIdAddress) SetRuianReferenceNil(b bool)`

 SetRuianReferenceNil sets the value for RuianReference to be an explicit nil

### UnsetRuianReference
`func (o *CzechBankIdAddress) UnsetRuianReference()`

UnsetRuianReference ensures that no value is present for RuianReference, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


