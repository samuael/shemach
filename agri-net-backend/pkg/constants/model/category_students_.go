package model

type CategoryStudentCountReponse struct {
	CategoryID     uint `json:"category_id"`
	AllStudents    uint `json:"students_count"`
	ActiveStudents uint `json:"active_students"`
}
