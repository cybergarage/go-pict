# Copyright (C) 2023 The go-pict Authors. All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#    http:#www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

SHELL := bash

GOBIN := $(shell go env GOPATH)/bin
PATH := $(GOBIN):$(PATH)

PKG_NAME=pict
PKG_VER=$(shell git describe --abbrev=0 --tags)
PKG_COVER=${PKG_NAME}-cover

MODULE_ROOT=github.com/cybergarage/go-pict

PKG_SRC_ROOT=${PKG_NAME}
PKG=\
	${MODULE_ROOT}/${PKG_SRC_ROOT}

TEST_SRC_ROOT=picttest
TEST_PKG=\
	${MODULE_ROOT}/${TEST_SRC_ROOT}

.PHONY: version test format vet lint clean docker
.IGNORE: lint

all: test

version:
	@pushd ${PKG_SRC_ROOT} && ./version.gen > version.go && popd
	-git commit ${PKG_SRC_ROOT}/version.go -m "Update version"

format: version
	gofmt -s -w ${PKG_SRC_ROOT} ${TEST_SRC_ROOT}

vet: format
	go vet ${PKG}

lint: format
	golangci-lint run ${PKG_SRC_ROOT}/... ${TEST_SRC_ROOT}

%.pict : %.mod
	pict $< > $@

picts := $(patsubst %.mod,%.pict,$(wildcard picttest/picts/*.mod))

picttest/picts/embed.go : picttest/picts/embed.pl $(wildcard picttest/picts/*.pict)
	perl $< > $@

test: lint $(picts) picttest/picts/embed.go
	go test -v -p 1 -timeout 10m -cover -coverpkg=${PKG}/... -coverprofile=${PKG_COVER}.out ${PKG}/... ${TEST_PKG}/...
	go tool cover -html=${PKG_COVER}.out -o ${PKG_COVER}.html

cover: test
	open ${PKG_COVER}.html || xdg-open ${PKG_COVER}.html || gnome-open ${PKG_COVER}.html

godoc:
	go install golang.org/x/tools/cmd/godoc@latest
	open http://localhost:6060/pkg/${PKG}/ || xdg-open http://localhost:6060/pkg/${PKG}/ || gnome-open http://localhost:6060/pkg/${PKG}/
	godoc -http=:6060 -play

clean:
	go clean -i ${PKG}
