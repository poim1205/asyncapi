package asyncapi2

import (
	"fmt"
)

type Message struct {
	Ref           string
	Headers       *Schema
	Payload       interface{}
	CorrelationId *CorrelationID
	SchemaFormat  string // strick list oneof
	ContentType   string
	Name          string
	Title         string
	Summary       string
	Description   string
	Tags          Tags
	ExternalDocs  *ExternalDocs
	Bindings      MessageBindings
	Traits        MessageTraits
}

func NewMessage() *Message {
	return &Message{}
}

func (message *Message) SetValues(v interface{}) *Message {
	switch mapOpTrait := v.(type) {
	case map[interface{}]interface{}:
		for key, val := range mapOpTrait {
			keyStr := fmt.Sprintf("%v", key)
			if keyStr == "$ref" {
				message.Ref = fmt.Sprintf("%v", val)
			}
			if keyStr == "headers" {
				sch := NewSchema()
				message.Headers = sch.SetValues(val)
			}
			if keyStr == "payload" {
				message.Payload = val
			}
			if keyStr == "correlationId" {
				corID := NewCorrelationId()

				message.CorrelationId = corID.SetValues(val)
			}
			if keyStr == "schemaFormat" {
				message.SchemaFormat = fmt.Sprintf("%v", val)
			}
			if keyStr == "contentType" {
				message.ContentType = fmt.Sprintf("%v", val)
			}
			if keyStr == "name" {
				message.Name = fmt.Sprintf("%v", val)
			}
			if keyStr == "title" {
				message.Title = fmt.Sprintf("%v", val)
			}
			if keyStr == "summary" {
				message.Summary = fmt.Sprintf("%v", val)
			}
			if keyStr == "description" {
				message.Description = fmt.Sprintf("%v", val)
			}
			if keyStr == "tags" {
				newTags := NewTags()
				message.Tags = newTags.SetValues(val)
			}
			if keyStr == "externalDocs" {
				newExternalDocs := NewExternalDocs()
				message.ExternalDocs = newExternalDocs.SetValues(val)
			}
			if keyStr == "bindings" {
				newMesBind := NewMessageBindings()
				message.Bindings = newMesBind
			}
			if keyStr == "traits" {
				newTraits := NewMessageTraits()

				message.Traits = newTraits.SetValues(val)
			}
		}

	default:
	}
	return message
}

type MessageTraits []*MessageTrait

func NewMessageTraits() MessageTraits {
	return make(MessageTraits, 0)
}

func (messTraits MessageTraits) SetValues(v interface{}) MessageTraits {
	switch sliceMessTraits := v.(type) {
	case []interface{}:
		for _, messTraitVal := range sliceMessTraits {

			newMessTrait := NewMessageTrait()
			messTraits = append(messTraits, newMessTrait.SetValues(messTraitVal))
		}
	default:
	}
	return messTraits
}

type MessageTrait struct {
	Ref           string
	Headers       *Schema
	CorrelationId *CorrelationID
	SchemaFormat  string // strick list
	ContentType   string
	Name          string
	Title         string
	Summary       string
	Description   string
	Tags          Tags
	ExternalDocs  *ExternalDocs
	Bindings      MessageBindings
}

func NewMessageTrait() *MessageTrait {
	return &MessageTrait{}
}

func (messTrait *MessageTrait) SetValues(v interface{}) *MessageTrait {
	switch mapMessTrait := v.(type) {
	case map[interface{}]interface{}:
		for key, val := range mapMessTrait {
			keyStr := fmt.Sprintf("%v", key)
			if keyStr == "$ref" {
				messTrait.Ref = fmt.Sprintf("%v", val)
			}
			if keyStr == "headers" {
				sch := NewSchema()
				messTrait.Headers = sch.SetValues(val)
			}
			if keyStr == "correlationId" {
				corID := NewCorrelationId()

				messTrait.CorrelationId = corID.SetValues(val)
			}
			if keyStr == "schemaFormat" {
				messTrait.SchemaFormat = fmt.Sprintf("%v", val)
			}
			if keyStr == "contentType" {
				messTrait.ContentType = fmt.Sprintf("%v", val)
			}
			if keyStr == "name" {
				messTrait.Name = fmt.Sprintf("%v", val)
			}
			if keyStr == "title" {
				messTrait.Title = fmt.Sprintf("%v", val)
			}
			if keyStr == "summary" {
				messTrait.Summary = fmt.Sprintf("%v", val)
			}
			if keyStr == "description" {
				messTrait.Description = fmt.Sprintf("%v", val)
			}
			if keyStr == "tags" {
				newTags := NewTags()
				messTrait.Tags = newTags.SetValues(val)
			}
			if keyStr == "externalDocs" {
				newExternalDocs := NewExternalDocs()
				messTrait.ExternalDocs = newExternalDocs.SetValues(val)
			}
			if keyStr == "bindings" {
				newMesBind := NewMessageBindings()
				messTrait.Bindings = newMesBind.SetValues(val)
			}
		}

	default:
	}
	return messTrait
}

type CorrelationID struct {
	Description string
	Location    string
}

func NewCorrelationId() *CorrelationID {
	return &CorrelationID{}
}

func (corID *CorrelationID) SetValues(v interface{}) *CorrelationID {
	switch mapCorID := v.(type) {
	case map[interface{}]interface{}:
		for key, val := range mapCorID {
			keyStr := fmt.Sprintf("%v", key)
			if keyStr == "description" {
				corID.Description = fmt.Sprintf("%v", val)
			}
			if keyStr == "location" {
				corID.Location = fmt.Sprintf("%v", val)
			}
		}
	default:
	}
	return corID
}
