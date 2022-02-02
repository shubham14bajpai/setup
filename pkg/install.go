package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

type Candidate struct {
	Name     string   `yaml:"name"`
	Commands []string `yaml:"commands"`
}

type Config struct {
	Setup []Candidate `yaml:"setup"`
}

func Install() {

	var config Config
	filename, err := filepath.Abs("./config/ubuntu.yml")
	if err != nil {
		fmt.Println("failed to find config yaml:", err.Error())
		return
	}
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("failed to read config yaml:", err.Error())
		return
	}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		fmt.Println("failed to unmarshall config yaml:", err.Error())
		return
	}
	s, _ := json.MarshalIndent(config, "", "\t")
	fmt.Println("Configuring the system for the below config:")
	fmt.Printf("%s\n\n", string(s))

	for _, c := range config.Setup {
		install(c)
	}

}

func install(c Candidate) {
	fmt.Printf("Installing %s...\n", c.Name)

	for _, command := range c.Commands {
		fmt.Println("Running:", command)
		command := strings.Split(command, " ")
		if len(command) == 2 && command[0] == "wget" {
			url := strings.Split(command[1], "/")
			name := url[len(url)-1]
			if _, err := os.Stat(name); err == nil {
				continue
			}
		}
		cmd := exec.Command(command[0], command[1:]...)
		_, err := cmd.Output()
		if err != nil {
			fmt.Println("failed to install", c.Name, ":", err.Error())
			return
		}
	}
	fmt.Printf("Installed %s successfully!!\n", c.Name)
}
