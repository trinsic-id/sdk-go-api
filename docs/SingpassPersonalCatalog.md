# SingpassPersonalCatalog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdentityNumber** | Pointer to **NullableString** | The Unique Identifier Number (UINFIN) for the user.              The number can be either a National Registration Identity Card or a Foreign Identification Number (NRIC/FIN). Follows the format @xxxxxxx#: - @ is the status of the holder.     - Singapore citizens and permanent residents born before 1 January 2000 are assigned the letter \&quot;S\&quot;.     - Singapore citizens and permanent residents born on or after 1 January 2000 are assigned the letter \&quot;T\&quot;.     - Foreigners issued with long-term passes before 1 January 2000 are assigned the letter \&quot;F\&quot;.     - Foreigners issued with long-term passes from 1 January 2000 to 31 December 2021 are assigned the letter \&quot;G\&quot;.     - Foreigners issued with long-term passes on or after 1 January 2022 are assigned the letter \&quot;M\&quot;. - xxxxxxx is seven digit serial number. - # is the checksum letter. | [optional] 
**PartialIdentityNumber** | Pointer to **NullableString** | The partial National Registration Identity Card or Foreign Identification Number (NRIC/FIN) shown by Singpass.              The value hides all but the last three digits and checksum. | [optional] 
**Name** | Pointer to **NullableString** | The name of verified individual composed of all the possible names.              The name is assembled in the following format (where applicable): 1. Principal Name (The first name section) 2. Hanyu Pinyin Name 3. Alias Name 4. Hanyu Pinyin Alias Name 5. Married Name (optional): Available only for females, prefixed with \&quot;MRS\&quot; | [optional] 
**AliasName** | Pointer to **NullableString** | The individual&#39;s alternate legally recognized name. | [optional] 
**HanyuPinyinName** | Pointer to **NullableString** | The individual&#39;s Hanyu Pinyin name. | [optional] 
**HanyuPinyinAliasName** | Pointer to **NullableString** | The individual&#39;s Hanyu Pinyin alias name. | [optional] 
**MarriedName** | Pointer to **NullableString** | The individual&#39;s married name.              This appears when Singpass has a legally recognized married surname on file. | [optional] 
**Email** | Pointer to **NullableString** | The user&#39;s email address. | [optional] 
**MobileNumber** | Pointer to [**NullableSingpassPhoneNumber**](SingpassPhoneNumber.md) | The user&#39;s mobile number, split into prefix, area code, and subscriber number. | [optional] 
**Dialect** | Pointer to [**NullableSingpassMyInfoDescription**](SingpassMyInfoDescription.md) | The user&#39;s dialect. | [optional] 
**DateOfBirth** | Pointer to **NullableString** | The user&#39;s date of birth. | [optional] 
**ResidentialStatus** | Pointer to [**NullableSingpassMyInfoDescription**](SingpassMyInfoDescription.md) | The user&#39;s residential status.              Possible values (code description): - A Alien - C Citizen - P PR - U Unknown - N NOT APPLICABLE | [optional] 
**Nationality** | Pointer to **NullableString** | The user&#39;s nationality in IS0-2 country code format. | [optional] 
**BirthCountry** | Pointer to **NullableString** | The user&#39;s birth country in IS0-2 country code format. | [optional] 
**PassportNumber** | Pointer to **NullableString** | The user&#39;s passport number. | [optional] 
**PassportExpiryDate** | Pointer to **NullableString** | The user&#39;s passport expiry date. | [optional] 
**PassType** | Pointer to [**NullableSingpassMyInfoDescription**](SingpassMyInfoDescription.md) | The Pass type of the individual.              Possible values: (code description)              - RPass Work Permit - SPass S Pass - P1Pass Employment Pass - P2Pass Employment Pass - QPass Employment Pass - PEP Personalised Employment Pass - WHP Work Holiday Pass - TEP Training Employment Pass - Entre EntrePass - OVE Overseas Networks &amp; Expertise Pass - DP Dependent Pass - LTVP Long-Term Visit Pass - LOC Letter of Consent - MWP Miscellaneous Work Pass - STP Student&#39;s Pass - LTVP+ Long-Term Visit Pass Plus - IEO Immigration Exemption Order              Note that this only applies to a FIN holder. | [optional] 
**PassStatus** | Pointer to **NullableString** | Pass status.              Possible values: - Live - Approved (Interim status in which the FIN holder has yet to receive the pass)              Note that this only applies to a FIN holder. | [optional] 
**PassExpiryDate** | Pointer to **NullableString** | Pass expiration date.              Note that this only applies to a FIN holder. | [optional] 
**EmploymentSector** | Pointer to **NullableString** | The user&#39;s employment sector. | [optional] 
**RegisteredAddress** | Pointer to [**NullableSingpassMyInfoRegisteredAddressOutput**](SingpassMyInfoRegisteredAddressOutput.md) | The user&#39;s registered address. | [optional] 
**HouseDevelopmenetBoardType** | Pointer to [**NullableSingpassMyInfoDescription**](SingpassMyInfoDescription.md) | Housing and Development Board (HBD) flat type code and description.              This value will be null if not HBD Possible values (code description): - 111 1-ROOM FLAT (HDB) - 112 2-ROOM FLAT (HDB) - 113 3-ROOM FLAT (HDB) - 114 4-ROOM FLAT (HDB) - 115 5-ROOM FLAT (HDB) - 116 EXECUTIVE FLAT (HDB) - 118 STUDIO APARTMENT (HDB) | [optional] 
**HousingType** | Pointer to [**NullableSingpassMyInfoDescription**](SingpassMyInfoDescription.md) | The user&#39;s housing type.              Possible values (code description): - 121 DETACHED HOUSE - 122 SEMI-DETACHED HOUSE - 123 TERRACE HOUSE - 131 CONDOMINIUM - 132 EXECUTIVE CONDOMINIUM - 139 APARTMENT | [optional] 
**Sex** | Pointer to **NullableString** | Sex of the individual.              Possible values: - Female - Male - Unknown | [optional] 

