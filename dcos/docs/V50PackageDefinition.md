# V50PackageDefinition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Command** | [**Command**](command.md) |  | [optional] 
**Description** | **string** |  | 
**DowngradesTo** | **[]string** | List of versions that this package can downgrade to. If the property is a list containing the string &#39;*&#39;, this package can downgrade to any version. If the property is not set or the empty list, this package cannot downgrade. | [optional] 
**Framework** | **bool** | True if this package installs a new Mesos framework. | [optional] [default to false]
**Licenses** | [**[]Licence**](licence.md) |  | [optional] 
**Maintainer** | **string** |  | 
**Marathon** | [**Marathon**](marathon.md) |  | [optional] 
**MinDcosReleaseVersion** | **string** | The minimum DC/OS Release Version the package can run on. | [optional] 
**Name** | **string** |  | 
**PackagingVersion** | **string** |  | 
**PostInstallNotes** | **string** | Post installation notes that would be useful to the user of this package. | [optional] 
**PostUninstallNotes** | **string** | Post uninstallation notes that would be useful to the user of this package. | [optional] 
**PreInstallNotes** | **string** | Pre installation notes that would be useful to the user of this package. | [optional] 
**ReleaseVersion** | **int32** | Corresponds to the revision index from the universe directory structure | 
**Resource** | [**V30Resource**](v30Resource.md) |  | [optional] 
**Scm** | **string** |  | [optional] 
**Selected** | **bool** | Flag indicating if the package is selected in search results | [optional] [default to false]
**Tags** | **[]string** |  | 
**UpgradesFrom** | **[]string** | List of versions that can upgrade to this package. If the property is a list containing the string &#39;*&#39;, any version can upgrade to this package. If the property is not set or the empty list, no version can upgrade to this package. | [optional] 
**Version** | **string** |  | 
**Website** | **string** |  | [optional] 
**Manager** | [**Manager**](manager.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


