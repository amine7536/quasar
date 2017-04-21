package conf

import "github.com/amine7536/quasar/event"

// OutputHandler generic output handler
type OutputHandler interface {
	Send(event event.Event) error
}

var (
	MapOutputs = map[string]OutputHandler{}
)

// RegisterOutput with its name and handler
func RegisterOutput(name string, output OutputHandler) {
	MapOutputs[name] = output
}
