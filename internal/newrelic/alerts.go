package newrelic

import (
	"go.uber.org/zap"
	"github.com/newrelic/newrelic-client-go/pkg/alerts"
)

// Client - Interface for network calls
type Client interface {
	CreatePolicy(alerts.Policy) (*alerts.Policy, error)
	ListPolicies(*alerts.ListPoliciesParams) ([]alerts.Policy, error)
	UpdatePolicy(alerts.Policy) (*alerts.Policy, error)
}

// Alerts - NewRelic Alerts wrapper
type Alerts struct {
	client Client
}

// NewAlerts - Creates an Alerts client wrapper
func NewAlerts(client Client) Alerts {
	return Alerts{client}
}

// Find -- Search for policy
func (f *Alerts) Find(name string) (*alerts.Policy, error) {
	policyList, err := f.client.ListPolicies(&alerts.ListPoliciesParams {
		Name: name,
	})

	if err != nil {
		zap.S().Error("Error fetching list of dashboards", err)
		return nil, err
	}

	if (len(policyList) == 1) {
		return &policyList[0], nil
	}
	return nil, nil
}

// Manage - This will handle update or creating
func (f *Alerts) Manage(_policy *alerts.Policy) (*alerts.Policy, error) {
	if _policy != nil {
		policy, err := f.client.UpdatePolicy(*_policy)
		return policy, err
	}

	policy, err := f.client.CreatePolicy(*_policy)
	return policy, err
}