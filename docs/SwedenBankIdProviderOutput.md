# SwedenBankIdProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FullName** | Pointer to **NullableString** | The full name of the individual. | [optional] 
**DateOfBirth** | Pointer to **NullableString** | The date of birth of the individual. | [optional] 
**PersonalIdentityNumber** | Pointer to **NullableString** | The Swedish personal number (personnummer) associated with the verified individual. The value is typically a 12-digit number. This is in the format YYYYMMDDXXXX, where: - YYYYMMDD is the date of birth. - XXXX is an individual number, indicating gender.              If XXXX is even, the individual is female. If XXXX is odd, the individual is male. | [optional] 
**GivenName** | Pointer to **NullableString** | The given name of the individual | [optional] 
**FamilyName** | Pointer to **NullableString** | The family name of the individual | [optional] 
**Country** | Pointer to **NullableString** | The end user&#39;s country on the verified certificate. This is ISO 2-character code of the country. | [optional] 
**CertificatePolicy** | Pointer to **NullableString** | The object identifier policy of the individual&#39;s verified certificate | [optional] 
**CommonName** | Pointer to **NullableString** | The common name of the end user&#39;s verified certificate | [optional] 
**DistinguishedName** | Pointer to **NullableString** | The distinguished name of the end user&#39;s verified certificate | [optional] 
**PersonalIdentityNumberIssuingCountry** | Pointer to **NullableString** | The country that issued the Swedish personal number. This is an ISO 2-character code of the country and is extracted from the certificate. This will always be &#x60;SE&#x60;. | [optional] 
**AuthenticationDeviceIp** | Pointer to **NullableString** | The IP address of the device used for authentication | [optional] 
**NotBefore** | Pointer to **NullableTime** | The certificate validity start date (not before) in UTC | [optional] 
**NotAfter** | Pointer to **NullableTime** | The certificate validity end date (not after) in UTC | [optional] 

## Methods

### NewSwedenBankIdProviderOutput

`func NewSwedenBankIdProviderOutput() *SwedenBankIdProviderOutput`

NewSwedenBankIdProviderOutput instantiates a new SwedenBankIdProviderOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwedenBankIdProviderOutputWithDefaults

`func NewSwedenBankIdProviderOutputWithDefaults() *SwedenBankIdProviderOutput`

NewSwedenBankIdProviderOutputWithDefaults instantiates a new SwedenBankIdProviderOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFullName

`func (o *SwedenBankIdProviderOutput) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *SwedenBankIdProviderOutput) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *SwedenBankIdProviderOutput) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *SwedenBankIdProviderOutput) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### SetFullNameNil

`func (o *SwedenBankIdProviderOutput) SetFullNameNil(b bool)`

 SetFullNameNil sets the value for FullName to be an explicit nil

### UnsetFullName
`func (o *SwedenBankIdProviderOutput) UnsetFullName()`

UnsetFullName ensures that no value is present for FullName, not even an explicit nil
### GetDateOfBirth

`func (o *SwedenBankIdProviderOutput) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *SwedenBankIdProviderOutput) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *SwedenBankIdProviderOutput) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *SwedenBankIdProviderOutput) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### SetDateOfBirthNil

`func (o *SwedenBankIdProviderOutput) SetDateOfBirthNil(b bool)`

 SetDateOfBirthNil sets the value for DateOfBirth to be an explicit nil

### UnsetDateOfBirth
`func (o *SwedenBankIdProviderOutput) UnsetDateOfBirth()`

UnsetDateOfBirth ensures that no value is present for DateOfBirth, not even an explicit nil
### GetPersonalIdentityNumber

`func (o *SwedenBankIdProviderOutput) GetPersonalIdentityNumber() string`

GetPersonalIdentityNumber returns the PersonalIdentityNumber field if non-nil, zero value otherwise.

### GetPersonalIdentityNumberOk

`func (o *SwedenBankIdProviderOutput) GetPersonalIdentityNumberOk() (*string, bool)`

