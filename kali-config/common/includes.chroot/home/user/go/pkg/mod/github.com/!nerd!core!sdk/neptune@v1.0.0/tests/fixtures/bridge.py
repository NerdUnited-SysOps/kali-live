import pytest
from brownie import Bridge

from tests.fixtures.accounts import owner, notary, approver, feeReceiver

@pytest.fixture(scope="module", autouse=True)
def bridge(owner,notary, approver, feeReceiver):
    fee = 5* 10**6
    contract = owner.deploy(Bridge, owner, approver, notary, feeReceiver, fee)
    yield contract
