package conekta

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Customers Resource", func() {
	server := NewTestServer()
	client := NewTestClient(server)

	var customer *Customer

	BeforeEach(func() {
		customer = &Customer{
			Name:  "Thomas Logan",
			Email: "thomas.logan@xmen.org",
			Phone: "55-5555-5555",
		}
	})

	Describe("Create", func() {
		Context("Success", func() {
			It("Returns the created customer", func() {
				resp, err := client.Customers.Create(customer)

				Expect(err).ToNot(HaveOccurred())
				Expect(resp.Name).To(Equal("Thomas Logan"))
				Expect(resp.BillingAddress.Street1).To(Equal("77 Mystery Lane"))
				Expect(resp.ShippingAddress.Street1).To(Equal("250 Alexis St"))
				Expect(resp.Cards[0].Name).To(Equal("Thomas Logan"))
				Expect(resp.Subscription.PlanId).To(Equal("gold-plan"))
			})

		})
		Context("Failure", func() {})
	})

	Describe("Update", func() {
		Context("Success", func() {
			It("Returns the updated customer", func() {
				customer := &Customer{Name: "Logan", Email: "Whatever@me.com"}
				resp, err := client.Customers.Update("cus_k2D9DxlqdVTagmEd400001", customer)

				Expect(err).ToNot(HaveOccurred())
				Expect(resp.Name).To(Equal("Thomas Logan"))

			})
		})
	})

	Describe("Delete", func() {
		It("Returns the deleted customer", func() {
			resp, _ := client.Customers.Delete("cus_k2D9DxlqdVTagmEd400001")

			Expect(resp.Name).To(Equal("Thomas Logan"))

		})
	})

	Context("Handling Credit Cards", func() {
		Describe("Add  Credit Card", func() {
			It("Returns the created Card", func() {
				card := &CreditCard{Token: "tok_8kZwafM8IcN23Nd9"}
				resp, err := client.Customers.AddCreditCard("cus_k2D9DxlqdVTagmEd400001", card)

				Expect(err).NotTo(HaveOccurred())
				Expect(resp.Object).To(Equal("card"))
				Expect(resp.Address.Street1).To(Equal("250 Alexis St"))
			})
		})

		Describe("Update Credit Card", func() {
			It("Returns the updated Card", func() {
				card := &CreditCard{Token: "tok_8kZwafM8IcN23Nd9", Id: "xyz"}
				resp, err := client.Customers.UpdateCreditCard("cus_k2D9DxlqdVTagmEd400001", card)

				Expect(err).NotTo(HaveOccurred())
				Expect(resp.Object).To(Equal("card"))
				Expect(resp.Address.Street1).To(Equal("250 Alexis St"))

			})
		})
		Describe("Delete Credit Card", func() {
			It("Returns the deleted Card", func() {
				resp, err := client.Customers.DeleteCreditCard("cus_k2D9DxlqdVTagmEd400001", "card_TCuBjUEcy9r41Fk2")

				Expect(err).NotTo(HaveOccurred())
				Expect(resp.Object).To(Equal("card"))

			})
		})
	})

	Context("Handling Subscriptions", func() {

		const (
			customerId = "cus_Z9cVem5W3Rus2TAs7"
		)

		sub := &Subscription{
			PlanId: "opal-plan",
			CardId: "card_vow1u83899Rkj5LM",
		}

		Describe("Create", func() {
			It("Returns the created subscription", func() {
				resp, err := client.Customers.CreateSubscription(customerId, sub)

				Expect(err).NotTo(HaveOccurred())
				Expect(resp.Object).To(Equal("subscription"))
				Expect(resp.PlanId).To(Equal("gold-plan"))
			})
		})

		Describe("Change", func() {
			It("Returns the changed subscription", func() {
				resp, err := client.Customers.UpdateSubscription(customerId, sub)

				Expect(err).NotTo(HaveOccurred())
				Expect(resp.Object).To(Equal("subscription"))
				Expect(resp.PlanId).To(Equal("gold-plan"))
			})
		})

		Describe("Pause", func() {
			It("Returns the paused subscription", func() {
				resp, err := client.Customers.PauseSubscription(customerId)

				Expect(err).NotTo(HaveOccurred())
				Expect(resp.Object).To(Equal("subscription"))
				Expect(resp.PlanId).To(Equal("gold-plan"))
			})
		})

		Describe("Resume", func() {
			It("Returns the resumed subscription", func() {
				resp, err := client.Customers.ResumeSubscription(customerId)

				Expect(err).NotTo(HaveOccurred())
				Expect(resp.Object).To(Equal("subscription"))
				Expect(resp.PlanId).To(Equal("gold-plan"))

			})
		})

		Describe("Cancel", func() {
			It("Returns the cancelled subscription", func() {
				resp, err := client.Customers.CancelSubscription(customerId)

				Expect(err).NotTo(HaveOccurred())
				Expect(resp.Object).To(Equal("subscription"))
				Expect(resp.PlanId).To(Equal("gold-plan"))

			})
		})
	})
})
