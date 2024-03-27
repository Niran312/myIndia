package main

import (
	"myIndia/app"
)

func main() {
	//loggerConfig := zap.NewProductionConfig()
	//loggerConfig.EncoderConfig.TimeKey = "timestamp"
	//loggerConfig.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.DateTime)
	//logger, err := loggerConfig.Build()
	//if err != nil {
	//	panic("failed to initialize logger: " + err.Error())
	//}
	//defer func(logger *zap.Logger) {
	//	err := logger.Sync()
	//	if err != nil {
	//		panic("There is some error in logger: " + err.Error())
	//	}
	//}(logger)
	//sugar := logger.Sugar()
	//sugar.Infof("Initialized App Logger using zap")
	app.Run()
}
