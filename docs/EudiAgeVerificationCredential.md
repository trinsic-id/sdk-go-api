# EudiAgeVerificationCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgeOver** | Pointer to [**[]AgeOverOutput**](AgeOverOutput.md) | Processed age-over claims returned by the credential. | [optional] 
**ExpiryDate** | Pointer to **NullableString** | Date when the age verification data expires. | [optional] 
**IssuingCountry** | Pointer to **NullableString** | ISO 3166-1 alpha-2 country code of the country or territory of the issuer of the credential. | [optional] 

## Methods

### NewEudiAgeVerificationCredential

`func NewEudiAgeVerificationCredential() *EudiAgeVerificationCredential`

NewEudiAgeVerificationCredential instantiates a new EudiAgeVerificationCredential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEudiAgeVerificationCredentialWithDefaults

`func NewEudiAgeVerificationCredentialWithDefaults() *EudiAgeVerificationCredential`

NewEudiAgeVerificationCredentialWithDefaults instantiates a new EudiAgeVerificationCredential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgeOver

`func (o *EudiAgeVerificationCredential) GetAgeOver() []AgeOverOutput`

GetAgeOver returns the AgeOver field if non-nil, zero value otherwise.

### GetAgeOverOk

`func (o *EudiAgeVerificationCredential) GetAgeOverOk() (*[]AgeOverOutput, bool)`

GetAgeOverOk returns a tuple with the AgeOver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgeOver

`func (o *EudiAgeVerificationCredential) SetAgeOver(v []AgeOverOutput)`

SetAgeOver sets AgeOver field to given value.

### HasAgeOver

`func (o *EudiAgeVerificationCredential) HasAgeOver() bool`

HasAgeOver returns a boolean if a field has been set.

### SetAgeOverNil

`func (o *EudiAgeVerificationCredential) SetAgeOverNil(b bool)`

 SetAgeOverNil sets the value for AgeOver to be an explicit nil

### UnsetAgeOver
`func (o *EudiAgeVerificationCredential) UnsetAgeOver()`

UnsetAgeOver ensures that no value is present for AgeOver, not even an explicit nil
### GetExpiryDate

`func (o *EudiAgeVerificationCredential) GetExpiryDate() string`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *EudiAgeVerificationCredential) GetExpiryDateOk() (*string, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *EudiAgeVerificationCredential) SetExpiryDate(v string)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *EudiAgeVerificationCredential) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### SetExpiryDateNil

`func (o *EudiAgeVerificationCredential) SetExpiryDateNil(b bool)`

 SetExpiryDateNil sets the value for ExpiryDate to be an explicit nil

### UnsetExpiryDate
`func (o *EudiAgeVerificationCredential) UnsetExpiryDate()`

UnsetExpiryDate ensures that no value is present for ExpiryDate, not even an explicit nil
### GetIssuingCountry

`func (o *EudiAgeVerificationCredential) GetIssuingCountry() string`

GetIssuingCountry returns the IssuingCountry field if non-nil, zero value otherwise.

### GetIssuingCountryOk

`func (o *EudiAgeVerificationCredential) GetIssuingCountryOk() (*string, bool)`

GetIssuingCountryOk returns a tuple with the IssuingCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingCountry

`func (o *EudiAgeVerificationCredential) SetIssuingCountry(v string)`

SetIssuingCountry sets IssuingCountry field to given value.

### HasIssuingCountry

`func (o *EudiAgeVerificationCredential) HasIssuingCountry() bool`

HasIssuingCountry returns a boolean if a field has been set.

### SetIssuingCountryNil

`func (o *EudiAgeVerificationCredential) SetIssuingCountryNil(b bool)`

 SetIssuingCountryNil sets the value for IssuingCountry to be an explicit nil

### UnsetIssuingCountry
`func (o *EudiAgeVerificationCredential) UnsetIssuingCountry()`

UnsetIssuingCountry ensures that no value is present for IssuingCountry, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


