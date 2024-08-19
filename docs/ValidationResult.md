# ValidationResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MemberNames** | **[]string** |  | [readonly] 
**ErrorMessage** | Pointer to **string** |  | [optional] 

## Methods

### NewValidationResult

`func NewValidationResult(memberNames []string, ) *ValidationResult`

NewValidationResult instantiates a new ValidationResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidationResultWithDefaults

`func NewValidationResultWithDefaults() *ValidationResult`

NewValidationResultWithDefaults instantiates a new ValidationResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMemberNames

`func (o *ValidationResult) GetMemberNames() []string`

GetMemberNames returns the MemberNames field if non-nil, zero value otherwise.

### GetMemberNamesOk

`func (o *ValidationResult) GetMemberNamesOk() (*[]string, bool)`

GetMemberNamesOk returns a tuple with the MemberNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberNames

`func (o *ValidationResult) SetMemberNames(v []string)`

SetMemberNames sets MemberNames field to given value.


### GetErrorMessage

`func (o *ValidationResult) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ValidationResult) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ValidationResult) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *ValidationResult) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


