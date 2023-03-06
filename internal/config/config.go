package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Configs struct {
	Logging
	Server
	DB
}

type Logging struct {
	OutputFile string `mapstructure:"output-file"`
}

type Server struct {
	Port         string `mapstructure:"port"`
	ReadTimeout  int    `mapstructure:"read-timeout"`
	WriteTimeout int    `mapstructure:"write-timeout"`
}

type DB struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
	SslMode  string `mapstructure:"ssl-mode"`
}

func Init(configFile string) (*Configs, error) {
	viper.SetConfigName(configFile)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")

	if err := parseEnv(); err != nil {
		log.Fatalf("Error occured when parsing env file: %s", err.Error())
		return nil, err
	}

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error occured when reading config file: %s", err.Error())
		return nil, err
	}

	configs := &Configs{}
	if err := viper.Unmarshal(configs); err != nil {
		log.Fatalf("Error occured when decoding config file struct: %s", err.Error())
		return nil, err
	}
	return configs, nil
}

func parseEnv() error {
	configEnvMap := make(map[string]string)

	configEnvMap["project"] = "PROJECT_NAME"
	configEnvMap["server.port"] = "SERVER_PORT"
	configEnvMap["db.host"] = "DB_HOST"
	configEnvMap["db.port"] = "DB_PORT"
	configEnvMap["db.username"] = "DB_USERNAME"
	configEnvMap["db.password"] = "DB_PASSWORD"
	configEnvMap["db.database"] = "DB_DATABASE"

	for configKey, envKey := range configEnvMap {
		if err := viper.BindEnv(configKey, envKey); err != nil {
			return err
		}
	}
	return nil

}
