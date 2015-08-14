package test

import (
	"log"

	base "github.com/nubleer/conekta-go/conekta"
)

type Customer struct {
	Id           string `db:"-" json:"id"`
	Name         string `db:"-" json:"name"`
	Email        string `db:"-" json:"email"`
	Card         *base.CreditCard
	Charge       *base.Charge
	Subscription *base.Subscription
	*base.Client
}

func NewCustomer(customerId string, cardToken string) *Customer {
	customer := &Customer{}
	customer.Client = base.NewClient()
	customer.Card = &base.CreditCard{}
	customer.Subscription = &base.Subscription{}

	customer.Id = customerId
	customer.Card.Token = cardToken
	return customer
}

func (self *Customer) CreateSubscription() (err error) {
	result, err := self.Customers.CreateSubscription(self.Id, self.Subscription)
	if err != nil {
		return
	}
	log.Printf("Resultado: $v", result)
	return
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
	log.Printf("Activando cuenta: %v", self.Id)
	subscription, err := self.Customers.ResumeSubscription(self.Id)
	if err != nil {
		return
	}
	log.Printf("Subscripción activada: $v", subscription)
	return
}

func (self *Customer) Cancel() (err error) {
	log.Printf("Cancelando cuenta: %v", self.Id)
	subscription, err := self.Customers.CancelSubscription(self.Id)
	if err != nil {
		return
	}
	log.Printf("Subscripción Cancelada: $v", subscription)
	return
}

func (self *Customer) UpdateSubscription(subs *base.Subscription) (err error) {
	log.Printf("Actualizando suscripción")
	var subscription *base.Subscription
	subscription, err = self.Customers.UpdateSubscription(self.Id, subs)
	if err != nil {
		return
	}
	log.Printf("Subscripción actualizada! : %v", subscription)
	return
}
