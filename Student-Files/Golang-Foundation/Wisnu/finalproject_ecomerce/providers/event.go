package providers

import (
	"finalproject_ecomerce/engine"
	"github.com/chuckpreslar/emission"
)

type (
	event struct {
		emitter *emission.Emitter
	}
)

// NewEmitter instances new event
func NewEmitter() engine.Emitter {
	return &event{emission.NewEmitter()}
}

func (e *event) On(event, listener interface{}) {
	e.emitter.On(event, listener)
}

func (e *event) Emit(event interface{}, args ...interface{}) {
	e.emitter.Emit(event, args...)
}
