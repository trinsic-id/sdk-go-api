# RefreshStepContentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextStep** | [**IntegrationStep**](IntegrationStep.md) | The integration&#39;s next step with refreshed content | 

## Methods

### NewRefreshStepContentResponse

`func NewRefreshStepContentResponse(nextStep IntegrationStep, ) *RefreshStepContentResponse`

NewRefreshStepContentResponse instantiates a new RefreshStepContentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefreshStepContentResponseWithDefaults

`func NewRefreshStepContentResponseWithDefaults() *RefreshStepContentResponse`

NewRefreshStepContentResponseWithDefaults instantiates a new RefreshStepContentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextStep

`func (o *RefreshStepContentResponse) GetNextStep() IntegrationStep`

GetNextStep returns the NextStep field if non-nil, zero value otherwise.

### GetNextStepOk

`func (o *RefreshStepContentResponse) GetNextStepOk() (*IntegrationStep, bool)`

GetNextStepOk returns a tuple with the NextStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextStep

`func (o *RefreshStepContentResponse) SetNextStep(v IntegrationStep)`

SetNextStep sets NextStep field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


