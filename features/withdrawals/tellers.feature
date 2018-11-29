Feature: Teller Withdrawal
  
  Rules:
  - Desired withdrawal amount must be available in a customer's account   
  
  Scenario: A withdrawal request for more than a customer has will be declined
    Given Jack has an account balance of $5
    When Jack attempts to withdraw $100 from a teller
    Then the transaction must fail
