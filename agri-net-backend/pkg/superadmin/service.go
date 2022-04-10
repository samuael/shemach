package superadmin

type ISuperadminService interface{}

type SuperadminService struct {
	Repo ISuperadminRepo
}

func NewSuperadminService(repo ISuperadminRepo) ISuperadminService {
	return &SuperadminService{}
}
