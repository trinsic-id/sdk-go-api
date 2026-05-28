# InteracAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StreetAddress** | Pointer to **NullableString** | The address line or street address. | [optional] 
**Locality** | Pointer to **NullableString** | The city or locality. | [optional] 
**Region** | Pointer to **NullableString** | The province, state, territory, or region. | [optional] 
**PostalCode** | Pointer to **NullableString** | The postal code. | [optional] 
**Country** | Pointer to **NullableString** | The address&#39;s ISO 3166-1 alpha-2 country code. | [optional] 

## Methods

### NewInteracAddress

`func NewInteracAddress() *InteracAddress`

NewInteracAddress instantiates a new InteracAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInteracAddressWithDefaults

`func NewInteracAddressWithDefaults() *InteracAddress`

NewInteracAddressWithDefaults instantiates a new InteracAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStreetAddress

`func (o *InteracAddress) GetStreetAddress() string`

GetStreetAddress returns the StreetAddress field if non-nil, zero value otherwise.

### GetStreetAddressOk

`func (o *InteracAddress) GetStreetAddressOk() (*string, bool)`

GetStreetAddressOk returns a tuple with the StreetAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetAddress

`func (o *InteracAddress) SetStreetAddress(v string)`

SetStreetAddress sets StreetAddress field to given value.

### HasStreetAddress

`func (o *InteracAddress) HasStreetAddress() bool`

HasStreetAddress returns a boolean if a field has been set.

### SetStreetAddressNil

`func (o *InteracAddress) SetStreetAddressNil(b bool)`

 SetStreetAddressNil sets the value for StreetAddress to be an explicit nil

### UnsetStreetAddress
`func (o *InteracAddress) UnsetStreetAddress()`

UnsetStreetAddress ensures that no value is present for StreetAddress, not even an explicit nil
### GetLocality

`func (o *InteracAddress) GetLocality() string`

GetLocality returns the Locality field if non-nil, zero value otherwise.

### GetLocalityOk

`func (o *InteracAddress) GetLocalityOk() (*string, bool)`

GetLocalityOk returns a tuple with the Locality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocality

`func (o *InteracAddress) SetLocality(v string)`

SetLocality sets Locality field to given value.

### HasLocality

`func (o *InteracAddress) HasLocality() bool`

HasLocality returns a boolean if a field has been set.

### SetLocalityNil

`func (o *InteracAddress) SetLocalityNil(b bool)`

 SetLocalityNil sets the value for Locality to be an explicit nil

### UnsetLocality
`func (o *InteracAddress) UnsetLocality()`

UnsetLocality ensures that no value is present for Locality, not even an explicit nil
### GetRegion

`func (o *InteracAddress) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *InteracAddress) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *InteracAddress) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *InteracAddress) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *InteracAddress) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *InteracAddress) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetPostalCode

`func (o *InteracAddress) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *InteracAddress) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *InteracAddress) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *InteracAddress) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### SetPostalCodeNil

`func (o *InteracAddress) SetPostalCodeNil(b bool)`

 SetPostalCodeNil sets the value for PostalCode to be an explicit nil

### UnsetPostalCode
`func (o *InteracAddress) UnsetPostalCode()`

UnsetPostalCode ensures that no value is present for PostalCode, not even an explicit nil
### GetCountry

`func (o *InteracAddress) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *InteracAddress) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *InteracAddress) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *InteracAddress) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *InteracAddress) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *InteracAddress) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


