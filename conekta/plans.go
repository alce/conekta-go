package conekta

import (
	"fmt"
)

type plansResource struct {
	client *Client
	path   string
}

type Plan struct {
	Id                 string    `json:"id,omitempty"`
	Object             string    `json:"object,omitempty"`
	Livemode           bool      `json:"livemode,omitempty"`
	CreatedAt          timestamp `json:"created_at,omitempty"`
	Name               string    `json:"name,omitempty"`
	Amount             uint      `json:"amount,omitempty"`
	Currency           string    `json:"currency,omitempty"`
	Interval           string    `json:"interval,omitempty"`
	Frequency          int       `json:"frequency,omitempty"`
	IntervalTotalCount int       `json:"interval_total_count,omitempty"`
	TrialPeriodDays    int       `json:"trial_period_days,omitempty"`
	ExpiryCount        int       `json:"expiry_count,omitempty"`
}

func newPlansResource(c *Client) *plansResource {
	return &plansResource{
		client: c,
		path:   "plans",
	}
}

func (s *plansResource) Create(plan *Plan) (*Plan, error) {
	in := new(Plan)
	err := s.client.execute("POST", s.path, in, plan)
	if err != nil {
		return nil, err
	}
	return in, nil
}

func (s *plansResource) Update(planId string, plan *Plan) (*Plan, error) {
	in := new(Plan)
	path := fmt.Sprintf("%s/%s", s.path, planId)
	err := s.client.execute("PUT", path, in, plan)
	if err != nil {
		return nil, err
	}
	return in, nil
}

func (s *plansResource) Delete(planId string) (*Plan, error) {
	in := new(Plan)
	path := fmt.Sprintf("%s/%s", s.path, planId)
	err := s.client.execute("DELETE", path, in, nil)
	if err != nil {
		return nil, err
	}
	return in, nil
}
