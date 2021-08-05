package asyncapi2

import (
	"fmt"
	"strconv"
)

//                     *         *                  *
// Binding  | Operation | Channel | Message | Server
// http     |    yes    |   no    |   yes   |   no
// ws       |    no     |   yes   |   no    |   no
// kafka    |    yes    |   no    |   yes   |   no
// amqp     |    yes    |   yes   |   yes   |   no
// amqp1    |    no     |   no    |   no    |   no
// mqtt     |    yes    |   no    |   yes   |   yes
// mqtt5    |    no     |   no    |   no    |   no
// nats     |    no     |   no    |   no    |   no
// jms      |    no     |   no    |   no    |   no
// sns      |    no     |   no    |   no    |   no
// sqs      |    no     |   no    |   no    |   no
// stomp    |    no     |   no    |   no    |   no
// redis    |    no     |   no    |   no    |   no
// mercure  |    no     |   no    |   no    |   no
// ibmmq    |    no     |   yes   |   yes   |   yes

type OperationBinding interface {
	SetOperationBindingValues(interface{}) OperationBinding
}

type OperationBindings map[string]OperationBinding

func NewOperationBindings() OperationBindings {
	return make(OperationBindings)
}

func (b OperationBindings) SetValues(v interface{}) OperationBindings {
	switch mapOperationBindings := v.(type) {
	case map[interface{}]interface{}:
		for key, val := range mapOperationBindings {
			keyStr := fmt.Sprintf("%v", key)

			if "http" == keyStr {
				_, Ok := b[keyStr]

				if !Ok {
					newHttpBinding := NewHttpOperationBinding()
					b[keyStr] = newHttpBinding.SetOperationBindingValues(val)
				}
			}

			if "kafka" == keyStr {
				_, Ok := b[keyStr]

				if !Ok {
					newKafkaBinding := NewKafkaOperationBinding()
					b[keyStr] = newKafkaBinding.SetOperationBindingValues(val)
				}
			}

			if "amqp" == keyStr {
				_, Ok := b[keyStr]

				if !Ok {
					newAmqpBinding := NewAmqpOperationBinding()
					b[keyStr] = newAmqpBinding.SetOperationBindingValues(val)
				}
			}

			if "mqtt" == keyStr {
				_, Ok := b[keyStr]

				if !Ok {
					newMqttBinding := NewMqttOperationBinding()
					b[keyStr] = newMqttBinding.SetOperationBindingValues(val)
				}
			}

		}
	default:
	}
	return b
}

type HttpOperationBinding struct {
	Type           string `validate:"required"`
	Method         string `validate:"oneof=GET POST PUT PATCH DELETE HEAD OPTIONS CONNECT TRACE"`
	Query          *Schema
	BindingVersion string
}

func NewHttpOperationBinding() *HttpOperationBinding {
	return &HttpOperationBinding{}
}

func (h *HttpOperationBinding) SetOperationBindingValues(v interface{}) OperationBinding {

	switch mapOpBind := v.(type) {
	case map[interface{}]interface{}:
		for key, val := range mapOpBind {
			keyStr := fmt.Sprintf("%v", key)
			if keyStr == "type" {
				h.Type = fmt.Sprintf("%v", val)
			}
			if keyStr == "method" {
				h.Method = fmt.Sprintf("%v", val)
			}
			if keyStr == "bindingVersion" {
				h.BindingVersion = fmt.Sprintf("%v", val)
			}
			if keyStr == "query" {
				newSchema := NewSchema()
				h.Query = newSchema.SetValues(val)
			}
		}

	default:
	}

	return h
}

type KafkaOperationBinding struct {
	GroupId        *Schema
	ClientId       *Schema
	BindingVersion string
}

func NewKafkaOperationBinding() *KafkaOperationBinding {
	return &KafkaOperationBinding{}
}

func (h *KafkaOperationBinding) SetOperationBindingValues(v interface{}) OperationBinding {

	switch mapOpBind := v.(type) {
	case map[interface{}]interface{}:
		for key, val := range mapOpBind {
			keyStr := fmt.Sprintf("%v", key)
			if keyStr == "groupId" {
				newSchema := NewSchema()
				h.GroupId = newSchema.SetValues(val)
			}
			if keyStr == "clientId" {
				newSchema := NewSchema()
				h.ClientId = newSchema.SetValues(val)
			}
			if keyStr == "bindingVersion" {
				h.BindingVersion = fmt.Sprintf("%v", val)
			}
		}

	default:
	}

	return h
}

