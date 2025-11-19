# ProviderInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KenyaBiometric2** | Pointer to [**NullableKenyaNidBiometric2Input**](KenyaNidBiometric2Input.md) | Input for the &#x60;kenya-nid-match-biometric-2&#x60; provider | [optional] 
**IndonesiaNik** | Pointer to [**NullableIndonesiaNikInput**](IndonesiaNikInput.md) | Input for the &#x60;indonesia-nik-lookup&#x60; provider | [optional] 
**IndonesiaDukcapilMatch** | Pointer to [**NullableIndonesiaDukcapilMatchInput**](IndonesiaDukcapilMatchInput.md) | Input for the &#x60;indonesia-dukcapil-match&#x60; provider | [optional] 
**MexicoCurp** | Pointer to [**NullableMexicoCurpInput**](MexicoCurpInput.md) | Input for the &#x60;mexico-curp-lookup&#x60; provider | [optional] 
**SouthAfricaNid** | Pointer to [**NullableSouthAfricaNidInput**](SouthAfricaNidInput.md) | Input for the &#x60;south-africa-nid-lookup&#x60; provider | [optional] 
**KenyaNid** | Pointer to [**NullableKenyaNidInput**](KenyaNidInput.md) | Input for the &#x60;kenya-nid-lookup&#x60; provider | [optional] 
**KenyaNidMatch2** | Pointer to [**NullableKenyaNidMatch2Input**](KenyaNidMatch2Input.md) | Input for the &#x60;kenya-nid-match-2&#x60; provider | [optional] 
**KenyaNidLookup2** | Pointer to [**NullableKenyaNidLookup2Input**](KenyaNidLookup2Input.md) | Input for the &#x60;kenya-nid-lookup-2&#x60; provider | [optional] 
**SouthAfricaNidLookup2** | Pointer to [**NullableSouthAfricaNidLookup2Input**](SouthAfricaNidLookup2Input.md) | Input for the &#x60;south-africa-nid-lookup-2&#x60; provider | [optional] 
**NigeriaNin** | Pointer to [**NullableNigeriaNinInput**](NigeriaNinInput.md) | Input for the &#x60;nigeria-nin-lookup&#x60; provider | [optional] 
**Aadhaar** | Pointer to [**NullableAadhaarInput**](AadhaarInput.md) | Input for the &#x60;india-digilocker-aadhaar-match&#x60; provider | [optional] 
**BangladeshNationalId** | Pointer to [**NullableBangladeshNidInput**](BangladeshNidInput.md) | Input for the &#x60;bangladesh-nid&#x60; provider | [optional] 
**BrazilCpfCheck** | Pointer to [**NullableBrazilCpfCheckInput**](BrazilCpfCheckInput.md) | Input for the &#x60;brazil-cpf-lookup&#x60; provider | [optional] 
**BrazilDigitalCnh** | Pointer to [**NullableBrazilDigitalCnhInput**](BrazilDigitalCnhInput.md) | Input for the &#x60;brazil-digital-cnh&#x60; provider | [optional] 
**PhilippineMatch** | Pointer to [**NullablePhilippineMatchInput**](PhilippineMatchInput.md) | Input for the &#x60;philippines-philsys-match&#x60; provider | [optional] 
**PhilippineQR** | Pointer to [**NullablePhilippineQRInput**](PhilippineQRInput.md) | Input for the &#x60;philippines-digital-national-id-qr&#x60; and &#x60;philippines-physical-national-id-qr&#x60; providers | [optional] 
**SmartId** | Pointer to [**NullableSmartIdInput**](SmartIdInput.md) | Input for the &#x60;smart-id&#x60; provider | [optional] 
**MobileId** | Pointer to [**NullableMobileIdInput**](MobileIdInput.md) | Input for the &#x60;mobile-id&#x60; provider | [optional] 
**Idin** | Pointer to [**NullableIdinInput**](IdinInput.md) | Input for the &#x60;netherlands-idin&#x60; provider | [optional] 
**Spid** | Pointer to [**NullableSpidInput**](SpidInput.md) | Input for the &#x60;italy-spid&#x60; provider | [optional] 
**GoogleWallet** | Pointer to [**NullableGoogleWalletInput**](GoogleWalletInput.md) | Input for the &#x60;google-wallet&#x60; provider | [optional] 
**AppleWallet** | Pointer to [**NullableAppleWalletInput**](AppleWalletInput.md) | Input for the &#x60;apple-wallet&#x60; provider | [optional] 
**TrinsicTestDatabaseLookup** | Pointer to [**NullableTrinsicTestDatabaseLookupInput**](TrinsicTestDatabaseLookupInput.md) | *TEST MODE ONLY.*              Input for the &#x60;trinsic-test-database-lookup&#x60; provider | [optional] 
**TrinsicTestSubProviders** | Pointer to [**NullableTrinsicTestSubProvidersInput**](TrinsicTestSubProvidersInput.md) | *TEST MODE ONLY.*              Input for the &#x60;trinsic-test-sub-providers&#x60; provider | [optional] 