GetPersonalIdentityNumberOk returns a tuple with the PersonalIdentityNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalIdentityNumber

`func (o *SwedenBankIdProviderOutput) SetPersonalIdentityNumber(v string)`

SetPersonalIdentityNumber sets PersonalIdentityNumber field to given value.

### HasPersonalIdentityNumber

`func (o *SwedenBankIdProviderOutput) HasPersonalIdentityNumber() bool`

HasPersonalIdentityNumber returns a boolean if a field has been set.

### SetPersonalIdentityNumberNil

`func (o *SwedenBankIdProviderOutput) SetPersonalIdentityNumberNil(b bool)`

 SetPersonalIdentityNumberNil sets the value for PersonalIdentityNumber to be an explicit nil

### UnsetPersonalIdentityNumber
`func (o *SwedenBankIdProviderOutput) UnsetPersonalIdentityNumber()`

UnsetPersonalIdentityNumber ensures that no value is present for PersonalIdentityNumber, not even an explicit nil
### GetGivenName

`func (o *SwedenBankIdProviderOutput) GetGivenName() string`

GetGivenName returns the GivenName field if non-nil, zero value otherwise.

### GetGivenNameOk

`func (o *SwedenBankIdProviderOutput) GetGivenNameOk() (*string, bool)`

GetGivenNameOk returns a tuple with the GivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenName

`func (o *SwedenBankIdProviderOutput) SetGivenName(v string)`

SetGivenName sets GivenName field to given value.

### HasGivenName

`func (o *SwedenBankIdProviderOutput) HasGivenName() bool`

HasGivenName returns a boolean if a field has been set.

### SetGivenNameNil

`func (o *SwedenBankIdProviderOutput) SetGivenNameNil(b bool)`

 SetGivenNameNil sets the value for GivenName to be an explicit nil

### UnsetGivenName
`func (o *SwedenBankIdProviderOutput) UnsetGivenName()`

UnsetGivenName ensures that no value is present for GivenName, not even an explicit nil
### GetFamilyName

`func (o *SwedenBankIdProviderOutput) GetFamilyName() string`

GetFamilyName returns the FamilyName field if non-nil, zero value otherwise.

### GetFamilyNameOk

`func (o *SwedenBankIdProviderOutput) GetFamilyNameOk() (*string, bool)`

GetFamilyNameOk returns a tuple with the FamilyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyName

`func (o *SwedenBankIdProviderOutput) SetFamilyName(v string)`

SetFamilyName sets FamilyName field to given value.

### HasFamilyName

`func (o *SwedenBankIdProviderOutput) HasFamilyName() bool`

HasFamilyName returns a boolean if a field has been set.

### SetFamilyNameNil

`func (o *SwedenBankIdProviderOutput) SetFamilyNameNil(b bool)`

 SetFamilyNameNil sets the value for FamilyName to be an explicit nil

### UnsetFamilyName
`func (o *SwedenBankIdProviderOutput) UnsetFamilyName()`

UnsetFamilyName ensures that no value is present for FamilyName, not even an explicit nil
### GetCountry

`func (o *SwedenBankIdProviderOutput) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *SwedenBankIdProviderOutput) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *SwedenBankIdProviderOutput) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *SwedenBankIdProviderOutput) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *SwedenBankIdProviderOutput) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *SwedenBankIdProviderOutput) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetCertificatePolicy

`func (o *SwedenBankIdProviderOutput) GetCertificatePolicy() string`

GetCertificatePolicy returns the CertificatePolicy field if non-nil, zero value otherwise.

### GetCertificatePolicyOk

`func (o *SwedenBankIdProviderOutput) GetCertificatePolicyOk() (*string, bool)`

GetCertificatePolicyOk returns a tuple with the CertificatePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificatePolicy

`func (o *SwedenBankIdProviderOutput) SetCertificatePolicy(v string)`

SetCertificatePolicy sets CertificatePolicy field to given value.

### HasCertificatePolicy

`func (o *SwedenBankIdProviderOutput) HasCertificatePolicy() bool`

HasCertificatePolicy returns a boolean if a field has been set.

### SetCertificatePolicyNil

`func (o *SwedenBankIdProviderOutput) SetCertificatePolicyNil(b bool)`

 SetCertificatePolicyNil sets the value for CertificatePolicy to be an explicit nil

### UnsetCertificatePolicy
`func (o *SwedenBankIdProviderOutput) UnsetCertificatePolicy()`

UnsetCertificatePolicy ensures that no value is present for CertificatePolicy, not even an explicit nil
### GetCommonName

`func (o *SwedenBankIdProviderOutput) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *SwedenBankIdProviderOutput) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *SwedenBankIdProviderOutput) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.

