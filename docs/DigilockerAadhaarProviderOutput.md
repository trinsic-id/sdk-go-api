# DigilockerAadhaarProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocumentType** | Pointer to **NullableString** | The document type from which the identity data was retrieved. | [optional] 
**Timestamp** | Pointer to **NullableTime** | The timestamp when the signed document was generated and verified. | [optional] 
**TimeToLive** | Pointer to **NullableTime** | The validity expiration timestamp for the verification document. | [optional] 
**AadhaarNumberLastFour** | Pointer to **NullableString** | The Aadhaar number (UID) value for the individual.              This is only the last four digits of the Aadhaar number. | [optional] 
**Claims** | Pointer to [**NullableAadhaarClaims**](AadhaarClaims.md) | The claims extracted from the Aadhaar document. | [optional] 
**LocalizedClaims** | Pointer to [**NullableAadhaarLocalizedClaims**](AadhaarLocalizedClaims.md) | The localized claims extracted from the Aadhaar document. | [optional] 
**DocumentSignatureValidated** | Pointer to **NullableBool** | Whether our own validation of the Aadhaar document signature and certificate chain succeeded.              When the signed document (e.g. Digilocker XML) is available, we validate it using the standard CCA/SafeScrypt chain. When the document is not returned, the signature cannot be validated and this is false. Some providers (e.g. Signzy) also supply a separate DSC validation indicator in the webhook payload; that is independent of this flag, which reflects only our validation. | [optional] 

## Methods

### NewDigilockerAadhaarProviderOutput

`func NewDigilockerAadhaarProviderOutput() *DigilockerAadhaarProviderOutput`

NewDigilockerAadhaarProviderOutput instantiates a new DigilockerAadhaarProviderOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDigilockerAadhaarProviderOutputWithDefaults

`func NewDigilockerAadhaarProviderOutputWithDefaults() *DigilockerAadhaarProviderOutput`

NewDigilockerAadhaarProviderOutputWithDefaults instantiates a new DigilockerAadhaarProviderOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocumentType

`func (o *DigilockerAadhaarProviderOutput) GetDocumentType() string`

GetDocumentType returns the DocumentType field if non-nil, zero value otherwise.

### GetDocumentTypeOk

`func (o *DigilockerAadhaarProviderOutput) GetDocumentTypeOk() (*string, bool)`

GetDocumentTypeOk returns a tuple with the DocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentType

`func (o *DigilockerAadhaarProviderOutput) SetDocumentType(v string)`

SetDocumentType sets DocumentType field to given value.

### HasDocumentType

`func (o *DigilockerAadhaarProviderOutput) HasDocumentType() bool`

HasDocumentType returns a boolean if a field has been set.

### SetDocumentTypeNil

`func (o *DigilockerAadhaarProviderOutput) SetDocumentTypeNil(b bool)`

 SetDocumentTypeNil sets the value for DocumentType to be an explicit nil

### UnsetDocumentType
`func (o *DigilockerAadhaarProviderOutput) UnsetDocumentType()`

UnsetDocumentType ensures that no value is present for DocumentType, not even an explicit nil
### GetTimestamp

