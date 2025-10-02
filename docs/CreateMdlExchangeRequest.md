# CreateMdlExchangeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VerificationProfileId** | **string** | The ID of the VerificationProfile to use for this mDL exchange. | 
**Provider** | **string** | The ID of the provider to use for this mDL exchange. | 
**ExchangeMechanism** | [**MdlExchangeMechanism**](MdlExchangeMechanism.md) | The mechanism by which the mDL exchange will occur (web, native SDK, etc.) | 
**DocumentType** | **string** | The document type to request from the wallet.              Typically, this is one of the following values:              - &#x60;org.iso.18013.5.1.mDL&#x60; (Mobile Driver&#39;s License) - &#x60;com.google.wallet.idcard.1&#x60; (Google Wallet ID Pass) | 
**NameSpaces** | **map[string]map[string]bool** | The namespaces and fields to request from the mDL.              This is a nested map / dictionary. The outer dictionary&#39;s keys are namespaces (e.g. \&quot;org.iso.18013.5.1\&quot;). The inner dictionary&#39;s keys are field names within each namespace, with boolean values indicating whether the specified field will be retained post-verification. | 
**DigitalCredentialsApiHost** | Pointer to **NullableString** | If using the &#x60;DigitalCredentialsApi&#x60; exchange mechanism, this is the hostname on which the Digital Credentials API will be called.              For example, if the user is on the page &#x60;https://foo.example.com/verify-mdl&#x60;, the proper value to use is &#x60;foo.example.com&#x60;.              Present for ease of testing only. May be removed as this API is stabilized. | [optional] 
**AndroidNativeAppPackageName** | Pointer to **NullableString** | If using the &#x60;NativeApp&#x60; exchange mechanism with the &#x60;google-wallet&#x60; provider, this is the package name of the Android App which will execute the mDL exchange.              This should be set to the package name of your app.              Present for ease of testing only. May be removed as this API is stabilized. | [optional] 
**AndroidNativeAppSigningCertificateFingerprint** | Pointer to **NullableString** | If using the &#x60;NativeApp&#x60; exchange mechanism with the &#x60;google-wallet&#x60; provider, this is the SHA-256 fingerprint of the signing certificate used to sign the Android App which will execute the mDL exchange.              Present for ease of testing only. May be removed as this API is stabilized. | [optional] 

## Methods

### NewCreateMdlExchangeRequest

`func NewCreateMdlExchangeRequest(verificationProfileId string, provider string, exchangeMechanism MdlExchangeMechanism, documentType string, nameSpaces map[string]map[string]bool, ) *CreateMdlExchangeRequest`

NewCreateMdlExchangeRequest instantiates a new CreateMdlExchangeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMdlExchangeRequestWithDefaults

`func NewCreateMdlExchangeRequestWithDefaults() *CreateMdlExchangeRequest`

NewCreateMdlExchangeRequestWithDefaults instantiates a new CreateMdlExchangeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVerificationProfileId

`func (o *CreateMdlExchangeRequest) GetVerificationProfileId() string`

GetVerificationProfileId returns the VerificationProfileId field if non-nil, zero value otherwise.

### GetVerificationProfileIdOk

`func (o *CreateMdlExchangeRequest) GetVerificationProfileIdOk() (*string, bool)`

GetVerificationProfileIdOk returns a tuple with the VerificationProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationProfileId

`func (o *CreateMdlExchangeRequest) SetVerificationProfileId(v string)`

SetVerificationProfileId sets VerificationProfileId field to given value.


### GetProvider

