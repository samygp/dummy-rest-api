PREFIX?=$(shell pwd)
NAME := "dummy_rest_api"
PKG := bitbucket.org/volteo/$(NAME)
BUILDDIR := ${PREFIX}/dist
ENTRYPOINT := cmd/main.go
CTIMEVAR=-X $(PKG)/version.Name=$(NAME) -X $(PKG)/version.Version=$(VERSION)
GO_LDFLAGS=-ldflags "$(CTIMEVAR)"
GO_LDFLAGS_STATIC=-ldflags "-w $(CTIMEVAR) -extldflags -static"
SEMBUMP_IMG=chatu/sembump:0.1.0
SEMBUMP=docker run --rm -v `pwd`:/app $(SEMBUMP_IMG)

.PHONY: default
default: help

.PHONY: name
name: ## Output name of project
	@echo $(NAME)

.PHONY: version
version: ## Output current version
	@echo $(VERSION)

.PHONY: install-deps
install-deps: ## Install dependencies
	@echo "+ $@"
	@go get -u golang.org/x/lint/golint
	@go get -u github.com/kisielk/errcheck
	@go get -u honnef.co/go/tools/cmd/staticcheck
	@go get -u github.com/aws/aws-lambda-go/cmd/build-lambda-zip

.PHONY: local-build
local-build: ## Builds a dynamic executable or package
	@echo "+ $@"
	@go build \
	  ${GO_LDFLAGS} -o $(NAME) \
	  $(ENTRYPOINT)

.PROXY: run
run: ## Execute built binary as web
	@echo "+ $@"
	@$(shell grep -v ^# config.env | xargs) ./$(NAME)

.PHONY: docker-build
docker-build: ## Build docker image
	@echo "+ $@"
	@docker $$DOCKER_CONFIG build -t $(NAME) .

.PHONY: docker-run
docker-run: #docker-build ## Execute built docker image
	@echo "+ $@"
	@docker run -it --privileged --rm
	 $(NAME)
