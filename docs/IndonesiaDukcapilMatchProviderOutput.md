# IndonesiaDukcapilMatchProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NationalIdNumber** | Pointer to [**NullableIndonesiaDukcapilMatchNationalIdNumberField**](IndonesiaDukcapilMatchNationalIdNumberField.md) | NIK submitted for this Dukcapil match and the assessment result returned for that value. | [optional] 
**FullName** | Pointer to [**NullableIndonesiaDukcapilMatchFullNameField**](IndonesiaDukcapilMatchFullNameField.md) | Full name submitted for this Dukcapil match and the assessment result returned for that value. | [optional] 
**DateOfBirth** | Pointer to [**NullableIndonesiaDukcapilMatchDateOfBirthField**](IndonesiaDukcapilMatchDateOfBirthField.md) | Date of birth submitted for this Dukcapil match and the assessment result returned for that value. | [optional] 
**ProvinceCode** | Pointer to **NullableString** | Two-digit Indonesian government administrative region code for the province, extracted from digits 1-2 of the submitted NIK number.              Source system: Kode Wilayah Administrasi Pemerintahan, maintained by Indonesia&#39;s Ministry of Home Affairs. The first digit indicates the island group: 1-2 Sumatra, 3-4 Java, 5 Bali and Nusa Tenggara, 6 Kalimantan, 7 Sulawesi, 8 Maluku, and 9 Papua. The second digit follows the province creation order. | [optional] 
**RegencyOrCityCode** | Pointer to **NullableString** | Four-digit Indonesian government administrative region code for the regency or city, extracted from digits 1-4 of the submitted NIK number.              Source system: Kode Wilayah Administrasi Pemerintahan, maintained by Indonesia&#39;s Ministry of Home Affairs. Format: two-digit province code followed by a two-digit regency or city sequence; suffixes 01-69 identify regencies and suffixes 71-99 identify cities. The NIK stores this value without dot separators. | [optional] 
**DistrictCode** | Pointer to **NullableString** | Six-digit Indonesian government administrative region code for the district (kecamatan), extracted from digits 1-6 of the submitted NIK number.              Source system: Kode Wilayah Administrasi Pemerintahan, maintained by Indonesia&#39;s Ministry of Home Affairs. Format: two-digit province code, two-digit regency or city code, and two-digit district sequence. The NIK stores this value without dot separators. | [optional] 
**SexNationalIdNumber** | Pointer to **NullableString** | Sex extracted from the birth-day portion of the submitted NIK number.              Known values: - Male: The encoded day value is 40 or lower. - Female: The encoded day value is greater than 40. | [optional] 
**SerialNumber** | Pointer to **NullableString** | Four-digit issuance serial number extracted from digits 13-16 of the submitted NIK number. | [optional] 
**PhoneNumber** | Pointer to **NullableString** | The phone number submitted for this verification, if provided.              Format: international E.164 phone number. | [optional] 
**ConsentGivenAt** | Pointer to **NullableTime** | The consent timestamp submitted for this verification. | [optional] 
**Email** | Pointer to **NullableString** | The email address submitted for this verification, if provided. | [optional] 

## Methods

### NewIndonesiaDukcapilMatchProviderOutput

`func NewIndonesiaDukcapilMatchProviderOutput() *IndonesiaDukcapilMatchProviderOutput`

NewIndonesiaDukcapilMatchProviderOutput instantiates a new IndonesiaDukcapilMatchProviderOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndonesiaDukcapilMatchProviderOutputWithDefaults

`func NewIndonesiaDukcapilMatchProviderOutputWithDefaults() *IndonesiaDukcapilMatchProviderOutput`

NewIndonesiaDukcapilMatchProviderOutputWithDefaults instantiates a new IndonesiaDukcapilMatchProviderOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNationalIdNumber

`func (o *IndonesiaDukcapilMatchProviderOutput) GetNationalIdNumber() IndonesiaDukcapilMatchNationalIdNumberField`

GetNationalIdNumber returns the NationalIdNumber field if non-nil, zero value otherwise.

### GetNationalIdNumberOk

`func (o *IndonesiaDukcapilMatchProviderOutput) GetNationalIdNumberOk() (*IndonesiaDukcapilMatchNationalIdNumberField, bool)`

GetNationalIdNumberOk returns a tuple with the NationalIdNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationalIdNumber

`func (o *IndonesiaDukcapilMatchProviderOutput) SetNationalIdNumber(v IndonesiaDukcapilMatchNationalIdNumberField)`

SetNationalIdNumber sets NationalIdNumber field to given value.

### HasNationalIdNumber

`func (o *IndonesiaDukcapilMatchProviderOutput) HasNationalIdNumber() bool`

HasNationalIdNumber returns a boolean if a field has been set.

### SetNationalIdNumberNil

`func (o *IndonesiaDukcapilMatchProviderOutput) SetNationalIdNumberNil(b bool)`

 SetNationalIdNumberNil sets the value for NationalIdNumber to be an explicit nil

### UnsetNationalIdNumber
`func (o *IndonesiaDukcapilMatchProviderOutput) UnsetNationalIdNumber()`

