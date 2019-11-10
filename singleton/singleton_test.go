package singleton

import "testing"

func TestGetInstance(t *testing.T) {
	aCounter := GetInstance()

	if aCounter == nil {
		//Test of acceptance criteria 1 failed
		t.Error("a new counter object must have been made")
	}

	anotherCounter := GetInstance()
	if aCounter != anotherCounter {
		t.Error("the counters seem to be different")
	}

	mustBeOne := aCounter.AddOne()
	if mustBeOne != 1 {
		t.Errorf("after initialisation and adding one, the `aCounter.count`" +
			"count should be one (1) but it's instead %d", mustBeOne)
	}

	mustBeTwo := anotherCounter.AddOne()
	if mustBeTwo != 2 {
		t.Errorf("given that `aCounter` has invoked `AddOne` once, " +
			"`anotherCounter.AddOne()` should return 2 but returned %d",
			mustBeTwo)
	}
}
