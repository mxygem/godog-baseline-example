package example

import "github.com/jaysonesmith/godog-baseline-example/config"

// TestContext contains data that is specific to
// the scenario/test that is currently running
// It should be reinitialized between scenarios to ensure data
// from the previous scenario doesn't create conflicts.
type TestContext struct {
	Balance             float64
	ToDispense          float64
	TransactionResponse string
}

//NewTestContext returns a pointer to a new scenario context struct
func NewTestContext(config config.Config) *TestContext {
	return &TestContext{
		Balance: 0.00,
	}
}
