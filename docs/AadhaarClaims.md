# AadhaarClaims

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | The full name. | [optional] 
**DateOfBirth** | Pointer to **NullableString** | The date of birth.              The format is YYYY-MM-DD. | [optional] 
**Gender** | Pointer to **NullableString** | The gender of the individual.              Possible values: - M (Male) - F (Female) - T (Transgender) | [optional] 
**CareOf** | Pointer to **NullableString** | The address&#39;s care-of field. This indicates that mail should be delivered to an individual through another person or entity who is a known resident at the address. | [optional] 
**Country** | Pointer to **NullableString** | The address&#39;s country.              This is formatted as an alpha-2 country code. | [optional] 
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

### NewAadhaarClaims

`func NewAadhaarClaims() *AadhaarClaims`

NewAadhaarClaims instantiates a new AadhaarClaims object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAadhaarClaimsWithDefaults

`func NewAadhaarClaimsWithDefaults() *AadhaarClaims`

NewAadhaarClaimsWithDefaults instantiates a new AadhaarClaims object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AadhaarClaims) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AadhaarClaims) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AadhaarClaims) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AadhaarClaims) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AadhaarClaims) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AadhaarClaims) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDateOfBirth

`func (o *AadhaarClaims) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *AadhaarClaims) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *AadhaarClaims) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *AadhaarClaims) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### SetDateOfBirthNil

`func (o *AadhaarClaims) SetDateOfBirthNil(b bool)`

 SetDateOfBirthNil sets the value for DateOfBirth to be an explicit nil

### UnsetDateOfBirth
`func (o *AadhaarClaims) UnsetDateOfBirth()`

UnsetDateOfBirth ensures that no value is present for DateOfBirth, not even an explicit nil
### GetGender

`func (o *AadhaarClaims) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *AadhaarClaims) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *AadhaarClaims) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *AadhaarClaims) HasGender() bool`

HasGender returns a boolean if a field has been set.

### SetGenderNil

`func (o *AadhaarClaims) SetGenderNil(b bool)`

 SetGenderNil sets the value for Gender to be an explicit nil

### UnsetGender
`func (o *AadhaarClaims) UnsetGender()`

UnsetGender ensures that no value is present for Gender, not even an explicit nil
### GetCareOf

`func (o *AadhaarClaims) GetCareOf() string`

GetCareOf returns the CareOf field if non-nil, zero value otherwise.

### GetCareOfOk

`func (o *AadhaarClaims) GetCareOfOk() (*string, bool)`

GetCareOfOk returns a tuple with the CareOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCareOf

`func (o *AadhaarClaims) SetCareOf(v string)`

SetCareOf sets CareOf field to given value.

### HasCareOf

`func (o *AadhaarClaims) HasCareOf() bool`

HasCareOf returns a boolean if a field has been set.

### SetCareOfNil

`func (o *AadhaarClaims) SetCareOfNil(b bool)`

 SetCareOfNil sets the value for CareOf to be an explicit nil

### UnsetCareOf
`func (o *AadhaarClaims) UnsetCareOf()`

UnsetCareOf ensures that no value is present for CareOf, not even an explicit nil
### GetCountry

`func (o *AadhaarClaims) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *AadhaarClaims) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *AadhaarClaims) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *AadhaarClaims) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *AadhaarClaims) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *AadhaarClaims) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetDistrict

`func (o *AadhaarClaims) GetDistrict() string`

GetDistrict returns the District field if non-nil, zero value otherwise.

### GetDistrictOk

`func (o *AadhaarClaims) GetDistrictOk() (*string, bool)`

GetDistrictOk returns a tuple with the District field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistrict

`func (o *AadhaarClaims) SetDistrict(v string)`

SetDistrict sets District field to given value.

### HasDistrict

`func (o *AadhaarClaims) HasDistrict() bool`

HasDistrict returns a boolean if a field has been set.

### SetDistrictNil

`func (o *AadhaarClaims) SetDistrictNil(b bool)`

 SetDistrictNil sets the value for District to be an explicit nil

### UnsetDistrict
`func (o *AadhaarClaims) UnsetDistrict()`

UnsetDistrict ensures that no value is present for District, not even an explicit nil
### GetHouse

`func (o *AadhaarClaims) GetHouse() string`

GetHouse returns the House field if non-nil, zero value otherwise.

### GetHouseOk

`func (o *AadhaarClaims) GetHouseOk() (*string, bool)`

GetHouseOk returns a tuple with the House field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHouse

`func (o *AadhaarClaims) SetHouse(v string)`

SetHouse sets House field to given value.

### HasHouse

`func (o *AadhaarClaims) HasHouse() bool`

HasHouse returns a boolean if a field has been set.

### SetHouseNil

`func (o *AadhaarClaims) SetHouseNil(b bool)`

 SetHouseNil sets the value for House to be an explicit nil

### UnsetHouse
`func (o *AadhaarClaims) UnsetHouse()`

UnsetHouse ensures that no value is present for House, not even an explicit nil
### GetLocation

