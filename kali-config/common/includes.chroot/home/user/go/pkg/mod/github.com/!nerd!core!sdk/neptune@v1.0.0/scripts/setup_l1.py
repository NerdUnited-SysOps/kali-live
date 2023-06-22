from web3.auto import w3
from eth_account.messages import encode_defunct, defunct_hash_message
from brownie import Token, Bridge, BridgeMinter, Contract
from eth_abi import encode
from web3 import Web3
from scripts import accounts
from build import contracts
import json

def main():
    userAccounts = accounts.main()

    owner = userAccounts[0]
    approver = userAccounts[1]
    notary = userAccounts[2]
    feeReceiver = userAccounts[3]
    user = userAccounts[4]

    bridge_minter_address = "0x2140CF34D891D97D543360F3da67dc4087e70A47"
    token_address = "0xB0D98244FBf8c41b0A34A0B3845C4e2a131BE3a0"

    with open("build/contracts/Token.abi") as json_file:
        json_data = json.load(json_file)

    l1_token = Contract.from_abi("Token", token_address, json_data)
    l1_token.setIssuer(bridge_minter_address, {'from': owner})