# YotiStructuredPostalAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Country** | Pointer to **NullableString** | The address&#39;s country. This field does not have a specified format, but is recommended to use for the country field. | [optional] 
**CountryIso** | Pointer to **NullableString** | The address&#39;s ISO alpha-3 country code | [optional] 
**FormattedAddress** | Pointer to **NullableString** | The whole address in proper order and formatted with line breaks.              Examples: United Kingdom: \&quot;15a North Street\\nCARSHALTON\\nSM5 2HW\\nUK\&quot; India: \&quot;S/O: Vipen Kumar Aggarwal\\nHouse No.86-A\\nRajguru Nagar\\nLudhiana\\nPunjab\\n141012\\nIndia\&quot; United States of America: \&quot;1512 Ferry Street\\nAnniston AL 36201\\nUSA\&quot; | [optional] 
**Line1** | Pointer to **NullableString** | The first line of the address | [optional] 
**Line2** | Pointer to **NullableString** | The second line of the address | [optional] 
**Line3** | Pointer to **NullableString** | The third line of the address | [optional] 
**Line4** | Pointer to **NullableString** | The fourth line of the address | [optional] 
**TownCity** | Pointer to **NullableString** | The address&#39;s town or city | [optional] 
**State** | Pointer to **NullableString** | The address&#39;s state. This is an optional field that may only appear for US addresses. | [optional] 
**PostalCode** | Pointer to **NullableString** | The address&#39;s postal code | [optional] 
**BuildingNumber** | Pointer to **NullableString** | The address&#39;s building number | [optional] 
**Building** | Pointer to **NullableString** | The address&#39;s building name | [optional] 
**SubBuilding** | Pointer to **NullableString** | The address&#39;s sub building or specific part of the building | [optional] 
**Locality** | Pointer to **NullableString** | The address&#39;ss locality can be an area, a village, a region, or a known neighborhood | [optional] 
**DeliveryPointReferenceNumber** | Pointer to **NullableString** | The unique delivery point reference number. This is an 8-digit code used in the UK for identifying delivery addresses. | [optional] 
**Landmark** | Pointer to **NullableString** | The address&#39;s landmark. A landmark is a recognizable feature or place that helps locate the address. This optional field may appear only for Indian addresses. | [optional] 
**Subdistrict** | Pointer to **NullableString** | The address&#39;s subdistrict. This optional field may appear only for Indian addresses. | [optional] 
**District** | Pointer to **NullableString** | The address&#39;s district. This optional field may appear only for Indian addresses. | [optional] 
**PostOffice** | Pointer to **NullableString** | The address&#39;s post office. This optional field may appear only for Indian addresses. | [optional] 
**CareOf** | Pointer to **NullableString** | The address&#39;s care-of field. This indicates that mail should be delivered to an individual through another person or entity who is a known resident at the address. This optional field may appear only for Indian addresses. | [optional] 

## Methods

### NewYotiStructuredPostalAddress

`func NewYotiStructuredPostalAddress() *YotiStructuredPostalAddress`

NewYotiStructuredPostalAddress instantiates a new YotiStructuredPostalAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewYotiStructuredPostalAddressWithDefaults

`func NewYotiStructuredPostalAddressWithDefaults() *YotiStructuredPostalAddress`

NewYotiStructuredPostalAddressWithDefaults instantiates a new YotiStructuredPostalAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountry

