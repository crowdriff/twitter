package twitter_test

import (
	. "github.com/crowdriff/twitter"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Twitter", func() {
	Context("NewClient", func() {
		It("should return a new client", func() {
			consumerCreds := ConsumerCredentials{
				Key:    "somekey",
				Secret: "somesecret",
			}
			accessCreds := AccessCredentials{
				Token:  "sometoken",
				Secret: "somesecret",
			}

			client := NewClient(consumerCreds, accessCreds, nil)
			Ω(client).ShouldNot(Equal(nil))
		})
	})

})
