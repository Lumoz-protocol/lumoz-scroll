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

  if (!addressFileL1.get("L1ScrollMessenger.implementation")) {
    console.log(`>> Deploy L1ScrollMessenger implementation`);
    const ContractImpl = await ethers.getContractFactory("L1ScrollMessenger", deployer);
    const L2_SCROLL_MESSENGER_PROXY_ADDR = addressFileL2.get("L2ScrollMessenger.proxy");
    const L1_SCROLL_CHAIN_PROXY_ADDR = addressFileL1.get("ScrollChain.proxy");
    const L1_MESSAGE_QUEUE_PROXY_ADDR = addressFileL1.get("L1MessageQueue.proxy");
    const impl = await ContractImpl.deploy(
      L2_SCROLL_MESSENGER_PROXY_ADDR,
      L1_SCROLL_CHAIN_PROXY_ADDR,
      L1_MESSAGE_QUEUE_PROXY_ADDR
    );
    console.log(`>> waiting for transaction: ${impl.deployTransaction.hash}`);
    await impl.deployed();
    console.log(`✅ L1ScrollMessenger implementation deployed at ${impl.address}`);
    addressFileL1.set(`L1ScrollMessenger.implementation`, impl.address);
  }

  // Export contract address to testnet.
  console.log(
    `testnet-export: ${addressFileL1.get(`L1ScrollMessenger.implementation`)};${addressFileL1.get(
      `L1ScrollMessenger.proxy`
    )}`
  );
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
