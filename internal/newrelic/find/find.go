package find

import (
	"os"
	"go.uber.org/zap"
	"github.com/newrelic/newrelic-client-go/newrelic"
	"github.com/newrelic/newrelic-client-go/pkg/alerts"
	"github.com/newrelic/newrelic-client-go/pkg/dashboards"
)

func getClient() (*newrelic.NewRelic) {
	client, err := newrelic.New(newrelic.ConfigPersonalAPIKey(os.Getenv("NEW_RELIC_API_KEY")))
	if err != nil {
		zap.S().Error("error initializing client", err)
		panic(err)
	}

	return client
}
 
type Client interface {
	ListPolicies(*alerts.ListPoliciesParams) ([]alerts.Policy, error)
}

type Find struct {
	client Client
}

func New(client Client) Find {
	return Find{client}
}

// Dashboard -- Search for policy
func Dashboard(name string) (*dashboards.Dashboard) {
	dashList, err := getClient().Dashboards.ListDashboards(&dashboards.ListDashboardsParams {
		Title: name,
	})

	if err != nil {
		zap.S().Error("Error fetching list of policies", err)
		panic(err)
	}

	if len(dashList) == 1 {
		return dashList[0]
	}
	
	return nil
}


// Policy -- Search for policy
func (f *Find) Policy(name string) (*alerts.Policy) {
	policyList, err := f.client.ListPolicies(&alerts.ListPoliciesParams {
		Name: name,
	})

	if err != nil {
		zap.S().Error("Error fetching list of dashboards", err)
		panic(err)
	}

	if (len(policyList) == 1) {
		return &policyList[0]
	}
	return nil
}