package tracetest

import (
	"encoding/json"
	"sync"
	"time"

	"github.com/dop251/goja"
	"github.com/kubeshop/xk6-tracetest/models"
	"github.com/kubeshop/xk6-tracetest/openapi"
	"github.com/sirupsen/logrus"
	"go.k6.io/k6/js/common"
	"go.k6.io/k6/js/modules"
	"go.k6.io/k6/output"
)

type Tracetest struct {
	Vu              modules.VU
	bufferLock      sync.Mutex
	buffer          []models.Job
	processedBuffer sync.Map
	periodicFlusher *output.PeriodicFlusher
	logger          logrus.FieldLogger
	client          *openapi.APIClient
	apiOptions      models.ApiOptions
}

func New() *Tracetest {
	logger := *logrus.New()
	tracetest := &Tracetest{
		buffer:          []models.Job{},
		processedBuffer: sync.Map{},
		logger:          logger.WithField("component", "xk6-tracetest-tracing"),
		client:          NewAPIClient(models.ApiOptions{}),
	}

	duration := 1 * time.Second
	periodicFlusher, _ := output.NewPeriodicFlusher(duration, tracetest.processQueue)
	tracetest.periodicFlusher = periodicFlusher

	return tracetest
}

func (t *Tracetest) Constructor(call goja.ConstructorCall) *goja.Object {
	rt := t.Vu.Runtime()
	apiOptions, err := models.NewApiOptions(t.Vu, call.Argument(0))
	if err != nil {
		common.Throw(rt, err)
	}

	t.apiOptions = apiOptions
	t.client = NewAPIClient(apiOptions)

	return rt.ToValue(t).ToObject(rt)
}

func (t *Tracetest) RunTest(traceID string, options models.TracetestOptions, request models.Request) {
	t.queueJob(models.NewJob(traceID, options, request))
}

func (t *Tracetest) Summary() string {
	if len(t.buffer) != 0 {
		t.processQueue()
	}

	return t.stringSummary()
}

func (t *Tracetest) Json() string {
	rt := t.Vu.Runtime()
	jsonString, err := json.Marshal(t.jsonSummary())

	if err != nil {
		common.Throw(rt, err)
	}

	return string(jsonString)
}
