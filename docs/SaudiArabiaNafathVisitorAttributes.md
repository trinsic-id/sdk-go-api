# SaudiArabiaNafathVisitorAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoginType** | Pointer to **NullableString** | Credential the visitor used to log in to Nafath.              Possible values: - VisaNumber: the visitor authenticated with the visa number printed on their visa   document. - BorderNumber: the visitor authenticated with their border number (raqam al-hudud), the   ten-digit identifier the General Directorate of Passports (Jawazat) assigns to a   foreign national on entry into Saudi Arabia, used as their entry-linked identifier   before any residency is issued. | [optional] 
**FirstName** | Pointer to **NullableString** | First name in English.              Nafath describes firstName, secondName, thirdName, and lastName as English. translatedFirstName is the Arabic form. The session locale does not change either name. | [optional] 
**SecondName** | Pointer to **NullableString** | Second name in English. Same script as firstName. | [optional] 
**ThirdName** | Pointer to **NullableString** | Third name, from the Nafath thirdName attribute. See firstName for the script. | [optional] 
**LastName** | Pointer to **NullableString** | Last name, from the Nafath lastName attribute. See firstName for the script. | [optional] 
**TranslatedFirstName** | Pointer to **NullableString** | First name in Arabic, from the Nafath translatedFirstName attribute.              Nafath&#39;s descriptions of firstName, secondName, thirdName, and lastName call those fields English, so the translated names are treated as Arabic for both visa-number and border-number logins. The session locale does not change either script. | [optional] 
**TranslatedSecondName** | Pointer to **NullableString** | Second name in Arabic. Same script as translatedFirstName. | [optional] 
**TranslatedThirdName** | Pointer to **NullableString** | Third name in Arabic. Same script as translatedFirstName. | [optional] 
**TranslatedLastName** | Pointer to **NullableString** | Last name in Arabic. Same script as translatedFirstName. | [optional] 
**NationalityDescription** | Pointer to **NullableString** | English nationality name, from the Nafath nationalityDesc attribute. | [optional] 
**DateOfBirthHijri** | Pointer to **NullableString** | Date of birth in the Hijri (Umm al-Qura) calendar, in YYYY-MM-DD format, from the Nafath birthhDate attribute. Only returned for visa-number logins; the appendix defines no Hijri date of birth for border-number logins.              Nafath&#39;s examples also use DD-MM-YYYY and YYYY/MM/DD; Trinsic reorders those to YYYY-MM-DD and passes any other format through unchanged. | [optional] 
**VisaNumber** | Pointer to **NullableString** | Visa number, from the Nafath visaNumber attribute. Can contain letters as well as digits. | [optional] 
**VisaDurationDays** | Pointer to **NullableInt32** | Count of days from visa issue to visa expiry, parsed from the Nafath duration attribute. Null when duration is absent or not a valid integer. | [optional] 
**VisaIssueDate** | Pointer to **NullableString** | Date the visa was issued, from the Nafath visaIssueDate attribute. | [optional] 
**VisaExpiryDate** | Pointer to **NullableString** | Final date the visitor is permitted to enter Saudi Arabia, from the Nafath visaExpiryDate attribute. | [optional] 
**FinalDepartureDate** | Pointer to **NullableString** | Last day the visitor may legally remain in Saudi Arabia, from the Nafath finalDepartureDate attribute. | [optional] 
**NumberOfEntriesCode** | Pointer to **NullableString** | Number-of-entries code, from the Nafath numberOfEntriesCode attribute.              Known values: - 1: Single entry - 9: Multiple entry | [optional] 
**NumberOfEntriesDescriptionArabic** | Pointer to **NullableString** | Arabic description of the entries the visa allows, from the Nafath numberOfEntriesDescAr attribute. | [optional] 
**VisaTypeCode** | Pointer to **NullableString** | Visa type code, from the Nafath visaTypeCode attribute. The appendix publishes no lookup table; the sandbox returns BUS for a business visit visa. | [optional] 
**VisaTypeDescriptionArabic** | Pointer to **NullableString** | Arabic visa type description, from the Nafath visaTypeDescAr attribute. | [optional] 
**VisaTypeDescriptionEnglish** | Pointer to **NullableString** | English visa type description, from the Nafath visaTypeDescEn attribute. | [optional] 
**PassportIssueDate** | Pointer to **NullableString** | Date the visitor&#39;s passport was issued, from the Nafath passportIssueDate attribute. | [optional] 
**PassportExpiryDate** | Pointer to **NullableString** | Date the visitor&#39;s passport expires, from the Nafath passportExpiryDate attribute. | [optional] 
**OccupationDescriptionArabic** | Pointer to **NullableString** | Arabic occupation description, from the Nafath occupationDescAr attribute. | [optional] 
**OccupationDescriptionEnglish** | Pointer to **NullableString** | English occupation description, from the Nafath occupationDescEn attribute. | [optional] 
**BirthCountryArabic** | Pointer to **NullableString** | Arabic name of the visitor&#39;s birth country, from the Nafath birthCountryDescAr attribute. | [optional] 
**BirthCountryEnglish** | Pointer to **NullableString** | English name of the visitor&#39;s birth country, from the Nafath birthCountryDescEn attribute. | [optional] 
**EntryDate** | Pointer to **NullableString** | Date the visitor entered Saudi Arabia, from the Nafath entryDate attribute. Only returned for border-number logins. | [optional] 

