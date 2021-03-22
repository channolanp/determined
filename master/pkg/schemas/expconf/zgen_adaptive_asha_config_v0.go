// Code generated by gen.py. DO NOT EDIT.

package expconf

import (
	"github.com/santhosh-tekuri/jsonschema/v2"

	"github.com/determined-ai/determined/master/pkg/schemas"
)

func (a AdaptiveASHAConfigV0) GetMetric() string {
	return a.Metric
}

func (a AdaptiveASHAConfigV0) GetSmallerIsBetter() bool {
	if a.SmallerIsBetter == nil {
		panic("You must call WithDefaults on AdaptiveASHAConfigV0 before .GetSmallerIsBetter")
	}
	return *a.SmallerIsBetter
}

func (a AdaptiveASHAConfigV0) GetSourceTrialID() *int {
	return a.SourceTrialID
}

func (a AdaptiveASHAConfigV0) GetSourceCheckpointUUID() *string {
	return a.SourceCheckpointUUID
}

func (a AdaptiveASHAConfigV0) GetMaxLength() LengthV0 {
	return a.MaxLength
}

func (a AdaptiveASHAConfigV0) GetMaxTrials() int {
	return a.MaxTrials
}

func (a AdaptiveASHAConfigV0) GetBracketRungs() []int {
	if a.BracketRungs == nil {
		panic("You must call WithDefaults on AdaptiveASHAConfigV0 before .GetBracketRungs")
	}
	return *a.BracketRungs
}

func (a AdaptiveASHAConfigV0) GetDivisor() float64 {
	if a.Divisor == nil {
		panic("You must call WithDefaults on AdaptiveASHAConfigV0 before .GetDivisor")
	}
	return *a.Divisor
}

func (a AdaptiveASHAConfigV0) GetMode() AdaptiveMode {
	if a.Mode == nil {
		panic("You must call WithDefaults on AdaptiveASHAConfigV0 before .GetMode")
	}
	return *a.Mode
}

func (a AdaptiveASHAConfigV0) GetMaxRungs() int {
	if a.MaxRungs == nil {
		panic("You must call WithDefaults on AdaptiveASHAConfigV0 before .GetMaxRungs")
	}
	return *a.MaxRungs
}

func (a AdaptiveASHAConfigV0) GetMaxConcurrentTrials() int {
	if a.MaxConcurrentTrials == nil {
		panic("You must call WithDefaults on AdaptiveASHAConfigV0 before .GetMaxConcurrentTrials")
	}
	return *a.MaxConcurrentTrials
}

func (a AdaptiveASHAConfigV0) WithDefaults() AdaptiveASHAConfigV0 {
	return schemas.WithDefaults(a).(AdaptiveASHAConfigV0)
}

func (a AdaptiveASHAConfigV0) Merge(other AdaptiveASHAConfigV0) AdaptiveASHAConfigV0 {
	return schemas.Merge(a, other).(AdaptiveASHAConfigV0)
}

func (a AdaptiveASHAConfigV0) ParsedSchema() interface{} {
	return schemas.ParsedAdaptiveASHAConfigV0()
}

func (a AdaptiveASHAConfigV0) SanityValidator() *jsonschema.Schema {
	return schemas.GetSanityValidator("http://determined.ai/schemas/expconf/v0/searcher-adaptive-asha.json")
}

func (a AdaptiveASHAConfigV0) CompletenessValidator() *jsonschema.Schema {
	return schemas.GetCompletenessValidator("http://determined.ai/schemas/expconf/v0/searcher-adaptive-asha.json")
}
