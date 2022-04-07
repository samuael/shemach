package pgx_storage

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/pgxpool"
	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/model"
	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/state"
	"github.com/samuael/agri-net/agri-net-backend/pkg/round"
)

type RoundRepo struct {
	DB *pgxpool.Pool
}

func NewRoundRepo(db *pgxpool.Pool) round.IRoundRepository {
	return &RoundRepo{
		DB: db,
	}
}

// GetCategoryByRoundNumberAndCategoryID ... returns
func (rrepo *RoundRepo) GetRoundByRoundNumberAndCategoryID(ctx context.Context) (*model.Round, int, error) {
	categoryID := ctx.Value("category_id").(uint64)
	roundNumber := ctx.Value("round_number").(uint64)
	round := &model.Round{}
	round.CreatedAt = &model.Date{}
	era := rrepo.DB.QueryRow(ctx, "SELECT id ,categoryid , training_hour , round_no , students , active_amount , active , lang, start_date , end_date, created_at, fee FROM round WHERE round_no=$1 AND categoryid=$2", roundNumber, categoryID).
		Scan(&(round.ID), &(round.CategoryID), &(round.TrainingHour), &(round.RoundNo), &(round.Students), &(round.ActiveAmount), &(round.Active), &(round.Lang), &(round.StartDate), &(round.EndDate), &(round.CreatedAtRef), &(round.Fee))
	if era != nil {
		return nil, state.DT_STATUS_DBQUERY_ERROR, era
	}
	if era := rrepo.DB.QueryRow(ctx, "SELECT id, year, month, day,hour, minute, second from eth_date where id=$1", round.CreatedAtRef).
		Scan(&(round.CreatedAt.ID), &(round.CreatedAt.Years), &(round.CreatedAt.Months),
			&(round.CreatedAt.Days), &(round.CreatedAt.Hours), &(round.CreatedAt.Minutes), &(round.CreatedAt.Seconds)); era != nil {
		return nil, state.DT_STATUS_INCOMPLETE_DATA, era
	}
	return round, state.DT_STATUS_OK, nil
}

// CreateRound ...
func (rrepo *RoundRepo) CreateRound(ctx context.Context) (*model.Round, int, error) {
	round := ctx.Value("round").(*model.Round)
	era := rrepo.DB.QueryRow(ctx,
		"select id, categoryid, training_hour, round_no, students, active_amount, active, start_date, lang, end_date,created_at,fee from createRound( $1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13, $14,$15,$16 )",
		round.CategoryID, round.TrainingHour, round.RoundNo, round.Students, round.ActiveAmount, round.Active, round.StartDate, round.Lang, round.EndDate, round.Fee, round.CreatedAt.Years, round.CreatedAt.Months, round.CreatedAt.Days, round.CreatedAt.Hours, round.CreatedAt.Minutes, round.CreatedAt.Seconds,
	).Scan(&(round.ID), &(round.CategoryID), &(round.TrainingHour), &(round.RoundNo), &(round.Students), &(round.ActiveAmount), &(round.Active), &(round.StartDate), &(round.Lang), &(round.EndDate), &(round.CreatedAtRef), &(round.Fee))
	if era != nil {
		println("database query error ")
		return nil, state.DT_STATUS_DBQUERY_ERROR, era
	}
	date := &model.Date{}
	println("The Created At Reference ID: ", round.CreatedAtRef)
	er := rrepo.DB.QueryRow(ctx, "SELECT id,year,month,day,hour,minute,second from eth_date where id=$1", round.CreatedAtRef).
		Scan(&(date.ID), &(date.Years), &(date.Months), &(date.Days), &(date.Hours), &(date.Minutes), &(date.Seconds))
	if er != nil {
		println("data is not complete")
		println(er.Error())
		return nil, state.DT_STATUS_INCOMPLETE_DATA, er
	}
	round.CreatedAt = date
	return round, state.DT_STATUS_OK, era
}

// DeleteRoundByID(ctx context.Context) error
func (rrepo *RoundRepo) DeleteRoundByID(ctx context.Context) error {
	roundID := ctx.Value("round_id").(uint64)
	if cmd, er := rrepo.DB.Exec(ctx, "DELETE FROM round WHERE id=$1", roundID); cmd.RowsAffected() == 0 || er != nil {
		return errors.New("No round instance is deleted")
	}
	return nil
}

