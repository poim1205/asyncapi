package asyncapi2_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestAsyncapi2(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Asyncapi2 Suite")
}
