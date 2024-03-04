/* eslint-disable node/no-missing-import */
import * as dotenv from "dotenv";

import * as hre from "hardhat";
import { ethers } from "hardhat";
import { selectAddressFile } from "./utils";

dotenv.config();

async function main() {
  const addressFile = selectAddressFile(hre.network.name);
  const addressFileL2 = selectAddressFile("l2geth");

  const [deployer] = await ethers.getSigners();

  const L1StandardERC20Gateway = await ethers.getContractAt(
    "L1StandardERC20Gateway",
    addressFile.get("L1StandardERC20Gateway.proxy"),
    deployer
  );

  const L1_STANDARD_ERC20_GATEWAY_ADDR = addressFileL2.get("L2StandardERC20Gateway.proxy");
  const L1_GATEWAY_ROUTER_PROXY_ADDR = addressFile.get("L1GatewayRouter.proxy");
  const L1_SCROLL_MESSENGER_PROXY_ADDR = addressFile.get("L1ScrollMessenger.proxy");
  const SCROLL_STANDARD_ERC20 = addressFileL2.get("ScrollStandardERC20");
  const SCROLL_STANDARD_ERC20_FACTORY_ADDR = addressFileL2.get("ScrollStandardERC20Factory");
  const tx = await L1StandardERC20Gateway.initialize(
    L1_STANDARD_ERC20_GATEWAY_ADDR,
    L1_GATEWAY_ROUTER_PROXY_ADDR,
    L1_SCROLL_MESSENGER_PROXY_ADDR,
    SCROLL_STANDARD_ERC20,
    SCROLL_STANDARD_ERC20_FACTORY_ADDR
  );
  console.log("initialize L1ETHGateway, hash:", tx.hash);
  const receipt = await tx.wait();
  console.log(`✅ Done, gas used: ${receipt.gasUsed}`);
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
