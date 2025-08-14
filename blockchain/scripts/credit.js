const hre = require("hardhat");

async function main() {
  const [deployer] = await hre.ethers.getSigners();
  console.log("Using account:", deployer.address);

  // Substitua pelo endereço do seu contrato, que você já tem.
  const CONTRACT_ADDRESS = "0xAcF7391D07a68D10b810123f440A57b80fb67b8b"; 
  
  // Pegue o contrato já implantado
  const PixPrivado = await hre.ethers.getContractFactory("PixPrivado");
  const pix = PixPrivado.attach(CONTRACT_ADDRESS);

  // Endereço da conta que receberá o saldo (sua sender_id)
  const accountToCredit = "0x45DcA3089fE789B9127DD72d90E7e8E0200b6DE3"; 
  const amountToCredit = 1000; // 1000 tokens (ou 10.00 como você definiu 2 decimais)

  console.log(`Crediting ${amountToCredit} to ${accountToCredit}...`);

  // Chama a função 'credit' com a conta de operador
  const tx = await pix.credit(accountToCredit, amountToCredit, "Initial Balance");
  await tx.wait();

  console.log("Credit transaction successful!");
  
  // Opcional: Verifique o saldo para confirmar
  const balance = await pix.balanceOf(accountToCredit);
  console.log(`New balance of ${accountToCredit}:`, balance.toString());
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});