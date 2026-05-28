# DiiaSubjectOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommonName** | Pointer to **NullableString** | Full name. | [optional] 
**Country** | Pointer to **NullableString** | Country code in ISO 3166-1 alpha-2 format. | [optional] 
**SerialNumber** | Pointer to **NullableString** | The SERIALNUMBER attribute from the certificate Subject Distinguished Name. This is not the X.509 certificate serial number field.              For Ukrainian natural person certificates this commonly uses the ETSI natural person semantics identifier format \&quot;{identifier-type}{country-code}-{identifier}\&quot;.              Components: - identifier-type (3 chars): TIN (Tax Identification Number) for RNOKPP. - country-code (2 chars): ISO 3166-1 alpha-2 country code, typically UA. - identifier: RNOKPP, the Ukrainian individual taxpayer registration number. | [optional] 
**GivenName** | Pointer to **NullableString** | Given name. | [optional] 
**Surname** | Pointer to **NullableString** | Surname. | [optional] 

## Methods

### NewDiiaSubjectOutput

`func NewDiiaSubjectOutput() *DiiaSubjectOutput`

NewDiiaSubjectOutput instantiates a new DiiaSubjectOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiiaSubjectOutputWithDefaults

`func NewDiiaSubjectOutputWithDefaults() *DiiaSubjectOutput`

NewDiiaSubjectOutputWithDefaults instantiates a new DiiaSubjectOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommonName

`func (o *DiiaSubjectOutput) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *DiiaSubjectOutput) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *DiiaSubjectOutput) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.

### HasCommonName

`func (o *DiiaSubjectOutput) HasCommonName() bool`

HasCommonName returns a boolean if a field has been set.

### SetCommonNameNil

`func (o *DiiaSubjectOutput) SetCommonNameNil(b bool)`

 SetCommonNameNil sets the value for CommonName to be an explicit nil

### UnsetCommonName
`func (o *DiiaSubjectOutput) UnsetCommonName()`

UnsetCommonName ensures that no value is present for CommonName, not even an explicit nil
### GetCountry

`func (o *DiiaSubjectOutput) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *DiiaSubjectOutput) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *DiiaSubjectOutput) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *DiiaSubjectOutput) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *DiiaSubjectOutput) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *DiiaSubjectOutput) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetSerialNumber

`func (o *DiiaSubjectOutput) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *DiiaSubjectOutput) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *DiiaSubjectOutput) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *DiiaSubjectOutput) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### SetSerialNumberNil

`func (o *DiiaSubjectOutput) SetSerialNumberNil(b bool)`

 SetSerialNumberNil sets the value for SerialNumber to be an explicit nil

### UnsetSerialNumber
`func (o *DiiaSubjectOutput) UnsetSerialNumber()`

UnsetSerialNumber ensures that no value is present for SerialNumber, not even an explicit nil
### GetGivenName

`func (o *DiiaSubjectOutput) GetGivenName() string`

GetGivenName returns the GivenName field if non-nil, zero value otherwise.

### GetGivenNameOk

`func (o *DiiaSubjectOutput) GetGivenNameOk() (*string, bool)`

GetGivenNameOk returns a tuple with the GivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenName

`func (o *DiiaSubjectOutput) SetGivenName(v string)`

SetGivenName sets GivenName field to given value.

### HasGivenName

`func (o *DiiaSubjectOutput) HasGivenName() bool`

HasGivenName returns a boolean if a field has been set.

### SetGivenNameNil

`func (o *DiiaSubjectOutput) SetGivenNameNil(b bool)`

 SetGivenNameNil sets the value for GivenName to be an explicit nil

### UnsetGivenName
`func (o *DiiaSubjectOutput) UnsetGivenName()`

UnsetGivenName ensures that no value is present for GivenName, not even an explicit nil
### GetSurname

`func (o *DiiaSubjectOutput) GetSurname() string`

GetSurname returns the Surname field if non-nil, zero value otherwise.

### GetSurnameOk

`func (o *DiiaSubjectOutput) GetSurnameOk() (*string, bool)`

GetSurnameOk returns a tuple with the Surname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurname

`func (o *DiiaSubjectOutput) SetSurname(v string)`

SetSurname sets Surname field to given value.

### HasSurname

`func (o *DiiaSubjectOutput) HasSurname() bool`

HasSurname returns a boolean if a field has been set.

### SetSurnameNil

`func (o *DiiaSubjectOutput) SetSurnameNil(b bool)`

 SetSurnameNil sets the value for Surname to be an explicit nil

### UnsetSurname
`func (o *DiiaSubjectOutput) UnsetSurname()`

UnsetSurname ensures that no value is present for Surname, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