## Methods

### NewSingpassPersonalCatalog

`func NewSingpassPersonalCatalog() *SingpassPersonalCatalog`

NewSingpassPersonalCatalog instantiates a new SingpassPersonalCatalog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSingpassPersonalCatalogWithDefaults

`func NewSingpassPersonalCatalogWithDefaults() *SingpassPersonalCatalog`

NewSingpassPersonalCatalogWithDefaults instantiates a new SingpassPersonalCatalog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentityNumber

`func (o *SingpassPersonalCatalog) GetIdentityNumber() string`

GetIdentityNumber returns the IdentityNumber field if non-nil, zero value otherwise.

### GetIdentityNumberOk

`func (o *SingpassPersonalCatalog) GetIdentityNumberOk() (*string, bool)`

GetIdentityNumberOk returns a tuple with the IdentityNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityNumber

`func (o *SingpassPersonalCatalog) SetIdentityNumber(v string)`

SetIdentityNumber sets IdentityNumber field to given value.

### HasIdentityNumber

`func (o *SingpassPersonalCatalog) HasIdentityNumber() bool`

HasIdentityNumber returns a boolean if a field has been set.

### SetIdentityNumberNil

`func (o *SingpassPersonalCatalog) SetIdentityNumberNil(b bool)`

 SetIdentityNumberNil sets the value for IdentityNumber to be an explicit nil

### UnsetIdentityNumber
`func (o *SingpassPersonalCatalog) UnsetIdentityNumber()`

UnsetIdentityNumber ensures that no value is present for IdentityNumber, not even an explicit nil
### GetPartialIdentityNumber

`func (o *SingpassPersonalCatalog) GetPartialIdentityNumber() string`

GetPartialIdentityNumber returns the PartialIdentityNumber field if non-nil, zero value otherwise.

### GetPartialIdentityNumberOk