## Methods

### NewProviderInput

`func NewProviderInput() *ProviderInput`

NewProviderInput instantiates a new ProviderInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderInputWithDefaults

`func NewProviderInputWithDefaults() *ProviderInput`

NewProviderInputWithDefaults instantiates a new ProviderInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKenyaBiometric2

`func (o *ProviderInput) GetKenyaBiometric2() KenyaNidBiometric2Input`

GetKenyaBiometric2 returns the KenyaBiometric2 field if non-nil, zero value otherwise.

### GetKenyaBiometric2Ok

`func (o *ProviderInput) GetKenyaBiometric2Ok() (*KenyaNidBiometric2Input, bool)`

GetKenyaBiometric2Ok returns a tuple with the KenyaBiometric2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKenyaBiometric2

`func (o *ProviderInput) SetKenyaBiometric2(v KenyaNidBiometric2Input)`

SetKenyaBiometric2 sets KenyaBiometric2 field to given value.

### HasKenyaBiometric2

`func (o *ProviderInput) HasKenyaBiometric2() bool`

HasKenyaBiometric2 returns a boolean if a field has been set.

### SetKenyaBiometric2Nil

`func (o *ProviderInput) SetKenyaBiometric2Nil(b bool)`

 SetKenyaBiometric2Nil sets the value for KenyaBiometric2 to be an explicit nil

### UnsetKenyaBiometric2
`func (o *ProviderInput) UnsetKenyaBiometric2()`

UnsetKenyaBiometric2 ensures that no value is present for KenyaBiometric2, not even an explicit nil
### GetIndonesiaNik

`func (o *ProviderInput) GetIndonesiaNik() IndonesiaNikInput`

GetIndonesiaNik returns the IndonesiaNik field if non-nil, zero value otherwise.

### GetIndonesiaNikOk

`func (o *ProviderInput) GetIndonesiaNikOk() (*IndonesiaNikInput, bool)`

GetIndonesiaNikOk returns a tuple with the IndonesiaNik field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndonesiaNik

`func (o *ProviderInput) SetIndonesiaNik(v IndonesiaNikInput)`

SetIndonesiaNik sets IndonesiaNik field to given value.

### HasIndonesiaNik

`func (o *ProviderInput) HasIndonesiaNik() bool`

HasIndonesiaNik returns a boolean if a field has been set.

### SetIndonesiaNikNil

`func (o *ProviderInput) SetIndonesiaNikNil(b bool)`

 SetIndonesiaNikNil sets the value for IndonesiaNik to be an explicit nil

### UnsetIndonesiaNik
`func (o *ProviderInput) UnsetIndonesiaNik()`

UnsetIndonesiaNik ensures that no value is present for IndonesiaNik, not even an explicit nil
### GetIndonesiaDukcapilMatch

`func (o *ProviderInput) GetIndonesiaDukcapilMatch() IndonesiaDukcapilMatchInput`

GetIndonesiaDukcapilMatch returns the IndonesiaDukcapilMatch field if non-nil, zero value otherwise.

### GetIndonesiaDukcapilMatchOk

`func (o *ProviderInput) GetIndonesiaDukcapilMatchOk() (*IndonesiaDukcapilMatchInput, bool)`

GetIndonesiaDukcapilMatchOk returns a tuple with the IndonesiaDukcapilMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndonesiaDukcapilMatch

`func (o *ProviderInput) SetIndonesiaDukcapilMatch(v IndonesiaDukcapilMatchInput)`

