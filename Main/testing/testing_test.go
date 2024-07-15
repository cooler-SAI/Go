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

func TestSimpleTest2Func(t *testing.T) {
	expectedSimpleTestResult := 2
	if actualSimpleTestResult := SimpleTest2(); actualSimpleTestResult != expectedSimpleTestResult {
		t.Errorf("expected %d; got: %d", expectedSimpleTestResult, actualSimpleTestResult)
	}
}

func TestAdd(t *testing.T) {
	sum, err := Add(1, 2)
	if err != nil {
		t.Error("unexpected error")
	}
	if sum != 3 {
		t.Errorf("sum expected to be 3; got %d", sum)
	}
}

func TestAddZero(t *testing.T) {
	_, err := Add(0, 2)
	if err == nil {
		t.Error("first arg is zero  - expected error not be nil")
	}
	_, err = Add(1, 0)
	if err == nil {
		t.Error("second arg is zero  - expected error not be nil")
	}
	_, err = Add(0, 0)
	if err == nil {
		t.Error("all arg negative - expected error not be nil")
	}
}

func TestAdd2(t *testing.T) {
	tests := []struct {
		a, b       int
		want       int
		expectErr  bool
		errMessage string
	}{
		{1, -1, 0, true, "arg is zero"},
		{1, -2, 0, true, "arg is negative"},
		{1, 1, 2, false, ""},
	}

	for _, tt := range tests {
		got, err := Add2(tt.a, tt.b)
		if (err != nil) != tt.expectErr {
			t.Errorf("Add2(%d, %d) unexpected error status: got error = %v, expect error = %v", tt.a, tt.b, err, tt.expectErr)
		}
		if err != nil && err.Error() != tt.errMessage {
			t.Errorf("Add2(%d, %d) unexpected error message: got = %v, want = %v", tt.a, tt.b, err.Error(), tt.errMessage)
		}
		if got != tt.want {
			t.Errorf("Add2(%d, %d) = %v; want %v", tt.a, tt.b, got, tt.want)
		}
	}
}
