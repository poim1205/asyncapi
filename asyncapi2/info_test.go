package asyncapi2_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/poim1205/asyncapi/asyncapi2"
	"gopkg.in/yaml.v2"
)

var _ = Describe("Info", func() {
	var (
		infoInterface map[string]interface{}
	)
	BeforeEach(func() {
		infoInterface = make(map[string]interface{})
		yamlInfoByte := []byte(`info:
  title: my title
  version: '1.0.1'
  description: my description
  termsOfService: my terms of service
  contact:
    name: my contact name
    url: 'https://my.contact.url'
    email: my@contactemail.com
  license:
    name: my license name
    url: 'https://my.license.url'`)

		err := yaml.Unmarshal(yamlInfoByte, &infoInterface)
		Expect(err).ShouldNot(HaveOccurred())
	})
	Context("NewInfo()", func() {
		When("returns a structure of type *asyncapi2.Info", func() {
			It("should be `true`", func() {
				infoStruct := asyncapi2.NewInfo()

				Expect(fmt.Sprintf("%T", infoStruct)).Should(Equal("*asyncapi2.Info"))
			})
		})
	})
	Context("SetValues(interface{}) from yaml construct", func() {

		When("returns a structure of type *asyncapi2.Info", func() {
			It("should be `true`", func() {
				infoStruct := asyncapi2.NewInfo()

				Expect(fmt.Sprintf("%T", infoStruct.SetValues(infoInterface))).Should(Equal("*asyncapi2.Info"))
			})
		})
		When("values are assigned", func() {
			It("should be Equal to `my title` for info.title", func() {
				infoStruct := asyncapi2.NewInfo()
				testInfoStruct := infoStruct.SetValues(infoInterface["info"])

				Expect(testInfoStruct.Title).Should(Equal("my title"))
			})
			It("should be Equal to `1.0.1` for info.version", func() {
				infoStruct := asyncapi2.NewInfo()
				testInfoStruct := infoStruct.SetValues(infoInterface["info"])

				Expect(testInfoStruct.Version).Should(Equal("1.0.1"))
			})
			It("should be Equal to `my description` for info.description", func() {
				infoStruct := asyncapi2.NewInfo()
				testInfoStruct := infoStruct.SetValues(infoInterface["info"])

				Expect(testInfoStruct.Description).Should(Equal("my description"))
			})
			It("should be Equal to `my terms of service` for info.termsOfService", func() {
				infoStruct := asyncapi2.NewInfo()
				testInfoStruct := infoStruct.SetValues(infoInterface["info"])

				Expect(testInfoStruct.TermsOfService).Should(Equal("my terms of service"))
			})
			It("should be Equal to `my contact name` for info.contact.name", func() {
				infoStruct := asyncapi2.NewInfo()
				testInfoStruct := infoStruct.SetValues(infoInterface["info"])

				Expect(testInfoStruct.Contact.Name).Should(Equal("my contact name"))
			})
			It("should be Equal to `https://my.contact.url` for info.contact.url", func() {
				infoStruct := asyncapi2.NewInfo()
				testInfoStruct := infoStruct.SetValues(infoInterface["info"])

				Expect(testInfoStruct.Contact.Url).Should(Equal("https://my.contact.url"))
			})
			It("should be Equal to `my@contactemail.com` for info.contact.email", func() {
				infoStruct := asyncapi2.NewInfo()
				testInfoStruct := infoStruct.SetValues(infoInterface["info"])

				Expect(testInfoStruct.Contact.Email).Should(Equal("my@contactemail.com"))
			})
			It("should be Equal to `my license name` for info.license.name", func() {
				infoStruct := asyncapi2.NewInfo()
				testInfoStruct := infoStruct.SetValues(infoInterface["info"])

				Expect(testInfoStruct.License.Name).Should(Equal("my license name"))
			})
			It("should be Equal to `https://my.license.url` for info.license.url", func() {
				infoStruct := asyncapi2.NewInfo()
				testInfoStruct := infoStruct.SetValues(infoInterface["info"])

				Expect(testInfoStruct.License.Url).Should(Equal("https://my.license.url"))
			})
		})
	})
	// Context("SetValues(interface{}) from yaml construct", func() {
	// 	When("values are assigned", func() {
	// 		It("should be Equal to `my title` for info.title", func() {
	// 			infoStruct := asyncapi2.NewInfo()
	// 			testInfoStruct := infoStruct.SetValues(infoInterface["info"])

	// 			Expect(testInfoStruct.Title).Should(Equal("my title"))
	// 		})
	// 	})
	// })
})
