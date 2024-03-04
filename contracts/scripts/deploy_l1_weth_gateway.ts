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

  if (!addressFileL1.get("L1WETHGateway.implementation")) {
    console.log(`>> Deploy L1WETHGateway implementation`);
    const ContractImpl = await ethers.getContractFactory("L1WETHGateway", deployer);
    const L1WETH = addressFileL1.get("WETH");
    const L2WETH = process.env.L2_WETH_ADDR || "0x5300000000000000000000000000000000000004";
    const L2WETHGateway = addressFileL2.get("L2WETHGateway.proxy");
    const L1GatewayRouter = addressFileL1.get("L1GatewayRouter.proxy");
    const L1ScrollMessenger = addressFileL1.get("L1ScrollMessenger.proxy");
    const impl = await ContractImpl.deploy(L1WETH, L2WETH, L2WETHGateway, L1GatewayRouter, L1ScrollMessenger);
    console.log(`>> waiting for transaction: ${impl.deployTransaction.hash}`);
    await impl.deployed();
    console.log(`✅ L1WETHGateway implementation deployed at ${impl.address}`);
    addressFileL1.set(`L1WETHGateway.implementation`, impl.address);
  }

  // Export contract address to testnet.
  console.log(
    `testnet-export: ${addressFileL1.get(`L1WETHGateway.implementation`)};${addressFileL1.get(`L1WETHGateway.proxy`)}`
  );
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
