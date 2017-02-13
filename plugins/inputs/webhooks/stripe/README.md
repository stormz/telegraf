# stripe webhooks

## Events

### [charge succeeded](https://stripe.com/docs/api#event_types-charge.succeeded)

**Tags:**
* 'event' = `event.type` string

**Fields:**
* 'id' = `event.data.object.id` string
* 'amount' = `event.data.object.amount`,
* 'currency' = `event.data.object.currency`

### [charge refunded](https://stripe.com/docs/api#event_types-charge.refunded)

**Tags:**
* 'event' = `event.type` string

**Fields:**
* 'id' = `event.data.object.id` string
* 'amount' = `event.data.object.amount`,
* 'currency' = `event.data.object.currency`

### [charge failed](https://stripe.com/docs/api#event_types-charge.failed)

**Tags:**
* 'event' = `event.type` string

**Fields:**
* 'id' = `event.data.object.id` string
* 'amount' = `event.data.object.amount`,
* 'currency' = `event.data.object.currency`

### [subscription created](https://stripe.com/docs/api/curl#event_types-customer.subscription.created)

### [subscription deleted](https://stripe.com/docs/api/curl#event_types-customer.subscription.deleted)

### [subscription updated](https://stripe.com/docs/api/curl#event_types-customer.subscription.updated)


### transfer.created
### transfer.failed
### transfer.paid
### transfer.reversed
