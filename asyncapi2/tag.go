package asyncapi2

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

type Tags []*Tag

func NewTags(v interface{}) Tags {
	t := make(Tags, 1)
	switch arrayVal := v.(type) {
	case []interface{}:
		for _, val := range arrayVal {
			nt := NewTag()
			t = append(t, nt.SetValues(val))
		}
	default:
	}
	return t
}

type Tag struct {
	Name         string
	Description  string
	ExternalDocs *ExternalDocs
}

func NewTag() *Tag {
	return &Tag{}
}

func (value *Tag) MarshalYAML() ([]byte, error) {
	return yaml.Marshal(value)
}

func (value *Tag) UnmarshalYAML(data []byte) error {
	return yaml.Unmarshal(data, value)
}

func (value *Tag) SetValues(v interface{}) *Tag {
	switch mapVal := v.(type) {
	case map[interface{}]interface{}:
		for key, val := range mapVal {
			keyString := fmt.Sprintf("%v", key)
			if keyString == "name" {
				value.Name = fmt.Sprintf("%v", val)
			}
			if keyString == "description" {
				value.Description = fmt.Sprintf("%v", val)
			}
			if keyString == "externalDocs" {
				value.ExternalDocs = NewExternalDocs(val)
			}
		}
	default:
	}
	return value
}
