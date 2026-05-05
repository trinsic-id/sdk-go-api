# KenyaNidLookup2Input

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdNumber** | Pointer to **NullableString** | The Kenya National ID Number (Nambari ya Kitambulisho) or Unique Personal Identifier (Maisha Namba).              This is the primary unique identifier for Kenyan citizens in all government systems, issued by the National Registration Bureau (NRB). The format is either 8 digits for National ID or 9 digits for Maisha Namba UPI (the new format since 2023). | [optional] 

## Methods

### NewKenyaNidLookup2Input

`func NewKenyaNidLookup2Input() *KenyaNidLookup2Input`

NewKenyaNidLookup2Input instantiates a new KenyaNidLookup2Input object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKenyaNidLookup2InputWithDefaults

`func NewKenyaNidLookup2InputWithDefaults() *KenyaNidLookup2Input`

NewKenyaNidLookup2InputWithDefaults instantiates a new KenyaNidLookup2Input object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdNumber

`func (o *KenyaNidLookup2Input) GetIdNumber() string`

GetIdNumber returns the IdNumber field if non-nil, zero value otherwise.

### GetIdNumberOk

`func (o *KenyaNidLookup2Input) GetIdNumberOk() (*string, bool)`

GetIdNumberOk returns a tuple with the IdNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdNumber

`func (o *KenyaNidLookup2Input) SetIdNumber(v string)`

SetIdNumber sets IdNumber field to given value.

### HasIdNumber

`func (o *KenyaNidLookup2Input) HasIdNumber() bool`

HasIdNumber returns a boolean if a field has been set.

### SetIdNumberNil

`func (o *KenyaNidLookup2Input) SetIdNumberNil(b bool)`

 SetIdNumberNil sets the value for IdNumber to be an explicit nil

### UnsetIdNumber
`func (o *KenyaNidLookup2Input) UnsetIdNumber()`

UnsetIdNumber ensures that no value is present for IdNumber, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


