# Conekta

Package conekta is a wrapper for the [conekta API](https://www.conekta.io/docs/api)

## Important!!

This package is considered *alpha* and the public API will still change slightly before it's considered stable.

## Getting Started

First, get your account's private [API key](https://admin.conekta.io/#developers.keys). This package will need it in order to authenticate your requests.

    export CONEKTA_API_KEY=YOUR_PRIVATE_KEY

Get the package

    go get github.com/Boletia/conekta-go/conekta

## Usage

~~~ go
package main

import "github.com/Boletia/conekta"

func main() {
  client := conekta.NewClient()

  charge := conekta.Charge{
    Description: "Some description",
    Amount: 45000,
    Cash: PaymentOxxo{"type":"oxxo"},
  }

  charge, err = client.Charges.Create(charge)
~~~
