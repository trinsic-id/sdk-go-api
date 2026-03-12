# CzechMojeIdAddressOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StreetAddress** | Pointer to **NullableString** | The street address line, including house number. | [optional] 
**City** | Pointer to **NullableString** | The city or locality. | [optional] 
**Region** | Pointer to **NullableString** | The region or state. | [optional] 
**PostalCode** | Pointer to **NullableString** | The postal code. | [optional] 
**Country** | Pointer to **NullableString** | The ISO 3166-1 alpha-2 country code. | [optional] 
**FormattedAddress** | Pointer to **NullableString** | The full address as a formatted string, with components separated by newlines. | [optional] 

## Methods

### NewCzechMojeIdAddressOutput

`func NewCzechMojeIdAddressOutput() *CzechMojeIdAddressOutput`

NewCzechMojeIdAddressOutput instantiates a new CzechMojeIdAddressOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCzechMojeIdAddressOutputWithDefaults

`func NewCzechMojeIdAddressOutputWithDefaults() *CzechMojeIdAddressOutput`

NewCzechMojeIdAddressOutputWithDefaults instantiates a new CzechMojeIdAddressOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStreetAddress

`func (o *CzechMojeIdAddressOutput) GetStreetAddress() string`

GetStreetAddress returns the StreetAddress field if non-nil, zero value otherwise.

### GetStreetAddressOk

`func (o *CzechMojeIdAddressOutput) GetStreetAddressOk() (*string, bool)`

GetStreetAddressOk returns a tuple with the StreetAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetAddress

`func (o *CzechMojeIdAddressOutput) SetStreetAddress(v string)`

SetStreetAddress sets StreetAddress field to given value.

### HasStreetAddress

`func (o *CzechMojeIdAddressOutput) HasStreetAddress() bool`

HasStreetAddress returns a boolean if a field has been set.

### SetStreetAddressNil

`func (o *CzechMojeIdAddressOutput) SetStreetAddressNil(b bool)`

 SetStreetAddressNil sets the value for StreetAddress to be an explicit nil

### UnsetStreetAddress
`func (o *CzechMojeIdAddressOutput) UnsetStreetAddress()`

UnsetStreetAddress ensures that no value is present for StreetAddress, not even an explicit nil
### GetCity

`func (o *CzechMojeIdAddressOutput) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *CzechMojeIdAddressOutput) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *CzechMojeIdAddressOutput) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *CzechMojeIdAddressOutput) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *CzechMojeIdAddressOutput) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *CzechMojeIdAddressOutput) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetRegion

`func (o *CzechMojeIdAddressOutput) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CzechMojeIdAddressOutput) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CzechMojeIdAddressOutput) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *CzechMojeIdAddressOutput) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *CzechMojeIdAddressOutput) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *CzechMojeIdAddressOutput) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetPostalCode

`func (o *CzechMojeIdAddressOutput) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *CzechMojeIdAddressOutput) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *CzechMojeIdAddressOutput) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *CzechMojeIdAddressOutput) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### SetPostalCodeNil

`func (o *CzechMojeIdAddressOutput) SetPostalCodeNil(b bool)`

 SetPostalCodeNil sets the value for PostalCode to be an explicit nil

### UnsetPostalCode
`func (o *CzechMojeIdAddressOutput) UnsetPostalCode()`

UnsetPostalCode ensures that no value is present for PostalCode, not even an explicit nil
### GetCountry

`func (o *CzechMojeIdAddressOutput) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *CzechMojeIdAddressOutput) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *CzechMojeIdAddressOutput) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *CzechMojeIdAddressOutput) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *CzechMojeIdAddressOutput) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *CzechMojeIdAddressOutput) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetFormattedAddress

`func (o *CzechMojeIdAddressOutput) GetFormattedAddress() string`

GetFormattedAddress returns the FormattedAddress field if non-nil, zero value otherwise.

### GetFormattedAddressOk

`func (o *CzechMojeIdAddressOutput) GetFormattedAddressOk() (*string, bool)`

GetFormattedAddressOk returns a tuple with the FormattedAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedAddress

`func (o *CzechMojeIdAddressOutput) SetFormattedAddress(v string)`

SetFormattedAddress sets FormattedAddress field to given value.

### HasFormattedAddress

`func (o *CzechMojeIdAddressOutput) HasFormattedAddress() bool`

HasFormattedAddress returns a boolean if a field has been set.

### SetFormattedAddressNil

`func (o *CzechMojeIdAddressOutput) SetFormattedAddressNil(b bool)`

 SetFormattedAddressNil sets the value for FormattedAddress to be an explicit nil

### UnsetFormattedAddress
`func (o *CzechMojeIdAddressOutput) UnsetFormattedAddress()`

UnsetFormattedAddress ensures that no value is present for FormattedAddress, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


