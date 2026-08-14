# BrazilCpfProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PreviewRawErrorCode** | Pointer to **NullableString** | The raw error code for the verification, returned by Serpro.              This is a temporary preview to support detailed error handling. When building against this please let the Trinsic team know. Preview features can be turned off without an obsoletion notice. | [optional] 
**CpfNumber** | Pointer to **NullableString** | The individual&#39;s CPF number.              A CPF (Cadastro de Pessoas Físicas) is the Brazilian individual taxpayer identification number for residents and foreigners. It is an 11-digit number managed by the Federal Revenue Service, the last 2 digits are check digits, calculated based on the first 9 digits. The 9th digit indicates the fiscal region responsible for the registration, | [optional] 
**FiscalRegions** | Pointer to **[]string** | List of Brazilian fiscal regions that could have issued this number.              This is calculated based off of the CPF number. A CPF number encodes information about what region it is issued from, but some regions share the identifier.              Possible values: Formatted as two characters. The digit represents the region code this applies to. - AC (Acre) - 2 - AL (Alagoas) - 4 - AP (Amapá) - 2 - AM (Amazonas) - 2 - BA (Bahia) - 5 - CE (Ceará) - 3 - DF (Distrito Federal) - 1 - ES (Espírito Santo) - 7 - GO (Goiás) - 1 - MA (Maranhão) - 3 - MT (Mato Grosso) - 1 - MS (Mato Grosso do Sul) - 1 - MG (Minas Gerais) - 6 - PA (Pará) - 2 - PB (Paraíba) - 4 - PR (Paraná) - 9 - PE (Pernambuco) - 4 - PI (Piauí) - 3 - RJ (Rio de Janeiro) - 7 - RN (Rio Grande do Norte) - 4 - RO (Rondônia) - 2 - RS (Rio Grande do Sul) - 0 - RR (Roraima) - 2 - SC (Santa Catarina) - 9 - SP (São Paulo) - 8 - SE (Sergipe) - 5 - TO (Tocantins) - 1 | [optional] 
**StatusCode** | Pointer to **NullableInt32** | The CPF Status. Possible values: 0 - Regular. The CPF is valid and in good standing. There are no pending issues, registration inconsistencies, or missing mandatory Individual Income Tax Returns (DIRPF). 2 - Suspended. The CPF record contains incorrect or incomplete registration data (e.g., name, date of birth, or voter registration number). This status is typically caused by discrepancies with the Superior Electoral Court (TSE) database or failure to update personal information. 3 - Deceased Holder. The CPF has been registered as belonging to a deceased individual. This status is recorded to prevent post-mortem misuse. 4 - Pending Regularization. The taxpayer has failed to submit one or more mandatory Individual Income Tax Returns (DIRPF) within the last five years. This can occur for various reasons, such as failure to file a mandatory declaration, changes of address, name, marital status, or other important information that does not conform to the records. Regularization requires submission of the overdue returns. 5 - Canceled by Multiplicity. The CPF has been cancelled due to duplicate registration (multiple CPFs assigned to the same individual or a single CPF linked to multiple individuals). 8 - Null. The CPF registration was cancelled due to confirmed fraud or a serious issuance error. Regularization requires in-person verification at a Federal Revenue Service office and may not be guaranteed. 9 - Cancelled Ex Officio. The CPF was cancelled due to administrative or judicial decision, or other registration inconsistencies. | [optional] 
**StatusDescription** | Pointer to **NullableString** | Description of the CPF Status. | [optional] 
**Name** | Pointer to **NullableString** | The full legal name of the individual. | [optional] 
**SocialName** | Pointer to **NullableString** | The social name of the individual in case the individual would like to use a different name than their full legal name. | [optional] 
**DateOfBirth** | Pointer to **NullableString** | The date of birth of the individual. | [optional] 
**YearOfDeath** | Pointer to **NullableString** | Deprecated as of September 1, 2026.              The year of death of the individual, if applicable. | [optional] 
**DateOfRegistration** | Pointer to **NullableString** | The date of registration of the individual with the Federal Revenue Service. | [optional] 
**BiometricVerification** | Pointer to [**NullableBrazilBiometricVerificationOutput**](BrazilBiometricVerificationOutput.md) | The biometric verification results.              This is an optional configurable check. | [optional] 

