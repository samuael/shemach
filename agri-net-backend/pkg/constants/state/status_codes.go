package state

const (
	DT_STATUS_RECORD_NOT_FOUND = iota
	DT_STATUS_DBQUERY_ERROR
	DT_STATUS_OK
	DT_STATUS_NO_RECORD_UPDATED
	DT_STATUS_NO_RECORD_FOUND
	DT_STATUS_DELETION_FAILED
	DT_STATUS_NO_ROW_DELETED
	DT_STATUS_RECORD_ALREADY_EXIST
	DT_STATUS_INCOMPLETE_DATA
	DT_STATUS_DUPLICATE_PHONE_NUMBER
	DT_STATUS_DUPLICATE_EMAIL_ADDRESS
	DT_STATUS_INVALID_PHONE_NUMBER_VALUE
	DT_STATUS_DUPLICATE_INVALID_EMAIL_ADDRESS
)

var (
	STATUS = map[int]string{
		DT_STATUS_RECORD_NOT_FOUND:                "status record not found",
		DT_STATUS_DBQUERY_ERROR:                   "status database query error",
		DT_STATUS_OK:                              "db query ok",
		DT_STATUS_NO_RECORD_UPDATED:               " no record was updated",
		DT_STATUS_NO_RECORD_FOUND:                 " no record found",
		DT_STATUS_RECORD_ALREADY_EXIST:            "record already exist",
		DT_STATUS_INCOMPLETE_DATA:                 "querying incomplete data",
		DT_STATUS_DUPLICATE_EMAIL_ADDRESS:         "duplicate email address",
		DT_STATUS_DUPLICATE_PHONE_NUMBER:          "duplicate phome number",
		DT_STATUS_INVALID_PHONE_NUMBER_VALUE:      "invalid phone number value",
		DT_STATUS_DUPLICATE_INVALID_EMAIL_ADDRESS: "invalid email address value",
	}
)
