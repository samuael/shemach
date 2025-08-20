package pgx_storage

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/samuael/shemach/backend/pkg/constants/types"
)

var (
	ErrDeletionNotSuccesful            = errors.New("deletion was not succesful")
	ErrUpdateNotSuccesful              = errors.New("update was not succesful")
	ErrPropertyIDAndIntervalIDRequired = errors.New("property ID and interval ID is required")
)

type UserRepo struct {
	DB *pgxpool.Pool
}

func (repo *UserRepo) InsertReplaceUserProfileImage(ctx context.Context, path string, userid uint64) (status int8, oldPath string, imgID uint64, err error) {
	err = repo.DB.QueryRow(ctx, "select * from checkInsertUserImageURL($1, $2)", path, userid).Scan(&status, &oldPath, &imgID)
	return
}

// GetUserByID retrieves user by details
func (repo *UserRepo) GetUserByID(ctx context.Context, id uint64) (*types.User, error) {
	user := &types.User{}
	var err error
	var bio, telegram sql.NullString
	var profileImageID sql.NullInt64
	if err = repo.DB.QueryRow(ctx, "select id,firstname,lastname,phone,email,telegram,profile_img_id,created_at,password,role ,bio  from users where id=$1", id).
		Scan(&(user.ID), &(user.Firstname), &(user.Lastname), &(user.Phone), &(user.Email), &(telegram), &(profileImageID), &(user.CreatedAt), &(user.Password), &(user.Role), &(bio)); err != nil {
		return nil, err
	}
	user.Bio = bio.String
	user.Telegram = telegram.String
	if profileImageID.Valid {
		user.ProfileImageID = uint64(profileImageID.Int64)
	}
	return user, nil
}

// CheckVerifiedInfoAndRegisterUser checks if the verified phone number is found in the verified phone number informations and if found, it will register a new user with the specified detailed information
func (repo *UserRepo) CheckVerifiedInfoAndRegisterUser(ctx context.Context, session *types.TempoRegistrationSession, userData *types.RegistrationData) (int64, error) {
	var statusValueOrID int64
	return statusValueOrID, repo.DB.QueryRow(ctx, "select * from checkAndRegisterUserInformation($1, $2, $3, $4, $5, $6, $7, $8)", session.ID, session.Phone, userData.Email, userData.TelegramID, userData.FirstName, userData.LastName, userData.Password, userData.Role).Scan(&statusValueOrID)
}

// ConfirmPhoneShortCode checks if the user phone and secret code and returns a status indicator
// or if successful the status information
func (repo *UserRepo) ConfirmPhoneShortCode(ctx context.Context, phone, code string) (int64, error) {
	var selectedNumber int64
	return selectedNumber, repo.DB.QueryRow(ctx, "select checkVerificationCode($1, $2)", phone, code).Scan(&selectedNumber)
}

func (repo *UserRepo) DeleteExpiredVerifiedPhones(ctx context.Context, timestamp int64) (int, error) {
	var deletedCount int
	return deletedCount, repo.DB.QueryRow(ctx, `WITH deleted AS (
		delete from verified_phones where created_at <= $1 returning id
	)
	SELECT COUNT(*) FROM deleted;`, timestamp).Scan(&deletedCount)
}

func (repo *UserRepo) DeleteExpiredTempoUsers(ctx context.Context, timestamp int64) (int, error) {
	var deletedCount int
	return deletedCount, repo.DB.QueryRow(ctx, `WITH deleted AS (
		delete from tempo_registration_info where created_at <= $1 or trials < 0 returning id
	)
	SELECT COUNT(*) FROM deleted;`, timestamp).Scan(&deletedCount)
}

func (repo *UserRepo) GetTempoUserInfo(ctx context.Context, phone string) (*types.UserAuthData, error) {
	uData := new(types.UserAuthData)
	if err := repo.DB.QueryRow(ctx, "select id,phone,code,created_at,trials from tempo_registration_info where phone= $1", phone).Scan(&(uData.ID), &(uData.Phone), &(uData.Code), &(uData.CreatedAt), &(uData.Trials)); err != nil {
		return nil, err
	}
	return uData, nil
}

