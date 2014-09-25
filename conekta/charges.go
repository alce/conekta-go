package conekta

import (
	"fmt"
)

type chargesResource struct {
	client *Client
	path   string
}

type Charge struct {
	Description         string         `json:"description"`
	Amount              int            `json:"amount"`
	Capture							*bool					 `json:"capture,omitempty"`	
	Currency            string         `json:"currency"`
	Card                string         `json:"card,omitempty"`
	MonthlyInstallments int            `json:"monthly_installments,omitempty"`
	ReferenceId         string         `json:"reference_id,omitempty"`
	Cash                CashPayment    `json:"cash",omitempty"`
	Bank                BankPayment    `json:"bank,omitempty"`
	Id                  *string        `json:"id,omitempty"`
	Livemode            *bool          `json:"livemode,omitempty"`
	CreatedAt           *timestamp     `json:"created_at,omitempty"`
	Status              *string        `json:"status,omitempty"`
	FailureCode         *string        `json:"failure_code,omitempty"`
	FailureMessage      *string        `json:"failure_message,omitempty"`
	Object              *string        `json:"object,omitempty"`
	AmountRefunded      *int           `json:"amount_refunded,omitempty"`
	Fee                 *int           `json:"fee,omitempty"`
	PaymentMethod       *PaymentMethod `json:"payment_method,omitempty"`
	Details             *Details       `json:"details,omitempty"`
	Refunds             []Refund       `json:"refunds,omitempty"`
	Customer            *Customer      `json:"customer,omitempty"`
}

type Refund struct {
	CreatedAt   *timestamp `json:"created_at,omitempty"`
	Amount      int        `json:"amount,omitempty"`
	Currency    string     `json:"currency,omitempty"`
	Transaction string     `json:"transaction,omitempty"`
}

type Shipment struct {
	Carrier    string   `json:"carrier,omitempty"`
	Service    string   `json:"service,omitempty"`
	TrackingId string   `json:"tracking_id,omitempty"`
	Price      int      `json:"price,omitempty"`
	Address    *Address `json:"address,omitempty"`
}

type LineItem struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Sku         string `json:"sku,omitempty"`
	UnitPrice   int    `json:"unit_price,omitempty"`
	Quantity    int    `json:"quantity,omitempty"`
	Type        string `json:"type,omiempty"`
}

type Details struct {
	Name           string          `json:"name,omitempty"`
	Phone          string          `json:"phone,omitempty"`
	Email          string          `json:"email,omitempty"`
	DateOfBirth    string          `json:"date_of_birth,omitempty"`
	LineItems      []LineItem      `json:"line_items,omitempty"`
	BillingAddress *BillingAddress `json:"billing_address,omitempty"`
	Shipment       *Shipment       `json:"shipment,omitempty"`
}

type PaymentMethod struct {
	Object        string     `json:"object,omitempty"`
	Type          string     `json:"type,omitempty"`
	ServiceName   string     `json:"service_name,omitempty"` //Bank
	ServiceNumber string     `json:"service_number,omitempty"`
	Reference     string     `json:"reference,omitempty"`
	ExpiryDate    string     `json:"expiry_date,omitempty"` //Oxxo
	Barcode       string     `json:"barcode,omitempty"`
	BarcodeUrl    string     `json:"barcode_url,omitempty"`
	ExpiresAt     *timestamp `json:"expires_at,omitempty"`
	Brand         string     `json:"brand,omitempty"` // CC
	AuthCode      string     `json:"auth_code,omitempty"`
	Last4         string     `json:"last4,omitempty"`
	ExpMonth      string     `json:"exp_month,omitempty"`
	ExpYear       string     `json:"exp_year,omitempty"`
	Name          string     `json:"name,omitempty"`
	Address       *Address   `json:"address,omitempty"`
}

type CashPayment map[string]string
type BankPayment map[string]string

func newChargesResource(c *Client) *chargesResource {
	return &chargesResource{
		client: c,
		path:   "charges",
	}
}

func (s *chargesResource) Create(out *Charge) (*Charge, error) {
	in := new(Charge)
	err := s.client.execute("POST", s.path, in, out)
	if err != nil {
		return nil, err
	}
	return in, nil
}

func (s *chargesResource) Get(id string) (*Charge, error) {
	path := fmt.Sprintf("%s/%s", s.path, id)
	in := new(Charge)
	err := s.client.execute("GET", path, in, nil)
	if err != nil {
		return nil, err
	}
	return in, nil
}

// If amount is  empty  a full refund will be issued
func (s *chargesResource) Refund(id string, amount Param) (*Charge, error) {
	path := fmt.Sprintf("%s/%s/refund", s.path, id)
	in := new(Charge)
	err := s.client.execute("POST", path, in, amount)
	if err != nil {
		return nil, err
	}
	return in, nil
}

func (s *chargesResource) Filter(options FilterOptions) ([]Charge, error) {
	req, err := s.client.prepareRequest("GET", s.path, nil)
	if err != nil {
		return nil, err
	}

	var r []Charge
	err = s.client.executeRequest(req, &r)
	if err != nil {
		return nil, err
	}
	return r, nil
}
