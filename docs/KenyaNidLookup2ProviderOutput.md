# KenyaNidLookup2ProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | Pointer to **NullableString** | The first name (given name) of the ID holder as recorded in IPRS. | [optional] 
**Surname** | Pointer to **NullableString** | The surname (family name) of the ID holder as recorded in IPRS. | [optional] 
**OtherName** | Pointer to **NullableString** | The other name (middle name) of the ID holder as recorded in IPRS. | [optional] 
**Sex** | Pointer to **NullableString** | The sex of the ID holder as recorded on the National ID.              Possible values: - Male - Female | [optional] 
**DateOfBirth** | Pointer to **NullableString** | The date of birth of the ID holder as recorded in IPRS. | [optional] 
**Citizenship** | Pointer to **NullableString** | Citizenship status as recorded in the IPRS civil registry database.              For the Kenya National ID lookup, this value will always be \&quot;Kenyan\&quot; as the National ID is only issued to Kenyan citizens. Non-citizens residing in Kenya are issued different identification documents (Alien ID cards, refugee documentation, etc.) which are not supported by this provider. | [optional] 
**IdNumber** | Pointer to **NullableString** | The Kenya National ID Number (Nambari ya Kitambulisho) or Unique Personal Identifier (Maisha Namba).              This is the primary unique identifier for Kenyan citizens in all government systems, issued by the National Registration Bureau (NRB). The format is either 8 digits for National ID or 9 digits for Maisha Namba UPI (the new format since 2023). | [optional] 
**SerialNumber** | Pointer to **NullableString** | The physical card serial number printed on the Kenya National ID card.              This is distinct from the ID Number and serves as a card issuance tracking identifier maintained by IPRS. This value changes each time a new physical card is issued (loss, damage, renewal). | [optional] 
**DateOfIssue** | Pointer to **NullableString** | The date the National ID was issued by the National Registration Bureau (NRB). | [optional] 
**PlaceOfBirth** | Pointer to [**NullableKenyaNidLookup2Address**](KenyaNidLookup2Address.md) | Place of birth as recorded in Kenya&#39;s civil registry (IPRS).              This is structured according to Kenya&#39;s pre-2010 administrative hierarchy (District &gt; Division &gt; Location). | [optional] 
**PlaceOfResidence** | Pointer to [**NullableKenyaNidLookup2Address**](KenyaNidLookup2Address.md) | Current residence address as registered in IPRS.              This represents the address on file at the time of ID registration or last update, structured according to Kenya&#39;s pre-2010 administrative hierarchy (District &gt; Division &gt; Location). | [optional] 

## Methods

### NewKenyaNidLookup2ProviderOutput

`func NewKenyaNidLookup2ProviderOutput() *KenyaNidLookup2ProviderOutput`

NewKenyaNidLookup2ProviderOutput instantiates a new KenyaNidLookup2ProviderOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKenyaNidLookup2ProviderOutputWithDefaults

`func NewKenyaNidLookup2ProviderOutputWithDefaults() *KenyaNidLookup2ProviderOutput`

NewKenyaNidLookup2ProviderOutputWithDefaults instantiates a new KenyaNidLookup2ProviderOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *KenyaNidLookup2ProviderOutput) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *KenyaNidLookup2ProviderOutput) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *KenyaNidLookup2ProviderOutput) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *KenyaNidLookup2ProviderOutput) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *KenyaNidLookup2ProviderOutput) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *KenyaNidLookup2ProviderOutput) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetSurname

`func (o *KenyaNidLookup2ProviderOutput) GetSurname() string`

GetSurname returns the Surname field if non-nil, zero value otherwise.

### GetSurnameOk

`func (o *KenyaNidLookup2ProviderOutput) GetSurnameOk() (*string, bool)`

GetSurnameOk returns a tuple with the Surname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurname

`func (o *KenyaNidLookup2ProviderOutput) SetSurname(v string)`

SetSurname sets Surname field to given value.

### HasSurname

`func (o *KenyaNidLookup2ProviderOutput) HasSurname() bool`

HasSurname returns a boolean if a field has been set.

### SetSurnameNil

`func (o *KenyaNidLookup2ProviderOutput) SetSurnameNil(b bool)`

 SetSurnameNil sets the value for Surname to be an explicit nil

### UnsetSurname
`func (o *KenyaNidLookup2ProviderOutput) UnsetSurname()`

UnsetSurname ensures that no value is present for Surname, not even an explicit nil
### GetOtherName

`func (o *KenyaNidLookup2ProviderOutput) GetOtherName() string`

GetOtherName returns the OtherName field if non-nil, zero value otherwise.

### GetOtherNameOk

