package business

import (
	"testing"

	"github.com/avisiedo/go-test-demo/pkg/operations/mocks"
	"github.com/stretchr/testify/require"
	"gotest.tools/assert"
)

func TestMyBusinessLogic1(t *testing.T) {
	mockAddable := &mocks.MockAddable{}
	mockAddable.On("Add", 4, 9).Return(4 + 9)
	result := MyBusinessLogic1(mockAddable)
	require.Error(t, result)
	assert.Equal(t, "Operation failed", result.Error())
	mockAddable.AssertExpectations(t)
}
