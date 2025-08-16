// Em scripts/interact.js

const hre = require("hardhat");

async function main() {
  const contractAddress = "0xAcF7391D07a68D10b810123f440A57b80fb67b8b"; // O endereço do seu contrato
  
  // <<< ATUALIZADO com o endereço REAL do seu Ganache atual >>>
  const recipientAddress = "0xaB09b59B7E38D237bf909FcA29Ee8AFC46fDd2c0";
  
  const amount = hre.ethers.utils.parseUnits("100", 18); 
  const memo = "Teste final com tudo sincronizado!";

  console.log(`Anexando ao contrato no endereço: ${contractAddress}`);
  const pixPrivado = await hre.ethers.getContractAt("PixPrivado", contractAddress);

  console.log(`Creditando ${amount.toString()} tokens para o endereço ${recipientAddress}...`);
  
  const tx = await pixPrivado.credit(recipientAddress, amount, memo);
  
  await tx.wait(); 

  console.log("Transação concluída! Hash:", tx.hash);
  console.log("Agora, verifique o terminal da sua aplicação Go!");
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});