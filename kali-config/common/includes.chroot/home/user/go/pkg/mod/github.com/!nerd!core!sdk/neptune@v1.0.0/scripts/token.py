#
# Usage:
#
#       % brownie console
#       >>> from scripts import bridge
#       >>> bb = bridge.main()
#       >>> bb = bridge.transfer()
#

from web3.auto import w3
from eth_account.messages import encode_defunct, defunct_hash_message
from brownie import accounts, Token
from eth_abi import encode
from web3 import Web3

owner = accounts[0]
approver = accounts[2]
notary = accounts[4]
feeReceiver = accounts[15]

def main():
    name = "Elevate Bridged Token"
    symbol = "ELVb"
    decimals = 8
    tokenMaxSupply = 5*10**decimals*10**9
    print(5*10**decimals*10**9)

    # deploy contract
    token = Token.deploy(name, symbol, decimals, owner.address, owner.address, tokenMaxSupply, {'from': owner})

    return token