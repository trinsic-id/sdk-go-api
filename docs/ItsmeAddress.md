# ItsmeAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StreetAddress** | Pointer to **NullableString** | The full street address on a single line. | [optional] 
**Street** | Pointer to **NullableString** | The street name component of the address. | [optional] 
**HouseNumber** | Pointer to **NullableString** | The house number component of the address. | [optional] 
**PostalCode** | Pointer to **NullableString** | The postal code of the address. | [optional] 
**City** | Pointer to **NullableString** | The city of the address. | [optional] 
**Country** | Pointer to **NullableString** | The country of the address as an ISO 3166-1 alpha-2 code. | [optional] 

## Methods

### NewItsmeAddress

`func NewItsmeAddress() *ItsmeAddress`

NewItsmeAddress instantiates a new ItsmeAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItsmeAddressWithDefaults

`func NewItsmeAddressWithDefaults() *ItsmeAddress`

NewItsmeAddressWithDefaults instantiates a new ItsmeAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStreetAddress

`func (o *ItsmeAddress) GetStreetAddress() string`

GetStreetAddress returns the StreetAddress field if non-nil, zero value otherwise.

### GetStreetAddressOk

`func (o *ItsmeAddress) GetStreetAddressOk() (*string, bool)`

GetStreetAddressOk returns a tuple with the StreetAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetAddress

`func (o *ItsmeAddress) SetStreetAddress(v string)`

SetStreetAddress sets StreetAddress field to given value.

### HasStreetAddress

`func (o *ItsmeAddress) HasStreetAddress() bool`

HasStreetAddress returns a boolean if a field has been set.

### SetStreetAddressNil

`func (o *ItsmeAddress) SetStreetAddressNil(b bool)`

 SetStreetAddressNil sets the value for StreetAddress to be an explicit nil

### UnsetStreetAddress
`func (o *ItsmeAddress) UnsetStreetAddress()`

UnsetStreetAddress ensures that no value is present for StreetAddress, not even an explicit nil
### GetStreet

`func (o *ItsmeAddress) GetStreet() string`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *ItsmeAddress) GetStreetOk() (*string, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet

`func (o *ItsmeAddress) SetStreet(v string)`

SetStreet sets Street field to given value.

### HasStreet

`func (o *ItsmeAddress) HasStreet() bool`

HasStreet returns a boolean if a field has been set.

### SetStreetNil

`func (o *ItsmeAddress) SetStreetNil(b bool)`

 SetStreetNil sets the value for Street to be an explicit nil

### UnsetStreet
`func (o *ItsmeAddress) UnsetStreet()`

UnsetStreet ensures that no value is present for Street, not even an explicit nil
### GetHouseNumber

`func (o *ItsmeAddress) GetHouseNumber() string`

GetHouseNumber returns the HouseNumber field if non-nil, zero value otherwise.

### GetHouseNumberOk

`func (o *ItsmeAddress) GetHouseNumberOk() (*string, bool)`

GetHouseNumberOk returns a tuple with the HouseNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHouseNumber

`func (o *ItsmeAddress) SetHouseNumber(v string)`

SetHouseNumber sets HouseNumber field to given value.

### HasHouseNumber

`func (o *ItsmeAddress) HasHouseNumber() bool`

HasHouseNumber returns a boolean if a field has been set.

### SetHouseNumberNil

`func (o *ItsmeAddress) SetHouseNumberNil(b bool)`

 SetHouseNumberNil sets the value for HouseNumber to be an explicit nil

### UnsetHouseNumber
`func (o *ItsmeAddress) UnsetHouseNumber()`

UnsetHouseNumber ensures that no value is present for HouseNumber, not even an explicit nil
### GetPostalCode

`func (o *ItsmeAddress) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *ItsmeAddress) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *ItsmeAddress) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *ItsmeAddress) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### SetPostalCodeNil

`func (o *ItsmeAddress) SetPostalCodeNil(b bool)`

 SetPostalCodeNil sets the value for PostalCode to be an explicit nil

### UnsetPostalCode
`func (o *ItsmeAddress) UnsetPostalCode()`

UnsetPostalCode ensures that no value is present for PostalCode, not even an explicit nil
### GetCity

`func (o *ItsmeAddress) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *ItsmeAddress) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *ItsmeAddress) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *ItsmeAddress) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *ItsmeAddress) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *ItsmeAddress) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetCountry

`func (o *ItsmeAddress) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *ItsmeAddress) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *ItsmeAddress) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *ItsmeAddress) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *ItsmeAddress) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *ItsmeAddress) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


