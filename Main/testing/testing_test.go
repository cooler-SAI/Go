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
	sum, err := Add2(1, 2)
	if err != nil {
		t.Error("unexpected error")
	} else {
		if sum != 3 {
			t.Errorf("sum expected to be 3; got %d", sum)
		}
	}
}
