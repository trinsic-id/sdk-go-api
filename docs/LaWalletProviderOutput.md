# LaWalletProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DriversLicenseNumber** | Pointer to **NullableString** | The number of the driver&#39;s license used to create the LA Wallet credential | [optional] 
**IssueDate** | Pointer to **NullableString** | The issue date of the driver&#39;s license used to create the LA Wallet credential | [optional] 
**ExpirationDate** | Pointer to **NullableString** | The expiration date of the driver&#39;s license used to create the LA Wallet credential | [optional] 
**AuditNumber** | Pointer to **NullableString** | The 4-digit audit number of the driver&#39;s license used to create the LA Wallet credential | [optional] 
**LicenseStatus** | Pointer to **NullableString** | The license status from the LA Wallet credential | [optional] 
**LicenseClass** | Pointer to **NullableString** | The license class from the LA Wallet credential              Possible values: - \&quot;A\&quot;: Commercial Driver&#39;s License, Combination Vehicles - \&quot;B\&quot;: Commercial Driver&#39;s License, Heavy Straight Vehicle - \&quot;C\&quot;: Commercial Driver&#39;s License, Light Straight Vehicle - \&quot;D\&quot;: Chauffeur&#39;s Driver&#39;s License - \&quot;E\&quot;: Driver&#39;s License for Personal Vehicle | [optional] 
**FirstName** | Pointer to **NullableString** | The first name from the LA Wallet credential | [optional] 
**MiddleName** | Pointer to **NullableString** | The middle name from the LA Wallet credential | [optional] 
**LastName** | Pointer to **NullableString** | The last name from the LA Wallet credential | [optional] 
**DateOfBirth** | Pointer to **NullableString** | The date of birth from the LA Wallet credential | [optional] 
**Sex** | Pointer to **NullableString** | The sex from the LA Wallet credential | [optional] 
**AddressLine1** | Pointer to **NullableString** | The address&#39; line 1 from the LA Wallet credential | [optional] 
**AddressLine2** | Pointer to **NullableString** | The address&#39; line 2 from the LA Wallet credential | [optional] 
**AddressCity** | Pointer to **NullableString** | The address&#39; city from the LA Wallet credential | [optional] 
**AddressState** | Pointer to **NullableString** | The address&#39; state from the LA Wallet credential | [optional] 
**AddressZip** | Pointer to **NullableString** | The address&#39; ZIP from the LA Wallet credential | [optional] 
**County** | Pointer to **NullableString** | The county (\&quot;parish\&quot;) code from the LA Wallet credential.              This is a number from 1 to 64, representing one of Louisiana&#39;s 64 parishes. | [optional] 
**CoarseAge** | Pointer to **NullableString** | The coarse age returned by LA Wallet for this credential              Possible values: - \&quot;Under 18\&quot; - \&quot;Under 21\&quot; - \&quot;Over 21\&quot; | [optional] 

## Methods

### NewLaWalletProviderOutput

`func NewLaWalletProviderOutput() *LaWalletProviderOutput`

NewLaWalletProviderOutput instantiates a new LaWalletProviderOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLaWalletProviderOutputWithDefaults

`func NewLaWalletProviderOutputWithDefaults() *LaWalletProviderOutput`

NewLaWalletProviderOutputWithDefaults instantiates a new LaWalletProviderOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDriversLicenseNumber

`func (o *LaWalletProviderOutput) GetDriversLicenseNumber() string`

GetDriversLicenseNumber returns the DriversLicenseNumber field if non-nil, zero value otherwise.

### GetDriversLicenseNumberOk

`func (o *LaWalletProviderOutput) GetDriversLicenseNumberOk() (*string, bool)`

GetDriversLicenseNumberOk returns a tuple with the DriversLicenseNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriversLicenseNumber

`func (o *LaWalletProviderOutput) SetDriversLicenseNumber(v string)`

SetDriversLicenseNumber sets DriversLicenseNumber field to given value.

### HasDriversLicenseNumber

`func (o *LaWalletProviderOutput) HasDriversLicenseNumber() bool`

HasDriversLicenseNumber returns a boolean if a field has been set.

### SetDriversLicenseNumberNil

`func (o *LaWalletProviderOutput) SetDriversLicenseNumberNil(b bool)`

 SetDriversLicenseNumberNil sets the value for DriversLicenseNumber to be an explicit nil

### UnsetDriversLicenseNumber
`func (o *LaWalletProviderOutput) UnsetDriversLicenseNumber()`

