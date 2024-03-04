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

  if (!addressFileL2.get("L2ETHGateway.implementation")) {
    console.log(`>> Deploy L2ETHGateway implementation`);
    const ContractImpl = await ethers.getContractFactory("L2ETHGateway", deployer);
    const L1ETHGateway = addressFileL1.get("L1ETHGateway.proxy");
    const L2GatewayRouter = addressFileL2.get("L2GatewayRouter.proxy");
    const L2ScrollMessenger = addressFileL2.get("L2ScrollMessenger.proxy");
    const impl = await ContractImpl.deploy(L1ETHGateway, L2GatewayRouter, L2ScrollMessenger);
    console.log(`>> waiting for transaction: ${impl.deployTransaction.hash}`);
    await impl.deployed();
    console.log(`✅ L2ETHGateway implementation deployed at ${impl.address}`);
    addressFileL2.set(`L2ETHGateway.implementation`, impl.address);
  }

  // Export contract address to testnet.
  console.log(
    `testnet-export: ${addressFileL1.get(`L2ETHGateway.implementation`)};${addressFileL1.get(`L2ETHGateway.proxy`)}`
  );
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
