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
	{BeHTTPStatusOK, 200},
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
