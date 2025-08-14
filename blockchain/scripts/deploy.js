const hre = require("hardhat");

async function main() {
  const [deployer] = await hre.ethers.getSigners();
  console.log("Deploying contract with:", deployer.address);

  const PixPrivado = await hre.ethers.getContractFactory("PixPrivado");
  const pix = await PixPrivado.deploy(deployer.address);

  await pix.deployed();
  console.log("PixPrivado deployed to:", pix.address);
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
