# L1StandardERC20Gateway



> L1StandardERC20Gateway

The `L1StandardERC20Gateway` is used to deposit standard ERC20 tokens on layer 1 and finalize withdraw the tokens from layer 2.

*The deposited ERC20 tokens are held in this gateway. On finalizing withdraw, the corresponding token will be transfer to the recipient directly. Any ERC20 that requires non-standard functionality should use a separate gateway.*

## Methods

### counterpart

```solidity
function counterpart() external view returns (address)
```

The address of corresponding L1/L2 Gateway contract.




#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | address | undefined |

### depositERC20

```solidity
function depositERC20(address _token, uint256 _amount, uint256 _gasLimit) external payable
```

Deposit some token to a caller&#39;s account on L2.

*Make this function payable to send relayer fee in Ether.*

#### Parameters

| Name | Type | Description |
|---|---|---|
| _token | address | The address of token in L1. |
| _amount | uint256 | The amount of token to transfer. |
| _gasLimit | uint256 | Gas limit required to complete the deposit on L2. |

### depositERC20

```solidity
function depositERC20(address _token, address _to, uint256 _amount, uint256 _gasLimit) external payable
```

Deposit some token to a recipient&#39;s account on L2.

*Make this function payable to send relayer fee in Ether.*

#### Parameters

| Name | Type | Description |
|---|---|---|
| _token | address | The address of token in L1. |
| _to | address | The address of recipient&#39;s account on L2. |
| _amount | uint256 | The amount of token to transfer. |
| _gasLimit | uint256 | Gas limit required to complete the deposit on L2. |

### depositERC20AndCall

```solidity
function depositERC20AndCall(address _token, address _to, uint256 _amount, bytes _data, uint256 _gasLimit) external payable
```

Deposit some token to a recipient&#39;s account on L2 and call.

*Make this function payable to send relayer fee in Ether.*

#### Parameters

| Name | Type | Description |
|---|---|---|
| _token | address | The address of token in L1. |
| _to | address | The address of recipient&#39;s account on L2. |
| _amount | uint256 | The amount of token to transfer. |
| _data | bytes | Optional data to forward to recipient&#39;s account. |
| _gasLimit | uint256 | Gas limit required to complete the deposit on L2. |

### finalizeWithdrawERC20

```solidity
function finalizeWithdrawERC20(address _l1Token, address _l2Token, address _from, address _to, uint256 _amount, bytes _data) external payable
```

Complete ERC20 withdraw from L2 to L1 and send fund to recipient&#39;s account in L1.

*Make this function payable to handle WETH deposit/withdraw.      The function should only be called by L1ScrollMessenger.      The function should also only be called by L2ERC20Gateway in L2.*

#### Parameters

| Name | Type | Description |
|---|---|---|
| _l1Token | address | The address of corresponding L1 token. |
| _l2Token | address | The address of corresponding L2 token. |
| _from | address | The address of account who withdraw the token in L2. |
| _to | address | The address of recipient in L1 to receive the token. |
| _amount | uint256 | The amount of the token to withdraw. |
| _data | bytes | Optional data to forward to recipient&#39;s account. |

### getL2ERC20Address

```solidity
function getL2ERC20Address(address _l1Token) external view returns (address)
```

Return the corresponding l2 token address given l1 token address.



#### Parameters

| Name | Type | Description |
|---|---|---|
| _l1Token | address | The address of l1 token. |

#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | address | undefined |

### initialize

```solidity
function initialize(address _counterpart, address _router, address _messenger, address _l2TokenImplementation, address _l2TokenFactory) external nonpayable
```

Initialize the storage of L1StandardERC20Gateway.



#### Parameters

| Name | Type | Description |
|---|---|---|
| _counterpart | address | The address of L2StandardERC20Gateway in L2. |
| _router | address | The address of L1GatewayRouter. |
| _messenger | address | The address of L1ScrollMessenger. |
| _l2TokenImplementation | address | The address of ScrollStandardERC20 implementation in L2. |
| _l2TokenFactory | address | The address of ScrollStandardERC20Factory contract in L2. |

### l2TokenFactory

```solidity
function l2TokenFactory() external view returns (address)
```

The address of ScrollStandardERC20Factory contract in L2.




#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | address | undefined |

### l2TokenImplementation

```solidity
function l2TokenImplementation() external view returns (address)
```

The address of ScrollStandardERC20 implementation in L2.




#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | address | undefined |

### messenger

```solidity
function messenger() external view returns (address)
```

The address of corresponding L1ScrollMessenger/L2ScrollMessenger contract.




#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | address | undefined |

### onDropMessage

```solidity
function onDropMessage(bytes _message) external payable
```





#### Parameters

| Name | Type | Description |
|---|---|---|
| _message | bytes | undefined |

### owner

```solidity
function owner() external view returns (address)
```



*Returns the address of the current owner.*


#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | address | undefined |

### renounceOwnership

```solidity
function renounceOwnership() external nonpayable
```



*Leaves the contract without owner. It will not be possible to call `onlyOwner` functions. Can only be called by the current owner. NOTE: Renouncing ownership will leave the contract without an owner, thereby disabling any functionality that is only available to the owner.*


### router

```solidity
function router() external view returns (address)
```

The address of L1GatewayRouter/L2GatewayRouter contract.




#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | address | undefined |

### transferOwnership

```solidity
function transferOwnership(address newOwner) external nonpayable
```



*Transfers ownership of the contract to a new account (`newOwner`). Can only be called by the current owner.*

#### Parameters

| Name | Type | Description |
|---|---|---|
| newOwner | address | undefined |



## Events

### DepositERC20

```solidity
event DepositERC20(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data)
```

Emitted when someone deposit ERC20 token from L1 to L2.



#### Parameters

| Name | Type | Description |
|---|---|---|
| l1Token `indexed` | address | undefined |
| l2Token `indexed` | address | undefined |
| from `indexed` | address | undefined |
| to  | address | undefined |
| amount  | uint256 | undefined |
| data  | bytes | undefined |

### FinalizeWithdrawERC20

```solidity
event FinalizeWithdrawERC20(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data)
```

Emitted when ERC20 token is withdrawn from L2 to L1 and transfer to recipient.



#### Parameters

| Name | Type | Description |
|---|---|---|
| l1Token `indexed` | address | undefined |
| l2Token `indexed` | address | undefined |
| from `indexed` | address | undefined |
| to  | address | undefined |
| amount  | uint256 | undefined |
| data  | bytes | undefined |

### Initialized

```solidity
event Initialized(uint8 version)
```





#### Parameters

| Name | Type | Description |
|---|---|---|
| version  | uint8 | undefined |

### OwnershipTransferred

```solidity
event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
```





#### Parameters

| Name | Type | Description |
|---|---|---|
| previousOwner `indexed` | address | undefined |
| newOwner `indexed` | address | undefined |

### RefundERC20

```solidity
event RefundERC20(address indexed token, address indexed recipient, uint256 amount)
```

Emitted when some ERC20 token is refunded.



#### Parameters

| Name | Type | Description |
|---|---|---|
| token `indexed` | address | undefined |
| recipient `indexed` | address | undefined |
| amount  | uint256 | undefined |



