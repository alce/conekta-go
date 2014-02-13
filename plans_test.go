package conekta

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Plans", func() {

	server := NewTestServer()
	client := NewTestClient(server)

	plan := &Plan{
		Name:     "Gold Plan",
		Amount:   100000,
		Currency: "MXN",
	}

	Describe("Create", func() {
		It("Returns the created plan", func() {
			resp, err := client.Plans.Create(plan)

			Expect(err).ToNot(HaveOccurred())
			Expect(resp.Name).To(Equal("Gold Plan"))
		})
	})

	Describe("Update", func() {
		It("Returns the updated plan", func() {
			resp, err := client.Plans.Update("gold-plan", plan)

			Expect(err).ToNot(HaveOccurred())
			Expect(resp.Name).To(Equal("Gold Plan"))
		})
	})

	Describe("Delete", func() {
		It("Returns the deleted plan", func() {
			resp, err := client.Plans.Delete("gold-plan")

			Expect(err).ToNot(HaveOccurred())
			Expect(resp.Name).To(Equal("Gold Plan"))
		})
	})
})
