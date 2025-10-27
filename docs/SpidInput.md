# SpidInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubProviderId** | Pointer to **NullableString** | The ID of the specific IDP to invoke within SPID.              If not specified, the user will be prompted to select an IDP. | [optional] 
**BillingTrackingSecret** | Pointer to **NullableString** | Only applicable if period-based billing is enabled for your Verification Profile. Contact Trinsic to enable this.              A secret UTF-8 string between 32 and 64 characters in length, used to enable privacy-preserving tracking of unique user verifications during a billing period.              WARNING: This value must NOT change during the course of a billing period for a given Verification Profile, or double-billing may occur. If multiple Verification Profiles are configured to use the same Trinsic-managed SPID Service Provider, the same Billing Tracking Secret must be provided across all such Verification Profiles. | [optional] 

## Methods

### NewSpidInput

`func NewSpidInput() *SpidInput`

NewSpidInput instantiates a new SpidInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpidInputWithDefaults

`func NewSpidInputWithDefaults() *SpidInput`

NewSpidInputWithDefaults instantiates a new SpidInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubProviderId

`func (o *SpidInput) GetSubProviderId() string`

GetSubProviderId returns the SubProviderId field if non-nil, zero value otherwise.

### GetSubProviderIdOk

`func (o *SpidInput) GetSubProviderIdOk() (*string, bool)`

GetSubProviderIdOk returns a tuple with the SubProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubProviderId

`func (o *SpidInput) SetSubProviderId(v string)`

SetSubProviderId sets SubProviderId field to given value.

### HasSubProviderId

`func (o *SpidInput) HasSubProviderId() bool`

HasSubProviderId returns a boolean if a field has been set.

### SetSubProviderIdNil

`func (o *SpidInput) SetSubProviderIdNil(b bool)`

 SetSubProviderIdNil sets the value for SubProviderId to be an explicit nil

### UnsetSubProviderId
`func (o *SpidInput) UnsetSubProviderId()`

UnsetSubProviderId ensures that no value is present for SubProviderId, not even an explicit nil
### GetBillingTrackingSecret

`func (o *SpidInput) GetBillingTrackingSecret() string`

GetBillingTrackingSecret returns the BillingTrackingSecret field if non-nil, zero value otherwise.

### GetBillingTrackingSecretOk

`func (o *SpidInput) GetBillingTrackingSecretOk() (*string, bool)`

GetBillingTrackingSecretOk returns a tuple with the BillingTrackingSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingTrackingSecret

`func (o *SpidInput) SetBillingTrackingSecret(v string)`

SetBillingTrackingSecret sets BillingTrackingSecret field to given value.

### HasBillingTrackingSecret

`func (o *SpidInput) HasBillingTrackingSecret() bool`

HasBillingTrackingSecret returns a boolean if a field has been set.

### SetBillingTrackingSecretNil

`func (o *SpidInput) SetBillingTrackingSecretNil(b bool)`

 SetBillingTrackingSecretNil sets the value for BillingTrackingSecret to be an explicit nil

### UnsetBillingTrackingSecret
`func (o *SpidInput) UnsetBillingTrackingSecret()`

UnsetBillingTrackingSecret ensures that no value is present for BillingTrackingSecret, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