SetIndonesiaDukcapilMatch sets IndonesiaDukcapilMatch field to given value.

### HasIndonesiaDukcapilMatch

`func (o *ProviderInput) HasIndonesiaDukcapilMatch() bool`

HasIndonesiaDukcapilMatch returns a boolean if a field has been set.

### SetIndonesiaDukcapilMatchNil

`func (o *ProviderInput) SetIndonesiaDukcapilMatchNil(b bool)`

 SetIndonesiaDukcapilMatchNil sets the value for IndonesiaDukcapilMatch to be an explicit nil

### UnsetIndonesiaDukcapilMatch
`func (o *ProviderInput) UnsetIndonesiaDukcapilMatch()`

UnsetIndonesiaDukcapilMatch ensures that no value is present for IndonesiaDukcapilMatch, not even an explicit nil
### GetMexicoCurp

`func (o *ProviderInput) GetMexicoCurp() MexicoCurpInput`

GetMexicoCurp returns the MexicoCurp field if non-nil, zero value otherwise.

### GetMexicoCurpOk

`func (o *ProviderInput) GetMexicoCurpOk() (*MexicoCurpInput, bool)`

GetMexicoCurpOk returns a tuple with the MexicoCurp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMexicoCurp

`func (o *ProviderInput) SetMexicoCurp(v MexicoCurpInput)`

SetMexicoCurp sets MexicoCurp field to given value.

### HasMexicoCurp

`func (o *ProviderInput) HasMexicoCurp() bool`

HasMexicoCurp returns a boolean if a field has been set.

### SetMexicoCurpNil

`func (o *ProviderInput) SetMexicoCurpNil(b bool)`

 SetMexicoCurpNil sets the value for MexicoCurp to be an explicit nil

### UnsetMexicoCurp
`func (o *ProviderInput) UnsetMexicoCurp()`

UnsetMexicoCurp ensures that no value is present for MexicoCurp, not even an explicit nil
### GetSouthAfricaNid

`func (o *ProviderInput) GetSouthAfricaNid() SouthAfricaNidInput`

GetSouthAfricaNid returns the SouthAfricaNid field if non-nil, zero value otherwise.

### GetSouthAfricaNidOk

`func (o *ProviderInput) GetSouthAfricaNidOk() (*SouthAfricaNidInput, bool)`

GetSouthAfricaNidOk returns a tuple with the SouthAfricaNid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSouthAfricaNid

`func (o *ProviderInput) SetSouthAfricaNid(v SouthAfricaNidInput)`

SetSouthAfricaNid sets SouthAfricaNid field to given value.

### HasSouthAfricaNid

`func (o *ProviderInput) HasSouthAfricaNid() bool`

HasSouthAfricaNid returns a boolean if a field has been set.

### SetSouthAfricaNidNil

`func (o *ProviderInput) SetSouthAfricaNidNil(b bool)`

 SetSouthAfricaNidNil sets the value for SouthAfricaNid to be an explicit nil

### UnsetSouthAfricaNid
`func (o *ProviderInput) UnsetSouthAfricaNid()`

UnsetSouthAfricaNid ensures that no value is present for SouthAfricaNid, not even an explicit nil
### GetKenyaNid

`func (o *ProviderInput) GetKenyaNid() KenyaNidInput`

GetKenyaNid returns the KenyaNid field if non-nil, zero value otherwise.

### GetKenyaNidOk

`func (o *ProviderInput) GetKenyaNidOk() (*KenyaNidInput, bool)`

GetKenyaNidOk returns a tuple with the KenyaNid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKenyaNid

`func (o *ProviderInput) SetKenyaNid(v KenyaNidInput)`

SetKenyaNid sets KenyaNid field to given value.

### HasKenyaNid

`func (o *ProviderInput) HasKenyaNid() bool`

HasKenyaNid returns a boolean if a field has been set.

### SetKenyaNidNil

`func (o *ProviderInput) SetKenyaNidNil(b bool)`

 SetKenyaNidNil sets the value for KenyaNid to be an explicit nil

### UnsetKenyaNid
`func (o *ProviderInput) UnsetKenyaNid()`

UnsetKenyaNid ensures that no value is present for KenyaNid, not even an explicit nil
### GetKenyaNidMatch2

