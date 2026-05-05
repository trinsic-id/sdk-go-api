# PanamaCedulaProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FullName** | **string** | Full name from Tribunal Electoral records. | 
**GivenName** | **string** | Given name(s) from Tribunal Electoral records. | 
**FamilyName** | **string** | Family name(s) from Tribunal Electoral records. | 
**DateOfBirth** | Pointer to **NullableString** | Verified date of birth from Tribunal Electoral records. | [optional] 
**DocumentNumber** | **string** | The matched personal identity cédula (Cédula de Identidad Personal) number from Tribunal Electoral records.              On output, Trinsic applies the same normalization as for lookup input. Structure is always {firstSegment}-{libro}-{tomo}: libro is one to four digits and tomo is one to six digits, separated by hyphens.              Citizen category and format: - Born in Panama format: {province}-{libro}-{tomo} ({province} is official code 1 through 13). Examples   8-1234-12345, 4-56-789, 12-12-12345. - Panamanian born abroad format: PE-{libro}-{tomo}. Example PE-1234-12345. - Foreign national with cédula format: E-{libro}-{tomo}. Examples E-1234-12345, E-8-102017. - Naturalized citizen format: N-{libro}-{tomo}. Example N-1234-12345. - Pre-2006 civil registry (AV) format: {province}AV-{libro}-{tomo}. Example 10AV-1234-12345. - Indigenous (PI) format: {province}PI-{libro}-{tomo}. Example 1PI-1234-12345. | 
**NationalityOrResidenceType** | **string** | Inferred from DocumentNumber. Values are nationality or residence category.              Possible values: - BornInPanama - ForeignNational - BornAbroad - Naturalized - LegacyNumber - Indigenous - Unknown (we were unable to determine the category) | 
**SubdivisionOfOrigin** | Pointer to **NullableString** | ISO 3166-2 principal subdivision code. Only available for BornInPanama, LegacyNumber, and Indigenous.              Possible values, matching cédula province digits 1–13: 1. PA-1 - Bocas del Toro 2. PA-2 - Coclé 3. PA-3 - Colón 4. PA-4 - Chiriquí 5. PA-5 - Darién 6. PA-6 - Herrera 7. PA-7 - Los Santos 8. PA-8 - Panamá 9. PA-9 - Veraguas 10. PA-KY - Guna Yala 11. PA-EM - Emberá 12. PA-NB - Ngäbe-Buglé 13. PA-10 - Panamá Oeste | [optional] 
**SubdivisionOfOriginName** | Pointer to **NullableString** | Subdivision display name from the ISO 3166-2 registry, when available. | [optional] 

## Methods

### NewPanamaCedulaProviderOutput

`func NewPanamaCedulaProviderOutput(fullName string, givenName string, familyName string, documentNumber string, nationalityOrResidenceType string, ) *PanamaCedulaProviderOutput`

NewPanamaCedulaProviderOutput instantiates a new PanamaCedulaProviderOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPanamaCedulaProviderOutputWithDefaults

`func NewPanamaCedulaProviderOutputWithDefaults() *PanamaCedulaProviderOutput`

NewPanamaCedulaProviderOutputWithDefaults instantiates a new PanamaCedulaProviderOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFullName

`func (o *PanamaCedulaProviderOutput) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *PanamaCedulaProviderOutput) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *PanamaCedulaProviderOutput) SetFullName(v string)`

SetFullName sets FullName field to given value.


### GetGivenName

`func (o *PanamaCedulaProviderOutput) GetGivenName() string`

GetGivenName returns the GivenName field if non-nil, zero value otherwise.

### GetGivenNameOk

`func (o *PanamaCedulaProviderOutput) GetGivenNameOk() (*string, bool)`

GetGivenNameOk returns a tuple with the GivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenName

`func (o *PanamaCedulaProviderOutput) SetGivenName(v string)`

SetGivenName sets GivenName field to given value.


### GetFamilyName

`func (o *PanamaCedulaProviderOutput) GetFamilyName() string`

GetFamilyName returns the FamilyName field if non-nil, zero value otherwise.

### GetFamilyNameOk

`func (o *PanamaCedulaProviderOutput) GetFamilyNameOk() (*string, bool)`

GetFamilyNameOk returns a tuple with the FamilyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyName

`func (o *PanamaCedulaProviderOutput) SetFamilyName(v string)`

SetFamilyName sets FamilyName field to given value.


### GetDateOfBirth