// TempoRegisterNewUser checks the availability of the number in the system and returns status message if found or timestamp of creation if not found or at the stage of being verified.
//
// -- -1 if number found in already registered users
// -- -2 if insertion is not successful
// -- -{created_at} if number found in pending code confirmation
// -- {created_at} otherwise
func (repo *UserRepo) TempoRegisterNewUser(ctx context.Context, userData *types.UserAuthData) error {
	return repo.DB.QueryRow(ctx, "select registeruserforvarificationwithphonenumber($1,$2,$3)", userData.Phone, userData.Code, userData.Trials).Scan(&(userData.CreatedAt))
}

// UpdatePassword ...
func (repo *UserRepo) UpdatePassword(ctx context.Context) error {
	userid := ctx.Value("user_id").(uint64)
	password := ctx.Value("new_password").(string)
	uc, er := repo.DB.Exec(ctx, "Update users set password=$1 where id=$2", password, userid)
	if er != nil {
		return er
	} else if uc.RowsAffected() == 0 {
		return errors.New("no row affected")
	}
	return nil
}

// GetImageUrl ...  |
func (repo *UserRepo) GetImageUrl(ctx context.Context) string {
	userId := ctx.Value("user_id").(uint64)
	var imgurl string
	repo.DB.QueryRow(ctx, "select profile_image_id from users where id=$1", userId).Scan(&imgurl)
	return imgurl
}
func (repo *UserRepo) ChangeImageUrl(ctx context.Context) error {
	userId := ctx.Value("user_id").(uint64)
	imgurl := ctx.Value("image_url").(string)

	if er := repo.DB.QueryRow(ctx, "update users set imageurl=$1 where id=$2 returning id", imgurl, userId).Scan(&userId); er != nil {
		return er
	}
	return nil
}

// DeletePedingEmailConfirmation
// in this method if the account is new then not only the email in confirmation will be deleted
// but also the user account in the email will also be deleted so this functionality will be implemented withe the trigger i am going to write.
func (repo *UserRepo) DeletePendingEmailConfirmation(timestamp uint64) error {
	deleted := 0
	er := repo.DB.QueryRow(context.Background(), "delete from emailInConfirmation where created_at<$1  and is_new_account=$2", timestamp, false).Scan(&deleted)
	if er != nil || deleted == 0 {
		return errors.New("no row deleted ")
	}
	var ids []uint64

	rows, err := repo.DB.Query(context.Background(), "select userid from emailInConfirmation where created_at<$1  and is_new_account=$2", timestamp, true)
	if err != nil {
		return err
	}
	for rows.Next() {
		var id uint64
		r := rows.Scan(&id)
		if r != nil {
			continue
		}
		ids = append(ids, id)
	}

	var count uint64
	er = repo.DB.QueryRow(context.Background(), "select * from deleteUnconfirmedAdmins($1)", ids).Scan(&count)
	if er != nil || count == 0 {
		return er
	}
	return nil
}

// GetUserByPhone retrieves user detail by phone number
func (repo *UserRepo) GetUserByPhone(ctx context.Context, phone string) (*types.User, error) {
	user := &types.User{
		Phone: phone,
	}
	var bio, telegram, email sql.NullString
	var profileImageID sql.NullInt64
	err := repo.DB.QueryRow(ctx, "select  id,firstname,lastname,email,telegram,profile_img_id,created_at,password,role ,bio  from users where phone=$1", phone).
		Scan(&user.ID, &user.Firstname, &user.Lastname, &email, &telegram, &profileImageID, &user.CreatedAt, &user.Password, &user.Role, &bio)
	if err != nil {
		return nil, err
	}
	if bio.Valid {
		user.Bio = bio.String
	}
	user.Email = email.String
	if telegram.Valid {
		user.Telegram = telegram.String
	}
	if profileImageID.Valid {
		user.ProfileImageID = uint64(profileImageID.Int64)
	}
	return user, nil
}

// RemoveTempoCXP ...
func (repo *UserRepo) RemoveTempoCXP(ctx context.Context, phone string) error {
	count := 0
	ere := repo.DB.QueryRow(ctx, "delete from tempo_cxp where phone=$1 returning id", phone).Scan(&count)
	if count <= 0 {
		return errors.New("now row is deleted")
	}
	return ere
}

