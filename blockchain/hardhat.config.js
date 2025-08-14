require("@nomicfoundation/hardhat-toolbox");

/** @type import('hardhat/config').HardhatUserConfig */
module.exports = {
  solidity: "0.8.20",
  networks: {
    hardhat: {},
    ganache: {
      url: "http://127.0.0.1:7545", // Sua BLOCKCHAIN_URL
      chainId: 1337, // Adicionado o Chain ID do Ganache
    }
  },
};