package pgx_storage

import (
	"context"
	"strconv"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/model"
	"github.com/samuael/agri-net/agri-net-backend/pkg/crop"
)

type CropRepo struct {
	DB *pgxpool.Pool
}

func NewCropRepo(conn *pgxpool.Pool) crop.ICropRepo {
	return &CropRepo{
		DB: conn,
	}
}

func (repo *CropRepo) CreateCrop(ctx context.Context, crop *model.Crop) (int, error) {
	result := 0
	er := repo.DB.QueryRow(ctx, `select * from createProductPost( $1,$2,$3,$4,$5,$6,$7,$8,$9)`,
		crop.TypeID, crop.Description, crop.Negotiable, crop.RemainingQuantity, crop.SellingPrice,
		crop.AddressRef, crop.StoreID, crop.AgentID, crop.StoreOwned,
	).Scan(&result)
	if er != nil || result <= 0 {
		if er != nil {
			println(er.Error())
		}
		return result, er
	}
	crop.ID = uint64(result)
	return result, nil
}

func (repo *CropRepo) GetPostByID(ctx context.Context, postid uint64) (*model.Crop, error) {
	post := &model.Crop{}
	newval := &struct {
		StoreID interface{}
		AgentID interface{}
	}{}
	er := repo.DB.QueryRow(ctx, `select crop_id,type_id,description,negotiable,remaining_quantity,selling_price,address_id,images,created_at,store_id,agent_id,store_owned,closed from crop where crop_id=$1`, postid).
		Scan(
			&(post.ID), &(post.TypeID), &(post.Description), &(post.Negotiable),
			&(post.RemainingQuantity), &(post.SellingPrice),
			&(post.AddressRef), &(post.Images), &(post.CreatedAt),
			&(newval.StoreID), &(newval.AgentID), &(post.StoreOwned),
			&(post.Closed),
		)
	if er != nil {
		return nil, er
	}
	if newval.StoreID != nil {
		post.StoreID = uint64((newval.StoreID).(int32))
	}
	if newval.AgentID != nil {
		post.AgentID = uint64((newval.AgentID).(int32))
	}
	var address model.Address
	latitude := ""
	longitude := ""
	ers := repo.DB.QueryRow(ctx, `select address_id,kebele,woreda,city,region,unique_name,latitude,zone,longitude from address where address_id=$1`, post.AddressRef).
		Scan(&(address.ID), &(address.Kebele), &(address.Woreda), &(address.City), &(address.Region), &(address.UniqueAddressName), &(latitude), &(address.Zone), &(longitude))
	if ers == nil {
		address.Latitude, _ = strconv.ParseFloat(latitude, 64)
		address.Longitude, _ = strconv.ParseFloat(longitude, 64)
		post.Address = &address
	}
	return post, nil
}

// SaveNewPostImages save the images of the crop
func (repo *CropRepo) SaveNewPostImages(ctx context.Context, postid uint64, images []int) error {
	er := repo.DB.QueryRow(ctx, "update crop set images=$1 where crop_id=$2 returning crop_id", images, postid).Scan(&postid)
	if er != nil {
		return er
	}
	return nil
}

func (repo *CropRepo) GetPosts(ctx context.Context, offset, limit uint) ([]*model.Crop, error) {
	posts := []*model.Crop{}
	rows, er := repo.DB.Query(ctx, "select crop_id,type_id,description,negotiable,remaining_quantity,selling_price,address_id,images,created_at,store_id,agent_id,store_owned,closed from crop LIMIT $1  OFFSET $2", limit, offset)
	if er != nil {
		return nil, er
	}

	for rows.Next() {
		post := &model.Crop{}
		newval := &struct {
			StoreID interface{}
			AgentID interface{}
		}{}
		er := rows.
			Scan(
				&(post.ID), &(post.TypeID), &(post.Description), &(post.Negotiable),
				&(post.RemainingQuantity), &(post.SellingPrice),
				&(post.AddressRef), &(post.Images), &(post.CreatedAt),
				&(newval.StoreID), &(newval.AgentID), &(post.StoreOwned), &(post.Closed),
			)
		if er != nil {
			continue
		}
		if newval.StoreID != nil {
			post.StoreID = uint64((newval.StoreID).(int32))
		}
		if newval.AgentID != nil {
			post.AgentID = uint64((newval.AgentID).(int32))
		}
		var address model.Address
		latitude := ""
		longitude := ""
		ers := repo.DB.QueryRow(ctx, `select address_id,kebele,woreda,city,region,unique_name,latitude,zone,longitude from address where address_id=$1`, post.AddressRef).
			Scan(&(address.ID), &(address.Kebele), &(address.Woreda), &(address.City), &(address.Region), &(address.UniqueAddressName), &(latitude), &(address.Zone), &(longitude))
		if ers == nil {
			address.Latitude, _ = strconv.ParseFloat(latitude, 64)
			address.Longitude, _ = strconv.ParseFloat(longitude, 64)
			post.Address = &address
		}
		posts = append(posts, post)
	}
	return posts, nil
}
