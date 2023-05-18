package utils

import (
	"encoding/json"
	"fmt"

	"github.com/open-policy-agent/opa/rego"
)

func Logger(res rego.ResultSet) {

	jsonString, err := json.MarshalIndent(res[0].Expressions, "", "  ")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(jsonString))
}