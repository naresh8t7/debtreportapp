package reports

import (
	"debtreportapp/internal/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDebtReport(t *testing.T) {
	type test struct {
		desc         string
		Debts        []types.Debt
		Payments     []types.Payment
		PaymentPlans []types.PaymentPlans
		want         []OutReport
	}

	tests := []test{
		{
			desc: "Test with no payment plans",
			Debts: []types.Debt{
				{
					ID:     1,
					Amount: float64(100),
				},
			},
			PaymentPlans: []types.PaymentPlans{},
			Payments:     []types.Payment{},
			want: []OutReport{
				{
					ID:              1,
					Amount:          float64(100),
					IsInPaymentPlan: false,
					RemainingAmount: float64(100),
				},
			},
		},
		{
			desc: "Test with payment plans and having remaining amount",
			Debts: []types.Debt{
				{
					ID:     1,
					Amount: float64(100),
				},
			},
			PaymentPlans: []types.PaymentPlans{
				{
					DebtID:      1,
					ID:          1,
					AmountToPay: 50,
				},
			},
			Payments: []types.Payment{
				{
					PaymentPlanID: 1,
					Amount:        float64(25),
					Date:          "2020-12-22",
				},
			},
			want: []OutReport{
				{
					ID:              1,
					Amount:          float64(100),
					IsInPaymentPlan: true,
					RemainingAmount: float64(25),
				},
			},
		},
		{
			desc: "Test with payment plans and amount paid off",
			Debts: []types.Debt{
				{
					ID:     1,
					Amount: float64(100),
				},
			},
			PaymentPlans: []types.PaymentPlans{
				{
					DebtID:      1,
					ID:          1,
					AmountToPay: 50,
				},
			},
			Payments: []types.Payment{
				{
					PaymentPlanID: 1,
					Amount:        float64(25),
					Date:          "2020-12-22",
				},
				{
					PaymentPlanID: 1,
					Amount:        float64(25),
					Date:          "2021-01-22",
				},
			},
			want: []OutReport{
				{
					ID:              1,
					Amount:          float64(100),
					IsInPaymentPlan: false,
					RemainingAmount: float64(0),
				},
			},
		},
	}

	for _, tc := range tests {
		rpt := Report{
			Debts:        tc.Debts,
			Payments:     tc.Payments,
			PaymentPlans: tc.PaymentPlans,
		}
		got := rpt.DebtReport()
		assert.Equal(t, tc.want, got, tc.desc)
	}
}
