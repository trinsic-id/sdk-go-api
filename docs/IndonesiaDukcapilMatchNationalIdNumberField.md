# IndonesiaDukcapilMatchNationalIdNumberField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InputValue** | **string** | The NIK (Nomor Induk Kependudukan) submitted for this check.              NIK is Indonesia&#39;s unique population identity number, issued by Indonesia&#39;s population administration and civil registration authority (Dukcapil) under the Ministry of Home Affairs.              Format: - 16 numeric digits. - Digits 1-2 are the province code. - Digits 3-4 are the regency or city code within that province. - Digits 5-6 are the district code within that regency or city. - Digits 7-12 encode date of birth as DDMMYY. For female NIK holders, the day is increased by 40. - Digits 13-16 are an issuance serial number. | 
**Outcome** | **float64** | The provider assessment result for the submitted NIK.              Local integration code expects the provider to return 0 or 1 for this assessment. | 

## Methods

### NewIndonesiaDukcapilMatchNationalIdNumberField

`func NewIndonesiaDukcapilMatchNationalIdNumberField(inputValue string, outcome float64, ) *IndonesiaDukcapilMatchNationalIdNumberField`

NewIndonesiaDukcapilMatchNationalIdNumberField instantiates a new IndonesiaDukcapilMatchNationalIdNumberField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndonesiaDukcapilMatchNationalIdNumberFieldWithDefaults

`func NewIndonesiaDukcapilMatchNationalIdNumberFieldWithDefaults() *IndonesiaDukcapilMatchNationalIdNumberField`

NewIndonesiaDukcapilMatchNationalIdNumberFieldWithDefaults instantiates a new IndonesiaDukcapilMatchNationalIdNumberField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInputValue

`func (o *IndonesiaDukcapilMatchNationalIdNumberField) GetInputValue() string`

GetInputValue returns the InputValue field if non-nil, zero value otherwise.

### GetInputValueOk

`func (o *IndonesiaDukcapilMatchNationalIdNumberField) GetInputValueOk() (*string, bool)`

GetInputValueOk returns a tuple with the InputValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputValue

`func (o *IndonesiaDukcapilMatchNationalIdNumberField) SetInputValue(v string)`

SetInputValue sets InputValue field to given value.


### GetOutcome

`func (o *IndonesiaDukcapilMatchNationalIdNumberField) GetOutcome() float64`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *IndonesiaDukcapilMatchNationalIdNumberField) GetOutcomeOk() (*float64, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *IndonesiaDukcapilMatchNationalIdNumberField) SetOutcome(v float64)`

SetOutcome sets Outcome field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


