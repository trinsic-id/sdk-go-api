# KenyaNidLookup2Address

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**District** | Pointer to **NullableString** | The district where the person resides.              This is the highest level of the pre-2010 administrative hierarchy stored in IPRS. Districts were the primary administrative units under Kenya&#39;s provincial system before the change to counties in 2010. Districts roughly correspond to sub-counties in the current administrative structure. | [optional] 
**Division** | Pointer to **NullableString** | The division within the district.              This is the second level of the pre-2010 administrative hierarchy. Divisions were administrative units between districts and locations. | [optional] 
**Location** | Pointer to **NullableString** | The location within the division.              This is the third level of the pre-2010 administrative hierarchy. Locations were the smallest administrative units. | [optional] 
**AdditionalLines** | Pointer to **[]string** | Additional address lines that appear before the structured administrative parts.              May include P.O. Box (format: \&quot;P O BOX [number] [location]\&quot;), village names, estate names, or other informal locality information. | [optional] 
**Raw** | Pointer to **NullableString** | The raw, unparsed address string as returned from IPRS (Integrated Population Registration System).              This preserves the original newline-separated format containing both informal address components (village, P.O. Box) and labeled administrative divisions (LOCATION, DIVISION, DISTRICT). | [optional] 

## Methods

### NewKenyaNidLookup2Address

`func NewKenyaNidLookup2Address() *KenyaNidLookup2Address`

NewKenyaNidLookup2Address instantiates a new KenyaNidLookup2Address object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKenyaNidLookup2AddressWithDefaults

`func NewKenyaNidLookup2AddressWithDefaults() *KenyaNidLookup2Address`

NewKenyaNidLookup2AddressWithDefaults instantiates a new KenyaNidLookup2Address object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDistrict

`func (o *KenyaNidLookup2Address) GetDistrict() string`

GetDistrict returns the District field if non-nil, zero value otherwise.

### GetDistrictOk

`func (o *KenyaNidLookup2Address) GetDistrictOk() (*string, bool)`

GetDistrictOk returns a tuple with the District field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistrict

`func (o *KenyaNidLookup2Address) SetDistrict(v string)`

SetDistrict sets District field to given value.

### HasDistrict

`func (o *KenyaNidLookup2Address) HasDistrict() bool`

HasDistrict returns a boolean if a field has been set.

### SetDistrictNil

`func (o *KenyaNidLookup2Address) SetDistrictNil(b bool)`

 SetDistrictNil sets the value for District to be an explicit nil

### UnsetDistrict
`func (o *KenyaNidLookup2Address) UnsetDistrict()`

UnsetDistrict ensures that no value is present for District, not even an explicit nil
### GetDivision

`func (o *KenyaNidLookup2Address) GetDivision() string`

GetDivision returns the Division field if non-nil, zero value otherwise.

### GetDivisionOk

`func (o *KenyaNidLookup2Address) GetDivisionOk() (*string, bool)`

GetDivisionOk returns a tuple with the Division field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivision

`func (o *KenyaNidLookup2Address) SetDivision(v string)`

SetDivision sets Division field to given value.

### HasDivision

`func (o *KenyaNidLookup2Address) HasDivision() bool`

HasDivision returns a boolean if a field has been set.

### SetDivisionNil

`func (o *KenyaNidLookup2Address) SetDivisionNil(b bool)`

 SetDivisionNil sets the value for Division to be an explicit nil

### UnsetDivision
`func (o *KenyaNidLookup2Address) UnsetDivision()`

UnsetDivision ensures that no value is present for Division, not even an explicit nil
### GetLocation

`func (o *KenyaNidLookup2Address) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *KenyaNidLookup2Address) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *KenyaNidLookup2Address) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *KenyaNidLookup2Address) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *KenyaNidLookup2Address) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *KenyaNidLookup2Address) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetAdditionalLines

`func (o *KenyaNidLookup2Address) GetAdditionalLines() []string`

GetAdditionalLines returns the AdditionalLines field if non-nil, zero value otherwise.

### GetAdditionalLinesOk

`func (o *KenyaNidLookup2Address) GetAdditionalLinesOk() (*[]string, bool)`

GetAdditionalLinesOk returns a tuple with the AdditionalLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalLines

`func (o *KenyaNidLookup2Address) SetAdditionalLines(v []string)`

SetAdditionalLines sets AdditionalLines field to given value.

### HasAdditionalLines

`func (o *KenyaNidLookup2Address) HasAdditionalLines() bool`

HasAdditionalLines returns a boolean if a field has been set.

### SetAdditionalLinesNil

`func (o *KenyaNidLookup2Address) SetAdditionalLinesNil(b bool)`

 SetAdditionalLinesNil sets the value for AdditionalLines to be an explicit nil

### UnsetAdditionalLines
`func (o *KenyaNidLookup2Address) UnsetAdditionalLines()`

UnsetAdditionalLines ensures that no value is present for AdditionalLines, not even an explicit nil
### GetRaw

`func (o *KenyaNidLookup2Address) GetRaw() string`

GetRaw returns the Raw field if non-nil, zero value otherwise.

### GetRawOk

`func (o *KenyaNidLookup2Address) GetRawOk() (*string, bool)`

GetRawOk returns a tuple with the Raw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaw

`func (o *KenyaNidLookup2Address) SetRaw(v string)`

SetRaw sets Raw field to given value.

### HasRaw

`func (o *KenyaNidLookup2Address) HasRaw() bool`

HasRaw returns a boolean if a field has been set.

### SetRawNil

`func (o *KenyaNidLookup2Address) SetRawNil(b bool)`

 SetRawNil sets the value for Raw to be an explicit nil

### UnsetRaw
`func (o *KenyaNidLookup2Address) UnsetRaw()`

UnsetRaw ensures that no value is present for Raw, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


