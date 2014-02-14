package conekta

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const tokenTestVisa = "tok_test_visa_4242"

var _ = Describe("Charges Resource", func() {

	server := NewTestServer()
	client := NewTestClient(server)

	Describe("Creating a charge", func() {

		var charge *Charge

		BeforeEach(func() {
			charge = &Charge{
				Amount:      10000,
				Currency:    "MXN",
				Description: "DVD - Zorro",
				ReferenceId: "wahtever",
				Details: &Details{
					Name:  "Wolverine",
					Email: "logan@x-men.com",
					Phone: "345-1234-1234",
				},
			}
		})

		Describe("OXXO", func() {
			It("Should return a succesful charge", func() {
				charge.Cash = CashPayment{"type": "oxxo"}
				resp, err := client.Charges.Create(charge)

				Expect(err).ToNot(HaveOccurred())
				Expect(*resp.Id).To(Equal("52f88a63d7e1a0e1a20000b4"))
				Expect(resp.PaymentMethod.Type).To(Equal("oxxo"))
				Expect(*resp.Livemode).To(BeFalse())
			})
		})

		Describe("Credit Card", func() {
			It("Should return a succesful charge", func() {
				charge.Card = tokenTestVisa
				resp, err := client.Charges.Create(charge)

				Expect(err).ToNot(HaveOccurred())
				Expect(*resp.Id).To(Equal("52f89639d7e1a09657000007"))
				Expect(resp.PaymentMethod.Brand).To(Equal("visa"))
				Expect(resp.PaymentMethod.Last4).To(Equal("4242"))
			})

			PIt("Should return a token processing error", func() {})

			PIt("Should return a card declined error", func() {})
		})

		Describe("Bank deposit", func() {
			It("Should return a succesful charge", func() {
				charge.Bank = BankPayment{"type": "banorte"}
				resp, err := client.Charges.Create(charge)

				Expect(err).ToNot(HaveOccurred())
				Expect(*resp.Id).To(Equal("52f8901cd7e1a0e1a20000c7"))
				Expect(resp.PaymentMethod.ServiceName).To(Equal("Conekta"))
				Expect(resp.PaymentMethod.Type).To(Equal("banorte"))
				Expect(resp.PaymentMethod.Reference).To(Equal("0068916"))
			})
		})
	})

	Describe("Create charge with advanced call", func() {
		It("Returs the created charge", func() {
			charge := &Charge{
				Description: "Stogies",
				Amount:      20000,
				Currency:    "MXN",
				ReferenceId: "9839-wolf_pack",
				Card:        "tok_test_visa_4242",
				Details: &Details{
					Name:        "Wolverine",
					Email:       "logan@x-men.org",
					Phone:       "",
					DateOfBirth: "1980-09-24",
					BillingAddress: &BillingAddress{
						TaxId:       "wibble",
						CompanyName: "",
						Address: &Address{
							Street1: "Street1",
							Street2: "Street2",
							City:    "Springfield",
							State:   "NJ",
							Zip:     "1222",
						},
						Phone: "234-567-888",
						Email: "another@email.com",
					},
					LineItems: []LineItem{
						LineItem{
							Name:        "Box of Cohibe S1s",
							Sku:         "cohib_s1",
							UnitPrice:   200000,
							Description: "Imported from Mex",
							Quantity:    1,
							Type:        "whatever",
						},
					},
					Shipment: &Shipment{
						Carrier:    "estafeta",
						Service:    "international",
						TrackingId: "satoheusatohe",
						Price:      20000,
						Address: &Address{
							Street1: "Street1",
							Street2: "Street2",
							City:    "Springfield",
							State:   "NJ",
							Zip:     "1222",
						},
					},
				},
			}
			resp, err := client.Charges.Create(charge)

			Expect(err).ToNot(HaveOccurred())
			Expect(*resp.Id).To(Equal("52f89639d7e1a09657000007"))
			Expect(resp.PaymentMethod.Brand).To(Equal("visa"))
			Expect(resp.Details.Name).To(Equal("wolverine"))
			Expect(resp.Details.LineItems[0].Name).To(Equal("Box of Cohiba 51s"))
			Expect(resp.Details.Shipment.Carrier).To(Equal("estafeta"))
			Expect(resp.Details.Shipment.Address.Country).To(Equal("Canada"))

		})
	})

	Describe("Retrieve charge", func() {
		It("Returns the requested charge", func() {
			resp, err := client.Charges.Get("52f8901cd7e1a0e1a20000c7")

			Expect(err).ToNot(HaveOccurred())
			Expect(resp.PaymentMethod.ServiceName).To(Equal("Conekta"))
			Expect(resp.PaymentMethod.Type).To(Equal("banorte"))
			Expect(resp.PaymentMethod.Reference).To(Equal("0068916"))
			Expect(resp.Amount).To(Equal(30000))
		})
	})

	Describe("Refund", func() {
		It("Process a full refund", func() {
			resp, err := client.Charges.Refund("523df826aef8786485000001", nil)

			Expect(err).ToNot(HaveOccurred())
			Expect(resp.ReferenceId).To(Equal("9839-wolf_pack"))
			Expect(*resp.AmountRefunded).To(Equal(20000))
		})

		It("Process a partial refund", func() {
			resp, err := client.Charges.Refund("523df826aef8786485000001", Param{"amount": 10000})

			Expect(err).ToNot(HaveOccurred())
			Expect(resp.ReferenceId).To(Equal("9839-wolf_pack"))
			Expect(*resp.AmountRefunded).To(Equal(20000))
		})
	})

})
