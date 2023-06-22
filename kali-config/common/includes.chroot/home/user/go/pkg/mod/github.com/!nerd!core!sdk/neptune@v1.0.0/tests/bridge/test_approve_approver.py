import pytest
from brownie import accounts, reverts, chain
from brownie.convert import to_bytes

from tests.fixtures.bridge import bridge
from tests.fixtures.token import token
from tests.fixtures.accounts import owner, approver, notary, feeReceiver
from web3 import Web3
from web3.auto import w3
from eth_account.messages import encode_defunct, defunct_hash_message, encode_structured_data
from eth_utils import (
    keccak,
)
from eth_abi import encode_abi
from eip712.messages import EIP712Message, EIP712Type
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

def test_approve(bridge, approver):
    tx = accounts[11].transfer(bridge, amount="1 ether", data="0xa9059cbb00000000000000000000000095fccfaae554dcfcd40e46a7b2e9eccf7f8e34a800000000000000000000000000000000000000000000000000000000004c4b40")
    key = tx.events["Received"].items()[0][1]
    print(tx.events["Received"])
    sender = tx.events["Received"].items()[1][1]
    value = tx.events["Received"].items()[2][1]
    msg = build_message(key, sender, value, bridge.address)

    private_key = "0x8c1d4fb91836192c8a1ea0f1ff0d83c330884a6847bf9e1168e9eed3af18f848"
    signature="0xB2b54bf76156E617017C9575D35a78D991d700Fb"
    signed_message = w3.eth.account.sign_message(msg, private_key=private_key)
    print("SIGNED MESSAGE: ",signed_message)
    signed_hash = "0xd5682087ec540d8c30c3379f9725de8c77bc8754103b1d680cbc3ac55b04c0d3"
    sig = "0x0f29c5615c255eced47c42c3f7d29735598391a51e2b9b038908dc01bdbd930a4867b3554d5c595b2c9c0b64fdb57e588b1a97d8e39d6532648739000c4a927601"
    # print(Web3.toHex(msg))
    signed = bridge.approve(key ,signed_message.messageHash, signed_message.signature, {'from': approver})
    print(signed)
    assert(signed.events["Approved"]["key"] == tx.events["Received"]["key"])
    assert(signed.events["Approved"]["bridgedAmount"] == tx.events["Received"]["bridgedAmount"])
    assert(signed.events["Approved"]["sender"] == tx.events["Received"]["sender"])
    assert(signed.events["Approved"]["approver"] == approver.address)

def test_approve_twice(bridge, approver):
    tx = accounts[11].transfer(bridge, amount="1 ether", data="0xa9059cbb00000000000000000000000095fccfaae554dcfcd40e46a7b2e9eccf7f8e34a800000000000000000000000000000000000000000000000000000000004c4b40")
    key = tx.events["Received"].items()[0][1]
    sender = tx.events["Received"].items()[1][1]
    value = tx.events["Received"].items()[2][1]
    msg = build_message(key, sender, value, bridge.address)

    # msg = Web3.solidityKeccak(['bytes32','address','uint256'],[key,sender, value])
    private_key = "8c1d4fb91836192c8a1ea0f1ff0d83c330884a6847bf9e1168e9eed3af18f848"
    signature="0xB2b54bf76156E617017C9575D35a78D991d700Fb"
    # msg_ready = encode_defunct(hexstr=Web3.toHex(msg))
    signed_message = w3.eth.account.sign_message(msg, private_key=private_key)
    signed = bridge.approve(tx.events["Received"].items()[0][1],signed_message.messageHash, signed_message.signature, {'from': approver})
    with reverts("dev: already approved"):
        signed = bridge.approve(tx.events["Received"].items()[0][1],signed_message.messageHash, signed_message.signature, {'from': approver})


def test_approve_not_matching_key(bridge, approver):
    tx = accounts[11].transfer(bridge, amount="1 ether", data="0xa9059cbb00000000000000000000000095fccfaae554dcfcd40e46a7b2e9eccf7f8e34a800000000000000000000000000000000000000000000000000000000004c4b40")
    key = tx.events["Received"].items()[0][1]
    sender = tx.events["Received"].items()[1][1]
    value = tx.events["Received"].items()[2][1]
    msg = build_message(key, sender, value, bridge.address)

    # msg = Web3.solidityKeccak(['bytes32','address','uint256'],[key,sender, 0])
    private_key = "8c1d4fb91836192c8a1ea0f1ff0d83c330884a6847bf9e1168e9eed3af18f848"
    signature="0xB2b54bf76156E617017C9575D35a78D991d700Fb"
    # msg_ready = encode_defunct(hexstr=Web3.toHex(msg))
    signed_message = w3.eth.account.sign_message(msg, private_key=private_key)
    with reverts("dev: values do not match"):
        signed = bridge.approve(tx.events["Received"].items()[0][1],signed_message.messageHash, signed_message.signature, {'from': approver})