`func (o *KenyaNidLookup2ProviderOutput) GetOtherNameOk() (*string, bool)`

GetOtherNameOk returns a tuple with the OtherName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherName

`func (o *KenyaNidLookup2ProviderOutput) SetOtherName(v string)`

SetOtherName sets OtherName field to given value.

### HasOtherName

`func (o *KenyaNidLookup2ProviderOutput) HasOtherName() bool`

HasOtherName returns a boolean if a field has been set.

### SetOtherNameNil

`func (o *KenyaNidLookup2ProviderOutput) SetOtherNameNil(b bool)`

 SetOtherNameNil sets the value for OtherName to be an explicit nil

### UnsetOtherName
`func (o *KenyaNidLookup2ProviderOutput) UnsetOtherName()`

UnsetOtherName ensures that no value is present for OtherName, not even an explicit nil
### GetSex

`func (o *KenyaNidLookup2ProviderOutput) GetSex() string`

GetSex returns the Sex field if non-nil, zero value otherwise.

### GetSexOk

`func (o *KenyaNidLookup2ProviderOutput) GetSexOk() (*string, bool)`

GetSexOk returns a tuple with the Sex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSex

`func (o *KenyaNidLookup2ProviderOutput) SetSex(v string)`

SetSex sets Sex field to given value.

### HasSex

`func (o *KenyaNidLookup2ProviderOutput) HasSex() bool`

HasSex returns a boolean if a field has been set.

### SetSexNil

`func (o *KenyaNidLookup2ProviderOutput) SetSexNil(b bool)`

 SetSexNil sets the value for Sex to be an explicit nil

### UnsetSex
`func (o *KenyaNidLookup2ProviderOutput) UnsetSex()`

UnsetSex ensures that no value is present for Sex, not even an explicit nil
### GetDateOfBirth

`func (o *KenyaNidLookup2ProviderOutput) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *KenyaNidLookup2ProviderOutput) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *KenyaNidLookup2ProviderOutput) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *KenyaNidLookup2ProviderOutput) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### SetDateOfBirthNil

`func (o *KenyaNidLookup2ProviderOutput) SetDateOfBirthNil(b bool)`

 SetDateOfBirthNil sets the value for DateOfBirth to be an explicit nil

### UnsetDateOfBirth
`func (o *KenyaNidLookup2ProviderOutput) UnsetDateOfBirth()`

UnsetDateOfBirth ensures that no value is present for DateOfBirth, not even an explicit nil
### GetCitizenship

`func (o *KenyaNidLookup2ProviderOutput) GetCitizenship() string`

GetCitizenship returns the Citizenship field if non-nil, zero value otherwise.

### GetCitizenshipOk

`func (o *KenyaNidLookup2ProviderOutput) GetCitizenshipOk() (*string, bool)`

GetCitizenshipOk returns a tuple with the Citizenship field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitizenship

`func (o *KenyaNidLookup2ProviderOutput) SetCitizenship(v string)`

SetCitizenship sets Citizenship field to given value.

### HasCitizenship

`func (o *KenyaNidLookup2ProviderOutput) HasCitizenship() bool`

HasCitizenship returns a boolean if a field has been set.

### SetCitizenshipNil

`func (o *KenyaNidLookup2ProviderOutput) SetCitizenshipNil(b bool)`

 SetCitizenshipNil sets the value for Citizenship to be an explicit nil

### UnsetCitizenship
`func (o *KenyaNidLookup2ProviderOutput) UnsetCitizenship()`

UnsetCitizenship ensures that no value is present for Citizenship, not even an explicit nil
### GetIdNumber

`func (o *KenyaNidLookup2ProviderOutput) GetIdNumber() string`

GetIdNumber returns the IdNumber field if non-nil, zero value otherwise.

### GetIdNumberOk

`func (o *KenyaNidLookup2ProviderOutput) GetIdNumberOk() (*string, bool)`

GetIdNumberOk returns a tuple with the IdNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdNumber

`func (o *KenyaNidLookup2ProviderOutput) SetIdNumber(v string)`

SetIdNumber sets IdNumber field to given value.

### HasIdNumber

`func (o *KenyaNidLookup2ProviderOutput) HasIdNumber() bool`

HasIdNumber returns a boolean if a field has been set.

### SetIdNumberNil

`func (o *KenyaNidLookup2ProviderOutput) SetIdNumberNil(b bool)`

 SetIdNumberNil sets the value for IdNumber to be an explicit nil

### UnsetIdNumber
`func (o *KenyaNidLookup2ProviderOutput) UnsetIdNumber()`

UnsetIdNumber ensures that no value is present for IdNumber, not even an explicit nil
### GetSerialNumber

`func (o *KenyaNidLookup2ProviderOutput) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *KenyaNidLookup2ProviderOutput) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *KenyaNidLookup2ProviderOutput) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *KenyaNidLookup2ProviderOutput) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### SetSerialNumberNil

