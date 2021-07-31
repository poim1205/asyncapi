package asyncapi2_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/poim1205/asyncapi/asyncapi2"
	"gopkg.in/yaml.v2"
)

var _ = Describe("Message", func() {
	Context("MapMessage interface", func() {
		var (
			mapMessage          *asyncapi2.Message
			mapMessageInterface map[interface{}]interface{}
		)
		BeforeEach(func() {
			mapMessageInterface = make(map[interface{}]interface{})
			yamlMessageByte := []byte(`name: UserSignup
title: User signup
summary: Action to sign a user up.
description: A longer description
contentType: application/json
tags:
  - name: user
  - name: signup
  - name: register
headers:
  type: object
  properties:
    correlationId:
      description: Correlation ID set by application
      type: string
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
correlationId:
  description: Default Correlation ID
  location: $message.header#/correlationId
traits:
  - $ref: "#/components/messageTraits/commonHeaders"
examples:
  - name: SimpleSignup
    summary: A simple UserSignup example message
    headers:
      correlationId: my-correlation-id
      applicationInstanceId: myInstanceId
    payload:
      user:
        someUserKey: someUserValue
      signup:
        someSignupKey: someSignupValue`)

			err := yaml.Unmarshal(yamlMessageByte, &mapMessageInterface)
			Expect(err).ShouldNot(HaveOccurred())

			newMapMessage := asyncapi2.NewMessage()
			mapMessage = newMapMessage.SetValues(mapMessageInterface)
		})
		When("NewMessage is filled with a map structure", func() {
			It("should return a pointer to a structure of type Message", func() {

				Expect(fmt.Sprintf("%T", mapMessage)).Should(Equal("*asyncapi2.Message"))
			})
		})
		When("MapMessage is filled with a map structure", func() {
			It("should return a field `name` with value `UserSignup`", func() {
				Expect(mapMessage.Name).Should(Equal("UserSignup"))
			})
			It("should return a field `title` with value `User signup`", func() {
				Expect(mapMessage.Title).Should(Equal("User signup"))
			})
			It("should return a field `summary` with value `Action to sign a user up.`", func() {
				Expect(mapMessage.Summary).Should(Equal("Action to sign a user up."))
			})
			It("should return a field `description` with value `A longer description`", func() {
				Expect(mapMessage.Description).Should(Equal("A longer description"))
			})
			It("should return a field `contentType` with value `application/json`", func() {
				Expect(mapMessage.ContentType).Should(Equal("application/json"))
			})
			It("should return a field `tags` with first tag `name` value being `user`", func() {
				Expect(mapMessage.Tags[0].Name).Should(Equal("user"))
			})
			It("should return a field `headers` with sub field `type` value being `object`", func() {
				Expect(mapMessage.Headers.Type).Should(Equal("object"))
			})
			It("should return a field `payload` is not empty", func() {
				Expect(fmt.Sprintf("%T", mapMessage.Payload)).Should(Equal("map[interface {}]interface {}"))
			})
			It("should return a field `correlationId` with sub field `description` value being `Default Correlation ID`", func() {
				Expect(mapMessage.CorrelationId.Description).Should(Equal("Default Correlation ID"))
			})
			It("should return a field `traits` with first trait `$ref` value being `#/components/messageTraits/commonHeaders`", func() {
				Expect(mapMessage.Traits[0].Ref).Should(Equal("#/components/messageTraits/commonHeaders"))
			})
		})
	})
})

