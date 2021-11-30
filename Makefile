run:
	go run cmd/users-api/main.go

.PHONY: openapi-gen
openapi-gen:
	openapi-generator generate \
	--generator-name go \
	--input-spec docs/openapi.yaml \
	--output internal/openapi \
	--package-name openapi \
	--global-property models,modelDocs=false