`func (o *ProviderInput) GetKenyaNidMatch2() KenyaNidMatch2Input`

GetKenyaNidMatch2 returns the KenyaNidMatch2 field if non-nil, zero value otherwise.

### GetKenyaNidMatch2Ok

`func (o *ProviderInput) GetKenyaNidMatch2Ok() (*KenyaNidMatch2Input, bool)`

GetKenyaNidMatch2Ok returns a tuple with the KenyaNidMatch2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKenyaNidMatch2

`func (o *ProviderInput) SetKenyaNidMatch2(v KenyaNidMatch2Input)`

SetKenyaNidMatch2 sets KenyaNidMatch2 field to given value.

### HasKenyaNidMatch2

`func (o *ProviderInput) HasKenyaNidMatch2() bool`

HasKenyaNidMatch2 returns a boolean if a field has been set.

### SetKenyaNidMatch2Nil

`func (o *ProviderInput) SetKenyaNidMatch2Nil(b bool)`

 SetKenyaNidMatch2Nil sets the value for KenyaNidMatch2 to be an explicit nil

### UnsetKenyaNidMatch2
`func (o *ProviderInput) UnsetKenyaNidMatch2()`

UnsetKenyaNidMatch2 ensures that no value is present for KenyaNidMatch2, not even an explicit nil
### GetKenyaNidLookup2

`func (o *ProviderInput) GetKenyaNidLookup2() KenyaNidLookup2Input`

GetKenyaNidLookup2 returns the KenyaNidLookup2 field if non-nil, zero value otherwise.

### GetKenyaNidLookup2Ok

`func (o *ProviderInput) GetKenyaNidLookup2Ok() (*KenyaNidLookup2Input, bool)`

GetKenyaNidLookup2Ok returns a tuple with the KenyaNidLookup2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKenyaNidLookup2

`func (o *ProviderInput) SetKenyaNidLookup2(v KenyaNidLookup2Input)`

SetKenyaNidLookup2 sets KenyaNidLookup2 field to given value.

### HasKenyaNidLookup2

`func (o *ProviderInput) HasKenyaNidLookup2() bool`

HasKenyaNidLookup2 returns a boolean if a field has been set.

### SetKenyaNidLookup2Nil

`func (o *ProviderInput) SetKenyaNidLookup2Nil(b bool)`

 SetKenyaNidLookup2Nil sets the value for KenyaNidLookup2 to be an explicit nil

### UnsetKenyaNidLookup2
`func (o *ProviderInput) UnsetKenyaNidLookup2()`

UnsetKenyaNidLookup2 ensures that no value is present for KenyaNidLookup2, not even an explicit nil
### GetSouthAfricaNidLookup2

`func (o *ProviderInput) GetSouthAfricaNidLookup2() SouthAfricaNidLookup2Input`

GetSouthAfricaNidLookup2 returns the SouthAfricaNidLookup2 field if non-nil, zero value otherwise.

### GetSouthAfricaNidLookup2Ok

`func (o *ProviderInput) GetSouthAfricaNidLookup2Ok() (*SouthAfricaNidLookup2Input, bool)`

GetSouthAfricaNidLookup2Ok returns a tuple with the SouthAfricaNidLookup2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSouthAfricaNidLookup2

`func (o *ProviderInput) SetSouthAfricaNidLookup2(v SouthAfricaNidLookup2Input)`

SetSouthAfricaNidLookup2 sets SouthAfricaNidLookup2 field to given value.

### HasSouthAfricaNidLookup2

`func (o *ProviderInput) HasSouthAfricaNidLookup2() bool`

HasSouthAfricaNidLookup2 returns a boolean if a field has been set.

### SetSouthAfricaNidLookup2Nil

`func (o *ProviderInput) SetSouthAfricaNidLookup2Nil(b bool)`

 SetSouthAfricaNidLookup2Nil sets the value for SouthAfricaNidLookup2 to be an explicit nil

### UnsetSouthAfricaNidLookup2
`func (o *ProviderInput) UnsetSouthAfricaNidLookup2()`

UnsetSouthAfricaNidLookup2 ensures that no value is present for SouthAfricaNidLookup2, not even an explicit nil
### GetNigeriaNin

