# SaudiArabiaNafathCitizenAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArabicFirstName** | Pointer to **NullableString** | Given name in Arabic.              Nafath always returns this in Arabic. romanizedFirstName is the English form, and the session locale does not change either name. | [optional] 
**ArabicFatherName** | Pointer to **NullableString** | Father&#39;s given name in Arabic. Same script as arabicFirstName. | [optional] 
**ArabicGrandFatherName** | Pointer to **NullableString** | Grandfather&#39;s given name in Arabic. Same script as arabicFirstName. | [optional] 
**ArabicFamilyName** | Pointer to **NullableString** | Family name in Arabic. Same script as arabicFirstName. | [optional] 
**RomanizedFirstName** | Pointer to **NullableString** | Given name in English.              Nafath always returns this in English, including when the session locale is Arabic. | [optional] 
**RomanizedSecondName** | Pointer to **NullableString** | Father&#39;s given name in English. Same script as romanizedFirstName. | [optional] 
**RomanizedThirdName** | Pointer to **NullableString** | Grandfather&#39;s given name in English. Same script as romanizedFirstName. | [optional] 
**RomanizedLastName** | Pointer to **NullableString** | Family name in English. Same script as romanizedFirstName. | [optional] 
**DateOfBirthHijri** | Pointer to **NullableString** | Date of birth in the Hijri (Umm al-Qura) calendar, in YYYY-MM-DD format. | [optional] 
**IdIssueDateGregorian** | Pointer to **NullableString** | Date the national ID card was issued in the Gregorian calendar. | [optional] 
**IdIssueDateHijri** | Pointer to **NullableString** | Date the national ID card was issued in the Hijri calendar, in YYYY-MM-DD format. | [optional] 
**IdExpiryDateGregorian** | Pointer to **NullableString** | Date the national ID card expires in the Gregorian calendar. | [optional] 
**IdExpiryDateHijri** | Pointer to **NullableString** | Date the national ID card expires in the Hijri calendar, in YYYY-MM-DD format. | [optional] 
**IdVersionNumber** | Pointer to **NullableString** | Version number printed on the national ID card.              Represents the generation of national ID card this document was issued in. | [optional] 
**IdIssuePlace** | Pointer to **NullableString** | City or office where the national ID card was issued, in Arabic.              The session locale does not change it. | [optional] 
**SocialStatusCode** | Pointer to **NullableString** | Social-status (marital-status) code, from the Nafath socialStatusCode attribute. The numeric code is not unique on its own. The same code maps to different states depending on the value of the top-level sex field.              Code by sex: - 0, M: Single - 1, F: Single - 1, M: One wife - 2, F: Married - 2, M: Two wives - 3, F: Divorced - 3, M: Three wives - 4, F: Widow - 4, M: Four wives - 5, F: Pending divorced              Combinations not listed are not defined by Nafath. | [optional] 
**SocialStatus** | Pointer to **NullableString** | Social status description in Arabic, from the Nafath socialStatusDesc attribute.              The session locale does not change it. socialStatusCode is the separate numeric code. | [optional] 
**DependentCount** | Pointer to **NullableInt32** | Number of current dependents. | [optional] 
**PlaceOfBirth** | Pointer to **NullableString** | Place of birth in Arabic.              The session locale does not change it. | [optional] 
**LastTravelExitDateGregorian** | Pointer to **NullableString** | Most recent travel exit date in the Gregorian calendar. | [optional] 
**LastTravelExitDateHijri** | Pointer to **NullableString** | Most recent travel exit date in the Hijri calendar, in YYYY-MM-DD format. | [optional] 
**LastTravelEntryDateGregorian** | Pointer to **NullableString** | Most recent travel entry date in the Gregorian calendar. | [optional] 
**LastTravelEntryDateHijri** | Pointer to **NullableString** | Most recent travel entry date in the Hijri calendar, in YYYY-MM-DD format. | [optional] 
**IsMinor** | Pointer to **NullableBool** | True when the citizen is a minor. | [optional] 
**HeadOfHousehold** | Pointer to [**NullableSaudiArabiaNafathHeadOfHousehold**](SaudiArabiaNafathHeadOfHousehold.md) | Head of household information, intended for minor citizen workflows such as guardian consent for opening a bank account.              The sandbox also returns it for adults, so it is included whenever present. | [optional] 
**NationalAddresses** | Pointer to [**[]SaudiArabiaNafathAddress**](SaudiArabiaNafathAddress.md) | Registered Saudi national addresses, as structured fields.              Saudi Post (SPL) documents how to enter a national address on checkout and international forms: Line 1 is building number, unit number, and street name. Line 2 is the additional number and district. Then city and postal code.              The country is always Saudi Arabia because Nafath only ever returns Saudi national addresses.              https://splonline.com.sa/en/door-step/ | [optional] 

## Methods

### Named construction

Use a keyed struct literal (`SaudiArabiaNafathCitizenAttributes{FieldName: value}`), or call
`NewSaudiArabiaNafathCitizenAttributesWithDefaults()` and set fields by name. Positional field
factories and unkeyed struct literals are not supported.

### NewSaudiArabiaNafathCitizenAttributesWithDefaults

