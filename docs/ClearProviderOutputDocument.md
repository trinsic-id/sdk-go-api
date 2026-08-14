# ClearProviderOutputDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nationality** | Pointer to **NullableString** | The nationality scanned from the document, normalized to an ISO 3166-1 alpha-2 country code. | [optional] 
**DocumentType** | Pointer to **NullableString** | The document type scanned by CLEAR.              Known values: - drivers_license - paper_passport - passport_card - id_card - visa - health_care_card - other | [optional] 
**IssuingCountry** | Pointer to **NullableString** | The ISO 3166-1 alpha-2 country code of issue. | [optional] 
**IssuingSubdivision** | Pointer to **NullableString** | The ISO 3166 subdivision issuer of the document. | [optional] 
**DocumentNumber** | Pointer to **NullableString** | The document number scanned from the document. | [optional] 
**DateOfExpiry** | Pointer to **NullableString** | The document expiration date. | [optional] 
**Gender** | Pointer to **NullableString** | The gender scanned from the document. CLEAR does not publish a closed set of possible values for this field. | [optional] 
**Address** | Pointer to [**NullableClearProviderOutputAddress**](ClearProviderOutputAddress.md) | The address scanned from the document. | [optional] 
**DateOfBirth** | Pointer to **NullableString** | The date of birth scanned from the document. | [optional] 
**FirstName** | Pointer to **NullableString** | The first name scanned from the document. | [optional] 
**LastName** | Pointer to **NullableString** | The last name scanned from the document. | [optional] 
**MiddleName** | Pointer to **NullableString** | The middle name scanned from the document. | [optional] 
**FullName** | Pointer to **NullableString** | The full name scanned from the document. | [optional] 

## Methods

### NewClearProviderOutputDocument

`func NewClearProviderOutputDocument() *ClearProviderOutputDocument`

NewClearProviderOutputDocument instantiates a new ClearProviderOutputDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClearProviderOutputDocumentWithDefaults

`func NewClearProviderOutputDocumentWithDefaults() *ClearProviderOutputDocument`

NewClearProviderOutputDocumentWithDefaults instantiates a new ClearProviderOutputDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNationality

`func (o *ClearProviderOutputDocument) GetNationality() string`

GetNationality returns the Nationality field if non-nil, zero value otherwise.

### GetNationalityOk

`func (o *ClearProviderOutputDocument) GetNationalityOk() (*string, bool)`

GetNationalityOk returns a tuple with the Nationality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationality

`func (o *ClearProviderOutputDocument) SetNationality(v string)`

SetNationality sets Nationality field to given value.

### HasNationality

`func (o *ClearProviderOutputDocument) HasNationality() bool`

HasNationality returns a boolean if a field has been set.

### SetNationalityNil

`func (o *ClearProviderOutputDocument) SetNationalityNil(b bool)`

 SetNationalityNil sets the value for Nationality to be an explicit nil

### UnsetNationality
`func (o *ClearProviderOutputDocument) UnsetNationality()`

UnsetNationality ensures that no value is present for Nationality, not even an explicit nil
### GetDocumentType

`func (o *ClearProviderOutputDocument) GetDocumentType() string`

GetDocumentType returns the DocumentType field if non-nil, zero value otherwise.

### GetDocumentTypeOk

`func (o *ClearProviderOutputDocument) GetDocumentTypeOk() (*string, bool)`

GetDocumentTypeOk returns a tuple with the DocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentType

`func (o *ClearProviderOutputDocument) SetDocumentType(v string)`

SetDocumentType sets DocumentType field to given value.

### HasDocumentType

`func (o *ClearProviderOutputDocument) HasDocumentType() bool`

HasDocumentType returns a boolean if a field has been set.

### SetDocumentTypeNil

`func (o *ClearProviderOutputDocument) SetDocumentTypeNil(b bool)`

 SetDocumentTypeNil sets the value for DocumentType to be an explicit nil