## Methods

### Named construction

Use a keyed struct literal (`SaudiArabiaNafathVisitorAttributes{FieldName: value}`), or call
`NewSaudiArabiaNafathVisitorAttributesWithDefaults()` and set fields by name. Positional field
factories and unkeyed struct literals are not supported.

### NewSaudiArabiaNafathVisitorAttributesWithDefaults

`func NewSaudiArabiaNafathVisitorAttributesWithDefaults() *SaudiArabiaNafathVisitorAttributes`

NewSaudiArabiaNafathVisitorAttributesWithDefaults instantiates a new SaudiArabiaNafathVisitorAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoginType

`func (o *SaudiArabiaNafathVisitorAttributes) GetLoginType() string`

GetLoginType returns the LoginType field if non-nil, zero value otherwise.

### GetLoginTypeOk

`func (o *SaudiArabiaNafathVisitorAttributes) GetLoginTypeOk() (*string, bool)`

GetLoginTypeOk returns a tuple with the LoginType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginType

`func (o *SaudiArabiaNafathVisitorAttributes) SetLoginType(v string)`

SetLoginType sets LoginType field to given value.

### HasLoginType

`func (o *SaudiArabiaNafathVisitorAttributes) HasLoginType() bool`

HasLoginType returns a boolean if a field has been set.

### SetLoginTypeNil

`func (o *SaudiArabiaNafathVisitorAttributes) SetLoginTypeNil(b bool)`

 SetLoginTypeNil sets the value for LoginType to be an explicit nil

### UnsetLoginType
`func (o *SaudiArabiaNafathVisitorAttributes) UnsetLoginType()`

UnsetLoginType ensures that no value is present for LoginType, not even an explicit nil
### GetFirstName

`func (o *SaudiArabiaNafathVisitorAttributes) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *SaudiArabiaNafathVisitorAttributes) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *SaudiArabiaNafathVisitorAttributes) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *SaudiArabiaNafathVisitorAttributes) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *SaudiArabiaNafathVisitorAttributes) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *SaudiArabiaNafathVisitorAttributes) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetSecondName

`func (o *SaudiArabiaNafathVisitorAttributes) GetSecondName() string`

GetSecondName returns the SecondName field if non-nil, zero value otherwise.

### GetSecondNameOk

`func (o *SaudiArabiaNafathVisitorAttributes) GetSecondNameOk() (*string, bool)`

GetSecondNameOk returns a tuple with the SecondName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondName

`func (o *SaudiArabiaNafathVisitorAttributes) SetSecondName(v string)`

SetSecondName sets SecondName field to given value.

### HasSecondName

`func (o *SaudiArabiaNafathVisitorAttributes) HasSecondName() bool`

HasSecondName returns a boolean if a field has been set.

### SetSecondNameNil

`func (o *SaudiArabiaNafathVisitorAttributes) SetSecondNameNil(b bool)`

 SetSecondNameNil sets the value for SecondName to be an explicit nil

### UnsetSecondName
`func (o *SaudiArabiaNafathVisitorAttributes) UnsetSecondName()`

UnsetSecondName ensures that no value is present for SecondName, not even an explicit nil
### GetThirdName

`func (o *SaudiArabiaNafathVisitorAttributes) GetThirdName() string`

GetThirdName returns the ThirdName field if non-nil, zero value otherwise.

### GetThirdNameOk

`func (o *SaudiArabiaNafathVisitorAttributes) GetThirdNameOk() (*string, bool)`

GetThirdNameOk returns a tuple with the ThirdName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThirdName

`func (o *SaudiArabiaNafathVisitorAttributes) SetThirdName(v string)`

SetThirdName sets ThirdName field to given value.

### HasThirdName

`func (o *SaudiArabiaNafathVisitorAttributes) HasThirdName() bool`

HasThirdName returns a boolean if a field has been set.

### SetThirdNameNil

`func (o *SaudiArabiaNafathVisitorAttributes) SetThirdNameNil(b bool)`

 SetThirdNameNil sets the value for ThirdName to be an explicit nil

### UnsetThirdName
`func (o *SaudiArabiaNafathVisitorAttributes) UnsetThirdName()`

UnsetThirdName ensures that no value is present for ThirdName, not even an explicit nil
### GetLastName

`func (o *SaudiArabiaNafathVisitorAttributes) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *SaudiArabiaNafathVisitorAttributes) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *SaudiArabiaNafathVisitorAttributes) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *SaudiArabiaNafathVisitorAttributes) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *SaudiArabiaNafathVisitorAttributes) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *SaudiArabiaNafathVisitorAttributes) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetTranslatedFirstName

`func (o *SaudiArabiaNafathVisitorAttributes) GetTranslatedFirstName() string`

GetTranslatedFirstName returns the TranslatedFirstName field if non-nil, zero value otherwise.

### GetTranslatedFirstNameOk

`func (o *SaudiArabiaNafathVisitorAttributes) GetTranslatedFirstNameOk() (*string, bool)`

GetTranslatedFirstNameOk returns a tuple with the TranslatedFirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslatedFirstName

`func (o *SaudiArabiaNafathVisitorAttributes) SetTranslatedFirstName(v string)`

SetTranslatedFirstName sets TranslatedFirstName field to given value.

### HasTranslatedFirstName

`func (o *SaudiArabiaNafathVisitorAttributes) HasTranslatedFirstName() bool`

HasTranslatedFirstName returns a boolean if a field has been set.

### SetTranslatedFirstNameNil

`func (o *SaudiArabiaNafathVisitorAttributes) SetTranslatedFirstNameNil(b bool)`

 SetTranslatedFirstNameNil sets the value for TranslatedFirstName to be an explicit nil

### UnsetTranslatedFirstName
`func (o *SaudiArabiaNafathVisitorAttributes) UnsetTranslatedFirstName()`

UnsetTranslatedFirstName ensures that no value is present for TranslatedFirstName, not even an explicit nil
### GetTranslatedSecondName

`func (o *SaudiArabiaNafathVisitorAttributes) GetTranslatedSecondName() string`

GetTranslatedSecondName returns the TranslatedSecondName field if non-nil, zero value otherwise.

### GetTranslatedSecondNameOk

`func (o *SaudiArabiaNafathVisitorAttributes) GetTranslatedSecondNameOk() (*string, bool)`

GetTranslatedSecondNameOk returns a tuple with the TranslatedSecondName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslatedSecondName

`func (o *SaudiArabiaNafathVisitorAttributes) SetTranslatedSecondName(v string)`

SetTranslatedSecondName sets TranslatedSecondName field to given value.

### HasTranslatedSecondName

`func (o *SaudiArabiaNafathVisitorAttributes) HasTranslatedSecondName() bool`

HasTranslatedSecondName returns a boolean if a field has been set.

### SetTranslatedSecondNameNil

`func (o *SaudiArabiaNafathVisitorAttributes) SetTranslatedSecondNameNil(b bool)`

 SetTranslatedSecondNameNil sets the value for TranslatedSecondName to be an explicit nil

### UnsetTranslatedSecondName
`func (o *SaudiArabiaNafathVisitorAttributes) UnsetTranslatedSecondName()`

UnsetTranslatedSecondName ensures that no value is present for TranslatedSecondName, not even an explicit nil
### GetTranslatedThirdName

`func (o *SaudiArabiaNafathVisitorAttributes) GetTranslatedThirdName() string`

GetTranslatedThirdName returns the TranslatedThirdName field if non-nil, zero value otherwise.

### GetTranslatedThirdNameOk

`func (o *SaudiArabiaNafathVisitorAttributes) GetTranslatedThirdNameOk() (*string, bool)`

GetTranslatedThirdNameOk returns a tuple with the TranslatedThirdName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslatedThirdName

`func (o *SaudiArabiaNafathVisitorAttributes) SetTranslatedThirdName(v string)`

SetTranslatedThirdName sets TranslatedThirdName field to given value.

### HasTranslatedThirdName

`func (o *SaudiArabiaNafathVisitorAttributes) HasTranslatedThirdName() bool`

HasTranslatedThirdName returns a boolean if a field has been set.

### SetTranslatedThirdNameNil

`func (o *SaudiArabiaNafathVisitorAttributes) SetTranslatedThirdNameNil(b bool)`

 SetTranslatedThirdNameNil sets the value for TranslatedThirdName to be an explicit nil

### UnsetTranslatedThirdName
`func (o *SaudiArabiaNafathVisitorAttributes) UnsetTranslatedThirdName()`

UnsetTranslatedThirdName ensures that no value is present for TranslatedThirdName, not even an explicit nil
### GetTranslatedLastName

`func (o *SaudiArabiaNafathVisitorAttributes) GetTranslatedLastName() string`

GetTranslatedLastName returns the TranslatedLastName field if non-nil, zero value otherwise.

### GetTranslatedLastNameOk

`func (o *SaudiArabiaNafathVisitorAttributes) GetTranslatedLastNameOk() (*string, bool)`

GetTranslatedLastNameOk returns a tuple with the TranslatedLastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslatedLastName

`func (o *SaudiArabiaNafathVisitorAttributes) SetTranslatedLastName(v string)`

SetTranslatedLastName sets TranslatedLastName field to given value.

### HasTranslatedLastName

`func (o *SaudiArabiaNafathVisitorAttributes) HasTranslatedLastName() bool`

HasTranslatedLastName returns a boolean if a field has been set.

### SetTranslatedLastNameNil

`func (o *SaudiArabiaNafathVisitorAttributes) SetTranslatedLastNameNil(b bool)`

 SetTranslatedLastNameNil sets the value for TranslatedLastName to be an explicit nil

### UnsetTranslatedLastName
`func (o *SaudiArabiaNafathVisitorAttributes) UnsetTranslatedLastName()`

UnsetTranslatedLastName ensures that no value is present for TranslatedLastName, not even an explicit nil
### GetNationalityDescription

`func (o *SaudiArabiaNafathVisitorAttributes) GetNationalityDescription() string`

GetNationalityDescription returns the NationalityDescription field if non-nil, zero value otherwise.

### GetNationalityDescriptionOk

`func (o *SaudiArabiaNafathVisitorAttributes) GetNationalityDescriptionOk() (*string, bool)`

