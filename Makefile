.PHONY: fandom fandom-cross evm all test clean
.PHONY: fandom-linux fandom-linux-386 fandom-linux-amd64 fandom-linux-mips64 fandom-linux-mips64le
.PHONY: fandom-darwin fandom-darwin-386 fandom-darwin-amd64

GOBIN = $(shell pwd)/build/bin
GOFMT = gofmt
GO ?= latest
GO_PACKAGES = .
GO_FILES := $(shell find $(shell go list -f '{{.Dir}}' $(GO_PACKAGES)) -name \*.go)

GIT = git

fandom:
	build/env.sh go run build/ci.go install ./cmd/fandom
	@echo "Done building."
	@echo "Run \"$(GOBIN)/fandom\" to launch fandom."

gc:
	build/env.sh go run build/ci.go install ./cmd/gc
	@echo "Done building."
	@echo "Run \"$(GOBIN)/gc\" to launch gc."

bootnode:
	build/env.sh go run build/ci.go install ./cmd/bootnode
	@echo "Done building."
	@echo "Run \"$(GOBIN)/bootnode\" to launch a bootnode."

puppeth:
	build/env.sh go run build/ci.go install ./cmd/puppeth
	@echo "Done building."
	@echo "Run \"$(GOBIN)/puppeth\" to launch puppeth."

all:
	build/env.sh go run build/ci.go install

test: all
	build/env.sh go run build/ci.go test

clean:
	rm -fr build/_workspace/pkg/ $(GOBIN)/*

# Cross Compilation Targets (xgo)

fandom-cross: fandom-windows-amd64 fandom-darwin-amd64 fandom-linux
	@echo "Full cross compilation done:"
	@ls -ld $(GOBIN)/fandom-*

fandom-linux: fandom-linux-386 fandom-linux-amd64 fandom-linux-mips64 fandom-linux-mips64le
	@echo "Linux cross compilation done:"
	@ls -ld $(GOBIN)/fandom-linux-*

fandom-linux-386:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/386 -v ./cmd/fandom
	@echo "Linux 386 cross compilation done:"
	@ls -ld $(GOBIN)/fandom-linux-* | grep 386

fandom-linux-amd64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/amd64 -v ./cmd/fandom
	@echo "Linux amd64 cross compilation done:"
	@ls -ld $(GOBIN)/fandom-linux-* | grep amd64

fandom-linux-mips:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips --ldflags '-extldflags "-static"' -v ./cmd/fandom
	@echo "Linux MIPS cross compilation done:"
	@ls -ld $(GOBIN)/fandom-linux-* | grep mips

fandom-linux-mipsle:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mipsle --ldflags '-extldflags "-static"' -v ./cmd/fandom
	@echo "Linux MIPSle cross compilation done:"
	@ls -ld $(GOBIN)/fandom-linux-* | grep mipsle

fandom-linux-mips64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips64 --ldflags '-extldflags "-static"' -v ./cmd/fandom
	@echo "Linux MIPS64 cross compilation done:"
	@ls -ld $(GOBIN)/fandom-linux-* | grep mips64

fandom-linux-mips64le:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips64le --ldflags '-extldflags "-static"' -v ./cmd/fandom
	@echo "Linux MIPS64le cross compilation done:"
	@ls -ld $(GOBIN)/fandom-linux-* | grep mips64le

fandom-darwin: fandom-darwin-386 fandom-darwin-amd64
	@echo "Darwin cross compilation done:"
	@ls -ld $(GOBIN)/fandom-darwin-*

fandom-darwin-386:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=darwin/386 -v ./cmd/fandom
	@echo "Darwin 386 cross compilation done:"
	@ls -ld $(GOBIN)/fandom-darwin-* | grep 386

fandom-darwin-amd64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=darwin/amd64 -v ./cmd/fandom
	@echo "Darwin amd64 cross compilation done:"
	@ls -ld $(GOBIN)/fandom-darwin-* | grep amd64

fandom-windows-amd64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=windows/amd64 -v ./cmd/fandom
	@echo "Darwin amd64 cross compilation done:"
	@ls -ld $(GOBIN)/fandom-windows-* | grep amd64
gofmt:
	$(GOFMT) -s -w $(GO_FILES)
	$(GIT) checkout vendor
