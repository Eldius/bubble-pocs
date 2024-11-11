package config

import "github.com/spf13/viper"

const (
	DebugEnabled = "debug"
)

func GetDebug() bool {
	return viper.GetBool(DebugEnabled)
}
