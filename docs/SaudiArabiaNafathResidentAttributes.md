# SaudiArabiaNafathResidentAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArabicFirstName** | Pointer to **NullableString** | Given name in Arabic.              Nafath always returns this in Arabic. romanizedFirstName is the English form, and the session locale does not change either name. | [optional] 
**ArabicSecondName** | Pointer to **NullableString** | Second name in Arabic. Same script as arabicFirstName. | [optional] 
**ArabicThirdName** | Pointer to **NullableString** | Third name in Arabic. Same script as arabicFirstName. | [optional] 
**ArabicLastName** | Pointer to **NullableString** | Last name in Arabic. Same script as arabicFirstName. | [optional] 
**RomanizedFirstName** | Pointer to **NullableString** | Given name in English.              Nafath always returns this in English, including when the session locale is Arabic. | [optional] 
**RomanizedSecondName** | Pointer to **NullableString** | Second name in English. Same script as romanizedFirstName. | [optional] 
**RomanizedThirdName** | Pointer to **NullableString** | Third name in English. Same script as romanizedFirstName. | [optional] 
**RomanizedLastName** | Pointer to **NullableString** | Last name in English. Same script as romanizedFirstName. | [optional] 
**NationalityDescription** | Pointer to **NullableString** | English nationality name. | [optional] 
**DateOfBirthHijri** | Pointer to **NullableString** | Date of birth in the Hijri (Umm al-Qura) calendar, in YYYY-MM-DD format. | [optional] 
**IqamaIssueDateGregorian** | Pointer to **NullableString** | Date the Iqama was issued in the Gregorian calendar. | [optional] 
**IqamaIssueDateHijri** | Pointer to **NullableString** | Date the Iqama was issued in the Hijri calendar, in YYYY-MM-DD format. | [optional] 
**IqamaExpiryDateGregorian** | Pointer to **NullableString** | Date the Iqama expires in the Gregorian calendar. | [optional] 
**IqamaExpiryDateHijri** | Pointer to **NullableString** | Date the Iqama expires in the Hijri calendar, in YYYY-MM-DD format. | [optional] 
**IqamaVersionNumber** | Pointer to **NullableString** | Version number printed on the Iqama.              Represents the Iqama permit generation version. | [optional] 
**IqamaIssuePlace** | Pointer to **NullableString** | Name of the city or office that issued the Iqama, in Arabic.              The session locale does not change it. | [optional] 
**IqamaIssuePlaceCode** | Pointer to **NullableString** | Code for the office that issued the Iqama. | [optional] 
**OccupationDescription** | Pointer to **NullableString** | Free-text occupation description. | [optional] 
**SocialStatusCode** | Pointer to **NullableString** | Social-status (marital-status) code, from the Nafath socialStatusCode attribute. The numeric code is not unique on its own. The same code maps to different states depending on the value of the top-level sex field.              Code by sex: - 0, M: Single - 1, F: Single - 1, M: One wife - 2, F: Married - 2, M: Two wives - 3, F: Divorced - 3, M: Three wives - 4, F: Widow - 4, M: Four wives - 5, F: Pending divorced              Combinations not listed are not defined by Nafath. | [optional] 
**SocialStatus** | Pointer to **NullableString** | Social status in uppercase English, for example MARRIED, from the Nafath socialStatus attribute.              Possible English values include: Single, Married, Divorced, Widow, Pending Divorced, One Wife, Two Wives, Three Wives, Four Wives. | [optional] 
**PlaceOfBirth** | Pointer to **NullableString** | Place of birth name. | [optional] 
**PlaceOfBirthCode** | Pointer to **NullableString** | Code for the place of birth. | [optional] 
**SponsorName** | Pointer to **NullableString** | Name of the entity sponsoring the Iqama, typically an employer. | [optional] 
**LegalStatus** | Pointer to **NullableString** | Legal status. | [optional] 
**InOutIndicator** | Pointer to **NullableString** | Whether the resident is currently inside or outside Saudi Arabia based on their last recorded travel event.              Possible values include: - IN: most recent travel event was an entry into Saudi Arabia. - OUT: most recent travel event was an exit from Saudi Arabia. | [optional] 
**NationalAddresses** | Pointer to [**[]SaudiArabiaNafathAddress**](SaudiArabiaNafathAddress.md) | Registered Saudi national addresses.              Saudi Post (SPL) documents how to enter a national address on checkout and international forms: Line 1 is building number, unit number, and street name. Line 2 is the additional number and district. Then city and postal code. The country is always Saudi Arabia.              https://splonline.com.sa/en/door-step/ | [optional] 

## Methods

### Named construction

Use a keyed struct literal (`SaudiArabiaNafathResidentAttributes{FieldName: value}`), or call
`NewSaudiArabiaNafathResidentAttributesWithDefaults()` and set fields by name. Positional field
factories and unkeyed struct literals are not supported.

### NewSaudiArabiaNafathResidentAttributesWithDefaults

`func NewSaudiArabiaNafathResidentAttributesWithDefaults() *SaudiArabiaNafathResidentAttributes`

