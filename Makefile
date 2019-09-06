PREFIX?=$(shell pwd)
NAME := "dummy-rest-api"
PKG := github.com/samygp/$(NAME)
BUILDDIR := ${PREFIX}/dist
ENTRYPOINT := cmd/main.go
CTIMEVAR=-X $(PKG)/version.Name=$(NAME) -X $(PKG)/version.Version=$(VERSION)
GO_LDFLAGS=-ldflags "$(CTIMEVAR)"
GO_LDFLAGS_STATIC=-ldflags "-w $(CTIMEVAR) -extldflags -static"

.PHONY: default
default: help

.PHONY: name
name: ## Output name of project
	@echo $(NAME)

.PHONY: version
version: ## Output current version
	@echo $(VERSION)

.PHONY: local-build
local-build: ## Builds a dynamic executable or package
	@echo "+ $@"
	@go build \
	  ${GO_LDFLAGS} -o $(NAME) \
	  $(ENTRYPOINT)

.PROXY: static
static: ## Generate static binary
	@echo "+ $@"
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build \
	  -o $(NAME) \
	  -a -tags "static_build netgo" \
	  -installsuffix netgo ${GO_LDFLAGS_STATIC} \
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
