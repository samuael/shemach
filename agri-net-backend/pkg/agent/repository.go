package agent

import (
	"context"

	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/model"
)

type IAgentRepo interface {
	RegisterAgent(ctx context.Context, agent *model.Agent) (int, error)
}
