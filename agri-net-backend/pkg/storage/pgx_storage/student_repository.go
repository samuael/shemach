package pgx_storage

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/jackc/pgx/pgxpool"
	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/model"
	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/state"
)

type StudentRepo struct {
	DB *pgxpool.Pool
}

func NewStudentRepo(db *pgxpool.Pool) *StudentRepo {
	return &StudentRepo{
		DB: db,
	}
}

func (sturepo *StudentRepo) RegisterStudent(ctx context.Context) (*model.Student, int, error) {
	student := ctx.Value("student").(*model.Student)
	if ers := sturepo.DB.QueryRow(ctx, `select id, fullname, sex, age, birth_date, accamic_status, address, phone, paid, status, registered_by, round_id, imgurl, registered_at from registerStudent(
		$1, 
		cast ($2 as char), 
		cast ( $3 as decimal),
		cast ($4 as varchar),
		cast ($5 as varchar),
		cast ( $6 as decimal),
		cast ( $7 as integer),
		cast ( $8 as integer),
		cast ( $9 as integer),
		cast( $10 as varchar),
		cast( $11 as integer ),
		cast ( $12 as integer),
		cast ( $13 as integer),
		cast ( $14 as integer),
		cast ( $15 as integer),
		cast ( $16 as smallint),
		cast ( $17 as smallint),
		cast ( $18  as integer),
		cast ( $19  as integer),
		cast ( $20  as integer),
		cast ( $21  as integer),
		cast ( $22 as smallint),
		cast ( $23 as smallint),
		cast ($24 as varchar),
		cast ($25 as varchar),
		cast ($26 as varchar),
		cast ($27 as varchar),
		cast ($28 as varchar),
		cast ($29 as varchar)
	)`, student.Fullname, student.Sex, student.Age, student.AccStatus, student.Phone, student.PaidAmount,
		student.Status, student.RegisteredBy, student.RoundID, student.Imgurl, 0, student.BirthDate.Years, student.BirthDate.Months,
		student.BirthDate.Days, student.BirthDate.Hours, student.BirthDate.Minutes, student.BirthDate.Seconds,
		student.RegisteredAt.Years, student.RegisteredAt.Months, student.RegisteredAt.Days, student.RegisteredAt.Hours,
		student.RegisteredAt.Minutes, student.RegisteredAt.Seconds,
		student.Address.City, student.Address.Region, student.Address.Zone, student.Address.Woreda,
		student.Address.Kebele, student.Address.UniqueAddressName,
	).Scan(
		//id, fullname, sex, age, birth_date, accamic_status, address, phone,
		// paid, status, registered_by , round_id, imgurl, mark, registered_at
		&(student.ID),
		&(student.Fullname),
		&(student.Sex),
		&(student.Age),
		&(student.BirthDateRef),
		&(student.AccStatus),
		&(student.AddressRef),
		&(student.Phone),
		&(student.PaidAmount),
		&(student.Status),
		&(student.RegisteredBy),
		&(student.RoundID),
		&(student.Imgurl),
		&(student.RegisteredAtRef),
	); ers != nil {
		log.Println(ers.Error())
		if strings.Contains(ers.Error(), "duplicate key value violates unique constraint \"unique_phone_number\"") {
			return student, state.DT_STATUS_DUPLICATE_PHONE_NUMBER, ers
		}
		return student, state.DT_STATUS_DBQUERY_ERROR, ers
	}
	date := &model.Date{}
	er := sturepo.DB.QueryRow(ctx, "SELECT id,year,month,day,hour,minute,second from eth_date where id=$1", student.RegisteredAtRef).
		Scan(&(date.ID), &(date.Years), &(date.Months), &(date.Days), &(date.Hours), &(date.Minutes), &(date.Seconds))
	if er != nil {
		log.Println(er.Error())
		return student, state.DT_STATUS_INCOMPLETE_DATA, er
	}
	student.RegisteredAt = date
	bdate := &model.Date{}
	er = sturepo.DB.QueryRow(ctx, "SELECT id,year,month,day,hour,minute,second from eth_date where id=$1", student.BirthDateRef).
		Scan(&(bdate.ID), &(bdate.Years), &(bdate.Months), &(bdate.Days), &(bdate.Hours), &(bdate.Minutes), &(bdate.Seconds))
	if er != nil {
		log.Println(er.Error())
		return student, state.DT_STATUS_INCOMPLETE_DATA, er
	}
	student.BirthDate = bdate

	address := &model.Address{}
	er = sturepo.DB.QueryRow(ctx, "SELECT id, city, region, zone, woreda, kebele, unique_address from addresses where id=$1", student.AddressRef).
		Scan(&(address.ID), &(address.City), &(address.Region), &(address.Zone), &(address.Woreda), &(address.Kebele), &(address.UniqueAddressName))
	if er != nil {
		log.Println(er.Error())
		return student, state.DT_STATUS_INCOMPLETE_DATA, er
	}
	student.Address = address
	return student, state.DT_STATUS_OK, nil
}

