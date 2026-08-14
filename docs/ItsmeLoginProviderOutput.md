# ItsmeLoginProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | Pointer to **NullableString** | The first name of the verified individual | [optional] 
**LastName** | Pointer to **NullableString** | The last name of the verified individual | [optional] 
**DateOfBirth** | Pointer to **NullableString** | The date of birth of the verified individual | [optional] 
**HashedNationalRegisterNumber** | Pointer to **NullableString** | The hashed version of the Belgian National Register Number of the verified individual.              By default, itsme does not return the raw National Register Number of the individual; instead, only a hashed version is returned.              Your account must be approved by itsme to receive the raw, unhashed National Register Number. | [optional] 
**NationalRegisterNumber** | Pointer to **NullableString** | The raw (not hashed) Belgian National Register Number (\&quot;Rijksregisternummer\&quot;) of the verified individual.              Only returned if your account has been explicitly authorized to receive it by itsme; by law, this data is considered sensitive personal data.              This is an 11-digit number in the format YYMMDDXXXCC, where: - YYMMDD represents the individual&#39;s date of birth (year, month, day). - XXX is a sequential birth number, odd for females and even for males. - CC is a checksum, calculated with the equation: 97 - (YYMMDDXXX mod 97)              For births in the year 2000 or later, the digit &#39;2&#39; is prepended to the first 9 digits during checksum calculation. | [optional] 
**Email** | Pointer to **NullableString** | The individual&#39;s email address. | [optional] 
**PhoneNumber** | Pointer to **NullableString** | The individual&#39;s phone number in international format. | [optional] 
**Gender** | Pointer to **NullableString** | The individual&#39;s gender.              Possible values: - Male - Female - Unknown - Not Applicable | [optional] 
**Nationality** | Pointer to **NullableString** | The individual&#39;s nationality as an ISO 3166-1 alpha-3 code. | [optional] 
**BirthPlace** | Pointer to **NullableString** | The individual&#39;s place of birth. | [optional] 
**DocumentNumber** | Pointer to **NullableString** | The document number | [optional] 
**IdentityDocumentExpirationDate** | Pointer to **NullableString** | The expiration date of the identity document. | [optional] 
**Language** | Pointer to **NullableString** | The individual&#39;s language as an ISO 639-1 code. Expected values: NL, FR, DE, EN. | [optional] 
**Address** | Pointer to [**NullableItsmeAddress**](ItsmeAddress.md) | The individual&#39;s address | [optional] 

## Methods

### NewItsmeLoginProviderOutput

`func NewItsmeLoginProviderOutput() *ItsmeLoginProviderOutput`

NewItsmeLoginProviderOutput instantiates a new ItsmeLoginProviderOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItsmeLoginProviderOutputWithDefaults

`func NewItsmeLoginProviderOutputWithDefaults() *ItsmeLoginProviderOutput`

NewItsmeLoginProviderOutputWithDefaults instantiates a new ItsmeLoginProviderOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *ItsmeLoginProviderOutput) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *ItsmeLoginProviderOutput) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *ItsmeLoginProviderOutput) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *ItsmeLoginProviderOutput) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *ItsmeLoginProviderOutput) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *ItsmeLoginProviderOutput) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *ItsmeLoginProviderOutput) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *ItsmeLoginProviderOutput) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *ItsmeLoginProviderOutput) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *ItsmeLoginProviderOutput) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *ItsmeLoginProviderOutput) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *ItsmeLoginProviderOutput) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetDateOfBirth

`func (o *ItsmeLoginProviderOutput) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *ItsmeLoginProviderOutput) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *ItsmeLoginProviderOutput) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *ItsmeLoginProviderOutput) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### SetDateOfBirthNil

`func (o *ItsmeLoginProviderOutput) SetDateOfBirthNil(b bool)`

 SetDateOfBirthNil sets the value for DateOfBirth to be an explicit nil

### UnsetDateOfBirth
`func (o *ItsmeLoginProviderOutput) UnsetDateOfBirth()`

UnsetDateOfBirth ensures that no value is present for DateOfBirth, not even an explicit nil
### GetHashedNationalRegisterNumber

`func (o *ItsmeLoginProviderOutput) GetHashedNationalRegisterNumber() string`

GetHashedNationalRegisterNumber returns the HashedNationalRegisterNumber field if non-nil, zero value otherwise.

### GetHashedNationalRegisterNumberOk

`func (o *ItsmeLoginProviderOutput) GetHashedNationalRegisterNumberOk() (*string, bool)`

GetHashedNationalRegisterNumberOk returns a tuple with the HashedNationalRegisterNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashedNationalRegisterNumber

`func (o *ItsmeLoginProviderOutput) SetHashedNationalRegisterNumber(v string)`

SetHashedNationalRegisterNumber sets HashedNationalRegisterNumber field to given value.

### HasHashedNationalRegisterNumber

`func (o *ItsmeLoginProviderOutput) HasHashedNationalRegisterNumber() bool`

HasHashedNationalRegisterNumber returns a boolean if a field has been set.

### SetHashedNationalRegisterNumberNil

`func (o *ItsmeLoginProviderOutput) SetHashedNationalRegisterNumberNil(b bool)`

 SetHashedNationalRegisterNumberNil sets the value for HashedNationalRegisterNumber to be an explicit nil

### UnsetHashedNationalRegisterNumber
`func (o *ItsmeLoginProviderOutput) UnsetHashedNationalRegisterNumber()`

UnsetHashedNationalRegisterNumber ensures that no value is present for HashedNationalRegisterNumber, not even an explicit nil
### GetNationalRegisterNumber

`func (o *ItsmeLoginProviderOutput) GetNationalRegisterNumber() string`

GetNationalRegisterNumber returns the NationalRegisterNumber field if non-nil, zero value otherwise.

### GetNationalRegisterNumberOk

`func (o *ItsmeLoginProviderOutput) GetNationalRegisterNumberOk() (*string, bool)`

GetNationalRegisterNumberOk returns a tuple with the NationalRegisterNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationalRegisterNumber

`func (o *ItsmeLoginProviderOutput) SetNationalRegisterNumber(v string)`

SetNationalRegisterNumber sets NationalRegisterNumber field to given value.

### HasNationalRegisterNumber

`func (o *ItsmeLoginProviderOutput) HasNationalRegisterNumber() bool`

HasNationalRegisterNumber returns a boolean if a field has been set.

### SetNationalRegisterNumberNil

`func (o *ItsmeLoginProviderOutput) SetNationalRegisterNumberNil(b bool)`

 SetNationalRegisterNumberNil sets the value for NationalRegisterNumber to be an explicit nil

### UnsetNationalRegisterNumber
`func (o *ItsmeLoginProviderOutput) UnsetNationalRegisterNumber()`

UnsetNationalRegisterNumber ensures that no value is present for NationalRegisterNumber, not even an explicit nil
### GetEmail

`func (o *ItsmeLoginProviderOutput) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ItsmeLoginProviderOutput) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ItsmeLoginProviderOutput) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ItsmeLoginProviderOutput) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *ItsmeLoginProviderOutput) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *ItsmeLoginProviderOutput) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetPhoneNumber

`func (o *ItsmeLoginProviderOutput) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *ItsmeLoginProviderOutput) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *ItsmeLoginProviderOutput) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *ItsmeLoginProviderOutput) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### SetPhoneNumberNil

`func (o *ItsmeLoginProviderOutput) SetPhoneNumberNil(b bool)`

 SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil

### UnsetPhoneNumber
`func (o *ItsmeLoginProviderOutput) UnsetPhoneNumber()`

UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
### GetGender

`func (o *ItsmeLoginProviderOutput) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *ItsmeLoginProviderOutput) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *ItsmeLoginProviderOutput) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *ItsmeLoginProviderOutput) HasGender() bool`

HasGender returns a boolean if a field has been set.

### SetGenderNil

`func (o *ItsmeLoginProviderOutput) SetGenderNil(b bool)`

 SetGenderNil sets the value for Gender to be an explicit nil

### UnsetGender
`func (o *ItsmeLoginProviderOutput) UnsetGender()`

UnsetGender ensures that no value is present for Gender, not even an explicit nil
### GetNationality

`func (o *ItsmeLoginProviderOutput) GetNationality() string`

GetNationality returns the Nationality field if non-nil, zero value otherwise.

### GetNationalityOk

`func (o *ItsmeLoginProviderOutput) GetNationalityOk() (*string, bool)`

GetNationalityOk returns a tuple with the Nationality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationality

`func (o *ItsmeLoginProviderOutput) SetNationality(v string)`

SetNationality sets Nationality field to given value.

### HasNationality

`func (o *ItsmeLoginProviderOutput) HasNationality() bool`

HasNationality returns a boolean if a field has been set.

### SetNationalityNil

`func (o *ItsmeLoginProviderOutput) SetNationalityNil(b bool)`

 SetNationalityNil sets the value for Nationality to be an explicit nil

### UnsetNationality
`func (o *ItsmeLoginProviderOutput) UnsetNationality()`

UnsetNationality ensures that no value is present for Nationality, not even an explicit nil
### GetBirthPlace

`func (o *ItsmeLoginProviderOutput) GetBirthPlace() string`

GetBirthPlace returns the BirthPlace field if non-nil, zero value otherwise.

### GetBirthPlaceOk

`func (o *ItsmeLoginProviderOutput) GetBirthPlaceOk() (*string, bool)`

GetBirthPlaceOk returns a tuple with the BirthPlace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthPlace

`func (o *ItsmeLoginProviderOutput) SetBirthPlace(v string)`

SetBirthPlace sets BirthPlace field to given value.

### HasBirthPlace

`func (o *ItsmeLoginProviderOutput) HasBirthPlace() bool`

HasBirthPlace returns a boolean if a field has been set.

### SetBirthPlaceNil

`func (o *ItsmeLoginProviderOutput) SetBirthPlaceNil(b bool)`

 SetBirthPlaceNil sets the value for BirthPlace to be an explicit nil

### UnsetBirthPlace
`func (o *ItsmeLoginProviderOutput) UnsetBirthPlace()`

UnsetBirthPlace ensures that no value is present for BirthPlace, not even an explicit nil
### GetDocumentNumber

`func (o *ItsmeLoginProviderOutput) GetDocumentNumber() string`

GetDocumentNumber returns the DocumentNumber field if non-nil, zero value otherwise.

### GetDocumentNumberOk

`func (o *ItsmeLoginProviderOutput) GetDocumentNumberOk() (*string, bool)`

GetDocumentNumberOk returns a tuple with the DocumentNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentNumber

`func (o *ItsmeLoginProviderOutput) SetDocumentNumber(v string)`

SetDocumentNumber sets DocumentNumber field to given value.

### HasDocumentNumber

`func (o *ItsmeLoginProviderOutput) HasDocumentNumber() bool`

HasDocumentNumber returns a boolean if a field has been set.

### SetDocumentNumberNil

`func (o *ItsmeLoginProviderOutput) SetDocumentNumberNil(b bool)`

 SetDocumentNumberNil sets the value for DocumentNumber to be an explicit nil

### UnsetDocumentNumber
`func (o *ItsmeLoginProviderOutput) UnsetDocumentNumber()`

UnsetDocumentNumber ensures that no value is present for DocumentNumber, not even an explicit nil
### GetIdentityDocumentExpirationDate

`func (o *ItsmeLoginProviderOutput) GetIdentityDocumentExpirationDate() string`

GetIdentityDocumentExpirationDate returns the IdentityDocumentExpirationDate field if non-nil, zero value otherwise.

### GetIdentityDocumentExpirationDateOk

`func (o *ItsmeLoginProviderOutput) GetIdentityDocumentExpirationDateOk() (*string, bool)`

GetIdentityDocumentExpirationDateOk returns a tuple with the IdentityDocumentExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityDocumentExpirationDate

`func (o *ItsmeLoginProviderOutput) SetIdentityDocumentExpirationDate(v string)`

SetIdentityDocumentExpirationDate sets IdentityDocumentExpirationDate field to given value.

### HasIdentityDocumentExpirationDate

`func (o *ItsmeLoginProviderOutput) HasIdentityDocumentExpirationDate() bool`

HasIdentityDocumentExpirationDate returns a boolean if a field has been set.

### SetIdentityDocumentExpirationDateNil

`func (o *ItsmeLoginProviderOutput) SetIdentityDocumentExpirationDateNil(b bool)`

 SetIdentityDocumentExpirationDateNil sets the value for IdentityDocumentExpirationDate to be an explicit nil

### UnsetIdentityDocumentExpirationDate
`func (o *ItsmeLoginProviderOutput) UnsetIdentityDocumentExpirationDate()`

UnsetIdentityDocumentExpirationDate ensures that no value is present for IdentityDocumentExpirationDate, not even an explicit nil
### GetLanguage

`func (o *ItsmeLoginProviderOutput) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *ItsmeLoginProviderOutput) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *ItsmeLoginProviderOutput) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *ItsmeLoginProviderOutput) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### SetLanguageNil

`func (o *ItsmeLoginProviderOutput) SetLanguageNil(b bool)`

 SetLanguageNil sets the value for Language to be an explicit nil

### UnsetLanguage
`func (o *ItsmeLoginProviderOutput) UnsetLanguage()`

UnsetLanguage ensures that no value is present for Language, not even an explicit nil
### GetAddress

`func (o *ItsmeLoginProviderOutput) GetAddress() ItsmeAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ItsmeLoginProviderOutput) GetAddressOk() (*ItsmeAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ItsmeLoginProviderOutput) SetAddress(v ItsmeAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ItsmeLoginProviderOutput) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *ItsmeLoginProviderOutput) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *ItsmeLoginProviderOutput) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


