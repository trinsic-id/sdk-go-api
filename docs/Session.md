# Session

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**State** | [**SessionState**](SessionState.md) | The state of the session | 
**FailCode** | Pointer to [**SessionFailCode**](SessionFailCode.md) | If the session is in state &#x60;IdvFailed&#x60;, this field contains the reason for failure. | [optional] 
**Verification** | [**Verification**](Verification.md) | The underlying verification for this Session | 
**DisclosedFields** | [**DisclosedFields**](DisclosedFields.md) | The fields that were requested to be disclosed when the Session was created | 
**Created** | **int64** | The unix timestamp, in seconds, when this session was created | 
**Updated** | **int64** | The unix timestamp, in seconds, when this session&#39;s state last changed | 

## Methods

### NewSession

`func NewSession(id string, state SessionState, verification Verification, disclosedFields DisclosedFields, created int64, updated int64, ) *Session`

NewSession instantiates a new Session object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionWithDefaults

`func NewSessionWithDefaults() *Session`

NewSessionWithDefaults instantiates a new Session object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Session) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Session) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Session) SetId(v string)`

SetId sets Id field to given value.


### GetState

`func (o *Session) GetState() SessionState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Session) GetStateOk() (*SessionState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Session) SetState(v SessionState)`

SetState sets State field to given value.


### GetFailCode

`func (o *Session) GetFailCode() SessionFailCode`

GetFailCode returns the FailCode field if non-nil, zero value otherwise.

### GetFailCodeOk

`func (o *Session) GetFailCodeOk() (*SessionFailCode, bool)`

GetFailCodeOk returns a tuple with the FailCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailCode

`func (o *Session) SetFailCode(v SessionFailCode)`

SetFailCode sets FailCode field to given value.

### HasFailCode

`func (o *Session) HasFailCode() bool`

HasFailCode returns a boolean if a field has been set.

### GetVerification

`func (o *Session) GetVerification() Verification`

GetVerification returns the Verification field if non-nil, zero value otherwise.

### GetVerificationOk

`func (o *Session) GetVerificationOk() (*Verification, bool)`

GetVerificationOk returns a tuple with the Verification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerification

`func (o *Session) SetVerification(v Verification)`

SetVerification sets Verification field to given value.


### GetDisclosedFields

`func (o *Session) GetDisclosedFields() DisclosedFields`

GetDisclosedFields returns the DisclosedFields field if non-nil, zero value otherwise.

### GetDisclosedFieldsOk

`func (o *Session) GetDisclosedFieldsOk() (*DisclosedFields, bool)`

GetDisclosedFieldsOk returns a tuple with the DisclosedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisclosedFields

`func (o *Session) SetDisclosedFields(v DisclosedFields)`

SetDisclosedFields sets DisclosedFields field to given value.


### GetCreated

`func (o *Session) GetCreated() int64`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Session) GetCreatedOk() (*int64, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Session) SetCreated(v int64)`

SetCreated sets Created field to given value.


### GetUpdated

`func (o *Session) GetUpdated() int64`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *Session) GetUpdatedOk() (*int64, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *Session) SetUpdated(v int64)`

SetUpdated sets Updated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


