# ColombiaCcProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FullName** | Pointer to **NullableString** | Full name as it appears on the CC. | [optional] 
**GivenName** | Pointer to **NullableString** | Given name(s) of the holder as they appear on the CC. | [optional] 
**FamilyName** | Pointer to **NullableString** | Family name(s) of the holder as they appear on the CC. Space-separated when both paternal and maternal family names are present. | [optional] 
**DateOfBirth** | Pointer to **NullableString** | Date of birth as recorded in the civil registry (Registraduría Nacional). | [optional] 
**Sex** | Pointer to **NullableString** | Sex of the holder as recorded in the civil registry (Registraduría Nacional).              Possible values: - Male - Female - Unknown (when the sex is not recorded or cannot be confidently determined) | [optional] 
**IsAlive** | Pointer to **NullableBool** | Whether the person is reported as alive in Colombia&#39;s official civil registry (Registraduría Nacional).              Used to detect identity fraud when the holder is deceased. | [optional] 
**DocumentNumber** | Pointer to **NullableString** | The Cédula de Ciudadanía (CC) document number.              This is the unique identifier assigned by the Registraduría Nacional when the person is first issued a CC. It does not change when the person renews or receives a new physical card; it remains the same for the individual for life.              Cédulas issued after 2004 use the NUIP (Número Único de Identificación Personal), which is 10 digits. Older documents may have fewer than 10 digits and are still valid. | [optional] 
**ExpeditionDate** | Pointer to **NullableString** | Date the CC was issued (fecha de expedición).              Format: - yyyy-MM-dd | [optional] 
**ExpeditionPlace** | Pointer to [**NullableColombiaExpeditionPlace**](ColombiaExpeditionPlace.md) | Place where the CC was issued (lugar de expedición): municipality and department as recorded by the Registraduría Nacional. | [optional] 
**ArrayName** | Pointer to **[]string** | All names as they appear on the CC, as an array of strings.              Format: - Order follows the civil registry: typically family name(s) first, then given name(s). | [optional] 

## Methods

### NewColombiaCcProviderOutput

`func NewColombiaCcProviderOutput() *ColombiaCcProviderOutput`

NewColombiaCcProviderOutput instantiates a new ColombiaCcProviderOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewColombiaCcProviderOutputWithDefaults

`func NewColombiaCcProviderOutputWithDefaults() *ColombiaCcProviderOutput`

NewColombiaCcProviderOutputWithDefaults instantiates a new ColombiaCcProviderOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFullName

`func (o *ColombiaCcProviderOutput) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *ColombiaCcProviderOutput) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *ColombiaCcProviderOutput) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *ColombiaCcProviderOutput) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### SetFullNameNil

`func (o *ColombiaCcProviderOutput) SetFullNameNil(b bool)`

 SetFullNameNil sets the value for FullName to be an explicit nil

### UnsetFullName
`func (o *ColombiaCcProviderOutput) UnsetFullName()`

UnsetFullName ensures that no value is present for FullName, not even an explicit nil
### GetGivenName

`func (o *ColombiaCcProviderOutput) GetGivenName() string`

GetGivenName returns the GivenName field if non-nil, zero value otherwise.

### GetGivenNameOk

`func (o *ColombiaCcProviderOutput) GetGivenNameOk() (*string, bool)`

GetGivenNameOk returns a tuple with the GivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenName

`func (o *ColombiaCcProviderOutput) SetGivenName(v string)`

SetGivenName sets GivenName field to given value.

### HasGivenName

`func (o *ColombiaCcProviderOutput) HasGivenName() bool`

HasGivenName returns a boolean if a field has been set.

### SetGivenNameNil

`func (o *ColombiaCcProviderOutput) SetGivenNameNil(b bool)`

 SetGivenNameNil sets the value for GivenName to be an explicit nil

### UnsetGivenName
`func (o *ColombiaCcProviderOutput) UnsetGivenName()`

UnsetGivenName ensures that no value is present for GivenName, not even an explicit nil
### GetFamilyName

`func (o *ColombiaCcProviderOutput) GetFamilyName() string`

GetFamilyName returns the FamilyName field if non-nil, zero value otherwise.

### GetFamilyNameOk

`func (o *ColombiaCcProviderOutput) GetFamilyNameOk() (*string, bool)`

GetFamilyNameOk returns a tuple with the FamilyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyName

`func (o *ColombiaCcProviderOutput) SetFamilyName(v string)`

SetFamilyName sets FamilyName field to given value.

### HasFamilyName

`func (o *ColombiaCcProviderOutput) HasFamilyName() bool`

HasFamilyName returns a boolean if a field has been set.

### SetFamilyNameNil

`func (o *ColombiaCcProviderOutput) SetFamilyNameNil(b bool)`

 SetFamilyNameNil sets the value for FamilyName to be an explicit nil

### UnsetFamilyName
`func (o *ColombiaCcProviderOutput) UnsetFamilyName()`

UnsetFamilyName ensures that no value is present for FamilyName, not even an explicit nil
### GetDateOfBirth

`func (o *ColombiaCcProviderOutput) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *ColombiaCcProviderOutput) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *ColombiaCcProviderOutput) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *ColombiaCcProviderOutput) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### SetDateOfBirthNil