func (sturepo *StudentRepo) GetStudentByID(ctx context.Context) (*model.Student, int, error) {
	studentID := ctx.Value("student_id").(uint64)
	student := &model.Student{}
	if er := sturepo.DB.QueryRow(ctx, "select id, fullname, sex, age, birth_date, accamic_status, address, phone, paid, status, registered_by , round_id, imgurl, mark, registered_at from student where id=$1", studentID).
		Scan(
			&(student.ID),
			&(student.Fullname),
			&(student.Sex),
			&(student.Age),
			&(student.BirthDateRef),
			&(student.AccStatus),
			&(student.AddressRef),
			&(student.Phone),
			&(student.PaidAmount),
			&(student.Status),
			&(student.RegisteredBy),
			&(student.RoundID),
			&(student.Imgurl),
			&(student.MarkedRef),
			&(student.RegisteredAtRef),
		); er != nil {
		if strings.Contains(er.Error(), "no rows in result set") {
			return student, state.DT_STATUS_RECORD_NOT_FOUND, er
		}
		if strings.Contains(er.Error(), "duplicate key value violates unique constraint \"unique_phone_number\"") {
			return student, state.DT_STATUS_DUPLICATE_PHONE_NUMBER, er
		}
		return student, state.DT_STATUS_DBQUERY_ERROR, er
	}
	date := &model.Date{}
	er := sturepo.DB.QueryRow(ctx, "SELECT id,year,month,day,hour,minute,second from eth_date where id=$1", student.RegisteredAtRef).
		Scan(&(date.ID), &(date.Years), &(date.Months), &(date.Days), &(date.Hours), &(date.Minutes), &(date.Seconds))
	if er != nil {
		log.Println(er.Error())
		return student, state.DT_STATUS_INCOMPLETE_DATA, er
	}
	student.RegisteredAt = date
	bdate := &model.Date{}
	er = sturepo.DB.QueryRow(ctx, "SELECT id,year,month,day,hour,minute,second from eth_date where id=$1", student.BirthDateRef).
		Scan(&(bdate.ID), &(bdate.Years), &(bdate.Months), &(bdate.Days), &(bdate.Hours), &(bdate.Minutes), &(bdate.Seconds))
	if er != nil {
		log.Println(er.Error())
		return student, state.DT_STATUS_INCOMPLETE_DATA, er
	}
	student.BirthDate = bdate

	println(student.MarkedRef)
	if student.MarkedRef != 0 {
		mark := &model.SpecialCase{}
		er = sturepo.DB.QueryRow(ctx, "SELECT id, reason, covered_amount, complete_fee from special_case where id=$1", student.MarkedRef).
			Scan(&(mark.ID), &(mark.Reason), &(mark.Amount), &(mark.Total))
		if er != nil {
			log.Println(er.Error())
			return student, state.DT_STATUS_INCOMPLETE_DATA, er
		}
		student.Marked = mark
	}

	address := &model.Address{}
	er = sturepo.DB.QueryRow(ctx, "SELECT id, city, region, zone, woreda, kebele, unique_address from addresses where id=$1", student.AddressRef).
		Scan(&(address.ID), &(address.City), &(address.Region), &(address.Zone), &(address.Woreda), &(address.Kebele), &(address.UniqueAddressName))
	if er != nil {
		log.Println(er.Error())
		return student, state.DT_STATUS_INCOMPLETE_DATA, er
	}
	student.Address = address
	return student, state.DT_STATUS_OK, nil
}

