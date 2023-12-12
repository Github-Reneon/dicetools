package dicetools_test

import (
	"testing"

	dt "github.com/Github-Reneon/dicetools"
)

func TestMinMax(t *testing.T) {
	export := 0
	for i := 0; i < 100; i++ {
		export = dt.RollNotation("1d100")
		if export <= 0 || export > 100 {
			t.Error("Roll of 1d100 outside of bounds", export)
		}
	}
}

func TestModifiers(t *testing.T) {
	export := 0
	for i := 0; i < 100; i++ {
		export = dt.RollNotation("1d2+1")
		if export <= 1 || export > 3 {
			t.Error("Modifiers are not applying.", export)
		}
	}
	for i := 0; i < 100; i++ {
		export = dt.RollNotation("1d2-1")
		if export <= -1 || export > 1 {
			t.Error("Modifiers are not applying.", export)
		}
	}
}

func TestOutput(t *testing.T) {
	export := 0
	export = dt.RollNotation("1d20+1-1-20")
	if export > 0 {
		t.Error("Not adding up correctly")
	}

}