type AmqpOperationBinding struct {
	Expiration     int `validate:"gte=0"`
	UserId         string
	Cc             []string
	Priority       int
	DeliveryMode   int `validate:"oneof=1 2"`
	Mandatory      bool
	Bcc            []string
	ReplyTo        string
	Timestamp      bool
	Ack            bool
	BindingVersion string
}

func NewAmqpOperationBinding() *AmqpOperationBinding {
	return &AmqpOperationBinding{}
}

func (a *AmqpOperationBinding) SetOperationBindingValues(v interface{}) OperationBinding {

	switch mapOpBind := v.(type) {
	case map[interface{}]interface{}:
		for key, val := range mapOpBind {
			keyStr := fmt.Sprintf("%v", key)
			if keyStr == "expiration" {
				intVal, err := strconv.Atoi(fmt.Sprintf("%v", val))
				if err == nil {
					a.Expiration = intVal
				}
			}
			if keyStr == "userId" {
				a.UserId = fmt.Sprintf("%v", val)
			}
			if keyStr == "cc" {
				switch sliceCC := val.(type) {
				case []interface{}:
					c := make([]string, 0)
					for _, cc := range sliceCC {
						c = append(c, fmt.Sprintf("%v", cc))
					}
					a.Cc = c
				}
			}
			if keyStr == "priority" {
				intVal, err := strconv.Atoi(fmt.Sprintf("%v", val))
				if err == nil {
					a.Priority = intVal
				}
			}
			if keyStr == "deliveryMode" {
				intVal, err := strconv.Atoi(fmt.Sprintf("%v", val))
				if err == nil {
					a.DeliveryMode = intVal
				}
			}
			if keyStr == "mandatory" {
				boolVal, err := strconv.ParseBool(fmt.Sprintf("%v", val))
				if err == nil {
					a.Mandatory = boolVal
				}
			}
			if keyStr == "bcc" {
				switch sliceBCC := val.(type) {
				case []interface{}:
					bc := make([]string, 0)
					for _, bcc := range sliceBCC {
						bc = append(bc, fmt.Sprintf("%v", bcc))
					}
					a.Bcc = bc
				}
			}
			if keyStr == "replyTo" {
				a.ReplyTo = fmt.Sprintf("%v", val)
			}
			if keyStr == "timestamp" {
				boolVal, err := strconv.ParseBool(fmt.Sprintf("%v", val))
				if err == nil {
					a.Timestamp = boolVal
				}
			}
			if keyStr == "ack" {
				boolVal, err := strconv.ParseBool(fmt.Sprintf("%v", val))
				if err == nil {
					a.Ack = boolVal
				}
			}
			if keyStr == "bindingVersion" {
				a.BindingVersion = fmt.Sprintf("%v", val)
			}
		}

	default:
	}

	return a
}

type MqttOperationBinding struct {
	Qos            int `validate:"oneof=0 1 2"`
	Retain         bool
	BindingVersion string
}

func NewMqttOperationBinding() *MqttOperationBinding {
	return &MqttOperationBinding{}
}

func (m *MqttOperationBinding) SetOperationBindingValues(v interface{}) OperationBinding {

	switch mapOpBind := v.(type) {
	case map[interface{}]interface{}:
		for key, val := range mapOpBind {
			keyStr := fmt.Sprintf("%v", key)
			if keyStr == "qos" {
				intVal, err := strconv.Atoi(fmt.Sprintf("%v", val))
				if err == nil {
					m.Qos = intVal
				}
			}
			if keyStr == "retain" {
				boolVal, err := strconv.ParseBool(fmt.Sprintf("%v", val))
				if err == nil {
					m.Retain = boolVal
				}
			}
			if keyStr == "bindingVersion" {
				m.BindingVersion = fmt.Sprintf("%v", val)
			}
		}

	default:
	}

	return m
}

type MessageBinding interface {
	SetMessageBindingValues(interface{}) MessageBinding
}

type MessageBindings map[string]MessageBinding

func NewMessageBindings() MessageBindings {
	return make(MessageBindings)
}

