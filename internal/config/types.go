package config

type Configuration struct {
	Name       string            `yaml:"name"`
	Properties map[string]string `yaml:"properties"`
}

type ConfigFile struct {
	Configurations []Configuration `yaml:"configurations"`
}
