package testing

import (
	"testing"
)

func TestSimpleTestFunc(t *testing.T) {
	expectedSimpleTestResult := "good"
	if actualSimpleTestResult := SimpleTest(); actualSimpleTestResult != expectedSimpleTestResult {
		t.Errorf("expected %s; got: %s", expectedSimpleTestResult, actualSimpleTestResult)
	}
}
