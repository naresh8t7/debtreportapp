# Debt Report App

- [Instructions](#instruction-to-run)
- [Overview](#overview-of-time)
- [Approach](#approach)
- [Logic](#relevent-logic)
- [Assumptions](#design-assumptions)


# Instruction to run and test

## To run application
The app can be run in following ways
```
 cd DEBTREPORTAPP
 ./run.sh

```

OR

```
cd DEBTREPORTAPP
 go run cmd/main.go
```

## To test application
Executing runlcl.sh also runs the tests present in the app. We can also below command to test with the coverage.
```
cd DEBTREPORTAPP
go test ./... -coverprofile=c.out && go tool cover -html=c.out
```

# Overview
I basically setup a httpclient which deals with API endpoints to provide the relevent debts/paymemts/payment plans. And this info is fed to report generation which generates the report.

# Approach
As mentioned in the overview, setup interface of DebtAPI in httpclient  which deals with API endpoints to provide the relevent debts/paymemts/payment plans and generates debt report in reports package. 

# Relevent Logic
Package httpclient contains interfaces to Debts API and Package Reports contain debt report generation logic

# Design Assumptions
I made following assumptions during my development
- If there are no payment plans then there wont be any payments
- The installmemt frequency and start date has no affect on the payments
- payments can be more or less than the installment amount.