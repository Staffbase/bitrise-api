.PHONY: all build deps download generate help test validate
CHECK_FILES?=$$(go list ./... | grep -v /vendor/)
SWAGGER_SPEC=swagger.json
PATCH_FILE=staffbase.patch
PATCHED_SWAGGER_SPEC=swagger_patched.json

help: ## Show this help.
		@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {sub("\\\\n",sprintf("\n%22c"," "), $$2);printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

all: download patch deps generate build ## Validate the swagger spec, generate the code and build it.

build: ## Build the API Go client.
	go build ./go/...

deps: ## Download dependencies.
	go install github.com/myitcv/gobin@v0.0.14 && go mod download

download: ## Download bitrise swagger specification
	wget -q -O $(SWAGGER_SPEC) https://api-docs.bitrise.io/docs/swagger.json

patch: ## Apply swagger.json patches
	patch -u $(SWAGGER_SPEC) -i $(PATCH_FILE) -o $(PATCHED_SWAGGER_SPEC)

generate: validate ## Generate the API Go client and the JSON document for the UI.
	go generate

validate: deps ## Check that the swagger spec is valid.
	gobin -m -run github.com/go-swagger/go-swagger/cmd/swagger@v0.26.1 validate $(PATCHED_SWAGGER_SPEC)