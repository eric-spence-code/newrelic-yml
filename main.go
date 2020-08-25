package main

import (
	"fmt"

	newrelicWrapper "github.com/eric-spence-code/newrelic-yml/internal/newrelic"
	"github.com/eric-spence-code/newrelic-yml/internal/parser"
	"github.com/eric-spence-code/newrelic-yml/internal/templating"
	"github.com/newrelic/newrelic-client-go/newrelic"
	"github.com/newrelic/newrelic-client-go/pkg/alerts"
)

var testVariablesYaml = `
variables:
  dev:
    key: 1234
    role: dev
  test:
    key: 9876
    role: test
`

var testDataYaml = `
policies:
  - policy:
      name: Test Policy {{.role}}
    conditions:
      - name: Test Condition One
        enabled: true
dashboard:
  - title: New Dashboard
    id: 1234
  - title: New Dashboard 2
`

// var client, err = newrelic.New(
// 	newrelic.ConfigPersonalAPIKey(os.Getenv("NEW_RELIC_API_KEY")),
// 	newrelic.ConfigRegion(os.Getenv("NEW_RELIC_ENV")),
// 	newrelic.ConfigLogLevel(os.Getenv("NEW_RELIC_LOG_LEVEL")),
// )

var client, err = newrelic.New(
	newrelic.ConfigPersonalAPIKey("Foo"),
	newrelic.ConfigRegion("US"),
	newrelic.ConfigLogLevel("DEBUG"),
)


func main() {
	if err != nil {
		panic(err)
	}

	variables := parser.ParseVariables(testVariablesYaml)

	output := templating.Inject(testDataYaml, "dev", variables.Variables)

	newrelicObj := parser.Parse(output)

	// Create Find client
	alertHelper := newrelicWrapper.NewAlerts(&client.Alerts)

	for _, policyYml := range newrelicObj.Policies {
		// Find Policy
		policy, err := alertHelper.Find(policyYml.Policy.Name)
		policy, err = alertHelper.Manage(&policyYml.Policy)

		if err != nil {
			fmt.Print(err)
		}

		if policy != nil {
			fmt.Print(policy)
		}

		for _, condition := range policyYml.Conditions {
			fmt.Print(condition.Name)
		}
	}
}
