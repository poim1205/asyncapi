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

func (ts Tags) PrintTags(indentString string) {
	fmt.Printf("%stags:\n", indentString)
	for _, t := range ts {
		t.PrintTag(indentString)
	}
}

type Tag struct {
	Name         string
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

func (value *Tag) PrintTag(indentString string) {
	fmt.Printf("%s- name: %s", indentString, value.Name)
	fmt.Printf("%s%sdescription: %s", indentString, INDENT, value.Description)
	// if value.ExternalDocs != nil {
	// 	value.ExternalDocs.PrintExternalDocs(fmt.Sprintf("%s%s", indentString, INDENT))
	// }
}
