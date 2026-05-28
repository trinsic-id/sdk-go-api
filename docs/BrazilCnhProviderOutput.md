# BrazilCnhProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FullName** | Pointer to [**NullableBrazilCnhMatchFieldString**](BrazilCnhMatchFieldString.md) | Full name decoded from the CNH QR Code and whether it matched the government database. | [optional] 
**FullNameSimilarity** | Pointer to **NullableFloat64** | Similarity score for the decoded full name from the CNH QR Code.              Ranges from 0.0 to 1.0, where 1.0 is an exact match. | [optional] 
**OriginDocument** | Pointer to **NullableString** | Source identity document number/reference for the CNH record, decoded from the CNH QR Code. | [optional] 
**DateOfBirth** | Pointer to [**NullableBrazilCnhMatchFieldDateOnly**](BrazilCnhMatchFieldDateOnly.md) | Date of birth decoded from the CNH QR Code and whether it matched the government database. | [optional] 
**MotherFullName** | Pointer to **NullableString** | Mother&#39;s full name decoded from the CNH QR Code. | [optional] 
**FatherFullName** | Pointer to **NullableString** | Father&#39;s full name decoded from the CNH QR Code. | [optional] 
**Category** | Pointer to **NullableString** | Driver license category decoded from the CNH QR Code.              Known values: - A: Two- or three-wheeled motor vehicles, with or without a sidecar. - B: Motor vehicles outside category A up to 3,500 kg, with capacity for up to eight passengers besides the driver, and without articulation or trailer. - C: Category B vehicles and rigid cargo vehicles above 3,500 kg. - D: Category C vehicles and passenger vehicles with capacity above eight passengers besides the driver. - E: Category D vehicles, trailers, semi-trailers, and articulated buses. | [optional] 
**RegistrationNumber** | Pointer to [**NullableBrazilCnhMatchFieldString**](BrazilCnhMatchFieldString.md) | National Registry (Registro Nacional de Habilitacao) number decoded from the CNH QR Code and whether it matched the government database.              Is an 11-character ID with the format XXXXXXXXX@@: - XXXXXXXXX: Base characters. The value is unique to the driver for their lifetime. - @@: Security check digits | [optional] 
**RegistrationNumberCheckDigitsValid** | Pointer to **NullableBool** | Whether the National Registry number satisfies the CNH modulo-11 check digit routine. | [optional] 
**FirstIssueDate** | Pointer to **NullableString** | Date when the driver was first licensed, decoded from the CNH QR Code. | [optional] 
**Observations** | Pointer to **NullableString** | CNH observations decoded from the QR Code.              May include EAR for paid driving activity and one or more restriction codes: - A: Corrective lenses required. - B: Hearing prosthesis required. - C: Left-side accelerator required. - D: Automatic transmission vehicle required. - E: Steering-wheel grip, handle, or knob required. - F: Power steering vehicle required. - G: Manual clutch, automated clutch, or automatic transmission vehicle required. - H: Hand accelerator and brake required. - I: Dashboard control adaptations to the steering wheel required. - J: Dashboard control adaptations for lower limbs or other body parts required. - K: Gear-shift lever extension and/or fixed height/depth compensation cushions required. - L: Pedal extensions and floor elevation and/or fixed height/depth compensation cushions required. - M: Motorcycle with adapted gear-shift pedal required. - N: Motorcycle with adapted rear-brake pedal required. - O: Motorcycle with adapted front-brake handle required. - P: Motorcycle with adapted clutch handle required. - Q: Motorcycle with sidecar or tricycle required. - R: Scooter with sidecar or tricycle required. - S: Motorcycle with automated gear shifting required. - T: Driving on highways and high-speed roads prohibited. - U: Driving after sunset prohibited. - V: Helmet with visor and unrestricted field of vision required. - X: Other restrictions. | [optional] 
**IssueCity** | Pointer to **NullableString** | City where the CNH was issued, decoded from the QR Code. | [optional] 
**IssueState** | Pointer to **NullableString** | Brazilian state where the CNH was issued, decoded from the QR Code.              Format is a two-letter Brazilian state abbreviation when present. | [optional] 
**ValidationNumber** | Pointer to **NullableString** | Numeric security validation code decoded from the CNH QR Code.              This is an 11-digit code generated from Departamento Estadual de Trânsito (State Traffic Department or DETRAN) owned data for each CNH and used to validate the document. The algorithm is proprietary, so no further information is available on this number. | [optional] 
**RenachNumber** | Pointer to **NullableString** | Registro Nacional de Condutores Habilitados (RENACH) form number decoded from the CNH QR Code.              Is an 11-character ID with the format FFXXXXXXXX@: - FF: The issuing federation unit abbreviation. It identifies the state where the driver was licensed or last had record changes. - XXXXXXXXX: Serial number - @: Modulo-11 security check digit | [optional] 
**FacialBiometry** | Pointer to [**NullableBrazilCnhFacialBiometryOutput**](BrazilCnhFacialBiometryOutput.md) | Facial biometric comparison result if submitted and available. | [optional] 