UnsetDriversLicenseNumber ensures that no value is present for DriversLicenseNumber, not even an explicit nil
### GetIssueDate

`func (o *LaWalletProviderOutput) GetIssueDate() string`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *LaWalletProviderOutput) GetIssueDateOk() (*string, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *LaWalletProviderOutput) SetIssueDate(v string)`

SetIssueDate sets IssueDate field to given value.

### HasIssueDate

`func (o *LaWalletProviderOutput) HasIssueDate() bool`

HasIssueDate returns a boolean if a field has been set.

### SetIssueDateNil

`func (o *LaWalletProviderOutput) SetIssueDateNil(b bool)`

 SetIssueDateNil sets the value for IssueDate to be an explicit nil

### UnsetIssueDate
`func (o *LaWalletProviderOutput) UnsetIssueDate()`

UnsetIssueDate ensures that no value is present for IssueDate, not even an explicit nil
### GetExpirationDate

`func (o *LaWalletProviderOutput) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *LaWalletProviderOutput) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *LaWalletProviderOutput) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *LaWalletProviderOutput) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### SetExpirationDateNil

`func (o *LaWalletProviderOutput) SetExpirationDateNil(b bool)`

 SetExpirationDateNil sets the value for ExpirationDate to be an explicit nil

### UnsetExpirationDate
`func (o *LaWalletProviderOutput) UnsetExpirationDate()`

UnsetExpirationDate ensures that no value is present for ExpirationDate, not even an explicit nil
### GetAuditNumber

`func (o *LaWalletProviderOutput) GetAuditNumber() string`

GetAuditNumber returns the AuditNumber field if non-nil, zero value otherwise.

### GetAuditNumberOk

`func (o *LaWalletProviderOutput) GetAuditNumberOk() (*string, bool)`

GetAuditNumberOk returns a tuple with the AuditNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditNumber

`func (o *LaWalletProviderOutput) SetAuditNumber(v string)`

SetAuditNumber sets AuditNumber field to given value.

### HasAuditNumber

`func (o *LaWalletProviderOutput) HasAuditNumber() bool`

HasAuditNumber returns a boolean if a field has been set.

### SetAuditNumberNil

`func (o *LaWalletProviderOutput) SetAuditNumberNil(b bool)`

 SetAuditNumberNil sets the value for AuditNumber to be an explicit nil

### UnsetAuditNumber
`func (o *LaWalletProviderOutput) UnsetAuditNumber()`

UnsetAuditNumber ensures that no value is present for AuditNumber, not even an explicit nil
### GetLicenseStatus

`func (o *LaWalletProviderOutput) GetLicenseStatus() string`

GetLicenseStatus returns the LicenseStatus field if non-nil, zero value otherwise.

### GetLicenseStatusOk

`func (o *LaWalletProviderOutput) GetLicenseStatusOk() (*string, bool)`

GetLicenseStatusOk returns a tuple with the LicenseStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseStatus

`func (o *LaWalletProviderOutput) SetLicenseStatus(v string)`

SetLicenseStatus sets LicenseStatus field to given value.

### HasLicenseStatus

`func (o *LaWalletProviderOutput) HasLicenseStatus() bool`

HasLicenseStatus returns a boolean if a field has been set.

### SetLicenseStatusNil

`func (o *LaWalletProviderOutput) SetLicenseStatusNil(b bool)`

 SetLicenseStatusNil sets the value for LicenseStatus to be an explicit nil

### UnsetLicenseStatus
`func (o *LaWalletProviderOutput) UnsetLicenseStatus()`

UnsetLicenseStatus ensures that no value is present for LicenseStatus, not even an explicit nil
### GetLicenseClass

`func (o *LaWalletProviderOutput) GetLicenseClass() string`

GetLicenseClass returns the LicenseClass field if non-nil, zero value otherwise.

### GetLicenseClassOk

`func (o *LaWalletProviderOutput) GetLicenseClassOk() (*string, bool)`

GetLicenseClassOk returns a tuple with the LicenseClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseClass

`func (o *LaWalletProviderOutput) SetLicenseClass(v string)`

SetLicenseClass sets LicenseClass field to given value.

### HasLicenseClass

`func (o *LaWalletProviderOutput) HasLicenseClass() bool`

HasLicenseClass returns a boolean if a field has been set.

### SetLicenseClassNil

