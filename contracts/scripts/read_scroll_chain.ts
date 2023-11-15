/* eslint-disable node/no-missing-import */
import * as dotenv from "dotenv";
import * as hre from "hardhat";
import { ethers } from "hardhat";
import { selectAddressFile } from "./utils";

dotenv.config();

async function main() {
  const addressFile = selectAddressFile(hre.network.name);

  const signers = await ethers.getSigners();
  const prover = signers[3];

  const ScrollChain = await ethers.getContractAt("ScrollChain", addressFile.get("ScrollChain.proxy"), prover);

  // check if commit proof hash for a batch is allowed
  const batchIndex = 1;
  try {
    await ScrollChain.isCommitProofHashAllowed(batchIndex);
    console.log("batchIndex:", batchIndex, "commit proof hash allowed");
  } catch (e) {
    console.log("batchIndex", batchIndex, "commit proof hash not allowed");
  }

  // fetch provable batch index
  const step = 20;
  const batchToProve = await ScrollChain.getBatchToProve(step);
  console.log("batchToProve batchIndex:", batchToProve);

  // fetch ProofHashCommitEpoch
  const lastFinalizedBatchIndex = await ScrollChain.lastFinalizedBatchIndex();
  console.log("lastFinalizedBatchIndex:", lastFinalizedBatchIndex);

  // fetch ProofHashCommitEpoch
  const proofHashCommitEpoch = await ScrollChain.proofHashCommitEpoch();
  console.log("proofHashCommitEpoch:", proofHashCommitEpoch);

  // fetch ProofCommitEpoch
  const proofCommitEpoch = await ScrollChain.proofCommitEpoch();
  console.log("proofCommitEpoch:", proofCommitEpoch);

  // fetch IDEposit address
  const ideDeposit = await ScrollChain.ideDeposit();
  console.log("ideDeposit:", ideDeposit);

  // fetch slotAdapter address
  const slotAdapter = await ScrollChain.slotAdapter();
  console.log("slotAdapter:", slotAdapter);

  // fetch slotAdapter address
  const proverAddress = prover.address;
  const isProver = await ScrollChain.isProver(proverAddress);
  console.log("isProver:", isProver);

  // fetch slotAdapter address
  const proverCommitProofHashBatchIndex = "0x3e";
  const proverCommitProofHash = await ScrollChain.proverCommitProofHash(proverCommitProofHashBatchIndex, proverAddress);
  console.log("proverCommitProofHash:", proverCommitProofHash.blockNumber.toString());

  // fetch slotAdapter address
  const finalizedStateRootsBatchIndex = "0x3d";
  const finalizedStateRoots = await ScrollChain.finalizedStateRoots(finalizedStateRootsBatchIndex);
  console.log("finalizedStateRoots:", finalizedStateRoots);
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