## Methods

### NewBrazilCpfProviderOutput

`func NewBrazilCpfProviderOutput() *BrazilCpfProviderOutput`

NewBrazilCpfProviderOutput instantiates a new BrazilCpfProviderOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrazilCpfProviderOutputWithDefaults

`func NewBrazilCpfProviderOutputWithDefaults() *BrazilCpfProviderOutput`

NewBrazilCpfProviderOutputWithDefaults instantiates a new BrazilCpfProviderOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPreviewRawErrorCode

`func (o *BrazilCpfProviderOutput) GetPreviewRawErrorCode() string`

GetPreviewRawErrorCode returns the PreviewRawErrorCode field if non-nil, zero value otherwise.

### GetPreviewRawErrorCodeOk

`func (o *BrazilCpfProviderOutput) GetPreviewRawErrorCodeOk() (*string, bool)`

GetPreviewRawErrorCodeOk returns a tuple with the PreviewRawErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewRawErrorCode

`func (o *BrazilCpfProviderOutput) SetPreviewRawErrorCode(v string)`

SetPreviewRawErrorCode sets PreviewRawErrorCode field to given value.

### HasPreviewRawErrorCode

`func (o *BrazilCpfProviderOutput) HasPreviewRawErrorCode() bool`

HasPreviewRawErrorCode returns a boolean if a field has been set.

### SetPreviewRawErrorCodeNil

`func (o *BrazilCpfProviderOutput) SetPreviewRawErrorCodeNil(b bool)`

 SetPreviewRawErrorCodeNil sets the value for PreviewRawErrorCode to be an explicit nil

### UnsetPreviewRawErrorCode
`func (o *BrazilCpfProviderOutput) UnsetPreviewRawErrorCode()`

UnsetPreviewRawErrorCode ensures that no value is present for PreviewRawErrorCode, not even an explicit nil
### GetCpfNumber

`func (o *BrazilCpfProviderOutput) GetCpfNumber() string`

GetCpfNumber returns the CpfNumber field if non-nil, zero value otherwise.

### GetCpfNumberOk

`func (o *BrazilCpfProviderOutput) GetCpfNumberOk() (*string, bool)`

GetCpfNumberOk returns a tuple with the CpfNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpfNumber

`func (o *BrazilCpfProviderOutput) SetCpfNumber(v string)`

SetCpfNumber sets CpfNumber field to given value.

### HasCpfNumber

`func (o *BrazilCpfProviderOutput) HasCpfNumber() bool`

HasCpfNumber returns a boolean if a field has been set.

### SetCpfNumberNil

`func (o *BrazilCpfProviderOutput) SetCpfNumberNil(b bool)`

 SetCpfNumberNil sets the value for CpfNumber to be an explicit nil

### UnsetCpfNumber
`func (o *BrazilCpfProviderOutput) UnsetCpfNumber()`

UnsetCpfNumber ensures that no value is present for CpfNumber, not even an explicit nil
### GetFiscalRegions

`func (o *BrazilCpfProviderOutput) GetFiscalRegions() []string`

GetFiscalRegions returns the FiscalRegions field if non-nil, zero value otherwise.

### GetFiscalRegionsOk

`func (o *BrazilCpfProviderOutput) GetFiscalRegionsOk() (*[]string, bool)`

GetFiscalRegionsOk returns a tuple with the FiscalRegions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiscalRegions

`func (o *BrazilCpfProviderOutput) SetFiscalRegions(v []string)`