GetNationalityDescriptionOk returns a tuple with the NationalityDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationalityDescription

`func (o *SaudiArabiaNafathVisitorAttributes) SetNationalityDescription(v string)`

SetNationalityDescription sets NationalityDescription field to given value.

### HasNationalityDescription

`func (o *SaudiArabiaNafathVisitorAttributes) HasNationalityDescription() bool`

HasNationalityDescription returns a boolean if a field has been set.

### SetNationalityDescriptionNil

`func (o *SaudiArabiaNafathVisitorAttributes) SetNationalityDescriptionNil(b bool)`

 SetNationalityDescriptionNil sets the value for NationalityDescription to be an explicit nil

### UnsetNationalityDescription
`func (o *SaudiArabiaNafathVisitorAttributes) UnsetNationalityDescription()`

UnsetNationalityDescription ensures that no value is present for NationalityDescription, not even an explicit nil
### GetDateOfBirthHijri

`func (o *SaudiArabiaNafathVisitorAttributes) GetDateOfBirthHijri() string`

GetDateOfBirthHijri returns the DateOfBirthHijri field if non-nil, zero value otherwise.

### GetDateOfBirthHijriOk

`func (o *SaudiArabiaNafathVisitorAttributes) GetDateOfBirthHijriOk() (*string, bool)`

GetDateOfBirthHijriOk returns a tuple with the DateOfBirthHijri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirthHijri

`func (o *SaudiArabiaNafathVisitorAttributes) SetDateOfBirthHijri(v string)`

SetDateOfBirthHijri sets DateOfBirthHijri field to given value.

### HasDateOfBirthHijri

`func (o *SaudiArabiaNafathVisitorAttributes) HasDateOfBirthHijri() bool`

HasDateOfBirthHijri returns a boolean if a field has been set.

### SetDateOfBirthHijriNil

`func (o *SaudiArabiaNafathVisitorAttributes) SetDateOfBirthHijriNil(b bool)`

 SetDateOfBirthHijriNil sets the value for DateOfBirthHijri to be an explicit nil

### UnsetDateOfBirthHijri
`func (o *SaudiArabiaNafathVisitorAttributes) UnsetDateOfBirthHijri()`

UnsetDateOfBirthHijri ensures that no value is present for DateOfBirthHijri, not even an explicit nil
### GetVisaNumber

`func (o *SaudiArabiaNafathVisitorAttributes) GetVisaNumber() string`

GetVisaNumber returns the VisaNumber field if non-nil, zero value otherwise.

### GetVisaNumberOk

`func (o *SaudiArabiaNafathVisitorAttributes) GetVisaNumberOk() (*string, bool)`

GetVisaNumberOk returns a tuple with the VisaNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisaNumber

`func (o *SaudiArabiaNafathVisitorAttributes) SetVisaNumber(v string)`

SetVisaNumber sets VisaNumber field to given value.

### HasVisaNumber

`func (o *SaudiArabiaNafathVisitorAttributes) HasVisaNumber() bool`

HasVisaNumber returns a boolean if a field has been set.

### SetVisaNumberNil

`func (o *SaudiArabiaNafathVisitorAttributes) SetVisaNumberNil(b bool)`

 SetVisaNumberNil sets the value for VisaNumber to be an explicit nil

### UnsetVisaNumber
`func (o *SaudiArabiaNafathVisitorAttributes) UnsetVisaNumber()`

UnsetVisaNumber ensures that no value is present for VisaNumber, not even an explicit nil
### GetVisaDurationDays

`func (o *SaudiArabiaNafathVisitorAttributes) GetVisaDurationDays() int32`

GetVisaDurationDays returns the VisaDurationDays field if non-nil, zero value otherwise.

### GetVisaDurationDaysOk

`func (o *SaudiArabiaNafathVisitorAttributes) GetVisaDurationDaysOk() (*int32, bool)`

GetVisaDurationDaysOk returns a tuple with the VisaDurationDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisaDurationDays

`func (o *SaudiArabiaNafathVisitorAttributes) SetVisaDurationDays(v int32)`

SetVisaDurationDays sets VisaDurationDays field to given value.

### HasVisaDurationDays

`func (o *SaudiArabiaNafathVisitorAttributes) HasVisaDurationDays() bool`

HasVisaDurationDays returns a boolean if a field has been set.

### SetVisaDurationDaysNil

`func (o *SaudiArabiaNafathVisitorAttributes) SetVisaDurationDaysNil(b bool)`

 SetVisaDurationDaysNil sets the value for VisaDurationDays to be an explicit nil

### UnsetVisaDurationDays
`func (o *SaudiArabiaNafathVisitorAttributes) UnsetVisaDurationDays()`

UnsetVisaDurationDays ensures that no value is present for VisaDurationDays, not even an explicit nil
### GetVisaIssueDate

`func (o *SaudiArabiaNafathVisitorAttributes) GetVisaIssueDate() string`

GetVisaIssueDate returns the VisaIssueDate field if non-nil, zero value otherwise.

### GetVisaIssueDateOk

`func (o *SaudiArabiaNafathVisitorAttributes) GetVisaIssueDateOk() (*string, bool)`

GetVisaIssueDateOk returns a tuple with the VisaIssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisaIssueDate

`func (o *SaudiArabiaNafathVisitorAttributes) SetVisaIssueDate(v string)`

SetVisaIssueDate sets VisaIssueDate field to given value.

### HasVisaIssueDate

`func (o *SaudiArabiaNafathVisitorAttributes) HasVisaIssueDate() bool`

HasVisaIssueDate returns a boolean if a field has been set.

### SetVisaIssueDateNil

`func (o *SaudiArabiaNafathVisitorAttributes) SetVisaIssueDateNil(b bool)`

 SetVisaIssueDateNil sets the value for VisaIssueDate to be an explicit nil

### UnsetVisaIssueDate
`func (o *SaudiArabiaNafathVisitorAttributes) UnsetVisaIssueDate()`

UnsetVisaIssueDate ensures that no value is present for VisaIssueDate, not even an explicit nil
### GetVisaExpiryDate

`func (o *SaudiArabiaNafathVisitorAttributes) GetVisaExpiryDate() string`

GetVisaExpiryDate returns the VisaExpiryDate field if non-nil, zero value otherwise.

### GetVisaExpiryDateOk

`func (o *SaudiArabiaNafathVisitorAttributes) GetVisaExpiryDateOk() (*string, bool)`

GetVisaExpiryDateOk returns a tuple with the VisaExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisaExpiryDate

`func (o *SaudiArabiaNafathVisitorAttributes) SetVisaExpiryDate(v string)`

SetVisaExpiryDate sets VisaExpiryDate field to given value.

### HasVisaExpiryDate

`func (o *SaudiArabiaNafathVisitorAttributes) HasVisaExpiryDate() bool`

HasVisaExpiryDate returns a boolean if a field has been set.

### SetVisaExpiryDateNil

`func (o *SaudiArabiaNafathVisitorAttributes) SetVisaExpiryDateNil(b bool)`

 SetVisaExpiryDateNil sets the value for VisaExpiryDate to be an explicit nil

### UnsetVisaExpiryDate
`func (o *SaudiArabiaNafathVisitorAttributes) UnsetVisaExpiryDate()`

UnsetVisaExpiryDate ensures that no value is present for VisaExpiryDate, not even an explicit nil
### GetFinalDepartureDate

`func (o *SaudiArabiaNafathVisitorAttributes) GetFinalDepartureDate() string`

GetFinalDepartureDate returns the FinalDepartureDate field if non-nil, zero value otherwise.

### GetFinalDepartureDateOk

`func (o *SaudiArabiaNafathVisitorAttributes) GetFinalDepartureDateOk() (*string, bool)`

GetFinalDepartureDateOk returns a tuple with the FinalDepartureDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalDepartureDate

`func (o *SaudiArabiaNafathVisitorAttributes) SetFinalDepartureDate(v string)`

SetFinalDepartureDate sets FinalDepartureDate field to given value.

### HasFinalDepartureDate

`func (o *SaudiArabiaNafathVisitorAttributes) HasFinalDepartureDate() bool`

HasFinalDepartureDate returns a boolean if a field has been set.

### SetFinalDepartureDateNil

`func (o *SaudiArabiaNafathVisitorAttributes) SetFinalDepartureDateNil(b bool)`

 SetFinalDepartureDateNil sets the value for FinalDepartureDate to be an explicit nil

### UnsetFinalDepartureDate
`func (o *SaudiArabiaNafathVisitorAttributes) UnsetFinalDepartureDate()`

UnsetFinalDepartureDate ensures that no value is present for FinalDepartureDate, not even an explicit nil
### GetNumberOfEntriesCode

`func (o *SaudiArabiaNafathVisitorAttributes) GetNumberOfEntriesCode() string`

GetNumberOfEntriesCode returns the NumberOfEntriesCode field if non-nil, zero value otherwise.

### GetNumberOfEntriesCodeOk

`func (o *SaudiArabiaNafathVisitorAttributes) GetNumberOfEntriesCodeOk() (*string, bool)`

GetNumberOfEntriesCodeOk returns a tuple with the NumberOfEntriesCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfEntriesCode

`func (o *SaudiArabiaNafathVisitorAttributes) SetNumberOfEntriesCode(v string)`

SetNumberOfEntriesCode sets NumberOfEntriesCode field to given value.

### HasNumberOfEntriesCode

`func (o *SaudiArabiaNafathVisitorAttributes) HasNumberOfEntriesCode() bool`

HasNumberOfEntriesCode returns a boolean if a field has been set.

### SetNumberOfEntriesCodeNil

`func (o *SaudiArabiaNafathVisitorAttributes) SetNumberOfEntriesCodeNil(b bool)`

 SetNumberOfEntriesCodeNil sets the value for NumberOfEntriesCode to be an explicit nil

### UnsetNumberOfEntriesCode
`func (o *SaudiArabiaNafathVisitorAttributes) UnsetNumberOfEntriesCode()`

UnsetNumberOfEntriesCode ensures that no value is present for NumberOfEntriesCode, not even an explicit nil
### GetNumberOfEntriesDescriptionArabic

`func (o *SaudiArabiaNafathVisitorAttributes) GetNumberOfEntriesDescriptionArabic() string`

GetNumberOfEntriesDescriptionArabic returns the NumberOfEntriesDescriptionArabic field if non-nil, zero value otherwise.

