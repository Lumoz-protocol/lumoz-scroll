/* eslint-disable node/no-missing-import */
import * as dotenv from "dotenv";
import * as hre from "hardhat";
import { ethers } from "hardhat";
import { selectAddressFile } from "./utils";

dotenv.config();

async function main() {
  const addressFile = selectAddressFile(hre.network.name);
  const [deployer] = await ethers.getSigners();
  const ScrollChain = await ethers.getContractAt("ScrollChain", addressFile.get("ScrollChain.proxy"), deployer);

  const PROVER_ADDRESS = process.env.PROVER_ADDRESS || deployer.address;
  const tx = await ScrollChain.addProver(PROVER_ADDRESS);
  console.log(" ScrollChain addSequencer, hash:", tx.hash);
  const receipt = await tx.wait();
  console.log(`✅ Done, gas used: ${receipt.gasUsed}`);
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
