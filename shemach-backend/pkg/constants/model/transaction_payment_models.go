package model

import "time"

type InvoiceStatus string

const (
	Initializing = InvoiceStatus("INITIALIZING")
	PENDING      = InvoiceStatus("PENDING")
	EXPIRED      = InvoiceStatus("EXPIRED")
	PROCESSED    = InvoiceStatus("PROCESSED")
	CANCELED     = InvoiceStatus("CANCELED")
	FAILED       = InvoiceStatus("FAILED")
	NONE         = InvoiceStatus("")
)

type TransactionPayment struct {
	TransactionState
	SellerID           uint64  `json:"seller_id"`
	SellerInvoiceID    string  `json:"seller_invoice_id"`
	BuyerID            uint64  `json:"buyer_id"`
	BuyerInvoiceID     string  `json:"buyer_invoice_id"`
	KebdAmount         float64 `json:"kebd_amount"`
	GuaranteeAmount    float64 `json:"guarantee_amount"`
	KebdCompleted      bool    `json:"kebd_completed"`
	GuaranteeCompleted bool    `json:"guarantee_completed"`
	ServiceFee         float64 `json:"service_fee"`
}

// HelloCashInvoice
type HelloCashInvoice struct {
	ID           string      `json:"id"`
	Code         string      `json:"code"`
	Date         time.Time   `json:"date"`
	Expires      time.Time   `json:"expires"`
	From         string      `json:"from"`
	Fromname     interface{} `json:"fromname"`
	To           string      `json:"to"`
	Toname       string      `json:"toname"`
	Amount       int         `json:"amount"`
	Description  string      `json:"description"`
	Currency     string      `json:"currency"`
	Status       string      `json:"status"`
	Statusdetail string      `json:"statusdetail"`
}

// HellocashTransfferRequest
type HellocashTransfferRequest struct {
	Amount   int    `json:"amount"`
	To       string `json:"to"`
	Currency string `json:"currency"`
}

// HellocashTransffer
type HellocashTransffer struct {
	Amount        int         `json:"amount"`
	To            string      `json:"to"`
	Currency      string      `json:"currency"`
	From          string      `json:"from"`
	Fromname      string      `json:"fromname"`
	Toname        string      `json:"toname"`
	Fee           int         `json:"fee"`
	Date          time.Time   `json:"date"`
	Processdate   interface{} `json:"processdate"`
	Status        string      `json:"status"`
	Statusdetail  string      `json:"statusdetail"`
	Statuscomment interface{} `json:"statuscomment"`
}

// Credentials
type Credentials struct {
	Principal   string `json:"principal,omitempty"`
	Credentials string `json:"credentials,omitempty"`
	System      string `json:"system,omitempty"`
}

// HellocashInvoiceRequest
type HellocashInvoiceRequest struct {
	ID          string `json:"id,omitempty"`
	Amount      int    `json:"amount"`
	Description string `json:"description"`
	From        string `json:"from"`
	Currency    string `json:"currency"`
}
