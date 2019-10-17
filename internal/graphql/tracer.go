package graphql

import (
	"context"
	"sync"

	"github.com/99designs/gqlgen/graphql"
)

const (
	tracingExt = "tracing"

	tracingOperationParsing    = "parsing"
	tracingOperationValidation = "validation"

	tracingStart    = "startTime"
	tracingEnd      = "endTime"
	tracingDuration = "duration"
)

// todo implement tracer

func (s *server) newTracer() graphql.Tracer {
	return &tracer{
		timer: make(map[string]interface{}),
	}
}

type tracer struct {
	m     sync.Mutex
	timer map[string]interface{}
}

func (t *tracer) StartOperationParsing(ctx context.Context) context.Context {
	return ctx
}

func (t *tracer) EndOperationParsing(ctx context.Context) {
}

func (t *tracer) StartOperationValidation(ctx context.Context) context.Context {
	return ctx
}

func (t *tracer) EndOperationValidation(ctx context.Context) {
}

func (t *tracer) StartOperationExecution(ctx context.Context) context.Context {
	return ctx
}

func (t *tracer) StartFieldExecution(ctx context.Context, field graphql.CollectedField) context.Context {
	return ctx
}

func (t *tracer) StartFieldResolverExecution(ctx context.Context, rc *graphql.ResolverContext) context.Context {
	return ctx
}

func (t *tracer) StartFieldChildExecution(ctx context.Context) context.Context {
	return ctx
}

func (t *tracer) EndFieldExecution(ctx context.Context) {
}

func (t *tracer) EndOperationExecution(ctx context.Context) {
}
