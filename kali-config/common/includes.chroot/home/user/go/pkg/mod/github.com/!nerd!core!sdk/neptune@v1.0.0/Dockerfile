# ghcr.io/{organization}/{repository}
#
# Token
#
# Build:
#
#    % docker build -t bridge/1.0 .
#
# Run:
#
#    % docker run -it --rm brownie test
#
#    % docker run -it --rm --volume "$(pwd)":/code brownie test --coverage
#

FROM minty/brownie:1.19.2

ENV HOME=/tmp

RUN apk add curl
RUN python -c "from brownie.project.compiler import install_solc;install_solc('0.8.10')"
RUN ls
RUN brownie networks add live elevate-testnet host=https://rpc.testnet.elevate.blockfabric.net:8669/ chainid=7 

COPY . /code

RUN chmod -R 777 /code
RUN chmod -R 777 /tmp

WORKDIR /code
