package agent

import (
	"context"

	"github.com/samuael/shemach/shemach-backend/pkg/constants/model"
)

type IAgentService interface {
	RegisterAgent(ctx context.Context, agent *model.Agent) (int, error)
	GetAgentByID(ctx context.Context, id int) (*model.Agent, error)
	GetAgentsAddress(ctx context.Context, agent_id int) (*model.Address, error)
	SearchAgents(ctx context.Context, phone, name string, createdBy uint64, offset, limit uint) ([]*model.Agent, error)
	DeleteAgentByID(ctx context.Context, agentid uint64) error
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

// SearchAgent
func (service *AgentService) SearchAgents(ctx context.Context, phone, name string, createdBy uint64, offset, limit uint) ([]*model.Agent, error) {
	return service.Repo.SearchAgent(ctx, phone, name, createdBy, offset, limit)
}

func (service *AgentService) DeleteAgentByID(ctx context.Context, agentid uint64) error {
	return service.Repo.DeleteAgentByID(ctx, agentid)
}
