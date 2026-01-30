# LaWalletProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DriversLicenseNumber** | **string** | The number of the driver&#39;s license used to create the LA Wallet credential | 
**IssueDate** | **string** | The issue date of the driver&#39;s license used to create the LA Wallet credential | 
**ExpirationDate** | **string** | The expiration date of the driver&#39;s license used to create the LA Wallet credential | 
**AuditNumber** | **string** | The 4-digit audit number of the driver&#39;s license used to create the LA Wallet credential | 
**LicenseStatus** | **string** | The license status from the LA Wallet credential | 
**LicenseClass** | **string** | The license class from the LA Wallet credential              Possible values: - \&quot;A\&quot;: Commercial Driver&#39;s License, Combination Vehicles - \&quot;B\&quot;: Commercial Driver&#39;s License, Heavy Straight Vehicle - \&quot;C\&quot;: Commercial Driver&#39;s License, Light Straight Vehicle - \&quot;D\&quot;: Chauffeur&#39;s Driver&#39;s License - \&quot;E\&quot;: Driver&#39;s License for Personal Vehicle | 
**FirstName** | **string** | The first name from the LA Wallet credential | 
**MiddleName** | **string** | The middle name from the LA Wallet credential | 
**LastName** | **string** | The last name from the LA Wallet credential | 
**DateOfBirth** | **string** | The date of birth from the LA Wallet credential | 
**Sex** | **string** | The sex from the LA Wallet credential | 
**AddressLine1** | **string** | The address&#39; line 1 from the LA Wallet credential | 
**AddressLine2** | Pointer to **NullableString** | The address&#39; line 2 from the LA Wallet credential | [optional] 
**AddressCity** | **string** | The address&#39; city from the LA Wallet credential | 
**AddressState** | **string** | The address&#39; state from the LA Wallet credential | 
**AddressZip** | **string** | The address&#39; ZIP from the LA Wallet credential | 
**County** | **string** | The county (\&quot;parish\&quot;) code from the LA Wallet credential.              This is a number from 1 to 64, representing one of Louisiana&#39;s 64 parishes. | 
**CoarseAge** | **string** | The coarse age returned by LA Wallet for this credential              Possible values: - \&quot;Under 18\&quot; - \&quot;Under 21\&quot; - \&quot;Over 21\&quot; | 

## Methods

### NewLaWalletProviderOutput

`func NewLaWalletProviderOutput(driversLicenseNumber string, issueDate string, expirationDate string, auditNumber string, licenseStatus string, licenseClass string, firstName string, middleName string, lastName string, dateOfBirth string, sex string, addressLine1 string, addressCity string, addressState string, addressZip string, county string, coarseAge string, ) *LaWalletProviderOutput`

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


