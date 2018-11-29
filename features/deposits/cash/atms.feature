Feature: Depositing Cash Via ATMs
  For folks who want to deposit physical money
  without interacting with tellers will use ATMs.
  These scenarios assume a user has successfully
  authenticated and do not check this. Authentication
  is checked by the similarly named feature*

  *Does not exist in this project at the moment.

  Rules:
  - Only paper money can be deposited, not coins
  - A maximum of 20 bills are allowed in any single deposit

  Scenario: A deposit of a single $5 bill will be successful
    Given Jack has $5 in 1 bill
    When he tries to deposit the money into an ATM
    Then the transaction is successful

  Scenario: A deposit of $125 all in fives will not be successful
    Given Jack has $125 in 25 bills
    When he tries to deposit the money into an ATM
    Then the transaction must fail
    And he must be told to try again with fewer bills
