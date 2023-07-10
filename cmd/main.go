package main

import (
	"debtreportapp/internal/httpclient"
	"debtreportapp/internal/reports"
	"encoding/json"
	"fmt"

	log "github.com/sirupsen/logrus"
)

const (
	BaseURL = "https://my-json-server.typicode.com/druska/trueaccord-mock-payments-api"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	debtClient, err := httpclient.NewDebtClient(BaseURL)
	if err != nil {
		log.Fatalf(" Failed to setup debt client %v", err)
	}

	debts, err := debtClient.GetDebts()
	if err != nil {
		log.Errorf("unable to fetch debts from API %v", err)
	}
	payments, err := debtClient.GetPayments()
	if err != nil {
		log.Errorf("unable to fetch payments from API %v", err)
	}
	paymentPlans, err := debtClient.GetPaymentPlans()
	if err != nil {
		log.Errorf("unable to fetch paymentPlans from API %v", err)
	}
	report := reports.Report{
		Debts:        debts,
		PaymentPlans: paymentPlans,
		Payments:     payments,
	}

	opJSON, err := json.MarshalIndent(report.DebtReport(), "", "    ")
	if err != nil {
		log.Fatal("failed to generate json", err)
	}
	fmt.Printf("%s\n", string(opJSON))

}
