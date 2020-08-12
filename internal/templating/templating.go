package templating

import (
	"bytes"
	"text/template"
)

// Inject env variables into string
func Inject(data string, key string, variables map[string](map[string]interface{}) ) (val string)  {
	envVariables := variables[key]
	tmpl, err := template.New("data").Parse(data)
	if err != nil { 
		panic(err) 
	}
	var tpl bytes.Buffer
    err = tmpl.Execute(&tpl, envVariables)
	if err != nil { 
		panic(err) 
	}
	return tpl.String()
}