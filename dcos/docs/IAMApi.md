# \IAMApi

All URIs are relative to *http://your-dcos-cluster.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConfigureOIDCProvider**](IAMApi.md#ConfigureOIDCProvider) | **Put** /acs/api/v1/auth/oidc/providers/{provider-id} | Configure a new OIDC provider.
[**ConfigureSAMLProvider**](IAMApi.md#ConfigureSAMLProvider) | **Put** /acs/api/v1/auth/saml/providers/{provider-id} | Configure a new SAML provider.
[**CreateGroup**](IAMApi.md#CreateGroup) | **Put** /acs/api/v1/groups/{gid} | Create a group.
[**CreateGroupUser**](IAMApi.md#CreateGroupUser) | **Put** /acs/api/v1/groups/{gid}/users/{uid} | Add account to group.
[**CreateLDAPConfiguration**](IAMApi.md#CreateLDAPConfiguration) | **Put** /acs/api/v1/ldap/config | Set new LDAP configuration.
[**CreateResourceACL**](IAMApi.md#CreateResourceACL) | **Put** /acs/api/v1/acls/{rid} | Create ACL for a certain resource.
[**CreateUser**](IAMApi.md#CreateUser) | **Put** /acs/api/v1/users/{uid} | Create user account.
[**DeleteGroup**](IAMApi.md#DeleteGroup) | **Delete** /acs/api/v1/groups/{gid} | Delete group.
[**DeleteGroupUser**](IAMApi.md#DeleteGroupUser) | **Delete** /acs/api/v1/groups/{gid}/users/{uid} | Delete user account from group.
[**DeleteLDAPConfiguration**](IAMApi.md#DeleteLDAPConfiguration) | **Delete** /acs/api/v1/ldap/config | Delete current LDAP configuration.
[**DeleteOIDCProvider**](IAMApi.md#DeleteOIDCProvider) | **Delete** /acs/api/v1/auth/oidc/providers/{provider-id} | Delete provider.
[**DeleteResourceACL**](IAMApi.md#DeleteResourceACL) | **Delete** /acs/api/v1/acls/{rid} | Delete ACL for a certain resource.
[**DeleteSAMLProvider**](IAMApi.md#DeleteSAMLProvider) | **Delete** /acs/api/v1/auth/saml/providers/{provider-id} | Delete provider.
[**DeleteUser**](IAMApi.md#DeleteUser) | **Delete** /acs/api/v1/users/{uid} | Delete account.
[**ForbidResourceUserAction**](IAMApi.md#ForbidResourceUserAction) | **Delete** /acs/api/v1/acls/{rid}/users/{uid}/{action} | Forbid single action for given account and resource.
[**ForbidResourceUserActions**](IAMApi.md#ForbidResourceUserActions) | **Delete** /acs/api/v1/acls/{rid}/users/{uid} | Forbid all actions of given account to given resource.
[**ForbitResourceGroupAction**](IAMApi.md#ForbitResourceGroupAction) | **Delete** /acs/api/v1/acls/{rid}/groups/{gid}/{action} | Forbid single action for given resource and group.
[**ForbitResourceGroupActions**](IAMApi.md#ForbitResourceGroupActions) | **Delete** /acs/api/v1/acls/{rid}/groups/{gid} | Forbid all actions of given group to given resource.
[**GetACLs**](IAMApi.md#GetACLs) | **Get** /acs/api/v1/acls | Retrieve all ACL objects.
[**GetGroup**](IAMApi.md#GetGroup) | **Get** /acs/api/v1/groups/{gid} | Get single group object.
[**GetGroupPermissions**](IAMApi.md#GetGroupPermissions) | **Get** /acs/api/v1/groups/{gid}/permissions | Retrieve group permissions.
[**GetGroupUsers**](IAMApi.md#GetGroupUsers) | **Get** /acs/api/v1/groups/{gid}/users | Retrieve members of a group.
[**GetGroups**](IAMApi.md#GetGroups) | **Get** /acs/api/v1/groups | Retrieve all group objects.
[**GetJWKS**](IAMApi.md#GetJWKS) | **Get** /acs/api/v1/auth/jwks | Get the IAM&#39;s JSON Web Key Set (JWKS, according to RFCs 7517/7518).
[**GetLDAPConfiguration**](IAMApi.md#GetLDAPConfiguration) | **Get** /acs/api/v1/ldap/config | Retrieve current LDAP configuration.
[**GetOIDCProvider**](IAMApi.md#GetOIDCProvider) | **Get** /acs/api/v1/auth/oidc/providers/{provider-id} | Get configuration for a specific provider.
[**GetOIDCProviders**](IAMApi.md#GetOIDCProviders) | **Get** /acs/api/v1/auth/oidc/providers | Get an overview for the configured OIDC providers.
[**GetResourceACLs**](IAMApi.md#GetResourceACLs) | **Get** /acs/api/v1/acls/{rid} | Retrieve ACL for a certain resource.
[**GetResourceGroupAction**](IAMApi.md#GetResourceGroupAction) | **Get** /acs/api/v1/acls/{rid}/groups/{gid}/{action} | Query whether action is allowed or not.
[**GetResourceGroupActions**](IAMApi.md#GetResourceGroupActions) | **Get** /acs/api/v1/acls/{rid}/groups/{gid} | Get allowed actions for given resource and group.
[**GetResourcePermissions**](IAMApi.md#GetResourcePermissions) | **Get** /acs/api/v1/acls/{rid}/permissions | Retrieve all permissions for resource.
[**GetResourceUserAction**](IAMApi.md#GetResourceUserAction) | **Get** /acs/api/v1/acls/{rid}/users/{uid}/{action} | Query whether action is allowed or not.
[**GetResourceUserActions**](IAMApi.md#GetResourceUserActions) | **Get** /acs/api/v1/acls/{rid}/users/{uid} | Get allowed actions for given resource and user.
[**GetSAMLProvider**](IAMApi.md#GetSAMLProvider) | **Get** /acs/api/v1/auth/saml/providers/{provider-id} | Get configuration for a specific SAML provider.
[**GetSAMLProviderACSCallbackURL**](IAMApi.md#GetSAMLProviderACSCallbackURL) | **Get** /acs/api/v1/auth/saml/providers/{provider-id}/acs-callback-url | Get the authentication callback URL for this SP.
[**GetSAMLProviderSPMetadata**](IAMApi.md#GetSAMLProviderSPMetadata) | **Get** /acs/api/v1/auth/saml/providers/{provider-id}/sp-metadata | Get SP metadata (XML).
[**GetSAMLProviders**](IAMApi.md#GetSAMLProviders) | **Get** /acs/api/v1/auth/saml/providers | Get an overview for the configured SAML 2.0 providers.
[**GetUser**](IAMApi.md#GetUser) | **Get** /acs/api/v1/users/{uid} | Get single user object.
[**GetUserGroups**](IAMApi.md#GetUserGroups) | **Get** /acs/api/v1/users/{uid}/groups | Retrieve groups the user is member of.
[**GetUserPermissions**](IAMApi.md#GetUserPermissions) | **Get** /acs/api/v1/users/{uid}/permissions | Retrieve permissions an account has.
[**GetUsers**](IAMApi.md#GetUsers) | **Get** /acs/api/v1/users | Retrieve all regular user accounts or service user accounts.
[**ImportLDAPGroup**](IAMApi.md#ImportLDAPGroup) | **Post** /acs/api/v1/ldap/importgroup | Import an LDAP group.
[**ImportLDAPUser**](IAMApi.md#ImportLDAPUser) | **Post** /acs/api/v1/ldap/importuser | Import an LDAP user.
[**Login**](IAMApi.md#Login) | **Post** /acs/api/v1/auth/login | Log in (obtain a DC/OS authentication token).
[**LoginWithProvider**](IAMApi.md#LoginWithProvider) | **Get** /acs/api/v1/auth/login | Log in using an external identity provider.
[**OpenIDConnectCallbackURL**](IAMApi.md#OpenIDConnectCallbackURL) | **Get** /acs/api/v1/auth/oidc/callback | OpenID Connect callback URL.
[**PermitResourceGroupAction**](IAMApi.md#PermitResourceGroupAction) | **Put** /acs/api/v1/acls/{rid}/groups/{gid}/{action} | Permit single action for given resource and group.
[**PermitResourceUserAction**](IAMApi.md#PermitResourceUserAction) | **Put** /acs/api/v1/acls/{rid}/users/{uid}/{action} | Permit single action for given account and resource.
[**SAMLProviderACSCallbackEndpoint**](IAMApi.md#SAMLProviderACSCallbackEndpoint) | **Post** /acs/api/v1/auth/saml/providers/{provider-id}/acs-callback | The SP ACS callback endpoint.
[**TestLDAPBackendConnection**](IAMApi.md#TestLDAPBackendConnection) | **Post** /acs/api/v1/ldap/config/test | Test connection to the LDAP back-end.
[**UpdateGroup**](IAMApi.md#UpdateGroup) | **Patch** /acs/api/v1/groups/{gid} | Update group.
[**UpdateOIDCProvider**](IAMApi.md#UpdateOIDCProvider) | **Patch** /acs/api/v1/auth/oidc/providers/{provider-id} | Update OIDC provider config.
[**UpdateResourceACL**](IAMApi.md#UpdateResourceACL) | **Patch** /acs/api/v1/acls/{rid} | Update ACL for a certain resource.
[**UpdateSAMLProvider**](IAMApi.md#UpdateSAMLProvider) | **Patch** /acs/api/v1/auth/saml/providers/{provider-id} | Update SAML provider config.
[**UpdateUser**](IAMApi.md#UpdateUser) | **Patch** /acs/api/v1/users/{uid} | Update user account.


# **ConfigureOIDCProvider**
> ConfigureOIDCProvider(ctx, providerId, iamoidcProviderConfig)
Configure a new OIDC provider.

Set up OIDC provider with the ID as specified in the URL, and with the config as specified via JSON in the request body.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **providerId** | **string**| The ID of the provider to create. | 
  **iamoidcProviderConfig** | [**IamoidcProviderConfig**](IamoidcProviderConfig.md)| Provider config JSON object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConfigureSAMLProvider**
> ConfigureSAMLProvider(ctx, providerId, iamsamlProviderConfig)
Configure a new SAML provider.

Set up a SAML provider with the ID as specified in the URL, and with the config as given by the JSON document in the request body.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **providerId** | **string**| The ID of the provider to create. | 
  **iamsamlProviderConfig** | [**IamsamlProviderConfig**](IamsamlProviderConfig.md)| Provider config JSON object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateGroup**
> CreateGroup(ctx, gid, iamGroupCreate)
Create a group.

Create a group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gid** | **string**| The ID of the group. | 
  **iamGroupCreate** | [**IamGroupCreate**](IamGroupCreate.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateGroupUser**
> CreateGroupUser(ctx, gid, uid)
Add account to group.

Add account to group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gid** | **string**| The ID of the group to add the user account to. | 
  **uid** | **string**| The ID of the account to add. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateLDAPConfiguration**
> CreateLDAPConfiguration(ctx, iamldapConfiguration)
Set new LDAP configuration.

Set new directory (LDAP) back-end configuration. Replace current configuration, if existing.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **iamldapConfiguration** | [**IamldapConfiguration**](IamldapConfiguration.md)| JSON object containing the LDAP configuration details. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateResourceACL**
> CreateResourceACL(ctx, rid, iamaclCreate)
Create ACL for a certain resource.

Create new ACL for resource with ID `rid` (description in body, no permissions by default).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **rid** | **string**| The ID of the resource for which the ACL should be created. | 
  **iamaclCreate** | [**IamaclCreate**](IamaclCreate.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateUser**
> CreateUser(ctx, uid, iamUserCreate)
Create user account.

Create user (uid in url, details incl. credentials in body).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uid** | **string**| The ID of the user account to create. | 
  **iamUserCreate** | [**IamUserCreate**](IamUserCreate.md)| Password/description. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteGroup**
> DeleteGroup(ctx, gid)
Delete group.

Delete group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gid** | **string**| The ID of the group to delete. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteGroupUser**
> DeleteGroupUser(ctx, gid, uid)
Delete user account from group.

Delete user account from group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gid** | **string**| The ID of the group to delete from. | 
  **uid** | **string**| The ID of the user account. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteLDAPConfiguration**
> DeleteLDAPConfiguration(ctx, )
Delete current LDAP configuration.

Delete current directory (LDAP) back-end configuration. This deactivates the LDAP authentication.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOIDCProvider**
> DeleteOIDCProvider(ctx, providerId)
Delete provider.

Delete provider (disables authentication with that provider).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **providerId** | **string**| The ID of the OIDC provider to delete. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteResourceACL**
> DeleteResourceACL(ctx, rid)
Delete ACL for a certain resource.

Delete ACL of resource with ID `rid`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **rid** | **string**| The ID of resource for which the ACL should be deleted. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSAMLProvider**
> DeleteSAMLProvider(ctx, providerId)
Delete provider.

Delete provider (disables authentication with that provider).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **providerId** | **string**| The ID of the SAML provider to delete. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteUser**
> DeleteUser(ctx, uid)
Delete account.

Delete account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uid** | **string**| The ID of the user account to delete. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ForbidResourceUserAction**
> ForbidResourceUserAction(ctx, rid, uid, action)
Forbid single action for given account and resource.

Forbid single action for given account and resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **rid** | **string**| resource ID. | 
  **uid** | **string**| account ID. | 
  **action** | **string**| action name | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ForbidResourceUserActions**
> ForbidResourceUserActions(ctx, rid, uid)
Forbid all actions of given account to given resource.

Forbid all actions of given account to given resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **rid** | **string**| resource ID. | 
  **uid** | **string**| account ID. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ForbitResourceGroupAction**
> ForbitResourceGroupAction(ctx, rid, gid, action)
Forbid single action for given resource and group.

Forbid single action for given resource and group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **rid** | **string**| resource ID. | 
  **gid** | **string**| group ID. | 
  **action** | **string**| action name | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ForbitResourceGroupActions**
> ForbitResourceGroupActions(ctx, rid, gid)
Forbid all actions of given group to given resource.

Forbid all actions of given group to given resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **rid** | **string**| resource ID. | 
  **gid** | **string**| group ID. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetACLs**
> IamAcls GetACLs(ctx, )
Retrieve all ACL objects.

Get array of `ACL` objects.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**IamAcls**](IAMAcls.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGroup**
> IamGroup GetGroup(ctx, gid)
Get single group object.

Get specific `Group` object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gid** | **string**| The ID of the group to retrieve. | 

### Return type

[**IamGroup**](IAMGroup.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGroupPermissions**
> IamGroupPermissions GetGroupPermissions(ctx, gid)
Retrieve group permissions.

Retrieve permissions of this group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gid** | **string**| The group ID. | 

### Return type

[**IamGroupPermissions**](IAMGroupPermissions.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGroupUsers**
> IamGroupUsers GetGroupUsers(ctx, gid, optional)
Retrieve members of a group.

Retrieve users that are member of this group. Allows to query service accounts, defaults to list only user accounts.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gid** | **string**| The group ID. | 
 **optional** | ***GetGroupUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetGroupUsersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **optional.String**| If set to &#x60;service&#x60;, list only service accounts. If unset, default to only listing user accounts members of a group. | 

### Return type

[**IamGroupUsers**](IAMGroupUsers.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGroups**
> IamGroups GetGroups(ctx, )
Retrieve all group objects.

Retrieve array of `Group` objects.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**IamGroups**](IAMGroups.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetJWKS**
> GetJWKS(ctx, )
Get the IAM's JSON Web Key Set (JWKS, according to RFCs 7517/7518).

This endpoint provides the IAM's JSON Web Key Set (JWKS), exposing public key details required for the process of DC/OS authentication token verification: the public key material can be used for verifying authentication tokens signed by the IAM's private key. The DC/OS authentication token is a JSON Web Token (JWT) of type RS256, and is required to have the two claims `exp` and `uid`. For interpretation of the data provided by the JWKS endpoint see https://tools.ietf.org/html/rfc7517 and https://tools.ietf.org/html/rfc7518.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLDAPConfiguration**
> IamldapConfiguration GetLDAPConfiguration(ctx, )
Retrieve current LDAP configuration.

Retrieve current directory (LDAP) back-end configuration.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**IamldapConfiguration**](IAMLDAPConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOIDCProvider**
> IamoidcProviderConfig GetOIDCProvider(ctx, providerId)
Get configuration for a specific provider.

Get configuration for a specific provider.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **providerId** | **string**| The ID of the OIDC provider to retrieve the config for. | 

### Return type

[**IamoidcProviderConfig**](IAMOIDCProviderConfig.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOIDCProviders**
> GetOIDCProviders(ctx, )
Get an overview for the configured OIDC providers.

Get an overview for the configured OIDC providers. The response contains a JSON object, with each key being an OIDC provider ID, and each value being the corresponding provider description string. This endpoint does not expose sensitive provider configuration details.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetResourceACLs**
> Iamacl GetResourceACLs(ctx, rid)
Retrieve ACL for a certain resource.

Retrieve single `ACL` object, for a specific resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **rid** | **string**| The ID of the resource to retrieve the ACL for. | 

### Return type

[**Iamacl**](IAMACL.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetResourceGroupAction**
> IamActionAllowed GetResourceGroupAction(ctx, rid, gid, action)
Query whether action is allowed or not.

Query whether action is allowed or not.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **rid** | **string**| resource ID | 
  **gid** | **string**| group ID | 
  **action** | **string**| action name | 

### Return type

[**IamActionAllowed**](IAMActionAllowed.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetResourceGroupActions**
> IamActions GetResourceGroupActions(ctx, rid, gid)
Get allowed actions for given resource and group.

Get allowed actions for given resource and group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **rid** | **string**| resource ID | 
  **gid** | **string**| group ID | 

### Return type

[**IamActions**](IAMActions.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetResourcePermissions**
> IamaclPermissions GetResourcePermissions(ctx, rid)
Retrieve all permissions for resource.

Retrieve all permissions that are set for a specific resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **rid** | **string**| resource ID | 

### Return type

[**IamaclPermissions**](IAMACLPermissions.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetResourceUserAction**
> IamActionAllowed GetResourceUserAction(ctx, rid, uid, action)
Query whether action is allowed or not.

Query whether action is allowed or not.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **rid** | **string**| resource ID | 
  **uid** | **string**| account ID | 
  **action** | **string**| action name | 

### Return type

[**IamActionAllowed**](IAMActionAllowed.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetResourceUserActions**
> InlineResponse200 GetResourceUserActions(ctx, rid, uid)
Get allowed actions for given resource and user.

Get allowed actions for given resource and user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **rid** | **string**| resource ID | 
  **uid** | **string**| account ID | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSAMLProvider**
> IamsamlProviderConfig GetSAMLProvider(ctx, providerId)
Get configuration for a specific SAML provider.

Get configuration for a specific SAML provider.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **providerId** | **string**| The ID of the SAML provider to retrieve the config for. | 

### Return type

[**IamsamlProviderConfig**](IAMSAMLProviderConfig.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSAMLProviderACSCallbackURL**
> IamsamlacsCallbackUrlObject GetSAMLProviderACSCallbackURL(ctx, providerId)
Get the authentication callback URL for this SP.

The IAM acts as SAML service provider (SP). A SAML identity provider (IdP) usually requires to be configured with the Assertion Consumer Service (ACS) callback URL of the SP (which is where the IdP makes the end-user submit the authentication response).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **providerId** | **string**| The ID of the SAML provider to retrieve the ACS callback URL for. | 

### Return type

[**IamsamlacsCallbackUrlObject**](IAMSAMLACSCallbackUrlObject.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSAMLProviderSPMetadata**
> string GetSAMLProviderSPMetadata(ctx, providerId)
Get SP metadata (XML).

The IAM acts as SAML service provider (SP). This endpoint provides the SP metadata as an XML document. Certain identity providers (IdPs) may want to directly consume this document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **providerId** | **string**| The ID of the SAML provider to retrieve the metadata for. | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/samlmetadata+xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSAMLProviders**
> GetSAMLProviders(ctx, )
Get an overview for the configured SAML 2.0 providers.

Get an overview for the configured SAML 2.0 providers. The response contains a JSON object, with each key being a SAML provider ID, and each value being the corresponding provider description string. This endpoint does not expose sensitive provider configuration details.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUser**
> IamUser GetUser(ctx, uid)
Get single user object.

Get specific `User` object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uid** | **string**| The ID of the user object to retrieve. | 

### Return type

[**IamUser**](IAMUser.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserGroups**
> IamUserGroups GetUserGroups(ctx, uid)
Retrieve groups the user is member of.

Retrieve groups the user is member of.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uid** | **string**| The ID of the user. | 

### Return type

[**IamUserGroups**](IAMUserGroups.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserPermissions**
> IamUserPermissions GetUserPermissions(ctx, uid)
Retrieve permissions an account has.

Retrieve the permissions for this account with direct permissions distinguished from those that are obtained through group membership.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uid** | **string**| The id of the user. | 

### Return type

[**IamUserPermissions**](IAMUserPermissions.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUsers**
> IamUsers GetUsers(ctx, optional)
Retrieve all regular user accounts or service user accounts.

Retrieve `User` objects. By default the list consists of regular user accounts, only. Alternatively, service user accounts may be requested instead.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetUsersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **optional.String**| If set to &#x60;service&#x60;, list only service user accounts. If unset, default to only listing regular user accounts. | 

### Return type

[**IamUsers**](IAMUsers.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImportLDAPGroup**
> ImportLDAPGroup(ctx, iamldapImportGroupObject)
Import an LDAP group.

Attempt to import a group of users from the configured directory (LDAP) back-end. See docs/ldap.md for details on group import.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **iamldapImportGroupObject** | [**IamldapImportGroupObject**](IamldapImportGroupObject.md)| A JSON object specifying the name of the group to be imported. The meaning of the name depends on the group search settings. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImportLDAPUser**
> ImportLDAPUser(ctx, iamldapImportUserObject)
Import an LDAP user.

Attempt to import a user from the configured directory (LDAP) back-end.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **iamldapImportUserObject** | [**IamldapImportUserObject**](IamldapImportUserObject.md)| A JSON object specifying the username (read: \&quot;login\&quot; or \&quot;user ID\&quot;) of the user that should be imported. That string is equivalent to the &#x60;uid&#x60; the user is supposed to log in with after successful import. The exact meaning of this string depends on the configured LDAP authentication method. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Login**
> IamAuthToken Login(ctx, iamLoginObject)
Log in (obtain a DC/OS authentication token).

Exchange user credentials (regular user account: uid and password; service user account: uid and service login token) for a DC/OS authentication token. The resulting DC/OS authentication token is an RFC 7519 JSON Web Token (JWT) of type RS256. It has a limited lifetime which depends on the IAM configuration (only, i.e. the lifetime cannot be chosen as part of the login HTTP request). The DC/OS authentication token can be verified out-of-band using a standards-compliant RS256 JWT verification procedure based on the long-lived public key material presented by the IAM's /auth/jwks endpoint, and by requiring the two claims `exp` and `uid` to be present.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **iamLoginObject** | [**IamLoginObject**](IamLoginObject.md)| uid &amp; password or uid &amp; service login token. | 

### Return type

[**IamAuthToken**](IAMAuthToken.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LoginWithProvider**
> LoginWithProvider(ctx, optional)
Log in using an external identity provider.

Log in using an external identity provider (via e.g. OpenID Connect), as specified via query parameter. This request initiates a single sign-on flow.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LoginWithProviderOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LoginWithProviderOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oidcProvider** | **optional.String**| OIDC provider ID | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OpenIDConnectCallbackURL**
> OpenIDConnectCallbackURL(ctx, )
OpenID Connect callback URL.

After successfully logging in to an OpenID Connect identity provider, the end-user is being redirected back to the IAM via this callback URL. API consumers are not required to explicitly interact with this endpoint. This URL usually needs to be handed over to an OpenID Connect provider (often called \"redirect\" or \"callback\" URL).

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermitResourceGroupAction**
> PermitResourceGroupAction(ctx, rid, gid, action)
Permit single action for given resource and group.

Permit single action for given resource and group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **rid** | **string**| resource ID. | 
  **gid** | **string**| group ID. | 
  **action** | **string**| action name | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermitResourceUserAction**
> PermitResourceUserAction(ctx, rid, uid, action)
Permit single action for given account and resource.

Permit single action for given account and resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **rid** | **string**| resource ID. | 
  **uid** | **string**| account ID. | 
  **action** | **string**| action name | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SAMLProviderACSCallbackEndpoint**
> SAMLProviderACSCallbackEndpoint(ctx, providerId)
The SP ACS callback endpoint.

The IAM acts as SAML service provider (SP). As part of the authentication flow, a SAML identity provider (IdP) makes the end-user submit an authentication response to this endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **providerId** | **string**| The ID of the provider the authentication response is meant for. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TestLDAPBackendConnection**
> IamldapTestResultObject TestLDAPBackendConnection(ctx, iamldapTestCredentials)
Test connection to the LDAP back-end.

Perform basic feature tests. Verify that the current directory (LDAP) configuration parameters allow for a successful connection to the directory back-end. For instance, this endpoint simulates the procedure for authentication via LDAP, but provides more useful feedback upon failure than the actual login endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **iamldapTestCredentials** | [**IamldapTestCredentials**](IamldapTestCredentials.md)| JSON object containing &#x60;uid&#x60; and password of an LDAP user. For the most expressive test result, choose credentials different from the lookup credentials. The &#x60;uid&#x60; is the string the user is supposed to log in with after successful LDAP back-end configuration. | 

### Return type

[**IamldapTestResultObject**](IAMLDAPTestResultObject.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateGroup**
> UpdateGroup(ctx, gid, iamGroupUpdate)
Update group.

Update existing group (description).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gid** | **string**| The ID of the group to modify. | 
  **iamGroupUpdate** | [**IamGroupUpdate**](IamGroupUpdate.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOIDCProvider**
> UpdateOIDCProvider(ctx, providerId, iamoidcProviderConfig)
Update OIDC provider config.

Update config for existing OIDC provider.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **providerId** | **string**| The ID of the provider to modify. | 
  **iamoidcProviderConfig** | [**IamoidcProviderConfig**](IamoidcProviderConfig.md)| Provider config JSON object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateResourceACL**
> UpdateResourceACL(ctx, rid, iamaclUpdate)
Update ACL for a certain resource.

Update ACL for resource with ID `rid`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **rid** | **string**| The ID of the resource for which the ACL should be created. | 
  **iamaclUpdate** | [**IamaclUpdate**](IamaclUpdate.md)| New ACL. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSAMLProvider**
> UpdateSAMLProvider(ctx, providerId, iamsamlProviderConfig)
Update SAML provider config.

Update config for existing SAML provider.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **providerId** | **string**| The ID of the provider to modify. | 
  **iamsamlProviderConfig** | [**IamsamlProviderConfig**](IamsamlProviderConfig.md)| Provider config JSON object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateUser**
> UpdateUser(ctx, uid, iamUserUpdate)
Update user account.

Update existing user account (meta data and/or password).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uid** | **string**| The ID of the user account to modify. | 
  **iamUserUpdate** | [**IamUserUpdate**](IamUserUpdate.md)| Password/description. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

