# IdentityData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OriginatingProviderId** | Pointer to **NullableString** | The ID of the provider from which this data originated (eg \&quot;yoti\&quot;, \&quot;clear\&quot;) | [optional] 
**OriginatingSubProviderId** | Pointer to **NullableString** | The sub-provider ID of the provider from which this data originated (eg \&quot;rabo\&quot;, \&quot;poste-italiane\&quot;)              This is applicable only to federated Identity Providers such as SPID and IDIN. | [optional] 
**Person** | Pointer to [**NullablePersonData**](PersonData.md) | Identity data of the individual who was verified | [optional] 
**Document** | Pointer to [**NullableDocumentData**](DocumentData.md) | Identity data of the document involved in verification, if relevant | [optional] 
**Match** | Pointer to [**NullableMatchData**](MatchData.md) | Match results for the data being matched against.              This applies to Providers which operate based on matching data / biometrics against a government database, returning match scores or results as opposed to the data itself. | [optional] 
**Attachments** | [**[]AttachmentInfo**](AttachmentInfo.md) | Information for each attachment included with this set of identity data.              Use the Attachments API to fetch an attachment by its ID for a given Session. | 
**ProviderOutput** | Pointer to [**NullableProviderOutput**](ProviderOutput.md) | Provider-specific output data that doesn&#39;t fit the standard identity data schema.              The structure of this object varies by provider. | [optional] 

## Methods

### NewIdentityData

`func NewIdentityData(attachments []AttachmentInfo, ) *IdentityData`

NewIdentityData instantiates a new IdentityData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityDataWithDefaults

`func NewIdentityDataWithDefaults() *IdentityData`

NewIdentityDataWithDefaults instantiates a new IdentityData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOriginatingProviderId

`func (o *IdentityData) GetOriginatingProviderId() string`

GetOriginatingProviderId returns the OriginatingProviderId field if non-nil, zero value otherwise.

### GetOriginatingProviderIdOk

`func (o *IdentityData) GetOriginatingProviderIdOk() (*string, bool)`

GetOriginatingProviderIdOk returns a tuple with the OriginatingProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginatingProviderId

`func (o *IdentityData) SetOriginatingProviderId(v string)`

SetOriginatingProviderId sets OriginatingProviderId field to given value.

### HasOriginatingProviderId

`func (o *IdentityData) HasOriginatingProviderId() bool`

HasOriginatingProviderId returns a boolean if a field has been set.

### SetOriginatingProviderIdNil

`func (o *IdentityData) SetOriginatingProviderIdNil(b bool)`

 SetOriginatingProviderIdNil sets the value for OriginatingProviderId to be an explicit nil

### UnsetOriginatingProviderId
`func (o *IdentityData) UnsetOriginatingProviderId()`

UnsetOriginatingProviderId ensures that no value is present for OriginatingProviderId, not even an explicit nil
### GetOriginatingSubProviderId

`func (o *IdentityData) GetOriginatingSubProviderId() string`

GetOriginatingSubProviderId returns the OriginatingSubProviderId field if non-nil, zero value otherwise.

### GetOriginatingSubProviderIdOk

`func (o *IdentityData) GetOriginatingSubProviderIdOk() (*string, bool)`

GetOriginatingSubProviderIdOk returns a tuple with the OriginatingSubProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginatingSubProviderId

`func (o *IdentityData) SetOriginatingSubProviderId(v string)`

SetOriginatingSubProviderId sets OriginatingSubProviderId field to given value.

### HasOriginatingSubProviderId

`func (o *IdentityData) HasOriginatingSubProviderId() bool`

HasOriginatingSubProviderId returns a boolean if a field has been set.

### SetOriginatingSubProviderIdNil

`func (o *IdentityData) SetOriginatingSubProviderIdNil(b bool)`

 SetOriginatingSubProviderIdNil sets the value for OriginatingSubProviderId to be an explicit nil

