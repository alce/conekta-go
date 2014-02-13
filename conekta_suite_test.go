package conekta

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/codegangsta/martini"
	"github.com/martini-contrib/binding"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"
)

func TestGonekta(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gonekta Suite")
}

func createServer() martini.ClassicMartini {
	r := martini.NewRouter()
	m := martini.New()
	m.Action(r.Handle)
	return martini.ClassicMartini{m, r}
}

func NewTestServer() *httptest.Server {
	m := createServer()

	// Errors
	m.Get("/throw400", throw400)
	m.Get("/throw401", throw401)
	m.Get("/throw404", throw404)

	// Charges
	m.Post("/charges", binding.Json(Charge{}), createCharge)
	m.Get("/charges/:chargeId", getCharge)
	m.Post("/charges/:chargeId/refund", refundCharge)

	// Customers
	m.Post("/customers", binding.Json(Customer{}), createCustomer)
	m.Put("/customers/:customerId", binding.Json(Customer{}), updateCustomer)
	m.Delete("/customers/:customerId", deleteCustomer)
	m.Post("/customers/:customerId/cards", binding.Json(CreditCard{}), createCard)
	m.Put("/customers/:customerId/cards/:cardId", binding.Json(CreditCard{}), updateCard)
	m.Delete("/customers/:customerId/cards/:cardId", deleteCard)
	m.Post("/customers/:customerId/subscription", binding.Json(Subscription{}), createSubscription)
	m.Put("/customers/:customerId/subscription", binding.Json(Subscription{}), updateSubscription)
	m.Post("/customers/:customerId/subscription/pause", pauseSubscription)
	m.Post("/customers/:customerId/subscription/resume", resumeSubscription)
	m.Post("/customers/:customerId/subscription/cancel", cancelSubscription)

	// Plans
	m.Post("/plans", binding.Json(Plan{}), createPlan)
	m.Put("/plans/:planId", binding.Json(Plan{}), updatePlan)
	m.Delete("/plans/:planId", deletePlan)

	return httptest.NewServer(m)
}

func NewTestClient(s *httptest.Server) *Client {
	os.Setenv(envConektaAPIKey, "foo")
	c := NewClient()
	c.BaseURL, _ = url.Parse(s.URL)
	return c
}

func throw400(res http.ResponseWriter) {
	renderJSONFixture(res, 400, "400")
}

func throw401(res http.ResponseWriter) {
	renderJSONFixture(res, 401, "401")
}

func throw404(res http.ResponseWriter) {
	renderJSONFixture(res, 404, "404")
}

func createCharge(res http.ResponseWriter, charge Charge) {

	var fixturePrefix string

	switch {
	case charge.Cash != nil:
		fixturePrefix = "oxxo"
	case charge.Bank != nil:
		fixturePrefix = "bank"
	case len(charge.Card) > 0:
		fixturePrefix = "card"
	}
	renderJSONFixture(res, 200, fixturePrefix+"Charge")
}

func getCharge(res http.ResponseWriter, p martini.Params) {
	//fmt.Println(p["chargeId"])
	renderJSONFixture(res, 200, "bankCharge")
}

func refundCharge(res http.ResponseWriter, p martini.Params) {
	//fmt.Println(p["chargeId"])
	renderJSONFixture(res, 200, "refund")
}

func createCustomer(res http.ResponseWriter, customen Customer, req *http.Request) {
	//fmt.Println(customer)
	renderJSONFixture(res, 200, "customer")
}

func updateCustomer(res http.ResponseWriter, customer Customer, p martini.Params) {
	//fmt.Println(p["customerId"])
	renderJSONFixture(res, 200, "customer")
}

func deleteCustomer(res http.ResponseWriter, p martini.Params) {
	//fmt.Println(p["customerId"])
	renderJSONFixture(res, 200, "customer")
}

func createCard(res http.ResponseWriter, card CreditCard, p martini.Params) {
	//fmt.Println(p["customerId"])
	//fmt.Println(card.Token)
	renderJSONFixture(res, 200, "credit_card")
}

func updateCard(res http.ResponseWriter, card CreditCard, p martini.Params) {
	//fmt.Println(p["customerId"])
	//fmt.Println(p["cardId"])
	//fmt.Println(card.Active)
	renderJSONFixture(res, 200, "credit_card")
}

