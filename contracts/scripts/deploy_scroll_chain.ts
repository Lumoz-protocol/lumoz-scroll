/* eslint-disable node/no-missing-import */
import * as hre from "hardhat";
import { ethers } from "hardhat";
import { selectAddressFile } from "./utils";

async function main() {
  const addressFile = selectAddressFile(hre.network.name);

  const [deployer] = await ethers.getSigners();

  const CHAIN_ID_L2 = process.env.CHAIN_ID_L2 || "none";
  if (!addressFile.get("ScrollChain.verifier")) {
    console.log(">> Deploy ZkEvmVerifierV1");
    const RollupVerifier = await ethers.getContractFactory("ZkEvmVerifierV1", deployer);
    const verifier = await RollupVerifier.deploy(addressFile.get("ScrollChain.plonk_verifier"));
    console.log(`>> waiting for transaction: ${verifier.deployTransaction.hash}`);
    await verifier.deployed();
    console.log(`✅ RollupVerifier deployed at ${verifier.address}`);
    addressFile.set("ScrollChain.verifier", verifier.address);
  }

  if (!addressFile.get("ScrollChain.multiple_verifier")) {
    console.log(">> Deploy MultipleVersionRollupVerifier");
    const multipleRollupVerifier = await ethers.getContractFactory("MultipleVersionRollupVerifier", deployer);
    const verifier = await multipleRollupVerifier.deploy(addressFile.get("ScrollChain.verifier"));
    console.log(`>> waiting for transaction: ${verifier.deployTransaction.hash}`);
    await verifier.deployed();
    console.log(`✅ MultipleVersionRollupVerifier deployed at ${verifier.address}`);
    addressFile.set("ScrollChain.multiple_verifier", verifier.address);
  }

  if (!addressFile.get("ScrollChain.implementation")) {
    console.log(">> Deploy ScrollChain implementation");
    const ScrollChain = await ethers.getContractFactory("ScrollChain", {
      libraries: {},
      signer: deployer,
    });
    const L1MessageQueueAddress = addressFile.get("L1MessageQueue.proxy");
    const verifierAddress = addressFile.get("ScrollChain.multiple_verifier");
    const impl = await ScrollChain.deploy(CHAIN_ID_L2, L1MessageQueueAddress, verifierAddress);
    console.log(`>> waiting for transaction: ${impl.deployTransaction.hash}`);
    await impl.deployed();
    console.log(`✅ ScrollChain implementation deployed at ${impl.address}`);
    addressFile.set("ScrollChain.implementation", impl.address);
  }

  // Export contract address to testnet.
  console.log(
    `testnet-export: 
    ScrollChain.implementation: ${addressFile.get("ScrollChain.implementation")};
    ScrollChain.proxy: ${addressFile.get("ScrollChain.proxy")}`
  );
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
