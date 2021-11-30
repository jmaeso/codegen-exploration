run:
	go run cmd/users-api/main.go

.PHONY: openapi-gen
openapi-gen:
	@echo "-----\ndeprecated command, openapi-generator does not support 3.1.0 (yet)\n---\n"
	docker run --rm -v "$(PWD)/docs:/spec" openapitools/openapi-generator-cli:v5.3.0 generate \
		--generator-name go \
		--input-spec spec/openapi.yaml \
		--output internal/openapi \
		--package-name openapi \
		--global-property models,modelDocs=false

.PHONY: openapi-lint
openapi-lint:
	docker run --rm -v $(PWD)/docs:/spec redocly/openapi-cli:v1.0.0-beta.67 lint openapi.yaml