`func (o *CreateMdlExchangeRequest) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CreateMdlExchangeRequest) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CreateMdlExchangeRequest) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetExchangeMechanism

`func (o *CreateMdlExchangeRequest) GetExchangeMechanism() MdlExchangeMechanism`

GetExchangeMechanism returns the ExchangeMechanism field if non-nil, zero value otherwise.

### GetExchangeMechanismOk

`func (o *CreateMdlExchangeRequest) GetExchangeMechanismOk() (*MdlExchangeMechanism, bool)`

GetExchangeMechanismOk returns a tuple with the ExchangeMechanism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeMechanism

`func (o *CreateMdlExchangeRequest) SetExchangeMechanism(v MdlExchangeMechanism)`

SetExchangeMechanism sets ExchangeMechanism field to given value.


### GetDocumentType

`func (o *CreateMdlExchangeRequest) GetDocumentType() string`

GetDocumentType returns the DocumentType field if non-nil, zero value otherwise.

### GetDocumentTypeOk

`func (o *CreateMdlExchangeRequest) GetDocumentTypeOk() (*string, bool)`

GetDocumentTypeOk returns a tuple with the DocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentType

`func (o *CreateMdlExchangeRequest) SetDocumentType(v string)`

SetDocumentType sets DocumentType field to given value.


### GetNameSpaces

`func (o *CreateMdlExchangeRequest) GetNameSpaces() map[string]map[string]bool`

GetNameSpaces returns the NameSpaces field if non-nil, zero value otherwise.

### GetNameSpacesOk

`func (o *CreateMdlExchangeRequest) GetNameSpacesOk() (*map[string]map[string]bool, bool)`

GetNameSpacesOk returns a tuple with the NameSpaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameSpaces

`func (o *CreateMdlExchangeRequest) SetNameSpaces(v map[string]map[string]bool)`

SetNameSpaces sets NameSpaces field to given value.


### GetDigitalCredentialsApiHost

`func (o *CreateMdlExchangeRequest) GetDigitalCredentialsApiHost() string`

GetDigitalCredentialsApiHost returns the DigitalCredentialsApiHost field if non-nil, zero value otherwise.

### GetDigitalCredentialsApiHostOk

`func (o *CreateMdlExchangeRequest) GetDigitalCredentialsApiHostOk() (*string, bool)`

GetDigitalCredentialsApiHostOk returns a tuple with the DigitalCredentialsApiHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigitalCredentialsApiHost

`func (o *CreateMdlExchangeRequest) SetDigitalCredentialsApiHost(v string)`

SetDigitalCredentialsApiHost sets DigitalCredentialsApiHost field to given value.

### HasDigitalCredentialsApiHost

`func (o *CreateMdlExchangeRequest) HasDigitalCredentialsApiHost() bool`

HasDigitalCredentialsApiHost returns a boolean if a field has been set.

### SetDigitalCredentialsApiHostNil

`func (o *CreateMdlExchangeRequest) SetDigitalCredentialsApiHostNil(b bool)`

 SetDigitalCredentialsApiHostNil sets the value for DigitalCredentialsApiHost to be an explicit nil

### UnsetDigitalCredentialsApiHost
`func (o *CreateMdlExchangeRequest) UnsetDigitalCredentialsApiHost()`

UnsetDigitalCredentialsApiHost ensures that no value is present for DigitalCredentialsApiHost, not even an explicit nil
### GetAndroidNativeAppPackageName

`func (o *CreateMdlExchangeRequest) GetAndroidNativeAppPackageName() string`

GetAndroidNativeAppPackageName returns the AndroidNativeAppPackageName field if non-nil, zero value otherwise.

### GetAndroidNativeAppPackageNameOk

`func (o *CreateMdlExchangeRequest) GetAndroidNativeAppPackageNameOk() (*string, bool)`

GetAndroidNativeAppPackageNameOk returns a tuple with the AndroidNativeAppPackageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAndroidNativeAppPackageName

`func (o *CreateMdlExchangeRequest) SetAndroidNativeAppPackageName(v string)`

SetAndroidNativeAppPackageName sets AndroidNativeAppPackageName field to given value.

### HasAndroidNativeAppPackageName

`func (o *CreateMdlExchangeRequest) HasAndroidNativeAppPackageName() bool`

HasAndroidNativeAppPackageName returns a boolean if a field has been set.

### SetAndroidNativeAppPackageNameNil

`func (o *CreateMdlExchangeRequest) SetAndroidNativeAppPackageNameNil(b bool)`

 SetAndroidNativeAppPackageNameNil sets the value for AndroidNativeAppPackageName to be an explicit nil

### UnsetAndroidNativeAppPackageName
`func (o *CreateMdlExchangeRequest) UnsetAndroidNativeAppPackageName()`

UnsetAndroidNativeAppPackageName ensures that no value is present for AndroidNativeAppPackageName, not even an explicit nil
### GetAndroidNativeAppSigningCertificateFingerprint

`func (o *CreateMdlExchangeRequest) GetAndroidNativeAppSigningCertificateFingerprint() string`

GetAndroidNativeAppSigningCertificateFingerprint returns the AndroidNativeAppSigningCertificateFingerprint field if non-nil, zero value otherwise.

### GetAndroidNativeAppSigningCertificateFingerprintOk

`func (o *CreateMdlExchangeRequest) GetAndroidNativeAppSigningCertificateFingerprintOk() (*string, bool)`

GetAndroidNativeAppSigningCertificateFingerprintOk returns a tuple with the AndroidNativeAppSigningCertificateFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAndroidNativeAppSigningCertificateFingerprint

`func (o *CreateMdlExchangeRequest) SetAndroidNativeAppSigningCertificateFingerprint(v string)`

SetAndroidNativeAppSigningCertificateFingerprint sets AndroidNativeAppSigningCertificateFingerprint field to given value.

### HasAndroidNativeAppSigningCertificateFingerprint

`func (o *CreateMdlExchangeRequest) HasAndroidNativeAppSigningCertificateFingerprint() bool`

HasAndroidNativeAppSigningCertificateFingerprint returns a boolean if a field has been set.

### SetAndroidNativeAppSigningCertificateFingerprintNil

`func (o *CreateMdlExchangeRequest) SetAndroidNativeAppSigningCertificateFingerprintNil(b bool)`

 SetAndroidNativeAppSigningCertificateFingerprintNil sets the value for AndroidNativeAppSigningCertificateFingerprint to be an explicit nil

### UnsetAndroidNativeAppSigningCertificateFingerprint
`func (o *CreateMdlExchangeRequest) UnsetAndroidNativeAppSigningCertificateFingerprint()`

UnsetAndroidNativeAppSigningCertificateFingerprint ensures that no value is present for AndroidNativeAppSigningCertificateFingerprint, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