SetFiscalRegions sets FiscalRegions field to given value.

### HasFiscalRegions

`func (o *BrazilCpfProviderOutput) HasFiscalRegions() bool`

HasFiscalRegions returns a boolean if a field has been set.

### SetFiscalRegionsNil

`func (o *BrazilCpfProviderOutput) SetFiscalRegionsNil(b bool)`

 SetFiscalRegionsNil sets the value for FiscalRegions to be an explicit nil

### UnsetFiscalRegions
`func (o *BrazilCpfProviderOutput) UnsetFiscalRegions()`

UnsetFiscalRegions ensures that no value is present for FiscalRegions, not even an explicit nil
### GetStatusCode

`func (o *BrazilCpfProviderOutput) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *BrazilCpfProviderOutput) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *BrazilCpfProviderOutput) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *BrazilCpfProviderOutput) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### SetStatusCodeNil

`func (o *BrazilCpfProviderOutput) SetStatusCodeNil(b bool)`

 SetStatusCodeNil sets the value for StatusCode to be an explicit nil

### UnsetStatusCode
`func (o *BrazilCpfProviderOutput) UnsetStatusCode()`

UnsetStatusCode ensures that no value is present for StatusCode, not even an explicit nil
### GetStatusDescription

`func (o *BrazilCpfProviderOutput) GetStatusDescription() string`

GetStatusDescription returns the StatusDescription field if non-nil, zero value otherwise.

### GetStatusDescriptionOk

`func (o *BrazilCpfProviderOutput) GetStatusDescriptionOk() (*string, bool)`

GetStatusDescriptionOk returns a tuple with the StatusDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDescription

`func (o *BrazilCpfProviderOutput) SetStatusDescription(v string)`

SetStatusDescription sets StatusDescription field to given value.

### HasStatusDescription

`func (o *BrazilCpfProviderOutput) HasStatusDescription() bool`

HasStatusDescription returns a boolean if a field has been set.

### SetStatusDescriptionNil

`func (o *BrazilCpfProviderOutput) SetStatusDescriptionNil(b bool)`

 SetStatusDescriptionNil sets the value for StatusDescription to be an explicit nil

### UnsetStatusDescription
`func (o *BrazilCpfProviderOutput) UnsetStatusDescription()`

UnsetStatusDescription ensures that no value is present for StatusDescription, not even an explicit nil
### GetName

`func (o *BrazilCpfProviderOutput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BrazilCpfProviderOutput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BrazilCpfProviderOutput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BrazilCpfProviderOutput) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *BrazilCpfProviderOutput) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BrazilCpfProviderOutput) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSocialName

`func (o *BrazilCpfProviderOutput) GetSocialName() string`

GetSocialName returns the SocialName field if non-nil, zero value otherwise.

### GetSocialNameOk

`func (o *BrazilCpfProviderOutput) GetSocialNameOk() (*string, bool)`

GetSocialNameOk returns a tuple with the SocialName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialName

`func (o *BrazilCpfProviderOutput) SetSocialName(v string)`

SetSocialName sets SocialName field to given value.

### HasSocialName

`func (o *BrazilCpfProviderOutput) HasSocialName() bool`

HasSocialName returns a boolean if a field has been set.

### SetSocialNameNil

`func (o *BrazilCpfProviderOutput) SetSocialNameNil(b bool)`

 SetSocialNameNil sets the value for SocialName to be an explicit nil

### UnsetSocialName
`func (o *BrazilCpfProviderOutput) UnsetSocialName()`

UnsetSocialName ensures that no value is present for SocialName, not even an explicit nil
### GetDateOfBirth

