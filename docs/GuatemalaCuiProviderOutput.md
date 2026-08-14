# GuatemalaCuiProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FullName** | Pointer to **NullableString** | Full name, including given and family names of the CUI holder. | [optional] 
**GivenName** | Pointer to **NullableString** | Given name(s) of the CUI holder. | [optional] 
**FamilyName** | Pointer to **NullableString** | Family name(s) of the CUI holder, including maternal and paternal names. | [optional] 
**DateOfBirth** | Pointer to **NullableString** | Date of birth of the CUI holder. | [optional] 
**DocumentNumber** | Pointer to **NullableString** | The Guatemalan Código Único de Identificación (CUI) number.              Assigned and maintained by RENAP (Registro Nacional de las Personas). Official format: exactly 13 numeric digits. 8 RENAP-assigned serial digits, 1 verifier digit (dígito verificador), and 4 geographic digits for department and municipality of birth. The CUI is printed on the Documento Personal de Identificación (DPI) in three groups (4–5–4) separated by spaces.              Trinsic normalizes to digits-only when returning the result, automatically removing spaces, dots, hyphens, and other non-alphanumeric characters.              No verifier algorithm appears in publicly accessible RENAP resources. Community-maintained validators often use modulus-11 (non-official). | [optional] 
**SerialDigits** | Pointer to **NullableString** | The first 8 digits of the CUI: the portion RENAP assigns from its registration system, before the verifier digit and the four-place geographic codes. Called the “serial” numbers in community-maintained validators. | [optional] 
**VerifierDigit** | Pointer to **NullableString** | The 9th digit of the CUI (verifier digit / dígito verificador). No verifier algorithm appears in publicly accessible RENAP resources.              However, community-maintained validators often use modulus-11 (non-official). | [optional] 
**GeographicDigits** | Pointer to **NullableString** | The last four digits encode birthplace department and municipality. There is not a publicly accessible dataset from RENAP for geographic codes. The INE (Instituto Nacional de Estadística y Censos) provides a dictionary of variables (Educación Formal 2024) that includes statistical codes for decoding departments and municipalities. Be careful using these resources, as these codes are not official RENAP codes. They may not always map to the correct department and municipality for consular registrations, naturalizations, foreign residents, and other edge cases.              For the dataset, see: https://datos.ine.gob.gt/dataset/educacion-formal-2024 under \&quot;Diccionario de Variables (Educación Formal 2024)\&quot; | [optional] 
**Sex** | Pointer to **NullableString** | Sex of the CUI holder.              Possible values: - Male - Female - Unknown | [optional] 
**ArrayName** | Pointer to **[]string** | All names of the CUI holder, as an array of strings (e.g. given and family name parts). | [optional] 

## Methods

### NewGuatemalaCuiProviderOutput

`func NewGuatemalaCuiProviderOutput() *GuatemalaCuiProviderOutput`

NewGuatemalaCuiProviderOutput instantiates a new GuatemalaCuiProviderOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuatemalaCuiProviderOutputWithDefaults

`func NewGuatemalaCuiProviderOutputWithDefaults() *GuatemalaCuiProviderOutput`

NewGuatemalaCuiProviderOutputWithDefaults instantiates a new GuatemalaCuiProviderOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFullName

