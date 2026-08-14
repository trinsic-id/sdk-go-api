# MoldovaVehicleRegistrationCertificateCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlateNumber** | Pointer to **NullableString** | The registered vehicle&#39;s plate number. | [optional] 
**Idnv** | Pointer to **NullableString** | The registered vehicle&#39;s IDNV (Numărul de Identificare a Vehiculului), the identifier assigned to the vehicle in the Moldovan State Register of Transport.              This is a Moldovan registry identifier and is distinct from the manufacturer&#39;s VIN. | [optional] 
**Vin** | Pointer to **NullableString** | The vehicle&#39;s manufacturer-assigned Vehicle Identification Number (VIN), per ISO 3779. | [optional] 
**Make** | Pointer to **NullableString** | The make (manufacturer brand) of the vehicle. | [optional] 
**Model** | Pointer to **NullableString** | The commercial model name of the vehicle. | [optional] 
**Color** | Pointer to **NullableString** | The color of the vehicle, as recorded by the issuer.              TODO: Possible values? | [optional] 
**Category** | Pointer to **NullableString** | The vehicle&#39;s EU vehicle category code. | [optional] 
**Year** | Pointer to **NullableInt32** | The vehicle&#39;s year of manufacture. | [optional] 
**BodyNumber** | Pointer to **NullableString** | The body number of the vehicle, if recorded separately from the VIN. | [optional] 
**BodyType** | Pointer to **NullableString** | The body type of the vehicle, as recorded by the issuer.              This is a free text field with no guaranteed format. | [optional] 
**ChassisNumber** | Pointer to **NullableString** | The chassis number of the vehicle, if recorded separately from the VIN. | [optional] 
**EngineVolume** | Pointer to **NullableString** | The engine displacement volume of the vehicle, as recorded by the issuer (typically in cubic centimeters).              This is a free text field with no guaranteed format. | [optional] 
**EngineType** | Pointer to **NullableString** | The engine/fuel type of the vehicle, as recorded by the issuer.              This is a free text field with no guaranteed format. | [optional] 
**EngineNumber** | Pointer to **NullableString** | The engine number of the vehicle. | [optional] 
**AuthorizedWeight** | Pointer to **NullableInt32** | The maximum technically permissible laden mass of the vehicle, in kilograms. | [optional] 
**Weight** | Pointer to **NullableInt32** | The mass of the vehicle in service (unladen / kerb weight), in kilograms. | [optional] 
**Places** | Pointer to **NullableInt32** | The number of seating positions in the vehicle, including the driver. | [optional] 
**Idnp** | Pointer to **NullableString** | The IDNP (Numărul de Identificare Personal) of the vehicle&#39;s registered holder.              This is a 13-digit number which uniquely identifies all natural persons in the Republic of Moldova. It has the format &#x60;2YYYOOOSSSSSX&#x60;, where: - &#x60;2&#x60; is the literal number &#x60;2&#x60; to indicate that the identifier belongs to a natural person - &#x60;YYY&#x60; is the last 3 digits of the year in which the identifier was issued - &#x60;OOO&#x60; is the code of the registrar office which issued the identifier - &#x60;SSSSS&#x60; is the sequential birth number for that year - &#x60;X&#x60; is a check digit              Note that the year encoded in the identifier is not necessarily the same as the year of birth of the individual. | [optional] 
**FamilyName** | Pointer to **NullableString** | The family name (surname) of the vehicle&#39;s registered holder. | [optional] 
**GivenName** | Pointer to **NullableString** | The given name(s) of the vehicle&#39;s registered holder. | [optional] 
**Address** | Pointer to **NullableString** | The registered address of the vehicle&#39;s holder, as a single formatted string. | [optional] 
**VehicleRight** | Pointer to **NullableString** | The legal relationship the registration holder has to the vehicle (e.g. owner vs. authorized user), as recorded by the issuer.              TODO: Values | [optional] 
**SpecialRemarks** | Pointer to **[]string** | Free-text remarks or annotations recorded on the registration certificate. | [optional] 
**IssueDate** | Pointer to **NullableString** | The date the registration certificate was issued. | [optional] 
**ExpiryDate** | Pointer to **NullableString** | The date the registration certificate expires. | [optional] 
**IssuingAuthority** | Pointer to **NullableString** | The name of the authority that issued the registration certificate. | [optional] 
**DocumentNumber** | Pointer to **NullableString** | The number of the registration certificate document itself. | [optional] 

