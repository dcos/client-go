# EdgelbV2Frontend

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Defaults to frontend_{{bindAddress}}_{{bindPort}}. | [optional] 
**BindAddress** | **string** | Only use characters that are allowed in the frontend name. Known invalid frontend name characters include \&quot;*\&quot;, \&quot;[\&quot;, and \&quot;]\&quot;. | [optional] 
**BindPort** | **int32** | The port (e.g. 80 for HTTP or 443 for HTTPS) that this frontend will bind to. | [optional] 
**BindModifier** | **string** | Additional text to put in the bind field | [optional] 
**Certificates** | **[]string** |  | [optional] 
**RedirectToHttps** | Pointer to [**EdgelbV2FrontendRedirectToHttps**](EdgelbV2Frontend_redirectToHttps.md) |  | [optional] 
**MiscStrs** | **[]string** | Additional template lines inserted before use_backend | [optional] 
**Protocol** | [**EdgelbV2Protocol**](EdgelbV2Protocol.md) |  | [optional] 
**LinkBackend** | [**EdgelbV2FrontendLinkBackend**](EdgelbV2Frontend_linkBackend.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


