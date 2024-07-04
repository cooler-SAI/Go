package testing

import (
	"testing"
)

/*func TestSimpleTestFunc(t *testing.T) {
	expectedSimpleTestResult := "good"
	if actualSimpleTestResult := SimpleTest(); actualSimpleTestResult != expectedSimpleTestResult {
		t.Errorf("expected %s; got: %s", expectedSimpleTestResult, actualSimpleTestResult)
	}
}*/

func TestSimpleTest2Func(t *testing.T) {
	expectedSimpleTestResult := 2
	if actualSimpleTestResult := SimpleTest2(); actualSimpleTestResult != expectedSimpleTestResult {
		t.Errorf("expected %d; got: %d", expectedSimpleTestResult, actualSimpleTestResult)
	}
}
