# YotiDocumentDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **NullableString** | The type of document used to create the Yoti credential.              Possible values: - PASSPORT: A government-issued identity document - DRIVING_LICENCE: A government-issued identity card - NATIONAL_ID: A government-issued identity card - PASS_CARD: A physical or digital identity card that is accredited by the UK&#39;s national Proof of Age Standards Scheme (PASS) | [optional] 
**IssuingCountry** | Pointer to **NullableString** | The issuing country&#39;s ISO alpha-3 code | [optional] 
**DocumentNumber** | Pointer to **NullableString** | The arbitrary document number for the provided document type | [optional] 
**ExpirationDate** | Pointer to **NullableString** | The document&#39;s expiration date in YYYY-MM-DD format | [optional] 
**IssuingAuthority** | Pointer to **NullableString** | The document&#39;s issuing authority. This can either be a country code as a specified ISO alpha-3 or the name of the issuing authority. | [optional] 

## Methods

### NewYotiDocumentDetails

`func NewYotiDocumentDetails() *YotiDocumentDetails`

NewYotiDocumentDetails instantiates a new YotiDocumentDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewYotiDocumentDetailsWithDefaults

`func NewYotiDocumentDetailsWithDefaults() *YotiDocumentDetails`

NewYotiDocumentDetailsWithDefaults instantiates a new YotiDocumentDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *YotiDocumentDetails) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *YotiDocumentDetails) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *YotiDocumentDetails) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *YotiDocumentDetails) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *YotiDocumentDetails) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *YotiDocumentDetails) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetIssuingCountry

`func (o *YotiDocumentDetails) GetIssuingCountry() string`

GetIssuingCountry returns the IssuingCountry field if non-nil, zero value otherwise.

### GetIssuingCountryOk

`func (o *YotiDocumentDetails) GetIssuingCountryOk() (*string, bool)`

GetIssuingCountryOk returns a tuple with the IssuingCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingCountry

`func (o *YotiDocumentDetails) SetIssuingCountry(v string)`

SetIssuingCountry sets IssuingCountry field to given value.

### HasIssuingCountry

`func (o *YotiDocumentDetails) HasIssuingCountry() bool`

HasIssuingCountry returns a boolean if a field has been set.

### SetIssuingCountryNil

`func (o *YotiDocumentDetails) SetIssuingCountryNil(b bool)`

 SetIssuingCountryNil sets the value for IssuingCountry to be an explicit nil

### UnsetIssuingCountry
`func (o *YotiDocumentDetails) UnsetIssuingCountry()`

UnsetIssuingCountry ensures that no value is present for IssuingCountry, not even an explicit nil
### GetDocumentNumber

`func (o *YotiDocumentDetails) GetDocumentNumber() string`

GetDocumentNumber returns the DocumentNumber field if non-nil, zero value otherwise.

### GetDocumentNumberOk

`func (o *YotiDocumentDetails) GetDocumentNumberOk() (*string, bool)`

GetDocumentNumberOk returns a tuple with the DocumentNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentNumber

`func (o *YotiDocumentDetails) SetDocumentNumber(v string)`

SetDocumentNumber sets DocumentNumber field to given value.

### HasDocumentNumber

`func (o *YotiDocumentDetails) HasDocumentNumber() bool`

HasDocumentNumber returns a boolean if a field has been set.

### SetDocumentNumberNil

`func (o *YotiDocumentDetails) SetDocumentNumberNil(b bool)`

 SetDocumentNumberNil sets the value for DocumentNumber to be an explicit nil

### UnsetDocumentNumber
`func (o *YotiDocumentDetails) UnsetDocumentNumber()`

UnsetDocumentNumber ensures that no value is present for DocumentNumber, not even an explicit nil
### GetExpirationDate

`func (o *YotiDocumentDetails) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *YotiDocumentDetails) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *YotiDocumentDetails) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *YotiDocumentDetails) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### SetExpirationDateNil

`func (o *YotiDocumentDetails) SetExpirationDateNil(b bool)`

 SetExpirationDateNil sets the value for ExpirationDate to be an explicit nil

### UnsetExpirationDate
`func (o *YotiDocumentDetails) UnsetExpirationDate()`

UnsetExpirationDate ensures that no value is present for ExpirationDate, not even an explicit nil
### GetIssuingAuthority

`func (o *YotiDocumentDetails) GetIssuingAuthority() string`

GetIssuingAuthority returns the IssuingAuthority field if non-nil, zero value otherwise.

### GetIssuingAuthorityOk

`func (o *YotiDocumentDetails) GetIssuingAuthorityOk() (*string, bool)`

GetIssuingAuthorityOk returns a tuple with the IssuingAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingAuthority

`func (o *YotiDocumentDetails) SetIssuingAuthority(v string)`

SetIssuingAuthority sets IssuingAuthority field to given value.

### HasIssuingAuthority

`func (o *YotiDocumentDetails) HasIssuingAuthority() bool`

HasIssuingAuthority returns a boolean if a field has been set.

### SetIssuingAuthorityNil

`func (o *YotiDocumentDetails) SetIssuingAuthorityNil(b bool)`

 SetIssuingAuthorityNil sets the value for IssuingAuthority to be an explicit nil

### UnsetIssuingAuthority
`func (o *YotiDocumentDetails) UnsetIssuingAuthority()`

UnsetIssuingAuthority ensures that no value is present for IssuingAuthority, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


