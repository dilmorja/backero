package backero

import(
	"sync"
	"errors"
	"fmt"
)

// Unique Target Identification
type UTI struct {
	Name 		string
	Version string
	Hosts 	string
}

func (u UTI) String() string {
	return fmt.Sprintf("%s_%s_%s", u.Name, u.Version, u.Hosts)
}

type Target struct {
	ID UTI
}

type Cowboy struct {
	*Target
	m 			sync.RWMutex
	Targets map[string]*Target
}

func (c *Cowboy) NilTargets() bool { return bool(c.Targets == nil) }

func (c *Cowboy) Load(target *Target) error {
	if !c.NilTargets() {
		c.m.RLock()
		if _, ok := c.Targets[target.ID.String()]; !ok {
			c.m.RUnlock()
			c.m.Lock()
			c.Targets[target.ID.String()] = target
			c.m.Unlock()
			return nil
		}
		c.m.RUnlock()
		return nil
	}
	return errNilTargets
}

func New() *Cowboy {
	return &Cowboy{
		new(Target),
		sync.RWMutex{},
		make(map[string]*Target, 0),
	}
}

var errNilTargets error = errors.New("Nil pointer reference")
var errTargetNotExist error = errors.New("Target does not exist")

func (c *Cowboy) Use(tarname string) error {
	if !c.NilTargets() {
		c.m.RLock()
		if tar, ok := c.Targets[tarname]; ok {
			var tars map[string]*Target = c.Targets
			c.m.RUnlock()
			c = &Cowboy{
				tar,
				sync.RWMutex{},
				tars,
			}
			return nil
		}
		c.m.RUnlock()
		return errTargetNotExist
	}
	return errNilTargets
}