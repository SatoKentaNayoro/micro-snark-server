SHELL=/usr/bin/env bash

API_PROTO_FILES=$(shell find api -name *.proto)

.PHONY: dev
dev: init api snark_ffi_dev server

.PHONY: release
release: init api snark_ffi_release server

unexport GOFLAGS

ldflags=-X=github.com/filecoin-project/lotus/build.CurrentCommit=+git.$(subst -,.,$(shell git describe --always --match=NeVeRmAtCh --dirty 2>/dev/null || git rev-parse --short HEAD 2>/dev/null))
ifneq ($(strip $(LDFLAGS)),)
	ldflags+=-extldflags=$(LDFLAGS)
endif

GOFLAGS+=-ldflags="-s -w $(ldflags)"

GOCC?=go

.PHONY: init
init:
	@if [ ! -e "$(GOPATH)/bin/protoc-gen-go" ]; then \
		echo "go install google.golang.org/protobuf/cmd/protoc-gen-go@latest"; \
		go install google.golang.org/protobuf/cmd/protoc-gen-go@latest; \
	fi

	@if [ ! -e "$(GOPATH)/bin/protoc-gen-go-grpc" ]; then \
		echo "go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest"; \
		go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest; \
	fi

	@if [ ! -e "$(GOPATH)/bin/kratos" ]; then \
		echo "go install github.com/go-kratos/kratos/cmd/kratos/v2@latest"; \
		go install github.com/go-kratos/kratos/cmd/kratos/v2@latest; \
	fi

	@if [ ! -e "$(GOPATH)/bin/protoc-gen-go-http" ]; then \
		echo "go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest"; \
		go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest; \
	fi

	@if [ ! -e "$(GOPATH)/bin/protoc-gen-go-errors" ]; then \
		echo "go install github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2@latest"; \
		go install github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2@latest; \
	fi

	@if [ ! -e "$(GOPATH)/bin/protoc-gen-openapi" ]; then \
		echo "go install github.com/google/gnostic/cmd/protoc-gen-openapi@v0.6.1"; \
		go install github.com/google/gnostic/cmd/protoc-gen-openapi@v0.6.1; \
	fi


.PHONY: snark_ffi_dev
snark_ffi_dev:
	$(MAKE) -C ./snark-ffi/rust dev

.PHONY: server
server:
	$(GOCC) build $(GOFLAGS)

.PHONY: snark_ffi_release
snark_ffi_release:
	$(MAKE) -C ./snark-ffi/rust release

.PHONY: api
api:
	protoc --proto_path=./api \
		   --proto_path=./api/v1 \
		   --go_out=paths=source_relative:./api \
		   --go-grpc_out=paths=source_relative:./api \
		   --go-http_out=paths=source_relative:./api \
		   $(API_PROTO_FILES)

clean:
	$(MAKE) -C ./snark-ffi/rust clean
	rm -rf ./micro-snark-server
	rm ./api/*/*.go