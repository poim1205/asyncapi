package asyncapi2

import "gopkg.in/yaml.v2"

type Parameters map[string]Parameter

type Parameter struct {
	Description string
	Schema      Schema
	Location    string
}

func (value *Parameter) MarshalYAML() ([]byte, error) {
	return yaml.Marshal(value)
}

func (value *Parameter) UnmarshalYAML(data []byte) error {
	return yaml.Unmarshal(data, value)
}
