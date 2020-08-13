package find

import (
	"testing"

	"github.com/newrelic/newrelic-client-go/pkg/alerts"
)

type MockClient struct {
	response []alerts.Policy
}

func (mc *MockClient) ListPolicies(*alerts.ListPoliciesParams) ([]alerts.Policy, error) {
	return mc.response, nil
}

func TestPolicyFound(t *testing.T) {
	policy := alerts.Policy{
		ID:   1234,
		Name: "test-alert-policy-1",
	}
	policies := []alerts.Policy{policy}

	client := MockClient{
		response: policies,
	}

	find := New(&client)

	response := find.Policy("test-alert-policy-1")
	if response == nil {
		t.Errorf("Did not find alert policy")
	}
}

func TestPolicyNotFound(t *testing.T) {
	policies := []alerts.Policy{}

	client := MockClient{
		response: policies,
	}

	find := New(&client)

	response := find.Policy("test-alert-policy-1")
	if response != nil {
		t.Errorf("Found an policy when it should not have")
	}
}
