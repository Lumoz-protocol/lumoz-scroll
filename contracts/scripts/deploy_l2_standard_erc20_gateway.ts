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

  if (!addressFileL2.get("L2StandardERC20Gateway.implementation")) {
    console.log(`>> Deploy L2StandardERC20Gateway implementation`);
    const ContractImpl = await ethers.getContractFactory("L2StandardERC20Gateway", deployer);
    const L1StandardERC20Gateway = addressFileL1.get("L1StandardERC20Gateway.proxy");
    const L2GatewayRouter = addressFileL2.get("L2GatewayRouter.proxy");
    const L2ScrollMessenger = addressFileL2.get("L2ScrollMessenger.proxy");
    const ScrollStandardERC20Factory = addressFileL2.get("ScrollStandardERC20Factory");
    const impl = await ContractImpl.deploy(
      L1StandardERC20Gateway,
      L2GatewayRouter,
      L2ScrollMessenger,
      ScrollStandardERC20Factory
    );
    console.log(`>> waiting for transaction: ${impl.deployTransaction.hash}`);
    await impl.deployed();
    console.log(`✅ L2StandardERC20Gateway implementation deployed at ${impl.address}`);
    addressFileL2.set(`L2StandardERC20Gateway.implementation`, impl.address);
  }

  // Export contract address to testnet.
  console.log(
    `testnet-export: ${addressFileL2.get(`L2StandardERC20Gateway.implementation`)};${addressFileL2.get(
      `L2StandardERC20Gateway.proxy`
    )}`
  );
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
