package infrastructure

import (
	"fmt"
	"strings"

	"mosho-raspi/domain"
)

type MultiEnvRecorder struct {
	Recorders []domain.EnvRecorder
}

type MultiEnvRecorderError struct {
	errors []error
}

func (e *MultiEnvRecorderError) Error() string {
	strerrs := []string{}
	for _, err := range e.errors {
		strerrs = append(strerrs, err.Error())
	}
	return fmt.Sprintf("%d error(s): %s", len(e.errors), strings.Join(strerrs, ", "))
}

func (r *MultiEnvRecorder) RecordEnv(env *domain.Env) error {
	errs := []error{}
	for _, recorder := range r.Recorders {
		err := recorder.RecordEnv(env)
		if err != nil {
			errs = append(errs, err)
		}
	}
	if len(errs) > 0 {
		return &MultiEnvRecorderError{errs}
	}
	return nil
}