### HasCommonName

`func (o *SwedenBankIdProviderOutput) HasCommonName() bool`

HasCommonName returns a boolean if a field has been set.

### SetCommonNameNil

`func (o *SwedenBankIdProviderOutput) SetCommonNameNil(b bool)`

 SetCommonNameNil sets the value for CommonName to be an explicit nil

### UnsetCommonName
`func (o *SwedenBankIdProviderOutput) UnsetCommonName()`

UnsetCommonName ensures that no value is present for CommonName, not even an explicit nil
### GetDistinguishedName

`func (o *SwedenBankIdProviderOutput) GetDistinguishedName() string`

GetDistinguishedName returns the DistinguishedName field if non-nil, zero value otherwise.

### GetDistinguishedNameOk

`func (o *SwedenBankIdProviderOutput) GetDistinguishedNameOk() (*string, bool)`

GetDistinguishedNameOk returns a tuple with the DistinguishedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinguishedName

`func (o *SwedenBankIdProviderOutput) SetDistinguishedName(v string)`

SetDistinguishedName sets DistinguishedName field to given value.

### HasDistinguishedName

`func (o *SwedenBankIdProviderOutput) HasDistinguishedName() bool`

HasDistinguishedName returns a boolean if a field has been set.

### SetDistinguishedNameNil

`func (o *SwedenBankIdProviderOutput) SetDistinguishedNameNil(b bool)`

 SetDistinguishedNameNil sets the value for DistinguishedName to be an explicit nil

### UnsetDistinguishedName
`func (o *SwedenBankIdProviderOutput) UnsetDistinguishedName()`

UnsetDistinguishedName ensures that no value is present for DistinguishedName, not even an explicit nil
### GetPersonalIdentityNumberIssuingCountry

`func (o *SwedenBankIdProviderOutput) GetPersonalIdentityNumberIssuingCountry() string`

GetPersonalIdentityNumberIssuingCountry returns the PersonalIdentityNumberIssuingCountry field if non-nil, zero value otherwise.

### GetPersonalIdentityNumberIssuingCountryOk

`func (o *SwedenBankIdProviderOutput) GetPersonalIdentityNumberIssuingCountryOk() (*string, bool)`

GetPersonalIdentityNumberIssuingCountryOk returns a tuple with the PersonalIdentityNumberIssuingCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalIdentityNumberIssuingCountry

`func (o *SwedenBankIdProviderOutput) SetPersonalIdentityNumberIssuingCountry(v string)`

SetPersonalIdentityNumberIssuingCountry sets PersonalIdentityNumberIssuingCountry field to given value.

### HasPersonalIdentityNumberIssuingCountry

`func (o *SwedenBankIdProviderOutput) HasPersonalIdentityNumberIssuingCountry() bool`

HasPersonalIdentityNumberIssuingCountry returns a boolean if a field has been set.

### SetPersonalIdentityNumberIssuingCountryNil

`func (o *SwedenBankIdProviderOutput) SetPersonalIdentityNumberIssuingCountryNil(b bool)`

 SetPersonalIdentityNumberIssuingCountryNil sets the value for PersonalIdentityNumberIssuingCountry to be an explicit nil

