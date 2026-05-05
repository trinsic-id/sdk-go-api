# BoliviaCiProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FullName** | **string** | Full name as it appears on the CI (Cédula de Identidad). | 
**GivenName** | **string** | Given name(s) of the holder as they appear on the CI (Cédula de Identidad). | 
**FamilyName** | **string** | Family name(s) of the holder (paternal and maternal) as they appear on the CI (Cédula de Identidad). | 
**DateOfBirth** | Pointer to **NullableString** | Date of birth as recorded in official Bolivian identity records (RUI — Registro Único de Identificación). | [optional] 
**DocumentNumber** | **string** | The CI (Cédula de Identidad) document number for the matched record.              The document is officially called the Cédula de Identidad and is commonly called carnet or carnet de identidad in Bolivia. This is the identifier assigned by the Servicio General de Identificación Personal (SEGIP) in the Registro Único de Identificación (RUI). The value is entirely numeric. There is no verification digit or other data encoded in the number.              Published regulations do not define a fixed length; digit count may vary. | 

## Methods

### NewBoliviaCiProviderOutput

`func NewBoliviaCiProviderOutput(fullName string, givenName string, familyName string, documentNumber string, ) *BoliviaCiProviderOutput`

NewBoliviaCiProviderOutput instantiates a new BoliviaCiProviderOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBoliviaCiProviderOutputWithDefaults

`func NewBoliviaCiProviderOutputWithDefaults() *BoliviaCiProviderOutput`

NewBoliviaCiProviderOutputWithDefaults instantiates a new BoliviaCiProviderOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFullName

`func (o *BoliviaCiProviderOutput) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *BoliviaCiProviderOutput) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *BoliviaCiProviderOutput) SetFullName(v string)`

SetFullName sets FullName field to given value.


### GetGivenName

`func (o *BoliviaCiProviderOutput) GetGivenName() string`

GetGivenName returns the GivenName field if non-nil, zero value otherwise.

### GetGivenNameOk

`func (o *BoliviaCiProviderOutput) GetGivenNameOk() (*string, bool)`

GetGivenNameOk returns a tuple with the GivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenName

`func (o *BoliviaCiProviderOutput) SetGivenName(v string)`

SetGivenName sets GivenName field to given value.


### GetFamilyName

`func (o *BoliviaCiProviderOutput) GetFamilyName() string`

GetFamilyName returns the FamilyName field if non-nil, zero value otherwise.

### GetFamilyNameOk

`func (o *BoliviaCiProviderOutput) GetFamilyNameOk() (*string, bool)`

GetFamilyNameOk returns a tuple with the FamilyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyName

`func (o *BoliviaCiProviderOutput) SetFamilyName(v string)`

SetFamilyName sets FamilyName field to given value.


### GetDateOfBirth

`func (o *BoliviaCiProviderOutput) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *BoliviaCiProviderOutput) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *BoliviaCiProviderOutput) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *BoliviaCiProviderOutput) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### SetDateOfBirthNil

`func (o *BoliviaCiProviderOutput) SetDateOfBirthNil(b bool)`

 SetDateOfBirthNil sets the value for DateOfBirth to be an explicit nil

### UnsetDateOfBirth
`func (o *BoliviaCiProviderOutput) UnsetDateOfBirth()`

UnsetDateOfBirth ensures that no value is present for DateOfBirth, not even an explicit nil
### GetDocumentNumber

`func (o *BoliviaCiProviderOutput) GetDocumentNumber() string`

GetDocumentNumber returns the DocumentNumber field if non-nil, zero value otherwise.

### GetDocumentNumberOk

`func (o *BoliviaCiProviderOutput) GetDocumentNumberOk() (*string, bool)`

GetDocumentNumberOk returns a tuple with the DocumentNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentNumber

`func (o *BoliviaCiProviderOutput) SetDocumentNumber(v string)`

SetDocumentNumber sets DocumentNumber field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


