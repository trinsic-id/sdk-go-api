# IndonesiaNikMatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FullName** | Pointer to [**NullableIndonesiaNikMatchField**](IndonesiaNikMatchField.md) | Full name match result from Indonesia NIK verification. | [optional] 
**DateOfBirth** | Pointer to [**NullableIndonesiaNikMatchField**](IndonesiaNikMatchField.md) | Date of birth match result from Indonesia NIK verification. | [optional] 
**NationalIdNumber** | Pointer to [**NullableIndonesiaNikMatchField**](IndonesiaNikMatchField.md) | NIK number match result from Indonesia NIK verification.              NIK stands for Nomor Induk Kependudukan. It is Indonesia&#39;s unique population identity number, issued by Indonesia&#39;s population administration and civil registration authority (Dukcapil) under the Ministry of Home Affairs. A NIK has 16 digits: - Digits 1-2 are the province code. - Digits 3-4 are the regency or city code within that province. - Digits 5-6 are the district code within that regency or city. - Digits 7-12 encode date of birth as DDMMYY. For female NIK holders, the day is increased by 40. - Digits 13-16 are an issuance serial number. | [optional] 
**ProvinceCode** | Pointer to **NullableString** | Two-digit province code extracted from digits 1-2 of the NIK number. | [optional] 
**RegencyOrCityCode** | Pointer to **NullableString** | Four-digit regency or city code extracted from digits 1-4 of the NIK number, including the province code prefix. | [optional] 
**DistrictCode** | Pointer to **NullableString** | Six-digit district code extracted from digits 1-6 of the NIK number, including province and regency/city code prefixes. | [optional] 
**DateOfBirthNationalIdNumber** | Pointer to **NullableString** | Date of birth extracted from digits 7-12 of the NIK number.              The NIK encodes date of birth as DDMMYY. For female NIK holders, the encoded day is increased by 40 before being stored in the NIK number. | [optional] 
**SexNationalIdNumber** | Pointer to **NullableString** | Sex extracted from the birth-day portion of the NIK number.              The encoded day value is increased by 40 for female NIK holders. | [optional] 
**SerialNumber** | Pointer to **NullableString** | Four-digit issuance serial number extracted from digits 13-16 of the NIK number. | [optional] 

## Methods

### NewIndonesiaNikMatch

`func NewIndonesiaNikMatch() *IndonesiaNikMatch`

NewIndonesiaNikMatch instantiates a new IndonesiaNikMatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndonesiaNikMatchWithDefaults

`func NewIndonesiaNikMatchWithDefaults() *IndonesiaNikMatch`

NewIndonesiaNikMatchWithDefaults instantiates a new IndonesiaNikMatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFullName

`func (o *IndonesiaNikMatch) GetFullName() IndonesiaNikMatchField`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *IndonesiaNikMatch) GetFullNameOk() (*IndonesiaNikMatchField, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *IndonesiaNikMatch) SetFullName(v IndonesiaNikMatchField)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *IndonesiaNikMatch) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### SetFullNameNil

`func (o *IndonesiaNikMatch) SetFullNameNil(b bool)`

 SetFullNameNil sets the value for FullName to be an explicit nil

### UnsetFullName
`func (o *IndonesiaNikMatch) UnsetFullName()`

UnsetFullName ensures that no value is present for FullName, not even an explicit nil
### GetDateOfBirth

`func (o *IndonesiaNikMatch) GetDateOfBirth() IndonesiaNikMatchField`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *IndonesiaNikMatch) GetDateOfBirthOk() (*IndonesiaNikMatchField, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *IndonesiaNikMatch) SetDateOfBirth(v IndonesiaNikMatchField)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *IndonesiaNikMatch) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### SetDateOfBirthNil

`func (o *IndonesiaNikMatch) SetDateOfBirthNil(b bool)`

 SetDateOfBirthNil sets the value for DateOfBirth to be an explicit nil

### UnsetDateOfBirth
`func (o *IndonesiaNikMatch) UnsetDateOfBirth()`

UnsetDateOfBirth ensures that no value is present for DateOfBirth, not even an explicit nil
### GetNationalIdNumber

