package asyncapi2

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

type Operation struct {
	Ref          string
	OperationId  string
	Summary      string
	Description  string
	Tags         []*Tag
	ExternalDocs *ExternalDocs
	Bindings     map[string]OperationBindings
	Traits       []OperationTrait
	Message      []Message
}

func NewOperation(v interface{}) *Operation {
	o := Operation{}

	switch mapOperation := v.(type) {
	case map[interface{}]interface{}:
		for key, val := range mapOperation {
			keyStr := fmt.Sprintf("%v", key)
			if keyStr == "$ref" {
				o.Ref = fmt.Sprintf("%v", val)
			}
			if keyStr == "operationId" {
				o.OperationId = fmt.Sprintf("%v", val)
			}
			if keyStr == "summary" {
				o.Summary = fmt.Sprintf("%v", val)
			}
			if keyStr == "description" {
				o.Description = fmt.Sprintf("%v", val)
			}
			if keyStr == "tags" {
				o.Tags = NewTags(val)
			}
			if keyStr == "externalDocs" {
				o.ExternalDocs = NewExternalDocs(val)
			}
			if keyStr == "bindings" {
				o.ExternalDocs = NewExternalDocs(val)
			}
		}

	default:
	}
	return &o
}

func (value *Operation) MarshalYAML() ([]byte, error) {
	return yaml.Marshal(value)
}

func (value *Operation) UnmarshalYAML(data []byte) error {
	return yaml.Unmarshal(data, value)
}

type OperationBindings map[string]OperationBinding

type OperationBinding struct {
	_type          string
	method         string
	query          Schema
	bindingVersion string
}

func (value *OperationBinding) MarshalYAML() ([]byte, error) {
	return yaml.Marshal(value)
}

func (value *OperationBinding) UnmarshalYAML(data []byte) error {
	return yaml.Unmarshal(data, value)
}

type OperationTrait struct {
	operationId  string
	summary      string
	description  string
	tags         Tags
	externalDocs ExternalDocs
	bindings     OperationBindings
}

func NewOperationTrait() *OperationTrait {
	return &OperationTrait{}
}

func (value *OperationTrait) MarshalYAML() ([]byte, error) {
	return yaml.Marshal(value)
}

func (value *OperationTrait) UnmarshalYAML(data []byte) error {
	return yaml.Unmarshal(data, value)
}
