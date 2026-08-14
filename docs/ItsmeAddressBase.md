# ItsmeAddressBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StreetAddress** | Pointer to **NullableString** | The street address.              itsme documents this as an address member when the address claim is returned. | [optional] 
**PostalCode** | Pointer to **NullableString** | The postal code of the address.              itsme documents this as an address member when the address claim is returned. | [optional] 
**Country** | Pointer to **NullableString** | The country of the address as an ISO 3166-1 alpha-2 code, when returned.              This field is not part of itsme&#39;s documented address claim members; it is present only when itsme returns &#x60;address.country&#x60;. | [optional] 

## Methods

### NewItsmeAddressBase

`func NewItsmeAddressBase() *ItsmeAddressBase`

NewItsmeAddressBase instantiates a new ItsmeAddressBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItsmeAddressBaseWithDefaults

`func NewItsmeAddressBaseWithDefaults() *ItsmeAddressBase`

NewItsmeAddressBaseWithDefaults instantiates a new ItsmeAddressBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStreetAddress

`func (o *ItsmeAddressBase) GetStreetAddress() string`

GetStreetAddress returns the StreetAddress field if non-nil, zero value otherwise.

### GetStreetAddressOk

`func (o *ItsmeAddressBase) GetStreetAddressOk() (*string, bool)`

GetStreetAddressOk returns a tuple with the StreetAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetAddress

`func (o *ItsmeAddressBase) SetStreetAddress(v string)`

SetStreetAddress sets StreetAddress field to given value.

### HasStreetAddress

`func (o *ItsmeAddressBase) HasStreetAddress() bool`

HasStreetAddress returns a boolean if a field has been set.

### SetStreetAddressNil

`func (o *ItsmeAddressBase) SetStreetAddressNil(b bool)`

 SetStreetAddressNil sets the value for StreetAddress to be an explicit nil

### UnsetStreetAddress
`func (o *ItsmeAddressBase) UnsetStreetAddress()`

UnsetStreetAddress ensures that no value is present for StreetAddress, not even an explicit nil
### GetPostalCode

`func (o *ItsmeAddressBase) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *ItsmeAddressBase) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *ItsmeAddressBase) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *ItsmeAddressBase) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### SetPostalCodeNil

`func (o *ItsmeAddressBase) SetPostalCodeNil(b bool)`

 SetPostalCodeNil sets the value for PostalCode to be an explicit nil

### UnsetPostalCode
`func (o *ItsmeAddressBase) UnsetPostalCode()`

UnsetPostalCode ensures that no value is present for PostalCode, not even an explicit nil
### GetCountry

`func (o *ItsmeAddressBase) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *ItsmeAddressBase) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *ItsmeAddressBase) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *ItsmeAddressBase) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *ItsmeAddressBase) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *ItsmeAddressBase) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


