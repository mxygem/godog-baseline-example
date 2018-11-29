package example

import (
	"math"
)

// Withdraw enables the withdrawal of money
func Withdraw(amount float64, tc *TestContext) string {
	if amount > 300.00 {
		return "unable to process transaction as desired amount exceeds daily limit"
	} else if amount > tc.Balance {
		return "unable to process transaction as desired amount exceeds current balance"
	} else if math.Remainder(amount, 2) != 0 {
		return "unable to process transaction as desired amount is not a multiple of 20"
	}

	tc.Balance = tc.Balance - amount
	tc.ToDispense = amount

	return ""
}
