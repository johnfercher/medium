package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
	"regexp"
	"strings"
)

type Config struct {
	Env   string `yaml:"env"`
	Mysql struct {
		Url      string `yaml:"url"`
		Db       string `yaml:"db"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	} `yaml:"mysql"`
}

func Load(args []string) (*Config, error) {
	env, err := GetEnv(args)
	if err != nil {
		return nil, err
	}

	f, err := os.Open(fmt.Sprintf("configs/%s.yml", env))
	if err != nil {
		return nil, err
	}
	defer f.Close()

	cfg := &Config{}
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		return nil, err
	}

	cfg.Env = env

	return cfg, nil
}

func GetEnv(args []string) (string, error) {
	envRegex, err := regexp.Compile(`env=\w+`)
	if err != nil {
		return "", err
	}

	for _, arg := range args {
		env := envRegex.FindString(arg)
		if env != "" {
			return strings.Replace(env, "env=", "", -1), nil
		}
	}

	return "local", nil
}
