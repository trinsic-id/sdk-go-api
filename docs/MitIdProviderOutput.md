# MitIdProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FullName** | Pointer to **NullableString** | The full name of the individual. | [optional] 
**DateOfBirth** | Pointer to **NullableString** | The date of birth of the individual. | [optional] 
**UniqueIdentifier** | Pointer to **NullableString** | A stable unique identifier representing the user in MitID&#39;s system. It is recommended to use this identifier instead of using the civil registration number (CPR number) directly. | [optional] 
**AuthenticationAssuranceLevel** | Pointer to **NullableString** | The authentication assurance level (AAL) for the verification. The National Institute of Standards and Technology (NIST) measures the confidence in a digital identity&#39;s authentication strength set by a set of requirements for different levels. More information can be found here: https://pages.nist.gov/800-63-4/sp800-63b/aal/#AAL_SEC4 Possible values:              -  https://data.gov.dk/concept/core/nsis/Low - (Level 1) Provides basic confidence that the authenticator is either single factor or multifactor authentication. - https://data.gov.dk/concept/core/nsis/Substantial - (Level 2) Provides high confidence that the authenticator uses two distinct authentication factors that use secure authentication protocols. - https://data.gov.dk/concept/core/nsis/High - (Level 3) Provides very high confidence that the authenticator uses an non-exportable private key and another authentication factor.              **Note**: These urls are not valid anymore, but they still contain the level of assurance name. | [optional] 
**IdentityAuthenticationLevel** | Pointer to **NullableString** | The identity assurance level (IAL) for the verification. The National Institute of Standards and Technology (NIST) measures the confidence of the digital identity&#39;s verification set by a set of requirements for different levels. More information can be found here: https://pages.nist.gov/800-63-4/sp800-63b/aal/#AAL_SEC4 Possible values:              -  https://data.gov.dk/concept/core/nsis/Low - (Level 1) The user has self asserted their identity and is neither verified nor validated. - https://data.gov.dk/concept/core/nsis/Substantial - (Level 2) The user has performed either a remote or in-person identity proofing. - https://data.gov.dk/concept/core/nsis/High - (Level 3) The user has performed an in person identity proofing with an authorized representative.              **Note**: These urls are not valid anymore, but they still contain the level of assurance name. | [optional] 
**LevelOfAssurance** | Pointer to **NullableString** | The level of assurance (LOA) for the verification. This is an older retired model that is a combination of Identity Assurance Level and Authentication Assurance Level. The National Institute of Standards and Technology (NIST) measures the confidence of the digital identity&#39;s verification and authentication strength by a set of requirements for different levels.              Possible values: -  https://data.gov.dk/concept/core/nsis/Low - (Level 1) The user has self asserted their identity and multifactor authentication is not required. - https://data.gov.dk/concept/core/nsis/Substantial - (Level 3) The user has performed either a remote or in-person identity proofing and multifactor authentication is required. - https://data.gov.dk/concept/core/nsis/High - (Level 4) The user has performed an in person identity proofing with an authorized representative and has strong cryptographic authentication requirements.              **Note**: These urls are not valid anymore, but they still contain the level of assurance name. MitID only returns three of the four potential levels and does not include level 2.              More information can be found here: - https://pages.nist.gov/800-63-3/sp800-63-3.html - https://pages.nist.gov/800-63-3/ | [optional] 
**IdentificationSource** | Pointer to **NullableString** | The source of the identification for the verification              Possible values: - private - The verification is from a private individual. - professional - The verification is a user in behalf of a organization. | [optional] 
**CivilRegistrationNumber** | Pointer to **NullableString** | The civil registration number (CPR number) of the individual for the verification.              The CPR number is a 10-digit number in the format DDMMYYY-XXXX, where: - DDMMYYY represents the individual&#39;s date of birth (day, month, year). - XXXX is a unique identifier.              If XXXX is even, the individual is female. If XXXX is odd, the individual is male.              More information can be found here: https://international.kk.dk/live/cpr-registration-and-documents/cpr-number | [optional] 
**OrganizationName** | Pointer to **NullableString** | The name of the organization for the verification | [optional] 
**OrganizationNumber** | Pointer to **NullableString** | The organization number (CVR number). This is an 8-digit unique identifier for the organization. This will be prefixed with \&quot;DK\&quot; for VAT numbers. | [optional] 
**AuthorizedRepresentativeNumber** | Pointer to **NullableString** | The organization number (CVR number) of the organization the user is authorized to represent. This is only returned if the requested scope is for an organization verification and the user provides a private source in behalf of a company. | [optional] 

