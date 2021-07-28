package asyncapi2_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/poim1205/asyncapi/asyncapi2"
	"gopkg.in/yaml.v2"
)

var _ = Describe("Operation", func() {
	Context("MapOperation interface", func() {
		var (
			mapOperation          *asyncapi2.Operation
			mapOperationInterface map[interface{}]interface{}
		)
		BeforeEach(func() {
			mapOperationInterface = make(map[interface{}]interface{})
			yamlOperationByte := []byte(`operationId: registerUser
summary: Action to sign a user up.
description: A longer description
tags:
  - name: user
  - name: signup
  - name: register
message:
  headers:
    type: object
    properties:
      applicationInstanceId:
        description: Unique identifier for a given instance of the publishing application
        type: string
  payload:
    type: object
    properties:
      user:
        $ref: "#/components/schemas/userCreate"
      signup:
        $ref: "#/components/schemas/signup"
bindings:
  amqp:
    ack: false
traits:
  - $ref: "#/components/operationTraits/kafka"`)

			err := yaml.Unmarshal(yamlOperationByte, &mapOperationInterface)
			Expect(err).ShouldNot(HaveOccurred())

			newMapOperation := asyncapi2.NewOperation()
			mapOperation = newMapOperation.SetValues(mapOperationInterface)
		})
		When("NewOperation is filled with a map structure", func() {
			It("should return a pointer to a structure of type Operation", func() {

				Expect(fmt.Sprintf("%T", mapOperation)).Should(Equal("*asyncapi2.Operation"))
			})
		})
		When("MapOperation is filled with a map structure", func() {
			It("should return a field `operationId` with value `registerUser`", func() {
				Expect(mapOperation.OperationId).Should(Equal("registerUser"))
			})
			It("should return a field `summary` with value `Action to sign a user up.`", func() {
				Expect(mapOperation.Summary).Should(Equal("Action to sign a user up."))
			})
			It("should return a field `description` with value `A longer description`", func() {
				Expect(mapOperation.Description).Should(Equal("A longer description"))
			})
			It("should return a field `headers` with sub field `type` value being `object`", func() {
				Expect(mapOperation.Message.Headers.Type).Should(Equal("object"))
			})
			It("should return a field `headers.properties[applicationInstanceId]` with sub field `type` value being `string`", func() {
				Expect(mapOperation.Message.Headers.Properties["applicationInstanceId"].Type).Should(Equal("string"))
			})
			It("should return a field `traits[0].$ref` with value `#/components/operationTraits/kafka`", func() {
				Expect(mapOperation.Traits[0].Ref).Should(Equal("#/components/operationTraits/kafka"))
			})
		})
	})
})

var _ = Describe("OperationTraits", func() {
	Context("MapOperation interface", func() {
		var (
			mapOperationTraits          asyncapi2.OperationTraits
			mapOperationTraitsInterface map[interface{}]interface{}
		)
		BeforeEach(func() {
			mapOperationTraitsInterface = make(map[interface{}]interface{})
			yamlOperationTraitsByte := []byte(`traits:
  - $ref: "#/components/operationTraits/kafka"
  - operationId: Demo operation ID
    summary: Action to sign a user up.
    description: A longer description
    tags:
      - name: tag 1
      - name: tag 2
    bindinds:
      amqp:
        ack: false`)

			err := yaml.Unmarshal(yamlOperationTraitsByte, &mapOperationTraitsInterface)
			Expect(err).ShouldNot(HaveOccurred())

			newMapOperationTraits := asyncapi2.NewOperationTraits()
			mapOperationTraits = newMapOperationTraits.SetValues(mapOperationTraitsInterface["traits"])
		})
		When("NewOperationTraits is filled with a map structure", func() {
			It("should return a structure of type OperationTraits", func() {

				Expect(fmt.Sprintf("%T", mapOperationTraits)).Should(Equal("asyncapi2.OperationTraits"))
			})
		})
		When("MapOperationTraits is filled with a map structure", func() {
			It("should return a field `traits[0].$ref` with value `#/components/operationTraits/kafka`", func() {
				Expect(mapOperationTraits[0].Ref).Should(Equal("#/components/operationTraits/kafka"))
			})
			It("should return a field `traits[1].operationId` with value `Demo operation ID`", func() {
				Expect(mapOperationTraits[1].OperationId).Should(Equal("Demo operation ID"))
			})
			It("should return a field `traits[1].summary` with value `Action to sign a user up.`", func() {
				Expect(mapOperationTraits[1].Summary).Should(Equal("Action to sign a user up."))
			})
			It("should return a field `traits[1].description` with value `A longer description`", func() {
				Expect(mapOperationTraits[1].Description).Should(Equal("A longer description"))
			})
			It("should return a field `traits[1].tags[0]` with sub field `name` value being `tag 1`", func() {
				Expect(mapOperationTraits[1].Tags[0].Name).Should(Equal("tag 1"))
			})
		})
	})
})
