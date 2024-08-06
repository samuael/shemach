package pgx_storage

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/samuael/shemach/shemach-backend/pkg/constants/model"
	"github.com/samuael/shemach/shemach-backend/pkg/constants/state"
	"github.com/samuael/shemach/shemach-backend/pkg/product"
)

type ProductRepo struct {
	DB *pgxpool.Pool
}

func NewProductRepo(conn *pgxpool.Pool) product.IProductRepo {
	return &ProductRepo{
		DB: conn,
	}
}

func (repo *ProductRepo) CreateNewProduct(ctx context.Context) (*model.Product, int, error) {
	product := ctx.Value("product").(*model.Product)
	productID := 0
	if er := repo.DB.QueryRow(ctx, "select * from createProduct(cast ( $1 as varchar(200)) , cast ( $2 as varchar(200)),$7,cast ( $3 as decimal),$4,cast ($5 as integer), cast ( $6 as integer));",
		product.Name,
		product.ProductionArea,
		product.CurrentPrice,
		product.CreatedBy,
		product.CreatedAt,
		product.LastUpdateTime,
		product.UnitID,
	).Scan(&(productID)); er != nil {
		return product, state.STATUS_DBQUERY_ERROR, er
	}
	if productID == -1 {
		return product, state.STATUS_DUPLICATE_RECORD, errors.New("product with this id already exist")
	} else if productID == -2 {
		return product, state.STATUS_DBQUERY_ERROR, errors.New("can not create the product instance")
	} else if productID == -3 {
		return product, state.STATUS_RECORD_NOT_FOUND, errors.New("superadmin with the specified id does not exist")
	}
	product.ID = uint(productID)
	return product, state.STATUS_OK, nil
}

func (repo *ProductRepo) CheckTheExistanceOfProductInformation(ctx context.Context) bool {
	name := ctx.Value("name").(string)
	productionArea := ctx.Value("production_area").(string)
	unitID := ctx.Value("unit_id").(uint8)
	exists := false
	er := repo.DB.QueryRow(ctx, "select * from checkTheExistanceOfProduct( $1,$2,$3 )", name, productionArea, unitID).Scan(&exists)
	if er != nil {
		return false
	}

	return exists
}

func (repo *ProductRepo) GetProductByID(ctx context.Context) (*model.Product, int, error) {
	productID := ctx.Value("product_id").(uint8)
	product := &model.Product{}
	if er := repo.DB.QueryRow(ctx, `select 
	id,
	name,
	production_area,
	unit_id,
	current_price,
	created_by,
	created_at,
	last_updated_time from product where id=$1`, productID).Scan(
		&(product.ID),
		&(product.Name),
		&(product.ProductionArea),
		&(product.UnitID),
		&(product.CurrentPrice),
		&(product.CreatedBy),
		&(product.CreatedAt),
		&(product.LastUpdateTime),
	); er != nil {
		return product, state.STATUS_RECORD_NOT_FOUND, er
	}
	return product, state.STATUS_OK, nil
}

func (repo *ProductRepo) GetProducts(ctx context.Context) ([]*model.Product, int, error) {
	products := []*model.Product{}
	rows, er := repo.DB.Query(ctx, `select 
	id,
	name,
	production_area,
	unit_id,
	current_price,
	created_by,
	created_at,
	last_updated_time from product`)
	if rows == nil || er != nil {
		return products, state.STATUS_NO_RECORD_FOUND, errors.New("no record found")
	}
	for rows.Next() {
		product := &model.Product{}
		er := rows.Scan(
			&(product.ID),
			&(product.Name),
			&(product.ProductionArea),
			&(product.UnitID),
			&(product.CurrentPrice),
			&(product.CreatedBy),
			&(product.CreatedAt),
			&(product.LastUpdateTime),
		)
		if er != nil {
			continue
		}
		products = append(products, product)
	}
	if len(products) == 0 {
		return products, state.STATUS_NO_RECORD_FOUND, fmt.Errorf("no record was found")
	}
	return products, state.STATUS_OK, nil
}

func (repo *ProductRepo) CreateSubscriptions(ctx context.Context) (status int) {
	status = 0
	productID := ctx.Value("product_id").(uint8)
	subscriber := ctx.Value("subscriber_id").(uint64)
	if er := repo.DB.QueryRow(ctx, "select * from createNewSubscription( $1,$2::smallint )", subscriber, productID).Scan(&(status)); er != nil {
		return state.STATUS_DBQUERY_ERROR
	}
	return status
}

func (repo *ProductRepo) UnsubscribeProduct(ctx context.Context) (status int) {
	status = 0
	productID := ctx.Value("product_id").(uint8)
	subscriber := ctx.Value("subscriber_id").(uint64)
	if er := repo.DB.QueryRow(ctx, "select * from UnSubscribeToProduct( $1,$2::smallint )", subscriber, productID).Scan(&(status)); er != nil {
		return state.STATUS_DBQUERY_ERROR
	}
	return status
}

func (repo *ProductRepo) UpdateProductPrice(ctx context.Context) (int, int, error) {
	productID := ctx.Value("product_id").(uint8)
	price := ctx.Value("product_price").(float64)
	result := 0
	er := repo.DB.QueryRow(ctx, "select * from updateProductPrice($1, $2)", productID, price).Scan(&result)
	if er != nil {
		return -4, state.STATUS_DBQUERY_ERROR, er
	}
	if result == 0 {
		return result, state.STATUS_OK, nil
	} else if result == -1 {
		return result, state.STATUS_RECORD_NOT_FOUND, nil
	} else if result == -2 {
		return result, state.STATUS_NO_RECORD_UPDATED, nil
	} else {
		return result, state.STATUS_CONFLICT_ON_UPDATE, nil
	}
}

func (repo *ProductRepo) SearchProductsByText(ctx context.Context) ([]*model.Product, int, error) {
	text := ctx.Value("text").(string)
	products := []*model.Product{}
	rows, er := repo.DB.Query(ctx, `SELECT id,name,production_area,unit_id,current_price,created_by,created_at,last_updated_time 
	FROM product
	where name ILIKE $1 or production_area ILIKE $1`, "%"+text+"%")
	if er != nil {
		println(er.Error())
		return products, state.STATUS_DBQUERY_ERROR, er
	}
	for rows.Next() {
		product := &model.Product{}
		er := rows.Scan(&(product.ID), &(product.Name), &(product.ProductionArea), &(product.UnitID), &(product.CurrentPrice), &(product.CreatedBy), &(product.CreatedAt), &(product.LastUpdateTime))
		if er != nil {
			println(er.Error)
			continue
		}
		products = append(products, product)
	}
	if len(products) == 0 {
		return products, state.STATUS_NO_RECORD_FOUND, fmt.Errorf("no product instance was found")
	}
	return products, state.STATUS_OK, nil
}
