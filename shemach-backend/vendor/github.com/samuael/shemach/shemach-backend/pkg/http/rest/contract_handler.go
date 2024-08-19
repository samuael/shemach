package rest

type IContractHandler interface {
}

type ContractHandler struct {
}

func NewContractHandler() IContractHandler {
	return &ContractHandler{}
}