### UnsetDocumentType
`func (o *ClearProviderOutputDocument) UnsetDocumentType()`

UnsetDocumentType ensures that no value is present for DocumentType, not even an explicit nil
### GetIssuingCountry

`func (o *ClearProviderOutputDocument) GetIssuingCountry() string`

GetIssuingCountry returns the IssuingCountry field if non-nil, zero value otherwise.

### GetIssuingCountryOk

`func (o *ClearProviderOutputDocument) GetIssuingCountryOk() (*string, bool)`

GetIssuingCountryOk returns a tuple with the IssuingCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingCountry

`func (o *ClearProviderOutputDocument) SetIssuingCountry(v string)`

SetIssuingCountry sets IssuingCountry field to given value.

### HasIssuingCountry

`func (o *ClearProviderOutputDocument) HasIssuingCountry() bool`

HasIssuingCountry returns a boolean if a field has been set.

### SetIssuingCountryNil

`func (o *ClearProviderOutputDocument) SetIssuingCountryNil(b bool)`

 SetIssuingCountryNil sets the value for IssuingCountry to be an explicit nil

### UnsetIssuingCountry
`func (o *ClearProviderOutputDocument) UnsetIssuingCountry()`

UnsetIssuingCountry ensures that no value is present for IssuingCountry, not even an explicit nil
### GetIssuingSubdivision

`func (o *ClearProviderOutputDocument) GetIssuingSubdivision() string`

GetIssuingSubdivision returns the IssuingSubdivision field if non-nil, zero value otherwise.

### GetIssuingSubdivisionOk

`func (o *ClearProviderOutputDocument) GetIssuingSubdivisionOk() (*string, bool)`

GetIssuingSubdivisionOk returns a tuple with the IssuingSubdivision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingSubdivision

`func (o *ClearProviderOutputDocument) SetIssuingSubdivision(v string)`

SetIssuingSubdivision sets IssuingSubdivision field to given value.

### HasIssuingSubdivision

`func (o *ClearProviderOutputDocument) HasIssuingSubdivision() bool`

HasIssuingSubdivision returns a boolean if a field has been set.

### SetIssuingSubdivisionNil

`func (o *ClearProviderOutputDocument) SetIssuingSubdivisionNil(b bool)`

 SetIssuingSubdivisionNil sets the value for IssuingSubdivision to be an explicit nil

### UnsetIssuingSubdivision
`func (o *ClearProviderOutputDocument) UnsetIssuingSubdivision()`

UnsetIssuingSubdivision ensures that no value is present for IssuingSubdivision, not even an explicit nil
### GetDocumentNumber

`func (o *ClearProviderOutputDocument) GetDocumentNumber() string`

GetDocumentNumber returns the DocumentNumber field if non-nil, zero value otherwise.

### GetDocumentNumberOk

`func (o *ClearProviderOutputDocument) GetDocumentNumberOk() (*string, bool)`

GetDocumentNumberOk returns a tuple with the DocumentNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentNumber

`func (o *ClearProviderOutputDocument) SetDocumentNumber(v string)`

SetDocumentNumber sets DocumentNumber field to given value.

### HasDocumentNumber

`func (o *ClearProviderOutputDocument) HasDocumentNumber() bool`

HasDocumentNumber returns a boolean if a field has been set.

### SetDocumentNumberNil

`func (o *ClearProviderOutputDocument) SetDocumentNumberNil(b bool)`

 SetDocumentNumberNil sets the value for DocumentNumber to be an explicit nil

### UnsetDocumentNumber
`func (o *ClearProviderOutputDocument) UnsetDocumentNumber()`

UnsetDocumentNumber ensures that no value is present for DocumentNumber, not even an explicit nil
### GetDateOfExpiry

`func (o *ClearProviderOutputDocument) GetDateOfExpiry() string`

GetDateOfExpiry returns the DateOfExpiry field if non-nil, zero value otherwise.

### GetDateOfExpiryOk

`func (o *ClearProviderOutputDocument) GetDateOfExpiryOk() (*string, bool)`

