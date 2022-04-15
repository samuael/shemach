package state

const (
	STATUS_OK = iota
	STATUS_RECORD_NOT_FOUND
	STATUS_DBQUERY_ERROR
	STATUS_NO_RECORD_UPDATED
	STATUS_NO_RECORD_FOUND
	STATUS_DELETION_FAILED
	STATUS_NO_ROW_DELETED
	STATUS_RECORD_ALREADY_EXIST
	STATUS_INCOMPLETE_DATA
	STATUS_DUPLICATE_PHONE_NUMBER
	STATUS_DUPLICATE_EMAIL_ADDRESS
	STATUS_INVALID_PHONE_NUMBER_VALUE
	STATUS_DUPLICATE_INVALID_EMAIL_ADDRESS
	STATUS_DUPLICATE_RECORD
)

var (
	STATUS = map[int]string{
		STATUS_RECORD_NOT_FOUND:                "status record not found",
		STATUS_DBQUERY_ERROR:                   "status database query error",
		STATUS_OK:                              "db query ok",
		STATUS_NO_RECORD_UPDATED:               " no record was updated",
		STATUS_NO_RECORD_FOUND:                 " no record found",
		STATUS_RECORD_ALREADY_EXIST:            "record already exist",
		STATUS_INCOMPLETE_DATA:                 "querying incomplete data",
		STATUS_DUPLICATE_EMAIL_ADDRESS:         "duplicate email address",
		STATUS_DUPLICATE_PHONE_NUMBER:          "duplicate phome number",
		STATUS_INVALID_PHONE_NUMBER_VALUE:      "invalid phone number value",
		STATUS_DUPLICATE_INVALID_EMAIL_ADDRESS: "invalid email address value",
		STATUS_DUPLICATE_RECORD:                "DUPLICATE RECODS",
	}
)
