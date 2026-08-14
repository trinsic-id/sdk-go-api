# ClearProviderOutputTraits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to [**NullableClearProviderOutputAddress**](ClearProviderOutputAddress.md) | The individual&#39;s address. | [optional] 
**DateOfBirth** | Pointer to **NullableString** | The individual&#39;s date of birth. | [optional] 
**Email** | Pointer to **NullableString** | The individual&#39;s email address. | [optional] 
**FirstName** | Pointer to **NullableString** | The individual&#39;s first name. | [optional] 
**LastName** | Pointer to **NullableString** | The individual&#39;s last name. | [optional] 
**MiddleName** | Pointer to **NullableString** | The individual&#39;s middle name. | [optional] 
**SecondFamilyName** | Pointer to **NullableString** | The individual&#39;s second family name. | [optional] 
**FullLastName** | Pointer to **NullableString** | The individual&#39;s full last name. | [optional] 
**Phone** | Pointer to **NullableString** | The individual&#39;s phone number, normalized to international E.164 format. | [optional] 
**Ssn4** | Pointer to **NullableString** | The last four digits of the individual&#39;s SSN. | [optional] 
**Ssn9** | Pointer to **NullableString** | The individual&#39;s SSN. | [optional] 
**IdentificationNumber** | Pointer to **NullableString** | The national identification number. | [optional] 
**IdentificationType** | Pointer to **NullableString** | The national identification number type. | [optional] 
**Document** | Pointer to [**NullableClearProviderOutputDocument**](ClearProviderOutputDocument.md) | Parsed data from the individual&#39;s scanned document. | [optional] 
**HealthInsurance** | Pointer to [**NullableClearProviderOutputHealthInsurance**](ClearProviderOutputHealthInsurance.md) | Health insurance information. | [optional] 
**VerifiedEmail** | Pointer to [**NullableClearProviderOutputVerifiedEmail**](ClearProviderOutputVerifiedEmail.md) | Verified email information. | [optional] 
**HistoricalData** | Pointer to [**NullableClearProviderOutputHistoricalData**](ClearProviderOutputHistoricalData.md) | Historical identity traits. | [optional] 

## Methods

### NewClearProviderOutputTraits

`func NewClearProviderOutputTraits() *ClearProviderOutputTraits`

NewClearProviderOutputTraits instantiates a new ClearProviderOutputTraits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClearProviderOutputTraitsWithDefaults

`func NewClearProviderOutputTraitsWithDefaults() *ClearProviderOutputTraits`

NewClearProviderOutputTraitsWithDefaults instantiates a new ClearProviderOutputTraits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *ClearProviderOutputTraits) GetAddress() ClearProviderOutputAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ClearProviderOutputTraits) GetAddressOk() (*ClearProviderOutputAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ClearProviderOutputTraits) SetAddress(v ClearProviderOutputAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ClearProviderOutputTraits) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *ClearProviderOutputTraits) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *ClearProviderOutputTraits) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetDateOfBirth

`func (o *ClearProviderOutputTraits) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *ClearProviderOutputTraits) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *ClearProviderOutputTraits) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *ClearProviderOutputTraits) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### SetDateOfBirthNil

`func (o *ClearProviderOutputTraits) SetDateOfBirthNil(b bool)`

 SetDateOfBirthNil sets the value for DateOfBirth to be an explicit nil

### UnsetDateOfBirth
`func (o *ClearProviderOutputTraits) UnsetDateOfBirth()`

UnsetDateOfBirth ensures that no value is present for DateOfBirth, not even an explicit nil
### GetEmail

`func (o *ClearProviderOutputTraits) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ClearProviderOutputTraits) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ClearProviderOutputTraits) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ClearProviderOutputTraits) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *ClearProviderOutputTraits) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *ClearProviderOutputTraits) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetFirstName

