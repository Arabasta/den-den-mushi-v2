package config

import (
	"den-den-mushi-v2/pkg/config"
)

type Config struct {
	App *config.App

	Ssl *config.Ssl

	Cors *config.Cors

	Logger *config.Logger

	Ssh *config.Ssh

	Websocket struct {
		ReadBufferSize  int      `json:"ReadBufferSize"`
		WriteBufferSize int      `json:"WriteBufferSize"`
		AllowedOrigins  []string `json:"AllowedOrigins"`
		Subprotocols    string   `json:"Subprotocols"`
	} `json:"Websocket"`
}
