# FinnishTrustNetworkProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FullName** | Pointer to **NullableString** | The full name of the individual. | [optional] 
**DateOfBirth** | Pointer to **NullableString** | The date of birth of the individual. | [optional] 
**PersonalIdentificationCode** | Pointer to **NullableString** | The 11-digit Finnish Personal Identification Code (Henkilötunnus) of the verified individual.              This is in the format DDMMYYCZZZQ, where: - DDMMYY is the date of birth - C is a symbol which determines the century of birth - ZZZ is an individual number, indicating gender - Q is a checksum character              If ZZZ is even, the individual is female. If ZZZ is odd, the individual is male.              If C is &#39;+&#39;, the individual was born in the 19th century (1800-1899). If C is &#39;-&#39;, &#39;U&#39;, &#39;V&#39;, &#39;W&#39;, &#39;X&#39;, or &#39;Y&#39;, the individual was born in the 20th century (1900-1999). If C is &#39;A&#39;, &#39;B&#39;, &#39;C&#39;, &#39;D&#39;, &#39;E&#39;, or &#39;F&#39;, the individual was born in the 21st century (2000-2099). | [optional] 
**UniqueIdentificationNumber** | Pointer to **NullableString** | The 9-digit Finnish Unique Identification Number (FINUID, or sähköinen asiointitunnus SATU) of the verified individual. This number is typically used for online transactions and unlike the Finnish Personal Identification Code, does not reveal personal information such as birthdate.              The first 8 digits are randomly generated and the last character is a check control. | [optional] 
**GivenName** | Pointer to **NullableString** | The given name of the individual. | [optional] 
**FamilyName** | Pointer to **NullableString** | The family name of the individual. | [optional] 
**LevelOfAssurance** | Pointer to **NullableString** | The level of assurance (LOA) for the verification.              The LOA refers to the degree of confidence in the claimed identity of a person. The European Digital Identity Framework (EUDI) measures the confidence of the digital identity&#39;s verification and authentication strength by a set of requirements for different levels. To learn more, see: https://ec.europa.eu/digital-building-blocks/sites/spaces/DIGITAL/pages/467110081/eIDAS+Levels+of+Assurance              Possible values: - Low: The individual has self asserted their identity and multifactor authentication is not required. - Substantial: The individual has performed either a remote or in-person identity verification and multifactor authentication is required. - High: The individual has performed an in-person identity proofing with an authorized representative and has strong cryptographic authentication requirements such as using a smart card. | [optional] 
**Bank** | Pointer to **NullableString** | The bank used to perform the identification for the verified individual. | [optional] 
**PhoneNumber** | Pointer to **NullableString** | The phone number of the verified individual.              This is only returned if the individual authenticated with MobileID (Mobiilivarmenne). | [optional] 
**OrganizationName** | Pointer to **NullableString** | The organization name that the individual is associated with.              This is returned for requests that contain the organization scope. | [optional] 
**VatNumber** | Pointer to **NullableString** | The VAT number of the organization that the individual is associated with.              This is returned for requests that contain the organization scope. | [optional] 

## Methods

### NewFinnishTrustNetworkProviderOutput

`func NewFinnishTrustNetworkProviderOutput() *FinnishTrustNetworkProviderOutput`

NewFinnishTrustNetworkProviderOutput instantiates a new FinnishTrustNetworkProviderOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFinnishTrustNetworkProviderOutputWithDefaults

`func NewFinnishTrustNetworkProviderOutputWithDefaults() *FinnishTrustNetworkProviderOutput`

NewFinnishTrustNetworkProviderOutputWithDefaults instantiates a new FinnishTrustNetworkProviderOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFullName

`func (o *FinnishTrustNetworkProviderOutput) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *FinnishTrustNetworkProviderOutput) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *FinnishTrustNetworkProviderOutput) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *FinnishTrustNetworkProviderOutput) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### SetFullNameNil

`func (o *FinnishTrustNetworkProviderOutput) SetFullNameNil(b bool)`

 SetFullNameNil sets the value for FullName to be an explicit nil

### UnsetFullName
`func (o *FinnishTrustNetworkProviderOutput) UnsetFullName()`

UnsetFullName ensures that no value is present for FullName, not even an explicit nil
### GetDateOfBirth

`func (o *FinnishTrustNetworkProviderOutput) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *FinnishTrustNetworkProviderOutput) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *FinnishTrustNetworkProviderOutput) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *FinnishTrustNetworkProviderOutput) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### SetDateOfBirthNil

`func (o *FinnishTrustNetworkProviderOutput) SetDateOfBirthNil(b bool)`

 SetDateOfBirthNil sets the value for DateOfBirth to be an explicit nil

### UnsetDateOfBirth
`func (o *FinnishTrustNetworkProviderOutput) UnsetDateOfBirth()`

UnsetDateOfBirth ensures that no value is present for DateOfBirth, not even an explicit nil
### GetPersonalIdentificationCode

