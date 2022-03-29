package backero

import "testing"

func Verify[TEG comparable](t *testing.T, expected TEG, got TEG) {
	if expected != got {
		t.Errorf("\nExpected: %v\nGot: %v\n", expected, got)
	}
}

func Test_NilTargets(t *testing.T) {
	Verify(t, new(Cowboy).NilTargets(), true)
}