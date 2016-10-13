package twitter_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestTwitter(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Twitter Suite")
}