`func (o *LaWalletProviderOutput) SetLicenseClassNil(b bool)`

 SetLicenseClassNil sets the value for LicenseClass to be an explicit nil

### UnsetLicenseClass
`func (o *LaWalletProviderOutput) UnsetLicenseClass()`

UnsetLicenseClass ensures that no value is present for LicenseClass, not even an explicit nil
### GetFirstName

`func (o *LaWalletProviderOutput) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *LaWalletProviderOutput) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *LaWalletProviderOutput) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *LaWalletProviderOutput) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *LaWalletProviderOutput) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *LaWalletProviderOutput) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetMiddleName

`func (o *LaWalletProviderOutput) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *LaWalletProviderOutput) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *LaWalletProviderOutput) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *LaWalletProviderOutput) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### SetMiddleNameNil

`func (o *LaWalletProviderOutput) SetMiddleNameNil(b bool)`

 SetMiddleNameNil sets the value for MiddleName to be an explicit nil

### UnsetMiddleName
`func (o *LaWalletProviderOutput) UnsetMiddleName()`

UnsetMiddleName ensures that no value is present for MiddleName, not even an explicit nil
### GetLastName

`func (o *LaWalletProviderOutput) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *LaWalletProviderOutput) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *LaWalletProviderOutput) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *LaWalletProviderOutput) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *LaWalletProviderOutput) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *LaWalletProviderOutput) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetDateOfBirth

`func (o *LaWalletProviderOutput) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *LaWalletProviderOutput) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *LaWalletProviderOutput) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *LaWalletProviderOutput) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### SetDateOfBirthNil

`func (o *LaWalletProviderOutput) SetDateOfBirthNil(b bool)`

 SetDateOfBirthNil sets the value for DateOfBirth to be an explicit nil

### UnsetDateOfBirth
`func (o *LaWalletProviderOutput) UnsetDateOfBirth()`

UnsetDateOfBirth ensures that no value is present for DateOfBirth, not even an explicit nil
### GetSex

`func (o *LaWalletProviderOutput) GetSex() string`

GetSex returns the Sex field if non-nil, zero value otherwise.

### GetSexOk

`func (o *LaWalletProviderOutput) GetSexOk() (*string, bool)`

GetSexOk returns a tuple with the Sex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSex

`func (o *LaWalletProviderOutput) SetSex(v string)`

SetSex sets Sex field to given value.

### HasSex

`func (o *LaWalletProviderOutput) HasSex() bool`

HasSex returns a boolean if a field has been set.

### SetSexNil

`func (o *LaWalletProviderOutput) SetSexNil(b bool)`

 SetSexNil sets the value for Sex to be an explicit nil

### UnsetSex
`func (o *LaWalletProviderOutput) UnsetSex()`

UnsetSex ensures that no value is present for Sex, not even an explicit nil
### GetAddressLine1

`func (o *LaWalletProviderOutput) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *LaWalletProviderOutput) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *LaWalletProviderOutput) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.

### HasAddressLine1

`func (o *LaWalletProviderOutput) HasAddressLine1() bool`

HasAddressLine1 returns a boolean if a field has been set.

### SetAddressLine1Nil

`func (o *LaWalletProviderOutput) SetAddressLine1Nil(b bool)`

 SetAddressLine1Nil sets the value for AddressLine1 to be an explicit nil

### UnsetAddressLine1
`func (o *LaWalletProviderOutput) UnsetAddressLine1()`

UnsetAddressLine1 ensures that no value is present for AddressLine1, not even an explicit nil
### GetAddressLine2

`func (o *LaWalletProviderOutput) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *LaWalletProviderOutput) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *LaWalletProviderOutput) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *LaWalletProviderOutput) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### SetAddressLine2Nil

`func (o *LaWalletProviderOutput) SetAddressLine2Nil(b bool)`

 SetAddressLine2Nil sets the value for AddressLine2 to be an explicit nil

### UnsetAddressLine2
`func (o *LaWalletProviderOutput) UnsetAddressLine2()`

UnsetAddressLine2 ensures that no value is present for AddressLine2, not even an explicit nil
### GetAddressCity

`func (o *LaWalletProviderOutput) GetAddressCity() string`

GetAddressCity returns the AddressCity field if non-nil, zero value otherwise.

### GetAddressCityOk

`func (o *LaWalletProviderOutput) GetAddressCityOk() (*string, bool)`

GetAddressCityOk returns a tuple with the AddressCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressCity

`func (o *LaWalletProviderOutput) SetAddressCity(v string)`