## Methods

### NewMoldovaVehicleRegistrationCertificateCredential

`func NewMoldovaVehicleRegistrationCertificateCredential() *MoldovaVehicleRegistrationCertificateCredential`

NewMoldovaVehicleRegistrationCertificateCredential instantiates a new MoldovaVehicleRegistrationCertificateCredential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoldovaVehicleRegistrationCertificateCredentialWithDefaults

`func NewMoldovaVehicleRegistrationCertificateCredentialWithDefaults() *MoldovaVehicleRegistrationCertificateCredential`

NewMoldovaVehicleRegistrationCertificateCredentialWithDefaults instantiates a new MoldovaVehicleRegistrationCertificateCredential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlateNumber

`func (o *MoldovaVehicleRegistrationCertificateCredential) GetPlateNumber() string`

GetPlateNumber returns the PlateNumber field if non-nil, zero value otherwise.

### GetPlateNumberOk

`func (o *MoldovaVehicleRegistrationCertificateCredential) GetPlateNumberOk() (*string, bool)`

GetPlateNumberOk returns a tuple with the PlateNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlateNumber

`func (o *MoldovaVehicleRegistrationCertificateCredential) SetPlateNumber(v string)`

SetPlateNumber sets PlateNumber field to given value.

### HasPlateNumber

`func (o *MoldovaVehicleRegistrationCertificateCredential) HasPlateNumber() bool`

HasPlateNumber returns a boolean if a field has been set.

### SetPlateNumberNil

`func (o *MoldovaVehicleRegistrationCertificateCredential) SetPlateNumberNil(b bool)`

 SetPlateNumberNil sets the value for PlateNumber to be an explicit nil

### UnsetPlateNumber
`func (o *MoldovaVehicleRegistrationCertificateCredential) UnsetPlateNumber()`

UnsetPlateNumber ensures that no value is present for PlateNumber, not even an explicit nil
### GetIdnv

`func (o *MoldovaVehicleRegistrationCertificateCredential) GetIdnv() string`

GetIdnv returns the Idnv field if non-nil, zero value otherwise.

### GetIdnvOk

`func (o *MoldovaVehicleRegistrationCertificateCredential) GetIdnvOk() (*string, bool)`

GetIdnvOk returns a tuple with the Idnv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdnv

`func (o *MoldovaVehicleRegistrationCertificateCredential) SetIdnv(v string)`

SetIdnv sets Idnv field to given value.

### HasIdnv

`func (o *MoldovaVehicleRegistrationCertificateCredential) HasIdnv() bool`

HasIdnv returns a boolean if a field has been set.

### SetIdnvNil

`func (o *MoldovaVehicleRegistrationCertificateCredential) SetIdnvNil(b bool)`

 SetIdnvNil sets the value for Idnv to be an explicit nil

### UnsetIdnv
`func (o *MoldovaVehicleRegistrationCertificateCredential) UnsetIdnv()`

UnsetIdnv ensures that no value is present for Idnv, not even an explicit nil
### GetVin

`func (o *MoldovaVehicleRegistrationCertificateCredential) GetVin() string`

GetVin returns the Vin field if non-nil, zero value otherwise.

### GetVinOk

`func (o *MoldovaVehicleRegistrationCertificateCredential) GetVinOk() (*string, bool)`

GetVinOk returns a tuple with the Vin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVin

`func (o *MoldovaVehicleRegistrationCertificateCredential) SetVin(v string)`

SetVin sets Vin field to given value.

### HasVin

`func (o *MoldovaVehicleRegistrationCertificateCredential) HasVin() bool`

HasVin returns a boolean if a field has been set.

### SetVinNil

`func (o *MoldovaVehicleRegistrationCertificateCredential) SetVinNil(b bool)`

 SetVinNil sets the value for Vin to be an explicit nil

