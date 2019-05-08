package usecase

import (
	"fmt"

	"mosho-raspi/domain"
)

type RecordEnv struct {
	EnvGetter   domain.EnvGetter
	EnvRecorder domain.EnvRecorder
}

func (u *RecordEnv) Invoke() error {
	env, err := u.EnvGetter.GetEnv()
	if err != nil {
		return fmt.Errorf("Failed to get env: %s", err)
	}
	err = u.EnvRecorder.RecordEnv(env)
	return err
}
