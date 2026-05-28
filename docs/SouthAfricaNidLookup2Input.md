# SouthAfricaNidLookup2Input

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdNumber** | Pointer to **NullableString** | The South African National Identity Number (13 digits).              Issued for life by the Department of Home Affairs (DHA) and stored in the HANIS (Home Affairs National Identification System) database. The same number is mandatory for banking, employment, taxation, and voting, and is printed on both the legacy green ID book and the Smart ID Card (rolled out from 2013 onward).              Format: - YYMMDD G(4) C A Z - YYMMDD is the date of birth - G(4) is the gender code (below 5000 female, 5000 or above male) - C is the citizenship indicator (0 citizen, 1 permanent resident) - A is reserved (it had a politically sensitive meaning in the past, but is currently   semantically meaningless) - Z is a Luhn check digit | [optional] 

## Methods

### NewSouthAfricaNidLookup2Input

`func NewSouthAfricaNidLookup2Input() *SouthAfricaNidLookup2Input`

NewSouthAfricaNidLookup2Input instantiates a new SouthAfricaNidLookup2Input object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSouthAfricaNidLookup2InputWithDefaults

`func NewSouthAfricaNidLookup2InputWithDefaults() *SouthAfricaNidLookup2Input`

NewSouthAfricaNidLookup2InputWithDefaults instantiates a new SouthAfricaNidLookup2Input object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdNumber

`func (o *SouthAfricaNidLookup2Input) GetIdNumber() string`

GetIdNumber returns the IdNumber field if non-nil, zero value otherwise.

### GetIdNumberOk

`func (o *SouthAfricaNidLookup2Input) GetIdNumberOk() (*string, bool)`

GetIdNumberOk returns a tuple with the IdNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdNumber

`func (o *SouthAfricaNidLookup2Input) SetIdNumber(v string)`

SetIdNumber sets IdNumber field to given value.

### HasIdNumber

`func (o *SouthAfricaNidLookup2Input) HasIdNumber() bool`

HasIdNumber returns a boolean if a field has been set.

### SetIdNumberNil

`func (o *SouthAfricaNidLookup2Input) SetIdNumberNil(b bool)`

 SetIdNumberNil sets the value for IdNumber to be an explicit nil

### UnsetIdNumber
`func (o *SouthAfricaNidLookup2Input) UnsetIdNumber()`

UnsetIdNumber ensures that no value is present for IdNumber, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


