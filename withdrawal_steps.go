package example

import "github.com/DATA-DOG/godog"

func (t *tester) JackAttemptsToWithdrawFromAnATM(amount float64) error {
	t.TestContext.TransactionResponse = Withdraw(amount, t.TestContext)
	return nil
}

func (t *tester) JackAttemptsToWithdrawFromATeller(arg1 int) error {
	return godog.ErrPending
}
