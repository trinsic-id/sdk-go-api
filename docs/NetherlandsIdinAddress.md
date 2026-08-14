# NetherlandsIdinAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Street** | Pointer to **NullableString** | Street name of the individual&#39;s domestic address. | [optional] 
**HouseNumber** | Pointer to **NullableString** | House number of the individual&#39;s domestic address. | [optional] 
**HouseNumberSuffix** | Pointer to **NullableString** | House number suffix of the individual&#39;s domestic address. | [optional] 
**AddressExtra** | Pointer to **NullableString** | Additional domestic address information returned by iDIN. | [optional] 
**PostalCode** | Pointer to **NullableString** | Postal code of the individual&#39;s domestic address.              A postal code in The Netherlands is a 4-digit number and two letters. | [optional] 
**City** | Pointer to **NullableString** | City of the individual&#39;s domestic address. | [optional] 
**CountryCode** | Pointer to **NullableString** | Country code of the individual&#39;s address in ISO 3166-1 alpha-2 format. | [optional] 
**InternationalAddressLine1** | Pointer to **NullableString** | First line of the individual&#39;s international address. | [optional] 
**InternationalAddressLine2** | Pointer to **NullableString** | Second line of the individual&#39;s international address. | [optional] 
**InternationalAddressLine3** | Pointer to **NullableString** | Third line of the individual&#39;s international address. | [optional] 

## Methods

### NewNetherlandsIdinAddress

`func NewNetherlandsIdinAddress() *NetherlandsIdinAddress`

NewNetherlandsIdinAddress instantiates a new NetherlandsIdinAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetherlandsIdinAddressWithDefaults

`func NewNetherlandsIdinAddressWithDefaults() *NetherlandsIdinAddress`

NewNetherlandsIdinAddressWithDefaults instantiates a new NetherlandsIdinAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStreet

`func (o *NetherlandsIdinAddress) GetStreet() string`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *NetherlandsIdinAddress) GetStreetOk() (*string, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet

`func (o *NetherlandsIdinAddress) SetStreet(v string)`

SetStreet sets Street field to given value.

### HasStreet

`func (o *NetherlandsIdinAddress) HasStreet() bool`

HasStreet returns a boolean if a field has been set.

### SetStreetNil

`func (o *NetherlandsIdinAddress) SetStreetNil(b bool)`

 SetStreetNil sets the value for Street to be an explicit nil

### UnsetStreet
`func (o *NetherlandsIdinAddress) UnsetStreet()`

UnsetStreet ensures that no value is present for Street, not even an explicit nil
### GetHouseNumber

`func (o *NetherlandsIdinAddress) GetHouseNumber() string`

GetHouseNumber returns the HouseNumber field if non-nil, zero value otherwise.

### GetHouseNumberOk

`func (o *NetherlandsIdinAddress) GetHouseNumberOk() (*string, bool)`

GetHouseNumberOk returns a tuple with the HouseNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHouseNumber

`func (o *NetherlandsIdinAddress) SetHouseNumber(v string)`

SetHouseNumber sets HouseNumber field to given value.

### HasHouseNumber

`func (o *NetherlandsIdinAddress) HasHouseNumber() bool`

HasHouseNumber returns a boolean if a field has been set.

### SetHouseNumberNil

`func (o *NetherlandsIdinAddress) SetHouseNumberNil(b bool)`

 SetHouseNumberNil sets the value for HouseNumber to be an explicit nil

### UnsetHouseNumber
`func (o *NetherlandsIdinAddress) UnsetHouseNumber()`

UnsetHouseNumber ensures that no value is present for HouseNumber, not even an explicit nil
### GetHouseNumberSuffix

`func (o *NetherlandsIdinAddress) GetHouseNumberSuffix() string`

GetHouseNumberSuffix returns the HouseNumberSuffix field if non-nil, zero value otherwise.

### GetHouseNumberSuffixOk

`func (o *NetherlandsIdinAddress) GetHouseNumberSuffixOk() (*string, bool)`

GetHouseNumberSuffixOk returns a tuple with the HouseNumberSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHouseNumberSuffix

`func (o *NetherlandsIdinAddress) SetHouseNumberSuffix(v string)`

SetHouseNumberSuffix sets HouseNumberSuffix field to given value.

### HasHouseNumberSuffix

`func (o *NetherlandsIdinAddress) HasHouseNumberSuffix() bool`

HasHouseNumberSuffix returns a boolean if a field has been set.

### SetHouseNumberSuffixNil

`func (o *NetherlandsIdinAddress) SetHouseNumberSuffixNil(b bool)`

 SetHouseNumberSuffixNil sets the value for HouseNumberSuffix to be an explicit nil

### UnsetHouseNumberSuffix
`func (o *NetherlandsIdinAddress) UnsetHouseNumberSuffix()`

UnsetHouseNumberSuffix ensures that no value is present for HouseNumberSuffix, not even an explicit nil
### GetAddressExtra

`func (o *NetherlandsIdinAddress) GetAddressExtra() string`

GetAddressExtra returns the AddressExtra field if non-nil, zero value otherwise.

### GetAddressExtraOk

`func (o *NetherlandsIdinAddress) GetAddressExtraOk() (*string, bool)`

GetAddressExtraOk returns a tuple with the AddressExtra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressExtra

`func (o *NetherlandsIdinAddress) SetAddressExtra(v string)`

SetAddressExtra sets AddressExtra field to given value.

### HasAddressExtra

`func (o *NetherlandsIdinAddress) HasAddressExtra() bool`

HasAddressExtra returns a boolean if a field has been set.

### SetAddressExtraNil

`func (o *NetherlandsIdinAddress) SetAddressExtraNil(b bool)`

 SetAddressExtraNil sets the value for AddressExtra to be an explicit nil

### UnsetAddressExtra
`func (o *NetherlandsIdinAddress) UnsetAddressExtra()`

UnsetAddressExtra ensures that no value is present for AddressExtra, not even an explicit nil
### GetPostalCode

`func (o *NetherlandsIdinAddress) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *NetherlandsIdinAddress) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *NetherlandsIdinAddress) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *NetherlandsIdinAddress) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### SetPostalCodeNil

`func (o *NetherlandsIdinAddress) SetPostalCodeNil(b bool)`

 SetPostalCodeNil sets the value for PostalCode to be an explicit nil

### UnsetPostalCode
`func (o *NetherlandsIdinAddress) UnsetPostalCode()`

UnsetPostalCode ensures that no value is present for PostalCode, not even an explicit nil
### GetCity

`func (o *NetherlandsIdinAddress) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *NetherlandsIdinAddress) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *NetherlandsIdinAddress) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *NetherlandsIdinAddress) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *NetherlandsIdinAddress) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *NetherlandsIdinAddress) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetCountryCode

