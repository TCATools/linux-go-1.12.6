package multimain_test

import "testing"

func TestMain(m *testing.M) {
	// Some users run m.Run multiple times, changing
	// some kind of global state between runs.
	// This used to work so I guess now it has to keep working.
	// See golang.org/issue/23129.
	m.Run()
	m.Run()
}

func Test(t *testing.T) {
	t.Log("notwithstanding")
}
