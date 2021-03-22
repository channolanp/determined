// Code generated by gen.py. DO NOT EDIT.

package expconf

import (
	"github.com/santhosh-tekuri/jsonschema/v2"

	"github.com/determined-ai/determined/master/pkg/schemas"
)

func (e EnvironmentImageMapV0) GetCPU() *string {
	return e.CPU
}

func (e EnvironmentImageMapV0) GetGPU() *string {
	return e.GPU
}

func (e EnvironmentImageMapV0) WithDefaults() EnvironmentImageMapV0 {
	return schemas.WithDefaults(e).(EnvironmentImageMapV0)
}

func (e EnvironmentImageMapV0) Merge(other EnvironmentImageMapV0) EnvironmentImageMapV0 {
	return schemas.Merge(e, other).(EnvironmentImageMapV0)
}

func (e EnvironmentImageMapV0) ParsedSchema() interface{} {
	return schemas.ParsedEnvironmentImageMapV0()
}

func (e EnvironmentImageMapV0) SanityValidator() *jsonschema.Schema {
	return schemas.GetSanityValidator("http://determined.ai/schemas/expconf/v0/environment-image-map.json")
}

func (e EnvironmentImageMapV0) CompletenessValidator() *jsonschema.Schema {
	return schemas.GetCompletenessValidator("http://determined.ai/schemas/expconf/v0/environment-image-map.json")
}
