package state

const (
	TS_CREATED = iota + 1
	TS_AMENDMENT_REQUESTED
	TS_AMENDED
	TS_KEBD_REQUESTED
	TS_KEBD_REQUEST_AMENDMENT_REQUEST_SENT //NIMP
	TS_KEBD_AMENDED                        //NIMP
	TS_GUARANTEE_AMOUNT_REQUEST_SENT
	TS_GUARANTEE_AMOUNT_AMEND_REQUEST_SENT //NIMP
	TS_GUARANTEE_AMOUNT_AMENDED            //NIMP
	TS_SELLER_ACCEPTED
	TS_BUYER_ACCEPTED
	TS_ACCEPTED
	TS_DECLINED
	TS_PAYMENT_INSTANTIATED
	TS_SELLER_PAYMENT_COMPLETED
	TS_BUYER_PAYMENT_COMPLETED
	TS_ERROR
)
const ServiceFee = 2.00

var TransactionStateMaps = map[uint]string{
	TS_CREATED:                             "transaction_created",
	TS_AMENDMENT_REQUESTED:                 "transaction_amendment_requested",
	TS_AMENDED:                             "transaction_amended",
	TS_KEBD_REQUESTED:                      "transaction_kebd_requested",
	TS_KEBD_REQUEST_AMENDMENT_REQUEST_SENT: "transaction_kebd_amendment_request_sent",
	TS_KEBD_AMENDED:                        "transaction_kebd_amended",
	TS_GUARANTEE_AMOUNT_REQUEST_SENT:       "transaction_guarantee_amount_requested",
	TS_GUARANTEE_AMOUNT_AMEND_REQUEST_SENT: "transaction_gurantee_amount_amendment_request_sent",
	TS_GUARANTEE_AMOUNT_AMENDED:            "transaction_guarantee_amount_amended",
	TS_SELLER_ACCEPTED:                     "seller_accepted_the_transaction",
	TS_BUYER_ACCEPTED:                      "buyer_accepted_the_transaction",
	TS_ACCEPTED:                            "transaction_accepted",
	TS_DECLINED:                            "transaction_declined",
	TS_SELLER_PAYMENT_COMPLETED:            "seller payment completed",
	TS_BUYER_PAYMENT_COMPLETED:             "buyer payment completed",
}