`func (o *NetherlandsIdinAddress) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *NetherlandsIdinAddress) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *NetherlandsIdinAddress) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *NetherlandsIdinAddress) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### SetCountryCodeNil

`func (o *NetherlandsIdinAddress) SetCountryCodeNil(b bool)`

 SetCountryCodeNil sets the value for CountryCode to be an explicit nil

### UnsetCountryCode
`func (o *NetherlandsIdinAddress) UnsetCountryCode()`

UnsetCountryCode ensures that no value is present for CountryCode, not even an explicit nil
### GetInternationalAddressLine1

`func (o *NetherlandsIdinAddress) GetInternationalAddressLine1() string`

GetInternationalAddressLine1 returns the InternationalAddressLine1 field if non-nil, zero value otherwise.

### GetInternationalAddressLine1Ok

`func (o *NetherlandsIdinAddress) GetInternationalAddressLine1Ok() (*string, bool)`

GetInternationalAddressLine1Ok returns a tuple with the InternationalAddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternationalAddressLine1

`func (o *NetherlandsIdinAddress) SetInternationalAddressLine1(v string)`

SetInternationalAddressLine1 sets InternationalAddressLine1 field to given value.

### HasInternationalAddressLine1

`func (o *NetherlandsIdinAddress) HasInternationalAddressLine1() bool`

HasInternationalAddressLine1 returns a boolean if a field has been set.

### SetInternationalAddressLine1Nil

`func (o *NetherlandsIdinAddress) SetInternationalAddressLine1Nil(b bool)`

 SetInternationalAddressLine1Nil sets the value for InternationalAddressLine1 to be an explicit nil

### UnsetInternationalAddressLine1
`func (o *NetherlandsIdinAddress) UnsetInternationalAddressLine1()`

UnsetInternationalAddressLine1 ensures that no value is present for InternationalAddressLine1, not even an explicit nil
### GetInternationalAddressLine2

`func (o *NetherlandsIdinAddress) GetInternationalAddressLine2() string`

GetInternationalAddressLine2 returns the InternationalAddressLine2 field if non-nil, zero value otherwise.

### GetInternationalAddressLine2Ok

`func (o *NetherlandsIdinAddress) GetInternationalAddressLine2Ok() (*string, bool)`

GetInternationalAddressLine2Ok returns a tuple with the InternationalAddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternationalAddressLine2

`func (o *NetherlandsIdinAddress) SetInternationalAddressLine2(v string)`

SetInternationalAddressLine2 sets InternationalAddressLine2 field to given value.

### HasInternationalAddressLine2

`func (o *NetherlandsIdinAddress) HasInternationalAddressLine2() bool`

HasInternationalAddressLine2 returns a boolean if a field has been set.

### SetInternationalAddressLine2Nil

`func (o *NetherlandsIdinAddress) SetInternationalAddressLine2Nil(b bool)`

 SetInternationalAddressLine2Nil sets the value for InternationalAddressLine2 to be an explicit nil

### UnsetInternationalAddressLine2
`func (o *NetherlandsIdinAddress) UnsetInternationalAddressLine2()`

UnsetInternationalAddressLine2 ensures that no value is present for InternationalAddressLine2, not even an explicit nil
### GetInternationalAddressLine3

`func (o *NetherlandsIdinAddress) GetInternationalAddressLine3() string`

GetInternationalAddressLine3 returns the InternationalAddressLine3 field if non-nil, zero value otherwise.

### GetInternationalAddressLine3Ok

`func (o *NetherlandsIdinAddress) GetInternationalAddressLine3Ok() (*string, bool)`

GetInternationalAddressLine3Ok returns a tuple with the InternationalAddressLine3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternationalAddressLine3

`func (o *NetherlandsIdinAddress) SetInternationalAddressLine3(v string)`

SetInternationalAddressLine3 sets InternationalAddressLine3 field to given value.

### HasInternationalAddressLine3

`func (o *NetherlandsIdinAddress) HasInternationalAddressLine3() bool`

HasInternationalAddressLine3 returns a boolean if a field has been set.

### SetInternationalAddressLine3Nil

`func (o *NetherlandsIdinAddress) SetInternationalAddressLine3Nil(b bool)`

 SetInternationalAddressLine3Nil sets the value for InternationalAddressLine3 to be an explicit nil

### UnsetInternationalAddressLine3
`func (o *NetherlandsIdinAddress) UnsetInternationalAddressLine3()`

UnsetInternationalAddressLine3 ensures that no value is present for InternationalAddressLine3, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


