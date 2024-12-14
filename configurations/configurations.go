package configurations

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

type DbPsxConfig struct {
	User         string `yaml:"user"`
	Password     string `yaml:"password"`
	Dbname       string `yaml:"dbname"`
	Host         string `yaml:"host"`
	Port         int    `yaml:"port"`
	Sslmode      string `yaml:"sslmode"`
	MaxOpenConns int    `yaml:"max_open_conns"`
	Timer        int    `yaml:"timer"`
}

type DbRedisCfg struct {
	Host     string `yaml:"host"`
	Password string `yaml:"password"`
	DbNumber int    `yaml:"db"`
	Timer    int    `yaml:"timer"`
}

type DbService struct {
	Address          string `yaml:"address"`
	GraphqlPort      string `yaml:"graphql_port"`
	GrpcListenerPort string `yaml:"grpc_listener_port"`
	ConnectionType   string `yaml:"connection_type"`
}

func ParseRedisConfig(path string) (*DbRedisCfg, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("os.ReadFile: %v", err)
	}

	var config DbRedisCfg
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		return nil, fmt.Errorf("yaml.Unmarshal: %v", err)
	}

	return &config, nil
}

func ParsePostgresConfig(path string) (*DbPsxConfig, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("os.ReadFile: %v", err)
	}

	var config DbPsxConfig
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		return nil, fmt.Errorf("yaml.Unmarshal: %v", err)
	}

	return &config, nil
}

func ParseServiceConfig(path string) (*DbService, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("os.ReadFile: %v", err)
	}

	var config DbService
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		return nil, fmt.Errorf("yaml.Unmarshal: %v", err)
	}

	return &config, nil
}
