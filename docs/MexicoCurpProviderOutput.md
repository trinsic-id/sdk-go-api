# MexicoCurpProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurpStatus** | Pointer to **NullableString** | Curp status for the subject.              Possible values: - AN: Alta Normal (Normal registration) - Active - AH: Alta con Homonimia (Registration with homonymy) - Active - RCC: Registro de cambio afectando a CURP (Change affecting CURP) - Active - RCN: Registro de cambio no afectando a CURP (Change not affecting CURP) - Active - BAP: Baja por documento apócrifo (Low due to apocryphal document) - Inactive - BSU: Baja sin uso (Low curp without use) - Inactive - BD: Baja por defunción (Low curp due to death) - Inactive - BDM: Baja administrativa (Low, due to administrative process) - Inactive - BDP: Baja por adopción (Low, due to adoption) - Inactive - BJD: Baja Judicial (Low for judicial reasons) - Inactive | [optional] 

## Methods

### NewMexicoCurpProviderOutput

`func NewMexicoCurpProviderOutput() *MexicoCurpProviderOutput`

NewMexicoCurpProviderOutput instantiates a new MexicoCurpProviderOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMexicoCurpProviderOutputWithDefaults

`func NewMexicoCurpProviderOutputWithDefaults() *MexicoCurpProviderOutput`

NewMexicoCurpProviderOutputWithDefaults instantiates a new MexicoCurpProviderOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurpStatus

`func (o *MexicoCurpProviderOutput) GetCurpStatus() string`

GetCurpStatus returns the CurpStatus field if non-nil, zero value otherwise.

### GetCurpStatusOk

`func (o *MexicoCurpProviderOutput) GetCurpStatusOk() (*string, bool)`

GetCurpStatusOk returns a tuple with the CurpStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurpStatus

`func (o *MexicoCurpProviderOutput) SetCurpStatus(v string)`

SetCurpStatus sets CurpStatus field to given value.

### HasCurpStatus

`func (o *MexicoCurpProviderOutput) HasCurpStatus() bool`

HasCurpStatus returns a boolean if a field has been set.

### SetCurpStatusNil

`func (o *MexicoCurpProviderOutput) SetCurpStatusNil(b bool)`

 SetCurpStatusNil sets the value for CurpStatus to be an explicit nil

### UnsetCurpStatus
`func (o *MexicoCurpProviderOutput) UnsetCurpStatus()`

UnsetCurpStatus ensures that no value is present for CurpStatus, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


