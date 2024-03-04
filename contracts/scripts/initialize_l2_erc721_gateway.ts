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

  const L2ERC721Gateway = await ethers.getContractAt(
    "L2ERC721Gateway",
    addressFile.get("L2ERC721Gateway.proxy"),
    deployer
  );

  const L2_SCROLL_MESSENGER_ADDR = addressFile.get("L2ScrollMessenger.proxy");
  const L1_ERC721_GATEWAY_ADDR = addressFileL1.get("L1ERC721Gateway.proxy");

  const tx = await L2ERC721Gateway.initialize(L1_ERC721_GATEWAY_ADDR, L2_SCROLL_MESSENGER_ADDR);
  console.log("initialize L2ERC721Gateway, hash:", tx.hash);
  const receipt = await tx.wait();
  console.log(`✅ Done, gas used: ${receipt.gasUsed}`);
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
