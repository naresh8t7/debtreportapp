package httpclient

import (
	"debtreportapp/internal/types"

	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetPaymentPlans_Success(t *testing.T) {
	mockHTTPClient, dClient := setUp(t)
	paymentPlans := `[{"amount_to_pay": 123.46,"id": 0},{"amount_to_pay": 100,"id": 1}]`
	expectedPaymentPlans := []types.PaymentPlans{
		{
			ID:          0,
			AmountToPay: float64(123.46),
		},
		{
			ID:          1,
			AmountToPay: float64(100),
		},
	}
	expectedResp := &http.Response{
		StatusCode: 200,
		Status:     "200 OK",
		Body:       ioutil.NopCloser(strings.NewReader(paymentPlans)),
	}
	mockHTTPClient.EXPECT().Do(gomock.Any()).Return(expectedResp, nil)
	out, err := dClient.GetPaymentPlans()
	assert.Nil(t, err, "No error expected")
	assert.Equal(t, expectedPaymentPlans, out, "response expected")
}

func TestGetPaymentPlans_Failure(t *testing.T) {
	mockHTTPClient, dClient := setUp(t)

	expectedResp := &http.Response{
		StatusCode: 500,
		Status:     "500  Error",
		Body:       ioutil.NopCloser(strings.NewReader(`{}`)),
	}
	mockHTTPClient.EXPECT().Do(gomock.Any()).Return(expectedResp, nil)
	_, err := dClient.GetPaymentPlans()
	assert.NotNil(t, err, "error expected")
}
