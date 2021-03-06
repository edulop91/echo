package web

import (
	env "github.com/blendlabs/go-util/env"
	uuid "github.com/blendlabs/go-util/uuid"
)

// ViewModel is a wrapping viewmodel.
type ViewModel struct {
	Ctx       *Ctx
	ViewModel interface{}
}

// HasEnv returns if an env var is set.
func (vm *ViewModel) HasEnv(key string) bool {
	return env.Env().HasVar(key)
}

// Env returns a value from the environment.
func (vm *ViewModel) Env(key string, defaults ...string) string {
	return env.Env().String(key, defaults...)
}

// UUIDv4 returns a uuidv4 as a string.
func (vm *ViewModel) UUIDv4() string {
	return uuid.V4().String()
}
