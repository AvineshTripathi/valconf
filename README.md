# ValConf

The project is a kubernetes config validator that uses the custom rego policy to validate kubernetes objects. For now we support objects like pods, deployements, services, service accounts, roles. Custom rego policy checks for the fields in their respective config files are and returns any violations.

Config files are retrieved from the runnning k8s cluster using the kubernetes client-go library. 

## Working

main.go file gets list of objects and calls the object.<function_name> func from the object pkg. These func are responsible return array of map([]map[string]interface{}). The returned array consists of respective lists of objects along with fields from their configs that are to be validated against the rego.

All the returned values are then stored in a map(called as input) which is then sent to the rego Engine present in the validate pkg. The rego Engine uses the utils.RegoUtils variable as a policy and validates the config.

## Scope of Improvements

- The current codebase is slow and can be improved using the golang threads.
- Validation for now is based on the fields of connfig files and hence is not the perfect solution to detect complex misconfigurations.