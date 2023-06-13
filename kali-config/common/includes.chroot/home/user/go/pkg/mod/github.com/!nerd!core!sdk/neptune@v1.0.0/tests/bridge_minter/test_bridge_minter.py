import pytest
from brownie import accounts, reverts, chain

from tests.fixtures.token import token
from tests.fixtures.bridge import bridge
from tests.fixtures.bridge_minter import bridge_minter
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

def test_bridge_minter(bridge, approver, bridge_minter, token, notary):
    tx = accounts[11].transfer(bridge, amount="1 ether", data="0xa9059cbb00000000000000000000000095fccfaae554dcfcd40e46a7b2e9eccf7f8e34a800000000000000000000000000000000000000000000000000000000004c4b40")
    key = tx.events["Received"].items()[0][1]
    sender = tx.events["Received"].items()[1][1]
    value = tx.events["Received"].items()[2][1]
    msg = build_message(key, sender, value, bridge.address)
    private_key = "8c1d4fb91836192c8a1ea0f1ff0d83c330884a6847bf9e1168e9eed3af18f848"
    signature="0xB2b54bf76156E617017C9575D35a78D991d700Fb"
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
    print(notarized)
    txn = bridge_minter.bridge(notarized.events["Notarized"].items()[0][1],notarized.events["Notarized"].items()[1][1],notarized.events["Notarized"].items()[2][1],notarized.events["Notarized"].items()[3][1],notarized.events["Notarized"].items()[4][1],notarized.events["Notarized"].items()[5][1])
    print(txn)
    assert(txn.events["Bridged"].items()[0][1] == accounts[11].address)
    assert(txn.events["Bridged"].items()[1][1] == value)
    assert(token.balanceOf(accounts[11].address) == value)

# invalid approved message
def test_bridge_minter_invalid_approver(bridge, approver, bridge_minter, token, notary):
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
    private_key = "ab2665304c146db371be6b3fc640d55a7237a1d68bda4b657c65b8d814774f40"
    signature="0x32eF24805bb2020773e0C530885e01Cf21dD75bB"
    # msg_ready = encode_defunct(hexstr=Web3.toHex(msg))
    signed_message = w3.eth.account.sign_message(msg, private_key=private_key)
    with reverts(""):
        txn = bridge_minter.bridge(notarized.events["Notarized"].items()[0][1],notarized.events["Notarized"].items()[1][1],notarized.events["Notarized"].items()[2][1],notarized.events["Notarized"].items()[3][1],signed_message.signature,notarized.events["Notarized"].items()[5][1])

def test_bridge_minter_invalid_notary(bridge, approver, bridge_minter, token, notary):
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
    private_key = "ab2665304c146db371be6b3fc640d55a7237a1d68bda4b657c65b8d814774f40"
    signature="0x32eF24805bb2020773e0C530885e01Cf21dD75bB"
    # msg_ready = encode_defunct(hexstr=Web3.toHex(msg))
    signed_message = w3.eth.account.sign_message(msg, private_key=private_key)
    with reverts(""):
        txn = bridge_minter.bridge(notarized.events["Notarized"].items()[0][1],notarized.events["Notarized"].items()[1][1],notarized.events["Notarized"].items()[2][1],notarized.events["Notarized"].items()[3][1],notarized.events["Notarized"].items()[4][1],signed_message.signature)

def test_bridge_minter_invalid_hash(bridge, approver, bridge_minter, token, notary):
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
    private_key = "ab2665304c146db371be6b3fc640d55a7237a1d68bda4b657c65b8d814774f40"
    signature="0x32eF24805bb2020773e0C530885e01Cf21dD75bB"
    msg = build_message(key, signature, value, bridge.address)
    signed_message = w3.eth.account.sign_message(msg, private_key=private_key)
    with reverts("dev: values do not match"):
        txn = bridge_minter.bridge(notarized.events["Notarized"].items()[0][1],notarized.events["Notarized"].items()[1][1],notarized.events["Notarized"].items()[2][1],signed_message.messageHash,notarized.events["Notarized"].items()[4][1],notarized.events["Notarized"].items()[5][1])

# invalid notarized messages

#