### UnsetVin
`func (o *MoldovaVehicleRegistrationCertificateCredential) UnsetVin()`

UnsetVin ensures that no value is present for Vin, not even an explicit nil
### GetMake

`func (o *MoldovaVehicleRegistrationCertificateCredential) GetMake() string`

GetMake returns the Make field if non-nil, zero value otherwise.

### GetMakeOk

`func (o *MoldovaVehicleRegistrationCertificateCredential) GetMakeOk() (*string, bool)`

GetMakeOk returns a tuple with the Make field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMake

`func (o *MoldovaVehicleRegistrationCertificateCredential) SetMake(v string)`

SetMake sets Make field to given value.

### HasMake

`func (o *MoldovaVehicleRegistrationCertificateCredential) HasMake() bool`

HasMake returns a boolean if a field has been set.

### SetMakeNil

`func (o *MoldovaVehicleRegistrationCertificateCredential) SetMakeNil(b bool)`

 SetMakeNil sets the value for Make to be an explicit nil

### UnsetMake
`func (o *MoldovaVehicleRegistrationCertificateCredential) UnsetMake()`

UnsetMake ensures that no value is present for Make, not even an explicit nil
### GetModel

`func (o *MoldovaVehicleRegistrationCertificateCredential) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *MoldovaVehicleRegistrationCertificateCredential) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *MoldovaVehicleRegistrationCertificateCredential) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *MoldovaVehicleRegistrationCertificateCredential) HasModel() bool`

HasModel returns a boolean if a field has been set.

### SetModelNil

`func (o *MoldovaVehicleRegistrationCertificateCredential) SetModelNil(b bool)`

 SetModelNil sets the value for Model to be an explicit nil

### UnsetModel
`func (o *MoldovaVehicleRegistrationCertificateCredential) UnsetModel()`

UnsetModel ensures that no value is present for Model, not even an explicit nil
### GetColor

`func (o *MoldovaVehicleRegistrationCertificateCredential) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *MoldovaVehicleRegistrationCertificateCredential) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *MoldovaVehicleRegistrationCertificateCredential) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *MoldovaVehicleRegistrationCertificateCredential) HasColor() bool`

HasColor returns a boolean if a field has been set.

### SetColorNil

`func (o *MoldovaVehicleRegistrationCertificateCredential) SetColorNil(b bool)`

 SetColorNil sets the value for Color to be an explicit nil

### UnsetColor
`func (o *MoldovaVehicleRegistrationCertificateCredential) UnsetColor()`

UnsetColor ensures that no value is present for Color, not even an explicit nil
### GetCategory