func (sturepo *StudentRepo) CheckWhetherTheStudentWithThisPhoneNumberExists(ctx context.Context) (int, error) {
	studentPhone := ctx.Value("student_phone").(string)
	studentsQuantity := 0
	if er := sturepo.DB.QueryRow(ctx, "select count(*) as id from student where phone=$1", studentPhone).
		Scan(&(studentsQuantity)); er != nil {
		return -1, er
	}
	return studentsQuantity, nil
}

func (sturepo *StudentRepo) UpdateStudent(ctx context.Context) (*model.Student, int, error) {
	student := ctx.Value("updated_student").(*model.Student)
	if uc, er := sturepo.DB.Exec(ctx, "update student set fullname=$1,age=$2, sex=$3, accamic_status=$4, phone=$5, status=$6, round_id=$7,imgurl=$8 where id =$9",
		student.Fullname, student.Age, student.Sex, student.AccStatus, student.Phone, student.Status, student.RoundID, student.Imgurl, student.ID,
	); er != nil || uc.RowsAffected() == 0 {
		if er != nil {
			return nil, state.DT_STATUS_DBQUERY_ERROR, er
		}
		return nil, state.DT_STATUS_NO_RECORD_UPDATED, fmt.Errorf("no row is updated")
	}
	return student, state.DT_STATUS_OK, nil
}

func (sturepo *StudentRepo) GetStudentsOfRound(ctx context.Context) ([]*model.Student, int, error) {
	roundID := ctx.Value("round_id").(uint)
	offset := ctx.Value("offset").(uint)
	limit := ctx.Value("limit").(uint)
	students := []*model.Student{}
	uc, er := sturepo.DB.Query(ctx, "SELECT id, fullname, sex, age, birth_date, accamic_status, address, phone, paid, status, registered_by, round_id, imgurl, mark, registered_at from student where round_id=$1 OFFSET $2 limit $3", roundID, offset, limit)
	if er != nil {
		println("  ---   ---- ", er.Error())
		return students, state.DT_STATUS_DBQUERY_ERROR, er
	}
	for uc.Next() {
		student := &model.Student{}
		if ers := uc.Scan(&(student.ID),
			&(student.Fullname),
			&(student.Sex),
			&(student.Age),
			&(student.BirthDateRef),
			&(student.AccStatus),
			&(student.AddressRef),
			&(student.Phone),
			&(student.PaidAmount),
			&(student.Status),
			&(student.RegisteredBy),
			&(student.RoundID),
			&(student.Imgurl),
			&(student.MarkedRef),
			&(student.RegisteredAtRef)); ers != nil {
			println(ers.Error())
			continue
		}
		date := &model.Date{}
		er := sturepo.DB.QueryRow(ctx, "SELECT id,year,month,day,hour,minute,second from eth_date where id=$1", student.RegisteredAtRef).
			Scan(&(date.ID), &(date.Years), &(date.Months), &(date.Days), &(date.Hours), &(date.Minutes), &(date.Seconds))
		if er != nil {
			println("  ---  1 ---- ", er.Error())
			continue
		}
		student.RegisteredAt = date
		bdate := &model.Date{}
		er = sturepo.DB.QueryRow(ctx, "SELECT id,year,month,day,hour,minute,second from eth_date where id=$1", student.BirthDateRef).
			Scan(&(bdate.ID), &(bdate.Years), &(bdate.Months), &(bdate.Days), &(bdate.Hours), &(bdate.Minutes), &(bdate.Seconds))
		if er != nil {
			println("  ---  2  ---- ", er.Error())
			continue
		}
		student.BirthDate = bdate
		if student.MarkedRef != 0 {
			mark := &model.SpecialCase{}
			er = sturepo.DB.QueryRow(ctx, "SELECT id, reason, covered_amount, complete_fee from special_case where id=$1", student.MarkedRef).
				Scan(&(mark.ID), &(mark.Reason), &(mark.Amount), &(mark.Total))
			if er != nil {
				println("  ---  3 ---- ", er.Error())
				continue
			}
			student.Marked = mark
		}
		address := &model.Address{}
		er = sturepo.DB.QueryRow(ctx, "SELECT id, city, region, zone, woreda, kebele, unique_address from addresses where id=$1", student.AddressRef).
			Scan(&(address.ID), &(address.City), &(address.Region), &(address.Zone), &(address.Woreda), &(address.Kebele), &(address.UniqueAddressName))
		if er != nil {
			println("  ---  4 ---- ", er.Error())
			continue
		}
		student.Address = address
		students = append(students, student)
	}
	if len(students) == 0 {
		return students, state.DT_STATUS_RECORD_NOT_FOUND, fmt.Errorf("no students found")
	}
	return students, state.DT_STATUS_OK, nil
}

