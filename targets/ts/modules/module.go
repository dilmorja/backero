package modules

type Module struct {
	Name string
	Memory Mem_t
}

type Mem_t map[string]map[string]interface{}

func InitModule(moduleName string) *Module {
	return &Module{
		Name: moduleName,
		Memory: make(Mem_t, 0),
	}
}