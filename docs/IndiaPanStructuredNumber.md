# IndiaPanStructuredNumber

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SeriesCode** | **string** | Three-letter block issued by the Income Tax Department of India.              Ties the number to a specific Income Tax Office or jurisdiction using an internal labeling system.              Position: - Characters 1 through 3 of the PAN              Possible values: - Any three-letter block of uppercase Latin letters | 
**AssesseeCategoryCode** | **string** | Code that represents the entity that is subject to income tax in India.              Position: - Character 4 of the PAN              Possible values: - \&quot;A\&quot; for Association of Persons (AOP) - \&quot;B\&quot; for Body of Individuals (BOI) - \&quot;C\&quot; for Company - \&quot;F\&quot; for Firm - \&quot;G\&quot; for Government - \&quot;H\&quot; for Hindu Undivided Family (HUF) - \&quot;J\&quot; for Artificial Juridical Person - \&quot;L\&quot; for Local Authority - \&quot;P\&quot; for Individual - \&quot;T\&quot; for Trust | 
**AssesseeCategoryName** | Pointer to **NullableString** | Human-readable label for the assessee category when it matches a known code.              Possible values: - Association of Persons (AOP) - Body of Individuals (BOI) - Company - Firm - Government - Hindu Undivided Family (HUF) - Artificial Juridical Person - Local Authority - Individual - Trust              Omitted when the letter is not a standard category code. | [optional] 
**NamePrefixLetter** | **string** | First letter of the name of the entity that is subject to income tax in India.              Position: - Character 5 of the PAN.              Format: - When AssesseeCategoryCode is P, this character is the first letter of the individual&#39;s   family name as recorded for the PAN. - When AssesseeCategoryCode is A, B, C, F, G, H, J, L, or T, this character is the first letter   of the name of the association, body, company, firm, government, HUF, local authority, or   trust.              Possible values: - Any uppercase Latin letter | 
**SerialNumber** | **string** | Sequential four-digit field assigned by the Income Tax Department of India.              Each PAN receives a unique serial within the given series code and assessee category.              Position: - Characters 6 through 9 of the PAN.              Possible values: - Four-digit strings 0001 through 9999 inclusive, using only digits 0 through 9 | 
**CheckLetter** | **string** | Alphabetic check character for the first nine characters of the PAN.              The algorithm has not been made publicly available by the Income Tax Department of India.              Position: - Character 10 of the PAN.              Possible values: - One uppercase Latin letter A through Z | 

## Methods

### NewIndiaPanStructuredNumber

`func NewIndiaPanStructuredNumber(seriesCode string, assesseeCategoryCode string, namePrefixLetter string, serialNumber string, checkLetter string, ) *IndiaPanStructuredNumber`

NewIndiaPanStructuredNumber instantiates a new IndiaPanStructuredNumber object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndiaPanStructuredNumberWithDefaults

`func NewIndiaPanStructuredNumberWithDefaults() *IndiaPanStructuredNumber`

NewIndiaPanStructuredNumberWithDefaults instantiates a new IndiaPanStructuredNumber object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSeriesCode

`func (o *IndiaPanStructuredNumber) GetSeriesCode() string`

GetSeriesCode returns the SeriesCode field if non-nil, zero value otherwise.

### GetSeriesCodeOk

`func (o *IndiaPanStructuredNumber) GetSeriesCodeOk() (*string, bool)`

GetSeriesCodeOk returns a tuple with the SeriesCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesCode

`func (o *IndiaPanStructuredNumber) SetSeriesCode(v string)`

SetSeriesCode sets SeriesCode field to given value.


### GetAssesseeCategoryCode

`func (o *IndiaPanStructuredNumber) GetAssesseeCategoryCode() string`

GetAssesseeCategoryCode returns the AssesseeCategoryCode field if non-nil, zero value otherwise.

### GetAssesseeCategoryCodeOk

`func (o *IndiaPanStructuredNumber) GetAssesseeCategoryCodeOk() (*string, bool)`

GetAssesseeCategoryCodeOk returns a tuple with the AssesseeCategoryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssesseeCategoryCode

`func (o *IndiaPanStructuredNumber) SetAssesseeCategoryCode(v string)`

SetAssesseeCategoryCode sets AssesseeCategoryCode field to given value.


### GetAssesseeCategoryName

`func (o *IndiaPanStructuredNumber) GetAssesseeCategoryName() string`

GetAssesseeCategoryName returns the AssesseeCategoryName field if non-nil, zero value otherwise.

### GetAssesseeCategoryNameOk

`func (o *IndiaPanStructuredNumber) GetAssesseeCategoryNameOk() (*string, bool)`

GetAssesseeCategoryNameOk returns a tuple with the AssesseeCategoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssesseeCategoryName

`func (o *IndiaPanStructuredNumber) SetAssesseeCategoryName(v string)`

SetAssesseeCategoryName sets AssesseeCategoryName field to given value.

### HasAssesseeCategoryName

`func (o *IndiaPanStructuredNumber) HasAssesseeCategoryName() bool`

HasAssesseeCategoryName returns a boolean if a field has been set.

### SetAssesseeCategoryNameNil

`func (o *IndiaPanStructuredNumber) SetAssesseeCategoryNameNil(b bool)`

 SetAssesseeCategoryNameNil sets the value for AssesseeCategoryName to be an explicit nil

### UnsetAssesseeCategoryName
`func (o *IndiaPanStructuredNumber) UnsetAssesseeCategoryName()`

UnsetAssesseeCategoryName ensures that no value is present for AssesseeCategoryName, not even an explicit nil
### GetNamePrefixLetter

`func (o *IndiaPanStructuredNumber) GetNamePrefixLetter() string`

GetNamePrefixLetter returns the NamePrefixLetter field if non-nil, zero value otherwise.

### GetNamePrefixLetterOk

`func (o *IndiaPanStructuredNumber) GetNamePrefixLetterOk() (*string, bool)`

GetNamePrefixLetterOk returns a tuple with the NamePrefixLetter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamePrefixLetter

`func (o *IndiaPanStructuredNumber) SetNamePrefixLetter(v string)`

SetNamePrefixLetter sets NamePrefixLetter field to given value.


### GetSerialNumber

`func (o *IndiaPanStructuredNumber) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *IndiaPanStructuredNumber) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *IndiaPanStructuredNumber) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.


### GetCheckLetter

`func (o *IndiaPanStructuredNumber) GetCheckLetter() string`

GetCheckLetter returns the CheckLetter field if non-nil, zero value otherwise.

### GetCheckLetterOk

`func (o *IndiaPanStructuredNumber) GetCheckLetterOk() (*string, bool)`

GetCheckLetterOk returns a tuple with the CheckLetter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckLetter

`func (o *IndiaPanStructuredNumber) SetCheckLetter(v string)`

SetCheckLetter sets CheckLetter field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


