# OutputFrejaAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Line1** | Pointer to **NullableString** | The line1 portion of the address. Incorporates all address details that are not included by city, postal code, and country. | [optional] 
**City** | Pointer to **NullableString** | The name of the city associated with the address. | [optional] 
**PostalCode** | Pointer to **NullableString** | The postal code associated with the address. | [optional] 
**Country** | Pointer to **NullableString** | The ISO 3166-1 alpha-2 country code associated with the address. | [optional] 

## Methods

### NewOutputFrejaAddress

`func NewOutputFrejaAddress() *OutputFrejaAddress`

NewOutputFrejaAddress instantiates a new OutputFrejaAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputFrejaAddressWithDefaults

`func NewOutputFrejaAddressWithDefaults() *OutputFrejaAddress`

NewOutputFrejaAddressWithDefaults instantiates a new OutputFrejaAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLine1

`func (o *OutputFrejaAddress) GetLine1() string`

GetLine1 returns the Line1 field if non-nil, zero value otherwise.

### GetLine1Ok

`func (o *OutputFrejaAddress) GetLine1Ok() (*string, bool)`

GetLine1Ok returns a tuple with the Line1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine1

`func (o *OutputFrejaAddress) SetLine1(v string)`

SetLine1 sets Line1 field to given value.

### HasLine1

`func (o *OutputFrejaAddress) HasLine1() bool`

HasLine1 returns a boolean if a field has been set.

### SetLine1Nil

`func (o *OutputFrejaAddress) SetLine1Nil(b bool)`

 SetLine1Nil sets the value for Line1 to be an explicit nil

### UnsetLine1
`func (o *OutputFrejaAddress) UnsetLine1()`

UnsetLine1 ensures that no value is present for Line1, not even an explicit nil
### GetCity

`func (o *OutputFrejaAddress) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *OutputFrejaAddress) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *OutputFrejaAddress) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *OutputFrejaAddress) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *OutputFrejaAddress) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *OutputFrejaAddress) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetPostalCode

`func (o *OutputFrejaAddress) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *OutputFrejaAddress) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *OutputFrejaAddress) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *OutputFrejaAddress) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### SetPostalCodeNil

`func (o *OutputFrejaAddress) SetPostalCodeNil(b bool)`

 SetPostalCodeNil sets the value for PostalCode to be an explicit nil

### UnsetPostalCode
`func (o *OutputFrejaAddress) UnsetPostalCode()`

UnsetPostalCode ensures that no value is present for PostalCode, not even an explicit nil
### GetCountry

`func (o *OutputFrejaAddress) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *OutputFrejaAddress) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *OutputFrejaAddress) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *OutputFrejaAddress) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *OutputFrejaAddress) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *OutputFrejaAddress) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


