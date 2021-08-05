package asyncapi2

import (
	"fmt"
)

type ExternalDocs struct {
	Description string
	URL         string `validate:"required"`
}

func NewExternalDocs() *ExternalDocs {
	return &ExternalDocs{}
}

func (e *ExternalDocs) SetValues(v interface{}) *ExternalDocs {
	switch mapVal := v.(type) {
	case map[interface{}]interface{}:
		for key, val := range mapVal {
			keyString := fmt.Sprintf("%v", key)
			if keyString == "url" {
				e.URL = fmt.Sprintf("%v", val)
			}
			if keyString == "description" {
				e.Description = fmt.Sprintf("%v", val)
			}

		}
	default:
	}
	return e
}
