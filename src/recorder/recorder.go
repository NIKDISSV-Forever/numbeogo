// Package recorder for GUI integration
package recorder

import "github.com/hashicorp/go-set"

type Params struct {
	Titles, Regions string
}

type Container struct {
	Titles, Regions map[string]string
	Countries       *set.Set[string]
}

var record = false
var Recorded Container

func init() {
	Reset()
}

func SetRecord(state bool) {
	record = state
}

func Reset() {
	Recorded = Container{
		Titles:    make(map[string]string),
		Regions:   make(map[string]string),
		Countries: set.New[string](200),
	}
}
