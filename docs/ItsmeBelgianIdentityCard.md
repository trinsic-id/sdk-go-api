# ItsmeBelgianIdentityCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | Pointer to **NullableString** | The Belgian eID card number. This identifies the card document, not the person&#39;s Belgian National Register Number.              Availability by ID document issuing country: returned only for Belgian-issued ID documents.              Belgian citizen examples are 12 digits in the form &#x60;xxx-xxxxxxx-yy&#x60;; EU, EEA, and Swiss resident card examples can start with a letter followed by digits. | [optional] 
**IssuingCountry** | Pointer to **NullableString** | The issuing country as an ISO 3166-1 alpha-2 country code.              Availability by ID document issuing country: always returned for all supported issuing countries. | [optional] 
**ValidityFrom** | Pointer to **NullableTime** | The identity card validity start date time.              Availability by ID document issuing country: best effort for Belgian-issued ID documents; not returned for other supported issuing countries. | [optional] 
**ValidityTo** | Pointer to **NullableTime** | The identity card validity end date time.              Availability by ID document issuing country: always returned for supported issuing countries except Belgium, where it is best effort. | [optional] 
**IssuanceLocality** | Pointer to **NullableString** | The locality that issued the identity card.              Availability by ID document issuing country: best effort for Belgian-issued ID documents; not returned for other supported issuing countries. | [optional] 

## Methods

### NewItsmeBelgianIdentityCard

`func NewItsmeBelgianIdentityCard() *ItsmeBelgianIdentityCard`

NewItsmeBelgianIdentityCard instantiates a new ItsmeBelgianIdentityCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItsmeBelgianIdentityCardWithDefaults

`func NewItsmeBelgianIdentityCardWithDefaults() *ItsmeBelgianIdentityCard`

NewItsmeBelgianIdentityCardWithDefaults instantiates a new ItsmeBelgianIdentityCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *ItsmeBelgianIdentityCard) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *ItsmeBelgianIdentityCard) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *ItsmeBelgianIdentityCard) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *ItsmeBelgianIdentityCard) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### SetNumberNil

`func (o *ItsmeBelgianIdentityCard) SetNumberNil(b bool)`

 SetNumberNil sets the value for Number to be an explicit nil

### UnsetNumber
`func (o *ItsmeBelgianIdentityCard) UnsetNumber()`

UnsetNumber ensures that no value is present for Number, not even an explicit nil
### GetIssuingCountry

`func (o *ItsmeBelgianIdentityCard) GetIssuingCountry() string`

GetIssuingCountry returns the IssuingCountry field if non-nil, zero value otherwise.

### GetIssuingCountryOk

`func (o *ItsmeBelgianIdentityCard) GetIssuingCountryOk() (*string, bool)`

GetIssuingCountryOk returns a tuple with the IssuingCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingCountry

`func (o *ItsmeBelgianIdentityCard) SetIssuingCountry(v string)`

SetIssuingCountry sets IssuingCountry field to given value.

### HasIssuingCountry

`func (o *ItsmeBelgianIdentityCard) HasIssuingCountry() bool`

HasIssuingCountry returns a boolean if a field has been set.

### SetIssuingCountryNil

`func (o *ItsmeBelgianIdentityCard) SetIssuingCountryNil(b bool)`

 SetIssuingCountryNil sets the value for IssuingCountry to be an explicit nil

### UnsetIssuingCountry
`func (o *ItsmeBelgianIdentityCard) UnsetIssuingCountry()`

UnsetIssuingCountry ensures that no value is present for IssuingCountry, not even an explicit nil
### GetValidityFrom

`func (o *ItsmeBelgianIdentityCard) GetValidityFrom() time.Time`

GetValidityFrom returns the ValidityFrom field if non-nil, zero value otherwise.

### GetValidityFromOk

`func (o *ItsmeBelgianIdentityCard) GetValidityFromOk() (*time.Time, bool)`

GetValidityFromOk returns a tuple with the ValidityFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidityFrom

`func (o *ItsmeBelgianIdentityCard) SetValidityFrom(v time.Time)`

SetValidityFrom sets ValidityFrom field to given value.

### HasValidityFrom

`func (o *ItsmeBelgianIdentityCard) HasValidityFrom() bool`

HasValidityFrom returns a boolean if a field has been set.

### SetValidityFromNil

`func (o *ItsmeBelgianIdentityCard) SetValidityFromNil(b bool)`

 SetValidityFromNil sets the value for ValidityFrom to be an explicit nil

### UnsetValidityFrom
`func (o *ItsmeBelgianIdentityCard) UnsetValidityFrom()`

UnsetValidityFrom ensures that no value is present for ValidityFrom, not even an explicit nil
### GetValidityTo

`func (o *ItsmeBelgianIdentityCard) GetValidityTo() time.Time`

GetValidityTo returns the ValidityTo field if non-nil, zero value otherwise.

### GetValidityToOk

`func (o *ItsmeBelgianIdentityCard) GetValidityToOk() (*time.Time, bool)`

GetValidityToOk returns a tuple with the ValidityTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidityTo

`func (o *ItsmeBelgianIdentityCard) SetValidityTo(v time.Time)`

SetValidityTo sets ValidityTo field to given value.

### HasValidityTo

`func (o *ItsmeBelgianIdentityCard) HasValidityTo() bool`

HasValidityTo returns a boolean if a field has been set.

### SetValidityToNil

`func (o *ItsmeBelgianIdentityCard) SetValidityToNil(b bool)`

 SetValidityToNil sets the value for ValidityTo to be an explicit nil

### UnsetValidityTo
`func (o *ItsmeBelgianIdentityCard) UnsetValidityTo()`

UnsetValidityTo ensures that no value is present for ValidityTo, not even an explicit nil
### GetIssuanceLocality

`func (o *ItsmeBelgianIdentityCard) GetIssuanceLocality() string`

GetIssuanceLocality returns the IssuanceLocality field if non-nil, zero value otherwise.

### GetIssuanceLocalityOk

`func (o *ItsmeBelgianIdentityCard) GetIssuanceLocalityOk() (*string, bool)`

GetIssuanceLocalityOk returns a tuple with the IssuanceLocality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuanceLocality

`func (o *ItsmeBelgianIdentityCard) SetIssuanceLocality(v string)`

SetIssuanceLocality sets IssuanceLocality field to given value.

### HasIssuanceLocality

`func (o *ItsmeBelgianIdentityCard) HasIssuanceLocality() bool`

HasIssuanceLocality returns a boolean if a field has been set.

### SetIssuanceLocalityNil

`func (o *ItsmeBelgianIdentityCard) SetIssuanceLocalityNil(b bool)`

 SetIssuanceLocalityNil sets the value for IssuanceLocality to be an explicit nil

### UnsetIssuanceLocality
`func (o *ItsmeBelgianIdentityCard) UnsetIssuanceLocality()`

UnsetIssuanceLocality ensures that no value is present for IssuanceLocality, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


