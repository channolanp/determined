// Code generated by gen.py. DO NOT EDIT.

package expconf

import (
	"github.com/santhosh-tekuri/jsonschema/v2"

	"github.com/determined-ai/determined/master/pkg/schemas"
)

func (s SearcherConfigV0) GetUnionMember() interface{} {
	if s.SingleConfig != nil {
		return nil
	}
	if s.RandomConfig != nil {
		return nil
	}
	if s.GridConfig != nil {
		return nil
	}
	if s.SyncHalvingConfig != nil {
		return nil
	}
	if s.AsyncHalvingConfig != nil {
		return nil
	}
	if s.AdaptiveConfig != nil {
		return nil
	}
	if s.AdaptiveSimpleConfig != nil {
		return nil
	}
	if s.AdaptiveASHAConfig != nil {
		return nil
	}
	if s.PBTConfig != nil {
		return nil
	}
	panic("no union member defined")
}

func (s SearcherConfigV0) GetMetric() string {
	if s.SingleConfig != nil {
		return s.SingleConfig.GetMetric()
	}
	if s.RandomConfig != nil {
		return s.RandomConfig.GetMetric()
	}
	if s.GridConfig != nil {
		return s.GridConfig.GetMetric()
	}
	if s.SyncHalvingConfig != nil {
		return s.SyncHalvingConfig.GetMetric()
	}
	if s.AsyncHalvingConfig != nil {
		return s.AsyncHalvingConfig.GetMetric()
	}
	if s.AdaptiveConfig != nil {
		return s.AdaptiveConfig.GetMetric()
	}
	if s.AdaptiveSimpleConfig != nil {
		return s.AdaptiveSimpleConfig.GetMetric()
	}
	if s.AdaptiveASHAConfig != nil {
		return s.AdaptiveASHAConfig.GetMetric()
	}
	if s.PBTConfig != nil {
		return s.PBTConfig.GetMetric()
	}
	panic("no union member defined")
}

func (s SearcherConfigV0) GetSmallerIsBetter() bool {
	if s.SingleConfig != nil {
		return s.SingleConfig.GetSmallerIsBetter()
	}
	if s.RandomConfig != nil {
		return s.RandomConfig.GetSmallerIsBetter()
	}
	if s.GridConfig != nil {
		return s.GridConfig.GetSmallerIsBetter()
	}
	if s.SyncHalvingConfig != nil {
		return s.SyncHalvingConfig.GetSmallerIsBetter()
	}
	if s.AsyncHalvingConfig != nil {
		return s.AsyncHalvingConfig.GetSmallerIsBetter()
	}
	if s.AdaptiveConfig != nil {
		return s.AdaptiveConfig.GetSmallerIsBetter()
	}
	if s.AdaptiveSimpleConfig != nil {
		return s.AdaptiveSimpleConfig.GetSmallerIsBetter()
	}
	if s.AdaptiveASHAConfig != nil {
		return s.AdaptiveASHAConfig.GetSmallerIsBetter()
	}
	if s.PBTConfig != nil {
		return s.PBTConfig.GetSmallerIsBetter()
	}
	panic("no union member defined")
}

func (s SearcherConfigV0) GetSourceCheckpointUUID() *string {
	if s.SingleConfig != nil {
		return s.SingleConfig.GetSourceCheckpointUUID()
	}
	if s.RandomConfig != nil {
		return s.RandomConfig.GetSourceCheckpointUUID()
	}
	if s.GridConfig != nil {
		return s.GridConfig.GetSourceCheckpointUUID()
	}
	if s.SyncHalvingConfig != nil {
		return s.SyncHalvingConfig.GetSourceCheckpointUUID()
	}
	if s.AsyncHalvingConfig != nil {
		return s.AsyncHalvingConfig.GetSourceCheckpointUUID()
	}
	if s.AdaptiveConfig != nil {
		return s.AdaptiveConfig.GetSourceCheckpointUUID()
	}
	if s.AdaptiveSimpleConfig != nil {
		return s.AdaptiveSimpleConfig.GetSourceCheckpointUUID()
	}
	if s.AdaptiveASHAConfig != nil {
		return s.AdaptiveASHAConfig.GetSourceCheckpointUUID()
	}
	if s.PBTConfig != nil {
		return s.PBTConfig.GetSourceCheckpointUUID()
	}
	panic("no union member defined")
}

func (s SearcherConfigV0) GetSourceTrialID() *int {
	if s.SingleConfig != nil {
		return s.SingleConfig.GetSourceTrialID()
	}
	if s.RandomConfig != nil {
		return s.RandomConfig.GetSourceTrialID()
	}
	if s.GridConfig != nil {
		return s.GridConfig.GetSourceTrialID()
	}
	if s.SyncHalvingConfig != nil {
		return s.SyncHalvingConfig.GetSourceTrialID()
	}
	if s.AsyncHalvingConfig != nil {
		return s.AsyncHalvingConfig.GetSourceTrialID()
	}
	if s.AdaptiveConfig != nil {
		return s.AdaptiveConfig.GetSourceTrialID()
	}
	if s.AdaptiveSimpleConfig != nil {
		return s.AdaptiveSimpleConfig.GetSourceTrialID()
	}
	if s.AdaptiveASHAConfig != nil {
		return s.AdaptiveASHAConfig.GetSourceTrialID()
	}
	if s.PBTConfig != nil {
		return s.PBTConfig.GetSourceTrialID()
	}
	panic("no union member defined")
}

func (s SearcherConfigV0) WithDefaults() SearcherConfigV0 {
	return schemas.WithDefaults(s).(SearcherConfigV0)
}

func (s SearcherConfigV0) Merge(other SearcherConfigV0) SearcherConfigV0 {
	return schemas.Merge(s, other).(SearcherConfigV0)
}

func (s SearcherConfigV0) ParsedSchema() interface{} {
	return schemas.ParsedSearcherConfigV0()
}

func (s SearcherConfigV0) SanityValidator() *jsonschema.Schema {
	return schemas.GetSanityValidator("http://determined.ai/schemas/expconf/v0/searcher.json")
}

func (s SearcherConfigV0) CompletenessValidator() *jsonschema.Schema {
	return schemas.GetCompletenessValidator("http://determined.ai/schemas/expconf/v0/searcher.json")
}
