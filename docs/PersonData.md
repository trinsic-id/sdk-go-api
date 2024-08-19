# PersonData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GivenName** | Pointer to **string** | Given (first) name of the individual | [optional] 
**FamilyName** | Pointer to **string** | Family (last) name of the individual | [optional] 
**MiddleName** | Pointer to **string** | Middle name of the individual | [optional] 
**FullName** | Pointer to **string** | The individual&#39;s full name as a single string.                Useful for names which do not fit into a \&quot;first middle last\&quot; structure. | [optional] 
**Nationality** | Pointer to **string** |  | [optional] 
**Gender** | Pointer to **string** |  | [optional] 
**PhoneNumber** | Pointer to **string** |  | [optional] 
**Address** | Pointer to [**Address**](Address.md) | Address information for an individual | [optional] 
**DateOfBirth** | Pointer to **string** |  | [optional] 

## Methods

### NewPersonData

`func NewPersonData() *PersonData`

NewPersonData instantiates a new PersonData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersonDataWithDefaults

`func NewPersonDataWithDefaults() *PersonData`

NewPersonDataWithDefaults instantiates a new PersonData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGivenName

`func (o *PersonData) GetGivenName() string`

GetGivenName returns the GivenName field if non-nil, zero value otherwise.

### GetGivenNameOk

`func (o *PersonData) GetGivenNameOk() (*string, bool)`

GetGivenNameOk returns a tuple with the GivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenName

`func (o *PersonData) SetGivenName(v string)`

SetGivenName sets GivenName field to given value.

### HasGivenName

`func (o *PersonData) HasGivenName() bool`

HasGivenName returns a boolean if a field has been set.

### GetFamilyName

`func (o *PersonData) GetFamilyName() string`

GetFamilyName returns the FamilyName field if non-nil, zero value otherwise.

### GetFamilyNameOk

`func (o *PersonData) GetFamilyNameOk() (*string, bool)`

GetFamilyNameOk returns a tuple with the FamilyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyName

`func (o *PersonData) SetFamilyName(v string)`

SetFamilyName sets FamilyName field to given value.

### HasFamilyName

`func (o *PersonData) HasFamilyName() bool`

HasFamilyName returns a boolean if a field has been set.

### GetMiddleName

`func (o *PersonData) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *PersonData) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *PersonData) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *PersonData) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### GetFullName

`func (o *PersonData) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *PersonData) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *PersonData) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *PersonData) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetNationality

`func (o *PersonData) GetNationality() string`

GetNationality returns the Nationality field if non-nil, zero value otherwise.

### GetNationalityOk

`func (o *PersonData) GetNationalityOk() (*string, bool)`

GetNationalityOk returns a tuple with the Nationality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationality

`func (o *PersonData) SetNationality(v string)`

SetNationality sets Nationality field to given value.

### HasNationality

`func (o *PersonData) HasNationality() bool`

HasNationality returns a boolean if a field has been set.

### GetGender

`func (o *PersonData) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *PersonData) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *PersonData) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *PersonData) HasGender() bool`

HasGender returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *PersonData) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *PersonData) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *PersonData) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *PersonData) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetAddress

`func (o *PersonData) GetAddress() Address`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *PersonData) GetAddressOk() (*Address, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *PersonData) SetAddress(v Address)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *PersonData) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetDateOfBirth

`func (o *PersonData) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *PersonData) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *PersonData) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *PersonData) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


