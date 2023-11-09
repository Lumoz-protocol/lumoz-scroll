/* eslint-disable node/no-missing-import */
import * as dotenv from "dotenv";
import * as hre from "hardhat";
import { ethers } from "hardhat";
import { selectAddressFile } from "./utils";

dotenv.config();

async function main() {
  const addressFile = selectAddressFile(hre.network.name);

  const signers = await ethers.getSigners();
  const prover = signers[2];

  const ScrollChain = await ethers.getContractAt("ScrollChain", addressFile.get("ScrollChain.proxy"), prover);

  // check if commit proof hash for a batch is allowed
  try {
    const batchIndex = 1;
    const isCommitProofHashAllowed = await ScrollChain.isCommitProofHashAllowed(batchIndex);
    console.log("isCommitProofHashAllowed batchIndex:", batchIndex, "result:", isCommitProofHashAllowed);
  } catch (e) {
    console.log(e);
  }

  // fetch provable batch index
  const fromIndex = 10;
  const toIndex = 0;
  const batchToProve = await ScrollChain.getBatchToProve(fromIndex, toIndex);
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
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
