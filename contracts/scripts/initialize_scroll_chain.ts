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
  const sequencerAddress = process.env.SEQUENCER_ADDRESS || "0x0000000000000000000000000000000000000000";
  const proofHashCommitEpoch = process.env.PROOf_HASH_COMMIT_EPOCH || 32;
  const proofCommitEpoch = process.env.PROOf_COMMIT_EPOCH || 20;
  const minDeposit = process.env.MIN_DEPOSIT || 0;
  const noProofPunishAmount = process.env.NO_PROOF_PUNISH_AMOUNT || 0;
  const incorrectProofPunishAmount = process.env.INCORRECT_PROOF_PUNISH_AMOUNT || 0;
  const slotAdapter = process.env.SLOT_ADAPTER_ADDR || "0x0000000000000000000000000000000000000000";
  const ideDeposit = process.env.MINER_DEPOSIT_ADDRESS || "0x0000000000000000000000000000000000000000";

  const tx = await ScrollChain.initialize(
    L1MessageQueueAddress,
    verifierAddress,
    maxNumTxInChunk,
    sequencerAddress,
    proofHashCommitEpoch,
    proofCommitEpoch,
    minDeposit,
    noProofPunishAmount,
    incorrectProofPunishAmount,
    slotAdapter,
    ideDeposit
  );
  console.log("initialize ScrollChain, hash:", tx.hash);
  const receipt = await tx.wait();
  console.log(`✅ Done, gas used: ${receipt.gasUsed}`);
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
