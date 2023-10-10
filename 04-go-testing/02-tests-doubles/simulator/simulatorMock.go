package simulator

type CatchSimulatorMock struct {
	CanCatchResponse          bool
	GetLinearDistanceCalled   bool
	GetLinearDistanceResponse float64
}

func (m *CatchSimulatorMock) CanCatch(distance float64, speed float64, catchSpeed float64) bool {
	return m.CanCatchResponse
}

func (m *CatchSimulatorMock) GetLinearDistance(position [2]float64) float64 {
	m.GetLinearDistanceCalled = true
	return m.GetLinearDistanceResponse
}

func NewCatchSimulatorMock(CanCatchResponse bool, GetLinearDistanceResponse float64) CatchSimulatorMock {
	return CatchSimulatorMock{
		CanCatchResponse:          CanCatchResponse,
		GetLinearDistanceResponse: GetLinearDistanceResponse,
	}
}
