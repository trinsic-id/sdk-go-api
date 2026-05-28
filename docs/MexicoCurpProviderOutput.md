# MexicoCurpProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | Pointer to **NullableString** | First name of the individual. | [optional] 
**FatherLastName** | Pointer to **NullableString** | The last name of the father of the individual. | [optional] 
**MotherLastName** | Pointer to **NullableString** | The last name of the mother of the individual. | [optional] 
**Gender** | Pointer to **NullableString** | The gender of the individual.   List of possible values:   - Male  - Female | [optional] 
**DateOfBirth** | Pointer to **NullableString** | The date of birth of the individual. | [optional] 
**CountryOfBirth** | Pointer to **NullableString** | The country of birth of the individual. | [optional] 
**StateOfBirth** | Pointer to **NullableString** | The state of birth of the individual. | [optional] 
**Curp** | Pointer to **NullableString** | The CURP (Clave Única de Registro de Población) that was verified.              A CURP code is a unique identifier assigned to individuals in Mexico. It is 18 alphanumeric characters, with a structured meaning:              ABCD YYMMDD G SS XYZ M C The first 4 letters (ABCD): A: The first letter of the paternal last name. B: First internal vowel of paternal last name. C: First letter of maternal last name. D: First letter of first name.              YYMMDD: 2-digit year, month, day.              G: Gender, H for Hombre (male) and M for Mujer (female).              SS: State code (2 letters), e.g. NL for Nuevo León.              X: First internal consonant of paternal last name. Y: First internal consonant of maternal last name. Z: First internal consonant of given name.              M: Millennium indicator (0-9 for pre-2000 birth date, A-Z for post-2000&#39;s birth dates.              C: Checksum character | [optional] 
**CurpStatus** | Pointer to **NullableString** | Curp status for the subject.              Possible values: - AN: Alta Normal (Normal registration) - Active - AH: Alta con Homonimia (Registration with homonymy) - Active - RCC: Registro de cambio afectando a CURP (Change affecting CURP) - Active - RCN: Registro de cambio no afectando a CURP (Change not affecting CURP) - Active - BAP: Baja por documento apócrifo (Low due to apocryphal document) - Inactive - BSU: Baja sin uso (Low curp without use) - Inactive - BD: Baja por defunción (Low curp due to death) - Inactive - BDM: Baja administrativa (Low, due to administrative process) - Inactive - BDP: Baja por adopción (Low, due to adoption) - Inactive - BJD: Baja Judicial (Low for judicial reasons) - Inactive | [optional] 
**RegistrationYear** | Pointer to **NullableInt32** | The year the CURP number was registered in. | [optional] 
**RegistrationState** | Pointer to **NullableString** | The state the CURP number was registered in. | [optional] 
**ActNumber** | Pointer to **NullableString** | The ACT (Número de Acta) number of the individual.              The Act number is a civil registry index number. Various state have various formats of specifying these. | [optional] 

## Methods

### NewMexicoCurpProviderOutput

`func NewMexicoCurpProviderOutput() *MexicoCurpProviderOutput`

NewMexicoCurpProviderOutput instantiates a new MexicoCurpProviderOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMexicoCurpProviderOutputWithDefaults

`func NewMexicoCurpProviderOutputWithDefaults() *MexicoCurpProviderOutput`

NewMexicoCurpProviderOutputWithDefaults instantiates a new MexicoCurpProviderOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *MexicoCurpProviderOutput) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *MexicoCurpProviderOutput) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *MexicoCurpProviderOutput) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *MexicoCurpProviderOutput) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *MexicoCurpProviderOutput) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *MexicoCurpProviderOutput) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetFatherLastName

`func (o *MexicoCurpProviderOutput) GetFatherLastName() string`

GetFatherLastName returns the FatherLastName field if non-nil, zero value otherwise.

### GetFatherLastNameOk

`func (o *MexicoCurpProviderOutput) GetFatherLastNameOk() (*string, bool)`

GetFatherLastNameOk returns a tuple with the FatherLastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFatherLastName

`func (o *MexicoCurpProviderOutput) SetFatherLastName(v string)`

SetFatherLastName sets FatherLastName field to given value.

### HasFatherLastName

`func (o *MexicoCurpProviderOutput) HasFatherLastName() bool`

HasFatherLastName returns a boolean if a field has been set.

### SetFatherLastNameNil

`func (o *MexicoCurpProviderOutput) SetFatherLastNameNil(b bool)`

 SetFatherLastNameNil sets the value for FatherLastName to be an explicit nil

### UnsetFatherLastName
`func (o *MexicoCurpProviderOutput) UnsetFatherLastName()`

UnsetFatherLastName ensures that no value is present for FatherLastName, not even an explicit nil
### GetMotherLastName

`func (o *MexicoCurpProviderOutput) GetMotherLastName() string`

