import pytest
from brownie import accounts, reverts, chain

from tests.fixtures.bridge import bridge
from tests.fixtures.accounts import owner, approver, notary, feeReceiver
from web3 import Web3
from web3.auto import w3
from eth_account.messages import encode_defunct, defunct_hash_message

fees = 5 * 10**6

def test_fallback(bridge):
    tx = accounts[5].transfer(bridge, amount="1 ether", data="0xa9059cbb00000000000000000000000095fccfaae554dcfcd40e46a7b2e9eccf7f8e34a800000000000000000000000000000000000000000000000000000000004c4b40")
    amount = 10**18-bridge.getFee()
    assert(tx.events["Received"]["sender"] == accounts[5].address)
    assert(tx.events["Received"]["bridgedAmount"] == amount)
    assert(tx.events["Received"]["feePaid"] == bridge.getFee())

def test_receive(bridge):
    tx = accounts[6].transfer(bridge, amount="1 ether")
    amount = 10**18-bridge.getFee()
    assert(tx.events["Received"]["sender"] == accounts[6].address)
    assert(tx.events["Received"]["bridgedAmount"] == amount)
    assert(tx.events["Received"]["feePaid"] == bridge.getFee())

def test_not_sufficient_to_pay_fees(bridge):
    amount=1000 * 10**8 - 1
    with reverts("dev: insufficient funds"):
        tx = accounts[5].transfer(bridge, amount=amount, data="0xa9059cbb00000000000000000000000095fccfaae554dcfcd40e46a7b2e9eccf7f8e34a800000000000000000000000000000000000000000000000000000000004c4b40")

def test_sufficient_to_pay_after_fees(bridge):
    amount = 1000 * 10**8 + 1
    tx = accounts[5].transfer(bridge, amount=amount, data="0xa9059cbb00000000000000000000000095fccfaae554dcfcd40e46a7b2e9eccf7f8e34a800000000000000000000000000000000000000000000000000000000004c4b40")
    amount = amount-bridge.getFee()
    assert(tx.events["Received"]["sender"] == accounts[5].address)
    assert(tx.events["Received"]["bridgedAmount"] == amount)

def test_fallback_zero_amount(bridge):
    with reverts("dev: insufficient funds"):
        tx = accounts[7].transfer(bridge, amount="0 ether", data="0xa9059cbb00000000000000000000000095fccfaae554dcfcd40e46a7b2e9eccf7f8e34a800000000000000000000000000000000000000000000000000000000004c4b40")

def test_fallback_not_enough_gas(bridge):
    with pytest.raises(Exception) as e_info:
        tx = accounts[8].transfer(bridge, amount="100 ether", data="0xa9059cbb00000000000000000000000095fccfaae554dcfcd40e46a7b2e9eccf7f8e34a800000000000000000000000000000000000000000000000000000000004c4b40", gas_price=10)

    assert("sender doesn't have enough funds to send tx" in str(e_info.value))

def test_receive_zero_amount(bridge):
    with reverts("dev: insufficient funds"):
        tx = accounts[9].transfer(bridge, amount="0 ether")

# TODO: test when passed from a smart contract (tx.origin != msg.sender)

# TODO: test when you have the same value, sender, nonce for two transactions (test that the submission does exist with the same key)