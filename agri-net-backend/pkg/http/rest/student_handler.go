package rest

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"math"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/samuael/Project/CarInspection/platforms/form"
	"github.com/samuael/Project/RegistrationSystem/pkg/constants/model"
	"github.com/samuael/Project/RegistrationSystem/pkg/constants/state"
	"github.com/samuael/Project/RegistrationSystem/pkg/round"
	"github.com/samuael/Project/RegistrationSystem/pkg/student"
	"github.com/samuael/Project/RegistrationSystem/platforms"
	"github.com/samuael/Project/RegistrationSystem/platforms/helper"
)

// IStudentHandler  ....
type IStudentHandler interface {
	RegisterStudent(c *gin.Context)
	UpdateStudent(*gin.Context)
	CreateStudentsSpecialCaseForStudent(*gin.Context)
	UpdateStudentSpecialCaseInformation(*gin.Context)

	UpdateStudentProfilePicture(c *gin.Context)
	UpdateStudentBirthDate(*gin.Context)
	UpdateStudentAddress(*gin.Context)

	DeleteStudent(*gin.Context)
	GetStudentByID(*gin.Context)
	GetStudentsStatus(*gin.Context)
	SearchStudents(*gin.Context)
	GetStudentStatus(*gin.Context)

	GetStudentsOfRound(c *gin.Context)
	GetStudentsOfCategory(c *gin.Context)
}

// StudentHandler struct representing.
type StudentHandler struct {
	Service      student.IStudentService
	RoundService round.IRoundService
}

// NewStudentHandler creates a new student handler function.
func NewStudentHandler(sservice student.IStudentService, roundser round.IRoundService) IStudentHandler {
	return &StudentHandler{
		Service:      sservice,
		RoundService: roundser,
	}
}

var addressVariablesCombination = []string{
	"invalid region",
	"invalid zone",
	"invalid region and zone",
	"invalid woreda",
	"invalid region and woreda",
	"invalid zone and woreda",
	"invalid region, zone, and woreda",
	"invalid kebele",
	"invalid kebele and region",
	"invalid kebele and zone",
	"invalid kebele, region, and zone",
	"invalid kebele and woreda",
	"invalid kebele, woreda, region",
	"invalid kebele, woreda, and zone",
	"invalid kebele, woreda, region, and zone",
}
var birthdateCombinations = []string{
	"missing valid year",
	"missing valid month",
	"missing valid year and month",
	"missing valid day",
	"missing valid day and year ",
	"missing valid day and month",
	"missing valid year, month, and day",
}