NewSaudiArabiaNafathResidentAttributesWithDefaults instantiates a new SaudiArabiaNafathResidentAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArabicFirstName

`func (o *SaudiArabiaNafathResidentAttributes) GetArabicFirstName() string`

GetArabicFirstName returns the ArabicFirstName field if non-nil, zero value otherwise.

### GetArabicFirstNameOk

`func (o *SaudiArabiaNafathResidentAttributes) GetArabicFirstNameOk() (*string, bool)`

GetArabicFirstNameOk returns a tuple with the ArabicFirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArabicFirstName

`func (o *SaudiArabiaNafathResidentAttributes) SetArabicFirstName(v string)`

SetArabicFirstName sets ArabicFirstName field to given value.

### HasArabicFirstName

`func (o *SaudiArabiaNafathResidentAttributes) HasArabicFirstName() bool`

HasArabicFirstName returns a boolean if a field has been set.

### SetArabicFirstNameNil

`func (o *SaudiArabiaNafathResidentAttributes) SetArabicFirstNameNil(b bool)`

 SetArabicFirstNameNil sets the value for ArabicFirstName to be an explicit nil

### UnsetArabicFirstName
`func (o *SaudiArabiaNafathResidentAttributes) UnsetArabicFirstName()`

UnsetArabicFirstName ensures that no value is present for ArabicFirstName, not even an explicit nil
### GetArabicSecondName

`func (o *SaudiArabiaNafathResidentAttributes) GetArabicSecondName() string`

GetArabicSecondName returns the ArabicSecondName field if non-nil, zero value otherwise.

### GetArabicSecondNameOk

`func (o *SaudiArabiaNafathResidentAttributes) GetArabicSecondNameOk() (*string, bool)`

GetArabicSecondNameOk returns a tuple with the ArabicSecondName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArabicSecondName

`func (o *SaudiArabiaNafathResidentAttributes) SetArabicSecondName(v string)`

SetArabicSecondName sets ArabicSecondName field to given value.

### HasArabicSecondName

`func (o *SaudiArabiaNafathResidentAttributes) HasArabicSecondName() bool`

HasArabicSecondName returns a boolean if a field has been set.

### SetArabicSecondNameNil

`func (o *SaudiArabiaNafathResidentAttributes) SetArabicSecondNameNil(b bool)`

 SetArabicSecondNameNil sets the value for ArabicSecondName to be an explicit nil

### UnsetArabicSecondName
`func (o *SaudiArabiaNafathResidentAttributes) UnsetArabicSecondName()`

UnsetArabicSecondName ensures that no value is present for ArabicSecondName, not even an explicit nil
### GetArabicThirdName

`func (o *SaudiArabiaNafathResidentAttributes) GetArabicThirdName() string`

GetArabicThirdName returns the ArabicThirdName field if non-nil, zero value otherwise.

### GetArabicThirdNameOk

`func (o *SaudiArabiaNafathResidentAttributes) GetArabicThirdNameOk() (*string, bool)`

GetArabicThirdNameOk returns a tuple with the ArabicThirdName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArabicThirdName

`func (o *SaudiArabiaNafathResidentAttributes) SetArabicThirdName(v string)`

SetArabicThirdName sets ArabicThirdName field to given value.

### HasArabicThirdName

`func (o *SaudiArabiaNafathResidentAttributes) HasArabicThirdName() bool`

HasArabicThirdName returns a boolean if a field has been set.

### SetArabicThirdNameNil

`func (o *SaudiArabiaNafathResidentAttributes) SetArabicThirdNameNil(b bool)`

 SetArabicThirdNameNil sets the value for ArabicThirdName to be an explicit nil

### UnsetArabicThirdName
`func (o *SaudiArabiaNafathResidentAttributes) UnsetArabicThirdName()`

UnsetArabicThirdName ensures that no value is present for ArabicThirdName, not even an explicit nil
### GetArabicLastName

`func (o *SaudiArabiaNafathResidentAttributes) GetArabicLastName() string`

GetArabicLastName returns the ArabicLastName field if non-nil, zero value otherwise.

### GetArabicLastNameOk

`func (o *SaudiArabiaNafathResidentAttributes) GetArabicLastNameOk() (*string, bool)`

GetArabicLastNameOk returns a tuple with the ArabicLastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArabicLastName

`func (o *SaudiArabiaNafathResidentAttributes) SetArabicLastName(v string)`

SetArabicLastName sets ArabicLastName field to given value.

### HasArabicLastName

`func (o *SaudiArabiaNafathResidentAttributes) HasArabicLastName() bool`

HasArabicLastName returns a boolean if a field has been set.

### SetArabicLastNameNil

`func (o *SaudiArabiaNafathResidentAttributes) SetArabicLastNameNil(b bool)`

 SetArabicLastNameNil sets the value for ArabicLastName to be an explicit nil

### UnsetArabicLastName
`func (o *SaudiArabiaNafathResidentAttributes) UnsetArabicLastName()`

