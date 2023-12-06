#Copyright 2021, Staffbase GmbH and contributors.
#
#Licensed under the Apache License, Version 2.0 (the "License");
#you may not use this file except in compliance with the License.
#
#You may obtain a copy of the License at
#    http://www.apache.org/licenses/LICENSE-2.0
#
#Unless required by applicable law or agreed to in writing, software
#distributed under the License is distributed on an "AS IS" BASIS,
#WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
#See the License for the specific language governing permissions and
#limitations under the License.

.PHONY: download generate build help
SWAGGER_SPEC=swagger.json
SWAGGER_URL=https://api-docs.bitrise.io/docs/swagger.json

help: ## Show this help.
		@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {sub("\\\\n",sprintf("\n%22c"," "), $$2);printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

all: download generate build ## Validate the swagger spec, generate the code and build it.

build: ## Build the API Go client.
	go build ./go/...

download: ## Download bitrise swagger specification
	wget -q -O $(SWAGGER_SPEC) $(SWAGGER_URL)

clean:
	rm -rf go/*

generate: clean   ## Generate the API Go client and the JSON document for the UI.
	go run github.com/go-swagger/go-swagger/cmd/swagger@v0.30.5 generate client \
	--name bitrise-api \
	--spec $(SWAGGER_SPEC) \
	--target go \
	--default-scheme=https \
	--with-flatten=full \
	--skip-tag-packages

validate: ## Check that the swagger spec is valid.
	go run github.com/go-swagger/go-swagger/cmd/swagger@v0.30.5 validate $(SWAGGER_SPEC)

