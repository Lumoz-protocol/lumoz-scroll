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

  const verifierAddress = addressFile.get("ScrollChain.multiple_verifier");
  const L1MessageQueueAddress = addressFile.get("L1MessageQueue.proxy");
  const maxNumTxInChunk = 100;

  const tx = await ScrollChain.initialize(L1MessageQueueAddress, verifierAddress, maxNumTxInChunk);
  console.log("initialize ScrollChain, hash:", tx.hash);
  const receipt = await tx.wait();
  console.log(`✅ Done, gas used: ${receipt.gasUsed}`);

  const sequencerAddress = process.env.SEQUENCER_ADDRESS || "0x0000000000000000000000000000000000000000";
  const tx2 = await ScrollChain.addSequencer(sequencerAddress);
  console.log("initialize ScrollChain addSequencer, hash:", tx2.hash);
  const receipt2 = await tx2.wait();
  console.log(`✅ Done, gas used: ${receipt2.gasUsed}`);

  const proverAddress = process.env.PROVER_ADDRESS || "0x0000000000000000000000000000000000000000";
  const tx3 = await ScrollChain.addProver(proverAddress);
  console.log("initialize ScrollChain addProver, hash:", tx3.hash);
  const receipt3 = await tx3.wait();
  console.log(`✅ Done, gas used: ${receipt3.gasUsed}`);

  const proofHashCommitEpoch = process.env.PROOf_HASH_COMMIT_EPOCH || 32;
  const tx4 = await ScrollChain.setProofHashCommitEpoch(proofHashCommitEpoch);
  console.log("initialize ScrollChain setProofHashCommitEpoch, hash:", tx4.hash);
  const receipt4 = await tx4.wait();
  console.log(`✅ Done, gas used: ${receipt4.gasUsed}`);

  const proofCommitEpoch = process.env.PROOf_COMMIT_EPOCH || 20;
  const tx5 = await ScrollChain.setProofCommitEpoch(proofCommitEpoch);
  console.log("initialize ScrollChain setProofHashCommitEpoch, hash:", tx5.hash);
  const receipt5 = await tx5.wait();
  console.log(`✅ Done, gas used: ${receipt5.gasUsed}`);

  const noProofPunishAmount = process.env.NO_PROOF_PUNISH_AMOUNT || 0;
  const tx6 = await ScrollChain.setNoProofPunishAmount(noProofPunishAmount);
  console.log("initialize ScrollChain setNoProofPunishAmount, hash:", tx6.hash);
  const receipt6 = await tx6.wait();
  console.log(`✅ Done, gas used: ${receipt6.gasUsed}`);

  const incorrectProofPunishAmount = process.env.INCORRECT_PROOF_PUNISH_AMOUNT || 0;
  const tx7 = await ScrollChain.setIncorrectProofPunishAmount(incorrectProofPunishAmount);
  console.log("initialize ScrollChain setIncorrectProofPunishAmount, hash:", tx7.hash);
  const receipt7 = await tx7.wait();
  console.log(`✅ Done, gas used: ${receipt7.gasUsed}`);
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