## Methods

### NewBrazilCnhProviderOutput

`func NewBrazilCnhProviderOutput() *BrazilCnhProviderOutput`

NewBrazilCnhProviderOutput instantiates a new BrazilCnhProviderOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrazilCnhProviderOutputWithDefaults

`func NewBrazilCnhProviderOutputWithDefaults() *BrazilCnhProviderOutput`

NewBrazilCnhProviderOutputWithDefaults instantiates a new BrazilCnhProviderOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFullName

`func (o *BrazilCnhProviderOutput) GetFullName() BrazilCnhMatchFieldString`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *BrazilCnhProviderOutput) GetFullNameOk() (*BrazilCnhMatchFieldString, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *BrazilCnhProviderOutput) SetFullName(v BrazilCnhMatchFieldString)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *BrazilCnhProviderOutput) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### SetFullNameNil

`func (o *BrazilCnhProviderOutput) SetFullNameNil(b bool)`

 SetFullNameNil sets the value for FullName to be an explicit nil

### UnsetFullName
`func (o *BrazilCnhProviderOutput) UnsetFullName()`

UnsetFullName ensures that no value is present for FullName, not even an explicit nil
### GetFullNameSimilarity

`func (o *BrazilCnhProviderOutput) GetFullNameSimilarity() float64`

GetFullNameSimilarity returns the FullNameSimilarity field if non-nil, zero value otherwise.

### GetFullNameSimilarityOk

`func (o *BrazilCnhProviderOutput) GetFullNameSimilarityOk() (*float64, bool)`

GetFullNameSimilarityOk returns a tuple with the FullNameSimilarity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullNameSimilarity

`func (o *BrazilCnhProviderOutput) SetFullNameSimilarity(v float64)`

SetFullNameSimilarity sets FullNameSimilarity field to given value.

### HasFullNameSimilarity

`func (o *BrazilCnhProviderOutput) HasFullNameSimilarity() bool`

HasFullNameSimilarity returns a boolean if a field has been set.

### SetFullNameSimilarityNil

`func (o *BrazilCnhProviderOutput) SetFullNameSimilarityNil(b bool)`

 SetFullNameSimilarityNil sets the value for FullNameSimilarity to be an explicit nil

### UnsetFullNameSimilarity
`func (o *BrazilCnhProviderOutput) UnsetFullNameSimilarity()`

UnsetFullNameSimilarity ensures that no value is present for FullNameSimilarity, not even an explicit nil
### GetOriginDocument

`func (o *BrazilCnhProviderOutput) GetOriginDocument() string`

GetOriginDocument returns the OriginDocument field if non-nil, zero value otherwise.

### GetOriginDocumentOk

`func (o *BrazilCnhProviderOutput) GetOriginDocumentOk() (*string, bool)`

GetOriginDocumentOk returns a tuple with the OriginDocument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginDocument

`func (o *BrazilCnhProviderOutput) SetOriginDocument(v string)`

SetOriginDocument sets OriginDocument field to given value.

### HasOriginDocument

`func (o *BrazilCnhProviderOutput) HasOriginDocument() bool`

HasOriginDocument returns a boolean if a field has been set.

### SetOriginDocumentNil

`func (o *BrazilCnhProviderOutput) SetOriginDocumentNil(b bool)`

 SetOriginDocumentNil sets the value for OriginDocument to be an explicit nil

### UnsetOriginDocument
`func (o *BrazilCnhProviderOutput) UnsetOriginDocument()`

UnsetOriginDocument ensures that no value is present for OriginDocument, not even an explicit nil
### GetDateOfBirth

`func (o *BrazilCnhProviderOutput) GetDateOfBirth() BrazilCnhMatchFieldDateOnly`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *BrazilCnhProviderOutput) GetDateOfBirthOk() (*BrazilCnhMatchFieldDateOnly, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *BrazilCnhProviderOutput) SetDateOfBirth(v BrazilCnhMatchFieldDateOnly)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *BrazilCnhProviderOutput) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### SetDateOfBirthNil

