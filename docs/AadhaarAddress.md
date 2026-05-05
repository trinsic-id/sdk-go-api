# AadhaarAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CareOf** | Pointer to **NullableString** | The address&#39;s care-of field. This indicates that mail should be delivered to an individual through another person or entity who is a known resident at the address. | [optional] 
**Country** | Pointer to **NullableString** | The address&#39;s country.              This is usually formatted as an alpha-2 country code, but localized output may instead contain the country name as written on the document. | [optional] 
**District** | Pointer to **NullableString** | The address&#39;s district. | [optional] 
**House** | Pointer to **NullableString** | The identifier for the house address. | [optional] 
**Location** | Pointer to **NullableString** | The location or locality that helps identify where the address is within a city. | [optional] 
**Landmark** | Pointer to **NullableString** | A landmark near the address. | [optional] 
**PostalCode** | Pointer to **NullableString** | The address&#39;s postal code. | [optional] 
**PostOffice** | Pointer to **NullableString** | The address&#39;s post office. | [optional] 
**State** | Pointer to **NullableString** | The address&#39;s state or union territory. | [optional] 
**Street** | Pointer to **NullableString** | The address&#39;s street name. | [optional] 
**Subdistrict** | Pointer to **NullableString** | The address&#39;s subdistrict. | [optional] 
**VillageTownCity** | Pointer to **NullableString** | The address&#39;s village/town/city. | [optional] 

## Methods

### NewAadhaarAddress

`func NewAadhaarAddress() *AadhaarAddress`

NewAadhaarAddress instantiates a new AadhaarAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAadhaarAddressWithDefaults

`func NewAadhaarAddressWithDefaults() *AadhaarAddress`

NewAadhaarAddressWithDefaults instantiates a new AadhaarAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCareOf

`func (o *AadhaarAddress) GetCareOf() string`

GetCareOf returns the CareOf field if non-nil, zero value otherwise.

### GetCareOfOk

`func (o *AadhaarAddress) GetCareOfOk() (*string, bool)`

GetCareOfOk returns a tuple with the CareOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCareOf

`func (o *AadhaarAddress) SetCareOf(v string)`

SetCareOf sets CareOf field to given value.

### HasCareOf

`func (o *AadhaarAddress) HasCareOf() bool`

HasCareOf returns a boolean if a field has been set.

### SetCareOfNil

`func (o *AadhaarAddress) SetCareOfNil(b bool)`

 SetCareOfNil sets the value for CareOf to be an explicit nil

### UnsetCareOf
`func (o *AadhaarAddress) UnsetCareOf()`

UnsetCareOf ensures that no value is present for CareOf, not even an explicit nil
### GetCountry

`func (o *AadhaarAddress) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *AadhaarAddress) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *AadhaarAddress) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *AadhaarAddress) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *AadhaarAddress) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *AadhaarAddress) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetDistrict

`func (o *AadhaarAddress) GetDistrict() string`

GetDistrict returns the District field if non-nil, zero value otherwise.

### GetDistrictOk

`func (o *AadhaarAddress) GetDistrictOk() (*string, bool)`

GetDistrictOk returns a tuple with the District field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistrict

`func (o *AadhaarAddress) SetDistrict(v string)`

SetDistrict sets District field to given value.

### HasDistrict

`func (o *AadhaarAddress) HasDistrict() bool`

HasDistrict returns a boolean if a field has been set.

### SetDistrictNil

`func (o *AadhaarAddress) SetDistrictNil(b bool)`

 SetDistrictNil sets the value for District to be an explicit nil

### UnsetDistrict
`func (o *AadhaarAddress) UnsetDistrict()`

UnsetDistrict ensures that no value is present for District, not even an explicit nil
### GetHouse

`func (o *AadhaarAddress) GetHouse() string`

GetHouse returns the House field if non-nil, zero value otherwise.

### GetHouseOk

`func (o *AadhaarAddress) GetHouseOk() (*string, bool)`

GetHouseOk returns a tuple with the House field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHouse

`func (o *AadhaarAddress) SetHouse(v string)`

SetHouse sets House field to given value.

### HasHouse

`func (o *AadhaarAddress) HasHouse() bool`

HasHouse returns a boolean if a field has been set.

### SetHouseNil

`func (o *AadhaarAddress) SetHouseNil(b bool)`

 SetHouseNil sets the value for House to be an explicit nil

### UnsetHouse
`func (o *AadhaarAddress) UnsetHouse()`

UnsetHouse ensures that no value is present for House, not even an explicit nil
### GetLocation

`func (o *AadhaarAddress) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *AadhaarAddress) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *AadhaarAddress) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *AadhaarAddress) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *AadhaarAddress) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *AadhaarAddress) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetLandmark

`func (o *AadhaarAddress) GetLandmark() string`

GetLandmark returns the Landmark field if non-nil, zero value otherwise.

### GetLandmarkOk

`func (o *AadhaarAddress) GetLandmarkOk() (*string, bool)`

GetLandmarkOk returns a tuple with the Landmark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLandmark

`func (o *AadhaarAddress) SetLandmark(v string)`

SetLandmark sets Landmark field to given value.

### HasLandmark

`func (o *AadhaarAddress) HasLandmark() bool`

HasLandmark returns a boolean if a field has been set.

### SetLandmarkNil

`func (o *AadhaarAddress) SetLandmarkNil(b bool)`

 SetLandmarkNil sets the value for Landmark to be an explicit nil

### UnsetLandmark
`func (o *AadhaarAddress) UnsetLandmark()`

UnsetLandmark ensures that no value is present for Landmark, not even an explicit nil
### GetPostalCode

`func (o *AadhaarAddress) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *AadhaarAddress) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *AadhaarAddress) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *AadhaarAddress) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### SetPostalCodeNil

`func (o *AadhaarAddress) SetPostalCodeNil(b bool)`

 SetPostalCodeNil sets the value for PostalCode to be an explicit nil

### UnsetPostalCode
`func (o *AadhaarAddress) UnsetPostalCode()`

UnsetPostalCode ensures that no value is present for PostalCode, not even an explicit nil
### GetPostOffice

`func (o *AadhaarAddress) GetPostOffice() string`

GetPostOffice returns the PostOffice field if non-nil, zero value otherwise.

### GetPostOfficeOk

`func (o *AadhaarAddress) GetPostOfficeOk() (*string, bool)`

GetPostOfficeOk returns a tuple with the PostOffice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostOffice

`func (o *AadhaarAddress) SetPostOffice(v string)`

SetPostOffice sets PostOffice field to given value.

### HasPostOffice

`func (o *AadhaarAddress) HasPostOffice() bool`

HasPostOffice returns a boolean if a field has been set.

### SetPostOfficeNil

`func (o *AadhaarAddress) SetPostOfficeNil(b bool)`

 SetPostOfficeNil sets the value for PostOffice to be an explicit nil

### UnsetPostOffice
`func (o *AadhaarAddress) UnsetPostOffice()`

UnsetPostOffice ensures that no value is present for PostOffice, not even an explicit nil
### GetState

`func (o *AadhaarAddress) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AadhaarAddress) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AadhaarAddress) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *AadhaarAddress) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *AadhaarAddress) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *AadhaarAddress) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetStreet

`func (o *AadhaarAddress) GetStreet() string`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *AadhaarAddress) GetStreetOk() (*string, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet

`func (o *AadhaarAddress) SetStreet(v string)`

SetStreet sets Street field to given value.

### HasStreet

`func (o *AadhaarAddress) HasStreet() bool`

HasStreet returns a boolean if a field has been set.

### SetStreetNil

`func (o *AadhaarAddress) SetStreetNil(b bool)`

 SetStreetNil sets the value for Street to be an explicit nil

### UnsetStreet
`func (o *AadhaarAddress) UnsetStreet()`

UnsetStreet ensures that no value is present for Street, not even an explicit nil
### GetSubdistrict

`func (o *AadhaarAddress) GetSubdistrict() string`

GetSubdistrict returns the Subdistrict field if non-nil, zero value otherwise.

### GetSubdistrictOk

`func (o *AadhaarAddress) GetSubdistrictOk() (*string, bool)`

GetSubdistrictOk returns a tuple with the Subdistrict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdistrict

`func (o *AadhaarAddress) SetSubdistrict(v string)`

SetSubdistrict sets Subdistrict field to given value.

### HasSubdistrict

`func (o *AadhaarAddress) HasSubdistrict() bool`

HasSubdistrict returns a boolean if a field has been set.

### SetSubdistrictNil

`func (o *AadhaarAddress) SetSubdistrictNil(b bool)`

 SetSubdistrictNil sets the value for Subdistrict to be an explicit nil

### UnsetSubdistrict
`func (o *AadhaarAddress) UnsetSubdistrict()`

UnsetSubdistrict ensures that no value is present for Subdistrict, not even an explicit nil
### GetVillageTownCity

`func (o *AadhaarAddress) GetVillageTownCity() string`

GetVillageTownCity returns the VillageTownCity field if non-nil, zero value otherwise.

### GetVillageTownCityOk

`func (o *AadhaarAddress) GetVillageTownCityOk() (*string, bool)`

GetVillageTownCityOk returns a tuple with the VillageTownCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVillageTownCity

`func (o *AadhaarAddress) SetVillageTownCity(v string)`

SetVillageTownCity sets VillageTownCity field to given value.

### HasVillageTownCity

`func (o *AadhaarAddress) HasVillageTownCity() bool`

HasVillageTownCity returns a boolean if a field has been set.

### SetVillageTownCityNil

`func (o *AadhaarAddress) SetVillageTownCityNil(b bool)`

 SetVillageTownCityNil sets the value for VillageTownCity to be an explicit nil

### UnsetVillageTownCity
`func (o *AadhaarAddress) UnsetVillageTownCity()`

UnsetVillageTownCity ensures that no value is present for VillageTownCity, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


