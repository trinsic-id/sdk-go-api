# GuatemalaCuiProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FullName** | **string** | Full name, including given and family names of the CUI holder. | 
**GivenName** | **string** | Given name(s) of the CUI holder. | 
**FamilyName** | **string** | Family name(s) of the CUI holder, including maternal and paternal names. | 
**DateOfBirth** | Pointer to **NullableString** | Date of birth of the CUI holder. | [optional] 
**DocumentNumber** | **string** | The Guatemalan Código Único de Identificación (CUI) number.              Assigned and maintained by RENAP (Registro Nacional de las Personas). Official format: exactly 13 numeric digits. 8 RENAP-assigned serial digits, 1 verifier digit (dígito verificador), and 4 geographic digits for department and municipality of birth. The CUI is printed on the Documento Personal de Identificación (DPI) in three groups (4–5–4) separated by spaces.              Trinsic normalizes to digits-only when returning the result, automatically removing spaces, dots, hyphens, and other non-alphanumeric characters.              No verifier algorithm appears in publicly accessible RENAP resources. Community-maintained validators often use modulus-11 (non-official). | 
**SerialDigits** | **string** | The first 8 digits of the CUI: the portion RENAP assigns from its registration system, before the verifier digit and the four-place geographic codes. Called the “serial” numbers in community-maintained validators. | 
**VerifierDigit** | **string** | The 9th digit of the CUI (verifier digit / dígito verificador). No verifier algorithm appears in publicly accessible RENAP resources.              However, community-maintained validators often use modulus-11 (non-official). | 
**GeographicDigits** | **string** | The last four digits encode birthplace department and municipality. There is not a publicly accessible dataset from RENAP for geographic codes. The INE (Instituto Nacional de Estadística y Censos) provides a dictionary of variables (Educación Formal 2024) that includes statistical codes for decoding departments and municipalities. Be careful using these resources, as these codes are not official RENAP codes. They may not always map to the correct department and municipality for consular registrations, naturalizations, foreign residents, and other edge cases.              For the dataset, see: https://datos.ine.gob.gt/dataset/educacion-formal-2024 under \&quot;Diccionario de Variables (Educación Formal 2024)\&quot; | 
**Sex** | **string** | Sex of the CUI holder.              Possible values: - Male - Female - Unknown | 
**ArrayName** | **[]string** | All names of the CUI holder, as an array of strings (e.g. given and family name parts). | 

## Methods

### NewGuatemalaCuiProviderOutput

`func NewGuatemalaCuiProviderOutput(fullName string, givenName string, familyName string, documentNumber string, serialDigits string, verifierDigit string, geographicDigits string, sex string, arrayName []string, ) *GuatemalaCuiProviderOutput`

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


