# PeruDniProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to [**NullablePeruDniAddress**](PeruDniAddress.md) | Information about the general area in which the DNI holder resides.              This field is not always available. | [optional] 
**ArrayName** | **[]string** | All names that appear on DNI, as an array of strings.              Format: - All uppercase - Ordered by paternal family name, then maternal family name, then given names. | 
**CivilStatus** | Pointer to **NullableString** | Marital status as it appears on the DNI.              This field is not always available.              Valid values: - \&quot;Single\&quot; - \&quot;Married\&quot; - \&quot;Divorced\&quot; - \&quot;Widowed\&quot; | [optional] 
**DateOfBirth** | Pointer to **NullableString** | Date of birth as it appears on the DNI.              This field is not always available. | [optional] 
**DocumentNumber** | **string** | The user&#39;s National Identity Document number (Documento Nacional de Identidad or DNI).              Format: - 8 digits - Does NOT include verification digit. On the DNI card, a ninth digit appears next to the first eight. It is   0-9 or A-K. This is NOT included in the DNI number when returned from Peru&#39;s database. - Does NOT include dots or hyphens | 
**DocumentType** | **string** | Should always be \&quot;DNI\&quot;. | 
**ExpeditionDate** | Pointer to **NullableString** | Date that the DNI was issued.              This field is not always available.              Format: - yyyy-MM-dd | [optional] 
**ExpirationDate** | Pointer to **NullableString** | Date that the DNI will expire. DNI expires every 8 years, unless the citizen is &gt;&#x3D; 60 years old, in which case it never expires.              This field is not always available.              Format: - yyyy-MM-dd | [optional] 
**FirstName** | **string** | Given names as they appear on DNI.              Format: - All uppercase - Space-separated - Will include all given names | 
**FullName** | **string** | All names as they appear on DNI.              Format: - All uppercase - Space-separated - Will include all names, given and family - Ordered by given names first, then paternal family name, then maternal family name | 
**LastName** | **string** | Family names as they appear on DNI.              Format: - All uppercase - Will include all family names - Ordered by paternal family name first, then maternal family name | 
**MaternalLastName** | **string** | Maternal last name as it appears on DNI.              Format: - All uppercase | 
**PaternalLastName** | **string** | Paternal last name as it appears on DNI.              Format: - All uppercase | 
**Sex** | Pointer to **NullableString** | Sex as it appears on DNI.              This field is not always available.              Values: - \&quot;Male\&quot; - \&quot;Female\&quot; | [optional] 
**UbigeoReniec** | Pointer to **NullableString** | In Peru, geographical locations have an official geographical code called UBIGEO, from the spanish \&quot;UBIcación GEOgráfica\&quot; (Geographic Location). This is an administrative geocode, is different from a postal code (which Peru also has) and is used to specifically delineate the administrative region, province and district hierarchy.              There are two coding systems for UBIGEO: one from INEI (National Institute of Statistics and Informatics) and another from RENIEC (National Registry of Identification and Civil Status). The two coding systems are similar but are not 100% the same (some numbers will map to different geographic locations). This field follows the coding system from RENIEC.              This field is not always available.              Format: - Always 6 digits - First two digits represent region - Middle two are province - Last two are district              Given the example 081304, that would correspond to: - 08 - Cusco Region - 0813 - Urubamba Province - 081304 - Machupicchu District | [optional] 
**VerificationDigit** | Pointer to **NullableString** | The final (ninth) digit of the DNI, which serves as a checksum over the first eight digits.              This field is not always available.              Format: - Single character - Either 0-9 or A-K              Read more here:   https://elcomercio.pe/mag/respuestas/cual-es-el-digito-verificador-de-mi-dni-documento-nacional-de-identidad-reniec-peru-nnda-nnlt-noticia/ | [optional] 

## Methods

### NewPeruDniProviderOutput

`func NewPeruDniProviderOutput(arrayName []string, documentNumber string, documentType string, firstName string, fullName string, lastName string, maternalLastName string, paternalLastName string, ) *PeruDniProviderOutput`

NewPeruDniProviderOutput instantiates a new PeruDniProviderOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPeruDniProviderOutputWithDefaults

`func NewPeruDniProviderOutputWithDefaults() *PeruDniProviderOutput`

NewPeruDniProviderOutputWithDefaults instantiates a new PeruDniProviderOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *PeruDniProviderOutput) GetAddress() PeruDniAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *PeruDniProviderOutput) GetAddressOk() (*PeruDniAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *PeruDniProviderOutput) SetAddress(v PeruDniAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *PeruDniProviderOutput) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *PeruDniProviderOutput) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *PeruDniProviderOutput) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetArrayName

`func (o *PeruDniProviderOutput) GetArrayName() []string`