`func (o *MoldovaVehicleRegistrationCertificateCredential) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *MoldovaVehicleRegistrationCertificateCredential) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *MoldovaVehicleRegistrationCertificateCredential) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *MoldovaVehicleRegistrationCertificateCredential) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *MoldovaVehicleRegistrationCertificateCredential) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *MoldovaVehicleRegistrationCertificateCredential) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetYear

`func (o *MoldovaVehicleRegistrationCertificateCredential) GetYear() int32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *MoldovaVehicleRegistrationCertificateCredential) GetYearOk() (*int32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *MoldovaVehicleRegistrationCertificateCredential) SetYear(v int32)`

SetYear sets Year field to given value.

### HasYear

`func (o *MoldovaVehicleRegistrationCertificateCredential) HasYear() bool`

HasYear returns a boolean if a field has been set.

### SetYearNil

`func (o *MoldovaVehicleRegistrationCertificateCredential) SetYearNil(b bool)`

 SetYearNil sets the value for Year to be an explicit nil

### UnsetYear
`func (o *MoldovaVehicleRegistrationCertificateCredential) UnsetYear()`

UnsetYear ensures that no value is present for Year, not even an explicit nil
### GetBodyNumber

`func (o *MoldovaVehicleRegistrationCertificateCredential) GetBodyNumber() string`

GetBodyNumber returns the BodyNumber field if non-nil, zero value otherwise.

### GetBodyNumberOk

`func (o *MoldovaVehicleRegistrationCertificateCredential) GetBodyNumberOk() (*string, bool)`

GetBodyNumberOk returns a tuple with the BodyNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyNumber

`func (o *MoldovaVehicleRegistrationCertificateCredential) SetBodyNumber(v string)`

SetBodyNumber sets BodyNumber field to given value.

### HasBodyNumber

`func (o *MoldovaVehicleRegistrationCertificateCredential) HasBodyNumber() bool`

HasBodyNumber returns a boolean if a field has been set.

### SetBodyNumberNil

`func (o *MoldovaVehicleRegistrationCertificateCredential) SetBodyNumberNil(b bool)`

 SetBodyNumberNil sets the value for BodyNumber to be an explicit nil

### UnsetBodyNumber
`func (o *MoldovaVehicleRegistrationCertificateCredential) UnsetBodyNumber()`

UnsetBodyNumber ensures that no value is present for BodyNumber, not even an explicit nil
### GetBodyType

`func (o *MoldovaVehicleRegistrationCertificateCredential) GetBodyType() string`

GetBodyType returns the BodyType field if non-nil, zero value otherwise.

### GetBodyTypeOk

`func (o *MoldovaVehicleRegistrationCertificateCredential) GetBodyTypeOk() (*string, bool)`

GetBodyTypeOk returns a tuple with the BodyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyType

`func (o *MoldovaVehicleRegistrationCertificateCredential) SetBodyType(v string)`

SetBodyType sets BodyType field to given value.

### HasBodyType

`func (o *MoldovaVehicleRegistrationCertificateCredential) HasBodyType() bool`

HasBodyType returns a boolean if a field has been set.

### SetBodyTypeNil

`func (o *MoldovaVehicleRegistrationCertificateCredential) SetBodyTypeNil(b bool)`

 SetBodyTypeNil sets the value for BodyType to be an explicit nil

### UnsetBodyType
`func (o *MoldovaVehicleRegistrationCertificateCredential) UnsetBodyType()`

UnsetBodyType ensures that no value is present for BodyType, not even an explicit nil
### GetChassisNumber

`func (o *MoldovaVehicleRegistrationCertificateCredential) GetChassisNumber() string`

GetChassisNumber returns the ChassisNumber field if non-nil, zero value otherwise.

### GetChassisNumberOk

`func (o *MoldovaVehicleRegistrationCertificateCredential) GetChassisNumberOk() (*string, bool)`

GetChassisNumberOk returns a tuple with the ChassisNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassisNumber

`func (o *MoldovaVehicleRegistrationCertificateCredential) SetChassisNumber(v string)`

SetChassisNumber sets ChassisNumber field to given value.

### HasChassisNumber

`func (o *MoldovaVehicleRegistrationCertificateCredential) HasChassisNumber() bool`

HasChassisNumber returns a boolean if a field has been set.

### SetChassisNumberNil

`func (o *MoldovaVehicleRegistrationCertificateCredential) SetChassisNumberNil(b bool)`

 SetChassisNumberNil sets the value for ChassisNumber to be an explicit nil

### UnsetChassisNumber
`func (o *MoldovaVehicleRegistrationCertificateCredential) UnsetChassisNumber()`

UnsetChassisNumber ensures that no value is present for ChassisNumber, not even an explicit nil
### GetEngineVolume

`func (o *MoldovaVehicleRegistrationCertificateCredential) GetEngineVolume() string`

GetEngineVolume returns the EngineVolume field if non-nil, zero value otherwise.

### GetEngineVolumeOk

`func (o *MoldovaVehicleRegistrationCertificateCredential) GetEngineVolumeOk() (*string, bool)`

GetEngineVolumeOk returns a tuple with the EngineVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineVolume

`func (o *MoldovaVehicleRegistrationCertificateCredential) SetEngineVolume(v string)`

SetEngineVolume sets EngineVolume field to given value.

### HasEngineVolume

`func (o *MoldovaVehicleRegistrationCertificateCredential) HasEngineVolume() bool`

HasEngineVolume returns a boolean if a field has been set.

### SetEngineVolumeNil

`func (o *MoldovaVehicleRegistrationCertificateCredential) SetEngineVolumeNil(b bool)`

 SetEngineVolumeNil sets the value for EngineVolume to be an explicit nil

### UnsetEngineVolume
`func (o *MoldovaVehicleRegistrationCertificateCredential) UnsetEngineVolume()`

UnsetEngineVolume ensures that no value is present for EngineVolume, not even an explicit nil
### GetEngineType

`func (o *MoldovaVehicleRegistrationCertificateCredential) GetEngineType() string`

GetEngineType returns the EngineType field if non-nil, zero value otherwise.

### GetEngineTypeOk

`func (o *MoldovaVehicleRegistrationCertificateCredential) GetEngineTypeOk() (*string, bool)`

GetEngineTypeOk returns a tuple with the EngineType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineType

`func (o *MoldovaVehicleRegistrationCertificateCredential) SetEngineType(v string)`

SetEngineType sets EngineType field to given value.

### HasEngineType

`func (o *MoldovaVehicleRegistrationCertificateCredential) HasEngineType() bool`

HasEngineType returns a boolean if a field has been set.

### SetEngineTypeNil

`func (o *MoldovaVehicleRegistrationCertificateCredential) SetEngineTypeNil(b bool)`

 SetEngineTypeNil sets the value for EngineType to be an explicit nil

### UnsetEngineType
`func (o *MoldovaVehicleRegistrationCertificateCredential) UnsetEngineType()`

UnsetEngineType ensures that no value is present for EngineType, not even an explicit nil
### GetEngineNumber

`func (o *MoldovaVehicleRegistrationCertificateCredential) GetEngineNumber() string`

GetEngineNumber returns the EngineNumber field if non-nil, zero value otherwise.

### GetEngineNumberOk

`func (o *MoldovaVehicleRegistrationCertificateCredential) GetEngineNumberOk() (*string, bool)`

GetEngineNumberOk returns a tuple with the EngineNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineNumber

`func (o *MoldovaVehicleRegistrationCertificateCredential) SetEngineNumber(v string)`

SetEngineNumber sets EngineNumber field to given value.

### HasEngineNumber

`func (o *MoldovaVehicleRegistrationCertificateCredential) HasEngineNumber() bool`

HasEngineNumber returns a boolean if a field has been set.

### SetEngineNumberNil

`func (o *MoldovaVehicleRegistrationCertificateCredential) SetEngineNumberNil(b bool)`

 SetEngineNumberNil sets the value for EngineNumber to be an explicit nil

### UnsetEngineNumber
`func (o *MoldovaVehicleRegistrationCertificateCredential) UnsetEngineNumber()`

UnsetEngineNumber ensures that no value is present for EngineNumber, not even an explicit nil
### GetAuthorizedWeight

`func (o *MoldovaVehicleRegistrationCertificateCredential) GetAuthorizedWeight() int32`

GetAuthorizedWeight returns the AuthorizedWeight field if non-nil, zero value otherwise.

### GetAuthorizedWeightOk

`func (o *MoldovaVehicleRegistrationCertificateCredential) GetAuthorizedWeightOk() (*int32, bool)`

GetAuthorizedWeightOk returns a tuple with the AuthorizedWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedWeight

`func (o *MoldovaVehicleRegistrationCertificateCredential) SetAuthorizedWeight(v int32)`

SetAuthorizedWeight sets AuthorizedWeight field to given value.

### HasAuthorizedWeight

`func (o *MoldovaVehicleRegistrationCertificateCredential) HasAuthorizedWeight() bool`

HasAuthorizedWeight returns a boolean if a field has been set.

### SetAuthorizedWeightNil

`func (o *MoldovaVehicleRegistrationCertificateCredential) SetAuthorizedWeightNil(b bool)`

 SetAuthorizedWeightNil sets the value for AuthorizedWeight to be an explicit nil

### UnsetAuthorizedWeight
`func (o *MoldovaVehicleRegistrationCertificateCredential) UnsetAuthorizedWeight()`

UnsetAuthorizedWeight ensures that no value is present for AuthorizedWeight, not even an explicit nil
### GetWeight

`func (o *MoldovaVehicleRegistrationCertificateCredential) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *MoldovaVehicleRegistrationCertificateCredential) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *MoldovaVehicleRegistrationCertificateCredential) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *MoldovaVehicleRegistrationCertificateCredential) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### SetWeightNil

