package parser

import (
	"testing"
)

var testVariablesYaml = `
variables:
  dev:
    key: 1234
    role: blue
  test:
    key: 9876
    role: green
`

var testDataYaml = `
policies:
  - policy:
      name: Test Policy 
    conditions:
      - name: Test Condition One
        enabled: true
dashboard:
  - title: New Dashboard
    id: 1234
  - title: New Dashboard 2
`

func TestVariableParse(t *testing.T) {
	devVariableValue := float64(1234)
	devVariableRoleValue := "blue"
	testVariableValue := float64(9876)
	got := ParseVariables(testVariablesYaml)
	if got.Variables["dev"]["key"] != devVariableValue {
		t.Errorf("Dev Key :: Want: %g || Got: %g", devVariableValue, got.Variables["dev"]["key"])
	}

	if got.Variables["dev"]["role"] != devVariableRoleValue {
		t.Errorf("Dev Key :: Want: %s || Got: %s", devVariableRoleValue, got.Variables["dev"]["role"])
	}

	if got.Variables["test"]["key"] != testVariableValue {
		t.Errorf("Test Key :: Want: %g || Got: %g", testVariableValue, got.Variables["test"]["key"])
	}
}

func TestDashboardParse(t *testing.T) {
	wantTitle := "New Dashboard"
	wantId := 1234
	got := Parse(testDataYaml)
	gotDash := got.Dashboard[0]
	if gotDash.Title != wantTitle {
		t.Errorf("Title :: Want: %s || Got: %s", wantTitle, gotDash.Title)
	}
	if gotDash.ID != wantId {
		t.Errorf("ID :: Want: %d || Got: %d", wantId, gotDash.ID)
	}
}

func TestPolicyParse(t *testing.T) {
	wantTitle := "Test Policy"
	wantCondition := "Test Condition One"
	got := Parse(testDataYaml)
	gotPolicy := got.Policies[0]
	if gotPolicy.Policy.Name != wantTitle {
		t.Errorf("Name :: Want: %s || Got: %s", wantTitle, gotPolicy.Policy.Name)
	}
	if gotPolicy.Conditions[0].Name != wantCondition {
		t.Errorf("Name :: Want: %s || Got: %s", wantCondition, gotPolicy.Conditions[0].Name)
	}
	if gotPolicy.Conditions[0].Enabled != true {
		t.Errorf("Name :: Want: %t || Got: %t", true, gotPolicy.Conditions[0].Enabled)
	}
}
