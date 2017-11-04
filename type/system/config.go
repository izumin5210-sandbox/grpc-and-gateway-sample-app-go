package system

import (
	"github.com/creasty/apperrors"
	"github.com/creasty/configo"
)

// Config stores general setting parameters that are loaded from
// enviroment variables, a dotenv file, and yaml files
type Config struct {
	Env             string `envconfig:"env" valid:"required"`
	Host            string `envconfig:"host" valid:"required"`
	SentryKey       string `envconfig:"sentry_key"`
	SentrySecret    string `envconfig:"sentry_secret"`
	SentryProjectID string `envconfig:"sentry_project_id"`
}

func loadConfig() (*Config, error) {
	c := &Config{}
	if err := configo.Load(c, configo.Option{}); err != nil {
		return nil, apperrors.WithMessage(err, "failed to load configo.Load()")
	}
	return c, nil
}