`func (o *DigilockerAadhaarProviderOutput) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *DigilockerAadhaarProviderOutput) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *DigilockerAadhaarProviderOutput) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *DigilockerAadhaarProviderOutput) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *DigilockerAadhaarProviderOutput) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *DigilockerAadhaarProviderOutput) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetTimeToLive

`func (o *DigilockerAadhaarProviderOutput) GetTimeToLive() time.Time`

GetTimeToLive returns the TimeToLive field if non-nil, zero value otherwise.

### GetTimeToLiveOk

`func (o *DigilockerAadhaarProviderOutput) GetTimeToLiveOk() (*time.Time, bool)`

GetTimeToLiveOk returns a tuple with the TimeToLive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeToLive

`func (o *DigilockerAadhaarProviderOutput) SetTimeToLive(v time.Time)`

SetTimeToLive sets TimeToLive field to given value.

### HasTimeToLive

`func (o *DigilockerAadhaarProviderOutput) HasTimeToLive() bool`

HasTimeToLive returns a boolean if a field has been set.

### SetTimeToLiveNil

`func (o *DigilockerAadhaarProviderOutput) SetTimeToLiveNil(b bool)`

 SetTimeToLiveNil sets the value for TimeToLive to be an explicit nil

### UnsetTimeToLive
`func (o *DigilockerAadhaarProviderOutput) UnsetTimeToLive()`

UnsetTimeToLive ensures that no value is present for TimeToLive, not even an explicit nil
### GetAadhaarNumberLastFour

`func (o *DigilockerAadhaarProviderOutput) GetAadhaarNumberLastFour() string`

GetAadhaarNumberLastFour returns the AadhaarNumberLastFour field if non-nil, zero value otherwise.

### GetAadhaarNumberLastFourOk

`func (o *DigilockerAadhaarProviderOutput) GetAadhaarNumberLastFourOk() (*string, bool)`

GetAadhaarNumberLastFourOk returns a tuple with the AadhaarNumberLastFour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAadhaarNumberLastFour

`func (o *DigilockerAadhaarProviderOutput) SetAadhaarNumberLastFour(v string)`

SetAadhaarNumberLastFour sets AadhaarNumberLastFour field to given value.

### HasAadhaarNumberLastFour

`func (o *DigilockerAadhaarProviderOutput) HasAadhaarNumberLastFour() bool`

HasAadhaarNumberLastFour returns a boolean if a field has been set.

### SetAadhaarNumberLastFourNil

`func (o *DigilockerAadhaarProviderOutput) SetAadhaarNumberLastFourNil(b bool)`

 SetAadhaarNumberLastFourNil sets the value for AadhaarNumberLastFour to be an explicit nil

### UnsetAadhaarNumberLastFour
`func (o *DigilockerAadhaarProviderOutput) UnsetAadhaarNumberLastFour()`

UnsetAadhaarNumberLastFour ensures that no value is present for AadhaarNumberLastFour, not even an explicit nil
### GetClaims

`func (o *DigilockerAadhaarProviderOutput) GetClaims() AadhaarClaims`

GetClaims returns the Claims field if non-nil, zero value otherwise.

### GetClaimsOk

`func (o *DigilockerAadhaarProviderOutput) GetClaimsOk() (*AadhaarClaims, bool)`

GetClaimsOk returns a tuple with the Claims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaims

`func (o *DigilockerAadhaarProviderOutput) SetClaims(v AadhaarClaims)`

SetClaims sets Claims field to given value.

### HasClaims

`func (o *DigilockerAadhaarProviderOutput) HasClaims() bool`

HasClaims returns a boolean if a field has been set.

### SetClaimsNil

`func (o *DigilockerAadhaarProviderOutput) SetClaimsNil(b bool)`

 SetClaimsNil sets the value for Claims to be an explicit nil

### UnsetClaims
`func (o *DigilockerAadhaarProviderOutput) UnsetClaims()`

UnsetClaims ensures that no value is present for Claims, not even an explicit nil
### GetLocalizedClaims

`func (o *DigilockerAadhaarProviderOutput) GetLocalizedClaims() AadhaarLocalizedClaims`

GetLocalizedClaims returns the LocalizedClaims field if non-nil, zero value otherwise.

### GetLocalizedClaimsOk

`func (o *DigilockerAadhaarProviderOutput) GetLocalizedClaimsOk() (*AadhaarLocalizedClaims, bool)`

GetLocalizedClaimsOk returns a tuple with the LocalizedClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizedClaims

`func (o *DigilockerAadhaarProviderOutput) SetLocalizedClaims(v AadhaarLocalizedClaims)`

SetLocalizedClaims sets LocalizedClaims field to given value.

### HasLocalizedClaims

`func (o *DigilockerAadhaarProviderOutput) HasLocalizedClaims() bool`

HasLocalizedClaims returns a boolean if a field has been set.

### SetLocalizedClaimsNil

`func (o *DigilockerAadhaarProviderOutput) SetLocalizedClaimsNil(b bool)`

 SetLocalizedClaimsNil sets the value for LocalizedClaims to be an explicit nil

### UnsetLocalizedClaims
`func (o *DigilockerAadhaarProviderOutput) UnsetLocalizedClaims()`

UnsetLocalizedClaims ensures that no value is present for LocalizedClaims, not even an explicit nil
### GetDocumentSignatureValidated

`func (o *DigilockerAadhaarProviderOutput) GetDocumentSignatureValidated() bool`

GetDocumentSignatureValidated returns the DocumentSignatureValidated field if non-nil, zero value otherwise.

### GetDocumentSignatureValidatedOk

`func (o *DigilockerAadhaarProviderOutput) GetDocumentSignatureValidatedOk() (*bool, bool)`

GetDocumentSignatureValidatedOk returns a tuple with the DocumentSignatureValidated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentSignatureValidated

`func (o *DigilockerAadhaarProviderOutput) SetDocumentSignatureValidated(v bool)`

SetDocumentSignatureValidated sets DocumentSignatureValidated field to given value.

### HasDocumentSignatureValidated

`func (o *DigilockerAadhaarProviderOutput) HasDocumentSignatureValidated() bool`

HasDocumentSignatureValidated returns a boolean if a field has been set.

### SetDocumentSignatureValidatedNil

`func (o *DigilockerAadhaarProviderOutput) SetDocumentSignatureValidatedNil(b bool)`

 SetDocumentSignatureValidatedNil sets the value for DocumentSignatureValidated to be an explicit nil

### UnsetDocumentSignatureValidated
`func (o *DigilockerAadhaarProviderOutput) UnsetDocumentSignatureValidated()`

UnsetDocumentSignatureValidated ensures that no value is present for DocumentSignatureValidated, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


