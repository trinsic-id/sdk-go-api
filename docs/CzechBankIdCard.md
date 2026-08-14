# CzechBankIdCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **NullableString** | The identity document type code.              Possible values are:              - ID - Identity card - P - Passport - DL - Driving license - IR - Residence permit - VS - Visa permit label - PS - Residential label - IX - Book with residence permit - IE - Form with temporary residence - OP - Identity card – without machine readable zone - CA - Passport of the Czech Republic resident – without machine readable zone - UNKNOWN - Unknown id card type | [optional] 
**Description** | Pointer to **NullableString** | The localized identity document type description. | [optional] 
**Country** | Pointer to **NullableString** | The country for which the identity document is valid. | [optional] 
**Number** | Pointer to **NullableString** | The identity document number. | [optional] 
**ValidTo** | Pointer to **NullableString** | The identity document expiration date. | [optional] 
**Issuer** | Pointer to **NullableString** | The office that issued the identity document. | [optional] 
**IssueDate** | Pointer to **NullableString** | The identity document issue date. | [optional] 

## Methods

### NewCzechBankIdCard

`func NewCzechBankIdCard() *CzechBankIdCard`

NewCzechBankIdCard instantiates a new CzechBankIdCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCzechBankIdCardWithDefaults

`func NewCzechBankIdCardWithDefaults() *CzechBankIdCard`

NewCzechBankIdCardWithDefaults instantiates a new CzechBankIdCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CzechBankIdCard) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CzechBankIdCard) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CzechBankIdCard) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CzechBankIdCard) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *CzechBankIdCard) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *CzechBankIdCard) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetDescription

`func (o *CzechBankIdCard) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CzechBankIdCard) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CzechBankIdCard) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CzechBankIdCard) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CzechBankIdCard) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CzechBankIdCard) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCountry

`func (o *CzechBankIdCard) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *CzechBankIdCard) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *CzechBankIdCard) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *CzechBankIdCard) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *CzechBankIdCard) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *CzechBankIdCard) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetNumber

`func (o *CzechBankIdCard) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *CzechBankIdCard) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *CzechBankIdCard) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *CzechBankIdCard) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### SetNumberNil

`func (o *CzechBankIdCard) SetNumberNil(b bool)`

 SetNumberNil sets the value for Number to be an explicit nil

### UnsetNumber
`func (o *CzechBankIdCard) UnsetNumber()`

UnsetNumber ensures that no value is present for Number, not even an explicit nil
### GetValidTo

`func (o *CzechBankIdCard) GetValidTo() string`

GetValidTo returns the ValidTo field if non-nil, zero value otherwise.

### GetValidToOk

`func (o *CzechBankIdCard) GetValidToOk() (*string, bool)`

GetValidToOk returns a tuple with the ValidTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidTo

`func (o *CzechBankIdCard) SetValidTo(v string)`

SetValidTo sets ValidTo field to given value.

### HasValidTo

`func (o *CzechBankIdCard) HasValidTo() bool`

HasValidTo returns a boolean if a field has been set.

### SetValidToNil

`func (o *CzechBankIdCard) SetValidToNil(b bool)`

 SetValidToNil sets the value for ValidTo to be an explicit nil

### UnsetValidTo
`func (o *CzechBankIdCard) UnsetValidTo()`

UnsetValidTo ensures that no value is present for ValidTo, not even an explicit nil
### GetIssuer

`func (o *CzechBankIdCard) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *CzechBankIdCard) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *CzechBankIdCard) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *CzechBankIdCard) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### SetIssuerNil

`func (o *CzechBankIdCard) SetIssuerNil(b bool)`

 SetIssuerNil sets the value for Issuer to be an explicit nil

### UnsetIssuer
`func (o *CzechBankIdCard) UnsetIssuer()`

UnsetIssuer ensures that no value is present for Issuer, not even an explicit nil
### GetIssueDate

`func (o *CzechBankIdCard) GetIssueDate() string`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *CzechBankIdCard) GetIssueDateOk() (*string, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *CzechBankIdCard) SetIssueDate(v string)`

SetIssueDate sets IssueDate field to given value.

### HasIssueDate

`func (o *CzechBankIdCard) HasIssueDate() bool`

HasIssueDate returns a boolean if a field has been set.

### SetIssueDateNil

`func (o *CzechBankIdCard) SetIssueDateNil(b bool)`

 SetIssueDateNil sets the value for IssueDate to be an explicit nil

### UnsetIssueDate
`func (o *CzechBankIdCard) UnsetIssueDate()`

UnsetIssueDate ensures that no value is present for IssueDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


