package config

import "github.com/DaZZler12/sjf-be/pkg/entities/database/mongo"

type ServerConfig struct {
	Port string `yaml:"port"`
	Host string `yaml:"host"`
}

type Config struct {
	Database     mongo.DatabaseConfig `yaml:"database"`
	ServerConfig ServerConfig         `yaml:"serverconfig"`
}
