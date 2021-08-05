package asyncapi2

import (
	"fmt"
)

type Parameters map[string]*Parameter

func NewParameters() Parameters {
	return make(Parameters)
}

func (p Parameters) SetValues(v interface{}) Parameters {
	switch arrayVal := v.(type) {
	case map[interface{}]interface{}:
		for key, val := range arrayVal {
			keyString := fmt.Sprintf("%v", key)
			_, Ok := p[keyString]
			if !Ok {
				newParam := NewParameter()
				p[keyString] = newParam.SetValues(val)
			}
		}
	default:
	}
	return p
}

type Parameter struct {
	Ref         string
	Description string
	Schema      *Schema
	Location    string
}

func NewParameter() *Parameter {
	return &Parameter{}
}

func (p *Parameter) SetValues(v interface{}) *Parameter {

	switch arrayVal := v.(type) {
	case map[interface{}]interface{}:
		for key, val := range arrayVal {
			keyString := fmt.Sprintf("%v", key)
			if keyString == "description" {
				p.Description = fmt.Sprintf("%v", val)
			}
			if keyString == "location" {
				p.Location = fmt.Sprintf("%v", val)
			}
			if keyString == "schema" {
				sch := NewSchema()
				p.Schema = sch.SetValues(val)
			}
			if keyString == "$ref" {
				p.Ref = fmt.Sprintf("%v", val)
			}
		}
	default:
	}
	return p
}