`func (o *AadhaarClaims) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *AadhaarClaims) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *AadhaarClaims) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *AadhaarClaims) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *AadhaarClaims) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *AadhaarClaims) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetLandmark

`func (o *AadhaarClaims) GetLandmark() string`

GetLandmark returns the Landmark field if non-nil, zero value otherwise.

### GetLandmarkOk

`func (o *AadhaarClaims) GetLandmarkOk() (*string, bool)`

GetLandmarkOk returns a tuple with the Landmark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLandmark

`func (o *AadhaarClaims) SetLandmark(v string)`

SetLandmark sets Landmark field to given value.

### HasLandmark

`func (o *AadhaarClaims) HasLandmark() bool`

HasLandmark returns a boolean if a field has been set.

### SetLandmarkNil

`func (o *AadhaarClaims) SetLandmarkNil(b bool)`

 SetLandmarkNil sets the value for Landmark to be an explicit nil

### UnsetLandmark
`func (o *AadhaarClaims) UnsetLandmark()`

UnsetLandmark ensures that no value is present for Landmark, not even an explicit nil
### GetPostalCode

`func (o *AadhaarClaims) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *AadhaarClaims) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *AadhaarClaims) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *AadhaarClaims) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### SetPostalCodeNil

`func (o *AadhaarClaims) SetPostalCodeNil(b bool)`

 SetPostalCodeNil sets the value for PostalCode to be an explicit nil

### UnsetPostalCode
`func (o *AadhaarClaims) UnsetPostalCode()`

UnsetPostalCode ensures that no value is present for PostalCode, not even an explicit nil
### GetPostOffice

`func (o *AadhaarClaims) GetPostOffice() string`

GetPostOffice returns the PostOffice field if non-nil, zero value otherwise.

### GetPostOfficeOk

`func (o *AadhaarClaims) GetPostOfficeOk() (*string, bool)`

GetPostOfficeOk returns a tuple with the PostOffice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostOffice

`func (o *AadhaarClaims) SetPostOffice(v string)`

SetPostOffice sets PostOffice field to given value.

### HasPostOffice

`func (o *AadhaarClaims) HasPostOffice() bool`

HasPostOffice returns a boolean if a field has been set.

### SetPostOfficeNil

`func (o *AadhaarClaims) SetPostOfficeNil(b bool)`

 SetPostOfficeNil sets the value for PostOffice to be an explicit nil

### UnsetPostOffice
`func (o *AadhaarClaims) UnsetPostOffice()`

UnsetPostOffice ensures that no value is present for PostOffice, not even an explicit nil
### GetState

`func (o *AadhaarClaims) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AadhaarClaims) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AadhaarClaims) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *AadhaarClaims) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *AadhaarClaims) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *AadhaarClaims) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetStreet

`func (o *AadhaarClaims) GetStreet() string`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *AadhaarClaims) GetStreetOk() (*string, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet

`func (o *AadhaarClaims) SetStreet(v string)`

SetStreet sets Street field to given value.

### HasStreet

`func (o *AadhaarClaims) HasStreet() bool`

HasStreet returns a boolean if a field has been set.

### SetStreetNil

`func (o *AadhaarClaims) SetStreetNil(b bool)`

 SetStreetNil sets the value for Street to be an explicit nil

### UnsetStreet
`func (o *AadhaarClaims) UnsetStreet()`

UnsetStreet ensures that no value is present for Street, not even an explicit nil
### GetSubdistrict

`func (o *AadhaarClaims) GetSubdistrict() string`

GetSubdistrict returns the Subdistrict field if non-nil, zero value otherwise.

### GetSubdistrictOk

`func (o *AadhaarClaims) GetSubdistrictOk() (*string, bool)`

GetSubdistrictOk returns a tuple with the Subdistrict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdistrict

`func (o *AadhaarClaims) SetSubdistrict(v string)`

SetSubdistrict sets Subdistrict field to given value.

### HasSubdistrict

`func (o *AadhaarClaims) HasSubdistrict() bool`

HasSubdistrict returns a boolean if a field has been set.

### SetSubdistrictNil

`func (o *AadhaarClaims) SetSubdistrictNil(b bool)`

 SetSubdistrictNil sets the value for Subdistrict to be an explicit nil

### UnsetSubdistrict
`func (o *AadhaarClaims) UnsetSubdistrict()`

UnsetSubdistrict ensures that no value is present for Subdistrict, not even an explicit nil
### GetVillageTownCity

`func (o *AadhaarClaims) GetVillageTownCity() string`

GetVillageTownCity returns the VillageTownCity field if non-nil, zero value otherwise.

### GetVillageTownCityOk

`func (o *AadhaarClaims) GetVillageTownCityOk() (*string, bool)`

GetVillageTownCityOk returns a tuple with the VillageTownCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVillageTownCity

`func (o *AadhaarClaims) SetVillageTownCity(v string)`

SetVillageTownCity sets VillageTownCity field to given value.

### HasVillageTownCity

`func (o *AadhaarClaims) HasVillageTownCity() bool`

HasVillageTownCity returns a boolean if a field has been set.

### SetVillageTownCityNil

`func (o *AadhaarClaims) SetVillageTownCityNil(b bool)`

 SetVillageTownCityNil sets the value for VillageTownCity to be an explicit nil

### UnsetVillageTownCity
`func (o *AadhaarClaims) UnsetVillageTownCity()`

UnsetVillageTownCity ensures that no value is present for VillageTownCity, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


