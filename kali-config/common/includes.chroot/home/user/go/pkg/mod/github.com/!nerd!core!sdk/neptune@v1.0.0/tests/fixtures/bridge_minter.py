import pytest
from brownie import BridgeMinter, Token

from tests.fixtures.accounts import owner, notary, approver, feeReceiver
from tests.fixtures.token import token

@pytest.fixture(scope="module", autouse=True)
def bridge_minter(owner,notary, approver, feeReceiver, token):
    # deploy contract
    bridge_minter = BridgeMinter.deploy(owner.address, approver.address, notary.address, token.address, {'from': owner})
    token.setIssuer(bridge_minter.address)
    return bridge_minter
