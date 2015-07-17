// Package httpmatchers provides a set of custom matchers for
// github.com/onsi/gomega to use for testing responses from HTTP server
// handlers.
//
// Specifically it allows you to replace code like this
//
// 	Expect(w.Code).To(Equal(http.StatusOK), "HTTP Status should be OK")
//
// Which yields an error message like
//
// 	HTTP Status should be OK
// 	Expected
// 		<int>: 400
// 	to equal
// 		<int>: 200
//
// With this code
//
// 	Expect(w.Code).To(BeHTTPStatusOK())
//
// Which yields an error message like this
//
// 	Expected
// 		<string>: HTTP status 400 Bad Request
// 	to be
// 		<string>: HTTP status 200 OK
package httpmatchers
