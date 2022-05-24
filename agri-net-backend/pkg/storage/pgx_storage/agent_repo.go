package pgx_storage

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/samuael/agri-net/agri-net-backend/pkg/agent"
	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/model"
)

type AgentRepo struct {
	DB *pgxpool.Pool
}

func NewAgentRepo(conn *pgxpool.Pool) agent.IAgentRepo {
	return &AgentRepo{
		DB: conn,
	}
}

func (repo *AgentRepo) RegisterAgent(ctx context.Context, agent *model.Agent) (int, error) {
	addressID := 0
	agentID := 0
	if agent.FieldAddress == nil {
		agent.FieldAddress = &model.Address{}
	}
	er := repo.DB.QueryRow(ctx, `select * from createAgent(cast($1 as varchar),cast($2 as varchar),cast( $3 as varchar),cast($4 as varchar),cast($5 as text) ,cast ($6 as char(3)) ,cast ($7 as int),
	cast ($8 as varchar),cast ($9 as varchar) ,cast ($10 as varchar) ,cast ($11 as varchar) ,cast( $12 as varchar), cast ($13 as varchar),cast ($14 as varchar),cast($15 as varchar))`,
		agent.Firstname, agent.Lastname, agent.Phone, agent.Email, agent.Password, agent.Lang, agent.RegisteredBy,
		agent.FieldAddress.Kebele, agent.FieldAddress.Woreda, agent.FieldAddress.City, agent.FieldAddress.Region, agent.FieldAddress.Zone,
		agent.FieldAddress.UniqueAddressName, fmt.Sprint(agent.FieldAddress.Latitude), fmt.Sprint(agent.FieldAddress.Longitude),
	).Scan((&agentID))
	if er != nil {
		if strings.Contains(er.Error(), "duplicate key value violates unique constraint") {
			return -4, er
		}
		log.Println(er.Error())
		return agentID /*addressID, */, er
	}
	agent.FieldAddress.ID = uint(addressID)
	era := repo.DB.QueryRow(ctx, "select address_id from admin where id=$1", agent.ID).Scan(&addressID)
	if era == nil {
		addressID = 0
	}
	if agentID < -1 {
		return agentID, errors.New("unauthorized")
	} else if agentID == -2 {
		return agentID, errors.New("unacceptable address information")
	} else if agentID == -3 {
		return agentID, errors.New("error while creating the agent instance")
	}
	agent.ID = uint64(agentID)
	if agent.FieldAddress != nil {
		agent.FieldAddress.ID = uint(addressID)
	}
	return int(agent.ID) /*int(agent.FieldAddress.ID),*/, nil
}

func (repo *AgentRepo) GetAgentByID(ctx context.Context, id int) (*model.Agent, error) {
	agent := &model.Agent{}
	er := repo.DB.QueryRow(ctx, `select  id,firstname,lastname,phone,email,imageurl,created_at,password,lang,posts_count,field_address_ref,registered_by from agent where id=$1`, id).
		Scan(&(agent.ID), &(agent.Firstname), &(agent.Lastname), &(agent.Phone), &(agent.Email), &(agent.Imgurl), &(agent.CreatedAt), &(agent.Password), &(agent.Lang), &(agent.PostsCount), &(agent.FieldAddressRef), &(agent.RegisteredBy))
	if er != nil {
		return nil, er
	}
	var address model.Address
	latitude := ""
	longitude := ""
	ers := repo.DB.QueryRow(ctx, `select address_id,kebele,woreda,city,region,unique_name,latitude,zone,longitude from address where address_id=$1`, agent.FieldAddressRef).
		Scan(&(address.ID), &(address.Kebele), &(address.Woreda), &(address.City), &(address.Region), &(address.UniqueAddressName), &(latitude), &(address.Zone), &(longitude))
	if ers == nil {
		address.Latitude, _ = strconv.ParseFloat(latitude, 64)
		address.Longitude, _ = strconv.ParseFloat(longitude, 64)
		agent.FieldAddress = &address
	}
	return agent, nil
}

