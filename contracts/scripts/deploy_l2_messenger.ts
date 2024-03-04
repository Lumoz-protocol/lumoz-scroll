/* eslint-disable node/no-missing-import */
import * as dotenv from "dotenv";
import * as hre from "hardhat";
import { ethers } from "hardhat";
import { selectAddressFile } from "./utils";

dotenv.config();

async function main() {
  const addressFileL1 = selectAddressFile("l1geth");
  const addressFileL2 = selectAddressFile(hre.network.name);

  const [deployer] = await ethers.getSigners();

  if (!addressFileL2.get("L2ScrollMessenger.implementation")) {
    console.log(`>> Deploy L2ScrollMessenger implementation`);
    const ContractImpl = await ethers.getContractFactory("L2ScrollMessenger", deployer);
    const L1ScrollMessenger = addressFileL1.get("L1ScrollMessenger.proxy");
    const L2_MESSAGE_QUEUE = process.env.L2_MESSAGE_QUEUE_ADDR || "0x5300000000000000000000000000000000000000";
    const impl = await ContractImpl.deploy(L1ScrollMessenger, L2_MESSAGE_QUEUE);
    console.log(`>> waiting for transaction: ${impl.deployTransaction.hash}`);
    await impl.deployed();
    console.log(`✅ L2ScrollMessenger implementation deployed at ${impl.address}`);
    addressFileL2.set(`L2ScrollMessenger.implementation`, impl.address);
  }

  // Export contract address to testnet.
  console.log(
    `testnet-export: ${addressFileL1.get(`L2ScrollMessenger.implementation`)};${addressFileL1.get(
      `L2ScrollMessenger.proxy`
    )}`
  );
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
