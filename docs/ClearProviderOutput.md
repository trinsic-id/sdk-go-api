# ClearProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authenticated** | Pointer to **NullableBool** | Whether CLEAR authenticated the individual. | [optional] 
**AuthenticationMethods** | Pointer to **[]string** | The methods the individual authenticated with during the CLEAR flow.              Known values: - webauthn - email - sms_otp | [optional] 
**ActivatedAuthenticationMethods** | Pointer to **[]string** | The authentication methods the individual activated but did not use as a pre-existing credential during this session.              Known values: - webauthn - totp | [optional] 
**Checks** | Pointer to [**[]ClearProviderOutputCheck**](ClearProviderOutputCheck.md) | The checks CLEAR performed on the individual&#39;s data and each check&#39;s result. | [optional] 
**CheckMetadata** | Pointer to **[]string** | CLEAR error or metadata codes that influenced check results. | [optional] 
**CompletedAt** | Pointer to **NullableTime** | The time the CLEAR verification session completed, as a UTC timestamp. | [optional] 
**CreatedAt** | Pointer to **NullableTime** | The time the CLEAR verification session was created, as a UTC timestamp. | [optional] 
**Email** | Pointer to **NullableString** | The individual&#39;s verified email address from the session. | [optional] 
**ExpiresAt** | Pointer to **NullableTime** | The time the CLEAR verification session expires, as a UTC timestamp. | [optional] 
**FieldsToCollect** | Pointer to **[]string** | The fields CLEAR still required to complete the current verification step. | [optional] 
**Ip** | Pointer to **[]string** | The IP addresses used to access this CLEAR verification session. | [optional] 
**Phone** | Pointer to **NullableString** | The individual&#39;s verified phone number, normalized to international E.164 format. | [optional] 
**Status** | Pointer to **NullableString** | The CLEAR verification session status.              Known values: - success - fail - awaiting_user_input - deferred - processing_data - expired - awaiting_manual_review - manual_success - manual_fail - canceled | [optional] 
**UpdatedAt** | Pointer to **NullableTime** | The time the CLEAR verification session was most recently updated, as a UTC timestamp. | [optional] 
**UserAgent** | Pointer to **[]string** | The browser user agents that accessed this CLEAR verification session. | [optional] 
**UserCreated** | Pointer to **NullableBool** | Whether CLEAR created a new individual account during this session. | [optional] 
**UserId** | Pointer to **NullableString** | The CLEAR identifier for the individual. | [optional] 
**Traits** | Pointer to [**NullableClearProviderOutputTraits**](ClearProviderOutputTraits.md) | The individual traits CLEAR returned after collecting enough information to make a decision. | [optional] 
**IdvStatus** | Pointer to **NullableString** | The identity verification status.              Known documented value: - verified: CLEAR indicates the individual is verified. | [optional] 
**Sessions** | Pointer to [**[]ClearProviderOutputSessionInfo**](ClearProviderOutputSessionInfo.md) | Frontend session metadata CLEAR collected. | [optional] 

## Methods

### NewClearProviderOutput

`func NewClearProviderOutput() *ClearProviderOutput`

NewClearProviderOutput instantiates a new ClearProviderOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClearProviderOutputWithDefaults

`func NewClearProviderOutputWithDefaults() *ClearProviderOutput`

NewClearProviderOutputWithDefaults instantiates a new ClearProviderOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticated

`func (o *ClearProviderOutput) GetAuthenticated() bool`

GetAuthenticated returns the Authenticated field if non-nil, zero value otherwise.

### GetAuthenticatedOk

`func (o *ClearProviderOutput) GetAuthenticatedOk() (*bool, bool)`

GetAuthenticatedOk returns a tuple with the Authenticated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticated

`func (o *ClearProviderOutput) SetAuthenticated(v bool)`

SetAuthenticated sets Authenticated field to given value.

### HasAuthenticated

`func (o *ClearProviderOutput) HasAuthenticated() bool`

HasAuthenticated returns a boolean if a field has been set.

### SetAuthenticatedNil

`func (o *ClearProviderOutput) SetAuthenticatedNil(b bool)`

 SetAuthenticatedNil sets the value for Authenticated to be an explicit nil

### UnsetAuthenticated
`func (o *ClearProviderOutput) UnsetAuthenticated()`

