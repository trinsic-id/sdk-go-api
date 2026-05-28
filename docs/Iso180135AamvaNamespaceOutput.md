# Iso180135AamvaNamespaceOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomesticDrivingPrivileges** | Pointer to [**[]Iso180135AamvaDomesticDrivingPrivilege**](Iso180135AamvaDomesticDrivingPrivilege.md) | Domestic categories of vehicles, restrictions, and conditions, parsed per AAMVA Section 7.2.4. | [optional] 
**NameSuffix** | Pointer to **NullableString** | Name suffix of the individual that has been issued the credential. | [optional] 
**OrganDonor** | Pointer to **NullableBool** | Whether the individual is an organ donor. | [optional] 
**Veteran** | Pointer to **NullableBool** | Whether the individual is a veteran. | [optional] 
**FamilyNameTruncation** | Pointer to **NullableString** | Indicates whether the family name has been truncated.              Possible values: - &#39;T&#39;: Truncated - &#39;N&#39;: Not truncated - &#39;U&#39;: Unknown whether truncated | [optional] 
**GivenNameTruncation** | Pointer to **NullableString** | Indicates whether the given name has been truncated.              Possible values: - &#39;T&#39;: Truncated - &#39;N&#39;: Not truncated - &#39;U&#39;: Unknown whether truncated | [optional] 
**AkaFamilyNameV2** | Pointer to **NullableString** | Other family name by which the individual is known. | [optional] 
**AkaGivenNameV2** | Pointer to **NullableString** | Other given name by which the individual is known. | [optional] 
**AkaSuffix** | Pointer to **NullableString** | Other suffix by which the individual is known. | [optional] 
**WeightRange** | Pointer to [**NullableIso180135AamvaWeightRange**](Iso180135AamvaWeightRange.md) | Approximate weight range of the individual, in kilograms. | [optional] 
**RaceEthnicity** | Pointer to **NullableString** | AAMVA D20 race or ethnicity code of the individual. | [optional] 
**Sex** | Pointer to **NullableInt32** | The sex of the individual as an AAMVA-defined sex code.              This is distinct from an ISO/IEC 5218 sex code in two ways: there is no \&quot;Unknown\&quot; value, and \&quot;Not Applicable\&quot; is replaced with \&quot;Not Specified\&quot;.              Possible values: - 1: Male - 2: Female - 9: Not Specified | [optional] 
**FirstName** | Pointer to **NullableString** | First name of the individual. | [optional] 
**MiddleNames** | Pointer to **NullableString** | Middle name or names of the individual. | [optional] 
**FirstNameTruncation** | Pointer to **NullableString** | Whether the first name has been truncated. | [optional] 
**MiddleNamesTruncation** | Pointer to **NullableString** | Whether the middle name has been truncated. | [optional] 
**EdlCredential** | Pointer to **NullableInt32** | Deprecated AAMVA EDL (Enhanced Driver&#39;s License) credential indicator.              If present, indicates the type of the EDL credential.              Possible values: 1: Driver&#39;s License 2: Identification Card | [optional] 
**EdlCredentialV2** | Pointer to **NullableBool** | Whether the credential is an Enhanced Driver&#39;s License (EDL). | [optional] 
**DhsCompliance** | Pointer to **NullableBool** | Whether the credential is REAL ID compliant. &#x60;true&#x60; for fully compliant (\&quot;F\&quot;), &#x60;false&#x60; for non-compliant (\&quot;N\&quot;), &#x60;null&#x60; when not present. | [optional] 
**ResidentCounty** | Pointer to **NullableString** | Deprecated county code for the county where the individual lives. | [optional] 
**ResidentCountyV2** | Pointer to **NullableString** | County code for the county where the individual lives. | [optional] 
**HazmatEndorsementExpirationDate** | Pointer to **NullableString** | Date on which the hazardous material endorsement expires. | [optional] 
**CdlIndicator** | Pointer to **NullableBool** | Whether the credential is a Commercial Driver&#39;s License (CDL) per FMCSA. | [optional] 
**CdlNonDomiciled** | Pointer to **NullableBool** | Deprecated. Whether the CDL holder is non-domiciled in the issuing jurisdiction. | [optional] 
**CdlNonDomiciledV2** | Pointer to **NullableBool** | Whether the CDL holder is non-domiciled in the issuing jurisdiction. | [optional] 
**DhsComplianceText** | Pointer to **NullableString** | Text agreed on between the issuing authority and DHS for non-compliant credentials. | [optional] 
**DhsTemporaryLawfulStatus** | Pointer to **NullableBool** | Whether the individual has DHS temporary lawful status. | [optional] 

## Methods

### NewIso180135AamvaNamespaceOutput

`func NewIso180135AamvaNamespaceOutput() *Iso180135AamvaNamespaceOutput`

NewIso180135AamvaNamespaceOutput instantiates a new Iso180135AamvaNamespaceOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIso180135AamvaNamespaceOutputWithDefaults

`func NewIso180135AamvaNamespaceOutputWithDefaults() *Iso180135AamvaNamespaceOutput`

NewIso180135AamvaNamespaceOutputWithDefaults instantiates a new Iso180135AamvaNamespaceOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomesticDrivingPrivileges

`func (o *Iso180135AamvaNamespaceOutput) GetDomesticDrivingPrivileges() []Iso180135AamvaDomesticDrivingPrivilege`

GetDomesticDrivingPrivileges returns the DomesticDrivingPrivileges field if non-nil, zero value otherwise.

### GetDomesticDrivingPrivilegesOk

`func (o *Iso180135AamvaNamespaceOutput) GetDomesticDrivingPrivilegesOk() (*[]Iso180135AamvaDomesticDrivingPrivilege, bool)`

GetDomesticDrivingPrivilegesOk returns a tuple with the DomesticDrivingPrivileges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomesticDrivingPrivileges

`func (o *Iso180135AamvaNamespaceOutput) SetDomesticDrivingPrivileges(v []Iso180135AamvaDomesticDrivingPrivilege)`

SetDomesticDrivingPrivileges sets DomesticDrivingPrivileges field to given value.

### HasDomesticDrivingPrivileges

`func (o *Iso180135AamvaNamespaceOutput) HasDomesticDrivingPrivileges() bool`

HasDomesticDrivingPrivileges returns a boolean if a field has been set.

### SetDomesticDrivingPrivilegesNil

`func (o *Iso180135AamvaNamespaceOutput) SetDomesticDrivingPrivilegesNil(b bool)`

 SetDomesticDrivingPrivilegesNil sets the value for DomesticDrivingPrivileges to be an explicit nil

### UnsetDomesticDrivingPrivileges
`func (o *Iso180135AamvaNamespaceOutput) UnsetDomesticDrivingPrivileges()`

UnsetDomesticDrivingPrivileges ensures that no value is present for DomesticDrivingPrivileges, not even an explicit nil
### GetNameSuffix

`func (o *Iso180135AamvaNamespaceOutput) GetNameSuffix() string`

GetNameSuffix returns the NameSuffix field if non-nil, zero value otherwise.

### GetNameSuffixOk

`func (o *Iso180135AamvaNamespaceOutput) GetNameSuffixOk() (*string, bool)`

GetNameSuffixOk returns a tuple with the NameSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameSuffix

`func (o *Iso180135AamvaNamespaceOutput) SetNameSuffix(v string)`

SetNameSuffix sets NameSuffix field to given value.

### HasNameSuffix

`func (o *Iso180135AamvaNamespaceOutput) HasNameSuffix() bool`

HasNameSuffix returns a boolean if a field has been set.

### SetNameSuffixNil

`func (o *Iso180135AamvaNamespaceOutput) SetNameSuffixNil(b bool)`

 SetNameSuffixNil sets the value for NameSuffix to be an explicit nil

### UnsetNameSuffix
`func (o *Iso180135AamvaNamespaceOutput) UnsetNameSuffix()`

UnsetNameSuffix ensures that no value is present for NameSuffix, not even an explicit nil
### GetOrganDonor

`func (o *Iso180135AamvaNamespaceOutput) GetOrganDonor() bool`

GetOrganDonor returns the OrganDonor field if non-nil, zero value otherwise.

### GetOrganDonorOk

`func (o *Iso180135AamvaNamespaceOutput) GetOrganDonorOk() (*bool, bool)`

GetOrganDonorOk returns a tuple with the OrganDonor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganDonor

`func (o *Iso180135AamvaNamespaceOutput) SetOrganDonor(v bool)`

SetOrganDonor sets OrganDonor field to given value.

### HasOrganDonor

`func (o *Iso180135AamvaNamespaceOutput) HasOrganDonor() bool`

HasOrganDonor returns a boolean if a field has been set.

### SetOrganDonorNil

`func (o *Iso180135AamvaNamespaceOutput) SetOrganDonorNil(b bool)`

 SetOrganDonorNil sets the value for OrganDonor to be an explicit nil

### UnsetOrganDonor
`func (o *Iso180135AamvaNamespaceOutput) UnsetOrganDonor()`

UnsetOrganDonor ensures that no value is present for OrganDonor, not even an explicit nil
### GetVeteran

`func (o *Iso180135AamvaNamespaceOutput) GetVeteran() bool`

GetVeteran returns the Veteran field if non-nil, zero value otherwise.

### GetVeteranOk

`func (o *Iso180135AamvaNamespaceOutput) GetVeteranOk() (*bool, bool)`

GetVeteranOk returns a tuple with the Veteran field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVeteran

`func (o *Iso180135AamvaNamespaceOutput) SetVeteran(v bool)`

SetVeteran sets Veteran field to given value.

### HasVeteran

`func (o *Iso180135AamvaNamespaceOutput) HasVeteran() bool`

HasVeteran returns a boolean if a field has been set.

### SetVeteranNil

`func (o *Iso180135AamvaNamespaceOutput) SetVeteranNil(b bool)`

 SetVeteranNil sets the value for Veteran to be an explicit nil

### UnsetVeteran
`func (o *Iso180135AamvaNamespaceOutput) UnsetVeteran()`

UnsetVeteran ensures that no value is present for Veteran, not even an explicit nil
### GetFamilyNameTruncation

`func (o *Iso180135AamvaNamespaceOutput) GetFamilyNameTruncation() string`

GetFamilyNameTruncation returns the FamilyNameTruncation field if non-nil, zero value otherwise.

### GetFamilyNameTruncationOk

`func (o *Iso180135AamvaNamespaceOutput) GetFamilyNameTruncationOk() (*string, bool)`

GetFamilyNameTruncationOk returns a tuple with the FamilyNameTruncation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyNameTruncation

`func (o *Iso180135AamvaNamespaceOutput) SetFamilyNameTruncation(v string)`

SetFamilyNameTruncation sets FamilyNameTruncation field to given value.

### HasFamilyNameTruncation

`func (o *Iso180135AamvaNamespaceOutput) HasFamilyNameTruncation() bool`

HasFamilyNameTruncation returns a boolean if a field has been set.

### SetFamilyNameTruncationNil

`func (o *Iso180135AamvaNamespaceOutput) SetFamilyNameTruncationNil(b bool)`

 SetFamilyNameTruncationNil sets the value for FamilyNameTruncation to be an explicit nil

### UnsetFamilyNameTruncation
`func (o *Iso180135AamvaNamespaceOutput) UnsetFamilyNameTruncation()`

UnsetFamilyNameTruncation ensures that no value is present for FamilyNameTruncation, not even an explicit nil
### GetGivenNameTruncation

`func (o *Iso180135AamvaNamespaceOutput) GetGivenNameTruncation() string`

GetGivenNameTruncation returns the GivenNameTruncation field if non-nil, zero value otherwise.

### GetGivenNameTruncationOk

`func (o *Iso180135AamvaNamespaceOutput) GetGivenNameTruncationOk() (*string, bool)`

GetGivenNameTruncationOk returns a tuple with the GivenNameTruncation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenNameTruncation

`func (o *Iso180135AamvaNamespaceOutput) SetGivenNameTruncation(v string)`

SetGivenNameTruncation sets GivenNameTruncation field to given value.

### HasGivenNameTruncation

`func (o *Iso180135AamvaNamespaceOutput) HasGivenNameTruncation() bool`

HasGivenNameTruncation returns a boolean if a field has been set.

### SetGivenNameTruncationNil

`func (o *Iso180135AamvaNamespaceOutput) SetGivenNameTruncationNil(b bool)`

 SetGivenNameTruncationNil sets the value for GivenNameTruncation to be an explicit nil

### UnsetGivenNameTruncation
`func (o *Iso180135AamvaNamespaceOutput) UnsetGivenNameTruncation()`

UnsetGivenNameTruncation ensures that no value is present for GivenNameTruncation, not even an explicit nil
### GetAkaFamilyNameV2

`func (o *Iso180135AamvaNamespaceOutput) GetAkaFamilyNameV2() string`

GetAkaFamilyNameV2 returns the AkaFamilyNameV2 field if non-nil, zero value otherwise.

### GetAkaFamilyNameV2Ok

`func (o *Iso180135AamvaNamespaceOutput) GetAkaFamilyNameV2Ok() (*string, bool)`

GetAkaFamilyNameV2Ok returns a tuple with the AkaFamilyNameV2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAkaFamilyNameV2

`func (o *Iso180135AamvaNamespaceOutput) SetAkaFamilyNameV2(v string)`

SetAkaFamilyNameV2 sets AkaFamilyNameV2 field to given value.

### HasAkaFamilyNameV2

`func (o *Iso180135AamvaNamespaceOutput) HasAkaFamilyNameV2() bool`

HasAkaFamilyNameV2 returns a boolean if a field has been set.

### SetAkaFamilyNameV2Nil

`func (o *Iso180135AamvaNamespaceOutput) SetAkaFamilyNameV2Nil(b bool)`

 SetAkaFamilyNameV2Nil sets the value for AkaFamilyNameV2 to be an explicit nil

### UnsetAkaFamilyNameV2
`func (o *Iso180135AamvaNamespaceOutput) UnsetAkaFamilyNameV2()`

UnsetAkaFamilyNameV2 ensures that no value is present for AkaFamilyNameV2, not even an explicit nil
### GetAkaGivenNameV2

`func (o *Iso180135AamvaNamespaceOutput) GetAkaGivenNameV2() string`

GetAkaGivenNameV2 returns the AkaGivenNameV2 field if non-nil, zero value otherwise.

### GetAkaGivenNameV2Ok

`func (o *Iso180135AamvaNamespaceOutput) GetAkaGivenNameV2Ok() (*string, bool)`

GetAkaGivenNameV2Ok returns a tuple with the AkaGivenNameV2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAkaGivenNameV2

`func (o *Iso180135AamvaNamespaceOutput) SetAkaGivenNameV2(v string)`

SetAkaGivenNameV2 sets AkaGivenNameV2 field to given value.

### HasAkaGivenNameV2

`func (o *Iso180135AamvaNamespaceOutput) HasAkaGivenNameV2() bool`

HasAkaGivenNameV2 returns a boolean if a field has been set.

### SetAkaGivenNameV2Nil

`func (o *Iso180135AamvaNamespaceOutput) SetAkaGivenNameV2Nil(b bool)`

 SetAkaGivenNameV2Nil sets the value for AkaGivenNameV2 to be an explicit nil

### UnsetAkaGivenNameV2
`func (o *Iso180135AamvaNamespaceOutput) UnsetAkaGivenNameV2()`

UnsetAkaGivenNameV2 ensures that no value is present for AkaGivenNameV2, not even an explicit nil
### GetAkaSuffix

`func (o *Iso180135AamvaNamespaceOutput) GetAkaSuffix() string`

GetAkaSuffix returns the AkaSuffix field if non-nil, zero value otherwise.

### GetAkaSuffixOk

`func (o *Iso180135AamvaNamespaceOutput) GetAkaSuffixOk() (*string, bool)`

GetAkaSuffixOk returns a tuple with the AkaSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAkaSuffix

`func (o *Iso180135AamvaNamespaceOutput) SetAkaSuffix(v string)`

SetAkaSuffix sets AkaSuffix field to given value.

### HasAkaSuffix

`func (o *Iso180135AamvaNamespaceOutput) HasAkaSuffix() bool`

HasAkaSuffix returns a boolean if a field has been set.

### SetAkaSuffixNil

`func (o *Iso180135AamvaNamespaceOutput) SetAkaSuffixNil(b bool)`

 SetAkaSuffixNil sets the value for AkaSuffix to be an explicit nil

### UnsetAkaSuffix
`func (o *Iso180135AamvaNamespaceOutput) UnsetAkaSuffix()`

UnsetAkaSuffix ensures that no value is present for AkaSuffix, not even an explicit nil
### GetWeightRange

`func (o *Iso180135AamvaNamespaceOutput) GetWeightRange() Iso180135AamvaWeightRange`

GetWeightRange returns the WeightRange field if non-nil, zero value otherwise.

### GetWeightRangeOk

`func (o *Iso180135AamvaNamespaceOutput) GetWeightRangeOk() (*Iso180135AamvaWeightRange, bool)`

GetWeightRangeOk returns a tuple with the WeightRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightRange

`func (o *Iso180135AamvaNamespaceOutput) SetWeightRange(v Iso180135AamvaWeightRange)`

SetWeightRange sets WeightRange field to given value.

### HasWeightRange

`func (o *Iso180135AamvaNamespaceOutput) HasWeightRange() bool`

HasWeightRange returns a boolean if a field has been set.

### SetWeightRangeNil

`func (o *Iso180135AamvaNamespaceOutput) SetWeightRangeNil(b bool)`

 SetWeightRangeNil sets the value for WeightRange to be an explicit nil

### UnsetWeightRange
`func (o *Iso180135AamvaNamespaceOutput) UnsetWeightRange()`

UnsetWeightRange ensures that no value is present for WeightRange, not even an explicit nil
### GetRaceEthnicity

`func (o *Iso180135AamvaNamespaceOutput) GetRaceEthnicity() string`

GetRaceEthnicity returns the RaceEthnicity field if non-nil, zero value otherwise.

### GetRaceEthnicityOk

`func (o *Iso180135AamvaNamespaceOutput) GetRaceEthnicityOk() (*string, bool)`

GetRaceEthnicityOk returns a tuple with the RaceEthnicity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaceEthnicity

`func (o *Iso180135AamvaNamespaceOutput) SetRaceEthnicity(v string)`

SetRaceEthnicity sets RaceEthnicity field to given value.

### HasRaceEthnicity

`func (o *Iso180135AamvaNamespaceOutput) HasRaceEthnicity() bool`

HasRaceEthnicity returns a boolean if a field has been set.

### SetRaceEthnicityNil

`func (o *Iso180135AamvaNamespaceOutput) SetRaceEthnicityNil(b bool)`

 SetRaceEthnicityNil sets the value for RaceEthnicity to be an explicit nil

### UnsetRaceEthnicity
`func (o *Iso180135AamvaNamespaceOutput) UnsetRaceEthnicity()`

UnsetRaceEthnicity ensures that no value is present for RaceEthnicity, not even an explicit nil
### GetSex

`func (o *Iso180135AamvaNamespaceOutput) GetSex() int32`

GetSex returns the Sex field if non-nil, zero value otherwise.

### GetSexOk

`func (o *Iso180135AamvaNamespaceOutput) GetSexOk() (*int32, bool)`

GetSexOk returns a tuple with the Sex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSex

`func (o *Iso180135AamvaNamespaceOutput) SetSex(v int32)`

SetSex sets Sex field to given value.

### HasSex

`func (o *Iso180135AamvaNamespaceOutput) HasSex() bool`

HasSex returns a boolean if a field has been set.

### SetSexNil

`func (o *Iso180135AamvaNamespaceOutput) SetSexNil(b bool)`

 SetSexNil sets the value for Sex to be an explicit nil

### UnsetSex
`func (o *Iso180135AamvaNamespaceOutput) UnsetSex()`

UnsetSex ensures that no value is present for Sex, not even an explicit nil
### GetFirstName

`func (o *Iso180135AamvaNamespaceOutput) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *Iso180135AamvaNamespaceOutput) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *Iso180135AamvaNamespaceOutput) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *Iso180135AamvaNamespaceOutput) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *Iso180135AamvaNamespaceOutput) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *Iso180135AamvaNamespaceOutput) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetMiddleNames

`func (o *Iso180135AamvaNamespaceOutput) GetMiddleNames() string`

GetMiddleNames returns the MiddleNames field if non-nil, zero value otherwise.

### GetMiddleNamesOk

`func (o *Iso180135AamvaNamespaceOutput) GetMiddleNamesOk() (*string, bool)`

GetMiddleNamesOk returns a tuple with the MiddleNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleNames

`func (o *Iso180135AamvaNamespaceOutput) SetMiddleNames(v string)`

SetMiddleNames sets MiddleNames field to given value.

### HasMiddleNames

`func (o *Iso180135AamvaNamespaceOutput) HasMiddleNames() bool`

HasMiddleNames returns a boolean if a field has been set.

### SetMiddleNamesNil

`func (o *Iso180135AamvaNamespaceOutput) SetMiddleNamesNil(b bool)`

 SetMiddleNamesNil sets the value for MiddleNames to be an explicit nil

### UnsetMiddleNames
`func (o *Iso180135AamvaNamespaceOutput) UnsetMiddleNames()`

UnsetMiddleNames ensures that no value is present for MiddleNames, not even an explicit nil
### GetFirstNameTruncation

`func (o *Iso180135AamvaNamespaceOutput) GetFirstNameTruncation() string`

GetFirstNameTruncation returns the FirstNameTruncation field if non-nil, zero value otherwise.

### GetFirstNameTruncationOk

`func (o *Iso180135AamvaNamespaceOutput) GetFirstNameTruncationOk() (*string, bool)`

GetFirstNameTruncationOk returns a tuple with the FirstNameTruncation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstNameTruncation

`func (o *Iso180135AamvaNamespaceOutput) SetFirstNameTruncation(v string)`

SetFirstNameTruncation sets FirstNameTruncation field to given value.

### HasFirstNameTruncation

`func (o *Iso180135AamvaNamespaceOutput) HasFirstNameTruncation() bool`

HasFirstNameTruncation returns a boolean if a field has been set.

### SetFirstNameTruncationNil

`func (o *Iso180135AamvaNamespaceOutput) SetFirstNameTruncationNil(b bool)`

 SetFirstNameTruncationNil sets the value for FirstNameTruncation to be an explicit nil

### UnsetFirstNameTruncation
`func (o *Iso180135AamvaNamespaceOutput) UnsetFirstNameTruncation()`

UnsetFirstNameTruncation ensures that no value is present for FirstNameTruncation, not even an explicit nil
### GetMiddleNamesTruncation

`func (o *Iso180135AamvaNamespaceOutput) GetMiddleNamesTruncation() string`

GetMiddleNamesTruncation returns the MiddleNamesTruncation field if non-nil, zero value otherwise.

### GetMiddleNamesTruncationOk

`func (o *Iso180135AamvaNamespaceOutput) GetMiddleNamesTruncationOk() (*string, bool)`

GetMiddleNamesTruncationOk returns a tuple with the MiddleNamesTruncation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleNamesTruncation

`func (o *Iso180135AamvaNamespaceOutput) SetMiddleNamesTruncation(v string)`

SetMiddleNamesTruncation sets MiddleNamesTruncation field to given value.

### HasMiddleNamesTruncation

`func (o *Iso180135AamvaNamespaceOutput) HasMiddleNamesTruncation() bool`

HasMiddleNamesTruncation returns a boolean if a field has been set.

### SetMiddleNamesTruncationNil

`func (o *Iso180135AamvaNamespaceOutput) SetMiddleNamesTruncationNil(b bool)`

 SetMiddleNamesTruncationNil sets the value for MiddleNamesTruncation to be an explicit nil

### UnsetMiddleNamesTruncation
`func (o *Iso180135AamvaNamespaceOutput) UnsetMiddleNamesTruncation()`

UnsetMiddleNamesTruncation ensures that no value is present for MiddleNamesTruncation, not even an explicit nil
### GetEdlCredential

`func (o *Iso180135AamvaNamespaceOutput) GetEdlCredential() int32`

GetEdlCredential returns the EdlCredential field if non-nil, zero value otherwise.

### GetEdlCredentialOk

`func (o *Iso180135AamvaNamespaceOutput) GetEdlCredentialOk() (*int32, bool)`

GetEdlCredentialOk returns a tuple with the EdlCredential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdlCredential

`func (o *Iso180135AamvaNamespaceOutput) SetEdlCredential(v int32)`

SetEdlCredential sets EdlCredential field to given value.

### HasEdlCredential

`func (o *Iso180135AamvaNamespaceOutput) HasEdlCredential() bool`

HasEdlCredential returns a boolean if a field has been set.

### SetEdlCredentialNil

`func (o *Iso180135AamvaNamespaceOutput) SetEdlCredentialNil(b bool)`

 SetEdlCredentialNil sets the value for EdlCredential to be an explicit nil

### UnsetEdlCredential
`func (o *Iso180135AamvaNamespaceOutput) UnsetEdlCredential()`

UnsetEdlCredential ensures that no value is present for EdlCredential, not even an explicit nil
### GetEdlCredentialV2

`func (o *Iso180135AamvaNamespaceOutput) GetEdlCredentialV2() bool`

GetEdlCredentialV2 returns the EdlCredentialV2 field if non-nil, zero value otherwise.

### GetEdlCredentialV2Ok

`func (o *Iso180135AamvaNamespaceOutput) GetEdlCredentialV2Ok() (*bool, bool)`

GetEdlCredentialV2Ok returns a tuple with the EdlCredentialV2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdlCredentialV2

`func (o *Iso180135AamvaNamespaceOutput) SetEdlCredentialV2(v bool)`

SetEdlCredentialV2 sets EdlCredentialV2 field to given value.

### HasEdlCredentialV2

`func (o *Iso180135AamvaNamespaceOutput) HasEdlCredentialV2() bool`

HasEdlCredentialV2 returns a boolean if a field has been set.

### SetEdlCredentialV2Nil

`func (o *Iso180135AamvaNamespaceOutput) SetEdlCredentialV2Nil(b bool)`

 SetEdlCredentialV2Nil sets the value for EdlCredentialV2 to be an explicit nil

### UnsetEdlCredentialV2
`func (o *Iso180135AamvaNamespaceOutput) UnsetEdlCredentialV2()`

UnsetEdlCredentialV2 ensures that no value is present for EdlCredentialV2, not even an explicit nil
### GetDhsCompliance

`func (o *Iso180135AamvaNamespaceOutput) GetDhsCompliance() bool`

GetDhsCompliance returns the DhsCompliance field if non-nil, zero value otherwise.

### GetDhsComplianceOk

`func (o *Iso180135AamvaNamespaceOutput) GetDhsComplianceOk() (*bool, bool)`

GetDhsComplianceOk returns a tuple with the DhsCompliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhsCompliance

`func (o *Iso180135AamvaNamespaceOutput) SetDhsCompliance(v bool)`

SetDhsCompliance sets DhsCompliance field to given value.

### HasDhsCompliance

`func (o *Iso180135AamvaNamespaceOutput) HasDhsCompliance() bool`

HasDhsCompliance returns a boolean if a field has been set.

### SetDhsComplianceNil

`func (o *Iso180135AamvaNamespaceOutput) SetDhsComplianceNil(b bool)`

 SetDhsComplianceNil sets the value for DhsCompliance to be an explicit nil

### UnsetDhsCompliance
`func (o *Iso180135AamvaNamespaceOutput) UnsetDhsCompliance()`

UnsetDhsCompliance ensures that no value is present for DhsCompliance, not even an explicit nil
### GetResidentCounty

`func (o *Iso180135AamvaNamespaceOutput) GetResidentCounty() string`

GetResidentCounty returns the ResidentCounty field if non-nil, zero value otherwise.

### GetResidentCountyOk

`func (o *Iso180135AamvaNamespaceOutput) GetResidentCountyOk() (*string, bool)`

GetResidentCountyOk returns a tuple with the ResidentCounty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidentCounty

`func (o *Iso180135AamvaNamespaceOutput) SetResidentCounty(v string)`

SetResidentCounty sets ResidentCounty field to given value.

### HasResidentCounty

`func (o *Iso180135AamvaNamespaceOutput) HasResidentCounty() bool`

HasResidentCounty returns a boolean if a field has been set.

### SetResidentCountyNil

`func (o *Iso180135AamvaNamespaceOutput) SetResidentCountyNil(b bool)`

 SetResidentCountyNil sets the value for ResidentCounty to be an explicit nil

### UnsetResidentCounty
`func (o *Iso180135AamvaNamespaceOutput) UnsetResidentCounty()`

UnsetResidentCounty ensures that no value is present for ResidentCounty, not even an explicit nil
### GetResidentCountyV2

`func (o *Iso180135AamvaNamespaceOutput) GetResidentCountyV2() string`

GetResidentCountyV2 returns the ResidentCountyV2 field if non-nil, zero value otherwise.

### GetResidentCountyV2Ok

`func (o *Iso180135AamvaNamespaceOutput) GetResidentCountyV2Ok() (*string, bool)`

GetResidentCountyV2Ok returns a tuple with the ResidentCountyV2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidentCountyV2

`func (o *Iso180135AamvaNamespaceOutput) SetResidentCountyV2(v string)`

SetResidentCountyV2 sets ResidentCountyV2 field to given value.

### HasResidentCountyV2

`func (o *Iso180135AamvaNamespaceOutput) HasResidentCountyV2() bool`

HasResidentCountyV2 returns a boolean if a field has been set.

### SetResidentCountyV2Nil

`func (o *Iso180135AamvaNamespaceOutput) SetResidentCountyV2Nil(b bool)`

 SetResidentCountyV2Nil sets the value for ResidentCountyV2 to be an explicit nil

### UnsetResidentCountyV2
`func (o *Iso180135AamvaNamespaceOutput) UnsetResidentCountyV2()`

UnsetResidentCountyV2 ensures that no value is present for ResidentCountyV2, not even an explicit nil
### GetHazmatEndorsementExpirationDate

`func (o *Iso180135AamvaNamespaceOutput) GetHazmatEndorsementExpirationDate() string`

GetHazmatEndorsementExpirationDate returns the HazmatEndorsementExpirationDate field if non-nil, zero value otherwise.

### GetHazmatEndorsementExpirationDateOk

`func (o *Iso180135AamvaNamespaceOutput) GetHazmatEndorsementExpirationDateOk() (*string, bool)`

GetHazmatEndorsementExpirationDateOk returns a tuple with the HazmatEndorsementExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHazmatEndorsementExpirationDate

`func (o *Iso180135AamvaNamespaceOutput) SetHazmatEndorsementExpirationDate(v string)`

SetHazmatEndorsementExpirationDate sets HazmatEndorsementExpirationDate field to given value.

### HasHazmatEndorsementExpirationDate

`func (o *Iso180135AamvaNamespaceOutput) HasHazmatEndorsementExpirationDate() bool`

HasHazmatEndorsementExpirationDate returns a boolean if a field has been set.

### SetHazmatEndorsementExpirationDateNil

`func (o *Iso180135AamvaNamespaceOutput) SetHazmatEndorsementExpirationDateNil(b bool)`

 SetHazmatEndorsementExpirationDateNil sets the value for HazmatEndorsementExpirationDate to be an explicit nil

### UnsetHazmatEndorsementExpirationDate
`func (o *Iso180135AamvaNamespaceOutput) UnsetHazmatEndorsementExpirationDate()`

UnsetHazmatEndorsementExpirationDate ensures that no value is present for HazmatEndorsementExpirationDate, not even an explicit nil
### GetCdlIndicator

`func (o *Iso180135AamvaNamespaceOutput) GetCdlIndicator() bool`

GetCdlIndicator returns the CdlIndicator field if non-nil, zero value otherwise.

### GetCdlIndicatorOk

`func (o *Iso180135AamvaNamespaceOutput) GetCdlIndicatorOk() (*bool, bool)`

GetCdlIndicatorOk returns a tuple with the CdlIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdlIndicator

`func (o *Iso180135AamvaNamespaceOutput) SetCdlIndicator(v bool)`

SetCdlIndicator sets CdlIndicator field to given value.

### HasCdlIndicator

`func (o *Iso180135AamvaNamespaceOutput) HasCdlIndicator() bool`

HasCdlIndicator returns a boolean if a field has been set.

### SetCdlIndicatorNil

`func (o *Iso180135AamvaNamespaceOutput) SetCdlIndicatorNil(b bool)`

 SetCdlIndicatorNil sets the value for CdlIndicator to be an explicit nil

### UnsetCdlIndicator
`func (o *Iso180135AamvaNamespaceOutput) UnsetCdlIndicator()`

UnsetCdlIndicator ensures that no value is present for CdlIndicator, not even an explicit nil
### GetCdlNonDomiciled

`func (o *Iso180135AamvaNamespaceOutput) GetCdlNonDomiciled() bool`

GetCdlNonDomiciled returns the CdlNonDomiciled field if non-nil, zero value otherwise.

### GetCdlNonDomiciledOk

`func (o *Iso180135AamvaNamespaceOutput) GetCdlNonDomiciledOk() (*bool, bool)`

GetCdlNonDomiciledOk returns a tuple with the CdlNonDomiciled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdlNonDomiciled

`func (o *Iso180135AamvaNamespaceOutput) SetCdlNonDomiciled(v bool)`

SetCdlNonDomiciled sets CdlNonDomiciled field to given value.

### HasCdlNonDomiciled

`func (o *Iso180135AamvaNamespaceOutput) HasCdlNonDomiciled() bool`

HasCdlNonDomiciled returns a boolean if a field has been set.

### SetCdlNonDomiciledNil

`func (o *Iso180135AamvaNamespaceOutput) SetCdlNonDomiciledNil(b bool)`

 SetCdlNonDomiciledNil sets the value for CdlNonDomiciled to be an explicit nil

### UnsetCdlNonDomiciled
`func (o *Iso180135AamvaNamespaceOutput) UnsetCdlNonDomiciled()`

UnsetCdlNonDomiciled ensures that no value is present for CdlNonDomiciled, not even an explicit nil
### GetCdlNonDomiciledV2

`func (o *Iso180135AamvaNamespaceOutput) GetCdlNonDomiciledV2() bool`

GetCdlNonDomiciledV2 returns the CdlNonDomiciledV2 field if non-nil, zero value otherwise.

### GetCdlNonDomiciledV2Ok

`func (o *Iso180135AamvaNamespaceOutput) GetCdlNonDomiciledV2Ok() (*bool, bool)`

GetCdlNonDomiciledV2Ok returns a tuple with the CdlNonDomiciledV2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdlNonDomiciledV2

`func (o *Iso180135AamvaNamespaceOutput) SetCdlNonDomiciledV2(v bool)`

SetCdlNonDomiciledV2 sets CdlNonDomiciledV2 field to given value.

### HasCdlNonDomiciledV2

`func (o *Iso180135AamvaNamespaceOutput) HasCdlNonDomiciledV2() bool`

HasCdlNonDomiciledV2 returns a boolean if a field has been set.

### SetCdlNonDomiciledV2Nil

`func (o *Iso180135AamvaNamespaceOutput) SetCdlNonDomiciledV2Nil(b bool)`

 SetCdlNonDomiciledV2Nil sets the value for CdlNonDomiciledV2 to be an explicit nil

### UnsetCdlNonDomiciledV2
`func (o *Iso180135AamvaNamespaceOutput) UnsetCdlNonDomiciledV2()`

UnsetCdlNonDomiciledV2 ensures that no value is present for CdlNonDomiciledV2, not even an explicit nil
### GetDhsComplianceText

`func (o *Iso180135AamvaNamespaceOutput) GetDhsComplianceText() string`

GetDhsComplianceText returns the DhsComplianceText field if non-nil, zero value otherwise.

### GetDhsComplianceTextOk

`func (o *Iso180135AamvaNamespaceOutput) GetDhsComplianceTextOk() (*string, bool)`

GetDhsComplianceTextOk returns a tuple with the DhsComplianceText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhsComplianceText

`func (o *Iso180135AamvaNamespaceOutput) SetDhsComplianceText(v string)`

SetDhsComplianceText sets DhsComplianceText field to given value.

### HasDhsComplianceText

`func (o *Iso180135AamvaNamespaceOutput) HasDhsComplianceText() bool`

HasDhsComplianceText returns a boolean if a field has been set.

### SetDhsComplianceTextNil

`func (o *Iso180135AamvaNamespaceOutput) SetDhsComplianceTextNil(b bool)`

 SetDhsComplianceTextNil sets the value for DhsComplianceText to be an explicit nil

### UnsetDhsComplianceText
`func (o *Iso180135AamvaNamespaceOutput) UnsetDhsComplianceText()`

UnsetDhsComplianceText ensures that no value is present for DhsComplianceText, not even an explicit nil
### GetDhsTemporaryLawfulStatus

`func (o *Iso180135AamvaNamespaceOutput) GetDhsTemporaryLawfulStatus() bool`

GetDhsTemporaryLawfulStatus returns the DhsTemporaryLawfulStatus field if non-nil, zero value otherwise.

### GetDhsTemporaryLawfulStatusOk

`func (o *Iso180135AamvaNamespaceOutput) GetDhsTemporaryLawfulStatusOk() (*bool, bool)`

GetDhsTemporaryLawfulStatusOk returns a tuple with the DhsTemporaryLawfulStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhsTemporaryLawfulStatus

`func (o *Iso180135AamvaNamespaceOutput) SetDhsTemporaryLawfulStatus(v bool)`

SetDhsTemporaryLawfulStatus sets DhsTemporaryLawfulStatus field to given value.

### HasDhsTemporaryLawfulStatus

`func (o *Iso180135AamvaNamespaceOutput) HasDhsTemporaryLawfulStatus() bool`

HasDhsTemporaryLawfulStatus returns a boolean if a field has been set.

### SetDhsTemporaryLawfulStatusNil

`func (o *Iso180135AamvaNamespaceOutput) SetDhsTemporaryLawfulStatusNil(b bool)`

 SetDhsTemporaryLawfulStatusNil sets the value for DhsTemporaryLawfulStatus to be an explicit nil

### UnsetDhsTemporaryLawfulStatus
`func (o *Iso180135AamvaNamespaceOutput) UnsetDhsTemporaryLawfulStatus()`

UnsetDhsTemporaryLawfulStatus ensures that no value is present for DhsTemporaryLawfulStatus, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


