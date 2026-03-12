# AadhaarLocalizedClaims

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Language** | Pointer to [**NullableAadhaarLanguage**](AadhaarLanguage.md) | The language code for the localized claims. | [optional] 
**Name** | Pointer to **NullableString** | The full name. | [optional] 
**CareOf** | Pointer to **NullableString** | The address&#39;s care-of field. This indicates that mail should be delivered to an individual through another person or entity who is a known resident at the address. | [optional] 
**Country** | Pointer to **NullableString** | The address&#39;s country. | [optional] 
**District** | Pointer to **NullableString** | The address&#39;s district. | [optional] 
**House** | Pointer to **NullableString** | The identifier for the house address. | [optional] 
**Location** | Pointer to **NullableString** | The location or locality that helps identify where the address is within a city. | [optional] 
**Landmark** | Pointer to **NullableString** | A landmark near the address. | [optional] 
**PostalCode** | Pointer to **NullableString** | The address&#39;s postal code. | [optional] 
**State** | Pointer to **NullableString** | The address&#39;s state or union territory. | [optional] 
**Street** | Pointer to **NullableString** | The address&#39;s street name. | [optional] 
**VillageTownCity** | Pointer to **NullableString** | The address&#39;s village/town/city. | [optional] 

## Methods

### NewAadhaarLocalizedClaims

`func NewAadhaarLocalizedClaims() *AadhaarLocalizedClaims`

NewAadhaarLocalizedClaims instantiates a new AadhaarLocalizedClaims object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAadhaarLocalizedClaimsWithDefaults

`func NewAadhaarLocalizedClaimsWithDefaults() *AadhaarLocalizedClaims`

NewAadhaarLocalizedClaimsWithDefaults instantiates a new AadhaarLocalizedClaims object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLanguage

`func (o *AadhaarLocalizedClaims) GetLanguage() AadhaarLanguage`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *AadhaarLocalizedClaims) GetLanguageOk() (*AadhaarLanguage, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *AadhaarLocalizedClaims) SetLanguage(v AadhaarLanguage)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *AadhaarLocalizedClaims) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### SetLanguageNil

`func (o *AadhaarLocalizedClaims) SetLanguageNil(b bool)`

 SetLanguageNil sets the value for Language to be an explicit nil

### UnsetLanguage
`func (o *AadhaarLocalizedClaims) UnsetLanguage()`

UnsetLanguage ensures that no value is present for Language, not even an explicit nil
### GetName

`func (o *AadhaarLocalizedClaims) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AadhaarLocalizedClaims) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AadhaarLocalizedClaims) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AadhaarLocalizedClaims) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AadhaarLocalizedClaims) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AadhaarLocalizedClaims) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCareOf

`func (o *AadhaarLocalizedClaims) GetCareOf() string`

GetCareOf returns the CareOf field if non-nil, zero value otherwise.

### GetCareOfOk

`func (o *AadhaarLocalizedClaims) GetCareOfOk() (*string, bool)`

GetCareOfOk returns a tuple with the CareOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCareOf

`func (o *AadhaarLocalizedClaims) SetCareOf(v string)`

SetCareOf sets CareOf field to given value.

### HasCareOf

`func (o *AadhaarLocalizedClaims) HasCareOf() bool`

HasCareOf returns a boolean if a field has been set.

### SetCareOfNil

`func (o *AadhaarLocalizedClaims) SetCareOfNil(b bool)`

 SetCareOfNil sets the value for CareOf to be an explicit nil

### UnsetCareOf
`func (o *AadhaarLocalizedClaims) UnsetCareOf()`

UnsetCareOf ensures that no value is present for CareOf, not even an explicit nil
### GetCountry

`func (o *AadhaarLocalizedClaims) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *AadhaarLocalizedClaims) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *AadhaarLocalizedClaims) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *AadhaarLocalizedClaims) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *AadhaarLocalizedClaims) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *AadhaarLocalizedClaims) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetDistrict

`func (o *AadhaarLocalizedClaims) GetDistrict() string`

GetDistrict returns the District field if non-nil, zero value otherwise.

### GetDistrictOk

`func (o *AadhaarLocalizedClaims) GetDistrictOk() (*string, bool)`

GetDistrictOk returns a tuple with the District field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistrict

`func (o *AadhaarLocalizedClaims) SetDistrict(v string)`

SetDistrict sets District field to given value.

### HasDistrict

`func (o *AadhaarLocalizedClaims) HasDistrict() bool`

HasDistrict returns a boolean if a field has been set.

### SetDistrictNil

`func (o *AadhaarLocalizedClaims) SetDistrictNil(b bool)`

 SetDistrictNil sets the value for District to be an explicit nil

### UnsetDistrict
`func (o *AadhaarLocalizedClaims) UnsetDistrict()`

UnsetDistrict ensures that no value is present for District, not even an explicit nil
### GetHouse