`func (o *ProviderInput) GetNigeriaNin() NigeriaNinInput`

GetNigeriaNin returns the NigeriaNin field if non-nil, zero value otherwise.

### GetNigeriaNinOk

`func (o *ProviderInput) GetNigeriaNinOk() (*NigeriaNinInput, bool)`

GetNigeriaNinOk returns a tuple with the NigeriaNin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNigeriaNin

`func (o *ProviderInput) SetNigeriaNin(v NigeriaNinInput)`

SetNigeriaNin sets NigeriaNin field to given value.

### HasNigeriaNin

`func (o *ProviderInput) HasNigeriaNin() bool`

HasNigeriaNin returns a boolean if a field has been set.

### SetNigeriaNinNil

`func (o *ProviderInput) SetNigeriaNinNil(b bool)`

 SetNigeriaNinNil sets the value for NigeriaNin to be an explicit nil

### UnsetNigeriaNin
`func (o *ProviderInput) UnsetNigeriaNin()`

UnsetNigeriaNin ensures that no value is present for NigeriaNin, not even an explicit nil
### GetAadhaar

`func (o *ProviderInput) GetAadhaar() AadhaarInput`

GetAadhaar returns the Aadhaar field if non-nil, zero value otherwise.

### GetAadhaarOk

`func (o *ProviderInput) GetAadhaarOk() (*AadhaarInput, bool)`

GetAadhaarOk returns a tuple with the Aadhaar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAadhaar

`func (o *ProviderInput) SetAadhaar(v AadhaarInput)`

SetAadhaar sets Aadhaar field to given value.

### HasAadhaar

`func (o *ProviderInput) HasAadhaar() bool`

HasAadhaar returns a boolean if a field has been set.

### SetAadhaarNil

`func (o *ProviderInput) SetAadhaarNil(b bool)`

 SetAadhaarNil sets the value for Aadhaar to be an explicit nil

### UnsetAadhaar
`func (o *ProviderInput) UnsetAadhaar()`

UnsetAadhaar ensures that no value is present for Aadhaar, not even an explicit nil
### GetBangladeshNationalId

`func (o *ProviderInput) GetBangladeshNationalId() BangladeshNidInput`

GetBangladeshNationalId returns the BangladeshNationalId field if non-nil, zero value otherwise.

### GetBangladeshNationalIdOk

`func (o *ProviderInput) GetBangladeshNationalIdOk() (*BangladeshNidInput, bool)`

GetBangladeshNationalIdOk returns a tuple with the BangladeshNationalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBangladeshNationalId

`func (o *ProviderInput) SetBangladeshNationalId(v BangladeshNidInput)`

SetBangladeshNationalId sets BangladeshNationalId field to given value.

### HasBangladeshNationalId

`func (o *ProviderInput) HasBangladeshNationalId() bool`

HasBangladeshNationalId returns a boolean if a field has been set.

### SetBangladeshNationalIdNil

`func (o *ProviderInput) SetBangladeshNationalIdNil(b bool)`

 SetBangladeshNationalIdNil sets the value for BangladeshNationalId to be an explicit nil

### UnsetBangladeshNationalId
`func (o *ProviderInput) UnsetBangladeshNationalId()`

UnsetBangladeshNationalId ensures that no value is present for BangladeshNationalId, not even an explicit nil
### GetBrazilCpfCheck

`func (o *ProviderInput) GetBrazilCpfCheck() BrazilCpfCheckInput`

GetBrazilCpfCheck returns the BrazilCpfCheck field if non-nil, zero value otherwise.

### GetBrazilCpfCheckOk

`func (o *ProviderInput) GetBrazilCpfCheckOk() (*BrazilCpfCheckInput, bool)`

GetBrazilCpfCheckOk returns a tuple with the BrazilCpfCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrazilCpfCheck

`func (o *ProviderInput) SetBrazilCpfCheck(v BrazilCpfCheckInput)`

SetBrazilCpfCheck sets BrazilCpfCheck field to given value.

### HasBrazilCpfCheck

`func (o *ProviderInput) HasBrazilCpfCheck() bool`

HasBrazilCpfCheck returns a boolean if a field has been set.

### SetBrazilCpfCheckNil

