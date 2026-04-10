package config

import (
	"sync"
	"sync/atomic"

	"github.com/spf13/viper"
)

var (
	once sync.Once
	cfg  atomic.Pointer[EnvConfig]
)

type EnvConfig struct {
	AppEnv         string `mapstructure:"APP_ENV"`
	AppName        string `mapstructure:"APP_NAME"`
	RestPort       string `mapstructure:"REST_PORT"`
	AppVersion     string `mapstructure:"APP_VERSION"`
	GracefulPeriod int    `mapstructure:"GRACEFUL_PERIOD"`
	DatabaseURL    string `mapstructure:"DATABASE_URL"`
	DSN            string `mapstructure:"MYSQL_DSN"`
	SQLitePath     string `mapstructure:"SQLITE_PATH"`
}

func GetConfig() *EnvConfig {
	once.Do(func() {
		c := loadConfig()
		cfg.Store(c)
	})
	return cfg.Load()
}

func loadConfig() *EnvConfig {
	var envCfg EnvConfig
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			panic(err)
		}
		// Continue with environment variables only
	}
	if err := viper.Unmarshal(&envCfg); err != nil {
		panic(err)
	}
	return &envCfg
}
