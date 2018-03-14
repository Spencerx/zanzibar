// Code generated by zanzibar
// @generated

// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package bazendpoint

import (
	"context"

	zanzibar "github.com/uber/zanzibar/runtime"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	clientsBazBase "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients/baz/base"
	clientsBazBaz "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients/baz/baz"
	endpointsBazBaz "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/endpoints/baz/baz"

	module "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/baz/module"
)

// SimpleServiceSillyNoopHandler is the handler for "/baz/silly-noop"
type SimpleServiceSillyNoopHandler struct {
	Clients  *module.ClientDependencies
	endpoint *zanzibar.RouterEndpoint
}

// NewSimpleServiceSillyNoopHandler creates a handler
func NewSimpleServiceSillyNoopHandler(deps *module.Dependencies) *SimpleServiceSillyNoopHandler {
	handler := &SimpleServiceSillyNoopHandler{
		Clients: deps.Client,
	}
	handler.endpoint = zanzibar.NewRouterEndpoint(
		deps.Default.Logger, deps.Default.Scope,
		"baz", "sillyNoop",
		handler.HandleRequest,
	)
	return handler
}

// Register adds the http handler to the gateway's http router
func (h *SimpleServiceSillyNoopHandler) Register(g *zanzibar.Gateway) error {
	g.HTTPRouter.Register(
		"GET", "/baz/silly-noop",
		h.endpoint,
	)
	// TODO: register should return errors on route conflicts
	return nil
}

// HandleRequest handles "/baz/silly-noop".
func (h *SimpleServiceSillyNoopHandler) HandleRequest(
	ctx context.Context,
	req *zanzibar.ServerHTTPRequest,
	res *zanzibar.ServerHTTPResponse,
) {

	// log endpoint request to downstream services
	zfields := []zapcore.Field{
		zap.String("endpoint", h.endpoint.EndpointName),
	}

	req.Logger.Debug("Endpoint request to downstream", zfields...)

	workflow := SimpleServiceSillyNoopEndpoint{
		Clients: h.Clients,
		Logger:  req.Logger,
		Request: req,
	}

	cliRespHeaders, err := workflow.Handle(ctx, req.Header)
	if err != nil {
		switch errValue := err.(type) {

		case *endpointsBazBaz.AuthErr:
			res.WriteJSON(
				403, cliRespHeaders, errValue,
			)
			return

		case *endpointsBazBaz.ServerErr:
			res.WriteJSON(
				500, cliRespHeaders, errValue,
			)
			return

		default:
			res.SendError(500, "Unexpected server error", err)
			return
		}

	}

	res.WriteJSONBytes(204, cliRespHeaders, nil)
}

// SimpleServiceSillyNoopEndpoint calls thrift client Baz.DeliberateDiffNoop
type SimpleServiceSillyNoopEndpoint struct {
	Clients *module.ClientDependencies
	Logger  *zap.Logger
	Request *zanzibar.ServerHTTPRequest
}

// Handle calls thrift client.
func (w SimpleServiceSillyNoopEndpoint) Handle(
	ctx context.Context,
	reqHeaders zanzibar.Header,
) (zanzibar.Header, error) {

	clientHeaders := map[string]string{}

	_, err := w.Clients.Baz.DeliberateDiffNoop(ctx, clientHeaders)

	if err != nil {
		switch errValue := err.(type) {

		case *clientsBazBaz.AuthErr:
			serverErr := convertSillyNoopAuthErr(
				errValue,
			)
			// TODO(sindelar): Consider returning partial headers

			return nil, serverErr

		case *clientsBazBase.ServerErr:
			serverErr := convertSillyNoopServerErr(
				errValue,
			)
			// TODO(sindelar): Consider returning partial headers

			return nil, serverErr

		default:
			w.Logger.Warn("Could not make client request",
				zap.Error(errValue),
				zap.String("client", "Baz"),
			)

			// TODO(sindelar): Consider returning partial headers

			return nil, err

		}
	}

	// Filter and map response headers from client to server response.

	// TODO: Add support for TChannel Headers with a switch here
	resHeaders := zanzibar.ServerHTTPHeader{}

	return resHeaders, nil
}

func convertSillyNoopAuthErr(
	clientError *clientsBazBaz.AuthErr,
) *endpointsBazBaz.AuthErr {
	// TODO: Add error fields mapping here.
	serverError := &endpointsBazBaz.AuthErr{}
	return serverError
}
func convertSillyNoopServerErr(
	clientError *clientsBazBase.ServerErr,
) *endpointsBazBaz.ServerErr {
	// TODO: Add error fields mapping here.
	serverError := &endpointsBazBaz.ServerErr{}
	return serverError
}