`func (o *FinnishTrustNetworkProviderOutput) GetPersonalIdentificationCode() string`

GetPersonalIdentificationCode returns the PersonalIdentificationCode field if non-nil, zero value otherwise.

### GetPersonalIdentificationCodeOk

`func (o *FinnishTrustNetworkProviderOutput) GetPersonalIdentificationCodeOk() (*string, bool)`

GetPersonalIdentificationCodeOk returns a tuple with the PersonalIdentificationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalIdentificationCode

`func (o *FinnishTrustNetworkProviderOutput) SetPersonalIdentificationCode(v string)`

SetPersonalIdentificationCode sets PersonalIdentificationCode field to given value.

### HasPersonalIdentificationCode

`func (o *FinnishTrustNetworkProviderOutput) HasPersonalIdentificationCode() bool`

HasPersonalIdentificationCode returns a boolean if a field has been set.

### SetPersonalIdentificationCodeNil

`func (o *FinnishTrustNetworkProviderOutput) SetPersonalIdentificationCodeNil(b bool)`

 SetPersonalIdentificationCodeNil sets the value for PersonalIdentificationCode to be an explicit nil

### UnsetPersonalIdentificationCode
`func (o *FinnishTrustNetworkProviderOutput) UnsetPersonalIdentificationCode()`

UnsetPersonalIdentificationCode ensures that no value is present for PersonalIdentificationCode, not even an explicit nil
### GetUniqueIdentificationNumber

`func (o *FinnishTrustNetworkProviderOutput) GetUniqueIdentificationNumber() string`

GetUniqueIdentificationNumber returns the UniqueIdentificationNumber field if non-nil, zero value otherwise.

### GetUniqueIdentificationNumberOk

`func (o *FinnishTrustNetworkProviderOutput) GetUniqueIdentificationNumberOk() (*string, bool)`

GetUniqueIdentificationNumberOk returns a tuple with the UniqueIdentificationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueIdentificationNumber

`func (o *FinnishTrustNetworkProviderOutput) SetUniqueIdentificationNumber(v string)`

SetUniqueIdentificationNumber sets UniqueIdentificationNumber field to given value.

### HasUniqueIdentificationNumber

`func (o *FinnishTrustNetworkProviderOutput) HasUniqueIdentificationNumber() bool`

HasUniqueIdentificationNumber returns a boolean if a field has been set.

### SetUniqueIdentificationNumberNil

`func (o *FinnishTrustNetworkProviderOutput) SetUniqueIdentificationNumberNil(b bool)`

 SetUniqueIdentificationNumberNil sets the value for UniqueIdentificationNumber to be an explicit nil

### UnsetUniqueIdentificationNumber
`func (o *FinnishTrustNetworkProviderOutput) UnsetUniqueIdentificationNumber()`

UnsetUniqueIdentificationNumber ensures that no value is present for UniqueIdentificationNumber, not even an explicit nil
### GetGivenName

`func (o *FinnishTrustNetworkProviderOutput) GetGivenName() string`

GetGivenName returns the GivenName field if non-nil, zero value otherwise.

### GetGivenNameOk

`func (o *FinnishTrustNetworkProviderOutput) GetGivenNameOk() (*string, bool)`

GetGivenNameOk returns a tuple with the GivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenName

`func (o *FinnishTrustNetworkProviderOutput) SetGivenName(v string)`

SetGivenName sets GivenName field to given value.

### HasGivenName

`func (o *FinnishTrustNetworkProviderOutput) HasGivenName() bool`

HasGivenName returns a boolean if a field has been set.

### SetGivenNameNil

`func (o *FinnishTrustNetworkProviderOutput) SetGivenNameNil(b bool)`

 SetGivenNameNil sets the value for GivenName to be an explicit nil

### UnsetGivenName
`func (o *FinnishTrustNetworkProviderOutput) UnsetGivenName()`

UnsetGivenName ensures that no value is present for GivenName, not even an explicit nil
### GetFamilyName

`func (o *FinnishTrustNetworkProviderOutput) GetFamilyName() string`

GetFamilyName returns the FamilyName field if non-nil, zero value otherwise.

### GetFamilyNameOk

`func (o *FinnishTrustNetworkProviderOutput) GetFamilyNameOk() (*string, bool)`

GetFamilyNameOk returns a tuple with the FamilyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyName

`func (o *FinnishTrustNetworkProviderOutput) SetFamilyName(v string)`

SetFamilyName sets FamilyName field to given value.

### HasFamilyName

`func (o *FinnishTrustNetworkProviderOutput) HasFamilyName() bool`

HasFamilyName returns a boolean if a field has been set.

### SetFamilyNameNil

`func (o *FinnishTrustNetworkProviderOutput) SetFamilyNameNil(b bool)`

 SetFamilyNameNil sets the value for FamilyName to be an explicit nil