GetArrayName returns the ArrayName field if non-nil, zero value otherwise.

### GetArrayNameOk

`func (o *PeruDniProviderOutput) GetArrayNameOk() (*[]string, bool)`

GetArrayNameOk returns a tuple with the ArrayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrayName

`func (o *PeruDniProviderOutput) SetArrayName(v []string)`

SetArrayName sets ArrayName field to given value.


### GetCivilStatus

`func (o *PeruDniProviderOutput) GetCivilStatus() string`

GetCivilStatus returns the CivilStatus field if non-nil, zero value otherwise.

### GetCivilStatusOk

`func (o *PeruDniProviderOutput) GetCivilStatusOk() (*string, bool)`

GetCivilStatusOk returns a tuple with the CivilStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCivilStatus

`func (o *PeruDniProviderOutput) SetCivilStatus(v string)`

SetCivilStatus sets CivilStatus field to given value.

### HasCivilStatus

`func (o *PeruDniProviderOutput) HasCivilStatus() bool`

HasCivilStatus returns a boolean if a field has been set.

### SetCivilStatusNil

`func (o *PeruDniProviderOutput) SetCivilStatusNil(b bool)`

 SetCivilStatusNil sets the value for CivilStatus to be an explicit nil

### UnsetCivilStatus
`func (o *PeruDniProviderOutput) UnsetCivilStatus()`

UnsetCivilStatus ensures that no value is present for CivilStatus, not even an explicit nil
### GetDateOfBirth

`func (o *PeruDniProviderOutput) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *PeruDniProviderOutput) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *PeruDniProviderOutput) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *PeruDniProviderOutput) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### SetDateOfBirthNil

`func (o *PeruDniProviderOutput) SetDateOfBirthNil(b bool)`

 SetDateOfBirthNil sets the value for DateOfBirth to be an explicit nil

### UnsetDateOfBirth
`func (o *PeruDniProviderOutput) UnsetDateOfBirth()`

UnsetDateOfBirth ensures that no value is present for DateOfBirth, not even an explicit nil
### GetDocumentNumber

`func (o *PeruDniProviderOutput) GetDocumentNumber() string`

GetDocumentNumber returns the DocumentNumber field if non-nil, zero value otherwise.

### GetDocumentNumberOk

`func (o *PeruDniProviderOutput) GetDocumentNumberOk() (*string, bool)`

GetDocumentNumberOk returns a tuple with the DocumentNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentNumber

`func (o *PeruDniProviderOutput) SetDocumentNumber(v string)`

SetDocumentNumber sets DocumentNumber field to given value.


### GetDocumentType

`func (o *PeruDniProviderOutput) GetDocumentType() string`

GetDocumentType returns the DocumentType field if non-nil, zero value otherwise.

### GetDocumentTypeOk

`func (o *PeruDniProviderOutput) GetDocumentTypeOk() (*string, bool)`

GetDocumentTypeOk returns a tuple with the DocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentType

`func (o *PeruDniProviderOutput) SetDocumentType(v string)`

SetDocumentType sets DocumentType field to given value.


### GetExpeditionDate

`func (o *PeruDniProviderOutput) GetExpeditionDate() string`

GetExpeditionDate returns the ExpeditionDate field if non-nil, zero value otherwise.

### GetExpeditionDateOk

`func (o *PeruDniProviderOutput) GetExpeditionDateOk() (*string, bool)`

GetExpeditionDateOk returns a tuple with the ExpeditionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpeditionDate

`func (o *PeruDniProviderOutput) SetExpeditionDate(v string)`

SetExpeditionDate sets ExpeditionDate field to given value.

### HasExpeditionDate

`func (o *PeruDniProviderOutput) HasExpeditionDate() bool`

HasExpeditionDate returns a boolean if a field has been set.

### SetExpeditionDateNil

`func (o *PeruDniProviderOutput) SetExpeditionDateNil(b bool)`

 SetExpeditionDateNil sets the value for ExpeditionDate to be an explicit nil

### UnsetExpeditionDate
`func (o *PeruDniProviderOutput) UnsetExpeditionDate()`

UnsetExpeditionDate ensures that no value is present for ExpeditionDate, not even an explicit nil
### GetExpirationDate

`func (o *PeruDniProviderOutput) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *PeruDniProviderOutput) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *PeruDniProviderOutput) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *PeruDniProviderOutput) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### SetExpirationDateNil

`func (o *PeruDniProviderOutput) SetExpirationDateNil(b bool)`

 SetExpirationDateNil sets the value for ExpirationDate to be an explicit nil

### UnsetExpirationDate
`func (o *PeruDniProviderOutput) UnsetExpirationDate()`