GetMotherLastName returns the MotherLastName field if non-nil, zero value otherwise.

### GetMotherLastNameOk

`func (o *MexicoCurpProviderOutput) GetMotherLastNameOk() (*string, bool)`

GetMotherLastNameOk returns a tuple with the MotherLastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMotherLastName

`func (o *MexicoCurpProviderOutput) SetMotherLastName(v string)`

SetMotherLastName sets MotherLastName field to given value.

### HasMotherLastName

`func (o *MexicoCurpProviderOutput) HasMotherLastName() bool`

HasMotherLastName returns a boolean if a field has been set.

### SetMotherLastNameNil

`func (o *MexicoCurpProviderOutput) SetMotherLastNameNil(b bool)`

 SetMotherLastNameNil sets the value for MotherLastName to be an explicit nil

### UnsetMotherLastName
`func (o *MexicoCurpProviderOutput) UnsetMotherLastName()`

UnsetMotherLastName ensures that no value is present for MotherLastName, not even an explicit nil
### GetGender

`func (o *MexicoCurpProviderOutput) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *MexicoCurpProviderOutput) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *MexicoCurpProviderOutput) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *MexicoCurpProviderOutput) HasGender() bool`

HasGender returns a boolean if a field has been set.

### SetGenderNil

`func (o *MexicoCurpProviderOutput) SetGenderNil(b bool)`

 SetGenderNil sets the value for Gender to be an explicit nil

### UnsetGender
`func (o *MexicoCurpProviderOutput) UnsetGender()`

UnsetGender ensures that no value is present for Gender, not even an explicit nil
### GetDateOfBirth

