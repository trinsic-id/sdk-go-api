# ItsmeProviderAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StreetAddress** | Pointer to **NullableString** | The street address.              itsme documents this as an address member when the address claim is returned. | [optional] 
**PostalCode** | Pointer to **NullableString** | The postal code of the address.              itsme documents this as an address member when the address claim is returned. | [optional] 
**Country** | Pointer to **NullableString** | The country of the address as an ISO 3166-1 alpha-2 code, when returned.              This field is not part of itsme&#39;s documented address claim members; it is present only when itsme returns &#x60;address.country&#x60;. | [optional] 
**Locality** | Pointer to **NullableString** | The locality or city of the address.              itsme documents this as an address member when the address claim is returned. | [optional] 
**Formatted** | Pointer to **NullableString** | The full address formatted by itsme.              itsme documents this as an address member when the address claim is returned. | [optional] 

## Methods

### NewItsmeProviderAddress

`func NewItsmeProviderAddress() *ItsmeProviderAddress`

NewItsmeProviderAddress instantiates a new ItsmeProviderAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItsmeProviderAddressWithDefaults

`func NewItsmeProviderAddressWithDefaults() *ItsmeProviderAddress`

NewItsmeProviderAddressWithDefaults instantiates a new ItsmeProviderAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStreetAddress

`func (o *ItsmeProviderAddress) GetStreetAddress() string`

GetStreetAddress returns the StreetAddress field if non-nil, zero value otherwise.

### GetStreetAddressOk

`func (o *ItsmeProviderAddress) GetStreetAddressOk() (*string, bool)`

GetStreetAddressOk returns a tuple with the StreetAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetAddress

`func (o *ItsmeProviderAddress) SetStreetAddress(v string)`

SetStreetAddress sets StreetAddress field to given value.

### HasStreetAddress

`func (o *ItsmeProviderAddress) HasStreetAddress() bool`

HasStreetAddress returns a boolean if a field has been set.

### SetStreetAddressNil

`func (o *ItsmeProviderAddress) SetStreetAddressNil(b bool)`

 SetStreetAddressNil sets the value for StreetAddress to be an explicit nil

### UnsetStreetAddress
`func (o *ItsmeProviderAddress) UnsetStreetAddress()`

UnsetStreetAddress ensures that no value is present for StreetAddress, not even an explicit nil
### GetPostalCode

`func (o *ItsmeProviderAddress) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *ItsmeProviderAddress) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *ItsmeProviderAddress) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *ItsmeProviderAddress) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### SetPostalCodeNil

`func (o *ItsmeProviderAddress) SetPostalCodeNil(b bool)`

 SetPostalCodeNil sets the value for PostalCode to be an explicit nil

### UnsetPostalCode
`func (o *ItsmeProviderAddress) UnsetPostalCode()`

UnsetPostalCode ensures that no value is present for PostalCode, not even an explicit nil
### GetCountry

`func (o *ItsmeProviderAddress) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *ItsmeProviderAddress) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *ItsmeProviderAddress) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *ItsmeProviderAddress) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *ItsmeProviderAddress) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *ItsmeProviderAddress) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetLocality

`func (o *ItsmeProviderAddress) GetLocality() string`

GetLocality returns the Locality field if non-nil, zero value otherwise.

### GetLocalityOk

`func (o *ItsmeProviderAddress) GetLocalityOk() (*string, bool)`

GetLocalityOk returns a tuple with the Locality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocality

`func (o *ItsmeProviderAddress) SetLocality(v string)`

SetLocality sets Locality field to given value.

### HasLocality

`func (o *ItsmeProviderAddress) HasLocality() bool`

HasLocality returns a boolean if a field has been set.

### SetLocalityNil

`func (o *ItsmeProviderAddress) SetLocalityNil(b bool)`

 SetLocalityNil sets the value for Locality to be an explicit nil

### UnsetLocality
`func (o *ItsmeProviderAddress) UnsetLocality()`

UnsetLocality ensures that no value is present for Locality, not even an explicit nil
### GetFormatted

`func (o *ItsmeProviderAddress) GetFormatted() string`

GetFormatted returns the Formatted field if non-nil, zero value otherwise.

### GetFormattedOk

`func (o *ItsmeProviderAddress) GetFormattedOk() (*string, bool)`

GetFormattedOk returns a tuple with the Formatted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatted

`func (o *ItsmeProviderAddress) SetFormatted(v string)`

SetFormatted sets Formatted field to given value.

### HasFormatted

`func (o *ItsmeProviderAddress) HasFormatted() bool`

HasFormatted returns a boolean if a field has been set.

### SetFormattedNil

`func (o *ItsmeProviderAddress) SetFormattedNil(b bool)`

 SetFormattedNil sets the value for Formatted to be an explicit nil

### UnsetFormatted
`func (o *ItsmeProviderAddress) UnsetFormatted()`

UnsetFormatted ensures that no value is present for Formatted, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


