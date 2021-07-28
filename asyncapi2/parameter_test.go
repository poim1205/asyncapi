package asyncapi2_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/poim1205/asyncapi/asyncapi2"
	"gopkg.in/yaml.v2"
)

var _ = Describe("Parameter", func() {
	Context("MapParameter interface", func() {
		var (
			mapParameters          asyncapi2.Parameters
			mapParametersInterface map[interface{}]interface{}
		)
		BeforeEach(func() {
			mapObjectInterface := make(map[interface{}]interface{})
			yamlParameterByte := []byte(`user/{userId}/signup:
  parameters:
    userId:
      description: Id of the user.
      schema:
        type: string
      location: $message.payload#/user/id
    subscribe:
      $ref: "#/components/messages/userSignedUp"`)

			err := yaml.Unmarshal(yamlParameterByte, &mapObjectInterface)
			Expect(err).ShouldNot(HaveOccurred())

			parametersInterface, _ := mapObjectInterface["user/{userId}/signup"]

			switch arrayVal := parametersInterface.(type) {
			case map[interface{}]interface{}:
				mapParametersInterface = arrayVal

			}

			newMapParameters := asyncapi2.NewParameters()
			mapParameters = newMapParameters.SetValues(mapParametersInterface["parameters"])
		})
		When("NewParameters is filled with a map structure", func() {
			It("should return a structure of type Parameters", func() {

				Expect(fmt.Sprintf("%T", mapParameters)).Should(Equal("asyncapi2.Parameters"))
			})
		})
		When("MapParameters is filled with a map structure", func() {
			It("should return 2 instances of Parameters", func() {
				Expect(len(mapParameters)).Should(Equal(2))
			})
			It("should return a Parameter named `userId`", func() {
				_, Ok := mapParameters["userId"]
				Expect(Ok).Should(Equal(true))
			})
			It("should return a Parameter named `subscribe`", func() {
				_, Ok := mapParameters["subscribe"]
				Expect(Ok).Should(Equal(true))
			})
			It("should return a description `Id of the user.` for parameter `userId`", func() {
				userIdParam, _ := mapParameters["userId"]
				Expect(userIdParam.Description).Should(Equal("Id of the user."))
			})
			It("should return a $ref `#/components/messages/userSignedUp` for parameter `subscribe`", func() {
				subscribeParam, _ := mapParameters["subscribe"]
				Expect(subscribeParam.Ref).Should(Equal("#/components/messages/userSignedUp"))
			})
			It("should return a schema.type `string` for parameter `userId`", func() {
				userIdParam, _ := mapParameters["userId"]
				Expect(userIdParam.Schema.Type).Should(Equal("string"))
			})
		})
	})
})
