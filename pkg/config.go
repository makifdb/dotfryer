package cli

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Name     string `yaml:"name"`
	Username string `yaml:"username"`
	Pc       string `yaml:"pc"`
	Repo     string `yaml:"repo"`
	Location string `yaml:"location"`
}

var C Config = Config{"", "", "", "", ".dotfiles/"}

func ReadConfigFile() Config {
	lc := C.Location + "config.yaml"
	_, err := os.OpenFile(lc, os.O_RDONLY, 0644)
	if os.IsNotExist(err) {
		fmt.Println("Config file not found. Creating...")
		WriteConfigFile(lc)
	}

	f, err := ioutil.ReadFile(lc)
	if err != nil {
		fmt.Printf("Error reading YAML file: %s\n", err)
	}

	err = yaml.Unmarshal(f, &C)
	if err != nil {
		fmt.Printf("Error parsing YAML file: %s\n", err)
	}
	return C
}

func WriteConfigFile(configLocation string) {
	data, err := yaml.Marshal(&C)
	if err != nil {
		fmt.Printf("Error unparsing YAML file: %s\n", err)
	}

	err2 := ioutil.WriteFile(configLocation, data, 0)
	if err2 != nil {
		fmt.Printf("Error writing YAML file: %s\n", err)
	}
}
