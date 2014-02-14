/*
Package conekta implements a client that wraps the conekta API (v 0.3.0)

Create a new client to make requests to the resources exposed by the API.
For example, to create a charge to be paid in OXXO (more examples are included in
the examples directory):

	client := conekta.NewClient()
	charge := &conekta.Charge{
		Amount: 20000, //in cents
		Currency: "MXN",
		Description: "Some useless widgets",
		Cash: Oxxo
	}
    res, err := client.Charges.Create(charge)


Authenticating requests.

All requests to conekta must be authenticated. The client expects to find
the CONEKTA_API_KEY environment variable with your account's API key:

 	export CONEKTA_API_KEY=your_api_key

or, if you prefer:

 	os.Setenv("CONEKTA_API_KEY", your_api_key)

Handling responses

Handling errors

*/
package conekta

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"runtime"
	"strconv"
	"time"
)

const (
	baseURLString       = "https://api.conekta.io"
	apiVersion          = "0.3.0"
	gonektaVersion      = "0.1"
	userAgent           = "gonekta-" + gonektaVersion
	mimeType            = "application/vnd.conekta." + apiVersion + "+json"
	jsonMimeType        = "application/json"
	headerUserAgent     = "User-Agent"
	headerContentType   = "Content-Type"
	headerAccept        = "Accept"
	headerConektaClient = "X-Conekta-Client-User-Agent"
	envConektaAPIKey    = "CONEKTA_API_KEY"
)

type Client struct {
	client    *http.Client
	userAgent string
	BaseURL   *url.URL
	Charges   *chargesResource
	Customers *customersResource
	Plans     *plansResource
}

type ConektaError struct {
	Response *http.Response
	Type     string `json:"type"`
	Code     string `json:"code"`
	Param    string `json:"param"`
	Message  string `json:"message"`
}

type GonektaError struct {
	Message string `json:"message"`
}

type timestamp struct {
	time.Time
}

type Param map[string]interface{}

func (t timestamp) String() string {
	return t.Time.String()
}

func (ts *timestamp) UnmarshalJSON(b []byte) error {
	result, err := strconv.ParseInt(string(b), 10, 64)
	if err == nil {
		(*ts).Time = time.Unix(result, 0)
	} else {
		(*ts).Time, err = time.Parse(`"`+time.RFC3339+`"`, string(b))
	}
	return err
}

// NewClient returns a configured conekta client. All requests to the API
// go through this value.
func NewClient() *Client {
	baseUrl, _ := url.Parse(baseURLString)
	cli := &Client{
		client:    http.DefaultClient,
		BaseURL:   baseUrl,
		userAgent: userAgent,
	}
	cli.Charges = newChargesResource(cli)
	cli.Customers = newCustomersResource(cli)
	cli.Plans = newPlansResource(cli)

	return cli
}

func (c *Client) execute(method, path string, resBody, reqBody interface{}) error {
	req, err := c.prepareRequest(method, path, reqBody)
	if err != nil {
		return err
	}
	err = c.executeRequest(req, resBody)
	return err
}

func (c *Client) prepareRequest(method, path string, body interface{}) (*http.Request, error) {
	relative, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	baseUrl := c.BaseURL.ResolveReference(relative)
	buf := new(bytes.Buffer)

	if body != nil {
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, baseUrl.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add(headerContentType, jsonMimeType)
	req.Header.Add(headerAccept, mimeType)
	req.Header.Add(headerUserAgent, userAgent)
	req.Header.Add(headerConektaClient, func() string {
		j, _ := json.Marshal(map[string]string{
			"lang":         "Go",
			"lang_version": runtime.Version(),
			"uname":        runtime.GOOS,
		})
		return string(j)
	}())

	apiKey := os.Getenv(envConektaAPIKey)

	if len(apiKey) == 0 {
		return nil, GonektaError{"Missing CONEKTA_API_KEY"}
	}
	req.SetBasicAuth(apiKey, "")
	return req, nil
}

func (c *Client) executeRequest(req *http.Request, val interface{}) error {

	res, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	err = handleConektaError(res)

	if err != nil {
		return err
	}

	if val != nil {
		err = json.NewDecoder(res.Body).Decode(val)
	}
	return err
}

func (e *ConektaError) Error() string {
	return fmt.Sprintf("[%d] %s %s %s %s",
		e.Response.StatusCode,
		e.Type,
		e.Code,
		e.Param,
		e.Message,
	)
}

func (e GonektaError) Error() string {
	return e.Message
}

func handleConektaError(r *http.Response) error {
	if code := r.StatusCode; 200 <= code && code <= 299 {
		return nil
	}
	e := &ConektaError{Response: r}
	body, err := ioutil.ReadAll(r.Body)
	if err == nil && body != nil {
		json.Unmarshal(body, e)
	}
	return e
}
