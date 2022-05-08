package rest

type ITransactionHandler interface {
	//
}

type TransactionHandler struct {
	//
}

func NewTransactionHandler() ITransactionHandler {
	return &TransactionHandler{}
}