## Methods

### NewMitIdProviderOutput

`func NewMitIdProviderOutput() *MitIdProviderOutput`

NewMitIdProviderOutput instantiates a new MitIdProviderOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMitIdProviderOutputWithDefaults

`func NewMitIdProviderOutputWithDefaults() *MitIdProviderOutput`

NewMitIdProviderOutputWithDefaults instantiates a new MitIdProviderOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFullName

`func (o *MitIdProviderOutput) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *MitIdProviderOutput) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *MitIdProviderOutput) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *MitIdProviderOutput) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### SetFullNameNil

`func (o *MitIdProviderOutput) SetFullNameNil(b bool)`

 SetFullNameNil sets the value for FullName to be an explicit nil

### UnsetFullName
`func (o *MitIdProviderOutput) UnsetFullName()`

UnsetFullName ensures that no value is present for FullName, not even an explicit nil
### GetDateOfBirth

`func (o *MitIdProviderOutput) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *MitIdProviderOutput) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *MitIdProviderOutput) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *MitIdProviderOutput) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### SetDateOfBirthNil

`func (o *MitIdProviderOutput) SetDateOfBirthNil(b bool)`

 SetDateOfBirthNil sets the value for DateOfBirth to be an explicit nil

### UnsetDateOfBirth
`func (o *MitIdProviderOutput) UnsetDateOfBirth()`

UnsetDateOfBirth ensures that no value is present for DateOfBirth, not even an explicit nil
### GetUniqueIdentifier

`func (o *MitIdProviderOutput) GetUniqueIdentifier() string`

GetUniqueIdentifier returns the UniqueIdentifier field if non-nil, zero value otherwise.

### GetUniqueIdentifierOk

`func (o *MitIdProviderOutput) GetUniqueIdentifierOk() (*string, bool)`

GetUniqueIdentifierOk returns a tuple with the UniqueIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueIdentifier

`func (o *MitIdProviderOutput) SetUniqueIdentifier(v string)`

SetUniqueIdentifier sets UniqueIdentifier field to given value.

### HasUniqueIdentifier

`func (o *MitIdProviderOutput) HasUniqueIdentifier() bool`

HasUniqueIdentifier returns a boolean if a field has been set.

### SetUniqueIdentifierNil

`func (o *MitIdProviderOutput) SetUniqueIdentifierNil(b bool)`

 SetUniqueIdentifierNil sets the value for UniqueIdentifier to be an explicit nil

### UnsetUniqueIdentifier
`func (o *MitIdProviderOutput) UnsetUniqueIdentifier()`

UnsetUniqueIdentifier ensures that no value is present for UniqueIdentifier, not even an explicit nil
### GetAuthenticationAssuranceLevel

`func (o *MitIdProviderOutput) GetAuthenticationAssuranceLevel() string`

GetAuthenticationAssuranceLevel returns the AuthenticationAssuranceLevel field if non-nil, zero value otherwise.

### GetAuthenticationAssuranceLevelOk

`func (o *MitIdProviderOutput) GetAuthenticationAssuranceLevelOk() (*string, bool)`

GetAuthenticationAssuranceLevelOk returns a tuple with the AuthenticationAssuranceLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationAssuranceLevel

`func (o *MitIdProviderOutput) SetAuthenticationAssuranceLevel(v string)`

SetAuthenticationAssuranceLevel sets AuthenticationAssuranceLevel field to given value.

### HasAuthenticationAssuranceLevel

`func (o *MitIdProviderOutput) HasAuthenticationAssuranceLevel() bool`

HasAuthenticationAssuranceLevel returns a boolean if a field has been set.

### SetAuthenticationAssuranceLevelNil

`func (o *MitIdProviderOutput) SetAuthenticationAssuranceLevelNil(b bool)`

 SetAuthenticationAssuranceLevelNil sets the value for AuthenticationAssuranceLevel to be an explicit nil

### UnsetAuthenticationAssuranceLevel
`func (o *MitIdProviderOutput) UnsetAuthenticationAssuranceLevel()`

