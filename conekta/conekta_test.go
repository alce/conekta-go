package conekta

import (
	"encoding/json"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

var _ = Describe("Conekta", func() {

	server := NewTestServer()
	client := NewTestClient(server)

	Context("Setup", func() {

		var c *Client

		BeforeEach(func() {
			c = NewClient()
			os.Setenv(envConektaAPIKey, "foo")
		})

		Describe("NewClient", func() {
			It("Should have the correct base URL", func() {
				Expect(c.BaseURL.String()).To(Equal(baseURLString))
			})
		})

		Describe("NewRequest", func() {

			var r *http.Request
			var err error

			BeforeEach(func() {
				r, err = c.prepareRequest("GET", "/wibble", nil)
			})

			It("Should return an error if the API key is not found", func() {
				os.Setenv(envConektaAPIKey, "")
				_, e := c.prepareRequest("GET", "/wibble", nil)
				Expect(e).To(HaveOccurred())
				Expect(e.Error()).To(Equal("Missing CONEKTA_API_KEY"))
			})

			It("Should set the correct User Agent", func() {
				Expect(r.Header.Get(headerUserAgent)).To(Equal(userAgent))
			})

			It("Should set the correct Accept header", func() {
				Expect(r.Header.Get(headerAccept)).To(Equal(mimeType))
			})

			It("Should set the absolute URL", func() {
				Expect(r.URL.String()).To(Equal(baseURLString + "/wibble"))
			})

			It("Should encode the body", func() {
				rawBody := map[string]string{"SomeKey": "SomeValue"}
				parsedBody := `{"SomeKey":"SomeValue"}` + "\n"
				r, _ = c.prepareRequest("POST", "/wibbles", rawBody)
				body, _ := ioutil.ReadAll(r.Body)
				Expect(string(body)).To(Equal(parsedBody))
			})

			It("Should return an error if the body is invalid JSON", func() {
				type T struct {
					A map[int]interface{}
				}
				_, err = c.prepareRequest("POST", "/wibbles", &T{})
				_, ok := err.(*json.UnsupportedTypeError)
				Expect(err).To(HaveOccurred())
				Expect(ok).To(BeTrue())
			})

			It("Should return an error if the URL is not valid", func() {
				_, err := c.prepareRequest("GET", ":", nil)
				_, ok := err.(*url.Error)
				Expect(err).To(HaveOccurred())
				Expect(ok).To(BeTrue())
			})
		})
	})

	Context("Making requests", func() {

		Describe("executeRequest", func() {
			It("Should return a not found error", func() {
				req, err := client.prepareRequest("GET", "/throw404", nil)
				err = client.executeRequest(req, nil)
				err, ok := err.(*ConektaError)
				Expect(err).To(HaveOccurred())
				Expect(ok).To(BeTrue())
			})

		})
	})
})
