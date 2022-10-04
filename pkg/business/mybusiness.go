package business

import (
	"fmt"

	"github.com/avisiedo/go-test-demo/pkg/operations"
)

func MyBusinessLogic1(operation operations.Addable) error {
	if operation.Add(4, 9) > 10 {
		return fmt.Errorf("Operation failed")
	}
	return nil
}