`func NewSaudiArabiaNafathCitizenAttributesWithDefaults() *SaudiArabiaNafathCitizenAttributes`

NewSaudiArabiaNafathCitizenAttributesWithDefaults instantiates a new SaudiArabiaNafathCitizenAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArabicFirstName

`func (o *SaudiArabiaNafathCitizenAttributes) GetArabicFirstName() string`

GetArabicFirstName returns the ArabicFirstName field if non-nil, zero value otherwise.

### GetArabicFirstNameOk

`func (o *SaudiArabiaNafathCitizenAttributes) GetArabicFirstNameOk() (*string, bool)`

GetArabicFirstNameOk returns a tuple with the ArabicFirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArabicFirstName

`func (o *SaudiArabiaNafathCitizenAttributes) SetArabicFirstName(v string)`

SetArabicFirstName sets ArabicFirstName field to given value.

### HasArabicFirstName

`func (o *SaudiArabiaNafathCitizenAttributes) HasArabicFirstName() bool`

HasArabicFirstName returns a boolean if a field has been set.

### SetArabicFirstNameNil

`func (o *SaudiArabiaNafathCitizenAttributes) SetArabicFirstNameNil(b bool)`

 SetArabicFirstNameNil sets the value for ArabicFirstName to be an explicit nil

### UnsetArabicFirstName
`func (o *SaudiArabiaNafathCitizenAttributes) UnsetArabicFirstName()`

UnsetArabicFirstName ensures that no value is present for ArabicFirstName, not even an explicit nil
### GetArabicFatherName

`func (o *SaudiArabiaNafathCitizenAttributes) GetArabicFatherName() string`

GetArabicFatherName returns the ArabicFatherName field if non-nil, zero value otherwise.

### GetArabicFatherNameOk

`func (o *SaudiArabiaNafathCitizenAttributes) GetArabicFatherNameOk() (*string, bool)`

GetArabicFatherNameOk returns a tuple with the ArabicFatherName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArabicFatherName

`func (o *SaudiArabiaNafathCitizenAttributes) SetArabicFatherName(v string)`

SetArabicFatherName sets ArabicFatherName field to given value.

### HasArabicFatherName

`func (o *SaudiArabiaNafathCitizenAttributes) HasArabicFatherName() bool`

HasArabicFatherName returns a boolean if a field has been set.

### SetArabicFatherNameNil

`func (o *SaudiArabiaNafathCitizenAttributes) SetArabicFatherNameNil(b bool)`

 SetArabicFatherNameNil sets the value for ArabicFatherName to be an explicit nil

### UnsetArabicFatherName
`func (o *SaudiArabiaNafathCitizenAttributes) UnsetArabicFatherName()`

UnsetArabicFatherName ensures that no value is present for ArabicFatherName, not even an explicit nil
### GetArabicGrandFatherName

`func (o *SaudiArabiaNafathCitizenAttributes) GetArabicGrandFatherName() string`

GetArabicGrandFatherName returns the ArabicGrandFatherName field if non-nil, zero value otherwise.

### GetArabicGrandFatherNameOk

`func (o *SaudiArabiaNafathCitizenAttributes) GetArabicGrandFatherNameOk() (*string, bool)`

GetArabicGrandFatherNameOk returns a tuple with the ArabicGrandFatherName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArabicGrandFatherName

`func (o *SaudiArabiaNafathCitizenAttributes) SetArabicGrandFatherName(v string)`

SetArabicGrandFatherName sets ArabicGrandFatherName field to given value.

### HasArabicGrandFatherName

`func (o *SaudiArabiaNafathCitizenAttributes) HasArabicGrandFatherName() bool`

HasArabicGrandFatherName returns a boolean if a field has been set.

### SetArabicGrandFatherNameNil

`func (o *SaudiArabiaNafathCitizenAttributes) SetArabicGrandFatherNameNil(b bool)`

 SetArabicGrandFatherNameNil sets the value for ArabicGrandFatherName to be an explicit nil

### UnsetArabicGrandFatherName
`func (o *SaudiArabiaNafathCitizenAttributes) UnsetArabicGrandFatherName()`

UnsetArabicGrandFatherName ensures that no value is present for ArabicGrandFatherName, not even an explicit nil
### GetArabicFamilyName

`func (o *SaudiArabiaNafathCitizenAttributes) GetArabicFamilyName() string`

GetArabicFamilyName returns the ArabicFamilyName field if non-nil, zero value otherwise.

### GetArabicFamilyNameOk

`func (o *SaudiArabiaNafathCitizenAttributes) GetArabicFamilyNameOk() (*string, bool)`

GetArabicFamilyNameOk returns a tuple with the ArabicFamilyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArabicFamilyName

`func (o *SaudiArabiaNafathCitizenAttributes) SetArabicFamilyName(v string)`

SetArabicFamilyName sets ArabicFamilyName field to given value.

### HasArabicFamilyName

`func (o *SaudiArabiaNafathCitizenAttributes) HasArabicFamilyName() bool`

HasArabicFamilyName returns a boolean if a field has been set.

### SetArabicFamilyNameNil

