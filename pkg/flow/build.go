package flow

import (
	"errors"
)

// Options to build a ExistenceDependentOperation.
type Options struct {
	TargetObjectName string
	Execute          execute
	ExistenceCheck   ExistenceChecker
	ExecOnExistence  bool
}

// NewOperation build an ExistenceDependentOperation by the definition given as Options.
func NewOperation(o *Options) *ExistenceDependentOperation {
	if o.Execute == nil {
		return &ExistenceDependentOperation{
			buildErr: errors.New("missing required options 'Execute'"),
		}
	}

	// TODO: add further validations

	return &ExistenceDependentOperation{
		targetObjectName: o.TargetObjectName,
		execute:          o.Execute,
		existenceChecker: o.ExistenceCheck,
		execOnExistence:  o.ExecOnExistence,
	}
}
