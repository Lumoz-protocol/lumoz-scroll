/* eslint-disable node/no-missing-import */
import * as dotenv from "dotenv";
import * as hre from "hardhat";
import { ethers } from "hardhat";
import { selectAddressFile } from "./utils";

dotenv.config();

async function main() {
  const addressFileL1 = selectAddressFile(hre.network.name);

  const [deployer] = await ethers.getSigners();

  if (!addressFileL1.get("L1MessageQueue.implementation")) {
    console.log(`>> Deploy L1MessageQueue implementation`);
    const ContractImpl = await ethers.getContractFactory("L1MessageQueue", deployer);
    const L1ScrollMessenger = addressFileL1.get("L1ScrollMessenger.proxy");
    const ScrollChain = addressFileL1.get("ScrollChain.proxy");
    const EnforcedTxGateway = addressFileL1.get("EnforcedTxGateway.proxy");
    const impl = await ContractImpl.deploy(L1ScrollMessenger, ScrollChain, EnforcedTxGateway);
    console.log(`>> waiting for transaction: ${impl.deployTransaction.hash}`);
    await impl.deployed();
    console.log(`✅ L1MessageQueue implementation deployed at ${impl.address}`);
    addressFileL1.set(`L1MessageQueue.implementation`, impl.address);
  }

  // Export contract address to testnet.
  console.log(
    `testnet-export: ${addressFileL1.get(`L1MessageQueue.implementation`)};${addressFileL1.get(`L1MessageQueue.proxy`)}`
  );
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
