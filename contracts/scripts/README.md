# Deployment scripts of Scroll contracts

## Deployment using Hardhat

The scripts should run as below sequence:

```bash

# generate alloc(storage for pre contract) from l2getg
npx hardhat run scripts/generate_precontract_storage.ts

export layer1=l1geth # change to actual network name
export layer2=l2geth # change to actual network name
export owner=0x0000000000000000000000000000000000000000 # change to actual owner

# deploy contracts in layer 1
npx hardhat --network $layer1 run scripts/deploy_proxy_admin.ts
npx hardhat --network $layer1 run scripts/deploy_scroll_chain.ts
npx hardhat --network $layer1 run scripts/deploy_weth.ts
npx hardhat --network $layer1 run scripts/deploy_whitelist.ts
env CONTRACT_NAME=L1ScrollMessenger npx hardhat run --network $layer1 scripts/deploy_proxy_contract.ts
env CONTRACT_NAME=L1GatewayRouter npx hardhat run --network $layer1 scripts/deploy_proxy_contract.ts
env CONTRACT_NAME=L1StandardERC20Gateway npx hardhat run --network $layer1 scripts/deploy_proxy_contract.ts
env CONTRACT_NAME=L1CustomERC20Gateway npx hardhat run --network $layer1 scripts/deploy_proxy_contract.ts
env CONTRACT_NAME=L1ERC721Gateway npx hardhat run --network $layer1 scripts/deploy_proxy_contract.ts
env CONTRACT_NAME=L1ERC1155Gateway npx hardhat run --network $layer1 scripts/deploy_proxy_contract.ts
env CONTRACT_NAME=L1WETHGateway npx hardhat run --network $layer1 scripts/deploy_proxy_contract.ts
env CONTRACT_NAME=L1ETHGateway npx hardhat run --network $layer1 scripts/deploy_proxy_contract.ts
env CONTRACT_NAME=L1MessageQueue npx hardhat run --network $layer1 scripts/deploy_proxy_contract.ts
env CONTRACT_NAME=L2GasPriceOracle npx hardhat run --network $layer1 scripts/deploy_proxy_contract.ts
env CONTRACT_NAME=EnforcedTxGateway npx hardhat run --network $layer1 scripts/deploy_proxy_contract.ts

# deploy contracts in layer 2, note: l2_messenger is predeployed
npx hardhat --network $layer2 run scripts/deploy_whitelist.ts
npx hardhat --network $layer2 run scripts/deploy_proxy_admin.ts
npx hardhat --network $layer2 run scripts/deploy_l2_messenger.ts
npx hardhat --network $layer2 run scripts/deploy_l2_token_factory.ts
env CONTRACT_NAME=L2GatewayRouter npx hardhat run --network $layer2 scripts/deploy_proxy_contract.ts
env CONTRACT_NAME=L2StandardERC20Gateway npx hardhat run --network $layer2 scripts/deploy_proxy_contract.ts
env CONTRACT_NAME=L2CustomERC20Gateway npx hardhat run --network $layer2 scripts/deploy_proxy_contract.ts
env CONTRACT_NAME=L2ERC721Gateway npx hardhat run --network $layer2 scripts/deploy_proxy_contract.ts
env CONTRACT_NAME=L2ERC1155Gateway npx hardhat run --network $layer2 scripts/deploy_proxy_contract.ts
env CONTRACT_NAME=L2ETHGateway npx hardhat run --network $layer2 scripts/deploy_proxy_contract.ts
env CONTRACT_NAME=L2WETHGateway npx hardhat run --network $layer2 scripts/deploy_proxy_contract.ts


# initalize contracts in layer 1, should set proper bash env variables first
npx hardhat --network $layer1 run scripts/initialize_scroll_chain.ts
npx hardhat --network $layer1 run scripts/initializeL2GasPriceOracle.ts
npx hardhat --network $layer1 run scripts/initializeL1MessageQueue.ts
npx hardhat --network $layer1 run scripts/initializeL1Messager.ts
npx hardhat --network $layer1 run scripts/initializeL1ETHGateway.ts
npx hardhat --network $layer1 run ./scripts/initialize_l1_gateway_router.ts

<!-- npx hardhat --network $layer1 run scripts/initialize_l1_erc20_gateway.ts
npx hardhat --network $layer1 run scripts/initialize_l1_gateway_router.ts
npx hardhat --network $layer1 run scripts/initialize_l1_messenger.ts
npx hardhat --network $layer1 run scripts/initialize_l1_custom_erc20_gateway.ts
npx hardhat --network $layer1 run scripts/initialize_l1_erc1155_gateway.ts
npx hardhat --network $layer1 run scripts/initialize_l1_erc721_gateway.ts -->

#initalize contracts in layer 2, should set proper bash env variables first
npx hardhat --network $layer2 run scripts/initialize_l2_predeploys.ts
npx hardhat --network $layer2 run scripts/initialize_l2_messenger.ts
npx hardhat --network $layer2 run scripts/initialize_l2_eth_gateway.ts
npx hardhat --network $layer2 run scripts/initialize_l2_gateway_router.ts

<!-- npx hardhat --network $layer2 run scripts/initialize_l2_erc20_gateway.ts
npx hardhat --network $layer2 run scripts/initialize_l2_gateway_router.ts
npx hardhat --network $layer2 run scripts/initialize_l2_custom_erc20_gateway.ts
npx hardhat --network $layer2 run scripts/initialize_l2_erc1155_gateway.ts
npx hardhat --network $layer2 run scripts/initialize_l2_erc721_gateway.ts
npx hardhat --network $layer2 run scripts/initialize_l2_token_factory.ts -->

# transfer ownership in layer 1

env CONTRACT_NAME=ProxyAdmin CONTRACT_OWNER=$owner npx hardhat run --network $layer1 scripts/transfer_ownership.ts
env CONTRACT_NAME=L1ScrollMessenger CONTRACT_OWNER=$owner npx hardhat run --network $layer1 scripts/transfer_ownership.ts
env CONTRACT_NAME=ZKRollup CONTRACT_OWNER=$owner npx hardhat run --network $layer1 scripts/transfer_ownership.ts
env CONTRACT_NAME=L1GatewayRouter CONTRACT_OWNER=$owner npx hardhat run --network $layer1 scripts/transfer_ownership.ts
env CONTRACT_NAME=L1CustomERC20Gateway CONTRACT_OWNER=$owner npx hardhat run --network $layer1 scripts/transfer_ownership.ts
env CONTRACT_NAME=L1ERC721Gateway CONTRACT_OWNER=$owner npx hardhat run --network $layer1 scripts/transfer_ownership.ts
env CONTRACT_NAME=L1ERC1155Gateway CONTRACT_OWNER=$owner npx hardhat run --network $layer1 scripts/transfer_ownership.ts

# transfer ownership in layer 2

env CONTRACT_NAME=ProxyAdmin CONTRACT_OWNER=$owner npx hardhat run --network $layer2 scripts/transfer_ownership.ts
env CONTRACT_NAME=L2ScrollMessenger CONTRACT_OWNER=$owner npx hardhat run --network $layer2 scripts/transfer_ownership.ts
env CONTRACT_NAME=L2GatewayRouter CONTRACT_OWNER=$owner npx hardhat run --network $layer2 scripts/transfer_ownership.ts
env CONTRACT_NAME=L2CustomERC20Gateway CONTRACT_OWNER=$owner npx hardhat run --network $layer2 scripts/transfer_ownership.ts
env CONTRACT_NAME=L2ERC721Gateway CONTRACT_OWNER=$owner npx hardhat run --network $layer2 scripts/transfer_ownership.ts
env CONTRACT_NAME=L2ERC1155Gateway CONTRACT_OWNER=$owner npx hardhat run --network $layer2 scripts/transfer_ownership.ts
```
