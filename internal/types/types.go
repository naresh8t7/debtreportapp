package types

type Payment struct {
	Amount        float64 `json:"amount"`
	Date          string  `json:"date"`
	PaymentPlanID int     `json:"payment_plan_id"`
}

type PaymentPlans struct {
	AmountToPay          float64 `json:"amount_to_pay"`
	DebtID               int     `json:"debt_id"`
	ID                   int     `json:"id"`
	InstallmentAmount    float64 `json:"installment_amount"`
	InstallmentFrequency string  `json:"installment_frequency"`
	StartDate            string  `json:"start_date"`
}

type Debt struct {
	Amount float64 `json:"amount"`
	ID     int     `json:"id"`
}
