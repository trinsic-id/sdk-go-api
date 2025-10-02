# CreateDirectProviderSessionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SessionId** | **string** | The ID of the newly-created Acceptance Session | 
**ResultCollection** | [**ResultCollection**](ResultCollection.md) | The method by which you must collect the results of the Acceptance Session. | 
**NextStep** | [**IntegrationStep**](IntegrationStep.md) | The next step you must take to launch the user into the integration | 

## Methods

### NewCreateDirectProviderSessionResponse

`func NewCreateDirectProviderSessionResponse(sessionId string, resultCollection ResultCollection, nextStep IntegrationStep, ) *CreateDirectProviderSessionResponse`

NewCreateDirectProviderSessionResponse instantiates a new CreateDirectProviderSessionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDirectProviderSessionResponseWithDefaults

`func NewCreateDirectProviderSessionResponseWithDefaults() *CreateDirectProviderSessionResponse`

NewCreateDirectProviderSessionResponseWithDefaults instantiates a new CreateDirectProviderSessionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSessionId

`func (o *CreateDirectProviderSessionResponse) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *CreateDirectProviderSessionResponse) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *CreateDirectProviderSessionResponse) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.


### GetResultCollection

`func (o *CreateDirectProviderSessionResponse) GetResultCollection() ResultCollection`

GetResultCollection returns the ResultCollection field if non-nil, zero value otherwise.

### GetResultCollectionOk

`func (o *CreateDirectProviderSessionResponse) GetResultCollectionOk() (*ResultCollection, bool)`

GetResultCollectionOk returns a tuple with the ResultCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultCollection

`func (o *CreateDirectProviderSessionResponse) SetResultCollection(v ResultCollection)`

SetResultCollection sets ResultCollection field to given value.


### GetNextStep

`func (o *CreateDirectProviderSessionResponse) GetNextStep() IntegrationStep`

GetNextStep returns the NextStep field if non-nil, zero value otherwise.

### GetNextStepOk

`func (o *CreateDirectProviderSessionResponse) GetNextStepOk() (*IntegrationStep, bool)`

GetNextStepOk returns a tuple with the NextStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextStep

`func (o *CreateDirectProviderSessionResponse) SetNextStep(v IntegrationStep)`

SetNextStep sets NextStep field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


