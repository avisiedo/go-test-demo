package operations

type MyOperand int

func (i *MyOperand) Add(operand1 int, operand2 int) int {
	var op1 int = int(operand1)
	var op2 int = int(operand2)
	var result int = op1 + op2
	return result
}