UnsetAuthenticationAssuranceLevel ensures that no value is present for AuthenticationAssuranceLevel, not even an explicit nil
### GetIdentityAuthenticationLevel

`func (o *MitIdProviderOutput) GetIdentityAuthenticationLevel() string`

GetIdentityAuthenticationLevel returns the IdentityAuthenticationLevel field if non-nil, zero value otherwise.

### GetIdentityAuthenticationLevelOk

`func (o *MitIdProviderOutput) GetIdentityAuthenticationLevelOk() (*string, bool)`

GetIdentityAuthenticationLevelOk returns a tuple with the IdentityAuthenticationLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityAuthenticationLevel

`func (o *MitIdProviderOutput) SetIdentityAuthenticationLevel(v string)`

SetIdentityAuthenticationLevel sets IdentityAuthenticationLevel field to given value.

### HasIdentityAuthenticationLevel

`func (o *MitIdProviderOutput) HasIdentityAuthenticationLevel() bool`

HasIdentityAuthenticationLevel returns a boolean if a field has been set.

### SetIdentityAuthenticationLevelNil

`func (o *MitIdProviderOutput) SetIdentityAuthenticationLevelNil(b bool)`

 SetIdentityAuthenticationLevelNil sets the value for IdentityAuthenticationLevel to be an explicit nil

### UnsetIdentityAuthenticationLevel
`func (o *MitIdProviderOutput) UnsetIdentityAuthenticationLevel()`

UnsetIdentityAuthenticationLevel ensures that no value is present for IdentityAuthenticationLevel, not even an explicit nil
### GetLevelOfAssurance

`func (o *MitIdProviderOutput) GetLevelOfAssurance() string`

GetLevelOfAssurance returns the LevelOfAssurance field if non-nil, zero value otherwise.

### GetLevelOfAssuranceOk

`func (o *MitIdProviderOutput) GetLevelOfAssuranceOk() (*string, bool)`

GetLevelOfAssuranceOk returns a tuple with the LevelOfAssurance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelOfAssurance

`func (o *MitIdProviderOutput) SetLevelOfAssurance(v string)`

SetLevelOfAssurance sets LevelOfAssurance field to given value.

### HasLevelOfAssurance

`func (o *MitIdProviderOutput) HasLevelOfAssurance() bool`

HasLevelOfAssurance returns a boolean if a field has been set.

### SetLevelOfAssuranceNil

`func (o *MitIdProviderOutput) SetLevelOfAssuranceNil(b bool)`

 SetLevelOfAssuranceNil sets the value for LevelOfAssurance to be an explicit nil

### UnsetLevelOfAssurance
`func (o *MitIdProviderOutput) UnsetLevelOfAssurance()`

UnsetLevelOfAssurance ensures that no value is present for LevelOfAssurance, not even an explicit nil
### GetIdentificationSource

`func (o *MitIdProviderOutput) GetIdentificationSource() string`

GetIdentificationSource returns the IdentificationSource field if non-nil, zero value otherwise.

### GetIdentificationSourceOk

`func (o *MitIdProviderOutput) GetIdentificationSourceOk() (*string, bool)`

GetIdentificationSourceOk returns a tuple with the IdentificationSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentificationSource

`func (o *MitIdProviderOutput) SetIdentificationSource(v string)`

SetIdentificationSource sets IdentificationSource field to given value.

### HasIdentificationSource

`func (o *MitIdProviderOutput) HasIdentificationSource() bool`

HasIdentificationSource returns a boolean if a field has been set.

### SetIdentificationSourceNil

`func (o *MitIdProviderOutput) SetIdentificationSourceNil(b bool)`

 SetIdentificationSourceNil sets the value for IdentificationSource to be an explicit nil

### UnsetIdentificationSource
`func (o *MitIdProviderOutput) UnsetIdentificationSource()`

UnsetIdentificationSource ensures that no value is present for IdentificationSource, not even an explicit nil
### GetCivilRegistrationNumber

`func (o *MitIdProviderOutput) GetCivilRegistrationNumber() string`

GetCivilRegistrationNumber returns the CivilRegistrationNumber field if non-nil, zero value otherwise.

### GetCivilRegistrationNumberOk

`func (o *MitIdProviderOutput) GetCivilRegistrationNumberOk() (*string, bool)`

