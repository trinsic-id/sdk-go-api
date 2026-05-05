# SignzyIndiaAadhaarFetchProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocumentType** | Pointer to **NullableString** | The document type from which the identity data was retrieved. | [optional] 
**Timestamp** | Pointer to **NullableTime** | The timestamp when the signed document was generated and verified. | [optional] 
**TimeToLive** | Pointer to **NullableTime** | The validity expiration timestamp for the verification document. | [optional] 
**AadhaarNumberLastFour** | Pointer to **NullableString** | The Aadhaar number (UID) value for the individual.              This is only the last four digits of the Aadhaar number. | [optional] 
**Claims** | Pointer to [**NullableAadhaarClaims**](AadhaarClaims.md) | The claims extracted from the Aadhaar document. | [optional] 
**LocalizedClaims** | Pointer to [**NullableAadhaarLocalizedClaims**](AadhaarLocalizedClaims.md) | The localized claims extracted from the Aadhaar document. | [optional] 
**DocumentSignatureValidated** | **bool** | Whether our own validation of the Aadhaar document signature and certificate chain succeeded.              When the signed document (e.g. Digilocker XML) is available, we validate it using the standard CCA/SafeScrypt chain. When the document is not returned, the signature cannot be validated and this is false. Some providers (e.g. Signzy) also supply a separate DSC validation indicator in the webhook payload; that is independent of this flag, which reflects only our validation. | 
**DigilockerId** | Pointer to **NullableString** | DigiLocker&#39;s 36-character stable account identifier.              This identifier is deemed safe to use to reference the individual. *Note, the format is not guaranteed to be a UUID. | [optional] 
**IssuerId** | Pointer to **NullableString** | The identifier for the issuer of the DigiLocker document. | [optional] 
**Issuer** | Pointer to **NullableString** | Issuer name for the DigiLocker document. | [optional] 
**MobilePhone** | Pointer to **NullableString** | The individual&#39;s mobile phone number from DigiLocker&#39;s account details. | [optional] 
**Scope** | **[]string** | DigiLocker consent scopes that the individual actually consented to for this session. Included in provider output so customers can verify what was granted: the customer does not control these—the individual chooses scopes in the DigiLocker consent UI, and there is no way to pre-select or enforce them. This is a common source of error (e.g. the individual skips a scope), so surfacing the granted scopes lets customers confirm the session had the expected consent. | 

## Methods

### NewSignzyIndiaAadhaarFetchProviderOutput

`func NewSignzyIndiaAadhaarFetchProviderOutput(documentSignatureValidated bool, scope []string, ) *SignzyIndiaAadhaarFetchProviderOutput`

NewSignzyIndiaAadhaarFetchProviderOutput instantiates a new SignzyIndiaAadhaarFetchProviderOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignzyIndiaAadhaarFetchProviderOutputWithDefaults

`func NewSignzyIndiaAadhaarFetchProviderOutputWithDefaults() *SignzyIndiaAadhaarFetchProviderOutput`

NewSignzyIndiaAadhaarFetchProviderOutputWithDefaults instantiates a new SignzyIndiaAadhaarFetchProviderOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocumentType

`func (o *SignzyIndiaAadhaarFetchProviderOutput) GetDocumentType() string`

GetDocumentType returns the DocumentType field if non-nil, zero value otherwise.

### GetDocumentTypeOk

`func (o *SignzyIndiaAadhaarFetchProviderOutput) GetDocumentTypeOk() (*string, bool)`

GetDocumentTypeOk returns a tuple with the DocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentType

`func (o *SignzyIndiaAadhaarFetchProviderOutput) SetDocumentType(v string)`

SetDocumentType sets DocumentType field to given value.

### HasDocumentType

`func (o *SignzyIndiaAadhaarFetchProviderOutput) HasDocumentType() bool`

HasDocumentType returns a boolean if a field has been set.

