# SaudiArabiaNafathProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NationalIdentityNumber** | Pointer to **NullableString** | Saudi Arabia national identity number used to authenticate the individual with Nafath.              Nafath supports three user types and the leading digit of this number identifies which one signed in: - 1: Saudi citizen national ID, issued by Civil Affairs (Al-Ahwal al-Madaniyya) under   the Ministry of Interior. - 2: Resident (Iqama or muqeem) number on the residence permit, issued by the General   Directorate of Passports (Al-Jawazat) under the Ministry of Interior. - 3 through 6: Visitor identity number (Hawiyyat Za&#39;ir) for foreign nationals admitted on   short-term visas, including tourist, business, family visit, and Umrah visas.              Format: - Exactly ten decimal digits. - First digit is the user type as listed above. - Digits two through nine are assigned serially and have no semantic meaning. - For citizen and resident numbers, the tenth digit is a check digit derived from the   first nine digits with a modulus-ten style algorithm. No check digit is documented for   visitor numbers, and the Nafath sandbox&#39;s visitor numbers do not follow it.              Nafath returns no identifier for visitors, so for them this is the number the user submitted. | [optional] 
**UserType** | Pointer to **NullableString** | Nafath user type derived from the leading digit of nationalIdentityNumber. The matching attributes object carries the fields Nafath defines only for that user type.              Possible values: - Citizen: Saudi citizen national ID (1-series). See citizenSpecificAttributes. - Resident: legally resident foreign national, Iqama (2-series). See   residentSpecificAttributes. - Visitor: visitor identity (3 through 6-series). See visitorSpecificAttributes. | [optional] 
**Sex** | Pointer to **NullableString** | Sex as returned by Nafath, normalized to the platform sex code.              Possible values: - Male - Female | [optional] 
**Nationality** | Pointer to **NullableString** | Nationality as an ISO 3166-1 alpha-2 country code.              Not every nationality passed by back Nafath corresponds to an ISO 3166-1 alpha-2 code, for example code 335 (Tribes Emigrated) or 336 (Turkestan). In these cases, this field is null. See nationalityCode for the raw nationality code. | [optional] 
**NationalityCode** | Pointer to **NullableString** | Nationality code as returned by Nafath under nationalityCode.              This is the Saudi government&#39;s three-digit nationality code, grouped by region: - 1xx Arab states - 2xx North Africa - 3xx Asia, Middle East, and Oceania - 4xx Sub-Saharan Africa - 5xx Europe - 6xx Americas - 7xx Oceania              It is not an ISO 3166 numeric code. For example, 504 is Ireland according to Nafath but Morocco in ISO 3166.              See https://docs.trinsic.id/docs/saudi-arabia-nafath for a full list. | [optional] 
**DateOfBirthGregorian** | Pointer to **NullableString** | Date of birth in the Gregorian calendar. | [optional] 
**OccupationCode** | Pointer to **NullableString** | Occupation code reported by Nafath.              Known values: - O: Other - G: Government | [optional] 
**PassportNumber** | Pointer to **NullableString** | Passport number reported by Nafath. | [optional] 
**CitizenSpecificAttributes** | Pointer to [**NullableSaudiArabiaNafathCitizenAttributes**](SaudiArabiaNafathCitizenAttributes.md) | Attributes Nafath defines only for Saudi citizens.              Present only when userType is Citizen. | [optional] 
**ResidentSpecificAttributes** | Pointer to [**NullableSaudiArabiaNafathResidentAttributes**](SaudiArabiaNafathResidentAttributes.md) | Attributes Nafath defines only for residents holding an Iqama.              Present only when userType is Resident. | [optional] 
**VisitorSpecificAttributes** | Pointer to [**NullableSaudiArabiaNafathVisitorAttributes**](SaudiArabiaNafathVisitorAttributes.md) | Attributes Nafath defines only for visitors on short-term visas.              Present only when userType is Visitor. | [optional] 

## Methods

### Named construction

Use a keyed struct literal (`SaudiArabiaNafathProviderOutput{FieldName: value}`), or call
`NewSaudiArabiaNafathProviderOutputWithDefaults()` and set fields by name. Positional field
factories and unkeyed struct literals are not supported.

### NewSaudiArabiaNafathProviderOutputWithDefaults

`func NewSaudiArabiaNafathProviderOutputWithDefaults() *SaudiArabiaNafathProviderOutput`

NewSaudiArabiaNafathProviderOutputWithDefaults instantiates a new SaudiArabiaNafathProviderOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNationalIdentityNumber

`func (o *SaudiArabiaNafathProviderOutput) GetNationalIdentityNumber() string`

