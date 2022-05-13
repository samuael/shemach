package pgx_storage

import (
	"context"
	"errors"
	"fmt"

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
