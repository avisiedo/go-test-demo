package mocks

import (
	"github.com/stretchr/testify/mock"
)

type MockAddable struct {
	mock.Mock
}

func (m *MockAddable) Add(operand1 int, operand2 int) int {
	// This method only implement the interface
	// We use the mock methods to:
	// - register the call to the method.
	// - retrieve the expected returns for this call.
	args := m.Called(operand1, operand2)

	// The index is the positional returning value
	return args.Int(0)
}
