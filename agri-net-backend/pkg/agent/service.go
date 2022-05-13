package agent

import (
	"context"

	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/model"
)

type IAgentService interface {
	RegisterAgent(ctx context.Context, agent *model.Agent) (int, error)
	GetAgentByID(ctx context.Context, id int) (*model.Agent, error)
	GetAgentsAddress(ctx context.Context, agent_id int) (*model.Address, error)
}

type AgentService struct {
	Repo IAgentRepo
}

func NewAgentService(repo IAgentRepo) IAgentService {
	return &AgentService{
		Repo: repo,
	}
}

// RegisterAgent ...
func (service *AgentService) RegisterAgent(ctx context.Context, agent *model.Agent) (int, error) {
	return service.Repo.RegisterAgent(ctx, agent)
}

func (service *AgentService) GetAgentByID(ctx context.Context, id int) (*model.Agent, error) {
	return service.Repo.GetAgentByID(ctx, id)
}

func (service *AgentService) GetAgentsAddress(ctx context.Context, agent_id int) (*model.Address, error) {
	return service.Repo.GetAgentsAddress(ctx, agent_id)
}