`func (o *SaudiArabiaNafathCitizenAttributes) SetArabicFamilyNameNil(b bool)`

 SetArabicFamilyNameNil sets the value for ArabicFamilyName to be an explicit nil

### UnsetArabicFamilyName
`func (o *SaudiArabiaNafathCitizenAttributes) UnsetArabicFamilyName()`

UnsetArabicFamilyName ensures that no value is present for ArabicFamilyName, not even an explicit nil
### GetRomanizedFirstName

`func (o *SaudiArabiaNafathCitizenAttributes) GetRomanizedFirstName() string`

GetRomanizedFirstName returns the RomanizedFirstName field if non-nil, zero value otherwise.

### GetRomanizedFirstNameOk

`func (o *SaudiArabiaNafathCitizenAttributes) GetRomanizedFirstNameOk() (*string, bool)`

GetRomanizedFirstNameOk returns a tuple with the RomanizedFirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRomanizedFirstName

`func (o *SaudiArabiaNafathCitizenAttributes) SetRomanizedFirstName(v string)`

SetRomanizedFirstName sets RomanizedFirstName field to given value.

### HasRomanizedFirstName

`func (o *SaudiArabiaNafathCitizenAttributes) HasRomanizedFirstName() bool`

HasRomanizedFirstName returns a boolean if a field has been set.

### SetRomanizedFirstNameNil

`func (o *SaudiArabiaNafathCitizenAttributes) SetRomanizedFirstNameNil(b bool)`

 SetRomanizedFirstNameNil sets the value for RomanizedFirstName to be an explicit nil

### UnsetRomanizedFirstName
`func (o *SaudiArabiaNafathCitizenAttributes) UnsetRomanizedFirstName()`

UnsetRomanizedFirstName ensures that no value is present for RomanizedFirstName, not even an explicit nil
### GetRomanizedSecondName

`func (o *SaudiArabiaNafathCitizenAttributes) GetRomanizedSecondName() string`

GetRomanizedSecondName returns the RomanizedSecondName field if non-nil, zero value otherwise.

### GetRomanizedSecondNameOk

`func (o *SaudiArabiaNafathCitizenAttributes) GetRomanizedSecondNameOk() (*string, bool)`

GetRomanizedSecondNameOk returns a tuple with the RomanizedSecondName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRomanizedSecondName

`func (o *SaudiArabiaNafathCitizenAttributes) SetRomanizedSecondName(v string)`

SetRomanizedSecondName sets RomanizedSecondName field to given value.

### HasRomanizedSecondName

`func (o *SaudiArabiaNafathCitizenAttributes) HasRomanizedSecondName() bool`

HasRomanizedSecondName returns a boolean if a field has been set.

### SetRomanizedSecondNameNil

`func (o *SaudiArabiaNafathCitizenAttributes) SetRomanizedSecondNameNil(b bool)`

 SetRomanizedSecondNameNil sets the value for RomanizedSecondName to be an explicit nil

### UnsetRomanizedSecondName
`func (o *SaudiArabiaNafathCitizenAttributes) UnsetRomanizedSecondName()`

UnsetRomanizedSecondName ensures that no value is present for RomanizedSecondName, not even an explicit nil
### GetRomanizedThirdName

`func (o *SaudiArabiaNafathCitizenAttributes) GetRomanizedThirdName() string`

GetRomanizedThirdName returns the RomanizedThirdName field if non-nil, zero value otherwise.

### GetRomanizedThirdNameOk

`func (o *SaudiArabiaNafathCitizenAttributes) GetRomanizedThirdNameOk() (*string, bool)`

GetRomanizedThirdNameOk returns a tuple with the RomanizedThirdName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRomanizedThirdName

`func (o *SaudiArabiaNafathCitizenAttributes) SetRomanizedThirdName(v string)`

SetRomanizedThirdName sets RomanizedThirdName field to given value.

### HasRomanizedThirdName

`func (o *SaudiArabiaNafathCitizenAttributes) HasRomanizedThirdName() bool`

HasRomanizedThirdName returns a boolean if a field has been set.

### SetRomanizedThirdNameNil

`func (o *SaudiArabiaNafathCitizenAttributes) SetRomanizedThirdNameNil(b bool)`

 SetRomanizedThirdNameNil sets the value for RomanizedThirdName to be an explicit nil

### UnsetRomanizedThirdName
`func (o *SaudiArabiaNafathCitizenAttributes) UnsetRomanizedThirdName()`

UnsetRomanizedThirdName ensures that no value is present for RomanizedThirdName, not even an explicit nil
### GetRomanizedLastName

`func (o *SaudiArabiaNafathCitizenAttributes) GetRomanizedLastName() string`

GetRomanizedLastName returns the RomanizedLastName field if non-nil, zero value otherwise.

### GetRomanizedLastNameOk

`func (o *SaudiArabiaNafathCitizenAttributes) GetRomanizedLastNameOk() (*string, bool)`

GetRomanizedLastNameOk returns a tuple with the RomanizedLastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRomanizedLastName

`func (o *SaudiArabiaNafathCitizenAttributes) SetRomanizedLastName(v string)`

SetRomanizedLastName sets RomanizedLastName field to given value.

