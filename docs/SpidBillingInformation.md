# SpidBillingInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsBillable** | **bool** | Whether this SPID verification has resulted in a billable event. | 
**VerificationType** | **string** | The billable verification type for this SPID verification.              Possible values: \&quot;Authentication\&quot; | \&quot;Registration\&quot; | 
**VerificationLevel** | **int32** | The billable verification level for this SPID verification.              Possible values: 1 | 2 | 3 | 

## Methods

### NewSpidBillingInformation

`func NewSpidBillingInformation(isBillable bool, verificationType string, verificationLevel int32, ) *SpidBillingInformation`

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