func (m MessageBindings) SetValues(v interface{}) MessageBindings {
	switch mapMessageBindings := v.(type) {
	case map[interface{}]interface{}:
		for key, val := range mapMessageBindings {
			keyStr := fmt.Sprintf("%v", key)

			if "http" == keyStr {
				_, Ok := m[keyStr]

				if !Ok {
					newHttpMessageBinding := NewHttpMessageBinding()
					m[keyStr] = newHttpMessageBinding.SetMessageBindingValues(val)
				}
			}

			if "kafka" == keyStr {
				_, Ok := m[keyStr]

				if !Ok {
					newKafkaBinding := NewKafkaMessageBinding()
					m[keyStr] = newKafkaBinding.SetMessageBindingValues(val)
				}
			}

			if "amqp" == keyStr {
				_, Ok := m[keyStr]

				if !Ok {
					newAmqpBinding := NewAmqpMessageBinding()
					m[keyStr] = newAmqpBinding.SetMessageBindingValues(val)
				}
			}

			if "mqtt" == keyStr {
				_, Ok := m[keyStr]

				if !Ok {
					newMqttBinding := NewMqttMessageBinding()
					m[keyStr] = newMqttBinding.SetMessageBindingValues(val)
				}
			}

			if "ibmmq" == keyStr {
				_, Ok := m[keyStr]

				if !Ok {
					newIbmmqBinding := NewIbmmqMessageBinding()
					m[keyStr] = newIbmmqBinding.SetMessageBindingValues(val)
				}
			}
		}
	default:
	}
	return m
}

type HttpMessageBinding struct {
	Headers        *Schema
	BindingVersion string
}

func NewHttpMessageBinding() *HttpMessageBinding {
	return &HttpMessageBinding{}
}

func (h *HttpMessageBinding) SetMessageBindingValues(v interface{}) MessageBinding {
	switch mapMesBind := v.(type) {
	case map[interface{}]interface{}:
		for key, val := range mapMesBind {
			keyStr := fmt.Sprintf("%v", key)
			if keyStr == "headers" {
				newSchema := NewSchema()
				h.Headers = newSchema.SetValues(val)
			}
			if keyStr == "bindingVersion" {
				h.BindingVersion = fmt.Sprintf("%v", val)
			}
		}
	default:
	}
	return h
}

type KafkaMessageBinding struct {
	Key            *Schema
	BindingVersion string
}

func NewKafkaMessageBinding() *KafkaMessageBinding {
	return &KafkaMessageBinding{}
}

func (h *KafkaMessageBinding) SetMessageBindingValues(v interface{}) MessageBinding {
	switch mapMesBind := v.(type) {
	case map[interface{}]interface{}:
		for key, val := range mapMesBind {
			keyStr := fmt.Sprintf("%v", key)
			if keyStr == "key" {
				newSchema := NewSchema()
				h.Key = newSchema.SetValues(val)
			}
			if keyStr == "bindingVersion" {
				h.BindingVersion = fmt.Sprintf("%v", val)
			}
		}
	default:
	}
	return h
}

type AmqpMessageBinding struct {
	ContentEncoding string
	MessageType     string
	BindingVersion  string
}

func NewAmqpMessageBinding() *AmqpMessageBinding {
	return &AmqpMessageBinding{}
}

func (h *AmqpMessageBinding) SetMessageBindingValues(v interface{}) MessageBinding {
	switch mapMesBind := v.(type) {
	case map[interface{}]interface{}:
		for key, val := range mapMesBind {
			keyStr := fmt.Sprintf("%v", key)
			if keyStr == "contentEncoding" {
				h.ContentEncoding = fmt.Sprintf("%v", val)
			}
			if keyStr == "messageType" {
				h.MessageType = fmt.Sprintf("%v", val)
			}
			if keyStr == "bindingVersion" {
				h.BindingVersion = fmt.Sprintf("%v", val)
			}
		}
	default:
	}
	return h
}

type MqttMessageBinding struct {
	BindingVersion string
}

func NewMqttMessageBinding() *MqttMessageBinding {
	return &MqttMessageBinding{}
}

func (h *MqttMessageBinding) SetMessageBindingValues(v interface{}) MessageBinding {
	switch mapMesBind := v.(type) {
	case map[interface{}]interface{}:
		for key, val := range mapMesBind {
			keyStr := fmt.Sprintf("%v", key)
			if keyStr == "bindingVersion" {
				h.BindingVersion = fmt.Sprintf("%v", val)
			}
		}
	default:
	}
	return h
}

