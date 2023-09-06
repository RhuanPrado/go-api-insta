package logger

import "go.uber.org/zap"

var (
	Development *zap.Logger
	Production  *zap.Logger
)

func InitializeLogger() {
	var err interface{}

	Development, err = zap.NewDevelopment()

	if err != nil {
		panic(0)
	}

	Production, err = zap.NewProduction()

	if err != nil {
		panic(0)
	}
}
