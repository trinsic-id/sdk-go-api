# IndonesiaDukcapilMatchProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NationalIdNumber** | [**IndonesiaDukcapilMatchNationalIdNumberField**](IndonesiaDukcapilMatchNationalIdNumberField.md) | NIK submitted for this Dukcapil match and the assessment result returned for that value. | 
**FullName** | [**IndonesiaDukcapilMatchFullNameField**](IndonesiaDukcapilMatchFullNameField.md) | Full name submitted for this Dukcapil match and the assessment result returned for that value. | 
**DateOfBirth** | [**IndonesiaDukcapilMatchDateOfBirthField**](IndonesiaDukcapilMatchDateOfBirthField.md) | Date of birth submitted for this Dukcapil match and the assessment result returned for that value. | 
**ProvinceCode** | **string** | Two-digit Indonesian government administrative region code for the province, extracted from digits 1-2 of the submitted NIK number.              Source system: Kode Wilayah Administrasi Pemerintahan, maintained by Indonesia&#39;s Ministry of Home Affairs. The first digit indicates the island group: 1-2 Sumatra, 3-4 Java, 5 Bali and Nusa Tenggara, 6 Kalimantan, 7 Sulawesi, 8 Maluku, and 9 Papua. The second digit follows the province creation order. | 
**RegencyOrCityCode** | **string** | Four-digit Indonesian government administrative region code for the regency or city, extracted from digits 1-4 of the submitted NIK number.              Source system: Kode Wilayah Administrasi Pemerintahan, maintained by Indonesia&#39;s Ministry of Home Affairs. Format: two-digit province code followed by a two-digit regency or city sequence; suffixes 01-69 identify regencies and suffixes 71-99 identify cities. The NIK stores this value without dot separators. | 
**DistrictCode** | **string** | Six-digit Indonesian government administrative region code for the district (kecamatan), extracted from digits 1-6 of the submitted NIK number.              Source system: Kode Wilayah Administrasi Pemerintahan, maintained by Indonesia&#39;s Ministry of Home Affairs. Format: two-digit province code, two-digit regency or city code, and two-digit district sequence. The NIK stores this value without dot separators. | 
**SexNationalIdNumber** | **string** | Sex extracted from the birth-day portion of the submitted NIK number.              Known values: - Male: The encoded day value is 40 or lower. - Female: The encoded day value is greater than 40. | 
**SerialNumber** | **string** | Four-digit issuance serial number extracted from digits 13-16 of the submitted NIK number. | 
**PhoneNumber** | Pointer to **NullableString** | The phone number submitted for this verification, if provided.              Format: international E.164 phone number. | [optional] 
**ConsentGivenAt** | **time.Time** | The consent timestamp submitted for this verification. | 
**Email** | Pointer to **NullableString** | The email address submitted for this verification, if provided. | [optional] 

## Methods

### NewIndonesiaDukcapilMatchProviderOutput

`func NewIndonesiaDukcapilMatchProviderOutput(nationalIdNumber IndonesiaDukcapilMatchNationalIdNumberField, fullName IndonesiaDukcapilMatchFullNameField, dateOfBirth IndonesiaDukcapilMatchDateOfBirthField, provinceCode string, regencyOrCityCode string, districtCode string, sexNationalIdNumber string, serialNumber string, consentGivenAt time.Time, ) *IndonesiaDukcapilMatchProviderOutput`

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


