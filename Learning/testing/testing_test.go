package testing

import (
	"testing"
)

func TestSimpleTestFunc(t *testing.T) {
	expectedFooResult := "bar"
	if actualFooResult := SimpleTest(); actualFooResult != expectedFooResult {
		t.Errorf("expected %s; got: %s", expectedFooResult, actualFooResult)
	}
}
