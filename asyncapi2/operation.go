package asyncapi2

import (
	"fmt"
)

type Operation struct {
	Ref          string // ???
	OperationId  string
	Summary      string
	Description  string
	Tags         Tags
	ExternalDocs *ExternalDocs
	Bindings     OperationBindings
	Traits       OperationTraits
	Message      *Message
}

func NewOperation() *Operation {
	return &Operation{}
}

func (o *Operation) SetValues(v interface{}) *Operation {
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
				newTags := NewTags()

				o.Tags = newTags.SetValues(val)
			}
			if keyStr == "externalDocs" {
				newExternalDocs := NewExternalDocs()
				o.ExternalDocs = newExternalDocs.SetValues(val)
			}
			if keyStr == "bindings" {
				opBindings := NewOperationBindings()
				o.Bindings = opBindings.SetValues(val)
			}
			if keyStr == "traits" {
				traits := NewOperationTraits()
				o.Traits = traits.SetValues(val)
			}
			if keyStr == "message" {
				message := NewMessage()
				o.Message = message.SetValues(val)
			}
		}

	default:
	}
	return o
}

type OperationTraits []*OperationTrait

func NewOperationTraits() OperationTraits {
	return make(OperationTraits, 0)
}

func (traits OperationTraits) SetValues(v interface{}) OperationTraits {
	switch operationTraitsVal := v.(type) {
	case []interface{}:
		for _, operationTraitVal := range operationTraitsVal {

			newTrait := NewOperationTrait()
			traits = append(traits, newTrait.SetValues(operationTraitVal))
		}
	default:
	}

	return traits
}

type OperationTrait struct {
	Ref          string
	OperationId  string
	Summary      string
	Description  string
	Tags         Tags
	ExternalDocs *ExternalDocs
	Bindings     OperationBindings
}

func NewOperationTrait() *OperationTrait {
	return &OperationTrait{}
}

func (ot *OperationTrait) SetValues(v interface{}) *OperationTrait {
	switch mapOpTrait := v.(type) {
	case map[interface{}]interface{}:
		for key, val := range mapOpTrait {
			keyStr := fmt.Sprintf("%v", key)
			if keyStr == "$ref" {
				ot.Ref = fmt.Sprintf("%v", val)
			}
			if keyStr == "operationId" {
				ot.OperationId = fmt.Sprintf("%v", val)
			}
			if keyStr == "summary" {
				ot.Summary = fmt.Sprintf("%v", val)
			}
			if keyStr == "description" {
				ot.Description = fmt.Sprintf("%v", val)
			}
			if keyStr == "tags" {
				newTags := NewTags()
				ot.Tags = newTags.SetValues(val)
			}
			if keyStr == "externalDocs" {
				newExternalDocs := NewExternalDocs()
				ot.ExternalDocs = newExternalDocs.SetValues(val)
			}
			if keyStr == "bindings" {
				opBindings := NewOperationBindings()

				ot.Bindings = opBindings.SetValues(val)
			}
		}

	default:
	}

	return ot
}
