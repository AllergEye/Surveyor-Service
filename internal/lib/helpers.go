package lib

import "github.com/google/uuid"

type Helpers interface {
	GenerateUUID() uuid.UUID
}

type HelpersImplementation struct{}

func NewHelpers() Helpers {
	return HelpersImplementation{}
}

func (h HelpersImplementation) GenerateUUID() uuid.UUID {
	return uuid.New()
}