### HasRomanizedLastName

`func (o *SaudiArabiaNafathCitizenAttributes) HasRomanizedLastName() bool`

HasRomanizedLastName returns a boolean if a field has been set.

### SetRomanizedLastNameNil

`func (o *SaudiArabiaNafathCitizenAttributes) SetRomanizedLastNameNil(b bool)`

 SetRomanizedLastNameNil sets the value for RomanizedLastName to be an explicit nil

### UnsetRomanizedLastName
`func (o *SaudiArabiaNafathCitizenAttributes) UnsetRomanizedLastName()`

UnsetRomanizedLastName ensures that no value is present for RomanizedLastName, not even an explicit nil
### GetDateOfBirthHijri

`func (o *SaudiArabiaNafathCitizenAttributes) GetDateOfBirthHijri() string`

GetDateOfBirthHijri returns the DateOfBirthHijri field if non-nil, zero value otherwise.

### GetDateOfBirthHijriOk

`func (o *SaudiArabiaNafathCitizenAttributes) GetDateOfBirthHijriOk() (*string, bool)`

GetDateOfBirthHijriOk returns a tuple with the DateOfBirthHijri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirthHijri

`func (o *SaudiArabiaNafathCitizenAttributes) SetDateOfBirthHijri(v string)`

SetDateOfBirthHijri sets DateOfBirthHijri field to given value.

### HasDateOfBirthHijri

`func (o *SaudiArabiaNafathCitizenAttributes) HasDateOfBirthHijri() bool`

HasDateOfBirthHijri returns a boolean if a field has been set.

### SetDateOfBirthHijriNil

`func (o *SaudiArabiaNafathCitizenAttributes) SetDateOfBirthHijriNil(b bool)`

 SetDateOfBirthHijriNil sets the value for DateOfBirthHijri to be an explicit nil

### UnsetDateOfBirthHijri
`func (o *SaudiArabiaNafathCitizenAttributes) UnsetDateOfBirthHijri()`

UnsetDateOfBirthHijri ensures that no value is present for DateOfBirthHijri, not even an explicit nil
### GetIdIssueDateGregorian

`func (o *SaudiArabiaNafathCitizenAttributes) GetIdIssueDateGregorian() string`

GetIdIssueDateGregorian returns the IdIssueDateGregorian field if non-nil, zero value otherwise.

### GetIdIssueDateGregorianOk

`func (o *SaudiArabiaNafathCitizenAttributes) GetIdIssueDateGregorianOk() (*string, bool)`

GetIdIssueDateGregorianOk returns a tuple with the IdIssueDateGregorian field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdIssueDateGregorian

`func (o *SaudiArabiaNafathCitizenAttributes) SetIdIssueDateGregorian(v string)`

SetIdIssueDateGregorian sets IdIssueDateGregorian field to given value.

### HasIdIssueDateGregorian

`func (o *SaudiArabiaNafathCitizenAttributes) HasIdIssueDateGregorian() bool`

HasIdIssueDateGregorian returns a boolean if a field has been set.

### SetIdIssueDateGregorianNil

`func (o *SaudiArabiaNafathCitizenAttributes) SetIdIssueDateGregorianNil(b bool)`

 SetIdIssueDateGregorianNil sets the value for IdIssueDateGregorian to be an explicit nil

### UnsetIdIssueDateGregorian
`func (o *SaudiArabiaNafathCitizenAttributes) UnsetIdIssueDateGregorian()`

UnsetIdIssueDateGregorian ensures that no value is present for IdIssueDateGregorian, not even an explicit nil
### GetIdIssueDateHijri

`func (o *SaudiArabiaNafathCitizenAttributes) GetIdIssueDateHijri() string`

GetIdIssueDateHijri returns the IdIssueDateHijri field if non-nil, zero value otherwise.

### GetIdIssueDateHijriOk

`func (o *SaudiArabiaNafathCitizenAttributes) GetIdIssueDateHijriOk() (*string, bool)`

GetIdIssueDateHijriOk returns a tuple with the IdIssueDateHijri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdIssueDateHijri

`func (o *SaudiArabiaNafathCitizenAttributes) SetIdIssueDateHijri(v string)`

SetIdIssueDateHijri sets IdIssueDateHijri field to given value.

### HasIdIssueDateHijri

`func (o *SaudiArabiaNafathCitizenAttributes) HasIdIssueDateHijri() bool`

HasIdIssueDateHijri returns a boolean if a field has been set.

### SetIdIssueDateHijriNil

`func (o *SaudiArabiaNafathCitizenAttributes) SetIdIssueDateHijriNil(b bool)`

 SetIdIssueDateHijriNil sets the value for IdIssueDateHijri to be an explicit nil

### UnsetIdIssueDateHijri
`func (o *SaudiArabiaNafathCitizenAttributes) UnsetIdIssueDateHijri()`

UnsetIdIssueDateHijri ensures that no value is present for IdIssueDateHijri, not even an explicit nil
### GetIdExpiryDateGregorian

`func (o *SaudiArabiaNafathCitizenAttributes) GetIdExpiryDateGregorian() string`

