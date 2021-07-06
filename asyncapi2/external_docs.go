package asyncapi2

import "gopkg.in/yaml.v2"

type ExternalDocs struct {
	Description string
	URL         string
}

func (value *ExternalDocs) MarshalYAML() ([]byte, error) {
	return yaml.Marshal(value)
}

func (value *ExternalDocs) UnmarshalYAML(data []byte) error {
	return yaml.Unmarshal(data, value)
}