`func (o *IndonesiaNikMatch) GetNationalIdNumber() IndonesiaNikMatchField`

GetNationalIdNumber returns the NationalIdNumber field if non-nil, zero value otherwise.

### GetNationalIdNumberOk

`func (o *IndonesiaNikMatch) GetNationalIdNumberOk() (*IndonesiaNikMatchField, bool)`

GetNationalIdNumberOk returns a tuple with the NationalIdNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationalIdNumber

`func (o *IndonesiaNikMatch) SetNationalIdNumber(v IndonesiaNikMatchField)`

SetNationalIdNumber sets NationalIdNumber field to given value.

### HasNationalIdNumber

`func (o *IndonesiaNikMatch) HasNationalIdNumber() bool`

HasNationalIdNumber returns a boolean if a field has been set.

### SetNationalIdNumberNil

`func (o *IndonesiaNikMatch) SetNationalIdNumberNil(b bool)`

 SetNationalIdNumberNil sets the value for NationalIdNumber to be an explicit nil

### UnsetNationalIdNumber
`func (o *IndonesiaNikMatch) UnsetNationalIdNumber()`

UnsetNationalIdNumber ensures that no value is present for NationalIdNumber, not even an explicit nil
### GetProvinceCode

`func (o *IndonesiaNikMatch) GetProvinceCode() string`

GetProvinceCode returns the ProvinceCode field if non-nil, zero value otherwise.

### GetProvinceCodeOk

`func (o *IndonesiaNikMatch) GetProvinceCodeOk() (*string, bool)`

GetProvinceCodeOk returns a tuple with the ProvinceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvinceCode

`func (o *IndonesiaNikMatch) SetProvinceCode(v string)`

SetProvinceCode sets ProvinceCode field to given value.

### HasProvinceCode

`func (o *IndonesiaNikMatch) HasProvinceCode() bool`

HasProvinceCode returns a boolean if a field has been set.

### SetProvinceCodeNil

`func (o *IndonesiaNikMatch) SetProvinceCodeNil(b bool)`

 SetProvinceCodeNil sets the value for ProvinceCode to be an explicit nil

### UnsetProvinceCode
`func (o *IndonesiaNikMatch) UnsetProvinceCode()`

UnsetProvinceCode ensures that no value is present for ProvinceCode, not even an explicit nil
### GetRegencyOrCityCode

`func (o *IndonesiaNikMatch) GetRegencyOrCityCode() string`

GetRegencyOrCityCode returns the RegencyOrCityCode field if non-nil, zero value otherwise.

### GetRegencyOrCityCodeOk

`func (o *IndonesiaNikMatch) GetRegencyOrCityCodeOk() (*string, bool)`

GetRegencyOrCityCodeOk returns a tuple with the RegencyOrCityCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegencyOrCityCode

`func (o *IndonesiaNikMatch) SetRegencyOrCityCode(v string)`

SetRegencyOrCityCode sets RegencyOrCityCode field to given value.

### HasRegencyOrCityCode

`func (o *IndonesiaNikMatch) HasRegencyOrCityCode() bool`

HasRegencyOrCityCode returns a boolean if a field has been set.

### SetRegencyOrCityCodeNil

`func (o *IndonesiaNikMatch) SetRegencyOrCityCodeNil(b bool)`

 SetRegencyOrCityCodeNil sets the value for RegencyOrCityCode to be an explicit nil

### UnsetRegencyOrCityCode
`func (o *IndonesiaNikMatch) UnsetRegencyOrCityCode()`

UnsetRegencyOrCityCode ensures that no value is present for RegencyOrCityCode, not even an explicit nil
### GetDistrictCode

`func (o *IndonesiaNikMatch) GetDistrictCode() string`

GetDistrictCode returns the DistrictCode field if non-nil, zero value otherwise.

### GetDistrictCodeOk

`func (o *IndonesiaNikMatch) GetDistrictCodeOk() (*string, bool)`

GetDistrictCodeOk returns a tuple with the DistrictCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistrictCode

`func (o *IndonesiaNikMatch) SetDistrictCode(v string)`

SetDistrictCode sets DistrictCode field to given value.

### HasDistrictCode

`func (o *IndonesiaNikMatch) HasDistrictCode() bool`

HasDistrictCode returns a boolean if a field has been set.

### SetDistrictCodeNil

`func (o *IndonesiaNikMatch) SetDistrictCodeNil(b bool)`

 SetDistrictCodeNil sets the value for DistrictCode to be an explicit nil

### UnsetDistrictCode
`func (o *IndonesiaNikMatch) UnsetDistrictCode()`

UnsetDistrictCode ensures that no value is present for DistrictCode, not even an explicit nil
### GetDateOfBirthNationalIdNumber

`func (o *IndonesiaNikMatch) GetDateOfBirthNationalIdNumber() string`

GetDateOfBirthNationalIdNumber returns the DateOfBirthNationalIdNumber field if non-nil, zero value otherwise.

### GetDateOfBirthNationalIdNumberOk

`func (o *IndonesiaNikMatch) GetDateOfBirthNationalIdNumberOk() (*string, bool)`

GetDateOfBirthNationalIdNumberOk returns a tuple with the DateOfBirthNationalIdNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirthNationalIdNumber

`func (o *IndonesiaNikMatch) SetDateOfBirthNationalIdNumber(v string)`

SetDateOfBirthNationalIdNumber sets DateOfBirthNationalIdNumber field to given value.

### HasDateOfBirthNationalIdNumber

`func (o *IndonesiaNikMatch) HasDateOfBirthNationalIdNumber() bool`

HasDateOfBirthNationalIdNumber returns a boolean if a field has been set.

### SetDateOfBirthNationalIdNumberNil

`func (o *IndonesiaNikMatch) SetDateOfBirthNationalIdNumberNil(b bool)`

 SetDateOfBirthNationalIdNumberNil sets the value for DateOfBirthNationalIdNumber to be an explicit nil

### UnsetDateOfBirthNationalIdNumber
`func (o *IndonesiaNikMatch) UnsetDateOfBirthNationalIdNumber()`

UnsetDateOfBirthNationalIdNumber ensures that no value is present for DateOfBirthNationalIdNumber, not even an explicit nil
### GetSexNationalIdNumber

`func (o *IndonesiaNikMatch) GetSexNationalIdNumber() string`

GetSexNationalIdNumber returns the SexNationalIdNumber field if non-nil, zero value otherwise.

### GetSexNationalIdNumberOk

`func (o *IndonesiaNikMatch) GetSexNationalIdNumberOk() (*string, bool)`

GetSexNationalIdNumberOk returns a tuple with the SexNationalIdNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSexNationalIdNumber

`func (o *IndonesiaNikMatch) SetSexNationalIdNumber(v string)`

SetSexNationalIdNumber sets SexNationalIdNumber field to given value.

### HasSexNationalIdNumber

`func (o *IndonesiaNikMatch) HasSexNationalIdNumber() bool`

HasSexNationalIdNumber returns a boolean if a field has been set.

### SetSexNationalIdNumberNil

`func (o *IndonesiaNikMatch) SetSexNationalIdNumberNil(b bool)`

 SetSexNationalIdNumberNil sets the value for SexNationalIdNumber to be an explicit nil

### UnsetSexNationalIdNumber
`func (o *IndonesiaNikMatch) UnsetSexNationalIdNumber()`

UnsetSexNationalIdNumber ensures that no value is present for SexNationalIdNumber, not even an explicit nil
### GetSerialNumber

`func (o *IndonesiaNikMatch) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *IndonesiaNikMatch) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *IndonesiaNikMatch) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *IndonesiaNikMatch) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### SetSerialNumberNil

`func (o *IndonesiaNikMatch) SetSerialNumberNil(b bool)`

 SetSerialNumberNil sets the value for SerialNumber to be an explicit nil

### UnsetSerialNumber
`func (o *IndonesiaNikMatch) UnsetSerialNumber()`

UnsetSerialNumber ensures that no value is present for SerialNumber, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


