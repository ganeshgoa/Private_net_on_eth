// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0; 
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
contract Escrow {
    address public seller;
    address public buyer;
    bool private accepted;
    address public spender;
    event Released(address, uint);
    constructor () payable {
        seller = 0x1eC5259ed79a8C54d8F8D7BC4f64D40415BDbE57;
        payable(address(this)).transfer(msg.value);
        accepted = false;
        spender = 0xAf3789e113806F795991B508BE455dE3C285F2C2;
    } 
    function getBalance() external view returns(uint) {
    return address(this).balance;    
    }
    function increaseAllowance(address spender, uint256 addedValue) public virtual returns (bool) {
        _approve(seller, spender, allowance(owner, spender) + 1);
        return true;
    }

    function transferToMe(address _owner, address _token, uint amount) public {
        ERC20(_token).transferFrom(_owner, address(this), amount);
    }

    function accept() public {
        buyer = msg.sender;
        accepted = true;
    }
    function release() public payable {
        require(seller == msg.sender, "Only seller can do that");
        require(accepted == true, "Only accepted");
        payable(buyer).transfer(address(this).balance);
        emit Released(buyer, address(this).balance);
    }
    function cancel () public payable {
        require(seller == msg.sender, "Only seller can do that");
        require(address(this).balance != 0, "Balance is 0");
        payable(seller).transfer(address(this).balance);

    }
       
}