`func (o *ClearProviderOutputTraits) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *ClearProviderOutputTraits) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *ClearProviderOutputTraits) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *ClearProviderOutputTraits) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *ClearProviderOutputTraits) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *ClearProviderOutputTraits) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *ClearProviderOutputTraits) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *ClearProviderOutputTraits) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *ClearProviderOutputTraits) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *ClearProviderOutputTraits) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *ClearProviderOutputTraits) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *ClearProviderOutputTraits) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetMiddleName

`func (o *ClearProviderOutputTraits) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *ClearProviderOutputTraits) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *ClearProviderOutputTraits) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *ClearProviderOutputTraits) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### SetMiddleNameNil

`func (o *ClearProviderOutputTraits) SetMiddleNameNil(b bool)`

 SetMiddleNameNil sets the value for MiddleName to be an explicit nil

### UnsetMiddleName
`func (o *ClearProviderOutputTraits) UnsetMiddleName()`

UnsetMiddleName ensures that no value is present for MiddleName, not even an explicit nil
### GetSecondFamilyName

`func (o *ClearProviderOutputTraits) GetSecondFamilyName() string`

GetSecondFamilyName returns the SecondFamilyName field if non-nil, zero value otherwise.

### GetSecondFamilyNameOk

`func (o *ClearProviderOutputTraits) GetSecondFamilyNameOk() (*string, bool)`

GetSecondFamilyNameOk returns a tuple with the SecondFamilyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondFamilyName

`func (o *ClearProviderOutputTraits) SetSecondFamilyName(v string)`

SetSecondFamilyName sets SecondFamilyName field to given value.

### HasSecondFamilyName

`func (o *ClearProviderOutputTraits) HasSecondFamilyName() bool`

HasSecondFamilyName returns a boolean if a field has been set.

### SetSecondFamilyNameNil

`func (o *ClearProviderOutputTraits) SetSecondFamilyNameNil(b bool)`

 SetSecondFamilyNameNil sets the value for SecondFamilyName to be an explicit nil

### UnsetSecondFamilyName
`func (o *ClearProviderOutputTraits) UnsetSecondFamilyName()`

UnsetSecondFamilyName ensures that no value is present for SecondFamilyName, not even an explicit nil
### GetFullLastName

`func (o *ClearProviderOutputTraits) GetFullLastName() string`

GetFullLastName returns the FullLastName field if non-nil, zero value otherwise.

### GetFullLastNameOk

`func (o *ClearProviderOutputTraits) GetFullLastNameOk() (*string, bool)`

GetFullLastNameOk returns a tuple with the FullLastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullLastName

`func (o *ClearProviderOutputTraits) SetFullLastName(v string)`

SetFullLastName sets FullLastName field to given value.

### HasFullLastName

`func (o *ClearProviderOutputTraits) HasFullLastName() bool`

HasFullLastName returns a boolean if a field has been set.

### SetFullLastNameNil

`func (o *ClearProviderOutputTraits) SetFullLastNameNil(b bool)`

 SetFullLastNameNil sets the value for FullLastName to be an explicit nil

### UnsetFullLastName
`func (o *ClearProviderOutputTraits) UnsetFullLastName()`

UnsetFullLastName ensures that no value is present for FullLastName, not even an explicit nil
### GetPhone

`func (o *ClearProviderOutputTraits) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *ClearProviderOutputTraits) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *ClearProviderOutputTraits) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *ClearProviderOutputTraits) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### SetPhoneNil

`func (o *ClearProviderOutputTraits) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *ClearProviderOutputTraits) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil
### GetSsn4

`func (o *ClearProviderOutputTraits) GetSsn4() string`

GetSsn4 returns the Ssn4 field if non-nil, zero value otherwise.

### GetSsn4Ok

`func (o *ClearProviderOutputTraits) GetSsn4Ok() (*string, bool)`

GetSsn4Ok returns a tuple with the Ssn4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsn4

`func (o *ClearProviderOutputTraits) SetSsn4(v string)`

SetSsn4 sets Ssn4 field to given value.

### HasSsn4

`func (o *ClearProviderOutputTraits) HasSsn4() bool`

HasSsn4 returns a boolean if a field has been set.

### SetSsn4Nil

`func (o *ClearProviderOutputTraits) SetSsn4Nil(b bool)`

 SetSsn4Nil sets the value for Ssn4 to be an explicit nil

