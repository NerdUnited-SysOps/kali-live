import pytest
from brownie import Lockup

#from tests.fixtures.accounts import owner, issuer, receiver, users
from tests.fixtures.accounts import owner, issuer, users

@pytest.fixture(scope="module", autouse=True)
def lockup(owner, issuer, receiver, users):
    #yield owner.deploy(Lockup, owner, issuer, receiver, 50)
    yield owner.deploy(Lockup, owner, issuer, 50)
