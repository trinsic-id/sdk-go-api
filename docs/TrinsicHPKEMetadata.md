# TrinsicHPKEMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | **int32** | The HPKE mode in use. Trinsic only currently supports the Basic mode.              There are four modes defined in the HPKE specification:              - Basic (0): Sender is not authenticated - PSK (1): Sender is authenticated by proving possession of a Pre-Shared Key (byte array) which the recipient also knows. - Auth (2): Sender is authenticated by proving possession of a Private Key, for which the recipient knows the corresponding Public Key. - AuthPSK (3): Sender is authenticated by proving possession of both a Pre-Shared Key and a Private Key. | 
**Kem** | **int32** | The Key Encapsulation Mechanism (KEM) used for the HPKE handshake.              Trinsic currently only supports the DHKEM(P256, HKDF-SHA256) KEM, meaning that only P-256 keys are supported. | 
**Kdf** | **int32** | The Key Derivation Function (KDF) used for the HPKE handshake.              Trinsic currently only supports HKDF-SHA256. | 
**Aead** | **int32** | The symmetric encryption cipher used for ciphertext encryption and decryption.              Trinsic currently only supports AES-128-GCM. | 
**InfoHash** | **string** | SHA-256 hash of the &#x60;info&#x60; parameter used during HPKE encryption.              This is not used in HPKE directly.              This is provided as a helper for debugging and development purposes, to verify that the same &#x60;info&#x60; hash is being used by both Trinsic and your integration. | 
**RecipientPublicKeyHash** | **string** | SHA-256 hash of the recipient&#39;s P-256 public key in raw uncompressed point encoding ([0x04][X][Y]).              This is not used in HPKE directly.              Use this to verify that an HPKE message was encrypted against your public key, and/or to identify which of your keys a message was encrypted against, before proceeding with decryption.              This and &#x60;recipientPublicKeyJwkThumbprint&#x60; are both effectively hashes of the recipient&#39;s (your) public key. Either is appropriate to use for the same purpose. | 
**RecipientPublicKeyJwkThumbprint** | **string** | JWK Thumbprint of the recipient&#39;s P-256 public key, in base64url encoding.              This is not used in HPKE directly.              Use this to verify that an HPKE message was encrypted against your public key, and/or to identify which of your keys a message was encrypted against, before proceeding with decryption.              This and &#x60;recipientPublicKeyHash&#x60; are both effectively hashes of the recipient&#39;s (your) public key. Either is appropriate to use for the same purpose. | 

## Methods

### Named construction

Use a keyed struct literal (`TrinsicHPKEMetadata{FieldName: value}`), or call
`NewTrinsicHPKEMetadataWithDefaults()` and set fields by name. Positional field
factories and unkeyed struct literals are not supported.

### NewTrinsicHPKEMetadataWithDefaults

`func NewTrinsicHPKEMetadataWithDefaults() *TrinsicHPKEMetadata`

NewTrinsicHPKEMetadataWithDefaults instantiates a new TrinsicHPKEMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *TrinsicHPKEMetadata) GetMode() int32`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *TrinsicHPKEMetadata) GetModeOk() (*int32, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *TrinsicHPKEMetadata) SetMode(v int32)`

SetMode sets Mode field to given value.


### GetKem

`func (o *TrinsicHPKEMetadata) GetKem() int32`

GetKem returns the Kem field if non-nil, zero value otherwise.

### GetKemOk

`func (o *TrinsicHPKEMetadata) GetKemOk() (*int32, bool)`

GetKemOk returns a tuple with the Kem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKem

`func (o *TrinsicHPKEMetadata) SetKem(v int32)`

SetKem sets Kem field to given value.


### GetKdf

`func (o *TrinsicHPKEMetadata) GetKdf() int32`

GetKdf returns the Kdf field if non-nil, zero value otherwise.

### GetKdfOk

`func (o *TrinsicHPKEMetadata) GetKdfOk() (*int32, bool)`

GetKdfOk returns a tuple with the Kdf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKdf

`func (o *TrinsicHPKEMetadata) SetKdf(v int32)`

SetKdf sets Kdf field to given value.


### GetAead

`func (o *TrinsicHPKEMetadata) GetAead() int32`

GetAead returns the Aead field if non-nil, zero value otherwise.

### GetAeadOk

`func (o *TrinsicHPKEMetadata) GetAeadOk() (*int32, bool)`

GetAeadOk returns a tuple with the Aead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAead

`func (o *TrinsicHPKEMetadata) SetAead(v int32)`

SetAead sets Aead field to given value.


### GetInfoHash

`func (o *TrinsicHPKEMetadata) GetInfoHash() string`

GetInfoHash returns the InfoHash field if non-nil, zero value otherwise.

### GetInfoHashOk

`func (o *TrinsicHPKEMetadata) GetInfoHashOk() (*string, bool)`

GetInfoHashOk returns a tuple with the InfoHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfoHash

`func (o *TrinsicHPKEMetadata) SetInfoHash(v string)`

SetInfoHash sets InfoHash field to given value.


### GetRecipientPublicKeyHash

`func (o *TrinsicHPKEMetadata) GetRecipientPublicKeyHash() string`

GetRecipientPublicKeyHash returns the RecipientPublicKeyHash field if non-nil, zero value otherwise.

### GetRecipientPublicKeyHashOk

`func (o *TrinsicHPKEMetadata) GetRecipientPublicKeyHashOk() (*string, bool)`

GetRecipientPublicKeyHashOk returns a tuple with the RecipientPublicKeyHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientPublicKeyHash

`func (o *TrinsicHPKEMetadata) SetRecipientPublicKeyHash(v string)`

SetRecipientPublicKeyHash sets RecipientPublicKeyHash field to given value.


### GetRecipientPublicKeyJwkThumbprint

`func (o *TrinsicHPKEMetadata) GetRecipientPublicKeyJwkThumbprint() string`

GetRecipientPublicKeyJwkThumbprint returns the RecipientPublicKeyJwkThumbprint field if non-nil, zero value otherwise.

### GetRecipientPublicKeyJwkThumbprintOk

`func (o *TrinsicHPKEMetadata) GetRecipientPublicKeyJwkThumbprintOk() (*string, bool)`

GetRecipientPublicKeyJwkThumbprintOk returns a tuple with the RecipientPublicKeyJwkThumbprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientPublicKeyJwkThumbprint

`func (o *TrinsicHPKEMetadata) SetRecipientPublicKeyJwkThumbprint(v string)`

SetRecipientPublicKeyJwkThumbprint sets RecipientPublicKeyJwkThumbprint field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