### UnsetFamilyName
`func (o *FinnishTrustNetworkProviderOutput) UnsetFamilyName()`

UnsetFamilyName ensures that no value is present for FamilyName, not even an explicit nil
### GetLevelOfAssurance

`func (o *FinnishTrustNetworkProviderOutput) GetLevelOfAssurance() string`

GetLevelOfAssurance returns the LevelOfAssurance field if non-nil, zero value otherwise.

### GetLevelOfAssuranceOk

`func (o *FinnishTrustNetworkProviderOutput) GetLevelOfAssuranceOk() (*string, bool)`

GetLevelOfAssuranceOk returns a tuple with the LevelOfAssurance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelOfAssurance

`func (o *FinnishTrustNetworkProviderOutput) SetLevelOfAssurance(v string)`

SetLevelOfAssurance sets LevelOfAssurance field to given value.

### HasLevelOfAssurance

`func (o *FinnishTrustNetworkProviderOutput) HasLevelOfAssurance() bool`

HasLevelOfAssurance returns a boolean if a field has been set.

### SetLevelOfAssuranceNil

`func (o *FinnishTrustNetworkProviderOutput) SetLevelOfAssuranceNil(b bool)`

 SetLevelOfAssuranceNil sets the value for LevelOfAssurance to be an explicit nil

### UnsetLevelOfAssurance
`func (o *FinnishTrustNetworkProviderOutput) UnsetLevelOfAssurance()`

UnsetLevelOfAssurance ensures that no value is present for LevelOfAssurance, not even an explicit nil
### GetBank

`func (o *FinnishTrustNetworkProviderOutput) GetBank() string`

GetBank returns the Bank field if non-nil, zero value otherwise.

### GetBankOk

`func (o *FinnishTrustNetworkProviderOutput) GetBankOk() (*string, bool)`

GetBankOk returns a tuple with the Bank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBank

`func (o *FinnishTrustNetworkProviderOutput) SetBank(v string)`

SetBank sets Bank field to given value.

### HasBank

`func (o *FinnishTrustNetworkProviderOutput) HasBank() bool`

HasBank returns a boolean if a field has been set.

### SetBankNil

`func (o *FinnishTrustNetworkProviderOutput) SetBankNil(b bool)`

 SetBankNil sets the value for Bank to be an explicit nil

### UnsetBank
`func (o *FinnishTrustNetworkProviderOutput) UnsetBank()`

UnsetBank ensures that no value is present for Bank, not even an explicit nil
### GetPhoneNumber

`func (o *FinnishTrustNetworkProviderOutput) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *FinnishTrustNetworkProviderOutput) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *FinnishTrustNetworkProviderOutput) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *FinnishTrustNetworkProviderOutput) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### SetPhoneNumberNil

`func (o *FinnishTrustNetworkProviderOutput) SetPhoneNumberNil(b bool)`

 SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil

### UnsetPhoneNumber
`func (o *FinnishTrustNetworkProviderOutput) UnsetPhoneNumber()`

UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
### GetOrganizationName

`func (o *FinnishTrustNetworkProviderOutput) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *FinnishTrustNetworkProviderOutput) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *FinnishTrustNetworkProviderOutput) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.

### HasOrganizationName

`func (o *FinnishTrustNetworkProviderOutput) HasOrganizationName() bool`

HasOrganizationName returns a boolean if a field has been set.

### SetOrganizationNameNil

`func (o *FinnishTrustNetworkProviderOutput) SetOrganizationNameNil(b bool)`

 SetOrganizationNameNil sets the value for OrganizationName to be an explicit nil

### UnsetOrganizationName
`func (o *FinnishTrustNetworkProviderOutput) UnsetOrganizationName()`

UnsetOrganizationName ensures that no value is present for OrganizationName, not even an explicit nil
### GetVatNumber

`func (o *FinnishTrustNetworkProviderOutput) GetVatNumber() string`

GetVatNumber returns the VatNumber field if non-nil, zero value otherwise.

### GetVatNumberOk

`func (o *FinnishTrustNetworkProviderOutput) GetVatNumberOk() (*string, bool)`

GetVatNumberOk returns a tuple with the VatNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatNumber

`func (o *FinnishTrustNetworkProviderOutput) SetVatNumber(v string)`

SetVatNumber sets VatNumber field to given value.

### HasVatNumber

`func (o *FinnishTrustNetworkProviderOutput) HasVatNumber() bool`

HasVatNumber returns a boolean if a field has been set.

### SetVatNumberNil

`func (o *FinnishTrustNetworkProviderOutput) SetVatNumberNil(b bool)`

 SetVatNumberNil sets the value for VatNumber to be an explicit nil

### UnsetVatNumber
`func (o *FinnishTrustNetworkProviderOutput) UnsetVatNumber()`

UnsetVatNumber ensures that no value is present for VatNumber, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


