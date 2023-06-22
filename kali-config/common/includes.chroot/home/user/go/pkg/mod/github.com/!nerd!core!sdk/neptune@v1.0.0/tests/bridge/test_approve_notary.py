import pytest
from brownie import accounts, reverts, chain

from tests.fixtures.bridge import bridge
from tests.fixtures.accounts import owner, approver, notary, feeReceiver
from web3 import Web3
from web3.auto import w3
from eth_account.messages import encode_defunct, defunct_hash_message, encode_structured_data
import json

def build_message(key, sender, value, bridge):
    print(type(key))
    print(key)
    data = {
        "types": {
            "EIP712Domain": [
                {"name": "name", "type": "string"},
                {"name": "version", "type": "string"},
                {"name": "chainId", "type": "uint256"},
            ],
            "SignedMessage": [
                # {"name": "key", "type": "bytes32"},
                {"name": "sender", "type": "address"},
                {"name": "amount", "type": "uint256"},
            ],
        },
        "domain": {
            "name": "Neptune Bridge", 
            "version": "0.0.1",
            "chainId": 7,
        },
        "primaryType": "SignedMessage",
        "message": {
            # "key": key,
            "sender": sender,
            "amount": value,
        },
    }
    
    new_msg = json.loads(json.dumps(data))
    return encode_structured_data(data)

def test_approve_notarize(bridge, approver, notary):
    ogbalance = accounts[15].balance()
    tx = accounts[11].transfer(bridge, amount="1 ether", data="0xa9059cbb00000000000000000000000095fccfaae554dcfcd40e46a7b2e9eccf7f8e34a800000000000000000000000000000000000000000000000000000000004c4b40")
    key = tx.events["Received"].items()[0][1]
    sender = tx.events["Received"].items()[1][1]
    value = tx.events["Received"].items()[2][1]
    msg = build_message(key, sender, value, bridge.address)

    # msg = Web3.solidityKeccak(['bytes32','address','uint256'],[key, sender, value])
    private_key = "8c1d4fb91836192c8a1ea0f1ff0d83c330884a6847bf9e1168e9eed3af18f848"
    signature="0xB2b54bf76156E617017C9575D35a78D991d700Fb"
    # msg_ready = encode_defunct(hexstr=Web3.toHex(msg))
    signed_message = w3.eth.account.sign_message(msg, private_key=private_key)
    signed = bridge.approve(tx.events["Received"].items()[0][1],signed_message.messageHash, signed_message.signature, {'from': approver})
    key = signed.events["Approved"].items()[0][1]
    sender = signed.events["Approved"].items()[1][1]
    value = signed.events["Approved"].items()[2][1]
    msg = build_message(key, sender, value, bridge.address)

    # msg = Web3.solidityKeccak(['bytes32','address','uint256'],[key, sender, value])
    private_key = "9d218033dad4754a59c1f3475ea6392a8b69f7519dc15d4adff25412319eb3a4"
    signature="0xD634496E9EEA51887926D0a2f8Bd115Dc8B276f0"
    # msg_ready = encode_defunct(hexstr=Web3.toHex(msg))
    notarized_message = w3.eth.account.sign_message(msg, private_key=private_key)
    notarized = bridge.approve(signed.events["Approved"].items()[0][1],notarized_message.messageHash, notarized_message.signature, {'from': notary})
    assert(notarized.events["Notarized"]["nonce"] == tx.events["Received"]["key"])
    assert(notarized.events["Notarized"]["bridgedAmount"] == tx.events["Received"]["bridgedAmount"])
    assert(notarized.events["Notarized"]["sender"] == tx.events["Received"]["sender"])
    assert(notarized.events["Notarized"]["notary"] == notary.address)
    assert(notarized.events["Notarized"]["approvedMessage"] == (Web3.toHex(signed_message.signature)))
    assert(notarized.events["Notarized"]["notarizeMessage"] == (Web3.toHex(notarized_message.signature)))
    assert(notarized.events["Notarized"]["messageHash"] == (Web3.toHex(notarized_message.messageHash)))
    assert(notarized.events["Notarized"]["messageHash"] == (Web3.toHex(signed_message.messageHash)))


    assert(ogbalance+bridge.getFee() == accounts[15].balance())

