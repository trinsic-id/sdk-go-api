# DiiaProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subject** | Pointer to [**NullableDiiaSubjectOutput**](DiiaSubjectOutput.md) | Diia.Signature user data. | [optional] 
**Issuer** | Pointer to [**NullableDiiaIssuerOutput**](DiiaIssuerOutput.md) | Diia.Signature issuer data. | [optional] 

## Methods

### NewDiiaProviderOutput

`func NewDiiaProviderOutput() *DiiaProviderOutput`

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

### HasSubject

`func (o *DiiaProviderOutput) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### SetSubjectNil

`func (o *DiiaProviderOutput) SetSubjectNil(b bool)`

 SetSubjectNil sets the value for Subject to be an explicit nil

### UnsetSubject
`func (o *DiiaProviderOutput) UnsetSubject()`

UnsetSubject ensures that no value is present for Subject, not even an explicit nil
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

### HasIssuer

`func (o *DiiaProviderOutput) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### SetIssuerNil

`func (o *DiiaProviderOutput) SetIssuerNil(b bool)`

 SetIssuerNil sets the value for Issuer to be an explicit nil

### UnsetIssuer
`func (o *DiiaProviderOutput) UnsetIssuer()`

UnsetIssuer ensures that no value is present for Issuer, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


