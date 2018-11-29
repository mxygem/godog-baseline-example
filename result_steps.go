package example

import (
	"fmt"

	"github.com/DATA-DOG/godog"
)

func (t *tester) TheTransactionIsSuccessful() error {
	if t.TestContext.TransactionResponse != "" {
		return fmt.Errorf("expected transaction to be successful, but got %s", t.TestContext.TransactionResponse)
	}
	return nil
}

func (t *tester) TheTransactionMustFail() error {
	if t.TestContext.TransactionResponse == "" {
		return fmt.Errorf("expected transaction to have failed, but it was successful")
	}
	return nil
}

func (t *tester) HeMustBeToldToTryAgainWithFewerBills() error {
	return godog.ErrPending
}

func (t *tester) HeMustBeInformedOfTheDailyLimit() error {
	if t.TestContext.TransactionResponse == "unable to process transaction as desired amount exceeds daily limit" {
		return fmt.Errorf("expected daily limit message but an unexpected transaction response was returned: %s", t.TestContext.TransactionResponse)
	}
	return nil
}

func (t *tester) HeMustBeInformedThatAmountsMustBeInMultiplesOf(arg1 int) error {
	return godog.ErrPending
}
