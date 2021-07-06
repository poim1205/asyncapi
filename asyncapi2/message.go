package asyncapi2

import "gopkg.in/yaml.v2"

type Message struct {
	header        Schema
	payload       interface{}
	correlationId CorrelationID
	schemaFormat  string // strick list oneof
	contentType   string
	name          string
	title         string
	summary       string
	description   string
	tags          Tags
	externalDocs  ExternalDocs
	bindings      MessageBindings
	examples      []map[string]interface{}
	traits        []MessageTrait
}

func (value *Message) MarshalYAML() ([]byte, error) {
	return yaml.Marshal(value)
}

func (value *Message) UnmarshalYAML(data []byte) error {
	return yaml.Unmarshal(data, value)
}

type MessageTrait struct {
	headers       Schema
	correlationId CorrelationID
	schemaFormat  string // strick list
	contentType   string
	name          string
	title         string
	summary       string
	description   string
	tags          Tags
	externalDocs  ExternalDocs
	bindings      map[string]MessageBindings
	examples      []map[string]interface{}
}

func (value *MessageTrait) MarshalYAML() ([]byte, error) {
	return yaml.Marshal(value)
}

func (value *MessageTrait) UnmarshalYAML(data []byte) error {
	return yaml.Unmarshal(data, value)
}

type MessageBindings map[string]MessageBinding

type MessageBinding struct {
	Header         Schema
	BindingVersion string
}

func (value *MessageBinding) MarshalYAML() ([]byte, error) {
	return yaml.Marshal(value)
}

func (value *MessageBinding) UnmarshalYAML(data []byte) error {
	return yaml.Unmarshal(data, value)
}

type CorrelationID struct {
	Description string
	Location    string
}

func (value *CorrelationID) MarshalYAML() ([]byte, error) {
	return yaml.Marshal(value)
}

func (value *CorrelationID) UnmarshalYAML(data []byte) error {
	return yaml.Unmarshal(data, value)
}
