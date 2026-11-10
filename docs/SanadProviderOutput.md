# SanadProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NationalNumber** | Pointer to **NullableString** | The national number (الرقم الوطني).              This is a ten-digit number, with leading zeros preserved. The number is different from the identity-card number. | [optional] 
**DateOfBirth** | Pointer to **NullableString** | The date of birth of the individual. | [optional] 
**FullName** | Pointer to **NullableString** | The full name of the individual. | [optional] 
**DocumentNumber** | Pointer to **NullableString** | The identity-card number, separate from the national number.              The format is three letters and five digits. | [optional] 
**ExpirationDate** | Pointer to **NullableString** | The document expiration date. | [optional] 
**QrCreatedDate** | Pointer to **NullableString** | The QR creation date, not the document&#39;s issue date. | [optional] 

## Methods

### Named construction

Use a keyed struct literal (`SanadProviderOutput{FieldName: value}`), or call
`NewSanadProviderOutputWithDefaults()` and set fields by name. Positional field
factories and unkeyed struct literals are not supported.

### NewSanadProviderOutputWithDefaults

`func NewSanadProviderOutputWithDefaults() *SanadProviderOutput`

NewSanadProviderOutputWithDefaults instantiates a new SanadProviderOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNationalNumber

`func (o *SanadProviderOutput) GetNationalNumber() string`

GetNationalNumber returns the NationalNumber field if non-nil, zero value otherwise.

### GetNationalNumberOk

`func (o *SanadProviderOutput) GetNationalNumberOk() (*string, bool)`

GetNationalNumberOk returns a tuple with the NationalNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationalNumber

`func (o *SanadProviderOutput) SetNationalNumber(v string)`

SetNationalNumber sets NationalNumber field to given value.

### HasNationalNumber

`func (o *SanadProviderOutput) HasNationalNumber() bool`

HasNationalNumber returns a boolean if a field has been set.

### SetNationalNumberNil

`func (o *SanadProviderOutput) SetNationalNumberNil(b bool)`

 SetNationalNumberNil sets the value for NationalNumber to be an explicit nil

### UnsetNationalNumber
`func (o *SanadProviderOutput) UnsetNationalNumber()`

UnsetNationalNumber ensures that no value is present for NationalNumber, not even an explicit nil
### GetDateOfBirth

`func (o *SanadProviderOutput) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *SanadProviderOutput) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *SanadProviderOutput) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *SanadProviderOutput) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### SetDateOfBirthNil

`func (o *SanadProviderOutput) SetDateOfBirthNil(b bool)`

 SetDateOfBirthNil sets the value for DateOfBirth to be an explicit nil

### UnsetDateOfBirth
`func (o *SanadProviderOutput) UnsetDateOfBirth()`

UnsetDateOfBirth ensures that no value is present for DateOfBirth, not even an explicit nil
### GetFullName

`func (o *SanadProviderOutput) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *SanadProviderOutput) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *SanadProviderOutput) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *SanadProviderOutput) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### SetFullNameNil

`func (o *SanadProviderOutput) SetFullNameNil(b bool)`

 SetFullNameNil sets the value for FullName to be an explicit nil

### UnsetFullName
`func (o *SanadProviderOutput) UnsetFullName()`

UnsetFullName ensures that no value is present for FullName, not even an explicit nil
### GetDocumentNumber

`func (o *SanadProviderOutput) GetDocumentNumber() string`

GetDocumentNumber returns the DocumentNumber field if non-nil, zero value otherwise.

### GetDocumentNumberOk

`func (o *SanadProviderOutput) GetDocumentNumberOk() (*string, bool)`

GetDocumentNumberOk returns a tuple with the DocumentNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentNumber

`func (o *SanadProviderOutput) SetDocumentNumber(v string)`

SetDocumentNumber sets DocumentNumber field to given value.

### HasDocumentNumber

`func (o *SanadProviderOutput) HasDocumentNumber() bool`

HasDocumentNumber returns a boolean if a field has been set.

### SetDocumentNumberNil

`func (o *SanadProviderOutput) SetDocumentNumberNil(b bool)`

 SetDocumentNumberNil sets the value for DocumentNumber to be an explicit nil

### UnsetDocumentNumber
`func (o *SanadProviderOutput) UnsetDocumentNumber()`

UnsetDocumentNumber ensures that no value is present for DocumentNumber, not even an explicit nil
### GetExpirationDate

`func (o *SanadProviderOutput) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *SanadProviderOutput) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *SanadProviderOutput) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *SanadProviderOutput) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### SetExpirationDateNil

`func (o *SanadProviderOutput) SetExpirationDateNil(b bool)`

 SetExpirationDateNil sets the value for ExpirationDate to be an explicit nil

### UnsetExpirationDate
`func (o *SanadProviderOutput) UnsetExpirationDate()`

UnsetExpirationDate ensures that no value is present for ExpirationDate, not even an explicit nil
### GetQrCreatedDate

`func (o *SanadProviderOutput) GetQrCreatedDate() string`

GetQrCreatedDate returns the QrCreatedDate field if non-nil, zero value otherwise.

### GetQrCreatedDateOk

`func (o *SanadProviderOutput) GetQrCreatedDateOk() (*string, bool)`

GetQrCreatedDateOk returns a tuple with the QrCreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQrCreatedDate

`func (o *SanadProviderOutput) SetQrCreatedDate(v string)`

SetQrCreatedDate sets QrCreatedDate field to given value.

### HasQrCreatedDate

`func (o *SanadProviderOutput) HasQrCreatedDate() bool`

HasQrCreatedDate returns a boolean if a field has been set.

### SetQrCreatedDateNil

`func (o *SanadProviderOutput) SetQrCreatedDateNil(b bool)`

 SetQrCreatedDateNil sets the value for QrCreatedDate to be an explicit nil

### UnsetQrCreatedDate
`func (o *SanadProviderOutput) UnsetQrCreatedDate()`

UnsetQrCreatedDate ensures that no value is present for QrCreatedDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