UnsetArabicLastName ensures that no value is present for ArabicLastName, not even an explicit nil
### GetRomanizedFirstName

`func (o *SaudiArabiaNafathResidentAttributes) GetRomanizedFirstName() string`

GetRomanizedFirstName returns the RomanizedFirstName field if non-nil, zero value otherwise.

### GetRomanizedFirstNameOk

`func (o *SaudiArabiaNafathResidentAttributes) GetRomanizedFirstNameOk() (*string, bool)`

GetRomanizedFirstNameOk returns a tuple with the RomanizedFirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRomanizedFirstName

`func (o *SaudiArabiaNafathResidentAttributes) SetRomanizedFirstName(v string)`

SetRomanizedFirstName sets RomanizedFirstName field to given value.

### HasRomanizedFirstName

`func (o *SaudiArabiaNafathResidentAttributes) HasRomanizedFirstName() bool`

HasRomanizedFirstName returns a boolean if a field has been set.

### SetRomanizedFirstNameNil

`func (o *SaudiArabiaNafathResidentAttributes) SetRomanizedFirstNameNil(b bool)`

 SetRomanizedFirstNameNil sets the value for RomanizedFirstName to be an explicit nil

### UnsetRomanizedFirstName
`func (o *SaudiArabiaNafathResidentAttributes) UnsetRomanizedFirstName()`

UnsetRomanizedFirstName ensures that no value is present for RomanizedFirstName, not even an explicit nil
### GetRomanizedSecondName

`func (o *SaudiArabiaNafathResidentAttributes) GetRomanizedSecondName() string`

GetRomanizedSecondName returns the RomanizedSecondName field if non-nil, zero value otherwise.

### GetRomanizedSecondNameOk

`func (o *SaudiArabiaNafathResidentAttributes) GetRomanizedSecondNameOk() (*string, bool)`

GetRomanizedSecondNameOk returns a tuple with the RomanizedSecondName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRomanizedSecondName

`func (o *SaudiArabiaNafathResidentAttributes) SetRomanizedSecondName(v string)`

SetRomanizedSecondName sets RomanizedSecondName field to given value.

### HasRomanizedSecondName

`func (o *SaudiArabiaNafathResidentAttributes) HasRomanizedSecondName() bool`

HasRomanizedSecondName returns a boolean if a field has been set.

### SetRomanizedSecondNameNil

`func (o *SaudiArabiaNafathResidentAttributes) SetRomanizedSecondNameNil(b bool)`

 SetRomanizedSecondNameNil sets the value for RomanizedSecondName to be an explicit nil

### UnsetRomanizedSecondName
`func (o *SaudiArabiaNafathResidentAttributes) UnsetRomanizedSecondName()`

UnsetRomanizedSecondName ensures that no value is present for RomanizedSecondName, not even an explicit nil
### GetRomanizedThirdName

`func (o *SaudiArabiaNafathResidentAttributes) GetRomanizedThirdName() string`

GetRomanizedThirdName returns the RomanizedThirdName field if non-nil, zero value otherwise.

### GetRomanizedThirdNameOk

`func (o *SaudiArabiaNafathResidentAttributes) GetRomanizedThirdNameOk() (*string, bool)`

GetRomanizedThirdNameOk returns a tuple with the RomanizedThirdName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRomanizedThirdName

`func (o *SaudiArabiaNafathResidentAttributes) SetRomanizedThirdName(v string)`

SetRomanizedThirdName sets RomanizedThirdName field to given value.

### HasRomanizedThirdName

`func (o *SaudiArabiaNafathResidentAttributes) HasRomanizedThirdName() bool`

HasRomanizedThirdName returns a boolean if a field has been set.

### SetRomanizedThirdNameNil

`func (o *SaudiArabiaNafathResidentAttributes) SetRomanizedThirdNameNil(b bool)`

 SetRomanizedThirdNameNil sets the value for RomanizedThirdName to be an explicit nil

### UnsetRomanizedThirdName
`func (o *SaudiArabiaNafathResidentAttributes) UnsetRomanizedThirdName()`

UnsetRomanizedThirdName ensures that no value is present for RomanizedThirdName, not even an explicit nil
### GetRomanizedLastName

`func (o *SaudiArabiaNafathResidentAttributes) GetRomanizedLastName() string`

GetRomanizedLastName returns the RomanizedLastName field if non-nil, zero value otherwise.

### GetRomanizedLastNameOk

`func (o *SaudiArabiaNafathResidentAttributes) GetRomanizedLastNameOk() (*string, bool)`

GetRomanizedLastNameOk returns a tuple with the RomanizedLastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRomanizedLastName

`func (o *SaudiArabiaNafathResidentAttributes) SetRomanizedLastName(v string)`

SetRomanizedLastName sets RomanizedLastName field to given value.

### HasRomanizedLastName

`func (o *SaudiArabiaNafathResidentAttributes) HasRomanizedLastName() bool`

HasRomanizedLastName returns a boolean if a field has been set.

