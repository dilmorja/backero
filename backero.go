// Copyright (c) 2022, Daniel M. Jaén
// All rights reserved.

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

// Uses the values of the fields and returns a formatted string
func (u UTI) String() string {
	return fmt.Sprintf("%s_%s_%s", u.Name, u.Version, u.Hosts)
}

// An implementation that handles code generation.
type Target struct {
	ID UTI
}

// An implementation that handles targets
type Cowboy struct {
	*Target
	m 			sync.RWMutex
	Targets map[string]*Target
}

// Quickly check if the map that stores the targets is not nil
func (c *Cowboy) NilTargets() bool { return bool(c.Targets == nil) }

// Append a new target to the map
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

// Create an empty (but not nil) *Cowboy structure
func New() *Cowboy {
	return &Cowboy{
		new(Target),
		sync.RWMutex{},
		make(map[string]*Target, 0),
	}
}

var errNilTargets error = errors.New("Nil pointer reference")
var errTargetNotExist error = errors.New("Target does not exist")

// Returns a copy of the previous structure
// that only replaces the target to use.
func (c *Cowboy) Use(tarname string) error {
	if !c.NilTargets() {
		c.m.RLock()
		if tar, ok := c.Targets[tarname]; ok {
			c.m.RUnlock()
			c.m.Lock()
			c.Target = tar
			c.m.Unlock()
			return nil
		}
		c.m.RUnlock()
		return errTargetNotExist
	}
	return errNilTargets
}

// Load and use the given target
func (c *Cowboy) LoadAndUse(tar *Target) error {
	if err := c.Load(tar); err != nil {
		return err
	}

	c.m.Lock()
	// The corresponding function is not used 
	// to avoid the unnecessary computation that this entails.
	c.Target = tar
	c.m.Unlock()

	return nil
}