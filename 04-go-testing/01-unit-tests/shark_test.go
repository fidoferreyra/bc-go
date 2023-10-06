package hunt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	fasterPrey     = Prey{name: "dolphin", speed: 100}
	slowerPrey     = Prey{name: "tuna", speed: 50}
	hungryShark    = Shark{hungry: true, tired: false, speed: 80}
	tiredShark     = Shark{hungry: true, tired: true, speed: 80}
	notHungryShark = Shark{hungry: false, tired: false, speed: 80}
)

func TestSharkHuntsSuccessfully(t *testing.T) {
	// Arrange
	shark := hungryShark
	prey := slowerPrey
	// Act
	actualError := shark.Hunt(&prey)
	// Assert
	assert.Nil(t, actualError, "Must be able to hunt")
	assert.Equal(t, false, shark.hungry, "Must not be hungry")
	assert.Equal(t, true, shark.tired, "Must be tired")
}

func TestSharkCannotHuntBecauseIsTired(t *testing.T) {
	// Arrange
	shark := tiredShark
	prey := slowerPrey
	expectedError := ErrTired
	// Act
	actualError := shark.Hunt(&prey)
	// Assert
	assert.ErrorIs(t, actualError, expectedError, "Must not be able to hunt")
}

func TestSharkCannotHuntBecauseIsNotHungry(t *testing.T) {
	// Arrange
	shark := notHungryShark
	prey := slowerPrey
	expectedError := ErrNotHungry
	// Act
	actualError := shark.Hunt(&prey)
	// Assert
	assert.ErrorIs(t, actualError, expectedError, "Must not be able to hunt")

}

func TestSharkCannotReachThePrey(t *testing.T) {
	// Arrange
	shark := hungryShark
	prey := fasterPrey
	expectedError := ErrCantCatch
	// Act
	actualError := shark.Hunt(&prey)
	// Assert
	assert.ErrorIs(t, actualError, expectedError, "Must not be able to hunt")
}

func TestSharkHuntNilPrey(t *testing.T) {
	// Arrange
	shark := hungryShark

	expectedError := ErrNoPrey
	// Act
	actualError := shark.Hunt(nil)
	// Assert
	assert.ErrorIs(t, actualError, expectedError, "Must not be able to hunt")
}