func (shan *StudentHandler) RegisterStudent(c *gin.Context) {
	res := &model.Student{
		Address:   &model.Address{},
		BirthDate: &model.Date{},
	}
	eres := &struct {
		Errs map[string]string `json:"errs"`
		Code int               `json:"status_code"`
		Msg  string            `json:"msg"`
	}{
		Errs: map[string]string{},
	}
	ctx := c.Request.Context()
	jdeconder := json.NewDecoder(c.Request.Body)
	if err := jdeconder.Decode(res); err != nil {
		eres.Code = http.StatusBadRequest
		eres.Msg = "bad request body"
		c.JSON(http.StatusBadRequest, eres)
		return
	}
	passed := true
	count := 0
	res.Fullname = strings.Replace(res.Fullname, "  ", " ", -1)
	if len(res.Fullname) < 7 || len(strings.Split(res.Fullname, " ")) < 2 {
		passed = false
		count += 1
		eres.Errs["Full name"] = "Invalid fullname value"
	} else {
		println(res.Fullname, "The Last Name: ", (strings.Split(res.Fullname, " ")[1]))
		if len(strings.Split(res.Fullname, " ")) >= 2 && len(strings.Split(res.Fullname, " ")[0]) < 3 {
			passed = false
			count += 1
			eres.Errs["First name"] = "Invalid firstname value"
		}
		if len(strings.Split(res.Fullname, " ")) >= 2 && len(strings.Split(res.Fullname, " ")[1]) < 3 {
			passed = false
			count += 1
			eres.Errs["Last name"] = "Invalid lastname value"
		}
	}
	if res.Sex == "F" ||
		res.Sex == "f" ||
		res.Sex == "Female" ||
		strings.HasPrefix(res.Sex, "F") ||
		strings.HasPrefix(res.Sex, "f") {
		res.Sex = "F"
	} else if res.Sex == "M" ||
		res.Sex == "m" ||
		res.Sex == "Male" ||
		strings.HasPrefix(res.Sex, "M") ||
		strings.HasPrefix(res.Sex, "m") {
		res.Sex = "M"
	} else {
		passed = false
		count += 1
		eres.Errs["Sex"] = "invalid sex value"
	}
	if len(res.Phone) <= 13 && len(res.Phone) >= 10 && form.MatchesPattern(res.Phone, form.PhoneRX) {
		if strings.HasPrefix(res.Phone, "0") {
			res.Phone = strings.Replace(res.Phone, "0", "+251", 1)
		}
	} else {
		passed = false
		count += 1
		eres.Errs["Phone"] = "Invalid phone number value"
	}
	if res.RoundID <= 0 {
		passed = false
		count += 1
		eres.Errs["Round ID"] = "Invalid round id"
	} else {
		ctx = context.WithValue(ctx, "round_id", uint64(res.RoundID))
		result, statusCode := shan.RoundService.CheckTheExistanceAndActivenessOfRound(ctx)
		if result == -2 {
			if statusCode == state.DT_STATUS_DBQUERY_ERROR {
				log.Println("Database Query Error student_handler :: 146")
			}
			eres.Msg = "Internal problem, please try again"
			passed = false
			count += 1
		} else if result == -1 {
			eres.Msg = "Round with this ID does not exist"
			passed = false
			count += 1
		} else if result == 0 {
			eres.Msg = "Round is not active for registration"
			passed = false
			count += 1
		}
	}
	if res.Address == nil ||
		len(res.Address.Region) < 3 ||
		len(res.Address.Zone) < 3 ||
		len(res.Address.Woreda) < 3 ||
		len(res.Address.Kebele) < 1 {
		val := 0
		if len(res.Address.Region) < 3 {
			val |= 1
		}
		if len(res.Address.Zone) < 3 {
			val |= 2
		}
		if len(res.Address.Woreda) < 3 {
			val |= 4
		}
		if len(res.Address.Kebele) < 3 {
			val |= 8
		}
		if val == 0 {
			val = 16
		}
		passed = false
		count += 1
		eres.Errs["Address"] = addressVariablesCombination[val]
	}
	if res.BirthDate.Years == 0 || res.BirthDate.Months == 0 || res.BirthDate.Days == 0 {
		val := 0
		if res.BirthDate.Years == 0 || res.BirthDate.Years <= 1920 || res.BirthDate.Years > platforms.GetCurrentEthiopianTime().Years {
			val |= 1
		}
		if res.BirthDate.Months == 0 || res.BirthDate.Months <= 0 || res.BirthDate.Months > 13 {
			val |= 2
		}
		if res.BirthDate.Days == 0 || res.BirthDate.Days <= 0 || res.BirthDate.Days > 30 {
			val |= 4
		}
		eres.Errs["Birth Date"] = birthdateCombinations[val]
		passed = false
		count += 1
	} else {
		year, months := platforms.GetAgeUsingBirthDate(res.BirthDate)
		limit_years, _ := strconv.Atoi(os.Getenv("DRIVING_LICENSE_AGE_LIMIT"))
		if (year + int(math.Round(float64(months)/float64(12)))) < limit_years {
			passed = false
			count += 1
			eres.Errs["Birth Date"] = "student doesn't reach the specified lowest age limit\n\r\t" + func() string {
				year, month := platforms.GetAgeUsingBirthDate(res.BirthDate)
				if res.Sex == "M" {
					return "His age is " + strconv.Itoa(year) + " years and " + strconv.Itoa(month) + " Months , which is < 18"
				} else if res.Sex == "F" {
					return "Her age is " + strconv.Itoa(year) + " years and " + strconv.Itoa(month) + " Months , which is < 18"
				}
				return "Student's age is " + strconv.Itoa(year) + " years and " + strconv.Itoa(month) + " Months , which is < 18"

			}()
		}
		res.Age = uint(year + int(math.Round(float64(months)/float64(12))))
	}
	if !passed {
		eres.Code = http.StatusBadRequest
		if eres.Msg == "" {
			eres.Msg = "Missing important inputs"
		}
		c.JSON(http.StatusBadRequest, eres)
		return
	}
	res.RegisteredBy = uint(ctx.Value("session").(*model.Session).ID)
	res.PaidAmount = 0
	res.Status = state.STUDENT_STATUS_REGISTERED
	res.RegisteredAt = platforms.GetCurrentEthiopianTime()
	ctx = context.WithValue(ctx, "student", res)
	res, statusCode, er := shan.Service.RegisterStudent(ctx)
	if er != nil || statusCode != state.DT_STATUS_OK {
		if statusCode == state.DT_STATUS_DUPLICATE_PHONE_NUMBER {
			eres.Errs["Phone"] = "Account with this phone number already exist"
			eres.Msg = "Account with this phone number already exist"
			eres.Code = http.StatusConflict
			c.JSON(http.StatusConflict, eres)
		} else {
			eres.Msg = "Internal server problem please try again"
			eres.Code = http.StatusInternalServerError
			c.JSON(http.StatusInternalServerError, eres)
		}
		return
	}
	c.JSON(http.StatusOK, res)
}
func (shan *StudentHandler) UpdateStudent(c *gin.Context) {
	ctx := c.Request.Context()
	eres := &struct {
		Errs       map[string]string `json:"errs"`
		Msg        string            `json:"msg"`
		StatusCode int               `json:"status_code"`
	}{
		Errs: map[string]string{},
	}
	student := &model.Student{}
	jsonDecoder := json.NewDecoder(c.Request.Body)
	if ers := jsonDecoder.Decode(student); ers != nil {
		eres.Msg = "Bad request input"
		c.JSON(http.StatusBadRequest, ers)
		return
	}
	if student.ID <= 0 {
		eres.Msg = "Student with this id doesn't exist"
		c.JSON(http.StatusNotFound, eres)
		return
	}
	ctx = context.WithValue(ctx, "student_id", uint64(student.ID))
	oldstudent, status, er := shan.Service.GetStudentByID(ctx)
	if oldstudent == nil || status != state.DT_STATUS_OK || er != nil {
		if status == state.DT_STATUS_RECORD_NOT_FOUND {
			eres.Msg = "Student with this id doesn't exist"
			eres.StatusCode = http.StatusNotFound
			c.JSON(http.StatusNotFound, eres)
			return
		} else if status == state.DT_STATUS_DBQUERY_ERROR {
			eres.Msg = "Internal problem"
			eres.StatusCode = http.StatusInternalServerError
			c.JSON(http.StatusInternalServerError, eres)
			return
		} else {
			log.Println("Internal Problem in the else statement.")
			eres.Msg = "Internal problem"
			eres.StatusCode = http.StatusInternalServerError
			c.JSON(http.StatusInternalServerError, eres)
			return
		}
	}
	session := ctx.Value("session").(*model.Session)
	if !(session.ID == uint64(oldstudent.RegisteredBy) || (session.Role == state.SUPERADMIN)) {
		eres.Msg = "You are not authorized to update this student information"
		eres.StatusCode = http.StatusUnauthorized
		c.JSON(http.StatusUnauthorized, eres)
		return
	}
	updated := false
	passed := true
	if student.Fullname != "" && student.Fullname != oldstudent.Fullname {
		if len(student.Fullname) < 7 || len(strings.Split(student.Fullname, " ")) < 2 {
			passed = false
			eres.Errs["Full name"] = "Invalid fullname value"
		} else {
			erross := false
			if len(strings.Split(student.Fullname, " ")) >= 2 && len(strings.Split(student.Fullname, " ")[0]) < 3 {
				passed = false
				erross = true
				eres.Errs["First name"] = "Invalid firstname value"
			}
			if len(strings.Split(student.Fullname, " ")) >= 2 && len(strings.Split(student.Fullname, " ")[1]) < 3 {
				passed = false
				erross = true
				eres.Errs["Last name"] = "Invalid lastname value"
			}
			if !erross {
				updated = true
				oldstudent.Fullname = student.Fullname
			}
		}
	}
	if student.Sex != "" {
		if student.Sex == "F" ||
			student.Sex == "f" ||
			student.Sex == "Female" ||
			strings.HasPrefix(student.Sex, "F") ||
			strings.HasPrefix(student.Sex, "f") {
			oldstudent.Sex = "F"
			updated = true
		} else if student.Sex == "M" ||
			student.Sex == "m" ||
			student.Sex == "Male" ||
			strings.HasPrefix(student.Sex, "M") ||
			strings.HasPrefix(student.Sex, "m") {
			oldstudent.Sex = "M"
			updated = true
		} else {
			passed = false
			eres.Errs["sex"] = "invalid sex value"
		}
	}
	if student.AccStatus != "" {
		if student.AccStatus != oldstudent.AccStatus {
			oldstudent.AccStatus = student.AccStatus
			updated = true
		}
	}
	if student.Phone != "" {
		if len(student.Phone) <= 13 && len(student.Phone) >= 10 && form.MatchesPattern(student.Phone, form.PhoneRX) {
			if strings.HasPrefix(student.Phone, "0") {
				student.Phone = strings.Replace(student.Phone, "0", "+251", 1)
			}
			if !(student.Phone == oldstudent.Phone) { // check the existance of a student with this id.
				ctx = context.WithValue(ctx, "student_phone", student.Phone)
				status, er := shan.Service.CheckWhetherTheStudentWithThisPhoneNumberExists(ctx)
				if status == -1 || er != nil {
					eres.StatusCode = http.StatusInternalServerError
					eres.Msg = "internal problem happened"
					c.JSON(http.StatusInternalServerError, eres)
					return
				} else if status >= 1 {
					eres.StatusCode = http.StatusConflict
					eres.Msg = "account with this phone already exist"
					eres.Errs["phone"] = "account with this phone already exist"
					c.JSON(http.StatusConflict, eres)
					return
				}
				if student.Phone != oldstudent.Phone {
					oldstudent.Phone = student.Phone
					updated = true
				}
			}
		} else {
			passed = false
			eres.Errs["phone"] = "invalid phone number value"
		}
	}
	if student.RoundID <= 0 {
		passed = false
		eres.Errs["Round ID"] = "invalid round id"
	} else {
		ctx = context.WithValue(ctx, "round_id", uint64(student.RoundID))
		result, statusCode := shan.RoundService.CheckTheExistanceAndActivenessOfRound(ctx)
		if result == -2 {
			if statusCode == state.DT_STATUS_DBQUERY_ERROR {
				// log.Println("Database Query Error student_handler :: 146")
			}
			eres.Msg = "Internal problem, please try again"
			passed = false
			// count += 1
		} else if result == -1 {
			eres.Msg = "Round with this ID does not exist"
			passed = false
			// count += 1
		} else if result == 0 {
			eres.Msg = "Round is not active for registration"
			passed = false
			// count += 1
		} else {
			oldstudent.RoundID = student.RoundID
			updated = true
		}
	}
	if updated && passed {
		ctx = context.WithValue(ctx, "updated_student", oldstudent)
		newstudent, status, er := shan.Service.UpdateStudent(ctx)
		if er != nil || status != state.DT_STATUS_OK {
			eres.Msg = "internal server problem"
			eres.StatusCode = http.StatusInternalServerError
			c.JSON(http.StatusInternalServerError, eres)
			return
		}
		c.JSON(http.StatusOK, newstudent)
		return
	} else {
		eres.StatusCode = http.StatusNotModified
		if !passed {
			eres.Msg = "invalid update content is included"
		} else if !updated {
			eres.Msg = "no update was made"
		}
		c.JSON(http.StatusNotModified, eres)
		return
	}
}