### UnsetPersonalIdentityNumberIssuingCountry
`func (o *SwedenBankIdProviderOutput) UnsetPersonalIdentityNumberIssuingCountry()`

UnsetPersonalIdentityNumberIssuingCountry ensures that no value is present for PersonalIdentityNumberIssuingCountry, not even an explicit nil
### GetAuthenticationDeviceIp

`func (o *SwedenBankIdProviderOutput) GetAuthenticationDeviceIp() string`

GetAuthenticationDeviceIp returns the AuthenticationDeviceIp field if non-nil, zero value otherwise.

### GetAuthenticationDeviceIpOk

`func (o *SwedenBankIdProviderOutput) GetAuthenticationDeviceIpOk() (*string, bool)`

GetAuthenticationDeviceIpOk returns a tuple with the AuthenticationDeviceIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationDeviceIp

`func (o *SwedenBankIdProviderOutput) SetAuthenticationDeviceIp(v string)`

SetAuthenticationDeviceIp sets AuthenticationDeviceIp field to given value.

### HasAuthenticationDeviceIp

`func (o *SwedenBankIdProviderOutput) HasAuthenticationDeviceIp() bool`

HasAuthenticationDeviceIp returns a boolean if a field has been set.

### SetAuthenticationDeviceIpNil

`func (o *SwedenBankIdProviderOutput) SetAuthenticationDeviceIpNil(b bool)`

 SetAuthenticationDeviceIpNil sets the value for AuthenticationDeviceIp to be an explicit nil

### UnsetAuthenticationDeviceIp
`func (o *SwedenBankIdProviderOutput) UnsetAuthenticationDeviceIp()`

UnsetAuthenticationDeviceIp ensures that no value is present for AuthenticationDeviceIp, not even an explicit nil
### GetNotBefore

`func (o *SwedenBankIdProviderOutput) GetNotBefore() time.Time`

GetNotBefore returns the NotBefore field if non-nil, zero value otherwise.

### GetNotBeforeOk

`func (o *SwedenBankIdProviderOutput) GetNotBeforeOk() (*time.Time, bool)`

GetNotBeforeOk returns a tuple with the NotBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotBefore

`func (o *SwedenBankIdProviderOutput) SetNotBefore(v time.Time)`

SetNotBefore sets NotBefore field to given value.

### HasNotBefore

`func (o *SwedenBankIdProviderOutput) HasNotBefore() bool`

HasNotBefore returns a boolean if a field has been set.

### SetNotBeforeNil

`func (o *SwedenBankIdProviderOutput) SetNotBeforeNil(b bool)`

 SetNotBeforeNil sets the value for NotBefore to be an explicit nil

### UnsetNotBefore
`func (o *SwedenBankIdProviderOutput) UnsetNotBefore()`

UnsetNotBefore ensures that no value is present for NotBefore, not even an explicit nil
### GetNotAfter

`func (o *SwedenBankIdProviderOutput) GetNotAfter() time.Time`

GetNotAfter returns the NotAfter field if non-nil, zero value otherwise.

### GetNotAfterOk

`func (o *SwedenBankIdProviderOutput) GetNotAfterOk() (*time.Time, bool)`

GetNotAfterOk returns a tuple with the NotAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAfter

`func (o *SwedenBankIdProviderOutput) SetNotAfter(v time.Time)`

SetNotAfter sets NotAfter field to given value.

### HasNotAfter

`func (o *SwedenBankIdProviderOutput) HasNotAfter() bool`

HasNotAfter returns a boolean if a field has been set.

### SetNotAfterNil

`func (o *SwedenBankIdProviderOutput) SetNotAfterNil(b bool)`

 SetNotAfterNil sets the value for NotAfter to be an explicit nil

### UnsetNotAfter
`func (o *SwedenBankIdProviderOutput) UnsetNotAfter()`

UnsetNotAfter ensures that no value is present for NotAfter, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