`func (o *BrazilCpfProviderOutput) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *BrazilCpfProviderOutput) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *BrazilCpfProviderOutput) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *BrazilCpfProviderOutput) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### SetDateOfBirthNil

`func (o *BrazilCpfProviderOutput) SetDateOfBirthNil(b bool)`

 SetDateOfBirthNil sets the value for DateOfBirth to be an explicit nil

### UnsetDateOfBirth
`func (o *BrazilCpfProviderOutput) UnsetDateOfBirth()`

UnsetDateOfBirth ensures that no value is present for DateOfBirth, not even an explicit nil
### GetYearOfDeath

`func (o *BrazilCpfProviderOutput) GetYearOfDeath() string`

GetYearOfDeath returns the YearOfDeath field if non-nil, zero value otherwise.

### GetYearOfDeathOk

`func (o *BrazilCpfProviderOutput) GetYearOfDeathOk() (*string, bool)`

GetYearOfDeathOk returns a tuple with the YearOfDeath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYearOfDeath

`func (o *BrazilCpfProviderOutput) SetYearOfDeath(v string)`

SetYearOfDeath sets YearOfDeath field to given value.

### HasYearOfDeath

`func (o *BrazilCpfProviderOutput) HasYearOfDeath() bool`

HasYearOfDeath returns a boolean if a field has been set.

### SetYearOfDeathNil

`func (o *BrazilCpfProviderOutput) SetYearOfDeathNil(b bool)`

 SetYearOfDeathNil sets the value for YearOfDeath to be an explicit nil

### UnsetYearOfDeath
`func (o *BrazilCpfProviderOutput) UnsetYearOfDeath()`

UnsetYearOfDeath ensures that no value is present for YearOfDeath, not even an explicit nil
### GetDateOfRegistration

`func (o *BrazilCpfProviderOutput) GetDateOfRegistration() string`

GetDateOfRegistration returns the DateOfRegistration field if non-nil, zero value otherwise.

### GetDateOfRegistrationOk

`func (o *BrazilCpfProviderOutput) GetDateOfRegistrationOk() (*string, bool)`

GetDateOfRegistrationOk returns a tuple with the DateOfRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfRegistration

`func (o *BrazilCpfProviderOutput) SetDateOfRegistration(v string)`

SetDateOfRegistration sets DateOfRegistration field to given value.

### HasDateOfRegistration

`func (o *BrazilCpfProviderOutput) HasDateOfRegistration() bool`

HasDateOfRegistration returns a boolean if a field has been set.

### SetDateOfRegistrationNil

`func (o *BrazilCpfProviderOutput) SetDateOfRegistrationNil(b bool)`

 SetDateOfRegistrationNil sets the value for DateOfRegistration to be an explicit nil

### UnsetDateOfRegistration
`func (o *BrazilCpfProviderOutput) UnsetDateOfRegistration()`

UnsetDateOfRegistration ensures that no value is present for DateOfRegistration, not even an explicit nil
### GetBiometricVerification

`func (o *BrazilCpfProviderOutput) GetBiometricVerification() BrazilBiometricVerificationOutput`

GetBiometricVerification returns the BiometricVerification field if non-nil, zero value otherwise.

### GetBiometricVerificationOk

`func (o *BrazilCpfProviderOutput) GetBiometricVerificationOk() (*BrazilBiometricVerificationOutput, bool)`

GetBiometricVerificationOk returns a tuple with the BiometricVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiometricVerification

`func (o *BrazilCpfProviderOutput) SetBiometricVerification(v BrazilBiometricVerificationOutput)`

SetBiometricVerification sets BiometricVerification field to given value.

### HasBiometricVerification

`func (o *BrazilCpfProviderOutput) HasBiometricVerification() bool`

HasBiometricVerification returns a boolean if a field has been set.

### SetBiometricVerificationNil

`func (o *BrazilCpfProviderOutput) SetBiometricVerificationNil(b bool)`

 SetBiometricVerificationNil sets the value for BiometricVerification to be an explicit nil

### UnsetBiometricVerification
`func (o *BrazilCpfProviderOutput) UnsetBiometricVerification()`

UnsetBiometricVerification ensures that no value is present for BiometricVerification, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


