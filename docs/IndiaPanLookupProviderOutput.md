# IndiaPanLookupProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PermanentAccountNumber** | Pointer to **NullableString** | Permanent Account Number (PAN) that was looked up.              PAN is a ten-character identifier issued by the Income Tax Department of India. | [optional] 
**StructuredPermanentAccountNumber** | Pointer to [**NullableIndiaPanStructuredNumber**](IndiaPanStructuredNumber.md) | PAN split into structural parts for the canonical ten-character format.              Covers alphabetic series, assessee category, name prefix letter, serial number, and check letter. | [optional] 
**FullName** | Pointer to **NullableString** | Full name of the entity subject to income tax.              NOTE: Only available when PAN status is VALID. | [optional] 
**GivenName** | Pointer to **NullableString** | Given name of the individual.              NOTE: Only available when PAN status is VALID and the verification is for an individual person. | [optional] 
**MiddleName** | Pointer to **NullableString** | Middle name of the individual.              NOTE: Only available when PAN status is VALID and the verification is for an individual person. | [optional] 
**FamilyName** | Pointer to **NullableString** | Family name of the individual.              NOTE: Only available when PAN status is VALID and the verification is for an individual person. | [optional] 
**DateOfBirth** | Pointer to **NullableString** | Date of birth of the individual.              NOTE: Only available when PAN status is VALID and the verification is for an individual person. | [optional] 
**PanStatus** | Pointer to **NullableString** | PAN card status from Signzy for the lookup.              Possible values: - VALID - FAKE - DEACTIVATED - DELETED - INVALID - AMALGAMATION - ACQUISITION - DEATH - DISSOLUTION - LIQUIDATED - MERGER - PARTITION - SPLIT - UNDER LIQUIDATION - INOPERATIVE | [optional] 

## Methods

### NewIndiaPanLookupProviderOutput

`func NewIndiaPanLookupProviderOutput() *IndiaPanLookupProviderOutput`

NewIndiaPanLookupProviderOutput instantiates a new IndiaPanLookupProviderOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndiaPanLookupProviderOutputWithDefaults

`func NewIndiaPanLookupProviderOutputWithDefaults() *IndiaPanLookupProviderOutput`

NewIndiaPanLookupProviderOutputWithDefaults instantiates a new IndiaPanLookupProviderOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermanentAccountNumber

`func (o *IndiaPanLookupProviderOutput) GetPermanentAccountNumber() string`

GetPermanentAccountNumber returns the PermanentAccountNumber field if non-nil, zero value otherwise.

### GetPermanentAccountNumberOk

`func (o *IndiaPanLookupProviderOutput) GetPermanentAccountNumberOk() (*string, bool)`

GetPermanentAccountNumberOk returns a tuple with the PermanentAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermanentAccountNumber

`func (o *IndiaPanLookupProviderOutput) SetPermanentAccountNumber(v string)`

SetPermanentAccountNumber sets PermanentAccountNumber field to given value.

### HasPermanentAccountNumber

`func (o *IndiaPanLookupProviderOutput) HasPermanentAccountNumber() bool`

HasPermanentAccountNumber returns a boolean if a field has been set.

### SetPermanentAccountNumberNil

`func (o *IndiaPanLookupProviderOutput) SetPermanentAccountNumberNil(b bool)`

 SetPermanentAccountNumberNil sets the value for PermanentAccountNumber to be an explicit nil

### UnsetPermanentAccountNumber
`func (o *IndiaPanLookupProviderOutput) UnsetPermanentAccountNumber()`

UnsetPermanentAccountNumber ensures that no value is present for PermanentAccountNumber, not even an explicit nil
### GetStructuredPermanentAccountNumber

`func (o *IndiaPanLookupProviderOutput) GetStructuredPermanentAccountNumber() IndiaPanStructuredNumber`

GetStructuredPermanentAccountNumber returns the StructuredPermanentAccountNumber field if non-nil, zero value otherwise.

### GetStructuredPermanentAccountNumberOk

`func (o *IndiaPanLookupProviderOutput) GetStructuredPermanentAccountNumberOk() (*IndiaPanStructuredNumber, bool)`

GetStructuredPermanentAccountNumberOk returns a tuple with the StructuredPermanentAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructuredPermanentAccountNumber

`func (o *IndiaPanLookupProviderOutput) SetStructuredPermanentAccountNumber(v IndiaPanStructuredNumber)`

SetStructuredPermanentAccountNumber sets StructuredPermanentAccountNumber field to given value.

### HasStructuredPermanentAccountNumber

`func (o *IndiaPanLookupProviderOutput) HasStructuredPermanentAccountNumber() bool`

HasStructuredPermanentAccountNumber returns a boolean if a field has been set.

### SetStructuredPermanentAccountNumberNil

`func (o *IndiaPanLookupProviderOutput) SetStructuredPermanentAccountNumberNil(b bool)`

 SetStructuredPermanentAccountNumberNil sets the value for StructuredPermanentAccountNumber to be an explicit nil

### UnsetStructuredPermanentAccountNumber
`func (o *IndiaPanLookupProviderOutput) UnsetStructuredPermanentAccountNumber()`