### SetRomanizedLastNameNil

`func (o *SaudiArabiaNafathResidentAttributes) SetRomanizedLastNameNil(b bool)`

 SetRomanizedLastNameNil sets the value for RomanizedLastName to be an explicit nil

### UnsetRomanizedLastName
`func (o *SaudiArabiaNafathResidentAttributes) UnsetRomanizedLastName()`

UnsetRomanizedLastName ensures that no value is present for RomanizedLastName, not even an explicit nil
### GetNationalityDescription

`func (o *SaudiArabiaNafathResidentAttributes) GetNationalityDescription() string`

GetNationalityDescription returns the NationalityDescription field if non-nil, zero value otherwise.

### GetNationalityDescriptionOk

`func (o *SaudiArabiaNafathResidentAttributes) GetNationalityDescriptionOk() (*string, bool)`

GetNationalityDescriptionOk returns a tuple with the NationalityDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationalityDescription

`func (o *SaudiArabiaNafathResidentAttributes) SetNationalityDescription(v string)`

SetNationalityDescription sets NationalityDescription field to given value.

### HasNationalityDescription

`func (o *SaudiArabiaNafathResidentAttributes) HasNationalityDescription() bool`

HasNationalityDescription returns a boolean if a field has been set.

### SetNationalityDescriptionNil

`func (o *SaudiArabiaNafathResidentAttributes) SetNationalityDescriptionNil(b bool)`

 SetNationalityDescriptionNil sets the value for NationalityDescription to be an explicit nil

### UnsetNationalityDescription
`func (o *SaudiArabiaNafathResidentAttributes) UnsetNationalityDescription()`

UnsetNationalityDescription ensures that no value is present for NationalityDescription, not even an explicit nil
### GetDateOfBirthHijri

`func (o *SaudiArabiaNafathResidentAttributes) GetDateOfBirthHijri() string`

GetDateOfBirthHijri returns the DateOfBirthHijri field if non-nil, zero value otherwise.

### GetDateOfBirthHijriOk

`func (o *SaudiArabiaNafathResidentAttributes) GetDateOfBirthHijriOk() (*string, bool)`

GetDateOfBirthHijriOk returns a tuple with the DateOfBirthHijri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirthHijri

`func (o *SaudiArabiaNafathResidentAttributes) SetDateOfBirthHijri(v string)`

SetDateOfBirthHijri sets DateOfBirthHijri field to given value.

### HasDateOfBirthHijri

`func (o *SaudiArabiaNafathResidentAttributes) HasDateOfBirthHijri() bool`

HasDateOfBirthHijri returns a boolean if a field has been set.

### SetDateOfBirthHijriNil

`func (o *SaudiArabiaNafathResidentAttributes) SetDateOfBirthHijriNil(b bool)`

 SetDateOfBirthHijriNil sets the value for DateOfBirthHijri to be an explicit nil

### UnsetDateOfBirthHijri
`func (o *SaudiArabiaNafathResidentAttributes) UnsetDateOfBirthHijri()`

UnsetDateOfBirthHijri ensures that no value is present for DateOfBirthHijri, not even an explicit nil
### GetIqamaIssueDateGregorian

`func (o *SaudiArabiaNafathResidentAttributes) GetIqamaIssueDateGregorian() string`

GetIqamaIssueDateGregorian returns the IqamaIssueDateGregorian field if non-nil, zero value otherwise.

### GetIqamaIssueDateGregorianOk

`func (o *SaudiArabiaNafathResidentAttributes) GetIqamaIssueDateGregorianOk() (*string, bool)`

GetIqamaIssueDateGregorianOk returns a tuple with the IqamaIssueDateGregorian field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIqamaIssueDateGregorian

`func (o *SaudiArabiaNafathResidentAttributes) SetIqamaIssueDateGregorian(v string)`

SetIqamaIssueDateGregorian sets IqamaIssueDateGregorian field to given value.

### HasIqamaIssueDateGregorian

`func (o *SaudiArabiaNafathResidentAttributes) HasIqamaIssueDateGregorian() bool`

HasIqamaIssueDateGregorian returns a boolean if a field has been set.

### SetIqamaIssueDateGregorianNil

`func (o *SaudiArabiaNafathResidentAttributes) SetIqamaIssueDateGregorianNil(b bool)`

 SetIqamaIssueDateGregorianNil sets the value for IqamaIssueDateGregorian to be an explicit nil

### UnsetIqamaIssueDateGregorian
`func (o *SaudiArabiaNafathResidentAttributes) UnsetIqamaIssueDateGregorian()`

UnsetIqamaIssueDateGregorian ensures that no value is present for IqamaIssueDateGregorian, not even an explicit nil
### GetIqamaIssueDateHijri

`func (o *SaudiArabiaNafathResidentAttributes) GetIqamaIssueDateHijri() string`

GetIqamaIssueDateHijri returns the IqamaIssueDateHijri field if non-nil, zero value otherwise.

### GetIqamaIssueDateHijriOk

