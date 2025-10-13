// libs/config/loader.go
package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

func Load[T any](path string) (*T, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	expandedData := []byte(os.ExpandEnv(string(data)))

	var cfg T
	if err := yaml.Unmarshal(expandedData, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
