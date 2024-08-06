package model

type Transaction struct {
	ID                uint    `json:"transaction_id"`
	RequestingPrice   float64 `json:"price"`
	Quantity          uint64  `json:"quantity"`
	State             uint    `json:"state"`
	Deadline          uint64  `json:"deadline_ts"`
	Description       string  `json:"description"`
	ProductID         uint64  `json:"product_id"`
	RequesterID       uint64  `json:"requester_id"`
	RequesterStoreRef uint64  `json:"requester_store_ref"`
	SellerID          uint64  `json:"seller_id"`
	SellerStoreRef    uint64  `json:"seller_store_ref"`
	CreatedAt         uint64  `json:"created_at"`
	// ----
	KebdAmount      float64 `json:"kebd_amount"`
	GuaranteeAmount float64 `json:"guarantee_amount"`
}

// TransactionState
type TransactionState struct {
	ID            uint64 `json:"id"`
	State         uint8  `json:"state"`
	CreatedAt     uint64 `json:"created_at"`
	TransactionID uint64 `json:"tr_id"`
}

// TransactionRequest
type TransactionRequest struct {
	TransactionState
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    float64 `json:"quantity"`
}

// KebdAmountRequest
type KebdAmountRequest struct {
	TransactionState
	KebdAmount  float64 `json:"kebd_amount"`
	Deadline    uint64  `json:"deadline"`
	Description string  `json:"description"`
}

// GuaranteeAmountRequest
type GuaranteeAmountRequest struct {
	TransactionState
	Description string  `json:"description"`
	Amount      float64 `json:"amount"`
}

type TransactionNotificationResponse struct {
	Msg           string                     `json:"msg"`
	StatusCode    int                        `json:"status_code"`
	Notifications []*TransactionNotification `json:"results"`
}

type TransactionNotification struct {
	TransactionID           uint64                  `json:"tr_id"`
	TransactionNotification *TransactionRequest     `json:"transaction_request"`
	KebdNotification        *KebdAmountRequest      `json:"kebd_request"`
	PaymentNotification     *TransactionPayment     `json:"payment_notification"`
	ContractNotification    *Contract               `json:"contract_notification"`
	GuaranteeNotification   *GuaranteeAmountRequest `json:"guarantee_request"`
}
