# MitIdAddressOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CityName** | Pointer to **NullableString** | The city name returned by MitID. | [optional] 
**ZipName** | Pointer to **NullableString** | The postal district name returned by MitID. | [optional] 
**StreetNumber** | Pointer to **NullableString** | The street number. | [optional] 
**StreetName** | Pointer to **NullableString** | The street name. | [optional] 
**ZipCode** | Pointer to **NullableString** | The zip code. | [optional] 
**CountryCode** | Pointer to **NullableString** | The ISO 3166-1 alpha-2 country code, set to DK (Denmark). | [optional] 

## Methods

### Named construction

Use a keyed struct literal (`MitIdAddressOutput{FieldName: value}`), or call
`NewMitIdAddressOutputWithDefaults()` and set fields by name. Positional field
factories and unkeyed struct literals are not supported.

### NewMitIdAddressOutputWithDefaults

`func NewMitIdAddressOutputWithDefaults() *MitIdAddressOutput`

NewMitIdAddressOutputWithDefaults instantiates a new MitIdAddressOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCityName

`func (o *MitIdAddressOutput) GetCityName() string`

GetCityName returns the CityName field if non-nil, zero value otherwise.

### GetCityNameOk

`func (o *MitIdAddressOutput) GetCityNameOk() (*string, bool)`

GetCityNameOk returns a tuple with the CityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityName

`func (o *MitIdAddressOutput) SetCityName(v string)`

SetCityName sets CityName field to given value.

### HasCityName

`func (o *MitIdAddressOutput) HasCityName() bool`

HasCityName returns a boolean if a field has been set.

### SetCityNameNil

`func (o *MitIdAddressOutput) SetCityNameNil(b bool)`

 SetCityNameNil sets the value for CityName to be an explicit nil

### UnsetCityName
`func (o *MitIdAddressOutput) UnsetCityName()`

UnsetCityName ensures that no value is present for CityName, not even an explicit nil
### GetZipName

`func (o *MitIdAddressOutput) GetZipName() string`

GetZipName returns the ZipName field if non-nil, zero value otherwise.

### GetZipNameOk

`func (o *MitIdAddressOutput) GetZipNameOk() (*string, bool)`

GetZipNameOk returns a tuple with the ZipName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipName

`func (o *MitIdAddressOutput) SetZipName(v string)`

SetZipName sets ZipName field to given value.

### HasZipName

`func (o *MitIdAddressOutput) HasZipName() bool`

HasZipName returns a boolean if a field has been set.

### SetZipNameNil

`func (o *MitIdAddressOutput) SetZipNameNil(b bool)`

 SetZipNameNil sets the value for ZipName to be an explicit nil

### UnsetZipName
`func (o *MitIdAddressOutput) UnsetZipName()`

UnsetZipName ensures that no value is present for ZipName, not even an explicit nil
### GetStreetNumber

`func (o *MitIdAddressOutput) GetStreetNumber() string`

GetStreetNumber returns the StreetNumber field if non-nil, zero value otherwise.

### GetStreetNumberOk

`func (o *MitIdAddressOutput) GetStreetNumberOk() (*string, bool)`

GetStreetNumberOk returns a tuple with the StreetNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetNumber

`func (o *MitIdAddressOutput) SetStreetNumber(v string)`

SetStreetNumber sets StreetNumber field to given value.

### HasStreetNumber

`func (o *MitIdAddressOutput) HasStreetNumber() bool`

HasStreetNumber returns a boolean if a field has been set.

### SetStreetNumberNil

`func (o *MitIdAddressOutput) SetStreetNumberNil(b bool)`

 SetStreetNumberNil sets the value for StreetNumber to be an explicit nil

### UnsetStreetNumber
`func (o *MitIdAddressOutput) UnsetStreetNumber()`

UnsetStreetNumber ensures that no value is present for StreetNumber, not even an explicit nil
### GetStreetName

`func (o *MitIdAddressOutput) GetStreetName() string`

GetStreetName returns the StreetName field if non-nil, zero value otherwise.

### GetStreetNameOk

`func (o *MitIdAddressOutput) GetStreetNameOk() (*string, bool)`

GetStreetNameOk returns a tuple with the StreetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetName

`func (o *MitIdAddressOutput) SetStreetName(v string)`

SetStreetName sets StreetName field to given value.

### HasStreetName

`func (o *MitIdAddressOutput) HasStreetName() bool`

HasStreetName returns a boolean if a field has been set.

### SetStreetNameNil

`func (o *MitIdAddressOutput) SetStreetNameNil(b bool)`

 SetStreetNameNil sets the value for StreetName to be an explicit nil

### UnsetStreetName
`func (o *MitIdAddressOutput) UnsetStreetName()`

UnsetStreetName ensures that no value is present for StreetName, not even an explicit nil
### GetZipCode

`func (o *MitIdAddressOutput) GetZipCode() string`

GetZipCode returns the ZipCode field if non-nil, zero value otherwise.

### GetZipCodeOk

`func (o *MitIdAddressOutput) GetZipCodeOk() (*string, bool)`

GetZipCodeOk returns a tuple with the ZipCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCode

`func (o *MitIdAddressOutput) SetZipCode(v string)`

SetZipCode sets ZipCode field to given value.

### HasZipCode

`func (o *MitIdAddressOutput) HasZipCode() bool`

HasZipCode returns a boolean if a field has been set.

### SetZipCodeNil

`func (o *MitIdAddressOutput) SetZipCodeNil(b bool)`

 SetZipCodeNil sets the value for ZipCode to be an explicit nil

### UnsetZipCode
`func (o *MitIdAddressOutput) UnsetZipCode()`

UnsetZipCode ensures that no value is present for ZipCode, not even an explicit nil
### GetCountryCode

`func (o *MitIdAddressOutput) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *MitIdAddressOutput) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *MitIdAddressOutput) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *MitIdAddressOutput) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### SetCountryCodeNil

`func (o *MitIdAddressOutput) SetCountryCodeNil(b bool)`

 SetCountryCodeNil sets the value for CountryCode to be an explicit nil

### UnsetCountryCode
`func (o *MitIdAddressOutput) UnsetCountryCode()`

UnsetCountryCode ensures that no value is present for CountryCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