### GetNumberOfEntriesDescriptionArabicOk

`func (o *SaudiArabiaNafathVisitorAttributes) GetNumberOfEntriesDescriptionArabicOk() (*string, bool)`

GetNumberOfEntriesDescriptionArabicOk returns a tuple with the NumberOfEntriesDescriptionArabic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfEntriesDescriptionArabic

`func (o *SaudiArabiaNafathVisitorAttributes) SetNumberOfEntriesDescriptionArabic(v string)`

SetNumberOfEntriesDescriptionArabic sets NumberOfEntriesDescriptionArabic field to given value.

### HasNumberOfEntriesDescriptionArabic

`func (o *SaudiArabiaNafathVisitorAttributes) HasNumberOfEntriesDescriptionArabic() bool`

HasNumberOfEntriesDescriptionArabic returns a boolean if a field has been set.

### SetNumberOfEntriesDescriptionArabicNil

`func (o *SaudiArabiaNafathVisitorAttributes) SetNumberOfEntriesDescriptionArabicNil(b bool)`

 SetNumberOfEntriesDescriptionArabicNil sets the value for NumberOfEntriesDescriptionArabic to be an explicit nil

### UnsetNumberOfEntriesDescriptionArabic
`func (o *SaudiArabiaNafathVisitorAttributes) UnsetNumberOfEntriesDescriptionArabic()`

UnsetNumberOfEntriesDescriptionArabic ensures that no value is present for NumberOfEntriesDescriptionArabic, not even an explicit nil
### GetVisaTypeCode

`func (o *SaudiArabiaNafathVisitorAttributes) GetVisaTypeCode() string`

GetVisaTypeCode returns the VisaTypeCode field if non-nil, zero value otherwise.

### GetVisaTypeCodeOk

`func (o *SaudiArabiaNafathVisitorAttributes) GetVisaTypeCodeOk() (*string, bool)`

GetVisaTypeCodeOk returns a tuple with the VisaTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisaTypeCode

`func (o *SaudiArabiaNafathVisitorAttributes) SetVisaTypeCode(v string)`

SetVisaTypeCode sets VisaTypeCode field to given value.

### HasVisaTypeCode

`func (o *SaudiArabiaNafathVisitorAttributes) HasVisaTypeCode() bool`

HasVisaTypeCode returns a boolean if a field has been set.

### SetVisaTypeCodeNil

`func (o *SaudiArabiaNafathVisitorAttributes) SetVisaTypeCodeNil(b bool)`

 SetVisaTypeCodeNil sets the value for VisaTypeCode to be an explicit nil

### UnsetVisaTypeCode
`func (o *SaudiArabiaNafathVisitorAttributes) UnsetVisaTypeCode()`

UnsetVisaTypeCode ensures that no value is present for VisaTypeCode, not even an explicit nil
### GetVisaTypeDescriptionArabic

`func (o *SaudiArabiaNafathVisitorAttributes) GetVisaTypeDescriptionArabic() string`

GetVisaTypeDescriptionArabic returns the VisaTypeDescriptionArabic field if non-nil, zero value otherwise.

### GetVisaTypeDescriptionArabicOk

`func (o *SaudiArabiaNafathVisitorAttributes) GetVisaTypeDescriptionArabicOk() (*string, bool)`

GetVisaTypeDescriptionArabicOk returns a tuple with the VisaTypeDescriptionArabic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisaTypeDescriptionArabic

`func (o *SaudiArabiaNafathVisitorAttributes) SetVisaTypeDescriptionArabic(v string)`

SetVisaTypeDescriptionArabic sets VisaTypeDescriptionArabic field to given value.

### HasVisaTypeDescriptionArabic

`func (o *SaudiArabiaNafathVisitorAttributes) HasVisaTypeDescriptionArabic() bool`

HasVisaTypeDescriptionArabic returns a boolean if a field has been set.

### SetVisaTypeDescriptionArabicNil

`func (o *SaudiArabiaNafathVisitorAttributes) SetVisaTypeDescriptionArabicNil(b bool)`

 SetVisaTypeDescriptionArabicNil sets the value for VisaTypeDescriptionArabic to be an explicit nil

### UnsetVisaTypeDescriptionArabic
`func (o *SaudiArabiaNafathVisitorAttributes) UnsetVisaTypeDescriptionArabic()`

UnsetVisaTypeDescriptionArabic ensures that no value is present for VisaTypeDescriptionArabic, not even an explicit nil
### GetVisaTypeDescriptionEnglish

`func (o *SaudiArabiaNafathVisitorAttributes) GetVisaTypeDescriptionEnglish() string`

GetVisaTypeDescriptionEnglish returns the VisaTypeDescriptionEnglish field if non-nil, zero value otherwise.

### GetVisaTypeDescriptionEnglishOk

`func (o *SaudiArabiaNafathVisitorAttributes) GetVisaTypeDescriptionEnglishOk() (*string, bool)`

GetVisaTypeDescriptionEnglishOk returns a tuple with the VisaTypeDescriptionEnglish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisaTypeDescriptionEnglish

`func (o *SaudiArabiaNafathVisitorAttributes) SetVisaTypeDescriptionEnglish(v string)`