### SetDocumentTypeNil

`func (o *SignzyIndiaAadhaarFetchProviderOutput) SetDocumentTypeNil(b bool)`

 SetDocumentTypeNil sets the value for DocumentType to be an explicit nil

### UnsetDocumentType
`func (o *SignzyIndiaAadhaarFetchProviderOutput) UnsetDocumentType()`

UnsetDocumentType ensures that no value is present for DocumentType, not even an explicit nil
### GetTimestamp

`func (o *SignzyIndiaAadhaarFetchProviderOutput) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SignzyIndiaAadhaarFetchProviderOutput) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SignzyIndiaAadhaarFetchProviderOutput) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *SignzyIndiaAadhaarFetchProviderOutput) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *SignzyIndiaAadhaarFetchProviderOutput) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *SignzyIndiaAadhaarFetchProviderOutput) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetTimeToLive

`func (o *SignzyIndiaAadhaarFetchProviderOutput) GetTimeToLive() time.Time`

GetTimeToLive returns the TimeToLive field if non-nil, zero value otherwise.

### GetTimeToLiveOk

`func (o *SignzyIndiaAadhaarFetchProviderOutput) GetTimeToLiveOk() (*time.Time, bool)`

GetTimeToLiveOk returns a tuple with the TimeToLive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeToLive

`func (o *SignzyIndiaAadhaarFetchProviderOutput) SetTimeToLive(v time.Time)`

SetTimeToLive sets TimeToLive field to given value.

### HasTimeToLive

`func (o *SignzyIndiaAadhaarFetchProviderOutput) HasTimeToLive() bool`

HasTimeToLive returns a boolean if a field has been set.

### SetTimeToLiveNil

`func (o *SignzyIndiaAadhaarFetchProviderOutput) SetTimeToLiveNil(b bool)`

 SetTimeToLiveNil sets the value for TimeToLive to be an explicit nil

### UnsetTimeToLive
`func (o *SignzyIndiaAadhaarFetchProviderOutput) UnsetTimeToLive()`

UnsetTimeToLive ensures that no value is present for TimeToLive, not even an explicit nil
### GetAadhaarNumberLastFour

`func (o *SignzyIndiaAadhaarFetchProviderOutput) GetAadhaarNumberLastFour() string`

GetAadhaarNumberLastFour returns the AadhaarNumberLastFour field if non-nil, zero value otherwise.

### GetAadhaarNumberLastFourOk

`func (o *SignzyIndiaAadhaarFetchProviderOutput) GetAadhaarNumberLastFourOk() (*string, bool)`

GetAadhaarNumberLastFourOk returns a tuple with the AadhaarNumberLastFour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAadhaarNumberLastFour

`func (o *SignzyIndiaAadhaarFetchProviderOutput) SetAadhaarNumberLastFour(v string)`

SetAadhaarNumberLastFour sets AadhaarNumberLastFour field to given value.

### HasAadhaarNumberLastFour

`func (o *SignzyIndiaAadhaarFetchProviderOutput) HasAadhaarNumberLastFour() bool`

HasAadhaarNumberLastFour returns a boolean if a field has been set.

### SetAadhaarNumberLastFourNil

`func (o *SignzyIndiaAadhaarFetchProviderOutput) SetAadhaarNumberLastFourNil(b bool)`

 SetAadhaarNumberLastFourNil sets the value for AadhaarNumberLastFour to be an explicit nil

### UnsetAadhaarNumberLastFour
`func (o *SignzyIndiaAadhaarFetchProviderOutput) UnsetAadhaarNumberLastFour()`

UnsetAadhaarNumberLastFour ensures that no value is present for AadhaarNumberLastFour, not even an explicit nil
### GetClaims

`func (o *SignzyIndiaAadhaarFetchProviderOutput) GetClaims() AadhaarClaims`

GetClaims returns the Claims field if non-nil, zero value otherwise.

### GetClaimsOk

`func (o *SignzyIndiaAadhaarFetchProviderOutput) GetClaimsOk() (*AadhaarClaims, bool)`

GetClaimsOk returns a tuple with the Claims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaims

`func (o *SignzyIndiaAadhaarFetchProviderOutput) SetClaims(v AadhaarClaims)`

SetClaims sets Claims field to given value.

### HasClaims

`func (o *SignzyIndiaAadhaarFetchProviderOutput) HasClaims() bool`

HasClaims returns a boolean if a field has been set.

### SetClaimsNil

`func (o *SignzyIndiaAadhaarFetchProviderOutput) SetClaimsNil(b bool)`

 SetClaimsNil sets the value for Claims to be an explicit nil

### UnsetClaims
`func (o *SignzyIndiaAadhaarFetchProviderOutput) UnsetClaims()`

UnsetClaims ensures that no value is present for Claims, not even an explicit nil
### GetLocalizedClaims

`func (o *SignzyIndiaAadhaarFetchProviderOutput) GetLocalizedClaims() AadhaarLocalizedClaims`

GetLocalizedClaims returns the LocalizedClaims field if non-nil, zero value otherwise.

### GetLocalizedClaimsOk

`func (o *SignzyIndiaAadhaarFetchProviderOutput) GetLocalizedClaimsOk() (*AadhaarLocalizedClaims, bool)`

GetLocalizedClaimsOk returns a tuple with the LocalizedClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizedClaims

`func (o *SignzyIndiaAadhaarFetchProviderOutput) SetLocalizedClaims(v AadhaarLocalizedClaims)`

SetLocalizedClaims sets LocalizedClaims field to given value.

### HasLocalizedClaims

`func (o *SignzyIndiaAadhaarFetchProviderOutput) HasLocalizedClaims() bool`

HasLocalizedClaims returns a boolean if a field has been set.

### SetLocalizedClaimsNil

`func (o *SignzyIndiaAadhaarFetchProviderOutput) SetLocalizedClaimsNil(b bool)`

 SetLocalizedClaimsNil sets the value for LocalizedClaims to be an explicit nil

### UnsetLocalizedClaims
`func (o *SignzyIndiaAadhaarFetchProviderOutput) UnsetLocalizedClaims()`

UnsetLocalizedClaims ensures that no value is present for LocalizedClaims, not even an explicit nil
### GetDocumentSignatureValidated

`func (o *SignzyIndiaAadhaarFetchProviderOutput) GetDocumentSignatureValidated() bool`

GetDocumentSignatureValidated returns the DocumentSignatureValidated field if non-nil, zero value otherwise.

### GetDocumentSignatureValidatedOk

`func (o *SignzyIndiaAadhaarFetchProviderOutput) GetDocumentSignatureValidatedOk() (*bool, bool)`

GetDocumentSignatureValidatedOk returns a tuple with the DocumentSignatureValidated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentSignatureValidated

`func (o *SignzyIndiaAadhaarFetchProviderOutput) SetDocumentSignatureValidated(v bool)`

SetDocumentSignatureValidated sets DocumentSignatureValidated field to given value.


### GetDigilockerId

`func (o *SignzyIndiaAadhaarFetchProviderOutput) GetDigilockerId() string`

GetDigilockerId returns the DigilockerId field if non-nil, zero value otherwise.

### GetDigilockerIdOk

`func (o *SignzyIndiaAadhaarFetchProviderOutput) GetDigilockerIdOk() (*string, bool)`

GetDigilockerIdOk returns a tuple with the DigilockerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigilockerId

`func (o *SignzyIndiaAadhaarFetchProviderOutput) SetDigilockerId(v string)`

SetDigilockerId sets DigilockerId field to given value.

### HasDigilockerId

`func (o *SignzyIndiaAadhaarFetchProviderOutput) HasDigilockerId() bool`

HasDigilockerId returns a boolean if a field has been set.

### SetDigilockerIdNil

`func (o *SignzyIndiaAadhaarFetchProviderOutput) SetDigilockerIdNil(b bool)`

 SetDigilockerIdNil sets the value for DigilockerId to be an explicit nil

### UnsetDigilockerId
`func (o *SignzyIndiaAadhaarFetchProviderOutput) UnsetDigilockerId()`

UnsetDigilockerId ensures that no value is present for DigilockerId, not even an explicit nil
### GetIssuerId

`func (o *SignzyIndiaAadhaarFetchProviderOutput) GetIssuerId() string`

GetIssuerId returns the IssuerId field if non-nil, zero value otherwise.

### GetIssuerIdOk

`func (o *SignzyIndiaAadhaarFetchProviderOutput) GetIssuerIdOk() (*string, bool)`

GetIssuerIdOk returns a tuple with the IssuerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerId

`func (o *SignzyIndiaAadhaarFetchProviderOutput) SetIssuerId(v string)`

SetIssuerId sets IssuerId field to given value.

### HasIssuerId

`func (o *SignzyIndiaAadhaarFetchProviderOutput) HasIssuerId() bool`

HasIssuerId returns a boolean if a field has been set.

### SetIssuerIdNil

`func (o *SignzyIndiaAadhaarFetchProviderOutput) SetIssuerIdNil(b bool)`

 SetIssuerIdNil sets the value for IssuerId to be an explicit nil

### UnsetIssuerId
`func (o *SignzyIndiaAadhaarFetchProviderOutput) UnsetIssuerId()`

UnsetIssuerId ensures that no value is present for IssuerId, not even an explicit nil
### GetIssuer

`func (o *SignzyIndiaAadhaarFetchProviderOutput) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *SignzyIndiaAadhaarFetchProviderOutput) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *SignzyIndiaAadhaarFetchProviderOutput) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *SignzyIndiaAadhaarFetchProviderOutput) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### SetIssuerNil

