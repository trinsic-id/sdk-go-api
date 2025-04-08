# TrinsicTestDatabaseLookupInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GivenName** | Pointer to **NullableString** | The given name to use for the output of the test Session.              This is required; if not provided, Trinsic&#39;s Fallback UI will be invoked to collect it from the user.              Can be any non-empty value. | [optional] 
**FamilyName** | Pointer to **NullableString** | The family name to use for the output of the test Session.              This is required; if not provided, Trinsic&#39;s Fallback UI will be invoked to collect it from the user.              Can be any non-empty value. | [optional] 
**IdentityCode** | Pointer to **NullableString** | A 6-digit code; must be \&quot;123456\&quot; for the Session to succeed.              This is required; if not provided, Trinsic&#39;s Fallback UI will be invoked to collect it from the user.              Any other value will cause the Session to fail. | [optional] 
**SelfieBase64** | Pointer to **NullableString** | An optional selfie image, base64-encoded.              Will replace the existing test selfie attachment output if provided. | [optional] 

## Methods

### NewTrinsicTestDatabaseLookupInput

`func NewTrinsicTestDatabaseLookupInput() *TrinsicTestDatabaseLookupInput`

NewTrinsicTestDatabaseLookupInput instantiates a new TrinsicTestDatabaseLookupInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrinsicTestDatabaseLookupInputWithDefaults

`func NewTrinsicTestDatabaseLookupInputWithDefaults() *TrinsicTestDatabaseLookupInput`

NewTrinsicTestDatabaseLookupInputWithDefaults instantiates a new TrinsicTestDatabaseLookupInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGivenName

`func (o *TrinsicTestDatabaseLookupInput) GetGivenName() string`

GetGivenName returns the GivenName field if non-nil, zero value otherwise.

### GetGivenNameOk

`func (o *TrinsicTestDatabaseLookupInput) GetGivenNameOk() (*string, bool)`

GetGivenNameOk returns a tuple with the GivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenName

`func (o *TrinsicTestDatabaseLookupInput) SetGivenName(v string)`

SetGivenName sets GivenName field to given value.

### HasGivenName

`func (o *TrinsicTestDatabaseLookupInput) HasGivenName() bool`

HasGivenName returns a boolean if a field has been set.

### SetGivenNameNil

`func (o *TrinsicTestDatabaseLookupInput) SetGivenNameNil(b bool)`

 SetGivenNameNil sets the value for GivenName to be an explicit nil

### UnsetGivenName
`func (o *TrinsicTestDatabaseLookupInput) UnsetGivenName()`

UnsetGivenName ensures that no value is present for GivenName, not even an explicit nil
### GetFamilyName

`func (o *TrinsicTestDatabaseLookupInput) GetFamilyName() string`

GetFamilyName returns the FamilyName field if non-nil, zero value otherwise.

### GetFamilyNameOk

`func (o *TrinsicTestDatabaseLookupInput) GetFamilyNameOk() (*string, bool)`

GetFamilyNameOk returns a tuple with the FamilyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyName

`func (o *TrinsicTestDatabaseLookupInput) SetFamilyName(v string)`

SetFamilyName sets FamilyName field to given value.

### HasFamilyName

`func (o *TrinsicTestDatabaseLookupInput) HasFamilyName() bool`

HasFamilyName returns a boolean if a field has been set.

### SetFamilyNameNil

`func (o *TrinsicTestDatabaseLookupInput) SetFamilyNameNil(b bool)`

 SetFamilyNameNil sets the value for FamilyName to be an explicit nil

### UnsetFamilyName
`func (o *TrinsicTestDatabaseLookupInput) UnsetFamilyName()`

UnsetFamilyName ensures that no value is present for FamilyName, not even an explicit nil
### GetIdentityCode

`func (o *TrinsicTestDatabaseLookupInput) GetIdentityCode() string`

GetIdentityCode returns the IdentityCode field if non-nil, zero value otherwise.

### GetIdentityCodeOk

`func (o *TrinsicTestDatabaseLookupInput) GetIdentityCodeOk() (*string, bool)`

GetIdentityCodeOk returns a tuple with the IdentityCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityCode

`func (o *TrinsicTestDatabaseLookupInput) SetIdentityCode(v string)`

SetIdentityCode sets IdentityCode field to given value.

### HasIdentityCode

`func (o *TrinsicTestDatabaseLookupInput) HasIdentityCode() bool`

HasIdentityCode returns a boolean if a field has been set.

### SetIdentityCodeNil

`func (o *TrinsicTestDatabaseLookupInput) SetIdentityCodeNil(b bool)`

 SetIdentityCodeNil sets the value for IdentityCode to be an explicit nil

### UnsetIdentityCode
`func (o *TrinsicTestDatabaseLookupInput) UnsetIdentityCode()`

UnsetIdentityCode ensures that no value is present for IdentityCode, not even an explicit nil
### GetSelfieBase64

`func (o *TrinsicTestDatabaseLookupInput) GetSelfieBase64() string`

GetSelfieBase64 returns the SelfieBase64 field if non-nil, zero value otherwise.

### GetSelfieBase64Ok

`func (o *TrinsicTestDatabaseLookupInput) GetSelfieBase64Ok() (*string, bool)`

GetSelfieBase64Ok returns a tuple with the SelfieBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfieBase64

`func (o *TrinsicTestDatabaseLookupInput) SetSelfieBase64(v string)`

SetSelfieBase64 sets SelfieBase64 field to given value.

### HasSelfieBase64

`func (o *TrinsicTestDatabaseLookupInput) HasSelfieBase64() bool`

HasSelfieBase64 returns a boolean if a field has been set.

### SetSelfieBase64Nil

`func (o *TrinsicTestDatabaseLookupInput) SetSelfieBase64Nil(b bool)`

 SetSelfieBase64Nil sets the value for SelfieBase64 to be an explicit nil

### UnsetSelfieBase64
`func (o *TrinsicTestDatabaseLookupInput) UnsetSelfieBase64()`

UnsetSelfieBase64 ensures that no value is present for SelfieBase64, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