type IbmmqMessageBinding struct {
	Type           string
	Headers        string
	Description    string
	Expiry         int
	BindingVersion string
}

func NewIbmmqMessageBinding() *IbmmqMessageBinding {
	return &IbmmqMessageBinding{}
}

func (h *IbmmqMessageBinding) SetMessageBindingValues(v interface{}) MessageBinding {
	switch mapMesBind := v.(type) {
	case map[interface{}]interface{}:
		for key, val := range mapMesBind {
			keyStr := fmt.Sprintf("%v", key)
			if keyStr == "type" {
				h.Type = fmt.Sprintf("%v", val)
			}
			if keyStr == "headers" {
				h.Headers = fmt.Sprintf("%v", val)
			}
			if keyStr == "description" {
				h.Description = fmt.Sprintf("%v", val)
			}
			if keyStr == "expiry" {
				intVal, err := strconv.Atoi(fmt.Sprintf("%v", val))
				if err == nil {
					h.Expiry = intVal
				}
			}
			if keyStr == "bindingVersion" {
				h.BindingVersion = fmt.Sprintf("%v", val)
			}
		}
	default:
	}
	return h
}

type ChannelBinding interface {
	SetChannelBindingValues(interface{}) ChannelBinding
}

type ChannelBindings map[string]ChannelBinding

func NewChannelBindings() ChannelBindings {
	return make(ChannelBindings)
}

func (b ChannelBindings) SetValues(v interface{}) ChannelBindings {
	switch mapChannelBindings := v.(type) {
	case map[interface{}]interface{}:
		for key, val := range mapChannelBindings {
			keyStr := fmt.Sprintf("%v", key)

			if "ws" == keyStr {
				_, Ok := b[keyStr]

				if !Ok {
					newWsBinding := NewWsChannelBinding()
					b[keyStr] = newWsBinding.SetChannelBindingValues(val)
				}
			}

			if "amqp" == keyStr {
				_, Ok := b[keyStr]

				if !Ok {
					newAmqpBinding := NewAmqpChannelBinding()
					b[keyStr] = newAmqpBinding.SetChannelBindingValues(val)
				}
			}

			if "ibmmq" == keyStr {
				_, Ok := b[keyStr]

				if !Ok {
					newIbmmqBinding := NewIbmmqChannelBinding()
					b[keyStr] = newIbmmqBinding.SetChannelBindingValues(val)
				}
			}

		}
	default:
	}
	return b
}

type WsChannelBinding struct {
	Method         string `validate:"oneof=GET POST"`
	Query          *Schema
	Headers        *Schema
	BindingVersion string
}

func NewWsChannelBinding() *WsChannelBinding {
	return &WsChannelBinding{}
}

func (cb *WsChannelBinding) SetChannelBindingValues(v interface{}) ChannelBinding {

	switch mapChannelBind := v.(type) {
	case map[interface{}]interface{}:
		for key, val := range mapChannelBind {
			keyStr := fmt.Sprintf("%v", key)
			if keyStr == "method" {
				cb.Method = fmt.Sprintf("%v", val)
			}
			if keyStr == "query" {
				newQuerySchema := NewSchema()

				cb.Query = newQuerySchema.SetValues(val)
			}
			if keyStr == "headers" {
				newQuerySchema := NewSchema()

				cb.Headers = newQuerySchema.SetValues(val)
			}
			if keyStr == "bindingVersion" {
				cb.BindingVersion = fmt.Sprintf("%v", val)
			}
		}

	default:
	}
	return cb
}

type AmqpChannelBinding struct {
	Is             string `validate:"oneof=queue routingKey"`
	ChannelSpec    map[string]AmqpChannelSpecification
	BindingVersion string
}

func NewAmqpChannelBinding() *AmqpChannelBinding {
	return &AmqpChannelBinding{}
}

