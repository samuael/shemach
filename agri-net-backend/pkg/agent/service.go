package agent

import (
	"context"

	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/model"
)

type IAgentService interface {
	RegisterAgent(ctx context.Context, agent *model.Agent) (int, error)
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
