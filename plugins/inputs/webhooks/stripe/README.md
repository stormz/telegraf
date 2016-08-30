# stripe webhooks

## Events

### [charge succeeded](https://stripe.com/docs/api#event_types-charge.succeeded)

**Tags:**
* 'event' = `event.type` string

**Fields:**
* 'id' = `event.id` string
* 'amount' = `event.data.object.amount`,

### [charge refunded](https://stripe.com/docs/api#event_types-charge.refunded)

### [charge failed](https://stripe.com/docs/api#event_types-charge.failed)
