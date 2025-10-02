# AppleWalletInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExchangeMechanism** | Pointer to [**NullableMdlExchangeMechanism**](MdlExchangeMechanism.md) | The exchange mechanism to use for this Apple Wallet verification.              Use &#x60;DigitalCredentialsApi&#x60; for Digital Credentials API on web, or &#x60;NativeApp&#x60; for a native iOS app. | [optional] 

## Methods

### NewAppleWalletInput

`func NewAppleWalletInput() *AppleWalletInput`

NewAppleWalletInput instantiates a new AppleWalletInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppleWalletInputWithDefaults

`func NewAppleWalletInputWithDefaults() *AppleWalletInput`

NewAppleWalletInputWithDefaults instantiates a new AppleWalletInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExchangeMechanism

`func (o *AppleWalletInput) GetExchangeMechanism() MdlExchangeMechanism`

GetExchangeMechanism returns the ExchangeMechanism field if non-nil, zero value otherwise.

### GetExchangeMechanismOk

`func (o *AppleWalletInput) GetExchangeMechanismOk() (*MdlExchangeMechanism, bool)`

GetExchangeMechanismOk returns a tuple with the ExchangeMechanism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeMechanism

`func (o *AppleWalletInput) SetExchangeMechanism(v MdlExchangeMechanism)`

SetExchangeMechanism sets ExchangeMechanism field to given value.

### HasExchangeMechanism

`func (o *AppleWalletInput) HasExchangeMechanism() bool`

HasExchangeMechanism returns a boolean if a field has been set.

### SetExchangeMechanismNil

`func (o *AppleWalletInput) SetExchangeMechanismNil(b bool)`

 SetExchangeMechanismNil sets the value for ExchangeMechanism to be an explicit nil

### UnsetExchangeMechanism
`func (o *AppleWalletInput) UnsetExchangeMechanism()`

UnsetExchangeMechanism ensures that no value is present for ExchangeMechanism, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


