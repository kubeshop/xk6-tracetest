package tracetest

import (
	"fmt"

	"github.com/kubeshop/xk6-tracetest/models"
	"github.com/kubeshop/xk6-tracetest/modules/instance"
	tracetestOutput "github.com/kubeshop/xk6-tracetest/modules/output"
	"github.com/kubeshop/xk6-tracetest/modules/tracetest"
	"go.k6.io/k6/js/modules"
	"go.k6.io/k6/output"
)

func init() {
	tracetest := tracetest.New()
	output.RegisterExtension("xk6-tracetest", func(params output.Params) (output.Output, error) {
		cfg, err := models.NewConfig(params)
		if err != nil {
			return nil, fmt.Errorf("could not get tracetest config from output params: %w", err)
		}

		tracetest.UpdateFromConfig(cfg)

		return tracetestOutput.New(params, tracetest)
	})

	modules.Register("k6/x/tracetest", New(tracetest))
}

type RootModule struct {
	tracetest *tracetest.Tracetest
}

var _ modules.Module = &RootModule{}

func New(tracetest *tracetest.Tracetest) *RootModule {
	return &RootModule{
		tracetest: tracetest,
	}
}

func (r *RootModule) NewModuleInstance(vu modules.VU) modules.Instance {
	r.tracetest.Vu = vu
	return instance.New(vu, r.tracetest)
}
