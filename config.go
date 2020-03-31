package mongodb

import (
	"emperror.dev/errors"
)

// Config holds information necessary for connecting to a database.
type Config struct {
	Hosts      []string
	ReplicaSet *string
	User       string
	Pass       string
	Name       string

	Params map[string]string
}

// Validate checks that the configuration is valid.
func (c Config) Validate() error {
	if len(c.Hosts) == 0 {
		return errors.New("database hosts is required")
	}

	if c.User == "" {
		return errors.New("database user is required")
	}

	if c.Name == "" {
		return errors.New("database name is required")
	}

	return nil
}