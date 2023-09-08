package mascot_test

import (
	"testing"

	"example.com/P01/mascot"
)

func TestMascot(t *testing.T) {
	if mascot.BestMascot() != "Go Gopher"{
		t.Fatal("Wrong Mascot :(")
	}
}
