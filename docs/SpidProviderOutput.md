# SpidProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingInformation** | Pointer to [**NullableSpidBillingInformation**](SpidBillingInformation.md) | Information about the billable status of this SPID Verification.              Present only if your account has period-based billing enabled for SPID. Contact Trinsic to enable this. | [optional] 
**FiscalNumber** | Pointer to **NullableString** | Fiscal tax number for the subject. | [optional] 
**SpidCode** | Pointer to **NullableString** | Unique user identifier contained within the SPID identity. | [optional] 
**IvaCode** | Pointer to **NullableString** | VAT number for the subject. | [optional] 
**SpidCredentialExpirationDate** | Pointer to **NullableString** | Expiration date of the user&#39;s SPID credential.              This is not the same as the expiration date of the underlying identity document (such as a passport) which was used to create the SPID identity. | [optional] 

## Methods

### NewSpidProviderOutput

`func NewSpidProviderOutput() *SpidProviderOutput`

NewSpidProviderOutput instantiates a new SpidProviderOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpidProviderOutputWithDefaults

`func NewSpidProviderOutputWithDefaults() *SpidProviderOutput`

NewSpidProviderOutputWithDefaults instantiates a new SpidProviderOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingInformation

`func (o *SpidProviderOutput) GetBillingInformation() SpidBillingInformation`

GetBillingInformation returns the BillingInformation field if non-nil, zero value otherwise.

### GetBillingInformationOk

`func (o *SpidProviderOutput) GetBillingInformationOk() (*SpidBillingInformation, bool)`

GetBillingInformationOk returns a tuple with the BillingInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingInformation

`func (o *SpidProviderOutput) SetBillingInformation(v SpidBillingInformation)`

SetBillingInformation sets BillingInformation field to given value.

### HasBillingInformation

`func (o *SpidProviderOutput) HasBillingInformation() bool`

HasBillingInformation returns a boolean if a field has been set.

### SetBillingInformationNil

`func (o *SpidProviderOutput) SetBillingInformationNil(b bool)`

 SetBillingInformationNil sets the value for BillingInformation to be an explicit nil

### UnsetBillingInformation
`func (o *SpidProviderOutput) UnsetBillingInformation()`

UnsetBillingInformation ensures that no value is present for BillingInformation, not even an explicit nil
### GetFiscalNumber

`func (o *SpidProviderOutput) GetFiscalNumber() string`

GetFiscalNumber returns the FiscalNumber field if non-nil, zero value otherwise.

### GetFiscalNumberOk

`func (o *SpidProviderOutput) GetFiscalNumberOk() (*string, bool)`

GetFiscalNumberOk returns a tuple with the FiscalNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiscalNumber

`func (o *SpidProviderOutput) SetFiscalNumber(v string)`

SetFiscalNumber sets FiscalNumber field to given value.

### HasFiscalNumber

`func (o *SpidProviderOutput) HasFiscalNumber() bool`

HasFiscalNumber returns a boolean if a field has been set.

### SetFiscalNumberNil

`func (o *SpidProviderOutput) SetFiscalNumberNil(b bool)`

 SetFiscalNumberNil sets the value for FiscalNumber to be an explicit nil

### UnsetFiscalNumber
`func (o *SpidProviderOutput) UnsetFiscalNumber()`

UnsetFiscalNumber ensures that no value is present for FiscalNumber, not even an explicit nil
### GetSpidCode

`func (o *SpidProviderOutput) GetSpidCode() string`

GetSpidCode returns the SpidCode field if non-nil, zero value otherwise.

### GetSpidCodeOk

`func (o *SpidProviderOutput) GetSpidCodeOk() (*string, bool)`

GetSpidCodeOk returns a tuple with the SpidCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpidCode

`func (o *SpidProviderOutput) SetSpidCode(v string)`

SetSpidCode sets SpidCode field to given value.

### HasSpidCode

`func (o *SpidProviderOutput) HasSpidCode() bool`

HasSpidCode returns a boolean if a field has been set.

### SetSpidCodeNil

`func (o *SpidProviderOutput) SetSpidCodeNil(b bool)`

 SetSpidCodeNil sets the value for SpidCode to be an explicit nil

### UnsetSpidCode
`func (o *SpidProviderOutput) UnsetSpidCode()`

UnsetSpidCode ensures that no value is present for SpidCode, not even an explicit nil
### GetIvaCode

`func (o *SpidProviderOutput) GetIvaCode() string`

GetIvaCode returns the IvaCode field if non-nil, zero value otherwise.

### GetIvaCodeOk

`func (o *SpidProviderOutput) GetIvaCodeOk() (*string, bool)`

GetIvaCodeOk returns a tuple with the IvaCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIvaCode

`func (o *SpidProviderOutput) SetIvaCode(v string)`

SetIvaCode sets IvaCode field to given value.

### HasIvaCode

`func (o *SpidProviderOutput) HasIvaCode() bool`

HasIvaCode returns a boolean if a field has been set.

### SetIvaCodeNil

`func (o *SpidProviderOutput) SetIvaCodeNil(b bool)`

 SetIvaCodeNil sets the value for IvaCode to be an explicit nil

### UnsetIvaCode
`func (o *SpidProviderOutput) UnsetIvaCode()`

UnsetIvaCode ensures that no value is present for IvaCode, not even an explicit nil
### GetSpidCredentialExpirationDate

`func (o *SpidProviderOutput) GetSpidCredentialExpirationDate() string`

GetSpidCredentialExpirationDate returns the SpidCredentialExpirationDate field if non-nil, zero value otherwise.

### GetSpidCredentialExpirationDateOk

`func (o *SpidProviderOutput) GetSpidCredentialExpirationDateOk() (*string, bool)`

GetSpidCredentialExpirationDateOk returns a tuple with the SpidCredentialExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpidCredentialExpirationDate

`func (o *SpidProviderOutput) SetSpidCredentialExpirationDate(v string)`

SetSpidCredentialExpirationDate sets SpidCredentialExpirationDate field to given value.

### HasSpidCredentialExpirationDate

`func (o *SpidProviderOutput) HasSpidCredentialExpirationDate() bool`

HasSpidCredentialExpirationDate returns a boolean if a field has been set.

### SetSpidCredentialExpirationDateNil

`func (o *SpidProviderOutput) SetSpidCredentialExpirationDateNil(b bool)`

 SetSpidCredentialExpirationDateNil sets the value for SpidCredentialExpirationDate to be an explicit nil

### UnsetSpidCredentialExpirationDate
`func (o *SpidProviderOutput) UnsetSpidCredentialExpirationDate()`

UnsetSpidCredentialExpirationDate ensures that no value is present for SpidCredentialExpirationDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


