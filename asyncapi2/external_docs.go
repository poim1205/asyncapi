package asyncapi2

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

type ExternalDocs struct {
	Description string
	URL         string
}

func NewExternalDocs(v interface{}) *ExternalDocs {
	value := ExternalDocs{}
	switch mapVal := v.(type) {
	case map[interface{}]interface{}:
		for key, val := range mapVal {
			keyString := fmt.Sprintf("%v", key)
			if keyString == "url" {
				value.URL = fmt.Sprintf("%v", val)
			}
			if keyString == "description" {
				value.Description = fmt.Sprintf("%v", val)
			}

		}
	default:
	}
	return &value
}

func (value *ExternalDocs) MarshalYAML() ([]byte, error) {
	return yaml.Marshal(value)
}

func (value *ExternalDocs) UnmarshalYAML(data []byte) error {
	return yaml.Unmarshal(data, value)
}

func (value *ExternalDocs) PrintExternalDocs() {
	fmt.Printf("ExternalDocs.url: %s\n", value.URL)
	fmt.Printf("ExternalDocs.description: %s\n", value.Description)
}