func (sturepo *StudentRepo) GetStudentsOfCategory(ctx context.Context) ([]*model.Student, int, error) {
	categoryID := ctx.Value("category_id").(uint)
	offset := ctx.Value("offset").(uint)
	limit := ctx.Value("limit").(uint)
	students := []*model.Student{}
	uc, er := sturepo.DB.Query(ctx, `select 
	student.id,
	student.fullname,
	student.sex,
	student.age,
	student.birth_date,
	student.accamic_status,
	student.address,
	student.phone,
	student.paid,
	student.status,
	student.registered_by,
	student.round_id,
	student.imgurl,
	student.mark,
	student.registered_at from student INNER JOIN round ON round.id=student.round_id
	INNER JOIN category ON category.id = round.categoryid WHERE category.id= $1 OFFSET $2 LIMIT $3;`, categoryID, offset, limit)
	if er != nil {
		println("  ---   ---- ", er.Error())
		return students, state.DT_STATUS_DBQUERY_ERROR, er
	}
	for uc.Next() {
		student := &model.Student{}
		if ers := uc.Scan(&(student.ID),
			&(student.Fullname),
			&(student.Sex),
			&(student.Age),
			&(student.BirthDateRef),
			&(student.AccStatus),
			&(student.AddressRef),
			&(student.Phone),
			&(student.PaidAmount),
			&(student.Status),
			&(student.RegisteredBy),
			&(student.RoundID),
			&(student.Imgurl),
			&(student.MarkedRef),
			&(student.RegisteredAtRef)); ers != nil {
			println(ers.Error())
			continue
		}
		date := &model.Date{}
		er := sturepo.DB.QueryRow(ctx, "SELECT id,year,month,day,hour,minute,second from eth_date where id=$1", student.RegisteredAtRef).
			Scan(&(date.ID), &(date.Years), &(date.Months), &(date.Days), &(date.Hours), &(date.Minutes), &(date.Seconds))
		if er != nil {
			println("  ---  1 ---- ", er.Error())
			continue
		}
		student.RegisteredAt = date
		bdate := &model.Date{}
		er = sturepo.DB.QueryRow(ctx, "SELECT id,year,month,day,hour,minute,second from eth_date where id=$1", student.BirthDateRef).
			Scan(&(bdate.ID), &(bdate.Years), &(bdate.Months), &(bdate.Days), &(bdate.Hours), &(bdate.Minutes), &(bdate.Seconds))
		if er != nil {
			println("  ---  2  ---- ", er.Error())
			continue
		}
		student.BirthDate = bdate
		if student.MarkedRef != 0 {
			mark := &model.SpecialCase{}
			er = sturepo.DB.QueryRow(ctx, "SELECT id, reason, covered_amount, complete_fee from special_case where id=$1", student.MarkedRef).
				Scan(&(mark.ID), &(mark.Reason), &(mark.Amount), &(mark.Total))
			if er != nil {
				continue
			}
			student.Marked = mark
		}
		address := &model.Address{}
		er = sturepo.DB.QueryRow(ctx, "SELECT id, city, region, zone, woreda, kebele, unique_address from addresses where id=$1", student.AddressRef).
			Scan(&(address.ID), &(address.City), &(address.Region), &(address.Zone), &(address.Woreda), &(address.Kebele), &(address.UniqueAddressName))
		if er != nil {
			println("  ---  4 ---- ", er.Error())
			continue
		}
		student.Address = address
		students = append(students, student)
	}
	if len(students) == 0 {
		return students, state.DT_STATUS_RECORD_NOT_FOUND, fmt.Errorf("no students found")
	}
	return students, state.DT_STATUS_OK, nil
}