`func (o *SaudiArabiaNafathResidentAttributes) GetIqamaIssueDateHijriOk() (*string, bool)`

GetIqamaIssueDateHijriOk returns a tuple with the IqamaIssueDateHijri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIqamaIssueDateHijri

`func (o *SaudiArabiaNafathResidentAttributes) SetIqamaIssueDateHijri(v string)`

SetIqamaIssueDateHijri sets IqamaIssueDateHijri field to given value.

### HasIqamaIssueDateHijri

`func (o *SaudiArabiaNafathResidentAttributes) HasIqamaIssueDateHijri() bool`

HasIqamaIssueDateHijri returns a boolean if a field has been set.

### SetIqamaIssueDateHijriNil

`func (o *SaudiArabiaNafathResidentAttributes) SetIqamaIssueDateHijriNil(b bool)`

 SetIqamaIssueDateHijriNil sets the value for IqamaIssueDateHijri to be an explicit nil

### UnsetIqamaIssueDateHijri
`func (o *SaudiArabiaNafathResidentAttributes) UnsetIqamaIssueDateHijri()`

UnsetIqamaIssueDateHijri ensures that no value is present for IqamaIssueDateHijri, not even an explicit nil
### GetIqamaExpiryDateGregorian

`func (o *SaudiArabiaNafathResidentAttributes) GetIqamaExpiryDateGregorian() string`

GetIqamaExpiryDateGregorian returns the IqamaExpiryDateGregorian field if non-nil, zero value otherwise.

### GetIqamaExpiryDateGregorianOk

`func (o *SaudiArabiaNafathResidentAttributes) GetIqamaExpiryDateGregorianOk() (*string, bool)`

GetIqamaExpiryDateGregorianOk returns a tuple with the IqamaExpiryDateGregorian field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIqamaExpiryDateGregorian

`func (o *SaudiArabiaNafathResidentAttributes) SetIqamaExpiryDateGregorian(v string)`

SetIqamaExpiryDateGregorian sets IqamaExpiryDateGregorian field to given value.

### HasIqamaExpiryDateGregorian

`func (o *SaudiArabiaNafathResidentAttributes) HasIqamaExpiryDateGregorian() bool`

HasIqamaExpiryDateGregorian returns a boolean if a field has been set.

### SetIqamaExpiryDateGregorianNil

`func (o *SaudiArabiaNafathResidentAttributes) SetIqamaExpiryDateGregorianNil(b bool)`

 SetIqamaExpiryDateGregorianNil sets the value for IqamaExpiryDateGregorian to be an explicit nil

### UnsetIqamaExpiryDateGregorian
`func (o *SaudiArabiaNafathResidentAttributes) UnsetIqamaExpiryDateGregorian()`

UnsetIqamaExpiryDateGregorian ensures that no value is present for IqamaExpiryDateGregorian, not even an explicit nil
### GetIqamaExpiryDateHijri

`func (o *SaudiArabiaNafathResidentAttributes) GetIqamaExpiryDateHijri() string`

GetIqamaExpiryDateHijri returns the IqamaExpiryDateHijri field if non-nil, zero value otherwise.

### GetIqamaExpiryDateHijriOk

`func (o *SaudiArabiaNafathResidentAttributes) GetIqamaExpiryDateHijriOk() (*string, bool)`

GetIqamaExpiryDateHijriOk returns a tuple with the IqamaExpiryDateHijri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIqamaExpiryDateHijri

`func (o *SaudiArabiaNafathResidentAttributes) SetIqamaExpiryDateHijri(v string)`

SetIqamaExpiryDateHijri sets IqamaExpiryDateHijri field to given value.

### HasIqamaExpiryDateHijri

`func (o *SaudiArabiaNafathResidentAttributes) HasIqamaExpiryDateHijri() bool`

HasIqamaExpiryDateHijri returns a boolean if a field has been set.

### SetIqamaExpiryDateHijriNil

`func (o *SaudiArabiaNafathResidentAttributes) SetIqamaExpiryDateHijriNil(b bool)`

 SetIqamaExpiryDateHijriNil sets the value for IqamaExpiryDateHijri to be an explicit nil

### UnsetIqamaExpiryDateHijri
`func (o *SaudiArabiaNafathResidentAttributes) UnsetIqamaExpiryDateHijri()`

UnsetIqamaExpiryDateHijri ensures that no value is present for IqamaExpiryDateHijri, not even an explicit nil
### GetIqamaVersionNumber

`func (o *SaudiArabiaNafathResidentAttributes) GetIqamaVersionNumber() string`

GetIqamaVersionNumber returns the IqamaVersionNumber field if non-nil, zero value otherwise.

### GetIqamaVersionNumberOk

`func (o *SaudiArabiaNafathResidentAttributes) GetIqamaVersionNumberOk() (*string, bool)`

GetIqamaVersionNumberOk returns a tuple with the IqamaVersionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIqamaVersionNumber

`func (o *SaudiArabiaNafathResidentAttributes) SetIqamaVersionNumber(v string)`