`func (o *MoldovaVehicleRegistrationCertificateCredential) SetWeightNil(b bool)`

 SetWeightNil sets the value for Weight to be an explicit nil

### UnsetWeight
`func (o *MoldovaVehicleRegistrationCertificateCredential) UnsetWeight()`

UnsetWeight ensures that no value is present for Weight, not even an explicit nil
### GetPlaces

`func (o *MoldovaVehicleRegistrationCertificateCredential) GetPlaces() int32`

GetPlaces returns the Places field if non-nil, zero value otherwise.

### GetPlacesOk

`func (o *MoldovaVehicleRegistrationCertificateCredential) GetPlacesOk() (*int32, bool)`

GetPlacesOk returns a tuple with the Places field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaces

`func (o *MoldovaVehicleRegistrationCertificateCredential) SetPlaces(v int32)`

SetPlaces sets Places field to given value.

### HasPlaces

`func (o *MoldovaVehicleRegistrationCertificateCredential) HasPlaces() bool`

HasPlaces returns a boolean if a field has been set.

### SetPlacesNil

`func (o *MoldovaVehicleRegistrationCertificateCredential) SetPlacesNil(b bool)`

 SetPlacesNil sets the value for Places to be an explicit nil

### UnsetPlaces
`func (o *MoldovaVehicleRegistrationCertificateCredential) UnsetPlaces()`

