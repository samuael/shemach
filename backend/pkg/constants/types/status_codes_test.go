package types

import "testing"

func TestStatusCodes(t *testing.T) {
	t.Parallel()
	var statsCodes = []int{ScOK, ScCreated, ScIncorrect, ScFailed, ScExpired, ScTrialExceeded, ScNotFound, ScConflict, ScBadRequest, ScInternalError}
	for _, a := range statsCodes {
		println(a)
	}
}
