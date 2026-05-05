# CoteDIvoireNidLookup2ProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FullName** | **string** | Full name as returned by ONECI (National Civil Registry and Identification Office). | 
**GivenName** | **string** | Given name of the ID holder as returned by ONECI (National Civil Registry and Identification Office). | 
**FamilyName** | **string** | Family name of the ID holder as returned by ONECI (National Civil Registry and Identification Office). | 
**DateOfBirth** | **string** | Date of birth as returned by ONECI (National Civil Registry and Identification Office). | 
**Sex** | **string** | Sex of the ID holder as returned by ONECI (National Civil Registry and Identification Office). Possible values: - Male - Female | 
**Nationality** | **string** | Nationality as ISO 3166-1 alpha-2 country code (e.g. \&quot;CI\&quot; for Côte d&#39;Ivoire). | 
**Address** | Pointer to **NullableString** | Address as returned from ONECI (National Civil Registry and Identification Office). Format is LOCALITY,COMMUNE: the locality (village, neighborhood, or sous-quartier) followed by the commune. Not a full street address. | [optional] 
**DocumentNumber** | Pointer to **NullableString** | The document&#39;s own identifier (printed on the card; often labeled \&quot;Immatriculation\&quot; on older cards or \&quot;Numéro CNI\&quot; in post-2020 cards). On older cards this is usually one letter followed by 10 digits. On new cards it appears as 1 or 2 leading letters plus 9 digits. | [optional] 
**NationalIdNumber** | Pointer to **NullableString** | The NNI (Numéro National d&#39;Identification): the person&#39;s 11-digit national ID, printed on the back of the new national id card and assigned by ONECI. It is always exactly 11 digits with no letters, and is semi-random, non-repetitive, and does not encode any extra data, such as date of birth, gender, or other readable attributes. | [optional] 

## Methods

### NewCoteDIvoireNidLookup2ProviderOutput

`func NewCoteDIvoireNidLookup2ProviderOutput(fullName string, givenName string, familyName string, dateOfBirth string, sex string, nationality string, ) *CoteDIvoireNidLookup2ProviderOutput`

NewCoteDIvoireNidLookup2ProviderOutput instantiates a new CoteDIvoireNidLookup2ProviderOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoteDIvoireNidLookup2ProviderOutputWithDefaults

`func NewCoteDIvoireNidLookup2ProviderOutputWithDefaults() *CoteDIvoireNidLookup2ProviderOutput`

NewCoteDIvoireNidLookup2ProviderOutputWithDefaults instantiates a new CoteDIvoireNidLookup2ProviderOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFullName

`func (o *CoteDIvoireNidLookup2ProviderOutput) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *CoteDIvoireNidLookup2ProviderOutput) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *CoteDIvoireNidLookup2ProviderOutput) SetFullName(v string)`

SetFullName sets FullName field to given value.


### GetGivenName

`func (o *CoteDIvoireNidLookup2ProviderOutput) GetGivenName() string`

GetGivenName returns the GivenName field if non-nil, zero value otherwise.

### GetGivenNameOk

`func (o *CoteDIvoireNidLookup2ProviderOutput) GetGivenNameOk() (*string, bool)`

GetGivenNameOk returns a tuple with the GivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenName

`func (o *CoteDIvoireNidLookup2ProviderOutput) SetGivenName(v string)`

SetGivenName sets GivenName field to given value.


### GetFamilyName

`func (o *CoteDIvoireNidLookup2ProviderOutput) GetFamilyName() string`

GetFamilyName returns the FamilyName field if non-nil, zero value otherwise.

### GetFamilyNameOk

`func (o *CoteDIvoireNidLookup2ProviderOutput) GetFamilyNameOk() (*string, bool)`

GetFamilyNameOk returns a tuple with the FamilyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyName

`func (o *CoteDIvoireNidLookup2ProviderOutput) SetFamilyName(v string)`

SetFamilyName sets FamilyName field to given value.


### GetDateOfBirth

`func (o *CoteDIvoireNidLookup2ProviderOutput) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *CoteDIvoireNidLookup2ProviderOutput) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *CoteDIvoireNidLookup2ProviderOutput) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.


### GetSex

`func (o *CoteDIvoireNidLookup2ProviderOutput) GetSex() string`

GetSex returns the Sex field if non-nil, zero value otherwise.

### GetSexOk

`func (o *CoteDIvoireNidLookup2ProviderOutput) GetSexOk() (*string, bool)`

GetSexOk returns a tuple with the Sex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSex

`func (o *CoteDIvoireNidLookup2ProviderOutput) SetSex(v string)`

SetSex sets Sex field to given value.


### GetNationality

`func (o *CoteDIvoireNidLookup2ProviderOutput) GetNationality() string`

GetNationality returns the Nationality field if non-nil, zero value otherwise.

### GetNationalityOk

`func (o *CoteDIvoireNidLookup2ProviderOutput) GetNationalityOk() (*string, bool)`

GetNationalityOk returns a tuple with the Nationality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationality

`func (o *CoteDIvoireNidLookup2ProviderOutput) SetNationality(v string)`

SetNationality sets Nationality field to given value.


### GetAddress

`func (o *CoteDIvoireNidLookup2ProviderOutput) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *CoteDIvoireNidLookup2ProviderOutput) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *CoteDIvoireNidLookup2ProviderOutput) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *CoteDIvoireNidLookup2ProviderOutput) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *CoteDIvoireNidLookup2ProviderOutput) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *CoteDIvoireNidLookup2ProviderOutput) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetDocumentNumber

`func (o *CoteDIvoireNidLookup2ProviderOutput) GetDocumentNumber() string`

GetDocumentNumber returns the DocumentNumber field if non-nil, zero value otherwise.

### GetDocumentNumberOk

`func (o *CoteDIvoireNidLookup2ProviderOutput) GetDocumentNumberOk() (*string, bool)`

GetDocumentNumberOk returns a tuple with the DocumentNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentNumber

`func (o *CoteDIvoireNidLookup2ProviderOutput) SetDocumentNumber(v string)`

SetDocumentNumber sets DocumentNumber field to given value.

### HasDocumentNumber

`func (o *CoteDIvoireNidLookup2ProviderOutput) HasDocumentNumber() bool`

HasDocumentNumber returns a boolean if a field has been set.

### SetDocumentNumberNil

`func (o *CoteDIvoireNidLookup2ProviderOutput) SetDocumentNumberNil(b bool)`

 SetDocumentNumberNil sets the value for DocumentNumber to be an explicit nil

### UnsetDocumentNumber
`func (o *CoteDIvoireNidLookup2ProviderOutput) UnsetDocumentNumber()`

UnsetDocumentNumber ensures that no value is present for DocumentNumber, not even an explicit nil
### GetNationalIdNumber

`func (o *CoteDIvoireNidLookup2ProviderOutput) GetNationalIdNumber() string`

GetNationalIdNumber returns the NationalIdNumber field if non-nil, zero value otherwise.

### GetNationalIdNumberOk

`func (o *CoteDIvoireNidLookup2ProviderOutput) GetNationalIdNumberOk() (*string, bool)`

GetNationalIdNumberOk returns a tuple with the NationalIdNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationalIdNumber

`func (o *CoteDIvoireNidLookup2ProviderOutput) SetNationalIdNumber(v string)`

SetNationalIdNumber sets NationalIdNumber field to given value.

### HasNationalIdNumber

`func (o *CoteDIvoireNidLookup2ProviderOutput) HasNationalIdNumber() bool`

HasNationalIdNumber returns a boolean if a field has been set.

### SetNationalIdNumberNil

`func (o *CoteDIvoireNidLookup2ProviderOutput) SetNationalIdNumberNil(b bool)`

 SetNationalIdNumberNil sets the value for NationalIdNumber to be an explicit nil

### UnsetNationalIdNumber
`func (o *CoteDIvoireNidLookup2ProviderOutput) UnsetNationalIdNumber()`

UnsetNationalIdNumber ensures that no value is present for NationalIdNumber, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