`func (o *ProviderInput) SetBrazilCpfCheckNil(b bool)`

 SetBrazilCpfCheckNil sets the value for BrazilCpfCheck to be an explicit nil

### UnsetBrazilCpfCheck
`func (o *ProviderInput) UnsetBrazilCpfCheck()`

UnsetBrazilCpfCheck ensures that no value is present for BrazilCpfCheck, not even an explicit nil
### GetBrazilDigitalCnh

`func (o *ProviderInput) GetBrazilDigitalCnh() BrazilDigitalCnhInput`

GetBrazilDigitalCnh returns the BrazilDigitalCnh field if non-nil, zero value otherwise.

### GetBrazilDigitalCnhOk

`func (o *ProviderInput) GetBrazilDigitalCnhOk() (*BrazilDigitalCnhInput, bool)`

GetBrazilDigitalCnhOk returns a tuple with the BrazilDigitalCnh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrazilDigitalCnh

`func (o *ProviderInput) SetBrazilDigitalCnh(v BrazilDigitalCnhInput)`

SetBrazilDigitalCnh sets BrazilDigitalCnh field to given value.

### HasBrazilDigitalCnh

`func (o *ProviderInput) HasBrazilDigitalCnh() bool`

HasBrazilDigitalCnh returns a boolean if a field has been set.

### SetBrazilDigitalCnhNil

`func (o *ProviderInput) SetBrazilDigitalCnhNil(b bool)`

 SetBrazilDigitalCnhNil sets the value for BrazilDigitalCnh to be an explicit nil

### UnsetBrazilDigitalCnh
`func (o *ProviderInput) UnsetBrazilDigitalCnh()`

UnsetBrazilDigitalCnh ensures that no value is present for BrazilDigitalCnh, not even an explicit nil
### GetPhilippineMatch

`func (o *ProviderInput) GetPhilippineMatch() PhilippineMatchInput`

GetPhilippineMatch returns the PhilippineMatch field if non-nil, zero value otherwise.

### GetPhilippineMatchOk

`func (o *ProviderInput) GetPhilippineMatchOk() (*PhilippineMatchInput, bool)`

GetPhilippineMatchOk returns a tuple with the PhilippineMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhilippineMatch

`func (o *ProviderInput) SetPhilippineMatch(v PhilippineMatchInput)`

SetPhilippineMatch sets PhilippineMatch field to given value.

### HasPhilippineMatch

`func (o *ProviderInput) HasPhilippineMatch() bool`

HasPhilippineMatch returns a boolean if a field has been set.

### SetPhilippineMatchNil

`func (o *ProviderInput) SetPhilippineMatchNil(b bool)`

 SetPhilippineMatchNil sets the value for PhilippineMatch to be an explicit nil

### UnsetPhilippineMatch
`func (o *ProviderInput) UnsetPhilippineMatch()`

UnsetPhilippineMatch ensures that no value is present for PhilippineMatch, not even an explicit nil
### GetPhilippineQR

`func (o *ProviderInput) GetPhilippineQR() PhilippineQRInput`

GetPhilippineQR returns the PhilippineQR field if non-nil, zero value otherwise.

### GetPhilippineQROk

`func (o *ProviderInput) GetPhilippineQROk() (*PhilippineQRInput, bool)`

GetPhilippineQROk returns a tuple with the PhilippineQR field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhilippineQR

`func (o *ProviderInput) SetPhilippineQR(v PhilippineQRInput)`

SetPhilippineQR sets PhilippineQR field to given value.

### HasPhilippineQR

`func (o *ProviderInput) HasPhilippineQR() bool`

HasPhilippineQR returns a boolean if a field has been set.

### SetPhilippineQRNil

`func (o *ProviderInput) SetPhilippineQRNil(b bool)`

 SetPhilippineQRNil sets the value for PhilippineQR to be an explicit nil

### UnsetPhilippineQR
`func (o *ProviderInput) UnsetPhilippineQR()`

UnsetPhilippineQR ensures that no value is present for PhilippineQR, not even an explicit nil
### GetSmartId

`func (o *ProviderInput) GetSmartId() SmartIdInput`

GetSmartId returns the SmartId field if non-nil, zero value otherwise.

### GetSmartIdOk

`func (o *ProviderInput) GetSmartIdOk() (*SmartIdInput, bool)`

