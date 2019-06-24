# EdgelbV2Backend

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | This is name that frontends refer to. | [optional] 
**Protocol** | [**EdgelbV2Protocol**](EdgelbV2Protocol.md) |  | [optional] 
**RewriteHttp** | [**EdgelbV2RewriteHttp**](EdgelbV2RewriteHttp.md) |  | [optional] 
**Balance** | **string** | Load balancing strategy. e.g. roundrobin, leastconn, etc. | [optional] 
**CustomCheck** | [**EdgelbV2BackendCustomCheck**](EdgelbV2Backend_customCheck.md) |  | [optional] 
**MiscStrs** | **[]string** | Additional template lines inserted before servers | [optional] 
**Services** | [**[]EdgelbV2Service**](EdgelbV2Service.md) | Array of backend service selectors. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


