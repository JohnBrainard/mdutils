package config

import (
	"io"

	"github.com/naoina/toml"
)

type Config struct {
	Template Template
	Colors   map[string]HexColor
	Layers   []Layer
	Edges    Edges
}

func Read(rdr io.Reader) (Config, error) {
	var config Config
	if err := toml.NewDecoder(rdr).Decode(&config); err != nil {
		return Config{}, err
	}
	return config, nil
}

func (c *Config) Validate() {
	for _, layer := range c.Layers {
		layer.Validate()
	}
}
