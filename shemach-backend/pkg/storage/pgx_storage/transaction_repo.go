package pgx_storage

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/samuael/shemach/shemach-backend/pkg/constants/model"
	"github.com/samuael/shemach/shemach-backend/pkg/transaction"
)

type TransactionRepo struct {
	DB *pgxpool.Pool
}

func NewTransactionRepo(conn *pgxpool.Pool) transaction.ITransactionRepo {
	return &TransactionRepo{
		DB: conn,
	}
}

func (repo *TransactionRepo) CreateNewTransaction(ctx context.Context, transaction *model.Transaction) error {
	return repo.DB.QueryRow(ctx, `insert into transaction(
		price,
		quantity,
		state,
		description,
		crop_id,
		requester_id,
		requester_store_ref,
		seller_id,
		seller_store_ref
	) values($1,$2,$3,$4,$5,$6,$7,$8,$9) returning transaction_id`, transaction.RequestingPrice, transaction.Quantity, transaction.State,
		transaction.Description, transaction.ProductID, transaction.RequesterID,
		transaction.RequesterStoreRef, transaction.SellerID, transaction.SellerStoreRef).Scan(&(transaction.ID))
}

// GetMyActiveTransactions
func (repo *TransactionRepo) GetMyActiveTransactions(ctx context.Context, userid uint64) ([]*model.Transaction, error) {
	transactions := []*model.Transaction{}
	rows, er := repo.DB.Query(ctx, `select transaction_id,price,quantity,state,deadline,description,crop_id,requester_id,
	requester_store_ref,seller_id,seller_store_ref,created_at,kebd_amount,guarantee_amount from transaction where requester_id=$1 or seller_id=$1`, userid)
	if er != nil {
		return nil, er
	}
	for rows.Next() {
		t := &model.Transaction{}
		er := rows.Scan(
			&(t.ID), &(t.RequestingPrice), &(t.Quantity), &(t.State), &(t.Deadline), &(t.Description), &(t.ProductID), &(t.RequesterID),
			&(t.RequesterStoreRef), &(t.SellerID), &(t.SellerStoreRef), &(t.CreatedAt), &(t.KebdAmount),
			&(t.GuaranteeAmount))
		if er != nil {
			continue
		}
		transactions = append(transactions, t)
	}
	return transactions, nil
}

// DeclineTransaction
func (repo *TransactionRepo) DeclineTransaction(ctx context.Context, transactionid uint64, userid uint64) (int, error) {
	var state int
	er := repo.DB.QueryRow(ctx, "select * from declineTransaction($1,$2)", transactionid, userid).Scan(&state)
	return state, er
}

func (repo *TransactionRepo) GetTransactionByID(ctx context.Context, transactionid uint64) (*model.Transaction, error) {
	t := &model.Transaction{}
	er := repo.DB.QueryRow(ctx, `select transaction_id,price,quantity,state,deadline,description,crop_id,requester_id,
	requester_store_ref,seller_id,seller_store_ref,created_at,kebd_amount,guarantee_amount from transaction where transaction_id=$1`, transactionid).Scan(
		&(t.ID), &(t.RequestingPrice), &(t.Quantity), &(t.State), &(t.Deadline), &(t.Description), &(t.ProductID), &(t.RequesterID),
		&(t.RequesterStoreRef), &(t.SellerID), &(t.SellerStoreRef), &(t.CreatedAt), &(t.KebdAmount),
		&(t.GuaranteeAmount))
	if er != nil {
		return nil, er
	}
	return t, nil
}

func (repo *TransactionRepo) SaveTransactionAmendmentRequest(ctx context.Context, transaction *model.TransactionRequest) error {
	var state int
	println(transaction.State, transaction.TransactionID, transaction.Description, transaction.Price, transaction.Quantity)
	if er := repo.DB.QueryRow(ctx, `select * from createNewTrasactionAmendmentRequest($1,$2,$3,$4,$5)`, transaction.State, transaction.TransactionID, transaction.Description, transaction.Price, transaction.Quantity).Scan(&(state)); er != nil {
		if er != nil {
			println(er.Error())
			return er
		}
		return errors.New("problem while creating a transaction information")
	}
	if state > 0 {
		println(state)
		transaction.ID = uint64(state)
	}
	return nil
}

func (repo *TransactionRepo) GetTransactionRequestByID(ctx context.Context, id uint64) (*model.TransactionRequest, error) {
	t := &model.TransactionRequest{}
	er := repo.DB.QueryRow(ctx, `select 
	transaction_changes_id,state,transaction_id,description,price,qty,created_at from transaction_changes where transaction_changes_id=$1`, id).
		Scan(
			&(t.ID), &(t.State), &(t.TransactionID), &(t.Description), &(t.Price), &(t.Quantity),
			&(t.CreatedAt),
		)
	if er != nil {
		println(er.Error())
		return nil, er
	}
	return t, nil
}

