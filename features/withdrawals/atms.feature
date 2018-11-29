Feature: ATM Withdrawal
  
  Rules:
  - Desired withdrawal amount must be available in a customer's account
  - Maximum of $300 per day can be withdrawn
  - Cash is only dispensed in multiples of $20    
  
  Scenario: A withdrawal request for more than a customer has will be declined
    Given Jack has an account balance of $5
    When Jack attempts to withdraw $100 from an ATM
    Then the transaction must fail
    And he must be informed of the daily limit

  Scenario: A withdrawal request over $300 will be declined
    Given Jack has an account balance of $400
    When Jack attempts to withdraw $400 from an ATM
    Then the transaction must fail
    And he must be informed of the daily limit

  Scenario: A withdrawal request for $25 will be declined as it is not a multiple of 20
    Given Jack has an account balance of $100
    When Jack attempts to withdraw $25 from an ATM
    Then the transaction must fail
    And he must be informed that amounts must be in multiples of 20
