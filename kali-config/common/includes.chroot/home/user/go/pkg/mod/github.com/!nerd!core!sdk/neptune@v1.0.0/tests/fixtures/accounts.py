import pytest
from brownie import accounts

zero_address = "0x0000000000000000000000000000000000000000"

@pytest.fixture(scope="session", autouse=True)
def owner():
    yield accounts[0]

@pytest.fixture(scope="session", autouse=True)
def issuer():
    yield accounts[1]

@pytest.fixture(scope="session", autouse=True)
def receiver():
    yield accounts[2]

@pytest.fixture(scope="session", autouse=True)
def approver():
    yield accounts[3]

@pytest.fixture(scope="session", autouse=True)
def notary():
    yield accounts[4]

@pytest.fixture(scope="session", autouse=True)
def users():
    yield accounts[5:15]

@pytest.fixture(scope="session", autouse=True)
def feeReceiver():
    yield accounts[15]
