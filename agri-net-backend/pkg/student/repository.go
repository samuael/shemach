package student

import (
	"context"

	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/model"
)

type IStudentRepo interface {
	RegisterStudent(ctx context.Context) (*model.Student, int, error)
	GetStudentByID(ctx context.Context) (*model.Student, int, error)
	CheckWhetherTheStudentWithThisPhoneNumberExists(ctx context.Context) (int, error)
	UpdateStudent(ctx context.Context) (*model.Student, int, error)
	GetStudentsOfRound(ctx context.Context) ([]*model.Student, int, error)
	GetStudentImageUrl(context.Context) (string, uint8)
	ChangeStudentImageUrl(context.Context) error
	CreateSpecialCase(ctx context.Context) (*model.SpecialCase, int, error)
	GetSpecialCaseByID(ctx context.Context) (*model.SpecialCase, int, error)
	UpdateSpecialCase(ctx context.Context) (int, error)
	GetStudentsOfCategory(ctx context.Context) ([]*model.Student, int, error)
}
