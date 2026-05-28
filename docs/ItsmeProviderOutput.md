# ItsmeProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | **string** | The first name of the verified individual | 
**LastName** | **string** | The last name of the verified individual | 
**DateOfBirth** | **string** | The date of birth of the verified individual | 
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

### NewItsmeProviderOutput

`func NewItsmeProviderOutput(firstName string, lastName string, dateOfBirth string, ) *ItsmeProviderOutput`

NewItsmeProviderOutput instantiates a new ItsmeProviderOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItsmeProviderOutputWithDefaults

`func NewItsmeProviderOutputWithDefaults() *ItsmeProviderOutput`

NewItsmeProviderOutputWithDefaults instantiates a new ItsmeProviderOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *ItsmeProviderOutput) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *ItsmeProviderOutput) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *ItsmeProviderOutput) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *ItsmeProviderOutput) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *ItsmeProviderOutput) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *ItsmeProviderOutput) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetDateOfBirth

`func (o *ItsmeProviderOutput) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *ItsmeProviderOutput) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *ItsmeProviderOutput) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.


### GetHashedNationalRegisterNumber

`func (o *ItsmeProviderOutput) GetHashedNationalRegisterNumber() string`

GetHashedNationalRegisterNumber returns the HashedNationalRegisterNumber field if non-nil, zero value otherwise.

### GetHashedNationalRegisterNumberOk

`func (o *ItsmeProviderOutput) GetHashedNationalRegisterNumberOk() (*string, bool)`

GetHashedNationalRegisterNumberOk returns a tuple with the HashedNationalRegisterNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashedNationalRegisterNumber

`func (o *ItsmeProviderOutput) SetHashedNationalRegisterNumber(v string)`

SetHashedNationalRegisterNumber sets HashedNationalRegisterNumber field to given value.

### HasHashedNationalRegisterNumber

`func (o *ItsmeProviderOutput) HasHashedNationalRegisterNumber() bool`

HasHashedNationalRegisterNumber returns a boolean if a field has been set.

### SetHashedNationalRegisterNumberNil

`func (o *ItsmeProviderOutput) SetHashedNationalRegisterNumberNil(b bool)`

 SetHashedNationalRegisterNumberNil sets the value for HashedNationalRegisterNumber to be an explicit nil

### UnsetHashedNationalRegisterNumber
`func (o *ItsmeProviderOutput) UnsetHashedNationalRegisterNumber()`

UnsetHashedNationalRegisterNumber ensures that no value is present for HashedNationalRegisterNumber, not even an explicit nil
### GetNationalRegisterNumber

`func (o *ItsmeProviderOutput) GetNationalRegisterNumber() string`

GetNationalRegisterNumber returns the NationalRegisterNumber field if non-nil, zero value otherwise.

### GetNationalRegisterNumberOk

`func (o *ItsmeProviderOutput) GetNationalRegisterNumberOk() (*string, bool)`

GetNationalRegisterNumberOk returns a tuple with the NationalRegisterNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationalRegisterNumber

`func (o *ItsmeProviderOutput) SetNationalRegisterNumber(v string)`

SetNationalRegisterNumber sets NationalRegisterNumber field to given value.

### HasNationalRegisterNumber

`func (o *ItsmeProviderOutput) HasNationalRegisterNumber() bool`

HasNationalRegisterNumber returns a boolean if a field has been set.

### SetNationalRegisterNumberNil

`func (o *ItsmeProviderOutput) SetNationalRegisterNumberNil(b bool)`

 SetNationalRegisterNumberNil sets the value for NationalRegisterNumber to be an explicit nil

### UnsetNationalRegisterNumber
`func (o *ItsmeProviderOutput) UnsetNationalRegisterNumber()`

UnsetNationalRegisterNumber ensures that no value is present for NationalRegisterNumber, not even an explicit nil
### GetEmail

`func (o *ItsmeProviderOutput) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ItsmeProviderOutput) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ItsmeProviderOutput) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ItsmeProviderOutput) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *ItsmeProviderOutput) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *ItsmeProviderOutput) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetPhoneNumber

`func (o *ItsmeProviderOutput) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *ItsmeProviderOutput) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *ItsmeProviderOutput) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *ItsmeProviderOutput) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### SetPhoneNumberNil

`func (o *ItsmeProviderOutput) SetPhoneNumberNil(b bool)`

 SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil

### UnsetPhoneNumber
`func (o *ItsmeProviderOutput) UnsetPhoneNumber()`

UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
### GetGender

`func (o *ItsmeProviderOutput) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *ItsmeProviderOutput) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *ItsmeProviderOutput) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *ItsmeProviderOutput) HasGender() bool`

HasGender returns a boolean if a field has been set.

### SetGenderNil

`func (o *ItsmeProviderOutput) SetGenderNil(b bool)`

 SetGenderNil sets the value for Gender to be an explicit nil

### UnsetGender
`func (o *ItsmeProviderOutput) UnsetGender()`

UnsetGender ensures that no value is present for Gender, not even an explicit nil
### GetNationality

`func (o *ItsmeProviderOutput) GetNationality() string`

GetNationality returns the Nationality field if non-nil, zero value otherwise.

### GetNationalityOk

`func (o *ItsmeProviderOutput) GetNationalityOk() (*string, bool)`

GetNationalityOk returns a tuple with the Nationality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationality

`func (o *ItsmeProviderOutput) SetNationality(v string)`

SetNationality sets Nationality field to given value.

### HasNationality

`func (o *ItsmeProviderOutput) HasNationality() bool`

HasNationality returns a boolean if a field has been set.

### SetNationalityNil

`func (o *ItsmeProviderOutput) SetNationalityNil(b bool)`

 SetNationalityNil sets the value for Nationality to be an explicit nil

### UnsetNationality
`func (o *ItsmeProviderOutput) UnsetNationality()`

UnsetNationality ensures that no value is present for Nationality, not even an explicit nil
### GetBirthPlace

`func (o *ItsmeProviderOutput) GetBirthPlace() string`

GetBirthPlace returns the BirthPlace field if non-nil, zero value otherwise.

### GetBirthPlaceOk

`func (o *ItsmeProviderOutput) GetBirthPlaceOk() (*string, bool)`

GetBirthPlaceOk returns a tuple with the BirthPlace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthPlace

`func (o *ItsmeProviderOutput) SetBirthPlace(v string)`

SetBirthPlace sets BirthPlace field to given value.

### HasBirthPlace

`func (o *ItsmeProviderOutput) HasBirthPlace() bool`

HasBirthPlace returns a boolean if a field has been set.

### SetBirthPlaceNil

`func (o *ItsmeProviderOutput) SetBirthPlaceNil(b bool)`

 SetBirthPlaceNil sets the value for BirthPlace to be an explicit nil

### UnsetBirthPlace
`func (o *ItsmeProviderOutput) UnsetBirthPlace()`

UnsetBirthPlace ensures that no value is present for BirthPlace, not even an explicit nil
### GetDocumentNumber

`func (o *ItsmeProviderOutput) GetDocumentNumber() string`

GetDocumentNumber returns the DocumentNumber field if non-nil, zero value otherwise.

### GetDocumentNumberOk

`func (o *ItsmeProviderOutput) GetDocumentNumberOk() (*string, bool)`

GetDocumentNumberOk returns a tuple with the DocumentNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentNumber

`func (o *ItsmeProviderOutput) SetDocumentNumber(v string)`

SetDocumentNumber sets DocumentNumber field to given value.

### HasDocumentNumber

`func (o *ItsmeProviderOutput) HasDocumentNumber() bool`

HasDocumentNumber returns a boolean if a field has been set.

### SetDocumentNumberNil

`func (o *ItsmeProviderOutput) SetDocumentNumberNil(b bool)`

 SetDocumentNumberNil sets the value for DocumentNumber to be an explicit nil

### UnsetDocumentNumber
`func (o *ItsmeProviderOutput) UnsetDocumentNumber()`

UnsetDocumentNumber ensures that no value is present for DocumentNumber, not even an explicit nil
### GetIdentityDocumentExpirationDate

`func (o *ItsmeProviderOutput) GetIdentityDocumentExpirationDate() string`

GetIdentityDocumentExpirationDate returns the IdentityDocumentExpirationDate field if non-nil, zero value otherwise.

### GetIdentityDocumentExpirationDateOk

`func (o *ItsmeProviderOutput) GetIdentityDocumentExpirationDateOk() (*string, bool)`

GetIdentityDocumentExpirationDateOk returns a tuple with the IdentityDocumentExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityDocumentExpirationDate

`func (o *ItsmeProviderOutput) SetIdentityDocumentExpirationDate(v string)`

SetIdentityDocumentExpirationDate sets IdentityDocumentExpirationDate field to given value.

### HasIdentityDocumentExpirationDate

`func (o *ItsmeProviderOutput) HasIdentityDocumentExpirationDate() bool`

HasIdentityDocumentExpirationDate returns a boolean if a field has been set.

### SetIdentityDocumentExpirationDateNil

`func (o *ItsmeProviderOutput) SetIdentityDocumentExpirationDateNil(b bool)`

 SetIdentityDocumentExpirationDateNil sets the value for IdentityDocumentExpirationDate to be an explicit nil

### UnsetIdentityDocumentExpirationDate
`func (o *ItsmeProviderOutput) UnsetIdentityDocumentExpirationDate()`

UnsetIdentityDocumentExpirationDate ensures that no value is present for IdentityDocumentExpirationDate, not even an explicit nil
### GetLanguage

`func (o *ItsmeProviderOutput) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *ItsmeProviderOutput) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *ItsmeProviderOutput) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *ItsmeProviderOutput) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### SetLanguageNil

`func (o *ItsmeProviderOutput) SetLanguageNil(b bool)`

 SetLanguageNil sets the value for Language to be an explicit nil

### UnsetLanguage
`func (o *ItsmeProviderOutput) UnsetLanguage()`

UnsetLanguage ensures that no value is present for Language, not even an explicit nil
### GetAddress

`func (o *ItsmeProviderOutput) GetAddress() ItsmeAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ItsmeProviderOutput) GetAddressOk() (*ItsmeAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ItsmeProviderOutput) SetAddress(v ItsmeAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ItsmeProviderOutput) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *ItsmeProviderOutput) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *ItsmeProviderOutput) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