`func (o *YotiStructuredPostalAddress) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *YotiStructuredPostalAddress) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *YotiStructuredPostalAddress) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *YotiStructuredPostalAddress) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *YotiStructuredPostalAddress) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *YotiStructuredPostalAddress) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetCountryIso

`func (o *YotiStructuredPostalAddress) GetCountryIso() string`

GetCountryIso returns the CountryIso field if non-nil, zero value otherwise.

### GetCountryIsoOk

`func (o *YotiStructuredPostalAddress) GetCountryIsoOk() (*string, bool)`

GetCountryIsoOk returns a tuple with the CountryIso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryIso

`func (o *YotiStructuredPostalAddress) SetCountryIso(v string)`

SetCountryIso sets CountryIso field to given value.

### HasCountryIso

`func (o *YotiStructuredPostalAddress) HasCountryIso() bool`

HasCountryIso returns a boolean if a field has been set.

### SetCountryIsoNil

`func (o *YotiStructuredPostalAddress) SetCountryIsoNil(b bool)`

 SetCountryIsoNil sets the value for CountryIso to be an explicit nil

### UnsetCountryIso
`func (o *YotiStructuredPostalAddress) UnsetCountryIso()`

UnsetCountryIso ensures that no value is present for CountryIso, not even an explicit nil
### GetFormattedAddress

`func (o *YotiStructuredPostalAddress) GetFormattedAddress() string`

GetFormattedAddress returns the FormattedAddress field if non-nil, zero value otherwise.

### GetFormattedAddressOk

`func (o *YotiStructuredPostalAddress) GetFormattedAddressOk() (*string, bool)`

GetFormattedAddressOk returns a tuple with the FormattedAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedAddress

`func (o *YotiStructuredPostalAddress) SetFormattedAddress(v string)`

SetFormattedAddress sets FormattedAddress field to given value.

### HasFormattedAddress

`func (o *YotiStructuredPostalAddress) HasFormattedAddress() bool`

HasFormattedAddress returns a boolean if a field has been set.

### SetFormattedAddressNil

`func (o *YotiStructuredPostalAddress) SetFormattedAddressNil(b bool)`

 SetFormattedAddressNil sets the value for FormattedAddress to be an explicit nil

### UnsetFormattedAddress
`func (o *YotiStructuredPostalAddress) UnsetFormattedAddress()`

UnsetFormattedAddress ensures that no value is present for FormattedAddress, not even an explicit nil
### GetLine1

`func (o *YotiStructuredPostalAddress) GetLine1() string`

GetLine1 returns the Line1 field if non-nil, zero value otherwise.

### GetLine1Ok

`func (o *YotiStructuredPostalAddress) GetLine1Ok() (*string, bool)`

GetLine1Ok returns a tuple with the Line1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine1

`func (o *YotiStructuredPostalAddress) SetLine1(v string)`

SetLine1 sets Line1 field to given value.

### HasLine1

`func (o *YotiStructuredPostalAddress) HasLine1() bool`

HasLine1 returns a boolean if a field has been set.

### SetLine1Nil

`func (o *YotiStructuredPostalAddress) SetLine1Nil(b bool)`

 SetLine1Nil sets the value for Line1 to be an explicit nil

### UnsetLine1
`func (o *YotiStructuredPostalAddress) UnsetLine1()`

UnsetLine1 ensures that no value is present for Line1, not even an explicit nil
### GetLine2

`func (o *YotiStructuredPostalAddress) GetLine2() string`

GetLine2 returns the Line2 field if non-nil, zero value otherwise.

### GetLine2Ok

`func (o *YotiStructuredPostalAddress) GetLine2Ok() (*string, bool)`

GetLine2Ok returns a tuple with the Line2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine2

`func (o *YotiStructuredPostalAddress) SetLine2(v string)`

SetLine2 sets Line2 field to given value.

### HasLine2

`func (o *YotiStructuredPostalAddress) HasLine2() bool`

HasLine2 returns a boolean if a field has been set.

### SetLine2Nil

`func (o *YotiStructuredPostalAddress) SetLine2Nil(b bool)`

 SetLine2Nil sets the value for Line2 to be an explicit nil

### UnsetLine2
`func (o *YotiStructuredPostalAddress) UnsetLine2()`

UnsetLine2 ensures that no value is present for Line2, not even an explicit nil
### GetLine3

`func (o *YotiStructuredPostalAddress) GetLine3() string`

GetLine3 returns the Line3 field if non-nil, zero value otherwise.

### GetLine3Ok

`func (o *YotiStructuredPostalAddress) GetLine3Ok() (*string, bool)`

GetLine3Ok returns a tuple with the Line3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine3

`func (o *YotiStructuredPostalAddress) SetLine3(v string)`

SetLine3 sets Line3 field to given value.

### HasLine3

`func (o *YotiStructuredPostalAddress) HasLine3() bool`

HasLine3 returns a boolean if a field has been set.

### SetLine3Nil

`func (o *YotiStructuredPostalAddress) SetLine3Nil(b bool)`

 SetLine3Nil sets the value for Line3 to be an explicit nil

### UnsetLine3
`func (o *YotiStructuredPostalAddress) UnsetLine3()`

UnsetLine3 ensures that no value is present for Line3, not even an explicit nil
### GetLine4

`func (o *YotiStructuredPostalAddress) GetLine4() string`

GetLine4 returns the Line4 field if non-nil, zero value otherwise.

### GetLine4Ok

`func (o *YotiStructuredPostalAddress) GetLine4Ok() (*string, bool)`

GetLine4Ok returns a tuple with the Line4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine4

`func (o *YotiStructuredPostalAddress) SetLine4(v string)`

SetLine4 sets Line4 field to given value.

### HasLine4

`func (o *YotiStructuredPostalAddress) HasLine4() bool`

HasLine4 returns a boolean if a field has been set.

### SetLine4Nil

`func (o *YotiStructuredPostalAddress) SetLine4Nil(b bool)`

 SetLine4Nil sets the value for Line4 to be an explicit nil

### UnsetLine4
`func (o *YotiStructuredPostalAddress) UnsetLine4()`

UnsetLine4 ensures that no value is present for Line4, not even an explicit nil
### GetTownCity

`func (o *YotiStructuredPostalAddress) GetTownCity() string`

GetTownCity returns the TownCity field if non-nil, zero value otherwise.

### GetTownCityOk

`func (o *YotiStructuredPostalAddress) GetTownCityOk() (*string, bool)`

GetTownCityOk returns a tuple with the TownCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTownCity

`func (o *YotiStructuredPostalAddress) SetTownCity(v string)`

SetTownCity sets TownCity field to given value.

### HasTownCity

`func (o *YotiStructuredPostalAddress) HasTownCity() bool`

HasTownCity returns a boolean if a field has been set.

### SetTownCityNil

`func (o *YotiStructuredPostalAddress) SetTownCityNil(b bool)`

 SetTownCityNil sets the value for TownCity to be an explicit nil

### UnsetTownCity
`func (o *YotiStructuredPostalAddress) UnsetTownCity()`

UnsetTownCity ensures that no value is present for TownCity, not even an explicit nil
### GetState

`func (o *YotiStructuredPostalAddress) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *YotiStructuredPostalAddress) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *YotiStructuredPostalAddress) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *YotiStructuredPostalAddress) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *YotiStructuredPostalAddress) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *YotiStructuredPostalAddress) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetPostalCode

`func (o *YotiStructuredPostalAddress) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *YotiStructuredPostalAddress) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *YotiStructuredPostalAddress) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *YotiStructuredPostalAddress) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### SetPostalCodeNil

`func (o *YotiStructuredPostalAddress) SetPostalCodeNil(b bool)`

 SetPostalCodeNil sets the value for PostalCode to be an explicit nil

### UnsetPostalCode
`func (o *YotiStructuredPostalAddress) UnsetPostalCode()`

UnsetPostalCode ensures that no value is present for PostalCode, not even an explicit nil
### GetBuildingNumber

`func (o *YotiStructuredPostalAddress) GetBuildingNumber() string`

GetBuildingNumber returns the BuildingNumber field if non-nil, zero value otherwise.

### GetBuildingNumberOk

`func (o *YotiStructuredPostalAddress) GetBuildingNumberOk() (*string, bool)`

GetBuildingNumberOk returns a tuple with the BuildingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildingNumber

`func (o *YotiStructuredPostalAddress) SetBuildingNumber(v string)`

SetBuildingNumber sets BuildingNumber field to given value.

### HasBuildingNumber

`func (o *YotiStructuredPostalAddress) HasBuildingNumber() bool`

HasBuildingNumber returns a boolean if a field has been set.

### SetBuildingNumberNil

`func (o *YotiStructuredPostalAddress) SetBuildingNumberNil(b bool)`

 SetBuildingNumberNil sets the value for BuildingNumber to be an explicit nil

### UnsetBuildingNumber
`func (o *YotiStructuredPostalAddress) UnsetBuildingNumber()`

UnsetBuildingNumber ensures that no value is present for BuildingNumber, not even an explicit nil
### GetBuilding

`func (o *YotiStructuredPostalAddress) GetBuilding() string`

GetBuilding returns the Building field if non-nil, zero value otherwise.

### GetBuildingOk

`func (o *YotiStructuredPostalAddress) GetBuildingOk() (*string, bool)`

GetBuildingOk returns a tuple with the Building field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuilding

`func (o *YotiStructuredPostalAddress) SetBuilding(v string)`

SetBuilding sets Building field to given value.

### HasBuilding

`func (o *YotiStructuredPostalAddress) HasBuilding() bool`

HasBuilding returns a boolean if a field has been set.

### SetBuildingNil

`func (o *YotiStructuredPostalAddress) SetBuildingNil(b bool)`

 SetBuildingNil sets the value for Building to be an explicit nil

### UnsetBuilding
`func (o *YotiStructuredPostalAddress) UnsetBuilding()`

UnsetBuilding ensures that no value is present for Building, not even an explicit nil
### GetSubBuilding

`func (o *YotiStructuredPostalAddress) GetSubBuilding() string`

GetSubBuilding returns the SubBuilding field if non-nil, zero value otherwise.

### GetSubBuildingOk

`func (o *YotiStructuredPostalAddress) GetSubBuildingOk() (*string, bool)`

GetSubBuildingOk returns a tuple with the SubBuilding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubBuilding

`func (o *YotiStructuredPostalAddress) SetSubBuilding(v string)`

SetSubBuilding sets SubBuilding field to given value.

### HasSubBuilding

`func (o *YotiStructuredPostalAddress) HasSubBuilding() bool`

HasSubBuilding returns a boolean if a field has been set.

### SetSubBuildingNil

`func (o *YotiStructuredPostalAddress) SetSubBuildingNil(b bool)`

 SetSubBuildingNil sets the value for SubBuilding to be an explicit nil

### UnsetSubBuilding
`func (o *YotiStructuredPostalAddress) UnsetSubBuilding()`

UnsetSubBuilding ensures that no value is present for SubBuilding, not even an explicit nil
### GetLocality

`func (o *YotiStructuredPostalAddress) GetLocality() string`

GetLocality returns the Locality field if non-nil, zero value otherwise.

### GetLocalityOk

`func (o *YotiStructuredPostalAddress) GetLocalityOk() (*string, bool)`

GetLocalityOk returns a tuple with the Locality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocality

`func (o *YotiStructuredPostalAddress) SetLocality(v string)`

SetLocality sets Locality field to given value.

### HasLocality

`func (o *YotiStructuredPostalAddress) HasLocality() bool`

HasLocality returns a boolean if a field has been set.

### SetLocalityNil

`func (o *YotiStructuredPostalAddress) SetLocalityNil(b bool)`

 SetLocalityNil sets the value for Locality to be an explicit nil

### UnsetLocality
`func (o *YotiStructuredPostalAddress) UnsetLocality()`

UnsetLocality ensures that no value is present for Locality, not even an explicit nil
### GetDeliveryPointReferenceNumber

`func (o *YotiStructuredPostalAddress) GetDeliveryPointReferenceNumber() string`

GetDeliveryPointReferenceNumber returns the DeliveryPointReferenceNumber field if non-nil, zero value otherwise.

### GetDeliveryPointReferenceNumberOk

`func (o *YotiStructuredPostalAddress) GetDeliveryPointReferenceNumberOk() (*string, bool)`

GetDeliveryPointReferenceNumberOk returns a tuple with the DeliveryPointReferenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryPointReferenceNumber

`func (o *YotiStructuredPostalAddress) SetDeliveryPointReferenceNumber(v string)`

SetDeliveryPointReferenceNumber sets DeliveryPointReferenceNumber field to given value.

### HasDeliveryPointReferenceNumber

`func (o *YotiStructuredPostalAddress) HasDeliveryPointReferenceNumber() bool`

HasDeliveryPointReferenceNumber returns a boolean if a field has been set.

### SetDeliveryPointReferenceNumberNil

`func (o *YotiStructuredPostalAddress) SetDeliveryPointReferenceNumberNil(b bool)`

 SetDeliveryPointReferenceNumberNil sets the value for DeliveryPointReferenceNumber to be an explicit nil

### UnsetDeliveryPointReferenceNumber
`func (o *YotiStructuredPostalAddress) UnsetDeliveryPointReferenceNumber()`

UnsetDeliveryPointReferenceNumber ensures that no value is present for DeliveryPointReferenceNumber, not even an explicit nil
### GetLandmark

`func (o *YotiStructuredPostalAddress) GetLandmark() string`

GetLandmark returns the Landmark field if non-nil, zero value otherwise.

### GetLandmarkOk

`func (o *YotiStructuredPostalAddress) GetLandmarkOk() (*string, bool)`

GetLandmarkOk returns a tuple with the Landmark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLandmark

`func (o *YotiStructuredPostalAddress) SetLandmark(v string)`

SetLandmark sets Landmark field to given value.

### HasLandmark

`func (o *YotiStructuredPostalAddress) HasLandmark() bool`

HasLandmark returns a boolean if a field has been set.

### SetLandmarkNil

`func (o *YotiStructuredPostalAddress) SetLandmarkNil(b bool)`

 SetLandmarkNil sets the value for Landmark to be an explicit nil

### UnsetLandmark
`func (o *YotiStructuredPostalAddress) UnsetLandmark()`

UnsetLandmark ensures that no value is present for Landmark, not even an explicit nil
### GetSubdistrict

`func (o *YotiStructuredPostalAddress) GetSubdistrict() string`

GetSubdistrict returns the Subdistrict field if non-nil, zero value otherwise.

### GetSubdistrictOk

`func (o *YotiStructuredPostalAddress) GetSubdistrictOk() (*string, bool)`

GetSubdistrictOk returns a tuple with the Subdistrict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdistrict

`func (o *YotiStructuredPostalAddress) SetSubdistrict(v string)`

SetSubdistrict sets Subdistrict field to given value.

### HasSubdistrict

`func (o *YotiStructuredPostalAddress) HasSubdistrict() bool`

HasSubdistrict returns a boolean if a field has been set.

### SetSubdistrictNil

`func (o *YotiStructuredPostalAddress) SetSubdistrictNil(b bool)`

 SetSubdistrictNil sets the value for Subdistrict to be an explicit nil

### UnsetSubdistrict
`func (o *YotiStructuredPostalAddress) UnsetSubdistrict()`

UnsetSubdistrict ensures that no value is present for Subdistrict, not even an explicit nil
### GetDistrict

`func (o *YotiStructuredPostalAddress) GetDistrict() string`

GetDistrict returns the District field if non-nil, zero value otherwise.

### GetDistrictOk

`func (o *YotiStructuredPostalAddress) GetDistrictOk() (*string, bool)`

GetDistrictOk returns a tuple with the District field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistrict

`func (o *YotiStructuredPostalAddress) SetDistrict(v string)`

SetDistrict sets District field to given value.

### HasDistrict

`func (o *YotiStructuredPostalAddress) HasDistrict() bool`

HasDistrict returns a boolean if a field has been set.

### SetDistrictNil

`func (o *YotiStructuredPostalAddress) SetDistrictNil(b bool)`

 SetDistrictNil sets the value for District to be an explicit nil

### UnsetDistrict
`func (o *YotiStructuredPostalAddress) UnsetDistrict()`

UnsetDistrict ensures that no value is present for District, not even an explicit nil
### GetPostOffice

`func (o *YotiStructuredPostalAddress) GetPostOffice() string`

GetPostOffice returns the PostOffice field if non-nil, zero value otherwise.

### GetPostOfficeOk

`func (o *YotiStructuredPostalAddress) GetPostOfficeOk() (*string, bool)`

GetPostOfficeOk returns a tuple with the PostOffice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostOffice

`func (o *YotiStructuredPostalAddress) SetPostOffice(v string)`

SetPostOffice sets PostOffice field to given value.

### HasPostOffice

`func (o *YotiStructuredPostalAddress) HasPostOffice() bool`

HasPostOffice returns a boolean if a field has been set.

### SetPostOfficeNil

`func (o *YotiStructuredPostalAddress) SetPostOfficeNil(b bool)`

 SetPostOfficeNil sets the value for PostOffice to be an explicit nil

### UnsetPostOffice
`func (o *YotiStructuredPostalAddress) UnsetPostOffice()`

UnsetPostOffice ensures that no value is present for PostOffice, not even an explicit nil
### GetCareOf

`func (o *YotiStructuredPostalAddress) GetCareOf() string`

GetCareOf returns the CareOf field if non-nil, zero value otherwise.

### GetCareOfOk

`func (o *YotiStructuredPostalAddress) GetCareOfOk() (*string, bool)`

GetCareOfOk returns a tuple with the CareOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCareOf

`func (o *YotiStructuredPostalAddress) SetCareOf(v string)`

SetCareOf sets CareOf field to given value.

### HasCareOf

`func (o *YotiStructuredPostalAddress) HasCareOf() bool`

HasCareOf returns a boolean if a field has been set.

### SetCareOfNil

`func (o *YotiStructuredPostalAddress) SetCareOfNil(b bool)`

 SetCareOfNil sets the value for CareOf to be an explicit nil

### UnsetCareOf
`func (o *YotiStructuredPostalAddress) UnsetCareOf()`

UnsetCareOf ensures that no value is present for CareOf, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