UnsetAuthenticated ensures that no value is present for Authenticated, not even an explicit nil
### GetAuthenticationMethods

`func (o *ClearProviderOutput) GetAuthenticationMethods() []string`

GetAuthenticationMethods returns the AuthenticationMethods field if non-nil, zero value otherwise.

### GetAuthenticationMethodsOk

`func (o *ClearProviderOutput) GetAuthenticationMethodsOk() (*[]string, bool)`

GetAuthenticationMethodsOk returns a tuple with the AuthenticationMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationMethods

`func (o *ClearProviderOutput) SetAuthenticationMethods(v []string)`

SetAuthenticationMethods sets AuthenticationMethods field to given value.

### HasAuthenticationMethods

`func (o *ClearProviderOutput) HasAuthenticationMethods() bool`

HasAuthenticationMethods returns a boolean if a field has been set.

### SetAuthenticationMethodsNil

`func (o *ClearProviderOutput) SetAuthenticationMethodsNil(b bool)`

 SetAuthenticationMethodsNil sets the value for AuthenticationMethods to be an explicit nil

### UnsetAuthenticationMethods
`func (o *ClearProviderOutput) UnsetAuthenticationMethods()`

UnsetAuthenticationMethods ensures that no value is present for AuthenticationMethods, not even an explicit nil
### GetActivatedAuthenticationMethods

`func (o *ClearProviderOutput) GetActivatedAuthenticationMethods() []string`

GetActivatedAuthenticationMethods returns the ActivatedAuthenticationMethods field if non-nil, zero value otherwise.

### GetActivatedAuthenticationMethodsOk

`func (o *ClearProviderOutput) GetActivatedAuthenticationMethodsOk() (*[]string, bool)`

GetActivatedAuthenticationMethodsOk returns a tuple with the ActivatedAuthenticationMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivatedAuthenticationMethods

`func (o *ClearProviderOutput) SetActivatedAuthenticationMethods(v []string)`

SetActivatedAuthenticationMethods sets ActivatedAuthenticationMethods field to given value.

### HasActivatedAuthenticationMethods

`func (o *ClearProviderOutput) HasActivatedAuthenticationMethods() bool`

HasActivatedAuthenticationMethods returns a boolean if a field has been set.

### SetActivatedAuthenticationMethodsNil

`func (o *ClearProviderOutput) SetActivatedAuthenticationMethodsNil(b bool)`

 SetActivatedAuthenticationMethodsNil sets the value for ActivatedAuthenticationMethods to be an explicit nil

### UnsetActivatedAuthenticationMethods
`func (o *ClearProviderOutput) UnsetActivatedAuthenticationMethods()`

UnsetActivatedAuthenticationMethods ensures that no value is present for ActivatedAuthenticationMethods, not even an explicit nil
### GetChecks

`func (o *ClearProviderOutput) GetChecks() []ClearProviderOutputCheck`

GetChecks returns the Checks field if non-nil, zero value otherwise.

### GetChecksOk

`func (o *ClearProviderOutput) GetChecksOk() (*[]ClearProviderOutputCheck, bool)`

GetChecksOk returns a tuple with the Checks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecks

`func (o *ClearProviderOutput) SetChecks(v []ClearProviderOutputCheck)`

SetChecks sets Checks field to given value.

### HasChecks

`func (o *ClearProviderOutput) HasChecks() bool`

HasChecks returns a boolean if a field has been set.

### SetChecksNil

`func (o *ClearProviderOutput) SetChecksNil(b bool)`

 SetChecksNil sets the value for Checks to be an explicit nil

### UnsetChecks
`func (o *ClearProviderOutput) UnsetChecks()`

UnsetChecks ensures that no value is present for Checks, not even an explicit nil
### GetCheckMetadata

`func (o *ClearProviderOutput) GetCheckMetadata() []string`

GetCheckMetadata returns the CheckMetadata field if non-nil, zero value otherwise.

### GetCheckMetadataOk

`func (o *ClearProviderOutput) GetCheckMetadataOk() (*[]string, bool)`

GetCheckMetadataOk returns a tuple with the CheckMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckMetadata

`func (o *ClearProviderOutput) SetCheckMetadata(v []string)`

SetCheckMetadata sets CheckMetadata field to given value.

