# DocumentData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**NullableDocumentType**](DocumentType.md) | The type of the document. | [optional] 
**Number** | Pointer to **NullableString** | The primary identifying number of the document. | [optional] 
**IssueDate** | Pointer to **NullableString** | The date the document was issued. | [optional] 
**ExpirationDate** | Pointer to **NullableString** | The date the document expires. | [optional] 
**IssuingCountry** | Pointer to **NullableString** | The ISO 3166-1 alpha-2 country code of the country that issued the document. | [optional] 
**IssuingSubdivision** | Pointer to **NullableString** | The ISO 3166-2 subdivision code of the issuing authority which issued the document.              This is always in the form {CountryCode}-{SubdivisionCode}, where CountryCode is 2 letters and SubdivisionCode is 1-3 alphanumeric characters. | [optional] 
**IssuingAuthority** | Pointer to **NullableString** | The name of the authority which issued the document. | [optional] 

## Methods

### NewDocumentData

`func NewDocumentData() *DocumentData`

NewDocumentData instantiates a new DocumentData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentDataWithDefaults

`func NewDocumentDataWithDefaults() *DocumentData`

NewDocumentDataWithDefaults instantiates a new DocumentData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DocumentData) GetType() DocumentType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DocumentData) GetTypeOk() (*DocumentType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DocumentData) SetType(v DocumentType)`

SetType sets Type field to given value.

### HasType

`func (o *DocumentData) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *DocumentData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *DocumentData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetNumber

`func (o *DocumentData) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *DocumentData) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *DocumentData) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *DocumentData) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### SetNumberNil

`func (o *DocumentData) SetNumberNil(b bool)`

 SetNumberNil sets the value for Number to be an explicit nil

### UnsetNumber
`func (o *DocumentData) UnsetNumber()`

UnsetNumber ensures that no value is present for Number, not even an explicit nil
### GetIssueDate

`func (o *DocumentData) GetIssueDate() string`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *DocumentData) GetIssueDateOk() (*string, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *DocumentData) SetIssueDate(v string)`

SetIssueDate sets IssueDate field to given value.

### HasIssueDate

`func (o *DocumentData) HasIssueDate() bool`

HasIssueDate returns a boolean if a field has been set.

### SetIssueDateNil

`func (o *DocumentData) SetIssueDateNil(b bool)`

 SetIssueDateNil sets the value for IssueDate to be an explicit nil

### UnsetIssueDate
`func (o *DocumentData) UnsetIssueDate()`

UnsetIssueDate ensures that no value is present for IssueDate, not even an explicit nil
### GetExpirationDate

`func (o *DocumentData) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *DocumentData) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *DocumentData) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *DocumentData) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### SetExpirationDateNil

`func (o *DocumentData) SetExpirationDateNil(b bool)`

 SetExpirationDateNil sets the value for ExpirationDate to be an explicit nil

### UnsetExpirationDate
`func (o *DocumentData) UnsetExpirationDate()`

UnsetExpirationDate ensures that no value is present for ExpirationDate, not even an explicit nil
### GetIssuingCountry

`func (o *DocumentData) GetIssuingCountry() string`

GetIssuingCountry returns the IssuingCountry field if non-nil, zero value otherwise.

### GetIssuingCountryOk

`func (o *DocumentData) GetIssuingCountryOk() (*string, bool)`

GetIssuingCountryOk returns a tuple with the IssuingCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingCountry

`func (o *DocumentData) SetIssuingCountry(v string)`

SetIssuingCountry sets IssuingCountry field to given value.

### HasIssuingCountry

`func (o *DocumentData) HasIssuingCountry() bool`

HasIssuingCountry returns a boolean if a field has been set.

### SetIssuingCountryNil

`func (o *DocumentData) SetIssuingCountryNil(b bool)`

 SetIssuingCountryNil sets the value for IssuingCountry to be an explicit nil

### UnsetIssuingCountry
`func (o *DocumentData) UnsetIssuingCountry()`

UnsetIssuingCountry ensures that no value is present for IssuingCountry, not even an explicit nil
### GetIssuingSubdivision

`func (o *DocumentData) GetIssuingSubdivision() string`

GetIssuingSubdivision returns the IssuingSubdivision field if non-nil, zero value otherwise.

### GetIssuingSubdivisionOk

`func (o *DocumentData) GetIssuingSubdivisionOk() (*string, bool)`

GetIssuingSubdivisionOk returns a tuple with the IssuingSubdivision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingSubdivision

`func (o *DocumentData) SetIssuingSubdivision(v string)`

SetIssuingSubdivision sets IssuingSubdivision field to given value.

### HasIssuingSubdivision

`func (o *DocumentData) HasIssuingSubdivision() bool`

HasIssuingSubdivision returns a boolean if a field has been set.

### SetIssuingSubdivisionNil

`func (o *DocumentData) SetIssuingSubdivisionNil(b bool)`

 SetIssuingSubdivisionNil sets the value for IssuingSubdivision to be an explicit nil

### UnsetIssuingSubdivision
`func (o *DocumentData) UnsetIssuingSubdivision()`

UnsetIssuingSubdivision ensures that no value is present for IssuingSubdivision, not even an explicit nil
### GetIssuingAuthority

`func (o *DocumentData) GetIssuingAuthority() string`

GetIssuingAuthority returns the IssuingAuthority field if non-nil, zero value otherwise.

### GetIssuingAuthorityOk

`func (o *DocumentData) GetIssuingAuthorityOk() (*string, bool)`

GetIssuingAuthorityOk returns a tuple with the IssuingAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingAuthority

`func (o *DocumentData) SetIssuingAuthority(v string)`

SetIssuingAuthority sets IssuingAuthority field to given value.

### HasIssuingAuthority

`func (o *DocumentData) HasIssuingAuthority() bool`

HasIssuingAuthority returns a boolean if a field has been set.

### SetIssuingAuthorityNil

`func (o *DocumentData) SetIssuingAuthorityNil(b bool)`

 SetIssuingAuthorityNil sets the value for IssuingAuthority to be an explicit nil

### UnsetIssuingAuthority
`func (o *DocumentData) UnsetIssuingAuthority()`

UnsetIssuingAuthority ensures that no value is present for IssuingAuthority, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


