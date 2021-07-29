package asyncapi2_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/poim1205/asyncapi/asyncapi2"
	"gopkg.in/yaml.v2"
)

var _ = Describe("Components", func() {
	Context("MapComponents interface", func() {
		var (
			mapComponents          *asyncapi2.Components
			mapComponentsInterface map[interface{}]interface{}
		)
		BeforeEach(func() {
			mapComponentsInterface = make(map[interface{}]interface{})
			yamlComponentsByte := []byte(`components:
  messages:
    lightMeasured:
      name: lightMeasured
      title: Light measured
      summary: Inform about environmental lighting conditions of a particular streetlight.
      contentType: application/json
      traits:
        - $ref: '#/components/messageTraits/commonHeaders'
      payload:
        $ref: "#/components/schemas/lightMeasuredPayload"
    turnOnOff:
      name: turnOnOff
      title: Turn on/off
      summary: Command a particular streetlight to turn the lights on or off.
      traits:
        - $ref: '#/components/messageTraits/commonHeaders'
      payload:
        $ref: "#/components/schemas/turnOnOffPayload"
  schemas:
    lightMeasuredPayload:
      type: object
      properties:
        lumens:
          type: integer
          minimum: 0
          description: Light intensity measured in lumens.
        sentAt:
          $ref: "#/components/schemas/sentAt"
    turnOnOffPayload:
      type: object
      properties:
        command:
          type: string
          enum:
            - on
            - off
          description: Whether to turn on or off the light.
        sentAt:
          $ref: "#/components/schemas/sentAt"
    dimLightPayload:
      type: object
      properties:
        percentage:
          type: integer
          description: Percentage to which the light should be dimmed to.
          minimum: 0
          maximum: 100
        sentAt:
          $ref: "#/components/schemas/sentAt"
    sentAt:
      type: string
      format: date-time
      description: Date and time when the message was sent.
  correlationIds:
    demoCorrId:
      description: My demo description
      location: My demo location
  securitySchemes:
    saslScram:
      type: scramSha256
      description: Provide your username and password for SASL/SCRAM authentication
  parameters:
    streetlightId:
      description: The ID of the streetlight.
      schema:
        type: string
  messageTraits:
    commonHeaders:
      headers:
        type: object
        properties:
          my-app-header:
            type: integer
            minimum: 0
            maximum: 100
  operationTraits:
    kafka:
      bindings:
        kafka:
          clientId: my-app-id`)

			err := yaml.Unmarshal(yamlComponentsByte, &mapComponentsInterface)
			Expect(err).ShouldNot(HaveOccurred())

			newMapComponents := asyncapi2.NewComponents()
			mapComponents = newMapComponents.SetValues(mapComponentsInterface["components"])
		})
		When("NewComponents is filled with a map structure", func() {
			It("should return a structure of type Components", func() {

				Expect(fmt.Sprintf("%T", mapComponents)).Should(Equal("*asyncapi2.Components"))
			})
		})
		When("MapComponents is filled with a map structure", func() {
			It("should return a message named `lightMeasured`", func() {
				_, Ok := mapComponents.Messages["lightMeasured"]
				Expect(Ok).Should(Equal(true))
			})
			It("should return a Schema named `lightMeasuredPayload`", func() {
				_, Ok := mapComponents.Schemas["lightMeasuredPayload"]
				Expect(Ok).Should(Equal(true))
			})
			It("should return a correlationId named `demoCorrId`", func() {
				_, Ok := mapComponents.CorrelationIds["demoCorrId"]
				Expect(Ok).Should(Equal(true))
			})
			It("should return a securityScheme named `saslScram`", func() {
				_, Ok := mapComponents.SecuritySchemes["saslScram"]
				Expect(Ok).Should(Equal(true))
			})
			It("should return a parameter named `streetlightId`", func() {
				_, Ok := mapComponents.Parameters["streetlightId"]
				Expect(Ok).Should(Equal(true))
			})
			It("should return a messageTrait named `commonHeaders`", func() {
				_, Ok := mapComponents.MessageTraits["commonHeaders"]
				Expect(Ok).Should(Equal(true))
			})
			It("should return a operationTrait named `kafka`", func() {
				_, Ok := mapComponents.OperationTraits["kafka"]
				Expect(Ok).Should(Equal(true))
			})
		})
	})
	Context("MapComponents interface", func() {
		var (
			mapSecScheme          map[string]*asyncapi2.SecurityScheme
			mapSecSchemeInterface map[interface{}]interface{}
		)
		BeforeEach(func() {
			mapSecSchemeInterface = make(map[interface{}]interface{})
			yamlSecSchemeByte := []byte(`securitySchemes:
  saslScram:
    type: scramSha256
    description: Provide your username and password for SASL/SCRAM authentication
  myAuth:
    type: oauth2
    flows:
      implicit:
        authorizationUrl: https://example.com/api/oauth/dialog
        scopes:
          write:pets: modify pets in your account
          read:pets: read your pets
      authorizationCode:
        authorizationUrl: https://example.com/api/oauth/dialog
        tokenUrl: https://example.com/api/oauth/token
        scopes:
          write:pets: modify pets in your account
          read:pets: read your pets`)

			err := yaml.Unmarshal(yamlSecSchemeByte, &mapSecSchemeInterface)
			Expect(err).ShouldNot(HaveOccurred())

			switch mapSecuritySchemes := mapSecSchemeInterface["securitySchemes"].(type) {
			case map[interface{}]interface{}:
				secSchemeMap := make(map[string]*asyncapi2.SecurityScheme)
				for key, mapval := range mapSecuritySchemes {
					keySecSchemeStr := fmt.Sprintf("%v", key)
					_, Ok := secSchemeMap[keySecSchemeStr]
					if !Ok {
						secScheme := asyncapi2.NewSecurityScheme()
						secSchemeMap[keySecSchemeStr] = secScheme.SetValues(mapval)
					}
					mapSecScheme = secSchemeMap
				}
			default:
			}
		})
		When("NewSecurityScheme is filled with a map structure", func() {
			It("should return a structure of type SecurityScheme", func() {

				Expect(fmt.Sprintf("%T", mapSecScheme["saslScram"])).Should(Equal("*asyncapi2.SecurityScheme"))
			})
		})
		When("MapSecScheme is filled with a map structure", func() {
			It("should return a message named `saslScram`", func() {

				Expect(mapSecScheme["saslScram"].Type).Should(Equal("scramSha256"))
			})
		})
		When("MapSecScheme is filled with a map structure of type oauth2", func() {
			It("should return a message named `myAuth`", func() {
				_, Ok := mapSecScheme["myAuth"]
				Expect(Ok).Should(Equal(true))
			})
			It("should return a pointer to a flow structure", func() {
				Expect(fmt.Sprintf("%T", mapSecScheme["myAuth"].Flows)).Should(Equal("*asyncapi2.OAuthFlows"))
			})
			It("should return a field `write:pets` with a value of `modify pets in your account`", func() {
				Expect(mapSecScheme["myAuth"].Flows.AuthorizationCode.Scopes["write:pets"]).Should(Equal("modify pets in your account"))
			})
		})
	})
})
