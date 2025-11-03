package config

import (
	"fmt"
	"os"
	"time"

	"github.com/kelseyhightower/envconfig"

	"github.com/stackus/dotenv"
)

type (
	PGConfig struct {
		Conn string `required:"true"`
	}

	SMTPConfig struct {
		Host     string `required:"true"`
		Port     int    `required:"true"`
		Username string `required:"true"`
		Password string `required:"true"`
		Sender   string `required:"true"`
	}

	JWTSecret struct {
		Secret string `required:"true"`
	}

	AppConfig struct {
		Environment     string
		LogLevel        string `envconfig:"LOG_LEVEL" default:"DEBUG"`
		PG              PGConfig
		Web             WebConfig
		Rpc             RpcConfig
		SMTP            SMTPConfig
		JWT             JWTSecret
		ShutdownTimeout time.Duration `envconfig:"SHUTDOWN_TIMEOUT" default:"30s"`
	}
)

type WebConfig struct {
	Host string `default:"0.0.0.0"`
	Port string `default:":8080"`
}

func (c WebConfig) Address() string {
	return fmt.Sprintf("%s%s", c.Host, c.Port)
}

type RpcConfig struct {
	Host string `default:"0.0.0.0"`
	Port string `default:":8085"`
}

func (c RpcConfig) Address() string {
	return fmt.Sprintf("%s%s", c.Host, c.Port)
}

func InitConfig() (cfg AppConfig, err error) {
	if err = dotenv.Load(dotenv.EnvironmentFiles(os.Getenv("ENVIRONMENT"))); err != nil {
		return
	}

	err = envconfig.Process("", &cfg)

	return
}
