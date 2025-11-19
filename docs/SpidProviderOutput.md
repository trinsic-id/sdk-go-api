# SpidProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingInformation** | Pointer to [**NullableSpidBillingInformation**](SpidBillingInformation.md) | Information about the billable status of this SPID Verification.              Present only if your account has period-based billing enabled for SPID. Contact Trinsic to enable this. | [optional] 
**IdentityProviderEntityId** | **string** | The SPID Entity ID of the Identity Provider which issued the SPID identity.              This is an HTTPS URI which uniquely identifies the IdP within the SPID federation.              A normalized / simplified representation of this value is present in the &#x60;originatingSubProviderId&#x60; field in Trinsic&#39;s normalized data model. | 
**SpidCode** | **string** | The identifier of the user&#39;s SPID credential.              This uniquely identifies the credential within the SPID federation. | 
**SpidCredentialExpirationDate** | Pointer to **NullableString** | Expiration date of the user&#39;s SPID credential.              This is not the same as the expiration date of the underlying identity document (such as a passport) which was used to create the SPID identity. | [optional] 
**PlaceOfBirth** | Pointer to **NullableString** | The user&#39;s place of birth. | [optional] 
**CountyOfBirth** | Pointer to **NullableString** | The user&#39;s county of birth. | [optional] 
**RawIdCard** | Pointer to **NullableString** | The raw, space-separated string value for the \&quot;IdCard\&quot; field from the SPID identity.              Trinsic additionally parses this field and uses it to populate the &#x60;Document&#x60; object in the normalized data model. | [optional] 
**Email** | Pointer to **NullableString** | The email address of the user. | [optional] 
**DigitalAddress** | Pointer to **NullableString** | The digital address of the user. | [optional] 
**FiscalNumber** | Pointer to **NullableString** | Fiscal tax number for the subject. | [optional] 
**IvaCode** | Pointer to **NullableString** | VAT number for the subject. | [optional] 
**CompanyName** | Pointer to **NullableString** | The name of the company which the user is associated with. | [optional] 
**CompanyFiscalNumber** | Pointer to **NullableString** | The fiscal tax number of the company which the user is associated with. | [optional] 
**RegisteredOffice** | Pointer to **NullableString** | The registered office address of the company which the user is associated with. | [optional] 

## Methods

### NewSpidProviderOutput

`func NewSpidProviderOutput(identityProviderEntityId string, spidCode string, ) *SpidProviderOutput`

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
### GetIdentityProviderEntityId

`func (o *SpidProviderOutput) GetIdentityProviderEntityId() string`

GetIdentityProviderEntityId returns the IdentityProviderEntityId field if non-nil, zero value otherwise.

### GetIdentityProviderEntityIdOk

`func (o *SpidProviderOutput) GetIdentityProviderEntityIdOk() (*string, bool)`

GetIdentityProviderEntityIdOk returns a tuple with the IdentityProviderEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityProviderEntityId

`func (o *SpidProviderOutput) SetIdentityProviderEntityId(v string)`

SetIdentityProviderEntityId sets IdentityProviderEntityId field to given value.


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
### GetPlaceOfBirth

`func (o *SpidProviderOutput) GetPlaceOfBirth() string`

GetPlaceOfBirth returns the PlaceOfBirth field if non-nil, zero value otherwise.

### GetPlaceOfBirthOk

`func (o *SpidProviderOutput) GetPlaceOfBirthOk() (*string, bool)`

GetPlaceOfBirthOk returns a tuple with the PlaceOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceOfBirth

`func (o *SpidProviderOutput) SetPlaceOfBirth(v string)`

SetPlaceOfBirth sets PlaceOfBirth field to given value.

### HasPlaceOfBirth

`func (o *SpidProviderOutput) HasPlaceOfBirth() bool`

HasPlaceOfBirth returns a boolean if a field has been set.

### SetPlaceOfBirthNil

`func (o *SpidProviderOutput) SetPlaceOfBirthNil(b bool)`

 SetPlaceOfBirthNil sets the value for PlaceOfBirth to be an explicit nil

### UnsetPlaceOfBirth
`func (o *SpidProviderOutput) UnsetPlaceOfBirth()`

UnsetPlaceOfBirth ensures that no value is present for PlaceOfBirth, not even an explicit nil
### GetCountyOfBirth

`func (o *SpidProviderOutput) GetCountyOfBirth() string`

GetCountyOfBirth returns the CountyOfBirth field if non-nil, zero value otherwise.

### GetCountyOfBirthOk

`func (o *SpidProviderOutput) GetCountyOfBirthOk() (*string, bool)`

GetCountyOfBirthOk returns a tuple with the CountyOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountyOfBirth

`func (o *SpidProviderOutput) SetCountyOfBirth(v string)`

SetCountyOfBirth sets CountyOfBirth field to given value.

### HasCountyOfBirth

`func (o *SpidProviderOutput) HasCountyOfBirth() bool`

HasCountyOfBirth returns a boolean if a field has been set.

### SetCountyOfBirthNil

`func (o *SpidProviderOutput) SetCountyOfBirthNil(b bool)`

 SetCountyOfBirthNil sets the value for CountyOfBirth to be an explicit nil

### UnsetCountyOfBirth
`func (o *SpidProviderOutput) UnsetCountyOfBirth()`

