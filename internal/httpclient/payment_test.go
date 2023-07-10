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

func TestGetPayments_Success(t *testing.T) {
	mockHTTPClient, dClient := setUp(t)
	payments := `[{"amount": 123.46,"payment_plan_id": 0},{"amount": 100,"payment_plan_id": 1}]`
	expectedPayments := []types.Payment{
		{
			PaymentPlanID: 0,
			Amount:        float64(123.46),
		},
		{
			PaymentPlanID: 1,
			Amount:        float64(100),
		},
	}
	expectedResp := &http.Response{
		StatusCode: 200,
		Status:     "200 OK",
		Body:       ioutil.NopCloser(strings.NewReader(payments)),
	}
	mockHTTPClient.EXPECT().Do(gomock.Any()).Return(expectedResp, nil)
	out, err := dClient.GetPayments()
	assert.Nil(t, err, "No error expected")
	assert.Equal(t, expectedPayments, out, "response expected")
}

func TestGetPayments_Failure(t *testing.T) {
	mockHTTPClient, dClient := setUp(t)

	expectedResp := &http.Response{
		StatusCode: 500,
		Status:     "500  Error",
		Body:       ioutil.NopCloser(strings.NewReader(`{}`)),
	}
	mockHTTPClient.EXPECT().Do(gomock.Any()).Return(expectedResp, nil)
	_, err := dClient.GetPayments()
	assert.NotNil(t, err, "error expected")
}
