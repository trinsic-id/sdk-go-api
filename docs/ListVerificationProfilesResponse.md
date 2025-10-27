# ListVerificationProfilesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VerificationProfiles** | [**[]VerificationProfileResponse**](VerificationProfileResponse.md) |  | 
**More** | **bool** | Whether there are additional pages of verification profiles to retrieve | 

## Methods

### NewListVerificationProfilesResponse

`func NewListVerificationProfilesResponse(verificationProfiles []VerificationProfileResponse, more bool, ) *ListVerificationProfilesResponse`

NewListVerificationProfilesResponse instantiates a new ListVerificationProfilesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListVerificationProfilesResponseWithDefaults

`func NewListVerificationProfilesResponseWithDefaults() *ListVerificationProfilesResponse`

NewListVerificationProfilesResponseWithDefaults instantiates a new ListVerificationProfilesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVerificationProfiles

`func (o *ListVerificationProfilesResponse) GetVerificationProfiles() []VerificationProfileResponse`

GetVerificationProfiles returns the VerificationProfiles field if non-nil, zero value otherwise.

### GetVerificationProfilesOk

`func (o *ListVerificationProfilesResponse) GetVerificationProfilesOk() (*[]VerificationProfileResponse, bool)`

GetVerificationProfilesOk returns a tuple with the VerificationProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationProfiles

`func (o *ListVerificationProfilesResponse) SetVerificationProfiles(v []VerificationProfileResponse)`

SetVerificationProfiles sets VerificationProfiles field to given value.


### GetMore

`func (o *ListVerificationProfilesResponse) GetMore() bool`

GetMore returns the More field if non-nil, zero value otherwise.

### GetMoreOk

`func (o *ListVerificationProfilesResponse) GetMoreOk() (*bool, bool)`

GetMoreOk returns a tuple with the More field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMore

`func (o *ListVerificationProfilesResponse) SetMore(v bool)`

SetMore sets More field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


