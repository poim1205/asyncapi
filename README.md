# Introduction

A [Go](https://golang.org) project for handling [AsyncAPI](https://github.com/asyncapi/spec/blob/v2.1.0/spec/asyncapi.md) files. We target the latest AsyncAPI version (currently 2.1.0), it remains to be tested with the previous version 2.0.0 files.

## Contributors and users

So far I've been the only contributor for this project. But it was greatly inspired by [_kin-openapi_](https://github.com/getkin/kin-openapi)

Here's some of the librairies I use in my project:

- [Validator](github.com/go-playground/validator/v10)  
- [Ginkgo as the testing framework](github.com/onsi/ginkgo)  
- [Omega for assertions](github.com/onsi/gomega)  
- [Logrus as logging framework](github.com/sirupsen/logrus)
- [Yaml](gopkg.in/yaml.v2)

## Alternatives

- Best I found was in javascript... which led me to create this projet

## Structure

- _asyncapi2_ (TODO: Create the go doc)
  - Support for AsyncAPI 2.1.0 files, including deserialization and validation.  

## How to use it

## Loading AsyncAPI document

Use `asyncapi2.Loader`, which provided support for json or yaml file and byte slice:

```go  
doc, err := asyncapi2.NewLoader().LoadFromFile("swagger.json")
```

## Getting OpenAPI operation that matches request

```go
loader := asyncapi2.NewLoader()
asyncApiStruct, _ := loader.LoadFromData([]byte(`...`))
// or asyncApiStruct, _ := loader.LoadFromFile("async.json")
_ := asyncApiStruct.Validate()
```

## Current version

### v1.0.0

- Initial official release
