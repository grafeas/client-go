# V1beta1Signature

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Signature** | **string** | The content of the signature, an opaque bytestring. The payload that this signature verifies MUST be unambiguously provided with the Signature during verification. A wrapper message might provide the payload explicitly. Alternatively, a message might have a canonical serialization that can always be unambiguously computed to derive the payload. | [optional] [default to null]
**PublicKeyId** | **string** | The identifier for the public key that verifies this signature.   * The &#x60;public_key_id&#x60; is required.   * The &#x60;public_key_id&#x60; MUST be an RFC3986 conformant URI.   * When possible, the &#x60;public_key_id&#x60; SHOULD be an immutable reference,     such as a cryptographic digest.  Examples of valid &#x60;public_key_id&#x60;s:  OpenPGP V4 public key fingerprint:   * \&quot;openpgp4fpr:74FAF3B861BDA0870C7B6DEF607E48D2A663AEEA\&quot; See https://www.iana.org/assignments/uri-schemes/prov/openpgp4fpr for more details on this scheme.  RFC6920 digest-named SubjectPublicKeyInfo (digest of the DER serialization):   * \&quot;ni:///sha-256;cD9o9Cq6LG3jD0iKXqEi_vdjJGecm_iXkbqVoScViaU\&quot;   * \&quot;nih:///sha-256;703f68f42aba2c6de30f488a5ea122fef76324679c9bf89791ba95a1271589a5\&quot; | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


