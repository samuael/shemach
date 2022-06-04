package contract

type IContractService interface{}

type ContractService struct {
	Repo IContractRepo
}

func NewContractService(repo IContractRepo) IContractService {
	return &ContractService{
		Repo: repo,
	}
}
