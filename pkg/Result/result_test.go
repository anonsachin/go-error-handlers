package Result_test

import (
	"fmt"
	"testing"

	"github.com/anonsachin/go-error-handlers/pkg/Result"
)


func TestValue(t *testing.T) {
	res := Result.Value("hi")

	if *res.Value() != "hi"{
		t.Fatalf("The result value doesn't match. ==> %#v == %#v",*res.Value(),"hi")
	}

	if res.Error() != nil {
		t.Fatal("The error was not false.")
	}


	resInt := Result.Value(1)
	if *resInt.Value() != 1 {
		t.Fatalf("The result value doesn't match. ==> %#v == %#v",resInt,1)
	}
	if res.IsErr() != false {
		t.Fatal("The IsErr atribute is not false.")
	}
	if resInt.Error() != nil {
		t.Fatal("The IsNone atribute is not none.")
	}
}

func TestError(t *testing.T) {
	err := fmt.Errorf("Error")
	res := Result.Err[any](err)
	if res.Value() != nil {
		t.Fatal("The result value doesn't match. its not nil.")
	}
	if res.Error() != err {
		t.Fatal("The IsNone atribute is not true.")
	}
	if res.IsErr() != true {
		t.Fatal("The IsErr atribute is not true.")
	}
}