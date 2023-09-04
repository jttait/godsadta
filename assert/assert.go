package assert

import "testing"

func AssertTrue(i bool, t *testing.T) {
	if !i {
		t.Fatalf("Got false. Want true.")
	}
}

func AssertFalse(i bool, t *testing.T) {
	if i {
		t.Fatalf("Got true. Want false.")
	}
}

func AssertEqual[T comparable](got T, want T, t *testing.T) {
	if got != want {
		t.Fatalf("Got %v. Want %v.\n", got, want)
	}
}