`func (o *PanamaCedulaProviderOutput) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *PanamaCedulaProviderOutput) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *PanamaCedulaProviderOutput) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *PanamaCedulaProviderOutput) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### SetDateOfBirthNil

`func (o *PanamaCedulaProviderOutput) SetDateOfBirthNil(b bool)`

 SetDateOfBirthNil sets the value for DateOfBirth to be an explicit nil

### UnsetDateOfBirth
`func (o *PanamaCedulaProviderOutput) UnsetDateOfBirth()`

UnsetDateOfBirth ensures that no value is present for DateOfBirth, not even an explicit nil
### GetDocumentNumber

`func (o *PanamaCedulaProviderOutput) GetDocumentNumber() string`

GetDocumentNumber returns the DocumentNumber field if non-nil, zero value otherwise.

### GetDocumentNumberOk

`func (o *PanamaCedulaProviderOutput) GetDocumentNumberOk() (*string, bool)`

GetDocumentNumberOk returns a tuple with the DocumentNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentNumber

`func (o *PanamaCedulaProviderOutput) SetDocumentNumber(v string)`

SetDocumentNumber sets DocumentNumber field to given value.


### GetNationalityOrResidenceType

`func (o *PanamaCedulaProviderOutput) GetNationalityOrResidenceType() string`

GetNationalityOrResidenceType returns the NationalityOrResidenceType field if non-nil, zero value otherwise.

### GetNationalityOrResidenceTypeOk

`func (o *PanamaCedulaProviderOutput) GetNationalityOrResidenceTypeOk() (*string, bool)`

GetNationalityOrResidenceTypeOk returns a tuple with the NationalityOrResidenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationalityOrResidenceType

`func (o *PanamaCedulaProviderOutput) SetNationalityOrResidenceType(v string)`

SetNationalityOrResidenceType sets NationalityOrResidenceType field to given value.


### GetSubdivisionOfOrigin

`func (o *PanamaCedulaProviderOutput) GetSubdivisionOfOrigin() string`

GetSubdivisionOfOrigin returns the SubdivisionOfOrigin field if non-nil, zero value otherwise.

### GetSubdivisionOfOriginOk

`func (o *PanamaCedulaProviderOutput) GetSubdivisionOfOriginOk() (*string, bool)`

GetSubdivisionOfOriginOk returns a tuple with the SubdivisionOfOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdivisionOfOrigin

`func (o *PanamaCedulaProviderOutput) SetSubdivisionOfOrigin(v string)`

SetSubdivisionOfOrigin sets SubdivisionOfOrigin field to given value.

### HasSubdivisionOfOrigin

`func (o *PanamaCedulaProviderOutput) HasSubdivisionOfOrigin() bool`

HasSubdivisionOfOrigin returns a boolean if a field has been set.

### SetSubdivisionOfOriginNil

`func (o *PanamaCedulaProviderOutput) SetSubdivisionOfOriginNil(b bool)`

 SetSubdivisionOfOriginNil sets the value for SubdivisionOfOrigin to be an explicit nil

### UnsetSubdivisionOfOrigin
`func (o *PanamaCedulaProviderOutput) UnsetSubdivisionOfOrigin()`

UnsetSubdivisionOfOrigin ensures that no value is present for SubdivisionOfOrigin, not even an explicit nil
### GetSubdivisionOfOriginName

`func (o *PanamaCedulaProviderOutput) GetSubdivisionOfOriginName() string`

GetSubdivisionOfOriginName returns the SubdivisionOfOriginName field if non-nil, zero value otherwise.

### GetSubdivisionOfOriginNameOk

`func (o *PanamaCedulaProviderOutput) GetSubdivisionOfOriginNameOk() (*string, bool)`

GetSubdivisionOfOriginNameOk returns a tuple with the SubdivisionOfOriginName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdivisionOfOriginName

`func (o *PanamaCedulaProviderOutput) SetSubdivisionOfOriginName(v string)`

SetSubdivisionOfOriginName sets SubdivisionOfOriginName field to given value.

### HasSubdivisionOfOriginName

`func (o *PanamaCedulaProviderOutput) HasSubdivisionOfOriginName() bool`

HasSubdivisionOfOriginName returns a boolean if a field has been set.

### SetSubdivisionOfOriginNameNil

`func (o *PanamaCedulaProviderOutput) SetSubdivisionOfOriginNameNil(b bool)`

 SetSubdivisionOfOriginNameNil sets the value for SubdivisionOfOriginName to be an explicit nil

### UnsetSubdivisionOfOriginName
`func (o *PanamaCedulaProviderOutput) UnsetSubdivisionOfOriginName()`

UnsetSubdivisionOfOriginName ensures that no value is present for SubdivisionOfOriginName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


