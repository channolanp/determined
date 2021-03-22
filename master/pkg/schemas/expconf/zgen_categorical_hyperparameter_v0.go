// Code generated by gen.py. DO NOT EDIT.

package expconf

import (
	"github.com/santhosh-tekuri/jsonschema/v2"

	"github.com/determined-ai/determined/master/pkg/schemas"
)

func (c CategoricalHyperparameterV0) GetVals() []interface{} {
	return c.Vals
}

func (c CategoricalHyperparameterV0) WithDefaults() CategoricalHyperparameterV0 {
	return schemas.WithDefaults(c).(CategoricalHyperparameterV0)
}

func (c CategoricalHyperparameterV0) Merge(other CategoricalHyperparameterV0) CategoricalHyperparameterV0 {
	return schemas.Merge(c, other).(CategoricalHyperparameterV0)
}

func (c CategoricalHyperparameterV0) ParsedSchema() interface{} {
	return schemas.ParsedCategoricalHyperparameterV0()
}

func (c CategoricalHyperparameterV0) SanityValidator() *jsonschema.Schema {
	return schemas.GetSanityValidator("http://determined.ai/schemas/expconf/v0/hyperparameter-categorical.json")
}

func (c CategoricalHyperparameterV0) CompletenessValidator() *jsonschema.Schema {
	return schemas.GetCompletenessValidator("http://determined.ai/schemas/expconf/v0/hyperparameter-categorical.json")
}
