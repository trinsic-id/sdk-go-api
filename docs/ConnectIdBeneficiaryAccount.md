# ConnectIdBeneficiaryAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BeneficiaryName** | Pointer to **NullableString** | The beneficiary&#39;s name. | [optional] 
**AccountBankStateBranch** | Pointer to **NullableString** | The bank state branch (BSB) which the account belongs to.              This will be a six-digit number. | [optional] 
**AccountNumber** | Pointer to **NullableString** | The account number. | [optional] 
**TrustFramework** | Pointer to **NullableString** | The authority that verified the claim. | [optional] 

## Methods

### NewConnectIdBeneficiaryAccount

`func NewConnectIdBeneficiaryAccount() *ConnectIdBeneficiaryAccount`

NewConnectIdBeneficiaryAccount instantiates a new ConnectIdBeneficiaryAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectIdBeneficiaryAccountWithDefaults

`func NewConnectIdBeneficiaryAccountWithDefaults() *ConnectIdBeneficiaryAccount`

NewConnectIdBeneficiaryAccountWithDefaults instantiates a new ConnectIdBeneficiaryAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBeneficiaryName

`func (o *ConnectIdBeneficiaryAccount) GetBeneficiaryName() string`

GetBeneficiaryName returns the BeneficiaryName field if non-nil, zero value otherwise.

### GetBeneficiaryNameOk

`func (o *ConnectIdBeneficiaryAccount) GetBeneficiaryNameOk() (*string, bool)`

GetBeneficiaryNameOk returns a tuple with the BeneficiaryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiaryName

`func (o *ConnectIdBeneficiaryAccount) SetBeneficiaryName(v string)`

SetBeneficiaryName sets BeneficiaryName field to given value.

### HasBeneficiaryName

`func (o *ConnectIdBeneficiaryAccount) HasBeneficiaryName() bool`

HasBeneficiaryName returns a boolean if a field has been set.

### SetBeneficiaryNameNil

`func (o *ConnectIdBeneficiaryAccount) SetBeneficiaryNameNil(b bool)`

 SetBeneficiaryNameNil sets the value for BeneficiaryName to be an explicit nil

### UnsetBeneficiaryName
`func (o *ConnectIdBeneficiaryAccount) UnsetBeneficiaryName()`

UnsetBeneficiaryName ensures that no value is present for BeneficiaryName, not even an explicit nil
### GetAccountBankStateBranch

`func (o *ConnectIdBeneficiaryAccount) GetAccountBankStateBranch() string`

GetAccountBankStateBranch returns the AccountBankStateBranch field if non-nil, zero value otherwise.

### GetAccountBankStateBranchOk

`func (o *ConnectIdBeneficiaryAccount) GetAccountBankStateBranchOk() (*string, bool)`

GetAccountBankStateBranchOk returns a tuple with the AccountBankStateBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountBankStateBranch

`func (o *ConnectIdBeneficiaryAccount) SetAccountBankStateBranch(v string)`

SetAccountBankStateBranch sets AccountBankStateBranch field to given value.

### HasAccountBankStateBranch

`func (o *ConnectIdBeneficiaryAccount) HasAccountBankStateBranch() bool`

HasAccountBankStateBranch returns a boolean if a field has been set.

### SetAccountBankStateBranchNil

`func (o *ConnectIdBeneficiaryAccount) SetAccountBankStateBranchNil(b bool)`

 SetAccountBankStateBranchNil sets the value for AccountBankStateBranch to be an explicit nil

### UnsetAccountBankStateBranch
`func (o *ConnectIdBeneficiaryAccount) UnsetAccountBankStateBranch()`

UnsetAccountBankStateBranch ensures that no value is present for AccountBankStateBranch, not even an explicit nil
### GetAccountNumber

`func (o *ConnectIdBeneficiaryAccount) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *ConnectIdBeneficiaryAccount) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *ConnectIdBeneficiaryAccount) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *ConnectIdBeneficiaryAccount) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### SetAccountNumberNil

`func (o *ConnectIdBeneficiaryAccount) SetAccountNumberNil(b bool)`

 SetAccountNumberNil sets the value for AccountNumber to be an explicit nil

### UnsetAccountNumber
`func (o *ConnectIdBeneficiaryAccount) UnsetAccountNumber()`

UnsetAccountNumber ensures that no value is present for AccountNumber, not even an explicit nil
### GetTrustFramework

`func (o *ConnectIdBeneficiaryAccount) GetTrustFramework() string`

GetTrustFramework returns the TrustFramework field if non-nil, zero value otherwise.

### GetTrustFrameworkOk

`func (o *ConnectIdBeneficiaryAccount) GetTrustFrameworkOk() (*string, bool)`

GetTrustFrameworkOk returns a tuple with the TrustFramework field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustFramework

`func (o *ConnectIdBeneficiaryAccount) SetTrustFramework(v string)`

SetTrustFramework sets TrustFramework field to given value.

### HasTrustFramework

`func (o *ConnectIdBeneficiaryAccount) HasTrustFramework() bool`

HasTrustFramework returns a boolean if a field has been set.

### SetTrustFrameworkNil

`func (o *ConnectIdBeneficiaryAccount) SetTrustFrameworkNil(b bool)`

 SetTrustFrameworkNil sets the value for TrustFramework to be an explicit nil

### UnsetTrustFramework
`func (o *ConnectIdBeneficiaryAccount) UnsetTrustFramework()`

UnsetTrustFramework ensures that no value is present for TrustFramework, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


