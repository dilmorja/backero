package ts

import(
	"github.com/dilmorja/backero"
	"github.com/dilmorja/backero/targets/ts/modules"
)

type TS_t backero.Target

func (t *TS_t) InitModule(mn string) *modules.Module {
	return modules.InitModule(mn)
}

func TypeScript() *backero.Target {
	this := TS_t{
		ID: backero.UTI{
			Name: "typescript",
			Version: "v0",
			Hosts: "ts",
		},
	}

	r := backero.Target(this)

	return &r
}