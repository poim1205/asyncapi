package asyncapi2

import (
	"fmt"
)

type Tags []*Tag

func NewTags() Tags {
	t := make(Tags, 0)
	return t
}

func (t Tags) SetValues(v interface{}) Tags {
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
	Name         string `validate:"required"`
	Description  string
	ExternalDocs *ExternalDocs
}

func NewTag() *Tag {
	return &Tag{}
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
				newExternalDocs := NewExternalDocs()
				value.ExternalDocs = newExternalDocs.SetValues(val)
			}
		}
	default:
	}
	return value
}
