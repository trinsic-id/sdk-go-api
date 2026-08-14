# ItsmeBelgianNationalNumber

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | Pointer to **NullableString** | The raw Belgian National Register Number of the verified individual.              Availability by ID document issuing country: returned only for Belgian-issued ID documents.              This is an 11-digit number in the format YYMMDDXXXCC: - YYMMDD is the date of birth. - XXX is a sequential birth number, odd for males and even for females. - CC is the checksum. | [optional] 
**IssuingCountry** | Pointer to **NullableString** | The issuing country as an ISO 3166-1 alpha-2 country code.              Availability by ID document issuing country: always returned for all supported issuing countries. | [optional] 
**VerificationDate** | Pointer to **NullableTime** | The date time when the document was last read.              Availability by ID document issuing country: always returned for supported issuing countries except Belgium, where it is best effort. | [optional] 

## Methods

### NewItsmeBelgianNationalNumber

`func NewItsmeBelgianNationalNumber() *ItsmeBelgianNationalNumber`

NewItsmeBelgianNationalNumber instantiates a new ItsmeBelgianNationalNumber object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItsmeBelgianNationalNumberWithDefaults

`func NewItsmeBelgianNationalNumberWithDefaults() *ItsmeBelgianNationalNumber`

NewItsmeBelgianNationalNumberWithDefaults instantiates a new ItsmeBelgianNationalNumber object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *ItsmeBelgianNationalNumber) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *ItsmeBelgianNationalNumber) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *ItsmeBelgianNationalNumber) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *ItsmeBelgianNationalNumber) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### SetNumberNil

`func (o *ItsmeBelgianNationalNumber) SetNumberNil(b bool)`

 SetNumberNil sets the value for Number to be an explicit nil

### UnsetNumber
`func (o *ItsmeBelgianNationalNumber) UnsetNumber()`

UnsetNumber ensures that no value is present for Number, not even an explicit nil
### GetIssuingCountry

`func (o *ItsmeBelgianNationalNumber) GetIssuingCountry() string`

GetIssuingCountry returns the IssuingCountry field if non-nil, zero value otherwise.

### GetIssuingCountryOk

`func (o *ItsmeBelgianNationalNumber) GetIssuingCountryOk() (*string, bool)`

GetIssuingCountryOk returns a tuple with the IssuingCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingCountry

`func (o *ItsmeBelgianNationalNumber) SetIssuingCountry(v string)`

SetIssuingCountry sets IssuingCountry field to given value.

### HasIssuingCountry

`func (o *ItsmeBelgianNationalNumber) HasIssuingCountry() bool`

HasIssuingCountry returns a boolean if a field has been set.

### SetIssuingCountryNil

`func (o *ItsmeBelgianNationalNumber) SetIssuingCountryNil(b bool)`

 SetIssuingCountryNil sets the value for IssuingCountry to be an explicit nil

### UnsetIssuingCountry
`func (o *ItsmeBelgianNationalNumber) UnsetIssuingCountry()`

UnsetIssuingCountry ensures that no value is present for IssuingCountry, not even an explicit nil
### GetVerificationDate

`func (o *ItsmeBelgianNationalNumber) GetVerificationDate() time.Time`

GetVerificationDate returns the VerificationDate field if non-nil, zero value otherwise.

### GetVerificationDateOk

`func (o *ItsmeBelgianNationalNumber) GetVerificationDateOk() (*time.Time, bool)`

GetVerificationDateOk returns a tuple with the VerificationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationDate

`func (o *ItsmeBelgianNationalNumber) SetVerificationDate(v time.Time)`

SetVerificationDate sets VerificationDate field to given value.

### HasVerificationDate

`func (o *ItsmeBelgianNationalNumber) HasVerificationDate() bool`

HasVerificationDate returns a boolean if a field has been set.

### SetVerificationDateNil

`func (o *ItsmeBelgianNationalNumber) SetVerificationDateNil(b bool)`

 SetVerificationDateNil sets the value for VerificationDate to be an explicit nil

### UnsetVerificationDate
`func (o *ItsmeBelgianNationalNumber) UnsetVerificationDate()`

UnsetVerificationDate ensures that no value is present for VerificationDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


