package asyncapi2

import (
	"fmt"
)

type Channels map[string]*Channel

func NewChannels() Channels {
	return make(Channels)
}

func (c Channels) SetValues(v interface{}) Channels {
	switch arrayVal := v.(type) {
	case map[interface{}]interface{}:
		for key, val := range arrayVal {
			keyString := fmt.Sprintf("%v", key)
			_, Ok := c[keyString]
			if !Ok {
				nc := NewChannel()

				c[keyString] = nc.SetValues(val)
			}
		}
	default:
	}
	return c
}

type Channel struct {
	Ref         string
	Description string
	Subscribe   *Operation
	Publish     *Operation
	Parameters  Parameters
	Bindings    ChannelBindings
}

func NewChannel() *Channel {
	return &Channel{}
}

func (value *Channel) SetValues(v interface{}) *Channel {
	switch mapVal := v.(type) {
	case map[interface{}]interface{}:
		for key, val := range mapVal {
			keyString := fmt.Sprintf("%v", key)
			if keyString == "$ref" {
				value.Ref = fmt.Sprintf("%v", val)
			}
			if keyString == "description" {
				value.Description = fmt.Sprintf("%v", val)
			}
			if keyString == "subscribe" {
				newOperation := NewOperation()
				value.Subscribe = newOperation.SetValues(val)
			}
			if keyString == "publish" {

				newOperation := NewOperation()
				value.Publish = newOperation.SetValues(val)
			}
			if keyString == "parameters" {

				p := NewParameters()
				value.Parameters = p.SetValues(val)
			}
		}
	default:
	}

	return value
}
