# ResultCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | [**ResultCollectionMethod**](ResultCollectionMethod.md) | The method by which the results of the Acceptance Session should be collected. | 
**ResultsAccessKey** | Pointer to **NullableString** | If the method is &#x60;PollResult&#x60;, this is the key that should be used to poll for the results. | [optional] 

## Methods

### NewResultCollection

`func NewResultCollection(method ResultCollectionMethod, ) *ResultCollection`

NewResultCollection instantiates a new ResultCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultCollectionWithDefaults

`func NewResultCollectionWithDefaults() *ResultCollection`

NewResultCollectionWithDefaults instantiates a new ResultCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethod

`func (o *ResultCollection) GetMethod() ResultCollectionMethod`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *ResultCollection) GetMethodOk() (*ResultCollectionMethod, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *ResultCollection) SetMethod(v ResultCollectionMethod)`

SetMethod sets Method field to given value.


### GetResultsAccessKey

`func (o *ResultCollection) GetResultsAccessKey() string`

GetResultsAccessKey returns the ResultsAccessKey field if non-nil, zero value otherwise.

### GetResultsAccessKeyOk

`func (o *ResultCollection) GetResultsAccessKeyOk() (*string, bool)`

GetResultsAccessKeyOk returns a tuple with the ResultsAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsAccessKey

`func (o *ResultCollection) SetResultsAccessKey(v string)`

SetResultsAccessKey sets ResultsAccessKey field to given value.

### HasResultsAccessKey

`func (o *ResultCollection) HasResultsAccessKey() bool`

HasResultsAccessKey returns a boolean if a field has been set.

### SetResultsAccessKeyNil

`func (o *ResultCollection) SetResultsAccessKeyNil(b bool)`

 SetResultsAccessKeyNil sets the value for ResultsAccessKey to be an explicit nil

### UnsetResultsAccessKey
`func (o *ResultCollection) UnsetResultsAccessKey()`

UnsetResultsAccessKey ensures that no value is present for ResultsAccessKey, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


