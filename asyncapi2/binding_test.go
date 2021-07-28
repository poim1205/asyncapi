package asyncapi2_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/poim1205/asyncapi/asyncapi2"
	"gopkg.in/yaml.v2"
)

var _ = Describe("ChannelBindings", func() {
	Context("MapChannels-queue interface", func() {
		var (
			mapChannelBindings asyncapi2.ChannelBindings
			mapCBsInterface    map[interface{}]interface{}
		)
		BeforeEach(func() {
			mapCBsInterface = make(map[interface{}]interface{})
			yamlCBsByte := []byte(`bindings:
  ws:
    method: my-method
    query:
      type: string
    headers:
      type: object
    bindingVersion: 1.0.0
  amqp:
    is: queue
    queue:
      name: my-amqp-name
      durable: true
      exclusive: false
      autoDelete: true
      vhost: my-amqp-vhost
    bindingVersion: 2.0.0
  ibmmq:
    destinationType: queue
    queue:
      objectName: ibm-mq-object-name
      isPartitioned: true
      exclusive: false
    maxMsgLength: 100
    bindingVersion: 3.0.0`)

			err := yaml.Unmarshal(yamlCBsByte, &mapCBsInterface)
			Expect(err).ShouldNot(HaveOccurred())

			newMapChannelBindings := asyncapi2.NewChannelBindings()
			mapChannelBindings = newMapChannelBindings.SetValues(mapCBsInterface["bindings"])
		})
		When("NewChannelBindings is filled with a map structure", func() {
			It("should return a structure of type ChannelBindings", func() {
				Expect(fmt.Sprintf("%T", mapChannelBindings)).Should(Equal("asyncapi2.ChannelBindings"))
			})
		})
		When("MapChannelBindings is filled with a map structure", func() {
			It("should return a ChannelBinding named `ws`", func() {
				_, Ok := mapChannelBindings["ws"]
				Expect(Ok).Should(Equal(true))
			})
			It("should return a ChannelBinding named `amqp`", func() {
				_, Ok := mapChannelBindings["amqp"]
				Expect(Ok).Should(Equal(true))
			})
			It("should return a Channel named `ibmmq`", func() {
				_, Ok := mapChannelBindings["ibmmq"]
				Expect(Ok).Should(Equal(true))
			})
			It("should return a field `method` with value `my-method`", func() {
				binding, _ := mapChannelBindings["ws"]
				wsBinding := binding.(*asyncapi2.WsChannelBinding)

				Expect(wsBinding.Method).Should(Equal("my-method"))
			})
			It("should return a field `query.type` with value `string`", func() {
				binding, _ := mapChannelBindings["ws"]
				wsBinding := binding.(*asyncapi2.WsChannelBinding)

				Expect(wsBinding.Query.Type).Should(Equal("string"))
			})
			It("should return a field `bindingVersion` with value `1.0.0`", func() {
				binding, _ := mapChannelBindings["ws"]
				wsBinding := binding.(*asyncapi2.WsChannelBinding)

				Expect(wsBinding.BindingVersion).Should(Equal("1.0.0"))
			})
			It("should return a field `is` with value `queue`", func() {
				binding, _ := mapChannelBindings["amqp"]
				amqpBinding := binding.(*asyncapi2.AmqpChannelBinding)

				Expect(amqpBinding.Is).Should(Equal("queue"))
			})
			It("should return a field `queue.name` with value `my-amqp-name`", func() {
				binding, _ := mapChannelBindings["amqp"]
				amqpBinding := binding.(*asyncapi2.AmqpChannelBinding)
				amqpChannelSpec := amqpBinding.ChannelSpec["queue"].(asyncapi2.AmqpChannelSpecification)
				amqpChannelSpecQueue := amqpChannelSpec.(*asyncapi2.AmqpChannelQueue)

				Expect(amqpChannelSpecQueue.Name).Should(Equal("my-amqp-name"))
			})
			It("should return a field `queue.autoDelete` with value `true`", func() {
				binding, _ := mapChannelBindings["amqp"]
				amqpBinding := binding.(*asyncapi2.AmqpChannelBinding)
				amqpChannelSpec := amqpBinding.ChannelSpec["queue"].(asyncapi2.AmqpChannelSpecification)
				amqpChannelSpecQueue := amqpChannelSpec.(*asyncapi2.AmqpChannelQueue)

				Expect(amqpChannelSpecQueue.AutoDelete).Should(Equal(true))
			})
			It("should return a field `destinationType` with value `queue`", func() {
				binding, _ := mapChannelBindings["ibmmq"]
				ibmmqBinding := binding.(*asyncapi2.IbmmqChannelBinding)

				Expect(ibmmqBinding.DestinationType).Should(Equal("queue"))
			})
			It("should return a field `queue.objectName` with value `ibm-mq-object-name`", func() {
				binding, _ := mapChannelBindings["ibmmq"]
				ibmmqBinding := binding.(*asyncapi2.IbmmqChannelBinding)
				ibmmqChannelSpec := ibmmqBinding.ChannelSpec["queue"].(asyncapi2.IbmmqChannelSpecification)
				ibmmqChannelSpecQueue := ibmmqChannelSpec.(*asyncapi2.IbmmqChannelQueue)

				Expect(ibmmqChannelSpecQueue.ObjectName).Should(Equal("ibm-mq-object-name"))
			})
			It("should return a field `queue.exclusive` with value `false`", func() {
				binding, _ := mapChannelBindings["ibmmq"]
				ibmmqBinding := binding.(*asyncapi2.IbmmqChannelBinding)
				ibmmqChannelSpec := ibmmqBinding.ChannelSpec["queue"].(asyncapi2.IbmmqChannelSpecification)
				ibmmqChannelSpecQueue := ibmmqChannelSpec.(*asyncapi2.IbmmqChannelQueue)

				Expect(ibmmqChannelSpecQueue.Exclusive).Should(Equal(false))
			})
			It("should return a field `maxMsgLength` with value `100`", func() {
				binding, _ := mapChannelBindings["ibmmq"]
				ibmmqBinding := binding.(*asyncapi2.IbmmqChannelBinding)

				Expect(ibmmqBinding.MaxMsgLength).Should(Equal(100))
			})
			It("should return a field `bindingVersion` with value `3.0.0`", func() {
				binding, _ := mapChannelBindings["ibmmq"]
				ibmmqBinding := binding.(*asyncapi2.IbmmqChannelBinding)

				Expect(ibmmqBinding.BindingVersion).Should(Equal("3.0.0"))
			})
		})
	})

	Context("MapChannels-others interface", func() {
		var (
			mapChannelBindings asyncapi2.ChannelBindings
			mapCBsInterface    map[interface{}]interface{}
		)
		BeforeEach(func() {
			mapCBsInterface = make(map[interface{}]interface{})
			yamlCBsByte := []byte(`bindings:
  ws:
    method: my-method
    query:
      type: string
    headers:
      type: object
    bindingVersion: 1.0.0
  amqp:
    is: routingKey
    exchange:
      name: my-amqp-name
      type: topic
      durable: false
      autoDelete: true
      vhost: my-amqp-vhost
    bindingVersion: 2.0.0
  ibmmq:
    destinationType: topic
    topic:
      string: ibm-topic-string
      objectName: ibm-mq-object-name
      durablePermitted: true
      lastMsgRetained: false
    maxMsgLength: 100
    bindingVersion: 3.0.0`)

			err := yaml.Unmarshal(yamlCBsByte, &mapCBsInterface)
			Expect(err).ShouldNot(HaveOccurred())

			newMapChannelBindings := asyncapi2.NewChannelBindings()
			mapChannelBindings = newMapChannelBindings.SetValues(mapCBsInterface["bindings"])
		})
		When("NewChannelBindings is filled with a map structure", func() {
			It("should return a structure of type ChannelBindings", func() {
				Expect(fmt.Sprintf("%T", mapChannelBindings)).Should(Equal("asyncapi2.ChannelBindings"))
			})
		})
		When("MapChannelBindings is filled with a map structure", func() {
			It("should return a ChannelBinding named `ws`", func() {
				_, Ok := mapChannelBindings["ws"]
				Expect(Ok).Should(Equal(true))
			})
			It("should return a ChannelBinding named `amqp`", func() {
				_, Ok := mapChannelBindings["amqp"]
				Expect(Ok).Should(Equal(true))
			})
			It("should return a Channel named `ibmmq`", func() {
				_, Ok := mapChannelBindings["ibmmq"]
				Expect(Ok).Should(Equal(true))
			})
			It("should return a field `method` with value `my-method`", func() {
				binding, _ := mapChannelBindings["ws"]
				wsBinding := binding.(*asyncapi2.WsChannelBinding)

				Expect(wsBinding.Method).Should(Equal("my-method"))
			})
			It("should return a field `query.type` with value `string`", func() {
				binding, _ := mapChannelBindings["ws"]
				wsBinding := binding.(*asyncapi2.WsChannelBinding)

				Expect(wsBinding.Query.Type).Should(Equal("string"))
			})
			It("should return a field `bindingVersion` with value `1.0.0`", func() {
				binding, _ := mapChannelBindings["ws"]
				wsBinding := binding.(*asyncapi2.WsChannelBinding)

				Expect(wsBinding.BindingVersion).Should(Equal("1.0.0"))
			})
			It("should return a field `is` with value `routingKey`", func() {
				binding, _ := mapChannelBindings["amqp"]
				amqpBinding := binding.(*asyncapi2.AmqpChannelBinding)

				Expect(amqpBinding.Is).Should(Equal("routingKey"))
			})
			It("should return a field `exchange.name` with value `my-amqp-name`", func() {
				binding, _ := mapChannelBindings["amqp"]
				amqpBinding := binding.(*asyncapi2.AmqpChannelBinding)
				amqpChannelSpec := amqpBinding.ChannelSpec["exchange"].(asyncapi2.AmqpChannelSpecification)
				amqpChannelSpecExchange := amqpChannelSpec.(*asyncapi2.AmqpChannelExchange)

				Expect(amqpChannelSpecExchange.Name).Should(Equal("my-amqp-name"))
			})
			It("should return a field `exchange.autoDelete` with value `true`", func() {
				binding, _ := mapChannelBindings["amqp"]
				amqpBinding := binding.(*asyncapi2.AmqpChannelBinding)
				amqpChannelSpec := amqpBinding.ChannelSpec["exchange"].(asyncapi2.AmqpChannelSpecification)
				amqpChannelSpecExchange := amqpChannelSpec.(*asyncapi2.AmqpChannelExchange)

				Expect(amqpChannelSpecExchange.AutoDelete).Should(Equal(true))
			})
			It("should return a field `destinationType` with value `topic`", func() {
				binding, _ := mapChannelBindings["ibmmq"]
				ibmmqBinding := binding.(*asyncapi2.IbmmqChannelBinding)

				Expect(ibmmqBinding.DestinationType).Should(Equal("topic"))
			})
			It("should return a field `topic.topicString` with value `ibm-topic-string`", func() {
				binding, _ := mapChannelBindings["ibmmq"]
				ibmmqBinding := binding.(*asyncapi2.IbmmqChannelBinding)
				ibmmqChannelSpec := ibmmqBinding.ChannelSpec["topic"].(asyncapi2.IbmmqChannelSpecification)
				ibmmqChannelSpecTopic := ibmmqChannelSpec.(*asyncapi2.IbmmqChannelTopic)

				Expect(ibmmqChannelSpecTopic.TopicString).Should(Equal("ibm-topic-string"))
			})
			It("should return a field `queue.objectName` with value `ibm-mq-object-name`", func() {
				binding, _ := mapChannelBindings["ibmmq"]
				ibmmqBinding := binding.(*asyncapi2.IbmmqChannelBinding)
				ibmmqChannelSpec := ibmmqBinding.ChannelSpec["topic"].(asyncapi2.IbmmqChannelSpecification)
				ibmmqChannelSpecTopic := ibmmqChannelSpec.(*asyncapi2.IbmmqChannelTopic)

				Expect(ibmmqChannelSpecTopic.ObjectName).Should(Equal("ibm-mq-object-name"))
			})
			It("should return a field `topic.durablePermitted` with value `true`", func() {
				binding, _ := mapChannelBindings["ibmmq"]
				ibmmqBinding := binding.(*asyncapi2.IbmmqChannelBinding)
				ibmmqChannelSpec := ibmmqBinding.ChannelSpec["topic"].(asyncapi2.IbmmqChannelSpecification)
				ibmmqChannelSpecTopic := ibmmqChannelSpec.(*asyncapi2.IbmmqChannelTopic)

				Expect(ibmmqChannelSpecTopic.DurablePermitted).Should(Equal(true))
			})
			It("should return a field `topic.lastMsgRetained` with value `false`", func() {
				binding, _ := mapChannelBindings["ibmmq"]
				ibmmqBinding := binding.(*asyncapi2.IbmmqChannelBinding)
				ibmmqChannelSpec := ibmmqBinding.ChannelSpec["topic"].(asyncapi2.IbmmqChannelSpecification)
				ibmmqChannelSpecTopic := ibmmqChannelSpec.(*asyncapi2.IbmmqChannelTopic)

				Expect(ibmmqChannelSpecTopic.LastMsgRetained).Should(Equal(false))
			})
			It("should return a field `maxMsgLength` with value `100`", func() {
				binding, _ := mapChannelBindings["ibmmq"]
				ibmmqBinding := binding.(*asyncapi2.IbmmqChannelBinding)

				Expect(ibmmqBinding.MaxMsgLength).Should(Equal(100))
			})
			It("should return a field `bindingVersion` with value `3.0.0`", func() {
				binding, _ := mapChannelBindings["ibmmq"]
				ibmmqBinding := binding.(*asyncapi2.IbmmqChannelBinding)

				Expect(ibmmqBinding.BindingVersion).Should(Equal("3.0.0"))
			})
		})
	})
	Context("MapOperation interface", func() {
		var (
			mapOperationBindings asyncapi2.OperationBindings
			mapOBsInterface      map[interface{}]interface{}
		)
		BeforeEach(func() {
			mapOBsInterface = make(map[interface{}]interface{})
			yamlOBsByte := []byte(`bindings:
  http:
    type: request
    method: GET
    query:
      type: object
      required:
        - companyId
      properties:
        companyId:
          type: number
          minimum: 1
          description: The Id of the company.
      additionalProperties: false
    bindingVersion: '0.1.0'
  kafka:
    groupId:
      type: string
      enum: ['myGroupId']
    clientId:
      type: string
      enum: ['myClientId']
    bindingVersion: '0.1.0'
  amqp:
    expiration: 100000
    userId: guest
    cc: ['user.logs']
    priority: 10
    deliveryMode: 2
    mandatory: false
    bcc: ['external.audit']
    replyTo: user.signedup
    timestamp: true
    ack: false
    bindingVersion: 0.2.0
  mqtt:
    qos: 2
    retain: true
    bindingVersion: 0.1.0`)

			err := yaml.Unmarshal(yamlOBsByte, &mapOBsInterface)
			Expect(err).ShouldNot(HaveOccurred())

			newMapOperationBindings := asyncapi2.NewOperationBindings()
			mapOperationBindings = newMapOperationBindings.SetValues(mapOBsInterface["bindings"])
		})
		When("NewOperationBindings is filled with a map structure", func() {
			It("should return a structure of type OperationBindings", func() {
				Expect(fmt.Sprintf("%T", mapOperationBindings)).Should(Equal("asyncapi2.OperationBindings"))
			})
		})
		When("MapOperationBindings is filled with a map structure", func() {
			It("should return a OperationBinding named `http`", func() {
				_, Ok := mapOperationBindings["http"]
				Expect(Ok).Should(Equal(true))
			})
			It("should return a OperationBinding named `kafka`", func() {
				_, Ok := mapOperationBindings["kafka"]
				Expect(Ok).Should(Equal(true))
			})
			It("should return a ChannelBinding named `amqp`", func() {
				_, Ok := mapOperationBindings["amqp"]
				Expect(Ok).Should(Equal(true))
			})
			It("should return a Channel named `mqtt`", func() {
				_, Ok := mapOperationBindings["mqtt"]
				Expect(Ok).Should(Equal(true))
			})
			It("should return a field `type` with value `request`", func() {
				binding, _ := mapOperationBindings["http"]
				httpBinding := binding.(*asyncapi2.HttpOperationBinding)

				Expect(httpBinding.Type).Should(Equal("request"))
			})
			It("should return a field `query.type` with value `object`", func() {
				binding, _ := mapOperationBindings["http"]
				httpBinding := binding.(*asyncapi2.HttpOperationBinding)

				Expect(httpBinding.Query.Type).Should(Equal("object"))
			})
			It("should return a field `bindingVersion` with value `0.1.0`", func() {
				binding, _ := mapOperationBindings["http"]
				httpBinding := binding.(*asyncapi2.HttpOperationBinding)

				Expect(httpBinding.BindingVersion).Should(Equal("0.1.0"))
			})
			It("should return a field `groupId.type` with value `string`", func() {
				binding, _ := mapOperationBindings["kafka"]
				kafkaBinding := binding.(*asyncapi2.KafkaOperationBinding)

				Expect(kafkaBinding.GroupId.Type).Should(Equal("string"))
			})
			It("should return a field `clientId.enum[0]` with value `myClientId`", func() {
				binding, _ := mapOperationBindings["kafka"]
				kafkaBinding := binding.(*asyncapi2.KafkaOperationBinding)

				Expect(kafkaBinding.ClientId.Enum[0]).Should(Equal("myClientId"))
			})
			It("should return a field `expiration` with value `100000`", func() {
				binding, _ := mapOperationBindings["amqp"]
				amqpBinding := binding.(*asyncapi2.AmqpOperationBinding)

				Expect(amqpBinding.Expiration).Should(Equal(100000))
			})
			It("should return a field `cc[0]` with value `user.logs`", func() {
				binding, _ := mapOperationBindings["amqp"]
				amqpBinding := binding.(*asyncapi2.AmqpOperationBinding)

				Expect(amqpBinding.Cc[0]).Should(Equal("user.logs"))
			})
			It("should return a field `timestamp` with value `true`", func() {
				binding, _ := mapOperationBindings["amqp"]
				amqpBinding := binding.(*asyncapi2.AmqpOperationBinding)

				Expect(amqpBinding.Timestamp).Should(Equal(true))
			})
			It("should return a field `qos` with value `2`", func() {
				binding, _ := mapOperationBindings["mqtt"]
				mqttBinding := binding.(*asyncapi2.MqttOperationBinding)

				Expect(mqttBinding.Qos).Should(Equal(2))
			})
			It("should return a field `retain` with value `true`", func() {
				binding, _ := mapOperationBindings["mqtt"]
				mqttBinding := binding.(*asyncapi2.MqttOperationBinding)

				Expect(mqttBinding.Retain).Should(Equal(true))
			})
			It("should return a field `bindingVersion` with value `0.1.0`", func() {
				binding, _ := mapOperationBindings["mqtt"]
				mqttBinding := binding.(*asyncapi2.MqttOperationBinding)

				Expect(mqttBinding.BindingVersion).Should(Equal("0.1.0"))
			})
		})
	})
	Context("MapMessage interface", func() {
		var (
			mapMessageBindings asyncapi2.MessageBindings
			mapMBsInterface    map[interface{}]interface{}
		)
		BeforeEach(func() {
			mapMBsInterface = make(map[interface{}]interface{})
			yamlMBsByte := []byte(`bindings:
  http:
    headers:
      type: object
      properties:
        Content-Type:
          type: string
          enum: ['application/json']
    bindingVersion: '0.1.0'
  kafka:
    key:
      type: string
      enum: ['myKey']
    bindingVersion: '0.1.0'
  amqp:
    contentEncoding: gzip
    messageType: 'user.signup'
    bindingVersion: 0.2.0
  mqtt:
    bindingVersion: 0.1.0
  ibmmq:
    type: jms
    description: JMS stream message
    bindingVersion: 0.1.0`)

			err := yaml.Unmarshal(yamlMBsByte, &mapMBsInterface)
			Expect(err).ShouldNot(HaveOccurred())

			newMapMessageBindings := asyncapi2.NewMessageBindings()
			mapMessageBindings = newMapMessageBindings.SetValues(mapMBsInterface["bindings"])
		})
		When("NewMessageBindings is filled with a map structure", func() {
			It("should return a structure of type OperationBindings", func() {
				Expect(fmt.Sprintf("%T", mapMessageBindings)).Should(Equal("asyncapi2.MessageBindings"))
			})
		})
		When("MapMessageBindings is filled with a map structure", func() {
			It("should return a MessageBinding named `http`", func() {
				_, Ok := mapMessageBindings["http"]
				Expect(Ok).Should(Equal(true))
			})
			It("should return a MessageBinding named `kafka`", func() {
				_, Ok := mapMessageBindings["kafka"]
				Expect(Ok).Should(Equal(true))
			})
			It("should return a ChannelBinding named `amqp`", func() {
				_, Ok := mapMessageBindings["amqp"]
				Expect(Ok).Should(Equal(true))
			})
			It("should return a MessageBinding named `mqtt`", func() {
				_, Ok := mapMessageBindings["mqtt"]
				Expect(Ok).Should(Equal(true))
			})
			It("should return a MessageBinding named `ibmmq`", func() {
				_, Ok := mapMessageBindings["ibmmq"]
				Expect(Ok).Should(Equal(true))
			})
			It("should return a field `headers.type` with value `object`", func() {
				binding, _ := mapMessageBindings["http"]
				httpBinding := binding.(*asyncapi2.HttpMessageBinding)

				Expect(httpBinding.Headers.Type).Should(Equal("object"))
			})
			It("should return a field `bindingVersion` with value `0.1.0`", func() {
				binding, _ := mapMessageBindings["http"]
				httpBinding := binding.(*asyncapi2.HttpMessageBinding)

				Expect(httpBinding.BindingVersion).Should(Equal("0.1.0"))
			})
			It("should return a field `key.type` with value `string`", func() {
				binding, _ := mapMessageBindings["kafka"]
				kafkaBinding := binding.(*asyncapi2.KafkaMessageBinding)

				Expect(kafkaBinding.Key.Type).Should(Equal("string"))
			})
			It("should return a field `key.enum[0]` with value `myKey`", func() {
				binding, _ := mapMessageBindings["kafka"]
				kafkaBinding := binding.(*asyncapi2.KafkaMessageBinding)

				Expect(kafkaBinding.Key.Enum[0]).Should(Equal("myKey"))
			})
			It("should return a field `contentEncoding` with value `gzip`", func() {
				binding, _ := mapMessageBindings["amqp"]
				amqpBinding := binding.(*asyncapi2.AmqpMessageBinding)

				Expect(amqpBinding.ContentEncoding).Should(Equal("gzip"))
			})
			It("should return a field `timestamp` with value `true`", func() {
				binding, _ := mapMessageBindings["amqp"]
				amqpBinding := binding.(*asyncapi2.AmqpMessageBinding)

				Expect(amqpBinding.MessageType).Should(Equal("user.signup"))
			})
			It("should return a field `bindingVersion` with value `0.1.0`", func() {
				binding, _ := mapMessageBindings["mqtt"]
				mqttBinding := binding.(*asyncapi2.MqttMessageBinding)

				Expect(mqttBinding.BindingVersion).Should(Equal("0.1.0"))
			})
			It("should return a field `type` with value `jms`", func() {
				binding, _ := mapMessageBindings["ibmmq"]
				ibmmqBinding := binding.(*asyncapi2.IbmmqMessageBinding)

				Expect(ibmmqBinding.Type).Should(Equal("jms"))
			})
			It("should return a field `description` with value `JMS stream message`", func() {
				binding, _ := mapMessageBindings["ibmmq"]
				ibmmqBinding := binding.(*asyncapi2.IbmmqMessageBinding)

				Expect(ibmmqBinding.Description).Should(Equal("JMS stream message"))
			})
		})
	})
	Context("MapServer interface", func() {
		var (
			mapServerBindings asyncapi2.ServerBindings
			mapSBsInterface   map[interface{}]interface{}
		)
		BeforeEach(func() {
			mapSBsInterface = make(map[interface{}]interface{})
			yamlSBsByte := []byte(`bindings:
  mqtt:
    clientId: guest
    cleanSession: true
    lastWill:
      topic: /last-wills
      qos: 2
      message: Guest gone offline.
      retain: false
    keepAlive: 60
    bindingVersion: 0.1.0
  ibmmq:
    groupId: PRODCLSTR1
    cipherSpec: ANY_TLS12_OR_HIGHER
    bindingVersion: 0.1.0`)

			err := yaml.Unmarshal(yamlSBsByte, &mapSBsInterface)
			Expect(err).ShouldNot(HaveOccurred())

			newMapServerBindings := asyncapi2.NewServerBindings()
			mapServerBindings = newMapServerBindings.SetValues(mapSBsInterface["bindings"])
		})
		When("NewServerBindings is filled with a map structure", func() {
			It("should return a structure of type ServerBindings", func() {
				Expect(fmt.Sprintf("%T", mapServerBindings)).Should(Equal("asyncapi2.ServerBindings"))
			})
		})
		When("MapMessageBindings is filled with a map structure", func() {
			It("should return a MessageBinding named `mqtt`", func() {
				_, Ok := mapServerBindings["mqtt"]
				Expect(Ok).Should(Equal(true))
			})
			It("should return a MessageBinding named `ibmmq`", func() {
				_, Ok := mapServerBindings["ibmmq"]
				Expect(Ok).Should(Equal(true))
			})
			It("should return a field `clientId` with value `guest`", func() {
				binding, _ := mapServerBindings["mqtt"]
				mqttBinding := binding.(*asyncapi2.MqttServerBinding)

				Expect(mqttBinding.ClientId).Should(Equal("guest"))
			})
			It("should return a field `lastWill.message` with value `Guest gone offline.`", func() {
				binding, _ := mapServerBindings["mqtt"]
				mqttBinding := binding.(*asyncapi2.MqttServerBinding)

				Expect(mqttBinding.LastWill.Message).Should(Equal("Guest gone offline."))
			})
			It("should return a field `bindingVersion` with value `0.1.0`", func() {
				binding, _ := mapServerBindings["mqtt"]
				mqttBinding := binding.(*asyncapi2.MqttServerBinding)

				Expect(mqttBinding.BindingVersion).Should(Equal("0.1.0"))
			})
			It("should return a field `groupId` with value `PRODCLSTR1`", func() {
				binding, _ := mapServerBindings["ibmmq"]
				ibmmqBinding := binding.(*asyncapi2.IbmmqServerBinding)

				Expect(ibmmqBinding.GroupId).Should(Equal("PRODCLSTR1"))
			})
			It("should return a field `cipherSpec` with value `ANY_TLS12_OR_HIGHER`", func() {
				binding, _ := mapServerBindings["ibmmq"]
				ibmmqBinding := binding.(*asyncapi2.IbmmqServerBinding)

				Expect(ibmmqBinding.CipherSpec).Should(Equal("ANY_TLS12_OR_HIGHER"))
			})
		})
	})
})
