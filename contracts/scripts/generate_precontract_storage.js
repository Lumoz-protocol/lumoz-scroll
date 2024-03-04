const fs = require("fs");
const { ethers } = require("hardhat");

async function main() {
  // Load deployer
  const currentProvider = new ethers.providers.FallbackProvider([ethers.provider], 1);
  const deployer = new ethers.Wallet(process.env.L2_IMPL_DEPLOYER_PRIVATE_KEY || "", currentProvider);
  console.log(await deployer.getAddress());

  const L1GasPriceOracleFactory = await ethers.getContractFactory("L1GasPriceOracle", deployer);
  let contract_ = await L1GasPriceOracleFactory.deploy(deployer.address);

  console.log(`>> waiting for transaction: ${contract_.deployTransaction.hash}`);
  await contract_.deployed();
  const L1GasPriceOracleImplementationAddress = contract_.address;

  const L2MessageQueueFactory = await ethers.getContractFactory("L2MessageQueue", deployer);
  contract_ = await L2MessageQueueFactory.deploy(deployer.address);

  console.log(`>> waiting for transaction: ${contract_.deployTransaction.hash}`);
  await contract_.deployed();
  const L2MessageQueueImplementationAddress = contract_.address;

  const L2TxFeeVaultFactory = await ethers.getContractFactory("L2TxFeeVault", deployer);
  contract_ = await L2TxFeeVaultFactory.deploy(deployer.address, deployer.address, 0);

  console.log(`>> waiting for transaction: ${contract_.deployTransaction.hash}`);
  await contract_.deployed();
  const L2TxFeeVaultmplementationAddress = contract_.address;

  const WhitelistFactory = await ethers.getContractFactory("Whitelist", deployer);
  contract_ = await WhitelistFactory.deploy(deployer.address);

  console.log(`>> waiting for transaction: ${contract_.deployTransaction.hash}`);
  await contract_.deployed();
  const WhitelistImplementationAddress = contract_.address;

  const L1BlockContainerFactory = await ethers.getContractFactory("L1BlockContainer", deployer);
  contract_ = await L1BlockContainerFactory.deploy(deployer.address);

  console.log(`>> waiting for transaction: ${contract_.deployTransaction.hash}`);
  await contract_.deployed();
  const L1BlockContainerImplementationAddress = contract_.address;

  const WrappedEtherFactory = await ethers.getContractFactory("WrappedEther", deployer);
  contract_ = await WrappedEtherFactory.deploy();

  console.log(`>> waiting for transaction: ${contract_.deployTransaction.hash}`);
  await contract_.deployed();
  const WrappedEtherImplementationAddress = contract_.address;

  const L2MessageQueueInfo = await getAddressInfo(L2MessageQueueImplementationAddress);
  const L1BlockContainerInfo = await getAddressInfo(L1BlockContainerImplementationAddress);
  const L1GasPriceOracleInfo = await getAddressInfo(L1GasPriceOracleImplementationAddress);
  const WhitelistInfo = await getAddressInfo(WhitelistImplementationAddress);
  const WETHInfo = await getAddressInfo(WrappedEtherImplementationAddress);
  const L2TxFeeVaultInfo = await getAddressInfo(L2TxFeeVaultmplementationAddress);

  const alloc = {
    "5300000000000000000000000000000000000000": {
      balance: "0x0",
      code: L2MessageQueueInfo.bytecode,
      storage: L2MessageQueueInfo.storage,
    },
    "0x5300000000000000000000000000000000000001": {
      balance: "0x0",
      code: L1BlockContainerInfo.bytecode,
      storage: L1BlockContainerInfo.storage,
    },
    "0x5300000000000000000000000000000000000002": {
      balance: "0x0",
      code: L1GasPriceOracleInfo.bytecode,
      storage: L1GasPriceOracleInfo.storage,
    },
    "0x5300000000000000000000000000000000000003": {
      balance: "0x0",
      code: WhitelistInfo.bytecode,
      storage: WhitelistInfo.storage,
    },
    "0x5300000000000000000000000000000000000004": {
      balance: "0x0",
      code: WETHInfo.bytecode,
      storage: WETHInfo.storage,
    },
    "0x5300000000000000000000000000000000000005": {
      balance: "0x0",
      code: L2TxFeeVaultInfo.bytecode,
      storage: L2TxFeeVaultInfo.storage,
    },
  };

  fs.writeFileSync("./genesis.json", JSON.stringify({ alloc: alloc }, null, 4));
}

async function getAddressInfo(address) {
  const nonce = await ethers.provider.getTransactionCount(address);
  const bytecode = await ethers.provider.getCode(address);

  const storage = {};
  for (let i = 0; i < 120; i++) {
    const storageValue = await ethers.provider.getStorageAt(address, i);
    if (storageValue !== "0x0000000000000000000000000000000000000000000000000000000000000000") {
      storage[ethers.utils.hexZeroPad(ethers.utils.hexlify(i), 32)] = storageValue;
    }
  }

  return { nonce, bytecode, storage };
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
