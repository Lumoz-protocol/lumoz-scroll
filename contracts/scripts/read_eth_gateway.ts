/* eslint-disable node/no-missing-import */
import * as dotenv from "dotenv";
import * as hre from "hardhat";
import { ethers } from "hardhat";
import { selectAddressFile } from "./utils";

dotenv.config();

async function main() {
  const addressFile = selectAddressFile(hre.network.name);

  const signers = await ethers.getSigners();

  const L2ETHGateway = await ethers.getContractAt("L2ETHGateway", addressFile.get("L2ETHGateway.proxy"), signers[0]);

  const messenger = await L2ETHGateway["withdrawETH(uint256,uint256)"](1000, 21000, { value: 10000 });
  console.log("messenger", messenger);
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