GetSmartIdOk returns a tuple with the SmartId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartId

`func (o *ProviderInput) SetSmartId(v SmartIdInput)`

SetSmartId sets SmartId field to given value.

### HasSmartId

`func (o *ProviderInput) HasSmartId() bool`

HasSmartId returns a boolean if a field has been set.

### SetSmartIdNil

`func (o *ProviderInput) SetSmartIdNil(b bool)`

 SetSmartIdNil sets the value for SmartId to be an explicit nil

### UnsetSmartId
`func (o *ProviderInput) UnsetSmartId()`

UnsetSmartId ensures that no value is present for SmartId, not even an explicit nil
### GetMobileId

`func (o *ProviderInput) GetMobileId() MobileIdInput`

GetMobileId returns the MobileId field if non-nil, zero value otherwise.

### GetMobileIdOk

`func (o *ProviderInput) GetMobileIdOk() (*MobileIdInput, bool)`

GetMobileIdOk returns a tuple with the MobileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileId

`func (o *ProviderInput) SetMobileId(v MobileIdInput)`

SetMobileId sets MobileId field to given value.

### HasMobileId

`func (o *ProviderInput) HasMobileId() bool`

HasMobileId returns a boolean if a field has been set.

### SetMobileIdNil

`func (o *ProviderInput) SetMobileIdNil(b bool)`

 SetMobileIdNil sets the value for MobileId to be an explicit nil

### UnsetMobileId
`func (o *ProviderInput) UnsetMobileId()`

UnsetMobileId ensures that no value is present for MobileId, not even an explicit nil
### GetIdin

`func (o *ProviderInput) GetIdin() IdinInput`

GetIdin returns the Idin field if non-nil, zero value otherwise.

### GetIdinOk

`func (o *ProviderInput) GetIdinOk() (*IdinInput, bool)`

GetIdinOk returns a tuple with the Idin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdin

`func (o *ProviderInput) SetIdin(v IdinInput)`

SetIdin sets Idin field to given value.

### HasIdin

`func (o *ProviderInput) HasIdin() bool`

HasIdin returns a boolean if a field has been set.

### SetIdinNil

`func (o *ProviderInput) SetIdinNil(b bool)`

 SetIdinNil sets the value for Idin to be an explicit nil

### UnsetIdin
`func (o *ProviderInput) UnsetIdin()`

UnsetIdin ensures that no value is present for Idin, not even an explicit nil
### GetSpid

`func (o *ProviderInput) GetSpid() SpidInput`

GetSpid returns the Spid field if non-nil, zero value otherwise.

### GetSpidOk

`func (o *ProviderInput) GetSpidOk() (*SpidInput, bool)`

GetSpidOk returns a tuple with the Spid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpid

`func (o *ProviderInput) SetSpid(v SpidInput)`

SetSpid sets Spid field to given value.

### HasSpid

`func (o *ProviderInput) HasSpid() bool`

HasSpid returns a boolean if a field has been set.

### SetSpidNil

`func (o *ProviderInput) SetSpidNil(b bool)`

 SetSpidNil sets the value for Spid to be an explicit nil

### UnsetSpid
`func (o *ProviderInput) UnsetSpid()`

UnsetSpid ensures that no value is present for Spid, not even an explicit nil
### GetGoogleWallet

`func (o *ProviderInput) GetGoogleWallet() GoogleWalletInput`

GetGoogleWallet returns the GoogleWallet field if non-nil, zero value otherwise.

### GetGoogleWalletOk

`func (o *ProviderInput) GetGoogleWalletOk() (*GoogleWalletInput, bool)`

GetGoogleWalletOk returns a tuple with the GoogleWallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleWallet

`func (o *ProviderInput) SetGoogleWallet(v GoogleWalletInput)`

SetGoogleWallet sets GoogleWallet field to given value.

### HasGoogleWallet

`func (o *ProviderInput) HasGoogleWallet() bool`

HasGoogleWallet returns a boolean if a field has been set.

### SetGoogleWalletNil

`func (o *ProviderInput) SetGoogleWalletNil(b bool)`

 SetGoogleWalletNil sets the value for GoogleWallet to be an explicit nil

### UnsetGoogleWallet
`func (o *ProviderInput) UnsetGoogleWallet()`

