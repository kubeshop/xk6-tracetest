package models

import (
	"go.k6.io/k6/output"
)

var (
	ENV_SERVER_URL  = "XK6_TRACETEST_SERVER_URL"
	ENV_SERVER_PATH = "XK6_TRACETEST_SERVER_PATH"
)

type OutputConfig struct {
	ServerUrl  string
	ServerPath string
}

func NewConfig(params output.Params) (OutputConfig, error) {
	cfg := OutputConfig{
		ServerUrl:  ServerURL,
		ServerPath: ServerPath,
	}

	if params.ConfigArgument != "" {
		cfg.ServerUrl = params.ConfigArgument
	} else if val, ok := params.Environment[ENV_SERVER_URL]; ok {
		cfg.ServerUrl = val
	}

	if val, ok := params.Environment[ENV_SERVER_PATH]; ok {
		cfg.ServerPath = val
	}

	return cfg, nil
}
