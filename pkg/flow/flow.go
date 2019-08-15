// Package flow provides abstractions for execution flow, e.g. executing code depending on the result of another operation which was done before.
package flow

import "github.com/pkg/errors"

// ExistenceDependentOperation provides an abstraction to execute an operation depending on the existence or absence of the target object (e.g. create kubernetes deployment if it does not yet exist).
type ExistenceDependentOperation struct {
	targetObjectName string
	execute          execute
	existenceChecker ExistenceChecker
	execOnExistence  bool

	// buildErr holds a possible error which occurred while building the object
	buildErr error
}

// NamedObject represents an object with a name.
type NamedObject interface {
	GetName() string
}

type execute func(targetName string) error

// ExistenceChecker allows to fetch an object from a backend (e.g. a REST API where a kubernetes deployment is fetched via a GET request).
type ExistenceChecker interface {

	// Get fetches an object of the given name.
	Get(name string) (NamedObject, error)

	// IsNotFoundError is the hook were one can provide the check if a backend error which occurred during Get represents a NOT FOUND error kind.
	IsNotFoundError(err error) bool
}

// Run executes the operation.
func (o *ExistenceDependentOperation) Run() error {
	if o.buildErr != nil {
		return errors.Wrap(o.buildErr, "could not build operation")
	}

	objExists, err := o.resourceExists(o.existenceChecker.IsNotFoundError)
	if err != nil {
		return errors.Wrap(err, "check if resource exists failed")
	}

	if o.shouldExec(objExists) {
		return o.execute(o.targetObjectName)
	}

	return nil
}

func (o *ExistenceDependentOperation) resourceExists(isNotFound func(error) bool) (bool, error) {
	namedObj, err := o.existenceChecker.Get(o.targetObjectName)
	if err != nil {
		if isNotFound(err) {
			return false, nil
		}
		return true, err
	}

	if namedObj == nil {
		return false, nil
	}

	return true, nil
}

func (o *ExistenceDependentOperation) shouldExec(objectExists bool) bool {
	if o.execOnExistence {
		return objectExists
	}
	return !objectExists
}
