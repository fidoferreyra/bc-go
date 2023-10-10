package prey

type tunaStub struct {
	// max speed in m/s
	maxSpeed float64
}

func (t *tunaStub) GetSpeed() float64 {
	return t.maxSpeed
}

func CreateTunaStubMaxSpeed() Prey {
	return &tunaStub{
		maxSpeed: 252,
	}
}

func CreateTunaStubZeroSpeed() Prey {
	return &tunaStub{
		maxSpeed: 0,
	}
}

func CreateTunaStubMidSpeed() Prey {
	return &tunaStub{
		maxSpeed: 126,
	}
}
