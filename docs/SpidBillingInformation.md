# SpidBillingInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsBillable** | Pointer to **NullableBool** | Whether this SPID verification has resulted in a billable event. | [optional] 
**VerificationType** | Pointer to **NullableString** | The billable verification type for this SPID verification.              Possible values: \&quot;Authentication\&quot; | \&quot;Registration\&quot; | [optional] 
**VerificationLevel** | Pointer to **NullableInt32** | The billable verification level for this SPID verification.              Possible values: 1 | 2 | 3 | [optional] 

## Methods

### NewSpidBillingInformation

`func NewSpidBillingInformation() *SpidBillingInformation`

NewSpidBillingInformation instantiates a new SpidBillingInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpidBillingInformationWithDefaults

`func NewSpidBillingInformationWithDefaults() *SpidBillingInformation`

NewSpidBillingInformationWithDefaults instantiates a new SpidBillingInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsBillable

`func (o *SpidBillingInformation) GetIsBillable() bool`

GetIsBillable returns the IsBillable field if non-nil, zero value otherwise.

### GetIsBillableOk

`func (o *SpidBillingInformation) GetIsBillableOk() (*bool, bool)`

GetIsBillableOk returns a tuple with the IsBillable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBillable

`func (o *SpidBillingInformation) SetIsBillable(v bool)`

SetIsBillable sets IsBillable field to given value.

### HasIsBillable

`func (o *SpidBillingInformation) HasIsBillable() bool`

HasIsBillable returns a boolean if a field has been set.

### SetIsBillableNil

`func (o *SpidBillingInformation) SetIsBillableNil(b bool)`

 SetIsBillableNil sets the value for IsBillable to be an explicit nil

### UnsetIsBillable
`func (o *SpidBillingInformation) UnsetIsBillable()`

UnsetIsBillable ensures that no value is present for IsBillable, not even an explicit nil
### GetVerificationType

`func (o *SpidBillingInformation) GetVerificationType() string`

GetVerificationType returns the VerificationType field if non-nil, zero value otherwise.

### GetVerificationTypeOk

`func (o *SpidBillingInformation) GetVerificationTypeOk() (*string, bool)`

GetVerificationTypeOk returns a tuple with the VerificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationType

`func (o *SpidBillingInformation) SetVerificationType(v string)`

SetVerificationType sets VerificationType field to given value.

### HasVerificationType

`func (o *SpidBillingInformation) HasVerificationType() bool`

HasVerificationType returns a boolean if a field has been set.

### SetVerificationTypeNil

`func (o *SpidBillingInformation) SetVerificationTypeNil(b bool)`

 SetVerificationTypeNil sets the value for VerificationType to be an explicit nil

### UnsetVerificationType
`func (o *SpidBillingInformation) UnsetVerificationType()`

UnsetVerificationType ensures that no value is present for VerificationType, not even an explicit nil
### GetVerificationLevel

`func (o *SpidBillingInformation) GetVerificationLevel() int32`

GetVerificationLevel returns the VerificationLevel field if non-nil, zero value otherwise.

### GetVerificationLevelOk

`func (o *SpidBillingInformation) GetVerificationLevelOk() (*int32, bool)`

GetVerificationLevelOk returns a tuple with the VerificationLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationLevel

`func (o *SpidBillingInformation) SetVerificationLevel(v int32)`

SetVerificationLevel sets VerificationLevel field to given value.

### HasVerificationLevel

`func (o *SpidBillingInformation) HasVerificationLevel() bool`

HasVerificationLevel returns a boolean if a field has been set.

### SetVerificationLevelNil

`func (o *SpidBillingInformation) SetVerificationLevelNil(b bool)`

 SetVerificationLevelNil sets the value for VerificationLevel to be an explicit nil

### UnsetVerificationLevel
`func (o *SpidBillingInformation) UnsetVerificationLevel()`

UnsetVerificationLevel ensures that no value is present for VerificationLevel, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