### HasCheckMetadata

`func (o *ClearProviderOutput) HasCheckMetadata() bool`

HasCheckMetadata returns a boolean if a field has been set.

### SetCheckMetadataNil

`func (o *ClearProviderOutput) SetCheckMetadataNil(b bool)`

 SetCheckMetadataNil sets the value for CheckMetadata to be an explicit nil

### UnsetCheckMetadata
`func (o *ClearProviderOutput) UnsetCheckMetadata()`

UnsetCheckMetadata ensures that no value is present for CheckMetadata, not even an explicit nil
### GetCompletedAt

`func (o *ClearProviderOutput) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *ClearProviderOutput) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *ClearProviderOutput) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *ClearProviderOutput) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### SetCompletedAtNil

`func (o *ClearProviderOutput) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *ClearProviderOutput) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
### GetCreatedAt

`func (o *ClearProviderOutput) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ClearProviderOutput) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ClearProviderOutput) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ClearProviderOutput) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *ClearProviderOutput) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *ClearProviderOutput) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetEmail

`func (o *ClearProviderOutput) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ClearProviderOutput) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ClearProviderOutput) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ClearProviderOutput) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *ClearProviderOutput) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *ClearProviderOutput) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetExpiresAt

`func (o *ClearProviderOutput) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *ClearProviderOutput) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *ClearProviderOutput) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *ClearProviderOutput) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *ClearProviderOutput) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *ClearProviderOutput) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetFieldsToCollect

`func (o *ClearProviderOutput) GetFieldsToCollect() []string`

GetFieldsToCollect returns the FieldsToCollect field if non-nil, zero value otherwise.

### GetFieldsToCollectOk

`func (o *ClearProviderOutput) GetFieldsToCollectOk() (*[]string, bool)`

GetFieldsToCollectOk returns a tuple with the FieldsToCollect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldsToCollect

`func (o *ClearProviderOutput) SetFieldsToCollect(v []string)`

SetFieldsToCollect sets FieldsToCollect field to given value.

### HasFieldsToCollect

`func (o *ClearProviderOutput) HasFieldsToCollect() bool`

HasFieldsToCollect returns a boolean if a field has been set.

### SetFieldsToCollectNil

`func (o *ClearProviderOutput) SetFieldsToCollectNil(b bool)`

 SetFieldsToCollectNil sets the value for FieldsToCollect to be an explicit nil

### UnsetFieldsToCollect
`func (o *ClearProviderOutput) UnsetFieldsToCollect()`

UnsetFieldsToCollect ensures that no value is present for FieldsToCollect, not even an explicit nil
### GetIp

`func (o *ClearProviderOutput) GetIp() []string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *ClearProviderOutput) GetIpOk() (*[]string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *ClearProviderOutput) SetIp(v []string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *ClearProviderOutput) HasIp() bool`

HasIp returns a boolean if a field has been set.

### SetIpNil

`func (o *ClearProviderOutput) SetIpNil(b bool)`

 SetIpNil sets the value for Ip to be an explicit nil

### UnsetIp
`func (o *ClearProviderOutput) UnsetIp()`

UnsetIp ensures that no value is present for Ip, not even an explicit nil
### GetPhone

`func (o *ClearProviderOutput) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *ClearProviderOutput) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *ClearProviderOutput) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *ClearProviderOutput) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### SetPhoneNil

`func (o *ClearProviderOutput) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *ClearProviderOutput) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil
### GetStatus

`func (o *ClearProviderOutput) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ClearProviderOutput) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ClearProviderOutput) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ClearProviderOutput) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ClearProviderOutput) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ClearProviderOutput) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetUpdatedAt

`func (o *ClearProviderOutput) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ClearProviderOutput) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ClearProviderOutput) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ClearProviderOutput) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *ClearProviderOutput) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *ClearProviderOutput) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetUserAgent

`func (o *ClearProviderOutput) GetUserAgent() []string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *ClearProviderOutput) GetUserAgentOk() (*[]string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *ClearProviderOutput) SetUserAgent(v []string)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *ClearProviderOutput) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.

### SetUserAgentNil

`func (o *ClearProviderOutput) SetUserAgentNil(b bool)`

 SetUserAgentNil sets the value for UserAgent to be an explicit nil