func (cb *AmqpChannelBinding) SetChannelBindingValues(v interface{}) ChannelBinding {

	switch mapChannelBind := v.(type) {
	case map[interface{}]interface{}:
		for key, val := range mapChannelBind {
			keyStr := fmt.Sprintf("%v", key)
			if keyStr == "is" {
				cb.Is = fmt.Sprintf("%v", val)
			}
			if keyStr == "queue" {
				chanSpec := make(map[string]AmqpChannelSpecification)
				newAmqpChannelQueue := NewAmqpChannelQueue()
				chanSpec[keyStr] = newAmqpChannelQueue.SetAmqpChannelSpecValues(val)
				cb.ChannelSpec = chanSpec

			}
			if keyStr == "exchange" {
				chanSpec := make(map[string]AmqpChannelSpecification)
				newAmqpChannelExchange := NewAmqpChannelExchange()
				chanSpec[keyStr] = newAmqpChannelExchange.SetAmqpChannelSpecValues(val)
				cb.ChannelSpec = chanSpec
			}
			if keyStr == "bindingVersion" {
				cb.BindingVersion = fmt.Sprintf("%v", val)
			}
		}

	default:
	}
	return cb
}

type AmqpChannelSpecification interface {
	SetAmqpChannelSpecValues(v interface{}) AmqpChannelSpecification
}

type AmqpChannelExchange struct {
	Name       string `validate:"max=255"`
	Type       string `validate:"oneof=topic direct fanout default headers"`
	Durable    bool
	AutoDelete bool
	VHost      string
}

func NewAmqpChannelExchange() *AmqpChannelExchange {
	return &AmqpChannelExchange{}
}

func (q *AmqpChannelExchange) SetAmqpChannelSpecValues(v interface{}) AmqpChannelSpecification {

	switch mapChannelBind := v.(type) {
	case map[interface{}]interface{}:
		for key, val := range mapChannelBind {
			keyStr := fmt.Sprintf("%v", key)
			if keyStr == "name" {
				q.Name = fmt.Sprintf("%v", val)

			}
			if keyStr == "type" {
				q.Type = fmt.Sprintf("%v", val)

			}
			if keyStr == "durable" {
				boolVal, err := strconv.ParseBool(fmt.Sprintf("%v", val))
				if err == nil {
					q.Durable = boolVal
				}
			}
			if keyStr == "autoDelete" {
				boolVal, err := strconv.ParseBool(fmt.Sprintf("%v", val))
				if err == nil {
					q.AutoDelete = boolVal
				}
			}
			if keyStr == "vhost" {
				q.VHost = fmt.Sprintf("%v", val)
			}
		}

	default:
	}
	return q
}

type AmqpChannelQueue struct {
	Name       string `validate:"max=255"`
	Durable    bool
	Exclusive  bool
	AutoDelete bool
	VHost      string
}

func NewAmqpChannelQueue() *AmqpChannelQueue {
	return &AmqpChannelQueue{}
}

func (q *AmqpChannelQueue) SetAmqpChannelSpecValues(v interface{}) AmqpChannelSpecification {
	switch mapChannelBind := v.(type) {
	case map[interface{}]interface{}:
		for key, val := range mapChannelBind {
			keyStr := fmt.Sprintf("%v", key)
			if keyStr == "name" {
				q.Name = fmt.Sprintf("%v", val)
			}
			if keyStr == "durable" {
				boolVal, err := strconv.ParseBool(fmt.Sprintf("%v", val))
				if err == nil {
					q.Durable = boolVal
				}
			}
			if keyStr == "exclusive" {
				boolVal, err := strconv.ParseBool(fmt.Sprintf("%v", val))
				if err == nil {
					q.Exclusive = boolVal
				}
			}
			if keyStr == "autoDelete" {
				boolVal, err := strconv.ParseBool(fmt.Sprintf("%v", val))
				if err == nil {
					q.AutoDelete = boolVal
				}
			}
			if keyStr == "vhost" {
				q.VHost = fmt.Sprintf("%v", val)
			}
		}

	default:
	}
	return q
}

type IbmmqChannelSpecification interface {
	SetIbmmqChannelSpecValues(v interface{}) IbmmqChannelSpecification
}

type IbmmqChannelQueue struct {
	ObjectName    string `validate:"required"`
	IsPartitioned bool
	Exclusive     bool
}

func NewIbmmqChannelQueue() *IbmmqChannelQueue {
	return &IbmmqChannelQueue{}
}

