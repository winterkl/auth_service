package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"time"
)

const defaultConfigPath = "./config.yaml"

type Config struct {
	Env        string `yaml:"env" env-required:"local"`
	GRPCServer `yaml:"grpc_server"`
	DataBase   `yaml:"db"`
	Redis      `yaml:"redis"`
	JwtAuth    `yaml:"jwt_auth"`
}

type GRPCServer struct {
	PORT    int           `yaml:"port" env-required:"true"`
	Host    string        `yaml:"host" env-required:"localhost"`
	Timeout time.Duration `yaml:"timeout" env-required:"10s"`
}

type DataBase struct {
	Host     string `yaml:"host" env-required:"true"`
	Port     int    `yaml:"port" env-required:"true"`
	User     string `yaml:"user" env-required:"true"`
	Password string `yaml:"password" env-required:"true"`
	DB       string `yaml:"db" env-required:"true"`
	SSLMode  string `yaml:"ssl_mode" env-required:"true"`
}

type Redis struct {
	Host     string `yaml:"host" env-required:"true"`
	Port     int    `yaml:"port" env-required:"true"`
	DB       int    `yaml:"db"`
	Password string `yaml:"password" env-required:"true"`
}

type JwtAuth struct {
	Key string `yaml:"key" env-required:"true"`
}

func MustLoad() *Config {
	configPath := fetchConfigPath()
	return MustLoadPath(configPath)
}

func MustLoadPath(configPath string) *Config {
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("config file does not found: " + configPath)
	}

	var cfg = new(Config)

	if err := cleanenv.ReadConfig(configPath, cfg); err != nil {
		panic("cannot read config: " + err.Error())
	}

	return cfg
}

func fetchConfigPath() string {

	var res string

	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	if res == "" {
		if envPath := os.Getenv("CONFIG_PATH"); envPath != "" {
			return envPath
		}
	}

	return defaultConfigPath
}
