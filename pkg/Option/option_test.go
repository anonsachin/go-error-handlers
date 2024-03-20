package Option_test

import (
	"testing"

	"github.com/anonsachin/go-error-handlers/pkg/Option"
)


func TestOptionSome(t *testing.T) {
	res := Option.Some("hi")
	if *res.Value() != "hi"{
		t.Fatalf("The result value doesn't match. ==> %#v == %#v",res,"hi")
	}
	if res.IsNone() != false {
		t.Fatal("The IsNone atribute is not none.")
	}

	resInt := Option.Some(1)
	if *resInt.Value() != 1 {
		t.Fatalf("The result value doesn't match. ==> %#v == %#v",resInt,1)
	}
	if resInt.IsNone() != false {
		t.Fatal("The IsNone atribute is not none.")
	}
}

func TestOptionNone(t *testing.T) {
	res := Option.None[any]()
	if res.Value() != nil {
		t.Fatal("The result value doesn't match. its not nil.")
	}
	if res.IsNone() != true {
		t.Fatal("The IsNone atribute is not true.")
	}
}