UnsetNationalIdNumber ensures that no value is present for NationalIdNumber, not even an explicit nil
### GetFullName

`func (o *IndonesiaDukcapilMatchProviderOutput) GetFullName() IndonesiaDukcapilMatchFullNameField`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *IndonesiaDukcapilMatchProviderOutput) GetFullNameOk() (*IndonesiaDukcapilMatchFullNameField, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *IndonesiaDukcapilMatchProviderOutput) SetFullName(v IndonesiaDukcapilMatchFullNameField)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *IndonesiaDukcapilMatchProviderOutput) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### SetFullNameNil

`func (o *IndonesiaDukcapilMatchProviderOutput) SetFullNameNil(b bool)`

 SetFullNameNil sets the value for FullName to be an explicit nil

### UnsetFullName
`func (o *IndonesiaDukcapilMatchProviderOutput) UnsetFullName()`

UnsetFullName ensures that no value is present for FullName, not even an explicit nil
### GetDateOfBirth

`func (o *IndonesiaDukcapilMatchProviderOutput) GetDateOfBirth() IndonesiaDukcapilMatchDateOfBirthField`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *IndonesiaDukcapilMatchProviderOutput) GetDateOfBirthOk() (*IndonesiaDukcapilMatchDateOfBirthField, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *IndonesiaDukcapilMatchProviderOutput) SetDateOfBirth(v IndonesiaDukcapilMatchDateOfBirthField)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *IndonesiaDukcapilMatchProviderOutput) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### SetDateOfBirthNil

`func (o *IndonesiaDukcapilMatchProviderOutput) SetDateOfBirthNil(b bool)`

 SetDateOfBirthNil sets the value for DateOfBirth to be an explicit nil

### UnsetDateOfBirth
`func (o *IndonesiaDukcapilMatchProviderOutput) UnsetDateOfBirth()`

UnsetDateOfBirth ensures that no value is present for DateOfBirth, not even an explicit nil
### GetProvinceCode

`func (o *IndonesiaDukcapilMatchProviderOutput) GetProvinceCode() string`

GetProvinceCode returns the ProvinceCode field if non-nil, zero value otherwise.

### GetProvinceCodeOk

`func (o *IndonesiaDukcapilMatchProviderOutput) GetProvinceCodeOk() (*string, bool)`

GetProvinceCodeOk returns a tuple with the ProvinceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvinceCode

`func (o *IndonesiaDukcapilMatchProviderOutput) SetProvinceCode(v string)`

SetProvinceCode sets ProvinceCode field to given value.

### HasProvinceCode

`func (o *IndonesiaDukcapilMatchProviderOutput) HasProvinceCode() bool`

HasProvinceCode returns a boolean if a field has been set.

### SetProvinceCodeNil

`func (o *IndonesiaDukcapilMatchProviderOutput) SetProvinceCodeNil(b bool)`

 SetProvinceCodeNil sets the value for ProvinceCode to be an explicit nil

### UnsetProvinceCode
`func (o *IndonesiaDukcapilMatchProviderOutput) UnsetProvinceCode()`

UnsetProvinceCode ensures that no value is present for ProvinceCode, not even an explicit nil
### GetRegencyOrCityCode

`func (o *IndonesiaDukcapilMatchProviderOutput) GetRegencyOrCityCode() string`

GetRegencyOrCityCode returns the RegencyOrCityCode field if non-nil, zero value otherwise.

### GetRegencyOrCityCodeOk

`func (o *IndonesiaDukcapilMatchProviderOutput) GetRegencyOrCityCodeOk() (*string, bool)`

GetRegencyOrCityCodeOk returns a tuple with the RegencyOrCityCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegencyOrCityCode

`func (o *IndonesiaDukcapilMatchProviderOutput) SetRegencyOrCityCode(v string)`

SetRegencyOrCityCode sets RegencyOrCityCode field to given value.

### HasRegencyOrCityCode

`func (o *IndonesiaDukcapilMatchProviderOutput) HasRegencyOrCityCode() bool`

HasRegencyOrCityCode returns a boolean if a field has been set.

### SetRegencyOrCityCodeNil

`func (o *IndonesiaDukcapilMatchProviderOutput) SetRegencyOrCityCodeNil(b bool)`

 SetRegencyOrCityCodeNil sets the value for RegencyOrCityCode to be an explicit nil

### UnsetRegencyOrCityCode
`func (o *IndonesiaDukcapilMatchProviderOutput) UnsetRegencyOrCityCode()`

UnsetRegencyOrCityCode ensures that no value is present for RegencyOrCityCode, not even an explicit nil
### GetDistrictCode

`func (o *IndonesiaDukcapilMatchProviderOutput) GetDistrictCode() string`

GetDistrictCode returns the DistrictCode field if non-nil, zero value otherwise.

### GetDistrictCodeOk

`func (o *IndonesiaDukcapilMatchProviderOutput) GetDistrictCodeOk() (*string, bool)`

GetDistrictCodeOk returns a tuple with the DistrictCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistrictCode

`func (o *IndonesiaDukcapilMatchProviderOutput) SetDistrictCode(v string)`

