# ClearProviderOutputSessionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Location** | Pointer to [**NullableClearProviderOutputLocationInfo**](ClearProviderOutputLocationInfo.md) | Location information gathered from the individual&#39;s frontend session. | [optional] 
**UserAgent** | Pointer to **NullableString** | The browser user agent for this frontend session. | [optional] 
**CreatedAt** | Pointer to **NullableTime** | The time this frontend session was created, as a UTC timestamp. | [optional] 
**Ip** | Pointer to **NullableString** | The IP address for this frontend session. | [optional] 
**Locale** | Pointer to **NullableString** | The individual&#39;s locale for this frontend session. | [optional] 
**ZoneInfo** | Pointer to **NullableString** | The individual&#39;s IANA time zone for this frontend session. | [optional] 

## Methods

### NewClearProviderOutputSessionInfo

`func NewClearProviderOutputSessionInfo() *ClearProviderOutputSessionInfo`

NewClearProviderOutputSessionInfo instantiates a new ClearProviderOutputSessionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClearProviderOutputSessionInfoWithDefaults

`func NewClearProviderOutputSessionInfoWithDefaults() *ClearProviderOutputSessionInfo`

NewClearProviderOutputSessionInfoWithDefaults instantiates a new ClearProviderOutputSessionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocation

`func (o *ClearProviderOutputSessionInfo) GetLocation() ClearProviderOutputLocationInfo`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ClearProviderOutputSessionInfo) GetLocationOk() (*ClearProviderOutputLocationInfo, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ClearProviderOutputSessionInfo) SetLocation(v ClearProviderOutputLocationInfo)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ClearProviderOutputSessionInfo) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *ClearProviderOutputSessionInfo) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *ClearProviderOutputSessionInfo) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetUserAgent

`func (o *ClearProviderOutputSessionInfo) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *ClearProviderOutputSessionInfo) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *ClearProviderOutputSessionInfo) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *ClearProviderOutputSessionInfo) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.

### SetUserAgentNil

`func (o *ClearProviderOutputSessionInfo) SetUserAgentNil(b bool)`

 SetUserAgentNil sets the value for UserAgent to be an explicit nil

### UnsetUserAgent
`func (o *ClearProviderOutputSessionInfo) UnsetUserAgent()`

UnsetUserAgent ensures that no value is present for UserAgent, not even an explicit nil
### GetCreatedAt

`func (o *ClearProviderOutputSessionInfo) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ClearProviderOutputSessionInfo) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ClearProviderOutputSessionInfo) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ClearProviderOutputSessionInfo) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *ClearProviderOutputSessionInfo) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *ClearProviderOutputSessionInfo) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetIp

`func (o *ClearProviderOutputSessionInfo) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *ClearProviderOutputSessionInfo) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *ClearProviderOutputSessionInfo) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *ClearProviderOutputSessionInfo) HasIp() bool`

HasIp returns a boolean if a field has been set.

### SetIpNil

`func (o *ClearProviderOutputSessionInfo) SetIpNil(b bool)`

 SetIpNil sets the value for Ip to be an explicit nil

### UnsetIp
`func (o *ClearProviderOutputSessionInfo) UnsetIp()`

UnsetIp ensures that no value is present for Ip, not even an explicit nil
### GetLocale

`func (o *ClearProviderOutputSessionInfo) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *ClearProviderOutputSessionInfo) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *ClearProviderOutputSessionInfo) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *ClearProviderOutputSessionInfo) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### SetLocaleNil

`func (o *ClearProviderOutputSessionInfo) SetLocaleNil(b bool)`

 SetLocaleNil sets the value for Locale to be an explicit nil

### UnsetLocale
`func (o *ClearProviderOutputSessionInfo) UnsetLocale()`

UnsetLocale ensures that no value is present for Locale, not even an explicit nil
### GetZoneInfo

`func (o *ClearProviderOutputSessionInfo) GetZoneInfo() string`

GetZoneInfo returns the ZoneInfo field if non-nil, zero value otherwise.

### GetZoneInfoOk

`func (o *ClearProviderOutputSessionInfo) GetZoneInfoOk() (*string, bool)`

GetZoneInfoOk returns a tuple with the ZoneInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneInfo

`func (o *ClearProviderOutputSessionInfo) SetZoneInfo(v string)`

SetZoneInfo sets ZoneInfo field to given value.

### HasZoneInfo

`func (o *ClearProviderOutputSessionInfo) HasZoneInfo() bool`

HasZoneInfo returns a boolean if a field has been set.

### SetZoneInfoNil

`func (o *ClearProviderOutputSessionInfo) SetZoneInfoNil(b bool)`

 SetZoneInfoNil sets the value for ZoneInfo to be an explicit nil

### UnsetZoneInfo
`func (o *ClearProviderOutputSessionInfo) UnsetZoneInfo()`

UnsetZoneInfo ensures that no value is present for ZoneInfo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


