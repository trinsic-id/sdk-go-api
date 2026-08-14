# Iso180135DrivingPrivilege

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VehicleCategoryCode** | Pointer to **NullableString** | Vehicle category code that the holder is authorized to operate. | [optional] 
**IssueDate** | Pointer to **NullableString** | Date when this driving privilege was issued. | [optional] 
**ExpiryDate** | Pointer to **NullableString** | Date when this driving privilege expires. | [optional] 
**Codes** | Pointer to [**[]Iso180135DrivingPrivilegeCode**](Iso180135DrivingPrivilegeCode.md) | Restriction, condition, or special privilege codes attached to this entry. | [optional] 

## Methods

### NewIso180135DrivingPrivilege

`func NewIso180135DrivingPrivilege() *Iso180135DrivingPrivilege`

NewIso180135DrivingPrivilege instantiates a new Iso180135DrivingPrivilege object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIso180135DrivingPrivilegeWithDefaults

`func NewIso180135DrivingPrivilegeWithDefaults() *Iso180135DrivingPrivilege`

NewIso180135DrivingPrivilegeWithDefaults instantiates a new Iso180135DrivingPrivilege object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVehicleCategoryCode

`func (o *Iso180135DrivingPrivilege) GetVehicleCategoryCode() string`

GetVehicleCategoryCode returns the VehicleCategoryCode field if non-nil, zero value otherwise.

### GetVehicleCategoryCodeOk

`func (o *Iso180135DrivingPrivilege) GetVehicleCategoryCodeOk() (*string, bool)`

GetVehicleCategoryCodeOk returns a tuple with the VehicleCategoryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicleCategoryCode

`func (o *Iso180135DrivingPrivilege) SetVehicleCategoryCode(v string)`

SetVehicleCategoryCode sets VehicleCategoryCode field to given value.

### HasVehicleCategoryCode

`func (o *Iso180135DrivingPrivilege) HasVehicleCategoryCode() bool`

HasVehicleCategoryCode returns a boolean if a field has been set.

### SetVehicleCategoryCodeNil

`func (o *Iso180135DrivingPrivilege) SetVehicleCategoryCodeNil(b bool)`

 SetVehicleCategoryCodeNil sets the value for VehicleCategoryCode to be an explicit nil

### UnsetVehicleCategoryCode
`func (o *Iso180135DrivingPrivilege) UnsetVehicleCategoryCode()`

UnsetVehicleCategoryCode ensures that no value is present for VehicleCategoryCode, not even an explicit nil
### GetIssueDate

`func (o *Iso180135DrivingPrivilege) GetIssueDate() string`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *Iso180135DrivingPrivilege) GetIssueDateOk() (*string, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *Iso180135DrivingPrivilege) SetIssueDate(v string)`

SetIssueDate sets IssueDate field to given value.

### HasIssueDate

`func (o *Iso180135DrivingPrivilege) HasIssueDate() bool`

HasIssueDate returns a boolean if a field has been set.

### SetIssueDateNil

`func (o *Iso180135DrivingPrivilege) SetIssueDateNil(b bool)`

 SetIssueDateNil sets the value for IssueDate to be an explicit nil

### UnsetIssueDate
`func (o *Iso180135DrivingPrivilege) UnsetIssueDate()`

UnsetIssueDate ensures that no value is present for IssueDate, not even an explicit nil
### GetExpiryDate

`func (o *Iso180135DrivingPrivilege) GetExpiryDate() string`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *Iso180135DrivingPrivilege) GetExpiryDateOk() (*string, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *Iso180135DrivingPrivilege) SetExpiryDate(v string)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *Iso180135DrivingPrivilege) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### SetExpiryDateNil

`func (o *Iso180135DrivingPrivilege) SetExpiryDateNil(b bool)`

 SetExpiryDateNil sets the value for ExpiryDate to be an explicit nil

### UnsetExpiryDate
`func (o *Iso180135DrivingPrivilege) UnsetExpiryDate()`

UnsetExpiryDate ensures that no value is present for ExpiryDate, not even an explicit nil
### GetCodes

`func (o *Iso180135DrivingPrivilege) GetCodes() []Iso180135DrivingPrivilegeCode`

GetCodes returns the Codes field if non-nil, zero value otherwise.

### GetCodesOk

`func (o *Iso180135DrivingPrivilege) GetCodesOk() (*[]Iso180135DrivingPrivilegeCode, bool)`

GetCodesOk returns a tuple with the Codes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodes

`func (o *Iso180135DrivingPrivilege) SetCodes(v []Iso180135DrivingPrivilegeCode)`

SetCodes sets Codes field to given value.

### HasCodes

`func (o *Iso180135DrivingPrivilege) HasCodes() bool`

HasCodes returns a boolean if a field has been set.

### SetCodesNil

`func (o *Iso180135DrivingPrivilege) SetCodesNil(b bool)`

 SetCodesNil sets the value for Codes to be an explicit nil

### UnsetCodes
`func (o *Iso180135DrivingPrivilege) UnsetCodes()`

UnsetCodes ensures that no value is present for Codes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


