package example

import "github.com/DATA-DOG/godog"

func (t *tester) JackHasAnAccountBalanceOf(amount float64) error {

	return nil
}

func (t *tester) JackHasInBill(arg1, arg2 int) error {
	return godog.ErrPending
}

func (t *tester) JackHasInBills(arg1, arg2 int) error {
	return godog.ErrPending
}

func (t *tester) JackHasInBillAndCoin(arg1, arg2, arg3, arg4 int) error {
	return godog.ErrPending
}
