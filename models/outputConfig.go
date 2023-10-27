package models

import (
	"go.k6.io/k6/output"
)

var (
	ENV_SERVER_URL  = "XK6_TRACETEST_SERVER_URL"
	ENV_SERVER_PATH = "XK6_TRACETEST_SERVER_PATH"
	ENV_API_TOKEN   = "XK6_TRACETEST_API_TOKEN"
)

type OutputConfig struct {
	ServerUrl  string
	ServerPath string
	APIToken   string
}

func NewConfig(params output.Params) (OutputConfig, error) {
	cfg := OutputConfig{
		ServerUrl:  ServerURL,
		ServerPath: ServerPath,
		APIToken:   APIToken,
	}

	if params.ConfigArgument != "" {
		cfg.ServerUrl = params.ConfigArgument
	} else if val, ok := params.Environment[ENV_SERVER_URL]; ok {
		cfg.ServerUrl = val
	}

	if val, ok := params.Environment[ENV_SERVER_PATH]; ok {
		cfg.ServerPath = val
	}

	if val, ok := params.Environment[ENV_API_TOKEN]; ok {
		cfg.APIToken = val
	}

	return cfg, nil
}
