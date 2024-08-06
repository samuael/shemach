package pgx_storage

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/samuael/shemach/shemach-backend/pkg/constants/model"
	"github.com/samuael/shemach/shemach-backend/pkg/store"
)

type StoreRepo struct {
	DB *pgxpool.Pool
}

func NewStoreRepo(conn *pgxpool.Pool) store.IStoreRepo {
	return &StoreRepo{
		DB: conn,
	}
}

// CreateStore
func (repo *StoreRepo) CreateStore(ctx context.Context, store *model.Store) (int, error) {
	response := 0
	er := repo.DB.QueryRow(ctx, "select * from createStore( cast ( $1 as varchar(100)) , cast($2 as integer) , cast ($3 as integer),cast($4 as varchar(100)),cast( $5 as varchar(100)),cast($6 as varchar(100)),cast ($7  as varchar(100)), cast ($8 as varchar(100)), cast($9 as varchar(100)),$10,$11)",
		store.StoreName, store.OwnerID, store.CreatedBy, store.Address.Kebele, store.Address.Woreda, store.Address.City, store.Address.Region, store.Address.Zone, store.Address.UniqueAddressName, fmt.Sprint(store.Address.Latitude), fmt.Sprint(store.Address.Longitude),
	).Scan(&response)
	if er != nil {
		if strings.Contains(er.Error(), "value too long for type character varying") {
			return -6, er
		}
		return -5, er
	}
	if response == -4 {
		return response, errors.New("internal problem")
	} else if response == -3 {
		return response, errors.New("invalid merchant information")
	} else if response == -2 {
		return response, errors.New("invalid address information")
	} else if response == -1 {
		return response, errors.New("you are not authorized to take this action")
	}
	store.ID = uint64(response)
	return response, nil
}

// GetMerchantStores ...
func (repo *StoreRepo) GetMerchantStores(ctx context.Context, merchantID uint64) ([]*model.Store, error) {
	stores := []*model.Store{}
	println(merchantID)
	rows, er := repo.DB.Query(ctx, `select store_id,owner_id,address_id,active_products,store_name,active_contracts,created_at,created_by from store where owner_id=$1`, merchantID)
	if er != nil {
		return stores, er
	}
	for rows.Next() {
		var era error
		var store model.Store
		era = rows.Scan(&(store.ID), &(store.OwnerID), &(store.AddressRef), &(store.ActiveProducts), &(store.StoreName), &(store.ActiveContracts), &(store.CreatedAt), &(store.CreatedBy))
		if era != nil {
			continue
		}
		var address model.Address
		latitude := ""
		longitude := ""
		ers := repo.DB.QueryRow(ctx, `select address_id,kebele,woreda,city,region,unique_name,latitude,zone,longitude from address where address_id=$1`, store.AddressRef).
			Scan(&(address.ID), &(address.Kebele), &(address.Woreda), &(address.City), &(address.Region), &(address.UniqueAddressName), &(latitude), &(address.Zone), &(longitude))
		if ers == nil {
			address.Latitude, _ = strconv.ParseFloat(latitude, 64)
			address.Longitude, _ = strconv.ParseFloat(longitude, 64)
			store.Address = &address
		}
		stores = append(stores, &store)
	}
	return stores, nil
}

func (repo *StoreRepo) GetStoreByID(ctx context.Context, id uint64) (*model.Store, error) {
	store := &model.Store{}
	ers := repo.DB.QueryRow(ctx, `select store_id,owner_id,address_id,active_products,store_name,active_contracts,created_at,created_by from store where store_id=$1`, id).Scan(&(store.ID), &(store.OwnerID), &(store.AddressRef), &(store.ActiveProducts), &(store.StoreName), &(store.ActiveContracts), &(store.CreatedAt), &(store.CreatedBy))
	if ers != nil {
		return nil, nil
	}
	var address model.Address
	latitude := ""
	longitude := ""
	ers = repo.DB.QueryRow(ctx, `select address_id,kebele,woreda,city,region,unique_name,latitude,zone,longitude from address where address_id=$1`, store.AddressRef).
		Scan(&(address.ID), &(address.Kebele), &(address.Woreda), &(address.City), &(address.Region), &(address.UniqueAddressName), &(latitude), &(address.Zone), &(longitude))
	if ers == nil {
		address.Latitude, _ = strconv.ParseFloat(latitude, 64)
		address.Longitude, _ = strconv.ParseFloat(longitude, 64)
		store.Address = &address
	}
	return store, nil
}