func deleteCard(res http.ResponseWriter, p martini.Params) {
	//fmt.Println(p["customerId"])
	//fmt.Println(p["cardId"])
	renderJSONFixture(res, 200, "credit_card")
}

func createPlan(res http.ResponseWriter, plan Plan) {
	// fmt.Println(plan)
	renderJSONFixture(res, 200, "plan")
}

func updatePlan(res http.ResponseWriter, plan Plan) {
	// fmt.Println(plan)
	renderJSONFixture(res, 200, "plan")
}

func deletePlan(res http.ResponseWriter, p martini.Params) {
	//fmt.println(p["planId"])
	renderJSONFixture(res, 200, "plan")
}

func createSubscription(res http.ResponseWriter, sub Subscription, p martini.Params) {
	//fmt.Println(sub)
	//fmt.Pristln(p["customerId"])
	renderJSONFixture(res, 200, "subscription")
}

func updateSubscription(res http.ResponseWriter, sub Subscription, p martini.Params) {
	//fmt.Println(sub)
	//fmt.Pristln(p["customerId"])
	renderJSONFixture(res, 200, "subscription")
}

func pauseSubscription(res http.ResponseWriter, p martini.Params) {
	//fmt.Println(p["customerId"])
	renderJSONFixture(res, 200, "subscription")
}

func resumeSubscription(res http.ResponseWriter, p martini.Params) {
	//fmt.Println(p["customerId"])
	renderJSONFixture(res, 200, "subscription")
}

func cancelSubscription(res http.ResponseWriter, p martini.Params) {
	//fmt.Println(p["customerId"])
	renderJSONFixture(res, 200, "subscription")
}

func renderJSONFixture(res http.ResponseWriter, status int, fixtureName string) {
	res.WriteHeader(status)
	res.Header().Set("Content-Type", "application/json; charset=UTF-8")
	res.Write([]byte(Fixtures[fixtureName]))
}

