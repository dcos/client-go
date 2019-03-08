# \IamApi

All URIs are relative to *http://your-dcos-cluster.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConfigureOIDCProvider**](IamApi.md#ConfigureOIDCProvider) | **Put** /acs/api/v1/auth/oidc/providers/{provider-id} | Configure a new OIDC provider.
[**ConfigureSAMLProvider**](IamApi.md#ConfigureSAMLProvider) | **Put** /acs/api/v1/auth/saml/providers/{provider-id} | Configure a new SAML provider.
[**CreateGroup**](IamApi.md#CreateGroup) | **Put** /acs/api/v1/groups/{gid} | Create a group.
[**CreateGroupUser**](IamApi.md#CreateGroupUser) | **Put** /acs/api/v1/groups/{gid}/users/{uid} | Add account to group.
[**CreateLDAPConfiguration**](IamApi.md#CreateLDAPConfiguration) | **Put** /acs/api/v1/ldap/config | Set new LDAP configuration.
[**CreateResourceACL**](IamApi.md#CreateResourceACL) | **Put** /acs/api/v1/acls/{rid} | Create ACL for a certain resource.
[**CreateUser**](IamApi.md#CreateUser) | **Put** /acs/api/v1/users/{uid} | Create user account.
[**DeleteGroup**](IamApi.md#DeleteGroup) | **Delete** /acs/api/v1/groups/{gid} | Delete group.
[**DeleteGroupUser**](IamApi.md#DeleteGroupUser) | **Delete** /acs/api/v1/groups/{gid}/users/{uid} | Delete user account from group.
[**DeleteLDAPConfiguration**](IamApi.md#DeleteLDAPConfiguration) | **Delete** /acs/api/v1/ldap/config | Delete current LDAP configuration.
[**DeleteOIDCProvider**](IamApi.md#DeleteOIDCProvider) | **Delete** /acs/api/v1/auth/oidc/providers/{provider-id} | Delete provider.
[**DeleteResourceACL**](IamApi.md#DeleteResourceACL) | **Delete** /acs/api/v1/acls/{rid} | Delete ACL for a certain resource.
[**DeleteSAMLProvider**](IamApi.md#DeleteSAMLProvider) | **Delete** /acs/api/v1/auth/saml/providers/{provider-id} | Delete provider.
[**DeleteUser**](IamApi.md#DeleteUser) | **Delete** /acs/api/v1/users/{uid} | Delete account.
[**ForbidResourceUserAction**](IamApi.md#ForbidResourceUserAction) | **Delete** /acs/api/v1/acls/{rid}/users/{uid}/{action} | Forbid single action for given account and resource.
[**ForbidResourceUserActions**](IamApi.md#ForbidResourceUserActions) | **Delete** /acs/api/v1/acls/{rid}/users/{uid} | Forbid all actions of given account to given resource.
[**ForbitResourceGroupAction**](IamApi.md#ForbitResourceGroupAction) | **Delete** /acs/api/v1/acls/{rid}/groups/{gid}/{action} | Forbid single action for given resource and group.
[**ForbitResourceGroupActions**](IamApi.md#ForbitResourceGroupActions) | **Delete** /acs/api/v1/acls/{rid}/groups/{gid} | Forbid all actions of given group to given resource.
[**GetACLs**](IamApi.md#GetACLs) | **Get** /acs/api/v1/acls | Retrieve all ACL objects.
[**GetGroup**](IamApi.md#GetGroup) | **Get** /acs/api/v1/groups/{gid} | Get single group object.
[**GetGroupPermissions**](IamApi.md#GetGroupPermissions) | **Get** /acs/api/v1/groups/{gid}/permissions | Retrieve group permissions.
[**GetGroupUsers**](IamApi.md#GetGroupUsers) | **Get** /acs/api/v1/groups/{gid}/users | Retrieve members of a group.
[**GetGroups**](IamApi.md#GetGroups) | **Get** /acs/api/v1/groups | Retrieve all group objects.
[**GetJWKS**](IamApi.md#GetJWKS) | **Get** /acs/api/v1/auth/jwks | Get the IAM&#39;s JSON Web Key Set (JWKS, according to RFCs 7517/7518).
[**GetLDAPConfiguration**](IamApi.md#GetLDAPConfiguration) | **Get** /acs/api/v1/ldap/config | Retrieve current LDAP configuration.
[**GetOIDCProvider**](IamApi.md#GetOIDCProvider) | **Get** /acs/api/v1/auth/oidc/providers/{provider-id} | Get configuration for a specific provider.
[**GetOIDCProviders**](IamApi.md#GetOIDCProviders) | **Get** /acs/api/v1/auth/oidc/providers | Get an overview for the configured OIDC providers.
[**GetResourceACLs**](IamApi.md#GetResourceACLs) | **Get** /acs/api/v1/acls/{rid} | Retrieve ACL for a certain resource.
[**GetResourceGroupAction**](IamApi.md#GetResourceGroupAction) | **Get** /acs/api/v1/acls/{rid}/groups/{gid}/{action} | Query whether action is allowed or not.
[**GetResourceGroupActions**](IamApi.md#GetResourceGroupActions) | **Get** /acs/api/v1/acls/{rid}/groups/{gid} | Get allowed actions for given resource and group.
[**GetResourcePermissions**](IamApi.md#GetResourcePermissions) | **Get** /acs/api/v1/acls/{rid}/permissions | Retrieve all permissions for resource.
[**GetResourceUserAction**](IamApi.md#GetResourceUserAction) | **Get** /acs/api/v1/acls/{rid}/users/{uid}/{action} | Query whether action is allowed or not.
[**GetResourceUserActions**](IamApi.md#GetResourceUserActions) | **Get** /acs/api/v1/acls/{rid}/users/{uid} | Get allowed actions for given resource and user.
[**GetSAMLProvider**](IamApi.md#GetSAMLProvider) | **Get** /acs/api/v1/auth/saml/providers/{provider-id} | Get configuration for a specific SAML provider.
[**GetSAMLProviderACSCallbackURL**](IamApi.md#GetSAMLProviderACSCallbackURL) | **Get** /acs/api/v1/auth/saml/providers/{provider-id}/acs-callback-url | Get the authentication callback URL for this SP.
[**GetSAMLProviderSPMetadata**](IamApi.md#GetSAMLProviderSPMetadata) | **Get** /acs/api/v1/auth/saml/providers/{provider-id}/sp-metadata | Get SP metadata (XML).
[**GetSAMLProviders**](IamApi.md#GetSAMLProviders) | **Get** /acs/api/v1/auth/saml/providers | Get an overview for the configured SAML 2.0 providers.
[**GetUser**](IamApi.md#GetUser) | **Get** /acs/api/v1/users/{uid} | Get single user object.
[**GetUserGroups**](IamApi.md#GetUserGroups) | **Get** /acs/api/v1/users/{uid}/groups | Retrieve groups the user is member of.
[**GetUserPermissions**](IamApi.md#GetUserPermissions) | **Get** /acs/api/v1/users/{uid}/permissions | Retrieve permissions an account has.
[**GetUsers**](IamApi.md#GetUsers) | **Get** /acs/api/v1/users | Retrieve all regular user accounts or service user accounts.
[**ImportLDAPGroup**](IamApi.md#ImportLDAPGroup) | **Post** /acs/api/v1/ldap/importgroup | Import an LDAP group.
[**ImportLDAPUser**](IamApi.md#ImportLDAPUser) | **Post** /acs/api/v1/ldap/importuser | Import an LDAP user.
[**Login**](IamApi.md#Login) | **Post** /acs/api/v1/auth/login | Log in (obtain a DC/OS authentication token).
[**LoginWithProvider**](IamApi.md#LoginWithProvider) | **Get** /acs/api/v1/auth/login | Log in using an external identity provider.
[**OpenIDConnectCallbackURL**](IamApi.md#OpenIDConnectCallbackURL) | **Get** /acs/api/v1/auth/oidc/callback | OpenID Connect callback URL.
[**PermitResourceGroupAction**](IamApi.md#PermitResourceGroupAction) | **Put** /acs/api/v1/acls/{rid}/groups/{gid}/{action} | Permit single action for given resource and group.
[**PermitResourceUserAction**](IamApi.md#PermitResourceUserAction) | **Put** /acs/api/v1/acls/{rid}/users/{uid}/{action} | Permit single action for given account and resource.
[**SAMLProviderACSCallbackEndpoint**](IamApi.md#SAMLProviderACSCallbackEndpoint) | **Post** /acs/api/v1/auth/saml/providers/{provider-id}/acs-callback | The SP ACS callback endpoint.
[**TestLDAPBackendConnection**](IamApi.md#TestLDAPBackendConnection) | **Post** /acs/api/v1/ldap/config/test | Test connection to the LDAP back-end.
[**UpdateGroup**](IamApi.md#UpdateGroup) | **Patch** /acs/api/v1/groups/{gid} | Update group.
[**UpdateOIDCProvider**](IamApi.md#UpdateOIDCProvider) | **Patch** /acs/api/v1/auth/oidc/providers/{provider-id} | Update OIDC provider config.
[**UpdateResourceACL**](IamApi.md#UpdateResourceACL) | **Patch** /acs/api/v1/acls/{rid} | Update ACL for a certain resource.
[**UpdateSAMLProvider**](IamApi.md#UpdateSAMLProvider) | **Patch** /acs/api/v1/auth/saml/providers/{provider-id} | Update SAML provider config.
[**UpdateUser**](IamApi.md#UpdateUser) | **Patch** /acs/api/v1/users/{uid} | Update user account.


# **ConfigureOIDCProvider**
> ConfigureOIDCProvider(ctx, providerId, oidcProviderConfig)
Configure a new OIDC provider.

Set up OIDC provider with the ID as specified in the URL, and with the config as specified via JSON in the request body.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **providerId** | **string**| The ID of the provider to create. | 
  **oidcProviderConfig** | [**OidcProviderConfig**](OidcProviderConfig.md)| Provider config JSON object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConfigureSAMLProvider**
> ConfigureSAMLProvider(ctx, providerId, samlProviderConfig)
Configure a new SAML provider.

Set up a SAML provider with the ID as specified in the URL, and with the config as given by the JSON document in the request body.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **providerId** | **string**| The ID of the provider to create. | 
  **samlProviderConfig** | [**SamlProviderConfig**](SamlProviderConfig.md)| Provider config JSON object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateGroup**
> CreateGroup(ctx, gid, groupCreate)
Create a group.

Create a group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gid** | **string**| The ID of the group. | 
  **groupCreate** | [**GroupCreate**](GroupCreate.md)|  | 

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
> CreateLDAPConfiguration(ctx, ldapConfiguration)
Set new LDAP configuration.

Set new directory (LDAP) back-end configuration. Replace current configuration, if existing.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ldapConfiguration** | [**LdapConfiguration**](LdapConfiguration.md)| JSON object containing the LDAP configuration details. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateResourceACL**
> CreateResourceACL(ctx, rid, aclCreate)
Create ACL for a certain resource.

Create new ACL for resource with ID `rid` (description in body, no permissions by default).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **rid** | **string**| The ID of the resource for which the ACL should be created. | 
  **aclCreate** | [**AclCreate**](AclCreate.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateUser**
> CreateUser(ctx, uid, userCreate)
Create user account.

Create user (uid in url, details incl. credentials in body).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uid** | **string**| The ID of the user account to create. | 
  **userCreate** | [**UserCreate**](UserCreate.md)| Password/description. | 

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
> InlineResponse200 GetACLs(ctx, )
Retrieve all ACL objects.

Get array of `ACL` objects.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGroup**
> Group GetGroup(ctx, gid)
Get single group object.

Get specific `Group` object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gid** | **string**| The ID of the group to retrieve. | 

### Return type

[**Group**](Group.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGroupPermissions**
> GroupPermissions GetGroupPermissions(ctx, gid)
Retrieve group permissions.

Retrieve permissions of this group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gid** | **string**| The group ID. | 

### Return type

[**GroupPermissions**](GroupPermissions.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGroupUsers**
> GroupUsers GetGroupUsers(ctx, gid, optional)
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

[**GroupUsers**](GroupUsers.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGroups**
> InlineResponse2002 GetGroups(ctx, )
Retrieve all group objects.

Retrieve array of `Group` objects.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

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
> LdapConfiguration GetLDAPConfiguration(ctx, )
Retrieve current LDAP configuration.

Retrieve current directory (LDAP) back-end configuration.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**LdapConfiguration**](LDAPConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOIDCProvider**
> OidcProviderConfig GetOIDCProvider(ctx, providerId)
Get configuration for a specific provider.

Get configuration for a specific provider.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **providerId** | **string**| The ID of the OIDC provider to retrieve the config for. | 

### Return type

[**OidcProviderConfig**](OIDCProviderConfig.md)

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
> Acl GetResourceACLs(ctx, rid)
Retrieve ACL for a certain resource.

Retrieve single `ACL` object, for a specific resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **rid** | **string**| The ID of the resource to retrieve the ACL for. | 

### Return type

[**Acl**](ACL.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetResourceGroupAction**
> ActionAllowed GetResourceGroupAction(ctx, rid, gid, action)
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

[**ActionAllowed**](ActionAllowed.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetResourceGroupActions**
> InlineResponse2001 GetResourceGroupActions(ctx, rid, gid)
Get allowed actions for given resource and group.

Get allowed actions for given resource and group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **rid** | **string**| resource ID | 
  **gid** | **string**| group ID | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetResourcePermissions**
> AclPermissions GetResourcePermissions(ctx, rid)
Retrieve all permissions for resource.

Retrieve all permissions that are set for a specific resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **rid** | **string**| resource ID | 

### Return type

[**AclPermissions**](ACLPermissions.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetResourceUserAction**
> ActionAllowed GetResourceUserAction(ctx, rid, uid, action)
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

[**ActionAllowed**](ActionAllowed.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetResourceUserActions**
> InlineResponse2001 GetResourceUserActions(ctx, rid, uid)
Get allowed actions for given resource and user.

Get allowed actions for given resource and user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **rid** | **string**| resource ID | 
  **uid** | **string**| account ID | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSAMLProvider**
> SamlProviderConfig GetSAMLProvider(ctx, providerId)
Get configuration for a specific SAML provider.

Get configuration for a specific SAML provider.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **providerId** | **string**| The ID of the SAML provider to retrieve the config for. | 

### Return type

[**SamlProviderConfig**](SAMLProviderConfig.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSAMLProviderACSCallbackURL**
> SamlacsCallbackUrlObject GetSAMLProviderACSCallbackURL(ctx, providerId)
Get the authentication callback URL for this SP.

The IAM acts as SAML service provider (SP). A SAML identity provider (IdP) usually requires to be configured with the Assertion Consumer Service (ACS) callback URL of the SP (which is where the IdP makes the end-user submit the authentication response).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **providerId** | **string**| The ID of the SAML provider to retrieve the ACS callback URL for. | 

### Return type

[**SamlacsCallbackUrlObject**](SAMLACSCallbackUrlObject.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSAMLProviderSPMetadata**
> GetSAMLProviderSPMetadata(ctx, providerId)
Get SP metadata (XML).

The IAM acts as SAML service provider (SP). This endpoint provides the SP metadata as an XML document. Certain identity providers (IdPs) may want to directly consume this document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **providerId** | **string**| The ID of the SAML provider to retrieve the metadata for. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

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
> User GetUser(ctx, uid)
Get single user object.

Get specific `User` object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uid** | **string**| The ID of the user object to retrieve. | 

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserGroups**
> UserGroups GetUserGroups(ctx, uid)
Retrieve groups the user is member of.

Retrieve groups the user is member of.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uid** | **string**| The ID of the user. | 

### Return type

[**UserGroups**](UserGroups.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserPermissions**
> UserPermissions GetUserPermissions(ctx, uid)
Retrieve permissions an account has.

Retrieve the permissions for this account with direct permissions distinguished from those that are obtained through group membership.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uid** | **string**| The id of the user. | 

### Return type

[**UserPermissions**](UserPermissions.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUsers**
> InlineResponse2003 GetUsers(ctx, optional)
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

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImportLDAPGroup**
> ImportLDAPGroup(ctx, ldapImportGroupObject)
Import an LDAP group.

Attempt to import a group of users from the configured directory (LDAP) back-end. See docs/ldap.md for details on group import.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ldapImportGroupObject** | [**LdapImportGroupObject**](LdapImportGroupObject.md)| A JSON object specifying the name of the group to be imported. The meaning of the name depends on the group search settings. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImportLDAPUser**
> ImportLDAPUser(ctx, ldapImportUserObject)
Import an LDAP user.

Attempt to import a user from the configured directory (LDAP) back-end.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ldapImportUserObject** | [**LdapImportUserObject**](LdapImportUserObject.md)| A JSON object specifying the username (read: \&quot;login\&quot; or \&quot;user ID\&quot;) of the user that should be imported. That string is equivalent to the &#x60;uid&#x60; the user is supposed to log in with after successful import. The exact meaning of this string depends on the configured LDAP authentication method. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Login**
> AuthToken Login(ctx, loginObject)
Log in (obtain a DC/OS authentication token).

Exchange user credentials (regular user account: uid and password; service user account: uid and service login token) for a DC/OS authentication token. The resulting DC/OS authentication token is an RFC 7519 JSON Web Token (JWT) of type RS256. It has a limited lifetime which depends on the IAM configuration (only, i.e. the lifetime cannot be chosen as part of the login HTTP request). The DC/OS authentication token can be verified out-of-band using a standards-compliant RS256 JWT verification procedure based on the long-lived public key material presented by the IAM's /auth/jwks endpoint, and by requiring the two claims `exp` and `uid` to be present.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **loginObject** | [**LoginObject**](LoginObject.md)| uid &amp; password or uid &amp; service login token. | 

### Return type

[**AuthToken**](AuthToken.md)

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
> LdapTestResultObject TestLDAPBackendConnection(ctx, ldapTestCredentials)
Test connection to the LDAP back-end.

Perform basic feature tests. Verify that the current directory (LDAP) configuration parameters allow for a successful connection to the directory back-end. For instance, this endpoint simulates the procedure for authentication via LDAP, but provides more useful feedback upon failure than the actual login endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ldapTestCredentials** | [**LdapTestCredentials**](LdapTestCredentials.md)| JSON object containing &#x60;uid&#x60; and password of an LDAP user. For the most expressive test result, choose credentials different from the lookup credentials. The &#x60;uid&#x60; is the string the user is supposed to log in with after successful LDAP back-end configuration. | 

### Return type

[**LdapTestResultObject**](LDAPTestResultObject.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateGroup**
> UpdateGroup(ctx, gid, groupUpdate)
Update group.

Update existing group (description).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gid** | **string**| The ID of the group to modify. | 
  **groupUpdate** | [**GroupUpdate**](GroupUpdate.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOIDCProvider**
> UpdateOIDCProvider(ctx, providerId, oidcProviderConfig)
Update OIDC provider config.

Update config for existing OIDC provider.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **providerId** | **string**| The ID of the provider to modify. | 
  **oidcProviderConfig** | [**OidcProviderConfig**](OidcProviderConfig.md)| Provider config JSON object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateResourceACL**
> UpdateResourceACL(ctx, rid, aclUpdate)
Update ACL for a certain resource.

Update ACL for resource with ID `rid`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **rid** | **string**| The ID of the resource for which the ACL should be created. | 
  **aclUpdate** | [**AclUpdate**](AclUpdate.md)| New ACL. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSAMLProvider**
> UpdateSAMLProvider(ctx, providerId, samlProviderConfig)
Update SAML provider config.

Update config for existing SAML provider.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **providerId** | **string**| The ID of the provider to modify. | 
  **samlProviderConfig** | [**SamlProviderConfig**](SamlProviderConfig.md)| Provider config JSON object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateUser**
> UpdateUser(ctx, uid, userUpdate)
Update user account.

Update existing user account (meta data and/or password).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uid** | **string**| The ID of the user account to modify. | 
  **userUpdate** | [**UserUpdate**](UserUpdate.md)| Password/description. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

