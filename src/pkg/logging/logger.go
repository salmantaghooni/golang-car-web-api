package logging

import "github.com/salmantaghooni/golang-car-web-api/src/config"

type Logger interface {
	Init()

	Info(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{})
	Infof(template string, args ...interface{})

	Debug(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{})
	Debugf(template string, args ...interface{})

	Warn(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{})
	Warnf(template string, args ...interface{})

	Error(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{})
	Errorf(template string, args ...interface{})

	Fatal(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{})
	Fatalf(template string, args ...interface{})
}

func NewLogger(cfg *config.Config) Logger {
	return newZapLogger(cfg)
}