GetNationalIdentityNumber returns the NationalIdentityNumber field if non-nil, zero value otherwise.

### GetNationalIdentityNumberOk

`func (o *SaudiArabiaNafathProviderOutput) GetNationalIdentityNumberOk() (*string, bool)`

GetNationalIdentityNumberOk returns a tuple with the NationalIdentityNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationalIdentityNumber

`func (o *SaudiArabiaNafathProviderOutput) SetNationalIdentityNumber(v string)`

SetNationalIdentityNumber sets NationalIdentityNumber field to given value.

### HasNationalIdentityNumber

`func (o *SaudiArabiaNafathProviderOutput) HasNationalIdentityNumber() bool`

HasNationalIdentityNumber returns a boolean if a field has been set.

### SetNationalIdentityNumberNil

`func (o *SaudiArabiaNafathProviderOutput) SetNationalIdentityNumberNil(b bool)`

 SetNationalIdentityNumberNil sets the value for NationalIdentityNumber to be an explicit nil

### UnsetNationalIdentityNumber
`func (o *SaudiArabiaNafathProviderOutput) UnsetNationalIdentityNumber()`

UnsetNationalIdentityNumber ensures that no value is present for NationalIdentityNumber, not even an explicit nil
### GetUserType

`func (o *SaudiArabiaNafathProviderOutput) GetUserType() string`

GetUserType returns the UserType field if non-nil, zero value otherwise.

### GetUserTypeOk

`func (o *SaudiArabiaNafathProviderOutput) GetUserTypeOk() (*string, bool)`

GetUserTypeOk returns a tuple with the UserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserType

`func (o *SaudiArabiaNafathProviderOutput) SetUserType(v string)`

SetUserType sets UserType field to given value.

### HasUserType

`func (o *SaudiArabiaNafathProviderOutput) HasUserType() bool`

HasUserType returns a boolean if a field has been set.

### SetUserTypeNil

`func (o *SaudiArabiaNafathProviderOutput) SetUserTypeNil(b bool)`

 SetUserTypeNil sets the value for UserType to be an explicit nil

### UnsetUserType
`func (o *SaudiArabiaNafathProviderOutput) UnsetUserType()`

UnsetUserType ensures that no value is present for UserType, not even an explicit nil
### GetSex

`func (o *SaudiArabiaNafathProviderOutput) GetSex() string`

GetSex returns the Sex field if non-nil, zero value otherwise.

### GetSexOk

`func (o *SaudiArabiaNafathProviderOutput) GetSexOk() (*string, bool)`

GetSexOk returns a tuple with the Sex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSex

`func (o *SaudiArabiaNafathProviderOutput) SetSex(v string)`

SetSex sets Sex field to given value.

### HasSex

`func (o *SaudiArabiaNafathProviderOutput) HasSex() bool`

HasSex returns a boolean if a field has been set.

### SetSexNil

`func (o *SaudiArabiaNafathProviderOutput) SetSexNil(b bool)`

 SetSexNil sets the value for Sex to be an explicit nil

### UnsetSex
`func (o *SaudiArabiaNafathProviderOutput) UnsetSex()`

UnsetSex ensures that no value is present for Sex, not even an explicit nil
### GetNationality

`func (o *SaudiArabiaNafathProviderOutput) GetNationality() string`

GetNationality returns the Nationality field if non-nil, zero value otherwise.

### GetNationalityOk

`func (o *SaudiArabiaNafathProviderOutput) GetNationalityOk() (*string, bool)`

GetNationalityOk returns a tuple with the Nationality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationality

`func (o *SaudiArabiaNafathProviderOutput) SetNationality(v string)`

SetNationality sets Nationality field to given value.

### HasNationality

`func (o *SaudiArabiaNafathProviderOutput) HasNationality() bool`

HasNationality returns a boolean if a field has been set.

### SetNationalityNil

`func (o *SaudiArabiaNafathProviderOutput) SetNationalityNil(b bool)`

 SetNationalityNil sets the value for Nationality to be an explicit nil

### UnsetNationality
`func (o *SaudiArabiaNafathProviderOutput) UnsetNationality()`

UnsetNationality ensures that no value is present for Nationality, not even an explicit nil
### GetNationalityCode

`func (o *SaudiArabiaNafathProviderOutput) GetNationalityCode() string`

GetNationalityCode returns the NationalityCode field if non-nil, zero value otherwise.

### GetNationalityCodeOk

`func (o *SaudiArabiaNafathProviderOutput) GetNationalityCodeOk() (*string, bool)`

GetNationalityCodeOk returns a tuple with the NationalityCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationalityCode

