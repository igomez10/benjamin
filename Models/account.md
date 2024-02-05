# account
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **id** | **String** | The account identification value | [optional] [default to null] |
| **accountId** | **String** | The account number | [optional] [default to null] |
| **accountVan** | **String** | The accountAlias | [optional] [default to null] |
| **accountTitle** | **String** | Title of the account | [optional] [default to null] |
| **displayName** | **String** | Whichever value is not null in this priority | [optional] [default to null] |
| **accountAlias** | **String** | User customizable account alias. Refer to [Configure Account Alias](https://guides.interactivebrokers.com/cp/cp.htm#am/settings/accountalias.htm) for details. | [optional] [default to null] |
| **accountStatus** | **BigDecimal** | When the account was opened in unix time. | [optional] [default to null] |
| **currency** | **String** | Base currency of the account. | [optional] [default to null] |
| **type** | **String** | Account Type | [optional] [default to null] |
| **tradingType** | **String** | UNI - Deprecated property | [optional] [default to null] |
| **faclient** | **Boolean** | If an account is a sub-account to a Financial Advisor. | [optional] [default to null] |
| **clearingStatus** | **String** | Status of the Account   * O &#x3D; Open   * P or N &#x3D; Pending   * A &#x3D; Abandoned   * R &#x3D; Rejected   * C &#x3D; Closed  | [optional] [default to null] |
| **covestor** | **Boolean** | Is a Covestor Account | [optional] [default to null] |
| **parent** | [**account_parent**](account_parent.md) |  | [optional] [default to null] |
| **desc** | **String** | Formatted \&quot;accountId - accountAlias\&quot; | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