### UnsetUserAgent
`func (o *ClearProviderOutput) UnsetUserAgent()`

UnsetUserAgent ensures that no value is present for UserAgent, not even an explicit nil
### GetUserCreated

`func (o *ClearProviderOutput) GetUserCreated() bool`

GetUserCreated returns the UserCreated field if non-nil, zero value otherwise.

### GetUserCreatedOk

`func (o *ClearProviderOutput) GetUserCreatedOk() (*bool, bool)`

GetUserCreatedOk returns a tuple with the UserCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserCreated

`func (o *ClearProviderOutput) SetUserCreated(v bool)`

SetUserCreated sets UserCreated field to given value.

### HasUserCreated

`func (o *ClearProviderOutput) HasUserCreated() bool`

HasUserCreated returns a boolean if a field has been set.

### SetUserCreatedNil

`func (o *ClearProviderOutput) SetUserCreatedNil(b bool)`

 SetUserCreatedNil sets the value for UserCreated to be an explicit nil

### UnsetUserCreated
`func (o *ClearProviderOutput) UnsetUserCreated()`

UnsetUserCreated ensures that no value is present for UserCreated, not even an explicit nil
### GetUserId

`func (o *ClearProviderOutput) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ClearProviderOutput) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ClearProviderOutput) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ClearProviderOutput) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *ClearProviderOutput) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *ClearProviderOutput) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetTraits

`func (o *ClearProviderOutput) GetTraits() ClearProviderOutputTraits`

GetTraits returns the Traits field if non-nil, zero value otherwise.

### GetTraitsOk

`func (o *ClearProviderOutput) GetTraitsOk() (*ClearProviderOutputTraits, bool)`

GetTraitsOk returns a tuple with the Traits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraits

`func (o *ClearProviderOutput) SetTraits(v ClearProviderOutputTraits)`

SetTraits sets Traits field to given value.

### HasTraits

`func (o *ClearProviderOutput) HasTraits() bool`

HasTraits returns a boolean if a field has been set.

### SetTraitsNil

`func (o *ClearProviderOutput) SetTraitsNil(b bool)`

 SetTraitsNil sets the value for Traits to be an explicit nil

### UnsetTraits
`func (o *ClearProviderOutput) UnsetTraits()`

UnsetTraits ensures that no value is present for Traits, not even an explicit nil
### GetIdvStatus

`func (o *ClearProviderOutput) GetIdvStatus() string`

GetIdvStatus returns the IdvStatus field if non-nil, zero value otherwise.

### GetIdvStatusOk

`func (o *ClearProviderOutput) GetIdvStatusOk() (*string, bool)`

GetIdvStatusOk returns a tuple with the IdvStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdvStatus

`func (o *ClearProviderOutput) SetIdvStatus(v string)`

SetIdvStatus sets IdvStatus field to given value.

### HasIdvStatus

`func (o *ClearProviderOutput) HasIdvStatus() bool`

HasIdvStatus returns a boolean if a field has been set.

### SetIdvStatusNil

`func (o *ClearProviderOutput) SetIdvStatusNil(b bool)`

 SetIdvStatusNil sets the value for IdvStatus to be an explicit nil

### UnsetIdvStatus
`func (o *ClearProviderOutput) UnsetIdvStatus()`

UnsetIdvStatus ensures that no value is present for IdvStatus, not even an explicit nil
### GetSessions

`func (o *ClearProviderOutput) GetSessions() []ClearProviderOutputSessionInfo`

GetSessions returns the Sessions field if non-nil, zero value otherwise.

### GetSessionsOk

`func (o *ClearProviderOutput) GetSessionsOk() (*[]ClearProviderOutputSessionInfo, bool)`

GetSessionsOk returns a tuple with the Sessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessions

`func (o *ClearProviderOutput) SetSessions(v []ClearProviderOutputSessionInfo)`

SetSessions sets Sessions field to given value.

### HasSessions

`func (o *ClearProviderOutput) HasSessions() bool`

HasSessions returns a boolean if a field has been set.

### SetSessionsNil

`func (o *ClearProviderOutput) SetSessionsNil(b bool)`

 SetSessionsNil sets the value for Sessions to be an explicit nil

### UnsetSessions
`func (o *ClearProviderOutput) UnsetSessions()`

UnsetSessions ensures that no value is present for Sessions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


