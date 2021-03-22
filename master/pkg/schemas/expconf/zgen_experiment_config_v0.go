// Code generated by gen.py. DO NOT EDIT.

package expconf

import (
	"github.com/santhosh-tekuri/jsonschema/v2"

	"github.com/determined-ai/determined/master/pkg/schemas"
)

func (e ExperimentConfigV0) GetBindMounts() BindMountsConfigV0 {
	if e.BindMounts == nil {
		panic("You must call WithDefaults on ExperimentConfigV0 before .GetBindMounts")
	}
	return *e.BindMounts
}

func (e ExperimentConfigV0) GetCheckpointPolicy() string {
	if e.CheckpointPolicy == nil {
		panic("You must call WithDefaults on ExperimentConfigV0 before .GetCheckpointPolicy")
	}
	return *e.CheckpointPolicy
}

func (e ExperimentConfigV0) GetCheckpointStorage() *CheckpointStorageConfigV0 {
	return e.CheckpointStorage
}

func (e ExperimentConfigV0) GetDataLayer() DataLayerConfigV0 {
	if e.DataLayer == nil {
		panic("You must call WithDefaults on ExperimentConfigV0 before .GetDataLayer")
	}
	return *e.DataLayer
}

func (e ExperimentConfigV0) GetData() map[string]interface{} {
	if e.Data == nil {
		panic("You must call WithDefaults on ExperimentConfigV0 before .GetData")
	}
	return *e.Data
}

func (e ExperimentConfigV0) GetDebug() bool {
	if e.Debug == nil {
		panic("You must call WithDefaults on ExperimentConfigV0 before .GetDebug")
	}
	return *e.Debug
}

func (e ExperimentConfigV0) GetDescription() *string {
	return e.Description
}

func (e ExperimentConfigV0) GetEntrypoint() *string {
	return e.Entrypoint
}

func (e ExperimentConfigV0) GetEnvironment() EnvironmentConfigV0 {
	if e.Environment == nil {
		panic("You must call WithDefaults on ExperimentConfigV0 before .GetEnvironment")
	}
	return *e.Environment
}

func (e ExperimentConfigV0) GetInternal() *InternalConfigV0 {
	return e.Internal
}

func (e ExperimentConfigV0) GetLabels() LabelsV0 {
	if e.Labels == nil {
		panic("You must call WithDefaults on ExperimentConfigV0 before .GetLabels")
	}
	return *e.Labels
}

func (e ExperimentConfigV0) GetMaxRestarts() int {
	if e.MaxRestarts == nil {
		panic("You must call WithDefaults on ExperimentConfigV0 before .GetMaxRestarts")
	}
	return *e.MaxRestarts
}

func (e ExperimentConfigV0) GetMinCheckpointPeriod() LengthV0 {
	if e.MinCheckpointPeriod == nil {
		panic("You must call WithDefaults on ExperimentConfigV0 before .GetMinCheckpointPeriod")
	}
	return *e.MinCheckpointPeriod
}

func (e ExperimentConfigV0) GetMinValidationPeriod() LengthV0 {
	if e.MinValidationPeriod == nil {
		panic("You must call WithDefaults on ExperimentConfigV0 before .GetMinValidationPeriod")
	}
	return *e.MinValidationPeriod
}

func (e ExperimentConfigV0) GetOptimizations() OptimizationsConfigV0 {
	if e.Optimizations == nil {
		panic("You must call WithDefaults on ExperimentConfigV0 before .GetOptimizations")
	}
	return *e.Optimizations
}

func (e ExperimentConfigV0) GetPerformInitialValidation() bool {
	if e.PerformInitialValidation == nil {
		panic("You must call WithDefaults on ExperimentConfigV0 before .GetPerformInitialValidation")
	}
	return *e.PerformInitialValidation
}

func (e ExperimentConfigV0) GetRecordsPerEpoch() int {
	if e.RecordsPerEpoch == nil {
		panic("You must call WithDefaults on ExperimentConfigV0 before .GetRecordsPerEpoch")
	}
	return *e.RecordsPerEpoch
}

func (e ExperimentConfigV0) GetReproducibility() ReproducibilityConfigV0 {
	if e.Reproducibility == nil {
		panic("You must call WithDefaults on ExperimentConfigV0 before .GetReproducibility")
	}
	return *e.Reproducibility
}

func (e ExperimentConfigV0) GetResources() ResourcesConfigV0 {
	if e.Resources == nil {
		panic("You must call WithDefaults on ExperimentConfigV0 before .GetResources")
	}
	return *e.Resources
}

func (e ExperimentConfigV0) GetSchedulingUnit() int {
	if e.SchedulingUnit == nil {
		panic("You must call WithDefaults on ExperimentConfigV0 before .GetSchedulingUnit")
	}
	return *e.SchedulingUnit
}

func (e ExperimentConfigV0) GetSecurity() *SecurityConfigV0 {
	return e.Security
}

func (e ExperimentConfigV0) GetTensorboardStorage() *TensorboardStorageConfigV0 {
	return e.TensorboardStorage
}

func (e ExperimentConfigV0) GetHyperparameters() HyperparametersV0 {
	return e.Hyperparameters
}

func (e ExperimentConfigV0) GetSearcher() SearcherConfigV0 {
	return e.Searcher
}

func (e ExperimentConfigV0) WithDefaults() ExperimentConfigV0 {
	return schemas.WithDefaults(e).(ExperimentConfigV0)
}

func (e ExperimentConfigV0) Merge(other ExperimentConfigV0) ExperimentConfigV0 {
	return schemas.Merge(e, other).(ExperimentConfigV0)
}

func (e ExperimentConfigV0) ParsedSchema() interface{} {
	return schemas.ParsedExperimentConfigV0()
}

func (e ExperimentConfigV0) SanityValidator() *jsonschema.Schema {
	return schemas.GetSanityValidator("http://determined.ai/schemas/expconf/v0/experiment.json")
}

func (e ExperimentConfigV0) CompletenessValidator() *jsonschema.Schema {
	return schemas.GetCompletenessValidator("http://determined.ai/schemas/expconf/v0/experiment.json")
}