def test_approve_invalid_signer(bridge, approver):
    tx = accounts[11].transfer(bridge, amount="1 ether", data="0xa9059cbb00000000000000000000000095fccfaae554dcfcd40e46a7b2e9eccf7f8e34a800000000000000000000000000000000000000000000000000000000004c4b40")
    key = tx.events["Received"].items()[0][1]
    sender = tx.events["Received"].items()[1][1]
    value = tx.events["Received"].items()[2][1]
    msg = build_message(key, sender, value, bridge.address)

    # msg = Web3.solidityKeccak(['bytes32','address','uint256'],[key,sender, value])
    private_key = "8c1d4fb91836192c8a1ea0f1ff0d83c330884a6847bf9e1168e9eed3af18f848"
    signature="0xB2b54bf76156E617017C9575D35a78D991d700Fb"
    # msg_ready = encode_defunct(hexstr=Web3.toHex(msg))
    signed_message = w3.eth.account.sign_message(msg, private_key=private_key)
    with reverts("dev: invalid signer"):
        signed = bridge.approve(tx.events["Received"].items()[0][1],signed_message.messageHash, signed_message.signature, {'from': accounts[10]})

def test_approve_invalid_signer_not_approver(bridge, approver):
    tx = accounts[11].transfer(bridge, amount="1 ether", data="0xa9059cbb00000000000000000000000095fccfaae554dcfcd40e46a7b2e9eccf7f8e34a800000000000000000000000000000000000000000000000000000000004c4b40")
    key = tx.events["Received"].items()[0][1]
    sender = tx.events["Received"].items()[1][1]
    value = tx.events["Received"].items()[2][1]
    msg = build_message(key, sender, value, bridge.address)

    # msg = Web3.solidityKeccak(['bytes32','address','uint256'],[key,sender, value])
    private_key = "2b020437510fd3e98dcd41610e17fac57b2a1f22342ad692c33a40e62bce19a5"
    signature="0x826De92f47021b69f9b93090346395AEcd50f92C"
    # msg_ready = encode_defunct(hexstr=Web3.toHex(msg))
    signed_message = w3.eth.account.sign_message(msg, private_key=private_key)
    with reverts("dev: invalid signed message"):
        signed = bridge.approve(tx.events["Received"].items()[0][1],signed_message.messageHash, signed_message.signature, {'from': approver})     

def test_approve_invalid_signer_signature_not_approver(bridge, approver):
    tx = accounts[11].transfer(bridge, amount="1 ether", data="0xa9059cbb00000000000000000000000095fccfaae554dcfcd40e46a7b2e9eccf7f8e34a800000000000000000000000000000000000000000000000000000000004c4b40")
    key = tx.events["Received"].items()[0][1]
    sender = tx.events["Received"].items()[1][1]
    value = tx.events["Received"].items()[2][1]
    msg = build_message(key, sender, value, bridge.address)

    # msg = Web3.solidityKeccak(['bytes32','address','uint256'],[key, sender, value])
    private_key = "2b020437510fd3e98dcd41610e17fac57b2a1f22342ad692c33a40e62bce19a5"
    signature1="0xB2b54bf76156E617017C9575D35a78D991d700Fb"
    signature="0x826De92f47021b69f9b93090346395AEcd50f92C"
    # msg_ready = encode_defunct(hexstr=Web3.toHex(msg))
    signed_message = w3.eth.account.sign_message(msg, private_key=private_key)
    with reverts("dev: invalid signed message"):
        signed = bridge.approve(tx.events["Received"].items()[0][1],signed_message.messageHash, signed_message.signature, {'from': approver})

def test_approve_submission_not_found(bridge, approver):
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
    with reverts("dev: invalid submission"):
        signed = bridge.approve(0,signed_message.messageHash, signed_message.signature, {'from': approver})