GetIdExpiryDateGregorian returns the IdExpiryDateGregorian field if non-nil, zero value otherwise.

### GetIdExpiryDateGregorianOk

`func (o *SaudiArabiaNafathCitizenAttributes) GetIdExpiryDateGregorianOk() (*string, bool)`

GetIdExpiryDateGregorianOk returns a tuple with the IdExpiryDateGregorian field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdExpiryDateGregorian

`func (o *SaudiArabiaNafathCitizenAttributes) SetIdExpiryDateGregorian(v string)`

SetIdExpiryDateGregorian sets IdExpiryDateGregorian field to given value.

### HasIdExpiryDateGregorian

`func (o *SaudiArabiaNafathCitizenAttributes) HasIdExpiryDateGregorian() bool`

HasIdExpiryDateGregorian returns a boolean if a field has been set.

### SetIdExpiryDateGregorianNil

`func (o *SaudiArabiaNafathCitizenAttributes) SetIdExpiryDateGregorianNil(b bool)`

 SetIdExpiryDateGregorianNil sets the value for IdExpiryDateGregorian to be an explicit nil

### UnsetIdExpiryDateGregorian
`func (o *SaudiArabiaNafathCitizenAttributes) UnsetIdExpiryDateGregorian()`

UnsetIdExpiryDateGregorian ensures that no value is present for IdExpiryDateGregorian, not even an explicit nil
### GetIdExpiryDateHijri

`func (o *SaudiArabiaNafathCitizenAttributes) GetIdExpiryDateHijri() string`

GetIdExpiryDateHijri returns the IdExpiryDateHijri field if non-nil, zero value otherwise.

### GetIdExpiryDateHijriOk

`func (o *SaudiArabiaNafathCitizenAttributes) GetIdExpiryDateHijriOk() (*string, bool)`

GetIdExpiryDateHijriOk returns a tuple with the IdExpiryDateHijri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdExpiryDateHijri

`func (o *SaudiArabiaNafathCitizenAttributes) SetIdExpiryDateHijri(v string)`

SetIdExpiryDateHijri sets IdExpiryDateHijri field to given value.

### HasIdExpiryDateHijri

`func (o *SaudiArabiaNafathCitizenAttributes) HasIdExpiryDateHijri() bool`

HasIdExpiryDateHijri returns a boolean if a field has been set.

### SetIdExpiryDateHijriNil

`func (o *SaudiArabiaNafathCitizenAttributes) SetIdExpiryDateHijriNil(b bool)`

 SetIdExpiryDateHijriNil sets the value for IdExpiryDateHijri to be an explicit nil

### UnsetIdExpiryDateHijri
`func (o *SaudiArabiaNafathCitizenAttributes) UnsetIdExpiryDateHijri()`

UnsetIdExpiryDateHijri ensures that no value is present for IdExpiryDateHijri, not even an explicit nil
### GetIdVersionNumber

`func (o *SaudiArabiaNafathCitizenAttributes) GetIdVersionNumber() string`

GetIdVersionNumber returns the IdVersionNumber field if non-nil, zero value otherwise.

### GetIdVersionNumberOk

`func (o *SaudiArabiaNafathCitizenAttributes) GetIdVersionNumberOk() (*string, bool)`

GetIdVersionNumberOk returns a tuple with the IdVersionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdVersionNumber

`func (o *SaudiArabiaNafathCitizenAttributes) SetIdVersionNumber(v string)`

SetIdVersionNumber sets IdVersionNumber field to given value.

### HasIdVersionNumber

`func (o *SaudiArabiaNafathCitizenAttributes) HasIdVersionNumber() bool`

HasIdVersionNumber returns a boolean if a field has been set.

### SetIdVersionNumberNil

`func (o *SaudiArabiaNafathCitizenAttributes) SetIdVersionNumberNil(b bool)`

 SetIdVersionNumberNil sets the value for IdVersionNumber to be an explicit nil

### UnsetIdVersionNumber
`func (o *SaudiArabiaNafathCitizenAttributes) UnsetIdVersionNumber()`

UnsetIdVersionNumber ensures that no value is present for IdVersionNumber, not even an explicit nil
### GetIdIssuePlace

`func (o *SaudiArabiaNafathCitizenAttributes) GetIdIssuePlace() string`

GetIdIssuePlace returns the IdIssuePlace field if non-nil, zero value otherwise.

### GetIdIssuePlaceOk

`func (o *SaudiArabiaNafathCitizenAttributes) GetIdIssuePlaceOk() (*string, bool)`

GetIdIssuePlaceOk returns a tuple with the IdIssuePlace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdIssuePlace

`func (o *SaudiArabiaNafathCitizenAttributes) SetIdIssuePlace(v string)`

SetIdIssuePlace sets IdIssuePlace field to given value.

### HasIdIssuePlace

`func (o *SaudiArabiaNafathCitizenAttributes) HasIdIssuePlace() bool`

HasIdIssuePlace returns a boolean if a field has been set.

### SetIdIssuePlaceNil

`func (o *SaudiArabiaNafathCitizenAttributes) SetIdIssuePlaceNil(b bool)`

 SetIdIssuePlaceNil sets the value for IdIssuePlace to be an explicit nil

