/* eslint-disable node/no-missing-import */
import * as dotenv from "dotenv";

import * as hre from "hardhat";
import { ethers } from "hardhat";
import { selectAddressFile } from "./utils";

dotenv.config();

async function main() {
  const addressFileL1 = selectAddressFile("l1geth");
  const addressFile = selectAddressFile(hre.network.name);

  const [deployer] = await ethers.getSigners();

  const L2CustomERC20Gateway = await ethers.getContractAt(
    "L2CustomERC20Gateway",
    addressFile.get("L2CustomERC20Gateway.proxy"),
    deployer
  );

  const L2_GATEWAY_ROUTER_ADDR = addressFile.get("L2GatewayRouter.proxy");
  const L2_SCROLL_MESSENGER_ADDR = addressFile.get("L2ScrollMessenger.proxy");
  const L1_CUSTOM_ERC20_GATEWAY_ADDR = addressFileL1.get("L1CustomERC20Gateway.proxy");

  const tx = await L2CustomERC20Gateway.initialize(
    L1_CUSTOM_ERC20_GATEWAY_ADDR,
    L2_GATEWAY_ROUTER_ADDR,
    L2_SCROLL_MESSENGER_ADDR
  );
  console.log("initialize L2CustomERC20Gateway, hash:", tx.hash);
  const receipt = await tx.wait();
  console.log(`✅ Done, gas used: ${receipt.gasUsed}`);
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
