package shark

import (
	"testdoubles/prey"
	"testdoubles/simulator"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
Realizar test unitarios del método Hunt del tiburón blanco, cubriendo todos los casos posibles, usando los stubs y mocks creados anteriormente:
1. El tiburón logra cazar el atún al ser más veloz y al estar en una distancia corta. Hacer un assert de que el método GetLinearDistance fue llamado.
2. El tiburón no logra cazar el atún al ser más lento.
3. El tiburón no logra cazar el atún por estar a una distancia muy larga, a pesar de ser
más veloz.
*/

func TestCatchWhenFasterAndShortDistance(t *testing.T) {
	// Arrange
	mockSim := simulator.NewCatchSimulatorMock(true, 100)
	var sim simulator.CatchSimulator = &mockSim
	shark := CreateWhiteShark(sim)
	prey := prey.CreateTunaStubZeroSpeed()

	// Act
	err := shark.Hunt(prey)

	// Assert
	assert.Nil(t, err)
	assert.True(t, mockSim.GetLinearDistanceCalled, "GetLinearDistance should be called")
}

func TestCantCatchWhenSlowerThanPrey(t *testing.T) {
	// Arrange
	mockSim := simulator.NewCatchSimulatorMock(false, 100)
	var sim simulator.CatchSimulator = &mockSim
	shark := CreateWhiteShark(sim)
	prey := prey.CreateTunaStubMaxSpeed()

	// Act
	err := shark.Hunt(prey)
	// Assert
	assert.NotNil(t, err, err.Error())
	assert.ErrorIs(t, err, ErrCouldntHunt)
}

func TestCantCatchWhenTooFar(t *testing.T) {
	// Arrange
	mockSim := simulator.NewCatchSimulatorMock(false, 1000)
	var sim simulator.CatchSimulator = &mockSim
	shark := CreateWhiteShark(sim)
	prey := prey.CreateTunaStubZeroSpeed()

	// Act
	err := shark.Hunt(prey)

	// Arrange
	assert.NotNil(t, err, err.Error())
	assert.ErrorIs(t, err, ErrCouldntHunt)
}