SetVisaTypeDescriptionEnglish sets VisaTypeDescriptionEnglish field to given value.

### HasVisaTypeDescriptionEnglish

`func (o *SaudiArabiaNafathVisitorAttributes) HasVisaTypeDescriptionEnglish() bool`

HasVisaTypeDescriptionEnglish returns a boolean if a field has been set.

### SetVisaTypeDescriptionEnglishNil

`func (o *SaudiArabiaNafathVisitorAttributes) SetVisaTypeDescriptionEnglishNil(b bool)`

 SetVisaTypeDescriptionEnglishNil sets the value for VisaTypeDescriptionEnglish to be an explicit nil

### UnsetVisaTypeDescriptionEnglish
`func (o *SaudiArabiaNafathVisitorAttributes) UnsetVisaTypeDescriptionEnglish()`

UnsetVisaTypeDescriptionEnglish ensures that no value is present for VisaTypeDescriptionEnglish, not even an explicit nil
### GetPassportIssueDate

`func (o *SaudiArabiaNafathVisitorAttributes) GetPassportIssueDate() string`

GetPassportIssueDate returns the PassportIssueDate field if non-nil, zero value otherwise.

### GetPassportIssueDateOk

`func (o *SaudiArabiaNafathVisitorAttributes) GetPassportIssueDateOk() (*string, bool)`

GetPassportIssueDateOk returns a tuple with the PassportIssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassportIssueDate

`func (o *SaudiArabiaNafathVisitorAttributes) SetPassportIssueDate(v string)`

SetPassportIssueDate sets PassportIssueDate field to given value.

### HasPassportIssueDate

`func (o *SaudiArabiaNafathVisitorAttributes) HasPassportIssueDate() bool`

HasPassportIssueDate returns a boolean if a field has been set.

### SetPassportIssueDateNil

`func (o *SaudiArabiaNafathVisitorAttributes) SetPassportIssueDateNil(b bool)`

 SetPassportIssueDateNil sets the value for PassportIssueDate to be an explicit nil

### UnsetPassportIssueDate
`func (o *SaudiArabiaNafathVisitorAttributes) UnsetPassportIssueDate()`

UnsetPassportIssueDate ensures that no value is present for PassportIssueDate, not even an explicit nil
### GetPassportExpiryDate

`func (o *SaudiArabiaNafathVisitorAttributes) GetPassportExpiryDate() string`

GetPassportExpiryDate returns the PassportExpiryDate field if non-nil, zero value otherwise.

### GetPassportExpiryDateOk

`func (o *SaudiArabiaNafathVisitorAttributes) GetPassportExpiryDateOk() (*string, bool)`

GetPassportExpiryDateOk returns a tuple with the PassportExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassportExpiryDate

`func (o *SaudiArabiaNafathVisitorAttributes) SetPassportExpiryDate(v string)`

SetPassportExpiryDate sets PassportExpiryDate field to given value.

### HasPassportExpiryDate

`func (o *SaudiArabiaNafathVisitorAttributes) HasPassportExpiryDate() bool`

HasPassportExpiryDate returns a boolean if a field has been set.

### SetPassportExpiryDateNil

`func (o *SaudiArabiaNafathVisitorAttributes) SetPassportExpiryDateNil(b bool)`

 SetPassportExpiryDateNil sets the value for PassportExpiryDate to be an explicit nil

### UnsetPassportExpiryDate
`func (o *SaudiArabiaNafathVisitorAttributes) UnsetPassportExpiryDate()`

UnsetPassportExpiryDate ensures that no value is present for PassportExpiryDate, not even an explicit nil
### GetOccupationDescriptionArabic

`func (o *SaudiArabiaNafathVisitorAttributes) GetOccupationDescriptionArabic() string`

GetOccupationDescriptionArabic returns the OccupationDescriptionArabic field if non-nil, zero value otherwise.

### GetOccupationDescriptionArabicOk

`func (o *SaudiArabiaNafathVisitorAttributes) GetOccupationDescriptionArabicOk() (*string, bool)`

GetOccupationDescriptionArabicOk returns a tuple with the OccupationDescriptionArabic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccupationDescriptionArabic

`func (o *SaudiArabiaNafathVisitorAttributes) SetOccupationDescriptionArabic(v string)`

SetOccupationDescriptionArabic sets OccupationDescriptionArabic field to given value.

### HasOccupationDescriptionArabic

`func (o *SaudiArabiaNafathVisitorAttributes) HasOccupationDescriptionArabic() bool`

HasOccupationDescriptionArabic returns a boolean if a field has been set.

### SetOccupationDescriptionArabicNil

`func (o *SaudiArabiaNafathVisitorAttributes) SetOccupationDescriptionArabicNil(b bool)`

 SetOccupationDescriptionArabicNil sets the value for OccupationDescriptionArabic to be an explicit nil

### UnsetOccupationDescriptionArabic
`func (o *SaudiArabiaNafathVisitorAttributes) UnsetOccupationDescriptionArabic()`

UnsetOccupationDescriptionArabic ensures that no value is present for OccupationDescriptionArabic, not even an explicit nil
### GetOccupationDescriptionEnglish

`func (o *SaudiArabiaNafathVisitorAttributes) GetOccupationDescriptionEnglish() string`

