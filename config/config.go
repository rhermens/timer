package config

import (
	"path"

	"github.com/adrg/xdg"
)

type Config struct {
	Directory string
	Format    string
}

func DefaultConfig() Config {
	return Config{
		Directory: path.Join(xdg.DataHome, "timer"),
		Format:    "%s,%s\n",
	}
}