func (q *IbmmqChannelQueue) SetIbmmqChannelSpecValues(v interface{}) IbmmqChannelSpecification {

	switch mapChannelBind := v.(type) {
	case map[interface{}]interface{}:
		for key, val := range mapChannelBind {
			keyStr := fmt.Sprintf("%v", key)
			if keyStr == "objectName" {
				q.ObjectName = fmt.Sprintf("%v", val)

			}
			if keyStr == "durablePermitted" {
				boolVal, err := strconv.ParseBool(fmt.Sprintf("%v", val))
				if err == nil {
					q.IsPartitioned = boolVal
				}
			}
			if keyStr == "lastMsgRetained" {
				boolVal, err := strconv.ParseBool(fmt.Sprintf("%v", val))
				if err == nil {
					q.Exclusive = boolVal
				}
			}
		}

	default:
	}
	return q
}

type IbmmqChannelTopic struct {
	TopicString      string
	ObjectName       string
	DurablePermitted bool
	LastMsgRetained  bool
}

func NewIbmmqChannelTopic() *IbmmqChannelTopic {
	return &IbmmqChannelTopic{}
}

func (t *IbmmqChannelTopic) SetIbmmqChannelSpecValues(v interface{}) IbmmqChannelSpecification {

	switch mapChannelBind := v.(type) {
	case map[interface{}]interface{}:
		for key, val := range mapChannelBind {
			keyStr := fmt.Sprintf("%v", key)
			if keyStr == "string" {
				t.TopicString = fmt.Sprintf("%v", val)
			}
			if keyStr == "objectName" {
				t.ObjectName = fmt.Sprintf("%v", val)

			}
			if keyStr == "durablePermitted" {
				boolVal, err := strconv.ParseBool(fmt.Sprintf("%v", val))
				if err == nil {
					t.DurablePermitted = boolVal
				}
			}
			if keyStr == "lastMsgRetained" {
				boolVal, err := strconv.ParseBool(fmt.Sprintf("%v", val))
				if err == nil {
					t.LastMsgRetained = boolVal
				}
			}
		}

	default:
	}
	return t
}

type IbmmqChannelBinding struct {
	DestinationType string
	ChannelSpec     map[string]IbmmqChannelSpecification
	MaxMsgLength    int
	BindingVersion  string
}

func NewIbmmqChannelBinding() *IbmmqChannelBinding {
	return &IbmmqChannelBinding{}
}

func (cb *IbmmqChannelBinding) SetChannelBindingValues(v interface{}) ChannelBinding {

	switch mapChannelBind := v.(type) {
	case map[interface{}]interface{}:
		for key, val := range mapChannelBind {
			keyStr := fmt.Sprintf("%v", key)
			if keyStr == "destinationType" {
				cb.DestinationType = fmt.Sprintf("%v", val)
			}
			if keyStr == "queue" {
				channelSpec := make(map[string]IbmmqChannelSpecification)
				newIbmmqChannelQueue := NewIbmmqChannelQueue()

				channelSpec[keyStr] = newIbmmqChannelQueue.SetIbmmqChannelSpecValues(val)
				cb.ChannelSpec = channelSpec

			}
			if keyStr == "topic" {
				channelSpec := make(map[string]IbmmqChannelSpecification)
				newIbmmqChannelTopic := NewIbmmqChannelTopic()

				channelSpec[keyStr] = newIbmmqChannelTopic.SetIbmmqChannelSpecValues(val)
				cb.ChannelSpec = channelSpec
			}
			if keyStr == "maxMsgLength" {
				intVal, err := strconv.Atoi(fmt.Sprintf("%v", val))
				if err == nil {
					cb.MaxMsgLength = intVal
				}
			}
			if keyStr == "bindingVersion" {
				cb.BindingVersion = fmt.Sprintf("%v", val)
			}
		}

	default:
	}

	return cb
}

type ServerBinding interface {
	SetServerBindingValues(interface{}) ServerBinding
}

type ServerBindings map[string]ServerBinding

func NewServerBindings() ServerBindings {
	return make(ServerBindings)
}

func (b ServerBindings) SetValues(v interface{}) ServerBindings {
	switch mapServerBindings := v.(type) {
	case map[interface{}]interface{}:
		for key, val := range mapServerBindings {
			keyStr := fmt.Sprintf("%v", key)

			if "mqtt" == keyStr {
				_, Ok := b[keyStr]

				if !Ok {
					newMqttBinding := NewMqttServerBinding()
					b[keyStr] = newMqttBinding.SetServerBindingValues(val)
				}
			}

			if "ibmmq" == keyStr {
				_, Ok := b[keyStr]

				if !Ok {
					newIbmmqBinding := NewIbmmqServerBinding()
					b[keyStr] = newIbmmqBinding.SetServerBindingValues(val)
				}
			}

		}
	default:
	}
	return b
}

