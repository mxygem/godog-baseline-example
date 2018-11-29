package example

import (
	"github.com/DATA-DOG/godog"
	"github.com/jaysonesmith/godog-baseline-example/config"
)

func FeatureContext(s *godog.Suite) {
	config := config.NewConfig()
	context := NewTestContext(config)

	t := &tester{
		TestContext: context,
	}

	s.BeforeScenario(func(interface{}) {
		// reset context between scenarios to avoid
		// cross contamination of data
		context = NewTestContext(config)
	})

	s.Step(`^Jack has \$(\d+) in (\d+) bill$`, t.JackHasInBill)
	s.Step(`^he tries to deposit the money into an ATM$`, t.HeTriesToDepositTheMoneyIntoAnATM)
	s.Step(`^the transaction is successful$`, t.TheTransactionIsSuccessful)
	s.Step(`^Jack has \$(\d+) in (\d+) bills$`, t.JackHasInBills)
	s.Step(`^the transaction must fail$`, t.TheTransactionMustFail)
	s.Step(`^he must be told to try again with fewer bills$`, t.HeMustBeToldToTryAgainWithFewerBills)
	s.Step(`^Jack has an account balance of \$(\d+)$`, t.JackHasAnAccountBalanceOf)
	s.Step(`^he tries to deposit the money with a teller$`, t.HeTriesToDepositTheMoneyWithATeller)
	s.Step(`^his account must accurately reflect the increased balance$`, t.HisAccountMustAccuratelyReflectTheIncreasedBalance)
	s.Step(`^Jack has \$(\d+)\.(\d+) in (\d+) bill and (\d+) coin$`, t.JackHasInBillAndCoin)
	s.Step(`^Jack attempts to withdraw \$(\d+) from an ATM$`, t.JackAttemptsToWithdrawFromAnATM)
	s.Step(`^he must be informed of the daily limit$`, t.HeMustBeInformedOfTheDailyLimit)
	s.Step(`^he must be informed that amounts must be in multiples of (\d+)$`, t.HeMustBeInformedThatAmountsMustBeInMultiplesOf)
	s.Step(`^Jack attempts to withdraw \$(\d+) from a teller$`, t.JackAttemptsToWithdrawFromATeller)
}
