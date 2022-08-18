package config

const (
	migDir = "MIG_DIR"
	migExt = "MIG_EXT"
)

type (
	// Mig is the migration configuration.
	Mig interface {
		Dir() string
		Ext() string
	}

	// migImpl is the implementation of the Mig interface.
	migImpl struct{}
)

// Dir returns the directory of the migration.
func (m *migImpl) Dir() string {
	return getEnv(migDir)
}

// Ext returns the file extension of the migration.
func (m *migImpl) Ext() string {
	return getEnv(migExt)
}

// NewMig returns the migration configuration.
func NewMig() Mig {
	return &migImpl{}
}
