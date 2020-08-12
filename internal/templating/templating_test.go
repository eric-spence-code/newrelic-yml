package templating

import (
	"testing"
)

var testDataYaml = `
Test -- {{.env}}
`

func TestVariableParse(t *testing.T) {
	testString := "Env -- {{.env}}"
	wantString := "Env -- dev"

	testData := make(map[string]map[string]interface{})

	devData := map[string]interface{}{ "env" : "dev"}

	testData["dev"] = devData
	
	if val := Inject(testString, "dev", testData); val != wantString {
		t.Errorf("Inject:: Want: %s || Got: %s", wantString, val)
	}
}