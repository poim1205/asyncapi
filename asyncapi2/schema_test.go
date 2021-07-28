package asyncapi2_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/poim1205/asyncapi/asyncapi2"
	"gopkg.in/yaml.v2"
)

var _ = Describe("Schema", func() {
	Context("MapSchema interface", func() {
		var (
			mapSchema          *asyncapi2.Schema
			mapSchemaInterface map[interface{}]interface{}
		)
		BeforeEach(func() {
			mapSchemaInterface = make(map[interface{}]interface{})
			yamlSchemaByte := []byte(`const: demo const
description: demo description
discriminator: demo discriminator
format: demo format
$id: demo id
property: demo property
$ref: demo $ref
title: demo title
type: object
enum:
  - val1
  - val2
  - val3
required:
  - block required 1
  - block required 2
exclusiveMaximum: 12
exclusiveMinimum: 2
maximum: 100
maxItems: 50
maxLength: 500
maxProperties: 5
minimum: 10
minItems: 5
minLength: 5
minProperties: 2
deprecated: false
itemsChildrenIsSingleSchema: true
pass: true
uniqueItems: false
contains:
  type: object
  title: contains schema
else:
  type: object
  title: else schema
if:
  type: object
  title: if schema
not:
  type: object
  title: not schema
parent:
  type: object
  title: parent schema
propertyNames:
  type: object
  title: property name schema
refSchema: 
  type: object
  title: refSchema schema
then: 
  type: object
  title: then schema
allOf:
  - type: object
    title: allOf schema 1
  - type: object
    title: allOf schema 2
anyOf:
  - type: object
    title: anyOf schema 1
  - type: object
    title: anyOf schema 2
itemsChildren:
  - type: object
    title: itemsChildren schema 1
  - type: object
    title: itemsChildren schema 2
oneOf:
  - type: object
    title: oneOf schema 1
  - type: object
    title: oneOf schema 2
propertiesChildren:
  - type: object
    title: propertiesChildren schema 1
  - type: object
    title: propertiesChildren schema 2
patternProperties:
  patterlumens:
    type: integer
    minimum: 0
    description: Light intensity measured in lumens.
  pattersentAt:
    $ref: "#/components/schemas/sentAt"
properties:
  lumens:
    type: integer
    minimum: 0
    description: Light intensity measured in lumens.
  sentAt:
    $ref: "#/components/schemas/sentAt"
externalDocs:
  description: API Documentation
  url: https://openweathermap.org/api
pattern: "foo.?"
additionalItems: false
additionalProperties:
  type: string
dependencies:
  lumens: 
    - sentAt
`)

			err := yaml.Unmarshal(yamlSchemaByte, &mapSchemaInterface)
			Expect(err).ShouldNot(HaveOccurred())

			newMapSchema := asyncapi2.NewSchema()
			mapSchema = newMapSchema.SetValues(mapSchemaInterface)
		})
		When("NewSchema is filled with a map structure", func() {
			It("should return a pointer to a structure of type Schema", func() {

				Expect(fmt.Sprintf("%T", mapSchema)).Should(Equal("*asyncapi2.Schema"))
			})
		})
		When("MapSchema is filled with a map structure", func() {
			It("should return a field `const` with value `demo const`", func() {
				Expect(mapSchema.Const).Should(Equal("demo const"))
			})
			It("should return a field `description` with value `demo description`", func() {
				Expect(mapSchema.Description).Should(Equal("demo description"))
			})
			It("should return a field `discriminator` with value `demo discriminator`", func() {
				Expect(mapSchema.Discriminator).Should(Equal("demo discriminator"))
			})
			It("should return a field `$id` with value `demo id`", func() {
				Expect(mapSchema.Id).Should(Equal("demo id"))
			})
			It("should return a field `property` with value `demo property`", func() {
				Expect(mapSchema.Property).Should(Equal("demo property"))
			})
			It("should return a field `$ref` with value `demo $ref`", func() {
				Expect(mapSchema.Ref).Should(Equal("demo $ref"))
			})
			It("should return a field `title` with value being `demo title`", func() {
				Expect(mapSchema.Title).Should(Equal("demo title"))
			})
			It("should return a field `type` with value being `object`", func() {
				Expect(mapSchema.Type).Should(Equal("object"))
			})
			It("should return 3 values in `enum` field", func() {
				Expect(len(mapSchema.Enum)).Should(Equal(3))
			})
			It("should return 2 values in `required` field", func() {
				Expect(len(mapSchema.Required)).Should(Equal(2))
			})
			It("should return a field `exclusiveMaximum` with a values of `12`", func() {
				Expect(mapSchema.ExclusiveMaximum).Should(Equal(12))
			})
			It("should return a field `exclusiveMinimum` with a values of `2`", func() {
				Expect(mapSchema.ExclusiveMinimum).Should(Equal(2))
			})
			It("should return a field `maximum` with a values of `100`", func() {
				Expect(mapSchema.Maximum).Should(Equal(100))
			})
			It("should return a field `maxItems` with a values of `50`", func() {
				Expect(mapSchema.MaxItems).Should(Equal(50))
			})
			It("should return a field `maxLength` with a values of `500`", func() {
				Expect(mapSchema.MaxLength).Should(Equal(500))
			})
			It("should return a field `maxProperties` with a values of `5`", func() {
				Expect(mapSchema.MaxProperties).Should(Equal(5))
			})
			It("should return a field `minimum` with a values of `10`", func() {
				Expect(mapSchema.Minimum).Should(Equal(10))
			})
			It("should return a field `minItems` with a values of `5`", func() {
				Expect(mapSchema.MinItems).Should(Equal(5))
			})
			It("should return a field `minLength` with a values of `5`", func() {
				Expect(mapSchema.MinLength).Should(Equal(5))
			})
			It("should return a field `minProperties` with a values of `2`", func() {
				Expect(mapSchema.MinProperties).Should(Equal(2))
			})
			It("should return a field `deprecated` with a values of `false`", func() {
				Expect(mapSchema.Deprecated).Should(Equal(false))
			})
			It("should return a field `itemsChildrenIsSingleSchema` with a values of `true`", func() {
				Expect(mapSchema.ItemsChildrenIsSingleSchema).Should(Equal(true))
			})
			It("should return a field `pass` with a values of `true`", func() {
				Expect(mapSchema.Pass).Should(Equal(true))
			})
			It("should return a field `uniqueItems` with a values of `false`", func() {
				Expect(mapSchema.UniqueItems).Should(Equal(false))
			})
			It("should return a field `contains.title` with value being `contains schema`", func() {
				Expect(mapSchema.Contains.Title).Should(Equal("contains schema"))
			})
			It("should return a field `else.title` with value being `else schema`", func() {
				Expect(mapSchema.Else.Title).Should(Equal("else schema"))
			})
			It("should return a field `if.title` with value being `if schema`", func() {
				Expect(mapSchema.If.Title).Should(Equal("if schema"))
			})
			It("should return a field `not.title` with value being `not schema`", func() {
				Expect(mapSchema.Not.Title).Should(Equal("not schema"))
			})
			It("should return a field `parent.title` with value being `parent schema`", func() {
				Expect(mapSchema.Parent.Title).Should(Equal("parent schema"))
			})
			It("should return a field `propertyNames.title` with value being `property name schema`", func() {
				Expect(mapSchema.PropertyNames.Title).Should(Equal("property name schema"))
			})
			It("should return a field `refSchema.title` with value being `refSchema schema`", func() {
				Expect(mapSchema.RefSchema.Title).Should(Equal("refSchema schema"))
			})
			It("should return a field `then.title` with value being `then schema`", func() {
				Expect(mapSchema.Then.Title).Should(Equal("then schema"))
			})
			It("should return 2 values in `allOf` slice", func() {
				Expect(len(mapSchema.AllOf)).Should(Equal(2))
			})
			It("should return 2 values in `anyOf` slice", func() {
				Expect(len(mapSchema.AnyOf)).Should(Equal(2))
			})
			It("should return 2 values in `itemsChildren` slice", func() {
				Expect(len(mapSchema.ItemsChildren)).Should(Equal(2))
			})
			It("should return 2 values in `oneOf` slice", func() {
				Expect(len(mapSchema.OneOf)).Should(Equal(2))
			})
			It("should return 2 values in `propertiesChildren` slice", func() {
				Expect(len(mapSchema.PropertiesChildren)).Should(Equal(2))
			})
			It("should return a field `propertiesChildren[1].title` with value being `propertiesChildren schema 2`", func() {
				Expect(mapSchema.PropertiesChildren[1].Title).Should(Equal("propertiesChildren schema 2"))
			})
			It("should return 2 values in `patternProperties` map", func() {
				Expect(len(mapSchema.PatternProperties)).Should(Equal(2))
			})
			It("should return 2 values in `properties` map", func() {
				Expect(len(mapSchema.Properties)).Should(Equal(2))
			})
			It("should return a field `properties[lumens].description` with value being `Light intensity measured in lumens.`", func() {
				mapVal, _ := mapSchema.Properties["lumens"]
				Expect(mapVal.Description).Should(Equal("Light intensity measured in lumens."))
			})
			It("should return a field `externalDocs.description` with value being `API Documentation`", func() {
				Expect(mapSchema.ExternalDocs.Description).Should(Equal("API Documentation"))
			})
			It("should return a field `pattern` with value being `food`", func() {
				Expect(fmt.Sprintf("%s", mapSchema.Pattern.Find([]byte(`seafood fool`)))).Should(Equal("food"))
			})
			It("should return a true value when searching for `dependencies[lumens]`", func() {
				_, Ok := mapSchema.Dependencies["lumens"]
				Expect(Ok).Should(Equal(true))
			})
		})
	})
})

/*

pattern: "^(\\([0-9]{3}\\))?[0-9]{3}-[0-9]{4}$"
additionalItems: false
additionalProperties:
  type: string
dependencies:
  lumens:
    - sentAt
*/
