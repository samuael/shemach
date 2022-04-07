package student

import (
	"context"

	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/model"
)

type IStudentService interface {
	// RegisterStudent uses "student"  type *model.Student as a parameter
	// and returns *model.Student , integer (status code) , error
	RegisterStudent(context.Context) (*model.Student, int, error)
	// GetStudentByID uses "student_id" of type uint64 to return a student instance.
	GetStudentByID(ctx context.Context) (*model.Student, int, error)
	// uses "student_phone" type string to get a student infor using its phone number
	CheckWhetherTheStudentWithThisPhoneNumberExists(ctx context.Context) (int, error)
	// UpdateStudent uses "updated_student" type *model.Student instance to update the student instance.
	UpdateStudent(ctx context.Context) (*model.Student, int, error)
	// GetStudentsOfRound uses
	// "round_id" type uint ,
	//  offset type uint ,
	// limit type uint and returns a list of students if the query was succesful.
	GetStudentsOfRound(ctx context.Context) ([]*model.Student, int, error)
	// GetStudentsOfCategory uses
	// "category_id" type uint,
	//  offset type uint,
	// limit type uint and returns a list of students if the query was succesful.
	GetStudentsOfCategory(ctx context.Context) ([]*model.Student, int, error)

	// GetStudentImageUrl uses "student_id" type uint64
	// returns "" image url and statusCOde uint8
	GetStudentImageUrl(context.Context) (string, uint8)

	// ChangeStudentImageUrl uses "student_id" type uint64 and
	// "image_url" type string
	ChangeStudentImageUrl(context.Context) bool
	// CreateSpecialCase uses "special_case" type of *model.SpecialCase to create a special case instance
	// the Student ID is included in the special case model parameter with a Field name SpecialCase.
	CreateSpecialCase(ctx context.Context) (*model.SpecialCase, int, error)
	// GetSpecialCaseByID uses "special_case_id" of type uint64
	GetSpecialCaseByID(ctx context.Context) (*model.SpecialCase, int, error)
	// UpdateSpecialCase uses "special_case" of type *model.SpecialCase
	// to update the special case instance.
	// and returns an integer - representing status and error
	UpdateSpecialCase(ctx context.Context) (int, error)
}

type StudentService struct {
	Repo IStudentRepo
}

func NewStudentService(repo IStudentRepo) IStudentService {
	return &StudentService{
		Repo: repo,
	}
}

func (sservice *StudentService) RegisterStudent(ctx context.Context) (*model.Student, int, error) {
	return sservice.Repo.RegisterStudent(ctx)
}

func (sservice *StudentService) GetStudentByID(ctx context.Context) (*model.Student, int, error) {
	return sservice.Repo.GetStudentByID(ctx)
}

func (sservice *StudentService) CheckWhetherTheStudentWithThisPhoneNumberExists(ctx context.Context) (int, error) {
	return sservice.Repo.CheckWhetherTheStudentWithThisPhoneNumberExists(ctx)
}
func (sservice *StudentService) UpdateStudent(ctx context.Context) (*model.Student, int, error) {
	return sservice.Repo.UpdateStudent(ctx)
}
func (sservice *StudentService) GetStudentsOfRound(ctx context.Context) ([]*model.Student, int, error) {
	return sservice.Repo.GetStudentsOfRound(ctx)
}
func (sservice *StudentService) GetStudentImageUrl(ctx context.Context) (string, uint8) {
	return sservice.Repo.GetStudentImageUrl(ctx)
}

func (sservice *StudentService) ChangeStudentImageUrl(ctx context.Context) bool {
	return sservice.Repo.ChangeStudentImageUrl(ctx) == nil
}

func (sservice *StudentService) CreateSpecialCase(ctx context.Context) (*model.SpecialCase, int, error) {
	return sservice.Repo.CreateSpecialCase(ctx)
}
func (sservice *StudentService) GetSpecialCaseByID(ctx context.Context) (*model.SpecialCase, int, error) {
	return sservice.Repo.GetSpecialCaseByID(ctx)
}

func (sservice *StudentService) UpdateSpecialCase(ctx context.Context) (int, error) {
	return sservice.Repo.UpdateSpecialCase(ctx)
}
func (sservice *StudentService) GetStudentsOfCategory(ctx context.Context) ([]*model.Student, int, error) {
	return sservice.Repo.GetStudentsOfCategory(ctx)
}