// GetRoundByID(ctx context.Context) (*model.Round, error)
func (rrepo *RoundRepo) GetRoundByID(ctx context.Context) (*model.Round, error) {
	roundID := ctx.Value("round_id").(uint64)
	round := &model.Round{}
	er := rrepo.DB.QueryRow(ctx, " SELECT id ,categoryid , training_hour , round_no , students , active_amount , active , lang, start_date ,end_date, fee, created_at FROM round WHERE id=$1 ", roundID).
		Scan(&(round.ID), &(round.CategoryID), &(round.TrainingHour), &(round.RoundNo), &(round.Students), &(round.ActiveAmount), &(round.Active), &(round.Lang), &(round.StartDate), &(round.EndDate), &(round.Fee), &(round.CreatedAtRef))
	if er == nil {
		date := &model.Date{}
		er := rrepo.DB.QueryRow(ctx, "SELECT id,year,month,day,hour,minute,second from eth_date where id=$1", round.CreatedAtRef).
			Scan(&(date.ID), &(date.Years), &(date.Months), &(date.Days), &(date.Hours), &(date.Minutes), &(date.Seconds))
		if er != nil {
			return nil, er
		}
		round.CreatedAt = date
	}
	return round, er
}

func (rrepo *RoundRepo) UpdateRound(ctx context.Context) (int, error) {
	round := ctx.Value("round").(*model.Round)
	uc, er := rrepo.DB.Exec(ctx, "UPDATE round SET categoryid=$1,training_hour=$2 , round_no=$3, students=$4, active_amount=$5, active=$6, start_date=$7, lang=$8, end_date=$9, fee=$11 WHERE id=$10",
		round.CategoryID, round.TrainingHour, round.RoundNo, round.Students, round.ActiveAmount, round.Active, round.StartDate, round.Lang, round.EndDate, round.ID, round.Fee)
	if er != nil {
		return state.DT_STATUS_DBQUERY_ERROR, er
	} else if uc.RowsAffected() == 0 {
		return state.DT_STATUS_NO_RECORD_UPDATED, fmt.Errorf("%d rows affected", uc.RowsAffected())
	}
	return state.DT_STATUS_OK, nil
}

func (rrepo *RoundRepo) GetRoundsOfCategory(ctx context.Context) ([]*model.Round, int, error) {
	categoryID := ctx.Value("category_id").(uint)
	rounds := []*model.Round{}

	if cursor, er := rrepo.DB.Query(ctx,
		"SELECT id ,categoryid , training_hour , round_no , students , active_amount , active , lang, start_date,end_date,created_at,fee FROM round WHERE categoryid=$1 ", categoryID); er != nil {
		if cursor.Err() != nil {
			return rounds, state.DT_STATUS_DBQUERY_ERROR, fmt.Errorf(state.STATUS[state.DT_STATUS_DBQUERY_ERROR])
		} else {
			return rounds, state.DT_STATUS_NO_RECORD_FOUND, fmt.Errorf(" no record was found ")
		}
	} else {

		for cursor.Next() {
			round := &model.Round{}
			er := cursor.Scan(&(round.ID), &(round.CategoryID),
				&(round.TrainingHour), &(round.RoundNo), &(round.Students),
				&(round.ActiveAmount), &(round.Active), &(round.Lang),
				&(round.StartDate), &(round.EndDate), &(round.CreatedAtRef), &(round.Fee))
			if er != nil {
				continue
			}
			date := &model.Date{}
			if er := rrepo.DB.QueryRow(ctx, "SELECT id, year,month,day,hour,minute,second from eth_date where id=$1", round.CreatedAtRef).
				Scan(&(date.ID), &(date.Years), &(date.Months), &(date.Days), &(date.Hours), &(date.Minutes), &(date.Seconds)); er != nil {
				continue
			}
			round.CreatedAt = date
			rounds = append(rounds, round)
		}
		if len(rounds) == 0 {
			return rounds, state.DT_STATUS_NO_RECORD_FOUND, fmt.Errorf(state.STATUS[state.DT_STATUS_NO_RECORD_FOUND])
		}
		return rounds, state.DT_STATUS_OK, nil
	}
}

func (rrepo *RoundRepo) CountCategoryRounds(ctx context.Context) (int, error) {
	categoryID := ctx.Value("category_id").(uint64)
	counts := 0
	if er := rrepo.DB.QueryRow(ctx, "SELECT count(*) as count FROM round WHERE categoryid=$1", categoryID).Scan(&categoryID); er != nil {
		return counts, er
	}
	return counts, nil
}

func (rrepo *RoundRepo) CheckTheExistanceAndActivenessOfRound(ctx context.Context) (int, int) {
	roundid := ctx.Value("round_id").(uint64)
	var result int
	if er := rrepo.DB.QueryRow(ctx, "select * from checkExistanceAndActivenessOrRound($1)", roundid).Scan(&result); er != nil {
		println(er.Error())
		return -2, state.DT_STATUS_DBQUERY_ERROR
	}
	return result, state.DT_STATUS_OK
}

func (rrepo *RoundRepo) CheckTheExistanceOfCategory(ctx context.Context) (int, int) {
	categoryID := ctx.Value("category_id").(uint64)
	var result int
	if er := rrepo.DB.QueryRow(ctx, "select * from checkExistanceOfCategory($1)", categoryID).Scan(&result); er != nil {
		println(er.Error())
		return -1, state.DT_STATUS_DBQUERY_ERROR
	}
	return result, state.DT_STATUS_OK
}
