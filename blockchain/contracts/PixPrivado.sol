// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/access/AccessControl.sol";

contract PixPrivado is AccessControl {
    bytes32 public constant OPERATOR_ROLE = keccak256("OPERATOR_ROLE");
    uint8 public decimals = 2; // centavos

    mapping(address => uint256) private _balances;
    mapping(address => bool) public frozen;

    event Transfer(address indexed from, address indexed to, uint256 amount, string memo);
    event Credit(address indexed to, uint256 amount, string reason);
    event Debit(address indexed from, uint256 amount, string reason);
    event Frozen(address indexed account, bool isFrozen);

   constructor(address admin) {
    _grantRole(DEFAULT_ADMIN_ROLE, admin);
    _grantRole(OPERATOR_ROLE, admin);
}

    modifier notFrozen(address acc) {
        require(!frozen[acc], "Account frozen");
        _;
    }

    function credit(address to, uint256 amount, string calldata reason)
        external
        onlyRole(OPERATOR_ROLE)
    {
        _balances[to] += amount;
        emit Credit(to, amount, reason);
    }

    function debit(address from, uint256 amount, string calldata reason)
        external
        onlyRole(OPERATOR_ROLE)
    {
        require(_balances[from] >= amount, "Insufficient balance");
        _balances[from] -= amount;
        emit Debit(from, amount, reason);
    }

    function transfer(address to, uint256 amount, string calldata memo)
        external
        notFrozen(msg.sender)
        notFrozen(to)
    {
        require(_balances[msg.sender] >= amount, "Insufficient balance");
        _balances[msg.sender] -= amount;
        _balances[to] += amount;
        emit Transfer(msg.sender, to, amount, memo);
    }

    function balanceOf(address acc) external view returns (uint256) {
        return _balances[acc];
    }

    function setFrozen(address acc, bool f)
        external
        onlyRole(DEFAULT_ADMIN_ROLE)
    {
        frozen[acc] = f;
        emit Frozen(acc, f);
    }
}