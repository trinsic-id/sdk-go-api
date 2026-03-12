# OutputFrejaDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **NullableString** | Type of document. Possible values are: * \&quot;PASS\&quot;: Passport * \&quot;DRILIC\&quot;: Driver&#39;s License * \&quot;NATID\&quot;: National ID * \&quot;IDSIS\&quot;: SiS certified ID document * \&quot;TAXID\&quot;: Tax Agency ID card * \&quot;OTHERID\&quot;: Other IDs | [optional] 
**SerialNumber** | Pointer to **NullableString** | The document serial number. The structure of this number depends on the type and nationality of the document. | [optional] 
**ExpirationDate** | Pointer to **NullableString** | Expiration date of the document. Formatted as an ISO 8601 Date. | [optional] 
**Country** | Pointer to **NullableString** | The ISO 3166-1 alpha-2 country code associated with the document. | [optional] 

## Methods

### NewOutputFrejaDocument

`func NewOutputFrejaDocument() *OutputFrejaDocument`

NewOutputFrejaDocument instantiates a new OutputFrejaDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputFrejaDocumentWithDefaults

`func NewOutputFrejaDocumentWithDefaults() *OutputFrejaDocument`

NewOutputFrejaDocumentWithDefaults instantiates a new OutputFrejaDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *OutputFrejaDocument) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OutputFrejaDocument) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OutputFrejaDocument) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OutputFrejaDocument) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *OutputFrejaDocument) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *OutputFrejaDocument) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetSerialNumber

`func (o *OutputFrejaDocument) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *OutputFrejaDocument) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *OutputFrejaDocument) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *OutputFrejaDocument) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### SetSerialNumberNil

`func (o *OutputFrejaDocument) SetSerialNumberNil(b bool)`

 SetSerialNumberNil sets the value for SerialNumber to be an explicit nil

### UnsetSerialNumber
`func (o *OutputFrejaDocument) UnsetSerialNumber()`

UnsetSerialNumber ensures that no value is present for SerialNumber, not even an explicit nil
### GetExpirationDate

`func (o *OutputFrejaDocument) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *OutputFrejaDocument) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *OutputFrejaDocument) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *OutputFrejaDocument) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### SetExpirationDateNil

`func (o *OutputFrejaDocument) SetExpirationDateNil(b bool)`

 SetExpirationDateNil sets the value for ExpirationDate to be an explicit nil

### UnsetExpirationDate
`func (o *OutputFrejaDocument) UnsetExpirationDate()`

UnsetExpirationDate ensures that no value is present for ExpirationDate, not even an explicit nil
### GetCountry

`func (o *OutputFrejaDocument) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *OutputFrejaDocument) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *OutputFrejaDocument) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *OutputFrejaDocument) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *OutputFrejaDocument) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *OutputFrejaDocument) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