### UnsetOriginatingSubProviderId
`func (o *IdentityData) UnsetOriginatingSubProviderId()`

UnsetOriginatingSubProviderId ensures that no value is present for OriginatingSubProviderId, not even an explicit nil
### GetPerson

`func (o *IdentityData) GetPerson() PersonData`

GetPerson returns the Person field if non-nil, zero value otherwise.

### GetPersonOk

`func (o *IdentityData) GetPersonOk() (*PersonData, bool)`

GetPersonOk returns a tuple with the Person field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerson

`func (o *IdentityData) SetPerson(v PersonData)`

SetPerson sets Person field to given value.

### HasPerson

`func (o *IdentityData) HasPerson() bool`

HasPerson returns a boolean if a field has been set.

### SetPersonNil

`func (o *IdentityData) SetPersonNil(b bool)`

 SetPersonNil sets the value for Person to be an explicit nil

### UnsetPerson
`func (o *IdentityData) UnsetPerson()`

UnsetPerson ensures that no value is present for Person, not even an explicit nil
### GetDocument

`func (o *IdentityData) GetDocument() DocumentData`

GetDocument returns the Document field if non-nil, zero value otherwise.

### GetDocumentOk

`func (o *IdentityData) GetDocumentOk() (*DocumentData, bool)`

GetDocumentOk returns a tuple with the Document field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocument

`func (o *IdentityData) SetDocument(v DocumentData)`

SetDocument sets Document field to given value.

### HasDocument

`func (o *IdentityData) HasDocument() bool`

HasDocument returns a boolean if a field has been set.

### SetDocumentNil

`func (o *IdentityData) SetDocumentNil(b bool)`

 SetDocumentNil sets the value for Document to be an explicit nil

### UnsetDocument
`func (o *IdentityData) UnsetDocument()`

UnsetDocument ensures that no value is present for Document, not even an explicit nil
### GetMatch

`func (o *IdentityData) GetMatch() MatchData`

GetMatch returns the Match field if non-nil, zero value otherwise.

### GetMatchOk

`func (o *IdentityData) GetMatchOk() (*MatchData, bool)`

GetMatchOk returns a tuple with the Match field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatch

`func (o *IdentityData) SetMatch(v MatchData)`

SetMatch sets Match field to given value.

### HasMatch

`func (o *IdentityData) HasMatch() bool`

HasMatch returns a boolean if a field has been set.

### SetMatchNil

`func (o *IdentityData) SetMatchNil(b bool)`

 SetMatchNil sets the value for Match to be an explicit nil

### UnsetMatch
`func (o *IdentityData) UnsetMatch()`

UnsetMatch ensures that no value is present for Match, not even an explicit nil
### GetAttachments

`func (o *IdentityData) GetAttachments() []AttachmentInfo`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *IdentityData) GetAttachmentsOk() (*[]AttachmentInfo, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *IdentityData) SetAttachments(v []AttachmentInfo)`

SetAttachments sets Attachments field to given value.


### GetProviderOutput

`func (o *IdentityData) GetProviderOutput() ProviderOutput`

GetProviderOutput returns the ProviderOutput field if non-nil, zero value otherwise.

### GetProviderOutputOk

`func (o *IdentityData) GetProviderOutputOk() (*ProviderOutput, bool)`

GetProviderOutputOk returns a tuple with the ProviderOutput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderOutput

`func (o *IdentityData) SetProviderOutput(v ProviderOutput)`

SetProviderOutput sets ProviderOutput field to given value.

### HasProviderOutput

`func (o *IdentityData) HasProviderOutput() bool`

HasProviderOutput returns a boolean if a field has been set.

### SetProviderOutputNil

`func (o *IdentityData) SetProviderOutputNil(b bool)`

 SetProviderOutputNil sets the value for ProviderOutput to be an explicit nil

### UnsetProviderOutput
`func (o *IdentityData) UnsetProviderOutput()`

UnsetProviderOutput ensures that no value is present for ProviderOutput, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


