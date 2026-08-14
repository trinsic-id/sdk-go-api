# ClearProviderOutputCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VerificationCheckName** | Pointer to **NullableString** | The verification check name as specified in the CLEAR verification configuration. | [optional] 
**Value** | Pointer to **NullableBool** | The boolean check value.              This can be null when a check is skipped or unable to evaluate to a deterministic result. | [optional] 
**Status** | Pointer to **NullableString** | Whether CLEAR performed the check.              Known values: - completed - skipped - error | [optional] 
**CheckResult** | Pointer to **NullableString** | The granular CLEAR result for this check.              Known values: - success - failure - indeterminate - not_applicable - awaiting_async_response | [optional] 
**AdditionalDetails** | Pointer to [**NullableClearProviderOutputCheckAdditionalDetails**](ClearProviderOutputCheckAdditionalDetails.md) | Additional structured details CLEAR used for this check. Currently modeled details: watchlistHits. | [optional] 
**Params** | Pointer to **map[string]interface{}** | Custom parameters configured for this CLEAR check. | [optional] 

## Methods

### NewClearProviderOutputCheck

`func NewClearProviderOutputCheck() *ClearProviderOutputCheck`

NewClearProviderOutputCheck instantiates a new ClearProviderOutputCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClearProviderOutputCheckWithDefaults

`func NewClearProviderOutputCheckWithDefaults() *ClearProviderOutputCheck`

NewClearProviderOutputCheckWithDefaults instantiates a new ClearProviderOutputCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVerificationCheckName

`func (o *ClearProviderOutputCheck) GetVerificationCheckName() string`

GetVerificationCheckName returns the VerificationCheckName field if non-nil, zero value otherwise.

### GetVerificationCheckNameOk

`func (o *ClearProviderOutputCheck) GetVerificationCheckNameOk() (*string, bool)`

GetVerificationCheckNameOk returns a tuple with the VerificationCheckName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationCheckName

`func (o *ClearProviderOutputCheck) SetVerificationCheckName(v string)`

SetVerificationCheckName sets VerificationCheckName field to given value.

### HasVerificationCheckName

`func (o *ClearProviderOutputCheck) HasVerificationCheckName() bool`

HasVerificationCheckName returns a boolean if a field has been set.

### SetVerificationCheckNameNil

`func (o *ClearProviderOutputCheck) SetVerificationCheckNameNil(b bool)`

 SetVerificationCheckNameNil sets the value for VerificationCheckName to be an explicit nil

### UnsetVerificationCheckName
`func (o *ClearProviderOutputCheck) UnsetVerificationCheckName()`

UnsetVerificationCheckName ensures that no value is present for VerificationCheckName, not even an explicit nil
### GetValue

`func (o *ClearProviderOutputCheck) GetValue() bool`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ClearProviderOutputCheck) GetValueOk() (*bool, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ClearProviderOutputCheck) SetValue(v bool)`

SetValue sets Value field to given value.

### HasValue

`func (o *ClearProviderOutputCheck) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *ClearProviderOutputCheck) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *ClearProviderOutputCheck) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetStatus

`func (o *ClearProviderOutputCheck) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ClearProviderOutputCheck) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ClearProviderOutputCheck) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ClearProviderOutputCheck) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ClearProviderOutputCheck) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ClearProviderOutputCheck) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetCheckResult

`func (o *ClearProviderOutputCheck) GetCheckResult() string`

GetCheckResult returns the CheckResult field if non-nil, zero value otherwise.

### GetCheckResultOk

`func (o *ClearProviderOutputCheck) GetCheckResultOk() (*string, bool)`

GetCheckResultOk returns a tuple with the CheckResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckResult

`func (o *ClearProviderOutputCheck) SetCheckResult(v string)`

SetCheckResult sets CheckResult field to given value.

### HasCheckResult

`func (o *ClearProviderOutputCheck) HasCheckResult() bool`

HasCheckResult returns a boolean if a field has been set.

### SetCheckResultNil

`func (o *ClearProviderOutputCheck) SetCheckResultNil(b bool)`

 SetCheckResultNil sets the value for CheckResult to be an explicit nil

### UnsetCheckResult
`func (o *ClearProviderOutputCheck) UnsetCheckResult()`

UnsetCheckResult ensures that no value is present for CheckResult, not even an explicit nil
### GetAdditionalDetails

`func (o *ClearProviderOutputCheck) GetAdditionalDetails() ClearProviderOutputCheckAdditionalDetails`

GetAdditionalDetails returns the AdditionalDetails field if non-nil, zero value otherwise.

### GetAdditionalDetailsOk

`func (o *ClearProviderOutputCheck) GetAdditionalDetailsOk() (*ClearProviderOutputCheckAdditionalDetails, bool)`

GetAdditionalDetailsOk returns a tuple with the AdditionalDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalDetails

`func (o *ClearProviderOutputCheck) SetAdditionalDetails(v ClearProviderOutputCheckAdditionalDetails)`

SetAdditionalDetails sets AdditionalDetails field to given value.

### HasAdditionalDetails

`func (o *ClearProviderOutputCheck) HasAdditionalDetails() bool`

HasAdditionalDetails returns a boolean if a field has been set.

### SetAdditionalDetailsNil

`func (o *ClearProviderOutputCheck) SetAdditionalDetailsNil(b bool)`

 SetAdditionalDetailsNil sets the value for AdditionalDetails to be an explicit nil

### UnsetAdditionalDetails
`func (o *ClearProviderOutputCheck) UnsetAdditionalDetails()`

UnsetAdditionalDetails ensures that no value is present for AdditionalDetails, not even an explicit nil
### GetParams

`func (o *ClearProviderOutputCheck) GetParams() map[string]interface{}`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *ClearProviderOutputCheck) GetParamsOk() (*map[string]interface{}, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *ClearProviderOutputCheck) SetParams(v map[string]interface{})`

SetParams sets Params field to given value.

### HasParams

`func (o *ClearProviderOutputCheck) HasParams() bool`

HasParams returns a boolean if a field has been set.

### SetParamsNil

`func (o *ClearProviderOutputCheck) SetParamsNil(b bool)`

 SetParamsNil sets the value for Params to be an explicit nil

### UnsetParams
`func (o *ClearProviderOutputCheck) UnsetParams()`

UnsetParams ensures that no value is present for Params, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


