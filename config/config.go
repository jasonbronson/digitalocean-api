package config

import "github.com/digitalocean/godo"

type Config struct {
	GodoClient *godo.Client
}
