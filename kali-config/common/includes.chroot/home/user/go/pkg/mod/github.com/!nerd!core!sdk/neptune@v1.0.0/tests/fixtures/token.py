import pytest
from brownie import Token

from tests.fixtures.accounts import owner

@pytest.fixture(scope="module", autouse=True)
def token(owner):    
    
    name = "Elevate Bridged Token Test"
    symbol = "ELVbt"
    decimals = 8
    tokenMaxSupply = 50*10**decimals*10**9

    # deploy contract
    token = Token.deploy(name, symbol, decimals, owner.address, owner.address, tokenMaxSupply, {'from': owner})

    return token