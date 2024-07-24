RELAYER_TAG := $(shell grep '^ENV RELAYER_TAG' cosmos-relayer/Dockerfile | cut -f3 -d\ )
BABYLON_FULL_PATH := $(shell git rev-parse --show-toplevel)

all: babylond cosmos-relayer

babylond: babylond-rmi
	docker build --tag babylonchain/babylond -f babylond/Dockerfile ${BABYLON_FULL_PATH}

babylond-before-upgrade:
	docker rmi babylonchain/babylond-before-upgrade 2>/dev/null; true && \
	docker build --tag babylonchain/babylond-before-upgrade -f babylond/Dockerfile \
    --build-arg VERSION="a33a3344bb44bde2f3374d3cbf919abb942c341a" ${BABYLON_FULL_PATH}

babylond-rmi:
	docker rmi babylonchain/babylond 2>/dev/null; true

e2e-init-chain:
	@DOCKER_BUILDKIT=1 docker build -t babylonchain/babylond-e2e-init-chain --build-arg E2E_SCRIPT_NAME=chain --platform=linux/x86_64 -f e2e-initialization/init.Dockerfile ${BABYLON_FULL_PATH}

cosmos-relayer: cosmos-relayer-rmi
	docker build --tag babylonchain/cosmos-relayer:${RELAYER_TAG} -f cosmos-relayer/Dockerfile \
				${BABYLON_FULL_PATH}/contrib/images/cosmos-relayer
	docker tag babylonchain/cosmos-relayer:${RELAYER_TAG} babylonchain/cosmos-relayer:latest

cosmos-relayer-rmi:
	docker rmi babylonchain/cosmos-relayer 2>/dev/null; true

.PHONY: all babylond cosmos-relayer babylond-rmi cosmos-relayer-rmi
