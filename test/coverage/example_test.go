package coverage

import "testing"

func TestAdd(t *testing.T) {
	n := Add(1, 2)
	if n == 3 {
		t.Log("success")
	} else {
		t.Error("fail")
	}
}
func TestIsFindName(t *testing.T) {
	if isSuccess := IsFindName("Daniel", []string{"Andy", "Jack", "Daniel", "Sam"}); isSuccess {
		t.Log("success")
	} else {
		t.Error("fail")
	}
}

func TestIsNotFindNames(t *testing.T) {
	if isSuccess := IsFindName("Lucy", []string{"Andy", "Jack", "Daniel", "Sam"}); isSuccess {
		t.Error("fail")
	} else {
		t.Log("success")
	}
}
