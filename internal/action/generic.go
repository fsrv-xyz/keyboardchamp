package action

import (
	"keyboardchamp/internal/keycode"
	"sync"
)

type Action interface {
	Execute()
	GetRequirements() []int
}

type GenericAction struct {
	Requirements []int
}

type Registry struct {
	Checks map[string]func() Action
	mux    sync.Mutex
}

func (gen *GenericAction) GetRequirements() []int {
	return gen.Requirements
}

func CheckRequirements(states map[int]int, requirements []int) bool {
	for _, requirement := range requirements {
		if states[requirement] == keycode.EventTypeKeyUp {
			return false
		}
	}
	return true
}

var RegistryInstance Registry

func (r *Registry) Register(name string, factory func() Action) {
	if r.Checks == nil {
		r.Checks = make(map[string]func() Action)
	}
	r.mux.Lock()
	r.Checks[name] = factory
	r.mux.Unlock()
}
