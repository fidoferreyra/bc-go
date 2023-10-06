package hunt

import "fmt"

type Shark struct {
	hungry bool
	tired  bool
	speed  int
}

type Prey struct {
	name  string
	speed int
}

// Errors
var (
	ErrTired     = fmt.Errorf("cannot hunt, i am really tired")
	ErrNotHungry = fmt.Errorf("cannot hunt, i am not hungry")
	ErrCantCatch = fmt.Errorf("could not catch it")
	ErrNoPrey    = fmt.Errorf("cannot hunt, there is no prey")
)

func (s *Shark) Hunt(p *Prey) error {
	if p == nil {
		return ErrNoPrey
	}
	if s.tired {
		return ErrTired
	}
	if !s.hungry {
		return ErrNotHungry
	}
	if p.speed >= s.speed {
		s.tired = true
		return ErrCantCatch
	}

	s.hungry = false
	s.tired = true
	return nil
}
