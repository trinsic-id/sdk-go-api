# StepRefreshInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiresAt** | **time.Time** | The UTC date/time at which the step &#x60;content&#x60; will expire and should no longer be used.              Use the &#x60;Refresh Step Content&#x60; API to obtain a new value for &#x60;content&#x60;. | 
**RefreshAfter** | **time.Time** | The UTC date/time after which Trinsic recommends you refresh the step &#x60;content&#x60;. | 
**TimeToLiveSeconds** | **int32** | The total lifetime of the step &#x60;content&#x60;. | 

## Methods

### NewStepRefreshInfo

`func NewStepRefreshInfo(expiresAt time.Time, refreshAfter time.Time, timeToLiveSeconds int32, ) *StepRefreshInfo`

NewStepRefreshInfo instantiates a new StepRefreshInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStepRefreshInfoWithDefaults

`func NewStepRefreshInfoWithDefaults() *StepRefreshInfo`

NewStepRefreshInfoWithDefaults instantiates a new StepRefreshInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiresAt

`func (o *StepRefreshInfo) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *StepRefreshInfo) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *StepRefreshInfo) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.


### GetRefreshAfter

`func (o *StepRefreshInfo) GetRefreshAfter() time.Time`

GetRefreshAfter returns the RefreshAfter field if non-nil, zero value otherwise.

### GetRefreshAfterOk

`func (o *StepRefreshInfo) GetRefreshAfterOk() (*time.Time, bool)`

GetRefreshAfterOk returns a tuple with the RefreshAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshAfter

`func (o *StepRefreshInfo) SetRefreshAfter(v time.Time)`

SetRefreshAfter sets RefreshAfter field to given value.


### GetTimeToLiveSeconds

`func (o *StepRefreshInfo) GetTimeToLiveSeconds() int32`

GetTimeToLiveSeconds returns the TimeToLiveSeconds field if non-nil, zero value otherwise.

### GetTimeToLiveSecondsOk

`func (o *StepRefreshInfo) GetTimeToLiveSecondsOk() (*int32, bool)`

GetTimeToLiveSecondsOk returns a tuple with the TimeToLiveSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeToLiveSeconds

`func (o *StepRefreshInfo) SetTimeToLiveSeconds(v int32)`

SetTimeToLiveSeconds sets TimeToLiveSeconds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