### UnsetIdIssuePlace
`func (o *SaudiArabiaNafathCitizenAttributes) UnsetIdIssuePlace()`

UnsetIdIssuePlace ensures that no value is present for IdIssuePlace, not even an explicit nil
### GetSocialStatusCode

`func (o *SaudiArabiaNafathCitizenAttributes) GetSocialStatusCode() string`

GetSocialStatusCode returns the SocialStatusCode field if non-nil, zero value otherwise.

### GetSocialStatusCodeOk

`func (o *SaudiArabiaNafathCitizenAttributes) GetSocialStatusCodeOk() (*string, bool)`

GetSocialStatusCodeOk returns a tuple with the SocialStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialStatusCode

`func (o *SaudiArabiaNafathCitizenAttributes) SetSocialStatusCode(v string)`

SetSocialStatusCode sets SocialStatusCode field to given value.

### HasSocialStatusCode

`func (o *SaudiArabiaNafathCitizenAttributes) HasSocialStatusCode() bool`

HasSocialStatusCode returns a boolean if a field has been set.

### SetSocialStatusCodeNil

`func (o *SaudiArabiaNafathCitizenAttributes) SetSocialStatusCodeNil(b bool)`

 SetSocialStatusCodeNil sets the value for SocialStatusCode to be an explicit nil

### UnsetSocialStatusCode
`func (o *SaudiArabiaNafathCitizenAttributes) UnsetSocialStatusCode()`

UnsetSocialStatusCode ensures that no value is present for SocialStatusCode, not even an explicit nil
### GetSocialStatus

`func (o *SaudiArabiaNafathCitizenAttributes) GetSocialStatus() string`

GetSocialStatus returns the SocialStatus field if non-nil, zero value otherwise.

### GetSocialStatusOk

`func (o *SaudiArabiaNafathCitizenAttributes) GetSocialStatusOk() (*string, bool)`

GetSocialStatusOk returns a tuple with the SocialStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialStatus

`func (o *SaudiArabiaNafathCitizenAttributes) SetSocialStatus(v string)`

SetSocialStatus sets SocialStatus field to given value.

### HasSocialStatus

`func (o *SaudiArabiaNafathCitizenAttributes) HasSocialStatus() bool`

HasSocialStatus returns a boolean if a field has been set.

### SetSocialStatusNil

`func (o *SaudiArabiaNafathCitizenAttributes) SetSocialStatusNil(b bool)`

 SetSocialStatusNil sets the value for SocialStatus to be an explicit nil

### UnsetSocialStatus
`func (o *SaudiArabiaNafathCitizenAttributes) UnsetSocialStatus()`

UnsetSocialStatus ensures that no value is present for SocialStatus, not even an explicit nil
### GetDependentCount

`func (o *SaudiArabiaNafathCitizenAttributes) GetDependentCount() int32`

GetDependentCount returns the DependentCount field if non-nil, zero value otherwise.

### GetDependentCountOk

`func (o *SaudiArabiaNafathCitizenAttributes) GetDependentCountOk() (*int32, bool)`

GetDependentCountOk returns a tuple with the DependentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependentCount

`func (o *SaudiArabiaNafathCitizenAttributes) SetDependentCount(v int32)`

SetDependentCount sets DependentCount field to given value.

### HasDependentCount

`func (o *SaudiArabiaNafathCitizenAttributes) HasDependentCount() bool`

HasDependentCount returns a boolean if a field has been set.

### SetDependentCountNil

`func (o *SaudiArabiaNafathCitizenAttributes) SetDependentCountNil(b bool)`

 SetDependentCountNil sets the value for DependentCount to be an explicit nil

### UnsetDependentCount
`func (o *SaudiArabiaNafathCitizenAttributes) UnsetDependentCount()`

UnsetDependentCount ensures that no value is present for DependentCount, not even an explicit nil
### GetPlaceOfBirth

`func (o *SaudiArabiaNafathCitizenAttributes) GetPlaceOfBirth() string`

GetPlaceOfBirth returns the PlaceOfBirth field if non-nil, zero value otherwise.

### GetPlaceOfBirthOk

`func (o *SaudiArabiaNafathCitizenAttributes) GetPlaceOfBirthOk() (*string, bool)`

GetPlaceOfBirthOk returns a tuple with the PlaceOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceOfBirth

`func (o *SaudiArabiaNafathCitizenAttributes) SetPlaceOfBirth(v string)`

SetPlaceOfBirth sets PlaceOfBirth field to given value.

### HasPlaceOfBirth

`func (o *SaudiArabiaNafathCitizenAttributes) HasPlaceOfBirth() bool`

HasPlaceOfBirth returns a boolean if a field has been set.

### SetPlaceOfBirthNil

`func (o *SaudiArabiaNafathCitizenAttributes) SetPlaceOfBirthNil(b bool)`

 SetPlaceOfBirthNil sets the value for PlaceOfBirth to be an explicit nil

### UnsetPlaceOfBirth
`func (o *SaudiArabiaNafathCitizenAttributes) UnsetPlaceOfBirth()`

UnsetPlaceOfBirth ensures that no value is present for PlaceOfBirth, not even an explicit nil
### GetLastTravelExitDateGregorian

