.DEFAULT_GOAL := all

SRC_FILES := $(shell git ls-files | grep -E "\.go$$" | grep -v -E "\.pb(:?\.gw)?\.go$$")
GO_TEST_FLAGS  := -v
GO_BUILD_FLAGS := -v -ldflags="-s -w"

DEP_COMMANDS := \
	github.com/golang/protobuf/protoc-gen-go \
	github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway \
	github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger \
	github.com/golang/lint/golint

define depcmdtmpl
$(1):
	@echo "Install $1"
	@cd vendor/$1 && GOBIN="$(shell pwd)/bin" go install .
endef

$(foreach cmd,$(DEP_COMMANDS),$(eval $(call depcmdtmpl,$(cmd))))

define section
  @printf "\e[34m--> $1\e[0m\n"
endef

#  App
#-----------------------------------------------
BIN := bin/app

$(BIN): $(SRC_FIELS)
	@echo "Building $(BIN)"
	@go build $(GO_BUILD_FLAGS) -o $(BIN) main.go


#  Commands
#-----------------------------------------------
.PHONY: all
all: $(BIN)

.PHONY: dep
dep: Gopkg.toml Gopkg.lock
	@dep ensure -v -vendor-only

.PHONY: cmds
cmds: $(DEP_COMMANDS)

.PHONY: gen
gen:
	$(call section,Generate)
	@PATH=$$PWD/bin:$$PATH go generate ./...

.PHONY: setup
setup: dep cmds gen

.PHONY: lint
lint:
	$(call section,Lint)
	@gofmt -e -d -s $(SRC_FILES) | awk '{ e = 1; print $0 } END { if (e) exit(1) }'
	@echo $(SRC_FILES) | xargs -n1 golint -set_exit_status
	@go vet ./...

.PHONY: test
test: gen lint
	$(call section,Test)
	@go test $(GO_TEST_FLAGS) ./...

.PHONY: ci-test
ci-test: lint
	$(call section,Test)
	@go test $(GO_TEST_FLAGS) -race ./...

.PHONY: run
run: gen $(BIN)
	$(call section,Start server)
	@$(BIN)