var Fixtures = map[string]string{
	"subscription": `{
		"id":"sub_EfhFCp5SKvp5XzXQk",
  		"status":"in_trial",
	  	"object":"subscription",
	  	"created_at":1385696776,
	  	"billing_cycle_start":1385696776,
	  	"billing_cycle_end":1386301576,
	  	"plan_id":"gold-plan",
	  	"card_id":"card_vow1u83899Rkj5LM"
	}`,
	"plan": `{
		"id":"gold-plan",
		"object":"plan",
		"livemode":false,
		"created_at":1385481591,
		"name":"Gold Plan",
		"amount":10000,
		"currency":"MXN",
		"interval":"month",
		"frequency":1,
		"interval_total_count":12,
		"trial_period_days":15
	}`,
	"credit_card": `{
		"id":"card_TCuBjUEcy9r41Fk2",
		"object":"card",
		"active":true,
		"brand":"VISA",
		"last4":"4242",
		"name":"James Howlett",
		"exp_month":"12",
		"exp_year":"2013",
		"address":{
			"street1":"250 Alexis St",
			"street2": null,
			"street3": null,
			"city":"Red Deer",
			"state":"Alberta",
			"zip":"T4N 0B8",
			"country":"Canada"
  		}
	}`,
	"customer": `{
		"id":"cus_k2D9DxlqdVTagmEd400001",
	  	"object":"customer",
	  	"livemode": false,
	  	"created_at": 1379784950,
	  	"name":"Thomas Logan",
	  	"email":"thomas.logan@xmen.org",
	  	"phone":"55-5555-5555",
	  	"default_card":"card_9kWcdlL7xbvQu5jd3",
	"billing_address": {
		"street1":"77 Mystery Lane",
		"street2":"Suite 124",
		"street3": null,
		"city":"Darlington",
		"state":"NJ",
		"zip":"10192",
		"country": null,
		"tax_id":"xmn671212drx",
		"company_name":"X-Men Inc.",
		"phone":"77-777-7777",
		"email":"purshasing@x-men.org"
	  },
	"shipping_address": {
		"street1":"250 Alexis St",
		"street2": null,
		"street3": null,
		"city":"Red Deer",
		"state":"Alberta",
		"zip":"T4N 0B8",
		"country":"Canada"
	  },
	 "cards": [{
	 	"id":"card_9kWcdlL7xbvQu5jd3",
  		"name":"Thomas Logan",
  		"last4":"4242",
  		"exp_month":"12",
  		"exp_year":"17",
  		"active":true
	 }],
	"subscription":{
		"id":"sub_ls9dklD9sAxW29dSmF",
		"card_id":"card_9kWcdlL7xbvQu5jd3",
		"plan_id":"gold-plan",
		"status":"active",
		"start":1379784950,
		"billing_cycle_start":1379784950,
		"billing_cycle_end":1379384950
  		}
	}`,
	"refund": ` {
	  	"id":"523df826aef8786485000001",
		  "livemode": false,
		  "created_at": 1379792934,
		  "status":"refunded",
		  "currency":"MXN",
		  "description":"Stogies",
		  "reference_id":"9839-wolf_pack",
		  "failure_code": null,
		  "failure_message": null,
		  "object":"charge",
		  "amount": 20000,
		  "amount_refunded":20000,
	  	"refunds":[{
			"created_at": 1379792934,
			"amount":20000,
			"currency": "MXN",
			"transaction": "5254d0f026c605054b0015ea"
		  }],
	  	"payment_method": {
			"object":"card_payment",
			"name":"Thomas Logan",
			"exp_month":"12",
			"exp_year":"15",
			"auth_code": "813038",
			"last4":"1111",
			"brand":"visa"
	  }
	}`,
	"bankCharge": `{
		"id":"52f8901cd7e1a0e1a20000c7",
		"livemode":false,
		"created_at":1392021532,
		"status":"pending_payment",
		"currency":"MXN",
		"description":"some useless shit",
		"reference_id":null,
		"failure_code":null,
		"failure_message":null,
		"monthly_installments":null,
		"object":"charge",
		"amount":30000,
		"fee":1670,
		"refunds":[],
		"payment_method": {
			"service_name":"Conekta",
			"service_number":"127589",
			"object":"bank_transfer_payment",
			"type":"banorte",
			"reference":"0068916"
		},
		"details":{
			"name":null,
			"phone":null,
			"email":null,
			"line_items":[]
		}
	}`,
	"oxxoCharge": `{
		"id": "52f88a63d7e1a0e1a20000b4",
		"livemode": false,
		"created_at": 1392020067,
		"status": "pending_payment",
		"currency": "MXN",
		"description": "some useless shit",
		"reference_id": null,
		"failure_code": null,
		"failure_message": null,
		"monthly_installments": null,
		"object": "charge",
		"amount": 30000,
		"fee": 1705,
		"refunds": [],
		"payment_method": {
			"expiry_date": "100314",
			"barcode": "38100000000042290121213001160013",
			"barcode_url": "https://www2.oxxo.com:8443/HTP/barcode/genbc?data=38100000000042290121213001160013&height=50&width=1&type=Code128",
			"object": "cash_payment",
			"type": "oxxo",
			"expires_at": 1394409600
		},
		"details": {
			"name": null,
			"phone": null,
			"email": null,
			"line_items": []
		}
	}`,
	"cardCharge": `{
		"id":"52f89639d7e1a09657000007",
		"livemode":false,
		"created_at":1392023097,
		"status":"paid",
		"currency":"MXN",
		"description":"some useless shit",
		"reference_id":null,
		"failure_code":null,
		"failure_message":null,
		"monthly_installments":null,
		"object":"charge",
		"amount":30000,
		"fee":1345,
		"refunds":[],
		"payment_method":{
			"name":"Jorge Lopez",
			"exp_month":"12",
			"exp_year":"19",
			"auth_code":"390678",
			"object":"card_payment",
			"last4":"4242",
			"brand":"visa"
		},
		"details":{
		"name": "wolverine",
		"phone":null,
		"email":null,
		"billing_address" : { "street2" : "Suite 124" },
		"line_items":[{"name": "Box of Cohiba 51s"}],
		"shipment" : {
			"carrier" : "estafeta",
			"address": {
				"country" : "Canada"
		   }
		}
	  }
	}`,
	"400": `{
    	"object":"error",
    	"type":"resource_not_found_url",
    	"message":"The requested resource could not be found"
	}`,
	"401": `{
    	"object":"error",
    	"type":"resource_not_found_url",
    	"message":"The requested resource could not be found"
	}`,
	"404": `{
    	"object":"error",
    	"type":"resource_not_found_url",
    	"message":"The requested resource could not be found"
	}`,
	"422": `{
		"object": "error",
		"type": "invalid_parameter_error",
		"code": "invalid_amount",
		"param": "amount",
		"message": "Invalid amount or incorrect format (must be an integer in cents)"
	}`,
}
