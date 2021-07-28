package asyncapi2_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/poim1205/asyncapi/asyncapi2"
	"gopkg.in/yaml.v2"
)

var _ = Describe("Server", func() {

	Context("MapServer interface", func() {
		var (
			mapServers         asyncapi2.Servers
			mapServerInterface map[interface{}]interface{}
		)
		BeforeEach(func() {
			mapServerInterface = make(map[interface{}]interface{})
			yamlServersByte := []byte(`servers:
  development:
    url: '{username}.gigantic-server.com:{port}/{basePath}'
    description: The development API server
    protocol: secure-mqtt
    variables:
      username:
        default: demo
        description: demo description
      port:
        enum:
          - '8883'
          - '8884'
        default: '8883'
      basePath:
        default: 'v2'
  staging:
    url: staging.gigantic-server.com
    description: Staging server
    protocol: amqp
    protocolVersion: 0.9.1
  production:
    url: api.gigantic-server.com
    description: Production server
    protocol: amqp
    protocolVersion: 0.9.1`)

			err := yaml.Unmarshal(yamlServersByte, &mapServerInterface)
			Expect(err).ShouldNot(HaveOccurred())

			newMapServers := asyncapi2.NewServers()
			mapServers = newMapServers.SetValues(mapServerInterface["servers"])

		})
		When("NewServers is filled with a map structure", func() {
			It("should return a structure of type Servers", func() {

				Expect(fmt.Sprintf("%T", mapServers)).Should(Equal("asyncapi2.Servers"))
			})
		})

		When("MapServers is filled with a map structure", func() {
			It("should return 3 instances of Server", func() {
				Expect(len(mapServers)).Should(Equal(3))
			})
			It("should return a Server named `development`", func() {
				_, Ok := mapServers["development"]
				Expect(Ok).Should(Equal(true))
			})
			It("should return a Server named `staging`", func() {
				_, Ok := mapServers["staging"]
				Expect(Ok).Should(Equal(true))
			})
			It("should return a Server named `production`", func() {
				_, Ok := mapServers["production"]
				Expect(Ok).Should(Equal(true))
			})
			It("should return an URL named `api.gigantic-server.com` for server `production`", func() {
				servProduction, _ := mapServers["production"]
				Expect(servProduction.Url).Should(Equal("api.gigantic-server.com"))
			})
			It("should return a description `Production server` for server `production`", func() {
				servProduction, _ := mapServers["production"]
				Expect(servProduction.Description).Should(Equal("Production server"))
			})
			It("should return a protocol named `amqp` for server `production`", func() {
				servProduction, _ := mapServers["production"]
				Expect(servProduction.Protocol).Should(Equal("amqp"))
			})
			It("should return a protocolVersion `0.9.1` for server `production`", func() {
				servProduction, _ := mapServers["production"]
				Expect(servProduction.ProtocolVersion).Should(Equal("0.9.1"))
			})
			It("should return a description `demo description` for the variable `username` for server `development`", func() {
				servDevelopment, _ := mapServers["development"]
				Expect(servDevelopment.Variables["username"].Description).Should(Equal("demo description"))
			})
			It("should return an enum `8883 and 8884` for the variable `port` for server `development`", func() {
				servDevelopment, _ := mapServers["development"]
				Expect(len(servDevelopment.Variables["port"].Enum)).Should(Equal(2))
			})
			It("should return an default `v2` for the variable `basePath` for server `development`", func() {
				servDevelopment, _ := mapServers["development"]
				Expect(servDevelopment.Variables["basePath"].Default).Should(Equal("v2"))
			})
		})
	})
})
