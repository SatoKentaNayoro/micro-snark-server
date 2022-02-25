SHELL=/usr/bin/env bash
INTERNAL_PROTO_FILES=$(shell find internal -name *.proto)
API_PROTO_FILES=$(shell find api -name *.proto)

.PHONY: dev
dev: init api snark_ffi_dev server

.PHONY: release
release: init api snark_ffi_release server

unexport GOFLAGS

ldflags=-X=micro-snark-server/build.CurrentCommit=git.$(subst -,.,$(shell git describe --always --match=NeVeRmAtCh --dirty 2>/dev/null || git rev-parse --short HEAD 2>/dev/null))
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
	$(MAKE) -C ./internal/snark-ffi/rust dev

.PHONY: server
server:
	$(GOCC) build $(GOFLAGS) -o ./bins/micro-snark-server ./cmd/snark_server

.PHONY: snark_ffi_release
snark_ffi_release:
	$(MAKE) -C ./internal/snark-ffi/rust release

.PHONY: api
api:
	protoc --proto_path=. \
		   --proto_path=./third_party \
		   --go_out=paths=source_relative:. \
		   --go-http_out=paths=source_relative:. \
		   --go-grpc_out=paths=source_relative:. \
		   $(API_PROTO_FILES)

clean_rust:
	$(MAKE) -C ./internal/snark-ffi/rust clean

clean_pb_go:
	rm ./api/*/*.go

bins:
	rm -rf ./bins/micro-snark-server

clean: clean_rust clean_pb_go bins