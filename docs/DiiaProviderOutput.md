# DiiaProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subject** | [**DiiaSubjectOutput**](DiiaSubjectOutput.md) | Diia.Signature user data. | 
**Issuer** | [**DiiaIssuerOutput**](DiiaIssuerOutput.md) | Diia.Signature issuer data. | 

## Methods

### NewDiiaProviderOutput

`func NewDiiaProviderOutput(subject DiiaSubjectOutput, issuer DiiaIssuerOutput, ) *DiiaProviderOutput`

NewDiiaProviderOutput instantiates a new DiiaProviderOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiiaProviderOutputWithDefaults

`func NewDiiaProviderOutputWithDefaults() *DiiaProviderOutput`

NewDiiaProviderOutputWithDefaults instantiates a new DiiaProviderOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubject

`func (o *DiiaProviderOutput) GetSubject() DiiaSubjectOutput`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *DiiaProviderOutput) GetSubjectOk() (*DiiaSubjectOutput, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *DiiaProviderOutput) SetSubject(v DiiaSubjectOutput)`

SetSubject sets Subject field to given value.


### GetIssuer

`func (o *DiiaProviderOutput) GetIssuer() DiiaIssuerOutput`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *DiiaProviderOutput) GetIssuerOk() (*DiiaIssuerOutput, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *DiiaProviderOutput) SetIssuer(v DiiaIssuerOutput)`

SetIssuer sets Issuer field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


