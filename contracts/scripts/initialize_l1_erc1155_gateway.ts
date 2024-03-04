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

  const L1ERC1155Gateway = await ethers.getContractAt(
    "L1ERC1155Gateway",
    addressFile.get("L1ERC1155Gateway.proxy"),
    deployer
  );

  const L1_SCROLL_MESSENGER_ADDR = addressFile.get("L1ScrollMessenger.proxy");
  const L2_ERC1155_GATEWAY_ADDR = addressFileL2.get("L2ERC1155Gateway.proxy");

  const tx = await L1ERC1155Gateway.initialize(L2_ERC1155_GATEWAY_ADDR, L1_SCROLL_MESSENGER_ADDR);
  console.log("initialize L1ERC1155Gateway, hash:", tx.hash);
  const receipt = await tx.wait();
  console.log(`✅ Done, gas used: ${receipt.gasUsed}`);
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
