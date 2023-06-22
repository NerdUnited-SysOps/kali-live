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
from brownie import Bridge
from eth_abi import encode
from web3 import Web3
from scripts import accounts

def main():
    userAccounts = accounts.main()

    owner = userAccounts[0]
    approver = userAccounts[1]
    notary = userAccounts[2]
    feeReceiver = userAccounts[3]
    fee = 5 * 10**6
    chain_id = 80

    # deploy contract
    bridge = Bridge.deploy(approver, notary, feeReceiver,fee, chain_id, {'from': owner})
    Bridge.publish_source(bridge)

    return bridge

