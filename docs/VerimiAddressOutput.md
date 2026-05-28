# VerimiAddressOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FullAddress** | Pointer to **NullableString** | The formatted address. | [optional] 
**CountryCode** | Pointer to **NullableString** | The ISO 3166-1 alpha-2 country code. | [optional] 
**PostalCode** | Pointer to **NullableString** | The postal code. | [optional] 
**Street** | Pointer to **NullableString** | The street name. | [optional] 
**BuildingNumber** | Pointer to **NullableString** | The building number. | [optional] 
**City** | Pointer to **NullableString** | The city. | [optional] 

## Methods

### NewVerimiAddressOutput

`func NewVerimiAddressOutput() *VerimiAddressOutput`

NewVerimiAddressOutput instantiates a new VerimiAddressOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerimiAddressOutputWithDefaults

`func NewVerimiAddressOutputWithDefaults() *VerimiAddressOutput`

NewVerimiAddressOutputWithDefaults instantiates a new VerimiAddressOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFullAddress

`func (o *VerimiAddressOutput) GetFullAddress() string`

GetFullAddress returns the FullAddress field if non-nil, zero value otherwise.

### GetFullAddressOk

`func (o *VerimiAddressOutput) GetFullAddressOk() (*string, bool)`

GetFullAddressOk returns a tuple with the FullAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullAddress

`func (o *VerimiAddressOutput) SetFullAddress(v string)`

SetFullAddress sets FullAddress field to given value.

### HasFullAddress

`func (o *VerimiAddressOutput) HasFullAddress() bool`

HasFullAddress returns a boolean if a field has been set.

### SetFullAddressNil

`func (o *VerimiAddressOutput) SetFullAddressNil(b bool)`

 SetFullAddressNil sets the value for FullAddress to be an explicit nil

### UnsetFullAddress
`func (o *VerimiAddressOutput) UnsetFullAddress()`

UnsetFullAddress ensures that no value is present for FullAddress, not even an explicit nil
### GetCountryCode

`func (o *VerimiAddressOutput) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *VerimiAddressOutput) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *VerimiAddressOutput) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *VerimiAddressOutput) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### SetCountryCodeNil

`func (o *VerimiAddressOutput) SetCountryCodeNil(b bool)`

 SetCountryCodeNil sets the value for CountryCode to be an explicit nil

### UnsetCountryCode
`func (o *VerimiAddressOutput) UnsetCountryCode()`

UnsetCountryCode ensures that no value is present for CountryCode, not even an explicit nil
### GetPostalCode

`func (o *VerimiAddressOutput) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *VerimiAddressOutput) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *VerimiAddressOutput) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *VerimiAddressOutput) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### SetPostalCodeNil

`func (o *VerimiAddressOutput) SetPostalCodeNil(b bool)`

 SetPostalCodeNil sets the value for PostalCode to be an explicit nil

### UnsetPostalCode
`func (o *VerimiAddressOutput) UnsetPostalCode()`

UnsetPostalCode ensures that no value is present for PostalCode, not even an explicit nil
### GetStreet

`func (o *VerimiAddressOutput) GetStreet() string`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *VerimiAddressOutput) GetStreetOk() (*string, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet

`func (o *VerimiAddressOutput) SetStreet(v string)`

SetStreet sets Street field to given value.

### HasStreet

`func (o *VerimiAddressOutput) HasStreet() bool`

HasStreet returns a boolean if a field has been set.

### SetStreetNil

`func (o *VerimiAddressOutput) SetStreetNil(b bool)`

 SetStreetNil sets the value for Street to be an explicit nil

### UnsetStreet
`func (o *VerimiAddressOutput) UnsetStreet()`

UnsetStreet ensures that no value is present for Street, not even an explicit nil
### GetBuildingNumber

`func (o *VerimiAddressOutput) GetBuildingNumber() string`

GetBuildingNumber returns the BuildingNumber field if non-nil, zero value otherwise.

### GetBuildingNumberOk

`func (o *VerimiAddressOutput) GetBuildingNumberOk() (*string, bool)`

GetBuildingNumberOk returns a tuple with the BuildingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildingNumber

`func (o *VerimiAddressOutput) SetBuildingNumber(v string)`

SetBuildingNumber sets BuildingNumber field to given value.

### HasBuildingNumber

`func (o *VerimiAddressOutput) HasBuildingNumber() bool`

HasBuildingNumber returns a boolean if a field has been set.

### SetBuildingNumberNil

`func (o *VerimiAddressOutput) SetBuildingNumberNil(b bool)`

 SetBuildingNumberNil sets the value for BuildingNumber to be an explicit nil

### UnsetBuildingNumber
`func (o *VerimiAddressOutput) UnsetBuildingNumber()`

UnsetBuildingNumber ensures that no value is present for BuildingNumber, not even an explicit nil
### GetCity

`func (o *VerimiAddressOutput) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *VerimiAddressOutput) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *VerimiAddressOutput) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *VerimiAddressOutput) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *VerimiAddressOutput) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *VerimiAddressOutput) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