`func (o *GuatemalaCuiProviderOutput) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *GuatemalaCuiProviderOutput) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *GuatemalaCuiProviderOutput) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *GuatemalaCuiProviderOutput) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### SetFullNameNil

`func (o *GuatemalaCuiProviderOutput) SetFullNameNil(b bool)`

 SetFullNameNil sets the value for FullName to be an explicit nil

### UnsetFullName
`func (o *GuatemalaCuiProviderOutput) UnsetFullName()`

UnsetFullName ensures that no value is present for FullName, not even an explicit nil
### GetGivenName

`func (o *GuatemalaCuiProviderOutput) GetGivenName() string`

GetGivenName returns the GivenName field if non-nil, zero value otherwise.

### GetGivenNameOk

`func (o *GuatemalaCuiProviderOutput) GetGivenNameOk() (*string, bool)`

GetGivenNameOk returns a tuple with the GivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenName

`func (o *GuatemalaCuiProviderOutput) SetGivenName(v string)`

SetGivenName sets GivenName field to given value.

### HasGivenName

`func (o *GuatemalaCuiProviderOutput) HasGivenName() bool`

HasGivenName returns a boolean if a field has been set.

### SetGivenNameNil

`func (o *GuatemalaCuiProviderOutput) SetGivenNameNil(b bool)`

 SetGivenNameNil sets the value for GivenName to be an explicit nil

### UnsetGivenName
`func (o *GuatemalaCuiProviderOutput) UnsetGivenName()`

UnsetGivenName ensures that no value is present for GivenName, not even an explicit nil
### GetFamilyName

`func (o *GuatemalaCuiProviderOutput) GetFamilyName() string`

GetFamilyName returns the FamilyName field if non-nil, zero value otherwise.

### GetFamilyNameOk

`func (o *GuatemalaCuiProviderOutput) GetFamilyNameOk() (*string, bool)`

GetFamilyNameOk returns a tuple with the FamilyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyName

`func (o *GuatemalaCuiProviderOutput) SetFamilyName(v string)`

SetFamilyName sets FamilyName field to given value.

### HasFamilyName

`func (o *GuatemalaCuiProviderOutput) HasFamilyName() bool`

HasFamilyName returns a boolean if a field has been set.

### SetFamilyNameNil

`func (o *GuatemalaCuiProviderOutput) SetFamilyNameNil(b bool)`

 SetFamilyNameNil sets the value for FamilyName to be an explicit nil

### UnsetFamilyName
`func (o *GuatemalaCuiProviderOutput) UnsetFamilyName()`

UnsetFamilyName ensures that no value is present for FamilyName, not even an explicit nil
### GetDateOfBirth

`func (o *GuatemalaCuiProviderOutput) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *GuatemalaCuiProviderOutput) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *GuatemalaCuiProviderOutput) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *GuatemalaCuiProviderOutput) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### SetDateOfBirthNil

`func (o *GuatemalaCuiProviderOutput) SetDateOfBirthNil(b bool)`

 SetDateOfBirthNil sets the value for DateOfBirth to be an explicit nil

### UnsetDateOfBirth
`func (o *GuatemalaCuiProviderOutput) UnsetDateOfBirth()`

UnsetDateOfBirth ensures that no value is present for DateOfBirth, not even an explicit nil
### GetDocumentNumber

`func (o *GuatemalaCuiProviderOutput) GetDocumentNumber() string`

GetDocumentNumber returns the DocumentNumber field if non-nil, zero value otherwise.

### GetDocumentNumberOk

`func (o *GuatemalaCuiProviderOutput) GetDocumentNumberOk() (*string, bool)`

GetDocumentNumberOk returns a tuple with the DocumentNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentNumber

`func (o *GuatemalaCuiProviderOutput) SetDocumentNumber(v string)`

SetDocumentNumber sets DocumentNumber field to given value.

### HasDocumentNumber

`func (o *GuatemalaCuiProviderOutput) HasDocumentNumber() bool`

HasDocumentNumber returns a boolean if a field has been set.

### SetDocumentNumberNil

`func (o *GuatemalaCuiProviderOutput) SetDocumentNumberNil(b bool)`

 SetDocumentNumberNil sets the value for DocumentNumber to be an explicit nil

### UnsetDocumentNumber
`func (o *GuatemalaCuiProviderOutput) UnsetDocumentNumber()`

UnsetDocumentNumber ensures that no value is present for DocumentNumber, not even an explicit nil
### GetSerialDigits

`func (o *GuatemalaCuiProviderOutput) GetSerialDigits() string`

GetSerialDigits returns the SerialDigits field if non-nil, zero value otherwise.

### GetSerialDigitsOk

`func (o *GuatemalaCuiProviderOutput) GetSerialDigitsOk() (*string, bool)`

GetSerialDigitsOk returns a tuple with the SerialDigits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialDigits

`func (o *GuatemalaCuiProviderOutput) SetSerialDigits(v string)`

SetSerialDigits sets SerialDigits field to given value.

### HasSerialDigits

`func (o *GuatemalaCuiProviderOutput) HasSerialDigits() bool`

HasSerialDigits returns a boolean if a field has been set.

### SetSerialDigitsNil

`func (o *GuatemalaCuiProviderOutput) SetSerialDigitsNil(b bool)`

 SetSerialDigitsNil sets the value for SerialDigits to be an explicit nil

### UnsetSerialDigits
`func (o *GuatemalaCuiProviderOutput) UnsetSerialDigits()`

