#!/bin/sh

HEADER_ACCEPT='Accept: application/vnd.conekta-v0.3.0+json'
HEADER_CONTENT_TYPE='Content-type: application/json'
API_KEY=':'
BASE_URL='https://api.conekta.io'
#BASE_URL='http://localhost:3000'


# bash => "{\n   \"name\":\"juan\",\n   \"email\":\"uno@masuno.com\"\n   }"

# ruby => "{\"name\":\"juan\",\"email\":\"whatever@com.com\"}"

# go => "{\"name\":\"Wasup\",\"email\":\"uno@masuno.com\",\"phone\":\"222-333-444\"}\n"

createCustomer() {
  REQUEST_BODY=`printf '{
   "name":"juan",
   "email":"uno@masuno.com"
   }'`

  curl -H "${HEADER_ACCEPT}" -H \
   "${HEADER_CONTENT_TYPE}" \
   -u ${API_KEY} \
   -X "POST" \
   -d "${REQUEST_BODY}" \
    ${BASE_URL}/customers
}

createCharge() {
  case $1 in
    "card")
        PAYMENT_TYPE='"card":"tok_test_visa_4242"'
        ;;
    "oxxo")
        PAYMENT_TYPE='"cash":{"type":"oxxo"}'
        ;;
    "deposit")
        PAYMENT_TYPE='"bank":{"type":"banorte"}'
        ;;
    *)
        PAYMENT_TYPE='"cash":{"type":"oxxo"}'
  esac

  REQUEST_BODY=`printf '{
   "description":"some useless shit",
   "amount":30000, %s
   }' ${PAYMENT_TYPE}`

  curl -H "${HEADER_ACCEPT}" -H \
   "${HEADER_CONTENT_TYPE}" \
   -u ${API_KEY} \
   -X "POST" \
   -d "${REQUEST_BODY}" \
    ${BASE_URL}/charges
}

createCustomer

