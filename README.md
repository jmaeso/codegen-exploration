# codegen-exploration
This repository hosts some simple code since the focus will be in the details and the processes

## Hypothesis / Goal

Be able to **build APIs** with a schema-first approach, which would allow us to:

- Autogenerate most of the code (at least request, params and response types) in the "HTTP layer".
- Documentation always up-to-date by definition.
- Allow the service to serve its own documentation to facilitate integrations.
- From a single schema file.
- Eventually generate client code for that API.

## Process of the exploration

1. Defined [docs/openapi.yaml](https://github.com/jmaeso/codegen-exploration/blob/dccf58ac9fe26ac5930ecfabdca5eb300fa17426/docs/openapi.yaml) (using version 3.0.2).
2. Added `make openapi-gen` to generate go models for the API (using openapi 3.0.2).
3. Updated [docs/openapi.yaml](https://github.com/jmaeso/codegen-exploration/blob/44daf2c1bd5442601b00b3e9fd02ff4c1d5bec8a/docs/openapi.yaml) (now version 3.1.0).
4. Added `make openapi-lint` to validate new version is valid.
5. Added `make jsonschema-gen` (using https://github.com/atombender/go-jsonschema).

## Notes

- Adding `$schema: https://json-schema.org/draft/2020-12/schema` to the openapi.yaml (3.1.0) file avoids the `gojsonschema` generator to fail, but generates an empty go file.
- Interesting feature for `gojsonschema`: Be able to provide a custom type (and package) for types with format. It would be nice to have `UUID` or `ISO 3166` validations...
    - Idea: Support `x-go-customType: {package.type}` for properties so a custom type can be used for them. Example: 
    ```yaml
        id:
            x-go-customType: github.com/google/uuid.UUID
            description: Unique identifier of the user.
            type: string
            format: uuid
            example: e22896bb-de9c-47d1-b5ae-005c31d49a6f
    ```