func (repo *TransactionRepo) AcceptTransactionAmendmentRequest(ctx context.Context, merchantid, requestid uint64) (int, error) {
	var state int
	er := repo.DB.QueryRow(ctx, "select * from acceptTransactionAmendmentRequest($1,$2)", merchantid, requestid).Scan(&state)
	if er != nil {
		return -6, er
	}
	return state, nil
}

func (repo *TransactionRepo) SetTransactionAmendment(ctx context.Context, userid uint64, req *model.TransactionRequest) (int, error) {
	if req.ID <= 0 {
		return -4, errors.New("id has to be specified")
	}
	if req.Price <= 0 || req.Quantity <= 0 {
		return -5, errors.New("price or Quantity has to be specified")
	}
	var state int
	er := repo.DB.QueryRow(ctx, "select * from amendTransaction($1, $2,$3,$4,$5)", userid, req.ID, req.Price, req.Quantity, req.Description).Scan(&state)
	return state, er
}

func (repo *TransactionRepo) CreateKebdRequest(ctx context.Context, cxpid uint64, input *model.KebdAmountRequest) (int, error) {
	var state int
	er := repo.DB.QueryRow(ctx, "select * from createKebdRequest( $1, $2,$3,$4,$5)", cxpid, input.KebdAmount, input.Deadline, input.Description, input.TransactionID).Scan(&state)
	return state, er
}

func (repo *TransactionRepo) CreateKebdAmendmentRequest(ctx context.Context, merchantid uint64, input *model.KebdAmountRequest) (int, error) {
	var state int
	er := repo.DB.QueryRow(ctx, "select * from createKebdAmendmentRequest( $1, $2,$3,$4,$5)", merchantid, input.KebdAmount, input.Deadline, input.Description, input.TransactionID).Scan(&state)
	return state, er
}
func (repo *TransactionRepo) AmendKebdRequest(ctx context.Context, cxpid uint64, input *model.KebdAmountRequest) (int, error) {
	var state int
	er := repo.DB.QueryRow(ctx, "select * from ammendKebdRequest( $1, $2,$3,$4,$5)", cxpid, input.KebdAmount, input.Deadline, input.Description, input.TransactionID).Scan(&state)
	return state, er
}

func (repo *TransactionRepo) CreateGuaranteeRequest(ctx context.Context, cxpid uint64, input *model.GuaranteeAmountRequest) (int, error) {
	var state int
	er := repo.DB.QueryRow(ctx, "select * from createGuaranteeRequest($1,$2, $3,$4)", cxpid, input.Amount, input.Description, input.TransactionID).Scan(&state)
	return state, er
}

func (repo *TransactionRepo) SellerAcceptTransaction(ctx context.Context, sellerid, transactionID uint64) error {
	er := repo.DB.QueryRow(ctx, "update transaction set state=10 where transaction_id=$1 and seller_id=$2 returning transaction_id", transactionID, sellerid).Scan(&(transactionID))
	return er
}
func (repo *TransactionRepo) BuyerAcceptTransaction(ctx context.Context, sellerid, transactionID uint64) error {
	er := repo.DB.QueryRow(ctx, "update transaction set state=11 where transaction_id=$1 and seller_id=$2 returning transaction_id", transactionID, sellerid).Scan(&(transactionID))
	return er
}

func (repo *TransactionRepo) GetTransactionNotificationByTransactionID(ctx context.Context, trid uint64) (*model.TransactionRequest, error) {
	treq := &model.TransactionRequest{}
	er := repo.DB.QueryRow(ctx, `select transaction_changes_id
	,state
	,transaction_id
	,description
	,price
	,qty
	,created_at from transaction_changes where transaction_id=$1`, trid).Scan(
		&(treq.ID), &(treq.State), &(treq.TransactionID), &(treq.Description), &(treq.Price),
		&(treq.Quantity), &(treq.CreatedAt),
	)
	return treq, er
}
func (repo *TransactionRepo) GetKebdNotificationByTransactionID(ctx context.Context, trid uint64) (*model.KebdAmountRequest, error) {
	kreq := &model.KebdAmountRequest{}
	er := repo.DB.QueryRow(ctx, `select kebd_transaction_info_id
	,transaction_id
	,state
	,kebd_amount
	,deadline
	,description
	,created_at  from kebd_transaction_info where transaction_id=$1`, trid).Scan(
		&(kreq.ID), &(kreq.TransactionID), &(kreq.State), &(kreq.KebdAmount), &(kreq.Deadline),
		&(kreq.Description), &(kreq.CreatedAt),
	)
	return kreq, er
}
func (repo *TransactionRepo) GetGuaranteeNotificationByTransactionID(ctx context.Context, trid uint64) (*model.GuaranteeAmountRequest, error) {
	greq := &model.GuaranteeAmountRequest{}
	er := repo.DB.QueryRow(ctx, ` select transaction_guarantee_info_id
	,transaction_id
	,state
	,description
	,amount
	,created_at from transaction_guarantee_info where transaction_id=$1`, trid).Scan(
		&(greq.ID), &(greq.TransactionID), &(greq.State),
		&(greq.Description), &(greq.Amount),
		&(greq.CreatedAt),
	)
	return greq, er
}