`func (o *SaudiArabiaNafathCitizenAttributes) GetLastTravelExitDateGregorian() string`

GetLastTravelExitDateGregorian returns the LastTravelExitDateGregorian field if non-nil, zero value otherwise.

### GetLastTravelExitDateGregorianOk

`func (o *SaudiArabiaNafathCitizenAttributes) GetLastTravelExitDateGregorianOk() (*string, bool)`

GetLastTravelExitDateGregorianOk returns a tuple with the LastTravelExitDateGregorian field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTravelExitDateGregorian

`func (o *SaudiArabiaNafathCitizenAttributes) SetLastTravelExitDateGregorian(v string)`

SetLastTravelExitDateGregorian sets LastTravelExitDateGregorian field to given value.

### HasLastTravelExitDateGregorian

`func (o *SaudiArabiaNafathCitizenAttributes) HasLastTravelExitDateGregorian() bool`

HasLastTravelExitDateGregorian returns a boolean if a field has been set.

### SetLastTravelExitDateGregorianNil

`func (o *SaudiArabiaNafathCitizenAttributes) SetLastTravelExitDateGregorianNil(b bool)`

 SetLastTravelExitDateGregorianNil sets the value for LastTravelExitDateGregorian to be an explicit nil

### UnsetLastTravelExitDateGregorian
`func (o *SaudiArabiaNafathCitizenAttributes) UnsetLastTravelExitDateGregorian()`

UnsetLastTravelExitDateGregorian ensures that no value is present for LastTravelExitDateGregorian, not even an explicit nil
### GetLastTravelExitDateHijri

`func (o *SaudiArabiaNafathCitizenAttributes) GetLastTravelExitDateHijri() string`

GetLastTravelExitDateHijri returns the LastTravelExitDateHijri field if non-nil, zero value otherwise.

### GetLastTravelExitDateHijriOk

`func (o *SaudiArabiaNafathCitizenAttributes) GetLastTravelExitDateHijriOk() (*string, bool)`

GetLastTravelExitDateHijriOk returns a tuple with the LastTravelExitDateHijri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTravelExitDateHijri

`func (o *SaudiArabiaNafathCitizenAttributes) SetLastTravelExitDateHijri(v string)`

SetLastTravelExitDateHijri sets LastTravelExitDateHijri field to given value.

### HasLastTravelExitDateHijri

`func (o *SaudiArabiaNafathCitizenAttributes) HasLastTravelExitDateHijri() bool`

HasLastTravelExitDateHijri returns a boolean if a field has been set.

### SetLastTravelExitDateHijriNil

`func (o *SaudiArabiaNafathCitizenAttributes) SetLastTravelExitDateHijriNil(b bool)`

 SetLastTravelExitDateHijriNil sets the value for LastTravelExitDateHijri to be an explicit nil

### UnsetLastTravelExitDateHijri
`func (o *SaudiArabiaNafathCitizenAttributes) UnsetLastTravelExitDateHijri()`

UnsetLastTravelExitDateHijri ensures that no value is present for LastTravelExitDateHijri, not even an explicit nil
### GetLastTravelEntryDateGregorian

`func (o *SaudiArabiaNafathCitizenAttributes) GetLastTravelEntryDateGregorian() string`

GetLastTravelEntryDateGregorian returns the LastTravelEntryDateGregorian field if non-nil, zero value otherwise.

### GetLastTravelEntryDateGregorianOk

`func (o *SaudiArabiaNafathCitizenAttributes) GetLastTravelEntryDateGregorianOk() (*string, bool)`

GetLastTravelEntryDateGregorianOk returns a tuple with the LastTravelEntryDateGregorian field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTravelEntryDateGregorian

`func (o *SaudiArabiaNafathCitizenAttributes) SetLastTravelEntryDateGregorian(v string)`

SetLastTravelEntryDateGregorian sets LastTravelEntryDateGregorian field to given value.

### HasLastTravelEntryDateGregorian

`func (o *SaudiArabiaNafathCitizenAttributes) HasLastTravelEntryDateGregorian() bool`

HasLastTravelEntryDateGregorian returns a boolean if a field has been set.

### SetLastTravelEntryDateGregorianNil

`func (o *SaudiArabiaNafathCitizenAttributes) SetLastTravelEntryDateGregorianNil(b bool)`

 SetLastTravelEntryDateGregorianNil sets the value for LastTravelEntryDateGregorian to be an explicit nil

### UnsetLastTravelEntryDateGregorian
`func (o *SaudiArabiaNafathCitizenAttributes) UnsetLastTravelEntryDateGregorian()`

UnsetLastTravelEntryDateGregorian ensures that no value is present for LastTravelEntryDateGregorian, not even an explicit nil
### GetLastTravelEntryDateHijri

`func (o *SaudiArabiaNafathCitizenAttributes) GetLastTravelEntryDateHijri() string`

GetLastTravelEntryDateHijri returns the LastTravelEntryDateHijri field if non-nil, zero value otherwise.

### GetLastTravelEntryDateHijriOk

