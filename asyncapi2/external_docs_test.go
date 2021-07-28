package asyncapi2_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/poim1205/asyncapi/asyncapi2"
	"gopkg.in/yaml.v2"
)

var _ = Describe("ExternalDocs", func() {
	Context("MapExternalDocs interface", func() {
		var (
			mapExternalDocs          *asyncapi2.ExternalDocs
			mapExternalDocsInterface map[interface{}]interface{}
		)
		BeforeEach(func() {
			mapExternalDocsInterface = make(map[interface{}]interface{})
			yamlExternalDocsByte := []byte(`externalDocs:
  description: API Documentation
  url: https://openweathermap.org/api`)

			err := yaml.Unmarshal(yamlExternalDocsByte, &mapExternalDocsInterface)
			Expect(err).ShouldNot(HaveOccurred())

			newMapExternalDocs := asyncapi2.NewExternalDocs()
			mapExternalDocs = newMapExternalDocs.SetValues(mapExternalDocsInterface["externalDocs"])
		})
		When("NewExternalDocs is filled with a map structure", func() {
			It("should return a pointer to a structure of type ExternalDocs", func() {

				Expect(fmt.Sprintf("%T", mapExternalDocs)).Should(Equal("*asyncapi2.ExternalDocs"))
			})
		})
		When("MapExternalDocs is filled with a map structure", func() {
			It("should return an `API Documentation` in the `Description` field", func() {
				Expect(mapExternalDocs.Description).Should(Equal("API Documentation"))
			})
			It("should return an `https://openweathermap.org/api` in the `Url` field", func() {
				Expect(mapExternalDocs.URL).Should(Equal("https://openweathermap.org/api"))
			})
		})
	})
})
