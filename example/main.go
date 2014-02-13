package main

import (
	"github.com/alce/conekta"
	"log"
)

var client *conekta.Client

func init() {
	// To test on a sandbox account, set CONEKTA_API_KEY environment variable
	// os.Setenv("CONEKTA_API_KEY", "THE_KEY")
	client = conekta.NewClient()

	// To test on local server, uncomment to set the base url:
	// u, _ := url.Parse("http://localhost:3000")
	// client.BaseURL(u)
}

func main() {
	createAndUpdateCustomer()
	createCardCharge()
	createPlan()
}

func createAndUpdateCustomer() {
	// Create a customer
	c := &conekta.Customer{
		Name:  "Logan",
		Email: "no@email.com",
		Phone: "222-333-444",
	}

	customer, err := client.Customers.Create(c)
	if err != nil {
		log.Println(err)
	}
	log.Println(customer.Name)

	// Update the customer
	customer.Name = "Xavier"
	updatedCustomer, err := client.Customers.Update(customer.Id, customer)
	if err != nil {
		log.Println(err)
	}
	log.Println(updatedCustomer.Name)
}

func createCardCharge() {
	c := &conekta.Charge{
		Description: "Stogies",
		Amount:      20000,
		Currency:    "MXN",
		ReferenceId: "9839-wolf_pack",
		Card:        "tok_test_visa_4242",
	}

	charge, err := client.Charges.Create(c)
	if err != nil {
		log.Println(err)
	}

	log.Println(charge)
}

func createPlan() {
	p := &conekta.Plan{
		Name: "Golden Boy",
		Amount: 333333,
	}

	plan, err := client.Plans.Create(p)
	if err != nil {
		log.Println(err)
	}

	log.Println(plan)
}