### UnsetSsn4
`func (o *ClearProviderOutputTraits) UnsetSsn4()`

UnsetSsn4 ensures that no value is present for Ssn4, not even an explicit nil
### GetSsn9

`func (o *ClearProviderOutputTraits) GetSsn9() string`

GetSsn9 returns the Ssn9 field if non-nil, zero value otherwise.

### GetSsn9Ok

`func (o *ClearProviderOutputTraits) GetSsn9Ok() (*string, bool)`

GetSsn9Ok returns a tuple with the Ssn9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsn9

`func (o *ClearProviderOutputTraits) SetSsn9(v string)`

SetSsn9 sets Ssn9 field to given value.

### HasSsn9

`func (o *ClearProviderOutputTraits) HasSsn9() bool`

HasSsn9 returns a boolean if a field has been set.

### SetSsn9Nil

`func (o *ClearProviderOutputTraits) SetSsn9Nil(b bool)`

 SetSsn9Nil sets the value for Ssn9 to be an explicit nil

### UnsetSsn9
`func (o *ClearProviderOutputTraits) UnsetSsn9()`

UnsetSsn9 ensures that no value is present for Ssn9, not even an explicit nil
### GetIdentificationNumber

`func (o *ClearProviderOutputTraits) GetIdentificationNumber() string`

GetIdentificationNumber returns the IdentificationNumber field if non-nil, zero value otherwise.

### GetIdentificationNumberOk

`func (o *ClearProviderOutputTraits) GetIdentificationNumberOk() (*string, bool)`

GetIdentificationNumberOk returns a tuple with the IdentificationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentificationNumber

`func (o *ClearProviderOutputTraits) SetIdentificationNumber(v string)`

SetIdentificationNumber sets IdentificationNumber field to given value.

### HasIdentificationNumber

`func (o *ClearProviderOutputTraits) HasIdentificationNumber() bool`

HasIdentificationNumber returns a boolean if a field has been set.

### SetIdentificationNumberNil

`func (o *ClearProviderOutputTraits) SetIdentificationNumberNil(b bool)`

 SetIdentificationNumberNil sets the value for IdentificationNumber to be an explicit nil

### UnsetIdentificationNumber
`func (o *ClearProviderOutputTraits) UnsetIdentificationNumber()`

UnsetIdentificationNumber ensures that no value is present for IdentificationNumber, not even an explicit nil
### GetIdentificationType

`func (o *ClearProviderOutputTraits) GetIdentificationType() string`

GetIdentificationType returns the IdentificationType field if non-nil, zero value otherwise.

### GetIdentificationTypeOk

`func (o *ClearProviderOutputTraits) GetIdentificationTypeOk() (*string, bool)`

GetIdentificationTypeOk returns a tuple with the IdentificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentificationType

`func (o *ClearProviderOutputTraits) SetIdentificationType(v string)`

SetIdentificationType sets IdentificationType field to given value.

### HasIdentificationType

`func (o *ClearProviderOutputTraits) HasIdentificationType() bool`

HasIdentificationType returns a boolean if a field has been set.

### SetIdentificationTypeNil

`func (o *ClearProviderOutputTraits) SetIdentificationTypeNil(b bool)`

 SetIdentificationTypeNil sets the value for IdentificationType to be an explicit nil

### UnsetIdentificationType
`func (o *ClearProviderOutputTraits) UnsetIdentificationType()`

UnsetIdentificationType ensures that no value is present for IdentificationType, not even an explicit nil
### GetDocument

`func (o *ClearProviderOutputTraits) GetDocument() ClearProviderOutputDocument`

GetDocument returns the Document field if non-nil, zero value otherwise.

### GetDocumentOk

`func (o *ClearProviderOutputTraits) GetDocumentOk() (*ClearProviderOutputDocument, bool)`

GetDocumentOk returns a tuple with the Document field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocument

`func (o *ClearProviderOutputTraits) SetDocument(v ClearProviderOutputDocument)`

SetDocument sets Document field to given value.

### HasDocument

`func (o *ClearProviderOutputTraits) HasDocument() bool`

