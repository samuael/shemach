package agent

import (
	"context"

	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/model"
)

type IAgentRepo interface {
	RegisterAgent(ctx context.Context, agent *model.Agent) (int, error)
	GetAgentByID(ctx context.Context, id int) (*model.Agent, error)
	GetAgentsAddress(ctx context.Context, agent_id int) (*model.Address, error)
	SearchAgent(ctx context.Context, phone, name string, createdBy uint64, offset, limit uint) ([]*model.Agent, error)
	DeleteAgentByID(ctx context.Context, agentid uint64) error
}