`func (o *SignzyIndiaAadhaarFetchProviderOutput) SetIssuerNil(b bool)`

 SetIssuerNil sets the value for Issuer to be an explicit nil

### UnsetIssuer
`func (o *SignzyIndiaAadhaarFetchProviderOutput) UnsetIssuer()`

UnsetIssuer ensures that no value is present for Issuer, not even an explicit nil
### GetMobilePhone

`func (o *SignzyIndiaAadhaarFetchProviderOutput) GetMobilePhone() string`

GetMobilePhone returns the MobilePhone field if non-nil, zero value otherwise.

### GetMobilePhoneOk

`func (o *SignzyIndiaAadhaarFetchProviderOutput) GetMobilePhoneOk() (*string, bool)`

GetMobilePhoneOk returns a tuple with the MobilePhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobilePhone

`func (o *SignzyIndiaAadhaarFetchProviderOutput) SetMobilePhone(v string)`

SetMobilePhone sets MobilePhone field to given value.

### HasMobilePhone

`func (o *SignzyIndiaAadhaarFetchProviderOutput) HasMobilePhone() bool`

HasMobilePhone returns a boolean if a field has been set.

### SetMobilePhoneNil

`func (o *SignzyIndiaAadhaarFetchProviderOutput) SetMobilePhoneNil(b bool)`

 SetMobilePhoneNil sets the value for MobilePhone to be an explicit nil

### UnsetMobilePhone
`func (o *SignzyIndiaAadhaarFetchProviderOutput) UnsetMobilePhone()`

UnsetMobilePhone ensures that no value is present for MobilePhone, not even an explicit nil
### GetScope

`func (o *SignzyIndiaAadhaarFetchProviderOutput) GetScope() []string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *SignzyIndiaAadhaarFetchProviderOutput) GetScopeOk() (*[]string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *SignzyIndiaAadhaarFetchProviderOutput) SetScope(v []string)`

SetScope sets Scope field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


