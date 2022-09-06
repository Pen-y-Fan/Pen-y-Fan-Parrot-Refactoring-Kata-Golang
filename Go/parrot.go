package parrot

import (
	"errors"
	"math"
)

type parrotType int

const (
	baseSpeed 				float64 = 12.0
	loadFactor 				float64 = 9.0
	norwegianBlueMaxSpeed 	float64 = 24.0
)

const (
	TypeEuropean      parrotType = 1
	TypeAfrican       parrotType = 2
	TypeNorwegianBlue parrotType = 3
)

// Parrot has a Speed.
type Parrot interface {
	Speed() (float64, error)
}

type europeanParrot struct {
}

type unknownParrot struct {
}

type africanParrot struct {
	numberOfCoconuts int
}

type norwegianBlueParrot struct {
	voltage	float64
	nailed  bool
}

func CreateParrot(t parrotType, numberOfCoconuts int, voltage float64, nailed bool) Parrot {
	switch t {
		case TypeEuropean:
			return europeanParrot{}
		case TypeAfrican:
			return africanParrot{numberOfCoconuts}
		case TypeNorwegianBlue:
			return norwegianBlueParrot{voltage, nailed}
		default:
			return unknownParrot{}
	}
}

func (parrot europeanParrot) Speed() (float64, error) {
	return baseSpeed, nil
}


func (parrot africanParrot) Speed() (float64, error) {
	return math.Max(0, baseSpeed-loadFactor*float64(parrot.numberOfCoconuts)), nil
}

func (parrot norwegianBlueParrot) Speed() (float64, error) {
	if parrot.nailed {
		return 0, nil
	}
	return parrot.computeBaseSpeedForVoltage(parrot.voltage), nil
}

func (parrot unknownParrot) Speed() (float64, error) {
	return 0, errors.New("should be unreachable")
}

func (parrot norwegianBlueParrot) computeBaseSpeedForVoltage(voltage float64) float64 {
	return math.Min(norwegianBlueMaxSpeed, voltage*baseSpeed)
}
