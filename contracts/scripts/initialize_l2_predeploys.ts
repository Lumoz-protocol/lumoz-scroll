/* eslint-disable node/no-missing-import */
import * as dotenv from "dotenv";
import * as hre from "hardhat";
import { ethers } from "hardhat";
import { selectAddressFile } from "./utils";

dotenv.config();

async function main() {
  const addressFile = selectAddressFile(hre.network.name);
  const [deployer] = await ethers.getSigners();

  // initialize L2MessageQueue
  const L2MessageQueueAddress = process.env.L2_MESSAGE_QUEUE_ADDR || "0x5300000000000000000000000000000000000000";
  const L2MessageQueue = await ethers.getContractAt("L2MessageQueue", L2MessageQueueAddress, deployer);
  const L2_SCROLL_MESSENGER_PROXY_ADDR = addressFile.get("L2ScrollMessenger.proxy");
  const tx = await L2MessageQueue.initialize(L2_SCROLL_MESSENGER_PROXY_ADDR);
  console.log("initialize L2MessageQueue, hash:", tx.hash);
  const receipt = await tx.wait();
  console.log(`✅ Done, gas used: ${receipt.gasUsed}`);

  // initialize L2TxFeeVault
  const L2TxFeeVaultAddress = process.env.scroll_feevault_address || "0x5300000000000000000000000000000000000005";
  const L2TxFeeVault = await ethers.getContractAt("L2TxFeeVault", L2TxFeeVaultAddress, deployer);
  const tx2 = await L2TxFeeVault.updateMessenger(L2_SCROLL_MESSENGER_PROXY_ADDR);
  console.log("initialize L2TxFeeVault, hash:", tx2.hash);
  const receipt2 = await tx2.wait();
  console.log(`✅ Done, gas used: ${receipt2.gasUsed}`);

  // initialize L1GasPriceOracle
  const L1GasPriceOracleAddress = process.env.L1_GAS_PRICE_ORACLE || "0x5300000000000000000000000000000000000002";
  const L1GasPriceOracle = await ethers.getContractAt("L1GasPriceOracle", L1GasPriceOracleAddress, deployer);
  const L2_WHITELIST_ADDR = process.env.L2_WHITE_LIST || "0x5300000000000000000000000000000000000003";
  const tx3 = await L1GasPriceOracle.updateWhitelist(L2_WHITELIST_ADDR);
  console.log("initialize L1GasPriceOracle, hash:", tx3.hash);
  const receipt3 = await tx3.wait();
  console.log(`✅ Done, gas used: ${receipt3.gasUsed}`);

  // l2Whitelist updateWhitelistStatus
  const L2WhiteList = await ethers.getContractAt("Whitelist", L2_WHITELIST_ADDR, deployer);
  const GasOracleSender = process.env.GAS_PRICE_ORACLE_SENDER || "0x0000000000000000000000000000000000000000";
  const tx4 = await L2WhiteList.updateWhitelistStatus([deployer.address, GasOracleSender], true);
  console.log("L2WhiteList updateWhitelistStatus, hash:", tx4.hash);
  const receipt4 = await tx4.wait();
  console.log(`✅ Done, gas used: ${receipt4.gasUsed}`);
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
