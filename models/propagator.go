package models

import (
	"context"
	"net/http"

	"github.com/kubeshop/tracetest/extensions/k6/utils"
	"go.opentelemetry.io/contrib/propagators/aws/xray"
	"go.opentelemetry.io/contrib/propagators/b3"
	"go.opentelemetry.io/contrib/propagators/jaeger"
	"go.opentelemetry.io/contrib/propagators/ot"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

type PropagatorName string

type Propagator struct {
	propagators []PropagatorName
}

const (
	TraceContext PropagatorName = "tracecontext"
	Baggage      PropagatorName = "baggage"
	PropagatorB3 PropagatorName = "b3"
	OT           PropagatorName = "ot"
	Jaeger       PropagatorName = "jaeger"
	XRay         PropagatorName = "xray"
)

func NewPropagator(propagators []PropagatorName) Propagator {
	return Propagator{
		propagators: propagators,
	}
}

func (p Propagator) GenerateHeaders(traceID string) http.Header {
	ctx := context.Background()
	spanContext := NewSpanContext(traceID)
	ctx = trace.ContextWithSpanContext(ctx, spanContext)
	header := http.Header{}

	carrier := propagation.MapCarrier{}
	otel.GetTextMapPropagator().Inject(ctx, carrier)
	propagator := p.getTextMapPropagator()
	propagator.Inject(ctx, propagation.HeaderCarrier(header))

	return header
}

func (p Propagator) getTextMapPropagator() propagation.TextMapPropagator {
	propagators := make([]propagation.TextMapPropagator, len(p.propagators))
	for i, propagator := range p.propagators {
		switch propagator {
		case Jaeger:
			propagators[i] = jaeger.Jaeger{}
		case OT:
			propagators[i] = ot.OT{}
		case PropagatorB3:
			propagators[i] = b3.New()
		case XRay:
			propagators[i] = xray.Propagator{}
		case TraceContext:
			propagators[i] = propagation.TraceContext{}
		case Baggage:
			propagators[i] = propagation.Baggage{}
		}
	}

	return propagation.NewCompositeTextMapPropagator(propagators...)
}

func NewSpanContext(traceID string) trace.SpanContext {
	parsedTraceID, _ := trace.TraceIDFromHex(traceID)
	var tf trace.TraceFlags
	return trace.NewSpanContext(trace.SpanContextConfig{
		TraceID:    parsedTraceID,
		SpanID:     utils.SpanID(),
		TraceFlags: tf.WithSampled(true),
		Remote:     true,
	})
}
