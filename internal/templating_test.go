package templating

import (
	"testing"
)

var testDataYaml = `
Test -- {{.env}}
`

func TestVariableParse(t *testing.T) {
	testString := "Test -- {{.env}}"

}