UnsetGoogleWallet ensures that no value is present for GoogleWallet, not even an explicit nil
### GetAppleWallet

`func (o *ProviderInput) GetAppleWallet() AppleWalletInput`

GetAppleWallet returns the AppleWallet field if non-nil, zero value otherwise.

### GetAppleWalletOk

`func (o *ProviderInput) GetAppleWalletOk() (*AppleWalletInput, bool)`

GetAppleWalletOk returns a tuple with the AppleWallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppleWallet

`func (o *ProviderInput) SetAppleWallet(v AppleWalletInput)`

SetAppleWallet sets AppleWallet field to given value.

### HasAppleWallet

`func (o *ProviderInput) HasAppleWallet() bool`

HasAppleWallet returns a boolean if a field has been set.

### SetAppleWalletNil

`func (o *ProviderInput) SetAppleWalletNil(b bool)`

 SetAppleWalletNil sets the value for AppleWallet to be an explicit nil

### UnsetAppleWallet
`func (o *ProviderInput) UnsetAppleWallet()`

UnsetAppleWallet ensures that no value is present for AppleWallet, not even an explicit nil
### GetTrinsicTestDatabaseLookup

`func (o *ProviderInput) GetTrinsicTestDatabaseLookup() TrinsicTestDatabaseLookupInput`

GetTrinsicTestDatabaseLookup returns the TrinsicTestDatabaseLookup field if non-nil, zero value otherwise.

### GetTrinsicTestDatabaseLookupOk

`func (o *ProviderInput) GetTrinsicTestDatabaseLookupOk() (*TrinsicTestDatabaseLookupInput, bool)`

GetTrinsicTestDatabaseLookupOk returns a tuple with the TrinsicTestDatabaseLookup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrinsicTestDatabaseLookup

`func (o *ProviderInput) SetTrinsicTestDatabaseLookup(v TrinsicTestDatabaseLookupInput)`

SetTrinsicTestDatabaseLookup sets TrinsicTestDatabaseLookup field to given value.

### HasTrinsicTestDatabaseLookup

`func (o *ProviderInput) HasTrinsicTestDatabaseLookup() bool`

HasTrinsicTestDatabaseLookup returns a boolean if a field has been set.

### SetTrinsicTestDatabaseLookupNil

`func (o *ProviderInput) SetTrinsicTestDatabaseLookupNil(b bool)`

 SetTrinsicTestDatabaseLookupNil sets the value for TrinsicTestDatabaseLookup to be an explicit nil

### UnsetTrinsicTestDatabaseLookup
`func (o *ProviderInput) UnsetTrinsicTestDatabaseLookup()`

UnsetTrinsicTestDatabaseLookup ensures that no value is present for TrinsicTestDatabaseLookup, not even an explicit nil
### GetTrinsicTestSubProviders

`func (o *ProviderInput) GetTrinsicTestSubProviders() TrinsicTestSubProvidersInput`

GetTrinsicTestSubProviders returns the TrinsicTestSubProviders field if non-nil, zero value otherwise.

### GetTrinsicTestSubProvidersOk

`func (o *ProviderInput) GetTrinsicTestSubProvidersOk() (*TrinsicTestSubProvidersInput, bool)`

GetTrinsicTestSubProvidersOk returns a tuple with the TrinsicTestSubProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrinsicTestSubProviders

`func (o *ProviderInput) SetTrinsicTestSubProviders(v TrinsicTestSubProvidersInput)`

SetTrinsicTestSubProviders sets TrinsicTestSubProviders field to given value.

### HasTrinsicTestSubProviders

`func (o *ProviderInput) HasTrinsicTestSubProviders() bool`

HasTrinsicTestSubProviders returns a boolean if a field has been set.

### SetTrinsicTestSubProvidersNil

`func (o *ProviderInput) SetTrinsicTestSubProvidersNil(b bool)`

 SetTrinsicTestSubProvidersNil sets the value for TrinsicTestSubProviders to be an explicit nil

### UnsetTrinsicTestSubProviders
`func (o *ProviderInput) UnsetTrinsicTestSubProviders()`

UnsetTrinsicTestSubProviders ensures that no value is present for TrinsicTestSubProviders, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


