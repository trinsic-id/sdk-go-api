# ItsmePlaceOfBirth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**City** | Pointer to **NullableString** | The city component of the place of birth.              Availability by ID document issuing country: best effort for Belgian-issued ID documents; not returned for other supported issuing countries. | [optional] 
**Formatted** | Pointer to **NullableString** | The formatted place of birth.              Availability by ID document issuing country: best effort for Belgian-issued ID documents; not returned for other supported issuing countries. | [optional] 

## Methods

### NewItsmePlaceOfBirth

`func NewItsmePlaceOfBirth() *ItsmePlaceOfBirth`

NewItsmePlaceOfBirth instantiates a new ItsmePlaceOfBirth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItsmePlaceOfBirthWithDefaults

`func NewItsmePlaceOfBirthWithDefaults() *ItsmePlaceOfBirth`

NewItsmePlaceOfBirthWithDefaults instantiates a new ItsmePlaceOfBirth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCity

`func (o *ItsmePlaceOfBirth) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *ItsmePlaceOfBirth) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *ItsmePlaceOfBirth) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *ItsmePlaceOfBirth) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *ItsmePlaceOfBirth) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *ItsmePlaceOfBirth) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetFormatted

`func (o *ItsmePlaceOfBirth) GetFormatted() string`

GetFormatted returns the Formatted field if non-nil, zero value otherwise.

### GetFormattedOk

`func (o *ItsmePlaceOfBirth) GetFormattedOk() (*string, bool)`

GetFormattedOk returns a tuple with the Formatted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatted

`func (o *ItsmePlaceOfBirth) SetFormatted(v string)`

SetFormatted sets Formatted field to given value.

### HasFormatted

`func (o *ItsmePlaceOfBirth) HasFormatted() bool`

HasFormatted returns a boolean if a field has been set.

### SetFormattedNil

`func (o *ItsmePlaceOfBirth) SetFormattedNil(b bool)`

 SetFormattedNil sets the value for Formatted to be an explicit nil

### UnsetFormatted
`func (o *ItsmePlaceOfBirth) UnsetFormatted()`

UnsetFormatted ensures that no value is present for Formatted, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


