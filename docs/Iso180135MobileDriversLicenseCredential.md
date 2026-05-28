# Iso180135MobileDriversLicenseCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Standard** | Pointer to [**NullableIso180135StandardNamespaceOutput**](Iso180135StandardNamespaceOutput.md) | Fields from the standard &#x60;org.iso.18013.5.1&#x60; namespace within the Mobile Driver&#39;s License.              This namespace is the primary dataset defined for Mobile Driver&#39;s License credentials. | [optional] 
**Aamva** | Pointer to [**NullableIso180135AamvaNamespaceOutput**](Iso180135AamvaNamespaceOutput.md) | Fields from the &#x60;org.iso.18013.5.1.aamva&#x60; namespace.              This namespace is an additional dataset defined by AAMVA (American Association of Motor Vehicle Administrators). | [optional] 

## Methods

### NewIso180135MobileDriversLicenseCredential

`func NewIso180135MobileDriversLicenseCredential() *Iso180135MobileDriversLicenseCredential`

NewIso180135MobileDriversLicenseCredential instantiates a new Iso180135MobileDriversLicenseCredential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIso180135MobileDriversLicenseCredentialWithDefaults

`func NewIso180135MobileDriversLicenseCredentialWithDefaults() *Iso180135MobileDriversLicenseCredential`

NewIso180135MobileDriversLicenseCredentialWithDefaults instantiates a new Iso180135MobileDriversLicenseCredential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStandard

`func (o *Iso180135MobileDriversLicenseCredential) GetStandard() Iso180135StandardNamespaceOutput`

GetStandard returns the Standard field if non-nil, zero value otherwise.

### GetStandardOk

`func (o *Iso180135MobileDriversLicenseCredential) GetStandardOk() (*Iso180135StandardNamespaceOutput, bool)`

GetStandardOk returns a tuple with the Standard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandard

`func (o *Iso180135MobileDriversLicenseCredential) SetStandard(v Iso180135StandardNamespaceOutput)`

SetStandard sets Standard field to given value.

### HasStandard

`func (o *Iso180135MobileDriversLicenseCredential) HasStandard() bool`

HasStandard returns a boolean if a field has been set.

### SetStandardNil

`func (o *Iso180135MobileDriversLicenseCredential) SetStandardNil(b bool)`

 SetStandardNil sets the value for Standard to be an explicit nil

### UnsetStandard
`func (o *Iso180135MobileDriversLicenseCredential) UnsetStandard()`

UnsetStandard ensures that no value is present for Standard, not even an explicit nil
### GetAamva

`func (o *Iso180135MobileDriversLicenseCredential) GetAamva() Iso180135AamvaNamespaceOutput`

GetAamva returns the Aamva field if non-nil, zero value otherwise.

### GetAamvaOk

`func (o *Iso180135MobileDriversLicenseCredential) GetAamvaOk() (*Iso180135AamvaNamespaceOutput, bool)`

GetAamvaOk returns a tuple with the Aamva field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAamva

`func (o *Iso180135MobileDriversLicenseCredential) SetAamva(v Iso180135AamvaNamespaceOutput)`

SetAamva sets Aamva field to given value.

### HasAamva

`func (o *Iso180135MobileDriversLicenseCredential) HasAamva() bool`

HasAamva returns a boolean if a field has been set.

### SetAamvaNil

`func (o *Iso180135MobileDriversLicenseCredential) SetAamvaNil(b bool)`

 SetAamvaNil sets the value for Aamva to be an explicit nil

### UnsetAamva
`func (o *Iso180135MobileDriversLicenseCredential) UnsetAamva()`

UnsetAamva ensures that no value is present for Aamva, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


