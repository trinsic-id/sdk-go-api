# TrinsicHPKEBundle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrinsicHpkeVersion** | **int32** | The version of the Trinsic HPKE Exchange Protocol in use.              Always &#x60;1&#x60; at this time. | 
**Metadata** | [**TrinsicHPKEMetadata**](TrinsicHPKEMetadata.md) | Information about the HPKE parameters in use.              Trinsic currently supports a fixed set of HPKE parameters: - Mode: Basic - KEM: DHKEM(P256, HKDF-SHA256) - KDF: HKDF-SHA256 - AEAD: AES-128-GCM | 
**EphemeralSenderPublicKey** | **string** | Corresponds to the &#x60;enc&#x60; parameter of HPKE.              The raw bytes, in a format unique to the chosen HPKE KEM, of the ephemeral sender public key for this single-shot HPKE exchange.              Because Trinsic only supports DHKEM(P256, HKDF-SHA256), this is always a 65-byte array with the following format: [0x04][XCoord][YCoord]              Where [0x04] is a single byte with value &#x60;4&#x60;, and [XCoord] / [YCoord] are each 32-byte X/Y coordinates of a P-256 public key. | 
**Ciphertext** | **string** | Corresponds to the &#x60;ct&#x60; parameter of HPKE.              The encrypted ciphertext for this HPKE single-shot message.              Because Trinsic only supports AES-128-GCM, this byte array always has an exact format: [CipherText][Tag]              Where [ciphertext] is a variable-length array of bytes, and [tag] is a 16-length byte array. | 

## Methods

### Named construction

Use a keyed struct literal (`TrinsicHPKEBundle{FieldName: value}`), or call
`NewTrinsicHPKEBundleWithDefaults()` and set fields by name. Positional field
factories and unkeyed struct literals are not supported.

### NewTrinsicHPKEBundleWithDefaults

`func NewTrinsicHPKEBundleWithDefaults() *TrinsicHPKEBundle`

NewTrinsicHPKEBundleWithDefaults instantiates a new TrinsicHPKEBundle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrinsicHpkeVersion

`func (o *TrinsicHPKEBundle) GetTrinsicHpkeVersion() int32`

GetTrinsicHpkeVersion returns the TrinsicHpkeVersion field if non-nil, zero value otherwise.

### GetTrinsicHpkeVersionOk

`func (o *TrinsicHPKEBundle) GetTrinsicHpkeVersionOk() (*int32, bool)`

GetTrinsicHpkeVersionOk returns a tuple with the TrinsicHpkeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrinsicHpkeVersion

`func (o *TrinsicHPKEBundle) SetTrinsicHpkeVersion(v int32)`

SetTrinsicHpkeVersion sets TrinsicHpkeVersion field to given value.


### GetMetadata

`func (o *TrinsicHPKEBundle) GetMetadata() TrinsicHPKEMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *TrinsicHPKEBundle) GetMetadataOk() (*TrinsicHPKEMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *TrinsicHPKEBundle) SetMetadata(v TrinsicHPKEMetadata)`

SetMetadata sets Metadata field to given value.


### GetEphemeralSenderPublicKey

`func (o *TrinsicHPKEBundle) GetEphemeralSenderPublicKey() string`

GetEphemeralSenderPublicKey returns the EphemeralSenderPublicKey field if non-nil, zero value otherwise.

### GetEphemeralSenderPublicKeyOk

`func (o *TrinsicHPKEBundle) GetEphemeralSenderPublicKeyOk() (*string, bool)`

GetEphemeralSenderPublicKeyOk returns a tuple with the EphemeralSenderPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEphemeralSenderPublicKey

`func (o *TrinsicHPKEBundle) SetEphemeralSenderPublicKey(v string)`

SetEphemeralSenderPublicKey sets EphemeralSenderPublicKey field to given value.


### GetCiphertext

`func (o *TrinsicHPKEBundle) GetCiphertext() string`

GetCiphertext returns the Ciphertext field if non-nil, zero value otherwise.

### GetCiphertextOk

`func (o *TrinsicHPKEBundle) GetCiphertextOk() (*string, bool)`

GetCiphertextOk returns a tuple with the Ciphertext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiphertext

`func (o *TrinsicHPKEBundle) SetCiphertext(v string)`

SetCiphertext sets Ciphertext field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