func (shan *StudentHandler) CreateStudentsSpecialCaseForStudent(c *gin.Context) {
	ctx := c.Request.Context()
	in := &struct {
		StudentID     int     `json:"student_id"`
		Reason        string  `json:"reason"`
		CoveredAmount float64 `json:"covered_amount"`
		CompleteFee   bool    `json:"complete_fee"`
	}{}
	resp := &struct {
		Errs        map[string]string  `json:"errors,omitempty"`
		Msg         string             `json:"msg"`
		StatusCode  int                `json:"status_code"`
		SpecialCase *model.SpecialCase `json:"special_case,omitempty"`
		StudentID   uint64             `json:"student_id"`
	}{
		Errs: map[string]string{},
	}
	jsonDec := json.NewDecoder(c.Request.Body)
	if jecErr := jsonDec.Decode(in); jecErr != nil || in.StudentID <= 0 || in.Reason == "" || (in.CompleteFee == false && in.CoveredAmount == 0) {
		resp.Msg = "bad request input"
		resp.StatusCode = http.StatusBadRequest
		if in.StudentID <= 0 {
			resp.Errs["student_id"] = "invalid student id"
		}
		if in.Reason == "" {
			resp.Errs["reason"] = "missing important information \"reason\""
		}

		if in.CompleteFee == false && in.CoveredAmount == 0 {
			resp.Errs["covered_amount"] = "The amount of money covered for this student is not mentioned"
		}
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	ctx = context.WithValue(ctx, "student_id", uint64(in.StudentID))
	student, code, err := shan.Service.GetStudentByID(ctx)
	if student == nil || code != state.DT_STATUS_OK || err != nil {
		if code == state.DT_STATUS_RECORD_NOT_FOUND {
			resp.Msg = "student with id " + strconv.Itoa(in.StudentID) + " does not exist"
			resp.StatusCode = http.StatusNotFound
			c.JSON(http.StatusNotFound, resp)
			return
		} else {
			resp.Msg = "internal problem happeded"
			resp.StatusCode = http.StatusInternalServerError
			c.JSON(http.StatusInternalServerError, resp)
			return
		}
	}
	if student.Marked != nil && student.Marked.Reason != "" {
		resp.Msg = "Student, " + (strings.Split(student.Fullname, " "))[0] + " already has a special case information.\n please use the update option instead"
		resp.StatusCode = http.StatusConflict
		c.JSON(http.StatusConflict, resp)
		return
	}
	// if the amount is the total amount of the round payment should i set it as complte of not ?
	specialCase := &model.SpecialCase{
		StudentID: uint64(in.StudentID),
		Reason:    in.Reason,
		Amount:    in.CoveredAmount,
		Total:     in.CompleteFee,
	}
	// creating the special case instance ...
	ctx = context.WithValue(ctx, "special_case", specialCase)
	thecase, status, err := shan.Service.CreateSpecialCase(ctx)
	if status != state.DT_STATUS_OK || err != nil {
		resp.Msg = "unable to create a special case information for student " + student.Fullname
		resp.StatusCode = http.StatusInternalServerError
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	resp.SpecialCase = thecase
	resp.StudentID = uint64(in.StudentID)
	resp.Msg = "succesfuly created a special case instance for student " + student.Fullname
	resp.StatusCode = http.StatusOK
	c.JSON(http.StatusOK, resp)
}
func (shan *StudentHandler) UpdateStudentSpecialCaseInformation(c *gin.Context) {
	ctx := c.Request.Context()
	in := &struct {
		ID            int     `json:"id"`
		Reason        string  `json:"reason"`
		CoveredAmount float64 `json:"covered_amount"`
		CompleteFee   bool    `json:"complete_fee"`
	}{}
	resp := &struct {
		Errs        map[string]string  `json:"errors,omitempty"`
		Msg         string             `json:"msg"`
		StatusCode  int                `json:"status_code"`
		SpecialCase *model.SpecialCase `json:"special_case,omitempty"`
	}{
		Errs: map[string]string{},
	}
	jsonDec := json.NewDecoder(c.Request.Body)
	if jecErr := jsonDec.Decode(in); jecErr != nil ||
		in.ID <= 0 ||
		((in.CompleteFee == false && in.CoveredAmount == 0) ||
			(in.CompleteFee == true && in.CoveredAmount > 0)) {
		resp.Msg = "bad request input"
		resp.StatusCode = http.StatusBadRequest
		if in.ID <= 0 {
			resp.Errs["student_id"] = "invalid student id"
		}
		if in.Reason == "" {
			resp.Errs["reason"] = "missing important information \"reason\""
		}
		if in.CompleteFee == false && in.CoveredAmount == 0 {
			resp.Errs["covered_amount"] = "The amount of money covered for this student is not mentioned"
		}
		if in.CompleteFee == true && in.CoveredAmount > 0 {
			resp.Errs["covered_amount"] = "You shouldn't specify the amount of money waived if the total amount of payment is covered"
		}
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	ctx = context.WithValue(ctx, "special_case_id", uint64(in.ID))
	specialCase, st, errs := shan.Service.GetSpecialCaseByID(ctx)
	if errs != nil || st != state.DT_STATUS_OK || specialCase == nil {
		if st == state.DT_STATUS_NO_RECORD_FOUND {
			resp.Msg = "special case instance with this id doesn't exist"
			resp.StatusCode = http.StatusNotFound
			c.JSON(http.StatusNotFound, resp)
			return
		} else {
			resp.Msg = "internal problem, please try again"
			resp.StatusCode = http.StatusInternalServerError
			c.JSON(http.StatusInternalServerError, resp)
			return
		}
	}

	updated := false
	if specialCase.Reason != in.Reason && in.Reason != "" {
		updated = true
		specialCase.Reason = in.Reason
	}
	if in.CompleteFee == specialCase.Total {
		if !(specialCase.Total) {
			if specialCase.Amount != in.CoveredAmount {
				updated = true
				specialCase.Amount = in.CoveredAmount
			}
		} else {
			specialCase.Amount = 0
		}
	}
	if !updated {
		resp.Msg = "now update was made"
		resp.StatusCode = http.StatusNotModified
		c.JSON(http.StatusNotModified, resp)
		return
	}
	// creating the special case instance ...
	ctx = context.WithValue(ctx, "special_case", specialCase)
	status, err := shan.Service.UpdateSpecialCase(ctx)
	if status != state.DT_STATUS_OK || err != nil {
		resp.Msg = "unable to update a special case information "
		resp.StatusCode = http.StatusInternalServerError
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	resp.SpecialCase = specialCase
	resp.Msg = "succesfuly updated a special case instance "
	resp.StatusCode = http.StatusOK
	c.JSON(http.StatusOK, resp)
}

func (shan *StudentHandler) UpdateStudentProfilePicture(c *gin.Context) {
	ctx := c.Request.Context()
	var header *multipart.FileHeader
	var erro error
	var oldImage string

	eres := &struct {
		Msg        string `json:"msg,omitempty"`
		Imgurl     string `json:"imgurl,omitempty"`
		StatusCode int    `json:"status_code,omitempty"`
	}{}

	studentID, er := strconv.Atoi(c.Query("student_id"))
	if er != nil {
		eres.Msg = "bad request, missing \"student_id\" variable of type 'integer'"
		eres.StatusCode = http.StatusBadRequest
		c.JSON(http.StatusBadRequest, eres)
		return
	}
	erro = c.Request.ParseMultipartForm(99999999999)
	if erro != nil {
		eres.Msg = "bad request, missing \"student_id\" variable of type 'integer'"
		eres.StatusCode = http.StatusBadRequest
		c.JSON(http.StatusBadRequest, eres)
		return
	}
	image, header, erro := c.Request.FormFile("image")
	if erro != nil {
		eres.StatusCode = http.StatusBadRequest
		eres.Msg = "missing an image input \"image\""
		c.JSON(http.StatusBadRequest, eres)
		return
	}
	defer image.Close()
	if helper.IsImage(header.Filename) {
		newName := "images/students/" + helper.GenerateRandomString(5, helper.CHARACTERS) + "." + helper.GetExtension(header.Filename)
		var newImage *os.File
		if strings.HasSuffix(os.Getenv("ASSETS_DIRECTORY"), "/") {
			newImage, erro = os.Create(os.Getenv("ASSETS_DIRECTORY") + newName)
		} else {
			newImage, erro = os.Create(os.Getenv("ASSETS_DIRECTORY") + "/" + newName)
		}
		if erro != nil {
			log.Println(erro.Error())
			eres.StatusCode = http.StatusInternalServerError
			eres.Msg = "internal problem, please try again"
			c.JSON(http.StatusInternalServerError, eres)
			return
		}
		defer newImage.Close()
		ctx = context.WithValue(ctx, "student_id", studentID)
		var statusC uint8
		oldImage, statusC = shan.Service.GetStudentImageUrl(ctx)
		if statusC != http.StatusOK {
			log.Println()
		}
		_, er := io.Copy(newImage, image)
		if er != nil {
			log.Println(erro.Error())
			eres.Msg = "internal problem, please try again"
			eres.StatusCode = http.StatusInternalServerError
			c.JSON(http.StatusInternalServerError, eres)
			return
		}
		ncon := context.WithValue(c.Request.Context(), "student_id", uint64(studentID))
		ncon = context.WithValue(ncon, "image_url", string(newName))
		success := shan.Service.ChangeStudentImageUrl(ncon)
		if success {
			if oldImage != "" {
				if strings.HasSuffix(os.Getenv("ASSETS_DIRECTORY"), "/") {
					er = os.Remove(os.Getenv("ASSETS_DIRECTORY") + oldImage)
				} else {
					er = os.Remove(os.Getenv("ASSETS_DIRECTORY") + "/" + oldImage)
				}
			}
			eres.Msg = "image succesfuly uploaded"
			eres.StatusCode = http.StatusOK
			eres.Imgurl = newName
			c.JSON(http.StatusOK, eres)
			return
		}
		if strings.HasSuffix(os.Getenv("ASSETS_DIRECTORY"), "/") {
			er = os.Remove(os.Getenv("ASSETS_DIRECTORY") + newName)
		} else {
			er = os.Remove(os.Getenv("ASSETS_DIRECTORY") + "/" + newName)
		}
		eres.Msg = "internal problem, please try again"
		eres.StatusCode = http.StatusInternalServerError
		c.JSON(http.StatusInternalServerError, eres)
	} else {
		eres.Msg = "only images with type \"jpeg\", \"png\", \"jpg\", \"gif\", and \"btmp\" are accepted"
		eres.StatusCode = http.StatusUnsupportedMediaType
		c.JSON(http.StatusUnsupportedMediaType, eres)
	}
}
func (shan *StudentHandler) UpdateStudentBirthDate(*gin.Context) {}
func (shan *StudentHandler) UpdateStudentAddress(c *gin.Context) {}

func (shan *StudentHandler) DeleteStudent(c *gin.Context) {}
func (shan *StudentHandler) GetStudentByID(c *gin.Context) {
	ctx := c.Request.Context()
	studentID, ef := strconv.Atoi(c.Query("student_id"))
	eres := &struct {
		Msg        string         `json:"msg,omitempty"`
		StatusCode int            `json:"status_code"`
		Student    *model.Student `json:"student,omitempty"`
	}{}
	if ef != nil || studentID <= 0 {
		if c.Query("student_id") == "" {
			eres.Msg = " query value \"student_id\" of type integer , is not specified"
		} else if studentID <= 0 {
			eres.Msg = "student with this id does not exist"
		} else {
			eres.Msg = "bad query"
		}
		eres.StatusCode = http.StatusBadRequest
		c.JSON(http.StatusBadRequest, eres)
		return
	}
	ctx = context.WithValue(ctx, "student_id", uint64(studentID))
	student, code, er := shan.Service.GetStudentByID(ctx)
	if code != state.DT_STATUS_OK || student == nil || er != nil {
		eres.Msg = "student with this id does not exist"
		eres.StatusCode = http.StatusNotFound
		c.JSON(http.StatusNotFound, eres)
		return
	}
	eres.Student = student
	eres.StatusCode = http.StatusOK
	c.JSON(http.StatusOK, eres)
}
func (shan *StudentHandler) GetStudentsStatus(c *gin.Context) {}

var serchFiltersCombination = []string{
	"Name",
	"Phone",
	"Name and Phone",
	"Unique Address ",
	"Unique Address and Name",
	"Unique Address and Phone",
	"Unique Address , name , and Phone ",
	"age",
	"age and name",
	"age and phone",
	"age, name, and phone",
	"age and unique address",
	"age, name, and unique address",
	"age, phone, and unique address",
	"age, name, phone, and unique address",
}

func (shan *StudentHandler) SearchStudents(c *gin.Context) {
	// search by name
	// search by phone number
	// search by unique address
	// search by age
	// seatch by round number
	ctx := c.Request.Context()
	println(ctx)

}
func (shan *StudentHandler) GetStudentStatus(c *gin.Context) {}

func (shan *StudentHandler) GetStudentsOfRound(c *gin.Context) {
	ctx := c.Request.Context()
	roundID, er := strconv.Atoi(c.Query("round_id"))
	offset, era := strconv.Atoi(c.Query("offset"))
	if era != nil {
		offset = 0
	}
	limit, erb := strconv.Atoi(c.Query("limit"))
	if erb != nil {
		limit = offset + 5
	}
	res := &struct {
		StatusCode int               `json:"status_code"`
		Msg        string            `json:"msg"`
		Students   []*model.Student  `json:"students"`
		Errs       map[string]string `json:"errs"`
	}{Errs: map[string]string{}}
	if er != nil || roundID <= 0 {
		res.StatusCode = http.StatusBadRequest
		res.Msg = "invalid round id"
		res.Errs["round_id"] = "missing important parameter \"round_id\" type integer"
		c.JSON(http.StatusBadRequest, res)
		return
	}
	ctx = context.WithValue(ctx, "round_id", uint64(roundID))
	statk, _ := shan.RoundService.CheckTheExistanceAndActivenessOfRound(ctx)
	if statk == -2 {
		res.Msg = "internal problem"
		res.StatusCode = http.StatusInternalServerError
		c.JSON(http.StatusInternalServerError, res)
		return
	} else if statk == -1 {
		res.StatusCode = http.StatusNotFound
		res.Msg = "round with this id doesn't exist"
		c.JSON(http.StatusNotFound, res)
		return
	}
	ctx = context.WithValue(ctx, "round_id", uint(roundID))
	ctx = context.WithValue(ctx, "offset", uint(offset))
	ctx = context.WithValue(ctx, "limit", uint(limit))
	students, status, er := shan.Service.GetStudentsOfRound(ctx)
	if status != state.DT_STATUS_OK {
		if status == state.DT_STATUS_RECORD_NOT_FOUND {
			res.StatusCode = http.StatusNotFound
			res.Students = students
			res.Msg = " no student records found"
			c.JSON(http.StatusNotFound, res)
			return
		} else {
			res.StatusCode = http.StatusInternalServerError
			res.Msg = "internal problem had happeded, please try again"
			c.JSON(http.StatusInternalServerError, res)
			return
		}
	}
	res.Students = students
	res.StatusCode = http.StatusOK
	c.JSON(http.StatusOK, res)
}

func (shan *StudentHandler) DeleteStudentByID(c *gin.Context) {
	ctx := c.Request.Context()
	println(ctx)

}

func (shan *StudentHandler) GetStudentsOfCategory(c *gin.Context) {
	ctx := c.Request.Context()
	categoryID, er := strconv.Atoi(c.Query("category_id"))
	offset, era := strconv.Atoi(c.Query("offset"))
	if era != nil {
		offset = 0
	}
	limit, erb := strconv.Atoi(c.Query("limit"))
	if erb != nil {
		limit = offset + 5
	}
	res := &struct {
		StatusCode int               `json:"status_code"`
		Msg        string            `json:"msg"`
		Students   []*model.Student  `json:"students"`
		Errs       map[string]string `json:"errs"`
	}{Errs: map[string]string{}}
	if er != nil || categoryID <= 0 {
		res.StatusCode = http.StatusBadRequest
		res.Msg = "invalid category id"
		res.Errs["category_id"] = "missing important parameter \"category_id\" type positive integer"
		c.JSON(http.StatusBadRequest, res)
		return
	}
	ctx = context.WithValue(ctx, "category_id", uint64(categoryID))
	statk, _ := shan.RoundService.CheckTheExistanceOfCategory(ctx)
	if statk == -2 {
		res.Msg = "internal problem"
		res.StatusCode = http.StatusInternalServerError
		c.JSON(http.StatusInternalServerError, res)
		return
	} else if statk == -1 {
		res.StatusCode = http.StatusNotFound
		res.Msg = "round with this id doesn't exist"
		c.JSON(http.StatusNotFound, res)
		return
	}
	ctx = context.WithValue(ctx, "category_id", uint(categoryID))
	ctx = context.WithValue(ctx, "offset", uint(offset))
	ctx = context.WithValue(ctx, "limit", uint(limit))
	students, status, er := shan.Service.GetStudentsOfCategory(ctx)
	if status != state.DT_STATUS_OK {
		if status == state.DT_STATUS_RECORD_NOT_FOUND {
			res.StatusCode = http.StatusNotFound
			res.Students = students
			res.Msg = " no student records found"
			c.JSON(http.StatusNotFound, res)
			return
		} else {
			res.StatusCode = http.StatusInternalServerError
			res.Msg = "internal problem had happeded, please try again"
			c.JSON(http.StatusInternalServerError, res)
			return
		}
	}
	res.Students = students
	res.StatusCode = http.StatusOK
	c.JSON(http.StatusOK, res)
}