`func (o *SaudiArabiaNafathProviderOutput) SetNationalityCode(v string)`

SetNationalityCode sets NationalityCode field to given value.

### HasNationalityCode

`func (o *SaudiArabiaNafathProviderOutput) HasNationalityCode() bool`

HasNationalityCode returns a boolean if a field has been set.

### SetNationalityCodeNil

`func (o *SaudiArabiaNafathProviderOutput) SetNationalityCodeNil(b bool)`

 SetNationalityCodeNil sets the value for NationalityCode to be an explicit nil

### UnsetNationalityCode
`func (o *SaudiArabiaNafathProviderOutput) UnsetNationalityCode()`

UnsetNationalityCode ensures that no value is present for NationalityCode, not even an explicit nil
### GetDateOfBirthGregorian

`func (o *SaudiArabiaNafathProviderOutput) GetDateOfBirthGregorian() string`

GetDateOfBirthGregorian returns the DateOfBirthGregorian field if non-nil, zero value otherwise.

### GetDateOfBirthGregorianOk

`func (o *SaudiArabiaNafathProviderOutput) GetDateOfBirthGregorianOk() (*string, bool)`

GetDateOfBirthGregorianOk returns a tuple with the DateOfBirthGregorian field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirthGregorian

`func (o *SaudiArabiaNafathProviderOutput) SetDateOfBirthGregorian(v string)`

SetDateOfBirthGregorian sets DateOfBirthGregorian field to given value.

### HasDateOfBirthGregorian

`func (o *SaudiArabiaNafathProviderOutput) HasDateOfBirthGregorian() bool`

HasDateOfBirthGregorian returns a boolean if a field has been set.

### SetDateOfBirthGregorianNil

`func (o *SaudiArabiaNafathProviderOutput) SetDateOfBirthGregorianNil(b bool)`

 SetDateOfBirthGregorianNil sets the value for DateOfBirthGregorian to be an explicit nil

### UnsetDateOfBirthGregorian
`func (o *SaudiArabiaNafathProviderOutput) UnsetDateOfBirthGregorian()`

UnsetDateOfBirthGregorian ensures that no value is present for DateOfBirthGregorian, not even an explicit nil
### GetOccupationCode

`func (o *SaudiArabiaNafathProviderOutput) GetOccupationCode() string`

GetOccupationCode returns the OccupationCode field if non-nil, zero value otherwise.

### GetOccupationCodeOk

`func (o *SaudiArabiaNafathProviderOutput) GetOccupationCodeOk() (*string, bool)`

GetOccupationCodeOk returns a tuple with the OccupationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccupationCode

`func (o *SaudiArabiaNafathProviderOutput) SetOccupationCode(v string)`

SetOccupationCode sets OccupationCode field to given value.

### HasOccupationCode

`func (o *SaudiArabiaNafathProviderOutput) HasOccupationCode() bool`

HasOccupationCode returns a boolean if a field has been set.

### SetOccupationCodeNil

`func (o *SaudiArabiaNafathProviderOutput) SetOccupationCodeNil(b bool)`

 SetOccupationCodeNil sets the value for OccupationCode to be an explicit nil

### UnsetOccupationCode
`func (o *SaudiArabiaNafathProviderOutput) UnsetOccupationCode()`

UnsetOccupationCode ensures that no value is present for OccupationCode, not even an explicit nil
### GetPassportNumber

`func (o *SaudiArabiaNafathProviderOutput) GetPassportNumber() string`

GetPassportNumber returns the PassportNumber field if non-nil, zero value otherwise.

### GetPassportNumberOk

`func (o *SaudiArabiaNafathProviderOutput) GetPassportNumberOk() (*string, bool)`

GetPassportNumberOk returns a tuple with the PassportNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassportNumber

`func (o *SaudiArabiaNafathProviderOutput) SetPassportNumber(v string)`

SetPassportNumber sets PassportNumber field to given value.

### HasPassportNumber

`func (o *SaudiArabiaNafathProviderOutput) HasPassportNumber() bool`

HasPassportNumber returns a boolean if a field has been set.

### SetPassportNumberNil

`func (o *SaudiArabiaNafathProviderOutput) SetPassportNumberNil(b bool)`

 SetPassportNumberNil sets the value for PassportNumber to be an explicit nil

### UnsetPassportNumber
`func (o *SaudiArabiaNafathProviderOutput) UnsetPassportNumber()`

UnsetPassportNumber ensures that no value is present for PassportNumber, not even an explicit nil
### GetCitizenSpecificAttributes

`func (o *SaudiArabiaNafathProviderOutput) GetCitizenSpecificAttributes() SaudiArabiaNafathCitizenAttributes`