SetIqamaVersionNumber sets IqamaVersionNumber field to given value.

### HasIqamaVersionNumber

`func (o *SaudiArabiaNafathResidentAttributes) HasIqamaVersionNumber() bool`

HasIqamaVersionNumber returns a boolean if a field has been set.

### SetIqamaVersionNumberNil

`func (o *SaudiArabiaNafathResidentAttributes) SetIqamaVersionNumberNil(b bool)`

 SetIqamaVersionNumberNil sets the value for IqamaVersionNumber to be an explicit nil

### UnsetIqamaVersionNumber
`func (o *SaudiArabiaNafathResidentAttributes) UnsetIqamaVersionNumber()`

UnsetIqamaVersionNumber ensures that no value is present for IqamaVersionNumber, not even an explicit nil
### GetIqamaIssuePlace

`func (o *SaudiArabiaNafathResidentAttributes) GetIqamaIssuePlace() string`

GetIqamaIssuePlace returns the IqamaIssuePlace field if non-nil, zero value otherwise.

### GetIqamaIssuePlaceOk

`func (o *SaudiArabiaNafathResidentAttributes) GetIqamaIssuePlaceOk() (*string, bool)`

GetIqamaIssuePlaceOk returns a tuple with the IqamaIssuePlace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIqamaIssuePlace

`func (o *SaudiArabiaNafathResidentAttributes) SetIqamaIssuePlace(v string)`

SetIqamaIssuePlace sets IqamaIssuePlace field to given value.

### HasIqamaIssuePlace

`func (o *SaudiArabiaNafathResidentAttributes) HasIqamaIssuePlace() bool`

HasIqamaIssuePlace returns a boolean if a field has been set.

### SetIqamaIssuePlaceNil

`func (o *SaudiArabiaNafathResidentAttributes) SetIqamaIssuePlaceNil(b bool)`

 SetIqamaIssuePlaceNil sets the value for IqamaIssuePlace to be an explicit nil

### UnsetIqamaIssuePlace
`func (o *SaudiArabiaNafathResidentAttributes) UnsetIqamaIssuePlace()`

UnsetIqamaIssuePlace ensures that no value is present for IqamaIssuePlace, not even an explicit nil
### GetIqamaIssuePlaceCode

`func (o *SaudiArabiaNafathResidentAttributes) GetIqamaIssuePlaceCode() string`

GetIqamaIssuePlaceCode returns the IqamaIssuePlaceCode field if non-nil, zero value otherwise.

### GetIqamaIssuePlaceCodeOk

`func (o *SaudiArabiaNafathResidentAttributes) GetIqamaIssuePlaceCodeOk() (*string, bool)`

GetIqamaIssuePlaceCodeOk returns a tuple with the IqamaIssuePlaceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIqamaIssuePlaceCode

`func (o *SaudiArabiaNafathResidentAttributes) SetIqamaIssuePlaceCode(v string)`

SetIqamaIssuePlaceCode sets IqamaIssuePlaceCode field to given value.

### HasIqamaIssuePlaceCode

`func (o *SaudiArabiaNafathResidentAttributes) HasIqamaIssuePlaceCode() bool`

HasIqamaIssuePlaceCode returns a boolean if a field has been set.

### SetIqamaIssuePlaceCodeNil

`func (o *SaudiArabiaNafathResidentAttributes) SetIqamaIssuePlaceCodeNil(b bool)`

 SetIqamaIssuePlaceCodeNil sets the value for IqamaIssuePlaceCode to be an explicit nil

### UnsetIqamaIssuePlaceCode
`func (o *SaudiArabiaNafathResidentAttributes) UnsetIqamaIssuePlaceCode()`

UnsetIqamaIssuePlaceCode ensures that no value is present for IqamaIssuePlaceCode, not even an explicit nil
### GetOccupationDescription

`func (o *SaudiArabiaNafathResidentAttributes) GetOccupationDescription() string`

GetOccupationDescription returns the OccupationDescription field if non-nil, zero value otherwise.

### GetOccupationDescriptionOk

`func (o *SaudiArabiaNafathResidentAttributes) GetOccupationDescriptionOk() (*string, bool)`

GetOccupationDescriptionOk returns a tuple with the OccupationDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccupationDescription

`func (o *SaudiArabiaNafathResidentAttributes) SetOccupationDescription(v string)`

SetOccupationDescription sets OccupationDescription field to given value.

### HasOccupationDescription

`func (o *SaudiArabiaNafathResidentAttributes) HasOccupationDescription() bool`

HasOccupationDescription returns a boolean if a field has been set.

### SetOccupationDescriptionNil

`func (o *SaudiArabiaNafathResidentAttributes) SetOccupationDescriptionNil(b bool)`

 SetOccupationDescriptionNil sets the value for OccupationDescription to be an explicit nil

### UnsetOccupationDescription
`func (o *SaudiArabiaNafathResidentAttributes) UnsetOccupationDescription()`

UnsetOccupationDescription ensures that no value is present for OccupationDescription, not even an explicit nil
### GetSocialStatusCode

