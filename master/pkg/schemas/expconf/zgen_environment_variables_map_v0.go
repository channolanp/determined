// Code generated by gen.py. DO NOT EDIT.

package expconf

import (
	"github.com/santhosh-tekuri/jsonschema/v2"

	"github.com/determined-ai/determined/master/pkg/schemas"
)

func (e EnvironmentVariablesMapV0) GetCPU() []string {
	if e.CPU == nil {
		panic("You must call WithDefaults on EnvironmentVariablesMapV0 before .GetCPU")
	}
	return *e.CPU
}

func (e EnvironmentVariablesMapV0) GetGPU() []string {
	if e.GPU == nil {
		panic("You must call WithDefaults on EnvironmentVariablesMapV0 before .GetGPU")
	}
	return *e.GPU
}

func (e EnvironmentVariablesMapV0) WithDefaults() EnvironmentVariablesMapV0 {
	return schemas.WithDefaults(e).(EnvironmentVariablesMapV0)
}

func (e EnvironmentVariablesMapV0) Merge(other EnvironmentVariablesMapV0) EnvironmentVariablesMapV0 {
	return schemas.Merge(e, other).(EnvironmentVariablesMapV0)
}

func (e EnvironmentVariablesMapV0) ParsedSchema() interface{} {
	return schemas.ParsedEnvironmentVariablesMapV0()
}

func (e EnvironmentVariablesMapV0) SanityValidator() *jsonschema.Schema {
	return schemas.GetSanityValidator("http://determined.ai/schemas/expconf/v0/environment-variables-map.json")
}

func (e EnvironmentVariablesMapV0) CompletenessValidator() *jsonschema.Schema {
	return schemas.GetCompletenessValidator("http://determined.ai/schemas/expconf/v0/environment-variables-map.json")
}
