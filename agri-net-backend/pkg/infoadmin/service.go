package infoadmin

type IInfoadminService interface {
	// --
}

type InfoadminService struct {
	Repo IInfoadminRepo
}

func NewInfoadminService(repo IInfoadminRepo) IInfoadminService {
	return &InfoadminService{
		Repo: repo,
	}
}
