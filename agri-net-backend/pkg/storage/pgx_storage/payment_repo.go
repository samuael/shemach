package pgx_storage

import (
	"context"

	"github.com/jackc/pgx/pgxpool"
	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/model"
	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/state"
	"github.com/samuael/agri-net/agri-net-backend/pkg/payment"
)

type PaymentRepo struct {
	DB *pgxpool.Pool
}

func NewPaymentRepo(db *pgxpool.Pool) payment.IPaymentRepo {
	return &PaymentRepo{
		DB: db,
	}
}

func (payrepo *PaymentRepo) CheckTheExistanceOfPaymentInstance(ctx context.Context) int {
	payinInput := ctx.Value("payin_input").(*model.PayinInput)
	status := 0
	if er := payrepo.DB.QueryRow(ctx, `select * from checkTheExsistanceOfPayment($1,$2,$3,$4,$5,$6,$7,$8,$9)`, payinInput.UChars, payinInput.PayedBy, payinInput.RoundID, payinInput.CreatedAt.Years, payinInput.CreatedAt.Months, payinInput.CreatedAt.Days, payinInput.CreatedAt.Hours, payinInput.CreatedAt.Minutes, payinInput.CreatedAt.Seconds).Scan(&status); er != nil {
		return -4
	}
	return status
}

func (payrepo *PaymentRepo) GetPaymentUsingPayinInput(ctx context.Context) (*model.PayIn, int, error) {
	payment := &model.PayIn{}
	date := &model.Date{}
	payinInput := ctx.Value("payin_input").(*model.PayinInput)

	if er := payrepo.DB.QueryRow(ctx, `select id,amount,recieved_by,payed_by,created_at,roundid,status,uchars from GetPayinInstanceUsingItsInformation($1,$2,$3,$4,$5,$6,$7,$8)`,
		payinInput.UChars, payinInput.PayedBy, payinInput.CreatedAt.Years, payinInput.CreatedAt.Months, payinInput.CreatedAt.Days, payinInput.CreatedAt.Hours, payinInput.CreatedAt.Minutes, payinInput.CreatedAt.Seconds).
		Scan(&(payment.ID), &(payment.Amount), &(payment.RecievedBy), &(payment.StudentID), &(payment.CreatedAtRef), &(payment.RoundID), &(payment.Status), &(payment.UChars)); er != nil {
		println(er.Error())
		return nil, state.DT_STATUS_RECORD_NOT_FOUND, er
	}
	if er := payrepo.DB.QueryRow(ctx, "SELECT id, year,month,day,hour,minute,second from eth_date where id=$1", payment.CreatedAtRef).
		Scan(&(date.ID), &(date.Years), &(date.Months), &(date.Days), &(date.Hours), &(date.Minutes), &(date.Seconds)); er != nil {
		return nil, state.DT_STATUS_INCOMPLETE_DATA, er
	}
	payment.CreatedAt = date
	return payment, state.DT_STATUS_OK, nil
}

func (payrepo *PaymentRepo) CreatePayinInstance(ctx context.Context) (*model.PayIn, int, error) {
	payinInput := ctx.Value("payin_input").(*model.PayinInput)
	payment := &model.PayIn{}
	if er := payrepo.DB.QueryRow(ctx, `SELECT id,amount,recieved_by,payed_by,created_at,roundid,status,uchars 
		from createPayinTransaction( cast( $1 as char(2)),$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)`,
		payinInput.UChars, payinInput.PayedBy, payinInput.ReceivedBy, payinInput.RoundID, payinInput.CreatedAt.Years, payinInput.CreatedAt.Months, payinInput.CreatedAt.Days, payinInput.CreatedAt.Hours, payinInput.CreatedAt.Minutes, payinInput.CreatedAt.Seconds, payinInput.Amount,
	).
		Scan(&(payment.ID), &(payment.Amount), &(payment.RecievedBy), &(payment.StudentID), &(payment.CreatedAtRef), &(payment.RoundID), &(payment.Status), &(payment.UChars)); er != nil {
		println(er.Error())
		if er.Error() == "ERROR: value too long for type character(2) (SQLSTATE 22001)" {
			return nil, 999, er
		}
		return nil, state.DT_STATUS_DBQUERY_ERROR, er
	}
	date := &model.Date{}
	if er := payrepo.DB.QueryRow(ctx, "SELECT id, year,month,day,hour,minute,second from eth_date where id=$1", payment.CreatedAtRef).
		Scan(&(date.ID), &(date.Years), &(date.Months), &(date.Days), &(date.Hours), &(date.Minutes), &(date.Seconds)); er != nil {
		return nil, state.DT_STATUS_INCOMPLETE_DATA, er
	}
	payment.CreatedAt = date
	return payment, state.DT_STATUS_OK, nil
}