GetCitizenSpecificAttributes returns the CitizenSpecificAttributes field if non-nil, zero value otherwise.

### GetCitizenSpecificAttributesOk

`func (o *SaudiArabiaNafathProviderOutput) GetCitizenSpecificAttributesOk() (*SaudiArabiaNafathCitizenAttributes, bool)`

GetCitizenSpecificAttributesOk returns a tuple with the CitizenSpecificAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitizenSpecificAttributes

`func (o *SaudiArabiaNafathProviderOutput) SetCitizenSpecificAttributes(v SaudiArabiaNafathCitizenAttributes)`

SetCitizenSpecificAttributes sets CitizenSpecificAttributes field to given value.

### HasCitizenSpecificAttributes

`func (o *SaudiArabiaNafathProviderOutput) HasCitizenSpecificAttributes() bool`

HasCitizenSpecificAttributes returns a boolean if a field has been set.

### SetCitizenSpecificAttributesNil

`func (o *SaudiArabiaNafathProviderOutput) SetCitizenSpecificAttributesNil(b bool)`

 SetCitizenSpecificAttributesNil sets the value for CitizenSpecificAttributes to be an explicit nil

### UnsetCitizenSpecificAttributes
`func (o *SaudiArabiaNafathProviderOutput) UnsetCitizenSpecificAttributes()`

UnsetCitizenSpecificAttributes ensures that no value is present for CitizenSpecificAttributes, not even an explicit nil
### GetResidentSpecificAttributes

`func (o *SaudiArabiaNafathProviderOutput) GetResidentSpecificAttributes() SaudiArabiaNafathResidentAttributes`

GetResidentSpecificAttributes returns the ResidentSpecificAttributes field if non-nil, zero value otherwise.

### GetResidentSpecificAttributesOk

`func (o *SaudiArabiaNafathProviderOutput) GetResidentSpecificAttributesOk() (*SaudiArabiaNafathResidentAttributes, bool)`

GetResidentSpecificAttributesOk returns a tuple with the ResidentSpecificAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidentSpecificAttributes

`func (o *SaudiArabiaNafathProviderOutput) SetResidentSpecificAttributes(v SaudiArabiaNafathResidentAttributes)`

SetResidentSpecificAttributes sets ResidentSpecificAttributes field to given value.

### HasResidentSpecificAttributes

`func (o *SaudiArabiaNafathProviderOutput) HasResidentSpecificAttributes() bool`

HasResidentSpecificAttributes returns a boolean if a field has been set.

### SetResidentSpecificAttributesNil

`func (o *SaudiArabiaNafathProviderOutput) SetResidentSpecificAttributesNil(b bool)`

 SetResidentSpecificAttributesNil sets the value for ResidentSpecificAttributes to be an explicit nil

### UnsetResidentSpecificAttributes
`func (o *SaudiArabiaNafathProviderOutput) UnsetResidentSpecificAttributes()`

UnsetResidentSpecificAttributes ensures that no value is present for ResidentSpecificAttributes, not even an explicit nil
### GetVisitorSpecificAttributes

`func (o *SaudiArabiaNafathProviderOutput) GetVisitorSpecificAttributes() SaudiArabiaNafathVisitorAttributes`

GetVisitorSpecificAttributes returns the VisitorSpecificAttributes field if non-nil, zero value otherwise.

### GetVisitorSpecificAttributesOk

`func (o *SaudiArabiaNafathProviderOutput) GetVisitorSpecificAttributesOk() (*SaudiArabiaNafathVisitorAttributes, bool)`

GetVisitorSpecificAttributesOk returns a tuple with the VisitorSpecificAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisitorSpecificAttributes

`func (o *SaudiArabiaNafathProviderOutput) SetVisitorSpecificAttributes(v SaudiArabiaNafathVisitorAttributes)`

SetVisitorSpecificAttributes sets VisitorSpecificAttributes field to given value.

### HasVisitorSpecificAttributes

`func (o *SaudiArabiaNafathProviderOutput) HasVisitorSpecificAttributes() bool`

HasVisitorSpecificAttributes returns a boolean if a field has been set.

### SetVisitorSpecificAttributesNil

`func (o *SaudiArabiaNafathProviderOutput) SetVisitorSpecificAttributesNil(b bool)`

 SetVisitorSpecificAttributesNil sets the value for VisitorSpecificAttributes to be an explicit nil

### UnsetVisitorSpecificAttributes
`func (o *SaudiArabiaNafathProviderOutput) UnsetVisitorSpecificAttributes()`

UnsetVisitorSpecificAttributes ensures that no value is present for VisitorSpecificAttributes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