`func (o *SaudiArabiaNafathCitizenAttributes) GetLastTravelEntryDateHijriOk() (*string, bool)`

GetLastTravelEntryDateHijriOk returns a tuple with the LastTravelEntryDateHijri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTravelEntryDateHijri

`func (o *SaudiArabiaNafathCitizenAttributes) SetLastTravelEntryDateHijri(v string)`

SetLastTravelEntryDateHijri sets LastTravelEntryDateHijri field to given value.

### HasLastTravelEntryDateHijri

`func (o *SaudiArabiaNafathCitizenAttributes) HasLastTravelEntryDateHijri() bool`

HasLastTravelEntryDateHijri returns a boolean if a field has been set.

### SetLastTravelEntryDateHijriNil

`func (o *SaudiArabiaNafathCitizenAttributes) SetLastTravelEntryDateHijriNil(b bool)`

 SetLastTravelEntryDateHijriNil sets the value for LastTravelEntryDateHijri to be an explicit nil

### UnsetLastTravelEntryDateHijri
`func (o *SaudiArabiaNafathCitizenAttributes) UnsetLastTravelEntryDateHijri()`

UnsetLastTravelEntryDateHijri ensures that no value is present for LastTravelEntryDateHijri, not even an explicit nil
### GetIsMinor

`func (o *SaudiArabiaNafathCitizenAttributes) GetIsMinor() bool`

GetIsMinor returns the IsMinor field if non-nil, zero value otherwise.

### GetIsMinorOk

`func (o *SaudiArabiaNafathCitizenAttributes) GetIsMinorOk() (*bool, bool)`

GetIsMinorOk returns a tuple with the IsMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMinor

`func (o *SaudiArabiaNafathCitizenAttributes) SetIsMinor(v bool)`

SetIsMinor sets IsMinor field to given value.

### HasIsMinor

`func (o *SaudiArabiaNafathCitizenAttributes) HasIsMinor() bool`

HasIsMinor returns a boolean if a field has been set.

### SetIsMinorNil

`func (o *SaudiArabiaNafathCitizenAttributes) SetIsMinorNil(b bool)`

 SetIsMinorNil sets the value for IsMinor to be an explicit nil

### UnsetIsMinor
`func (o *SaudiArabiaNafathCitizenAttributes) UnsetIsMinor()`

UnsetIsMinor ensures that no value is present for IsMinor, not even an explicit nil
### GetHeadOfHousehold

`func (o *SaudiArabiaNafathCitizenAttributes) GetHeadOfHousehold() SaudiArabiaNafathHeadOfHousehold`

GetHeadOfHousehold returns the HeadOfHousehold field if non-nil, zero value otherwise.

### GetHeadOfHouseholdOk

`func (o *SaudiArabiaNafathCitizenAttributes) GetHeadOfHouseholdOk() (*SaudiArabiaNafathHeadOfHousehold, bool)`

GetHeadOfHouseholdOk returns a tuple with the HeadOfHousehold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadOfHousehold

`func (o *SaudiArabiaNafathCitizenAttributes) SetHeadOfHousehold(v SaudiArabiaNafathHeadOfHousehold)`

SetHeadOfHousehold sets HeadOfHousehold field to given value.

### HasHeadOfHousehold

`func (o *SaudiArabiaNafathCitizenAttributes) HasHeadOfHousehold() bool`

HasHeadOfHousehold returns a boolean if a field has been set.

### SetHeadOfHouseholdNil

`func (o *SaudiArabiaNafathCitizenAttributes) SetHeadOfHouseholdNil(b bool)`

 SetHeadOfHouseholdNil sets the value for HeadOfHousehold to be an explicit nil

### UnsetHeadOfHousehold
`func (o *SaudiArabiaNafathCitizenAttributes) UnsetHeadOfHousehold()`

UnsetHeadOfHousehold ensures that no value is present for HeadOfHousehold, not even an explicit nil
### GetNationalAddresses

`func (o *SaudiArabiaNafathCitizenAttributes) GetNationalAddresses() []SaudiArabiaNafathAddress`

GetNationalAddresses returns the NationalAddresses field if non-nil, zero value otherwise.

### GetNationalAddressesOk

`func (o *SaudiArabiaNafathCitizenAttributes) GetNationalAddressesOk() (*[]SaudiArabiaNafathAddress, bool)`

GetNationalAddressesOk returns a tuple with the NationalAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationalAddresses

`func (o *SaudiArabiaNafathCitizenAttributes) SetNationalAddresses(v []SaudiArabiaNafathAddress)`

SetNationalAddresses sets NationalAddresses field to given value.

### HasNationalAddresses

`func (o *SaudiArabiaNafathCitizenAttributes) HasNationalAddresses() bool`

HasNationalAddresses returns a boolean if a field has been set.

### SetNationalAddressesNil

`func (o *SaudiArabiaNafathCitizenAttributes) SetNationalAddressesNil(b bool)`

 SetNationalAddressesNil sets the value for NationalAddresses to be an explicit nil

### UnsetNationalAddresses
`func (o *SaudiArabiaNafathCitizenAttributes) UnsetNationalAddresses()`

UnsetNationalAddresses ensures that no value is present for NationalAddresses, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


