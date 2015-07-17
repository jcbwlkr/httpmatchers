package httpmatchers

import (
	"net/http"

	"github.com/onsi/gomega/types"
)

// BeHTTPStatusContinue matches the HTTP status code 100 Continue
func BeHTTPStatusContinue() types.GomegaMatcher {
	return &httpStatusMatcher{code: http.StatusContinue}
}

// BeHTTPStatusSwitchingProtocols matches the HTTP status code 101 Switching
// Protocols
func BeHTTPStatusSwitchingProtocols() types.GomegaMatcher {
	return &httpStatusMatcher{code: http.StatusSwitchingProtocols}
}

// BeHTTPStatusOK matches the HTTP status code 200 OK
func BeHTTPStatusOK() types.GomegaMatcher {
	return &httpStatusMatcher{code: http.StatusOK}
}

// BeHTTPStatusCreated matches the HTTP status code 201 Created
func BeHTTPStatusCreated() types.GomegaMatcher {
	return &httpStatusMatcher{code: http.StatusCreated}
}

// BeHTTPStatusAccepted matches the HTTP status code 202 Accepted
func BeHTTPStatusAccepted() types.GomegaMatcher {
	return &httpStatusMatcher{code: http.StatusAccepted}
}

// BeHTTPStatusNonAuthoritativeInfo matches the HTTP status code 203 Non
// Authoritative Info
func BeHTTPStatusNonAuthoritativeInfo() types.GomegaMatcher {
	return &httpStatusMatcher{code: http.StatusNonAuthoritativeInfo}
}

// BeHTTPStatusNoContent matches the HTTP status code 204 No Content
func BeHTTPStatusNoContent() types.GomegaMatcher {
	return &httpStatusMatcher{code: http.StatusNoContent}
}

// BeHTTPStatusResetContent matches the HTTP status code 205 Reset Content
func BeHTTPStatusResetContent() types.GomegaMatcher {
	return &httpStatusMatcher{code: http.StatusResetContent}
}

// BeHTTPStatusPartialContent matches the HTTP status code 206 Partial Content
func BeHTTPStatusPartialContent() types.GomegaMatcher {
	return &httpStatusMatcher{code: http.StatusPartialContent}
}

// BeHTTPStatusMultipleChoices matches the HTTP status code 300 Multiple
// Choices
func BeHTTPStatusMultipleChoices() types.GomegaMatcher {
	return &httpStatusMatcher{code: http.StatusMultipleChoices}
}

// BeHTTPStatusMovedPermanently matches the HTTP status code 301 Moved
// Permanently
func BeHTTPStatusMovedPermanently() types.GomegaMatcher {
	return &httpStatusMatcher{code: http.StatusMovedPermanently}
}

// BeHTTPStatusFound matches the HTTP status code 302 Found
func BeHTTPStatusFound() types.GomegaMatcher {
	return &httpStatusMatcher{code: http.StatusFound}
}

// BeHTTPStatusSeeOther matches the HTTP status code 303 See Other
func BeHTTPStatusSeeOther() types.GomegaMatcher {
	return &httpStatusMatcher{code: http.StatusSeeOther}
}

// BeHTTPStatusNotModified matches the HTTP status code 304 Not Modified
func BeHTTPStatusNotModified() types.GomegaMatcher {
	return &httpStatusMatcher{code: http.StatusNotModified}
}

// BeHTTPStatusUseProxy matches the HTTP status code 305 Use Proxy
func BeHTTPStatusUseProxy() types.GomegaMatcher {
	return &httpStatusMatcher{code: http.StatusUseProxy}
}

// BeHTTPStatusTemporaryRedirect matches the HTTP status code 307 Temporary
// Redirect
func BeHTTPStatusTemporaryRedirect() types.GomegaMatcher {
	return &httpStatusMatcher{code: http.StatusTemporaryRedirect}
}

// BeHTTPStatusBadRequest matches the HTTP status code 400 Bad Request
func BeHTTPStatusBadRequest() types.GomegaMatcher {
	return &httpStatusMatcher{code: http.StatusBadRequest}
}

// BeHTTPStatusUnauthorized matches the HTTP status code 401 Unauthorized
func BeHTTPStatusUnauthorized() types.GomegaMatcher {
	return &httpStatusMatcher{code: http.StatusUnauthorized}
}

// BeHTTPStatusPaymentRequired matches the HTTP status code 402 Payment
// Required
func BeHTTPStatusPaymentRequired() types.GomegaMatcher {
	return &httpStatusMatcher{code: http.StatusPaymentRequired}
}

// BeHTTPStatusForbidden matches the HTTP status code 403 Forbidden
func BeHTTPStatusForbidden() types.GomegaMatcher {
	return &httpStatusMatcher{code: http.StatusForbidden}
}

// BeHTTPStatusNotFound matches the HTTP status code 404 Not Found
func BeHTTPStatusNotFound() types.GomegaMatcher {
	return &httpStatusMatcher{code: http.StatusNotFound}
}

