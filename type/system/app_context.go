package system

import (
	"github.com/creasty/apperrors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// AppContext represents a context of application
type AppContext struct {
	*Config
	Logger *zap.Logger
}

// CreateAppContext loads application configs and creates objects to use in the application lifecycle
func CreateAppContext() (*AppContext, error) {
	var err error
	appCtx := &AppContext{}

	appCtx.Config, err = loadConfig()
	if err != nil {
		return nil, apperrors.WithMessage(err, "failed to load config")
	}

	if appCtx.IsDevelopment() {
		appCtx.Logger, err = zap.NewProduction()
	} else {
		appCtx.Logger, err = createZapConfigDev().Build()
	}
	if err != nil {
		return nil, apperrors.WithMessage(err, "failed to initialize logger")
	}

	return appCtx, nil
}

// Close cleans-up all connections
func (c *AppContext) Close() error {
	if c.Logger != nil {
		return c.Logger.Sync()
	}
	return nil
}

// IsDevelopment returns whether the application is running as a development mode
func (c *AppContext) IsDevelopment() bool {
	return c.Config.Env == "development"
}

// IsProduction returns whether the application is running as a production mode
func (c *AppContext) IsProduction() bool {
	return c.Config.Env == "production"
}

func createZapConfigDev() *zap.Config {
	lv := zap.NewAtomicLevel()
	lv.SetLevel(zapcore.DebugLevel)
	return &zap.Config{
		Level:       lv,
		Development: true,
		Encoding:    "console",
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:     "message",
			LevelKey:       "level",
			TimeKey:        "time",
			NameKey:        "name",
			CallerKey:      "caller",
			StacktraceKey:  "stack",
			EncodeLevel:    zapcore.CapitalColorLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}
}