UnsetSerialDigits ensures that no value is present for SerialDigits, not even an explicit nil
### GetVerifierDigit

`func (o *GuatemalaCuiProviderOutput) GetVerifierDigit() string`

GetVerifierDigit returns the VerifierDigit field if non-nil, zero value otherwise.

### GetVerifierDigitOk

`func (o *GuatemalaCuiProviderOutput) GetVerifierDigitOk() (*string, bool)`

GetVerifierDigitOk returns a tuple with the VerifierDigit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifierDigit

`func (o *GuatemalaCuiProviderOutput) SetVerifierDigit(v string)`

SetVerifierDigit sets VerifierDigit field to given value.

### HasVerifierDigit

`func (o *GuatemalaCuiProviderOutput) HasVerifierDigit() bool`

HasVerifierDigit returns a boolean if a field has been set.

### SetVerifierDigitNil

`func (o *GuatemalaCuiProviderOutput) SetVerifierDigitNil(b bool)`

 SetVerifierDigitNil sets the value for VerifierDigit to be an explicit nil

### UnsetVerifierDigit
`func (o *GuatemalaCuiProviderOutput) UnsetVerifierDigit()`

UnsetVerifierDigit ensures that no value is present for VerifierDigit, not even an explicit nil
### GetGeographicDigits

`func (o *GuatemalaCuiProviderOutput) GetGeographicDigits() string`

GetGeographicDigits returns the GeographicDigits field if non-nil, zero value otherwise.

### GetGeographicDigitsOk

`func (o *GuatemalaCuiProviderOutput) GetGeographicDigitsOk() (*string, bool)`

GetGeographicDigitsOk returns a tuple with the GeographicDigits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeographicDigits

`func (o *GuatemalaCuiProviderOutput) SetGeographicDigits(v string)`

SetGeographicDigits sets GeographicDigits field to given value.

### HasGeographicDigits

`func (o *GuatemalaCuiProviderOutput) HasGeographicDigits() bool`

HasGeographicDigits returns a boolean if a field has been set.

### SetGeographicDigitsNil

`func (o *GuatemalaCuiProviderOutput) SetGeographicDigitsNil(b bool)`

 SetGeographicDigitsNil sets the value for GeographicDigits to be an explicit nil

### UnsetGeographicDigits
`func (o *GuatemalaCuiProviderOutput) UnsetGeographicDigits()`

UnsetGeographicDigits ensures that no value is present for GeographicDigits, not even an explicit nil
### GetSex

`func (o *GuatemalaCuiProviderOutput) GetSex() string`

GetSex returns the Sex field if non-nil, zero value otherwise.

### GetSexOk

`func (o *GuatemalaCuiProviderOutput) GetSexOk() (*string, bool)`

GetSexOk returns a tuple with the Sex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSex

`func (o *GuatemalaCuiProviderOutput) SetSex(v string)`

SetSex sets Sex field to given value.

### HasSex

`func (o *GuatemalaCuiProviderOutput) HasSex() bool`

HasSex returns a boolean if a field has been set.

### SetSexNil

`func (o *GuatemalaCuiProviderOutput) SetSexNil(b bool)`

 SetSexNil sets the value for Sex to be an explicit nil

### UnsetSex
`func (o *GuatemalaCuiProviderOutput) UnsetSex()`

UnsetSex ensures that no value is present for Sex, not even an explicit nil
### GetArrayName

`func (o *GuatemalaCuiProviderOutput) GetArrayName() []string`

GetArrayName returns the ArrayName field if non-nil, zero value otherwise.

### GetArrayNameOk

`func (o *GuatemalaCuiProviderOutput) GetArrayNameOk() (*[]string, bool)`

GetArrayNameOk returns a tuple with the ArrayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrayName

`func (o *GuatemalaCuiProviderOutput) SetArrayName(v []string)`

SetArrayName sets ArrayName field to given value.

### HasArrayName

`func (o *GuatemalaCuiProviderOutput) HasArrayName() bool`

HasArrayName returns a boolean if a field has been set.

### SetArrayNameNil

`func (o *GuatemalaCuiProviderOutput) SetArrayNameNil(b bool)`

 SetArrayNameNil sets the value for ArrayName to be an explicit nil

### UnsetArrayName
`func (o *GuatemalaCuiProviderOutput) UnsetArrayName()`

UnsetArrayName ensures that no value is present for ArrayName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


