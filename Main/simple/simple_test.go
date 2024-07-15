package simple

import "testing"

func TestWordTest(t *testing.T) {
	expected := WordTest()
	if WordTest() != expected {
		t.Error("Expected", expected, "got", WordTest())
	}

}
