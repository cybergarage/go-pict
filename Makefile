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

PKG_NAME=pict
PKG_VER=$(shell git describe --abbrev=0 --tags)

MODULE_ROOT=github.com/cybergarage/go-pict

PKG_SRC_ROOT=${PKG_NAME}
PKG=\
	${MODULE_ROOT}/${PKG_SRC_ROOT}/...

TEST_SRC_ROOT=test
TEST_PKG=\
	${MODULE_ROOT}/${TEST_SRC_ROOT}/...

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

picts := $(patsubst %.mod,%.pict,$(wildcard test/picts/*.mod))

test/picts/embed.go : test/picts/embed.pl $(wildcard test/picts/*.pict)
	perl $< > $@

test: lint $(picts) test/picts/embed.go
	go test -v -p 1 -cover -timeout 60s ${PKG} ${TEST_PKG}

clean:
	go clean -i ${PKG}