func (payrepo *PaymentRepo) GetRemainingPaymentOfStudentForRound(ctx context.Context) (float64, int, error) {
	studentID := ctx.Value("student_id").(uint64)
	paymentRemaining := 0.0
	if er := payrepo.DB.QueryRow(ctx, "select * from getRemainingPaymentOfStudentForRound($1)", studentID).Scan(&paymentRemaining); er != nil {
		return paymentRemaining, state.DT_STATUS_DBQUERY_ERROR, er
	}
	return paymentRemaining, state.DT_STATUS_OK, nil
}

func (payrepo *PaymentRepo) GetPaymentsOfAStudent(ctx context.Context) ([]*model.PayIn, int, error) {
	studentID := ctx.Value("student_id").(uint64)
	payments := []*model.PayIn{}
	rows, ers := payrepo.DB.Query(ctx, "select id, amount, recieved_by, payed_by, created_at, roundid, status, uchars from payin where payed_by=$1", studentID)
	if ers != nil {
		return payments, state.DT_STATUS_RECORD_NOT_FOUND, ers
	}
	for rows.Next() {
		payment := &model.PayIn{}
		er := rows.Scan(&(payment.ID), &(payment.Amount), &(payment.RecievedBy), &(payment.StudentID), &(payment.CreatedAtRef), &(payment.RoundID), &(payment.Status), &(payment.UChars))
		if er != nil {
			continue
		}
		date := &model.Date{}
		if erf := payrepo.DB.QueryRow(ctx, "SELECT id, year,month,day,hour,minute,second from eth_date where id=$1", payment.CreatedAtRef).
			Scan(&(date.ID), &(date.Years), &(date.Months), &(date.Days), &(date.Hours), &(date.Minutes), &(date.Seconds)); erf != nil {
			continue
		}
		payment.CreatedAt = date
		payments = append(payments, payment)
	}
	if len(payments) == 0 {
		return payments, state.DT_STATUS_NO_RECORD_FOUND, nil
	}
	return payments, state.DT_STATUS_OK, nil
}

func (payrepo *PaymentRepo) DeletePaymentByID(ctx context.Context) (int, error) {
	paymentID := uint64(ctx.Value("payment_id").(uint64))
	dc, er := payrepo.DB.Exec(ctx, "DELETE from payin where id=$1", paymentID)
	if er != nil || dc.RowsAffected() == 0 {
		if dc.RowsAffected() == 0 {
			return state.DT_STATUS_NO_ROW_DELETED, er
		}
		return state.DT_STATUS_DELETION_FAILED, er
	}
	return state.DT_STATUS_OK, nil
}

func (payrepo *PaymentRepo) UpdatePaymentStatus(ctx context.Context) (int, error) {
	paymentID := ctx.Value("payment_id").(uint64)
	status := ctx.Value("payment_status").(uint8)
	var updateStatus int
	er := payrepo.DB.QueryRow(ctx, "select * updatePayInStatus($1,$2)", paymentID, status).Scan(&updateStatus)
	if er != nil {
		return -3, er
	}
	return updateStatus, nil
}