type MqttLastWill struct {
	Topic   string
	Qos     int `validate:"oneof=0 1 2"`
	Message string
	Retain  bool
}

func NewMqttLastWill() *MqttLastWill {
	return &MqttLastWill{}
}

func (l *MqttLastWill) SetValues(v interface{}) *MqttLastWill {
	switch mapLastWill := v.(type) {
	case map[interface{}]interface{}:
		for key, val := range mapLastWill {
			keyStr := fmt.Sprintf("%v", key)
			if keyStr == "topic" {
				l.Topic = fmt.Sprintf("%v", val)
			}
			if keyStr == "qos" {
				intVal, err := strconv.Atoi(fmt.Sprintf("%v", val))
				if err == nil {
					l.Qos = intVal
				}
			}
			if keyStr == "message" {
				l.Message = fmt.Sprintf("%v", val)
			}
			if keyStr == "retain" {
				boolVal, err := strconv.ParseBool(fmt.Sprintf("%v", val))
				if err == nil {
					l.Retain = boolVal
				} else {
					l.Retain = false
				}
			}
		}

	default:
	}

	return l
}

type MqttServerBinding struct {
	ClientId       string
	CleanSession   bool
	LastWill       *MqttLastWill
	KeepAlive      int
	BindingVersion string
}

func NewMqttServerBinding() *MqttServerBinding {
	return &MqttServerBinding{}
}

func (sb *MqttServerBinding) SetServerBindingValues(v interface{}) ServerBinding {

	switch mapServerBind := v.(type) {
	case map[interface{}]interface{}:
		for key, val := range mapServerBind {
			keyStr := fmt.Sprintf("%v", key)
			if keyStr == "clientId" {
				sb.ClientId = fmt.Sprintf("%v", val)
			}
			if keyStr == "cleanSession" {
				boolVal, err := strconv.ParseBool(fmt.Sprintf("%v", val))
				if err == nil {
					sb.CleanSession = boolVal
				}
			}
			if keyStr == "lastWill" {
				newLastWill := NewMqttLastWill()
				sb.LastWill = newLastWill.SetValues(val)
			}
			if keyStr == "keepAlive" {
				intVal, err := strconv.Atoi(fmt.Sprintf("%v", val))
				if err == nil {
					sb.KeepAlive = intVal
				}
			}
			if keyStr == "bindingVersion" {
				sb.BindingVersion = fmt.Sprintf("%v", val)
			}
		}

	default:
	}

	return sb
}

type IbmmqServerBinding struct {
	GroupId              string
	CcdtQueueManagerName string
	CipherSpec           string
	MultiEndpointServer  bool
	HeartBeatInterval    int
	BindingVersion       string
}

func NewIbmmqServerBinding() *IbmmqServerBinding {
	return &IbmmqServerBinding{}
}

func (sb *IbmmqServerBinding) SetServerBindingValues(v interface{}) ServerBinding {

	switch mapServerBind := v.(type) {
	case map[interface{}]interface{}:
		for key, val := range mapServerBind {
			keyStr := fmt.Sprintf("%v", key)
			if keyStr == "groupId" {
				sb.GroupId = fmt.Sprintf("%v", val)
			}
			if keyStr == "ccdtQueueManagerName" {
				sb.CcdtQueueManagerName = fmt.Sprintf("%v", val)
			}
			if keyStr == "cipherSpec" {
				sb.CipherSpec = fmt.Sprintf("%v", val)
			}
			if keyStr == "multiEndpointServer" {
				boolVal, err := strconv.ParseBool(fmt.Sprintf("%v", val))
				if err == nil {
					sb.MultiEndpointServer = boolVal
				}
			}
			if keyStr == "heartBeatInterval" {
				intVal, err := strconv.Atoi(fmt.Sprintf("%v", val))
				if err == nil {
					sb.HeartBeatInterval = intVal
				}
			}
			if keyStr == "bindingVersion" {
				sb.BindingVersion = fmt.Sprintf("%v", val)
			}
		}

	default:
	}

	return sb
}
