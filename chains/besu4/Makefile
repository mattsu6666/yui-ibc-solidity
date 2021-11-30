BESU_VERSION ?= 21.1
BESU ?= docker run --rm -v ${CURDIR}/networkFiles:/tmp/besu/networkFiles -v ${CURDIR}/ibftConfigFile.json:/tmp/besu/ibftConfigFile.json -u root hyperledger/besu:${BESU_VERSION}

.PHONY: generate-blockchain-config
generate-blockchain-config:
	rm -rf ${CURDIR}/networkFiles
	$(BESU) operator generate-blockchain-config --config-file=/tmp/besu/ibftConfigFile.json --to=/tmp/besu/networkFiles --private-key-file-name=key

# TODO: Fix docker-compose cmd to variable
.PHONY: network
network:
	docker-compose up -d --build

.PHONY: network-down
network-down:
	docker-compose down -v --remove-orphans

.PHONY: clean
clean:
	rm -rf networkFiles
