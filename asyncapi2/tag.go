package asyncapi2

import "gopkg.in/yaml.v2"

type Tags []Tag

type Tag struct {
	Name         string
	Description  string
	ExternalDocs *ExternalDocs
}

func (value *Tag) MarshalYAML() ([]byte, error) {
	return yaml.Marshal(value)
}

func (value *Tag) UnmarshalYAML(data []byte) error {
	return yaml.Unmarshal(data, value)
}
