# NorwegianBankIdProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FullName** | Pointer to **NullableString** | The full name of the individual. | [optional] 
**DateOfBirth** | Pointer to **NullableString** | The date of birth of the individual. | [optional] 
**PersonalIdentifier** | Pointer to **NullableString** | The personal identifier for Norwegian BankID.              This uniquely identifies the individual in the Norwegian BankID system and is considered a stable identifier to use. | [optional] 
**NationalIdentityNumber** | Pointer to **NullableString** | The 11-digit Norwegian National Identity Number (fødselsnummer) of the verified individual.              This is in the format DDMMYYZZZCC, where: - DDMMYY is the date of birth (In some cases, this is not the date of birth due to no available identity numbers for some dates - ZZZ is an individual number, indicating gender - CC is a checksum character              If ZZZ is even, the individual is female. If ZZZ is odd, the individual is male. | [optional] 
**GivenName** | Pointer to **NullableString** | The given name of the individual. | [optional] 
**FamilyName** | Pointer to **NullableString** | The family name of the individual. | [optional] 
**LevelOfAssurance** | Pointer to **NullableString** | The level of assurance (LOA) for the verification.              The LOA refers to the degree of confidence in the claimed identity of a person. The European Digital Identity Framework (EUDI) measures the confidence of the digital identity&#39;s verification and authentication strength by a set of requirements for different levels. To learn more, see: https://ec.europa.eu/digital-building-blocks/sites/spaces/DIGITAL/pages/467110081/eIDAS+Levels+of+Assurance              Possible values: - Low: The individual has self asserted their identity and multifactor authentication is not required. - Substantial: The individual has performed either a remote or in-person identity verification and multifactor authentication is required. - High: The individual has performed an in-person identity proofing with an authorized representative and has strong cryptographic authentication requirements such as using a smart card. | [optional] 
**AuthenticationMethod** | Pointer to **NullableString** | The authentication method used by the individual.              Possible values: - urn:bankid:bis - BankID with Biometrics - urn:bankid:bid - Standard BankID with High Assurance | [optional] 

## Methods

### NewNorwegianBankIdProviderOutput

`func NewNorwegianBankIdProviderOutput() *NorwegianBankIdProviderOutput`

NewNorwegianBankIdProviderOutput instantiates a new NorwegianBankIdProviderOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNorwegianBankIdProviderOutputWithDefaults

`func NewNorwegianBankIdProviderOutputWithDefaults() *NorwegianBankIdProviderOutput`

NewNorwegianBankIdProviderOutputWithDefaults instantiates a new NorwegianBankIdProviderOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFullName

`func (o *NorwegianBankIdProviderOutput) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *NorwegianBankIdProviderOutput) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *NorwegianBankIdProviderOutput) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *NorwegianBankIdProviderOutput) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### SetFullNameNil

`func (o *NorwegianBankIdProviderOutput) SetFullNameNil(b bool)`

 SetFullNameNil sets the value for FullName to be an explicit nil

### UnsetFullName
`func (o *NorwegianBankIdProviderOutput) UnsetFullName()`

UnsetFullName ensures that no value is present for FullName, not even an explicit nil
### GetDateOfBirth

`func (o *NorwegianBankIdProviderOutput) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *NorwegianBankIdProviderOutput) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *NorwegianBankIdProviderOutput) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *NorwegianBankIdProviderOutput) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### SetDateOfBirthNil

`func (o *NorwegianBankIdProviderOutput) SetDateOfBirthNil(b bool)`

 SetDateOfBirthNil sets the value for DateOfBirth to be an explicit nil

### UnsetDateOfBirth
`func (o *NorwegianBankIdProviderOutput) UnsetDateOfBirth()`

UnsetDateOfBirth ensures that no value is present for DateOfBirth, not even an explicit nil
### GetPersonalIdentifier

`func (o *NorwegianBankIdProviderOutput) GetPersonalIdentifier() string`

GetPersonalIdentifier returns the PersonalIdentifier field if non-nil, zero value otherwise.

### GetPersonalIdentifierOk

`func (o *NorwegianBankIdProviderOutput) GetPersonalIdentifierOk() (*string, bool)`

GetPersonalIdentifierOk returns a tuple with the PersonalIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalIdentifier

`func (o *NorwegianBankIdProviderOutput) SetPersonalIdentifier(v string)`

SetPersonalIdentifier sets PersonalIdentifier field to given value.

### HasPersonalIdentifier

`func (o *NorwegianBankIdProviderOutput) HasPersonalIdentifier() bool`

HasPersonalIdentifier returns a boolean if a field has been set.

### SetPersonalIdentifierNil

`func (o *NorwegianBankIdProviderOutput) SetPersonalIdentifierNil(b bool)`

 SetPersonalIdentifierNil sets the value for PersonalIdentifier to be an explicit nil

### UnsetPersonalIdentifier
`func (o *NorwegianBankIdProviderOutput) UnsetPersonalIdentifier()`

UnsetPersonalIdentifier ensures that no value is present for PersonalIdentifier, not even an explicit nil
### GetNationalIdentityNumber

`func (o *NorwegianBankIdProviderOutput) GetNationalIdentityNumber() string`

GetNationalIdentityNumber returns the NationalIdentityNumber field if non-nil, zero value otherwise.

### GetNationalIdentityNumberOk

`func (o *NorwegianBankIdProviderOutput) GetNationalIdentityNumberOk() (*string, bool)`

GetNationalIdentityNumberOk returns a tuple with the NationalIdentityNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationalIdentityNumber

`func (o *NorwegianBankIdProviderOutput) SetNationalIdentityNumber(v string)`

SetNationalIdentityNumber sets NationalIdentityNumber field to given value.

### HasNationalIdentityNumber

`func (o *NorwegianBankIdProviderOutput) HasNationalIdentityNumber() bool`

HasNationalIdentityNumber returns a boolean if a field has been set.

### SetNationalIdentityNumberNil

`func (o *NorwegianBankIdProviderOutput) SetNationalIdentityNumberNil(b bool)`

 SetNationalIdentityNumberNil sets the value for NationalIdentityNumber to be an explicit nil

### UnsetNationalIdentityNumber
`func (o *NorwegianBankIdProviderOutput) UnsetNationalIdentityNumber()`

UnsetNationalIdentityNumber ensures that no value is present for NationalIdentityNumber, not even an explicit nil
### GetGivenName

`func (o *NorwegianBankIdProviderOutput) GetGivenName() string`

GetGivenName returns the GivenName field if non-nil, zero value otherwise.

### GetGivenNameOk

`func (o *NorwegianBankIdProviderOutput) GetGivenNameOk() (*string, bool)`

GetGivenNameOk returns a tuple with the GivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenName

`func (o *NorwegianBankIdProviderOutput) SetGivenName(v string)`

SetGivenName sets GivenName field to given value.

### HasGivenName

`func (o *NorwegianBankIdProviderOutput) HasGivenName() bool`

HasGivenName returns a boolean if a field has been set.

### SetGivenNameNil

`func (o *NorwegianBankIdProviderOutput) SetGivenNameNil(b bool)`

 SetGivenNameNil sets the value for GivenName to be an explicit nil

### UnsetGivenName
`func (o *NorwegianBankIdProviderOutput) UnsetGivenName()`

UnsetGivenName ensures that no value is present for GivenName, not even an explicit nil
### GetFamilyName

`func (o *NorwegianBankIdProviderOutput) GetFamilyName() string`

GetFamilyName returns the FamilyName field if non-nil, zero value otherwise.

### GetFamilyNameOk

`func (o *NorwegianBankIdProviderOutput) GetFamilyNameOk() (*string, bool)`

GetFamilyNameOk returns a tuple with the FamilyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyName

`func (o *NorwegianBankIdProviderOutput) SetFamilyName(v string)`

SetFamilyName sets FamilyName field to given value.

### HasFamilyName

`func (o *NorwegianBankIdProviderOutput) HasFamilyName() bool`

HasFamilyName returns a boolean if a field has been set.

### SetFamilyNameNil

`func (o *NorwegianBankIdProviderOutput) SetFamilyNameNil(b bool)`

 SetFamilyNameNil sets the value for FamilyName to be an explicit nil

### UnsetFamilyName
`func (o *NorwegianBankIdProviderOutput) UnsetFamilyName()`

UnsetFamilyName ensures that no value is present for FamilyName, not even an explicit nil
### GetLevelOfAssurance

`func (o *NorwegianBankIdProviderOutput) GetLevelOfAssurance() string`

GetLevelOfAssurance returns the LevelOfAssurance field if non-nil, zero value otherwise.

### GetLevelOfAssuranceOk

`func (o *NorwegianBankIdProviderOutput) GetLevelOfAssuranceOk() (*string, bool)`

GetLevelOfAssuranceOk returns a tuple with the LevelOfAssurance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelOfAssurance

`func (o *NorwegianBankIdProviderOutput) SetLevelOfAssurance(v string)`

SetLevelOfAssurance sets LevelOfAssurance field to given value.

### HasLevelOfAssurance

`func (o *NorwegianBankIdProviderOutput) HasLevelOfAssurance() bool`

HasLevelOfAssurance returns a boolean if a field has been set.

### SetLevelOfAssuranceNil

`func (o *NorwegianBankIdProviderOutput) SetLevelOfAssuranceNil(b bool)`

 SetLevelOfAssuranceNil sets the value for LevelOfAssurance to be an explicit nil

### UnsetLevelOfAssurance
`func (o *NorwegianBankIdProviderOutput) UnsetLevelOfAssurance()`

UnsetLevelOfAssurance ensures that no value is present for LevelOfAssurance, not even an explicit nil
### GetAuthenticationMethod

`func (o *NorwegianBankIdProviderOutput) GetAuthenticationMethod() string`

GetAuthenticationMethod returns the AuthenticationMethod field if non-nil, zero value otherwise.

### GetAuthenticationMethodOk

`func (o *NorwegianBankIdProviderOutput) GetAuthenticationMethodOk() (*string, bool)`

GetAuthenticationMethodOk returns a tuple with the AuthenticationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationMethod

`func (o *NorwegianBankIdProviderOutput) SetAuthenticationMethod(v string)`

SetAuthenticationMethod sets AuthenticationMethod field to given value.

### HasAuthenticationMethod

`func (o *NorwegianBankIdProviderOutput) HasAuthenticationMethod() bool`

HasAuthenticationMethod returns a boolean if a field has been set.

### SetAuthenticationMethodNil

`func (o *NorwegianBankIdProviderOutput) SetAuthenticationMethodNil(b bool)`

 SetAuthenticationMethodNil sets the value for AuthenticationMethod to be an explicit nil

### UnsetAuthenticationMethod
`func (o *NorwegianBankIdProviderOutput) UnsetAuthenticationMethod()`

UnsetAuthenticationMethod ensures that no value is present for AuthenticationMethod, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


