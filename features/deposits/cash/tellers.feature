Feature: Depositing Cash Via Tellers
  For folks who want to enter one of our branches
  to deposit their money, they'll need to interact
  with a teller to do so. The following scenarios
  assume the user has a valid account.

  Rules:
  - Cash deposited is immediately available
  - Coins and paper bills are accepted for deposit

  Scenario: Deposited cash has no wait time for availability in a user's account
    Given Jack has an account balance of $0
    And Jack has $10 in 1 bill
    When he tries to deposit the money with a teller
    Then his account must accurately reflect the increased balance

  Scenario: Customers with coins can deposit them along with bills
    Given Jack has $1.25 in 1 bill and 1 coin
    When he tries to deposit the money with a teller
    Then his account must accurately reflect the increased balance
