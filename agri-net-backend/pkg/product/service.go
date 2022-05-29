package product

import (
	"context"
	"strings"

	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/model"
)

type IProductService interface {
	// CreateNewProduct uses "product" of type *model.Product to represent a product instance and return the production informatino with the ID
	// the integer status code respose representa the status code which is returned by the database.
	CreateNewProduct(ctx context.Context) (*model.Product, int, error)
	// CheckTheExistanceOfProductInformation uses "name" type string
	// "production_area" type string
	// "unit_id" type uint8
	// to check and return a boolean value representing the existance of a product.
	CheckTheExistanceOfProductInformation(ctx context.Context) bool
	// GetProductByID uses "product_id" type uint8
	GetProductByID(ctx context.Context) (*model.Product, int, error)
	// GetProducts
	GetProducts(ctx context.Context) ([]*model.Product, int, error)
	// GetProductInfoByID
	GetProductInfoByID(productID int) (string, string)
	SearchProduct(text string) []map[int]map[string]string
	// CreateSubscriptions uses "product_id" type uint8 and "subscriber_id" of type uint64
	CreateSubscriptions(ctx context.Context) (status int)
	// UnsubscribeProduct  uses "product_id" type uint8 and "subscriber_id" of type uint64
	UnsubscribeProduct(ctx context.Context) (status int)
	// UpdateProductPrice uses "product_id" of type uint8 and "product_price" of type float64
	UpdateProductPrice(ctx context.Context) (int, int, error)
	// SearchProductsByText uses "text" string to return list of products, status code int , and error
	SearchProductsByText(ctx context.Context) ([]*model.Product, int, error)
	GetProductUnits() map[string]map[int]map[string]string
}

type ProductService struct {
	Repo IProductRepo
}

func NewProductService(repo IProductRepo) IProductService {
	return &ProductService{
		Repo: repo,
	}
}

func (service *ProductService) GetProductUnits() map[string]map[int]map[string]string {
	return ProductTypes
}

var ProductTypes = map[string]map[int]map[string]string{
	"mass": {
		1: {"K": "killo Gram"},
		2: {"g": "gram"},
		3: {"Kun": "kuntal"},
		4: {"Ton": "ton"},
	},
	"volume": {
		5: {"L": "litter"},
		6: {"M3": "meter cube"},
		7: {"Gal": "gallon"},
	},
	"item": {
		8:  {"SIT": "single item"},
		9:  {"DZ": "dozen"},
		10: {"HDZ": "half dozen"},
		11: {"QDZ": "quarter dozen"},
	},
	"size": {
		12: {"SM": "small"},
		13: {"LG": "large"},
		14: {"MD": "medium"},
	},
	"length": {
		14: {"M": "meter"},
		15: {"KM": "killo meter"},
		16: {"cm": "centi meter"},
		17: {"Mil": "mile"},
	},
}

func (pser *ProductService) GetProductInfoByID(productID int) (string, string) {
	for _, product := range ProductTypes {
		for key, pdct := range product {
			if key == productID {
				for key, val := range pdct {
					return key, val
				}
			}
		}
	}
	return "", ""
}

func (pser *ProductService) SearchProduct(text string) []map[int]map[string]string {
	result := []map[int]map[string]string{}
	for mainTitle, product := range ProductTypes {
		if strings.Contains(mainTitle, text) {
			result = append(result, product)
			continue
		}
		for dkey, pdct := range product {
			for keyl, val := range pdct {
				if strings.Contains(keyl, text) || strings.Contains(val, text) {
					result = append(result, map[int]map[string]string{dkey: {keyl: val}})
				}
			}
		}
	}
	return result
}
func (pser *ProductService) CreateNewProduct(ctx context.Context) (*model.Product, int, error) {
	return pser.Repo.CreateNewProduct(ctx)
}

func (pser *ProductService) CheckTheExistanceOfProductInformation(ctx context.Context) bool {
	return pser.Repo.CheckTheExistanceOfProductInformation(ctx)
}
func (pser *ProductService) GetProductByID(ctx context.Context) (*model.Product, int, error) {
	return pser.Repo.GetProductByID(ctx)
}

func (pser *ProductService) GetProducts(ctx context.Context) ([]*model.Product, int, error) {
	return pser.Repo.GetProducts(ctx)
}

func (pser *ProductService) CreateSubscriptions(ctx context.Context) (status int) {
	return pser.Repo.CreateSubscriptions(ctx)
}

func (pser *ProductService) UnsubscribeProduct(ctx context.Context) (status int) {
	return pser.Repo.UnsubscribeProduct(ctx)
}

func (pser *ProductService) UpdateProductPrice(ctx context.Context) (int, int, error) {
	return pser.Repo.UpdateProductPrice(ctx)
}
func (pser *ProductService) SearchProductsByText(ctx context.Context) ([]*model.Product, int, error) {
	return pser.Repo.SearchProductsByText(ctx)
}
