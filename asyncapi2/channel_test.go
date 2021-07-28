package asyncapi2_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/poim1205/asyncapi/asyncapi2"
	"gopkg.in/yaml.v2"
)

var _ = Describe("Channels", func() {
	Context("MapChannels interface", func() {
		var (
			mapChannels          asyncapi2.Channels
			mapChannelsInterface map[interface{}]interface{}
		)
		BeforeEach(func() {
			mapChannelsInterface = make(map[interface{}]interface{})
			yamlChannelsByte := []byte(`channels:
  smartylighting.streetlights.1.0.event.{streetlightId}.lighting.measured:
    description: The topic on which measured values may be produced and consumed.
    parameters:
      streetlightId:
        $ref: '#/components/parameters/streetlightId'
    publish:
      summary: Inform about environmental lighting conditions of a particular streetlight.
      operationId: receiveLightMeasurement
      traits:
        - $ref: '#/components/operationTraits/kafka'
      message:
        $ref: '#/components/messages/lightMeasured'
  smartylighting.streetlights.1.0.action.{streetlightId}.turn.on:
    parameters:
      streetlightId:
        $ref: '#/components/parameters/streetlightId'
    subscribe:
      operationId: turnOn
      traits:
        - $ref: '#/components/operationTraits/kafka'
      message:
        $ref: '#/components/messages/turnOnOff'
  smartylighting.streetlights.1.0.action.{streetlightId}.turn.off:
    $ref: '#/components/schemas/turnOnOff'`)

			err := yaml.Unmarshal(yamlChannelsByte, &mapChannelsInterface)
			Expect(err).ShouldNot(HaveOccurred())

			newMapChannels := asyncapi2.NewChannels()
			mapChannels = newMapChannels.SetValues(mapChannelsInterface["channels"])
		})
		When("NewChannels is filled with a map structure", func() {
			It("should return a structure of type Channels", func() {

				Expect(fmt.Sprintf("%T", mapChannels)).Should(Equal("asyncapi2.Channels"))
			})
		})
		When("MapChannels is filled with a map structure", func() {
			It("should return 2 instances of Channel", func() {
				Expect(len(mapChannels)).Should(Equal(3))
			})
			It("should return a Channel named `...lighting.measured`", func() {
				_, Ok := mapChannels["smartylighting.streetlights.1.0.event.{streetlightId}.lighting.measured"]
				Expect(Ok).Should(Equal(true))
			})
			It("should return a Channel named `...turn.on`", func() {
				_, Ok := mapChannels["smartylighting.streetlights.1.0.action.{streetlightId}.turn.on"]
				Expect(Ok).Should(Equal(true))
			})
			It("should return a Channel named `...turn.off`", func() {
				_, Ok := mapChannels["smartylighting.streetlights.1.0.action.{streetlightId}.turn.off"]
				Expect(Ok).Should(Equal(true))
			})
			It("should return a field `$ref` with value `#/components/schemas/turnOnOff`", func() {
				channelOff, _ := mapChannels["smartylighting.streetlights.1.0.action.{streetlightId}.turn.off"]
				Expect(channelOff.Ref).Should(Equal("#/components/schemas/turnOnOff"))
			})
			It("should return a field `description` with value `The topic on which measured values may be produced and consumed.`", func() {
				channelMeasure, _ := mapChannels["smartylighting.streetlights.1.0.event.{streetlightId}.lighting.measured"]
				Expect(channelMeasure.Description).Should(Equal("The topic on which measured values may be produced and consumed."))
			})
			It("should return 1 `parameters` ", func() {
				channelMeasure, _ := mapChannels["smartylighting.streetlights.1.0.event.{streetlightId}.lighting.measured"]
				Expect(len(channelMeasure.Parameters)).Should(Equal(1))
			})
			It("should return a structure `publish` for the measure channel", func() {
				channelMeasure, _ := mapChannels["smartylighting.streetlights.1.0.event.{streetlightId}.lighting.measured"]
				Expect(channelMeasure.Publish).ShouldNot(Equal(nil))
			})
			It("should not return a structure `subscribe` for the measure channel", func() {
				channelMeasure, _ := mapChannels["smartylighting.streetlights.1.0.event.{streetlightId}.lighting.measured"]
				var emptyOperation *asyncapi2.Operation
				Expect(channelMeasure.Subscribe).Should(Equal(emptyOperation))
			})
			It("should return a structure `Subscribe` for the turn.on channel", func() {
				channelOn, _ := mapChannels["smartylighting.streetlights.1.0.action.{streetlightId}.turn.on"]
				Expect(channelOn.Subscribe.OperationId).Should(Equal("turnOn"))
			})
		})
	})
})
