package stripe

import (
	"fmt"
)

func EventChargeJSON(event_type string) string {
	return fmt.Sprintf(`
        {
          "id": "evt_19lsbWI4amPefGV6qFUNJ1OA",
           "object": "event",
           "api_version": "2016-03-07",
           "created": 1486787722,
           "livemode": false,
           "pending_webhooks": 0,
           "request": null,
           "type": "%s",
           "data": {
             "object": {
               "id": "ch_19lEGwI4amPefGV6kCxtZyQC",
               "object": "charge",
               "amount": 1800,
               "amount_refunded": 0,
               "application": null,
               "application_fee": null,
               "balance_transaction": "txn_19kQn6I4amPefGV6TV8PUbOs",
               "captured": false,
               "created": 1486632686,
               "currency": "eur",
               "customer": "cus_8ejmezQ64UNFVb",
               "description": null,
               "destination": null,
               "dispute": null,
               "failure_code": "expired_card",
               "failure_message": "Your card has expired.",
               "fraud_details": {
               },
               "invoice": "in_19kqr0I4amPefGV6MiQS7620",
               "livemode": false,
               "metadata": {
               },
               "on_behalf_of": null,
               "order": null,
               "outcome": {
                 "network_status": "declined_by_network",
                 "reason": "expired_card",
                 "risk_level": "normal",
                 "seller_message": "The bank returned the decline code expired_card.",
                 "type": "issuer_declined"
               },
               "paid": false,
               "receipt_email": null,
               "receipt_number": null,
               "refunded": false,
               "refunds": {
                 "object": "list",
                 "data": [

                 ],
                 "has_more": false,
                 "total_count": 0,
                 "url": "/v1/charges/ch_19lEGwI4amPefGV6kCxtZyQC/refunds"
               },
               "review": null,
               "shipping": null,
               "source": {
                 "id": "card_18g99dI4amPefGV6vqBo6q0n",
                 "object": "card",
                 "address_city": "",
                 "address_country": "",
                 "address_line1": "",
                 "address_line1_check": "pass",
                 "address_line2": null,
                 "address_state": null,
                 "address_zip": "",
                 "address_zip_check": "pass",
                 "brand": "Visa",
                 "country": "US",
                 "customer": "cus_8ejmezQ64UNFVb",
                 "cvc_check": null,
                 "dynamic_last4": null,
                 "exp_month": 1,
                 "exp_year": 2017,
                 "funding": "credit",
                 "last4": "4242",
                 "metadata": {
                 },
                 "name": null,
                 "tokenization_method": null
               },
               "source_transfer": null,
               "statement_descriptor": "Standard",
               "status": "failed",
               "transfer_group": null
             }
           }
         }`, event_type)
}

func EventPingJSON() string {
	return `
        {
           "id": "evt_19lsbWI4amPefGV6qFUNJ1OA",
           "object": "event",
           "api_version": "2016-03-07",
           "created": 1486787722,
           "livemode": false,
           "pending_webhooks": 0,
           "request": null,
           "type": "ping"
        }`
}
