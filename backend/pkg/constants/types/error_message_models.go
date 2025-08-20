package types

import "net/http"

// StatusMsg represents a status message detail
type StatusMsg struct {
	Success   bool        `json:"success,omitempty"`
	Message   string      `json:"msg,omitempty"`
	Error     string      `json:"err,omitempty"`
	Errors    []string    `json:"errors,omitempty"`
	Status    int64       `json:"status"`
	AuthToken string      `json:"auth_token,omitempty"`
	Data      interface{} `json:"data"`
}

// OTPExpirationInfo holds otp short code confirmation start-time and expiration timestamp information.
type OTPExpirationInfo struct {
	StartTimestamp        int64 `json:"start_timestamp"`      // starting timestamp in seconds
	ExpMinutes            int64 `json:"exp_minutes"`          // expiration minutes
	ResendDurationSeconds int64 `json:"resend_interval_secs"` // number of seconds until to request for a short code resend.
}

const (
	ScOK        = 200
	ScCreated   = http.StatusCreated
	ScIncorrect = 200 + iota
	ScFailed
	ScExpired
	ScTrialExceeded
	ScNotFound             = http.StatusNotFound
	ScConflict             = http.StatusConflict
	ScBadRequest           = http.StatusBadRequest
	ScInternalError        = http.StatusInternalServerError
	ScUnauthorized         = http.StatusUnauthorized
	ScPartialContent       = http.StatusPartialContent
	ScUnsupportedMediaType = http.StatusUnsupportedMediaType
	ScTooManyRequest       = http.StatusTooManyRequests

	ScFound = http.StatusFound

	// Specific Error Types
	ScMissingPropertyTypeInformation = 4000
	ScMissingFacilitySupport         = 4001
	ScInvalidFacilityIDProvided      = 4002

	ScMissingServiceSupport     = 4003
	ScInvalidServiceIDProvided  = 4004
	ScInvalidRoomTypeIDProvided = 4005

	ScMissingCompoundID         = 4006
	ScInvalidPropertyStatusInfo = 4007
	ScInvalidPropertyName       = 4008

	ScRoomsCreationError  = 4009
	ScMissingPropertySize = 4010
)

// Page requirement constants.
const (
	PAGE_REQUIRED       = 0
	PAGE_CAN_BE_SKIPPED = 1
	PAGE_NOT_REQUIRED   = 2
)
