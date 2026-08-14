# ItsmeIdentityDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SerialNumber** | Pointer to **NullableString** | The identity document serial number.              Availability by ID document issuing country: always returned for all supported issuing countries. | [optional] 
**TypeCode** | Pointer to **NullableString** | The identity document type code.              Availability by ID document issuing country: always returned for all supported issuing countries.              itsme documents this as a 1 or 2 character ICAO code. - Identity cards start with &#x60;I&#x60; - passports start with &#x60;P&#x60; | [optional] 
**IssuingCountry** | Pointer to **NullableString** | The issuing country as an ISO 3166-1 alpha-2 country code.              Availability by ID document issuing country: always returned for all supported issuing countries. | [optional] 
**ValidityFrom** | Pointer to **NullableTime** | The identity document validity start date time.              Availability by ID document issuing country: best effort for Belgian-issued ID documents; not returned for other supported issuing countries. | [optional] 
**ValidityTo** | Pointer to **NullableTime** | The identity document validity end date time.              Availability by ID document issuing country: always returned for supported issuing countries except Belgium, where it is best effort. | [optional] 

## Methods

### NewItsmeIdentityDocument

`func NewItsmeIdentityDocument() *ItsmeIdentityDocument`

NewItsmeIdentityDocument instantiates a new ItsmeIdentityDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItsmeIdentityDocumentWithDefaults

`func NewItsmeIdentityDocumentWithDefaults() *ItsmeIdentityDocument`

NewItsmeIdentityDocumentWithDefaults instantiates a new ItsmeIdentityDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerialNumber

`func (o *ItsmeIdentityDocument) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *ItsmeIdentityDocument) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *ItsmeIdentityDocument) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *ItsmeIdentityDocument) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### SetSerialNumberNil

`func (o *ItsmeIdentityDocument) SetSerialNumberNil(b bool)`

 SetSerialNumberNil sets the value for SerialNumber to be an explicit nil

### UnsetSerialNumber
`func (o *ItsmeIdentityDocument) UnsetSerialNumber()`

UnsetSerialNumber ensures that no value is present for SerialNumber, not even an explicit nil
### GetTypeCode

`func (o *ItsmeIdentityDocument) GetTypeCode() string`

GetTypeCode returns the TypeCode field if non-nil, zero value otherwise.

### GetTypeCodeOk

`func (o *ItsmeIdentityDocument) GetTypeCodeOk() (*string, bool)`

GetTypeCodeOk returns a tuple with the TypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeCode

`func (o *ItsmeIdentityDocument) SetTypeCode(v string)`

SetTypeCode sets TypeCode field to given value.

### HasTypeCode

`func (o *ItsmeIdentityDocument) HasTypeCode() bool`

HasTypeCode returns a boolean if a field has been set.

### SetTypeCodeNil

`func (o *ItsmeIdentityDocument) SetTypeCodeNil(b bool)`

 SetTypeCodeNil sets the value for TypeCode to be an explicit nil

### UnsetTypeCode
`func (o *ItsmeIdentityDocument) UnsetTypeCode()`

UnsetTypeCode ensures that no value is present for TypeCode, not even an explicit nil
### GetIssuingCountry

`func (o *ItsmeIdentityDocument) GetIssuingCountry() string`

GetIssuingCountry returns the IssuingCountry field if non-nil, zero value otherwise.

### GetIssuingCountryOk

`func (o *ItsmeIdentityDocument) GetIssuingCountryOk() (*string, bool)`

GetIssuingCountryOk returns a tuple with the IssuingCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingCountry

`func (o *ItsmeIdentityDocument) SetIssuingCountry(v string)`

SetIssuingCountry sets IssuingCountry field to given value.

### HasIssuingCountry

`func (o *ItsmeIdentityDocument) HasIssuingCountry() bool`

HasIssuingCountry returns a boolean if a field has been set.

### SetIssuingCountryNil

`func (o *ItsmeIdentityDocument) SetIssuingCountryNil(b bool)`

 SetIssuingCountryNil sets the value for IssuingCountry to be an explicit nil

### UnsetIssuingCountry
`func (o *ItsmeIdentityDocument) UnsetIssuingCountry()`

UnsetIssuingCountry ensures that no value is present for IssuingCountry, not even an explicit nil
### GetValidityFrom

`func (o *ItsmeIdentityDocument) GetValidityFrom() time.Time`

GetValidityFrom returns the ValidityFrom field if non-nil, zero value otherwise.

### GetValidityFromOk

`func (o *ItsmeIdentityDocument) GetValidityFromOk() (*time.Time, bool)`

GetValidityFromOk returns a tuple with the ValidityFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidityFrom

`func (o *ItsmeIdentityDocument) SetValidityFrom(v time.Time)`

SetValidityFrom sets ValidityFrom field to given value.

### HasValidityFrom

`func (o *ItsmeIdentityDocument) HasValidityFrom() bool`

HasValidityFrom returns a boolean if a field has been set.

### SetValidityFromNil

`func (o *ItsmeIdentityDocument) SetValidityFromNil(b bool)`

 SetValidityFromNil sets the value for ValidityFrom to be an explicit nil

### UnsetValidityFrom
`func (o *ItsmeIdentityDocument) UnsetValidityFrom()`

UnsetValidityFrom ensures that no value is present for ValidityFrom, not even an explicit nil
### GetValidityTo

`func (o *ItsmeIdentityDocument) GetValidityTo() time.Time`

GetValidityTo returns the ValidityTo field if non-nil, zero value otherwise.

### GetValidityToOk

`func (o *ItsmeIdentityDocument) GetValidityToOk() (*time.Time, bool)`

GetValidityToOk returns a tuple with the ValidityTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidityTo

`func (o *ItsmeIdentityDocument) SetValidityTo(v time.Time)`

SetValidityTo sets ValidityTo field to given value.

### HasValidityTo

`func (o *ItsmeIdentityDocument) HasValidityTo() bool`

HasValidityTo returns a boolean if a field has been set.

### SetValidityToNil

`func (o *ItsmeIdentityDocument) SetValidityToNil(b bool)`

 SetValidityToNil sets the value for ValidityTo to be an explicit nil

### UnsetValidityTo
`func (o *ItsmeIdentityDocument) UnsetValidityTo()`

UnsetValidityTo ensures that no value is present for ValidityTo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


