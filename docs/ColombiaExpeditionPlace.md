# ColombiaExpeditionPlace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Municipality** | Pointer to **NullableString** | Municipality (municipio) where the CC was issued.              This is the second-level administrative division in Colombia. | [optional] 
**Department** | Pointer to **NullableString** | Department (departamento) where the CC was issued.              This is the first-level administrative division in Colombia. | [optional] 

## Methods

### NewColombiaExpeditionPlace

`func NewColombiaExpeditionPlace() *ColombiaExpeditionPlace`

NewColombiaExpeditionPlace instantiates a new ColombiaExpeditionPlace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewColombiaExpeditionPlaceWithDefaults

`func NewColombiaExpeditionPlaceWithDefaults() *ColombiaExpeditionPlace`

NewColombiaExpeditionPlaceWithDefaults instantiates a new ColombiaExpeditionPlace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMunicipality

`func (o *ColombiaExpeditionPlace) GetMunicipality() string`

GetMunicipality returns the Municipality field if non-nil, zero value otherwise.

### GetMunicipalityOk

`func (o *ColombiaExpeditionPlace) GetMunicipalityOk() (*string, bool)`

GetMunicipalityOk returns a tuple with the Municipality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMunicipality

`func (o *ColombiaExpeditionPlace) SetMunicipality(v string)`

SetMunicipality sets Municipality field to given value.

### HasMunicipality

`func (o *ColombiaExpeditionPlace) HasMunicipality() bool`

HasMunicipality returns a boolean if a field has been set.

### SetMunicipalityNil

`func (o *ColombiaExpeditionPlace) SetMunicipalityNil(b bool)`

 SetMunicipalityNil sets the value for Municipality to be an explicit nil

### UnsetMunicipality
`func (o *ColombiaExpeditionPlace) UnsetMunicipality()`

UnsetMunicipality ensures that no value is present for Municipality, not even an explicit nil
### GetDepartment

`func (o *ColombiaExpeditionPlace) GetDepartment() string`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *ColombiaExpeditionPlace) GetDepartmentOk() (*string, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *ColombiaExpeditionPlace) SetDepartment(v string)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *ColombiaExpeditionPlace) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### SetDepartmentNil

`func (o *ColombiaExpeditionPlace) SetDepartmentNil(b bool)`

 SetDepartmentNil sets the value for Department to be an explicit nil

### UnsetDepartment
`func (o *ColombiaExpeditionPlace) UnsetDepartment()`

UnsetDepartment ensures that no value is present for Department, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


