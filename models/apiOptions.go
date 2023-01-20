package models

import (
	"fmt"

	"github.com/dop251/goja"
	"github.com/kubeshop/xk6-tracetest/utils"
	"go.k6.io/k6/js/modules"
)

type ApiOptions struct {
	ServerUrl  string
	ServerPath string
}

const (
	DefaultServerUrl = "http://localhost:3000"
	ServerURL        = "serverUrl"
	ServerPath       = "serverPath"
)

func NewApiOptions(vu modules.VU, val goja.Value) (ApiOptions, error) {
	rawOptions := utils.ParseOptions(vu, val)
	options := ApiOptions{
		ServerUrl:  DefaultServerUrl,
		ServerPath: "",
	}

	if len(rawOptions) == 0 {
		return options, nil
	}

	for key, value := range rawOptions {
		switch key {
		case ServerURL:
			options.ServerUrl = value.ToString().String()
		case ServerPath:
			options.ServerPath = value.ToString().String()
		default:
			return options, fmt.Errorf("unknown Tracetest option '%s'", key)
		}
	}

	return options, nil
}
