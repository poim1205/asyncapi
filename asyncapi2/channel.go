package asyncapi2

import "gopkg.in/yaml.v2"

type Channels map[string]*Channel

type Channel struct {
	ref         string
	description string
	subscribe   Operation
	publish     Operation
	parameters  Parameters
	bindings    ChannelBindings
}

func (value *Channel) MarshalYAML() ([]byte, error) {
	return yaml.Marshal(value)
}

func (value *Channel) UnmarshalYAML(data []byte) error {
	return yaml.Unmarshal(data, value)
}

type ChannelBindings map[string]ChannelBinding

type ChannelBinding struct {
}
