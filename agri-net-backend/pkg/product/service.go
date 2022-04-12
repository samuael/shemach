package product

type IProductService interface {
}

type ProductService struct {
	Repo IProductRepo
}

func NewProductService(repo IProductRepo) IProductService {
	return &ProductService{
		Repo: repo,
	}
}