// BeHTTPStatusMethodNotAllowed matches the HTTP status code 405 Method Not
// Allowed
func BeHTTPStatusMethodNotAllowed() types.GomegaMatcher {
	return &httpStatusMatcher{code: http.StatusMethodNotAllowed}
}

// BeHTTPStatusNotAcceptable matches the HTTP status code 406 Not Acceptable
func BeHTTPStatusNotAcceptable() types.GomegaMatcher {
	return &httpStatusMatcher{code: http.StatusNotAcceptable}
}

// BeHTTPStatusProxyAuthRequired matches the HTTP status code 407 Proxy Auth
// Required
func BeHTTPStatusProxyAuthRequired() types.GomegaMatcher {
	return &httpStatusMatcher{code: http.StatusProxyAuthRequired}
}

// BeHTTPStatusRequestTimeout matches the HTTP status code 408 Request Timeout
func BeHTTPStatusRequestTimeout() types.GomegaMatcher {
	return &httpStatusMatcher{code: http.StatusRequestTimeout}
}

// BeHTTPStatusConflict matches the HTTP status code 409 Conflict
func BeHTTPStatusConflict() types.GomegaMatcher {
	return &httpStatusMatcher{code: http.StatusConflict}
}

// BeHTTPStatusGone matches the HTTP status code 410 Gone
func BeHTTPStatusGone() types.GomegaMatcher {
	return &httpStatusMatcher{code: http.StatusGone}
}

// BeHTTPStatusLengthRequired matches the HTTP status code 411 Length Required
func BeHTTPStatusLengthRequired() types.GomegaMatcher {
	return &httpStatusMatcher{code: http.StatusLengthRequired}
}

// BeHTTPStatusPreconditionFailed matches the HTTP status code 412 Precondition
// Failed
func BeHTTPStatusPreconditionFailed() types.GomegaMatcher {
	return &httpStatusMatcher{code: http.StatusPreconditionFailed}
}

// BeHTTPStatusRequestEntityTooLarge matches the HTTP status code 413 Request
// Entity Too Large
func BeHTTPStatusRequestEntityTooLarge() types.GomegaMatcher {
	return &httpStatusMatcher{code: http.StatusRequestEntityTooLarge}
}

// BeHTTPStatusRequestURITooLong matches the HTTP status code 414 Request
// URIToo Long
func BeHTTPStatusRequestURITooLong() types.GomegaMatcher {
	return &httpStatusMatcher{code: http.StatusRequestURITooLong}
}

// BeHTTPStatusUnsupportedMediaType matches the HTTP status code 415
// Unsupported Media Type
func BeHTTPStatusUnsupportedMediaType() types.GomegaMatcher {
	return &httpStatusMatcher{code: http.StatusUnsupportedMediaType}
}

// BeHTTPStatusRequestedRangeNotSatisfiable matches the HTTP status code 416
// Requested Range Not Satisfiable
func BeHTTPStatusRequestedRangeNotSatisfiable() types.GomegaMatcher {
	return &httpStatusMatcher{code: http.StatusRequestedRangeNotSatisfiable}
}

// BeHTTPStatusExpectationFailed matches the HTTP status code 417 Expectation
// Failed
func BeHTTPStatusExpectationFailed() types.GomegaMatcher {
	return &httpStatusMatcher{code: http.StatusExpectationFailed}
}

// BeHTTPStatusTeapot matches the HTTP status code 418 Teapot
func BeHTTPStatusTeapot() types.GomegaMatcher {
	return &httpStatusMatcher{code: http.StatusTeapot}
}

// BeHTTPStatusInternalServerError matches the HTTP status code 500 Internal
// Server Error
func BeHTTPStatusInternalServerError() types.GomegaMatcher {
	return &httpStatusMatcher{code: http.StatusInternalServerError}
}

// BeHTTPStatusNotImplemented matches the HTTP status code 501 Not Implemented
func BeHTTPStatusNotImplemented() types.GomegaMatcher {
	return &httpStatusMatcher{code: http.StatusNotImplemented}
}

// BeHTTPStatusBadGateway matches the HTTP status code 502 Bad Gateway
func BeHTTPStatusBadGateway() types.GomegaMatcher {
	return &httpStatusMatcher{code: http.StatusBadGateway}
}

// BeHTTPStatusServiceUnavailable matches the HTTP status code 503 Service
// Unavailable
func BeHTTPStatusServiceUnavailable() types.GomegaMatcher {
	return &httpStatusMatcher{code: http.StatusServiceUnavailable}
}

// BeHTTPStatusGatewayTimeout matches the HTTP status code 504 Gateway Timeout
func BeHTTPStatusGatewayTimeout() types.GomegaMatcher {
	return &httpStatusMatcher{code: http.StatusGatewayTimeout}
}

// BeHTTPStatusHTTPVersionNotSupported matches the HTTP status code 505
// HTTP Version Not Supported
func BeHTTPStatusHTTPVersionNotSupported() types.GomegaMatcher {
	return &httpStatusMatcher{code: http.StatusHTTPVersionNotSupported}
}