UnsetCountyOfBirth ensures that no value is present for CountyOfBirth, not even an explicit nil
### GetRawIdCard

`func (o *SpidProviderOutput) GetRawIdCard() string`

GetRawIdCard returns the RawIdCard field if non-nil, zero value otherwise.

### GetRawIdCardOk

`func (o *SpidProviderOutput) GetRawIdCardOk() (*string, bool)`

GetRawIdCardOk returns a tuple with the RawIdCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawIdCard

`func (o *SpidProviderOutput) SetRawIdCard(v string)`

SetRawIdCard sets RawIdCard field to given value.

### HasRawIdCard

`func (o *SpidProviderOutput) HasRawIdCard() bool`

HasRawIdCard returns a boolean if a field has been set.

### SetRawIdCardNil

`func (o *SpidProviderOutput) SetRawIdCardNil(b bool)`

 SetRawIdCardNil sets the value for RawIdCard to be an explicit nil

### UnsetRawIdCard
`func (o *SpidProviderOutput) UnsetRawIdCard()`

UnsetRawIdCard ensures that no value is present for RawIdCard, not even an explicit nil
### GetEmail

`func (o *SpidProviderOutput) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *SpidProviderOutput) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *SpidProviderOutput) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *SpidProviderOutput) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *SpidProviderOutput) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *SpidProviderOutput) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetDigitalAddress

`func (o *SpidProviderOutput) GetDigitalAddress() string`

GetDigitalAddress returns the DigitalAddress field if non-nil, zero value otherwise.

### GetDigitalAddressOk

`func (o *SpidProviderOutput) GetDigitalAddressOk() (*string, bool)`

GetDigitalAddressOk returns a tuple with the DigitalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigitalAddress

`func (o *SpidProviderOutput) SetDigitalAddress(v string)`

SetDigitalAddress sets DigitalAddress field to given value.

### HasDigitalAddress

`func (o *SpidProviderOutput) HasDigitalAddress() bool`

HasDigitalAddress returns a boolean if a field has been set.

### SetDigitalAddressNil

`func (o *SpidProviderOutput) SetDigitalAddressNil(b bool)`

 SetDigitalAddressNil sets the value for DigitalAddress to be an explicit nil

### UnsetDigitalAddress
`func (o *SpidProviderOutput) UnsetDigitalAddress()`

UnsetDigitalAddress ensures that no value is present for DigitalAddress, not even an explicit nil
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
### GetCompanyName

`func (o *SpidProviderOutput) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *SpidProviderOutput) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *SpidProviderOutput) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *SpidProviderOutput) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### SetCompanyNameNil

`func (o *SpidProviderOutput) SetCompanyNameNil(b bool)`

 SetCompanyNameNil sets the value for CompanyName to be an explicit nil

### UnsetCompanyName
`func (o *SpidProviderOutput) UnsetCompanyName()`

UnsetCompanyName ensures that no value is present for CompanyName, not even an explicit nil
### GetCompanyFiscalNumber

`func (o *SpidProviderOutput) GetCompanyFiscalNumber() string`

GetCompanyFiscalNumber returns the CompanyFiscalNumber field if non-nil, zero value otherwise.

### GetCompanyFiscalNumberOk

`func (o *SpidProviderOutput) GetCompanyFiscalNumberOk() (*string, bool)`

GetCompanyFiscalNumberOk returns a tuple with the CompanyFiscalNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyFiscalNumber

`func (o *SpidProviderOutput) SetCompanyFiscalNumber(v string)`

SetCompanyFiscalNumber sets CompanyFiscalNumber field to given value.

### HasCompanyFiscalNumber

`func (o *SpidProviderOutput) HasCompanyFiscalNumber() bool`

HasCompanyFiscalNumber returns a boolean if a field has been set.

### SetCompanyFiscalNumberNil

`func (o *SpidProviderOutput) SetCompanyFiscalNumberNil(b bool)`

 SetCompanyFiscalNumberNil sets the value for CompanyFiscalNumber to be an explicit nil

### UnsetCompanyFiscalNumber
`func (o *SpidProviderOutput) UnsetCompanyFiscalNumber()`

UnsetCompanyFiscalNumber ensures that no value is present for CompanyFiscalNumber, not even an explicit nil
### GetRegisteredOffice

`func (o *SpidProviderOutput) GetRegisteredOffice() string`

GetRegisteredOffice returns the RegisteredOffice field if non-nil, zero value otherwise.

### GetRegisteredOfficeOk

`func (o *SpidProviderOutput) GetRegisteredOfficeOk() (*string, bool)`

GetRegisteredOfficeOk returns a tuple with the RegisteredOffice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredOffice

`func (o *SpidProviderOutput) SetRegisteredOffice(v string)`

SetRegisteredOffice sets RegisteredOffice field to given value.

### HasRegisteredOffice

`func (o *SpidProviderOutput) HasRegisteredOffice() bool`

HasRegisteredOffice returns a boolean if a field has been set.

### SetRegisteredOfficeNil

`func (o *SpidProviderOutput) SetRegisteredOfficeNil(b bool)`

 SetRegisteredOfficeNil sets the value for RegisteredOffice to be an explicit nil

### UnsetRegisteredOffice
`func (o *SpidProviderOutput) UnsetRegisteredOffice()`

UnsetRegisteredOffice ensures that no value is present for RegisteredOffice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


