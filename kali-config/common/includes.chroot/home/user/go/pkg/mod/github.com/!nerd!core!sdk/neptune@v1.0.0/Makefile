# BridgeMinter/Bridge Smart Contracts
#


.PHONY: build clean

GOJQ      = "github.com/itchyny/gojq/cmd/gojq@v0.12.9"
CONTRACTS = "./build/contracts"
VERSION   = "0.1.0"

.brownie:
	./dockit.sh compile

.abibin:
	cat $(CONTRACTS)/Bridge.json       	     | go run $(GOJQ) -c '.abi'              > $(CONTRACTS)/Bridge.abi
	cat $(CONTRACTS)/Bridge.json       		 | go run $(GOJQ) -r '.bytecode'         > $(CONTRACTS)/Bridge.bin
	cat $(CONTRACTS)/Bridge.json       		 | go run $(GOJQ) -r '.deployedBytecode' > $(CONTRACTS)/Bridge.bin-runtime
	cat $(CONTRACTS)/BridgeMinter.json       | go run $(GOJQ) -c '.abi'              > $(CONTRACTS)/BridgeMinter.abi
	cat $(CONTRACTS)/BridgeMinter.json       | go run $(GOJQ) -r '.bytecode'         > $(CONTRACTS)/BridgeMinter.bin
	cat $(CONTRACTS)/BridgeMinter.json       | go run $(GOJQ) -r '.deployedBytecode' > $(CONTRACTS)/BridgeMinter.bin-runtime
	cat $(CONTRACTS)/Token.json       	     | go run $(GOJQ) -c '.abi'              > $(CONTRACTS)/Token.abi
	cat $(CONTRACTS)/Token.json       	     | go run $(GOJQ) -r '.bytecode'         > $(CONTRACTS)/Token.bin
	cat $(CONTRACTS)/Token.json       		 | go run $(GOJQ) -r '.deployedBytecode' > $(CONTRACTS)/Token.bin-runtime

.go:
	abigen --abi $(CONTRACTS)/BridgeMinter.abi       --bin $(CONTRACTS)/BridgeMinter.bin       --pkg contracts --type BridgeMinter       --out pkg/contracts/bridge_minter.go
	abigen --abi $(CONTRACTS)/Bridge.abi       		 --bin $(CONTRACTS)/Bridge.bin       	   --pkg contracts --type Bridge       	     --out pkg/contracts/bridge.go
	abigen --abi $(CONTRACTS)/Token.abi       		 --bin $(CONTRACTS)/Token.bin       	   --pkg contracts --type Token       	     --out pkg/contracts/token.go


.archive:
	tar                                                 \
		-C $(CONTRACTS)                                 \
		-zcvf $(CONTRACTS)/contracts_v$(VERSION).tar.gz \
		BridgeMinter.abi                                \
		BridgeMinter.bin                                \
		BridgeMinter.bin-runtime                        \
		BridgeMinter.json                               \
	    Token.abi                                       \
		Token.bin                                       \
		Token.bin-runtime                               \
		Token.json                                      \
		Bridge.abi                                      \
		Bridge.bin                                      \
		Bridge.bin-runtime                              \
		Bridge.json

.checksum:
	sha256sum $(CONTRACTS)/contracts_v$(VERSION).tar.gz | awk '{print $1}' > $(CONTRACTS)/contracts_v$(VERSION).sha256.sig

build: .brownie .abibin .go .archive .checksum

clean:
	@rm -rf ./build/*
