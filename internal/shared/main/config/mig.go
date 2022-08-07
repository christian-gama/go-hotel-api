package config

import (
	"github.com/christian-gama/go-booking-api/internal/shared/infra/configger"
)

const (
	migDir = "MIG_DIR"
	migExt = "MIG_EXT"
)

type mig struct{}

// Dir returns the directory of the migration.
func (m *mig) Dir() string {
	return getEnv(migDir)
}

// Ext returns the file extension of the migration.
func (m *mig) Ext() string {
	return getEnv(migExt)
}

// NewMig returns the migration configuration.
func NewMig() configger.Mig {
	return &mig{}
}
