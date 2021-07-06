package asyncapi2

import "gopkg.in/yaml.v2"

type Operation struct {
	OperationId  string
	Summary      string
	Description  string
	Tags         []Tag
	ExternalDocs ExternalDocs
	Bindings     map[string]OperationBindings
	Traits       []OperationTrait
	Message      []Message
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

func (value *OperationTrait) MarshalYAML() ([]byte, error) {
	return yaml.Marshal(value)
}

func (value *OperationTrait) UnmarshalYAML(data []byte) error {
	return yaml.Unmarshal(data, value)
}