UnsetStructuredPermanentAccountNumber ensures that no value is present for StructuredPermanentAccountNumber, not even an explicit nil
### GetFullName

`func (o *IndiaPanLookupProviderOutput) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *IndiaPanLookupProviderOutput) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *IndiaPanLookupProviderOutput) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *IndiaPanLookupProviderOutput) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### SetFullNameNil

`func (o *IndiaPanLookupProviderOutput) SetFullNameNil(b bool)`

 SetFullNameNil sets the value for FullName to be an explicit nil

### UnsetFullName
`func (o *IndiaPanLookupProviderOutput) UnsetFullName()`

UnsetFullName ensures that no value is present for FullName, not even an explicit nil
### GetGivenName

`func (o *IndiaPanLookupProviderOutput) GetGivenName() string`

GetGivenName returns the GivenName field if non-nil, zero value otherwise.

### GetGivenNameOk

`func (o *IndiaPanLookupProviderOutput) GetGivenNameOk() (*string, bool)`

GetGivenNameOk returns a tuple with the GivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenName

`func (o *IndiaPanLookupProviderOutput) SetGivenName(v string)`

SetGivenName sets GivenName field to given value.

### HasGivenName

`func (o *IndiaPanLookupProviderOutput) HasGivenName() bool`

HasGivenName returns a boolean if a field has been set.

### SetGivenNameNil

`func (o *IndiaPanLookupProviderOutput) SetGivenNameNil(b bool)`

 SetGivenNameNil sets the value for GivenName to be an explicit nil

### UnsetGivenName
`func (o *IndiaPanLookupProviderOutput) UnsetGivenName()`

UnsetGivenName ensures that no value is present for GivenName, not even an explicit nil
### GetMiddleName

`func (o *IndiaPanLookupProviderOutput) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *IndiaPanLookupProviderOutput) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *IndiaPanLookupProviderOutput) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *IndiaPanLookupProviderOutput) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### SetMiddleNameNil

`func (o *IndiaPanLookupProviderOutput) SetMiddleNameNil(b bool)`

 SetMiddleNameNil sets the value for MiddleName to be an explicit nil

### UnsetMiddleName
`func (o *IndiaPanLookupProviderOutput) UnsetMiddleName()`

UnsetMiddleName ensures that no value is present for MiddleName, not even an explicit nil
### GetFamilyName

`func (o *IndiaPanLookupProviderOutput) GetFamilyName() string`

GetFamilyName returns the FamilyName field if non-nil, zero value otherwise.

### GetFamilyNameOk

`func (o *IndiaPanLookupProviderOutput) GetFamilyNameOk() (*string, bool)`

GetFamilyNameOk returns a tuple with the FamilyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyName

`func (o *IndiaPanLookupProviderOutput) SetFamilyName(v string)`

SetFamilyName sets FamilyName field to given value.

### HasFamilyName

`func (o *IndiaPanLookupProviderOutput) HasFamilyName() bool`

HasFamilyName returns a boolean if a field has been set.

### SetFamilyNameNil

`func (o *IndiaPanLookupProviderOutput) SetFamilyNameNil(b bool)`

 SetFamilyNameNil sets the value for FamilyName to be an explicit nil

### UnsetFamilyName
`func (o *IndiaPanLookupProviderOutput) UnsetFamilyName()`

UnsetFamilyName ensures that no value is present for FamilyName, not even an explicit nil
### GetDateOfBirth

`func (o *IndiaPanLookupProviderOutput) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *IndiaPanLookupProviderOutput) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *IndiaPanLookupProviderOutput) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *IndiaPanLookupProviderOutput) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### SetDateOfBirthNil

`func (o *IndiaPanLookupProviderOutput) SetDateOfBirthNil(b bool)`

 SetDateOfBirthNil sets the value for DateOfBirth to be an explicit nil

### UnsetDateOfBirth
`func (o *IndiaPanLookupProviderOutput) UnsetDateOfBirth()`

UnsetDateOfBirth ensures that no value is present for DateOfBirth, not even an explicit nil
### GetPanStatus

`func (o *IndiaPanLookupProviderOutput) GetPanStatus() string`

GetPanStatus returns the PanStatus field if non-nil, zero value otherwise.

### GetPanStatusOk

`func (o *IndiaPanLookupProviderOutput) GetPanStatusOk() (*string, bool)`

GetPanStatusOk returns a tuple with the PanStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPanStatus

`func (o *IndiaPanLookupProviderOutput) SetPanStatus(v string)`

SetPanStatus sets PanStatus field to given value.

### HasPanStatus

`func (o *IndiaPanLookupProviderOutput) HasPanStatus() bool`

HasPanStatus returns a boolean if a field has been set.

### SetPanStatusNil

`func (o *IndiaPanLookupProviderOutput) SetPanStatusNil(b bool)`

 SetPanStatusNil sets the value for PanStatus to be an explicit nil

### UnsetPanStatus
`func (o *IndiaPanLookupProviderOutput) UnsetPanStatus()`

UnsetPanStatus ensures that no value is present for PanStatus, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


