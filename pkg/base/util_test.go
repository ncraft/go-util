package base

import "testing"

func TestMustAtoiPanicOnEmptyString(t *testing.T) {
	defer checkPanic(t)
	MustAtoi("")
}

func TestMustAtoi(t *testing.T) {
	i := MustAtoi("3")
	if i != 3 {
		t.Error()
	}
}

func checkPanic(t *testing.T) {
	if r := recover(); r == nil {
		t.Errorf("The code did not panic")
	}
}
