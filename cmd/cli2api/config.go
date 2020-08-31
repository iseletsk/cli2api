package main

type CLIConfig struct {
	ID                string `yaml:"ID"`
	Command           string `yaml:"command"`
	APIKey            string `yaml:"api-key"`
	OutputContentType string `yaml:"output-content-type,omitempty"`
}

type Config struct {
	Server struct {
		Address  string `yaml:"address"`
		CertFile string `yaml:"cert-file"`
		KeyFile  string `yaml:"key-file"`
	} `yaml:"server"`
	CLI []CLIConfig `yaml:"cli"`
}