UnsetPlaces ensures that no value is present for Places, not even an explicit nil
### GetIdnp

`func (o *MoldovaVehicleRegistrationCertificateCredential) GetIdnp() string`

GetIdnp returns the Idnp field if non-nil, zero value otherwise.

### GetIdnpOk

`func (o *MoldovaVehicleRegistrationCertificateCredential) GetIdnpOk() (*string, bool)`

GetIdnpOk returns a tuple with the Idnp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdnp

`func (o *MoldovaVehicleRegistrationCertificateCredential) SetIdnp(v string)`

SetIdnp sets Idnp field to given value.

### HasIdnp

`func (o *MoldovaVehicleRegistrationCertificateCredential) HasIdnp() bool`

HasIdnp returns a boolean if a field has been set.

### SetIdnpNil

`func (o *MoldovaVehicleRegistrationCertificateCredential) SetIdnpNil(b bool)`

 SetIdnpNil sets the value for Idnp to be an explicit nil

### UnsetIdnp
`func (o *MoldovaVehicleRegistrationCertificateCredential) UnsetIdnp()`

UnsetIdnp ensures that no value is present for Idnp, not even an explicit nil
### GetFamilyName

`func (o *MoldovaVehicleRegistrationCertificateCredential) GetFamilyName() string`

GetFamilyName returns the FamilyName field if non-nil, zero value otherwise.

### GetFamilyNameOk

`func (o *MoldovaVehicleRegistrationCertificateCredential) GetFamilyNameOk() (*string, bool)`

GetFamilyNameOk returns a tuple with the FamilyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyName

`func (o *MoldovaVehicleRegistrationCertificateCredential) SetFamilyName(v string)`

SetFamilyName sets FamilyName field to given value.

### HasFamilyName

`func (o *MoldovaVehicleRegistrationCertificateCredential) HasFamilyName() bool`

HasFamilyName returns a boolean if a field has been set.

### SetFamilyNameNil

`func (o *MoldovaVehicleRegistrationCertificateCredential) SetFamilyNameNil(b bool)`

 SetFamilyNameNil sets the value for FamilyName to be an explicit nil

### UnsetFamilyName
`func (o *MoldovaVehicleRegistrationCertificateCredential) UnsetFamilyName()`

UnsetFamilyName ensures that no value is present for FamilyName, not even an explicit nil
### GetGivenName

`func (o *MoldovaVehicleRegistrationCertificateCredential) GetGivenName() string`

GetGivenName returns the GivenName field if non-nil, zero value otherwise.

