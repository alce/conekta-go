package conekta

import (
	"fmt"
)

type Customer struct {
	Id              string          `json:"id,omitempty"`
	Object          string          `json:"customer,omitempty"`
	Livemode        bool            `json:"livemode,omitempty"`
	CreatedAt       *timestamp      `json:"created_at,omitempty"`
	Name            string          `json:"name,omitempty"`
	Email           string          `json:"email,omitempty"`
	Phone           string          `json:"phone,omitempty"`
	DefaultCard     string          `json:"default_card,omitempty"`
	BillingAddress  *BillingAddress `json:"billing_address,omitempty"`
	ShippingAddress *Address        `json:"shipping_address,omitempty"`
	Cards           []CreditCard    `json:"cards,omitempty"`
	Subscription    *Subscription   `json:"subscription,omitempty"`
}

type Subscription struct {
	Id                string     `json:"id,omitempty"`
	Status            string     `json:"status,omitempty"`
	Object            string     `json:"object,omitempty"`
	CreatedAt         *timestamp `json:"created_at,omitempty"`
	BillingCycleStart *timestamp `json:"billing_cycle_start,omitempty"`
	BillingCycleEnd   *timestamp `json:"billing_cycle_end,omitempty"`
	TrialStart        *timestamp `json:"trial_start,omitempty"`
	TrialEnd          *timestamp `json:"trial_end,omitempty"`
	PlanId            string     `json:"plan_id,omitempty"`
	CardId            string     `json:"card_id,omitempty"`
}

type CreditCard struct {
	Id       string   `json:"id,omitempty"`
	Object   string   `json:"object,omitempty"`
	Brand    string   `json:"brand,omitempty"`
	Name     string   `json:"name,omitempty"`
	Last4    string   `json:"last4,omitempty"`
	ExpMonth string   `json:"exp_month,omitempty"`
	ExpYear  string   `json:"exp_year,omitempty"`
	Active   bool     `json:"active,omitempty"`
	Token    string   `json:"token,omitempty"`
	Address  *Address `json:"address,omitempty"`
}

type customersResource struct {
	client *Client
	path   string
}

func newCustomersResource(c *Client) *customersResource {
	return &customersResource{
		client: c,
		path:   "customers",
	}
}

func (s *customersResource) Create(customer *Customer) (*Customer, error) {
	in := new(Customer)
	err := s.client.execute("POST", s.path, in, customer)
	if err != nil {
		return nil, err
	}
	return in, nil
}

func (s *customersResource) Update(id string, customer *Customer) (*Customer, error) {
	in := new(Customer)
	path := fmt.Sprintf("%s/%s", s.path, id)
	err := s.client.execute("PUT", path, in, customer)
	if err != nil {
		return nil, err
	}
	return in, nil
}

func (s *customersResource) Delete(id string) (*Customer, error) {
	in := new(Customer)
	path := fmt.Sprintf("%s/%s", s.path, id)
	err := s.client.execute("DELETE", path, in, nil)
	if err != nil {
		return nil, err
	}
	return in, nil
}

func (s *customersResource) AddCreditCard(customerId string, card *CreditCard) (*CreditCard, error) {
	in := new(CreditCard)
	path := fmt.Sprintf("%s/%s/cards", s.path, customerId)
	err := s.client.execute("POST", path, in, card)
	if err != nil {
		return nil, err
	}
	return in, nil
}

func (s *customersResource) UpdateCreditCard(customerId string, card *CreditCard) (*CreditCard, error) {
	in := new(CreditCard)
	path := fmt.Sprintf("%s/%s/cards/%s", s.path, customerId, card.Id)
	err := s.client.execute("PUT", path, in, card)
	if err != nil {
		return nil, err
	}
	return in, nil
}

func (s *customersResource) DeleteCreditCard(customerId string, cardId string) (*CreditCard, error) {
	in := new(CreditCard)
	path := fmt.Sprintf("%s/%s/cards/%s", s.path, customerId, cardId)
	err := s.client.execute("DELETE", path, in, nil)
	if err != nil {
		return nil, err
	}
	return in, nil
}

func (s *customersResource) CreateSubscription(customerId string, sub *Subscription) (*Subscription, error) {
	in := new(Subscription)
	path := fmt.Sprintf("%s/%s/subscription", s.path, customerId)
	err := s.client.execute("POST", path, in, sub)
	if err != nil {
		return nil, err
	}
	return in, nil
}

func (s *customersResource) UpdateSubscription(customerId string, sub *Subscription) (*Subscription, error) {
	in := new(Subscription)
	path := fmt.Sprintf("%s/%s/subscription", s.path, customerId)
	err := s.client.execute("PUT", path, in, sub)
	if err != nil {
		return nil, err
	}
	return in, nil
}

func (s *customersResource) PauseSubscription(customerId string) (*Subscription, error) {
	in := new(Subscription)
	path := fmt.Sprintf("%s/%s/subscription/pause", s.path, customerId)
	err := s.client.execute("POST", path, in, nil)
	if err != nil {
		return nil, err
	}
	return in, nil
}

func (s *customersResource) ResumeSubscription(customerId string) (*Subscription, error) {
	in := new(Subscription)
	path := fmt.Sprintf("%s/%s/subscription/resume", s.path, customerId)
	err := s.client.execute("POST", path, in, nil)
	if err != nil {
		return nil, err
	}
	return in, nil
}

func (s *customersResource) CancelSubscription(customerId string) (*Subscription, error) {
	in := new(Subscription)
	path := fmt.Sprintf("%s/%s/subscription/cancel", s.path, customerId)
	err := s.client.execute("POST", path, in, nil)
	if err != nil {
		return nil, err
	}
	return in, nil
}
