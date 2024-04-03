package models

import (
	"github.com/dop251/goja"
)

type TracetestOptions struct {
	TestID       string
	ShouldWait   bool
	VariableName string
	RunGroupId   string
	Definition   string
}

func NewTracetestOptions(runTime *goja.Runtime, params *goja.Object) TracetestOptions {
	rawOptions := params.Get("tracetest")
	options := TracetestOptions{
		ShouldWait:   true,
		VariableName: "TRACE_ID",
	}

	if rawOptions == nil {
		return options
	}

	optionsObject := rawOptions.ToObject(runTime)
	for _, key := range optionsObject.Keys() {
		switch key {
		case "testId":
			options.TestID = optionsObject.Get(key).String()
		case "definition":
			options.Definition = optionsObject.Get(key).String()
		case "shouldWait":
			options.ShouldWait = optionsObject.Get(key).ToBoolean()
		case "variableName":
			options.VariableName = optionsObject.Get(key).ToString().String()
		}
	}

	return options
}
