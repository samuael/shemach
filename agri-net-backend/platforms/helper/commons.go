package helper

import "github.com/samuael/agri-net/agri-net-backend/pkg/constants/state"

func RoleIntFromStringRole(role string) int8 {
	switch role {

	case state.SUPERADMIN:
		return state.SUPERADMIN_ROLE_INT
	case state.ADMIN:
		return state.ADMIN_ROLE_INT
	case state.INFO_ADMIN:
		return state.INFOADMIN_ROLE_INT
	case state.MERCHANT:
		return state.MERCHANT_ROLE_INT
	case state.AGENT:
		return state.AGENT_ROLE_INT
	default:
		return 0
	}
}
