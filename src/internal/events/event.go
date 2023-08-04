package events

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
	newHandlers := make([]*func(), 0, len(s.handlers))
	for _, h := range s.handlers {
		if *h == nil {
			continue
		}
		newHandlers = append(newHandlers, h)
	}
	s.handlers = newHandlers
	for _, h := range s.handlers {
		go (*h)()
	}
}

func (s *Signaler) AddHandler(h *func()) {
	s.handlers = append(s.handlers, h)
}
