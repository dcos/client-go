# V2Frontend

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Defaults to frontend_{{bindAddress}}_{{bindPort}}. | [optional] 
**BindAddress** | **string** | Only use characters that are allowed in the frontend name. Known invalid frontend name characters include \&quot;*\&quot;, \&quot;[\&quot;, and \&quot;]\&quot;. | [optional] [default to 0.0.0.0]
**BindPort** | **int32** | The port (e.g. 80 for HTTP or 443 for HTTPS) that this frontend will bind to. | [optional] [default to -1]
**BindModifier** | **string** | Additional text to put in the bind field | [optional] 
**Certificates** | **[]string** |  | [optional] 
**RedirectToHttps** | [**V2FrontendRedirectToHttps**](V2Frontend_redirectToHttps.md) |  | [optional] 
**MiscStrs** | **[]string** | Additional template lines inserted before use_backend | [optional] 
**Protocol** | [**V2Protocol**](V2Protocol.md) |  | [optional] 
**LinkBackend** | [**V2FrontendLinkBackend**](V2Frontend_linkBackend.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


