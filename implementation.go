package httpmatchers

import (
	"fmt"
	"net/http"

	"github.com/onsi/gomega/format"
)

type httpStatusMatcher struct {
	code int
}

func (m *httpStatusMatcher) Match(actual interface{}) (bool, error) {
	code, ok := actual.(int)
	if !ok {
		return false, fmt.Errorf("BeHTTPStatus matchers expect an int")
	}

	return code == m.code, nil
}

func (m *httpStatusMatcher) FailureMessage(actual interface{}) string {
	return format.Message(statusString(actual.(int)), "to be", statusString(m.code))
}

func (m *httpStatusMatcher) NegatedFailureMessage(actual interface{}) string {
	return format.Message(statusString(actual.(int)), "not to be", statusString(m.code))
}

func statusString(code int) string {
	return fmt.Sprintf("HTTP status %d %s", code, http.StatusText(code))
}
