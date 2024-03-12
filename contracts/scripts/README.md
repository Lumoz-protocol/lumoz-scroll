# Deployment scripts of Scroll contracts

## Deployment using Hardhat

The scripts should run as below sequence:

```bash
export layer1=l1geth # change to actual network name
export layer2=l2geth # change to actual network name
export owner=0x0000000000000000000000000000000000000000 # change to real owner address

# generate alloc(storage for pre contract) from l2getg
npx hardhat run scripts/generate_precontract_storage.js

# prepare
npx hardhat --network $layer1 run scripts/deploy_proxy_admin.ts
npx hardhat --network $layer2 run scripts/deploy_proxy_admin.ts
npx hardhat --network $layer1 run scripts/compute_proxy.ts
npx hardhat --network $layer2 run scripts/compute_proxy.ts
npx hardhat --network $layer2 run scripts/deploy_l2_token_factory.ts

# deploy weth on L1 when env L1_WETH_ADDR not set
npx hardhat --network $layer1 run scripts/deploy_weth.ts

# deploy contracts in layer 1, set plonk verifier before deploy scroll chain, deploy weth only for test
npx hardhat --network $layer1 run scripts/deploy_scroll_chain.ts # deploy scrollchain
env MOCK_PROVER=true npx hardhat --network $layer1 run scripts/deploy_scroll_chain.ts # deploy mock scrollchain
npx hardhat --network $layer1 run scripts/deploy_whitelist.ts
npx hardhat --network $layer1 run scripts/deploy_l1_scroll_messenger.ts
env CONTRACT_NAME=L1GatewayRouter npx hardhat run --network $layer1 scripts/deploy_proxy_contract.ts
npx hardhat --network $layer1 run scripts/deploy_l1_standard_erc20_gateway.ts
npx hardhat --network $layer1 run scripts/deploy_l1_custom_erc20_gateway.ts
npx hardhat --network $layer1 run scripts/deploy_l1_erc721_gateway.ts
npx hardhat --network $layer1 run scripts/deploy_l1_erc1155_gateway.ts
npx hardhat --network $layer1 run scripts/deploy_l1_weth_gateway.ts
npx hardhat --network $layer1 run scripts/deploy_l1_eth_gateway.ts
npx hardhat --network $layer1 run scripts/deploy_l1_message_queue.ts
env CONTRACT_NAME=L2GasPriceOracle npx hardhat run --network $layer1 scripts/deploy_proxy_contract.ts
env CONTRACT_NAME=EnforcedTxGateway npx hardhat run --network $layer1 scripts/deploy_proxy_contract.ts
npx hardhat --network $layer1 run scripts/deploy_proxy.ts

# deploy contracts in layer 2
npx hardhat --network $layer2 run scripts/deploy_l2_messenger.ts
env CONTRACT_NAME=L2GatewayRouter npx hardhat run --network $layer2 scripts/deploy_proxy_contract.ts
npx hardhat --network $layer2 run scripts/deploy_l2_standard_erc20_gateway.ts
npx hardhat --network $layer2 run scripts/deploy_l2_eth_gateway.ts
npx hardhat --network $layer2 run scripts/deploy_l2_custom_erc20_gateway.ts
npx hardhat --network $layer2 run scripts/deploy_l2_erc721_gateway.ts
npx hardhat --network $layer2 run scripts/deploy_l2_erc1155_gateway.ts
npx hardhat --network $layer2 run scripts/deploy_l2_weth_gateway.ts
npx hardhat --network $layer2 run scripts/deploy_proxy.ts

# initalize contracts in layer 1, should set proper bash env variables first
npx hardhat --network $layer1 run scripts/initialize_scroll_chain.ts
npx hardhat --network $layer1 run scripts/initialize_l2_gasprice_oracle.ts
npx hardhat --network $layer1 run scripts/initialize_l1_message_queue.ts
npx hardhat --network $layer1 run scripts/initialize_l1_messenger.ts
npx hardhat --network $layer1 run scripts/initialize_l1_gateway_router.ts
npx hardhat --network $layer1 run scripts/initialize_l1_whitelist.ts
npx hardhat --network $layer1 run scripts/initialize_l1_eth_gateway.ts
npx hardhat --network $layer1 run scripts/initialize_l1_weth_gateway.ts
npx hardhat --network $layer1 run scripts/initialize_l1_erc20_gateway.ts
npx hardhat --network $layer1 run scripts/initialize_l1_custom_erc20_gateway.ts
npx hardhat --network $layer1 run scripts/initialize_l1_erc721_gateway.ts
npx hardhat --network $layer1 run scripts/initialize_l1_erc1155_gateway.ts
npx hardhat --network $layer1 run scripts/add_sequencer.ts
npx hardhat --network $layer1 run scripts/add_prover.ts

# initalize contracts in layer 2, should set proper bash env variables first
npx hardhat --network $layer2 run scripts/initialize_l2_predeploys.ts
npx hardhat --network $layer2 run scripts/initialize_l2_messenger.ts
npx hardhat --network $layer2 run scripts/initialize_l2_gateway_router.ts
npx hardhat --network $layer2 run scripts/initialize_l2_token_factory.ts
npx hardhat --network $layer2 run scripts/initialize_l2_eth_gateway.ts
npx hardhat --network $layer2 run scripts/initialize_l2_weth_gateway.ts
npx hardhat --network $layer2 run scripts/initialize_l2_erc20_gateway.ts
npx hardhat --network $layer2 run scripts/initialize_l2_custom_erc20_gateway.ts
npx hardhat --network $layer2 run scripts/initialize_l2_erc721_gateway.ts
npx hardhat --network $layer2 run scripts/initialize_l2_erc1155_gateway.ts

# transfer ownership in layer 1
env CONTRACT_NAME=ProxyAdmin CONTRACT_OWNER=$owner npx hardhat run --network $layer1 scripts/transfer_ownership.ts
env CONTRACT_NAME=L1ScrollMessenger CONTRACT_OWNER=$owner npx hardhat run --network $layer1 scripts/transfer_ownership.ts
env CONTRACT_NAME=ScrollChain CONTRACT_OWNER=$owner npx hardhat run --network $layer1 scripts/transfer_ownership.ts
env CONTRACT_NAME=L1GatewayRouter CONTRACT_OWNER=$owner npx hardhat run --network $layer1 scripts/transfer_ownership.ts
env CONTRACT_NAME=L1ETHGateway CONTRACT_OWNER=$owner npx hardhat run --network $layer1 scripts/transfer_ownership.ts
env CONTRACT_NAME=L1WETHGateway CONTRACT_OWNER=$owner npx hardhat run --network $layer1 scripts/transfer_ownership.ts
env CONTRACT_NAME=L1CustomERC20Gateway CONTRACT_OWNER=$owner npx hardhat run --network $layer1 scripts/transfer_ownership.ts
env CONTRACT_NAME=L1StandardERC20Gateway CONTRACT_OWNER=$owner npx hardhat run --network $layer1 scripts/transfer_ownership.ts
env CONTRACT_NAME=L1ERC721Gateway CONTRACT_OWNER=$owner npx hardhat run --network $layer1 scripts/transfer_ownership.ts
env CONTRACT_NAME=L1ERC1155Gateway CONTRACT_OWNER=$owner npx hardhat run --network $layer1 scripts/transfer_ownership.ts

# transfer ownership in layer 2
env CONTRACT_NAME=ProxyAdmin CONTRACT_OWNER=$owner npx hardhat run --network $layer2 scripts/transfer_ownership.ts
env CONTRACT_NAME=L2ScrollMessenger CONTRACT_OWNER=$owner npx hardhat run --network $layer2 scripts/transfer_ownership.ts
env CONTRACT_NAME=L2GatewayRouter CONTRACT_OWNER=$owner npx hardhat run --network $layer2 scripts/transfer_ownership.ts
env CONTRACT_NAME=L2ETHGateway CONTRACT_OWNER=$owner npx hardhat run --network $layer2 scripts/transfer_ownership.ts
env CONTRACT_NAME=L2WETHGateway CONTRACT_OWNER=$owner npx hardhat run --network $layer2 scripts/transfer_ownership.ts
env CONTRACT_NAME=L2CustomERC20Gateway CONTRACT_OWNER=$owner npx hardhat run --network $layer2 scripts/transfer_ownership.ts
env CONTRACT_NAME=L2StandardERC20Gateway CONTRACT_OWNER=$owner npx hardhat run --network $layer2 scripts/transfer_ownership.ts
env CONTRACT_NAME=L2ERC721Gateway CONTRACT_OWNER=$owner npx hardhat run --network $layer2 scripts/transfer_ownership.ts
env CONTRACT_NAME=L2ERC1155Gateway CONTRACT_OWNER=$owner npx hardhat run --network $layer2 scripts/transfer_ownership.ts
```
