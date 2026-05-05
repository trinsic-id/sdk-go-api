# ColombiaCcProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FullName** | **string** | Full name as it appears on the CC. | 
**GivenName** | **string** | Given name(s) of the holder as they appear on the CC. | 
**FamilyName** | **string** | Family name(s) of the holder as they appear on the CC. Space-separated when both paternal and maternal family names are present. | 
**DateOfBirth** | **string** | Date of birth as recorded in the civil registry (Registraduría Nacional). | 
**Sex** | **string** | Sex of the holder as recorded in the civil registry (Registraduría Nacional).              Possible values: - Male - Female - Unknown (when the sex is not recorded or cannot be confidently determined) | 
**IsAlive** | **bool** | Whether the person is reported as alive in Colombia&#39;s official civil registry (Registraduría Nacional).              Used to detect identity fraud when the holder is deceased. | 
**DocumentNumber** | **string** | The Cédula de Ciudadanía (CC) document number.              This is the unique identifier assigned by the Registraduría Nacional when the person is first issued a CC. It does not change when the person renews or receives a new physical card; it remains the same for the individual for life.              Cédulas issued after 2004 use the NUIP (Número Único de Identificación Personal), which is 10 digits. Older documents may have fewer than 10 digits and are still valid. | 
**ExpeditionDate** | **string** | Date the CC was issued (fecha de expedición).              Format: - yyyy-MM-dd | 
**ExpeditionPlace** | [**ColombiaExpeditionPlace**](ColombiaExpeditionPlace.md) | Place where the CC was issued (lugar de expedición): municipality and department as recorded by the Registraduría Nacional. | 
**ArrayName** | **[]string** | All names as they appear on the CC, as an array of strings.              Format: - Order follows the civil registry: typically family name(s) first, then given name(s). | 

## Methods

### NewColombiaCcProviderOutput

`func NewColombiaCcProviderOutput(fullName string, givenName string, familyName string, dateOfBirth string, sex string, isAlive bool, documentNumber string, expeditionDate string, expeditionPlace ColombiaExpeditionPlace, arrayName []string, ) *ColombiaCcProviderOutput`

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


