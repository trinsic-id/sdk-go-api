# DiiaIssuerOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommonName** | Pointer to **NullableString** | Certificate issuer common name. | [optional] 
**Organization** | Pointer to **NullableString** | Certificate issuer organization. | [optional] 
**Country** | Pointer to **NullableString** | Issuer country code in ISO 3166-1 alpha-2 format. | [optional] 
**Locality** | Pointer to **NullableString** | City or region where the issuer is located. | [optional] 
**SerialNumber** | Pointer to **NullableString** | Unique identifier assigned to the issuing certificate authority.              Typically formatted as \&quot;{country}-{registration code}-{sequence}\&quot;. | [optional] 
**NationalRegistrationNumber** | Pointer to **NullableString** | The organizationIdentifier attribute from the certificate Issuer Distinguished Name.              Ukrainian organization identifiers commonly use the ETSI format \&quot;{identifier-type}{country-code}-{identifier}\&quot;.              Components: - identifier-type (3 chars): NTR (National Trade Register). - country-code (2 chars): ISO 3166-1 alpha-2 country code, typically UA. - identifier: The legal entity registration code.              For Diia.Signature, this is issuer metadata rather than a natural person&#39;s RNOKPP. | [optional] 

## Methods

### NewDiiaIssuerOutput

`func NewDiiaIssuerOutput() *DiiaIssuerOutput`

NewDiiaIssuerOutput instantiates a new DiiaIssuerOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiiaIssuerOutputWithDefaults

`func NewDiiaIssuerOutputWithDefaults() *DiiaIssuerOutput`

NewDiiaIssuerOutputWithDefaults instantiates a new DiiaIssuerOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommonName

`func (o *DiiaIssuerOutput) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *DiiaIssuerOutput) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *DiiaIssuerOutput) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.

### HasCommonName

`func (o *DiiaIssuerOutput) HasCommonName() bool`

HasCommonName returns a boolean if a field has been set.

### SetCommonNameNil

`func (o *DiiaIssuerOutput) SetCommonNameNil(b bool)`

 SetCommonNameNil sets the value for CommonName to be an explicit nil

### UnsetCommonName
`func (o *DiiaIssuerOutput) UnsetCommonName()`

UnsetCommonName ensures that no value is present for CommonName, not even an explicit nil
### GetOrganization

`func (o *DiiaIssuerOutput) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *DiiaIssuerOutput) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *DiiaIssuerOutput) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *DiiaIssuerOutput) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### SetOrganizationNil

`func (o *DiiaIssuerOutput) SetOrganizationNil(b bool)`

 SetOrganizationNil sets the value for Organization to be an explicit nil

### UnsetOrganization
`func (o *DiiaIssuerOutput) UnsetOrganization()`

UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
### GetCountry

`func (o *DiiaIssuerOutput) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *DiiaIssuerOutput) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *DiiaIssuerOutput) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *DiiaIssuerOutput) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *DiiaIssuerOutput) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *DiiaIssuerOutput) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetLocality

`func (o *DiiaIssuerOutput) GetLocality() string`

GetLocality returns the Locality field if non-nil, zero value otherwise.

### GetLocalityOk

`func (o *DiiaIssuerOutput) GetLocalityOk() (*string, bool)`

GetLocalityOk returns a tuple with the Locality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocality

`func (o *DiiaIssuerOutput) SetLocality(v string)`

SetLocality sets Locality field to given value.

### HasLocality

`func (o *DiiaIssuerOutput) HasLocality() bool`

HasLocality returns a boolean if a field has been set.

### SetLocalityNil

`func (o *DiiaIssuerOutput) SetLocalityNil(b bool)`

 SetLocalityNil sets the value for Locality to be an explicit nil

### UnsetLocality
`func (o *DiiaIssuerOutput) UnsetLocality()`

UnsetLocality ensures that no value is present for Locality, not even an explicit nil
### GetSerialNumber

`func (o *DiiaIssuerOutput) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *DiiaIssuerOutput) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *DiiaIssuerOutput) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *DiiaIssuerOutput) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### SetSerialNumberNil

`func (o *DiiaIssuerOutput) SetSerialNumberNil(b bool)`

 SetSerialNumberNil sets the value for SerialNumber to be an explicit nil

### UnsetSerialNumber
`func (o *DiiaIssuerOutput) UnsetSerialNumber()`

UnsetSerialNumber ensures that no value is present for SerialNumber, not even an explicit nil
### GetNationalRegistrationNumber

`func (o *DiiaIssuerOutput) GetNationalRegistrationNumber() string`

GetNationalRegistrationNumber returns the NationalRegistrationNumber field if non-nil, zero value otherwise.

### GetNationalRegistrationNumberOk

`func (o *DiiaIssuerOutput) GetNationalRegistrationNumberOk() (*string, bool)`

GetNationalRegistrationNumberOk returns a tuple with the NationalRegistrationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationalRegistrationNumber

`func (o *DiiaIssuerOutput) SetNationalRegistrationNumber(v string)`

SetNationalRegistrationNumber sets NationalRegistrationNumber field to given value.

### HasNationalRegistrationNumber

`func (o *DiiaIssuerOutput) HasNationalRegistrationNumber() bool`

HasNationalRegistrationNumber returns a boolean if a field has been set.

### SetNationalRegistrationNumberNil

`func (o *DiiaIssuerOutput) SetNationalRegistrationNumberNil(b bool)`

 SetNationalRegistrationNumberNil sets the value for NationalRegistrationNumber to be an explicit nil

### UnsetNationalRegistrationNumber
`func (o *DiiaIssuerOutput) UnsetNationalRegistrationNumber()`

UnsetNationalRegistrationNumber ensures that no value is present for NationalRegistrationNumber, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


