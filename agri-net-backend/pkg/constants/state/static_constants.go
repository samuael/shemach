package state

var (
	// ImageExtensions list of valid image extensions
	ImageExtensions = []string{"jpeg", "png", "jpg", "gif", "btmp"}
)

const (
	// InvalidUsernameORPassword ...
	InvalidUsernameORPassword = " Invalid Username or Password "
	// InvalidEmailOrPassword
	InvalidEmailOrPassword = " Invalid email or password! "
	// SuccesfulyLoggedIn ...
	SuccesfulyLoggedIn = " Succesfuly Logged In "

	// This constant values listed below are roles of the system
	// those who directly interact with the system .

	LANGUAGE_ALL = "all"

	SUPERADMIN = "superadmin"
	// SUPERADMIN
	SUBSCRIBER = "subscriber"
	// ADMIN ...
	ADMIN = "admin"
	// INFO_ADMIN = "infoadmin"
	INFO_ADMIN = "infoadmin"
	// MERCHANT ...
	MERCHANT = "merchant"
	// AGENT
	AGENT = "agent"

	// // ROUND
	// ROUND = "round"
	// // CATEGORY
	// CATEGORY = "category"
	// HOST
	HOST = "http://192.168.1.7:8080"
)

// INTEGER REPRESENTATION of CXP(commodity exchange participant) ROLES.
const (
	ROLE_ALL = iota
	SUPERADMIN_ROLE_INT
	ADMIN_ROLE_INT
	INFOADMIN_ROLE_INT
	AGENT_ROLE_INT
	MERCHANT_ROLE_INT
	ALL_ROLES_BINARY = 31
)

const (
	POST_IMAGES_RELATIVE_PATH         = "post/image/"
	BLURRED_POST_IMAGES_RELATIVE_PATH = "post/image/blurred/"
)

const (
	PAYMENT_STATUS_CREATED = iota
	PAYMENT_STATUS_SEEN
	PAYMENT_STATUS_REJECTED
	PAYMENT_STATUS_ACCEPTED
)

const (
	STUDENT_STATUS_REGISTERED = iota
	STUDENT_STATUS_PAID
	STUDENT_STATUS_COMPLETED
)

var MAP_STUDENT_STATUS = map[int]string{
	STUDENT_STATUS_REGISTERED: "registered",
	STUDENT_STATUS_PAID:       "paid",
	STUDENT_STATUS_COMPLETED:  "completed",
}

var MAP_PAYMENT_STATUS = map[int]string{
	PAYMENT_STATUS_CREATED:  "created",
	PAYMENT_STATUS_SEEN:     "seen",
	PAYMENT_STATUS_ACCEPTED: "accepted",
	PAYMENT_STATUS_REJECTED: "rejected",
}
