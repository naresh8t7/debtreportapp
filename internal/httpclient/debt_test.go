package httpclient

import (
	"debtreportapp/internal/httpclient/mock"
	"debtreportapp/internal/types"

	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

const (
	Dummy_API_URL = "https://api.dummy.com"
)

func TestGetDebts_Success(t *testing.T) {
	mockHTTPClient, dClient := setUp(t)
	debts := `[{"amount": 123.46,"id": 0},{"amount": 100,"id": 1}]`
	expectedDebts := []types.Debt{
		{
			ID:     0,
			Amount: float64(123.46),
		},
		{
			ID:     1,
			Amount: float64(100),
		},
	}
	expectedResp := &http.Response{
		StatusCode: 200,
		Status:     "200 OK",
		Body:       ioutil.NopCloser(strings.NewReader(debts)),
	}
	mockHTTPClient.EXPECT().Do(gomock.Any()).Return(expectedResp, nil)
	out, err := dClient.GetDebts()
	assert.Nil(t, err, "No error expected")
	assert.Equal(t, expectedDebts, out, "response expected")
}

func TestGetDebts_Failure(t *testing.T) {
	mockHTTPClient, dClient := setUp(t)

	expectedResp := &http.Response{
		StatusCode: 500,
		Status:     "500  Error",
		Body:       ioutil.NopCloser(strings.NewReader(`{}`)),
	}
	mockHTTPClient.EXPECT().Do(gomock.Any()).Return(expectedResp, nil)
	_, err := dClient.GetDebts()
	assert.NotNil(t, err, "error expected")
}

func setUp(t *testing.T) (*mock.MockHTTPClient, *DebtClient) {
	ctrl := gomock.NewController(t)
	t.Cleanup(ctrl.Finish)
	mockHTTPClient := mock.NewMockHTTPClient(ctrl)

	baseURL, err := url.Parse(Dummy_API_URL)
	assert.Nil(t, err, "No error expected")

	debtClient := &DebtClient{
		baseURL: *baseURL,
		client:  mockHTTPClient,
	}

	return mockHTTPClient, debtClient
}
