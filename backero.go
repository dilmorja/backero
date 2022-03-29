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
	sync.RWMutex
	Targets map[string]*Target
}

func (c *Cowboy) NilTargets() bool { return bool(c.Targets == nil) }

func (c *Cowboy) Load(target *Target) error {
	if !c.NilTargets() {
		c.RLock()
		if _, ok := c.Targets[target.ID.String()]; !ok {
			c.RUnlock()
			c.Lock()
			c.Targets[target.ID.String()] = target
			c.Unlock()
			return nil
		}
		c.RUnlock()
		return nil
	}
	return errNilTargets
}

var errNilTargets error = errors.New("Nil pointer reference")