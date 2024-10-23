package config

import (
	"github.com/spf13/viper"
	"strings"
)

func GetDebugEnabled() bool {
	return viper.GetBool(DebugEnabled)
}

func GetLogFormat() string {
	return strings.ToLower(viper.GetString(LogFormatKey))
}

func GetLogLevel() string {
	return strings.ToLower(viper.GetString(LogLevelKey))
}

func GetLogOutput() string {
	return viper.GetString(LogOutputFileKey)
}
