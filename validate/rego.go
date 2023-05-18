package validate

import (
	"context"
	"fmt"

	"github.com/AvineshTripathi/valconf/utils"
	"github.com/open-policy-agent/opa/ast"
	"github.com/open-policy-agent/opa/rego"
)


func Rego(ctx context.Context, input map[string]interface{}) {
	policyKey := fmt.Sprintf("%s.rego", "eval")
	compiler, err := ast.CompileModules(map[string]string{
		policyKey: utils.RegoUtils,
	})
	if err != nil {
		fmt.Println(err)
	}

	pkgName := compiler.Modules[policyKey].Package.Path.String()

	regoFunc := make([]func(r *rego.Rego), 0)

	regoFunc = append(regoFunc, rego.Query(fmt.Sprintf("%s", pkgName)))

	regoFunc = append(regoFunc, rego.Compiler(compiler))
	regoFunc = append(regoFunc, rego.Input(input))

	rego := rego.New(regoFunc...)

	res, err := rego.Eval(ctx)

	if err != nil {
		fmt.Println(err)
	}
	utils.Logger(res)

}