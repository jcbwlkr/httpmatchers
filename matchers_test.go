package httpmatchers

import (
	"testing"

	"github.com/onsi/gomega/types"
)

type matchTest struct {
	matchFn     func() types.GomegaMatcher
	passingCode int
}

var matchTests = []matchTest{
	{BeHTTPStatusContinue, 100},
	{BeHTTPStatusSwitchingProtocols, 101},

	{BeHTTPStatusOK, 200},
	{BeHTTPStatusCreated, 201},
	{BeHTTPStatusAccepted, 202},
	{BeHTTPStatusNonAuthoritativeInfo, 203},
	{BeHTTPStatusNoContent, 204},
	{BeHTTPStatusResetContent, 205},
	{BeHTTPStatusPartialContent, 206},

	{BeHTTPStatusMultipleChoices, 300},
	{BeHTTPStatusMovedPermanently, 301},
	{BeHTTPStatusFound, 302},
	{BeHTTPStatusSeeOther, 303},
	{BeHTTPStatusNotModified, 304},
	{BeHTTPStatusUseProxy, 305},
	{BeHTTPStatusTemporaryRedirect, 307},

	{BeHTTPStatusBadRequest, 400},
	{BeHTTPStatusUnauthorized, 401},
	{BeHTTPStatusPaymentRequired, 402},
	{BeHTTPStatusForbidden, 403},
	{BeHTTPStatusNotFound, 404},
	{BeHTTPStatusMethodNotAllowed, 405},
	{BeHTTPStatusNotAcceptable, 406},
	{BeHTTPStatusProxyAuthRequired, 407},
	{BeHTTPStatusRequestTimeout, 408},
	{BeHTTPStatusConflict, 409},
	{BeHTTPStatusGone, 410},
	{BeHTTPStatusLengthRequired, 411},
	{BeHTTPStatusPreconditionFailed, 412},
	{BeHTTPStatusRequestEntityTooLarge, 413},
	{BeHTTPStatusRequestURITooLong, 414},
	{BeHTTPStatusUnsupportedMediaType, 415},
	{BeHTTPStatusRequestedRangeNotSatisfiable, 416},
	{BeHTTPStatusExpectationFailed, 417},
	{BeHTTPStatusTeapot, 418},

	{BeHTTPStatusInternalServerError, 500},
	{BeHTTPStatusNotImplemented, 501},
	{BeHTTPStatusBadGateway, 502},
	{BeHTTPStatusServiceUnavailable, 503},
	{BeHTTPStatusGatewayTimeout, 504},
	{BeHTTPStatusHTTPVersionNotSupported, 505},
}

func TestMatch(t *testing.T) {
	failCode := 0
	invalidCode := "banana"

	for i, test := range matchTests {
		matcher := test.matchFn()

		passSuccess, err := matcher.Match(test.passingCode)
		if !passSuccess {
			t.Errorf("Test %d: Match(%d): Expected true, actual false", i, test.passingCode)
		}
		if err != nil {
			t.Errorf("Test %d: Match(%d): Expected err to be nil, was %v", i, test.passingCode, err)
		}

		failSuccess, err := matcher.Match(failCode)
		if failSuccess {
			t.Errorf("Test %d: Match(%d): Expected false, actual true", i, failCode)
		}
		if err != nil {
			t.Errorf("Test %d: Match(%d): Expected err to be nil, was %v", i, failCode, err)
		}

		invalidMatch, err := matcher.Match(invalidCode)
		if invalidMatch {
			t.Errorf("Test %d: Match(%q): Expected false, actual true", i, invalidCode)
		}
		if err == nil {
			t.Errorf("Test %d: Match(%q): Expected err to not be nil", i, invalidCode)
		}
	}
}
