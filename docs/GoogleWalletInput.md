# GoogleWalletInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExchangeMechanism** | Pointer to [**NullableMdlExchangeMechanism**](MdlExchangeMechanism.md) | The exchange mechanism to use for this Google Wallet verification.              Use &#x60;DigitalCredentialsApi&#x60; for Digital Credentials API on web, or &#x60;NativeApp&#x60; for a native Android app. | [optional] 

## Methods

### NewGoogleWalletInput

`func NewGoogleWalletInput() *GoogleWalletInput`

NewGoogleWalletInput instantiates a new GoogleWalletInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoogleWalletInputWithDefaults

`func NewGoogleWalletInputWithDefaults() *GoogleWalletInput`

NewGoogleWalletInputWithDefaults instantiates a new GoogleWalletInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExchangeMechanism

`func (o *GoogleWalletInput) GetExchangeMechanism() MdlExchangeMechanism`

GetExchangeMechanism returns the ExchangeMechanism field if non-nil, zero value otherwise.

### GetExchangeMechanismOk

`func (o *GoogleWalletInput) GetExchangeMechanismOk() (*MdlExchangeMechanism, bool)`

GetExchangeMechanismOk returns a tuple with the ExchangeMechanism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeMechanism

`func (o *GoogleWalletInput) SetExchangeMechanism(v MdlExchangeMechanism)`

SetExchangeMechanism sets ExchangeMechanism field to given value.

### HasExchangeMechanism

`func (o *GoogleWalletInput) HasExchangeMechanism() bool`

HasExchangeMechanism returns a boolean if a field has been set.

### SetExchangeMechanismNil

`func (o *GoogleWalletInput) SetExchangeMechanismNil(b bool)`

 SetExchangeMechanismNil sets the value for ExchangeMechanism to be an explicit nil

### UnsetExchangeMechanism
`func (o *GoogleWalletInput) UnsetExchangeMechanism()`

UnsetExchangeMechanism ensures that no value is present for ExchangeMechanism, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


