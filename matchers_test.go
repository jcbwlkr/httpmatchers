package httpmatchers

import (
	"testing"

	"github.com/onsi/gomega/types"
)

type matchTest struct {
	matchFn     func() types.GomegaMatcher
	passingCode int
	failingCode int
}

var matchTests = []matchTest{
	{BeHTTPStatusOK, 200, 404},
}

func TestMatch(t *testing.T) {
	for i, test := range matchTests {
		matcher := test.matchFn()

		passSuccess, err := matcher.Match(test.passingCode)
		if !passSuccess {
			t.Errorf("Test %d: Match(%d): Expected true, actual false", i, test.passingCode)
		}
		if err != nil {
			t.Errorf("Test %d: Match(%d): Expected err to be nil, was %v", i, test.passingCode, err)
		}

		failSuccess, err := matcher.Match(test.failingCode)
		if failSuccess {
			t.Errorf("Test %d: Match(%d): Expected false, actual true", i, test.failingCode)
		}
		if err != nil {
			t.Errorf("Test %d: Match(%d): Expected err to be nil, was %v", i, test.failingCode, err)
		}

		invalidMatch, err := matcher.Match("banana")
		if invalidMatch {
			t.Errorf("Test %d: Match(%q): Expected false, actual true", i, "banana")
		}
		if err == nil {
			t.Errorf("Test %d: Match(%q): Expected err to not be nil", i, "banana")
		}
	}
}
