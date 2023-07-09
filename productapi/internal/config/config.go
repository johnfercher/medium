package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
	"regexp"
	"strings"
)

const filePath = "configs/%s.yml"

type Config struct {
	Env   string `yaml:"env"`
	Mysql struct {
		Url      string `yaml:"url"`
		Db       string `yaml:"db"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	} `yaml:"mysql"`
}

func (c *Config) Print() {
	fmt.Printf("loaded env=%s, mysql.url=%s, mysql.db=%s, mysql.user=%s, mysql.password=%s\n", c.Env, c.Mysql.Url, c.Mysql.Db, c.Mysql.User, c.Mysql.Password)
}

func Load(args []string) (*Config, error) {
	env, err := GetEnv(args)
	if err != nil {
		return nil, err
	}

	fmt.Printf("loading config file from env=%s\n", env)

	f, err := os.Open(fmt.Sprintf(filePath, env))
	if err != nil {
		fmt.Printf("could not load config file from env=%s\n", env)
		return nil, err
	}
	defer f.Close()

	cfg := &Config{}
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		fmt.Printf("could not parse config file from env=%s\n", env)
		return nil, err
	}

	cfg.Env = env
	cfg.Print()

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
