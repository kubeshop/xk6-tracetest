package models

import (
	"fmt"
	"strings"

	"github.com/dop251/goja"
	"github.com/kubeshop/xk6-tracetest/utils"
	"go.k6.io/k6/js/modules"
)

var defaultPropagatorList = []PropagatorName{
	TraceContext,
	Baggage,
	PropagatorB3,
	OT,
	Jaeger,
	XRay,
}

type HttpClientOptions struct {
	Propagator Propagator
	Tracetest  TracetestOptions
}

func NewHttpClientOptions(vu modules.VU, val goja.Value) (HttpClientOptions, error) {
	rawOptions := utils.ParseOptions(vu, val)
	options := HttpClientOptions{
		Propagator: NewPropagator(defaultPropagatorList),
		Tracetest: TracetestOptions{
			ShouldWait:   true,
			VariableName: "TRACE_ID",
		},
	}

	if len(rawOptions) == 0 {
		return options, nil
	}

	for key, value := range rawOptions {
		switch key {
		case "propagator":
			rawPropagatorList := strings.Split(value.ToString().String(), ",")
			propagatorList := make([]PropagatorName, len(rawPropagatorList))
			for i, propagator := range rawPropagatorList {
				propagatorList[i] = PropagatorName(propagator)
			}

			options.Propagator = NewPropagator(propagatorList)
		case "tracetest":
			options.Tracetest = NewTracetestOptions(vu.Runtime(), val.ToObject(vu.Runtime()))
		default:
			return options, fmt.Errorf("unknown HTTP tracing option '%s'", key)
		}
	}

	return options, nil
}
