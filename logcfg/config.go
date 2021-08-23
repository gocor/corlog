package logcfg

import "os"

// Config for logger
type Config struct {
	IsDebug bool
	App     string
	Env     string
}

// Defaults to debug
var sharedConfig Config

// GetSharedConfig ...
func GetSharedConfig() Config {
	return sharedConfig
}

// InitConfig ...
func InitConfig(config Config) {
	sharedConfig = config
}

// InitFromEnv ...
func InitFromEnv() {
	env := os.Getenv("ENV")
	app := os.Getenv("APP")

	isDebug := true
	if env == "prod" {
		isDebug = false
	}

	InitConfig(Config{
		IsDebug: isDebug,
		App:     app,
		Env:     env,
	})
}
