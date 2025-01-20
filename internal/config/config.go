package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type PVEConfig struct {
	Host     string `yaml:"host"`     // PVE主机地址
	Username string `yaml:"username"` // PVE用户名
	Password string `yaml:"password"` // PVE密码
	Realm    string `yaml:"realm"`    // 认证域（通常是pam或pve）
}

type Config struct {
	PVE PVEConfig `yaml:"pve"`
}

func LoadConfig() (*Config, error) {
	data, err := os.ReadFile("configs/config.yaml")
	if err != nil {
		return nil, fmt.Errorf("error reading config.yaml: %v", err)
	}
	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("error parsing config.yaml: %v", err)
	}
	return &cfg, nil
}