`func (o *SaudiArabiaNafathResidentAttributes) GetSocialStatusCode() string`

GetSocialStatusCode returns the SocialStatusCode field if non-nil, zero value otherwise.

### GetSocialStatusCodeOk

`func (o *SaudiArabiaNafathResidentAttributes) GetSocialStatusCodeOk() (*string, bool)`

GetSocialStatusCodeOk returns a tuple with the SocialStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialStatusCode

`func (o *SaudiArabiaNafathResidentAttributes) SetSocialStatusCode(v string)`

SetSocialStatusCode sets SocialStatusCode field to given value.

### HasSocialStatusCode

`func (o *SaudiArabiaNafathResidentAttributes) HasSocialStatusCode() bool`

HasSocialStatusCode returns a boolean if a field has been set.

### SetSocialStatusCodeNil

`func (o *SaudiArabiaNafathResidentAttributes) SetSocialStatusCodeNil(b bool)`

 SetSocialStatusCodeNil sets the value for SocialStatusCode to be an explicit nil

### UnsetSocialStatusCode
`func (o *SaudiArabiaNafathResidentAttributes) UnsetSocialStatusCode()`

UnsetSocialStatusCode ensures that no value is present for SocialStatusCode, not even an explicit nil
### GetSocialStatus

`func (o *SaudiArabiaNafathResidentAttributes) GetSocialStatus() string`

GetSocialStatus returns the SocialStatus field if non-nil, zero value otherwise.

### GetSocialStatusOk

`func (o *SaudiArabiaNafathResidentAttributes) GetSocialStatusOk() (*string, bool)`

GetSocialStatusOk returns a tuple with the SocialStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialStatus

`func (o *SaudiArabiaNafathResidentAttributes) SetSocialStatus(v string)`

SetSocialStatus sets SocialStatus field to given value.

### HasSocialStatus

`func (o *SaudiArabiaNafathResidentAttributes) HasSocialStatus() bool`

HasSocialStatus returns a boolean if a field has been set.

### SetSocialStatusNil

`func (o *SaudiArabiaNafathResidentAttributes) SetSocialStatusNil(b bool)`

 SetSocialStatusNil sets the value for SocialStatus to be an explicit nil

### UnsetSocialStatus
`func (o *SaudiArabiaNafathResidentAttributes) UnsetSocialStatus()`

UnsetSocialStatus ensures that no value is present for SocialStatus, not even an explicit nil
### GetPlaceOfBirth

`func (o *SaudiArabiaNafathResidentAttributes) GetPlaceOfBirth() string`

GetPlaceOfBirth returns the PlaceOfBirth field if non-nil, zero value otherwise.

### GetPlaceOfBirthOk

`func (o *SaudiArabiaNafathResidentAttributes) GetPlaceOfBirthOk() (*string, bool)`

GetPlaceOfBirthOk returns a tuple with the PlaceOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceOfBirth

`func (o *SaudiArabiaNafathResidentAttributes) SetPlaceOfBirth(v string)`

SetPlaceOfBirth sets PlaceOfBirth field to given value.

### HasPlaceOfBirth

`func (o *SaudiArabiaNafathResidentAttributes) HasPlaceOfBirth() bool`

HasPlaceOfBirth returns a boolean if a field has been set.

### SetPlaceOfBirthNil

`func (o *SaudiArabiaNafathResidentAttributes) SetPlaceOfBirthNil(b bool)`

 SetPlaceOfBirthNil sets the value for PlaceOfBirth to be an explicit nil

### UnsetPlaceOfBirth
`func (o *SaudiArabiaNafathResidentAttributes) UnsetPlaceOfBirth()`

UnsetPlaceOfBirth ensures that no value is present for PlaceOfBirth, not even an explicit nil
### GetPlaceOfBirthCode

`func (o *SaudiArabiaNafathResidentAttributes) GetPlaceOfBirthCode() string`

GetPlaceOfBirthCode returns the PlaceOfBirthCode field if non-nil, zero value otherwise.

### GetPlaceOfBirthCodeOk

`func (o *SaudiArabiaNafathResidentAttributes) GetPlaceOfBirthCodeOk() (*string, bool)`

GetPlaceOfBirthCodeOk returns a tuple with the PlaceOfBirthCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceOfBirthCode

`func (o *SaudiArabiaNafathResidentAttributes) SetPlaceOfBirthCode(v string)`

SetPlaceOfBirthCode sets PlaceOfBirthCode field to given value.

### HasPlaceOfBirthCode

`func (o *SaudiArabiaNafathResidentAttributes) HasPlaceOfBirthCode() bool`

HasPlaceOfBirthCode returns a boolean if a field has been set.

### SetPlaceOfBirthCodeNil

`func (o *SaudiArabiaNafathResidentAttributes) SetPlaceOfBirthCodeNil(b bool)`

 SetPlaceOfBirthCodeNil sets the value for PlaceOfBirthCode to be an explicit nil

