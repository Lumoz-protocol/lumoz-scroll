/* eslint-disable node/no-missing-import */
import * as hre from "hardhat";
import { ethers } from "hardhat";
import { selectAddressFile } from "./utils";

const l1Contracts = [
  "L1MessageQueue",
  "L2GasPriceOracle",
  "L1ScrollMessenger",
  "L1GatewayRouter",
  "L1StandardERC20Gateway",
  "L1WETHGateway",
  "ScrollChain",
  "L1CustomERC20Gateway",
  "L1ERC721Gateway",
  "L1ERC1155Gateway",
  "L1ETHGateway",
  "EnforcedTxGateway",
];
const l2Contracts = [
  "L2ScrollMessenger",
  "L2GatewayRouter",
  "L2StandardERC20Gateway",
  "L2WETHGateway",
  "L2CustomERC20Gateway",
  "L2ERC721Gateway",
  "L2ERC1155Gateway",
  "L2ETHGateway",
];

async function main() {
  const addressFile = selectAddressFile(hre.network.name);

  const [, proxyDeployer] = await ethers.getSigners();

  let contracts;
  if (hre.network.name === "l1geth") {
    contracts = l1Contracts;
  } else {
    contracts = l2Contracts;
  }

  let transactionCount = await proxyDeployer.getTransactionCount();

  for (const proxyContract of contracts) {
    const contractAddress = ethers.utils.getContractAddress({
      from: proxyDeployer.address,
      nonce: transactionCount,
    });
    transactionCount++;
    addressFile.set(`${proxyContract}.proxy`, contractAddress);
  }

  console.log("pre-computed all proxy addresses, and set to address file");
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
