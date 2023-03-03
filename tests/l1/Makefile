# Copyright (c) 2022 - 2023 Intel Corporation
# SPDX-License-Identifier: Apache-2.0

BINARY=main
GO := go
BINARY = main
OC := oc
OPERATOR_RELEASE ?= 0.1.0
SERVICE_ACCOUNT := intel-dgpu
NAMESPACE := intel-dgpu
PACKAGE = main.go

build:
	$(GO) build -o ${BINARY} $(PACKAGE) 

run:
	./${BINARY}

go-mod:
	$(GO) mod tidy

clean:
	$(GO) clean

build-image:
	$(OC) apply -f scc.yaml
	$(OC) adm policy add-scc-to-user container-scc -z $(SERVICE_ACCOUNT) -n $(NAMESPACE)
	$(OC) apply -f l1_build.yaml

deploy:
	$(OC) apply -f l1_deploy.yaml