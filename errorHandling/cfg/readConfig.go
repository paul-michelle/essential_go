package cfg

import (
	"os"
	"github.com/pkg/errors"
)

type Configuration struct {
	Host string
	Port int32
}

func ReadConfig(pathToConfigFile string) (*Configuration, error) {
	configFileOpened, err := os.Open(pathToConfigFile)
	if err != nil {
		return nil, errors.Wrap(err, "cannot open configuration file")
	}

	defer configFileOpened.Close()

	configs := &Configuration{}

	// cofig info parsed in here

	return configs, nil
}