SetAddressCity sets AddressCity field to given value.

### HasAddressCity

`func (o *LaWalletProviderOutput) HasAddressCity() bool`

HasAddressCity returns a boolean if a field has been set.

### SetAddressCityNil

`func (o *LaWalletProviderOutput) SetAddressCityNil(b bool)`

 SetAddressCityNil sets the value for AddressCity to be an explicit nil

### UnsetAddressCity
`func (o *LaWalletProviderOutput) UnsetAddressCity()`

UnsetAddressCity ensures that no value is present for AddressCity, not even an explicit nil
### GetAddressState

`func (o *LaWalletProviderOutput) GetAddressState() string`

GetAddressState returns the AddressState field if non-nil, zero value otherwise.

### GetAddressStateOk

`func (o *LaWalletProviderOutput) GetAddressStateOk() (*string, bool)`

GetAddressStateOk returns a tuple with the AddressState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressState

`func (o *LaWalletProviderOutput) SetAddressState(v string)`

SetAddressState sets AddressState field to given value.

### HasAddressState

`func (o *LaWalletProviderOutput) HasAddressState() bool`

HasAddressState returns a boolean if a field has been set.

### SetAddressStateNil

`func (o *LaWalletProviderOutput) SetAddressStateNil(b bool)`

 SetAddressStateNil sets the value for AddressState to be an explicit nil

### UnsetAddressState
`func (o *LaWalletProviderOutput) UnsetAddressState()`

UnsetAddressState ensures that no value is present for AddressState, not even an explicit nil
### GetAddressZip

`func (o *LaWalletProviderOutput) GetAddressZip() string`

GetAddressZip returns the AddressZip field if non-nil, zero value otherwise.

### GetAddressZipOk

`func (o *LaWalletProviderOutput) GetAddressZipOk() (*string, bool)`

GetAddressZipOk returns a tuple with the AddressZip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressZip

`func (o *LaWalletProviderOutput) SetAddressZip(v string)`

SetAddressZip sets AddressZip field to given value.

### HasAddressZip

`func (o *LaWalletProviderOutput) HasAddressZip() bool`

HasAddressZip returns a boolean if a field has been set.

### SetAddressZipNil

`func (o *LaWalletProviderOutput) SetAddressZipNil(b bool)`

 SetAddressZipNil sets the value for AddressZip to be an explicit nil

### UnsetAddressZip
`func (o *LaWalletProviderOutput) UnsetAddressZip()`

UnsetAddressZip ensures that no value is present for AddressZip, not even an explicit nil
### GetCounty

`func (o *LaWalletProviderOutput) GetCounty() string`

GetCounty returns the County field if non-nil, zero value otherwise.

### GetCountyOk

`func (o *LaWalletProviderOutput) GetCountyOk() (*string, bool)`

GetCountyOk returns a tuple with the County field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounty

`func (o *LaWalletProviderOutput) SetCounty(v string)`

SetCounty sets County field to given value.

### HasCounty

`func (o *LaWalletProviderOutput) HasCounty() bool`

HasCounty returns a boolean if a field has been set.

### SetCountyNil

`func (o *LaWalletProviderOutput) SetCountyNil(b bool)`

 SetCountyNil sets the value for County to be an explicit nil

### UnsetCounty
`func (o *LaWalletProviderOutput) UnsetCounty()`

UnsetCounty ensures that no value is present for County, not even an explicit nil
### GetCoarseAge

`func (o *LaWalletProviderOutput) GetCoarseAge() string`

GetCoarseAge returns the CoarseAge field if non-nil, zero value otherwise.

### GetCoarseAgeOk

`func (o *LaWalletProviderOutput) GetCoarseAgeOk() (*string, bool)`

GetCoarseAgeOk returns a tuple with the CoarseAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoarseAge

`func (o *LaWalletProviderOutput) SetCoarseAge(v string)`

SetCoarseAge sets CoarseAge field to given value.

### HasCoarseAge

`func (o *LaWalletProviderOutput) HasCoarseAge() bool`

HasCoarseAge returns a boolean if a field has been set.

### SetCoarseAgeNil

`func (o *LaWalletProviderOutput) SetCoarseAgeNil(b bool)`

 SetCoarseAgeNil sets the value for CoarseAge to be an explicit nil

### UnsetCoarseAge
`func (o *LaWalletProviderOutput) UnsetCoarseAge()`

UnsetCoarseAge ensures that no value is present for CoarseAge, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