HasDocument returns a boolean if a field has been set.

### SetDocumentNil

`func (o *ClearProviderOutputTraits) SetDocumentNil(b bool)`

 SetDocumentNil sets the value for Document to be an explicit nil

### UnsetDocument
`func (o *ClearProviderOutputTraits) UnsetDocument()`

UnsetDocument ensures that no value is present for Document, not even an explicit nil
### GetHealthInsurance

`func (o *ClearProviderOutputTraits) GetHealthInsurance() ClearProviderOutputHealthInsurance`

GetHealthInsurance returns the HealthInsurance field if non-nil, zero value otherwise.

### GetHealthInsuranceOk

`func (o *ClearProviderOutputTraits) GetHealthInsuranceOk() (*ClearProviderOutputHealthInsurance, bool)`

GetHealthInsuranceOk returns a tuple with the HealthInsurance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthInsurance

`func (o *ClearProviderOutputTraits) SetHealthInsurance(v ClearProviderOutputHealthInsurance)`

SetHealthInsurance sets HealthInsurance field to given value.

### HasHealthInsurance

`func (o *ClearProviderOutputTraits) HasHealthInsurance() bool`

HasHealthInsurance returns a boolean if a field has been set.

### SetHealthInsuranceNil

`func (o *ClearProviderOutputTraits) SetHealthInsuranceNil(b bool)`

 SetHealthInsuranceNil sets the value for HealthInsurance to be an explicit nil

### UnsetHealthInsurance
`func (o *ClearProviderOutputTraits) UnsetHealthInsurance()`

UnsetHealthInsurance ensures that no value is present for HealthInsurance, not even an explicit nil
### GetVerifiedEmail

`func (o *ClearProviderOutputTraits) GetVerifiedEmail() ClearProviderOutputVerifiedEmail`

GetVerifiedEmail returns the VerifiedEmail field if non-nil, zero value otherwise.

### GetVerifiedEmailOk

`func (o *ClearProviderOutputTraits) GetVerifiedEmailOk() (*ClearProviderOutputVerifiedEmail, bool)`

GetVerifiedEmailOk returns a tuple with the VerifiedEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedEmail

`func (o *ClearProviderOutputTraits) SetVerifiedEmail(v ClearProviderOutputVerifiedEmail)`

SetVerifiedEmail sets VerifiedEmail field to given value.

### HasVerifiedEmail

`func (o *ClearProviderOutputTraits) HasVerifiedEmail() bool`

HasVerifiedEmail returns a boolean if a field has been set.

### SetVerifiedEmailNil

`func (o *ClearProviderOutputTraits) SetVerifiedEmailNil(b bool)`

 SetVerifiedEmailNil sets the value for VerifiedEmail to be an explicit nil

### UnsetVerifiedEmail
`func (o *ClearProviderOutputTraits) UnsetVerifiedEmail()`

UnsetVerifiedEmail ensures that no value is present for VerifiedEmail, not even an explicit nil
### GetHistoricalData

`func (o *ClearProviderOutputTraits) GetHistoricalData() ClearProviderOutputHistoricalData`

GetHistoricalData returns the HistoricalData field if non-nil, zero value otherwise.

### GetHistoricalDataOk

`func (o *ClearProviderOutputTraits) GetHistoricalDataOk() (*ClearProviderOutputHistoricalData, bool)`

GetHistoricalDataOk returns a tuple with the HistoricalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoricalData

`func (o *ClearProviderOutputTraits) SetHistoricalData(v ClearProviderOutputHistoricalData)`

SetHistoricalData sets HistoricalData field to given value.

### HasHistoricalData

`func (o *ClearProviderOutputTraits) HasHistoricalData() bool`

HasHistoricalData returns a boolean if a field has been set.

### SetHistoricalDataNil

`func (o *ClearProviderOutputTraits) SetHistoricalDataNil(b bool)`

 SetHistoricalDataNil sets the value for HistoricalData to be an explicit nil

### UnsetHistoricalData
`func (o *ClearProviderOutputTraits) UnsetHistoricalData()`

UnsetHistoricalData ensures that no value is present for HistoricalData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


