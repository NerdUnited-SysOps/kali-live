from web3.auto import w3
from eth_account.messages import encode_defunct, defunct_hash_message
from brownie import Token, Bridge, BridgeMinter, Contract
from eth_abi import encode
from web3 import Web3
from scripts import accounts
import json

def main():
    userAccounts = accounts.main()

    owner = userAccounts[0]
    approver = userAccounts[1]
    notary = userAccounts[2]
    feeReceiver = userAccounts[3]
    # user = userAccounts[4]

    # This is what the indexer and signer will be looking at
    # def sign_message_approver(transaction):
    def sign_message_approver():
        # nonce = transaction.events["Received"].items()[0][1]
        # sender = transaction.events["Received"].items()[1][1]
        # value = transaction.events["Received"].items()[2][1]
        nonce = "0x9b7f1863b19652310d16a35b950b6dfefc26efa4efd2ff91037348e57196b997"
        sender="0x2C1C7F524c789c26caF6F12D681eaeb51EBD023d"
        value=100000000
        msg = Web3.solidityKeccak(['bytes32','address','uint256'],[nonce, sender, value])
        private_key = approver.private_key
        signature = approver.address
        msg_ready = encode_defunct(hexstr=Web3.toHex(msg))
        return w3.eth.account.sign_message(msg_ready, private_key=private_key)

    def sign_message_notary(transaction):
        nonce = transaction.events["Approved"].items()[0][1]
        sender = transaction.events["Approved"].items()[1][1]
        value = transaction.events["Approved"].items()[2][1]
        msg = Web3.solidityKeccak(['bytes32','address','uint256'],[nonce, sender, value])
        private_key = notary.private_key
        signature= notary.address
        msg_ready = encode_defunct(hexstr=Web3.toHex(msg))
        return w3.eth.account.sign_message(msg_ready, private_key=private_key)

    # Getting the ABI to be able to connect to the contract on L2
    with open("build/contracts/Bridge.abi") as json_file:
        json_data = json.load(json_file)

    deployed_bridge = Contract.from_abi("Bridge","0xF61db0C5ac21d6d3d19FA85E5Dce907525E9367D",json_data)

    # initiate transfer aka bridge
    # transaction = user.transfer(deployed_bridge, 10 * 10**6)

    # sign a message with info from approved event
    # approver_signed_message = sign_message_approver(transaction)
    approver_signed_message = sign_message_approver()

    print("SIGNED MESSAGE: APPROVED")
    print(approver_signed_message)

    # approve function call passing in the signed message from approver
    signed = deployed_bridge.approve("0x9b7f1863b19652310d16a35b950b6dfefc26efa4efd2ff91037348e57196b997",approver_signed_message.messageHash, approver_signed_message.signature, {'from': approver})
    print("APPROVED EVENT")
    print(signed.events)

    # sign message with info from signed event
    notary_signed_message = sign_message_notary(signed)
    print("SIGNED MESSAGE: NOTARIZED")
    print(notary_signed_message)

    # notarize function call passing in the signed message from notary
    notarized = deployed_bridge.approve(signed.events["Approved"].items()[0][1],notary_signed_message.messageHash, notary_signed_message.signature, {'from': notary})
    print("NOTARIZED EVENT")
    print(notarized.events)