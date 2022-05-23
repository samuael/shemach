package transaction

type ITransactionService interface {
}

type TransactionService struct {
	Repo ITransactionRepo
}

func NewTransactionService(repo ITransactionRepo) ITransactionService {
	return &TransactionService{
		Repo: repo,
	}
}
