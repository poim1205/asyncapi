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
				p[keyString] = NewParameter(val)
			}
		}
	default:
	}
	return p
}

func (ps Parameters) PrintParameters(indentString string) {
	fmt.Printf("%sparameters:\n", indentString)
	for k, v := range ps {
		fmt.Printf("%s%s%s:\n", indentString, INDENT, k)
		v.PrintParameter(fmt.Sprintf("%s%s", indentString, INDENT))
	}
}

type Parameter struct {
	Ref         string
	Description string
	Schema      *Schema
	Location    string
}

func NewParameter(v interface{}) *Parameter {
	p := Parameter{}

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

	return &p
}

func (value *Parameter) PrintParameter(indentString string) {
	fmt.Printf("%s%s$ref: %s\n", indentString, INDENT, value.Ref)
	fmt.Printf("%s%sdescription: %s\n", indentString, INDENT, value.Description)
	fmt.Printf("%s%slocation: %s\n", indentString, INDENT, value.Location)
}
