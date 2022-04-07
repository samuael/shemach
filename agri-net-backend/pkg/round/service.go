package round

import (
	"context"

	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/model"
)

type IRoundService interface {
	// GetRoundByRoundNumberAndCategoryID uses
	// 'category_id' and 'round_number'
	// to return a round using the round number and category id .
	GetRoundByRoundNumberAndCategoryID(ctx context.Context) (*model.Round, int, error)
	// CreateRound uses 'round' pointer to a round object to create a new round instance in the database.
	CreateRound(ctx context.Context) (*model.Round, int, error)
	// DeleteRoundByID uses "round_id" uint64 to delete a round
	DeleteRoundByID(ctx context.Context) error
	// GetRoundByID   users "round_id" uint64 to get round by id
	GetRoundByID(ctx context.Context) (*model.Round, error)
	// UpdateRound  uses "round" *model.Round to update a round instance.
	// this round instance must have an id
	UpdateRound(ctx context.Context) (int, error)
	// GetRoundsOfCategory  uses "category_id" uint instance to fetch all the rounds
	// of a category.
	GetRoundsOfCategory(context.Context) ([]*model.Round, int, error)
	// CheckTheExistanceAndActivenessOfRound uses "round_id" uint64
	// returns (  Methods Status Code , and Error Status Codes  )
	// return value
	// -2 : internal error
	// -1 : round doesn't exist
	//  0 : round is not active
	//  1 : round is active.
	CheckTheExistanceAndActivenessOfRound(ctx context.Context) (int, int)
	// CheckTheExistanceOfCategory users "category_id" type uint64 to
	//  returns int - the status code for representing the existance of the category
	//  -1 for internal problem , 0 for not found , and 1 for found.
	// int - to represent error type.
	CheckTheExistanceOfCategory(ctx context.Context) (int, int)
}

type RoundService struct {
	Repo IRoundRepository
}

func NewRoundService(repo IRoundRepository) IRoundService {
	return &RoundService{
		Repo: repo,
	}
}

func (rser *RoundService) GetRoundByRoundNumberAndCategoryID(ctx context.Context) (*model.Round, int, error) {
	return rser.Repo.GetRoundByRoundNumberAndCategoryID(ctx)
}

// CreateRound(ctx context.Context) (*model.Round, error)
func (rser *RoundService) CreateRound(ctx context.Context) (*model.Round, int, error) {
	return rser.Repo.CreateRound(ctx)
}
func (rser *RoundService) DeleteRoundByID(ctx context.Context) error {
	return rser.Repo.DeleteRoundByID(ctx)
}

// GetRoundByID(ctx context.Context) (*model.Round, error)
func (rser *RoundService) GetRoundByID(ctx context.Context) (*model.Round, error) {
	return rser.Repo.GetRoundByID(ctx)
}
func (rser *RoundService) UpdateRound(ctx context.Context) (int, error) {
	return rser.Repo.UpdateRound(ctx)
}

func (rser *RoundService) GetRoundsOfCategory(ctx context.Context) ([]*model.Round, int, error) {
	return rser.Repo.GetRoundsOfCategory(ctx)
}
func (rser *RoundService) CheckTheExistanceAndActivenessOfRound(ctx context.Context) (int, int) {
	return rser.Repo.CheckTheExistanceAndActivenessOfRound(ctx)
}
func (rser *RoundService) CheckTheExistanceOfCategory(ctx context.Context) (int, int) {
	return rser.Repo.CheckTheExistanceOfCategory(ctx)
}