GetDateOfExpiryOk returns a tuple with the DateOfExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfExpiry

`func (o *ClearProviderOutputDocument) SetDateOfExpiry(v string)`

SetDateOfExpiry sets DateOfExpiry field to given value.

### HasDateOfExpiry

`func (o *ClearProviderOutputDocument) HasDateOfExpiry() bool`

HasDateOfExpiry returns a boolean if a field has been set.

### SetDateOfExpiryNil

`func (o *ClearProviderOutputDocument) SetDateOfExpiryNil(b bool)`

 SetDateOfExpiryNil sets the value for DateOfExpiry to be an explicit nil

### UnsetDateOfExpiry
`func (o *ClearProviderOutputDocument) UnsetDateOfExpiry()`

UnsetDateOfExpiry ensures that no value is present for DateOfExpiry, not even an explicit nil
### GetGender

`func (o *ClearProviderOutputDocument) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *ClearProviderOutputDocument) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *ClearProviderOutputDocument) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *ClearProviderOutputDocument) HasGender() bool`

HasGender returns a boolean if a field has been set.

### SetGenderNil

`func (o *ClearProviderOutputDocument) SetGenderNil(b bool)`

 SetGenderNil sets the value for Gender to be an explicit nil

### UnsetGender
`func (o *ClearProviderOutputDocument) UnsetGender()`

UnsetGender ensures that no value is present for Gender, not even an explicit nil
### GetAddress

`func (o *ClearProviderOutputDocument) GetAddress() ClearProviderOutputAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ClearProviderOutputDocument) GetAddressOk() (*ClearProviderOutputAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ClearProviderOutputDocument) SetAddress(v ClearProviderOutputAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ClearProviderOutputDocument) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *ClearProviderOutputDocument) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *ClearProviderOutputDocument) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetDateOfBirth

`func (o *ClearProviderOutputDocument) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *ClearProviderOutputDocument) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *ClearProviderOutputDocument) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *ClearProviderOutputDocument) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### SetDateOfBirthNil

`func (o *ClearProviderOutputDocument) SetDateOfBirthNil(b bool)`

 SetDateOfBirthNil sets the value for DateOfBirth to be an explicit nil

### UnsetDateOfBirth
`func (o *ClearProviderOutputDocument) UnsetDateOfBirth()`

UnsetDateOfBirth ensures that no value is present for DateOfBirth, not even an explicit nil
### GetFirstName

`func (o *ClearProviderOutputDocument) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *ClearProviderOutputDocument) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *ClearProviderOutputDocument) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *ClearProviderOutputDocument) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *ClearProviderOutputDocument) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *ClearProviderOutputDocument) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *ClearProviderOutputDocument) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *ClearProviderOutputDocument) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *ClearProviderOutputDocument) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *ClearProviderOutputDocument) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *ClearProviderOutputDocument) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *ClearProviderOutputDocument) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetMiddleName

`func (o *ClearProviderOutputDocument) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *ClearProviderOutputDocument) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *ClearProviderOutputDocument) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *ClearProviderOutputDocument) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### SetMiddleNameNil

`func (o *ClearProviderOutputDocument) SetMiddleNameNil(b bool)`

 SetMiddleNameNil sets the value for MiddleName to be an explicit nil

### UnsetMiddleName
`func (o *ClearProviderOutputDocument) UnsetMiddleName()`

UnsetMiddleName ensures that no value is present for MiddleName, not even an explicit nil
### GetFullName

`func (o *ClearProviderOutputDocument) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *ClearProviderOutputDocument) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *ClearProviderOutputDocument) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *ClearProviderOutputDocument) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### SetFullNameNil

`func (o *ClearProviderOutputDocument) SetFullNameNil(b bool)`

 SetFullNameNil sets the value for FullName to be an explicit nil

### UnsetFullName
`func (o *ClearProviderOutputDocument) UnsetFullName()`

UnsetFullName ensures that no value is present for FullName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


