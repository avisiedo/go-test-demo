package operations

// This file is linked only for the tests
// We can launch from the command line by:
// go test ./pkg/operations/...

// TODO Fix the import to use testify

// Many times the test will be on this form
// func TestMyOperandAdd(t *testing.T) {
// 	var tu MyOperand
// 	result := tu.Add(4, 3)
// 	assert.Equal(t, 7, int(result))
// }

// When several scenarios can be represented in one table
// we could try to take advantage of iterating test on the
// table values; on this scenarios it could looks like this
// test
// func TestMyOperandAddTable(t *testing.T) {
// 	// Define local types
// 	type TestCaseGiven struct {
// 		Op1 int
// 		Op2 int
// 	}
// 	type TestCaseExpected int
// 	type TestCase struct {
// 		Given    TestCaseGiven
// 		Expected TestCaseExpected
// 	}

// 	// Define table
// 	testCases := []TestCase{
// 		{
// 			Given: TestCaseGiven{
// 				Op1: (1),
// 				Op2: (3),
// 			},
// 			Expected: TestCaseExpected(4),
// 		},
// 		{
// 			Given: TestCaseGiven{
// 				Op1: (8),
// 				Op2: (5),
// 			},
// 			Expected: TestCaseExpected(13),
// 		},
// 	}

// 	// Launch every test case
// 	for _, testCase := range testCases {
// 		var tu MyOperand
// 		result := tu.Add(testCase.Given.Op1, testCase.Given.Op2)
// 		assert.Equal(t, int(testCase.Expected), int(result))
// 		require.Equal(t, int(testCase.Expected), int(result))
// 	}
// }