`func (o *SingpassPersonalCatalog) GetPartialIdentityNumberOk() (*string, bool)`

GetPartialIdentityNumberOk returns a tuple with the PartialIdentityNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartialIdentityNumber

`func (o *SingpassPersonalCatalog) SetPartialIdentityNumber(v string)`

SetPartialIdentityNumber sets PartialIdentityNumber field to given value.

### HasPartialIdentityNumber

`func (o *SingpassPersonalCatalog) HasPartialIdentityNumber() bool`

HasPartialIdentityNumber returns a boolean if a field has been set.

### SetPartialIdentityNumberNil

`func (o *SingpassPersonalCatalog) SetPartialIdentityNumberNil(b bool)`

 SetPartialIdentityNumberNil sets the value for PartialIdentityNumber to be an explicit nil

### UnsetPartialIdentityNumber
`func (o *SingpassPersonalCatalog) UnsetPartialIdentityNumber()`

UnsetPartialIdentityNumber ensures that no value is present for PartialIdentityNumber, not even an explicit nil
### GetName

`func (o *SingpassPersonalCatalog) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SingpassPersonalCatalog) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SingpassPersonalCatalog) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SingpassPersonalCatalog) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *SingpassPersonalCatalog) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *SingpassPersonalCatalog) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetAliasName

`func (o *SingpassPersonalCatalog) GetAliasName() string`

GetAliasName returns the AliasName field if non-nil, zero value otherwise.

### GetAliasNameOk

`func (o *SingpassPersonalCatalog) GetAliasNameOk() (*string, bool)`

GetAliasNameOk returns a tuple with the AliasName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasName

`func (o *SingpassPersonalCatalog) SetAliasName(v string)`

SetAliasName sets AliasName field to given value.

### HasAliasName

`func (o *SingpassPersonalCatalog) HasAliasName() bool`

HasAliasName returns a boolean if a field has been set.

### SetAliasNameNil

`func (o *SingpassPersonalCatalog) SetAliasNameNil(b bool)`

 SetAliasNameNil sets the value for AliasName to be an explicit nil

### UnsetAliasName
`func (o *SingpassPersonalCatalog) UnsetAliasName()`

UnsetAliasName ensures that no value is present for AliasName, not even an explicit nil
### GetHanyuPinyinName

`func (o *SingpassPersonalCatalog) GetHanyuPinyinName() string`

GetHanyuPinyinName returns the HanyuPinyinName field if non-nil, zero value otherwise.

### GetHanyuPinyinNameOk

`func (o *SingpassPersonalCatalog) GetHanyuPinyinNameOk() (*string, bool)`

GetHanyuPinyinNameOk returns a tuple with the HanyuPinyinName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHanyuPinyinName

`func (o *SingpassPersonalCatalog) SetHanyuPinyinName(v string)`

SetHanyuPinyinName sets HanyuPinyinName field to given value.

### HasHanyuPinyinName

`func (o *SingpassPersonalCatalog) HasHanyuPinyinName() bool`

HasHanyuPinyinName returns a boolean if a field has been set.

### SetHanyuPinyinNameNil

`func (o *SingpassPersonalCatalog) SetHanyuPinyinNameNil(b bool)`

 SetHanyuPinyinNameNil sets the value for HanyuPinyinName to be an explicit nil

### UnsetHanyuPinyinName
`func (o *SingpassPersonalCatalog) UnsetHanyuPinyinName()`

UnsetHanyuPinyinName ensures that no value is present for HanyuPinyinName, not even an explicit nil
### GetHanyuPinyinAliasName

`func (o *SingpassPersonalCatalog) GetHanyuPinyinAliasName() string`

GetHanyuPinyinAliasName returns the HanyuPinyinAliasName field if non-nil, zero value otherwise.

### GetHanyuPinyinAliasNameOk

`func (o *SingpassPersonalCatalog) GetHanyuPinyinAliasNameOk() (*string, bool)`

GetHanyuPinyinAliasNameOk returns a tuple with the HanyuPinyinAliasName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHanyuPinyinAliasName

`func (o *SingpassPersonalCatalog) SetHanyuPinyinAliasName(v string)`

SetHanyuPinyinAliasName sets HanyuPinyinAliasName field to given value.

### HasHanyuPinyinAliasName

`func (o *SingpassPersonalCatalog) HasHanyuPinyinAliasName() bool`

HasHanyuPinyinAliasName returns a boolean if a field has been set.

### SetHanyuPinyinAliasNameNil

`func (o *SingpassPersonalCatalog) SetHanyuPinyinAliasNameNil(b bool)`

 SetHanyuPinyinAliasNameNil sets the value for HanyuPinyinAliasName to be an explicit nil

### UnsetHanyuPinyinAliasName
`func (o *SingpassPersonalCatalog) UnsetHanyuPinyinAliasName()`

UnsetHanyuPinyinAliasName ensures that no value is present for HanyuPinyinAliasName, not even an explicit nil
### GetMarriedName

`func (o *SingpassPersonalCatalog) GetMarriedName() string`

GetMarriedName returns the MarriedName field if non-nil, zero value otherwise.

### GetMarriedNameOk

`func (o *SingpassPersonalCatalog) GetMarriedNameOk() (*string, bool)`

GetMarriedNameOk returns a tuple with the MarriedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarriedName

`func (o *SingpassPersonalCatalog) SetMarriedName(v string)`

SetMarriedName sets MarriedName field to given value.

### HasMarriedName

`func (o *SingpassPersonalCatalog) HasMarriedName() bool`

HasMarriedName returns a boolean if a field has been set.

### SetMarriedNameNil

`func (o *SingpassPersonalCatalog) SetMarriedNameNil(b bool)`

 SetMarriedNameNil sets the value for MarriedName to be an explicit nil

### UnsetMarriedName
`func (o *SingpassPersonalCatalog) UnsetMarriedName()`

UnsetMarriedName ensures that no value is present for MarriedName, not even an explicit nil
### GetEmail

`func (o *SingpassPersonalCatalog) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *SingpassPersonalCatalog) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *SingpassPersonalCatalog) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *SingpassPersonalCatalog) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *SingpassPersonalCatalog) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *SingpassPersonalCatalog) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetMobileNumber

