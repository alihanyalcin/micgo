package configuration

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"project/internal/pkg/bootstrap/interfaces"

	"github.com/BurntSushi/toml"
)

// LoadFromFile attempts to read and unmarshal toml-based configuration into a configuration struct.
func LoadFromFile(configDir, profileDir, configFileName string, config interfaces.Configuration) error {
	if len(configDir) == 0 {
		configDir = os.Getenv("MICROSERVICE_CONF_DIR")
	}
	if len(configDir) == 0 {
		configDir = "./res"
	}

	if len(profileDir) == 0 {
		profileDir += "/"
	}

	fileName := configDir + "/" + profileDir + configFileName

	contents, err := ioutil.ReadFile(fileName)
	if err != nil {
		return errors.New(fmt.Sprintf("could not load configuration file (%s): %s", fileName, err.Error()))
	}
	if err = toml.Unmarshal(contents, config); err != nil {
		return errors.New(fmt.Sprintf("could not load configuration file (%s): %s", fileName, err.Error()))
	}
	return nil
}
