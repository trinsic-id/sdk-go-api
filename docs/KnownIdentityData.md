# KnownIdentityData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Person** | Pointer to [**KnownPersonData**](KnownPersonData.md) | Known identity data specific to the person being verified | [optional] 

## Methods

### NewKnownIdentityData

`func NewKnownIdentityData() *KnownIdentityData`

NewKnownIdentityData instantiates a new KnownIdentityData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKnownIdentityDataWithDefaults

`func NewKnownIdentityDataWithDefaults() *KnownIdentityData`

NewKnownIdentityDataWithDefaults instantiates a new KnownIdentityData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPerson

`func (o *KnownIdentityData) GetPerson() KnownPersonData`

GetPerson returns the Person field if non-nil, zero value otherwise.

### GetPersonOk

`func (o *KnownIdentityData) GetPersonOk() (*KnownPersonData, bool)`

GetPersonOk returns a tuple with the Person field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerson

`func (o *KnownIdentityData) SetPerson(v KnownPersonData)`

SetPerson sets Person field to given value.

### HasPerson

`func (o *KnownIdentityData) HasPerson() bool`

HasPerson returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