def test_approve_notarize_twice(bridge, approver, notary):
    tx = accounts[11].transfer(bridge, amount="1 ether", data="0xa9059cbb00000000000000000000000095fccfaae554dcfcd40e46a7b2e9eccf7f8e34a800000000000000000000000000000000000000000000000000000000004c4b40")
    key = tx.events["Received"].items()[0][1]
    sender = tx.events["Received"].items()[1][1]
    value = tx.events["Received"].items()[2][1]
    msg = build_message(key, sender, value, bridge.address)

    # msg = Web3.solidityKeccak(['bytes32','address','uint256'],[key, sender, value])
    private_key = "8c1d4fb91836192c8a1ea0f1ff0d83c330884a6847bf9e1168e9eed3af18f848"
    signature="0xB2b54bf76156E617017C9575D35a78D991d700Fb"
    # msg_ready = encode_defunct(hexstr=Web3.toHex(msg))
    signed_message = w3.eth.account.sign_message(msg, private_key=private_key)
    signed = bridge.approve(key ,signed_message.messageHash, signed_message.signature, {'from': approver})
    print(signed)
    key = signed.events["Approved"].items()[0][1]
    sender = signed.events["Approved"].items()[1][1]
    value = signed.events["Approved"].items()[2][1]
    msg = build_message(key, sender, value, bridge.address)

    # msg = Web3.solidityKeccak(['bytes32','address','uint256'],[key, sender, value])
    private_key = "9d218033dad4754a59c1f3475ea6392a8b69f7519dc15d4adff25412319eb3a4"
    signature="0xD634496E9EEA51887926D0a2f8Bd115Dc8B276f0"
    # msg_ready = encode_defunct(hexstr=Web3.toHex(msg))
    notarized_message = w3.eth.account.sign_message(msg, private_key=private_key)
    notarized = bridge.approve(key ,notarized_message.messageHash, notarized_message.signature, {'from': notary})
    with reverts(""):
        notarized = bridge.approve(signed.events["Approved"].items()[0][1],notarized_message.messageHash, notarized_message.signature, {'from': notary})

def test_notarize_before_approved(bridge, approver, notary):
    tx = accounts[11].transfer(bridge, amount="1 ether", data="0xa9059cbb00000000000000000000000095fccfaae554dcfcd40e46a7b2e9eccf7f8e34a800000000000000000000000000000000000000000000000000000000004c4b40")
    key = tx.events["Received"].items()[0][1]
    sender = tx.events["Received"].items()[1][1]
    value = tx.events["Received"].items()[2][1]
    msg = build_message(key, sender, value, bridge.address)
    
    # msg = Web3.solidityKeccak(['bytes32','address','uint256'],[key, sender, value])
    private_key = "8c1d4fb91836192c8a1ea0f1ff0d83c330884a6847bf9e1168e9eed3af18f848"
    signature="0xB2b54bf76156E617017C9575D35a78D991d700Fb"
    # msg_ready = encode_defunct(hexstr=Web3.toHex(msg))
    signed_message = w3.eth.account.sign_message(msg, private_key=private_key)
    with reverts("dev: requires prior approval"):
        signed = bridge.approve(tx.events["Received"].items()[0][1],signed_message.messageHash, signed_message.signature, {'from': notary})

# def test_approve_notarize(bridge, approver, notary):
#     tx = accounts[11].transfer(bridge, amount="1 ether", data="0xa9059cbb00000000000000000000000095fccfaae554dcfcd40e46a7b2e9eccf7f8e34a800000000000000000000000000000000000000000000000000000000004c4b40")
#     key = tx.events["Received"].items()[0][1]
#     sender = tx.events["Received"].items()[1][1]
#     value = tx.events["Received"].items()[2][1]
#     msg = Web3.solidityKeccak(['bytes32','address','uint256'],[key, sender, value])
#     private_key = "8c1d4fb91836192c8a1ea0f1ff0d83c330884a6847bf9e1168e9eed3af18f848"
#     signature="0xB2b54bf76156E617017C9575D35a78D991d700Fb"
#     msg_ready = encode_defunct(hexstr=Web3.toHex(msg))
#     signed_message = w3.eth.account.sign_message(msg_ready, private_key=private_key)
#     signed = bridge.approve(key, signed_message.messageHash, signed_message.signature, {'from': approver})
#     print(signed)
#     key = tx.events["Approved"].items()[0][1]
#     sender = tx.events["Approved"].items()[1][1]
#     value = tx.events["Approved"].items()[2][1]
#     msg = Web3.solidityKeccak(['bytes32','address','uint256'],[key, sender, value])
#     private_key = "9d218033dad4754a59c1f3475ea6392a8b69f7519dc15d4adff25412319eb3a4"
#     signature="0xD634496E9EEA51887926D0a2f8Bd115Dc8B276f0"
#     msg_ready = encode_defunct(hexstr=Web3.toHex(msg))
#     notarized_message = w3.eth.account.sign_message(msg_ready, private_key=private_key)
#     with reverts():
#         notarized = bridge.approve(signed.events["Approved"].items()[0][1],notarized_message.messageHash, notarized_message.signature, {'from': notary})

# TODO: check burned tokens go to 0x00 address

# TODO: pending[submission] is deleted

# TODO: approvedMessage[submission] is deleted
