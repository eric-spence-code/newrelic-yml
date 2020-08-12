package main

import (
	"fmt"
	"github.com/eric-spence-code/newrelic-yml/internal/parser"
	"github.com/eric-spence-code/newrelic-yml/internal/templating"
)

var testVariablesYaml = `
variables:
  dev:
    key: 1234
    role: dev-role
  test:
    key: 9876
    role: test-role
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

func main() {
	variables := parser.ParseVariables(testVariablesYaml)

	output := templating.Inject(testDataYaml, "dev", variables.Variables)

	fmt.Print(output)
}
