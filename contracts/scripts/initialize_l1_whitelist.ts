/* eslint-disable node/no-missing-import */
import * as dotenv from "dotenv";
import * as hre from "hardhat";
import { ethers } from "hardhat";
import { selectAddressFile } from "./utils";

dotenv.config();

async function main() {
  const addressFile = selectAddressFile(hre.network.name);
  const [deployer] = await ethers.getSigners();

  // L1 Whitelist updateWhitelistStatus
  const L1_WHITELIST_ADDR = addressFile.get("Whitelist");
  const L1WhiteList = await ethers.getContractAt("Whitelist", L1_WHITELIST_ADDR, deployer);
  const GasOracleSender = process.env.GAS_PRICE_ORACLE_SENDER || "0x0000000000000000000000000000000000000000";
  const tx = await L1WhiteList.updateWhitelistStatus([deployer.address, GasOracleSender], true);
  console.log("L1WhiteList updateWhitelistStatus, hash:", tx.hash);
  const receipt = await tx.wait();
  console.log(`✅ Done, gas used: ${receipt.gasUsed}`);
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
