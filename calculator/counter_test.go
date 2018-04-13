package calculator

import (
	"testing"
)

func TestCounterIncrease(t *testing.T) {
	expected := 2
	counter := NewCounter(1)
	counter.Increase()
	if counter.Count != expected {
		t.Errorf("expect %d but result is %d", expected, counter.Count)
	}
}
