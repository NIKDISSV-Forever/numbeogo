package events

import (
	"github.com/nikdissv-forever/numbeogo/internal/mutex"
)

type Signaler struct {
	Name     string
	handlers []*func()
}

func NewSignal(name string) Signaler {
	return Signaler{Name: name}
}

func (s *Signaler) String() string {
	return s.Name
}

func (s *Signaler) Bell() {
	nulCount := 0
	for _, h := range s.handlers {
		if (*h) == nil {
			nulCount++
		} else {
			go (*h)()
		}
	}
	if nulCount > 0 {
		mutex.Locker.Lock()
		newHandlers := make([]*func(), 0, len(s.handlers)-nulCount)
		for _, h := range s.handlers {
			if *h == nil {
				continue
			}
			newHandlers = append(newHandlers, h)
		}
		s.handlers = newHandlers
		mutex.Locker.Unlock()
	}
}

func (s *Signaler) AddHandler(h *func()) {
	s.handlers = append(s.handlers, h)
}