`func (o *MexicoCurpProviderOutput) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *MexicoCurpProviderOutput) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *MexicoCurpProviderOutput) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *MexicoCurpProviderOutput) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### SetDateOfBirthNil

`func (o *MexicoCurpProviderOutput) SetDateOfBirthNil(b bool)`

 SetDateOfBirthNil sets the value for DateOfBirth to be an explicit nil

### UnsetDateOfBirth
`func (o *MexicoCurpProviderOutput) UnsetDateOfBirth()`

UnsetDateOfBirth ensures that no value is present for DateOfBirth, not even an explicit nil
### GetCountryOfBirth

`func (o *MexicoCurpProviderOutput) GetCountryOfBirth() string`

GetCountryOfBirth returns the CountryOfBirth field if non-nil, zero value otherwise.

### GetCountryOfBirthOk

`func (o *MexicoCurpProviderOutput) GetCountryOfBirthOk() (*string, bool)`

GetCountryOfBirthOk returns a tuple with the CountryOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryOfBirth

`func (o *MexicoCurpProviderOutput) SetCountryOfBirth(v string)`

SetCountryOfBirth sets CountryOfBirth field to given value.

### HasCountryOfBirth

`func (o *MexicoCurpProviderOutput) HasCountryOfBirth() bool`

HasCountryOfBirth returns a boolean if a field has been set.

### SetCountryOfBirthNil

`func (o *MexicoCurpProviderOutput) SetCountryOfBirthNil(b bool)`

 SetCountryOfBirthNil sets the value for CountryOfBirth to be an explicit nil

### UnsetCountryOfBirth
`func (o *MexicoCurpProviderOutput) UnsetCountryOfBirth()`

UnsetCountryOfBirth ensures that no value is present for CountryOfBirth, not even an explicit nil
### GetStateOfBirth

`func (o *MexicoCurpProviderOutput) GetStateOfBirth() string`

GetStateOfBirth returns the StateOfBirth field if non-nil, zero value otherwise.

### GetStateOfBirthOk

`func (o *MexicoCurpProviderOutput) GetStateOfBirthOk() (*string, bool)`

GetStateOfBirthOk returns a tuple with the StateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateOfBirth

`func (o *MexicoCurpProviderOutput) SetStateOfBirth(v string)`

SetStateOfBirth sets StateOfBirth field to given value.

### HasStateOfBirth

`func (o *MexicoCurpProviderOutput) HasStateOfBirth() bool`

HasStateOfBirth returns a boolean if a field has been set.

### SetStateOfBirthNil

`func (o *MexicoCurpProviderOutput) SetStateOfBirthNil(b bool)`

 SetStateOfBirthNil sets the value for StateOfBirth to be an explicit nil

### UnsetStateOfBirth
`func (o *MexicoCurpProviderOutput) UnsetStateOfBirth()`

UnsetStateOfBirth ensures that no value is present for StateOfBirth, not even an explicit nil
### GetCurp

`func (o *MexicoCurpProviderOutput) GetCurp() string`

GetCurp returns the Curp field if non-nil, zero value otherwise.

### GetCurpOk

`func (o *MexicoCurpProviderOutput) GetCurpOk() (*string, bool)`

GetCurpOk returns a tuple with the Curp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurp

`func (o *MexicoCurpProviderOutput) SetCurp(v string)`

SetCurp sets Curp field to given value.

### HasCurp

`func (o *MexicoCurpProviderOutput) HasCurp() bool`

HasCurp returns a boolean if a field has been set.

### SetCurpNil

`func (o *MexicoCurpProviderOutput) SetCurpNil(b bool)`

 SetCurpNil sets the value for Curp to be an explicit nil

### UnsetCurp
`func (o *MexicoCurpProviderOutput) UnsetCurp()`

UnsetCurp ensures that no value is present for Curp, not even an explicit nil
### GetCurpStatus

`func (o *MexicoCurpProviderOutput) GetCurpStatus() string`

GetCurpStatus returns the CurpStatus field if non-nil, zero value otherwise.

### GetCurpStatusOk

`func (o *MexicoCurpProviderOutput) GetCurpStatusOk() (*string, bool)`

GetCurpStatusOk returns a tuple with the CurpStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurpStatus

`func (o *MexicoCurpProviderOutput) SetCurpStatus(v string)`

SetCurpStatus sets CurpStatus field to given value.

### HasCurpStatus

`func (o *MexicoCurpProviderOutput) HasCurpStatus() bool`

HasCurpStatus returns a boolean if a field has been set.

### SetCurpStatusNil

`func (o *MexicoCurpProviderOutput) SetCurpStatusNil(b bool)`

 SetCurpStatusNil sets the value for CurpStatus to be an explicit nil

### UnsetCurpStatus
`func (o *MexicoCurpProviderOutput) UnsetCurpStatus()`

UnsetCurpStatus ensures that no value is present for CurpStatus, not even an explicit nil
### GetRegistrationYear

`func (o *MexicoCurpProviderOutput) GetRegistrationYear() int32`

GetRegistrationYear returns the RegistrationYear field if non-nil, zero value otherwise.

### GetRegistrationYearOk

`func (o *MexicoCurpProviderOutput) GetRegistrationYearOk() (*int32, bool)`

GetRegistrationYearOk returns a tuple with the RegistrationYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationYear

`func (o *MexicoCurpProviderOutput) SetRegistrationYear(v int32)`

SetRegistrationYear sets RegistrationYear field to given value.

### HasRegistrationYear

`func (o *MexicoCurpProviderOutput) HasRegistrationYear() bool`

HasRegistrationYear returns a boolean if a field has been set.

### SetRegistrationYearNil

`func (o *MexicoCurpProviderOutput) SetRegistrationYearNil(b bool)`

 SetRegistrationYearNil sets the value for RegistrationYear to be an explicit nil

### UnsetRegistrationYear
`func (o *MexicoCurpProviderOutput) UnsetRegistrationYear()`

UnsetRegistrationYear ensures that no value is present for RegistrationYear, not even an explicit nil
### GetRegistrationState

`func (o *MexicoCurpProviderOutput) GetRegistrationState() string`

GetRegistrationState returns the RegistrationState field if non-nil, zero value otherwise.

### GetRegistrationStateOk

`func (o *MexicoCurpProviderOutput) GetRegistrationStateOk() (*string, bool)`

GetRegistrationStateOk returns a tuple with the RegistrationState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationState

`func (o *MexicoCurpProviderOutput) SetRegistrationState(v string)`

SetRegistrationState sets RegistrationState field to given value.

### HasRegistrationState

`func (o *MexicoCurpProviderOutput) HasRegistrationState() bool`

HasRegistrationState returns a boolean if a field has been set.

### SetRegistrationStateNil

`func (o *MexicoCurpProviderOutput) SetRegistrationStateNil(b bool)`

 SetRegistrationStateNil sets the value for RegistrationState to be an explicit nil

### UnsetRegistrationState
`func (o *MexicoCurpProviderOutput) UnsetRegistrationState()`

UnsetRegistrationState ensures that no value is present for RegistrationState, not even an explicit nil
### GetActNumber

`func (o *MexicoCurpProviderOutput) GetActNumber() string`

GetActNumber returns the ActNumber field if non-nil, zero value otherwise.

### GetActNumberOk

`func (o *MexicoCurpProviderOutput) GetActNumberOk() (*string, bool)`

GetActNumberOk returns a tuple with the ActNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActNumber

`func (o *MexicoCurpProviderOutput) SetActNumber(v string)`

SetActNumber sets ActNumber field to given value.

### HasActNumber

`func (o *MexicoCurpProviderOutput) HasActNumber() bool`

HasActNumber returns a boolean if a field has been set.

### SetActNumberNil

`func (o *MexicoCurpProviderOutput) SetActNumberNil(b bool)`

 SetActNumberNil sets the value for ActNumber to be an explicit nil

### UnsetActNumber
`func (o *MexicoCurpProviderOutput) UnsetActNumber()`

UnsetActNumber ensures that no value is present for ActNumber, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