### GetGivenNameOk

`func (o *MoldovaVehicleRegistrationCertificateCredential) GetGivenNameOk() (*string, bool)`

GetGivenNameOk returns a tuple with the GivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenName

`func (o *MoldovaVehicleRegistrationCertificateCredential) SetGivenName(v string)`

SetGivenName sets GivenName field to given value.

### HasGivenName

`func (o *MoldovaVehicleRegistrationCertificateCredential) HasGivenName() bool`

HasGivenName returns a boolean if a field has been set.

### SetGivenNameNil

`func (o *MoldovaVehicleRegistrationCertificateCredential) SetGivenNameNil(b bool)`

 SetGivenNameNil sets the value for GivenName to be an explicit nil

### UnsetGivenName
`func (o *MoldovaVehicleRegistrationCertificateCredential) UnsetGivenName()`

UnsetGivenName ensures that no value is present for GivenName, not even an explicit nil
### GetAddress

`func (o *MoldovaVehicleRegistrationCertificateCredential) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *MoldovaVehicleRegistrationCertificateCredential) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *MoldovaVehicleRegistrationCertificateCredential) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *MoldovaVehicleRegistrationCertificateCredential) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *MoldovaVehicleRegistrationCertificateCredential) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *MoldovaVehicleRegistrationCertificateCredential) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetVehicleRight

`func (o *MoldovaVehicleRegistrationCertificateCredential) GetVehicleRight() string`

GetVehicleRight returns the VehicleRight field if non-nil, zero value otherwise.

### GetVehicleRightOk

`func (o *MoldovaVehicleRegistrationCertificateCredential) GetVehicleRightOk() (*string, bool)`

GetVehicleRightOk returns a tuple with the VehicleRight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicleRight

`func (o *MoldovaVehicleRegistrationCertificateCredential) SetVehicleRight(v string)`

SetVehicleRight sets VehicleRight field to given value.

### HasVehicleRight

`func (o *MoldovaVehicleRegistrationCertificateCredential) HasVehicleRight() bool`

HasVehicleRight returns a boolean if a field has been set.

### SetVehicleRightNil

`func (o *MoldovaVehicleRegistrationCertificateCredential) SetVehicleRightNil(b bool)`

 SetVehicleRightNil sets the value for VehicleRight to be an explicit nil

### UnsetVehicleRight
`func (o *MoldovaVehicleRegistrationCertificateCredential) UnsetVehicleRight()`

UnsetVehicleRight ensures that no value is present for VehicleRight, not even an explicit nil
### GetSpecialRemarks

`func (o *MoldovaVehicleRegistrationCertificateCredential) GetSpecialRemarks() []string`

GetSpecialRemarks returns the SpecialRemarks field if non-nil, zero value otherwise.

### GetSpecialRemarksOk

`func (o *MoldovaVehicleRegistrationCertificateCredential) GetSpecialRemarksOk() (*[]string, bool)`

GetSpecialRemarksOk returns a tuple with the SpecialRemarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialRemarks

`func (o *MoldovaVehicleRegistrationCertificateCredential) SetSpecialRemarks(v []string)`

SetSpecialRemarks sets SpecialRemarks field to given value.

### HasSpecialRemarks

`func (o *MoldovaVehicleRegistrationCertificateCredential) HasSpecialRemarks() bool`

HasSpecialRemarks returns a boolean if a field has been set.

### SetSpecialRemarksNil

`func (o *MoldovaVehicleRegistrationCertificateCredential) SetSpecialRemarksNil(b bool)`

 SetSpecialRemarksNil sets the value for SpecialRemarks to be an explicit nil

### UnsetSpecialRemarks
`func (o *MoldovaVehicleRegistrationCertificateCredential) UnsetSpecialRemarks()`

UnsetSpecialRemarks ensures that no value is present for SpecialRemarks, not even an explicit nil
### GetIssueDate

`func (o *MoldovaVehicleRegistrationCertificateCredential) GetIssueDate() string`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *MoldovaVehicleRegistrationCertificateCredential) GetIssueDateOk() (*string, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *MoldovaVehicleRegistrationCertificateCredential) SetIssueDate(v string)`