`func (o *BrazilCnhProviderOutput) SetDateOfBirthNil(b bool)`

 SetDateOfBirthNil sets the value for DateOfBirth to be an explicit nil

### UnsetDateOfBirth
`func (o *BrazilCnhProviderOutput) UnsetDateOfBirth()`

UnsetDateOfBirth ensures that no value is present for DateOfBirth, not even an explicit nil
### GetMotherFullName

`func (o *BrazilCnhProviderOutput) GetMotherFullName() string`

GetMotherFullName returns the MotherFullName field if non-nil, zero value otherwise.

### GetMotherFullNameOk

`func (o *BrazilCnhProviderOutput) GetMotherFullNameOk() (*string, bool)`

GetMotherFullNameOk returns a tuple with the MotherFullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMotherFullName

`func (o *BrazilCnhProviderOutput) SetMotherFullName(v string)`

SetMotherFullName sets MotherFullName field to given value.

### HasMotherFullName

`func (o *BrazilCnhProviderOutput) HasMotherFullName() bool`

HasMotherFullName returns a boolean if a field has been set.

### SetMotherFullNameNil

`func (o *BrazilCnhProviderOutput) SetMotherFullNameNil(b bool)`

 SetMotherFullNameNil sets the value for MotherFullName to be an explicit nil

### UnsetMotherFullName
`func (o *BrazilCnhProviderOutput) UnsetMotherFullName()`

UnsetMotherFullName ensures that no value is present for MotherFullName, not even an explicit nil
### GetFatherFullName

`func (o *BrazilCnhProviderOutput) GetFatherFullName() string`

GetFatherFullName returns the FatherFullName field if non-nil, zero value otherwise.

### GetFatherFullNameOk

`func (o *BrazilCnhProviderOutput) GetFatherFullNameOk() (*string, bool)`

GetFatherFullNameOk returns a tuple with the FatherFullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFatherFullName

`func (o *BrazilCnhProviderOutput) SetFatherFullName(v string)`

SetFatherFullName sets FatherFullName field to given value.

### HasFatherFullName

`func (o *BrazilCnhProviderOutput) HasFatherFullName() bool`

HasFatherFullName returns a boolean if a field has been set.

### SetFatherFullNameNil

`func (o *BrazilCnhProviderOutput) SetFatherFullNameNil(b bool)`

 SetFatherFullNameNil sets the value for FatherFullName to be an explicit nil

### UnsetFatherFullName
`func (o *BrazilCnhProviderOutput) UnsetFatherFullName()`

UnsetFatherFullName ensures that no value is present for FatherFullName, not even an explicit nil
### GetCategory