`func (o *KenyaNidLookup2ProviderOutput) SetSerialNumberNil(b bool)`

 SetSerialNumberNil sets the value for SerialNumber to be an explicit nil

### UnsetSerialNumber
`func (o *KenyaNidLookup2ProviderOutput) UnsetSerialNumber()`

UnsetSerialNumber ensures that no value is present for SerialNumber, not even an explicit nil
### GetDateOfIssue

`func (o *KenyaNidLookup2ProviderOutput) GetDateOfIssue() string`

GetDateOfIssue returns the DateOfIssue field if non-nil, zero value otherwise.

### GetDateOfIssueOk

`func (o *KenyaNidLookup2ProviderOutput) GetDateOfIssueOk() (*string, bool)`

GetDateOfIssueOk returns a tuple with the DateOfIssue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfIssue

`func (o *KenyaNidLookup2ProviderOutput) SetDateOfIssue(v string)`

SetDateOfIssue sets DateOfIssue field to given value.

### HasDateOfIssue

`func (o *KenyaNidLookup2ProviderOutput) HasDateOfIssue() bool`

HasDateOfIssue returns a boolean if a field has been set.

### SetDateOfIssueNil

`func (o *KenyaNidLookup2ProviderOutput) SetDateOfIssueNil(b bool)`

 SetDateOfIssueNil sets the value for DateOfIssue to be an explicit nil

### UnsetDateOfIssue
`func (o *KenyaNidLookup2ProviderOutput) UnsetDateOfIssue()`

UnsetDateOfIssue ensures that no value is present for DateOfIssue, not even an explicit nil
### GetPlaceOfBirth

`func (o *KenyaNidLookup2ProviderOutput) GetPlaceOfBirth() KenyaNidLookup2Address`

GetPlaceOfBirth returns the PlaceOfBirth field if non-nil, zero value otherwise.

### GetPlaceOfBirthOk

`func (o *KenyaNidLookup2ProviderOutput) GetPlaceOfBirthOk() (*KenyaNidLookup2Address, bool)`

GetPlaceOfBirthOk returns a tuple with the PlaceOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceOfBirth

`func (o *KenyaNidLookup2ProviderOutput) SetPlaceOfBirth(v KenyaNidLookup2Address)`

SetPlaceOfBirth sets PlaceOfBirth field to given value.

### HasPlaceOfBirth

`func (o *KenyaNidLookup2ProviderOutput) HasPlaceOfBirth() bool`

HasPlaceOfBirth returns a boolean if a field has been set.

### SetPlaceOfBirthNil

`func (o *KenyaNidLookup2ProviderOutput) SetPlaceOfBirthNil(b bool)`

 SetPlaceOfBirthNil sets the value for PlaceOfBirth to be an explicit nil

### UnsetPlaceOfBirth
`func (o *KenyaNidLookup2ProviderOutput) UnsetPlaceOfBirth()`

UnsetPlaceOfBirth ensures that no value is present for PlaceOfBirth, not even an explicit nil
### GetPlaceOfResidence

`func (o *KenyaNidLookup2ProviderOutput) GetPlaceOfResidence() KenyaNidLookup2Address`

GetPlaceOfResidence returns the PlaceOfResidence field if non-nil, zero value otherwise.

### GetPlaceOfResidenceOk

`func (o *KenyaNidLookup2ProviderOutput) GetPlaceOfResidenceOk() (*KenyaNidLookup2Address, bool)`

GetPlaceOfResidenceOk returns a tuple with the PlaceOfResidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceOfResidence

`func (o *KenyaNidLookup2ProviderOutput) SetPlaceOfResidence(v KenyaNidLookup2Address)`

SetPlaceOfResidence sets PlaceOfResidence field to given value.

### HasPlaceOfResidence

`func (o *KenyaNidLookup2ProviderOutput) HasPlaceOfResidence() bool`

HasPlaceOfResidence returns a boolean if a field has been set.

### SetPlaceOfResidenceNil

`func (o *KenyaNidLookup2ProviderOutput) SetPlaceOfResidenceNil(b bool)`

 SetPlaceOfResidenceNil sets the value for PlaceOfResidence to be an explicit nil

### UnsetPlaceOfResidence
`func (o *KenyaNidLookup2ProviderOutput) UnsetPlaceOfResidence()`

UnsetPlaceOfResidence ensures that no value is present for PlaceOfResidence, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