`func (o *ColombiaCcProviderOutput) SetDateOfBirthNil(b bool)`

 SetDateOfBirthNil sets the value for DateOfBirth to be an explicit nil

### UnsetDateOfBirth
`func (o *ColombiaCcProviderOutput) UnsetDateOfBirth()`

UnsetDateOfBirth ensures that no value is present for DateOfBirth, not even an explicit nil
### GetSex

`func (o *ColombiaCcProviderOutput) GetSex() string`

GetSex returns the Sex field if non-nil, zero value otherwise.

### GetSexOk

`func (o *ColombiaCcProviderOutput) GetSexOk() (*string, bool)`

GetSexOk returns a tuple with the Sex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSex

`func (o *ColombiaCcProviderOutput) SetSex(v string)`

SetSex sets Sex field to given value.

### HasSex

`func (o *ColombiaCcProviderOutput) HasSex() bool`

HasSex returns a boolean if a field has been set.

### SetSexNil

`func (o *ColombiaCcProviderOutput) SetSexNil(b bool)`

 SetSexNil sets the value for Sex to be an explicit nil

### UnsetSex
`func (o *ColombiaCcProviderOutput) UnsetSex()`

UnsetSex ensures that no value is present for Sex, not even an explicit nil
### GetIsAlive

`func (o *ColombiaCcProviderOutput) GetIsAlive() bool`

GetIsAlive returns the IsAlive field if non-nil, zero value otherwise.

### GetIsAliveOk

`func (o *ColombiaCcProviderOutput) GetIsAliveOk() (*bool, bool)`

GetIsAliveOk returns a tuple with the IsAlive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAlive

`func (o *ColombiaCcProviderOutput) SetIsAlive(v bool)`

SetIsAlive sets IsAlive field to given value.

### HasIsAlive

`func (o *ColombiaCcProviderOutput) HasIsAlive() bool`

HasIsAlive returns a boolean if a field has been set.

### SetIsAliveNil

`func (o *ColombiaCcProviderOutput) SetIsAliveNil(b bool)`

 SetIsAliveNil sets the value for IsAlive to be an explicit nil

### UnsetIsAlive
`func (o *ColombiaCcProviderOutput) UnsetIsAlive()`

UnsetIsAlive ensures that no value is present for IsAlive, not even an explicit nil
### GetDocumentNumber

`func (o *ColombiaCcProviderOutput) GetDocumentNumber() string`

GetDocumentNumber returns the DocumentNumber field if non-nil, zero value otherwise.

### GetDocumentNumberOk

`func (o *ColombiaCcProviderOutput) GetDocumentNumberOk() (*string, bool)`

GetDocumentNumberOk returns a tuple with the DocumentNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentNumber

`func (o *ColombiaCcProviderOutput) SetDocumentNumber(v string)`

SetDocumentNumber sets DocumentNumber field to given value.