`func (o *SingpassPersonalCatalog) GetMobileNumber() SingpassPhoneNumber`

GetMobileNumber returns the MobileNumber field if non-nil, zero value otherwise.

### GetMobileNumberOk

`func (o *SingpassPersonalCatalog) GetMobileNumberOk() (*SingpassPhoneNumber, bool)`

GetMobileNumberOk returns a tuple with the MobileNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileNumber

`func (o *SingpassPersonalCatalog) SetMobileNumber(v SingpassPhoneNumber)`

SetMobileNumber sets MobileNumber field to given value.

### HasMobileNumber

`func (o *SingpassPersonalCatalog) HasMobileNumber() bool`

HasMobileNumber returns a boolean if a field has been set.

### SetMobileNumberNil

`func (o *SingpassPersonalCatalog) SetMobileNumberNil(b bool)`

 SetMobileNumberNil sets the value for MobileNumber to be an explicit nil

### UnsetMobileNumber
`func (o *SingpassPersonalCatalog) UnsetMobileNumber()`

UnsetMobileNumber ensures that no value is present for MobileNumber, not even an explicit nil
### GetDialect

`func (o *SingpassPersonalCatalog) GetDialect() SingpassMyInfoDescription`

GetDialect returns the Dialect field if non-nil, zero value otherwise.

### GetDialectOk

`func (o *SingpassPersonalCatalog) GetDialectOk() (*SingpassMyInfoDescription, bool)`

GetDialectOk returns a tuple with the Dialect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDialect

`func (o *SingpassPersonalCatalog) SetDialect(v SingpassMyInfoDescription)`

SetDialect sets Dialect field to given value.

### HasDialect

