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
from brownie import BridgeMinter
from eth_abi import encode
from web3 import Web3
from scripts import accounts

def main():
    token_address = "0xB0D98244FBf8c41b0A34A0B3845C4e2a131BE3a0"
    userAccounts = accounts.main()

    owner = userAccounts[0]
    approver = userAccounts[1]
    notary = userAccounts[2]
    feeReceiver = userAccounts[3]
    chain_id = 80

    # deploy contract
    bridge_minter = BridgeMinter.deploy(approver.address, notary.address, token_address, chain_id, {'from': owner})
    BridgeMinter.publish_source(bridge_minter)

    return bridge_minter