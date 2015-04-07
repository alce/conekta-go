package test

import (
	"log"

	base "github.com/nubleer/conekta-go/conekta"
)

type Customer struct {
	Id     string `db:"-" json:"id"`
	Name   string `db:"-" json:"name"`
	Email  string `db:"-" json:"email"`
	Card   *base.CreditCard
	Charge *base.Charge
	*base.Client
}

func NewCustomer(customerId string, cardToken string) *Customer {
	customer := &Customer{}
	customer.Client = base.NewClient()
	customer.Card = &base.CreditCard{}

	customer.Id = customerId
	customer.Card.Token = cardToken
	return customer
}

func (self *Customer) Pause() (err error) {
	subscription, err := self.Customers.PauseSubscription(self.Id)
	if err != nil {
		return
	}
	log.Printf("Subscripción pausada: $v", subscription)
	return
}

func (self *Customer) Resume() (err error) {
	subscription, err := self.Customers.ResumeSubscription(self.Id)
	if err != nil {
		return
	}
	log.Printf("Subscripción activada: $v", subscription)
	return
}
