# _portfolio_subaccounts2_get_200_response_subaccounts_inner
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
| **clearingStatus** | **String** | Status of the Account   * O &#x3D; Open   * P or N &#x3D; Pending   * A &#x3D; Abandoned   * R &#x3D; Rejected   * C &#x3D; Closed   covestor:     type: boolean     description: Is a Covestor Account   parent:     type: object     properties:       mmc:         type: array         items:           type: string           description: Money Manager Client (MMC) Account       accountId:         type: string         description: Account Number for Money Manager Client       isMParent:         type: boolean         description: Is MM a Parent Account       isMChild:         type: boolean         description: Is MM a Child Account       isMultiplex:         type: boolean         description: Is a Multiplex Account. These are account models with individual account being parent and managed account being child.   desc:     type: string     description: Formatted \&quot;accountId - accountAlias\&quot;  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

