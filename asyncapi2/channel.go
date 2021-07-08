package asyncapi2

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

type Channels map[string]*Channel

func NewChannels(v interface{}) Channels {
	var c Channels
	switch arrayVal := v.(type) {
	case map[interface{}]interface{}:
		mc := make(Channels, 1)
		for key, val := range arrayVal {
			keyString := fmt.Sprintf("%v", key)
			_, Ok := mc[keyString]
			if !Ok {
				nc := NewChannel()

				mc[keyString] = nc.SetValues(val)
			}
		}
		c = mc
	default:
	}
	return c
}

func (c Channels) PrintChannels() {
	fmt.Println("Print Channels")
	for k, v := range c {
		fmt.Printf("Channels.%s.subscribe.ref: %s\n", k, v.Subscribe.Ref)
	}
}

type Channel struct {
	Ref         string
	Description string
	Subscribe   *Operation
	Publish     Operation
	Parameters  Parameters
	Bindings    ChannelBindings
}

func NewChannel() *Channel {
	return &Channel{}
}

func (value *Channel) MarshalYAML() ([]byte, error) {
	return yaml.Marshal(value)
}

func (value *Channel) UnmarshalYAML(data []byte) error {
	return yaml.Unmarshal(data, value)
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
				fmt.Printf("Channel.subscribe.Type: %T\n", val)
				fmt.Printf("Channel.subscribe.Value: %v\n", val)

				o := NewOperation(val)
				value.Subscribe = o

				// switch subscribeVal := val.(type) {
				// case map[interface{}]interface{}:

				// 	for subKey, subVal := range subscribeVal {
				// 		keySubVar := fmt.Sprintf("%v", subKey)
				// 		nope := NewOperation()

				// 	}
				// 	value.Subscribe = sv
				// default:
				// }
			}
		}
	default:
	}

	return value
}

type ChannelBindings map[string]ChannelBinding

type ChannelBinding struct {
}
