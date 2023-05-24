// Copyright (C) 2022 Storj Labs, Inc.
// See LICENSE for copying information.

package http

import (
	"net/http"

	"github.com/spacemonkeygo/monkit/v3"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

// TraceHandler wraps a HTTPHandler and import trace information from header.
func TraceHandler(c http.Handler, scope *monkit.Scope, allowedBaggage ...string) http.Handler {
	return traceHandler{
		otelhttp.NewHandler(c, ""),
	}
}

type traceHandler struct {
	http.Handler
}

// ServeHTTP implements http.Handler with span propagation.
func (t traceHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	t.Handler.ServeHTTP(writer, request)
}
