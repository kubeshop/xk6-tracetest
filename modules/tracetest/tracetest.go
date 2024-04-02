package tracetest

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/dop251/goja"
	"github.com/kubeshop/tracetest/cli/openapi"
	"github.com/kubeshop/tracetest/server/pkg/id"
	"github.com/kubeshop/xk6-tracetest/models"
	"github.com/sirupsen/logrus"
	"go.k6.io/k6/js/common"
	"go.k6.io/k6/js/modules"
	"go.k6.io/k6/output"
)

type Tracetest struct {
	Vu              modules.VU
	bufferLock      sync.Mutex
	buffer          []models.Job
	runGroupId      string
	processedBuffer sync.Map
	periodicFlusher *output.PeriodicFlusher
	logger          logrus.FieldLogger
	client          *openapi.APIClient
	apiOptions      models.ApiOptions
	mutex           sync.Mutex
	jwt             string
}

func New() *Tracetest {
	logger := *logrus.New()
	client, jwt := NewAPIClient(models.ApiOptions{})
	tracetest := &Tracetest{
		buffer:          []models.Job{},
		processedBuffer: sync.Map{},
		logger:          logger.WithField("component", "xk6-tracetest-tracing"),
		client:          client,
		mutex:           sync.Mutex{},
		jwt:             jwt,
		runGroupId:      id.GenerateID().String(),
	}

	duration := 1 * time.Second
	periodicFlusher, _ := output.NewPeriodicFlusher(duration, tracetest.processQueue)
	tracetest.periodicFlusher = periodicFlusher

	return tracetest
}

func (t *Tracetest) UpdateFromConfig(config models.OutputConfig) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	if config.ServerUrl == "" {
		config.ServerUrl = models.DefaultServerUrl
	}

	apiOptions := models.ApiOptions{
		ServerUrl:  config.ServerUrl,
		ServerPath: config.ServerPath,
		APIToken:   config.APIToken,
	}

	t.apiOptions = apiOptions
	t.client, t.jwt = NewAPIClient(apiOptions)
}

func (t *Tracetest) Constructor(call goja.ConstructorCall) *goja.Object {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	rt := t.Vu.Runtime()

	return rt.ToValue(t).ToObject(rt)
}

func (t *Tracetest) RunTest(traceID string, options models.TracetestOptions, request models.Request) {
	if options.RunGroupId == "" {
		options.RunGroupId = t.runGroupId
	}

	t.queueJob(models.NewJob(traceID, options, request))
}

func (t *Tracetest) Summary() string {
	runGroup, _ := t.wait()

	return t.stringSummary(runGroup)
}

func (t *Tracetest) ValidateResult() {
	runGroup, _ := t.wait()

	summary := runGroup.GetSummary()

	if summary.GetFailed() > 0 {
		panic(fmt.Sprintf("Tracetest: %d jobs failed", summary.GetFailed()))
	}
}

func (t *Tracetest) Json() string {
	runGroup, _ := t.wait()

	rt := t.Vu.Runtime()
	jsonString, err := json.Marshal(t.jsonSummary(runGroup))

	if err != nil {
		common.Throw(rt, err)
	}

	return string(jsonString)
}

func (t *Tracetest) wait() (openapi.RunGroup, error) {
	if len(t.buffer) != 0 {
		t.processQueue()
	}

	return t.waitForRunGroup(t.runGroupId)
}
