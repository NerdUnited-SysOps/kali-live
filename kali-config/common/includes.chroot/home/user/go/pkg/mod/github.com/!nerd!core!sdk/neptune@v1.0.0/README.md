# Bridge Contracts

Bridge contract for L1 and L2, brand level tokens.

## Building

##### Prerequisites

 - [Docker 20+][docker_install]  
 - [Go 1.19+][go_install]  
 - [abigen](https://github.com/ethereum/go-ethereum/tree/master/cmd/abigen), Ethereum automatic code generation derived from a contracts ABI & compiled bytecode  
    ```go install github.com/ethereum/go-ethereum/cmd/abigen@v1.10.25```
 - Create an env file with two values: 
    - WEB3_INFURA_PROJECT_ID 
    - ETHERSCAN_TOKEN

Build the Docker image:

```
docker build -t ghcr.io/nerdunited/lockup .
```
_(the above container registery is not valid/active, build locally via: `docker build -t ghcr.io/nerdunited/lockup .`)_

## Running

Start a Brownie console with:

```
% ./dockit.sh console
```

_**NOTE**: the `dockit.sh` script mounts your local bridge code over the images, manually run `docker` command to run against image code_

Next, deploy the bridge to the running Ganache node:

```
>>> from scripts import sign_message
>>> from scripts import bridge
>>> from scripts import token
>>> from scripts import e2e
>>> from scripts import e2e_goerli
>>> from scripts import l1_bridge
```

Finally, you should now be able to execute all of the bridge functionality from the Brownie console.

## Testing

```
% ./dockit.sh test --coverage
```

## Accounts

The mnemonic for the ganache-cli accounts:

`mirror mixture pledge rail cover discover crouch shock fun mammal long next oblige happy knee merge lemon situate bike nature ankle stool seek below`

Account keys:

| Role          | Address                                    | Private Key                                                      |
| ------------- | ------------------------------------------ | ---------------------------------------------------------------- |
| Owner         | 0x3f1AE583e83260D239cD179886a095b96e67bc00 | e03fa64de89252ba2076748a38a38c766bed65b148708bc69ce0070283fbefed |
| Issuer        | 0x37d8c1257babb1C624C35Fb74BC0E5964f86aee8 | 37858f1ce82d3c09c935a98a0e7e0006ce391e1bdfce95f62b706347d53096c7 |
|               | 0x927dfe930072fF3BCd5526B2D6baF7903CCeaCbe | 17892092e6ed6a076a2b752ff5fedde56c58473c0d70613c110840a4ebc9f338 |
|               | 0xB2b54bf76156E617017C9575D35a78D991d700Fb | 8c1d4fb91836192c8a1ea0f1ff0d83c330884a6847bf9e1168e9eed3af18f848 |
|               | 0xD634496E9EEA51887926D0a2f8Bd115Dc8B276f0 | 9d218033dad4754a59c1f3475ea6392a8b69f7519dc15d4adff25412319eb3a4 |
|               | 0x9705dB198c95D4b545d0674727091Ee7af9af01F | 161dbc56b06992271f5dea4be0d813511ef19775e5b74796210044641aabdeb4 |
|               | 0x826De92f47021b69f9b93090346395AEcd50f92C | 2b020437510fd3e98dcd41610e17fac57b2a1f22342ad692c33a40e62bce19a5 |
|               | 0x5037021C6555e6167392B536a853d603959E1e72 | 1ac6f927d7e3bad3cd2f8612bea0e37631fe08ad50fe249461cfb4b1bed867d9 |
|               | 0xE67c7A0Ca87D64a3dC24aB51aDd0B839B7DeCF97 | 55566481b27c7abbe2dad356e2b4385e87814d790a724887f31d5c2b016797f7 |
|               | 0xc012fb177A6Fb90675183519A6D4884a243dd741 | f320a0ea1b5362d16b7797cf7a41ffa51a5ed088b797a06271e9a0a48f9d865a |
|               | 0x32eF24805bb2020773e0C530885e01Cf21dD75bB | ab2665304c146db371be6b3fc640d55a7237a1d68bda4b657c65b8d814774f40 |
|               | 0x3bE660dA2ccf5D00af284273acAB0251852b1cAb | 21d008a6e1be7b5271d389ee8623af04ca32cda87826c69191e9114332edabee |
|               | 0xc2784e233E136B34989382B192679F34FEEeE4b0 | 2900b4be05864f901cf7ff14d3e73556173b440ec09c7c16cf7477d23627e331 |
|               | 0x3101be197D448228afdCDAb3627dE7aBEBaad070 | 61989102a8a768f592b6826c94511ca2fd12f700bb0b66ac08544de0c06f0658 |
|               | 0xe1f423E826A6627DC1EcDaD6E51d42dA2f80559f | d13bf95b45848244741ed465293702fdf873026105b95cdbf7083d01b34df301 |
|               | 0x98Cf0e8E2ed62ec99fA22Fc2CC7e9Ac0fc543622 | 693804fe582d52bf68ad8b97667c312d4129637850ed2d19d171b45b74ad3abc |


[python3]: https://www.python.org/downloads/
[brownie]: https://github.com/eth-brownie/brownie/tree/v1.17.0
[brownie_docs]: https://eth-brownie.readthedocs.io/en/stable/
[brownie_install]: https://eth-brownie.readthedocs.io/en/stable/install.html
[docker_install]: https://docs.docker.com/get-docker/
[go_install]: https://golang.org/doc/install