`func (o *SingpassPersonalCatalog) HasDialect() bool`

HasDialect returns a boolean if a field has been set.

### SetDialectNil

`func (o *SingpassPersonalCatalog) SetDialectNil(b bool)`

 SetDialectNil sets the value for Dialect to be an explicit nil

### UnsetDialect
`func (o *SingpassPersonalCatalog) UnsetDialect()`

UnsetDialect ensures that no value is present for Dialect, not even an explicit nil
### GetDateOfBirth

`func (o *SingpassPersonalCatalog) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *SingpassPersonalCatalog) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *SingpassPersonalCatalog) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *SingpassPersonalCatalog) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### SetDateOfBirthNil

`func (o *SingpassPersonalCatalog) SetDateOfBirthNil(b bool)`

 SetDateOfBirthNil sets the value for DateOfBirth to be an explicit nil

### UnsetDateOfBirth
`func (o *SingpassPersonalCatalog) UnsetDateOfBirth()`

UnsetDateOfBirth ensures that no value is present for DateOfBirth, not even an explicit nil
### GetResidentialStatus

`func (o *SingpassPersonalCatalog) GetResidentialStatus() SingpassMyInfoDescription`

GetResidentialStatus returns the ResidentialStatus field if non-nil, zero value otherwise.

### GetResidentialStatusOk

`func (o *SingpassPersonalCatalog) GetResidentialStatusOk() (*SingpassMyInfoDescription, bool)`

GetResidentialStatusOk returns a tuple with the ResidentialStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidentialStatus

`func (o *SingpassPersonalCatalog) SetResidentialStatus(v SingpassMyInfoDescription)`

SetResidentialStatus sets ResidentialStatus field to given value.

### HasResidentialStatus

`func (o *SingpassPersonalCatalog) HasResidentialStatus() bool`

HasResidentialStatus returns a boolean if a field has been set.

### SetResidentialStatusNil

`func (o *SingpassPersonalCatalog) SetResidentialStatusNil(b bool)`

 SetResidentialStatusNil sets the value for ResidentialStatus to be an explicit nil

### UnsetResidentialStatus
`func (o *SingpassPersonalCatalog) UnsetResidentialStatus()`

UnsetResidentialStatus ensures that no value is present for ResidentialStatus, not even an explicit nil
### GetNationality

`func (o *SingpassPersonalCatalog) GetNationality() string`

GetNationality returns the Nationality field if non-nil, zero value otherwise.

### GetNationalityOk

`func (o *SingpassPersonalCatalog) GetNationalityOk() (*string, bool)`

GetNationalityOk returns a tuple with the Nationality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationality

`func (o *SingpassPersonalCatalog) SetNationality(v string)`

SetNationality sets Nationality field to given value.

### HasNationality

`func (o *SingpassPersonalCatalog) HasNationality() bool`

HasNationality returns a boolean if a field has been set.

### SetNationalityNil

`func (o *SingpassPersonalCatalog) SetNationalityNil(b bool)`

 SetNationalityNil sets the value for Nationality to be an explicit nil

### UnsetNationality
`func (o *SingpassPersonalCatalog) UnsetNationality()`

UnsetNationality ensures that no value is present for Nationality, not even an explicit nil
### GetBirthCountry

`func (o *SingpassPersonalCatalog) GetBirthCountry() string`

GetBirthCountry returns the BirthCountry field if non-nil, zero value otherwise.

### GetBirthCountryOk

`func (o *SingpassPersonalCatalog) GetBirthCountryOk() (*string, bool)`

GetBirthCountryOk returns a tuple with the BirthCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthCountry

`func (o *SingpassPersonalCatalog) SetBirthCountry(v string)`

SetBirthCountry sets BirthCountry field to given value.

### HasBirthCountry

`func (o *SingpassPersonalCatalog) HasBirthCountry() bool`

HasBirthCountry returns a boolean if a field has been set.

### SetBirthCountryNil

`func (o *SingpassPersonalCatalog) SetBirthCountryNil(b bool)`

 SetBirthCountryNil sets the value for BirthCountry to be an explicit nil

### UnsetBirthCountry
`func (o *SingpassPersonalCatalog) UnsetBirthCountry()`

UnsetBirthCountry ensures that no value is present for BirthCountry, not even an explicit nil
### GetPassportNumber

`func (o *SingpassPersonalCatalog) GetPassportNumber() string`

GetPassportNumber returns the PassportNumber field if non-nil, zero value otherwise.

### GetPassportNumberOk

`func (o *SingpassPersonalCatalog) GetPassportNumberOk() (*string, bool)`

GetPassportNumberOk returns a tuple with the PassportNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassportNumber

`func (o *SingpassPersonalCatalog) SetPassportNumber(v string)`

SetPassportNumber sets PassportNumber field to given value.

### HasPassportNumber

`func (o *SingpassPersonalCatalog) HasPassportNumber() bool`

HasPassportNumber returns a boolean if a field has been set.

### SetPassportNumberNil

`func (o *SingpassPersonalCatalog) SetPassportNumberNil(b bool)`

 SetPassportNumberNil sets the value for PassportNumber to be an explicit nil

### UnsetPassportNumber
`func (o *SingpassPersonalCatalog) UnsetPassportNumber()`

UnsetPassportNumber ensures that no value is present for PassportNumber, not even an explicit nil
### GetPassportExpiryDate

`func (o *SingpassPersonalCatalog) GetPassportExpiryDate() string`

GetPassportExpiryDate returns the PassportExpiryDate field if non-nil, zero value otherwise.

### GetPassportExpiryDateOk

`func (o *SingpassPersonalCatalog) GetPassportExpiryDateOk() (*string, bool)`

GetPassportExpiryDateOk returns a tuple with the PassportExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassportExpiryDate

`func (o *SingpassPersonalCatalog) SetPassportExpiryDate(v string)`

SetPassportExpiryDate sets PassportExpiryDate field to given value.

### HasPassportExpiryDate

`func (o *SingpassPersonalCatalog) HasPassportExpiryDate() bool`

HasPassportExpiryDate returns a boolean if a field has been set.

### SetPassportExpiryDateNil

`func (o *SingpassPersonalCatalog) SetPassportExpiryDateNil(b bool)`

 SetPassportExpiryDateNil sets the value for PassportExpiryDate to be an explicit nil

### UnsetPassportExpiryDate
`func (o *SingpassPersonalCatalog) UnsetPassportExpiryDate()`

UnsetPassportExpiryDate ensures that no value is present for PassportExpiryDate, not even an explicit nil
### GetPassType

`func (o *SingpassPersonalCatalog) GetPassType() SingpassMyInfoDescription`

GetPassType returns the PassType field if non-nil, zero value otherwise.

### GetPassTypeOk

`func (o *SingpassPersonalCatalog) GetPassTypeOk() (*SingpassMyInfoDescription, bool)`

GetPassTypeOk returns a tuple with the PassType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassType

`func (o *SingpassPersonalCatalog) SetPassType(v SingpassMyInfoDescription)`

SetPassType sets PassType field to given value.

### HasPassType

`func (o *SingpassPersonalCatalog) HasPassType() bool`

HasPassType returns a boolean if a field has been set.

### SetPassTypeNil

`func (o *SingpassPersonalCatalog) SetPassTypeNil(b bool)`

 SetPassTypeNil sets the value for PassType to be an explicit nil

### UnsetPassType
`func (o *SingpassPersonalCatalog) UnsetPassType()`

UnsetPassType ensures that no value is present for PassType, not even an explicit nil
### GetPassStatus

`func (o *SingpassPersonalCatalog) GetPassStatus() string`

GetPassStatus returns the PassStatus field if non-nil, zero value otherwise.

### GetPassStatusOk

`func (o *SingpassPersonalCatalog) GetPassStatusOk() (*string, bool)`

GetPassStatusOk returns a tuple with the PassStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassStatus

`func (o *SingpassPersonalCatalog) SetPassStatus(v string)`

SetPassStatus sets PassStatus field to given value.

### HasPassStatus

`func (o *SingpassPersonalCatalog) HasPassStatus() bool`

HasPassStatus returns a boolean if a field has been set.

### SetPassStatusNil

`func (o *SingpassPersonalCatalog) SetPassStatusNil(b bool)`

 SetPassStatusNil sets the value for PassStatus to be an explicit nil

### UnsetPassStatus
`func (o *SingpassPersonalCatalog) UnsetPassStatus()`

UnsetPassStatus ensures that no value is present for PassStatus, not even an explicit nil
### GetPassExpiryDate

`func (o *SingpassPersonalCatalog) GetPassExpiryDate() string`

GetPassExpiryDate returns the PassExpiryDate field if non-nil, zero value otherwise.

### GetPassExpiryDateOk

`func (o *SingpassPersonalCatalog) GetPassExpiryDateOk() (*string, bool)`

GetPassExpiryDateOk returns a tuple with the PassExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassExpiryDate

`func (o *SingpassPersonalCatalog) SetPassExpiryDate(v string)`

SetPassExpiryDate sets PassExpiryDate field to given value.

### HasPassExpiryDate

`func (o *SingpassPersonalCatalog) HasPassExpiryDate() bool`

HasPassExpiryDate returns a boolean if a field has been set.

### SetPassExpiryDateNil

`func (o *SingpassPersonalCatalog) SetPassExpiryDateNil(b bool)`

 SetPassExpiryDateNil sets the value for PassExpiryDate to be an explicit nil

### UnsetPassExpiryDate
`func (o *SingpassPersonalCatalog) UnsetPassExpiryDate()`

UnsetPassExpiryDate ensures that no value is present for PassExpiryDate, not even an explicit nil
### GetEmploymentSector

`func (o *SingpassPersonalCatalog) GetEmploymentSector() string`

GetEmploymentSector returns the EmploymentSector field if non-nil, zero value otherwise.

### GetEmploymentSectorOk

`func (o *SingpassPersonalCatalog) GetEmploymentSectorOk() (*string, bool)`

GetEmploymentSectorOk returns a tuple with the EmploymentSector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmploymentSector

`func (o *SingpassPersonalCatalog) SetEmploymentSector(v string)`

SetEmploymentSector sets EmploymentSector field to given value.

### HasEmploymentSector

`func (o *SingpassPersonalCatalog) HasEmploymentSector() bool`

HasEmploymentSector returns a boolean if a field has been set.

### SetEmploymentSectorNil

`func (o *SingpassPersonalCatalog) SetEmploymentSectorNil(b bool)`

 SetEmploymentSectorNil sets the value for EmploymentSector to be an explicit nil

### UnsetEmploymentSector
`func (o *SingpassPersonalCatalog) UnsetEmploymentSector()`

UnsetEmploymentSector ensures that no value is present for EmploymentSector, not even an explicit nil
### GetRegisteredAddress

`func (o *SingpassPersonalCatalog) GetRegisteredAddress() SingpassMyInfoRegisteredAddressOutput`

GetRegisteredAddress returns the RegisteredAddress field if non-nil, zero value otherwise.

### GetRegisteredAddressOk

`func (o *SingpassPersonalCatalog) GetRegisteredAddressOk() (*SingpassMyInfoRegisteredAddressOutput, bool)`

GetRegisteredAddressOk returns a tuple with the RegisteredAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredAddress

`func (o *SingpassPersonalCatalog) SetRegisteredAddress(v SingpassMyInfoRegisteredAddressOutput)`

SetRegisteredAddress sets RegisteredAddress field to given value.

### HasRegisteredAddress

`func (o *SingpassPersonalCatalog) HasRegisteredAddress() bool`

HasRegisteredAddress returns a boolean if a field has been set.

### SetRegisteredAddressNil

`func (o *SingpassPersonalCatalog) SetRegisteredAddressNil(b bool)`

 SetRegisteredAddressNil sets the value for RegisteredAddress to be an explicit nil

### UnsetRegisteredAddress
`func (o *SingpassPersonalCatalog) UnsetRegisteredAddress()`

UnsetRegisteredAddress ensures that no value is present for RegisteredAddress, not even an explicit nil
### GetHouseDevelopmenetBoardType

`func (o *SingpassPersonalCatalog) GetHouseDevelopmenetBoardType() SingpassMyInfoDescription`

GetHouseDevelopmenetBoardType returns the HouseDevelopmenetBoardType field if non-nil, zero value otherwise.

### GetHouseDevelopmenetBoardTypeOk

`func (o *SingpassPersonalCatalog) GetHouseDevelopmenetBoardTypeOk() (*SingpassMyInfoDescription, bool)`

GetHouseDevelopmenetBoardTypeOk returns a tuple with the HouseDevelopmenetBoardType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHouseDevelopmenetBoardType

`func (o *SingpassPersonalCatalog) SetHouseDevelopmenetBoardType(v SingpassMyInfoDescription)`

SetHouseDevelopmenetBoardType sets HouseDevelopmenetBoardType field to given value.

### HasHouseDevelopmenetBoardType

`func (o *SingpassPersonalCatalog) HasHouseDevelopmenetBoardType() bool`

HasHouseDevelopmenetBoardType returns a boolean if a field has been set.

### SetHouseDevelopmenetBoardTypeNil

`func (o *SingpassPersonalCatalog) SetHouseDevelopmenetBoardTypeNil(b bool)`

 SetHouseDevelopmenetBoardTypeNil sets the value for HouseDevelopmenetBoardType to be an explicit nil

### UnsetHouseDevelopmenetBoardType
`func (o *SingpassPersonalCatalog) UnsetHouseDevelopmenetBoardType()`

UnsetHouseDevelopmenetBoardType ensures that no value is present for HouseDevelopmenetBoardType, not even an explicit nil
### GetHousingType

`func (o *SingpassPersonalCatalog) GetHousingType() SingpassMyInfoDescription`

GetHousingType returns the HousingType field if non-nil, zero value otherwise.

### GetHousingTypeOk

`func (o *SingpassPersonalCatalog) GetHousingTypeOk() (*SingpassMyInfoDescription, bool)`

GetHousingTypeOk returns a tuple with the HousingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHousingType

`func (o *SingpassPersonalCatalog) SetHousingType(v SingpassMyInfoDescription)`

SetHousingType sets HousingType field to given value.

### HasHousingType

`func (o *SingpassPersonalCatalog) HasHousingType() bool`

HasHousingType returns a boolean if a field has been set.

### SetHousingTypeNil

`func (o *SingpassPersonalCatalog) SetHousingTypeNil(b bool)`

 SetHousingTypeNil sets the value for HousingType to be an explicit nil

### UnsetHousingType
`func (o *SingpassPersonalCatalog) UnsetHousingType()`

UnsetHousingType ensures that no value is present for HousingType, not even an explicit nil
### GetSex

`func (o *SingpassPersonalCatalog) GetSex() string`

GetSex returns the Sex field if non-nil, zero value otherwise.

### GetSexOk

`func (o *SingpassPersonalCatalog) GetSexOk() (*string, bool)`

GetSexOk returns a tuple with the Sex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSex

`func (o *SingpassPersonalCatalog) SetSex(v string)`

SetSex sets Sex field to given value.

### HasSex

`func (o *SingpassPersonalCatalog) HasSex() bool`

HasSex returns a boolean if a field has been set.

### SetSexNil

`func (o *SingpassPersonalCatalog) SetSexNil(b bool)`

 SetSexNil sets the value for Sex to be an explicit nil

### UnsetSex
`func (o *SingpassPersonalCatalog) UnsetSex()`

UnsetSex ensures that no value is present for Sex, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


