package asyncapi2_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/poim1205/asyncapi/asyncapi2"
)

var _ = Describe("Asyncapi2", func() {
	Context("NewAsyncAPI()", func() {
		When("returns a structure of type *asyncapi2.T", func() {
			It("should be `true`", func() {
				asyncApiStruct := asyncapi2.NewAsyncAPI()
				strTypeOfAsyncApiStruct := fmt.Sprintf("%T", asyncApiStruct)

				Expect(strTypeOfAsyncApiStruct).Should(Equal("*asyncapi2.T"))
			})
		})
	})
})
