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
	Variables map[string](map[string]interface{}) `json:"variables"`
	Policies  []PolicyCondition                   `json:"policies"`
	Dashboard []dashboards.Dashboard              `json:"dashboard"`
}

func Parse(data string) (newRelicYaml NewRelicYaml) {
	err := yaml.Unmarshal([]byte(data), &newRelicYaml)
	if err != nil {
		panic(err)
	}
	return newRelicYaml
}
