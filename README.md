# godog-baseline-example

A quick demo project showing how I set up godog suites for black-box style API testing.
This is not meant to be the end all be all for how a project should be configured. I'm quite open
to improvements and suggestions, so drop them my way if you've got them! **WARNING** This project may be out of date and not work without tweaking.

The project is centered around some VERY basic Banking activities.

## Prerequisites

* go
* godog

## Runnable Example

For a passing example run `godog ./features/withdrawals/atms.feature:8`

## Project Set Up

Here's a breakdown of the repo:

```bash
.
├── config
│  ├── config.go
│  └── config_dev.json
├── features
│  ├── deposits
│  │  ├── cash
│  │  │  ├── atms.feature
│  │  │  └── tellers.feature
│  │  └── checks
│  │     └── checks.feature
│  └── withdrawals
│     ├── atms.feature
│     └── tellers.feature
├── account_steps.go
├── account_support.go
├── account_support_test.go
├── deposit_steps.go
├── godog_baseline_example_test.go
├── README.md
├── result_steps.go
├── test_context.go
├── tester.go
└── withdrawal_steps.go
```

### Config

Directory for managing config files as needed. Currently uses viper for loading and reading files.

### Features

The home for all the Cucumber feature files for the suite. These are grouped together by feature and
then broken down further as needed to be more specific.

### godog_baseline_example_test.go

The main godog file. Typically this should be named something similar to the project's name, which
is how it's got its beautiful name!

* FeatureContext - This is the foundation/backbone of the suite. Set up happens here for things
  like configs, context, and where godog registers all the steps available. I've tried moving these
  to other directories/files but have decided for now to keep them here.
  * tester - This is the main hub if you will for the scenarios to be associated to. This assists for things like sharing data across steps of the same scenario (not multiples). Things like loggers can also be set up here to make them accessible. Make sure to reset the context before each scenario!

### ***_steps.go

These files are where the matching methods that godog has registered live. They should be named in a way that helps communicate the type of steps they contain if you'd like to separate them out.

The methods in these files should have as little code in them as possible as they're meant to be more of a bridge between the feature files and the support files.

Assertions for step definitions should live in these files.

### ***_support.go

These files are where the test suite should actually work. These files should be a collection of
helper methods to perform various setup and action steps. This model allows for easier changing of
the underlying processes in a more centralized place and allows step definitions to be expressed in
language that makes more sense to readers rather than trying worry about being DRY or having
incidental details show up in the scenarios.

### ***_support_test.go

*Who tests the tests?*

This section is meant to be where the support code's unit tests are.

### test_context.go & tester.go

These files are home to the structs behind the associated data objects.
