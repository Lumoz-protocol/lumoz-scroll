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

  if (!addressFileL2.get("L2CustomERC20Gateway.implementation")) {
    console.log(`>> Deploy L2CustomERC20Gateway implementation`);
    const ContractImpl = await ethers.getContractFactory("L2CustomERC20Gateway", deployer);
    const L1CustomERC20Gateway = addressFileL1.get("L1CustomERC20Gateway.proxy");
    const L2GatewayRouter = addressFileL2.get("L2GatewayRouter.proxy");
    const L2ScrollMessenger = addressFileL2.get("L2ScrollMessenger.proxy");
    const impl = await ContractImpl.deploy(L1CustomERC20Gateway, L2GatewayRouter, L2ScrollMessenger);
    console.log(`>> waiting for transaction: ${impl.deployTransaction.hash}`);
    await impl.deployed();
    console.log(`✅ L2CustomERC20Gateway implementation deployed at ${impl.address}`);
    addressFileL2.set(`L2CustomERC20Gateway.implementation`, impl.address);
  }

  // Export contract address to testnet.
  console.log(
    `testnet-export: ${addressFileL1.get(`L2CustomERC20Gateway.implementation`)};${addressFileL1.get(
      `L2CustomERC20Gateway.proxy`
    )}`
  );
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
