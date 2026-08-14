# NigeriaNinLookup2ProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | Pointer to **NullableString** | The first name (given name) of the ID holder. | [optional] 
**MiddleName** | Pointer to **NullableString** | The middle name of the ID holder. | [optional] 
**Surname** | Pointer to **NullableString** | The surname (family name) of the ID holder. | [optional] 
**Sex** | Pointer to **NullableString** | The sex of the ID holder.              Possible values: - Male - Female | [optional] 
**DateOfBirth** | Pointer to **NullableString** | The date of birth of the ID holder. | [optional] 
**BirthCountry** | Pointer to **NullableString** | Country of birth as an ISO 3166-1 alpha-2 code. | [optional] 
**NationalIdentityNumber** | Pointer to **NullableString** | National Identification Number (NIN).              This is a unique, permanent identifier assigned by the National Identity Management Commission (NIMC) upon enrollment.              Format: - 11 numeric digits - No publicly known encoding scheme is used to encode personal information in the NIN - Last digit is a checksum using the Verhoeff algorithm | [optional] 
**PhoneNumber** | Pointer to **NullableString** | Phone number registered with the National Identity Management Commission.              Format: - Valid Nigerian numbers are returned in international E.164 format   (for example, +2348031234567). - Otherwise, the value is returned in trunk notation (leading zero included)   as provided by the authority (for example, 0123456789). | [optional] 
**Email** | Pointer to **NullableString** | Email address registered with the National Identity Management Commission. | [optional] 
**Address** | Pointer to **NullableString** | Full residential address as a single string, normalized to lowercase without delimiters. | [optional] 
**LocalGovernmentArea** | Pointer to **NullableString** | Local Government Area of residence.              Nigeria is divided into 774 Local Government Areas (LGAs), which are the third-tier administrative divisions below states and the Federal Capital Territory. LGAs are roughly equivalent to counties or municipalities in other countries. | [optional] 
**State** | Pointer to **NullableString** | State of residence. | [optional] 

## Methods

### NewNigeriaNinLookup2ProviderOutput

`func NewNigeriaNinLookup2ProviderOutput() *NigeriaNinLookup2ProviderOutput`

NewNigeriaNinLookup2ProviderOutput instantiates a new NigeriaNinLookup2ProviderOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNigeriaNinLookup2ProviderOutputWithDefaults

`func NewNigeriaNinLookup2ProviderOutputWithDefaults() *NigeriaNinLookup2ProviderOutput`

NewNigeriaNinLookup2ProviderOutputWithDefaults instantiates a new NigeriaNinLookup2ProviderOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *NigeriaNinLookup2ProviderOutput) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *NigeriaNinLookup2ProviderOutput) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *NigeriaNinLookup2ProviderOutput) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *NigeriaNinLookup2ProviderOutput) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *NigeriaNinLookup2ProviderOutput) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *NigeriaNinLookup2ProviderOutput) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetMiddleName

`func (o *NigeriaNinLookup2ProviderOutput) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *NigeriaNinLookup2ProviderOutput) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *NigeriaNinLookup2ProviderOutput) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *NigeriaNinLookup2ProviderOutput) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### SetMiddleNameNil

`func (o *NigeriaNinLookup2ProviderOutput) SetMiddleNameNil(b bool)`

 SetMiddleNameNil sets the value for MiddleName to be an explicit nil

### UnsetMiddleName
`func (o *NigeriaNinLookup2ProviderOutput) UnsetMiddleName()`

UnsetMiddleName ensures that no value is present for MiddleName, not even an explicit nil
### GetSurname

`func (o *NigeriaNinLookup2ProviderOutput) GetSurname() string`

GetSurname returns the Surname field if non-nil, zero value otherwise.

### GetSurnameOk

`func (o *NigeriaNinLookup2ProviderOutput) GetSurnameOk() (*string, bool)`

GetSurnameOk returns a tuple with the Surname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurname

`func (o *NigeriaNinLookup2ProviderOutput) SetSurname(v string)`

SetSurname sets Surname field to given value.

### HasSurname

`func (o *NigeriaNinLookup2ProviderOutput) HasSurname() bool`

HasSurname returns a boolean if a field has been set.

### SetSurnameNil

`func (o *NigeriaNinLookup2ProviderOutput) SetSurnameNil(b bool)`

 SetSurnameNil sets the value for Surname to be an explicit nil

### UnsetSurname
`func (o *NigeriaNinLookup2ProviderOutput) UnsetSurname()`

UnsetSurname ensures that no value is present for Surname, not even an explicit nil
### GetSex

`func (o *NigeriaNinLookup2ProviderOutput) GetSex() string`

GetSex returns the Sex field if non-nil, zero value otherwise.

### GetSexOk

`func (o *NigeriaNinLookup2ProviderOutput) GetSexOk() (*string, bool)`

GetSexOk returns a tuple with the Sex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSex

`func (o *NigeriaNinLookup2ProviderOutput) SetSex(v string)`

SetSex sets Sex field to given value.

### HasSex

`func (o *NigeriaNinLookup2ProviderOutput) HasSex() bool`

HasSex returns a boolean if a field has been set.

### SetSexNil

`func (o *NigeriaNinLookup2ProviderOutput) SetSexNil(b bool)`

 SetSexNil sets the value for Sex to be an explicit nil

### UnsetSex
`func (o *NigeriaNinLookup2ProviderOutput) UnsetSex()`

UnsetSex ensures that no value is present for Sex, not even an explicit nil
### GetDateOfBirth

`func (o *NigeriaNinLookup2ProviderOutput) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *NigeriaNinLookup2ProviderOutput) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *NigeriaNinLookup2ProviderOutput) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *NigeriaNinLookup2ProviderOutput) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### SetDateOfBirthNil

`func (o *NigeriaNinLookup2ProviderOutput) SetDateOfBirthNil(b bool)`

 SetDateOfBirthNil sets the value for DateOfBirth to be an explicit nil

### UnsetDateOfBirth
`func (o *NigeriaNinLookup2ProviderOutput) UnsetDateOfBirth()`

UnsetDateOfBirth ensures that no value is present for DateOfBirth, not even an explicit nil
### GetBirthCountry

`func (o *NigeriaNinLookup2ProviderOutput) GetBirthCountry() string`

GetBirthCountry returns the BirthCountry field if non-nil, zero value otherwise.

### GetBirthCountryOk

`func (o *NigeriaNinLookup2ProviderOutput) GetBirthCountryOk() (*string, bool)`

GetBirthCountryOk returns a tuple with the BirthCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthCountry

`func (o *NigeriaNinLookup2ProviderOutput) SetBirthCountry(v string)`

SetBirthCountry sets BirthCountry field to given value.

### HasBirthCountry

`func (o *NigeriaNinLookup2ProviderOutput) HasBirthCountry() bool`

HasBirthCountry returns a boolean if a field has been set.

### SetBirthCountryNil

`func (o *NigeriaNinLookup2ProviderOutput) SetBirthCountryNil(b bool)`

 SetBirthCountryNil sets the value for BirthCountry to be an explicit nil

### UnsetBirthCountry
`func (o *NigeriaNinLookup2ProviderOutput) UnsetBirthCountry()`

UnsetBirthCountry ensures that no value is present for BirthCountry, not even an explicit nil
### GetNationalIdentityNumber

`func (o *NigeriaNinLookup2ProviderOutput) GetNationalIdentityNumber() string`

GetNationalIdentityNumber returns the NationalIdentityNumber field if non-nil, zero value otherwise.

### GetNationalIdentityNumberOk

`func (o *NigeriaNinLookup2ProviderOutput) GetNationalIdentityNumberOk() (*string, bool)`

GetNationalIdentityNumberOk returns a tuple with the NationalIdentityNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationalIdentityNumber

`func (o *NigeriaNinLookup2ProviderOutput) SetNationalIdentityNumber(v string)`

SetNationalIdentityNumber sets NationalIdentityNumber field to given value.

### HasNationalIdentityNumber

`func (o *NigeriaNinLookup2ProviderOutput) HasNationalIdentityNumber() bool`

HasNationalIdentityNumber returns a boolean if a field has been set.

### SetNationalIdentityNumberNil

`func (o *NigeriaNinLookup2ProviderOutput) SetNationalIdentityNumberNil(b bool)`

 SetNationalIdentityNumberNil sets the value for NationalIdentityNumber to be an explicit nil

### UnsetNationalIdentityNumber
`func (o *NigeriaNinLookup2ProviderOutput) UnsetNationalIdentityNumber()`

UnsetNationalIdentityNumber ensures that no value is present for NationalIdentityNumber, not even an explicit nil
### GetPhoneNumber

`func (o *NigeriaNinLookup2ProviderOutput) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *NigeriaNinLookup2ProviderOutput) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *NigeriaNinLookup2ProviderOutput) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *NigeriaNinLookup2ProviderOutput) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### SetPhoneNumberNil

`func (o *NigeriaNinLookup2ProviderOutput) SetPhoneNumberNil(b bool)`

 SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil

### UnsetPhoneNumber
`func (o *NigeriaNinLookup2ProviderOutput) UnsetPhoneNumber()`

UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
### GetEmail

`func (o *NigeriaNinLookup2ProviderOutput) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *NigeriaNinLookup2ProviderOutput) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *NigeriaNinLookup2ProviderOutput) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *NigeriaNinLookup2ProviderOutput) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *NigeriaNinLookup2ProviderOutput) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *NigeriaNinLookup2ProviderOutput) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetAddress

`func (o *NigeriaNinLookup2ProviderOutput) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *NigeriaNinLookup2ProviderOutput) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *NigeriaNinLookup2ProviderOutput) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *NigeriaNinLookup2ProviderOutput) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *NigeriaNinLookup2ProviderOutput) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *NigeriaNinLookup2ProviderOutput) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetLocalGovernmentArea

`func (o *NigeriaNinLookup2ProviderOutput) GetLocalGovernmentArea() string`

GetLocalGovernmentArea returns the LocalGovernmentArea field if non-nil, zero value otherwise.

### GetLocalGovernmentAreaOk

`func (o *NigeriaNinLookup2ProviderOutput) GetLocalGovernmentAreaOk() (*string, bool)`

GetLocalGovernmentAreaOk returns a tuple with the LocalGovernmentArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalGovernmentArea

`func (o *NigeriaNinLookup2ProviderOutput) SetLocalGovernmentArea(v string)`

SetLocalGovernmentArea sets LocalGovernmentArea field to given value.

### HasLocalGovernmentArea

`func (o *NigeriaNinLookup2ProviderOutput) HasLocalGovernmentArea() bool`

HasLocalGovernmentArea returns a boolean if a field has been set.

### SetLocalGovernmentAreaNil

`func (o *NigeriaNinLookup2ProviderOutput) SetLocalGovernmentAreaNil(b bool)`

 SetLocalGovernmentAreaNil sets the value for LocalGovernmentArea to be an explicit nil

### UnsetLocalGovernmentArea
`func (o *NigeriaNinLookup2ProviderOutput) UnsetLocalGovernmentArea()`

UnsetLocalGovernmentArea ensures that no value is present for LocalGovernmentArea, not even an explicit nil
### GetState

`func (o *NigeriaNinLookup2ProviderOutput) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *NigeriaNinLookup2ProviderOutput) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *NigeriaNinLookup2ProviderOutput) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *NigeriaNinLookup2ProviderOutput) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *NigeriaNinLookup2ProviderOutput) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *NigeriaNinLookup2ProviderOutput) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


