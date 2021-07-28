package asyncapi2

import (
	"fmt"
	"regexp"
	"strconv"
)

type Schema struct {
	Const                       string //const is a golang keyword
	Description                 string
	Discriminator               string
	Format                      string
	Id                          string
	Property                    string
	Ref                         string
	Title                       string
	Type                        string
	Enum                        []string
	Required                    []string
	ExclusiveMaximum            int
	ExclusiveMinimum            int
	Maximum                     int
	MaxItems                    int
	MaxLength                   int
	MaxProperties               int
	Minimum                     int
	MinItems                    int
	MinLength                   int
	MinProperties               int
	MultipleOf                  int
	Deprecated                  bool
	ItemsChildrenIsSingleSchema bool
	Pass                        bool
	UniqueItems                 bool
	Contains                    *Schema
	Else                        *Schema
	If                          *Schema // if/else are golang keywords
	Not                         *Schema
	Parent                      *Schema
	PropertyNames               *Schema
	RefSchema                   *Schema
	Then                        *Schema
	AllOf                       []*Schema
	AnyOf                       []*Schema
	ItemsChildren               []*Schema
	OneOf                       []*Schema
	PropertiesChildren          []*Schema
	PatternProperties           map[string]*Schema
	Properties                  map[string]*Schema /// string
	ExternalDocs                *ExternalDocs
	Pattern                     *regexp.Regexp
	AdditionalItems             interface{}
	AdditionalProperties        interface{}
	Dependencies                map[string]interface{}
}

func NewSchema() *Schema {
	return &Schema{}
}

