# CzechBankIdCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The identity document type code.              Possible values are:              - ID - Identity card - P - Passport - DL - Driving license - IR - Residence permit - VS - Visa permit label - PS - Residential label - IX - Book with residence permit - IE - Form with temporary residence - OP - Identity card – without machine readable zone - CA - Passport of the Czech Republic resident – without machine readable zone - UNKNOWN - Unknown id card type | 
**Description** | Pointer to **NullableString** | The localized identity document type description. | [optional] 
**Country** | **string** | The country for which the identity document is valid. | 
**Number** | **string** | The identity document number. | 
**ValidTo** | **string** | The identity document expiration date. | 
**Issuer** | Pointer to **NullableString** | The office that issued the identity document. | [optional] 
**IssueDate** | Pointer to **NullableString** | The identity document issue date. | [optional] 

## Methods

### NewCzechBankIdCard

`func NewCzechBankIdCard(type_ string, country string, number string, validTo string, ) *CzechBankIdCard`

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


