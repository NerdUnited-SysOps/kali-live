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
    # user = userAccounts[4]

    # Manual process: add in the Notarized Message here (will come from an indexer)
    message = [('sender', '0xCE89b5dceF3228EDf20A150A863A24956d3Bb660'), ('bridgedAmount', 1000000000), ('nonce', 0xb6455f198b206a5aeb4592bcf8433515cdc338fc39cdb75b7596b77777240008), ('messageHash', 0xca8069f6e8eedf43970269fdecf70cdecc00a3e492e0d71494c6a1c85275c6ce), ('approvedMessage', 0xf4adf315b270a011ff3698bfde2897ef49615e5f4f36c44706182902c3853bfa55a8248fa76188036de7fdc33e226a3fe1fd7c0868e87a6e133a00e7e868dde51c), ('notarizeMessage', 0x36ecaec8437206acc7875c349a03a46cdbf98071237bba53ceab796351e4ae1c49b7977cfa9d5aa840c15750513e95357d24700ba5947c3c1c0bf5a7fcd6ee601c), ('notary', '0x21c85810F4793F4E31F264C6e5bE5f4416183247')]

    # Getting the ABI to be able to connect to the token contract on L1
    with open("build/contracts/Token.abi") as json_file:
        json_data = json.load(json_file)

    token = Contract.from_abi("Token", "0x2C4dDAEE92F5268e85Fb8BFAA13e3C0f8c36719f", json_data)

    # Getting the ABI to be able to connect to the bridge minter contract on L1
    with open("build/contracts/BridgeMinter.abi") as json_file:
        json_data = json.load(json_file)

    bridge_minter = Contract.from_abi("BridgeMinter", "0x4ba6990eeecc3eF869Af1556Af97e52567a41809", json_data)

    # Check the token balance before bridging completed
    # print(token.balanceOf(user.address))

    # final bridge step, verify signatures
    result = bridge_minter.bridge(message[0][1], message[1][1], message[2][1], message[3][1], message[4][1], message[5][1], {"from":owner})
    
    # Print out a Bridged Event
    print(result.events)

    # Check the token balance has increased
    # print(token.balanceOf(user.address))