GetCivilRegistrationNumberOk returns a tuple with the CivilRegistrationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCivilRegistrationNumber

`func (o *MitIdProviderOutput) SetCivilRegistrationNumber(v string)`

SetCivilRegistrationNumber sets CivilRegistrationNumber field to given value.

### HasCivilRegistrationNumber

`func (o *MitIdProviderOutput) HasCivilRegistrationNumber() bool`

HasCivilRegistrationNumber returns a boolean if a field has been set.

### SetCivilRegistrationNumberNil

`func (o *MitIdProviderOutput) SetCivilRegistrationNumberNil(b bool)`

 SetCivilRegistrationNumberNil sets the value for CivilRegistrationNumber to be an explicit nil

### UnsetCivilRegistrationNumber
`func (o *MitIdProviderOutput) UnsetCivilRegistrationNumber()`

UnsetCivilRegistrationNumber ensures that no value is present for CivilRegistrationNumber, not even an explicit nil
### GetOrganizationName

`func (o *MitIdProviderOutput) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *MitIdProviderOutput) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *MitIdProviderOutput) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.

### HasOrganizationName

`func (o *MitIdProviderOutput) HasOrganizationName() bool`

HasOrganizationName returns a boolean if a field has been set.

### SetOrganizationNameNil

`func (o *MitIdProviderOutput) SetOrganizationNameNil(b bool)`

 SetOrganizationNameNil sets the value for OrganizationName to be an explicit nil

### UnsetOrganizationName
`func (o *MitIdProviderOutput) UnsetOrganizationName()`

UnsetOrganizationName ensures that no value is present for OrganizationName, not even an explicit nil
### GetOrganizationNumber

`func (o *MitIdProviderOutput) GetOrganizationNumber() string`

GetOrganizationNumber returns the OrganizationNumber field if non-nil, zero value otherwise.

### GetOrganizationNumberOk

`func (o *MitIdProviderOutput) GetOrganizationNumberOk() (*string, bool)`

GetOrganizationNumberOk returns a tuple with the OrganizationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationNumber

`func (o *MitIdProviderOutput) SetOrganizationNumber(v string)`

SetOrganizationNumber sets OrganizationNumber field to given value.

### HasOrganizationNumber

`func (o *MitIdProviderOutput) HasOrganizationNumber() bool`

HasOrganizationNumber returns a boolean if a field has been set.

### SetOrganizationNumberNil

`func (o *MitIdProviderOutput) SetOrganizationNumberNil(b bool)`

 SetOrganizationNumberNil sets the value for OrganizationNumber to be an explicit nil

### UnsetOrganizationNumber
`func (o *MitIdProviderOutput) UnsetOrganizationNumber()`

UnsetOrganizationNumber ensures that no value is present for OrganizationNumber, not even an explicit nil
### GetAuthorizedRepresentativeNumber

`func (o *MitIdProviderOutput) GetAuthorizedRepresentativeNumber() string`

GetAuthorizedRepresentativeNumber returns the AuthorizedRepresentativeNumber field if non-nil, zero value otherwise.

### GetAuthorizedRepresentativeNumberOk

`func (o *MitIdProviderOutput) GetAuthorizedRepresentativeNumberOk() (*string, bool)`

GetAuthorizedRepresentativeNumberOk returns a tuple with the AuthorizedRepresentativeNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedRepresentativeNumber

`func (o *MitIdProviderOutput) SetAuthorizedRepresentativeNumber(v string)`

SetAuthorizedRepresentativeNumber sets AuthorizedRepresentativeNumber field to given value.

### HasAuthorizedRepresentativeNumber

`func (o *MitIdProviderOutput) HasAuthorizedRepresentativeNumber() bool`

HasAuthorizedRepresentativeNumber returns a boolean if a field has been set.

### SetAuthorizedRepresentativeNumberNil

`func (o *MitIdProviderOutput) SetAuthorizedRepresentativeNumberNil(b bool)`

 SetAuthorizedRepresentativeNumberNil sets the value for AuthorizedRepresentativeNumber to be an explicit nil

### UnsetAuthorizedRepresentativeNumber
`func (o *MitIdProviderOutput) UnsetAuthorizedRepresentativeNumber()`

UnsetAuthorizedRepresentativeNumber ensures that no value is present for AuthorizedRepresentativeNumber, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