var _ = Describe("MessageTrait", func() {
	Context("MapMessage interface", func() {
		var (
			mapMessageTraits          asyncapi2.MessageTraits
			mapMessageTraitsInterface map[interface{}]interface{}
		)
		BeforeEach(func() {
			mapMessageTraitsInterface = make(map[interface{}]interface{})
			yamlMessageTraitsByte := []byte(`traits:
  - $ref: "#/components/messageTraits/commonHeaders"
  - headers:
      type: object
      properties:
        correlationId:
          description: Properties correlation ID set by application
          type: string
        applicationInstanceId:
          description: Unique identifier for a given instance of the publishing application
          type: string
    correlationId:
      description: Message trait correlation ID
      location: somewhere
    schemaFormat: Demo Schema format
    contentType: Demo content type
    name: Demo name
    title: Demo title
    summary: Demo summary
    description: Demo description
    tags:
      - name: tag 1
      - name: tag 2`)
			// TODO: Bindings...

			err := yaml.Unmarshal(yamlMessageTraitsByte, &mapMessageTraitsInterface)
			Expect(err).ShouldNot(HaveOccurred())

			newMapMessageTraits := asyncapi2.NewMessageTraits()
			mapMessageTraits = newMapMessageTraits.SetValues(mapMessageTraitsInterface["traits"])
		})
		When("NewMessageTraits is filled with a map structure", func() {
			It("should return a structure of type MessageTraits", func() {

				Expect(fmt.Sprintf("%T", mapMessageTraits)).Should(Equal("asyncapi2.MessageTraits"))
			})
		})
		When("MapMessageTraits is filled with a map structure", func() {
			It("should return 2 instances of MessageTrait", func() {
				Expect(len(mapMessageTraits)).Should(Equal(2))
			})
			It("should return a field `$ref` for the first item with the value `#/components/messageTraits/commonHeaders`", func() {
				Expect(mapMessageTraits[0].Ref).Should(Equal("#/components/messageTraits/commonHeaders"))
			})
			It("should return a field `headers.type` with value `object`", func() {
				Expect(mapMessageTraits[1].Headers.Type).Should(Equal("object"))
			})
			It("should return a field `correlationId.description` with value `Message trait correlation ID`", func() {
				Expect(mapMessageTraits[1].CorrelationId.Description).Should(Equal("Message trait correlation ID"))
			})
			It("should return a field `schemaFormat` with value `Demo Schema format`", func() {
				Expect(mapMessageTraits[1].SchemaFormat).Should(Equal("Demo Schema format"))
			})
			It("should return a field `contentType` with value `Demo content type`", func() {
				Expect(mapMessageTraits[1].ContentType).Should(Equal("Demo content type"))
			})
			It("should return a field `name` with value `Demo name`", func() {
				Expect(mapMessageTraits[1].Name).Should(Equal("Demo name"))
			})
			It("should return a field `title` with value `Demo title`", func() {
				Expect(mapMessageTraits[1].Title).Should(Equal("Demo title"))
			})
			It("should return a field `summary` with value `Demo summary`", func() {
				Expect(mapMessageTraits[1].Summary).Should(Equal("Demo summary"))
			})
			It("should return a field `description` with value `Demo description`", func() {
				Expect(mapMessageTraits[1].Description).Should(Equal("Demo description"))
			})
			It("should return a field `tags->name` with first trait `name` value being `tag 1`", func() {
				Expect(mapMessageTraits[1].Tags[0].Name).Should(Equal("tag 1"))
			})
		})
	})
})

var _ = Describe("CorrelationId", func() {
	Context("MapCorrelationId interface", func() {
		var (
			mapCorrelationId          *asyncapi2.CorrelationID
			mapCorrelationIdInterface map[interface{}]interface{}
		)
		BeforeEach(func() {
			mapCorrelationIdInterface = make(map[interface{}]interface{})
			yamlCorrelationIdByte := []byte(`correlationId:
  description: Correlation ID set by application
  location: Location of Correlation`)

			err := yaml.Unmarshal(yamlCorrelationIdByte, &mapCorrelationIdInterface)
			Expect(err).ShouldNot(HaveOccurred())

			newMapCorrelationId := asyncapi2.NewCorrelationId()
			mapCorrelationId = newMapCorrelationId.SetValues(mapCorrelationIdInterface["correlationId"])
		})
		When("NewCorrelationId is filled with a map structure", func() {
			It("should return a pointer to a structure of type CorrelationId", func() {

				Expect(fmt.Sprintf("%T", mapCorrelationId)).Should(Equal("*asyncapi2.CorrelationID"))
			})
		})
		When("MapMessageTraits is filled with a map structure", func() {
			It("should return a field `description` with the value `Correlation ID set by application`", func() {
				Expect(mapCorrelationId.Description).Should(Equal("Correlation ID set by application"))
			})
			It("should return a field `location` with the value `Location of Correlation`", func() {
				Expect(mapCorrelationId.Location).Should(Equal("Location of Correlation"))
			})
		})
	})
})
