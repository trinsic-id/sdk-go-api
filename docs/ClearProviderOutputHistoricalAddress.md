# ClearProviderOutputHistoricalAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Line1** | Pointer to **NullableString** | The first line of the address. | [optional] 
**Line2** | Pointer to **NullableString** | The second line of the address, such as an apartment or suite. | [optional] 
**City** | Pointer to **NullableString** | The city. | [optional] 
**State** | Pointer to **NullableString** | The state, province, or region. | [optional] 
**PostalCode** | Pointer to **NullableString** | The postal code. | [optional] 
**Country** | Pointer to **NullableString** | The ISO 3166-1 alpha-2 country code. | [optional] 
**StreetNumber** | Pointer to **NullableString** | The historical address street number. | [optional] 
**StreetName** | Pointer to **NullableString** | The historical address street name. | [optional] 
**FirstSeen** | Pointer to [**NullableClearProviderOutputPartialDate**](ClearProviderOutputPartialDate.md) | The first date CLEAR saw this historical address. | [optional] 
**LastSeen** | Pointer to [**NullableClearProviderOutputPartialDate**](ClearProviderOutputPartialDate.md) | The last date CLEAR saw this historical address. | [optional] 

## Methods

### NewClearProviderOutputHistoricalAddress

`func NewClearProviderOutputHistoricalAddress() *ClearProviderOutputHistoricalAddress`

NewClearProviderOutputHistoricalAddress instantiates a new ClearProviderOutputHistoricalAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClearProviderOutputHistoricalAddressWithDefaults

`func NewClearProviderOutputHistoricalAddressWithDefaults() *ClearProviderOutputHistoricalAddress`

NewClearProviderOutputHistoricalAddressWithDefaults instantiates a new ClearProviderOutputHistoricalAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLine1

`func (o *ClearProviderOutputHistoricalAddress) GetLine1() string`

GetLine1 returns the Line1 field if non-nil, zero value otherwise.

### GetLine1Ok

`func (o *ClearProviderOutputHistoricalAddress) GetLine1Ok() (*string, bool)`

GetLine1Ok returns a tuple with the Line1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine1

`func (o *ClearProviderOutputHistoricalAddress) SetLine1(v string)`

SetLine1 sets Line1 field to given value.

### HasLine1

`func (o *ClearProviderOutputHistoricalAddress) HasLine1() bool`

HasLine1 returns a boolean if a field has been set.

### SetLine1Nil

`func (o *ClearProviderOutputHistoricalAddress) SetLine1Nil(b bool)`

 SetLine1Nil sets the value for Line1 to be an explicit nil

### UnsetLine1
`func (o *ClearProviderOutputHistoricalAddress) UnsetLine1()`

UnsetLine1 ensures that no value is present for Line1, not even an explicit nil
### GetLine2

`func (o *ClearProviderOutputHistoricalAddress) GetLine2() string`

GetLine2 returns the Line2 field if non-nil, zero value otherwise.

### GetLine2Ok

`func (o *ClearProviderOutputHistoricalAddress) GetLine2Ok() (*string, bool)`

GetLine2Ok returns a tuple with the Line2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine2

`func (o *ClearProviderOutputHistoricalAddress) SetLine2(v string)`

SetLine2 sets Line2 field to given value.

### HasLine2

`func (o *ClearProviderOutputHistoricalAddress) HasLine2() bool`

HasLine2 returns a boolean if a field has been set.

### SetLine2Nil

`func (o *ClearProviderOutputHistoricalAddress) SetLine2Nil(b bool)`

 SetLine2Nil sets the value for Line2 to be an explicit nil

### UnsetLine2
`func (o *ClearProviderOutputHistoricalAddress) UnsetLine2()`

UnsetLine2 ensures that no value is present for Line2, not even an explicit nil
### GetCity

`func (o *ClearProviderOutputHistoricalAddress) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *ClearProviderOutputHistoricalAddress) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *ClearProviderOutputHistoricalAddress) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *ClearProviderOutputHistoricalAddress) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *ClearProviderOutputHistoricalAddress) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *ClearProviderOutputHistoricalAddress) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetState

`func (o *ClearProviderOutputHistoricalAddress) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ClearProviderOutputHistoricalAddress) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ClearProviderOutputHistoricalAddress) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ClearProviderOutputHistoricalAddress) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *ClearProviderOutputHistoricalAddress) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *ClearProviderOutputHistoricalAddress) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetPostalCode

`func (o *ClearProviderOutputHistoricalAddress) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *ClearProviderOutputHistoricalAddress) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *ClearProviderOutputHistoricalAddress) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *ClearProviderOutputHistoricalAddress) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### SetPostalCodeNil

`func (o *ClearProviderOutputHistoricalAddress) SetPostalCodeNil(b bool)`

 SetPostalCodeNil sets the value for PostalCode to be an explicit nil

### UnsetPostalCode
`func (o *ClearProviderOutputHistoricalAddress) UnsetPostalCode()`

UnsetPostalCode ensures that no value is present for PostalCode, not even an explicit nil
### GetCountry

`func (o *ClearProviderOutputHistoricalAddress) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *ClearProviderOutputHistoricalAddress) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *ClearProviderOutputHistoricalAddress) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *ClearProviderOutputHistoricalAddress) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *ClearProviderOutputHistoricalAddress) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *ClearProviderOutputHistoricalAddress) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetStreetNumber

`func (o *ClearProviderOutputHistoricalAddress) GetStreetNumber() string`

GetStreetNumber returns the StreetNumber field if non-nil, zero value otherwise.

### GetStreetNumberOk

`func (o *ClearProviderOutputHistoricalAddress) GetStreetNumberOk() (*string, bool)`

GetStreetNumberOk returns a tuple with the StreetNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetNumber

`func (o *ClearProviderOutputHistoricalAddress) SetStreetNumber(v string)`

SetStreetNumber sets StreetNumber field to given value.

### HasStreetNumber

`func (o *ClearProviderOutputHistoricalAddress) HasStreetNumber() bool`

HasStreetNumber returns a boolean if a field has been set.

### SetStreetNumberNil

`func (o *ClearProviderOutputHistoricalAddress) SetStreetNumberNil(b bool)`

 SetStreetNumberNil sets the value for StreetNumber to be an explicit nil

### UnsetStreetNumber
`func (o *ClearProviderOutputHistoricalAddress) UnsetStreetNumber()`

UnsetStreetNumber ensures that no value is present for StreetNumber, not even an explicit nil
### GetStreetName

`func (o *ClearProviderOutputHistoricalAddress) GetStreetName() string`

GetStreetName returns the StreetName field if non-nil, zero value otherwise.

### GetStreetNameOk

`func (o *ClearProviderOutputHistoricalAddress) GetStreetNameOk() (*string, bool)`

GetStreetNameOk returns a tuple with the StreetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetName

`func (o *ClearProviderOutputHistoricalAddress) SetStreetName(v string)`

SetStreetName sets StreetName field to given value.

### HasStreetName

`func (o *ClearProviderOutputHistoricalAddress) HasStreetName() bool`

HasStreetName returns a boolean if a field has been set.

### SetStreetNameNil

`func (o *ClearProviderOutputHistoricalAddress) SetStreetNameNil(b bool)`

 SetStreetNameNil sets the value for StreetName to be an explicit nil

### UnsetStreetName
`func (o *ClearProviderOutputHistoricalAddress) UnsetStreetName()`

UnsetStreetName ensures that no value is present for StreetName, not even an explicit nil
### GetFirstSeen

`func (o *ClearProviderOutputHistoricalAddress) GetFirstSeen() ClearProviderOutputPartialDate`

GetFirstSeen returns the FirstSeen field if non-nil, zero value otherwise.

### GetFirstSeenOk

`func (o *ClearProviderOutputHistoricalAddress) GetFirstSeenOk() (*ClearProviderOutputPartialDate, bool)`

GetFirstSeenOk returns a tuple with the FirstSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeen

`func (o *ClearProviderOutputHistoricalAddress) SetFirstSeen(v ClearProviderOutputPartialDate)`

SetFirstSeen sets FirstSeen field to given value.

### HasFirstSeen

`func (o *ClearProviderOutputHistoricalAddress) HasFirstSeen() bool`

HasFirstSeen returns a boolean if a field has been set.

### SetFirstSeenNil

`func (o *ClearProviderOutputHistoricalAddress) SetFirstSeenNil(b bool)`

 SetFirstSeenNil sets the value for FirstSeen to be an explicit nil

### UnsetFirstSeen
`func (o *ClearProviderOutputHistoricalAddress) UnsetFirstSeen()`

UnsetFirstSeen ensures that no value is present for FirstSeen, not even an explicit nil
### GetLastSeen

`func (o *ClearProviderOutputHistoricalAddress) GetLastSeen() ClearProviderOutputPartialDate`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *ClearProviderOutputHistoricalAddress) GetLastSeenOk() (*ClearProviderOutputPartialDate, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *ClearProviderOutputHistoricalAddress) SetLastSeen(v ClearProviderOutputPartialDate)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *ClearProviderOutputHistoricalAddress) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### SetLastSeenNil

`func (o *ClearProviderOutputHistoricalAddress) SetLastSeenNil(b bool)`

 SetLastSeenNil sets the value for LastSeen to be an explicit nil

### UnsetLastSeen
`func (o *ClearProviderOutputHistoricalAddress) UnsetLastSeen()`

UnsetLastSeen ensures that no value is present for LastSeen, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