UnsetExpirationDate ensures that no value is present for ExpirationDate, not even an explicit nil
### GetFirstName

`func (o *PeruDniProviderOutput) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *PeruDniProviderOutput) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *PeruDniProviderOutput) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetFullName

`func (o *PeruDniProviderOutput) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *PeruDniProviderOutput) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *PeruDniProviderOutput) SetFullName(v string)`

SetFullName sets FullName field to given value.


### GetLastName

`func (o *PeruDniProviderOutput) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *PeruDniProviderOutput) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *PeruDniProviderOutput) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetMaternalLastName

`func (o *PeruDniProviderOutput) GetMaternalLastName() string`

GetMaternalLastName returns the MaternalLastName field if non-nil, zero value otherwise.

### GetMaternalLastNameOk

`func (o *PeruDniProviderOutput) GetMaternalLastNameOk() (*string, bool)`

GetMaternalLastNameOk returns a tuple with the MaternalLastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaternalLastName

`func (o *PeruDniProviderOutput) SetMaternalLastName(v string)`

SetMaternalLastName sets MaternalLastName field to given value.


### GetPaternalLastName

`func (o *PeruDniProviderOutput) GetPaternalLastName() string`

GetPaternalLastName returns the PaternalLastName field if non-nil, zero value otherwise.

### GetPaternalLastNameOk

`func (o *PeruDniProviderOutput) GetPaternalLastNameOk() (*string, bool)`

GetPaternalLastNameOk returns a tuple with the PaternalLastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaternalLastName

`func (o *PeruDniProviderOutput) SetPaternalLastName(v string)`

SetPaternalLastName sets PaternalLastName field to given value.


### GetSex

`func (o *PeruDniProviderOutput) GetSex() string`

GetSex returns the Sex field if non-nil, zero value otherwise.

### GetSexOk

`func (o *PeruDniProviderOutput) GetSexOk() (*string, bool)`

GetSexOk returns a tuple with the Sex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSex

`func (o *PeruDniProviderOutput) SetSex(v string)`

SetSex sets Sex field to given value.

### HasSex

`func (o *PeruDniProviderOutput) HasSex() bool`

HasSex returns a boolean if a field has been set.

### SetSexNil

`func (o *PeruDniProviderOutput) SetSexNil(b bool)`

 SetSexNil sets the value for Sex to be an explicit nil

### UnsetSex
`func (o *PeruDniProviderOutput) UnsetSex()`

UnsetSex ensures that no value is present for Sex, not even an explicit nil
### GetUbigeoReniec

`func (o *PeruDniProviderOutput) GetUbigeoReniec() string`

GetUbigeoReniec returns the UbigeoReniec field if non-nil, zero value otherwise.

### GetUbigeoReniecOk

`func (o *PeruDniProviderOutput) GetUbigeoReniecOk() (*string, bool)`

GetUbigeoReniecOk returns a tuple with the UbigeoReniec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUbigeoReniec

`func (o *PeruDniProviderOutput) SetUbigeoReniec(v string)`

SetUbigeoReniec sets UbigeoReniec field to given value.

### HasUbigeoReniec

`func (o *PeruDniProviderOutput) HasUbigeoReniec() bool`

HasUbigeoReniec returns a boolean if a field has been set.

### SetUbigeoReniecNil

`func (o *PeruDniProviderOutput) SetUbigeoReniecNil(b bool)`

 SetUbigeoReniecNil sets the value for UbigeoReniec to be an explicit nil

### UnsetUbigeoReniec
`func (o *PeruDniProviderOutput) UnsetUbigeoReniec()`

UnsetUbigeoReniec ensures that no value is present for UbigeoReniec, not even an explicit nil
### GetVerificationDigit

`func (o *PeruDniProviderOutput) GetVerificationDigit() string`

GetVerificationDigit returns the VerificationDigit field if non-nil, zero value otherwise.

### GetVerificationDigitOk

`func (o *PeruDniProviderOutput) GetVerificationDigitOk() (*string, bool)`

GetVerificationDigitOk returns a tuple with the VerificationDigit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationDigit

`func (o *PeruDniProviderOutput) SetVerificationDigit(v string)`

SetVerificationDigit sets VerificationDigit field to given value.

### HasVerificationDigit

`func (o *PeruDniProviderOutput) HasVerificationDigit() bool`

HasVerificationDigit returns a boolean if a field has been set.

### SetVerificationDigitNil

`func (o *PeruDniProviderOutput) SetVerificationDigitNil(b bool)`

 SetVerificationDigitNil sets the value for VerificationDigit to be an explicit nil

### UnsetVerificationDigit
`func (o *PeruDniProviderOutput) UnsetVerificationDigit()`

UnsetVerificationDigit ensures that no value is present for VerificationDigit, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


