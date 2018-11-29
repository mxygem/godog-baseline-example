package example

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithdrawal(t *testing.T) {
	testCases := []struct {
		name              string
		withdrawalAmount  float64
		expectedDispensed float64
		startingBalance   float64
		expectedResponse  string
	}{
		{
			name:              "Amount is within all parameters",
			withdrawalAmount:  50.0,
			expectedDispensed: 50.0,
			startingBalance:   100.00,
			expectedResponse:  "",
		},
		{
			name:              "Amount is not multiple of 20",
			withdrawalAmount:  25.0,
			expectedDispensed: 0.0,
			startingBalance:   100.00,
			expectedResponse:  "unable to process transaction as desired amount is not a multiple of 20",
		},
		{
			name:              "Amount is greater than daily limit",
			withdrawalAmount:  500.0,
			expectedDispensed: 0.0,
			startingBalance:   100.00,
			expectedResponse:  "unable to process transaction as desired amount exceeds daily limit",
		},
		{
			name:              "Amount is greater than balance",
			withdrawalAmount:  101.00,
			expectedDispensed: 0.0,
			startingBalance:   100.00,
			expectedResponse:  "unable to process transaction as desired amount exceeds current balance",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testContext := &TestContext{Balance: tc.startingBalance}

			resp := Withdraw(tc.withdrawalAmount, testContext)

			assert.Equal(t, tc.expectedDispensed, testContext.ToDispense)
			assert.Equal(t, testContext.Balance, tc.startingBalance-tc.expectedDispensed)
			assert.Equal(t, tc.expectedResponse, resp)
		})
	}
}