func (repo *UserRepo) RemoveExpiredCXPConfirmations(timestamp uint64) (count int, er error) {
	phones := []string{}
	recs, er := repo.DB.Query(context.Background(), "select phone from tempo_cxp where created_at<$1", timestamp)
	if er != nil {
		return -1, er
	}
	for recs.Next() {
		var phone string
		er := recs.Scan(&phone)
		if er == nil {
			phones = append(phones, phone)
		}
	}
	var deletedCount int
	erf := repo.DB.QueryRow(context.Background(), "select * from deleteExpredCXPAccount($1)", phones).Scan(&deletedCount)
	if erf != nil {
		return 0, erf
	}
	return count, nil
}

func (repo *UserRepo) ConfirmUserEmailUpdate(ctx context.Context, id uint64, newemail, oldemail string) error {
	confirm, er := repo.GetEmailInConfirmationByID(ctx, id)
	if er != nil {
		return er
	}
	if !(confirm.IsNewAccount) && confirm.Email == newemail && confirm.OldEmail == oldemail {
		uc, er := repo.DB.Exec(ctx, "update users set email=$1 where email=$2 returning id", newemail, oldemail)
		if uc.RowsAffected() == 0 || er != nil {
			if er != nil {
				return er
			}
			return errors.New("no rows affected")
		}
	} else if (confirm.Email == newemail) && (confirm.IsNewAccount) {
		raff, erf := repo.DB.Exec(ctx, "delete from emailInConfirmation where id=$1", id)
		if erf != nil || raff.RowsAffected() == 0 {
			if erf != nil {
				return erf
			}
			return errors.New("no row was deleted")
		}
	}
	return errors.New("unauthorized")
}

func (repo *UserRepo) GetEmailInConfirmationByID(ctx context.Context, id uint64) (*types.EmailConfirmation, error) {
	confirm := &types.EmailConfirmation{}
	er := repo.DB.QueryRow(ctx, "select  id ,userid ,new_email, is_new_account, old_email,created_at from emailInConfirmation where id=$1", id).
		Scan(&(confirm.ID), &(confirm.UserID), &(confirm.Email), &(confirm.IsNewAccount), &(confirm.OldEmail), &(confirm.CreatedAt))
	if er != nil {
		return nil, er
	}
	return confirm, nil
}

func (repo *UserRepo) SaveForgotPasswordDetail(ctx context.Context, phone, code string, phoneUsed, emailUsed, useTelegram bool, maxTrials uint8) (statusCode uint64, shortCode string, err error) {
	boolToInt := func(val bool) uint8 {
		if val {
			return 1
		}
		return 0
	}
	err = repo.DB.QueryRow(ctx, "select status, msg from registerForgotPasswordShortCodePhoneNumber($1, $2,$3, $4, $5, $6)", phone, code, boolToInt(phoneUsed), boolToInt(emailUsed), boolToInt(useTelegram), maxTrials).Scan(&statusCode, &shortCode)
	return
}

func (repo *UserRepo) DeleteExpiredForgotPasswordSession(ctx context.Context, timestamp uint64) (int64, error) {
	tag, err := repo.DB.Exec(ctx, "delete from forgot_password_shortcode where created_at < $1", timestamp)
	if err != nil {
		return 0, err
	}
	if !tag.Delete() {
		return 0, ErrDeletionNotSuccesful
	}
	return tag.RowsAffected(), nil
}

func (repo *UserRepo) ConfirmForgotPasswordShortcode(ctx context.Context, phone, shortCode string) (int64, error) {
	var statusCode int64
	return statusCode, repo.DB.QueryRow(ctx, "select confirmForgotPasswordShortcode($1, $2)", phone, shortCode).Scan(&statusCode)
}

func (repo *UserRepo) ChangeUserPassword(ctx context.Context, userID uint64, newPassword string) error {
	tag, err := repo.DB.Exec(ctx, "update users set password=$1 where id=$2", newPassword, userID)
	if err != nil {
		return err
	}
	if !tag.Update() || tag.RowsAffected() == 0 {
		return ErrDeletionNotSuccesful
	}
	return nil
}
