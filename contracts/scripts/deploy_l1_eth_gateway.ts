/* eslint-disable node/no-missing-import */
import * as dotenv from "dotenv";
import * as hre from "hardhat";
import { ethers } from "hardhat";
import { selectAddressFile } from "./utils";

dotenv.config();

async function main() {
  const addressFileL1 = selectAddressFile(hre.network.name);
  const addressFileL2 = selectAddressFile("l2geth");

  const [deployer] = await ethers.getSigners();

  if (!addressFileL1.get("L1ETHGateway.implementation")) {
    console.log(`>> Deploy L1ETHGateway implementation`);
    const ContractImpl = await ethers.getContractFactory("L1ETHGateway", deployer);
    const L2ETHGateway = addressFileL2.get("L2ETHGateway.proxy");
    const L1GatewayRouter = addressFileL1.get("L1GatewayRouter.proxy");
    const L1ScrollMessenger = addressFileL1.get("L1ScrollMessenger.proxy");
    const impl = await ContractImpl.deploy(L2ETHGateway, L1GatewayRouter, L1ScrollMessenger);
    console.log(`>> waiting for transaction: ${impl.deployTransaction.hash}`);
    await impl.deployed();
    console.log(`✅ L1ETHGateway implementation deployed at ${impl.address}`);
    addressFileL1.set(`L1ETHGateway.implementation`, impl.address);
  }

  // Export contract address to testnet.
  console.log(
    `testnet-export: ${addressFileL1.get(`L1ETHGateway.implementation`)};${addressFileL1.get(`L1ETHGateway.proxy`)}`
  );
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
