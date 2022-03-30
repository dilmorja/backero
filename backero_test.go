package backero

import "testing"

func Verify[TEG comparable](t *testing.T, expected TEG, got TEG) {
	if expected != got {
		t.Errorf("\nExpected: %v\nGot: %v\n", expected, got)
	}
}

func Test_NilTargets(t *testing.T) {
	Verify(t, true, new(Cowboy).NilTargets())
}

func Test_UTI(t *testing.T) {
	x := UTI{
		Name: "test",
		Version: "v0",
		Hosts: "cp",
	}

	Verify(t, "test_v0_cp", x.String())
}