SetDistrictCode sets DistrictCode field to given value.

### HasDistrictCode

`func (o *IndonesiaDukcapilMatchProviderOutput) HasDistrictCode() bool`

HasDistrictCode returns a boolean if a field has been set.

### SetDistrictCodeNil

`func (o *IndonesiaDukcapilMatchProviderOutput) SetDistrictCodeNil(b bool)`

 SetDistrictCodeNil sets the value for DistrictCode to be an explicit nil

### UnsetDistrictCode
`func (o *IndonesiaDukcapilMatchProviderOutput) UnsetDistrictCode()`

UnsetDistrictCode ensures that no value is present for DistrictCode, not even an explicit nil
### GetSexNationalIdNumber

`func (o *IndonesiaDukcapilMatchProviderOutput) GetSexNationalIdNumber() string`

GetSexNationalIdNumber returns the SexNationalIdNumber field if non-nil, zero value otherwise.

### GetSexNationalIdNumberOk

`func (o *IndonesiaDukcapilMatchProviderOutput) GetSexNationalIdNumberOk() (*string, bool)`

GetSexNationalIdNumberOk returns a tuple with the SexNationalIdNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSexNationalIdNumber

`func (o *IndonesiaDukcapilMatchProviderOutput) SetSexNationalIdNumber(v string)`

SetSexNationalIdNumber sets SexNationalIdNumber field to given value.

### HasSexNationalIdNumber

`func (o *IndonesiaDukcapilMatchProviderOutput) HasSexNationalIdNumber() bool`

HasSexNationalIdNumber returns a boolean if a field has been set.

### SetSexNationalIdNumberNil

`func (o *IndonesiaDukcapilMatchProviderOutput) SetSexNationalIdNumberNil(b bool)`

 SetSexNationalIdNumberNil sets the value for SexNationalIdNumber to be an explicit nil

### UnsetSexNationalIdNumber
`func (o *IndonesiaDukcapilMatchProviderOutput) UnsetSexNationalIdNumber()`

UnsetSexNationalIdNumber ensures that no value is present for SexNationalIdNumber, not even an explicit nil
### GetSerialNumber

`func (o *IndonesiaDukcapilMatchProviderOutput) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *IndonesiaDukcapilMatchProviderOutput) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *IndonesiaDukcapilMatchProviderOutput) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *IndonesiaDukcapilMatchProviderOutput) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### SetSerialNumberNil

`func (o *IndonesiaDukcapilMatchProviderOutput) SetSerialNumberNil(b bool)`

 SetSerialNumberNil sets the value for SerialNumber to be an explicit nil

### UnsetSerialNumber
`func (o *IndonesiaDukcapilMatchProviderOutput) UnsetSerialNumber()`

UnsetSerialNumber ensures that no value is present for SerialNumber, not even an explicit nil
### GetPhoneNumber

`func (o *IndonesiaDukcapilMatchProviderOutput) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *IndonesiaDukcapilMatchProviderOutput) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *IndonesiaDukcapilMatchProviderOutput) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *IndonesiaDukcapilMatchProviderOutput) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### SetPhoneNumberNil

`func (o *IndonesiaDukcapilMatchProviderOutput) SetPhoneNumberNil(b bool)`

 SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil

### UnsetPhoneNumber
`func (o *IndonesiaDukcapilMatchProviderOutput) UnsetPhoneNumber()`

UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
### GetConsentGivenAt

`func (o *IndonesiaDukcapilMatchProviderOutput) GetConsentGivenAt() time.Time`

GetConsentGivenAt returns the ConsentGivenAt field if non-nil, zero value otherwise.

### GetConsentGivenAtOk

`func (o *IndonesiaDukcapilMatchProviderOutput) GetConsentGivenAtOk() (*time.Time, bool)`

GetConsentGivenAtOk returns a tuple with the ConsentGivenAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentGivenAt

`func (o *IndonesiaDukcapilMatchProviderOutput) SetConsentGivenAt(v time.Time)`

SetConsentGivenAt sets ConsentGivenAt field to given value.

### HasConsentGivenAt

`func (o *IndonesiaDukcapilMatchProviderOutput) HasConsentGivenAt() bool`

HasConsentGivenAt returns a boolean if a field has been set.

### SetConsentGivenAtNil

`func (o *IndonesiaDukcapilMatchProviderOutput) SetConsentGivenAtNil(b bool)`

 SetConsentGivenAtNil sets the value for ConsentGivenAt to be an explicit nil

### UnsetConsentGivenAt
`func (o *IndonesiaDukcapilMatchProviderOutput) UnsetConsentGivenAt()`

UnsetConsentGivenAt ensures that no value is present for ConsentGivenAt, not even an explicit nil
### GetEmail

`func (o *IndonesiaDukcapilMatchProviderOutput) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *IndonesiaDukcapilMatchProviderOutput) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *IndonesiaDukcapilMatchProviderOutput) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *IndonesiaDukcapilMatchProviderOutput) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *IndonesiaDukcapilMatchProviderOutput) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *IndonesiaDukcapilMatchProviderOutput) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