func (sturepo *StudentRepo) GetStudentImageUrl(ctx context.Context) (string, uint8) {
	studentID := ctx.Value("student-id")
	imgurl := ""
	if err := sturepo.DB.QueryRow(ctx, "SELECT imgurl from student where id=$1", studentID).Scan(&imgurl); err != nil {
		return imgurl, state.DT_STATUS_DBQUERY_ERROR
	}
	return imgurl, state.DT_STATUS_OK
}
func (sturepo *StudentRepo) ChangeStudentImageUrl(ctx context.Context) error {
	imgurl := ctx.Value("image_url").(string)
	userid := ctx.Value("student_id").(uint64)

	if cmd, err := sturepo.DB.Exec(ctx, "UPDATE student SET imgurl=$1  WHERE id=$2", imgurl, userid); cmd.RowsAffected() == 0 || err != nil {
		return errors.New("internal server error while updating the image url")
	}
	return nil
}

func (sturepo *StudentRepo) CreateSpecialCase(ctx context.Context) (*model.SpecialCase, int, error) {
	specialCase := ctx.Value("special_case").(*model.SpecialCase)
	if er := sturepo.DB.QueryRow(ctx, "SELECT id, reason, covered_amount,complete_fee from createSpecialCaseForStudent($1 , $2, $3 ,$4)", specialCase.Reason, specialCase.Amount, specialCase.Total, specialCase.StudentID).Scan(
		&(specialCase.ID),
		&(specialCase.Reason),
		&(specialCase.Amount),
		&(specialCase.Total),
	); er != nil {
		log.Println(er.Error())
		return specialCase, state.DT_STATUS_DBQUERY_ERROR, er
	} else {

		return specialCase, state.DT_STATUS_OK, nil
	}
}
func (sturepo *StudentRepo) GetSpecialCaseByID(ctx context.Context) (*model.SpecialCase, int, error) {
	specialCaseID := ctx.Value("special_case_id").(uint64)
	specialCase := &model.SpecialCase{}
	if er := sturepo.DB.QueryRow(ctx, "SELECT id, reason, covered_amount, complete_fee FROM special_case where id=$1", specialCaseID).Scan(
		&(specialCase.ID),
		&(specialCase.Reason),
		&(specialCase.Amount),
		&(specialCase.Total),
	); er != nil {
		log.Println(er.Error())
		return specialCase, state.DT_STATUS_DBQUERY_ERROR, er
	}
	return specialCase, state.DT_STATUS_OK, nil
}

func (sturepo *StudentRepo) UpdateSpecialCase(ctx context.Context) (int, error) {
	specialCase := ctx.Value("special_case").(*model.SpecialCase)
	if uc, er := sturepo.DB.Exec(ctx, "update special_case set reason=$2, covered_amount=$3, complete_fee=$4 where id=$1", specialCase.ID, specialCase.Reason, specialCase.Amount, specialCase.Total); uc.RowsAffected() == 0 || er != nil {
		if er != nil {
			log.Println(er.Error())
			return state.DT_STATUS_DBQUERY_ERROR, errors.New("database query error")
		}
		return state.DT_STATUS_NO_RECORD_UPDATED, nil
	}
	return state.DT_STATUS_OK, nil
}
