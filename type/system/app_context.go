package system

import (
	"github.com/creasty/apperrors"
)

// AppContext represents a context of application
type AppContext struct {
	*Config
}

// CreateAppContext loads application configs and creates objects to use in the application lifecycle
func CreateAppContext() (*AppContext, error) {
	cnf, err := loadConfig()
	if err != nil {
		return nil, apperrors.WithMessage(err, "failed to load config")
	}

	appCtx := &AppContext{
		Config: cnf,
	}

	return appCtx, nil
}
