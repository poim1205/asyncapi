package asyncapi2_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/poim1205/asyncapi/asyncapi2"
)

var _ = Describe("Asyncapi2", func() {
	Context("NewAsyncAPI()", func() {
		When("returns a structure of type *asyncapi2.T", func() {
			It("should be `true`", func() {
				asyncApiStruct := asyncapi2.NewAsyncAPI()
				strTypeOfAsyncApiStruct := fmt.Sprintf("%T", asyncApiStruct)

				Expect(strTypeOfAsyncApiStruct).Should(Equal("*asyncapi2.T"))
			})
		})
	})

	Context("Validate structure", func() {
		When("checking for all necessary structure to be filled", func() {
			It("should return no error", func() {
				minimalYamlByteArray := []byte(`asyncapi: '2.1.0'
id: 'urn:com:gitter:streaming:api'
info:
  title: Gitter Streaming API
  version: '1.0.0'

servers:
  production:
    url: https://stream.gitter.im/v1
    protocol: https
    protocolVersion: '1.1'
    security:
      - httpBearerToken: []

channels:
  /rooms/{roomId}/{resource}:
    parameters:
      roomId:
        description: Id of the Gitter room.
        schema:
          type: string
          examples:
            - 53307860c3599d1de448e19d
      resource:
        description: The resource to consume.
        schema:
          type: string
          enum:
            - chatMessages
            - events
    subscribe:
      bindings:
        http:
          type: response
      message:
        oneOf:
          - $ref: '#/components/messages/chatMessage'
          - $ref: '#/components/messages/heartbeat'
`)
				asyncApiLoader := asyncapi2.NewLoader()

				asyncApiStruct, err := asyncApiLoader.LoadFromData(minimalYamlByteArray)
				Expect(err).ShouldNot(HaveOccurred())

				err = asyncapi2.ValidateStruct(asyncApiStruct)
				Expect(err).ShouldNot(HaveOccurred())
			})
		})
		When("checking for all necessary structure to be filled but missing info", func() {
			It("should return no error", func() {
				minimalYamlByteArray := []byte(`asyncapi: '2.1.0'
id: 'urn:com:gitter:streaming:api'

servers:
  production:
    url: https://stream.gitter.im/v1
    protocol: https
    protocolVersion: '1.1'
    security:
      - httpBearerToken: []

channels:
  /rooms/{roomId}/{resource}:
    parameters:
      roomId:
        description: Id of the Gitter room.
        schema:
          type: string
          examples:
            - 53307860c3599d1de448e19d
      resource:
        description: The resource to consume.
        schema:
          type: string
          enum:
            - chatMessages
            - events
    subscribe:
      bindings:
        http:
          type: response
      message:
        oneOf:
          - $ref: '#/components/messages/chatMessage'
          - $ref: '#/components/messages/heartbeat'
`)
				asyncApiLoader := asyncapi2.NewLoader()

				asyncApiStruct, err := asyncApiLoader.LoadFromData(minimalYamlByteArray)
				Expect(err).ShouldNot(HaveOccurred())

				err = asyncapi2.ValidateStruct(asyncApiStruct)
				Expect(err).Should(HaveOccurred())
			})
		})
		When("checking for all necessary structure to be filled but wrong version format", func() {
			It("should return no error", func() {
				minimalYamlByteArray := []byte(`asyncapi: 'vers20'
id: 'urn:com:gitter:streaming:api'

servers:
  production:
    url: https://stream.gitter.im/v1
    protocol: https
    protocolVersion: '1.1'
    security:
      - httpBearerToken: []
info:
  title: Gitter Streaming API
  version: '1.0.0'
channels:
  /rooms/{roomId}/{resource}:
    parameters:
      roomId:
        description: Id of the Gitter room.
        schema:
          type: string
          examples:
            - 53307860c3599d1de448e19d
      resource:
        description: The resource to consume.
        schema:
          type: string
          enum:
            - chatMessages
            - events
    subscribe:
      bindings:
        http:
          type: response
      message:
        oneOf:
          - $ref: '#/components/messages/chatMessage'
          - $ref: '#/components/messages/heartbeat'
`)
				asyncApiLoader := asyncapi2.NewLoader()

				asyncApiStruct, err := asyncApiLoader.LoadFromData(minimalYamlByteArray)
				Expect(err).ShouldNot(HaveOccurred())

				err = asyncapi2.ValidateStruct(asyncApiStruct)
				Expect(err).Should(HaveOccurred())
			})
		})
		When("checking for all necessary structure to be filled but missing required from info", func() {
			It("should return no error", func() {
				minimalYamlByteArray := []byte(`asyncapi: '2.1.0'
id: 'urn:com:gitter:streaming:api'

servers:
  production:
    url: https://stream.gitter.im/v1
    protocol: https
    protocolVersion: '1.1'
    security:
      - httpBearerToken: []
info:
  title: Gitter Streaming API

channels:
  /rooms/{roomId}/{resource}:
    parameters:
      roomId:
        description: Id of the Gitter room.
        schema:
          type: string
          examples:
            - 53307860c3599d1de448e19d
      resource:
        description: The resource to consume.
        schema:
          type: string
          enum:
            - chatMessages
            - events
    subscribe:
      bindings:
        http:
          type: response
      message:
        oneOf:
          - $ref: '#/components/messages/chatMessage'
          - $ref: '#/components/messages/heartbeat'
`)
				asyncApiLoader := asyncapi2.NewLoader()

				asyncApiStruct, err := asyncApiLoader.LoadFromData(minimalYamlByteArray)
				Expect(err).ShouldNot(HaveOccurred())

				err = asyncapi2.ValidateStruct(asyncApiStruct)
				Expect(err).Should(HaveOccurred())
			})
		})
	})
})
