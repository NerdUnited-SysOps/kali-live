import pytest
from brownie import Distribution

from tests.fixtures.accounts import owner, issuer

@pytest.fixture(scope="module", autouse=True)
def distribution(owner, issuer):
    contract = owner.deploy(Distribution, owner, issuer)
    owner.transfer(contract.address, 500)
    yield contract