SetIssueDate sets IssueDate field to given value.

### HasIssueDate

`func (o *MoldovaVehicleRegistrationCertificateCredential) HasIssueDate() bool`

HasIssueDate returns a boolean if a field has been set.

### SetIssueDateNil

`func (o *MoldovaVehicleRegistrationCertificateCredential) SetIssueDateNil(b bool)`

 SetIssueDateNil sets the value for IssueDate to be an explicit nil

### UnsetIssueDate
`func (o *MoldovaVehicleRegistrationCertificateCredential) UnsetIssueDate()`

UnsetIssueDate ensures that no value is present for IssueDate, not even an explicit nil
### GetExpiryDate

`func (o *MoldovaVehicleRegistrationCertificateCredential) GetExpiryDate() string`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *MoldovaVehicleRegistrationCertificateCredential) GetExpiryDateOk() (*string, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *MoldovaVehicleRegistrationCertificateCredential) SetExpiryDate(v string)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *MoldovaVehicleRegistrationCertificateCredential) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### SetExpiryDateNil

`func (o *MoldovaVehicleRegistrationCertificateCredential) SetExpiryDateNil(b bool)`

 SetExpiryDateNil sets the value for ExpiryDate to be an explicit nil

### UnsetExpiryDate
`func (o *MoldovaVehicleRegistrationCertificateCredential) UnsetExpiryDate()`

UnsetExpiryDate ensures that no value is present for ExpiryDate, not even an explicit nil
### GetIssuingAuthority

`func (o *MoldovaVehicleRegistrationCertificateCredential) GetIssuingAuthority() string`

GetIssuingAuthority returns the IssuingAuthority field if non-nil, zero value otherwise.

### GetIssuingAuthorityOk

`func (o *MoldovaVehicleRegistrationCertificateCredential) GetIssuingAuthorityOk() (*string, bool)`

GetIssuingAuthorityOk returns a tuple with the IssuingAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingAuthority

`func (o *MoldovaVehicleRegistrationCertificateCredential) SetIssuingAuthority(v string)`

SetIssuingAuthority sets IssuingAuthority field to given value.

### HasIssuingAuthority

`func (o *MoldovaVehicleRegistrationCertificateCredential) HasIssuingAuthority() bool`

HasIssuingAuthority returns a boolean if a field has been set.

### SetIssuingAuthorityNil

`func (o *MoldovaVehicleRegistrationCertificateCredential) SetIssuingAuthorityNil(b bool)`

 SetIssuingAuthorityNil sets the value for IssuingAuthority to be an explicit nil

### UnsetIssuingAuthority
`func (o *MoldovaVehicleRegistrationCertificateCredential) UnsetIssuingAuthority()`

UnsetIssuingAuthority ensures that no value is present for IssuingAuthority, not even an explicit nil
### GetDocumentNumber

`func (o *MoldovaVehicleRegistrationCertificateCredential) GetDocumentNumber() string`

GetDocumentNumber returns the DocumentNumber field if non-nil, zero value otherwise.

### GetDocumentNumberOk

`func (o *MoldovaVehicleRegistrationCertificateCredential) GetDocumentNumberOk() (*string, bool)`

GetDocumentNumberOk returns a tuple with the DocumentNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentNumber

`func (o *MoldovaVehicleRegistrationCertificateCredential) SetDocumentNumber(v string)`

SetDocumentNumber sets DocumentNumber field to given value.

### HasDocumentNumber

`func (o *MoldovaVehicleRegistrationCertificateCredential) HasDocumentNumber() bool`

HasDocumentNumber returns a boolean if a field has been set.

### SetDocumentNumberNil

`func (o *MoldovaVehicleRegistrationCertificateCredential) SetDocumentNumberNil(b bool)`

 SetDocumentNumberNil sets the value for DocumentNumber to be an explicit nil

### UnsetDocumentNumber
`func (o *MoldovaVehicleRegistrationCertificateCredential) UnsetDocumentNumber()`

UnsetDocumentNumber ensures that no value is present for DocumentNumber, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


