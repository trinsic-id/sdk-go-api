# ResultCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | [**ResultCollectionMethod**](ResultCollectionMethod.md) | The method by which the results of the Acceptance Session should be collected. | 
**ResultsAccessKey** | **string** | The &#x60;resultsAccessKey&#x60; for the Acceptance Session.              This is an encrypted payload which contains the decryption key necessary to access the Session&#39;s Data Vault.              Save this securely in your systems; it must be passed back with any API call which requires access to the Session&#39;s Data Vault.              Trinsic cannot access a Session&#39;s Data Vault without this key. | 

## Methods

### NewResultCollection

`func NewResultCollection(method ResultCollectionMethod, resultsAccessKey string, ) *ResultCollection`

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


