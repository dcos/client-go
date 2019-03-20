# V2Backend

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | This is name that frontends refer to. | [optional] 
**Protocol** | [**V2Protocol**](V2Protocol.md) |  | [optional] 
**RewriteHttp** | [**V2RewriteHttp**](V2RewriteHttp.md) |  | [optional] 
**Balance** | **string** | Load balancing strategy. e.g. roundrobin, leastconn, etc. | [optional] [default to roundrobin]
**CustomCheck** | [**V2BackendCustomCheck**](V2Backend_customCheck.md) |  | [optional] 
**MiscStrs** | **[]string** | Additional template lines inserted before servers | [optional] 
**Services** | [**[]V2Service**](V2Service.md) | Array of backend service selectors. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