### UnsetPlaceOfBirthCode
`func (o *SaudiArabiaNafathResidentAttributes) UnsetPlaceOfBirthCode()`

UnsetPlaceOfBirthCode ensures that no value is present for PlaceOfBirthCode, not even an explicit nil
### GetSponsorName

`func (o *SaudiArabiaNafathResidentAttributes) GetSponsorName() string`

GetSponsorName returns the SponsorName field if non-nil, zero value otherwise.

### GetSponsorNameOk

`func (o *SaudiArabiaNafathResidentAttributes) GetSponsorNameOk() (*string, bool)`

GetSponsorNameOk returns a tuple with the SponsorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorName

`func (o *SaudiArabiaNafathResidentAttributes) SetSponsorName(v string)`

SetSponsorName sets SponsorName field to given value.

### HasSponsorName

`func (o *SaudiArabiaNafathResidentAttributes) HasSponsorName() bool`

HasSponsorName returns a boolean if a field has been set.

### SetSponsorNameNil

`func (o *SaudiArabiaNafathResidentAttributes) SetSponsorNameNil(b bool)`

 SetSponsorNameNil sets the value for SponsorName to be an explicit nil

### UnsetSponsorName
`func (o *SaudiArabiaNafathResidentAttributes) UnsetSponsorName()`

UnsetSponsorName ensures that no value is present for SponsorName, not even an explicit nil
### GetLegalStatus

`func (o *SaudiArabiaNafathResidentAttributes) GetLegalStatus() string`

GetLegalStatus returns the LegalStatus field if non-nil, zero value otherwise.

### GetLegalStatusOk

`func (o *SaudiArabiaNafathResidentAttributes) GetLegalStatusOk() (*string, bool)`

GetLegalStatusOk returns a tuple with the LegalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalStatus

`func (o *SaudiArabiaNafathResidentAttributes) SetLegalStatus(v string)`

SetLegalStatus sets LegalStatus field to given value.

### HasLegalStatus

`func (o *SaudiArabiaNafathResidentAttributes) HasLegalStatus() bool`

HasLegalStatus returns a boolean if a field has been set.

### SetLegalStatusNil

`func (o *SaudiArabiaNafathResidentAttributes) SetLegalStatusNil(b bool)`

 SetLegalStatusNil sets the value for LegalStatus to be an explicit nil

### UnsetLegalStatus
`func (o *SaudiArabiaNafathResidentAttributes) UnsetLegalStatus()`

UnsetLegalStatus ensures that no value is present for LegalStatus, not even an explicit nil
### GetInOutIndicator

`func (o *SaudiArabiaNafathResidentAttributes) GetInOutIndicator() string`

GetInOutIndicator returns the InOutIndicator field if non-nil, zero value otherwise.

### GetInOutIndicatorOk

`func (o *SaudiArabiaNafathResidentAttributes) GetInOutIndicatorOk() (*string, bool)`

GetInOutIndicatorOk returns a tuple with the InOutIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInOutIndicator

`func (o *SaudiArabiaNafathResidentAttributes) SetInOutIndicator(v string)`

SetInOutIndicator sets InOutIndicator field to given value.

### HasInOutIndicator

`func (o *SaudiArabiaNafathResidentAttributes) HasInOutIndicator() bool`

HasInOutIndicator returns a boolean if a field has been set.

### SetInOutIndicatorNil

`func (o *SaudiArabiaNafathResidentAttributes) SetInOutIndicatorNil(b bool)`

 SetInOutIndicatorNil sets the value for InOutIndicator to be an explicit nil

### UnsetInOutIndicator
`func (o *SaudiArabiaNafathResidentAttributes) UnsetInOutIndicator()`

UnsetInOutIndicator ensures that no value is present for InOutIndicator, not even an explicit nil
### GetNationalAddresses

`func (o *SaudiArabiaNafathResidentAttributes) GetNationalAddresses() []SaudiArabiaNafathAddress`

GetNationalAddresses returns the NationalAddresses field if non-nil, zero value otherwise.

### GetNationalAddressesOk

`func (o *SaudiArabiaNafathResidentAttributes) GetNationalAddressesOk() (*[]SaudiArabiaNafathAddress, bool)`

GetNationalAddressesOk returns a tuple with the NationalAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationalAddresses

`func (o *SaudiArabiaNafathResidentAttributes) SetNationalAddresses(v []SaudiArabiaNafathAddress)`

SetNationalAddresses sets NationalAddresses field to given value.

### HasNationalAddresses

`func (o *SaudiArabiaNafathResidentAttributes) HasNationalAddresses() bool`

HasNationalAddresses returns a boolean if a field has been set.

### SetNationalAddressesNil

`func (o *SaudiArabiaNafathResidentAttributes) SetNationalAddressesNil(b bool)`

 SetNationalAddressesNil sets the value for NationalAddresses to be an explicit nil

### UnsetNationalAddresses
`func (o *SaudiArabiaNafathResidentAttributes) UnsetNationalAddresses()`

UnsetNationalAddresses ensures that no value is present for NationalAddresses, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


