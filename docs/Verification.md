# Verification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | Pointer to **string** | The identity provider that was used to perform the Verification, if any | [optional] 
**FailCode** | Pointer to [**VerificationFailCode**](VerificationFailCode.md) | If the Verification is in state &#x60;VerificationFailed&#x60;, this field contains the reason for failure | [optional] 

## Methods

### NewVerification

`func NewVerification() *Verification`

NewVerification instantiates a new Verification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerificationWithDefaults

`func NewVerificationWithDefaults() *Verification`

NewVerificationWithDefaults instantiates a new Verification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *Verification) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *Verification) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *Verification) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *Verification) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetFailCode

`func (o *Verification) GetFailCode() VerificationFailCode`

GetFailCode returns the FailCode field if non-nil, zero value otherwise.

### GetFailCodeOk

`func (o *Verification) GetFailCodeOk() (*VerificationFailCode, bool)`

GetFailCodeOk returns a tuple with the FailCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailCode

`func (o *Verification) SetFailCode(v VerificationFailCode)`

SetFailCode sets FailCode field to given value.

### HasFailCode

`func (o *Verification) HasFailCode() bool`

HasFailCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