### HasDocumentNumber

`func (o *ColombiaCcProviderOutput) HasDocumentNumber() bool`

HasDocumentNumber returns a boolean if a field has been set.

### SetDocumentNumberNil

`func (o *ColombiaCcProviderOutput) SetDocumentNumberNil(b bool)`

 SetDocumentNumberNil sets the value for DocumentNumber to be an explicit nil

### UnsetDocumentNumber
`func (o *ColombiaCcProviderOutput) UnsetDocumentNumber()`

UnsetDocumentNumber ensures that no value is present for DocumentNumber, not even an explicit nil
### GetExpeditionDate

`func (o *ColombiaCcProviderOutput) GetExpeditionDate() string`

GetExpeditionDate returns the ExpeditionDate field if non-nil, zero value otherwise.

### GetExpeditionDateOk

`func (o *ColombiaCcProviderOutput) GetExpeditionDateOk() (*string, bool)`

GetExpeditionDateOk returns a tuple with the ExpeditionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpeditionDate

`func (o *ColombiaCcProviderOutput) SetExpeditionDate(v string)`

SetExpeditionDate sets ExpeditionDate field to given value.

### HasExpeditionDate

`func (o *ColombiaCcProviderOutput) HasExpeditionDate() bool`

HasExpeditionDate returns a boolean if a field has been set.

### SetExpeditionDateNil

`func (o *ColombiaCcProviderOutput) SetExpeditionDateNil(b bool)`

 SetExpeditionDateNil sets the value for ExpeditionDate to be an explicit nil

### UnsetExpeditionDate
`func (o *ColombiaCcProviderOutput) UnsetExpeditionDate()`

UnsetExpeditionDate ensures that no value is present for ExpeditionDate, not even an explicit nil
### GetExpeditionPlace

`func (o *ColombiaCcProviderOutput) GetExpeditionPlace() ColombiaExpeditionPlace`

GetExpeditionPlace returns the ExpeditionPlace field if non-nil, zero value otherwise.

### GetExpeditionPlaceOk

`func (o *ColombiaCcProviderOutput) GetExpeditionPlaceOk() (*ColombiaExpeditionPlace, bool)`

GetExpeditionPlaceOk returns a tuple with the ExpeditionPlace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpeditionPlace

`func (o *ColombiaCcProviderOutput) SetExpeditionPlace(v ColombiaExpeditionPlace)`

SetExpeditionPlace sets ExpeditionPlace field to given value.

### HasExpeditionPlace

`func (o *ColombiaCcProviderOutput) HasExpeditionPlace() bool`

HasExpeditionPlace returns a boolean if a field has been set.

### SetExpeditionPlaceNil

`func (o *ColombiaCcProviderOutput) SetExpeditionPlaceNil(b bool)`

 SetExpeditionPlaceNil sets the value for ExpeditionPlace to be an explicit nil

### UnsetExpeditionPlace
`func (o *ColombiaCcProviderOutput) UnsetExpeditionPlace()`

UnsetExpeditionPlace ensures that no value is present for ExpeditionPlace, not even an explicit nil
### GetArrayName

`func (o *ColombiaCcProviderOutput) GetArrayName() []string`

GetArrayName returns the ArrayName field if non-nil, zero value otherwise.

### GetArrayNameOk

`func (o *ColombiaCcProviderOutput) GetArrayNameOk() (*[]string, bool)`

GetArrayNameOk returns a tuple with the ArrayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrayName

`func (o *ColombiaCcProviderOutput) SetArrayName(v []string)`

SetArrayName sets ArrayName field to given value.

### HasArrayName

`func (o *ColombiaCcProviderOutput) HasArrayName() bool`

HasArrayName returns a boolean if a field has been set.

### SetArrayNameNil

`func (o *ColombiaCcProviderOutput) SetArrayNameNil(b bool)`

 SetArrayNameNil sets the value for ArrayName to be an explicit nil

### UnsetArrayName
`func (o *ColombiaCcProviderOutput) UnsetArrayName()`

UnsetArrayName ensures that no value is present for ArrayName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


