package reports

import (
	"debtreportapp/internal/types"
)

type Report struct {
	Debts             []types.Debt
	Payments          []types.Payment
	PaymentPlans      []types.PaymentPlans
	debtToPaymentPlan map[int]types.PaymentPlans
	planToPayments    map[int]float64
}

type OutReport struct {
	Amount          float64 `json:"amount"`
	ID              int     `json:"id"`
	IsInPaymentPlan bool    `json:"is_in_payment_plan"`
	RemainingAmount float64 `json:"remaining_amount"`
}

func (r *Report) DebtReport() []OutReport {
	r.debtToPaymentPlan = make(map[int]types.PaymentPlans)
	r.planToPayments = make(map[int]float64)

	for _, pp := range r.PaymentPlans {
		r.debtToPaymentPlan[pp.DebtID] = pp
	}

	for _, p := range r.Payments {
		r.planToPayments[p.PaymentPlanID] += p.Amount
	}
	OutReports := make([]OutReport, 0)
	for _, d := range r.Debts {
		out := OutReport{
			ID:              d.ID,
			Amount:          d.Amount,
			IsInPaymentPlan: false,
			RemainingAmount: d.Amount,
		}
		if len(r.PaymentPlans) > 0 {
			out.RemainingAmount = r.calculateRemainingAmount(d.ID)
			if out.RemainingAmount > 0 {
				out.IsInPaymentPlan = true
			}
		}
		OutReports = append(OutReports, out)

	}
	return OutReports
}

func (r *Report) calculateRemainingAmount(debtID int) float64 {
	var remainingAmount float64
	if pp, ok := r.debtToPaymentPlan[debtID]; ok {
		remainingAmount = pp.AmountToPay - r.planToPayments[pp.ID]
	}
	return remainingAmount
}
