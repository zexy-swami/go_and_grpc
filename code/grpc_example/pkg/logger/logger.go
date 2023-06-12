package logger

import "go.uber.org/zap"

var Logger *zap.Logger

func SetupLogger() error {
	if Logger != nil {
		return nil
	}

	zapCfg, err := getZapConfig()
	if err != nil {
		return err
	}

	localLogger, err := zapCfg.Build()
	if err != nil {
		return err
	}

	Logger = localLogger
	return nil
}
