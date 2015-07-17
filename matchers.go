package httpmatchers

import (
	"net/http"

	"github.com/onsi/gomega/types"
)

// BeHTTPStatusOK matches the HTTP status code 200 OK
func BeHTTPStatusOK() types.GomegaMatcher {
	return &httpStatusMatcher{code: http.StatusOK}
}
