package example

// SetBalance is a very basic illustration of a
// helper method to set a value in context
func SetBalance(amount float64, tc *TestContext) {
	tc.Balance = amount
}
