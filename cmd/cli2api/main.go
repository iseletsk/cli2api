package main

import (
	"flag"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

func failOnError(err error) {
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func initConfig(path string) Config {
	dat, err := ioutil.ReadFile(path)
	failOnError(err)
	config := Config{}
	err = yaml.Unmarshal(dat, &config)
	failOnError(err)
	for _, cli := range config.CLI {
		// Set default output content type if not defined
		if cli.OutputContentType == "" {
			cli.OutputContentType = "text/plain"
		}
	}
	return config
}

func main() {
	configFile := flag.String("config", "cli2api.yaml", "Path to config file")
	flag.Parse()
	config := initConfig(*configFile)
	StartHTTP(config)
}
