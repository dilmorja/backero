// Copyright (c) 2022, Daniel M. Jaén
// All rights reserved.

package backero

import(
	"testing"
	"sync"
)

func Verify[TEG comparable](t *testing.T, expected TEG, got TEG) {
	if expected != got {
		t.Errorf("\nExpected: %v\nGot: %v\n", expected, got)
	}
}

var testTarget *Target = &Target{
	ID: UTI{
		Name: "test",
		Version: "v0",
		Hosts: "cp",
	},
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

func Test_TargetID(t *testing.T) {
	x := &Target{
		ID: UTI{
			Name: "tt",
			Version: "v0",
			Hosts: "t",
		},
	}

	Verify(t, "tt_v0_t", x.ID.String())
}

func Test_Load(t *testing.T) {
	x := &Cowboy{
		new(Target),
		sync.RWMutex{},
		make(map[string]*Target, 0),
	}

	if err := x.Load(&Target{
		ID: UTI{
			Name: "test",
			Version: "v1",
			Hosts: "cp",
		},
	}); err != nil {
		t.Error(err)
	}

	expected := "test_v1_cp"

	Verify(t, expected, x.Targets[expected].ID.String())
}

func Test_New(t *testing.T) {
	instance := New()
	if err := instance.Load(&Target{
		ID: UTI{
			Name: "a",
			Version: "b",
			Hosts: "c",
		},
	}); err != nil {
		t.Error(err)
	}

	Verify(t, "a", instance.Targets["a_b_c"].ID.Name)
}

func Test_Use(t *testing.T) {
	x := New()

	if err := x.Load(testTarget); err != nil {
		t.Error(err)
	}

	if err := x.Use(testTarget.ID.String()); err != nil {
		t.Error(err)
	}


	Verify(t, "test", x.ID.Name)
}