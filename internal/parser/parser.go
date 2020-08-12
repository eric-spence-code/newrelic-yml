package parser

import (
	"github.com/ghodss/yaml"
	"github.com/newrelic/newrelic-client-go/pkg/alerts"
	"github.com/newrelic/newrelic-client-go/pkg/dashboards"
)

type PolicyCondition struct {
	Policy     alerts.Policy      `json:"policy"`
	Conditions []alerts.Condition `json:"conditions"`
}

type NewRelicYaml struct {
	Policies  []PolicyCondition      `json:"policies"`
	Dashboard []dashboards.Dashboard `json:"dashboard"`
}

type VariablesYml struct {
	Variables map[string](map[string]interface{}) `json:"variables"`
}

// Parse -- Parse out Newrelic data
func Parse(data string) (newRelicYaml NewRelicYaml) {
	err := yaml.Unmarshal([]byte(data), &newRelicYaml)
	if err != nil {
		panic(err)
	}
	return newRelicYaml
}

// ParseVariables -- Parse variables out
func ParseVariables(data string) (variables VariablesYml) {
	err := yaml.Unmarshal([]byte(data), &variables)
	if err != nil {
		panic(err)
	}
	return variables
}
