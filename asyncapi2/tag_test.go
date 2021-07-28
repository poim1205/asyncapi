package asyncapi2_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/poim1205/asyncapi/asyncapi2"
	"gopkg.in/yaml.v2"
)

var _ = Describe("Tag", func() {
	Context("MapTags interface", func() {
		var (
			mapTags          asyncapi2.Tags
			mapTagsInterface map[interface{}]interface{}
		)
		BeforeEach(func() {
			mapTagsInterface = make(map[interface{}]interface{})
			yamlTagByte := []byte(`tags:
  - name: tag name 1
    description: description of the tag 1
  - name: tag name 2
    description: description of the tag 2
    externalDocs:
      description: API Documentation
      url: https://openweathermap.org/api`)

			err := yaml.Unmarshal(yamlTagByte, &mapTagsInterface)
			Expect(err).ShouldNot(HaveOccurred())

			//logrus.Infof("Tags: %v", mapTagsInterface)
			newMapTags := asyncapi2.NewTags()
			mapTags = newMapTags.SetValues(mapTagsInterface["tags"])
		})
		When("NewTags is filled with a map structure", func() {
			It("should return a structure of type Tags", func() {

				Expect(fmt.Sprintf("%T", mapTags)).Should(Equal("asyncapi2.Tags"))
			})
		})
		When("MapTags is filled with a map structure", func() {
			It("should return 2 instances of Tag", func() {
				Expect(len(mapTags)).Should(Equal(2))
			})
			It("should return a name `tag name 1` for first tag", func() {
				Expect(mapTags[0].Name).Should(Equal("tag name 1"))
			})
			It("should return a name `tag name 2` for second tag", func() {
				Expect(mapTags[1].Name).Should(Equal("tag name 2"))
			})
			It("should return a description `description of the tag 1` for first tag", func() {
				Expect(mapTags[0].Description).Should(Equal("description of the tag 1"))
			})
			It("should return a description `description of the tag 2` for second tag", func() {
				Expect(mapTags[1].Description).Should(Equal("description of the tag 2"))
			})
			It("should return an external document for mapTags[1].externalDocs.description `API Documentation` for second tag", func() {
				Expect(mapTags[1].ExternalDocs.Description).Should(Equal("API Documentation"))
			})
		})
	})
})