`func (o *BrazilCnhProviderOutput) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *BrazilCnhProviderOutput) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *BrazilCnhProviderOutput) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *BrazilCnhProviderOutput) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *BrazilCnhProviderOutput) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *BrazilCnhProviderOutput) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetRegistrationNumber

`func (o *BrazilCnhProviderOutput) GetRegistrationNumber() BrazilCnhMatchFieldString`

GetRegistrationNumber returns the RegistrationNumber field if non-nil, zero value otherwise.

### GetRegistrationNumberOk

`func (o *BrazilCnhProviderOutput) GetRegistrationNumberOk() (*BrazilCnhMatchFieldString, bool)`

GetRegistrationNumberOk returns a tuple with the RegistrationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationNumber

`func (o *BrazilCnhProviderOutput) SetRegistrationNumber(v BrazilCnhMatchFieldString)`

SetRegistrationNumber sets RegistrationNumber field to given value.

### HasRegistrationNumber

`func (o *BrazilCnhProviderOutput) HasRegistrationNumber() bool`

HasRegistrationNumber returns a boolean if a field has been set.

### SetRegistrationNumberNil

`func (o *BrazilCnhProviderOutput) SetRegistrationNumberNil(b bool)`

 SetRegistrationNumberNil sets the value for RegistrationNumber to be an explicit nil

### UnsetRegistrationNumber
`func (o *BrazilCnhProviderOutput) UnsetRegistrationNumber()`

UnsetRegistrationNumber ensures that no value is present for RegistrationNumber, not even an explicit nil
### GetRegistrationNumberCheckDigitsValid

`func (o *BrazilCnhProviderOutput) GetRegistrationNumberCheckDigitsValid() bool`

GetRegistrationNumberCheckDigitsValid returns the RegistrationNumberCheckDigitsValid field if non-nil, zero value otherwise.

### GetRegistrationNumberCheckDigitsValidOk

`func (o *BrazilCnhProviderOutput) GetRegistrationNumberCheckDigitsValidOk() (*bool, bool)`

GetRegistrationNumberCheckDigitsValidOk returns a tuple with the RegistrationNumberCheckDigitsValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationNumberCheckDigitsValid

`func (o *BrazilCnhProviderOutput) SetRegistrationNumberCheckDigitsValid(v bool)`

SetRegistrationNumberCheckDigitsValid sets RegistrationNumberCheckDigitsValid field to given value.

### HasRegistrationNumberCheckDigitsValid

`func (o *BrazilCnhProviderOutput) HasRegistrationNumberCheckDigitsValid() bool`

HasRegistrationNumberCheckDigitsValid returns a boolean if a field has been set.

### SetRegistrationNumberCheckDigitsValidNil

`func (o *BrazilCnhProviderOutput) SetRegistrationNumberCheckDigitsValidNil(b bool)`

 SetRegistrationNumberCheckDigitsValidNil sets the value for RegistrationNumberCheckDigitsValid to be an explicit nil

### UnsetRegistrationNumberCheckDigitsValid
`func (o *BrazilCnhProviderOutput) UnsetRegistrationNumberCheckDigitsValid()`

UnsetRegistrationNumberCheckDigitsValid ensures that no value is present for RegistrationNumberCheckDigitsValid, not even an explicit nil
### GetFirstIssueDate

`func (o *BrazilCnhProviderOutput) GetFirstIssueDate() string`

GetFirstIssueDate returns the FirstIssueDate field if non-nil, zero value otherwise.

### GetFirstIssueDateOk

`func (o *BrazilCnhProviderOutput) GetFirstIssueDateOk() (*string, bool)`

GetFirstIssueDateOk returns a tuple with the FirstIssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstIssueDate

`func (o *BrazilCnhProviderOutput) SetFirstIssueDate(v string)`

SetFirstIssueDate sets FirstIssueDate field to given value.

### HasFirstIssueDate

`func (o *BrazilCnhProviderOutput) HasFirstIssueDate() bool`

HasFirstIssueDate returns a boolean if a field has been set.

### SetFirstIssueDateNil

`func (o *BrazilCnhProviderOutput) SetFirstIssueDateNil(b bool)`

 SetFirstIssueDateNil sets the value for FirstIssueDate to be an explicit nil

### UnsetFirstIssueDate
`func (o *BrazilCnhProviderOutput) UnsetFirstIssueDate()`

UnsetFirstIssueDate ensures that no value is present for FirstIssueDate, not even an explicit nil
### GetObservations

`func (o *BrazilCnhProviderOutput) GetObservations() string`

GetObservations returns the Observations field if non-nil, zero value otherwise.

### GetObservationsOk

`func (o *BrazilCnhProviderOutput) GetObservationsOk() (*string, bool)`

GetObservationsOk returns a tuple with the Observations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservations

`func (o *BrazilCnhProviderOutput) SetObservations(v string)`

SetObservations sets Observations field to given value.

### HasObservations

`func (o *BrazilCnhProviderOutput) HasObservations() bool`

HasObservations returns a boolean if a field has been set.

### SetObservationsNil

`func (o *BrazilCnhProviderOutput) SetObservationsNil(b bool)`

 SetObservationsNil sets the value for Observations to be an explicit nil

### UnsetObservations
`func (o *BrazilCnhProviderOutput) UnsetObservations()`

UnsetObservations ensures that no value is present for Observations, not even an explicit nil
### GetIssueCity

`func (o *BrazilCnhProviderOutput) GetIssueCity() string`

GetIssueCity returns the IssueCity field if non-nil, zero value otherwise.

### GetIssueCityOk

`func (o *BrazilCnhProviderOutput) GetIssueCityOk() (*string, bool)`

GetIssueCityOk returns a tuple with the IssueCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueCity

`func (o *BrazilCnhProviderOutput) SetIssueCity(v string)`

SetIssueCity sets IssueCity field to given value.

### HasIssueCity

`func (o *BrazilCnhProviderOutput) HasIssueCity() bool`

HasIssueCity returns a boolean if a field has been set.

### SetIssueCityNil

`func (o *BrazilCnhProviderOutput) SetIssueCityNil(b bool)`

 SetIssueCityNil sets the value for IssueCity to be an explicit nil

### UnsetIssueCity
`func (o *BrazilCnhProviderOutput) UnsetIssueCity()`

UnsetIssueCity ensures that no value is present for IssueCity, not even an explicit nil
### GetIssueState

`func (o *BrazilCnhProviderOutput) GetIssueState() string`

GetIssueState returns the IssueState field if non-nil, zero value otherwise.

### GetIssueStateOk

`func (o *BrazilCnhProviderOutput) GetIssueStateOk() (*string, bool)`

GetIssueStateOk returns a tuple with the IssueState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueState

`func (o *BrazilCnhProviderOutput) SetIssueState(v string)`

SetIssueState sets IssueState field to given value.

### HasIssueState

`func (o *BrazilCnhProviderOutput) HasIssueState() bool`

HasIssueState returns a boolean if a field has been set.

### SetIssueStateNil

`func (o *BrazilCnhProviderOutput) SetIssueStateNil(b bool)`

 SetIssueStateNil sets the value for IssueState to be an explicit nil

### UnsetIssueState
`func (o *BrazilCnhProviderOutput) UnsetIssueState()`

UnsetIssueState ensures that no value is present for IssueState, not even an explicit nil
### GetValidationNumber

`func (o *BrazilCnhProviderOutput) GetValidationNumber() string`

GetValidationNumber returns the ValidationNumber field if non-nil, zero value otherwise.

### GetValidationNumberOk

`func (o *BrazilCnhProviderOutput) GetValidationNumberOk() (*string, bool)`

GetValidationNumberOk returns a tuple with the ValidationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationNumber

`func (o *BrazilCnhProviderOutput) SetValidationNumber(v string)`

SetValidationNumber sets ValidationNumber field to given value.

### HasValidationNumber

`func (o *BrazilCnhProviderOutput) HasValidationNumber() bool`

HasValidationNumber returns a boolean if a field has been set.

### SetValidationNumberNil

`func (o *BrazilCnhProviderOutput) SetValidationNumberNil(b bool)`

 SetValidationNumberNil sets the value for ValidationNumber to be an explicit nil

### UnsetValidationNumber
`func (o *BrazilCnhProviderOutput) UnsetValidationNumber()`

UnsetValidationNumber ensures that no value is present for ValidationNumber, not even an explicit nil
### GetRenachNumber

`func (o *BrazilCnhProviderOutput) GetRenachNumber() string`

GetRenachNumber returns the RenachNumber field if non-nil, zero value otherwise.

### GetRenachNumberOk

`func (o *BrazilCnhProviderOutput) GetRenachNumberOk() (*string, bool)`

GetRenachNumberOk returns a tuple with the RenachNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenachNumber

`func (o *BrazilCnhProviderOutput) SetRenachNumber(v string)`

SetRenachNumber sets RenachNumber field to given value.

### HasRenachNumber

`func (o *BrazilCnhProviderOutput) HasRenachNumber() bool`

HasRenachNumber returns a boolean if a field has been set.

### SetRenachNumberNil

`func (o *BrazilCnhProviderOutput) SetRenachNumberNil(b bool)`

 SetRenachNumberNil sets the value for RenachNumber to be an explicit nil

### UnsetRenachNumber
`func (o *BrazilCnhProviderOutput) UnsetRenachNumber()`

UnsetRenachNumber ensures that no value is present for RenachNumber, not even an explicit nil
### GetFacialBiometry

`func (o *BrazilCnhProviderOutput) GetFacialBiometry() BrazilCnhFacialBiometryOutput`

GetFacialBiometry returns the FacialBiometry field if non-nil, zero value otherwise.

### GetFacialBiometryOk

`func (o *BrazilCnhProviderOutput) GetFacialBiometryOk() (*BrazilCnhFacialBiometryOutput, bool)`

GetFacialBiometryOk returns a tuple with the FacialBiometry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacialBiometry

`func (o *BrazilCnhProviderOutput) SetFacialBiometry(v BrazilCnhFacialBiometryOutput)`

SetFacialBiometry sets FacialBiometry field to given value.

### HasFacialBiometry

`func (o *BrazilCnhProviderOutput) HasFacialBiometry() bool`

HasFacialBiometry returns a boolean if a field has been set.

### SetFacialBiometryNil

`func (o *BrazilCnhProviderOutput) SetFacialBiometryNil(b bool)`

 SetFacialBiometryNil sets the value for FacialBiometry to be an explicit nil

### UnsetFacialBiometry
`func (o *BrazilCnhProviderOutput) UnsetFacialBiometry()`

UnsetFacialBiometry ensures that no value is present for FacialBiometry, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


