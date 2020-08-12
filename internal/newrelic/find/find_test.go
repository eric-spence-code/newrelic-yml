package find

import (-
	"fmt"
	"testing"
	"github.com/newrelic/newrelic-client-go/pkg/alerts"
	mock "github.com/newrelic/newrelic-client-go/pkg/testhelpers"
)

func newMockResponse(t *testing.T, mockJSONResponse string, statusCode int) alerts.Alerts {
	ts := mock.NewMockServer(t, mockJSONResponse, statusCode)
	tc := mock.NewTestConfig(t, ts)

	return alerts.New(tc)
}

func TestPolicy(t *testing.T) {
	fmt.Print("hit")
}
