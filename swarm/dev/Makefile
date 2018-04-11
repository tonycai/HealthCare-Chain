.PHONY: build cluster test

default: build

build:
	go build -o bin/swarm    github.com/tonycai/HealthCare-Chain/cmd/swarm
	go build -o bin/geth     github.com/tonycai/HealthCare-Chain/cmd/geth
	go build -o bin/bootnode github.com/tonycai/HealthCare-Chain/cmd/bootnode

cluster: build
	scripts/boot-cluster.sh

test:
	go test -v github.com/tonycai/HealthCare-Chain/swarm/...