GetOccupationDescriptionEnglish returns the OccupationDescriptionEnglish field if non-nil, zero value otherwise.

### GetOccupationDescriptionEnglishOk

`func (o *SaudiArabiaNafathVisitorAttributes) GetOccupationDescriptionEnglishOk() (*string, bool)`

GetOccupationDescriptionEnglishOk returns a tuple with the OccupationDescriptionEnglish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccupationDescriptionEnglish

`func (o *SaudiArabiaNafathVisitorAttributes) SetOccupationDescriptionEnglish(v string)`

SetOccupationDescriptionEnglish sets OccupationDescriptionEnglish field to given value.

### HasOccupationDescriptionEnglish

`func (o *SaudiArabiaNafathVisitorAttributes) HasOccupationDescriptionEnglish() bool`

HasOccupationDescriptionEnglish returns a boolean if a field has been set.

### SetOccupationDescriptionEnglishNil

`func (o *SaudiArabiaNafathVisitorAttributes) SetOccupationDescriptionEnglishNil(b bool)`

 SetOccupationDescriptionEnglishNil sets the value for OccupationDescriptionEnglish to be an explicit nil

### UnsetOccupationDescriptionEnglish
`func (o *SaudiArabiaNafathVisitorAttributes) UnsetOccupationDescriptionEnglish()`

UnsetOccupationDescriptionEnglish ensures that no value is present for OccupationDescriptionEnglish, not even an explicit nil
### GetBirthCountryArabic

`func (o *SaudiArabiaNafathVisitorAttributes) GetBirthCountryArabic() string`

GetBirthCountryArabic returns the BirthCountryArabic field if non-nil, zero value otherwise.

### GetBirthCountryArabicOk

`func (o *SaudiArabiaNafathVisitorAttributes) GetBirthCountryArabicOk() (*string, bool)`

GetBirthCountryArabicOk returns a tuple with the BirthCountryArabic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthCountryArabic

`func (o *SaudiArabiaNafathVisitorAttributes) SetBirthCountryArabic(v string)`

SetBirthCountryArabic sets BirthCountryArabic field to given value.

### HasBirthCountryArabic

`func (o *SaudiArabiaNafathVisitorAttributes) HasBirthCountryArabic() bool`

HasBirthCountryArabic returns a boolean if a field has been set.

### SetBirthCountryArabicNil

`func (o *SaudiArabiaNafathVisitorAttributes) SetBirthCountryArabicNil(b bool)`

 SetBirthCountryArabicNil sets the value for BirthCountryArabic to be an explicit nil

### UnsetBirthCountryArabic
`func (o *SaudiArabiaNafathVisitorAttributes) UnsetBirthCountryArabic()`

UnsetBirthCountryArabic ensures that no value is present for BirthCountryArabic, not even an explicit nil
### GetBirthCountryEnglish

`func (o *SaudiArabiaNafathVisitorAttributes) GetBirthCountryEnglish() string`

GetBirthCountryEnglish returns the BirthCountryEnglish field if non-nil, zero value otherwise.

### GetBirthCountryEnglishOk

`func (o *SaudiArabiaNafathVisitorAttributes) GetBirthCountryEnglishOk() (*string, bool)`

GetBirthCountryEnglishOk returns a tuple with the BirthCountryEnglish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthCountryEnglish

`func (o *SaudiArabiaNafathVisitorAttributes) SetBirthCountryEnglish(v string)`

SetBirthCountryEnglish sets BirthCountryEnglish field to given value.

### HasBirthCountryEnglish

`func (o *SaudiArabiaNafathVisitorAttributes) HasBirthCountryEnglish() bool`

HasBirthCountryEnglish returns a boolean if a field has been set.

### SetBirthCountryEnglishNil

`func (o *SaudiArabiaNafathVisitorAttributes) SetBirthCountryEnglishNil(b bool)`

 SetBirthCountryEnglishNil sets the value for BirthCountryEnglish to be an explicit nil

### UnsetBirthCountryEnglish
`func (o *SaudiArabiaNafathVisitorAttributes) UnsetBirthCountryEnglish()`

UnsetBirthCountryEnglish ensures that no value is present for BirthCountryEnglish, not even an explicit nil
### GetEntryDate

`func (o *SaudiArabiaNafathVisitorAttributes) GetEntryDate() string`

GetEntryDate returns the EntryDate field if non-nil, zero value otherwise.

### GetEntryDateOk

`func (o *SaudiArabiaNafathVisitorAttributes) GetEntryDateOk() (*string, bool)`

GetEntryDateOk returns a tuple with the EntryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryDate

`func (o *SaudiArabiaNafathVisitorAttributes) SetEntryDate(v string)`

SetEntryDate sets EntryDate field to given value.

### HasEntryDate

`func (o *SaudiArabiaNafathVisitorAttributes) HasEntryDate() bool`

HasEntryDate returns a boolean if a field has been set.

### SetEntryDateNil

`func (o *SaudiArabiaNafathVisitorAttributes) SetEntryDateNil(b bool)`

 SetEntryDateNil sets the value for EntryDate to be an explicit nil

### UnsetEntryDate
`func (o *SaudiArabiaNafathVisitorAttributes) UnsetEntryDate()`

UnsetEntryDate ensures that no value is present for EntryDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


