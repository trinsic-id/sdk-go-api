# ConnectIdAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StreetAddress** | Pointer to **NullableString** | The address&#39; line or street. | [optional] 
**Locality** | Pointer to **NullableString** | The address&#39; locality or city. | [optional] 
**Region** | Pointer to **NullableString** | The address&#39; region, state, or territory. | [optional] 
**Country** | Pointer to **NullableString** | The address&#39; country.              This value is not normalized and may be an alpha-3 code or full country name. | [optional] 
**PostalCode** | Pointer to **NullableString** | The address&#39; postal code. | [optional] 

## Methods

### NewConnectIdAddress

`func NewConnectIdAddress() *ConnectIdAddress`

NewConnectIdAddress instantiates a new ConnectIdAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectIdAddressWithDefaults

`func NewConnectIdAddressWithDefaults() *ConnectIdAddress`

NewConnectIdAddressWithDefaults instantiates a new ConnectIdAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStreetAddress

`func (o *ConnectIdAddress) GetStreetAddress() string`

GetStreetAddress returns the StreetAddress field if non-nil, zero value otherwise.

### GetStreetAddressOk

`func (o *ConnectIdAddress) GetStreetAddressOk() (*string, bool)`

GetStreetAddressOk returns a tuple with the StreetAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetAddress

`func (o *ConnectIdAddress) SetStreetAddress(v string)`

SetStreetAddress sets StreetAddress field to given value.

### HasStreetAddress

`func (o *ConnectIdAddress) HasStreetAddress() bool`

HasStreetAddress returns a boolean if a field has been set.

### SetStreetAddressNil

`func (o *ConnectIdAddress) SetStreetAddressNil(b bool)`

 SetStreetAddressNil sets the value for StreetAddress to be an explicit nil

### UnsetStreetAddress
`func (o *ConnectIdAddress) UnsetStreetAddress()`

UnsetStreetAddress ensures that no value is present for StreetAddress, not even an explicit nil
### GetLocality

`func (o *ConnectIdAddress) GetLocality() string`

GetLocality returns the Locality field if non-nil, zero value otherwise.

### GetLocalityOk

`func (o *ConnectIdAddress) GetLocalityOk() (*string, bool)`

GetLocalityOk returns a tuple with the Locality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocality

`func (o *ConnectIdAddress) SetLocality(v string)`

SetLocality sets Locality field to given value.

### HasLocality

`func (o *ConnectIdAddress) HasLocality() bool`

HasLocality returns a boolean if a field has been set.

### SetLocalityNil

`func (o *ConnectIdAddress) SetLocalityNil(b bool)`

 SetLocalityNil sets the value for Locality to be an explicit nil

### UnsetLocality
`func (o *ConnectIdAddress) UnsetLocality()`

UnsetLocality ensures that no value is present for Locality, not even an explicit nil
### GetRegion

`func (o *ConnectIdAddress) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ConnectIdAddress) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ConnectIdAddress) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *ConnectIdAddress) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *ConnectIdAddress) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *ConnectIdAddress) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetCountry

`func (o *ConnectIdAddress) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *ConnectIdAddress) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *ConnectIdAddress) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *ConnectIdAddress) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *ConnectIdAddress) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *ConnectIdAddress) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetPostalCode

`func (o *ConnectIdAddress) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *ConnectIdAddress) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *ConnectIdAddress) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *ConnectIdAddress) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### SetPostalCodeNil

`func (o *ConnectIdAddress) SetPostalCodeNil(b bool)`

 SetPostalCodeNil sets the value for PostalCode to be an explicit nil

### UnsetPostalCode
`func (o *ConnectIdAddress) UnsetPostalCode()`

UnsetPostalCode ensures that no value is present for PostalCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