func (repo *AgentRepo) GetAgentsAddress(ctx context.Context, agent_id int) (*model.Address, error) {
	addressid := 0
	er := repo.DB.QueryRow(ctx, `select field_address_ref from agent where id=$1`, agent_id).
		Scan(&(addressid))
	if er != nil {
		return nil, er
	}
	var address model.Address
	latitude := ""
	longitude := ""
	ers := repo.DB.QueryRow(ctx, `select address_id,kebele,woreda,city,region,unique_name,latitude,zone,longitude from address where address_id=$1`, addressid).
		Scan(&(address.ID), &(address.Kebele), &(address.Woreda), &(address.City), &(address.Region), &(address.UniqueAddressName), &(latitude), &(address.Zone), &(longitude))
	if ers == nil {
		address.Latitude, _ = strconv.ParseFloat(latitude, 64)
		address.Longitude, _ = strconv.ParseFloat(longitude, 64)
		return &address, nil

	}
	return nil, ers
}

// SearchAgent
func (repo *AgentRepo) SearchAgent(ctx context.Context, phone, name string, createdBy uint64, offset, limit uint) ([]*model.Agent, error) {
	agents := []*model.Agent{}
	values := []interface{}{}
	count := 1
	statement := `select id,firstname,lastname,phone,email,imageurl,created_at,password,lang,posts_count,field_address_ref,registered_by from agent where `
	if phone != "" {
		statement = fmt.Sprintf("%s phone ILIKE $%d", statement, count)
		values = append(values, "%"+strings.Trim(phone, " +")+"%")
		count++
	}
	name = strings.Trim(name, " ")
	if name != "" {
		if count > 1 {
			statement = fmt.Sprintf(" %s or ", statement)
		}
		statement = fmt.Sprintf("%s firstname ILIKE $"+strconv.Itoa(count), statement)
		values = append(values, "%"+name+"%")
		statement = fmt.Sprintf("%s or ", statement)
		count++
		statement = fmt.Sprintf("%s lastname ILIKE $"+strconv.Itoa(count), statement)
		values = append(values, "%"+name+"%")
		count++
	}
	if createdBy > 0 {
		if count > 1 {
			statement = fmt.Sprintf(" %s or ", statement)
		}
		statement = fmt.Sprintf(" %s registered_by=$"+strconv.Itoa(int(count)), statement)
		values = append(values, createdBy)
		count++
	}
	statement = fmt.Sprintf("%s ORDER BY id DESC OFFSET $%d LIMIT $%d ", statement, count, count+1)
	values = append(values, offset, limit)
	rows, er := repo.DB.Query(ctx, statement, values...)
	if er != nil {

		return nil, er
	}
	for rows.Next() {
		agent := &model.Agent{}
		erf := rows.Scan(
			&(agent.ID), &(agent.Firstname), &(agent.Lastname), &(agent.Phone), &(agent.Email), &(agent.Imgurl), &(agent.CreatedAt), &(agent.Password), &(agent.Lang),
			&(agent.PostsCount), &(agent.FieldAddressRef), &(agent.RegisteredBy))
		if erf != nil {
			println(erf.Error())
			continue
		}
		var address model.Address
		latitude := ""
		longitude := ""
		ers := repo.DB.QueryRow(ctx, `select address_id,kebele,woreda,city,region,unique_name,latitude,zone,longitude from address where address_id=$1`, agent.FieldAddressRef).
			Scan(&(address.ID), &(address.Kebele), &(address.Woreda), &(address.City), &(address.Region), &(address.UniqueAddressName), &(latitude), &(address.Zone), &(longitude))
		if ers == nil {
			address.Latitude, _ = strconv.ParseFloat(latitude, 64)
			address.Longitude, _ = strconv.ParseFloat(longitude, 64)
			agent.FieldAddress = &address
		}
		agents = append(agents, agent)
	}
	return agents, nil
}