func (s *Schema) SetValues(v interface{}) *Schema {
	switch mapSchema := v.(type) {
	case map[interface{}]interface{}:
		for key, val := range mapSchema {
			keyStr := fmt.Sprintf("%v", key)
			// Assign string values
			if keyStr == "const" {
				s.Const = fmt.Sprintf("%v", val)
			}
			if keyStr == "description" {
				s.Description = fmt.Sprintf("%v", val)
			}
			if keyStr == "discriminator" {
				s.Discriminator = fmt.Sprintf("%v", val)
			}
			if keyStr == "format" {
				s.Format = fmt.Sprintf("%v", val)
			}
			if keyStr == "$id" {
				s.Id = fmt.Sprintf("%v", val)
			}
			if keyStr == "property" {
				s.Property = fmt.Sprintf("%v", val)
			}
			if keyStr == "$ref" {
				s.Ref = fmt.Sprintf("%v", val)
			}
			if keyStr == "title" {
				s.Title = fmt.Sprintf("%v", val)
			}
			if keyStr == "type" {
				s.Type = fmt.Sprintf("%v", val)
			}
			// Assign array of strings
			if keyStr == "enum" {
				switch valSlice := val.(type) {
				case []interface{}:
					sliceEnum := make([]string, 0, 1)
					for _, valEnum := range valSlice {
						sliceEnum = append(sliceEnum, fmt.Sprintf("%v", valEnum))
					}
					s.Enum = sliceEnum
				default:

				}
			}
			if keyStr == "required" {
				switch valSlice := val.(type) {
				case []interface{}:
					sliceRequired := make([]string, 0, 1)
					for _, valRequired := range valSlice {
						sliceRequired = append(sliceRequired, fmt.Sprintf("%v", valRequired))
					}
					s.Required = sliceRequired
				default:

				}
			}
			// Assign int values
			if keyStr == "exclusiveMaximum" {
				intVal, err := strconv.Atoi(fmt.Sprintf("%v", val))
				if err == nil {
					s.ExclusiveMaximum = intVal
				}
			}
			if keyStr == "exclusiveMinimum" {
				intVal, err := strconv.Atoi(fmt.Sprintf("%v", val))
				if err == nil {
					s.ExclusiveMinimum = intVal
				}
			}
			if keyStr == "maximum" {
				intVal, err := strconv.Atoi(fmt.Sprintf("%v", val))
				if err == nil {
					s.Maximum = intVal
				}
			}
			if keyStr == "maxItems" {
				intVal, err := strconv.Atoi(fmt.Sprintf("%v", val))
				if err == nil {
					s.MaxItems = intVal
				}
			}
			if keyStr == "maxLength" {
				intVal, err := strconv.Atoi(fmt.Sprintf("%v", val))
				if err == nil {
					s.MaxLength = intVal
				}
			}
			if keyStr == "maxProperties" {
				intVal, err := strconv.Atoi(fmt.Sprintf("%v", val))
				if err == nil {
					s.MaxProperties = intVal
				}
			}
			if keyStr == "minimum" {
				intVal, err := strconv.Atoi(fmt.Sprintf("%v", val))
				if err == nil {
					s.Minimum = intVal
				}
			}
			if keyStr == "minItems" {
				intVal, err := strconv.Atoi(fmt.Sprintf("%v", val))
				if err == nil {
					s.MinItems = intVal
				}
			}
			if keyStr == "minLength" {
				intVal, err := strconv.Atoi(fmt.Sprintf("%v", val))
				if err == nil {
					s.MinLength = intVal
				}
			}
			if keyStr == "minProperties" {
				intVal, err := strconv.Atoi(fmt.Sprintf("%v", val))
				if err == nil {
					s.MinProperties = intVal
				}
			}
			// Assign bool values
			if keyStr == "deprecated" {
				boolVal, err := strconv.ParseBool(fmt.Sprintf("%v", val))
				if err == nil {
					s.Deprecated = boolVal
				} else {
					s.Deprecated = false
				}
			}
			if keyStr == "itemsChildrenIsSingleSchema" {
				boolVal, err := strconv.ParseBool(fmt.Sprintf("%v", val))
				if err == nil {
					s.ItemsChildrenIsSingleSchema = boolVal
				} else {
					s.ItemsChildrenIsSingleSchema = false
				}
			}
			if keyStr == "pass" {
				boolVal, err := strconv.ParseBool(fmt.Sprintf("%v", val))
				if err == nil {
					s.Pass = boolVal
				} else {
					s.Pass = false
				}
			}
			if keyStr == "uniqueItems" {
				boolVal, err := strconv.ParseBool(fmt.Sprintf("%v", val))
				if err == nil {
					s.UniqueItems = boolVal
				} else {
					s.UniqueItems = false
				}
			}
			// Assign inner Schema values
			if keyStr == "contains" {
				newSchema := NewSchema()
				s.Contains = newSchema.SetValues(val)
			}
			if keyStr == "else" {
				newSchema := NewSchema()
				s.Else = newSchema.SetValues(val)
			}
			if keyStr == "if" {
				newSchema := NewSchema()
				s.If = newSchema.SetValues(val)
			}
			if keyStr == "not" {
				newSchema := NewSchema()
				s.Not = newSchema.SetValues(val)
			}
			if keyStr == "parent" {
				newSchema := NewSchema()
				s.Parent = newSchema.SetValues(val)
			}
			if keyStr == "propertyNames" {
				newSchema := NewSchema()
				s.PropertyNames = newSchema.SetValues(val)
			}
			if keyStr == "refSchema" {
				newSchema := NewSchema()
				s.RefSchema = newSchema.SetValues(val)
			}
			if keyStr == "then" {
				newSchema := NewSchema()
				s.Then = newSchema.SetValues(val)
			}
			// Assign slices of inner Schema values
			if keyStr == "allOf" {
				switch sliceSchema := val.(type) {
				case []interface{}:
					sliceSch := make([]*Schema, 0)
					for _, schemaVal := range sliceSchema {
						newSchema := NewSchema()
						sliceSch = append(sliceSch, newSchema.SetValues(schemaVal))
					}
					s.AllOf = sliceSch
				default:

				}
			}
			if keyStr == "anyOf" {
				switch sliceSchema := val.(type) {
				case []interface{}:
					sliceSch := make([]*Schema, 0)
					for _, schemaVal := range sliceSchema {
						newSchema := NewSchema()
						sliceSch = append(sliceSch, newSchema.SetValues(schemaVal))
					}
					s.AnyOf = sliceSch
				default:

				}
			}
			if keyStr == "itemsChildren" {
				switch sliceSchema := val.(type) {
				case []interface{}:
					sliceSch := make([]*Schema, 0)
					for _, schemaVal := range sliceSchema {
						newSchema := NewSchema()
						sliceSch = append(sliceSch, newSchema.SetValues(schemaVal))
					}
					s.ItemsChildren = sliceSch
				default:

				}
			}
			if keyStr == "oneOf" {
				switch sliceSchema := val.(type) {
				case []interface{}:
					sliceSch := make([]*Schema, 0)
					for _, schemaVal := range sliceSchema {
						newSchema := NewSchema()
						sliceSch = append(sliceSch, newSchema.SetValues(schemaVal))
					}
					s.OneOf = sliceSch
				default:

				}
			}
			if keyStr == "propertiesChildren" {
				switch sliceSchema := val.(type) {
				case []interface{}:
					sliceSch := make([]*Schema, 0)
					for _, schemaVal := range sliceSchema {
						newSchema := NewSchema()
						sliceSch = append(sliceSch, newSchema.SetValues(schemaVal))
					}
					s.PropertiesChildren = sliceSch
				default:

				}
			}
			// Assign map of inner Schema values
			if keyStr == "patternProperties" {
				switch mapValues := val.(type) {
				case map[interface{}]interface{}:
					mapSchema := make(map[string]*Schema)
					for innerKey, innerVal := range mapValues {
						keyStr := fmt.Sprintf("%v", innerKey)

						_, Ok := mapSchema[keyStr]

						if !Ok {
							newSchema := NewSchema()
							mapSchema[keyStr] = newSchema.SetValues(innerVal)
						}

					}
					s.PatternProperties = mapSchema
				default:

				}
			}
			if keyStr == "properties" {
				switch mapValues := val.(type) {
				case map[interface{}]interface{}:
					mapSchema := make(map[string]*Schema)
					for innerKey, innerVal := range mapValues {
						keyStr := fmt.Sprintf("%v", innerKey)

						_, Ok := mapSchema[keyStr]

						if !Ok {
							newSchema := NewSchema()
							mapSchema[keyStr] = newSchema.SetValues(innerVal)
						}

					}
					s.Properties = mapSchema
				default:

				}
			}
			// Assigning external documents values
			if keyStr == "externalDocs" {
				newExternalDocs := NewExternalDocs()
				s.ExternalDocs = newExternalDocs.SetValues(val)
			}
			// Assigning regex pattern
			if keyStr == "pattern" {
				strPattern := fmt.Sprintf("%v", val)
				ptrRegex, err := regexp.Compile(strPattern)
				if err == nil {
					s.Pattern = ptrRegex
				}
			}
			// Assigning additional values
			if keyStr == "additionalItems" {
				s.AdditionalItems = val
			}
			if keyStr == "additionalProperties" {
				s.AdditionalProperties = val
			}
			// Assigning dependencies
			if keyStr == "dependencies" {
				switch mapValues := val.(type) {
				case map[interface{}]interface{}:
					mapInterface := make(map[string]interface{})
					for innerKey, innerVal := range mapValues {
						keyStr := fmt.Sprintf("%v", innerKey)

						_, Ok := mapInterface[keyStr]

						if !Ok {
							mapInterface[keyStr] = innerVal
						}

					}
					s.Dependencies = mapInterface
				default:

				}
			}
		}

	default:
	}

	return s
}