`func (o *AadhaarLocalizedClaims) GetHouse() string`

GetHouse returns the House field if non-nil, zero value otherwise.

### GetHouseOk

`func (o *AadhaarLocalizedClaims) GetHouseOk() (*string, bool)`

GetHouseOk returns a tuple with the House field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHouse

`func (o *AadhaarLocalizedClaims) SetHouse(v string)`

SetHouse sets House field to given value.

### HasHouse

`func (o *AadhaarLocalizedClaims) HasHouse() bool`

HasHouse returns a boolean if a field has been set.

### SetHouseNil

`func (o *AadhaarLocalizedClaims) SetHouseNil(b bool)`

 SetHouseNil sets the value for House to be an explicit nil

### UnsetHouse
`func (o *AadhaarLocalizedClaims) UnsetHouse()`

UnsetHouse ensures that no value is present for House, not even an explicit nil
### GetLocation

`func (o *AadhaarLocalizedClaims) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *AadhaarLocalizedClaims) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *AadhaarLocalizedClaims) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *AadhaarLocalizedClaims) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *AadhaarLocalizedClaims) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *AadhaarLocalizedClaims) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetLandmark

`func (o *AadhaarLocalizedClaims) GetLandmark() string`

GetLandmark returns the Landmark field if non-nil, zero value otherwise.

### GetLandmarkOk

`func (o *AadhaarLocalizedClaims) GetLandmarkOk() (*string, bool)`

GetLandmarkOk returns a tuple with the Landmark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLandmark

`func (o *AadhaarLocalizedClaims) SetLandmark(v string)`

SetLandmark sets Landmark field to given value.

### HasLandmark

`func (o *AadhaarLocalizedClaims) HasLandmark() bool`

HasLandmark returns a boolean if a field has been set.

### SetLandmarkNil

`func (o *AadhaarLocalizedClaims) SetLandmarkNil(b bool)`

 SetLandmarkNil sets the value for Landmark to be an explicit nil

### UnsetLandmark
`func (o *AadhaarLocalizedClaims) UnsetLandmark()`

UnsetLandmark ensures that no value is present for Landmark, not even an explicit nil
### GetPostalCode

`func (o *AadhaarLocalizedClaims) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *AadhaarLocalizedClaims) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *AadhaarLocalizedClaims) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *AadhaarLocalizedClaims) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### SetPostalCodeNil

`func (o *AadhaarLocalizedClaims) SetPostalCodeNil(b bool)`

 SetPostalCodeNil sets the value for PostalCode to be an explicit nil

### UnsetPostalCode
`func (o *AadhaarLocalizedClaims) UnsetPostalCode()`

UnsetPostalCode ensures that no value is present for PostalCode, not even an explicit nil
### GetState

`func (o *AadhaarLocalizedClaims) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AadhaarLocalizedClaims) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AadhaarLocalizedClaims) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *AadhaarLocalizedClaims) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *AadhaarLocalizedClaims) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *AadhaarLocalizedClaims) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetStreet

`func (o *AadhaarLocalizedClaims) GetStreet() string`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *AadhaarLocalizedClaims) GetStreetOk() (*string, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet

`func (o *AadhaarLocalizedClaims) SetStreet(v string)`

SetStreet sets Street field to given value.

### HasStreet

`func (o *AadhaarLocalizedClaims) HasStreet() bool`

HasStreet returns a boolean if a field has been set.

### SetStreetNil

`func (o *AadhaarLocalizedClaims) SetStreetNil(b bool)`

 SetStreetNil sets the value for Street to be an explicit nil

### UnsetStreet
`func (o *AadhaarLocalizedClaims) UnsetStreet()`

UnsetStreet ensures that no value is present for Street, not even an explicit nil
### GetVillageTownCity

`func (o *AadhaarLocalizedClaims) GetVillageTownCity() string`

GetVillageTownCity returns the VillageTownCity field if non-nil, zero value otherwise.

### GetVillageTownCityOk

`func (o *AadhaarLocalizedClaims) GetVillageTownCityOk() (*string, bool)`

GetVillageTownCityOk returns a tuple with the VillageTownCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVillageTownCity

`func (o *AadhaarLocalizedClaims) SetVillageTownCity(v string)`

SetVillageTownCity sets VillageTownCity field to given value.

### HasVillageTownCity

`func (o *AadhaarLocalizedClaims) HasVillageTownCity() bool`

HasVillageTownCity returns a boolean if a field has been set.

### SetVillageTownCityNil

`func (o *AadhaarLocalizedClaims) SetVillageTownCityNil(b bool)`

 SetVillageTownCityNil sets the value for VillageTownCity to be an explicit nil

### UnsetVillageTownCity
`func (o *AadhaarLocalizedClaims) UnsetVillageTownCity()`

UnsetVillageTownCity ensures that no value is present for VillageTownCity, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


